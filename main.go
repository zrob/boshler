package main

import (
	"fmt"
	"os/user"
	"path/filepath"
	"sync"

	"github.com/zrob/boshler/archiver"
	"github.com/zrob/boshler/bosh_cli"
	"github.com/zrob/boshler/bosh_file"
	"github.com/zrob/boshler/boshio"
)

func main() {
	boshfile := parseBoshFile()
	displayCurrentTarget()
	archiveDir := getArchiveDir()

	var wg sync.WaitGroup
	wg.Add(len(boshfile.Releases))

	for _, release := range boshfile.Releases {
		go func(release bosh_file.Release) {
			defer wg.Done()
			cacheAndUploadRelease(release, archiveDir)
		}(release)
	}

	wg.Wait()
}

func cacheAndUploadRelease(release bosh_file.Release, archiveDir string) {
	fetcher := boshio.NewMetadataFetcher()
	archiver := archiver.NewArchiver(archiveDir)

	metadata, err := fetcher.Fetch(release)
	if err != nil {
		panic(err.Error())
	}

	var releaseVersion boshio.ReleaseVersion
	if release.Version == "" {
		releaseVersion = metadata.Latest()
	} else {
		releaseVersion, err = metadata.Version(release.Version)
		if err != nil {
			panic(err.Error())
		}
	}

	path, err := archiver.Store(releaseVersion)
	if err != nil {
		panic(err.Error())
	}

	err = bosh_cli.UploadRelease(path)
	if err != nil {
		println(err.Error())
		panic(err.Error())
	}
}

func displayCurrentTarget() {
	target, err := bosh_cli.GetTarget()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(target)
}

func parseBoshFile() bosh_file.BoshFile {
	boshfile, err := bosh_file.ParseFile("BOSHFILE")
	if err != nil {
		panic(err.Error())
	}
	return boshfile
}

func getArchiveDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err.Error())
	}
	return filepath.Join(usr.HomeDir, ".boshler", "releases")
}
