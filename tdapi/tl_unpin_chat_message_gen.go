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

// UnpinChatMessageRequest represents TL type `unpinChatMessage#7b1c3ede`.
type UnpinChatMessageRequest struct {
	// Identifier of the chat
	ChatID int64
	// Identifier of the removed pinned message
	MessageID int64
}

// UnpinChatMessageRequestTypeID is TL type id of UnpinChatMessageRequest.
const UnpinChatMessageRequestTypeID = 0x7b1c3ede

// Ensuring interfaces in compile-time for UnpinChatMessageRequest.
var (
	_ bin.Encoder     = &UnpinChatMessageRequest{}
	_ bin.Decoder     = &UnpinChatMessageRequest{}
	_ bin.BareEncoder = &UnpinChatMessageRequest{}
	_ bin.BareDecoder = &UnpinChatMessageRequest{}
)

func (u *UnpinChatMessageRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.ChatID == 0) {
		return false
	}
	if !(u.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UnpinChatMessageRequest) String() string {
	if u == nil {
		return "UnpinChatMessageRequest(nil)"
	}
	type Alias UnpinChatMessageRequest
	return fmt.Sprintf("UnpinChatMessageRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UnpinChatMessageRequest) TypeID() uint32 {
	return UnpinChatMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UnpinChatMessageRequest) TypeName() string {
	return "unpinChatMessage"
}

// TypeInfo returns info about TL type.
func (u *UnpinChatMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "unpinChatMessage",
		ID:   UnpinChatMessageRequestTypeID,
	}
	if u == nil {
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UnpinChatMessageRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode unpinChatMessage#7b1c3ede as nil")
	}
	b.PutID(UnpinChatMessageRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UnpinChatMessageRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode unpinChatMessage#7b1c3ede as nil")
	}
	b.PutLong(u.ChatID)
	b.PutLong(u.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UnpinChatMessageRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode unpinChatMessage#7b1c3ede to nil")
	}
	if err := b.ConsumeID(UnpinChatMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode unpinChatMessage#7b1c3ede: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UnpinChatMessageRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode unpinChatMessage#7b1c3ede to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode unpinChatMessage#7b1c3ede: field chat_id: %w", err)
		}
		u.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode unpinChatMessage#7b1c3ede: field message_id: %w", err)
		}
		u.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON encodes u in TDLib API JSON format.
func (u *UnpinChatMessageRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode unpinChatMessage#7b1c3ede as nil")
	}
	b.ObjStart()
	b.PutID("unpinChatMessage")
	b.FieldStart("chat_id")
	b.PutLong(u.ChatID)
	b.FieldStart("message_id")
	b.PutLong(u.MessageID)
	b.ObjEnd()
	return nil
}

// GetChatID returns value of ChatID field.
func (u *UnpinChatMessageRequest) GetChatID() (value int64) {
	return u.ChatID
}

// GetMessageID returns value of MessageID field.
func (u *UnpinChatMessageRequest) GetMessageID() (value int64) {
	return u.MessageID
}

// UnpinChatMessage invokes method unpinChatMessage#7b1c3ede returning error if any.
func (c *Client) UnpinChatMessage(ctx context.Context, request *UnpinChatMessageRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
