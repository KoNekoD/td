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

// StatsGetStoryPublicForwardsRequest represents TL type `stats.getStoryPublicForwards#a6437ef6`.
// Obtain forwards of a story¹ as a message to public chats and reposts by public
// channels.
//
// Links:
//  1. https://core.telegram.org/api/stories
//
// See https://core.telegram.org/method/stats.getStoryPublicForwards for reference.
type StatsGetStoryPublicForwardsRequest struct {
	// Peer where the story was originally posted
	Peer InputPeerClass
	// Story¹ ID
	//
	// Links:
	//  1) https://core.telegram.org/api/stories
	ID int
	// Offset for pagination, from stats.PublicForwards¹.next_offset.
	//
	// Links:
	//  1) https://core.telegram.org/constructor/stats.publicForwards
	Offset string
	// Maximum number of results to return, see pagination¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets
	Limit int
}

// StatsGetStoryPublicForwardsRequestTypeID is TL type id of StatsGetStoryPublicForwardsRequest.
const StatsGetStoryPublicForwardsRequestTypeID = 0xa6437ef6

// Ensuring interfaces in compile-time for StatsGetStoryPublicForwardsRequest.
var (
	_ bin.Encoder     = &StatsGetStoryPublicForwardsRequest{}
	_ bin.Decoder     = &StatsGetStoryPublicForwardsRequest{}
	_ bin.BareEncoder = &StatsGetStoryPublicForwardsRequest{}
	_ bin.BareDecoder = &StatsGetStoryPublicForwardsRequest{}
)

func (g *StatsGetStoryPublicForwardsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Peer == nil) {
		return false
	}
	if !(g.ID == 0) {
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
func (g *StatsGetStoryPublicForwardsRequest) String() string {
	if g == nil {
		return "StatsGetStoryPublicForwardsRequest(nil)"
	}
	type Alias StatsGetStoryPublicForwardsRequest
	return fmt.Sprintf("StatsGetStoryPublicForwardsRequest%+v", Alias(*g))
}

// FillFrom fills StatsGetStoryPublicForwardsRequest from given interface.
func (g *StatsGetStoryPublicForwardsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetID() (value int)
	GetOffset() (value string)
	GetLimit() (value int)
}) {
	g.Peer = from.GetPeer()
	g.ID = from.GetID()
	g.Offset = from.GetOffset()
	g.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGetStoryPublicForwardsRequest) TypeID() uint32 {
	return StatsGetStoryPublicForwardsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGetStoryPublicForwardsRequest) TypeName() string {
	return "stats.getStoryPublicForwards"
}

// TypeInfo returns info about TL type.
func (g *StatsGetStoryPublicForwardsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.getStoryPublicForwards",
		ID:   StatsGetStoryPublicForwardsRequestTypeID,
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
func (g *StatsGetStoryPublicForwardsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getStoryPublicForwards#a6437ef6 as nil")
	}
	b.PutID(StatsGetStoryPublicForwardsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StatsGetStoryPublicForwardsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getStoryPublicForwards#a6437ef6 as nil")
	}
	if g.Peer == nil {
		return fmt.Errorf("unable to encode stats.getStoryPublicForwards#a6437ef6: field peer is nil")
	}
	if err := g.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getStoryPublicForwards#a6437ef6: field peer: %w", err)
	}
	b.PutInt(g.ID)
	b.PutString(g.Offset)
	b.PutInt(g.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (g *StatsGetStoryPublicForwardsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getStoryPublicForwards#a6437ef6 to nil")
	}
	if err := b.ConsumeID(StatsGetStoryPublicForwardsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getStoryPublicForwards#a6437ef6: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StatsGetStoryPublicForwardsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getStoryPublicForwards#a6437ef6 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getStoryPublicForwards#a6437ef6: field peer: %w", err)
		}
		g.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getStoryPublicForwards#a6437ef6: field id: %w", err)
		}
		g.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getStoryPublicForwards#a6437ef6: field offset: %w", err)
		}
		g.Offset = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stats.getStoryPublicForwards#a6437ef6: field limit: %w", err)
		}
		g.Limit = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (g *StatsGetStoryPublicForwardsRequest) GetPeer() (value InputPeerClass) {
	if g == nil {
		return
	}
	return g.Peer
}

// GetID returns value of ID field.
func (g *StatsGetStoryPublicForwardsRequest) GetID() (value int) {
	if g == nil {
		return
	}
	return g.ID
}

// GetOffset returns value of Offset field.
func (g *StatsGetStoryPublicForwardsRequest) GetOffset() (value string) {
	if g == nil {
		return
	}
	return g.Offset
}

// GetLimit returns value of Limit field.
func (g *StatsGetStoryPublicForwardsRequest) GetLimit() (value int) {
	if g == nil {
		return
	}
	return g.Limit
}

// StatsGetStoryPublicForwards invokes method stats.getStoryPublicForwards#a6437ef6 returning error if any.
// Obtain forwards of a story¹ as a message to public chats and reposts by public
// channels.
//
// Links:
//  1. https://core.telegram.org/api/stories
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/stats.getStoryPublicForwards for reference.
func (c *Client) StatsGetStoryPublicForwards(ctx context.Context, request *StatsGetStoryPublicForwardsRequest) (*StatsPublicForwards, error) {
	var result StatsPublicForwards

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
