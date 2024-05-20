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

// ChannelsGetParticipantsRequest represents TL type `channels.getParticipants#77ced9d0`.
// Get the participants of a supergroup/channel¹
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.getParticipants for reference.
type ChannelsGetParticipantsRequest struct {
	// Channel
	Channel InputChannelClass
	// Which participant types to fetch
	Filter ChannelParticipantsFilterClass
	// Offset¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Offset int
	// Limit¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
	// Hash¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Hash int64
}

// ChannelsGetParticipantsRequestTypeID is TL type id of ChannelsGetParticipantsRequest.
const ChannelsGetParticipantsRequestTypeID = 0x77ced9d0

// Ensuring interfaces in compile-time for ChannelsGetParticipantsRequest.
var (
	_ bin.Encoder     = &ChannelsGetParticipantsRequest{}
	_ bin.Decoder     = &ChannelsGetParticipantsRequest{}
	_ bin.BareEncoder = &ChannelsGetParticipantsRequest{}
	_ bin.BareDecoder = &ChannelsGetParticipantsRequest{}
)

func (g *ChannelsGetParticipantsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.Filter == nil) {
		return false
	}
	if !(g.Offset == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *ChannelsGetParticipantsRequest) String() string {
	if g == nil {
		return "ChannelsGetParticipantsRequest(nil)"
	}
	type Alias ChannelsGetParticipantsRequest
	return fmt.Sprintf("ChannelsGetParticipantsRequest%+v", Alias(*g))
}

// FillFrom fills ChannelsGetParticipantsRequest from given interface.
func (g *ChannelsGetParticipantsRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetFilter() (value ChannelParticipantsFilterClass)
	GetOffset() (value int)
	GetLimit() (value int)
	GetHash() (value int64)
}) {
	g.Channel = from.GetChannel()
	g.Filter = from.GetFilter()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsGetParticipantsRequest) TypeID() uint32 {
	return ChannelsGetParticipantsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsGetParticipantsRequest) TypeName() string {
	return "channels.getParticipants"
}

// TypeInfo returns info about TL type.
func (g *ChannelsGetParticipantsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.getParticipants",
		ID:   ChannelsGetParticipantsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *ChannelsGetParticipantsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getParticipants#77ced9d0 as nil")
	}
	b.PutID(ChannelsGetParticipantsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *ChannelsGetParticipantsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode channels.getParticipants#77ced9d0 as nil")
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode channels.getParticipants#77ced9d0: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getParticipants#77ced9d0: field channel: %w", err)
	}
	if g.Filter == nil {
		return fmt.Errorf("unable to encode channels.getParticipants#77ced9d0: field filter is nil")
	}
	if err := g.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.getParticipants#77ced9d0: field filter: %w", err)
	}
	b.PutInt(g.Offset)
	b.PutInt(g.Limit)
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *ChannelsGetParticipantsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getParticipants#77ced9d0 to nil")
	}
	if err := b.ConsumeID(ChannelsGetParticipantsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.getParticipants#77ced9d0: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *ChannelsGetParticipantsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode channels.getParticipants#77ced9d0 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#77ced9d0: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := DecodeChannelParticipantsFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#77ced9d0: field filter: %w", err)
		}
		g.Filter = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#77ced9d0: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#77ced9d0: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode channels.getParticipants#77ced9d0: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (g *ChannelsGetParticipantsRequest) GetChannel() (value InputChannelClass) {
	if g == nil {
		return
	}
	return g.Channel
}

// GetFilter returns value of Filter field.
func (g *ChannelsGetParticipantsRequest) GetFilter() (value ChannelParticipantsFilterClass) {
	if g == nil {
		return
	}
	return g.Filter
}

// GetOffset returns value of Offset field.
func (g *ChannelsGetParticipantsRequest) GetOffset() (value int) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *ChannelsGetParticipantsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetHash returns value of Hash field.
func (g *ChannelsGetParticipantsRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (g *ChannelsGetParticipantsRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return g.Channel.AsNotEmpty()
}

// ChannelsGetParticipants invokes method channels.getParticipants#77ced9d0 returning error if any.
// Get the participants of a supergroup/channel¹
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	406 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	403 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//
// See https://core.telegram.org/method/channels.getParticipants for reference.
// Can be used by bots.
func (c *Client) ChannelsGetParticipants(ctx context.Context, request *ChannelsGetParticipantsRequest) (ChannelsChannelParticipantsClass, error) {
	var result ChannelsChannelParticipantsBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ChannelParticipants, nil
}
