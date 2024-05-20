// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// BaseThemeClassic represents TL type `baseThemeClassic#c3a12462`.
// Classic theme
//
// See https://core.telegram.org/constructor/baseThemeClassic for reference.
type BaseThemeClassic struct {
}

// BaseThemeClassicTypeID is TL type id of BaseThemeClassic.
const BaseThemeClassicTypeID = 0xc3a12462

// construct implements constructor of BaseThemeClass.
func (b BaseThemeClassic) construct() BaseThemeClass { return &b }

// Ensuring interfaces in compile-time for BaseThemeClassic.
var (
	_ bin.Encoder     = &BaseThemeClassic{}
	_ bin.Decoder     = &BaseThemeClassic{}
	_ bin.BareEncoder = &BaseThemeClassic{}
	_ bin.BareDecoder = &BaseThemeClassic{}

	_ BaseThemeClass = &BaseThemeClassic{}
)

func (b *BaseThemeClassic) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BaseThemeClassic) String() string {
	if b == nil {
		return "BaseThemeClassic(nil)"
	}
	type Alias BaseThemeClassic
	return fmt.Sprintf("BaseThemeClassic%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BaseThemeClassic) TypeID() uint32 {
	return BaseThemeClassicTypeID
}

// TypeName returns name of type in TL schema.
func (*BaseThemeClassic) TypeName() string {
	return "baseThemeClassic"
}

// TypeInfo returns info about TL type.
func (b *BaseThemeClassic) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "baseThemeClassic",
		ID:   BaseThemeClassicTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BaseThemeClassic) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeClassic#c3a12462 as nil")
	}
	buf.PutID(BaseThemeClassicTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BaseThemeClassic) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeClassic#c3a12462 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BaseThemeClassic) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeClassic#c3a12462 to nil")
	}
	if err := buf.ConsumeID(BaseThemeClassicTypeID); err != nil {
		return fmt.Errorf("unable to decode baseThemeClassic#c3a12462: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BaseThemeClassic) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeClassic#c3a12462 to nil")
	}
	return nil
}

// BaseThemeDay represents TL type `baseThemeDay#fbd81688`.
// Day theme
//
// See https://core.telegram.org/constructor/baseThemeDay for reference.
type BaseThemeDay struct {
}

// BaseThemeDayTypeID is TL type id of BaseThemeDay.
const BaseThemeDayTypeID = 0xfbd81688

// construct implements constructor of BaseThemeClass.
func (b BaseThemeDay) construct() BaseThemeClass { return &b }

// Ensuring interfaces in compile-time for BaseThemeDay.
var (
	_ bin.Encoder     = &BaseThemeDay{}
	_ bin.Decoder     = &BaseThemeDay{}
	_ bin.BareEncoder = &BaseThemeDay{}
	_ bin.BareDecoder = &BaseThemeDay{}

	_ BaseThemeClass = &BaseThemeDay{}
)

func (b *BaseThemeDay) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BaseThemeDay) String() string {
	if b == nil {
		return "BaseThemeDay(nil)"
	}
	type Alias BaseThemeDay
	return fmt.Sprintf("BaseThemeDay%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BaseThemeDay) TypeID() uint32 {
	return BaseThemeDayTypeID
}

// TypeName returns name of type in TL schema.
func (*BaseThemeDay) TypeName() string {
	return "baseThemeDay"
}

// TypeInfo returns info about TL type.
func (b *BaseThemeDay) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "baseThemeDay",
		ID:   BaseThemeDayTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BaseThemeDay) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeDay#fbd81688 as nil")
	}
	buf.PutID(BaseThemeDayTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BaseThemeDay) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeDay#fbd81688 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BaseThemeDay) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeDay#fbd81688 to nil")
	}
	if err := buf.ConsumeID(BaseThemeDayTypeID); err != nil {
		return fmt.Errorf("unable to decode baseThemeDay#fbd81688: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BaseThemeDay) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeDay#fbd81688 to nil")
	}
	return nil
}

// BaseThemeNight represents TL type `baseThemeNight#b7b31ea8`.
// Night theme
//
// See https://core.telegram.org/constructor/baseThemeNight for reference.
type BaseThemeNight struct {
}

// BaseThemeNightTypeID is TL type id of BaseThemeNight.
const BaseThemeNightTypeID = 0xb7b31ea8

// construct implements constructor of BaseThemeClass.
func (b BaseThemeNight) construct() BaseThemeClass { return &b }

// Ensuring interfaces in compile-time for BaseThemeNight.
var (
	_ bin.Encoder     = &BaseThemeNight{}
	_ bin.Decoder     = &BaseThemeNight{}
	_ bin.BareEncoder = &BaseThemeNight{}
	_ bin.BareDecoder = &BaseThemeNight{}

	_ BaseThemeClass = &BaseThemeNight{}
)

func (b *BaseThemeNight) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BaseThemeNight) String() string {
	if b == nil {
		return "BaseThemeNight(nil)"
	}
	type Alias BaseThemeNight
	return fmt.Sprintf("BaseThemeNight%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BaseThemeNight) TypeID() uint32 {
	return BaseThemeNightTypeID
}

// TypeName returns name of type in TL schema.
func (*BaseThemeNight) TypeName() string {
	return "baseThemeNight"
}

// TypeInfo returns info about TL type.
func (b *BaseThemeNight) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "baseThemeNight",
		ID:   BaseThemeNightTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BaseThemeNight) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeNight#b7b31ea8 as nil")
	}
	buf.PutID(BaseThemeNightTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BaseThemeNight) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeNight#b7b31ea8 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BaseThemeNight) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeNight#b7b31ea8 to nil")
	}
	if err := buf.ConsumeID(BaseThemeNightTypeID); err != nil {
		return fmt.Errorf("unable to decode baseThemeNight#b7b31ea8: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BaseThemeNight) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeNight#b7b31ea8 to nil")
	}
	return nil
}

// BaseThemeTinted represents TL type `baseThemeTinted#6d5f77ee`.
// Tinted theme
//
// See https://core.telegram.org/constructor/baseThemeTinted for reference.
type BaseThemeTinted struct {
}

// BaseThemeTintedTypeID is TL type id of BaseThemeTinted.
const BaseThemeTintedTypeID = 0x6d5f77ee

// construct implements constructor of BaseThemeClass.
func (b BaseThemeTinted) construct() BaseThemeClass { return &b }

// Ensuring interfaces in compile-time for BaseThemeTinted.
var (
	_ bin.Encoder     = &BaseThemeTinted{}
	_ bin.Decoder     = &BaseThemeTinted{}
	_ bin.BareEncoder = &BaseThemeTinted{}
	_ bin.BareDecoder = &BaseThemeTinted{}

	_ BaseThemeClass = &BaseThemeTinted{}
)

func (b *BaseThemeTinted) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BaseThemeTinted) String() string {
	if b == nil {
		return "BaseThemeTinted(nil)"
	}
	type Alias BaseThemeTinted
	return fmt.Sprintf("BaseThemeTinted%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BaseThemeTinted) TypeID() uint32 {
	return BaseThemeTintedTypeID
}

// TypeName returns name of type in TL schema.
func (*BaseThemeTinted) TypeName() string {
	return "baseThemeTinted"
}

// TypeInfo returns info about TL type.
func (b *BaseThemeTinted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "baseThemeTinted",
		ID:   BaseThemeTintedTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BaseThemeTinted) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeTinted#6d5f77ee as nil")
	}
	buf.PutID(BaseThemeTintedTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BaseThemeTinted) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeTinted#6d5f77ee as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BaseThemeTinted) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeTinted#6d5f77ee to nil")
	}
	if err := buf.ConsumeID(BaseThemeTintedTypeID); err != nil {
		return fmt.Errorf("unable to decode baseThemeTinted#6d5f77ee: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BaseThemeTinted) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeTinted#6d5f77ee to nil")
	}
	return nil
}

// BaseThemeArctic represents TL type `baseThemeArctic#5b11125a`.
// Arctic theme
//
// See https://core.telegram.org/constructor/baseThemeArctic for reference.
type BaseThemeArctic struct {
}

// BaseThemeArcticTypeID is TL type id of BaseThemeArctic.
const BaseThemeArcticTypeID = 0x5b11125a

// construct implements constructor of BaseThemeClass.
func (b BaseThemeArctic) construct() BaseThemeClass { return &b }

// Ensuring interfaces in compile-time for BaseThemeArctic.
var (
	_ bin.Encoder     = &BaseThemeArctic{}
	_ bin.Decoder     = &BaseThemeArctic{}
	_ bin.BareEncoder = &BaseThemeArctic{}
	_ bin.BareDecoder = &BaseThemeArctic{}

	_ BaseThemeClass = &BaseThemeArctic{}
)

func (b *BaseThemeArctic) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BaseThemeArctic) String() string {
	if b == nil {
		return "BaseThemeArctic(nil)"
	}
	type Alias BaseThemeArctic
	return fmt.Sprintf("BaseThemeArctic%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BaseThemeArctic) TypeID() uint32 {
	return BaseThemeArcticTypeID
}

// TypeName returns name of type in TL schema.
func (*BaseThemeArctic) TypeName() string {
	return "baseThemeArctic"
}

// TypeInfo returns info about TL type.
func (b *BaseThemeArctic) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "baseThemeArctic",
		ID:   BaseThemeArcticTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BaseThemeArctic) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeArctic#5b11125a as nil")
	}
	buf.PutID(BaseThemeArcticTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BaseThemeArctic) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode baseThemeArctic#5b11125a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BaseThemeArctic) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeArctic#5b11125a to nil")
	}
	if err := buf.ConsumeID(BaseThemeArcticTypeID); err != nil {
		return fmt.Errorf("unable to decode baseThemeArctic#5b11125a: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BaseThemeArctic) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode baseThemeArctic#5b11125a to nil")
	}
	return nil
}

// BaseThemeClassName is schema name of BaseThemeClass.
const BaseThemeClassName = "BaseTheme"

// BaseThemeClass represents BaseTheme generic type.
//
// See https://core.telegram.org/type/BaseTheme for reference.
//
// Example:
//
//	g, err := tg.DecodeBaseTheme(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.BaseThemeClassic: // baseThemeClassic#c3a12462
//	case *tg.BaseThemeDay: // baseThemeDay#fbd81688
//	case *tg.BaseThemeNight: // baseThemeNight#b7b31ea8
//	case *tg.BaseThemeTinted: // baseThemeTinted#6d5f77ee
//	case *tg.BaseThemeArctic: // baseThemeArctic#5b11125a
//	default: panic(v)
//	}
type BaseThemeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() BaseThemeClass

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
}

// DecodeBaseTheme implements binary de-serialization for BaseThemeClass.
func DecodeBaseTheme(buf *bin.Buffer) (BaseThemeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BaseThemeClassicTypeID:
		// Decoding baseThemeClassic#c3a12462.
		v := BaseThemeClassic{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", err)
		}
		return &v, nil
	case BaseThemeDayTypeID:
		// Decoding baseThemeDay#fbd81688.
		v := BaseThemeDay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", err)
		}
		return &v, nil
	case BaseThemeNightTypeID:
		// Decoding baseThemeNight#b7b31ea8.
		v := BaseThemeNight{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", err)
		}
		return &v, nil
	case BaseThemeTintedTypeID:
		// Decoding baseThemeTinted#6d5f77ee.
		v := BaseThemeTinted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", err)
		}
		return &v, nil
	case BaseThemeArcticTypeID:
		// Decoding baseThemeArctic#5b11125a.
		v := BaseThemeArctic{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BaseThemeClass: %w", bin.NewUnexpectedID(id))
	}
}

// BaseTheme boxes the BaseThemeClass providing a helper.
type BaseThemeBox struct {
	BaseTheme BaseThemeClass
}

// Decode implements bin.Decoder for BaseThemeBox.
func (b *BaseThemeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BaseThemeBox to nil")
	}
	v, err := DecodeBaseTheme(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BaseTheme = v
	return nil
}

// Encode implements bin.Encode for BaseThemeBox.
func (b *BaseThemeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BaseTheme == nil {
		return fmt.Errorf("unable to encode BaseThemeClass as nil")
	}
	return b.BaseTheme.Encode(buf)
}
