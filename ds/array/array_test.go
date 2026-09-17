package array

import "testing"

func TestAppendGet(t *testing.T) {
	a := New[int](4)
	if !a.Append(10) || !a.Append(20) {
		t.Fatal("append")
	}
	if a.Len() != 2 {
		t.Fatalf("len=%d", a.Len())
	}
	v, ok := a.Get(1)
	if !ok || v != 20 {
		t.Fatalf("get=%d ok=%v", v, ok)
	}
}

func TestInsertDelete(t *testing.T) {
	b := New[int](8)
	for _, x := range []int{1, 3, 4} {
		b.Append(x)
	}
	if !b.Insert(1, 2) {
		t.Fatal("insert")
	}
	if got := b.Slice(); !equal(got, []int{1, 2, 3, 4}) {
		t.Fatalf("after insert %v", got)
	}
	if !b.Delete(2) {
		t.Fatal("delete")
	}
	if got := b.Slice(); !equal(got, []int{1, 2, 4}) {
		t.Fatalf("after delete %v", got)
	}
}

func TestFullAppendFails(t *testing.T) {
	a := New[int](1)
	if !a.Append(1) {
		t.Fatal("first")
	}
	if a.Append(2) {
		t.Fatal("expected full")
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
