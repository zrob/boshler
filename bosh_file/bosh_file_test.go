package bosh_file_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/zrob/boshler/bosh_file"
)

var _ = Describe("BoshFile", func() {
	Describe("ParseFile", func() {
		var tempFile *os.File
		var boshfileContent []byte

		BeforeEach(func() {
			var err error
			boshfileContent = []byte(`
			{
				"releases": [
					{
						"name": "release-name",
						"repository": "release-repository",
						"version": "1.2.3"
					}
				]
			}
			`)

			tempFile, err = ioutil.TempFile("", "boshler-parse-file")
			Expect(err).To(BeNil())

			err = ioutil.WriteFile(tempFile.Name(), boshfileContent, os.ModePerm)
			Expect(err).To(BeNil())
		})

		AfterEach(func() {
			os.Remove(tempFile.Name())
		})

		It("parses a bosh file", func() {
			boshfile, err := bosh_file.ParseFile(tempFile.Name())

			Expect(err).To(BeNil())
			Expect(len(boshfile.Releases)).To(Equal(1))

			release := boshfile.Releases[0]
			Expect(release.Name).To(Equal("release-name"))
			Expect(release.Repository).To(Equal("release-repository"))
			Expect(release.Version).To(Equal("1.2.3"))
		})

	})
})
