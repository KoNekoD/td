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

// MessageExtendedMediaClassArray is adapter for slice of MessageExtendedMediaClass.
type MessageExtendedMediaClassArray []MessageExtendedMediaClass

// Sort sorts slice of MessageExtendedMediaClass.
func (s MessageExtendedMediaClassArray) Sort(less func(a, b MessageExtendedMediaClass) bool) MessageExtendedMediaClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageExtendedMediaClass.
func (s MessageExtendedMediaClassArray) SortStable(less func(a, b MessageExtendedMediaClass) bool) MessageExtendedMediaClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageExtendedMediaClass.
func (s MessageExtendedMediaClassArray) Retain(keep func(x MessageExtendedMediaClass) bool) MessageExtendedMediaClassArray {
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
func (s MessageExtendedMediaClassArray) First() (v MessageExtendedMediaClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageExtendedMediaClassArray) Last() (v MessageExtendedMediaClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageExtendedMediaClassArray) PopFirst() (v MessageExtendedMediaClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageExtendedMediaClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageExtendedMediaClassArray) Pop() (v MessageExtendedMediaClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessageExtendedMediaPreview returns copy with only MessageExtendedMediaPreview constructors.
func (s MessageExtendedMediaClassArray) AsMessageExtendedMediaPreview() (to MessageExtendedMediaPreviewArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageExtendedMediaPreview)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsMessageExtendedMedia returns copy with only MessageExtendedMedia constructors.
func (s MessageExtendedMediaClassArray) AsMessageExtendedMedia() (to MessageExtendedMediaArray) {
	for _, elem := range s {
		value, ok := elem.(*MessageExtendedMedia)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// MessageExtendedMediaPreviewArray is adapter for slice of MessageExtendedMediaPreview.
type MessageExtendedMediaPreviewArray []MessageExtendedMediaPreview

// Sort sorts slice of MessageExtendedMediaPreview.
func (s MessageExtendedMediaPreviewArray) Sort(less func(a, b MessageExtendedMediaPreview) bool) MessageExtendedMediaPreviewArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageExtendedMediaPreview.
func (s MessageExtendedMediaPreviewArray) SortStable(less func(a, b MessageExtendedMediaPreview) bool) MessageExtendedMediaPreviewArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageExtendedMediaPreview.
func (s MessageExtendedMediaPreviewArray) Retain(keep func(x MessageExtendedMediaPreview) bool) MessageExtendedMediaPreviewArray {
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
func (s MessageExtendedMediaPreviewArray) First() (v MessageExtendedMediaPreview, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageExtendedMediaPreviewArray) Last() (v MessageExtendedMediaPreview, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageExtendedMediaPreviewArray) PopFirst() (v MessageExtendedMediaPreview, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageExtendedMediaPreview
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageExtendedMediaPreviewArray) Pop() (v MessageExtendedMediaPreview, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// MessageExtendedMediaArray is adapter for slice of MessageExtendedMedia.
type MessageExtendedMediaArray []MessageExtendedMedia

// Sort sorts slice of MessageExtendedMedia.
func (s MessageExtendedMediaArray) Sort(less func(a, b MessageExtendedMedia) bool) MessageExtendedMediaArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessageExtendedMedia.
func (s MessageExtendedMediaArray) SortStable(less func(a, b MessageExtendedMedia) bool) MessageExtendedMediaArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessageExtendedMedia.
func (s MessageExtendedMediaArray) Retain(keep func(x MessageExtendedMedia) bool) MessageExtendedMediaArray {
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
func (s MessageExtendedMediaArray) First() (v MessageExtendedMedia, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessageExtendedMediaArray) Last() (v MessageExtendedMedia, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessageExtendedMediaArray) PopFirst() (v MessageExtendedMedia, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessageExtendedMedia
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessageExtendedMediaArray) Pop() (v MessageExtendedMedia, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
