package etcd_test

import (
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf-experimental/destiny/etcd"
	"github.com/pivotal-cf-experimental/gomegamatchers"
)

var _ = Describe("ManifestV2", func() {
	Describe("NewManifestV2", func() {
		It("returns a YAML representation of the etcd manifest", func() {
			etcdManifest, err := ioutil.ReadFile("fixtures/etcd_manifest_v2.yml")
			Expect(err).NotTo(HaveOccurred())

			manifest, err := etcd.NewManifestV2(etcd.ConfigV2{
				DirectorUUID: "some-director-uuid",
				Name:         "etcd-some-random-guid",
				AZs:          []string{"z1", "z2"},
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(manifest).To(gomegamatchers.MatchYAML(etcdManifest))
		})
	})

	Describe("ApplyOp", func() {
		It("returns a manifest with an bosh ops change", func() {
			manifest, err := etcd.NewManifestV2(etcd.ConfigV2{
				DirectorUUID: "some-director-uuid",
				Name:         "etcd-some-random-guid",
				AZs:          []string{"z1", "z2"},
			})
			Expect(err).NotTo(HaveOccurred())
			manifest, err = etcd.ApplyOp(manifest, "replace", "/instance_groups/name=etcd/instances", 5)

			Expect(manifest).To(ContainSubstring("instances: 5"))
		})
	})
})
