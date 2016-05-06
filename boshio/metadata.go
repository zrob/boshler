package boshio

import (
	"errors"
	"fmt"
	"path/filepath"
)

type ReleaseVersion struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Url     string `json:"url"`
}

type ReleaseMetadata []ReleaseVersion

// lazily assume the boshio api is sorting this, not sure if that is true
func (m ReleaseMetadata) Latest() ReleaseVersion {
	return m[0]
}

func (m ReleaseMetadata) Version(searchVersion string) (ReleaseVersion, error) {
	for _, version := range m {
		if version.Version == searchVersion {
			return version, nil
		}
	}

	return ReleaseVersion{}, errors.New("version not found")
}

func (r ReleaseVersion) FileName() string {
	return fmt.Sprintf("%s-%s.tgz", r.ReleaseName(), r.Version)
}

func (r ReleaseVersion) ReleaseName() string {
	return filepath.Base(r.Name)
}
