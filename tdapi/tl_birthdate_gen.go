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

// Birthdate represents TL type `birthdate#61fe6d1e`.
type Birthdate struct {
	// Day of the month; 1-31
	Day int32
	// Month of the year; 1-12
	Month int32
	// Birth year; 0 if unknown
	Year int32
}

// BirthdateTypeID is TL type id of Birthdate.
const BirthdateTypeID = 0x61fe6d1e

// Ensuring interfaces in compile-time for Birthdate.
var (
	_ bin.Encoder     = &Birthdate{}
	_ bin.Decoder     = &Birthdate{}
	_ bin.BareEncoder = &Birthdate{}
	_ bin.BareDecoder = &Birthdate{}
)

func (b *Birthdate) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Day == 0) {
		return false
	}
	if !(b.Month == 0) {
		return false
	}
	if !(b.Year == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *Birthdate) String() string {
	if b == nil {
		return "Birthdate(nil)"
	}
	type Alias Birthdate
	return fmt.Sprintf("Birthdate%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Birthdate) TypeID() uint32 {
	return BirthdateTypeID
}

// TypeName returns name of type in TL schema.
func (*Birthdate) TypeName() string {
	return "birthdate"
}

// TypeInfo returns info about TL type.
func (b *Birthdate) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "birthdate",
		ID:   BirthdateTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Day",
			SchemaName: "day",
		},
		{
			Name:       "Month",
			SchemaName: "month",
		},
		{
			Name:       "Year",
			SchemaName: "year",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *Birthdate) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode birthdate#61fe6d1e as nil")
	}
	buf.PutID(BirthdateTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *Birthdate) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode birthdate#61fe6d1e as nil")
	}
	buf.PutInt32(b.Day)
	buf.PutInt32(b.Month)
	buf.PutInt32(b.Year)
	return nil
}

// Decode implements bin.Decoder.
func (b *Birthdate) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode birthdate#61fe6d1e to nil")
	}
	if err := buf.ConsumeID(BirthdateTypeID); err != nil {
		return fmt.Errorf("unable to decode birthdate#61fe6d1e: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *Birthdate) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode birthdate#61fe6d1e to nil")
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode birthdate#61fe6d1e: field day: %w", err)
		}
		b.Day = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode birthdate#61fe6d1e: field month: %w", err)
		}
		b.Month = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode birthdate#61fe6d1e: field year: %w", err)
		}
		b.Year = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *Birthdate) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode birthdate#61fe6d1e as nil")
	}
	buf.ObjStart()
	buf.PutID("birthdate")
	buf.Comma()
	buf.FieldStart("day")
	buf.PutInt32(b.Day)
	buf.Comma()
	buf.FieldStart("month")
	buf.PutInt32(b.Month)
	buf.Comma()
	buf.FieldStart("year")
	buf.PutInt32(b.Year)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *Birthdate) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode birthdate#61fe6d1e to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("birthdate"); err != nil {
				return fmt.Errorf("unable to decode birthdate#61fe6d1e: %w", err)
			}
		case "day":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode birthdate#61fe6d1e: field day: %w", err)
			}
			b.Day = value
		case "month":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode birthdate#61fe6d1e: field month: %w", err)
			}
			b.Month = value
		case "year":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode birthdate#61fe6d1e: field year: %w", err)
			}
			b.Year = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetDay returns value of Day field.
func (b *Birthdate) GetDay() (value int32) {
	if b == nil {
		return
	}
	return b.Day
}

// GetMonth returns value of Month field.
func (b *Birthdate) GetMonth() (value int32) {
	if b == nil {
		return
	}
	return b.Month
}

// GetYear returns value of Year field.
func (b *Birthdate) GetYear() (value int32) {
	if b == nil {
		return
	}
	return b.Year
}
