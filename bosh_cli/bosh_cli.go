package bosh_cli

import (
	"fmt"
	"os/exec"
)

func UploadRelease(releasePath string) error {
	fmt.Printf("Uploading release %s...\n", releasePath)
	cmd := exec.Command("bosh", "upload", "release", releasePath)
	err := cmd.Run()
	if err != nil {
		return err
	}
	fmt.Printf("Done uploading release %s.\n", releasePath)
	return nil
}

func GetTarget() (string, error) {
	cmd := exec.Command("bosh", "target")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
