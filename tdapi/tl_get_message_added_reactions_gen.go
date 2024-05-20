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

// GetMessageAddedReactionsRequest represents TL type `getMessageAddedReactions#7dc6ae52`.
type GetMessageAddedReactionsRequest struct {
	// Identifier of the chat to which the message belongs
	ChatID int64
	// Identifier of the message
	MessageID int64
	// Type of the reactions to return; pass null to return all added reactions
	ReactionType ReactionTypeClass
	// Offset of the first entry to return as received from the previous request; use empty
	// string to get the first chunk of results
	Offset string
	// The maximum number of reactions to be returned; must be positive and can't be greater
	// than 100
	Limit int32
}

// GetMessageAddedReactionsRequestTypeID is TL type id of GetMessageAddedReactionsRequest.
const GetMessageAddedReactionsRequestTypeID = 0x7dc6ae52

// Ensuring interfaces in compile-time for GetMessageAddedReactionsRequest.
var (
	_ bin.Encoder     = &GetMessageAddedReactionsRequest{}
	_ bin.Decoder     = &GetMessageAddedReactionsRequest{}
	_ bin.BareEncoder = &GetMessageAddedReactionsRequest{}
	_ bin.BareDecoder = &GetMessageAddedReactionsRequest{}
)

func (g *GetMessageAddedReactionsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.ReactionType == nil) {
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
func (g *GetMessageAddedReactionsRequest) String() string {
	if g == nil {
		return "GetMessageAddedReactionsRequest(nil)"
	}
	type Alias GetMessageAddedReactionsRequest
	return fmt.Sprintf("GetMessageAddedReactionsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessageAddedReactionsRequest) TypeID() uint32 {
	return GetMessageAddedReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessageAddedReactionsRequest) TypeName() string {
	return "getMessageAddedReactions"
}

// TypeInfo returns info about TL type.
func (g *GetMessageAddedReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessageAddedReactions",
		ID:   GetMessageAddedReactionsRequestTypeID,
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
			Name:       "ReactionType",
			SchemaName: "reaction_type",
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
func (g *GetMessageAddedReactionsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageAddedReactions#7dc6ae52 as nil")
	}
	b.PutID(GetMessageAddedReactionsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessageAddedReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageAddedReactions#7dc6ae52 as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	if g.ReactionType == nil {
		return fmt.Errorf("unable to encode getMessageAddedReactions#7dc6ae52: field reaction_type is nil")
	}
	if err := g.ReactionType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getMessageAddedReactions#7dc6ae52: field reaction_type: %w", err)
	}
	b.PutString(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessageAddedReactionsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageAddedReactions#7dc6ae52 to nil")
	}
	if err := b.ConsumeID(GetMessageAddedReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessageAddedReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageAddedReactions#7dc6ae52 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := DecodeReactionType(b)
		if err != nil {
			return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: field reaction_type: %w", err)
		}
		g.ReactionType = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetMessageAddedReactionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageAddedReactions#7dc6ae52 as nil")
	}
	b.ObjStart()
	b.PutID("getMessageAddedReactions")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.Comma()
	b.FieldStart("reaction_type")
	if g.ReactionType == nil {
		return fmt.Errorf("unable to encode getMessageAddedReactions#7dc6ae52: field reaction_type is nil")
	}
	if err := g.ReactionType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getMessageAddedReactions#7dc6ae52: field reaction_type: %w", err)
	}
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
func (g *GetMessageAddedReactionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageAddedReactions#7dc6ae52 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getMessageAddedReactions"); err != nil {
				return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: field message_id: %w", err)
			}
			g.MessageID = value
		case "reaction_type":
			value, err := DecodeTDLibJSONReactionType(b)
			if err != nil {
				return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: field reaction_type: %w", err)
			}
			g.ReactionType = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: field offset: %w", err)
			}
			g.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getMessageAddedReactions#7dc6ae52: field limit: %w", err)
			}
			g.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetMessageAddedReactionsRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetMessageAddedReactionsRequest) GetMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageID
}

// GetReactionType returns value of ReactionType field.
func (g *GetMessageAddedReactionsRequest) GetReactionType() (value ReactionTypeClass) {
	if g == nil {
		return
	}
	return g.ReactionType
}

// GetOffset returns value of Offset field.
func (g *GetMessageAddedReactionsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetMessageAddedReactionsRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetMessageAddedReactions invokes method getMessageAddedReactions#7dc6ae52 returning error if any.
func (c *Client) GetMessageAddedReactions(ctx context.Context, request *GetMessageAddedReactionsRequest) (*AddedReactions, error) {
	var result AddedReactions

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
