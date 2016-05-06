package bosh_cli

import "os/exec"

func UploadRelease(path string) error {
	cmd := exec.Command("bosh", "upload", "release", path)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func UploadStemcell(path string) error {
	cmd := exec.Command("bosh", "upload", "stemcell", path)
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
