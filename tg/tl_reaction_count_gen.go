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

// ReactionCount represents TL type `reactionCount#a3d1cb80`.
// Reactions
//
// See https://core.telegram.org/constructor/reactionCount for reference.
type ReactionCount struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, indicates that the current user also sent this reaction. The integer value
	// indicates when was the reaction added: the bigger the value, the newer the reaction.
	//
	// Use SetChosenOrder and GetChosenOrder helpers.
	ChosenOrder int
	// The reaction.
	Reaction ReactionClass
	// Number of users that reacted with this emoji.
	Count int
}

// ReactionCountTypeID is TL type id of ReactionCount.
const ReactionCountTypeID = 0xa3d1cb80

// Ensuring interfaces in compile-time for ReactionCount.
var (
	_ bin.Encoder     = &ReactionCount{}
	_ bin.Decoder     = &ReactionCount{}
	_ bin.BareEncoder = &ReactionCount{}
	_ bin.BareDecoder = &ReactionCount{}
)

func (r *ReactionCount) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.ChosenOrder == 0) {
		return false
	}
	if !(r.Reaction == nil) {
		return false
	}
	if !(r.Count == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionCount) String() string {
	if r == nil {
		return "ReactionCount(nil)"
	}
	type Alias ReactionCount
	return fmt.Sprintf("ReactionCount%+v", Alias(*r))
}

// FillFrom fills ReactionCount from given interface.
func (r *ReactionCount) FillFrom(from interface {
	GetChosenOrder() (value int, ok bool)
	GetReaction() (value ReactionClass)
	GetCount() (value int)
}) {
	if val, ok := from.GetChosenOrder(); ok {
		r.ChosenOrder = val
	}

	r.Reaction = from.GetReaction()
	r.Count = from.GetCount()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionCount) TypeID() uint32 {
	return ReactionCountTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionCount) TypeName() string {
	return "reactionCount"
}

// TypeInfo returns info about TL type.
func (r *ReactionCount) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionCount",
		ID:   ReactionCountTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChosenOrder",
			SchemaName: "chosen_order",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "Reaction",
			SchemaName: "reaction",
		},
		{
			Name:       "Count",
			SchemaName: "count",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *ReactionCount) SetFlags() {
	if !(r.ChosenOrder == 0) {
		r.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (r *ReactionCount) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionCount#a3d1cb80 as nil")
	}
	b.PutID(ReactionCountTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionCount) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionCount#a3d1cb80 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reactionCount#a3d1cb80: field flags: %w", err)
	}
	if r.Flags.Has(0) {
		b.PutInt(r.ChosenOrder)
	}
	if r.Reaction == nil {
		return fmt.Errorf("unable to encode reactionCount#a3d1cb80: field reaction is nil")
	}
	if err := r.Reaction.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reactionCount#a3d1cb80: field reaction: %w", err)
	}
	b.PutInt(r.Count)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionCount) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionCount#a3d1cb80 to nil")
	}
	if err := b.ConsumeID(ReactionCountTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionCount#a3d1cb80: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionCount) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionCount#a3d1cb80 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode reactionCount#a3d1cb80: field flags: %w", err)
		}
	}
	if r.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode reactionCount#a3d1cb80: field chosen_order: %w", err)
		}
		r.ChosenOrder = value
	}
	{
		value, err := DecodeReaction(b)
		if err != nil {
			return fmt.Errorf("unable to decode reactionCount#a3d1cb80: field reaction: %w", err)
		}
		r.Reaction = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode reactionCount#a3d1cb80: field count: %w", err)
		}
		r.Count = value
	}
	return nil
}

// SetChosenOrder sets value of ChosenOrder conditional field.
func (r *ReactionCount) SetChosenOrder(value int) {
	r.Flags.Set(0)
	r.ChosenOrder = value
}

// GetChosenOrder returns value of ChosenOrder conditional field and
// boolean which is true if field was set.
func (r *ReactionCount) GetChosenOrder() (value int, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.ChosenOrder, true
}

// GetReaction returns value of Reaction field.
func (r *ReactionCount) GetReaction() (value ReactionClass) {
	if r == nil {
		return
	}
	return r.Reaction
}

// GetCount returns value of Count field.
func (r *ReactionCount) GetCount() (value int) {
	if r == nil {
		return
	}
	return r.Count
}
