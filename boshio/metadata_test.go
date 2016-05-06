package boshio_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zrob/boshler/boshio"
)

var _ = Describe("Metadata", func() {
	Describe("ReleaseMetadata", func() {
		releaseVersion1 := boshio.ReleaseVersion{
			Version: "0.33.1",
		}
		releaseVersion2 := boshio.ReleaseVersion{
			Version: "0.33.0",
		}
		releaseVersion3 := boshio.ReleaseVersion{
			Version: "0.0.1",
		}
		metadata := boshio.ReleaseMetadata{releaseVersion1, releaseVersion2, releaseVersion3}

		Describe("Latest", func() {
			It("returns the most recent version", func() {
				Expect(metadata.Latest()).To(Equal(releaseVersion1))
			})
		})

		Describe("Version", func() {
			It("returns the matching version", func() {
				version, err := metadata.Version("0.33.1")
				Expect(err).To(BeNil())
				Expect(version).To(Equal(releaseVersion1))
			})
		})
	})
})
