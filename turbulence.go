package destiny

import "fmt"

type CPI struct {
	JobName     string
	ReleaseName string
}

func NewTurbulence(config Config) Manifest {
	turbulenceRelease := Release{
		Name:    "turbulence",
		Version: "latest",
	}

	ipRange := IPRange(config.IPRange)
	iaasConfig := IAASConfig(config)

	cloudProperties := iaasConfig.NetworkSubnet()
	cpi := iaasConfig.CPI()

	cpiRelease := Release{
		Name:    cpi.ReleaseName,
		Version: "latest",
	}

	turbulenceNetwork := Network{
		Name: "turbulence",
		Subnets: []NetworkSubnet{{
			CloudProperties: cloudProperties,
			Gateway:         ipRange.IP(1),
			Range:           string(ipRange),
			Reserved:        []string{ipRange.Range(2, 11), ipRange.Range(17, 254)},
			Static: []string{
				ipRange.IP(12),
				ipRange.IP(13),
			},
		}},
		Type: "manual",
	}

	compilation := Compilation{
		Network:             turbulenceNetwork.Name,
		ReuseCompilationVMs: true,
		Workers:             3,
		CloudProperties:     iaasConfig.Compilation(),
	}

	turbulenceResourcePool := ResourcePool{
		Name:    "turbulence",
		Network: turbulenceNetwork.Name,
		Stemcell: ResourcePoolStemcell{
			Name:    StemcellForIAAS(config.IAAS),
			Version: "latest",
		},
		CloudProperties: iaasConfig.ResourcePool(),
	}

	update := Update{
		Canaries:        1,
		CanaryWatchTime: "1000-180000",
		MaxInFlight:     1,
		Serial:          true,
		UpdateWatchTime: "1000-180000",
	}

	apiJob := Job{
		Instances: 1,
		Name:      "api",
		Networks: []JobNetwork{{
			Name:      turbulenceNetwork.Name,
			StaticIPs: turbulenceNetwork.StaticIPs(1),
		}},
		PersistentDisk: 1024,
		ResourcePool:   turbulenceResourcePool.Name,
		Templates: []JobTemplate{
			{
				Name:    "turbulence_api",
				Release: turbulenceRelease.Name,
			},
			{
				Name:    cpi.JobName,
				Release: cpiRelease.Name,
			},
		},
	}

	directorCACert := TurbulenceAPIDirectorCACert
	if config.BOSH.DirectorCACert != "" {
		directorCACert = config.BOSH.DirectorCACert
	}

	properties := Properties{
		TurbulenceAPI: &PropertiesTurbulenceAPI{
			Certificate: TurbulenceAPICertificate,
			CPIJobName:  cpi.JobName,
			Director: PropertiesTurbulenceAPIDirector{
				CACert:   directorCACert,
				Host:     config.BOSH.Target,
				Password: config.BOSH.Password,
				Username: config.BOSH.Username,
			},
			Password:   "turbulence-password",
			PrivateKey: TurbulenceAPIPrivateKey,
		},
	}

	//properties := properties.Merge(iaasConfig.Properties())

	switch config.IAAS {
	case Warden:
		properties.WardenCPI = &PropertiesWardenCPI{
			Agent: PropertiesWardenCPIAgent{
				Blobstore: PropertiesWardenCPIAgentBlobstore{
					Options: PropertiesWardenCPIAgentBlobstoreOptions{
						Endpoint: "http://10.254.50.4:25251",
						Password: "agent-password",
						User:     "agent",
					},
					Provider: "dav",
				},
				Mbus: "nats://nats:nats-password@10.254.50.4:4222",
			},
			Warden: PropertiesWardenCPIWarden{
				ConnectAddress: "10.254.50.4:7777",
				ConnectNetwork: "tcp",
			},
		}
	case AWS:
		properties.AWS = &PropertiesAWS{
			AccessKeyID:           config.AWS.AccessKeyID,
			SecretAccessKey:       config.AWS.SecretAccessKey,
			DefaultKeyName:        config.AWS.DefaultKeyName,
			DefaultSecurityGroups: config.AWS.DefaultSecurityGroups,
			Region:                config.AWS.Region,
		}
		properties.Registry = &PropertiesRegistry{
			Host:     config.Registry.Host,
			Password: config.Registry.Password,
			Port:     config.Registry.Port,
			Username: config.Registry.Username,
		}
		properties.Blobstore = &PropertiesBlobstore{
			Address: turbulenceNetwork.StaticIPs(1)[0],
			Port:    2520,
			Agent: PropertiesBlobstoreAgent{
				User:     "agent",
				Password: "agent-password",
			},
		}
		properties.Agent = &PropertiesAgent{
			Mbus: fmt.Sprintf("nats://nats:password@%s:4222", turbulenceNetwork.StaticIPs(1)[0]),
		}
	}

	return Manifest{
		DirectorUUID:  config.DirectorUUID,
		Name:          config.Name,
		Releases:      []Release{turbulenceRelease, cpiRelease},
		ResourcePools: []ResourcePool{turbulenceResourcePool},
		Compilation:   compilation,
		Update:        update,
		Jobs:          []Job{apiJob},
		Networks:      []Network{turbulenceNetwork},
		Properties:    properties,
	}
}
