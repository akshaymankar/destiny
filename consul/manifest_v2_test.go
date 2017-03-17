package consul_test

import (
	"io/ioutil"

	"github.com/pivotal-cf-experimental/destiny/consul"
	"github.com/pivotal-cf-experimental/gomegamatchers"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ManifestV2", func() {
	Describe("NewManifestV2", func() {
		It("returns a YAML representation of the consul manifest", func() {
			consulManifest, err := ioutil.ReadFile("fixtures/consul_manifest_v2.yml")
			Expect(err).NotTo(HaveOccurred())

			manifest, err := consul.NewManifestV2(consul.ConfigV2{
				DirectorUUID: "some-director-uuid",
				Name:         "some-manifest-name",
				AZs:          []string{"z1", "z2"},
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(manifest).To(gomegamatchers.MatchYAML(consulManifest))
		})
	})
})
