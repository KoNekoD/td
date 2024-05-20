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

// MessagesGetMessagesViewsRequest represents TL type `messages.getMessagesViews#5784d3e1`.
// Get and increase the view counter of a message sent or forwarded from a channel¹
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/messages.getMessagesViews for reference.
type MessagesGetMessagesViewsRequest struct {
	// Peer where the message was found
	Peer InputPeerClass
	// ID of message
	ID []int
	// Whether to mark the message as viewed and increment the view counter
	Increment bool
}

// MessagesGetMessagesViewsRequestTypeID is TL type id of MessagesGetMessagesViewsRequest.
const MessagesGetMessagesViewsRequestTypeID = 0x5784d3e1

// Ensuring interfaces in compile-time for MessagesGetMessagesViewsRequest.
var (
	_ bin.Encoder     = &MessagesGetMessagesViewsRequest{}
	_ bin.Decoder     = &MessagesGetMessagesViewsRequest{}
	_ bin.BareEncoder = &MessagesGetMessagesViewsRequest{}
	_ bin.BareDecoder = &MessagesGetMessagesViewsRequest{}
)

func (g *MessagesGetMessagesViewsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == nil) {
		return false
	}
	if !(g.Increment == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetMessagesViewsRequest) String() string {
	if g == nil {
		return "MessagesGetMessagesViewsRequest(nil)"
	}
	type Alias MessagesGetMessagesViewsRequest
	return fmt.Sprintf("MessagesGetMessagesViewsRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetMessagesViewsRequest from given interface.
func (g *MessagesGetMessagesViewsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value []int)
	GetIncrement() (value bool)
}) {
	g.Peer = from.GetPeer()
	g.ID = from.GetID()
	g.Increment = from.GetIncrement()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetMessagesViewsRequest) TypeID() uint32 {
	return MessagesGetMessagesViewsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetMessagesViewsRequest) TypeName() string {
	return "messages.getMessagesViews"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetMessagesViewsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getMessagesViews",
		ID:   MessagesGetMessagesViewsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Increment",
			SchemaName: "increment",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetMessagesViewsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessagesViews#5784d3e1 as nil")
	}
	b.PutID(MessagesGetMessagesViewsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetMessagesViewsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getMessagesViews#5784d3e1 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode messages.getMessagesViews#5784d3e1: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getMessagesViews#5784d3e1: field peer: %w", err)
	}
	b.PutVectorHeader(len(g.ID))
	for _, v := range g.ID {
		b.PutInt(v)
	}
	b.PutBool(g.Increment)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetMessagesViewsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessagesViews#5784d3e1 to nil")
	}
	if err := b.ConsumeID(MessagesGetMessagesViewsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getMessagesViews#5784d3e1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetMessagesViewsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getMessagesViews#5784d3e1 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessagesViews#5784d3e1: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessagesViews#5784d3e1: field id: %w", err)
		}

		if headerLen > 0 {
			g.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.getMessagesViews#5784d3e1: field id: %w", err)
			}
			g.ID = append(g.ID, value)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getMessagesViews#5784d3e1: field increment: %w", err)
		}
		g.Increment = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *MessagesGetMessagesViewsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetID returns value of ID field.
func (g *MessagesGetMessagesViewsRequest) GetID() (value []int) {
	if g == nil {
		return
	}
	return g.ID
}

// GetIncrement returns value of Increment field.
func (g *MessagesGetMessagesViewsRequest) GetIncrement() (value bool) {
	if g == nil {
		return
	}
	return g.Increment
}

// MessagesGetMessagesViews invokes method messages.getMessagesViews#5784d3e1 returning error if any.
// Get and increase the view counter of a message sent or forwarded from a channel¹
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	406 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ID_INVALID: The provided chat id is invalid.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.getMessagesViews for reference.
func (c *Client) MessagesGetMessagesViews(ctx context.Context, request *MessagesGetMessagesViewsRequest) (*MessagesMessageViews, error) {
	var result MessagesMessageViews

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
