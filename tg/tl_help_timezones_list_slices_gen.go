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

// HelpTimezonesListClassArray is adapter for slice of HelpTimezonesListClass.
type HelpTimezonesListClassArray []HelpTimezonesListClass

// Sort sorts slice of HelpTimezonesListClass.
func (s HelpTimezonesListClassArray) Sort(less func(a, b HelpTimezonesListClass) bool) HelpTimezonesListClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpTimezonesListClass.
func (s HelpTimezonesListClassArray) SortStable(less func(a, b HelpTimezonesListClass) bool) HelpTimezonesListClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpTimezonesListClass.
func (s HelpTimezonesListClassArray) Retain(keep func(x HelpTimezonesListClass) bool) HelpTimezonesListClassArray {
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
func (s HelpTimezonesListClassArray) First() (v HelpTimezonesListClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpTimezonesListClassArray) Last() (v HelpTimezonesListClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpTimezonesListClassArray) PopFirst() (v HelpTimezonesListClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpTimezonesListClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpTimezonesListClassArray) Pop() (v HelpTimezonesListClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpTimezonesList returns copy with only HelpTimezonesList constructors.
func (s HelpTimezonesListClassArray) AsHelpTimezonesList() (to HelpTimezonesListArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpTimezonesList)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s HelpTimezonesListClassArray) AppendOnlyModified(to []*HelpTimezonesList) []*HelpTimezonesList {
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
func (s HelpTimezonesListClassArray) AsModified() (to []*HelpTimezonesList) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s HelpTimezonesListClassArray) FirstAsModified() (v *HelpTimezonesList, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s HelpTimezonesListClassArray) LastAsModified() (v *HelpTimezonesList, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *HelpTimezonesListClassArray) PopFirstAsModified() (v *HelpTimezonesList, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *HelpTimezonesListClassArray) PopAsModified() (v *HelpTimezonesList, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// HelpTimezonesListArray is adapter for slice of HelpTimezonesList.
type HelpTimezonesListArray []HelpTimezonesList

// Sort sorts slice of HelpTimezonesList.
func (s HelpTimezonesListArray) Sort(less func(a, b HelpTimezonesList) bool) HelpTimezonesListArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpTimezonesList.
func (s HelpTimezonesListArray) SortStable(less func(a, b HelpTimezonesList) bool) HelpTimezonesListArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpTimezonesList.
func (s HelpTimezonesListArray) Retain(keep func(x HelpTimezonesList) bool) HelpTimezonesListArray {
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
func (s HelpTimezonesListArray) First() (v HelpTimezonesList, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpTimezonesListArray) Last() (v HelpTimezonesList, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpTimezonesListArray) PopFirst() (v HelpTimezonesList, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpTimezonesList
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpTimezonesListArray) Pop() (v HelpTimezonesList, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
