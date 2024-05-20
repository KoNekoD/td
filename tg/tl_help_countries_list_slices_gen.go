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

// HelpCountriesListClassArray is adapter for slice of HelpCountriesListClass.
type HelpCountriesListClassArray []HelpCountriesListClass

// Sort sorts slice of HelpCountriesListClass.
func (s HelpCountriesListClassArray) Sort(less func(a, b HelpCountriesListClass) bool) HelpCountriesListClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpCountriesListClass.
func (s HelpCountriesListClassArray) SortStable(less func(a, b HelpCountriesListClass) bool) HelpCountriesListClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpCountriesListClass.
func (s HelpCountriesListClassArray) Retain(keep func(x HelpCountriesListClass) bool) HelpCountriesListClassArray {
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
func (s HelpCountriesListClassArray) First() (v HelpCountriesListClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpCountriesListClassArray) Last() (v HelpCountriesListClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpCountriesListClassArray) PopFirst() (v HelpCountriesListClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpCountriesListClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpCountriesListClassArray) Pop() (v HelpCountriesListClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsHelpCountriesList returns copy with only HelpCountriesList constructors.
func (s HelpCountriesListClassArray) AsHelpCountriesList() (to HelpCountriesListArray) {
	for _, elem := range s {
		value, ok := elem.(*HelpCountriesList)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s HelpCountriesListClassArray) AppendOnlyModified(to []*HelpCountriesList) []*HelpCountriesList {
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
func (s HelpCountriesListClassArray) AsModified() (to []*HelpCountriesList) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s HelpCountriesListClassArray) FirstAsModified() (v *HelpCountriesList, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s HelpCountriesListClassArray) LastAsModified() (v *HelpCountriesList, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *HelpCountriesListClassArray) PopFirstAsModified() (v *HelpCountriesList, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *HelpCountriesListClassArray) PopAsModified() (v *HelpCountriesList, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// HelpCountriesListArray is adapter for slice of HelpCountriesList.
type HelpCountriesListArray []HelpCountriesList

// Sort sorts slice of HelpCountriesList.
func (s HelpCountriesListArray) Sort(less func(a, b HelpCountriesList) bool) HelpCountriesListArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of HelpCountriesList.
func (s HelpCountriesListArray) SortStable(less func(a, b HelpCountriesList) bool) HelpCountriesListArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of HelpCountriesList.
func (s HelpCountriesListArray) Retain(keep func(x HelpCountriesList) bool) HelpCountriesListArray {
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
func (s HelpCountriesListArray) First() (v HelpCountriesList, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s HelpCountriesListArray) Last() (v HelpCountriesList, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *HelpCountriesListArray) PopFirst() (v HelpCountriesList, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero HelpCountriesList
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *HelpCountriesListArray) Pop() (v HelpCountriesList, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
