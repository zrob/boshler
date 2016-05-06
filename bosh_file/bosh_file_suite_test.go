package bosh_file_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBoshFile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BoshFile Suite")
}
