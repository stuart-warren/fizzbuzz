package fizzbuzz_test

import (
	"fmt"
	"testing"

	. "github.com/stuart-warren/fizzbuzz"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fizzbuzz", func() {

	Context("With an integer that is a multiple of 3", func() {
		It("should return \"fizz\"", func() {
			Expect(FizzBuzz(3)).To(Equal("fizz"))
			Expect(FizzBuzz(99)).To(Equal("fizz"))
		})
	})

	Context("With an integer that is a multiple of 5", func() {
		It("should return \"buzz\"", func() {
			Expect(FizzBuzz(5)).To(Equal("buzz"))
			Expect(FizzBuzz(2000)).To(Equal("buzz"))
		})
	})

	Context("With an integer that is a multiple of 3 and 5", func() {
		It("should return \"fizzbuzz\"", func() {
			Expect(FizzBuzz(15)).To(Equal("fizzbuzz"))
			Expect(FizzBuzz(150)).To(Equal("fizzbuzz"))
		})
	})

	Context("Otherwise", func() {
		It("should return the number", func() {
			Expect(FizzBuzz(4)).To(Equal("4"))
			Expect(FizzBuzz(17)).To(Equal("17"))
		})
	})
})

func ExampleFizzBuzz() {
	fmt.Println(FizzBuzz(15))
	// Output: fizzbuzz
}

func BenchmarkFizzBuzz(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FizzBuzz(n)
	}
}
