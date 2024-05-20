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

// GetMessagePublicForwardsRequest represents TL type `getMessagePublicForwards#519da4b4`.
type GetMessagePublicForwardsRequest struct {
	// Chat identifier of the message
	ChatID int64
	// Message identifier
	MessageID int64
	// Offset of the first entry to return as received from the previous request; use empty
	// string to get the first chunk of results
	Offset string
	// The maximum number of messages and stories to be returned; must be positive and can't
	// be greater than 100. For optimal performance, the number of returned objects is chosen
	// by TDLib and can be smaller than the specified limit
	Limit int32
}

// GetMessagePublicForwardsRequestTypeID is TL type id of GetMessagePublicForwardsRequest.
const GetMessagePublicForwardsRequestTypeID = 0x519da4b4

// Ensuring interfaces in compile-time for GetMessagePublicForwardsRequest.
var (
	_ bin.Encoder     = &GetMessagePublicForwardsRequest{}
	_ bin.Decoder     = &GetMessagePublicForwardsRequest{}
	_ bin.BareEncoder = &GetMessagePublicForwardsRequest{}
	_ bin.BareDecoder = &GetMessagePublicForwardsRequest{}
)

func (g *GetMessagePublicForwardsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.Offset == "") {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetMessagePublicForwardsRequest) String() string {
	if g == nil {
		return "GetMessagePublicForwardsRequest(nil)"
	}
	type Alias GetMessagePublicForwardsRequest
	return fmt.Sprintf("GetMessagePublicForwardsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessagePublicForwardsRequest) TypeID() uint32 {
	return GetMessagePublicForwardsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessagePublicForwardsRequest) TypeName() string {
	return "getMessagePublicForwards"
}

// TypeInfo returns info about TL type.
func (g *GetMessagePublicForwardsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessagePublicForwards",
		ID:   GetMessagePublicForwardsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
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
func (g *GetMessagePublicForwardsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessagePublicForwards#519da4b4 as nil")
	}
	b.PutID(GetMessagePublicForwardsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessagePublicForwardsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessagePublicForwards#519da4b4 as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	b.PutString(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessagePublicForwardsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessagePublicForwards#519da4b4 to nil")
	}
	if err := b.ConsumeID(GetMessagePublicForwardsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessagePublicForwards#519da4b4: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessagePublicForwardsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessagePublicForwards#519da4b4 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessagePublicForwards#519da4b4: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessagePublicForwards#519da4b4: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getMessagePublicForwards#519da4b4: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getMessagePublicForwards#519da4b4: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMessagePublicForwardsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessagePublicForwards#519da4b4 as nil")
	}
	b.ObjStart()
	b.PutID("getMessagePublicForwards")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.Comma()
	b.FieldStart("offset")
	b.PutString(g.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetMessagePublicForwardsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessagePublicForwards#519da4b4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMessagePublicForwards"); err != nil {
				return fmt.Errorf("unable to decode getMessagePublicForwards#519da4b4: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessagePublicForwards#519da4b4: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessagePublicForwards#519da4b4: field message_id: %w", err)
			}
			g.MessageID = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getMessagePublicForwards#519da4b4: field offset: %w", err)
			}
			g.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getMessagePublicForwards#519da4b4: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetMessagePublicForwardsRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetMessagePublicForwardsRequest) GetMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageID
}

// GetOffset returns value of Offset field.
func (g *GetMessagePublicForwardsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetMessagePublicForwardsRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetMessagePublicForwards invokes method getMessagePublicForwards#519da4b4 returning error if any.
func (c *Client) GetMessagePublicForwards(ctx context.Context, request *GetMessagePublicForwardsRequest) (*PublicForwards, error) {
	var result PublicForwards

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
