package fizzbuzz_test

import (
	fizzbuzz "go-tdd/tcr"
	"testing"

	"github.com/stretchr/testify/suite"
)

type FizzBuzzSuite struct {
	suite.Suite
}

func TestFizzBuzzSuite(t *testing.T) {
	suite.Run(t, new(FizzBuzzSuite))
}

func (s *FizzBuzzSuite) TestOneUnchanged() {
	r := fizzbuzz.Run([]int{1})

	s.Equal([]string{"1"}, r)
}

func (s *FizzBuzzSuite) TestTwoUnchanged() {
	r := fizzbuzz.Run([]int{2})

	s.Equal([]string{"2"}, r)
}

func (s *FizzBuzzSuite) TestThreeChangedFizz() {
	r := fizzbuzz.Run([]int{3})

	s.Equal([]string{"Fizz"}, r)
}

func (s *FizzBuzzSuite) TestFiveChangedBuzz() {
	r := fizzbuzz.Run([]int{5})

	s.Equal([]string{"Buzz"}, r)
}

func (s *FizzBuzzSuite) TestFifteenChangedFizzBuzz() {
	r := fizzbuzz.Run([]int{15})

	s.Equal([]string{"FizzBuzz"}, r)
}

func (s *FizzBuzzSuite) TestManyNumbers() {
	r := fizzbuzz.Run([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})

	s.Equal([]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}, r)
}
