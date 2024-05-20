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

// MessagesReactionsNotModified represents TL type `messages.reactionsNotModified#b06fdbdf`.
// The server-side list of message reactions¹ hasn't changed
//
// Links:
//  1. https://core.telegram.org/api/reactions
//
// See https://core.telegram.org/constructor/messages.reactionsNotModified for reference.
type MessagesReactionsNotModified struct {
}

// MessagesReactionsNotModifiedTypeID is TL type id of MessagesReactionsNotModified.
const MessagesReactionsNotModifiedTypeID = 0xb06fdbdf

// construct implements constructor of MessagesReactionsClass.
func (r MessagesReactionsNotModified) construct() MessagesReactionsClass { return &r }

// Ensuring interfaces in compile-time for MessagesReactionsNotModified.
var (
	_ bin.Encoder     = &MessagesReactionsNotModified{}
	_ bin.Decoder     = &MessagesReactionsNotModified{}
	_ bin.BareEncoder = &MessagesReactionsNotModified{}
	_ bin.BareDecoder = &MessagesReactionsNotModified{}

	_ MessagesReactionsClass = &MessagesReactionsNotModified{}
)

func (r *MessagesReactionsNotModified) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReactionsNotModified) String() string {
	if r == nil {
		return "MessagesReactionsNotModified(nil)"
	}
	type Alias MessagesReactionsNotModified
	return fmt.Sprintf("MessagesReactionsNotModified%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReactionsNotModified) TypeID() uint32 {
	return MessagesReactionsNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReactionsNotModified) TypeName() string {
	return "messages.reactionsNotModified"
}

// TypeInfo returns info about TL type.
func (r *MessagesReactionsNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.reactionsNotModified",
		ID:   MessagesReactionsNotModifiedTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReactionsNotModified) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reactionsNotModified#b06fdbdf as nil")
	}
	b.PutID(MessagesReactionsNotModifiedTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReactionsNotModified) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reactionsNotModified#b06fdbdf as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReactionsNotModified) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reactionsNotModified#b06fdbdf to nil")
	}
	if err := b.ConsumeID(MessagesReactionsNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reactionsNotModified#b06fdbdf: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReactionsNotModified) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reactionsNotModified#b06fdbdf to nil")
	}
	return nil
}

// MessagesReactions represents TL type `messages.reactions#eafdf716`.
// List of message reactions¹
//
// Links:
//  1. https://core.telegram.org/api/reactions
//
// See https://core.telegram.org/constructor/messages.reactions for reference.
type MessagesReactions struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
	// Reactions
	Reactions []ReactionClass
}

// MessagesReactionsTypeID is TL type id of MessagesReactions.
const MessagesReactionsTypeID = 0xeafdf716

// construct implements constructor of MessagesReactionsClass.
func (r MessagesReactions) construct() MessagesReactionsClass { return &r }

// Ensuring interfaces in compile-time for MessagesReactions.
var (
	_ bin.Encoder     = &MessagesReactions{}
	_ bin.Decoder     = &MessagesReactions{}
	_ bin.BareEncoder = &MessagesReactions{}
	_ bin.BareDecoder = &MessagesReactions{}

	_ MessagesReactionsClass = &MessagesReactions{}
)

func (r *MessagesReactions) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Hash == 0) {
		return false
	}
	if !(r.Reactions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReactions) String() string {
	if r == nil {
		return "MessagesReactions(nil)"
	}
	type Alias MessagesReactions
	return fmt.Sprintf("MessagesReactions%+v", Alias(*r))
}

// FillFrom fills MessagesReactions from given interface.
func (r *MessagesReactions) FillFrom(from interface {
	GetHash() (value int64)
	GetReactions() (value []ReactionClass)
}) {
	r.Hash = from.GetHash()
	r.Reactions = from.GetReactions()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReactions) TypeID() uint32 {
	return MessagesReactionsTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReactions) TypeName() string {
	return "messages.reactions"
}

// TypeInfo returns info about TL type.
func (r *MessagesReactions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.reactions",
		ID:   MessagesReactionsTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Reactions",
			SchemaName: "reactions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReactions) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reactions#eafdf716 as nil")
	}
	b.PutID(MessagesReactionsTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReactions) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.reactions#eafdf716 as nil")
	}
	b.PutLong(r.Hash)
	b.PutVectorHeader(len(r.Reactions))
	for idx, v := range r.Reactions {
		if v == nil {
			return fmt.Errorf("unable to encode messages.reactions#eafdf716: field reactions element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.reactions#eafdf716: field reactions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReactions) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reactions#eafdf716 to nil")
	}
	if err := b.ConsumeID(MessagesReactionsTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.reactions#eafdf716: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReactions) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.reactions#eafdf716 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.reactions#eafdf716: field hash: %w", err)
		}
		r.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.reactions#eafdf716: field reactions: %w", err)
		}

		if headerLen > 0 {
			r.Reactions = make([]ReactionClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeReaction(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.reactions#eafdf716: field reactions: %w", err)
			}
			r.Reactions = append(r.Reactions, value)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (r *MessagesReactions) GetHash() (value int64) {
	if r == nil {
		return
	}
	return r.Hash
}

// GetReactions returns value of Reactions field.
func (r *MessagesReactions) GetReactions() (value []ReactionClass) {
	if r == nil {
		return
	}
	return r.Reactions
}

// MapReactions returns field Reactions wrapped in ReactionClassArray helper.
func (r *MessagesReactions) MapReactions() (value ReactionClassArray) {
	return ReactionClassArray(r.Reactions)
}

// MessagesReactionsClassName is schema name of MessagesReactionsClass.
const MessagesReactionsClassName = "messages.Reactions"

// MessagesReactionsClass represents messages.Reactions generic type.
//
// See https://core.telegram.org/type/messages.Reactions for reference.
//
// Example:
//
//	g, err := tg.DecodeMessagesReactions(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.MessagesReactionsNotModified: // messages.reactionsNotModified#b06fdbdf
//	case *tg.MessagesReactions: // messages.reactions#eafdf716
//	default: panic(v)
//	}
type MessagesReactionsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesReactionsClass

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

	// AsModified tries to map MessagesReactionsClass to MessagesReactions.
	AsModified() (*MessagesReactions, bool)
}

// AsModified tries to map MessagesReactionsNotModified to MessagesReactions.
func (r *MessagesReactionsNotModified) AsModified() (*MessagesReactions, bool) {
	return nil, false
}

// AsModified tries to map MessagesReactions to MessagesReactions.
func (r *MessagesReactions) AsModified() (*MessagesReactions, bool) {
	return r, true
}

// DecodeMessagesReactions implements binary de-serialization for MessagesReactionsClass.
func DecodeMessagesReactions(buf *bin.Buffer) (MessagesReactionsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesReactionsNotModifiedTypeID:
		// Decoding messages.reactionsNotModified#b06fdbdf.
		v := MessagesReactionsNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesReactionsClass: %w", err)
		}
		return &v, nil
	case MessagesReactionsTypeID:
		// Decoding messages.reactions#eafdf716.
		v := MessagesReactions{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesReactionsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesReactionsClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesReactions boxes the MessagesReactionsClass providing a helper.
type MessagesReactionsBox struct {
	Reactions MessagesReactionsClass
}

// Decode implements bin.Decoder for MessagesReactionsBox.
func (b *MessagesReactionsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesReactionsBox to nil")
	}
	v, err := DecodeMessagesReactions(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Reactions = v
	return nil
}

// Encode implements bin.Encode for MessagesReactionsBox.
func (b *MessagesReactionsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Reactions == nil {
		return fmt.Errorf("unable to encode MessagesReactionsClass as nil")
	}
	return b.Reactions.Encode(buf)
}
