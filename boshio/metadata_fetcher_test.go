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

			releaseVersion := releaseMetadata[0]
			Expect(releaseVersion.Name).To(Equal("github.com/cloudfoundry-community/ntp-release"))
			Expect(releaseVersion.Version).To(Equal("2"))
			Expect(releaseVersion.Url).To(Equal("https://bosh.io/d/github.com/cloudfoundry-community/ntp-release?v=2"))
		})
	})

	Describe("Latest", func() {
		It("returns the most recent version", func() {
			releaseVersion1 := boshio.ReleaseVersion{
				Version: "0.1.0",
			}
			releaseVersion2 := boshio.ReleaseVersion{
				Version: "0.1.1",
			}
			releaseVersion3 := boshio.ReleaseVersion{
				Version: "0.0.1",
			}

			metadata := boshio.ReleaseMetadata{releaseVersion1, releaseVersion2, releaseVersion3}

			Expect(metadata.Latest).To(Equal(releaseVersion2))
		})
	})
})
