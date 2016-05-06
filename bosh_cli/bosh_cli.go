package bosh_cli

import "os/exec"

func UploadRelease(releasePath string) error {
	cmd := exec.Command("bosh", "upload", "release", releasePath)
	err := cmd.Run()
	if err != nil {
		return err
	}
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
