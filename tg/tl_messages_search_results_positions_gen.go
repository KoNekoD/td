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

// MessagesSearchResultsPositions represents TL type `messages.searchResultsPositions#53b22baf`.
// Information about sparse positions of messages
//
// See https://core.telegram.org/constructor/messages.searchResultsPositions for reference.
type MessagesSearchResultsPositions struct {
	// Total number of found messages
	Count int
	// List of message positions
	Positions []SearchResultPosition
}

// MessagesSearchResultsPositionsTypeID is TL type id of MessagesSearchResultsPositions.
const MessagesSearchResultsPositionsTypeID = 0x53b22baf

// Ensuring interfaces in compile-time for MessagesSearchResultsPositions.
var (
	_ bin.Encoder     = &MessagesSearchResultsPositions{}
	_ bin.Decoder     = &MessagesSearchResultsPositions{}
	_ bin.BareEncoder = &MessagesSearchResultsPositions{}
	_ bin.BareDecoder = &MessagesSearchResultsPositions{}
)

func (s *MessagesSearchResultsPositions) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Count == 0) {
		return false
	}
	if !(s.Positions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSearchResultsPositions) String() string {
	if s == nil {
		return "MessagesSearchResultsPositions(nil)"
	}
	type Alias MessagesSearchResultsPositions
	return fmt.Sprintf("MessagesSearchResultsPositions%+v", Alias(*s))
}

// FillFrom fills MessagesSearchResultsPositions from given interface.
func (s *MessagesSearchResultsPositions) FillFrom(from interface {
	GetCount() (value int)
	GetPositions() (value []SearchResultPosition)
}) {
	s.Count = from.GetCount()
	s.Positions = from.GetPositions()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSearchResultsPositions) TypeID() uint32 {
	return MessagesSearchResultsPositionsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSearchResultsPositions) TypeName() string {
	return "messages.searchResultsPositions"
}

// TypeInfo returns info about TL type.
func (s *MessagesSearchResultsPositions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.searchResultsPositions",
		ID:   MessagesSearchResultsPositionsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Positions",
			SchemaName: "positions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSearchResultsPositions) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchResultsPositions#53b22baf as nil")
	}
	b.PutID(MessagesSearchResultsPositionsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSearchResultsPositions) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchResultsPositions#53b22baf as nil")
	}
	b.PutInt(s.Count)
	b.PutVectorHeader(len(s.Positions))
	for idx, v := range s.Positions {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.searchResultsPositions#53b22baf: field positions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSearchResultsPositions) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchResultsPositions#53b22baf to nil")
	}
	if err := b.ConsumeID(MessagesSearchResultsPositionsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.searchResultsPositions#53b22baf: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSearchResultsPositions) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchResultsPositions#53b22baf to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchResultsPositions#53b22baf: field count: %w", err)
		}
		s.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchResultsPositions#53b22baf: field positions: %w", err)
		}

		if headerLen > 0 {
			s.Positions = make([]SearchResultPosition, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SearchResultPosition
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.searchResultsPositions#53b22baf: field positions: %w", err)
			}
			s.Positions = append(s.Positions, value)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (s *MessagesSearchResultsPositions) GetCount() (value int) {
	if s == nil {
		return
	}
	return s.Count
}

// GetPositions returns value of Positions field.
func (s *MessagesSearchResultsPositions) GetPositions() (value []SearchResultPosition) {
	if s == nil {
		return
	}
	return s.Positions
}
