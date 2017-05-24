package bosh_manifest

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/zrob/boshler/bosh_file"
)

type release struct {
	Name    string `yaml:"name"`
	URL     string `yaml:"url"`
	Version string `yaml:"version"`
}

type boshManifest struct {
	Releases []release `yaml:"releases"`
}

func ParseFile(filePath string) (bosh_file.BoshFile, error) {
	readfile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return bosh_file.BoshFile{}, err
	}

	var manifest boshManifest
	if err = yaml.Unmarshal(readfile, &manifest); err != nil {
		return bosh_file.BoshFile{}, err
	}

	var boshfile bosh_file.BoshFile
	boshfile.Releases = transformReleases(manifest)

	return boshfile, nil
}

func transformReleases(manifest boshManifest) []bosh_file.Release {
	var releases []bosh_file.Release

	for _, manifestRelease := range manifest.Releases {
		if manifestRelease.URL == "" {
			continue
		}

		url, _ := url.Parse(manifestRelease.URL)

		releases = append(releases, bosh_file.Release{
			Version:    manifestRelease.Version,
			Name:       getNameFromUrl(url),
			Repository: getRepoFromUrl(url),
		})
	}

	return releases
}

func getRepoFromUrl(url *url.URL) string {
	var repo string
	parts := strings.Split(url.Path, "/")
	for index, part := range parts {
		if strings.EqualFold("github.com", part) {
			repo = parts[index+1]
			break
		}
	}
	return repo
}

func getNameFromUrl(url *url.URL) string {
	parts := strings.Split(url.Path, "/")
	return parts[len(parts)-1]
}
