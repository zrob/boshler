package boshio_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"io/ioutil"
	"path/filepath"

	"github.com/zrob/boshler/boshio"
)

var _ = Describe("Downloader", func() {
	downloader := boshio.NewDownloader()

	Describe("Download", func() {
		var tempdir string

		BeforeEach(func() {
			var err error
			tempdir, err = ioutil.TempDir("", "boshler-download-release")
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			err := os.RemoveAll(tempdir)
			Expect(err).To(BeNil())
		})

		It("downloads the release to the specified location", func() {
			url := "https://bosh.io/d/github.com/cloudfoundry-community/ntp-release?v=2"
			targetFile := filepath.Join(tempdir, "download-test-filename")

			err := downloader.Download(url, targetFile)
			Expect(err).To(BeNil())

			Expect(targetFile).To(BeAnExistingFile())
		})
	})
})
