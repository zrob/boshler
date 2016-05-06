package boshio_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"path/filepath"

	"github.com/zrob/boshler/boshio"
)

var _ = Describe("ReleaseDownloader", func() {
	releaseDownloader := boshio.NewReleaseDownloader()

	Describe("Download", func() {
		var tempdir string

		BeforeEach(func() {
			var err error
			tempdir, err = ioutil.TempDir("", "boshler-download-release")
			Expect(err).To(BeNil())
		})

		It("downloads the release to the specified location", func() {
			release := boshio.ReleaseVersion{
				Name:    "github.com/cloudfoundry-community/ntp-release",
				Version: "2",
				Url:     "https://bosh.io/d/github.com/cloudfoundry-community/ntp-release?v=2",
			}

			err := releaseDownloader.Download(release, tempdir)
			Expect(err).To(BeNil())

			Expect(filepath.Join(tempdir, "ntp-release-2.tgz")).To(BeAnExistingFile())
		})
	})
})
