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

// ThumbnailFormatJpeg represents TL type `thumbnailFormatJpeg#d90c5488`.
type ThumbnailFormatJpeg struct {
}

// ThumbnailFormatJpegTypeID is TL type id of ThumbnailFormatJpeg.
const ThumbnailFormatJpegTypeID = 0xd90c5488

// construct implements constructor of ThumbnailFormatClass.
func (t ThumbnailFormatJpeg) construct() ThumbnailFormatClass { return &t }

// Ensuring interfaces in compile-time for ThumbnailFormatJpeg.
var (
	_ bin.Encoder     = &ThumbnailFormatJpeg{}
	_ bin.Decoder     = &ThumbnailFormatJpeg{}
	_ bin.BareEncoder = &ThumbnailFormatJpeg{}
	_ bin.BareDecoder = &ThumbnailFormatJpeg{}

	_ ThumbnailFormatClass = &ThumbnailFormatJpeg{}
)

func (t *ThumbnailFormatJpeg) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThumbnailFormatJpeg) String() string {
	if t == nil {
		return "ThumbnailFormatJpeg(nil)"
	}
	type Alias ThumbnailFormatJpeg
	return fmt.Sprintf("ThumbnailFormatJpeg%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThumbnailFormatJpeg) TypeID() uint32 {
	return ThumbnailFormatJpegTypeID
}

// TypeName returns name of type in TL schema.
func (*ThumbnailFormatJpeg) TypeName() string {
	return "thumbnailFormatJpeg"
}

// TypeInfo returns info about TL type.
func (t *ThumbnailFormatJpeg) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "thumbnailFormatJpeg",
		ID:   ThumbnailFormatJpegTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThumbnailFormatJpeg) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatJpeg#d90c5488 as nil")
	}
	b.PutID(ThumbnailFormatJpegTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThumbnailFormatJpeg) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatJpeg#d90c5488 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ThumbnailFormatJpeg) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatJpeg#d90c5488 to nil")
	}
	if err := b.ConsumeID(ThumbnailFormatJpegTypeID); err != nil {
		return fmt.Errorf("unable to decode thumbnailFormatJpeg#d90c5488: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThumbnailFormatJpeg) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatJpeg#d90c5488 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ThumbnailFormatJpeg) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatJpeg#d90c5488 as nil")
	}
	b.ObjStart()
	b.PutID("thumbnailFormatJpeg")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ThumbnailFormatJpeg) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatJpeg#d90c5488 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("thumbnailFormatJpeg"); err != nil {
				return fmt.Errorf("unable to decode thumbnailFormatJpeg#d90c5488: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ThumbnailFormatGif represents TL type `thumbnailFormatGif#4aa3258a`.
type ThumbnailFormatGif struct {
}

// ThumbnailFormatGifTypeID is TL type id of ThumbnailFormatGif.
const ThumbnailFormatGifTypeID = 0x4aa3258a

// construct implements constructor of ThumbnailFormatClass.
func (t ThumbnailFormatGif) construct() ThumbnailFormatClass { return &t }

// Ensuring interfaces in compile-time for ThumbnailFormatGif.
var (
	_ bin.Encoder     = &ThumbnailFormatGif{}
	_ bin.Decoder     = &ThumbnailFormatGif{}
	_ bin.BareEncoder = &ThumbnailFormatGif{}
	_ bin.BareDecoder = &ThumbnailFormatGif{}

	_ ThumbnailFormatClass = &ThumbnailFormatGif{}
)

func (t *ThumbnailFormatGif) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThumbnailFormatGif) String() string {
	if t == nil {
		return "ThumbnailFormatGif(nil)"
	}
	type Alias ThumbnailFormatGif
	return fmt.Sprintf("ThumbnailFormatGif%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThumbnailFormatGif) TypeID() uint32 {
	return ThumbnailFormatGifTypeID
}

// TypeName returns name of type in TL schema.
func (*ThumbnailFormatGif) TypeName() string {
	return "thumbnailFormatGif"
}

// TypeInfo returns info about TL type.
func (t *ThumbnailFormatGif) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "thumbnailFormatGif",
		ID:   ThumbnailFormatGifTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThumbnailFormatGif) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatGif#4aa3258a as nil")
	}
	b.PutID(ThumbnailFormatGifTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThumbnailFormatGif) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatGif#4aa3258a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ThumbnailFormatGif) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatGif#4aa3258a to nil")
	}
	if err := b.ConsumeID(ThumbnailFormatGifTypeID); err != nil {
		return fmt.Errorf("unable to decode thumbnailFormatGif#4aa3258a: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThumbnailFormatGif) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatGif#4aa3258a to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ThumbnailFormatGif) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatGif#4aa3258a as nil")
	}
	b.ObjStart()
	b.PutID("thumbnailFormatGif")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ThumbnailFormatGif) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatGif#4aa3258a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("thumbnailFormatGif"); err != nil {
				return fmt.Errorf("unable to decode thumbnailFormatGif#4aa3258a: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ThumbnailFormatMpeg4 represents TL type `thumbnailFormatMpeg4#109b57fe`.
type ThumbnailFormatMpeg4 struct {
}

// ThumbnailFormatMpeg4TypeID is TL type id of ThumbnailFormatMpeg4.
const ThumbnailFormatMpeg4TypeID = 0x109b57fe

// construct implements constructor of ThumbnailFormatClass.
func (t ThumbnailFormatMpeg4) construct() ThumbnailFormatClass { return &t }

// Ensuring interfaces in compile-time for ThumbnailFormatMpeg4.
var (
	_ bin.Encoder     = &ThumbnailFormatMpeg4{}
	_ bin.Decoder     = &ThumbnailFormatMpeg4{}
	_ bin.BareEncoder = &ThumbnailFormatMpeg4{}
	_ bin.BareDecoder = &ThumbnailFormatMpeg4{}

	_ ThumbnailFormatClass = &ThumbnailFormatMpeg4{}
)

func (t *ThumbnailFormatMpeg4) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThumbnailFormatMpeg4) String() string {
	if t == nil {
		return "ThumbnailFormatMpeg4(nil)"
	}
	type Alias ThumbnailFormatMpeg4
	return fmt.Sprintf("ThumbnailFormatMpeg4%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThumbnailFormatMpeg4) TypeID() uint32 {
	return ThumbnailFormatMpeg4TypeID
}

// TypeName returns name of type in TL schema.
func (*ThumbnailFormatMpeg4) TypeName() string {
	return "thumbnailFormatMpeg4"
}

// TypeInfo returns info about TL type.
func (t *ThumbnailFormatMpeg4) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "thumbnailFormatMpeg4",
		ID:   ThumbnailFormatMpeg4TypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThumbnailFormatMpeg4) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatMpeg4#109b57fe as nil")
	}
	b.PutID(ThumbnailFormatMpeg4TypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThumbnailFormatMpeg4) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatMpeg4#109b57fe as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ThumbnailFormatMpeg4) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatMpeg4#109b57fe to nil")
	}
	if err := b.ConsumeID(ThumbnailFormatMpeg4TypeID); err != nil {
		return fmt.Errorf("unable to decode thumbnailFormatMpeg4#109b57fe: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThumbnailFormatMpeg4) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatMpeg4#109b57fe to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ThumbnailFormatMpeg4) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatMpeg4#109b57fe as nil")
	}
	b.ObjStart()
	b.PutID("thumbnailFormatMpeg4")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ThumbnailFormatMpeg4) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatMpeg4#109b57fe to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("thumbnailFormatMpeg4"); err != nil {
				return fmt.Errorf("unable to decode thumbnailFormatMpeg4#109b57fe: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ThumbnailFormatPng represents TL type `thumbnailFormatPng#5e0697f5`.
type ThumbnailFormatPng struct {
}

// ThumbnailFormatPngTypeID is TL type id of ThumbnailFormatPng.
const ThumbnailFormatPngTypeID = 0x5e0697f5

// construct implements constructor of ThumbnailFormatClass.
func (t ThumbnailFormatPng) construct() ThumbnailFormatClass { return &t }

// Ensuring interfaces in compile-time for ThumbnailFormatPng.
var (
	_ bin.Encoder     = &ThumbnailFormatPng{}
	_ bin.Decoder     = &ThumbnailFormatPng{}
	_ bin.BareEncoder = &ThumbnailFormatPng{}
	_ bin.BareDecoder = &ThumbnailFormatPng{}

	_ ThumbnailFormatClass = &ThumbnailFormatPng{}
)

func (t *ThumbnailFormatPng) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThumbnailFormatPng) String() string {
	if t == nil {
		return "ThumbnailFormatPng(nil)"
	}
	type Alias ThumbnailFormatPng
	return fmt.Sprintf("ThumbnailFormatPng%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThumbnailFormatPng) TypeID() uint32 {
	return ThumbnailFormatPngTypeID
}

// TypeName returns name of type in TL schema.
func (*ThumbnailFormatPng) TypeName() string {
	return "thumbnailFormatPng"
}

// TypeInfo returns info about TL type.
func (t *ThumbnailFormatPng) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "thumbnailFormatPng",
		ID:   ThumbnailFormatPngTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThumbnailFormatPng) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatPng#5e0697f5 as nil")
	}
	b.PutID(ThumbnailFormatPngTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThumbnailFormatPng) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatPng#5e0697f5 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ThumbnailFormatPng) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatPng#5e0697f5 to nil")
	}
	if err := b.ConsumeID(ThumbnailFormatPngTypeID); err != nil {
		return fmt.Errorf("unable to decode thumbnailFormatPng#5e0697f5: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThumbnailFormatPng) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatPng#5e0697f5 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ThumbnailFormatPng) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatPng#5e0697f5 as nil")
	}
	b.ObjStart()
	b.PutID("thumbnailFormatPng")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ThumbnailFormatPng) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatPng#5e0697f5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("thumbnailFormatPng"); err != nil {
				return fmt.Errorf("unable to decode thumbnailFormatPng#5e0697f5: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ThumbnailFormatTgs represents TL type `thumbnailFormatTgs#4e694852`.
type ThumbnailFormatTgs struct {
}

// ThumbnailFormatTgsTypeID is TL type id of ThumbnailFormatTgs.
const ThumbnailFormatTgsTypeID = 0x4e694852

// construct implements constructor of ThumbnailFormatClass.
func (t ThumbnailFormatTgs) construct() ThumbnailFormatClass { return &t }

// Ensuring interfaces in compile-time for ThumbnailFormatTgs.
var (
	_ bin.Encoder     = &ThumbnailFormatTgs{}
	_ bin.Decoder     = &ThumbnailFormatTgs{}
	_ bin.BareEncoder = &ThumbnailFormatTgs{}
	_ bin.BareDecoder = &ThumbnailFormatTgs{}

	_ ThumbnailFormatClass = &ThumbnailFormatTgs{}
)

func (t *ThumbnailFormatTgs) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThumbnailFormatTgs) String() string {
	if t == nil {
		return "ThumbnailFormatTgs(nil)"
	}
	type Alias ThumbnailFormatTgs
	return fmt.Sprintf("ThumbnailFormatTgs%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThumbnailFormatTgs) TypeID() uint32 {
	return ThumbnailFormatTgsTypeID
}

// TypeName returns name of type in TL schema.
func (*ThumbnailFormatTgs) TypeName() string {
	return "thumbnailFormatTgs"
}

// TypeInfo returns info about TL type.
func (t *ThumbnailFormatTgs) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "thumbnailFormatTgs",
		ID:   ThumbnailFormatTgsTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThumbnailFormatTgs) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatTgs#4e694852 as nil")
	}
	b.PutID(ThumbnailFormatTgsTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThumbnailFormatTgs) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatTgs#4e694852 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ThumbnailFormatTgs) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatTgs#4e694852 to nil")
	}
	if err := b.ConsumeID(ThumbnailFormatTgsTypeID); err != nil {
		return fmt.Errorf("unable to decode thumbnailFormatTgs#4e694852: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThumbnailFormatTgs) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatTgs#4e694852 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ThumbnailFormatTgs) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatTgs#4e694852 as nil")
	}
	b.ObjStart()
	b.PutID("thumbnailFormatTgs")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ThumbnailFormatTgs) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatTgs#4e694852 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("thumbnailFormatTgs"); err != nil {
				return fmt.Errorf("unable to decode thumbnailFormatTgs#4e694852: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ThumbnailFormatWebm represents TL type `thumbnailFormatWebm#d8a7e727`.
type ThumbnailFormatWebm struct {
}

// ThumbnailFormatWebmTypeID is TL type id of ThumbnailFormatWebm.
const ThumbnailFormatWebmTypeID = 0xd8a7e727

// construct implements constructor of ThumbnailFormatClass.
func (t ThumbnailFormatWebm) construct() ThumbnailFormatClass { return &t }

// Ensuring interfaces in compile-time for ThumbnailFormatWebm.
var (
	_ bin.Encoder     = &ThumbnailFormatWebm{}
	_ bin.Decoder     = &ThumbnailFormatWebm{}
	_ bin.BareEncoder = &ThumbnailFormatWebm{}
	_ bin.BareDecoder = &ThumbnailFormatWebm{}

	_ ThumbnailFormatClass = &ThumbnailFormatWebm{}
)

func (t *ThumbnailFormatWebm) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThumbnailFormatWebm) String() string {
	if t == nil {
		return "ThumbnailFormatWebm(nil)"
	}
	type Alias ThumbnailFormatWebm
	return fmt.Sprintf("ThumbnailFormatWebm%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThumbnailFormatWebm) TypeID() uint32 {
	return ThumbnailFormatWebmTypeID
}

// TypeName returns name of type in TL schema.
func (*ThumbnailFormatWebm) TypeName() string {
	return "thumbnailFormatWebm"
}

// TypeInfo returns info about TL type.
func (t *ThumbnailFormatWebm) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "thumbnailFormatWebm",
		ID:   ThumbnailFormatWebmTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThumbnailFormatWebm) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatWebm#d8a7e727 as nil")
	}
	b.PutID(ThumbnailFormatWebmTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThumbnailFormatWebm) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatWebm#d8a7e727 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ThumbnailFormatWebm) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatWebm#d8a7e727 to nil")
	}
	if err := b.ConsumeID(ThumbnailFormatWebmTypeID); err != nil {
		return fmt.Errorf("unable to decode thumbnailFormatWebm#d8a7e727: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThumbnailFormatWebm) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatWebm#d8a7e727 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ThumbnailFormatWebm) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatWebm#d8a7e727 as nil")
	}
	b.ObjStart()
	b.PutID("thumbnailFormatWebm")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ThumbnailFormatWebm) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatWebm#d8a7e727 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("thumbnailFormatWebm"); err != nil {
				return fmt.Errorf("unable to decode thumbnailFormatWebm#d8a7e727: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ThumbnailFormatWebp represents TL type `thumbnailFormatWebp#fcce4c12`.
type ThumbnailFormatWebp struct {
}

// ThumbnailFormatWebpTypeID is TL type id of ThumbnailFormatWebp.
const ThumbnailFormatWebpTypeID = 0xfcce4c12

// construct implements constructor of ThumbnailFormatClass.
func (t ThumbnailFormatWebp) construct() ThumbnailFormatClass { return &t }

// Ensuring interfaces in compile-time for ThumbnailFormatWebp.
var (
	_ bin.Encoder     = &ThumbnailFormatWebp{}
	_ bin.Decoder     = &ThumbnailFormatWebp{}
	_ bin.BareEncoder = &ThumbnailFormatWebp{}
	_ bin.BareDecoder = &ThumbnailFormatWebp{}

	_ ThumbnailFormatClass = &ThumbnailFormatWebp{}
)

func (t *ThumbnailFormatWebp) Zero() bool {
	if t == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (t *ThumbnailFormatWebp) String() string {
	if t == nil {
		return "ThumbnailFormatWebp(nil)"
	}
	type Alias ThumbnailFormatWebp
	return fmt.Sprintf("ThumbnailFormatWebp%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ThumbnailFormatWebp) TypeID() uint32 {
	return ThumbnailFormatWebpTypeID
}

// TypeName returns name of type in TL schema.
func (*ThumbnailFormatWebp) TypeName() string {
	return "thumbnailFormatWebp"
}

// TypeInfo returns info about TL type.
func (t *ThumbnailFormatWebp) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "thumbnailFormatWebp",
		ID:   ThumbnailFormatWebpTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (t *ThumbnailFormatWebp) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatWebp#fcce4c12 as nil")
	}
	b.PutID(ThumbnailFormatWebpTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ThumbnailFormatWebp) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatWebp#fcce4c12 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *ThumbnailFormatWebp) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatWebp#fcce4c12 to nil")
	}
	if err := b.ConsumeID(ThumbnailFormatWebpTypeID); err != nil {
		return fmt.Errorf("unable to decode thumbnailFormatWebp#fcce4c12: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ThumbnailFormatWebp) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatWebp#fcce4c12 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ThumbnailFormatWebp) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode thumbnailFormatWebp#fcce4c12 as nil")
	}
	b.ObjStart()
	b.PutID("thumbnailFormatWebp")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ThumbnailFormatWebp) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode thumbnailFormatWebp#fcce4c12 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("thumbnailFormatWebp"); err != nil {
				return fmt.Errorf("unable to decode thumbnailFormatWebp#fcce4c12: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ThumbnailFormatClassName is schema name of ThumbnailFormatClass.
const ThumbnailFormatClassName = "ThumbnailFormat"

// ThumbnailFormatClass represents ThumbnailFormat generic type.
//
// Example:
//
//	g, err := tdapi.DecodeThumbnailFormat(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.ThumbnailFormatJpeg: // thumbnailFormatJpeg#d90c5488
//	case *tdapi.ThumbnailFormatGif: // thumbnailFormatGif#4aa3258a
//	case *tdapi.ThumbnailFormatMpeg4: // thumbnailFormatMpeg4#109b57fe
//	case *tdapi.ThumbnailFormatPng: // thumbnailFormatPng#5e0697f5
//	case *tdapi.ThumbnailFormatTgs: // thumbnailFormatTgs#4e694852
//	case *tdapi.ThumbnailFormatWebm: // thumbnailFormatWebm#d8a7e727
//	case *tdapi.ThumbnailFormatWebp: // thumbnailFormatWebp#fcce4c12
//	default: panic(v)
//	}
type ThumbnailFormatClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ThumbnailFormatClass

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

// DecodeThumbnailFormat implements binary de-serialization for ThumbnailFormatClass.
func DecodeThumbnailFormat(buf *bin.Buffer) (ThumbnailFormatClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ThumbnailFormatJpegTypeID:
		// Decoding thumbnailFormatJpeg#d90c5488.
		v := ThumbnailFormatJpeg{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case ThumbnailFormatGifTypeID:
		// Decoding thumbnailFormatGif#4aa3258a.
		v := ThumbnailFormatGif{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case ThumbnailFormatMpeg4TypeID:
		// Decoding thumbnailFormatMpeg4#109b57fe.
		v := ThumbnailFormatMpeg4{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case ThumbnailFormatPngTypeID:
		// Decoding thumbnailFormatPng#5e0697f5.
		v := ThumbnailFormatPng{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case ThumbnailFormatTgsTypeID:
		// Decoding thumbnailFormatTgs#4e694852.
		v := ThumbnailFormatTgs{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case ThumbnailFormatWebmTypeID:
		// Decoding thumbnailFormatWebm#d8a7e727.
		v := ThumbnailFormatWebm{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case ThumbnailFormatWebpTypeID:
		// Decoding thumbnailFormatWebp#fcce4c12.
		v := ThumbnailFormatWebp{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONThumbnailFormat implements binary de-serialization for ThumbnailFormatClass.
func DecodeTDLibJSONThumbnailFormat(buf tdjson.Decoder) (ThumbnailFormatClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "thumbnailFormatJpeg":
		// Decoding thumbnailFormatJpeg#d90c5488.
		v := ThumbnailFormatJpeg{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case "thumbnailFormatGif":
		// Decoding thumbnailFormatGif#4aa3258a.
		v := ThumbnailFormatGif{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case "thumbnailFormatMpeg4":
		// Decoding thumbnailFormatMpeg4#109b57fe.
		v := ThumbnailFormatMpeg4{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case "thumbnailFormatPng":
		// Decoding thumbnailFormatPng#5e0697f5.
		v := ThumbnailFormatPng{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case "thumbnailFormatTgs":
		// Decoding thumbnailFormatTgs#4e694852.
		v := ThumbnailFormatTgs{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case "thumbnailFormatWebm":
		// Decoding thumbnailFormatWebm#d8a7e727.
		v := ThumbnailFormatWebm{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	case "thumbnailFormatWebp":
		// Decoding thumbnailFormatWebp#fcce4c12.
		v := ThumbnailFormatWebp{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ThumbnailFormatClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ThumbnailFormat boxes the ThumbnailFormatClass providing a helper.
type ThumbnailFormatBox struct {
	ThumbnailFormat ThumbnailFormatClass
}

// Decode implements bin.Decoder for ThumbnailFormatBox.
func (b *ThumbnailFormatBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ThumbnailFormatBox to nil")
	}
	v, err := DecodeThumbnailFormat(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ThumbnailFormat = v
	return nil
}

// Encode implements bin.Encode for ThumbnailFormatBox.
func (b *ThumbnailFormatBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ThumbnailFormat == nil {
		return fmt.Errorf("unable to encode ThumbnailFormatClass as nil")
	}
	return b.ThumbnailFormat.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ThumbnailFormatBox.
func (b *ThumbnailFormatBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ThumbnailFormatBox to nil")
	}
	v, err := DecodeTDLibJSONThumbnailFormat(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ThumbnailFormat = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ThumbnailFormatBox.
func (b *ThumbnailFormatBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ThumbnailFormat == nil {
		return fmt.Errorf("unable to encode ThumbnailFormatClass as nil")
	}
	return b.ThumbnailFormat.EncodeTDLibJSON(buf)
}
