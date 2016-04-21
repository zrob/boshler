package boshio_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zrob/boshler/bosh_file"
	"github.com/zrob/boshler/boshio"
)

var _ = Describe("MetadataFetcher", func() {
	fetcher := boshio.NewMetadataFetcher()

	Describe("Fetch", func() {
		It("fetches release metadata", func() {
			release := bosh_file.Release{
				Name:       "ntp-release",
				Repository: "cloudfoundry-community",
			}

			releaseMetadata, err := fetcher.Fetch(release)

			Expect(err).ToNot(HaveOccurred())
			Expect(releaseMetadata).To(HaveLen(1))

			releaseVersion := releaseMetadata.Versions[0]
			Expect(releaseVersion.Name).To(Equal("github.com/cloudfoundry-community/ntp-release"))
			Expect(releaseVersion.Version).To(Equal("2"))
			Expect(releaseVersion.Url).To(Equal("https://bosh.io/d/github.com/cloudfoundry-community/ntp-release?v=2"))
		})
	})
})
