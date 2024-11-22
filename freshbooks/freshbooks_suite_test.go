package freshbooks_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFreshbooksSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Freshbooks Client Suite")
}
