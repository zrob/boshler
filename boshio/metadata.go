package boshio

import "errors"

type ReleaseVersion struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Url     string `json:"url"`
}

type ReleaseMetadata []ReleaseVersion

func (m ReleaseMetadata) Latest() ReleaseVersion {
	var latestVersion ReleaseVersion

	for _, version := range m {
		if version.Version > latestVersion.Version {
			latestVersion = version
		}
	}

	return latestVersion
}

func (m ReleaseMetadata) Version(searchVersion string) (ReleaseVersion, error) {
	for _, version := range m {
		if version.Version == searchVersion {
			return version, nil
		}
	}

	return ReleaseVersion{}, errors.New("version not found")
}
