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

// StickerTypeRegular represents TL type `stickerTypeRegular#35bc575`.
type StickerTypeRegular struct {
}

// StickerTypeRegularTypeID is TL type id of StickerTypeRegular.
const StickerTypeRegularTypeID = 0x35bc575

// construct implements constructor of StickerTypeClass.
func (s StickerTypeRegular) construct() StickerTypeClass { return &s }

// Ensuring interfaces in compile-time for StickerTypeRegular.
var (
	_ bin.Encoder     = &StickerTypeRegular{}
	_ bin.Decoder     = &StickerTypeRegular{}
	_ bin.BareEncoder = &StickerTypeRegular{}
	_ bin.BareDecoder = &StickerTypeRegular{}

	_ StickerTypeClass = &StickerTypeRegular{}
)

func (s *StickerTypeRegular) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerTypeRegular) String() string {
	if s == nil {
		return "StickerTypeRegular(nil)"
	}
	type Alias StickerTypeRegular
	return fmt.Sprintf("StickerTypeRegular%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickerTypeRegular) TypeID() uint32 {
	return StickerTypeRegularTypeID
}

// TypeName returns name of type in TL schema.
func (*StickerTypeRegular) TypeName() string {
	return "stickerTypeRegular"
}

// TypeInfo returns info about TL type.
func (s *StickerTypeRegular) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickerTypeRegular",
		ID:   StickerTypeRegularTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickerTypeRegular) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerTypeRegular#35bc575 as nil")
	}
	b.PutID(StickerTypeRegularTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickerTypeRegular) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerTypeRegular#35bc575 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StickerTypeRegular) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerTypeRegular#35bc575 to nil")
	}
	if err := b.ConsumeID(StickerTypeRegularTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerTypeRegular#35bc575: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickerTypeRegular) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerTypeRegular#35bc575 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StickerTypeRegular) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerTypeRegular#35bc575 as nil")
	}
	b.ObjStart()
	b.PutID("stickerTypeRegular")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StickerTypeRegular) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerTypeRegular#35bc575 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("stickerTypeRegular"); err != nil {
				return fmt.Errorf("unable to decode stickerTypeRegular#35bc575: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StickerTypeMask represents TL type `stickerTypeMask#96c63694`.
type StickerTypeMask struct {
}

// StickerTypeMaskTypeID is TL type id of StickerTypeMask.
const StickerTypeMaskTypeID = 0x96c63694

// construct implements constructor of StickerTypeClass.
func (s StickerTypeMask) construct() StickerTypeClass { return &s }

// Ensuring interfaces in compile-time for StickerTypeMask.
var (
	_ bin.Encoder     = &StickerTypeMask{}
	_ bin.Decoder     = &StickerTypeMask{}
	_ bin.BareEncoder = &StickerTypeMask{}
	_ bin.BareDecoder = &StickerTypeMask{}

	_ StickerTypeClass = &StickerTypeMask{}
)

func (s *StickerTypeMask) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerTypeMask) String() string {
	if s == nil {
		return "StickerTypeMask(nil)"
	}
	type Alias StickerTypeMask
	return fmt.Sprintf("StickerTypeMask%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickerTypeMask) TypeID() uint32 {
	return StickerTypeMaskTypeID
}

// TypeName returns name of type in TL schema.
func (*StickerTypeMask) TypeName() string {
	return "stickerTypeMask"
}

// TypeInfo returns info about TL type.
func (s *StickerTypeMask) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickerTypeMask",
		ID:   StickerTypeMaskTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickerTypeMask) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerTypeMask#96c63694 as nil")
	}
	b.PutID(StickerTypeMaskTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickerTypeMask) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerTypeMask#96c63694 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StickerTypeMask) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerTypeMask#96c63694 to nil")
	}
	if err := b.ConsumeID(StickerTypeMaskTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerTypeMask#96c63694: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickerTypeMask) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerTypeMask#96c63694 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StickerTypeMask) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerTypeMask#96c63694 as nil")
	}
	b.ObjStart()
	b.PutID("stickerTypeMask")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StickerTypeMask) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerTypeMask#96c63694 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("stickerTypeMask"); err != nil {
				return fmt.Errorf("unable to decode stickerTypeMask#96c63694: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StickerTypeCustomEmoji represents TL type `stickerTypeCustomEmoji#f8cd7787`.
type StickerTypeCustomEmoji struct {
}

// StickerTypeCustomEmojiTypeID is TL type id of StickerTypeCustomEmoji.
const StickerTypeCustomEmojiTypeID = 0xf8cd7787

// construct implements constructor of StickerTypeClass.
func (s StickerTypeCustomEmoji) construct() StickerTypeClass { return &s }

// Ensuring interfaces in compile-time for StickerTypeCustomEmoji.
var (
	_ bin.Encoder     = &StickerTypeCustomEmoji{}
	_ bin.Decoder     = &StickerTypeCustomEmoji{}
	_ bin.BareEncoder = &StickerTypeCustomEmoji{}
	_ bin.BareDecoder = &StickerTypeCustomEmoji{}

	_ StickerTypeClass = &StickerTypeCustomEmoji{}
)

func (s *StickerTypeCustomEmoji) Zero() bool {
	if s == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerTypeCustomEmoji) String() string {
	if s == nil {
		return "StickerTypeCustomEmoji(nil)"
	}
	type Alias StickerTypeCustomEmoji
	return fmt.Sprintf("StickerTypeCustomEmoji%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickerTypeCustomEmoji) TypeID() uint32 {
	return StickerTypeCustomEmojiTypeID
}

// TypeName returns name of type in TL schema.
func (*StickerTypeCustomEmoji) TypeName() string {
	return "stickerTypeCustomEmoji"
}

// TypeInfo returns info about TL type.
func (s *StickerTypeCustomEmoji) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickerTypeCustomEmoji",
		ID:   StickerTypeCustomEmojiTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickerTypeCustomEmoji) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerTypeCustomEmoji#f8cd7787 as nil")
	}
	b.PutID(StickerTypeCustomEmojiTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickerTypeCustomEmoji) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerTypeCustomEmoji#f8cd7787 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StickerTypeCustomEmoji) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerTypeCustomEmoji#f8cd7787 to nil")
	}
	if err := b.ConsumeID(StickerTypeCustomEmojiTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerTypeCustomEmoji#f8cd7787: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickerTypeCustomEmoji) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerTypeCustomEmoji#f8cd7787 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StickerTypeCustomEmoji) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerTypeCustomEmoji#f8cd7787 as nil")
	}
	b.ObjStart()
	b.PutID("stickerTypeCustomEmoji")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StickerTypeCustomEmoji) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerTypeCustomEmoji#f8cd7787 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("stickerTypeCustomEmoji"); err != nil {
				return fmt.Errorf("unable to decode stickerTypeCustomEmoji#f8cd7787: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// StickerTypeClassName is schema name of StickerTypeClass.
const StickerTypeClassName = "StickerType"

// StickerTypeClass represents StickerType generic type.
//
// Example:
//
//	g, err := tdapi.DecodeStickerType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.StickerTypeRegular: // stickerTypeRegular#35bc575
//	case *tdapi.StickerTypeMask: // stickerTypeMask#96c63694
//	case *tdapi.StickerTypeCustomEmoji: // stickerTypeCustomEmoji#f8cd7787
//	default: panic(v)
//	}
type StickerTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StickerTypeClass

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

// DecodeStickerType implements binary de-serialization for StickerTypeClass.
func DecodeStickerType(buf *bin.Buffer) (StickerTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StickerTypeRegularTypeID:
		// Decoding stickerTypeRegular#35bc575.
		v := StickerTypeRegular{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StickerTypeClass: %w", err)
		}
		return &v, nil
	case StickerTypeMaskTypeID:
		// Decoding stickerTypeMask#96c63694.
		v := StickerTypeMask{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StickerTypeClass: %w", err)
		}
		return &v, nil
	case StickerTypeCustomEmojiTypeID:
		// Decoding stickerTypeCustomEmoji#f8cd7787.
		v := StickerTypeCustomEmoji{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StickerTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StickerTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONStickerType implements binary de-serialization for StickerTypeClass.
func DecodeTDLibJSONStickerType(buf tdjson.Decoder) (StickerTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "stickerTypeRegular":
		// Decoding stickerTypeRegular#35bc575.
		v := StickerTypeRegular{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StickerTypeClass: %w", err)
		}
		return &v, nil
	case "stickerTypeMask":
		// Decoding stickerTypeMask#96c63694.
		v := StickerTypeMask{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StickerTypeClass: %w", err)
		}
		return &v, nil
	case "stickerTypeCustomEmoji":
		// Decoding stickerTypeCustomEmoji#f8cd7787.
		v := StickerTypeCustomEmoji{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StickerTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StickerTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// StickerType boxes the StickerTypeClass providing a helper.
type StickerTypeBox struct {
	StickerType StickerTypeClass
}

// Decode implements bin.Decoder for StickerTypeBox.
func (b *StickerTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StickerTypeBox to nil")
	}
	v, err := DecodeStickerType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StickerType = v
	return nil
}

// Encode implements bin.Encode for StickerTypeBox.
func (b *StickerTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StickerType == nil {
		return fmt.Errorf("unable to encode StickerTypeClass as nil")
	}
	return b.StickerType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for StickerTypeBox.
func (b *StickerTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode StickerTypeBox to nil")
	}
	v, err := DecodeTDLibJSONStickerType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StickerType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for StickerTypeBox.
func (b *StickerTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.StickerType == nil {
		return fmt.Errorf("unable to encode StickerTypeClass as nil")
	}
	return b.StickerType.EncodeTDLibJSON(buf)
}
