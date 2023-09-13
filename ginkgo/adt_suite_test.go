package adt_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestADTSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ADT Suite")
}
