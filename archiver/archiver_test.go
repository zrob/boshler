package archiver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zrob/boshler/archiver"
	"github.com/zrob/boshler/boshio"
)

var _ = Describe("Archiver", func() {
	archiver := archiver.NewArchiver("/tmp/blah")

	Describe("StoreRelease", func() {
		It("saves the release", func() {
			release := boshio.ReleaseVersion{
				Name:    "github.com/cloudfoundry-community/ntp-release",
				Version: "2",
				Url:     "https://bosh.io/d/github.com/cloudfoundry-community/ntp-release?v=2",
			}

			path, err := archiver.StoreRelease(release)
			Expect(err).To(BeNil())

			Expect(path).To(Equal("/tmp/blah/releases/ntp-release/ntp-release-2.tgz"))
			Expect(path).To(BeARegularFile())
		})

	})

	Describe("StoreStemcell", func() {
		It("saves the stemcell", func() {
			release := boshio.StemcellVersion{
				Name:    "stemcell-name",
				Version: "3147",
				Regular: boshio.StemcellRegular{
					Url: "https://d26ekeud912fhb.cloudfront.net/bosh-stemcell/aws/light-bosh-stemcell-3232.1-aws-xen-hvm-ubuntu-trusty-go_agent.tgz",
				},
			}

			path, err := archiver.StoreStemcell(release)
			Expect(err).To(BeNil())

			Expect(path).To(Equal("/tmp/blah/stemcells/stemcell-name/bosh-stemcell-3147-stemcell-name.tgz"))
			Expect(path).To(BeARegularFile())
		})

		Context("light stemcell", func() {
			It("saves the stemcell", func() {

				release := boshio.StemcellVersion{
					Name:    "stemcell-name",
					Version: "3147",
					Light: boshio.StemcellLight{
						Url: "https://d26ekeud912fhb.cloudfront.net/bosh-stemcell/aws/light-bosh-stemcell-3232.1-aws-xen-hvm-ubuntu-trusty-go_agent.tgz",
					},
				}

				path, err := archiver.StoreStemcell(release)
				Expect(err).To(BeNil())

				Expect(path).To(Equal("/tmp/blah/stemcells/stemcell-name/light-bosh-stemcell-3147-stemcell-name.tgz"))
				Expect(path).To(BeARegularFile())
			})
		})
	})
})
