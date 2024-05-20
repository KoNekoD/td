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

// MessagesFeaturedStickersClassArray is adapter for slice of MessagesFeaturedStickersClass.
type MessagesFeaturedStickersClassArray []MessagesFeaturedStickersClass

// Sort sorts slice of MessagesFeaturedStickersClass.
func (s MessagesFeaturedStickersClassArray) Sort(less func(a, b MessagesFeaturedStickersClass) bool) MessagesFeaturedStickersClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFeaturedStickersClass.
func (s MessagesFeaturedStickersClassArray) SortStable(less func(a, b MessagesFeaturedStickersClass) bool) MessagesFeaturedStickersClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFeaturedStickersClass.
func (s MessagesFeaturedStickersClassArray) Retain(keep func(x MessagesFeaturedStickersClass) bool) MessagesFeaturedStickersClassArray {
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
func (s MessagesFeaturedStickersClassArray) First() (v MessagesFeaturedStickersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFeaturedStickersClassArray) Last() (v MessagesFeaturedStickersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersClassArray) PopFirst() (v MessagesFeaturedStickersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFeaturedStickersClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersClassArray) Pop() (v MessagesFeaturedStickersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesFeaturedStickersNotModified returns copy with only MessagesFeaturedStickersNotModified constructors.
func (s MessagesFeaturedStickersClassArray) AsMessagesFeaturedStickersNotModified() (to MessagesFeaturedStickersNotModifiedArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesFeaturedStickersNotModified)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessagesFeaturedStickers returns copy with only MessagesFeaturedStickers constructors.
func (s MessagesFeaturedStickersClassArray) AsMessagesFeaturedStickers() (to MessagesFeaturedStickersArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesFeaturedStickers)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesFeaturedStickersClassArray) AppendOnlyModified(to []*MessagesFeaturedStickers) []*MessagesFeaturedStickers {
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
func (s MessagesFeaturedStickersClassArray) AsModified() (to []*MessagesFeaturedStickers) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesFeaturedStickersClassArray) FirstAsModified() (v *MessagesFeaturedStickers, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesFeaturedStickersClassArray) LastAsModified() (v *MessagesFeaturedStickers, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesFeaturedStickersClassArray) PopFirstAsModified() (v *MessagesFeaturedStickers, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesFeaturedStickersClassArray) PopAsModified() (v *MessagesFeaturedStickers, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesFeaturedStickersNotModifiedArray is adapter for slice of MessagesFeaturedStickersNotModified.
type MessagesFeaturedStickersNotModifiedArray []MessagesFeaturedStickersNotModified

// Sort sorts slice of MessagesFeaturedStickersNotModified.
func (s MessagesFeaturedStickersNotModifiedArray) Sort(less func(a, b MessagesFeaturedStickersNotModified) bool) MessagesFeaturedStickersNotModifiedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFeaturedStickersNotModified.
func (s MessagesFeaturedStickersNotModifiedArray) SortStable(less func(a, b MessagesFeaturedStickersNotModified) bool) MessagesFeaturedStickersNotModifiedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFeaturedStickersNotModified.
func (s MessagesFeaturedStickersNotModifiedArray) Retain(keep func(x MessagesFeaturedStickersNotModified) bool) MessagesFeaturedStickersNotModifiedArray {
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
func (s MessagesFeaturedStickersNotModifiedArray) First() (v MessagesFeaturedStickersNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFeaturedStickersNotModifiedArray) Last() (v MessagesFeaturedStickersNotModified, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersNotModifiedArray) PopFirst() (v MessagesFeaturedStickersNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFeaturedStickersNotModified
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersNotModifiedArray) Pop() (v MessagesFeaturedStickersNotModified, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessagesFeaturedStickersArray is adapter for slice of MessagesFeaturedStickers.
type MessagesFeaturedStickersArray []MessagesFeaturedStickers

// Sort sorts slice of MessagesFeaturedStickers.
func (s MessagesFeaturedStickersArray) Sort(less func(a, b MessagesFeaturedStickers) bool) MessagesFeaturedStickersArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesFeaturedStickers.
func (s MessagesFeaturedStickersArray) SortStable(less func(a, b MessagesFeaturedStickers) bool) MessagesFeaturedStickersArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesFeaturedStickers.
func (s MessagesFeaturedStickersArray) Retain(keep func(x MessagesFeaturedStickers) bool) MessagesFeaturedStickersArray {
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
func (s MessagesFeaturedStickersArray) First() (v MessagesFeaturedStickers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesFeaturedStickersArray) Last() (v MessagesFeaturedStickers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersArray) PopFirst() (v MessagesFeaturedStickers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesFeaturedStickers
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesFeaturedStickersArray) Pop() (v MessagesFeaturedStickers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
