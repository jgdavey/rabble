package rabble_test

import (
	r "github.com/jgdavey/rabble"
	"testing"
)

func TestConsGetVector(t *testing.T) {
	v1 := r.NewVector()
	obj := r.EnsureObject("hello")
	v2 := v1.Cons(&obj)

	if v2.Count() != 1 {
		t.Fatalf("Expected count to be 1")
	}

	if _, ok := v1.GetNth(0); ok {
		t.Fatalf("Mutated original vector")
	}

	ret, _ := v2.GetNth(0)
	if ret == nil {
		t.Fatalf("Expected something, got nil")
	}

	if "hello" != ret.Value() {
		t.Fatalf("Expected 'hello', got: %v", ret)
	}

}

func TestVectorSetAppend(t *testing.T) {
	v1 := r.NewVector()
	obj := r.EnsureObject("hello")
	v2, ok := v1.SetNth(0, &obj)

	if !ok {
		t.Fatalf("Unable to set")
	}

	if v2.Count() != 1 {
		t.Fatalf("Expected count to be 1")
	}

	if v1.Count() != 0 {
		t.Fatalf("Mutated original vector")
	}

	ret, _ := v2.GetNth(0)
	if ret == nil {
		t.Fatalf("Expected something, got nil")
	}

	if "hello" != ret.Value() {
		t.Fatalf("Expected 'hello', got: %v", ret)
	}
}

func TestVectorMultipleCons(t *testing.T) {
	vec := r.NewVector()
	orig := vec
	max := 10
	for i := 0; i < max; i++ {
		obj := r.EnsureObject(i)
		vec = vec.Cons(&obj)
	}

	if vec.Count() != uint32(max) {
		t.Fatalf("Expected %v items, got: %v", max, vec.Count())
	}

	if orig.Count() != 0 || orig.RootArray()[0] != nil {
		t.Fatalf("Mutated original")
	}

	val, ok := vec.GetNth(uint32(max - 1))
	if !ok {
		t.Fatalf("Unable to get item at idx %v", max-1)
	}

	if val.Value() != (max - 1) {
		t.Fatalf("Expected %v, Got %v", (max - 1), val.Value())
	}
}

func TestVectorConsPastTail(t *testing.T) {
	vec := r.NewVector()
	orig := vec
	max := 40
	for i := 0; i < max; i++ {
		obj := r.EnsureObject(i)
		vec = vec.Cons(&obj)
	}

	if vec.Count() != uint32(max) {
		t.Fatalf("Expected %v items, got: %v", max, vec.Count())
	}

	if orig.Count() != 0 || orig.RootArray()[0] != nil {
		t.Fatalf("Mutated original")
	}

	val, ok := vec.GetNth(uint32(max - 1))
	if !ok {
		t.Fatalf("Unable to get item at idx %v", max-1)
	}

	if val.Value() != (max - 1) {
		t.Fatalf("Expected %v, Got %v", (max - 1), val.Value())
	}
}

func TestVectorConsFillTailAndRoot(t *testing.T) {
	vec := r.NewVector()
	max := 2000
	for i := 0; i < max; i++ {
		obj := r.EnsureObject(i)
		vec = vec.Cons(&obj)
	}

	for i := 0; i < max; i++ {
		val, ok := vec.GetNth(uint32(i))
		if !ok || val.Value() != i {
			t.Fatalf("Expected %v, got %v", i, val.Value())
		}
	}
}

func TestVectorSetNth(t *testing.T) {
	vec := r.NewVector()
	max := 64
	for i := 0; i < max; i++ {
		obj := r.EnsureObject(i)
		vec = vec.Cons(&obj)
	}
	orig1 := vec
	for i := 0; i < max; i++ {
		obj := r.EnsureObject(i)
		vec = vec.Cons(&obj)
	}
	orig2 := vec
	var ok bool

	obj := r.EnsureObject(999)
	vec, ok = vec.SetNth(33, &obj)

	val, ok := vec.GetNth(33)
	if !ok || val.Value() != 999 {
		t.Fatalf("Expected %v, got %v", 999, val.Value())
	}

	if !ok {
		t.Fatalf("Unable to SetNth")
	}

	val, ok = orig1.GetNth(33)
	if !ok || val.Value() == 999 {
		t.Fatalf("Expected not %v, got %v", 999, val.Value())
	}

	val, ok = orig2.GetNth(33)
	if !ok || val.Value() == 999 {
		t.Fatalf("Expected not %v, got %v", 999, val.Value())
	}
}
