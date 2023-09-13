package stack

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	s := NewStack()
	if s.Empty() != true {
		t.Error("stack was not empty")
	}
}


func TestNotEmpty(t *testing.T) {
	s := NewStack()
	s.Add("Bob")
	if s.Empty() != false {
		t.Error("stack was empty")
	}
}

func TestSizeZero(t *testing.T) {
	s := NewStack()
	if s.Size() != 0 {
		t.Error("size was not zero")
	}
}

func TestSizeOne(t *testing.T) {
	s := NewStack()
	s.Add("Bob")
	if s.Size() != 1 {
		t.Error("Incorrect size")
		t.Log("Expected 1")
		t.Logf("Actual: %d", s.Size())
	}
}

func TestSizeThree(t *testing.T) {
	s := NewStack()
	s.Add("Bob")
	s.Add("Victor")
	s.Add("Andrew")

	if s.Size() != 3 {
		t.Error("Incorrect size")
		t.Log("Expected 3")
		t.Logf("Actual: %d", s.Size())
	}
}

func TestPopOne(t *testing.T) {
	s := NewStack()
	s.Add("Bob")

	v := s.Pop()
	if v != "Bob" {
		t.Errorf("Expected Bob, found %s", v)
	}

	if s.Size() != 0 {
		t.Errorf("Expected 0, found %d", s.Size())
	}
}

func TestPopTwo(t *testing.T) {
	s := NewStack()
	s.Add("Bob")
	s.Add("Alice")

	v := s.Pop()
	if v != "Alice" {
		t.Errorf("Expected Alice, found %s", v)
	}

	v2 := s.Pop()
	if v2 != "Bob" {
		t.Errorf("Expected Bob, found %s", v2)
	}
}