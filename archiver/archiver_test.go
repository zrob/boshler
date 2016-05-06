package archiver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zrob/boshler/archiver"
	"github.com/zrob/boshler/boshio"
)

var _ = Describe("Archiver", func() {
	archiver := archiver.NewArchiver("/tmp/blah")

	Describe("Store", func() {
		It("saves the release", func() {
			release := boshio.ReleaseVersion{
				Name:    "github.com/cloudfoundry-community/ntp-release",
				Version: "2",
				Url:     "https://bosh.io/d/github.com/cloudfoundry-community/ntp-release?v=2",
			}

			path, err := archiver.Store(release)
			Expect(err).To(BeNil())

			Expect(path).To(Equal("/tmp/blah/ntp-release/ntp-release-2.tgz"))
			Expect(path).To(BeARegularFile())
		})

	})

})
