package boshio

import (
	"encoding/json"
	"fmt"
	"github.com/zrob/boshler/bosh_file"
	"net/http"
)

type MetadataFetcher interface {
	Fetch(bosh_file.Release) (ReleaseMetadata, error)
}

type fetcher struct{}

type ReleaseVersion struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Url     string `json:"url"`
}

const releaseMetadataUrl = "https://bosh.io/api/v1/releases/github.com"

type ReleaseMetadata struct {
	Versions []ReleaseVersion
}

func NewMetadataFetcher() MetadataFetcher {
	return &fetcher{}
}

func (f *fetcher) Fetch(release bosh_file.Release) (ReleaseMetadata, error) {
	url := fmt.Sprintf("%s/%s/%s", releaseMetadataUrl, release.Repository, release.Name)

	resp, err := http.Get(url)
	if err != nil {
		return ReleaseMetadata{}, err
	}
	defer resp.Body.Close()

	var metadata ReleaseMetadata
	jsonParser := json.NewDecoder(resp.Body)
	if err = jsonParser.Decode(&metadata); err != nil {
		return ReleaseMetadata{}, err
	}

	return metadata, nil
}
