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

// ChatClassArray is adapter for slice of ChatClass.
type ChatClassArray []ChatClass

// Sort sorts slice of ChatClass.
func (s ChatClassArray) Sort(less func(a, b ChatClass) bool) ChatClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatClass.
func (s ChatClassArray) SortStable(less func(a, b ChatClass) bool) ChatClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatClass.
func (s ChatClassArray) Retain(keep func(x ChatClass) bool) ChatClassArray {
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
func (s ChatClassArray) First() (v ChatClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatClassArray) Last() (v ChatClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatClassArray) PopFirst() (v ChatClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatClassArray) Pop() (v ChatClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of ChatClass by ID.
func (s ChatClassArray) SortByID() ChatClassArray {
	return s.Sort(func(a, b ChatClass) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of ChatClass by ID.
func (s ChatClassArray) SortStableByID() ChatClassArray {
	return s.SortStable(func(a, b ChatClass) bool {
		return a.GetID() < b.GetID()
	})
}

// FillChatEmptyMap fills only ChatEmpty constructors to given map.
func (s ChatClassArray) FillChatEmptyMap(to map[int64]*ChatEmpty) {
	for _, elem := range s {
		value, ok := elem.(*ChatEmpty)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// ChatEmptyToMap collects only ChatEmpty constructors to map.
func (s ChatClassArray) ChatEmptyToMap() map[int64]*ChatEmpty {
	r := make(map[int64]*ChatEmpty, len(s))
	s.FillChatEmptyMap(r)
	return r
}

// AsChatEmpty returns copy with only ChatEmpty constructors.
func (s ChatClassArray) AsChatEmpty() (to ChatEmptyArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatEmpty)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillChatMap fills only Chat constructors to given map.
func (s ChatClassArray) FillChatMap(to map[int64]*Chat) {
	for _, elem := range s {
		value, ok := elem.(*Chat)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// ChatToMap collects only Chat constructors to map.
func (s ChatClassArray) ChatToMap() map[int64]*Chat {
	r := make(map[int64]*Chat, len(s))
	s.FillChatMap(r)
	return r
}

// AsChat returns copy with only Chat constructors.
func (s ChatClassArray) AsChat() (to ChatArray) {
	for _, elem := range s {
		value, ok := elem.(*Chat)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillChatForbiddenMap fills only ChatForbidden constructors to given map.
func (s ChatClassArray) FillChatForbiddenMap(to map[int64]*ChatForbidden) {
	for _, elem := range s {
		value, ok := elem.(*ChatForbidden)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// ChatForbiddenToMap collects only ChatForbidden constructors to map.
func (s ChatClassArray) ChatForbiddenToMap() map[int64]*ChatForbidden {
	r := make(map[int64]*ChatForbidden, len(s))
	s.FillChatForbiddenMap(r)
	return r
}

// AsChatForbidden returns copy with only ChatForbidden constructors.
func (s ChatClassArray) AsChatForbidden() (to ChatForbiddenArray) {
	for _, elem := range s {
		value, ok := elem.(*ChatForbidden)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillChannelMap fills only Channel constructors to given map.
func (s ChatClassArray) FillChannelMap(to map[int64]*Channel) {
	for _, elem := range s {
		value, ok := elem.(*Channel)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// ChannelToMap collects only Channel constructors to map.
func (s ChatClassArray) ChannelToMap() map[int64]*Channel {
	r := make(map[int64]*Channel, len(s))
	s.FillChannelMap(r)
	return r
}

// AsChannel returns copy with only Channel constructors.
func (s ChatClassArray) AsChannel() (to ChannelArray) {
	for _, elem := range s {
		value, ok := elem.(*Channel)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillChannelForbiddenMap fills only ChannelForbidden constructors to given map.
func (s ChatClassArray) FillChannelForbiddenMap(to map[int64]*ChannelForbidden) {
	for _, elem := range s {
		value, ok := elem.(*ChannelForbidden)
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// ChannelForbiddenToMap collects only ChannelForbidden constructors to map.
func (s ChatClassArray) ChannelForbiddenToMap() map[int64]*ChannelForbidden {
	r := make(map[int64]*ChannelForbidden, len(s))
	s.FillChannelForbiddenMap(r)
	return r
}

// AsChannelForbidden returns copy with only ChannelForbidden constructors.
func (s ChatClassArray) AsChannelForbidden() (to ChannelForbiddenArray) {
	for _, elem := range s {
		value, ok := elem.(*ChannelForbidden)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// FillNotEmptyMap fills only NotEmpty constructors to given map.
func (s ChatClassArray) FillNotEmptyMap(to map[int64]NotEmptyChat) {
	for _, elem := range s {
		value, ok := elem.AsNotEmpty()
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// NotEmptyToMap collects only NotEmpty constructors to map.
func (s ChatClassArray) NotEmptyToMap() map[int64]NotEmptyChat {
	r := make(map[int64]NotEmptyChat, len(s))
	s.FillNotEmptyMap(r)
	return r
}

// AppendOnlyNotEmpty appends only NotEmpty constructors to
// given slice.
func (s ChatClassArray) AppendOnlyNotEmpty(to []NotEmptyChat) []NotEmptyChat {
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
func (s ChatClassArray) AsNotEmpty() (to []NotEmptyChat) {
	return s.AppendOnlyNotEmpty(to)
}

// FirstAsNotEmpty returns first element of slice (if exists).
func (s ChatClassArray) FirstAsNotEmpty() (v NotEmptyChat, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// LastAsNotEmpty returns last element of slice (if exists).
func (s ChatClassArray) LastAsNotEmpty() (v NotEmptyChat, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopFirstAsNotEmpty returns element of slice (if exists).
func (s *ChatClassArray) PopFirstAsNotEmpty() (v NotEmptyChat, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// PopAsNotEmpty returns element of slice (if exists).
func (s *ChatClassArray) PopAsNotEmpty() (v NotEmptyChat, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotEmpty()
}

// FillNotForbiddenMap fills only NotForbidden constructors to given map.
func (s ChatClassArray) FillNotForbiddenMap(to map[int64]NotForbiddenChat) {
	for _, elem := range s {
		value, ok := elem.AsNotForbidden()
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// NotForbiddenToMap collects only NotForbidden constructors to map.
func (s ChatClassArray) NotForbiddenToMap() map[int64]NotForbiddenChat {
	r := make(map[int64]NotForbiddenChat, len(s))
	s.FillNotForbiddenMap(r)
	return r
}

// AppendOnlyNotForbidden appends only NotForbidden constructors to
// given slice.
func (s ChatClassArray) AppendOnlyNotForbidden(to []NotForbiddenChat) []NotForbiddenChat {
	for _, elem := range s {
		value, ok := elem.AsNotForbidden()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsNotForbidden returns copy with only NotForbidden constructors.
func (s ChatClassArray) AsNotForbidden() (to []NotForbiddenChat) {
	return s.AppendOnlyNotForbidden(to)
}

// FirstAsNotForbidden returns first element of slice (if exists).
func (s ChatClassArray) FirstAsNotForbidden() (v NotForbiddenChat, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// LastAsNotForbidden returns last element of slice (if exists).
func (s ChatClassArray) LastAsNotForbidden() (v NotForbiddenChat, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// PopFirstAsNotForbidden returns element of slice (if exists).
func (s *ChatClassArray) PopFirstAsNotForbidden() (v NotForbiddenChat, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// PopAsNotForbidden returns element of slice (if exists).
func (s *ChatClassArray) PopAsNotForbidden() (v NotForbiddenChat, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsNotForbidden()
}

// FillFullMap fills only Full constructors to given map.
func (s ChatClassArray) FillFullMap(to map[int64]FullChat) {
	for _, elem := range s {
		value, ok := elem.AsFull()
		if !ok {
			continue
		}
		to[value.GetID()] = value
	}
}

// FullToMap collects only Full constructors to map.
func (s ChatClassArray) FullToMap() map[int64]FullChat {
	r := make(map[int64]FullChat, len(s))
	s.FillFullMap(r)
	return r
}

// AppendOnlyFull appends only Full constructors to
// given slice.
func (s ChatClassArray) AppendOnlyFull(to []FullChat) []FullChat {
	for _, elem := range s {
		value, ok := elem.AsFull()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsFull returns copy with only Full constructors.
func (s ChatClassArray) AsFull() (to []FullChat) {
	return s.AppendOnlyFull(to)
}

// FirstAsFull returns first element of slice (if exists).
func (s ChatClassArray) FirstAsFull() (v FullChat, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsFull()
}

// LastAsFull returns last element of slice (if exists).
func (s ChatClassArray) LastAsFull() (v FullChat, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsFull()
}

// PopFirstAsFull returns element of slice (if exists).
func (s *ChatClassArray) PopFirstAsFull() (v FullChat, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsFull()
}

// PopAsFull returns element of slice (if exists).
func (s *ChatClassArray) PopAsFull() (v FullChat, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsFull()
}

// ChatEmptyArray is adapter for slice of ChatEmpty.
type ChatEmptyArray []ChatEmpty

// Sort sorts slice of ChatEmpty.
func (s ChatEmptyArray) Sort(less func(a, b ChatEmpty) bool) ChatEmptyArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatEmpty.
func (s ChatEmptyArray) SortStable(less func(a, b ChatEmpty) bool) ChatEmptyArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatEmpty.
func (s ChatEmptyArray) Retain(keep func(x ChatEmpty) bool) ChatEmptyArray {
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
func (s ChatEmptyArray) First() (v ChatEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatEmptyArray) Last() (v ChatEmpty, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatEmptyArray) PopFirst() (v ChatEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatEmpty
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatEmptyArray) Pop() (v ChatEmpty, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of ChatEmpty by ID.
func (s ChatEmptyArray) SortByID() ChatEmptyArray {
	return s.Sort(func(a, b ChatEmpty) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of ChatEmpty by ID.
func (s ChatEmptyArray) SortStableByID() ChatEmptyArray {
	return s.SortStable(func(a, b ChatEmpty) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s ChatEmptyArray) FillMap(to map[int64]ChatEmpty) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s ChatEmptyArray) ToMap() map[int64]ChatEmpty {
	r := make(map[int64]ChatEmpty, len(s))
	s.FillMap(r)
	return r
}

// ChatArray is adapter for slice of Chat.
type ChatArray []Chat

// Sort sorts slice of Chat.
func (s ChatArray) Sort(less func(a, b Chat) bool) ChatArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of Chat.
func (s ChatArray) SortStable(less func(a, b Chat) bool) ChatArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of Chat.
func (s ChatArray) Retain(keep func(x Chat) bool) ChatArray {
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
func (s ChatArray) First() (v Chat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatArray) Last() (v Chat, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatArray) PopFirst() (v Chat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero Chat
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatArray) Pop() (v Chat, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of Chat by ID.
func (s ChatArray) SortByID() ChatArray {
	return s.Sort(func(a, b Chat) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of Chat by ID.
func (s ChatArray) SortStableByID() ChatArray {
	return s.SortStable(func(a, b Chat) bool {
		return a.GetID() < b.GetID()
	})
}

// SortByDate sorts slice of Chat by Date.
func (s ChatArray) SortByDate() ChatArray {
	return s.Sort(func(a, b Chat) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of Chat by Date.
func (s ChatArray) SortStableByDate() ChatArray {
	return s.SortStable(func(a, b Chat) bool {
		return a.GetDate() < b.GetDate()
	})
}

// FillMap fills constructors to given map.
func (s ChatArray) FillMap(to map[int64]Chat) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s ChatArray) ToMap() map[int64]Chat {
	r := make(map[int64]Chat, len(s))
	s.FillMap(r)
	return r
}

// ChatForbiddenArray is adapter for slice of ChatForbidden.
type ChatForbiddenArray []ChatForbidden

// Sort sorts slice of ChatForbidden.
func (s ChatForbiddenArray) Sort(less func(a, b ChatForbidden) bool) ChatForbiddenArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChatForbidden.
func (s ChatForbiddenArray) SortStable(less func(a, b ChatForbidden) bool) ChatForbiddenArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChatForbidden.
func (s ChatForbiddenArray) Retain(keep func(x ChatForbidden) bool) ChatForbiddenArray {
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
func (s ChatForbiddenArray) First() (v ChatForbidden, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChatForbiddenArray) Last() (v ChatForbidden, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChatForbiddenArray) PopFirst() (v ChatForbidden, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChatForbidden
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChatForbiddenArray) Pop() (v ChatForbidden, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of ChatForbidden by ID.
func (s ChatForbiddenArray) SortByID() ChatForbiddenArray {
	return s.Sort(func(a, b ChatForbidden) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of ChatForbidden by ID.
func (s ChatForbiddenArray) SortStableByID() ChatForbiddenArray {
	return s.SortStable(func(a, b ChatForbidden) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s ChatForbiddenArray) FillMap(to map[int64]ChatForbidden) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s ChatForbiddenArray) ToMap() map[int64]ChatForbidden {
	r := make(map[int64]ChatForbidden, len(s))
	s.FillMap(r)
	return r
}

// ChannelArray is adapter for slice of Channel.
type ChannelArray []Channel

// Sort sorts slice of Channel.
func (s ChannelArray) Sort(less func(a, b Channel) bool) ChannelArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of Channel.
func (s ChannelArray) SortStable(less func(a, b Channel) bool) ChannelArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of Channel.
func (s ChannelArray) Retain(keep func(x Channel) bool) ChannelArray {
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
func (s ChannelArray) First() (v Channel, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChannelArray) Last() (v Channel, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChannelArray) PopFirst() (v Channel, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero Channel
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChannelArray) Pop() (v Channel, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of Channel by ID.
func (s ChannelArray) SortByID() ChannelArray {
	return s.Sort(func(a, b Channel) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of Channel by ID.
func (s ChannelArray) SortStableByID() ChannelArray {
	return s.SortStable(func(a, b Channel) bool {
		return a.GetID() < b.GetID()
	})
}

// SortByDate sorts slice of Channel by Date.
func (s ChannelArray) SortByDate() ChannelArray {
	return s.Sort(func(a, b Channel) bool {
		return a.GetDate() < b.GetDate()
	})
}

// SortStableByDate sorts slice of Channel by Date.
func (s ChannelArray) SortStableByDate() ChannelArray {
	return s.SortStable(func(a, b Channel) bool {
		return a.GetDate() < b.GetDate()
	})
}

// FillMap fills constructors to given map.
func (s ChannelArray) FillMap(to map[int64]Channel) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s ChannelArray) ToMap() map[int64]Channel {
	r := make(map[int64]Channel, len(s))
	s.FillMap(r)
	return r
}

// ChannelForbiddenArray is adapter for slice of ChannelForbidden.
type ChannelForbiddenArray []ChannelForbidden

// Sort sorts slice of ChannelForbidden.
func (s ChannelForbiddenArray) Sort(less func(a, b ChannelForbidden) bool) ChannelForbiddenArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of ChannelForbidden.
func (s ChannelForbiddenArray) SortStable(less func(a, b ChannelForbidden) bool) ChannelForbiddenArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of ChannelForbidden.
func (s ChannelForbiddenArray) Retain(keep func(x ChannelForbidden) bool) ChannelForbiddenArray {
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
func (s ChannelForbiddenArray) First() (v ChannelForbidden, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s ChannelForbiddenArray) Last() (v ChannelForbidden, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *ChannelForbiddenArray) PopFirst() (v ChannelForbidden, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero ChannelForbidden
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *ChannelForbiddenArray) Pop() (v ChannelForbidden, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SortByID sorts slice of ChannelForbidden by ID.
func (s ChannelForbiddenArray) SortByID() ChannelForbiddenArray {
	return s.Sort(func(a, b ChannelForbidden) bool {
		return a.GetID() < b.GetID()
	})
}

// SortStableByID sorts slice of ChannelForbidden by ID.
func (s ChannelForbiddenArray) SortStableByID() ChannelForbiddenArray {
	return s.SortStable(func(a, b ChannelForbidden) bool {
		return a.GetID() < b.GetID()
	})
}

// FillMap fills constructors to given map.
func (s ChannelForbiddenArray) FillMap(to map[int64]ChannelForbidden) {
	for _, value := range s {
		to[value.GetID()] = value
	}
}

// ToMap collects constructors to map.
func (s ChannelForbiddenArray) ToMap() map[int64]ChannelForbidden {
	r := make(map[int64]ChannelForbidden, len(s))
	s.FillMap(r)
	return r
}
