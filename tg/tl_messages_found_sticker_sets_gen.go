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

// MessagesFoundStickerSetsNotModified represents TL type `messages.foundStickerSetsNotModified#d54b65d`.
// No further results were found
//
// See https://core.telegram.org/constructor/messages.foundStickerSetsNotModified for reference.
type MessagesFoundStickerSetsNotModified struct {
}

// MessagesFoundStickerSetsNotModifiedTypeID is TL type id of MessagesFoundStickerSetsNotModified.
const MessagesFoundStickerSetsNotModifiedTypeID = 0xd54b65d

// construct implements constructor of MessagesFoundStickerSetsClass.
func (f MessagesFoundStickerSetsNotModified) construct() MessagesFoundStickerSetsClass { return &f }

// Ensuring interfaces in compile-time for MessagesFoundStickerSetsNotModified.
var (
	_ bin.Encoder     = &MessagesFoundStickerSetsNotModified{}
	_ bin.Decoder     = &MessagesFoundStickerSetsNotModified{}
	_ bin.BareEncoder = &MessagesFoundStickerSetsNotModified{}
	_ bin.BareDecoder = &MessagesFoundStickerSetsNotModified{}

	_ MessagesFoundStickerSetsClass = &MessagesFoundStickerSetsNotModified{}
)

func (f *MessagesFoundStickerSetsNotModified) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFoundStickerSetsNotModified) String() string {
	if f == nil {
		return "MessagesFoundStickerSetsNotModified(nil)"
	}
	type Alias MessagesFoundStickerSetsNotModified
	return fmt.Sprintf("MessagesFoundStickerSetsNotModified%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFoundStickerSetsNotModified) TypeID() uint32 {
	return MessagesFoundStickerSetsNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFoundStickerSetsNotModified) TypeName() string {
	return "messages.foundStickerSetsNotModified"
}

// TypeInfo returns info about TL type.
func (f *MessagesFoundStickerSetsNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.foundStickerSetsNotModified",
		ID:   MessagesFoundStickerSetsNotModifiedTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *MessagesFoundStickerSetsNotModified) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.foundStickerSetsNotModified#d54b65d as nil")
	}
	b.PutID(MessagesFoundStickerSetsNotModifiedTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesFoundStickerSetsNotModified) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.foundStickerSetsNotModified#d54b65d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesFoundStickerSetsNotModified) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.foundStickerSetsNotModified#d54b65d to nil")
	}
	if err := b.ConsumeID(MessagesFoundStickerSetsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.foundStickerSetsNotModified#d54b65d: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesFoundStickerSetsNotModified) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.foundStickerSetsNotModified#d54b65d to nil")
	}
	return nil
}

// MessagesFoundStickerSets represents TL type `messages.foundStickerSets#8af09dd2`.
// Found stickersets
//
// See https://core.telegram.org/constructor/messages.foundStickerSets for reference.
type MessagesFoundStickerSets struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
	// Found stickersets
	Sets []StickerSetCoveredClass
}

// MessagesFoundStickerSetsTypeID is TL type id of MessagesFoundStickerSets.
const MessagesFoundStickerSetsTypeID = 0x8af09dd2

// construct implements constructor of MessagesFoundStickerSetsClass.
func (f MessagesFoundStickerSets) construct() MessagesFoundStickerSetsClass { return &f }

// Ensuring interfaces in compile-time for MessagesFoundStickerSets.
var (
	_ bin.Encoder     = &MessagesFoundStickerSets{}
	_ bin.Decoder     = &MessagesFoundStickerSets{}
	_ bin.BareEncoder = &MessagesFoundStickerSets{}
	_ bin.BareDecoder = &MessagesFoundStickerSets{}

	_ MessagesFoundStickerSetsClass = &MessagesFoundStickerSets{}
)

func (f *MessagesFoundStickerSets) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.Hash == 0) {
		return false
	}
	if !(f.Sets == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFoundStickerSets) String() string {
	if f == nil {
		return "MessagesFoundStickerSets(nil)"
	}
	type Alias MessagesFoundStickerSets
	return fmt.Sprintf("MessagesFoundStickerSets%+v", Alias(*f))
}

// FillFrom fills MessagesFoundStickerSets from given interface.
func (f *MessagesFoundStickerSets) FillFrom(from interface {
	GetHash() (value int64)
	GetSets() (value []StickerSetCoveredClass)
}) {
	f.Hash = from.GetHash()
	f.Sets = from.GetSets()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFoundStickerSets) TypeID() uint32 {
	return MessagesFoundStickerSetsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFoundStickerSets) TypeName() string {
	return "messages.foundStickerSets"
}

// TypeInfo returns info about TL type.
func (f *MessagesFoundStickerSets) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.foundStickerSets",
		ID:   MessagesFoundStickerSetsTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Sets",
			SchemaName: "sets",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *MessagesFoundStickerSets) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.foundStickerSets#8af09dd2 as nil")
	}
	b.PutID(MessagesFoundStickerSetsTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesFoundStickerSets) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.foundStickerSets#8af09dd2 as nil")
	}
	b.PutLong(f.Hash)
	b.PutVectorHeader(len(f.Sets))
	for idx, v := range f.Sets {
		if v == nil {
			return fmt.Errorf("unable to encode messages.foundStickerSets#8af09dd2: field sets element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.foundStickerSets#8af09dd2: field sets element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesFoundStickerSets) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.foundStickerSets#8af09dd2 to nil")
	}
	if err := b.ConsumeID(MessagesFoundStickerSetsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.foundStickerSets#8af09dd2: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesFoundStickerSets) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.foundStickerSets#8af09dd2 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.foundStickerSets#8af09dd2: field hash: %w", err)
		}
		f.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.foundStickerSets#8af09dd2: field sets: %w", err)
		}

		if headerLen > 0 {
			f.Sets = make([]StickerSetCoveredClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeStickerSetCovered(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.foundStickerSets#8af09dd2: field sets: %w", err)
			}
			f.Sets = append(f.Sets, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (f *MessagesFoundStickerSets) GetHash() (value int64) {
	return f.Hash
}

// GetSets returns value of Sets field.
func (f *MessagesFoundStickerSets) GetSets() (value []StickerSetCoveredClass) {
	return f.Sets
}

// MapSets returns field Sets wrapped in StickerSetCoveredClassArray helper.
func (f *MessagesFoundStickerSets) MapSets() (value StickerSetCoveredClassArray) {
	return StickerSetCoveredClassArray(f.Sets)
}

// MessagesFoundStickerSetsClass represents messages.FoundStickerSets generic type.
//
// See https://core.telegram.org/type/messages.FoundStickerSets for reference.
//
// Example:
//  g, err := tg.DecodeMessagesFoundStickerSets(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesFoundStickerSetsNotModified: // messages.foundStickerSetsNotModified#d54b65d
//  case *tg.MessagesFoundStickerSets: // messages.foundStickerSets#8af09dd2
//  default: panic(v)
//  }
type MessagesFoundStickerSetsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesFoundStickerSetsClass

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

	// AsModified tries to map MessagesFoundStickerSetsClass to MessagesFoundStickerSets.
	AsModified() (*MessagesFoundStickerSets, bool)
}

// AsModified tries to map MessagesFoundStickerSetsNotModified to MessagesFoundStickerSets.
func (f *MessagesFoundStickerSetsNotModified) AsModified() (*MessagesFoundStickerSets, bool) {
	return nil, false
}

// AsModified tries to map MessagesFoundStickerSets to MessagesFoundStickerSets.
func (f *MessagesFoundStickerSets) AsModified() (*MessagesFoundStickerSets, bool) {
	return f, true
}

// DecodeMessagesFoundStickerSets implements binary de-serialization for MessagesFoundStickerSetsClass.
func DecodeMessagesFoundStickerSets(buf *bin.Buffer) (MessagesFoundStickerSetsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesFoundStickerSetsNotModifiedTypeID:
		// Decoding messages.foundStickerSetsNotModified#d54b65d.
		v := MessagesFoundStickerSetsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFoundStickerSetsClass: %w", err)
		}
		return &v, nil
	case MessagesFoundStickerSetsTypeID:
		// Decoding messages.foundStickerSets#8af09dd2.
		v := MessagesFoundStickerSets{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesFoundStickerSetsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesFoundStickerSetsClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesFoundStickerSets boxes the MessagesFoundStickerSetsClass providing a helper.
type MessagesFoundStickerSetsBox struct {
	FoundStickerSets MessagesFoundStickerSetsClass
}

// Decode implements bin.Decoder for MessagesFoundStickerSetsBox.
func (b *MessagesFoundStickerSetsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesFoundStickerSetsBox to nil")
	}
	v, err := DecodeMessagesFoundStickerSets(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FoundStickerSets = v
	return nil
}

// Encode implements bin.Encode for MessagesFoundStickerSetsBox.
func (b *MessagesFoundStickerSetsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FoundStickerSets == nil {
		return fmt.Errorf("unable to encode MessagesFoundStickerSetsClass as nil")
	}
	return b.FoundStickerSets.Encode(buf)
}
