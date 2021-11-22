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

// StatsGraphClassArray is adapter for slice of StatsGraphClass.
type StatsGraphClassArray []StatsGraphClass

// Sort sorts slice of StatsGraphClass.
func (s StatsGraphClassArray) Sort(less func(a, b StatsGraphClass) bool) StatsGraphClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StatsGraphClass.
func (s StatsGraphClassArray) SortStable(less func(a, b StatsGraphClass) bool) StatsGraphClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StatsGraphClass.
func (s StatsGraphClassArray) Retain(keep func(x StatsGraphClass) bool) StatsGraphClassArray {
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
func (s StatsGraphClassArray) First() (v StatsGraphClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StatsGraphClassArray) Last() (v StatsGraphClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StatsGraphClassArray) PopFirst() (v StatsGraphClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StatsGraphClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StatsGraphClassArray) Pop() (v StatsGraphClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsStatsGraphAsync returns copy with only StatsGraphAsync constructors.
func (s StatsGraphClassArray) AsStatsGraphAsync() (to StatsGraphAsyncArray) {
	for _, elem := range s {
		value, ok := elem.(*StatsGraphAsync)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsStatsGraphError returns copy with only StatsGraphError constructors.
func (s StatsGraphClassArray) AsStatsGraphError() (to StatsGraphErrorArray) {
	for _, elem := range s {
		value, ok := elem.(*StatsGraphError)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsStatsGraph returns copy with only StatsGraph constructors.
func (s StatsGraphClassArray) AsStatsGraph() (to StatsGraphArray) {
	for _, elem := range s {
		value, ok := elem.(*StatsGraph)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// StatsGraphAsyncArray is adapter for slice of StatsGraphAsync.
type StatsGraphAsyncArray []StatsGraphAsync

// Sort sorts slice of StatsGraphAsync.
func (s StatsGraphAsyncArray) Sort(less func(a, b StatsGraphAsync) bool) StatsGraphAsyncArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StatsGraphAsync.
func (s StatsGraphAsyncArray) SortStable(less func(a, b StatsGraphAsync) bool) StatsGraphAsyncArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StatsGraphAsync.
func (s StatsGraphAsyncArray) Retain(keep func(x StatsGraphAsync) bool) StatsGraphAsyncArray {
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
func (s StatsGraphAsyncArray) First() (v StatsGraphAsync, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StatsGraphAsyncArray) Last() (v StatsGraphAsync, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StatsGraphAsyncArray) PopFirst() (v StatsGraphAsync, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StatsGraphAsync
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StatsGraphAsyncArray) Pop() (v StatsGraphAsync, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// StatsGraphErrorArray is adapter for slice of StatsGraphError.
type StatsGraphErrorArray []StatsGraphError

// Sort sorts slice of StatsGraphError.
func (s StatsGraphErrorArray) Sort(less func(a, b StatsGraphError) bool) StatsGraphErrorArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StatsGraphError.
func (s StatsGraphErrorArray) SortStable(less func(a, b StatsGraphError) bool) StatsGraphErrorArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StatsGraphError.
func (s StatsGraphErrorArray) Retain(keep func(x StatsGraphError) bool) StatsGraphErrorArray {
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
func (s StatsGraphErrorArray) First() (v StatsGraphError, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StatsGraphErrorArray) Last() (v StatsGraphError, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StatsGraphErrorArray) PopFirst() (v StatsGraphError, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StatsGraphError
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StatsGraphErrorArray) Pop() (v StatsGraphError, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// StatsGraphArray is adapter for slice of StatsGraph.
type StatsGraphArray []StatsGraph

// Sort sorts slice of StatsGraph.
func (s StatsGraphArray) Sort(less func(a, b StatsGraph) bool) StatsGraphArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of StatsGraph.
func (s StatsGraphArray) SortStable(less func(a, b StatsGraph) bool) StatsGraphArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of StatsGraph.
func (s StatsGraphArray) Retain(keep func(x StatsGraph) bool) StatsGraphArray {
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
func (s StatsGraphArray) First() (v StatsGraph, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s StatsGraphArray) Last() (v StatsGraph, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *StatsGraphArray) PopFirst() (v StatsGraph, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero StatsGraph
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *StatsGraphArray) Pop() (v StatsGraph, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
