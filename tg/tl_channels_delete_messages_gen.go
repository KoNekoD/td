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

// ChannelsDeleteMessagesRequest represents TL type `channels.deleteMessages#84c1fd4e`.
// Delete messages in a channel/supergroup¹
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.deleteMessages for reference.
type ChannelsDeleteMessagesRequest struct {
	// Channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// IDs of messages to delete
	ID []int
}

// ChannelsDeleteMessagesRequestTypeID is TL type id of ChannelsDeleteMessagesRequest.
const ChannelsDeleteMessagesRequestTypeID = 0x84c1fd4e

// Ensuring interfaces in compile-time for ChannelsDeleteMessagesRequest.
var (
	_ bin.Encoder     = &ChannelsDeleteMessagesRequest{}
	_ bin.Decoder     = &ChannelsDeleteMessagesRequest{}
	_ bin.BareEncoder = &ChannelsDeleteMessagesRequest{}
	_ bin.BareDecoder = &ChannelsDeleteMessagesRequest{}
)

func (d *ChannelsDeleteMessagesRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Channel == nil) {
		return false
	}
	if !(d.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChannelsDeleteMessagesRequest) String() string {
	if d == nil {
		return "ChannelsDeleteMessagesRequest(nil)"
	}
	type Alias ChannelsDeleteMessagesRequest
	return fmt.Sprintf("ChannelsDeleteMessagesRequest%+v", Alias(*d))
}

// FillFrom fills ChannelsDeleteMessagesRequest from given interface.
func (d *ChannelsDeleteMessagesRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetID() (value []int)
}) {
	d.Channel = from.GetChannel()
	d.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsDeleteMessagesRequest) TypeID() uint32 {
	return ChannelsDeleteMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsDeleteMessagesRequest) TypeName() string {
	return "channels.deleteMessages"
}

// TypeInfo returns info about TL type.
func (d *ChannelsDeleteMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.deleteMessages",
		ID:   ChannelsDeleteMessagesRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *ChannelsDeleteMessagesRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteMessages#84c1fd4e as nil")
	}
	b.PutID(ChannelsDeleteMessagesRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *ChannelsDeleteMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteMessages#84c1fd4e as nil")
	}
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteMessages#84c1fd4e: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteMessages#84c1fd4e: field channel: %w", err)
	}
	b.PutVectorHeader(len(d.ID))
	for _, v := range d.ID {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteMessagesRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteMessages#84c1fd4e to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *ChannelsDeleteMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteMessages#84c1fd4e to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: field id: %w", err)
		}

		if headerLen > 0 {
			d.ID = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode channels.deleteMessages#84c1fd4e: field id: %w", err)
			}
			d.ID = append(d.ID, value)
		}
	}
	return nil
}

// GetChannel returns value of Channel field.
func (d *ChannelsDeleteMessagesRequest) GetChannel() (value InputChannelClass) {
	if d == nil {
		return
	}
	return d.Channel
}

// GetID returns value of ID field.
func (d *ChannelsDeleteMessagesRequest) GetID() (value []int) {
	if d == nil {
		return
	}
	return d.ID
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (d *ChannelsDeleteMessagesRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return d.Channel.AsNotEmpty()
}

// ChannelsDeleteMessages invokes method channels.deleteMessages#84c1fd4e returning error if any.
// Delete messages in a channel/supergroup¹
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	406 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	403 MESSAGE_DELETE_FORBIDDEN: You can't delete one of the messages you tried to delete, most likely because it is a service message.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//
// See https://core.telegram.org/method/channels.deleteMessages for reference.
// Can be used by bots.
func (c *Client) ChannelsDeleteMessages(ctx context.Context, request *ChannelsDeleteMessagesRequest) (*MessagesAffectedMessages, error) {
	var result MessagesAffectedMessages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
