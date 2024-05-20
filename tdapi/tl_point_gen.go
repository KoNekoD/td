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

// Point represents TL type `point#1a13f5b9`.
type Point struct {
	// The point's first coordinate
	X float64
	// The point's second coordinate
	Y float64
}

// PointTypeID is TL type id of Point.
const PointTypeID = 0x1a13f5b9

// Ensuring interfaces in compile-time for Point.
var (
	_ bin.Encoder     = &Point{}
	_ bin.Decoder     = &Point{}
	_ bin.BareEncoder = &Point{}
	_ bin.BareDecoder = &Point{}
)

func (p *Point) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.X == 0) {
		return false
	}
	if !(p.Y == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *Point) String() string {
	if p == nil {
		return "Point(nil)"
	}
	type Alias Point
	return fmt.Sprintf("Point%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Point) TypeID() uint32 {
	return PointTypeID
}

// TypeName returns name of type in TL schema.
func (*Point) TypeName() string {
	return "point"
}

// TypeInfo returns info about TL type.
func (p *Point) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "point",
		ID:   PointTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "X",
			SchemaName: "x",
		},
		{
			Name:       "Y",
			SchemaName: "y",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *Point) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode point#1a13f5b9 as nil")
	}
	b.PutID(PointTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *Point) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode point#1a13f5b9 as nil")
	}
	b.PutDouble(p.X)
	b.PutDouble(p.Y)
	return nil
}

// Decode implements bin.Decoder.
func (p *Point) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode point#1a13f5b9 to nil")
	}
	if err := b.ConsumeID(PointTypeID); err != nil {
		return fmt.Errorf("unable to decode point#1a13f5b9: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *Point) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode point#1a13f5b9 to nil")
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode point#1a13f5b9: field x: %w", err)
		}
		p.X = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode point#1a13f5b9: field y: %w", err)
		}
		p.Y = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *Point) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode point#1a13f5b9 as nil")
	}
	b.ObjStart()
	b.PutID("point")
	b.Comma()
	b.FieldStart("x")
	b.PutDouble(p.X)
	b.Comma()
	b.FieldStart("y")
	b.PutDouble(p.Y)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *Point) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode point#1a13f5b9 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("point"); err != nil {
				return fmt.Errorf("unable to decode point#1a13f5b9: %w", err)
			}
		case "x":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode point#1a13f5b9: field x: %w", err)
			}
			p.X = value
		case "y":
			value, err := b.Double()
			if err != nil {
				return fmt.Errorf("unable to decode point#1a13f5b9: field y: %w", err)
			}
			p.Y = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetX returns value of X field.
func (p *Point) GetX() (value float64) {
	if p == nil {
		return
	}
	return p.X
}

// GetY returns value of Y field.
func (p *Point) GetY() (value float64) {
	if p == nil {
		return
	}
	return p.Y
}
