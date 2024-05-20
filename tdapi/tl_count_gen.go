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

// Count represents TL type `count#4d38f104`.
type Count struct {
	// Count
	Count int32
}

// CountTypeID is TL type id of Count.
const CountTypeID = 0x4d38f104

// Ensuring interfaces in compile-time for Count.
var (
	_ bin.Encoder     = &Count{}
	_ bin.Decoder     = &Count{}
	_ bin.BareEncoder = &Count{}
	_ bin.BareDecoder = &Count{}
)

func (c *Count) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *Count) String() string {
	if c == nil {
		return "Count(nil)"
	}
	type Alias Count
	return fmt.Sprintf("Count%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Count) TypeID() uint32 {
	return CountTypeID
}

// TypeName returns name of type in TL schema.
func (*Count) TypeName() string {
	return "count"
}

// TypeInfo returns info about TL type.
func (c *Count) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "count",
		ID:   CountTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *Count) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode count#4d38f104 as nil")
	}
	b.PutID(CountTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *Count) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode count#4d38f104 as nil")
	}
	b.PutInt32(c.Count)
	return nil
}

// Decode implements bin.Decoder.
func (c *Count) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode count#4d38f104 to nil")
	}
	if err := b.ConsumeID(CountTypeID); err != nil {
		return fmt.Errorf("unable to decode count#4d38f104: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *Count) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode count#4d38f104 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode count#4d38f104: field count: %w", err)
		}
		c.Count = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *Count) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode count#4d38f104 as nil")
	}
	b.ObjStart()
	b.PutID("count")
	b.Comma()
	b.FieldStart("count")
	b.PutInt32(c.Count)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *Count) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode count#4d38f104 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("count"); err != nil {
				return fmt.Errorf("unable to decode count#4d38f104: %w", err)
			}
		case "count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode count#4d38f104: field count: %w", err)
			}
			c.Count = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCount returns value of Count field.
func (c *Count) GetCount() (value int32) {
	if c == nil {
		return
	}
	return c.Count
}
