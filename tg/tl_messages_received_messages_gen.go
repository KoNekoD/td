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

// MessagesReceivedMessagesRequest represents TL type `messages.receivedMessages#5a954c0`.
// Confirms receipt of messages by a client, cancels PUSH-notification sending.
//
// See https://core.telegram.org/method/messages.receivedMessages for reference.
type MessagesReceivedMessagesRequest struct {
	// Maximum message ID available in a client.
	MaxID int
}

// MessagesReceivedMessagesRequestTypeID is TL type id of MessagesReceivedMessagesRequest.
const MessagesReceivedMessagesRequestTypeID = 0x5a954c0

// Ensuring interfaces in compile-time for MessagesReceivedMessagesRequest.
var (
	_ bin.Encoder     = &MessagesReceivedMessagesRequest{}
	_ bin.Decoder     = &MessagesReceivedMessagesRequest{}
	_ bin.BareEncoder = &MessagesReceivedMessagesRequest{}
	_ bin.BareDecoder = &MessagesReceivedMessagesRequest{}
)

func (r *MessagesReceivedMessagesRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.MaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *MessagesReceivedMessagesRequest) String() string {
	if r == nil {
		return "MessagesReceivedMessagesRequest(nil)"
	}
	type Alias MessagesReceivedMessagesRequest
	return fmt.Sprintf("MessagesReceivedMessagesRequest%+v", Alias(*r))
}

// FillFrom fills MessagesReceivedMessagesRequest from given interface.
func (r *MessagesReceivedMessagesRequest) FillFrom(from interface {
	GetMaxID() (value int)
}) {
	r.MaxID = from.GetMaxID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesReceivedMessagesRequest) TypeID() uint32 {
	return MessagesReceivedMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesReceivedMessagesRequest) TypeName() string {
	return "messages.receivedMessages"
}

// TypeInfo returns info about TL type.
func (r *MessagesReceivedMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.receivedMessages",
		ID:   MessagesReceivedMessagesRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *MessagesReceivedMessagesRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.receivedMessages#5a954c0 as nil")
	}
	b.PutID(MessagesReceivedMessagesRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *MessagesReceivedMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode messages.receivedMessages#5a954c0 as nil")
	}
	b.PutInt(r.MaxID)
	return nil
}

// Decode implements bin.Decoder.
func (r *MessagesReceivedMessagesRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.receivedMessages#5a954c0 to nil")
	}
	if err := b.ConsumeID(MessagesReceivedMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.receivedMessages#5a954c0: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *MessagesReceivedMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode messages.receivedMessages#5a954c0 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.receivedMessages#5a954c0: field max_id: %w", err)
		}
		r.MaxID = value
	}
	return nil
}

// GetMaxID returns value of MaxID field.
func (r *MessagesReceivedMessagesRequest) GetMaxID() (value int) {
	return r.MaxID
}

// MessagesReceivedMessages invokes method messages.receivedMessages#5a954c0 returning error if any.
// Confirms receipt of messages by a client, cancels PUSH-notification sending.
//
// See https://core.telegram.org/method/messages.receivedMessages for reference.
func (c *Client) MessagesReceivedMessages(ctx context.Context, maxid int) ([]ReceivedNotifyMessage, error) {
	var result ReceivedNotifyMessageVector

	request := &MessagesReceivedMessagesRequest{
		MaxID: maxid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []ReceivedNotifyMessage(result.Elems), nil
}
