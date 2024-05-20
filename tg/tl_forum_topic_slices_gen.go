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

// ForumTopicClassArray is adapter for slice of ForumTopicClass.
type ForumTopicClassArray []ForumTopicClass

// Sort sorts slice of ForumTopicClass.
func (s ForumTopicClassArray) Sort(less func(a, b ForumTopicClass) bool) ForumTopicClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ForumTopicClass.
func (s ForumTopicClassArray) SortStable(less func(a, b ForumTopicClass) bool) ForumTopicClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ForumTopicClass.
func (s ForumTopicClassArray) Retain(keep func(x ForumTopicClass) bool) ForumTopicClassArray {
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
func (s ForumTopicClassArray) First() (v ForumTopicClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ForumTopicClassArray) Last() (v ForumTopicClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ForumTopicClassArray) PopFirst() (v ForumTopicClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ForumTopicClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ForumTopicClassArray) Pop() (v ForumTopicClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of ForumTopicClass by ID.
func (s ForumTopicClassArray) SortByID() ForumTopicClassArray {
	return s.Sort(func(a, b ForumTopicClass) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of ForumTopicClass by ID.
func (s ForumTopicClassArray) SortStableByID() ForumTopicClassArray {
	return s.SortStable(func(a, b ForumTopicClass) bool {
		return a.GetID() < b.GetID()
	})
}

// FillForumTopicDeletedMap fills only ForumTopicDeleted constructors to given map.
func (s ForumTopicClassArray) FillForumTopicDeletedMap(to map[int]*ForumTopicDeleted) {
	for _, elem := range s {
		value, ok := elem.(*ForumTopicDeleted)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// ForumTopicDeletedToMap collects only ForumTopicDeleted constructors to map.
func (s ForumTopicClassArray) ForumTopicDeletedToMap() map[int]*ForumTopicDeleted {
	r := make(map[int]*ForumTopicDeleted, len(s))
	s.FillForumTopicDeletedMap(r)
	return r
}

// AsForumTopicDeleted returns copy with only ForumTopicDeleted constructors.
func (s ForumTopicClassArray) AsForumTopicDeleted() (to ForumTopicDeletedArray) {
	for _, elem := range s {
		value, ok := elem.(*ForumTopicDeleted)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillForumTopicMap fills only ForumTopic constructors to given map.
func (s ForumTopicClassArray) FillForumTopicMap(to map[int]*ForumTopic) {
	for _, elem := range s {
		value, ok := elem.(*ForumTopic)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// ForumTopicToMap collects only ForumTopic constructors to map.
func (s ForumTopicClassArray) ForumTopicToMap() map[int]*ForumTopic {
	r := make(map[int]*ForumTopic, len(s))
	s.FillForumTopicMap(r)
	return r
}

// AsForumTopic returns copy with only ForumTopic constructors.
func (s ForumTopicClassArray) AsForumTopic() (to ForumTopicArray) {
	for _, elem := range s {
		value, ok := elem.(*ForumTopic)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// ForumTopicDeletedArray is adapter for slice of ForumTopicDeleted.
type ForumTopicDeletedArray []ForumTopicDeleted

// Sort sorts slice of ForumTopicDeleted.
func (s ForumTopicDeletedArray) Sort(less func(a, b ForumTopicDeleted) bool) ForumTopicDeletedArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ForumTopicDeleted.
func (s ForumTopicDeletedArray) SortStable(less func(a, b ForumTopicDeleted) bool) ForumTopicDeletedArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ForumTopicDeleted.
func (s ForumTopicDeletedArray) Retain(keep func(x ForumTopicDeleted) bool) ForumTopicDeletedArray {
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
func (s ForumTopicDeletedArray) First() (v ForumTopicDeleted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ForumTopicDeletedArray) Last() (v ForumTopicDeleted, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ForumTopicDeletedArray) PopFirst() (v ForumTopicDeleted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ForumTopicDeleted
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ForumTopicDeletedArray) Pop() (v ForumTopicDeleted, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of ForumTopicDeleted by ID.
func (s ForumTopicDeletedArray) SortByID() ForumTopicDeletedArray {
	return s.Sort(func(a, b ForumTopicDeleted) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of ForumTopicDeleted by ID.
func (s ForumTopicDeletedArray) SortStableByID() ForumTopicDeletedArray {
	return s.SortStable(func(a, b ForumTopicDeleted) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s ForumTopicDeletedArray) FillMap(to map[int]ForumTopicDeleted) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s ForumTopicDeletedArray) ToMap() map[int]ForumTopicDeleted {
	r := make(map[int]ForumTopicDeleted, len(s))
	s.FillMap(r)
	return r
}

// ForumTopicArray is adapter for slice of ForumTopic.
type ForumTopicArray []ForumTopic

// Sort sorts slice of ForumTopic.
func (s ForumTopicArray) Sort(less func(a, b ForumTopic) bool) ForumTopicArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ForumTopic.
func (s ForumTopicArray) SortStable(less func(a, b ForumTopic) bool) ForumTopicArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ForumTopic.
func (s ForumTopicArray) Retain(keep func(x ForumTopic) bool) ForumTopicArray {
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
func (s ForumTopicArray) First() (v ForumTopic, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ForumTopicArray) Last() (v ForumTopic, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ForumTopicArray) PopFirst() (v ForumTopic, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ForumTopic
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ForumTopicArray) Pop() (v ForumTopic, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of ForumTopic by ID.
func (s ForumTopicArray) SortByID() ForumTopicArray {
	return s.Sort(func(a, b ForumTopic) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of ForumTopic by ID.
func (s ForumTopicArray) SortStableByID() ForumTopicArray {
	return s.SortStable(func(a, b ForumTopic) bool {
		return a.GetID() < b.GetID()
	})
}

// SortByDate sorts slice of ForumTopic by Date.
func (s ForumTopicArray) SortByDate() ForumTopicArray {
	return s.Sort(func(a, b ForumTopic) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of ForumTopic by Date.
func (s ForumTopicArray) SortStableByDate() ForumTopicArray {
	return s.SortStable(func(a, b ForumTopic) bool {
		return a.GetDate() < b.GetDate()
	})
}

// FillMap fills constructors to given map.
func (s ForumTopicArray) FillMap(to map[int]ForumTopic) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s ForumTopicArray) ToMap() map[int]ForumTopic {
	r := make(map[int]ForumTopic, len(s))
	s.FillMap(r)
	return r
}
