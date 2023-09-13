package stack_test

import (
	stack "go-tdd/stack/testify"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StackSuite struct {
	suite.Suite
	stack *stack.Stack
}

func (s *StackSuite) SetupTest() {
	s.stack = stack.NewStack()
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}

func (s *StackSuite) TestEmpty() {
	s.True(s.stack.IsEmpty())
}

func (s *StackSuite) TestNotEmpty() {
	s.stack.Bury("Red")
	s.False(s.stack.IsEmpty())
}

func (s *StackSuite) TestEmptySizeZero() {
	s.Zero(s.stack.Size())
}

func (s *StackSuite) TestSizeTwo() {
	s.stack.Bury("Red")
	s.stack.Bury("Blue")

	s.Equal(2, s.stack.Size())
}

func (s *StackSuite) TestUnburySingleItem() {
	s.stack.Bury("Red")

	s.Equal("Red", s.stack.Unbury())
	s.Zero(s.stack.Size())
	s.True(s.stack.IsEmpty())
}

func (s *StackSuite) TestUnburyMultipleItems() {
	s.stack.Bury("Red")
	s.stack.Bury("Blue")

	s.Equal("Blue", s.stack.Unbury())
	s.Equal(1, s.stack.Size())
	s.False(s.stack.IsEmpty())

	s.Equal("Red", s.stack.Unbury())
	s.Equal(0, s.stack.Size())
	s.True(s.stack.IsEmpty())
}