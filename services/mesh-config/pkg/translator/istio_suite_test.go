package translator_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIstio(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Istio RBAC Suite")
}