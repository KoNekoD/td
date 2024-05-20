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

// GroupCallID represents TL type `groupCallId#14e4bb45`.
type GroupCallID struct {
	// Group call identifier
	ID int32
}

// GroupCallIDTypeID is TL type id of GroupCallID.
const GroupCallIDTypeID = 0x14e4bb45

// Ensuring interfaces in compile-time for GroupCallID.
var (
	_ bin.Encoder     = &GroupCallID{}
	_ bin.Decoder     = &GroupCallID{}
	_ bin.BareEncoder = &GroupCallID{}
	_ bin.BareDecoder = &GroupCallID{}
)

func (g *GroupCallID) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GroupCallID) String() string {
	if g == nil {
		return "GroupCallID(nil)"
	}
	type Alias GroupCallID
	return fmt.Sprintf("GroupCallID%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GroupCallID) TypeID() uint32 {
	return GroupCallIDTypeID
}

// TypeName returns name of type in TL schema.
func (*GroupCallID) TypeName() string {
	return "groupCallId"
}

// TypeInfo returns info about TL type.
func (g *GroupCallID) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "groupCallId",
		ID:   GroupCallIDTypeID,
	}
	if g == nil {
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
func (g *GroupCallID) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallId#14e4bb45 as nil")
	}
	b.PutID(GroupCallIDTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GroupCallID) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallId#14e4bb45 as nil")
	}
	b.PutInt32(g.ID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GroupCallID) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallId#14e4bb45 to nil")
	}
	if err := b.ConsumeID(GroupCallIDTypeID); err != nil {
		return fmt.Errorf("unable to decode groupCallId#14e4bb45: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GroupCallID) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallId#14e4bb45 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode groupCallId#14e4bb45: field id: %w", err)
		}
		g.ID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GroupCallID) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode groupCallId#14e4bb45 as nil")
	}
	b.ObjStart()
	b.PutID("groupCallId")
	b.Comma()
	b.FieldStart("id")
	b.PutInt32(g.ID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GroupCallID) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode groupCallId#14e4bb45 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("groupCallId"); err != nil {
				return fmt.Errorf("unable to decode groupCallId#14e4bb45: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode groupCallId#14e4bb45: field id: %w", err)
			}
			g.ID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (g *GroupCallID) GetID() (value int32) {
	if g == nil {
		return
	}
	return g.ID
}
