package boshio

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type ReleaseDownloader interface {
	Download(ReleaseVersion, string) error
}

type downloader struct{}

func NewReleaseDownloader() ReleaseDownloader {
	return &downloader{}
}

func (d *downloader) Download(release ReleaseVersion, targetDir string) error {
	fileName := fmt.Sprintf("%s-%s.tgz", filepath.Base(release.Name), release.Version)
	fullPath := filepath.Join(targetDir, fileName)

	out, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(release.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
