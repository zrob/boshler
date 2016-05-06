package boshio

import (
	"io"
	"net/http"
	"os"
)

type Downloader interface {
	Download(string, string) error
}

type downloader struct{}

func NewDownloader() Downloader {
	return &downloader{}
}

func (d *downloader) Download(url string, targetFile string) error {
	out, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
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
