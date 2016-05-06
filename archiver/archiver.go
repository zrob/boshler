package archiver

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/zrob/boshler/boshio"
)

type Archiver interface {
	StoreRelease(boshio.ReleaseVersion) (string, error)
	StoreStemcell(boshio.StemcellVersion) (string, error)
}

type archiver struct {
	storePath string
}

func NewArchiver(storePath string) Archiver {
	return &archiver{
		storePath: storePath,
	}
}

func (a *archiver) StoreRelease(release boshio.ReleaseVersion) (string, error) {
	downloader := boshio.NewDownloader()

	releaseDir := filepath.Join(a.storePath, "releases", release.ReleaseName())
	targetFile := filepath.Join(releaseDir, release.FileName())

	_, err := os.Stat(targetFile)
	if err == nil {
		fmt.Printf("Using %s %s\n", release.ReleaseName(), release.Version)
		return targetFile, nil
	}

	err = os.MkdirAll(releaseDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	fmt.Printf("Downloading %s %s\n", release.ReleaseName(), release.Version)
	err = downloader.Download(release.Url, targetFile)
	if err != nil {
		return "", err
	}
	fmt.Printf("Done downloading %s %s\n", release.ReleaseName(), release.Version)

	return targetFile, nil
}

func (a *archiver) StoreStemcell(stemcell boshio.StemcellVersion) (string, error) {
	downloader := boshio.NewDownloader()

	releaseDir := filepath.Join(a.storePath, "stemcells", stemcell.Name)
	targetFile := filepath.Join(releaseDir, stemcell.FileName())

	_, err := os.Stat(targetFile)
	if err == nil {
		fmt.Printf("Using %s %s\n", stemcell.Name, stemcell.Version)
		return targetFile, nil
	}

	err = os.MkdirAll(releaseDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	fmt.Printf("Downloading %s %s\n", stemcell.Name, stemcell.Version)
	err = downloader.Download(stemcell.Url(), targetFile)
	if err != nil {
		return "", err
	}
	fmt.Printf("Done downloading %s %s\n", stemcell.Name, stemcell.Version)

	return targetFile, nil
}
