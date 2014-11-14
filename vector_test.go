package rabble_test

import (
	r "github.com/jgdavey/rabble"
	"testing"
)

func TestNode(t *testing.T) {
	n := r.NewNode()

	n.SetValue(1)

	if 1 != n.GetValue() {
		t.Fatalf("Expected 1, got: %v", n.GetValue())
	}

	n.SetValue("a string")

	if "a string" != n.GetValue() {
		t.Fatalf("Expected 'a string', got: %v", n.GetValue())
	}
}

func TestGetSetVector(t *testing.T) {
	v := r.NewVector()
	v.Cons("hello")

	ret := v.GetNth(0)
	if "hello" != ret {
		t.Fatalf("Expected 'hello', got: %v", ret)
	}
}

func TestGetSetVectorNode(t *testing.T) {
	vec := r.NewVector()
	node := r.NewNode()
	node.SetValue("hello")

	vec.Cons(&node)

	ret := vec.GetNth(0)
	if "hello" != ret {
		t.Fatalf("Expected 'hello', got: %v", ret)
	}
}
