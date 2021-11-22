// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/jsontd"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
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
	_ = jsontd.Encoder{}
)

// InputWallPaper represents TL type `inputWallPaper#e630b979`.
// Wallpaper
//
// See https://core.telegram.org/constructor/inputWallPaper for reference.
type InputWallPaper struct {
	// Wallpaper ID
	ID int64
	// Access hash
	AccessHash int64
}

// InputWallPaperTypeID is TL type id of InputWallPaper.
const InputWallPaperTypeID = 0xe630b979

// construct implements constructor of InputWallPaperClass.
func (i InputWallPaper) construct() InputWallPaperClass { return &i }

// Ensuring interfaces in compile-time for InputWallPaper.
var (
	_ bin.Encoder     = &InputWallPaper{}
	_ bin.Decoder     = &InputWallPaper{}
	_ bin.BareEncoder = &InputWallPaper{}
	_ bin.BareDecoder = &InputWallPaper{}

	_ InputWallPaperClass = &InputWallPaper{}
)

func (i *InputWallPaper) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputWallPaper) String() string {
	if i == nil {
		return "InputWallPaper(nil)"
	}
	type Alias InputWallPaper
	return fmt.Sprintf("InputWallPaper%+v", Alias(*i))
}

// FillFrom fills InputWallPaper from given interface.
func (i *InputWallPaper) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputWallPaper) TypeID() uint32 {
	return InputWallPaperTypeID
}

// TypeName returns name of type in TL schema.
func (*InputWallPaper) TypeName() string {
	return "inputWallPaper"
}

// TypeInfo returns info about TL type.
func (i *InputWallPaper) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputWallPaper",
		ID:   InputWallPaperTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputWallPaper) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWallPaper#e630b979 as nil")
	}
	b.PutID(InputWallPaperTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputWallPaper) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWallPaper#e630b979 as nil")
	}
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputWallPaper) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWallPaper#e630b979 to nil")
	}
	if err := b.ConsumeID(InputWallPaperTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWallPaper#e630b979: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputWallPaper) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWallPaper#e630b979 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputWallPaper#e630b979: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputWallPaper#e630b979: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputWallPaper) GetID() (value int64) {
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputWallPaper) GetAccessHash() (value int64) {
	return i.AccessHash
}

// InputWallPaperSlug represents TL type `inputWallPaperSlug#72091c80`.
// Wallpaper by slug (a unique ID)
//
// See https://core.telegram.org/constructor/inputWallPaperSlug for reference.
type InputWallPaperSlug struct {
	// Unique wallpaper ID
	Slug string
}

// InputWallPaperSlugTypeID is TL type id of InputWallPaperSlug.
const InputWallPaperSlugTypeID = 0x72091c80

// construct implements constructor of InputWallPaperClass.
func (i InputWallPaperSlug) construct() InputWallPaperClass { return &i }

// Ensuring interfaces in compile-time for InputWallPaperSlug.
var (
	_ bin.Encoder     = &InputWallPaperSlug{}
	_ bin.Decoder     = &InputWallPaperSlug{}
	_ bin.BareEncoder = &InputWallPaperSlug{}
	_ bin.BareDecoder = &InputWallPaperSlug{}

	_ InputWallPaperClass = &InputWallPaperSlug{}
)

func (i *InputWallPaperSlug) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Slug == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputWallPaperSlug) String() string {
	if i == nil {
		return "InputWallPaperSlug(nil)"
	}
	type Alias InputWallPaperSlug
	return fmt.Sprintf("InputWallPaperSlug%+v", Alias(*i))
}

// FillFrom fills InputWallPaperSlug from given interface.
func (i *InputWallPaperSlug) FillFrom(from interface {
	GetSlug() (value string)
}) {
	i.Slug = from.GetSlug()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputWallPaperSlug) TypeID() uint32 {
	return InputWallPaperSlugTypeID
}

// TypeName returns name of type in TL schema.
func (*InputWallPaperSlug) TypeName() string {
	return "inputWallPaperSlug"
}

// TypeInfo returns info about TL type.
func (i *InputWallPaperSlug) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputWallPaperSlug",
		ID:   InputWallPaperSlugTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Slug",
			SchemaName: "slug",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputWallPaperSlug) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWallPaperSlug#72091c80 as nil")
	}
	b.PutID(InputWallPaperSlugTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputWallPaperSlug) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWallPaperSlug#72091c80 as nil")
	}
	b.PutString(i.Slug)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputWallPaperSlug) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWallPaperSlug#72091c80 to nil")
	}
	if err := b.ConsumeID(InputWallPaperSlugTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWallPaperSlug#72091c80: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputWallPaperSlug) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWallPaperSlug#72091c80 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputWallPaperSlug#72091c80: field slug: %w", err)
		}
		i.Slug = value
	}
	return nil
}

// GetSlug returns value of Slug field.
func (i *InputWallPaperSlug) GetSlug() (value string) {
	return i.Slug
}

// InputWallPaperNoFile represents TL type `inputWallPaperNoFile#967a462e`.
// Wallpaper with no file
//
// See https://core.telegram.org/constructor/inputWallPaperNoFile for reference.
type InputWallPaperNoFile struct {
	// ID field of InputWallPaperNoFile.
	ID int64
}

// InputWallPaperNoFileTypeID is TL type id of InputWallPaperNoFile.
const InputWallPaperNoFileTypeID = 0x967a462e

// construct implements constructor of InputWallPaperClass.
func (i InputWallPaperNoFile) construct() InputWallPaperClass { return &i }

// Ensuring interfaces in compile-time for InputWallPaperNoFile.
var (
	_ bin.Encoder     = &InputWallPaperNoFile{}
	_ bin.Decoder     = &InputWallPaperNoFile{}
	_ bin.BareEncoder = &InputWallPaperNoFile{}
	_ bin.BareDecoder = &InputWallPaperNoFile{}

	_ InputWallPaperClass = &InputWallPaperNoFile{}
)

func (i *InputWallPaperNoFile) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputWallPaperNoFile) String() string {
	if i == nil {
		return "InputWallPaperNoFile(nil)"
	}
	type Alias InputWallPaperNoFile
	return fmt.Sprintf("InputWallPaperNoFile%+v", Alias(*i))
}

// FillFrom fills InputWallPaperNoFile from given interface.
func (i *InputWallPaperNoFile) FillFrom(from interface {
	GetID() (value int64)
}) {
	i.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputWallPaperNoFile) TypeID() uint32 {
	return InputWallPaperNoFileTypeID
}

// TypeName returns name of type in TL schema.
func (*InputWallPaperNoFile) TypeName() string {
	return "inputWallPaperNoFile"
}

// TypeInfo returns info about TL type.
func (i *InputWallPaperNoFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputWallPaperNoFile",
		ID:   InputWallPaperNoFileTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputWallPaperNoFile) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWallPaperNoFile#967a462e as nil")
	}
	b.PutID(InputWallPaperNoFileTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputWallPaperNoFile) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWallPaperNoFile#967a462e as nil")
	}
	b.PutLong(i.ID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputWallPaperNoFile) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWallPaperNoFile#967a462e to nil")
	}
	if err := b.ConsumeID(InputWallPaperNoFileTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWallPaperNoFile#967a462e: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputWallPaperNoFile) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWallPaperNoFile#967a462e to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputWallPaperNoFile#967a462e: field id: %w", err)
		}
		i.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputWallPaperNoFile) GetID() (value int64) {
	return i.ID
}

// InputWallPaperClass represents InputWallPaper generic type.
//
// See https://core.telegram.org/type/InputWallPaper for reference.
//
// Example:
//  g, err := tg.DecodeInputWallPaper(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputWallPaper: // inputWallPaper#e630b979
//  case *tg.InputWallPaperSlug: // inputWallPaperSlug#72091c80
//  case *tg.InputWallPaperNoFile: // inputWallPaperNoFile#967a462e
//  default: panic(v)
//  }
type InputWallPaperClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputWallPaperClass

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

// DecodeInputWallPaper implements binary de-serialization for InputWallPaperClass.
func DecodeInputWallPaper(buf *bin.Buffer) (InputWallPaperClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputWallPaperTypeID:
		// Decoding inputWallPaper#e630b979.
		v := InputWallPaper{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputWallPaperClass: %w", err)
		}
		return &v, nil
	case InputWallPaperSlugTypeID:
		// Decoding inputWallPaperSlug#72091c80.
		v := InputWallPaperSlug{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputWallPaperClass: %w", err)
		}
		return &v, nil
	case InputWallPaperNoFileTypeID:
		// Decoding inputWallPaperNoFile#967a462e.
		v := InputWallPaperNoFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputWallPaperClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputWallPaperClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputWallPaper boxes the InputWallPaperClass providing a helper.
type InputWallPaperBox struct {
	InputWallPaper InputWallPaperClass
}

// Decode implements bin.Decoder for InputWallPaperBox.
func (b *InputWallPaperBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputWallPaperBox to nil")
	}
	v, err := DecodeInputWallPaper(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputWallPaper = v
	return nil
}

// Encode implements bin.Encode for InputWallPaperBox.
func (b *InputWallPaperBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputWallPaper == nil {
		return fmt.Errorf("unable to encode InputWallPaperClass as nil")
	}
	return b.InputWallPaper.Encode(buf)
}
