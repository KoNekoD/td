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

// SearchResultsCalendarPeriod represents TL type `searchResultsCalendarPeriod#c9b0539f`.
// Information about found messages sent on a specific day, used to split the messages in
// messages.searchResultsCalendar¹ constructors by days.
// Multiple searchResultsCalendarPeriod constructors are returned in messages
// searchResultsCalendar¹, each containing information about the first, last and total
// number of messages matching the filter that were sent on a specific day.
//
// Links:
//  1. https://core.telegram.org/constructor/messages.searchResultsCalendar
//  2. https://core.telegram.org/constructor/messages.searchResultsCalendar
//
// See https://core.telegram.org/constructor/searchResultsCalendarPeriod for reference.
type SearchResultsCalendarPeriod struct {
	// The day this object is referring to.
	Date int
	// First message ID that was sent on this day.
	MinMsgID int
	// Last message ID that was sent on this day.
	MaxMsgID int
	// All messages that were sent on this day.
	Count int
}

// SearchResultsCalendarPeriodTypeID is TL type id of SearchResultsCalendarPeriod.
const SearchResultsCalendarPeriodTypeID = 0xc9b0539f

// Ensuring interfaces in compile-time for SearchResultsCalendarPeriod.
var (
	_ bin.Encoder     = &SearchResultsCalendarPeriod{}
	_ bin.Decoder     = &SearchResultsCalendarPeriod{}
	_ bin.BareEncoder = &SearchResultsCalendarPeriod{}
	_ bin.BareDecoder = &SearchResultsCalendarPeriod{}
)

func (s *SearchResultsCalendarPeriod) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.MinMsgID == 0) {
		return false
	}
	if !(s.MaxMsgID == 0) {
		return false
	}
	if !(s.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchResultsCalendarPeriod) String() string {
	if s == nil {
		return "SearchResultsCalendarPeriod(nil)"
	}
	type Alias SearchResultsCalendarPeriod
	return fmt.Sprintf("SearchResultsCalendarPeriod%+v", Alias(*s))
}

// FillFrom fills SearchResultsCalendarPeriod from given interface.
func (s *SearchResultsCalendarPeriod) FillFrom(from interface {
	GetDate() (value int)
	GetMinMsgID() (value int)
	GetMaxMsgID() (value int)
	GetCount() (value int)
}) {
	s.Date = from.GetDate()
	s.MinMsgID = from.GetMinMsgID()
	s.MaxMsgID = from.GetMaxMsgID()
	s.Count = from.GetCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchResultsCalendarPeriod) TypeID() uint32 {
	return SearchResultsCalendarPeriodTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchResultsCalendarPeriod) TypeName() string {
	return "searchResultsCalendarPeriod"
}

// TypeInfo returns info about TL type.
func (s *SearchResultsCalendarPeriod) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchResultsCalendarPeriod",
		ID:   SearchResultsCalendarPeriodTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "MinMsgID",
			SchemaName: "min_msg_id",
		},
		{
			Name:       "MaxMsgID",
			SchemaName: "max_msg_id",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchResultsCalendarPeriod) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchResultsCalendarPeriod#c9b0539f as nil")
	}
	b.PutID(SearchResultsCalendarPeriodTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchResultsCalendarPeriod) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchResultsCalendarPeriod#c9b0539f as nil")
	}
	b.PutInt(s.Date)
	b.PutInt(s.MinMsgID)
	b.PutInt(s.MaxMsgID)
	b.PutInt(s.Count)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchResultsCalendarPeriod) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchResultsCalendarPeriod#c9b0539f to nil")
	}
	if err := b.ConsumeID(SearchResultsCalendarPeriodTypeID); err != nil {
		return fmt.Errorf("unable to decode searchResultsCalendarPeriod#c9b0539f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchResultsCalendarPeriod) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchResultsCalendarPeriod#c9b0539f to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode searchResultsCalendarPeriod#c9b0539f: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode searchResultsCalendarPeriod#c9b0539f: field min_msg_id: %w", err)
		}
		s.MinMsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode searchResultsCalendarPeriod#c9b0539f: field max_msg_id: %w", err)
		}
		s.MaxMsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode searchResultsCalendarPeriod#c9b0539f: field count: %w", err)
		}
		s.Count = value
	}
	return nil
}

// GetDate returns value of Date field.
func (s *SearchResultsCalendarPeriod) GetDate() (value int) {
	if s == nil {
		return
	}
	return s.Date
}

// GetMinMsgID returns value of MinMsgID field.
func (s *SearchResultsCalendarPeriod) GetMinMsgID() (value int) {
	if s == nil {
		return
	}
	return s.MinMsgID
}

// GetMaxMsgID returns value of MaxMsgID field.
func (s *SearchResultsCalendarPeriod) GetMaxMsgID() (value int) {
	if s == nil {
		return
	}
	return s.MaxMsgID
}

// GetCount returns value of Count field.
func (s *SearchResultsCalendarPeriod) GetCount() (value int) {
	if s == nil {
		return
	}
	return s.Count
}
