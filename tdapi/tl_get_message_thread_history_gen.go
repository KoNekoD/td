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

// GetMessageThreadHistoryRequest represents TL type `getMessageThreadHistory#9435d428`.
type GetMessageThreadHistoryRequest struct {
	// Chat identifier
	ChatID int64
	// Message identifier, which thread history needs to be returned
	MessageID int64
	// Identifier of the message starting from which history must be fetched; use 0 to get
	// results from the last message
	FromMessageID int64
	// Specify 0 to get results from exactly the from_message_id or a negative offset up to
	// 99 to get additionally some newer messages
	Offset int32
	// The maximum number of messages to be returned; must be positive and can't be greater
	// than 100. If the offset is negative, the limit must be greater than or equal to
	// -offset. For optimal performance, the number of returned messages is chosen by TDLib
	// and can be smaller than the specified limit
	Limit int32
}

// GetMessageThreadHistoryRequestTypeID is TL type id of GetMessageThreadHistoryRequest.
const GetMessageThreadHistoryRequestTypeID = 0x9435d428

// Ensuring interfaces in compile-time for GetMessageThreadHistoryRequest.
var (
	_ bin.Encoder     = &GetMessageThreadHistoryRequest{}
	_ bin.Decoder     = &GetMessageThreadHistoryRequest{}
	_ bin.BareEncoder = &GetMessageThreadHistoryRequest{}
	_ bin.BareDecoder = &GetMessageThreadHistoryRequest{}
)

func (g *GetMessageThreadHistoryRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.FromMessageID == 0) {
		return false
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
func (g *GetMessageThreadHistoryRequest) String() string {
	if g == nil {
		return "GetMessageThreadHistoryRequest(nil)"
	}
	type Alias GetMessageThreadHistoryRequest
	return fmt.Sprintf("GetMessageThreadHistoryRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetMessageThreadHistoryRequest) TypeID() uint32 {
	return GetMessageThreadHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetMessageThreadHistoryRequest) TypeName() string {
	return "getMessageThreadHistory"
}

// TypeInfo returns info about TL type.
func (g *GetMessageThreadHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getMessageThreadHistory",
		ID:   GetMessageThreadHistoryRequestTypeID,
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
			Name:       "FromMessageID",
			SchemaName: "from_message_id",
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
func (g *GetMessageThreadHistoryRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageThreadHistory#9435d428 as nil")
	}
	b.PutID(GetMessageThreadHistoryRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetMessageThreadHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageThreadHistory#9435d428 as nil")
	}
	b.PutLong(g.ChatID)
	b.PutLong(g.MessageID)
	b.PutLong(g.FromMessageID)
	b.PutInt32(g.Offset)
	b.PutInt32(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetMessageThreadHistoryRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageThreadHistory#9435d428 to nil")
	}
	if err := b.ConsumeID(GetMessageThreadHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getMessageThreadHistory#9435d428: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetMessageThreadHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getMessageThreadHistory#9435d428 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageThreadHistory#9435d428: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageThreadHistory#9435d428: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageThreadHistory#9435d428: field from_message_id: %w", err)
		}
		g.FromMessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageThreadHistory#9435d428: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getMessageThreadHistory#9435d428: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// EncodeTDLibJSON encodes g in TDLib API JSON format.
func (g *GetMessageThreadHistoryRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getMessageThreadHistory#9435d428 as nil")
	}
	b.ObjStart()
	b.PutID("getMessageThreadHistory")
	b.FieldStart("chat_id")
	b.PutLong(g.ChatID)
	b.FieldStart("message_id")
	b.PutLong(g.MessageID)
	b.FieldStart("from_message_id")
	b.PutLong(g.FromMessageID)
	b.FieldStart("offset")
	b.PutInt32(g.Offset)
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.ObjEnd()
	return nil
}

// GetChatID returns value of ChatID field.
func (g *GetMessageThreadHistoryRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetMessageThreadHistoryRequest) GetMessageID() (value int64) {
	return g.MessageID
}

// GetFromMessageID returns value of FromMessageID field.
func (g *GetMessageThreadHistoryRequest) GetFromMessageID() (value int64) {
	return g.FromMessageID
}

// GetOffset returns value of Offset field.
func (g *GetMessageThreadHistoryRequest) GetOffset() (value int32) {
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *GetMessageThreadHistoryRequest) GetLimit() (value int32) {
	return g.Limit
}

// GetMessageThreadHistory invokes method getMessageThreadHistory#9435d428 returning error if any.
func (c *Client) GetMessageThreadHistory(ctx context.Context, request *GetMessageThreadHistoryRequest) (*Messages, error) {
	var result Messages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
