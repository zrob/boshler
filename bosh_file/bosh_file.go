package bosh_file

type Release struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
	Version    string `json:"version"`
}

type BoshFile struct {
	Releases []Release `json:"releases"`
}
