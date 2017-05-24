package bosh_manifest_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zrob/boshler/bosh_manifest"
)

var _ = Describe("BoshManifest", func() {
	Describe("ParseFile", func() {
		It("parses a bosh manifest", func() {
			boshfile, err := bosh_manifest.ParseFile("test.yml")

			Expect(err).To(BeNil())
			Expect(len(boshfile.Releases)).To(Equal(22))

			release := boshfile.Releases[0]
			Expect(release.Name).To(Equal("binary-buildpack-release"))
			Expect(release.Repository).To(Equal("cloudfoundry"))
			Expect(release.Version).To(Equal("1.0.11"))
		})
	})
})
