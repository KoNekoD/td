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

// StoriesAllStoriesClassArray is adapter for slice of StoriesAllStoriesClass.
type StoriesAllStoriesClassArray []StoriesAllStoriesClass

// Sort sorts slice of StoriesAllStoriesClass.
func (s StoriesAllStoriesClassArray) Sort(less func(a, b StoriesAllStoriesClass) bool) StoriesAllStoriesClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StoriesAllStoriesClass.
func (s StoriesAllStoriesClassArray) SortStable(less func(a, b StoriesAllStoriesClass) bool) StoriesAllStoriesClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StoriesAllStoriesClass.
func (s StoriesAllStoriesClassArray) Retain(keep func(x StoriesAllStoriesClass) bool) StoriesAllStoriesClassArray {
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
func (s StoriesAllStoriesClassArray) First() (v StoriesAllStoriesClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StoriesAllStoriesClassArray) Last() (v StoriesAllStoriesClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StoriesAllStoriesClassArray) PopFirst() (v StoriesAllStoriesClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StoriesAllStoriesClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StoriesAllStoriesClassArray) Pop() (v StoriesAllStoriesClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsStoriesAllStoriesNotModified returns copy with only StoriesAllStoriesNotModified constructors.
func (s StoriesAllStoriesClassArray) AsStoriesAllStoriesNotModified() (to StoriesAllStoriesNotModifiedArray) {
	for _, elem := range s {
		value, ok := elem.(*StoriesAllStoriesNotModified)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsStoriesAllStories returns copy with only StoriesAllStories constructors.
func (s StoriesAllStoriesClassArray) AsStoriesAllStories() (to StoriesAllStoriesArray) {
	for _, elem := range s {
		value, ok := elem.(*StoriesAllStories)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s StoriesAllStoriesClassArray) AppendOnlyModified(to []*StoriesAllStories) []*StoriesAllStories {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s StoriesAllStoriesClassArray) AsModified() (to []*StoriesAllStories) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s StoriesAllStoriesClassArray) FirstAsModified() (v *StoriesAllStories, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s StoriesAllStoriesClassArray) LastAsModified() (v *StoriesAllStories, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *StoriesAllStoriesClassArray) PopFirstAsModified() (v *StoriesAllStories, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *StoriesAllStoriesClassArray) PopAsModified() (v *StoriesAllStories, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// StoriesAllStoriesNotModifiedArray is adapter for slice of StoriesAllStoriesNotModified.
type StoriesAllStoriesNotModifiedArray []StoriesAllStoriesNotModified

// Sort sorts slice of StoriesAllStoriesNotModified.
func (s StoriesAllStoriesNotModifiedArray) Sort(less func(a, b StoriesAllStoriesNotModified) bool) StoriesAllStoriesNotModifiedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StoriesAllStoriesNotModified.
func (s StoriesAllStoriesNotModifiedArray) SortStable(less func(a, b StoriesAllStoriesNotModified) bool) StoriesAllStoriesNotModifiedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StoriesAllStoriesNotModified.
func (s StoriesAllStoriesNotModifiedArray) Retain(keep func(x StoriesAllStoriesNotModified) bool) StoriesAllStoriesNotModifiedArray {
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
func (s StoriesAllStoriesNotModifiedArray) First() (v StoriesAllStoriesNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StoriesAllStoriesNotModifiedArray) Last() (v StoriesAllStoriesNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StoriesAllStoriesNotModifiedArray) PopFirst() (v StoriesAllStoriesNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StoriesAllStoriesNotModified
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StoriesAllStoriesNotModifiedArray) Pop() (v StoriesAllStoriesNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// StoriesAllStoriesArray is adapter for slice of StoriesAllStories.
type StoriesAllStoriesArray []StoriesAllStories

// Sort sorts slice of StoriesAllStories.
func (s StoriesAllStoriesArray) Sort(less func(a, b StoriesAllStories) bool) StoriesAllStoriesArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StoriesAllStories.
func (s StoriesAllStoriesArray) SortStable(less func(a, b StoriesAllStories) bool) StoriesAllStoriesArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StoriesAllStories.
func (s StoriesAllStoriesArray) Retain(keep func(x StoriesAllStories) bool) StoriesAllStoriesArray {
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
func (s StoriesAllStoriesArray) First() (v StoriesAllStories, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StoriesAllStoriesArray) Last() (v StoriesAllStories, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StoriesAllStoriesArray) PopFirst() (v StoriesAllStories, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StoriesAllStories
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StoriesAllStoriesArray) Pop() (v StoriesAllStories, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
