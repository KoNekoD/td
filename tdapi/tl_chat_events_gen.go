// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// ChatEvents represents TL type `chatEvents#d73ecdc4`.
type ChatEvents struct {
	// List of events
	Events []ChatEvent
}

// ChatEventsTypeID is TL type id of ChatEvents.
const ChatEventsTypeID = 0xd73ecdc4

// Ensuring interfaces in compile-time for ChatEvents.
var (
	_ bin.Encoder     = &ChatEvents{}
	_ bin.Decoder     = &ChatEvents{}
	_ bin.BareEncoder = &ChatEvents{}
	_ bin.BareDecoder = &ChatEvents{}
)

func (c *ChatEvents) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Events == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatEvents) String() string {
	if c == nil {
		return "ChatEvents(nil)"
	}
	type Alias ChatEvents
	return fmt.Sprintf("ChatEvents%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatEvents) TypeID() uint32 {
	return ChatEventsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatEvents) TypeName() string {
	return "chatEvents"
}

// TypeInfo returns info about TL type.
func (c *ChatEvents) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatEvents",
		ID:   ChatEventsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Events",
			SchemaName: "events",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatEvents) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatEvents#d73ecdc4 as nil")
	}
	b.PutID(ChatEventsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatEvents) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatEvents#d73ecdc4 as nil")
	}
	b.PutInt(len(c.Events))
	for idx, v := range c.Events {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatEvents#d73ecdc4: field events element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatEvents) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatEvents#d73ecdc4 to nil")
	}
	if err := b.ConsumeID(ChatEventsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatEvents#d73ecdc4: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatEvents) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatEvents#d73ecdc4 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatEvents#d73ecdc4: field events: %w", err)
		}

		if headerLen > 0 {
			c.Events = make([]ChatEvent, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatEvent
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatEvents#d73ecdc4: field events: %w", err)
			}
			c.Events = append(c.Events, value)
		}
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatEvents) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatEvents#d73ecdc4 as nil")
	}
	b.ObjStart()
	b.PutID("chatEvents")
	b.FieldStart("events")
	b.ArrStart()
	for idx, v := range c.Events {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatEvents#d73ecdc4: field events element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// GetEvents returns value of Events field.
func (c *ChatEvents) GetEvents() (value []ChatEvent) {
	return c.Events
}
