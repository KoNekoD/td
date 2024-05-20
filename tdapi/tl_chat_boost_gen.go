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

// ChatBoost represents TL type `chatBoost#96bfccb2`.
type ChatBoost struct {
	// Unique identifier of the boost
	ID string
	// The number of identical boosts applied
	Count int32
	// Source of the boost
	Source ChatBoostSourceClass
	// Point in time (Unix timestamp) when the chat was boosted
	StartDate int32
	// Point in time (Unix timestamp) when the boost will expire
	ExpirationDate int32
}

// ChatBoostTypeID is TL type id of ChatBoost.
const ChatBoostTypeID = 0x96bfccb2

// Ensuring interfaces in compile-time for ChatBoost.
var (
	_ bin.Encoder     = &ChatBoost{}
	_ bin.Decoder     = &ChatBoost{}
	_ bin.BareEncoder = &ChatBoost{}
	_ bin.BareDecoder = &ChatBoost{}
)

func (c *ChatBoost) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ID == "") {
		return false
	}
	if !(c.Count == 0) {
		return false
	}
	if !(c.Source == nil) {
		return false
	}
	if !(c.StartDate == 0) {
		return false
	}
	if !(c.ExpirationDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatBoost) String() string {
	if c == nil {
		return "ChatBoost(nil)"
	}
	type Alias ChatBoost
	return fmt.Sprintf("ChatBoost%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatBoost) TypeID() uint32 {
	return ChatBoostTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatBoost) TypeName() string {
	return "chatBoost"
}

// TypeInfo returns info about TL type.
func (c *ChatBoost) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatBoost",
		ID:   ChatBoostTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Source",
			SchemaName: "source",
		},
		{
			Name:       "StartDate",
			SchemaName: "start_date",
		},
		{
			Name:       "ExpirationDate",
			SchemaName: "expiration_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatBoost) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoost#96bfccb2 as nil")
	}
	b.PutID(ChatBoostTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatBoost) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoost#96bfccb2 as nil")
	}
	b.PutString(c.ID)
	b.PutInt32(c.Count)
	if c.Source == nil {
		return fmt.Errorf("unable to encode chatBoost#96bfccb2: field source is nil")
	}
	if err := c.Source.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatBoost#96bfccb2: field source: %w", err)
	}
	b.PutInt32(c.StartDate)
	b.PutInt32(c.ExpirationDate)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatBoost) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoost#96bfccb2 to nil")
	}
	if err := b.ConsumeID(ChatBoostTypeID); err != nil {
		return fmt.Errorf("unable to decode chatBoost#96bfccb2: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatBoost) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoost#96bfccb2 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoost#96bfccb2: field id: %w", err)
		}
		c.ID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoost#96bfccb2: field count: %w", err)
		}
		c.Count = value
	}
	{
		value, err := DecodeChatBoostSource(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatBoost#96bfccb2: field source: %w", err)
		}
		c.Source = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoost#96bfccb2: field start_date: %w", err)
		}
		c.StartDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoost#96bfccb2: field expiration_date: %w", err)
		}
		c.ExpirationDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatBoost) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoost#96bfccb2 as nil")
	}
	b.ObjStart()
	b.PutID("chatBoost")
	b.Comma()
	b.FieldStart("id")
	b.PutString(c.ID)
	b.Comma()
	b.FieldStart("count")
	b.PutInt32(c.Count)
	b.Comma()
	b.FieldStart("source")
	if c.Source == nil {
		return fmt.Errorf("unable to encode chatBoost#96bfccb2: field source is nil")
	}
	if err := c.Source.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatBoost#96bfccb2: field source: %w", err)
	}
	b.Comma()
	b.FieldStart("start_date")
	b.PutInt32(c.StartDate)
	b.Comma()
	b.FieldStart("expiration_date")
	b.PutInt32(c.ExpirationDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatBoost) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoost#96bfccb2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatBoost"); err != nil {
				return fmt.Errorf("unable to decode chatBoost#96bfccb2: %w", err)
			}
		case "id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoost#96bfccb2: field id: %w", err)
			}
			c.ID = value
		case "count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoost#96bfccb2: field count: %w", err)
			}
			c.Count = value
		case "source":
			value, err := DecodeTDLibJSONChatBoostSource(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatBoost#96bfccb2: field source: %w", err)
			}
			c.Source = value
		case "start_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoost#96bfccb2: field start_date: %w", err)
			}
			c.StartDate = value
		case "expiration_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoost#96bfccb2: field expiration_date: %w", err)
			}
			c.ExpirationDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (c *ChatBoost) GetID() (value string) {
	if c == nil {
		return
	}
	return c.ID
}

// GetCount returns value of Count field.
func (c *ChatBoost) GetCount() (value int32) {
	if c == nil {
		return
	}
	return c.Count
}

// GetSource returns value of Source field.
func (c *ChatBoost) GetSource() (value ChatBoostSourceClass) {
	if c == nil {
		return
	}
	return c.Source
}

// GetStartDate returns value of StartDate field.
func (c *ChatBoost) GetStartDate() (value int32) {
	if c == nil {
		return
	}
	return c.StartDate
}

// GetExpirationDate returns value of ExpirationDate field.
func (c *ChatBoost) GetExpirationDate() (value int32) {
	if c == nil {
		return
	}
	return c.ExpirationDate
}
