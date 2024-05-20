// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// ChatLocation represents TL type `chatLocation#a29b8f21`.
type ChatLocation struct {
	// The location
	Location Location
	// Location address; 1-64 characters, as defined by the chat owner
	Address string
}

// ChatLocationTypeID is TL type id of ChatLocation.
const ChatLocationTypeID = 0xa29b8f21

// Ensuring interfaces in compile-time for ChatLocation.
var (
	_ bin.Encoder     = &ChatLocation{}
	_ bin.Decoder     = &ChatLocation{}
	_ bin.BareEncoder = &ChatLocation{}
	_ bin.BareDecoder = &ChatLocation{}
)

func (c *ChatLocation) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Location.Zero()) {
		return false
	}
	if !(c.Address == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatLocation) String() string {
	if c == nil {
		return "ChatLocation(nil)"
	}
	type Alias ChatLocation
	return fmt.Sprintf("ChatLocation%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatLocation) TypeID() uint32 {
	return ChatLocationTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatLocation) TypeName() string {
	return "chatLocation"
}

// TypeInfo returns info about TL type.
func (c *ChatLocation) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatLocation",
		ID:   ChatLocationTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Location",
			SchemaName: "location",
		},
		{
			Name:       "Address",
			SchemaName: "address",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatLocation) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatLocation#a29b8f21 as nil")
	}
	b.PutID(ChatLocationTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatLocation) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatLocation#a29b8f21 as nil")
	}
	if err := c.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatLocation#a29b8f21: field location: %w", err)
	}
	b.PutString(c.Address)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatLocation) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatLocation#a29b8f21 to nil")
	}
	if err := b.ConsumeID(ChatLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode chatLocation#a29b8f21: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatLocation) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatLocation#a29b8f21 to nil")
	}
	{
		if err := c.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatLocation#a29b8f21: field location: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatLocation#a29b8f21: field address: %w", err)
		}
		c.Address = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatLocation) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatLocation#a29b8f21 as nil")
	}
	b.ObjStart()
	b.PutID("chatLocation")
	b.Comma()
	b.FieldStart("location")
	if err := c.Location.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatLocation#a29b8f21: field location: %w", err)
	}
	b.Comma()
	b.FieldStart("address")
	b.PutString(c.Address)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatLocation) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatLocation#a29b8f21 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatLocation"); err != nil {
				return fmt.Errorf("unable to decode chatLocation#a29b8f21: %w", err)
			}
		case "location":
			if err := c.Location.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatLocation#a29b8f21: field location: %w", err)
			}
		case "address":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatLocation#a29b8f21: field address: %w", err)
			}
			c.Address = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLocation returns value of Location field.
func (c *ChatLocation) GetLocation() (value Location) {
	if c == nil {
		return
	}
	return c.Location
}

// GetAddress returns value of Address field.
func (c *ChatLocation) GetAddress() (value string) {
	if c == nil {
		return
	}
	return c.Address
}
