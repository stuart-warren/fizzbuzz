package fizzbuzz_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/ginkgo/reporters"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Fizzbuzz Suite", []Reporter{junitReporter})
}
