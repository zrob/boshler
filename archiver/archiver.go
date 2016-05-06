package archiver

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/zrob/boshler/boshio"
)

type Archiver interface {
	Store(boshio.ReleaseVersion) (string, error)
}

type archiver struct {
	storePath string
}

func NewArchiver(storePath string) Archiver {
	return &archiver{
		storePath: storePath,
	}
}

func (a *archiver) Store(release boshio.ReleaseVersion) (string, error) {
	downloader := boshio.NewDownloader()

	releaseDir := filepath.Join(a.storePath, release.ReleaseName())
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
