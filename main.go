package main

import (
	"fmt"
	"sync"

	"github.com/zrob/boshler/archiver"
	"github.com/zrob/boshler/bosh_cli"
	"github.com/zrob/boshler/bosh_file"
	"github.com/zrob/boshler/boshio"
)

func main() {
	boshfile, err := bosh_file.ParseFile("BOSHFILE")
	if err != nil {
		panic(err.Error())
	}

	target, err := bosh_cli.GetTarget()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(target)

	var wg sync.WaitGroup
	wg.Add(len(boshfile.Releases))

	for _, release := range boshfile.Releases {
		go func(release bosh_file.Release) {
			fetcher := boshio.NewMetadataFetcher()
			archiver := archiver.NewArchiver("/tmp/blah")

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

			wg.Done()
		}(release)
	}

	wg.Wait()
}
