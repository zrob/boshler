package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"sync"

	"github.com/zrob/boshler/archiver"
	"github.com/zrob/boshler/bosh_cli"
	"github.com/zrob/boshler/bosh_file"
	"github.com/zrob/boshler/boshio"
	"github.com/zrob/boshler/work_pool"
	"github.com/zrob/boshler/bosh_manifest"
)

func main() {
	boshfile := getBoshFile()
	displayCurrentTarget()
	archiveDir := getArchiveDir()
	downloadPool := work_pool.NewWorkPool(10)
	uploadPool := work_pool.NewWorkPool(5)

	downloadPool.Start()
	uploadPool.Start()

	for _, release := range boshfile.Releases {
		func(release bosh_file.Release) {
			downloadPool.Submit(func() {
				cacheAndUploadRelease(release, archiveDir, uploadPool)
			})
		}(release)
	}

	var wg sync.WaitGroup
	wg.Add(len(boshfile.Stemcells))

	for _, stemcell := range boshfile.Stemcells {
		go func(stemcell bosh_file.Stemcell) {
			defer wg.Done()

			cacheAndUploadStemcell(stemcell, archiveDir)
		}(stemcell)
	}

	wg.Wait()

	downloadPool.Stop()
	uploadPool.Stop()
}

func cacheAndUploadRelease(release bosh_file.Release, archiveDir string, pool work_pool.WorkPool) {
	fetcher := boshio.NewMetadataFetcher()
	archiver := archiver.NewArchiver(archiveDir)

	metadata, err := fetcher.FetchRelease(release)
	if err != nil {
		fmt.Printf("%v", release)
		panic(err.Error())
	}

	releaseVersion := selectReleaseVersion(release, metadata)

	path, err := archiver.StoreRelease(releaseVersion)
	if err != nil {
		panic(err.Error())
	}

	pool.Submit(func() {
		fmt.Printf("Uploading %s %s\n", releaseVersion.ReleaseName(), releaseVersion.Version)
		err := bosh_cli.UploadRelease(path)
		if err != nil {
			println(err.Error())
			panic(err.Error())
		}
		fmt.Printf("\x1b[32;1mDone uploading %s %s\x1b[0m\n", releaseVersion.ReleaseName(), releaseVersion.Version)
	})
}

func cacheAndUploadStemcell(stemcell bosh_file.Stemcell, archiveDir string) {
	fetcher := boshio.NewMetadataFetcher()
	archiver := archiver.NewArchiver(archiveDir)

	metadata, err := fetcher.FetchStemcell(stemcell.Name)
	if err != nil {
		panic(err.Error())
	}

	stemcellVersion := selectStemcellVersion(stemcell, metadata)

	path, err := archiver.StoreStemcell(stemcellVersion)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Uploading %s %s\n", stemcellVersion.Name, stemcellVersion.Version)
	err = bosh_cli.UploadStemcell(path)
	if err != nil {
		println(err.Error())
		panic(err.Error())
	}
	fmt.Printf("\x1b[32;1mDone uploading %s %s\x1b[0m\n", stemcellVersion.Name, stemcellVersion.Version)
}

func selectReleaseVersion(release bosh_file.Release, metadata boshio.ReleaseMetadata) boshio.ReleaseVersion {
	var releaseVersion boshio.ReleaseVersion

	if release.Version == "" {
		releaseVersion = metadata.Latest()
	} else {
		var err error
		releaseVersion, err = metadata.Version(release.Version)
		if err != nil {
			fmt.Printf("%v", release)
			panic(err.Error())
		}
	}

	return releaseVersion
}

func selectStemcellVersion(stemcell bosh_file.Stemcell, metadata boshio.StemcellMetadata) boshio.StemcellVersion {
	var stemcellVersion boshio.StemcellVersion

	if stemcell.Version == "" {
		stemcellVersion = metadata.Latest()
	} else {
		var err error
		stemcellVersion, err = metadata.Version(stemcell.Version)
		if err != nil {
			panic(err.Error())
		}
	}

	return stemcellVersion
}

func displayCurrentTarget() {
	target, err := bosh_cli.GetTarget()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(target)
}

func getBoshFile() bosh_file.BoshFile {
	var boshfile bosh_file.BoshFile
	var err error

	if _, err := os.Stat("BOSHFILE"); err == nil {
		boshfile, err = bosh_file.ParseFile("BOSHFILE")
	} else if _, err := os.Stat("cf-deployment.yml"); err == nil {
		boshfile, err = bosh_manifest.ParseFile("cf-deployment.yml")
	} else {
		fmt.Println("Cannot find BOSHFILE. Are you sure you have it in your working directory?")
		os.Exit(1)
	}

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
	return filepath.Join(usr.HomeDir, ".boshler")
}
