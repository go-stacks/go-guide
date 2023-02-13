package sync

import "fmt"

type ArrayList[T any] struct {
	vals []T
}

type List[T any] interface {
	Get(index int) (T, error)
	Append(val T) error
	Add(index int, val T) error
	Set(index int, val T) error
	DeleteAt(index int) error
	Len() int
	Cap() int
	Range(fn func(index int, val T) error) error
	AsSlice() []T
}

func NewArrayListOf[T any](ts []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: ts,
	}
}

var _ List[int] = &ArrayList[int]{}

// Append implements List
func (a *ArrayList[T]) Add(index int, val T) error {
	if index < 0 || index > len(a.vals) {
		return newErrIndexOutOfRange(len(a.vals), index)
	}
	a.vals = append(a.vals, val)
	copy(a.vals[index+1:], a.vals[index:])
	a.vals[index] = val
	return nil
}

// Append implements List
func (a *ArrayList[T]) Append(val T) error {
	a.vals = append(a.vals, val)
	return nil
}

// AsSlice implements List
func (a *ArrayList[T]) AsSlice() []T {
	slice := make([]T, len(a.vals))
	copy(slice, a.vals)
	return slice
}

// Cap implements List
func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
}

// DeleteAt implements List
func (a *ArrayList[T]) DeleteAt(index int) error {
	if index >= len(a.vals) || index < 0 {
		return fmt.Errorf("index out of bounds")
	}
	a.vals = append(a.vals[:index], a.vals[index+1:]...)
	return nil
}

// Get implements List
func (a *ArrayList[T]) Get(index int) (T, error) {
	if index >= len(a.vals) || index < 0 {
		var t T
		return t, fmt.Errorf("index out of bounds")
	}
	return a.vals[index], nil
}

// Len implements List
func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

// Range implements List
func (a *ArrayList[T]) Range(fn func(index int, val T) error) error {
	for key, value := range a.vals {
		e := fn(key, value)
		if e != nil {
			return e
		}
	}

	return nil
}

// Set implements List
func (a *ArrayList[T]) Set(index int, val T) error {
	if index >= len(a.vals) || index < 0 {
		return fmt.Errorf("index out of bounds")
	}
	a.vals[index] = val
	return nil
}
