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

// CheckStickerSetNameResultOk represents TL type `checkStickerSetNameResultOk#ac4bf258`.
type CheckStickerSetNameResultOk struct {
}

// CheckStickerSetNameResultOkTypeID is TL type id of CheckStickerSetNameResultOk.
const CheckStickerSetNameResultOkTypeID = 0xac4bf258

// construct implements constructor of CheckStickerSetNameResultClass.
func (c CheckStickerSetNameResultOk) construct() CheckStickerSetNameResultClass { return &c }

// Ensuring interfaces in compile-time for CheckStickerSetNameResultOk.
var (
	_ bin.Encoder     = &CheckStickerSetNameResultOk{}
	_ bin.Decoder     = &CheckStickerSetNameResultOk{}
	_ bin.BareEncoder = &CheckStickerSetNameResultOk{}
	_ bin.BareDecoder = &CheckStickerSetNameResultOk{}

	_ CheckStickerSetNameResultClass = &CheckStickerSetNameResultOk{}
)

func (c *CheckStickerSetNameResultOk) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckStickerSetNameResultOk) String() string {
	if c == nil {
		return "CheckStickerSetNameResultOk(nil)"
	}
	type Alias CheckStickerSetNameResultOk
	return fmt.Sprintf("CheckStickerSetNameResultOk%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckStickerSetNameResultOk) TypeID() uint32 {
	return CheckStickerSetNameResultOkTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckStickerSetNameResultOk) TypeName() string {
	return "checkStickerSetNameResultOk"
}

// TypeInfo returns info about TL type.
func (c *CheckStickerSetNameResultOk) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkStickerSetNameResultOk",
		ID:   CheckStickerSetNameResultOkTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckStickerSetNameResultOk) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetNameResultOk#ac4bf258 as nil")
	}
	b.PutID(CheckStickerSetNameResultOkTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckStickerSetNameResultOk) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetNameResultOk#ac4bf258 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckStickerSetNameResultOk) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetNameResultOk#ac4bf258 to nil")
	}
	if err := b.ConsumeID(CheckStickerSetNameResultOkTypeID); err != nil {
		return fmt.Errorf("unable to decode checkStickerSetNameResultOk#ac4bf258: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckStickerSetNameResultOk) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetNameResultOk#ac4bf258 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckStickerSetNameResultOk) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetNameResultOk#ac4bf258 as nil")
	}
	b.ObjStart()
	b.PutID("checkStickerSetNameResultOk")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckStickerSetNameResultOk) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetNameResultOk#ac4bf258 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkStickerSetNameResultOk"); err != nil {
				return fmt.Errorf("unable to decode checkStickerSetNameResultOk#ac4bf258: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// CheckStickerSetNameResultNameInvalid represents TL type `checkStickerSetNameResultNameInvalid#a9bf234`.
type CheckStickerSetNameResultNameInvalid struct {
}

// CheckStickerSetNameResultNameInvalidTypeID is TL type id of CheckStickerSetNameResultNameInvalid.
const CheckStickerSetNameResultNameInvalidTypeID = 0xa9bf234

// construct implements constructor of CheckStickerSetNameResultClass.
func (c CheckStickerSetNameResultNameInvalid) construct() CheckStickerSetNameResultClass { return &c }

// Ensuring interfaces in compile-time for CheckStickerSetNameResultNameInvalid.
var (
	_ bin.Encoder     = &CheckStickerSetNameResultNameInvalid{}
	_ bin.Decoder     = &CheckStickerSetNameResultNameInvalid{}
	_ bin.BareEncoder = &CheckStickerSetNameResultNameInvalid{}
	_ bin.BareDecoder = &CheckStickerSetNameResultNameInvalid{}

	_ CheckStickerSetNameResultClass = &CheckStickerSetNameResultNameInvalid{}
)

func (c *CheckStickerSetNameResultNameInvalid) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckStickerSetNameResultNameInvalid) String() string {
	if c == nil {
		return "CheckStickerSetNameResultNameInvalid(nil)"
	}
	type Alias CheckStickerSetNameResultNameInvalid
	return fmt.Sprintf("CheckStickerSetNameResultNameInvalid%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckStickerSetNameResultNameInvalid) TypeID() uint32 {
	return CheckStickerSetNameResultNameInvalidTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckStickerSetNameResultNameInvalid) TypeName() string {
	return "checkStickerSetNameResultNameInvalid"
}

// TypeInfo returns info about TL type.
func (c *CheckStickerSetNameResultNameInvalid) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkStickerSetNameResultNameInvalid",
		ID:   CheckStickerSetNameResultNameInvalidTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckStickerSetNameResultNameInvalid) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetNameResultNameInvalid#a9bf234 as nil")
	}
	b.PutID(CheckStickerSetNameResultNameInvalidTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckStickerSetNameResultNameInvalid) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetNameResultNameInvalid#a9bf234 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckStickerSetNameResultNameInvalid) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetNameResultNameInvalid#a9bf234 to nil")
	}
	if err := b.ConsumeID(CheckStickerSetNameResultNameInvalidTypeID); err != nil {
		return fmt.Errorf("unable to decode checkStickerSetNameResultNameInvalid#a9bf234: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckStickerSetNameResultNameInvalid) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetNameResultNameInvalid#a9bf234 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckStickerSetNameResultNameInvalid) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetNameResultNameInvalid#a9bf234 as nil")
	}
	b.ObjStart()
	b.PutID("checkStickerSetNameResultNameInvalid")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckStickerSetNameResultNameInvalid) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetNameResultNameInvalid#a9bf234 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkStickerSetNameResultNameInvalid"); err != nil {
				return fmt.Errorf("unable to decode checkStickerSetNameResultNameInvalid#a9bf234: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// CheckStickerSetNameResultNameOccupied represents TL type `checkStickerSetNameResultNameOccupied#3c60dc88`.
type CheckStickerSetNameResultNameOccupied struct {
}

// CheckStickerSetNameResultNameOccupiedTypeID is TL type id of CheckStickerSetNameResultNameOccupied.
const CheckStickerSetNameResultNameOccupiedTypeID = 0x3c60dc88

// construct implements constructor of CheckStickerSetNameResultClass.
func (c CheckStickerSetNameResultNameOccupied) construct() CheckStickerSetNameResultClass { return &c }

// Ensuring interfaces in compile-time for CheckStickerSetNameResultNameOccupied.
var (
	_ bin.Encoder     = &CheckStickerSetNameResultNameOccupied{}
	_ bin.Decoder     = &CheckStickerSetNameResultNameOccupied{}
	_ bin.BareEncoder = &CheckStickerSetNameResultNameOccupied{}
	_ bin.BareDecoder = &CheckStickerSetNameResultNameOccupied{}

	_ CheckStickerSetNameResultClass = &CheckStickerSetNameResultNameOccupied{}
)

func (c *CheckStickerSetNameResultNameOccupied) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckStickerSetNameResultNameOccupied) String() string {
	if c == nil {
		return "CheckStickerSetNameResultNameOccupied(nil)"
	}
	type Alias CheckStickerSetNameResultNameOccupied
	return fmt.Sprintf("CheckStickerSetNameResultNameOccupied%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckStickerSetNameResultNameOccupied) TypeID() uint32 {
	return CheckStickerSetNameResultNameOccupiedTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckStickerSetNameResultNameOccupied) TypeName() string {
	return "checkStickerSetNameResultNameOccupied"
}

// TypeInfo returns info about TL type.
func (c *CheckStickerSetNameResultNameOccupied) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkStickerSetNameResultNameOccupied",
		ID:   CheckStickerSetNameResultNameOccupiedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckStickerSetNameResultNameOccupied) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetNameResultNameOccupied#3c60dc88 as nil")
	}
	b.PutID(CheckStickerSetNameResultNameOccupiedTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckStickerSetNameResultNameOccupied) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetNameResultNameOccupied#3c60dc88 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckStickerSetNameResultNameOccupied) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetNameResultNameOccupied#3c60dc88 to nil")
	}
	if err := b.ConsumeID(CheckStickerSetNameResultNameOccupiedTypeID); err != nil {
		return fmt.Errorf("unable to decode checkStickerSetNameResultNameOccupied#3c60dc88: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckStickerSetNameResultNameOccupied) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetNameResultNameOccupied#3c60dc88 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckStickerSetNameResultNameOccupied) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetNameResultNameOccupied#3c60dc88 as nil")
	}
	b.ObjStart()
	b.PutID("checkStickerSetNameResultNameOccupied")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckStickerSetNameResultNameOccupied) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetNameResultNameOccupied#3c60dc88 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkStickerSetNameResultNameOccupied"); err != nil {
				return fmt.Errorf("unable to decode checkStickerSetNameResultNameOccupied#3c60dc88: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// CheckStickerSetNameResultClassName is schema name of CheckStickerSetNameResultClass.
const CheckStickerSetNameResultClassName = "CheckStickerSetNameResult"

// CheckStickerSetNameResultClass represents CheckStickerSetNameResult generic type.
//
// Example:
//
//	g, err := tdapi.DecodeCheckStickerSetNameResult(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.CheckStickerSetNameResultOk: // checkStickerSetNameResultOk#ac4bf258
//	case *tdapi.CheckStickerSetNameResultNameInvalid: // checkStickerSetNameResultNameInvalid#a9bf234
//	case *tdapi.CheckStickerSetNameResultNameOccupied: // checkStickerSetNameResultNameOccupied#3c60dc88
//	default: panic(v)
//	}
type CheckStickerSetNameResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() CheckStickerSetNameResultClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeCheckStickerSetNameResult implements binary de-serialization for CheckStickerSetNameResultClass.
func DecodeCheckStickerSetNameResult(buf *bin.Buffer) (CheckStickerSetNameResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case CheckStickerSetNameResultOkTypeID:
		// Decoding checkStickerSetNameResultOk#ac4bf258.
		v := CheckStickerSetNameResultOk{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckStickerSetNameResultClass: %w", err)
		}
		return &v, nil
	case CheckStickerSetNameResultNameInvalidTypeID:
		// Decoding checkStickerSetNameResultNameInvalid#a9bf234.
		v := CheckStickerSetNameResultNameInvalid{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckStickerSetNameResultClass: %w", err)
		}
		return &v, nil
	case CheckStickerSetNameResultNameOccupiedTypeID:
		// Decoding checkStickerSetNameResultNameOccupied#3c60dc88.
		v := CheckStickerSetNameResultNameOccupied{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckStickerSetNameResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode CheckStickerSetNameResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONCheckStickerSetNameResult implements binary de-serialization for CheckStickerSetNameResultClass.
func DecodeTDLibJSONCheckStickerSetNameResult(buf tdjson.Decoder) (CheckStickerSetNameResultClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "checkStickerSetNameResultOk":
		// Decoding checkStickerSetNameResultOk#ac4bf258.
		v := CheckStickerSetNameResultOk{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckStickerSetNameResultClass: %w", err)
		}
		return &v, nil
	case "checkStickerSetNameResultNameInvalid":
		// Decoding checkStickerSetNameResultNameInvalid#a9bf234.
		v := CheckStickerSetNameResultNameInvalid{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckStickerSetNameResultClass: %w", err)
		}
		return &v, nil
	case "checkStickerSetNameResultNameOccupied":
		// Decoding checkStickerSetNameResultNameOccupied#3c60dc88.
		v := CheckStickerSetNameResultNameOccupied{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode CheckStickerSetNameResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode CheckStickerSetNameResultClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// CheckStickerSetNameResult boxes the CheckStickerSetNameResultClass providing a helper.
type CheckStickerSetNameResultBox struct {
	CheckStickerSetNameResult CheckStickerSetNameResultClass
}

// Decode implements bin.Decoder for CheckStickerSetNameResultBox.
func (b *CheckStickerSetNameResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode CheckStickerSetNameResultBox to nil")
	}
	v, err := DecodeCheckStickerSetNameResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CheckStickerSetNameResult = v
	return nil
}

// Encode implements bin.Encode for CheckStickerSetNameResultBox.
func (b *CheckStickerSetNameResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.CheckStickerSetNameResult == nil {
		return fmt.Errorf("unable to encode CheckStickerSetNameResultClass as nil")
	}
	return b.CheckStickerSetNameResult.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for CheckStickerSetNameResultBox.
func (b *CheckStickerSetNameResultBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode CheckStickerSetNameResultBox to nil")
	}
	v, err := DecodeTDLibJSONCheckStickerSetNameResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.CheckStickerSetNameResult = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for CheckStickerSetNameResultBox.
func (b *CheckStickerSetNameResultBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.CheckStickerSetNameResult == nil {
		return fmt.Errorf("unable to encode CheckStickerSetNameResultClass as nil")
	}
	return b.CheckStickerSetNameResult.EncodeTDLibJSON(buf)
}
