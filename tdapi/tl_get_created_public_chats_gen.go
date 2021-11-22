// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// GetCreatedPublicChatsRequest represents TL type `getCreatedPublicChats#2a5725ef`.
type GetCreatedPublicChatsRequest struct {
	// Type of the public chats to return
	Type PublicChatTypeClass
}

// GetCreatedPublicChatsRequestTypeID is TL type id of GetCreatedPublicChatsRequest.
const GetCreatedPublicChatsRequestTypeID = 0x2a5725ef

// Ensuring interfaces in compile-time for GetCreatedPublicChatsRequest.
var (
	_ bin.Encoder     = &GetCreatedPublicChatsRequest{}
	_ bin.Decoder     = &GetCreatedPublicChatsRequest{}
	_ bin.BareEncoder = &GetCreatedPublicChatsRequest{}
	_ bin.BareDecoder = &GetCreatedPublicChatsRequest{}
)

func (g *GetCreatedPublicChatsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetCreatedPublicChatsRequest) String() string {
	if g == nil {
		return "GetCreatedPublicChatsRequest(nil)"
	}
	type Alias GetCreatedPublicChatsRequest
	return fmt.Sprintf("GetCreatedPublicChatsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetCreatedPublicChatsRequest) TypeID() uint32 {
	return GetCreatedPublicChatsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetCreatedPublicChatsRequest) TypeName() string {
	return "getCreatedPublicChats"
}

// TypeInfo returns info about TL type.
func (g *GetCreatedPublicChatsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getCreatedPublicChats",
		ID:   GetCreatedPublicChatsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetCreatedPublicChatsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCreatedPublicChats#2a5725ef as nil")
	}
	b.PutID(GetCreatedPublicChatsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetCreatedPublicChatsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getCreatedPublicChats#2a5725ef as nil")
	}
	if g.Type == nil {
		return fmt.Errorf("unable to encode getCreatedPublicChats#2a5725ef: field type is nil")
	}
	if err := g.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getCreatedPublicChats#2a5725ef: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetCreatedPublicChatsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCreatedPublicChats#2a5725ef to nil")
	}
	if err := b.ConsumeID(GetCreatedPublicChatsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getCreatedPublicChats#2a5725ef: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetCreatedPublicChatsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getCreatedPublicChats#2a5725ef to nil")
	}
	{
		value, err := DecodePublicChatType(b)
		if err != nil {
			return fmt.Errorf("unable to decode getCreatedPublicChats#2a5725ef: field type: %w", err)
		}
		g.Type = value
	}
	return nil
}

// EncodeTDLibJSON encodes g in TDLib API JSON format.
func (g *GetCreatedPublicChatsRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getCreatedPublicChats#2a5725ef as nil")
	}
	b.ObjStart()
	b.PutID("getCreatedPublicChats")
	b.FieldStart("type")
	if g.Type == nil {
		return fmt.Errorf("unable to encode getCreatedPublicChats#2a5725ef: field type is nil")
	}
	if err := g.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getCreatedPublicChats#2a5725ef: field type: %w", err)
	}
	b.ObjEnd()
	return nil
}

// GetType returns value of Type field.
func (g *GetCreatedPublicChatsRequest) GetType() (value PublicChatTypeClass) {
	return g.Type
}

// GetCreatedPublicChats invokes method getCreatedPublicChats#2a5725ef returning error if any.
func (c *Client) GetCreatedPublicChats(ctx context.Context, type_ PublicChatTypeClass) (*Chats, error) {
	var result Chats

	request := &GetCreatedPublicChatsRequest{
		Type: type_,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
