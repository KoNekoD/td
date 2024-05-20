//go:build !no_gotd_slices
// +build !no_gotd_slices

// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/KoNekoD/td/bin"
	"github.com/KoNekoD/td/tdjson"
	"github.com/KoNekoD/td/tdp"
	"github.com/KoNekoD/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
	_ = tdjson.Encoder{}
)

// SecureFileClassArray is adapter for slice of SecureFileClass.
type SecureFileClassArray []SecureFileClass

// Sort sorts slice of SecureFileClass.
func (s SecureFileClassArray) Sort(less func(a, b SecureFileClass) bool) SecureFileClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureFileClass.
func (s SecureFileClassArray) SortStable(less func(a, b SecureFileClass) bool) SecureFileClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureFileClass.
func (s SecureFileClassArray) Retain(keep func(x SecureFileClass) bool) SecureFileClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureFileClassArray) First() (v SecureFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureFileClassArray) Last() (v SecureFileClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureFileClassArray) PopFirst() (v SecureFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureFileClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureFileClassArray) Pop() (v SecureFileClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsSecureFile returns copy with only SecureFile constructors.
func (s SecureFileClassArray) AsSecureFile() (to SecureFileArray) {
	for _, elem := range s {
		value, ok := elem.(*SecureFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillNotEmptyMap fills only NotEmpty constructors to given map.
func (s SecureFileClassArray) FillNotEmptyMap(to map[int64]*SecureFile) {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// NotEmptyToMap collects only NotEmpty constructors to map.
func (s SecureFileClassArray) NotEmptyToMap() map[int64]*SecureFile {
	r := make(map[int64]*SecureFile, len(s))
	s.FillNotEmptyMap(r)
	return r
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s SecureFileClassArray) AppendOnlyNotEmpty(to []*SecureFile) []*SecureFile {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotEmpty returns copy with only NotEmpty constructors.
func (s SecureFileClassArray) AsNotEmpty() (to []*SecureFile) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s SecureFileClassArray) FirstAsNotEmpty() (v *SecureFile, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s SecureFileClassArray) LastAsNotEmpty() (v *SecureFile, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *SecureFileClassArray) PopFirstAsNotEmpty() (v *SecureFile, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *SecureFileClassArray) PopAsNotEmpty() (v *SecureFile, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// SecureFileArray is adapter for slice of SecureFile.
type SecureFileArray []SecureFile

// Sort sorts slice of SecureFile.
func (s SecureFileArray) Sort(less func(a, b SecureFile) bool) SecureFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SecureFile.
func (s SecureFileArray) SortStable(less func(a, b SecureFile) bool) SecureFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SecureFile.
func (s SecureFileArray) Retain(keep func(x SecureFile) bool) SecureFileArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SecureFileArray) First() (v SecureFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SecureFileArray) Last() (v SecureFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SecureFileArray) PopFirst() (v SecureFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SecureFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SecureFileArray) Pop() (v SecureFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of SecureFile by ID.
func (s SecureFileArray) SortByID() SecureFileArray {
	return s.Sort(func(a, b SecureFile) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of SecureFile by ID.
func (s SecureFileArray) SortStableByID() SecureFileArray {
	return s.SortStable(func(a, b SecureFile) bool {
		return a.GetID() < b.GetID()
	})
}

// SortByDate sorts slice of SecureFile by Date.
func (s SecureFileArray) SortByDate() SecureFileArray {
	return s.Sort(func(a, b SecureFile) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of SecureFile by Date.
func (s SecureFileArray) SortStableByDate() SecureFileArray {
	return s.SortStable(func(a, b SecureFile) bool {
		return a.GetDate() < b.GetDate()
	})
}

// FillMap fills constructors to given map.
func (s SecureFileArray) FillMap(to map[int64]SecureFile) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s SecureFileArray) ToMap() map[int64]SecureFile {
	r := make(map[int64]SecureFile, len(s))
	s.FillMap(r)
	return r
}
