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

// GetBlockedMessageSendersRequest represents TL type `getBlockedMessageSenders#740e1460`.
type GetBlockedMessageSendersRequest struct {
	// Number of users and chats to skip in the result; must be non-negative
	Offset int32
	// The maximum number of users and chats to return; up to 100
	Limit int32
}

// GetBlockedMessageSendersRequestTypeID is TL type id of GetBlockedMessageSendersRequest.
const GetBlockedMessageSendersRequestTypeID = 0x740e1460

// Ensuring interfaces in compile-time for GetBlockedMessageSendersRequest.
var (
	_ bin.Encoder     = &GetBlockedMessageSendersRequest{}
	_ bin.Decoder     = &GetBlockedMessageSendersRequest{}
	_ bin.BareEncoder = &GetBlockedMessageSendersRequest{}
	_ bin.BareDecoder = &GetBlockedMessageSendersRequest{}
)

func (g *GetBlockedMessageSendersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetBlockedMessageSendersRequest) String() string {
	if g == nil {
		return "GetBlockedMessageSendersRequest(nil)"
	}
	type Alias GetBlockedMessageSendersRequest
	return fmt.Sprintf("GetBlockedMessageSendersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetBlockedMessageSendersRequest) TypeID() uint32 {
	return GetBlockedMessageSendersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetBlockedMessageSendersRequest) TypeName() string {
	return "getBlockedMessageSenders"
}

// TypeInfo returns info about TL type.
func (g *GetBlockedMessageSendersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getBlockedMessageSenders",
		ID:   GetBlockedMessageSendersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetBlockedMessageSendersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBlockedMessageSenders#740e1460 as nil")
	}
	b.PutID(GetBlockedMessageSendersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetBlockedMessageSendersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBlockedMessageSenders#740e1460 as nil")
	}
	b.PutInt32(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetBlockedMessageSendersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBlockedMessageSenders#740e1460 to nil")
	}
	if err := b.ConsumeID(GetBlockedMessageSendersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getBlockedMessageSenders#740e1460: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetBlockedMessageSendersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBlockedMessageSenders#740e1460 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getBlockedMessageSenders#740e1460: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getBlockedMessageSenders#740e1460: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON encodes g in TDLib API JSON format.
func (g *GetBlockedMessageSendersRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getBlockedMessageSenders#740e1460 as nil")
	}
	b.ObjStart()
	b.PutID("getBlockedMessageSenders")
	b.FieldStart("offset")
	b.PutInt32(g.Offset)
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.ObjEnd()
	return nil
}

// GetOffset returns value of Offset field.
func (g *GetBlockedMessageSendersRequest) GetOffset() (value int32) {
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetBlockedMessageSendersRequest) GetLimit() (value int32) {
	return g.Limit
}

// GetBlockedMessageSenders invokes method getBlockedMessageSenders#740e1460 returning error if any.
func (c *Client) GetBlockedMessageSenders(ctx context.Context, request *GetBlockedMessageSendersRequest) (*MessageSenders, error) {
	var result MessageSenders

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
