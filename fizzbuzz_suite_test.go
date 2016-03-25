package fizzbuzz_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFizzbuzz(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fizzbuzz Suite")
}
