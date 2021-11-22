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

	"github.com/gotd/td/bin"
	"github.com/gotd/td/jsontd"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
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
	_ = jsontd.Encoder{}
)

// InputGameClassArray is adapter for slice of InputGameClass.
type InputGameClassArray []InputGameClass

// Sort sorts slice of InputGameClass.
func (s InputGameClassArray) Sort(less func(a, b InputGameClass) bool) InputGameClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputGameClass.
func (s InputGameClassArray) SortStable(less func(a, b InputGameClass) bool) InputGameClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputGameClass.
func (s InputGameClassArray) Retain(keep func(x InputGameClass) bool) InputGameClassArray {
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
func (s InputGameClassArray) First() (v InputGameClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputGameClassArray) Last() (v InputGameClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputGameClassArray) PopFirst() (v InputGameClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputGameClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputGameClassArray) Pop() (v InputGameClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputGameID returns copy with only InputGameID constructors.
func (s InputGameClassArray) AsInputGameID() (to InputGameIDArray) {
	for _, elem := range s {
		value, ok := elem.(*InputGameID)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputGameShortName returns copy with only InputGameShortName constructors.
func (s InputGameClassArray) AsInputGameShortName() (to InputGameShortNameArray) {
	for _, elem := range s {
		value, ok := elem.(*InputGameShortName)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputGameIDArray is adapter for slice of InputGameID.
type InputGameIDArray []InputGameID

// Sort sorts slice of InputGameID.
func (s InputGameIDArray) Sort(less func(a, b InputGameID) bool) InputGameIDArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputGameID.
func (s InputGameIDArray) SortStable(less func(a, b InputGameID) bool) InputGameIDArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputGameID.
func (s InputGameIDArray) Retain(keep func(x InputGameID) bool) InputGameIDArray {
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
func (s InputGameIDArray) First() (v InputGameID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputGameIDArray) Last() (v InputGameID, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputGameIDArray) PopFirst() (v InputGameID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputGameID
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputGameIDArray) Pop() (v InputGameID, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of InputGameID by ID.
func (s InputGameIDArray) SortByID() InputGameIDArray {
	return s.Sort(func(a, b InputGameID) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of InputGameID by ID.
func (s InputGameIDArray) SortStableByID() InputGameIDArray {
	return s.SortStable(func(a, b InputGameID) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s InputGameIDArray) FillMap(to map[int64]InputGameID) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s InputGameIDArray) ToMap() map[int64]InputGameID {
	r := make(map[int64]InputGameID, len(s))
	s.FillMap(r)
	return r
}

// InputGameShortNameArray is adapter for slice of InputGameShortName.
type InputGameShortNameArray []InputGameShortName

// Sort sorts slice of InputGameShortName.
func (s InputGameShortNameArray) Sort(less func(a, b InputGameShortName) bool) InputGameShortNameArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputGameShortName.
func (s InputGameShortNameArray) SortStable(less func(a, b InputGameShortName) bool) InputGameShortNameArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputGameShortName.
func (s InputGameShortNameArray) Retain(keep func(x InputGameShortName) bool) InputGameShortNameArray {
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
func (s InputGameShortNameArray) First() (v InputGameShortName, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputGameShortNameArray) Last() (v InputGameShortName, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputGameShortNameArray) PopFirst() (v InputGameShortName, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputGameShortName
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputGameShortNameArray) Pop() (v InputGameShortName, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
