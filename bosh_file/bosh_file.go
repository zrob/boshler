package bosh_file

import (
	"encoding/json"
	"os"
)

type Release struct {
	Name       string `json:"name"`
	Repository string `json:"repository"`
	Version    string `json:"version"`
}

type Stemcell struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type BoshFile struct {
	Releases  []Release  `json:"releases"`
	Stemcells []Stemcell `json:"stemcells"`
}

func ParseFile(filePath string) (BoshFile, error) {
	readfile, err := os.Open(filePath)
	if err != nil {
		return BoshFile{}, err
	}
	defer readfile.Close()

	var boshfile BoshFile
	jsonParser := json.NewDecoder(readfile)
	if err = jsonParser.Decode(&boshfile); err != nil {
		return BoshFile{}, err
	}

	return boshfile, nil
}
