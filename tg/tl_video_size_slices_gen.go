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

// VideoSizeClassArray is adapter for slice of VideoSizeClass.
type VideoSizeClassArray []VideoSizeClass

// Sort sorts slice of VideoSizeClass.
func (s VideoSizeClassArray) Sort(less func(a, b VideoSizeClass) bool) VideoSizeClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of VideoSizeClass.
func (s VideoSizeClassArray) SortStable(less func(a, b VideoSizeClass) bool) VideoSizeClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of VideoSizeClass.
func (s VideoSizeClassArray) Retain(keep func(x VideoSizeClass) bool) VideoSizeClassArray {
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
func (s VideoSizeClassArray) First() (v VideoSizeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s VideoSizeClassArray) Last() (v VideoSizeClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *VideoSizeClassArray) PopFirst() (v VideoSizeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero VideoSizeClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *VideoSizeClassArray) Pop() (v VideoSizeClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsVideoSize returns copy with only VideoSize constructors.
func (s VideoSizeClassArray) AsVideoSize() (to VideoSizeArray) {
	for _, elem := range s {
		value, ok := elem.(*VideoSize)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsVideoSizeEmojiMarkup returns copy with only VideoSizeEmojiMarkup constructors.
func (s VideoSizeClassArray) AsVideoSizeEmojiMarkup() (to VideoSizeEmojiMarkupArray) {
	for _, elem := range s {
		value, ok := elem.(*VideoSizeEmojiMarkup)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsVideoSizeStickerMarkup returns copy with only VideoSizeStickerMarkup constructors.
func (s VideoSizeClassArray) AsVideoSizeStickerMarkup() (to VideoSizeStickerMarkupArray) {
	for _, elem := range s {
		value, ok := elem.(*VideoSizeStickerMarkup)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// VideoSizeArray is adapter for slice of VideoSize.
type VideoSizeArray []VideoSize

// Sort sorts slice of VideoSize.
func (s VideoSizeArray) Sort(less func(a, b VideoSize) bool) VideoSizeArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of VideoSize.
func (s VideoSizeArray) SortStable(less func(a, b VideoSize) bool) VideoSizeArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of VideoSize.
func (s VideoSizeArray) Retain(keep func(x VideoSize) bool) VideoSizeArray {
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
func (s VideoSizeArray) First() (v VideoSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s VideoSizeArray) Last() (v VideoSize, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *VideoSizeArray) PopFirst() (v VideoSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero VideoSize
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *VideoSizeArray) Pop() (v VideoSize, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// VideoSizeEmojiMarkupArray is adapter for slice of VideoSizeEmojiMarkup.
type VideoSizeEmojiMarkupArray []VideoSizeEmojiMarkup

// Sort sorts slice of VideoSizeEmojiMarkup.
func (s VideoSizeEmojiMarkupArray) Sort(less func(a, b VideoSizeEmojiMarkup) bool) VideoSizeEmojiMarkupArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of VideoSizeEmojiMarkup.
func (s VideoSizeEmojiMarkupArray) SortStable(less func(a, b VideoSizeEmojiMarkup) bool) VideoSizeEmojiMarkupArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of VideoSizeEmojiMarkup.
func (s VideoSizeEmojiMarkupArray) Retain(keep func(x VideoSizeEmojiMarkup) bool) VideoSizeEmojiMarkupArray {
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
func (s VideoSizeEmojiMarkupArray) First() (v VideoSizeEmojiMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s VideoSizeEmojiMarkupArray) Last() (v VideoSizeEmojiMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *VideoSizeEmojiMarkupArray) PopFirst() (v VideoSizeEmojiMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero VideoSizeEmojiMarkup
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *VideoSizeEmojiMarkupArray) Pop() (v VideoSizeEmojiMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// VideoSizeStickerMarkupArray is adapter for slice of VideoSizeStickerMarkup.
type VideoSizeStickerMarkupArray []VideoSizeStickerMarkup

// Sort sorts slice of VideoSizeStickerMarkup.
func (s VideoSizeStickerMarkupArray) Sort(less func(a, b VideoSizeStickerMarkup) bool) VideoSizeStickerMarkupArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of VideoSizeStickerMarkup.
func (s VideoSizeStickerMarkupArray) SortStable(less func(a, b VideoSizeStickerMarkup) bool) VideoSizeStickerMarkupArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of VideoSizeStickerMarkup.
func (s VideoSizeStickerMarkupArray) Retain(keep func(x VideoSizeStickerMarkup) bool) VideoSizeStickerMarkupArray {
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
func (s VideoSizeStickerMarkupArray) First() (v VideoSizeStickerMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s VideoSizeStickerMarkupArray) Last() (v VideoSizeStickerMarkup, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *VideoSizeStickerMarkupArray) PopFirst() (v VideoSizeStickerMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero VideoSizeStickerMarkup
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *VideoSizeStickerMarkupArray) Pop() (v VideoSizeStickerMarkup, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
