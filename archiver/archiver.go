package archiver

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/zrob/boshler/boshio"
)

type Archiver interface {
	Store(boshio.ReleaseVersion) error
}

type archiver struct {
	storePath string
}

func NewArchiver(storePath string) Archiver {
	return &archiver{
		storePath: storePath,
	}
}

func (a *archiver) Store(release boshio.ReleaseVersion) error {
	downloader := boshio.NewReleaseDownloader()

	releaseDir := filepath.Join(a.storePath, release.ReleaseName())
	targetFile := filepath.Join(releaseDir, release.FileName())

	if _, err := os.Stat(targetFile); err == nil {
		fmt.Printf("Using %s.\n", release.FileName())
		return nil
	}

	os.MkdirAll(releaseDir, os.ModePerm)

	fmt.Printf("Downloading %s...\n", release.FileName())
	err := downloader.Download(release, targetFile)
	if err != nil {
		return err
	}
	fmt.Printf("Done downloading %s.\n", release.FileName())

	return nil
}
