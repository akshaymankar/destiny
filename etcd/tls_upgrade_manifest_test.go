package etcd_test

import (
	"io/ioutil"

	"github.com/pivotal-cf-experimental/destiny/consul"
	"github.com/pivotal-cf-experimental/destiny/core"
	"github.com/pivotal-cf-experimental/destiny/etcd"
	"github.com/pivotal-cf-experimental/destiny/iaas"
	"github.com/pivotal-cf-experimental/gomegamatchers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NewTLSUpgradeManifest", func() {
	It("generates a valid bosh-lite manifest", func() {
		manifest := etcd.NewTLSUpgradeManifest(etcd.Config{
			DirectorUUID: "some-director-uuid",
			Name:         "etcd-some-random-guid",
			IPRange:      "10.244.4.0/24",
		}, iaas.NewWardenConfig())

		Expect(manifest.DirectorUUID).To(Equal("some-director-uuid"))

		Expect(manifest.Name).To(Equal("etcd-some-random-guid"))

		Expect(manifest.Releases).To(HaveLen(2))

		Expect(manifest.Releases).To(ContainElement(core.Release{
			Name:    "etcd",
			Version: "latest",
		}))

		Expect(manifest.Releases).To(ContainElement(core.Release{
			Name:    "consul",
			Version: "latest",
		}))

		Expect(manifest.Compilation).To(Equal(core.Compilation{
			Network:             "etcd1",
			ReuseCompilationVMs: true,
			Workers:             3,
		}))

		Expect(manifest.Update).To(Equal(core.Update{
			Canaries:        1,
			CanaryWatchTime: "1000-180000",
			MaxInFlight:     1,
			Serial:          true,
			UpdateWatchTime: "1000-180000",
		}))

		Expect(manifest.ResourcePools).To(HaveLen(1))

		Expect(manifest.ResourcePools).To(ContainElement(core.ResourcePool{
			Name:    "etcd_z1",
			Network: "etcd1",
			Stemcell: core.ResourcePoolStemcell{
				Name:    "bosh-warden-boshlite-ubuntu-trusty-go_agent",
				Version: "latest",
			},
		}))

		Expect(manifest.Networks).To(HaveLen(1))

		Expect(manifest.Networks).To(ContainElement(core.Network{
			Name: "etcd1",
			Subnets: []core.NetworkSubnet{
				{
					CloudProperties: core.NetworkSubnetCloudProperties{Name: "random"},
					Gateway:         "10.244.4.1",
					Range:           "10.244.4.0/24",
					Reserved: []string{
						"10.244.4.2-10.244.4.3",
						"10.244.4.50-10.244.4.254",
					},
					Static: []string{
						"10.244.4.4",
						"10.244.4.5",
						"10.244.4.6",
						"10.244.4.7",
						"10.244.4.8",
						"10.244.4.9",
						"10.244.4.10",
						"10.244.4.11",
						"10.244.4.12",
						"10.244.4.13",
						"10.244.4.14",
						"10.244.4.15",
						"10.244.4.16",
						"10.244.4.17",
						"10.244.4.18",
						"10.244.4.19",
						"10.244.4.20",
						"10.244.4.21",
						"10.244.4.22",
						"10.244.4.23",
						"10.244.4.24",
						"10.244.4.25",
						"10.244.4.26",
						"10.244.4.27",
						"10.244.4.28",
						"10.244.4.29",
						"10.244.4.30",
						"10.244.4.31",
						"10.244.4.32",
						"10.244.4.33",
						"10.244.4.34",
						"10.244.4.35",
						"10.244.4.36",
						"10.244.4.37",
						"10.244.4.38",
						"10.244.4.39",
						"10.244.4.40",
					},
				},
			},
			Type: "manual",
		}))

		Expect(manifest.Jobs).To(HaveLen(4))

		Expect(manifest.Jobs[0]).To(Equal(core.Job{
			Name:      "consul_z1",
			Instances: 3,
			Networks: []core.JobNetwork{{
				Name: "etcd1",
				StaticIPs: []string{
					"10.244.4.9",
					"10.244.4.10",
					"10.244.4.11",
				},
			}},
			PersistentDisk: 1024,
			ResourcePool:   "etcd_z1",
			Templates: []core.JobTemplate{
				{
					Name:    "consul_agent",
					Release: "consul",
				},
			},
			Properties: &core.JobProperties{
				Consul: &core.JobPropertiesConsul{
					Agent: core.JobPropertiesConsulAgent{
						Mode: "server",
					},
				},
			},
		}))

		Expect(manifest.Jobs[1]).To(Equal(core.Job{
			Name:      "etcd_tls_z1",
			Instances: 3,
			Networks: []core.JobNetwork{{
				Name: "etcd1",
				StaticIPs: []string{
					"10.244.4.30",
					"10.244.4.31",
					"10.244.4.32",
				},
			}},
			PersistentDisk: 1024,
			ResourcePool:   "etcd_z1",
			Templates: []core.JobTemplate{
				{
					Name:    "consul_agent",
					Release: "consul",
				},
				{
					Name:    "etcd",
					Release: "etcd",
				},
			},
			Properties: &core.JobProperties{
				Consul: &core.JobPropertiesConsul{
					Agent: core.JobPropertiesConsulAgent{
						Services: core.JobPropertiesConsulAgentServices{
							"etcd": core.JobPropertiesConsulAgentService{},
						},
					},
				},
				Etcd: &core.JobPropertiesEtcd{
					Cluster: []core.JobPropertiesEtcdCluster{{
						Instances: 3,
						Name:      "etcd_tls_z1",
					}},
					Machines: []string{
						"etcd.service.cf.internal",
					},
					PeerRequireSSL:                  true,
					RequireSSL:                      true,
					HeartbeatIntervalInMilliseconds: 50,
					AdvertiseURLsDNSSuffix:          "etcd.service.cf.internal",
					CACert:                          etcd.CACert,
					ClientCert:                      etcd.ClientCert,
					ClientKey:                       etcd.ClientKey,
					PeerCACert:                      etcd.PeerCACert,
					PeerCert:                        etcd.PeerCert,
					PeerKey:                         etcd.PeerKey,
					ServerCert:                      etcd.ServerCert,
					ServerKey:                       etcd.ServerKey,
				},
			},
		}))

		Expect(manifest.Jobs[2]).To(Equal(core.Job{
			Name:      "etcd_z1",
			Instances: 1,
			Networks: []core.JobNetwork{{
				Name: "etcd1",
				StaticIPs: []string{
					"10.244.4.4",
				},
			}},
			PersistentDisk: 1024,
			ResourcePool:   "etcd_z1",
			Templates: []core.JobTemplate{
				{
					Name:    "consul_agent",
					Release: "consul",
				},
				{
					Name:    "etcd_proxy",
					Release: "etcd",
				},
			},
		}))

		Expect(manifest.Jobs[3]).To(Equal(core.Job{
			Name:      "testconsumer_z1",
			Instances: 5,
			Networks: []core.JobNetwork{{
				Name: "etcd1",
				StaticIPs: []string{
					"10.244.4.12",
					"10.244.4.13",
					"10.244.4.14",
					"10.244.4.15",
					"10.244.4.16",
				},
			}},
			PersistentDisk: 1024,
			ResourcePool:   "etcd_z1",
			Templates: []core.JobTemplate{
				{
					Name:    "consul_agent",
					Release: "consul",
				},
				{
					Name:    "etcd_testconsumer",
					Release: "etcd",
				},
			},
			Properties: &core.JobProperties{
				EtcdTestConsumer: &core.JobPropertiesEtcdTestConsumer{
					Etcd: core.JobPropertiesEtcdTestConsumerEtcd{
						Machines: []string{
							"etcd.service.cf.internal",
						},
						RequireSSL: true,
						CACert:     etcd.CACert,
						ClientCert: etcd.ClientCert,
						ClientKey:  etcd.ClientKey,
					},
				},
			},
		}))

		Expect(manifest.Properties.Consul).To(Equal(&consul.PropertiesConsul{
			Agent: consul.PropertiesConsulAgent{
				Domain: "cf.internal",
				Servers: consul.PropertiesConsulAgentServers{
					Lan: []string{
						"10.244.4.9",
						"10.244.4.10",
						"10.244.4.11",
					},
				},
			},
			CACert:      consul.CACert,
			AgentCert:   consul.DC1AgentCert,
			AgentKey:    consul.DC1AgentKey,
			ServerCert:  consul.DC1ServerCert,
			ServerKey:   consul.DC1ServerKey,
			EncryptKeys: []string{consul.EncryptKey},
		}))

		Expect(manifest.Properties.EtcdProxy).To(Equal(&etcd.PropertiesEtcdProxy{
			RequireSSL: true,
			Port:       4001,
			Etcd: etcd.PropertiesEtcdProxyEtcd{
				URL:        "https://etcd.service.cf.internal:4001",
				CACert:     etcd.CACert,
				ClientCert: etcd.ClientCert,
				ClientKey:  etcd.ClientKey,
			},
		}))
	})

	It("returns a YAML representation of the etcd manifest", func() {
		etcdManifest, err := ioutil.ReadFile("fixtures/tls_upgrade.yml")
		Expect(err).NotTo(HaveOccurred())

		manifest := etcd.NewTLSUpgradeManifest(etcd.Config{
			DirectorUUID: "some-director-uuid",
			Name:         "etcd",
			IPRange:      "10.244.4.0/24",
			Secrets: etcd.ConfigSecrets{
				Consul: consul.ConfigSecretsConsul{
					EncryptKey: "consul-encrypt-key",
					AgentKey:   "consul-agent-key",
					AgentCert:  "consul-agent-cert",
					ServerKey:  "consul-server-key",
					ServerCert: "consul-server-cert",
					CACert:     "consul-ca-cert",
				},
				Etcd: etcd.ConfigSecretsEtcd{
					CACert:     "etcd-ca-cert",
					ClientCert: "etcd-client-cert",
					ClientKey:  "etcd-client-key",
					PeerCACert: "etcd-peer-ca-cert",
					PeerCert:   "etcd-peer-cert",
					PeerKey:    "etcd-peer-key",
					ServerCert: "etcd-server-cert",
					ServerKey:  "etcd-server-key",
				},
			},
		}, iaas.NewWardenConfig())

		yaml, err := manifest.ToYAML()
		Expect(err).NotTo(HaveOccurred())
		Expect(yaml).To(gomegamatchers.MatchYAML(etcdManifest))
	})
})
