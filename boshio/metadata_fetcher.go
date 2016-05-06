package boshio

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/zrob/boshler/bosh_file"
)

type MetadataFetcher interface {
	FetchRelease(bosh_file.Release) (ReleaseMetadata, error)
	FetchStemcell(string) (StemcellMetadata, error)
}

type fetcher struct{}

const releaseMetadataUrl = "https://bosh.io/api/v1/releases/github.com"
const stemcellMetadataUrl = "https://bosh.io/api/v1/stemcells"

func NewMetadataFetcher() MetadataFetcher {
	return &fetcher{}
}

func (f *fetcher) FetchRelease(release bosh_file.Release) (ReleaseMetadata, error) {
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

func (f *fetcher) FetchStemcell(stemcell string) (StemcellMetadata, error) {
	url := fmt.Sprintf("%s/%s", stemcellMetadataUrl, stemcell)

	resp, err := http.Get(url)
	if err != nil {
		return StemcellMetadata{}, err
	}
	defer resp.Body.Close()

	var metadata StemcellMetadata
	jsonParser := json.NewDecoder(resp.Body)
	if err = jsonParser.Decode(&metadata); err != nil {
		return StemcellMetadata{}, err
	}

	return metadata, nil
}
