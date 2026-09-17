// Package array is a fixed-capacity contiguous list.
// Length is tracked separately from capacity.
package array

type Array[T any] struct {
	data []T
	n    int
}

func New[T any](cap int) *Array[T] {
	if cap < 0 {
		cap = 0
	}
	return &Array[T]{data: make([]T, cap), n: 0}
}

func From[T any](vals []T) *Array[T] {
	d := make([]T, len(vals))
	copy(d, vals)
	return &Array[T]{data: d, n: len(vals)}
}

func (a *Array[T]) Len() int { return a.n }
func (a *Array[T]) Cap() int { return len(a.data) }

func (a *Array[T]) Get(i int) (T, bool) {
	var zero T
	if i < 0 || i >= a.n {
		return zero, false
	}
	return a.data[i], true
}

func (a *Array[T]) Set(i int, v T) bool {
	if i < 0 || i >= a.n {
		return false
	}
	a.data[i] = v
	return true
}

// Append fails when capacity is full. Growing belongs in ds/dynarray.
func (a *Array[T]) Append(v T) bool {
	if a.n >= len(a.data) {
		return false
	}
	a.data[a.n] = v
	a.n++
	return true
}

func (a *Array[T]) Insert(i int, v T) bool {
	if i < 0 || i > a.n || a.n >= len(a.data) {
		return false
	}
	copy(a.data[i+1:a.n+1], a.data[i:a.n])
	a.data[i] = v
	a.n++
	return true
}

func (a *Array[T]) Delete(i int) bool {
	if i < 0 || i >= a.n {
		return false
	}
	copy(a.data[i:a.n-1], a.data[i+1:a.n])
	var zero T
	a.data[a.n-1] = zero
	a.n--
	return true
}

func (a *Array[T]) Slice() []T {
	out := make([]T, a.n)
	copy(out, a.data[:a.n])
	return out
}
