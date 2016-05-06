package boshio

import (
	"io"
	"net/http"
	"os"
)

type ReleaseDownloader interface {
	Download(ReleaseVersion, string) error
}

type downloader struct{}

func NewReleaseDownloader() ReleaseDownloader {
	return &downloader{}
}

func (d *downloader) Download(release ReleaseVersion, targetFile string) error {
	out, err := os.Create(targetFile)
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
