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

// TextEntity represents TL type `textEntity#8bab99a8`.
type TextEntity struct {
	// Offset of the entity, in UTF-16 code units
	Offset int32
	// Length of the entity, in UTF-16 code units
	Length int32
	// Type of the entity
	Type TextEntityTypeClass
}

// TextEntityTypeID is TL type id of TextEntity.
const TextEntityTypeID = 0x8bab99a8

// Ensuring interfaces in compile-time for TextEntity.
var (
	_ bin.Encoder     = &TextEntity{}
	_ bin.Decoder     = &TextEntity{}
	_ bin.BareEncoder = &TextEntity{}
	_ bin.BareDecoder = &TextEntity{}
)

func (t *TextEntity) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Offset == 0) {
		return false
	}
	if !(t.Length == 0) {
		return false
	}
	if !(t.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TextEntity) String() string {
	if t == nil {
		return "TextEntity(nil)"
	}
	type Alias TextEntity
	return fmt.Sprintf("TextEntity%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TextEntity) TypeID() uint32 {
	return TextEntityTypeID
}

// TypeName returns name of type in TL schema.
func (*TextEntity) TypeName() string {
	return "textEntity"
}

// TypeInfo returns info about TL type.
func (t *TextEntity) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "textEntity",
		ID:   TextEntityTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Length",
			SchemaName: "length",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TextEntity) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textEntity#8bab99a8 as nil")
	}
	b.PutID(TextEntityTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TextEntity) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode textEntity#8bab99a8 as nil")
	}
	b.PutInt32(t.Offset)
	b.PutInt32(t.Length)
	if t.Type == nil {
		return fmt.Errorf("unable to encode textEntity#8bab99a8: field type is nil")
	}
	if err := t.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode textEntity#8bab99a8: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TextEntity) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textEntity#8bab99a8 to nil")
	}
	if err := b.ConsumeID(TextEntityTypeID); err != nil {
		return fmt.Errorf("unable to decode textEntity#8bab99a8: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TextEntity) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode textEntity#8bab99a8 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode textEntity#8bab99a8: field offset: %w", err)
		}
		t.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode textEntity#8bab99a8: field length: %w", err)
		}
		t.Length = value
	}
	{
		value, err := DecodeTextEntityType(b)
		if err != nil {
			return fmt.Errorf("unable to decode textEntity#8bab99a8: field type: %w", err)
		}
		t.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TextEntity) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode textEntity#8bab99a8 as nil")
	}
	b.ObjStart()
	b.PutID("textEntity")
	b.Comma()
	b.FieldStart("offset")
	b.PutInt32(t.Offset)
	b.Comma()
	b.FieldStart("length")
	b.PutInt32(t.Length)
	b.Comma()
	b.FieldStart("type")
	if t.Type == nil {
		return fmt.Errorf("unable to encode textEntity#8bab99a8: field type is nil")
	}
	if err := t.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode textEntity#8bab99a8: field type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TextEntity) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode textEntity#8bab99a8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("textEntity"); err != nil {
				return fmt.Errorf("unable to decode textEntity#8bab99a8: %w", err)
			}
		case "offset":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode textEntity#8bab99a8: field offset: %w", err)
			}
			t.Offset = value
		case "length":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode textEntity#8bab99a8: field length: %w", err)
			}
			t.Length = value
		case "type":
			value, err := DecodeTDLibJSONTextEntityType(b)
			if err != nil {
				return fmt.Errorf("unable to decode textEntity#8bab99a8: field type: %w", err)
			}
			t.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetOffset returns value of Offset field.
func (t *TextEntity) GetOffset() (value int32) {
	if t == nil {
		return
	}
	return t.Offset
}

// GetLength returns value of Length field.
func (t *TextEntity) GetLength() (value int32) {
	if t == nil {
		return
	}
	return t.Length
}

// GetType returns value of Type field.
func (t *TextEntity) GetType() (value TextEntityTypeClass) {
	if t == nil {
		return
	}
	return t.Type
}
