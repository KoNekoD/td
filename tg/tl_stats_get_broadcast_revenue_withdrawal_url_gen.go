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

// StatsGetBroadcastRevenueWithdrawalURLRequest represents TL type `stats.getBroadcastRevenueWithdrawalUrl#2a65ef73`.
//
// See https://core.telegram.org/method/stats.getBroadcastRevenueWithdrawalUrl for reference.
type StatsGetBroadcastRevenueWithdrawalURLRequest struct {
	// Channel field of StatsGetBroadcastRevenueWithdrawalURLRequest.
	Channel InputChannelClass
	// Password field of StatsGetBroadcastRevenueWithdrawalURLRequest.
	Password InputCheckPasswordSRPClass
}

// StatsGetBroadcastRevenueWithdrawalURLRequestTypeID is TL type id of StatsGetBroadcastRevenueWithdrawalURLRequest.
const StatsGetBroadcastRevenueWithdrawalURLRequestTypeID = 0x2a65ef73

// Ensuring interfaces in compile-time for StatsGetBroadcastRevenueWithdrawalURLRequest.
var (
	_ bin.Encoder     = &StatsGetBroadcastRevenueWithdrawalURLRequest{}
	_ bin.Decoder     = &StatsGetBroadcastRevenueWithdrawalURLRequest{}
	_ bin.BareEncoder = &StatsGetBroadcastRevenueWithdrawalURLRequest{}
	_ bin.BareDecoder = &StatsGetBroadcastRevenueWithdrawalURLRequest{}
)

func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Channel == nil) {
		return false
	}
	if !(g.Password == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) String() string {
	if g == nil {
		return "StatsGetBroadcastRevenueWithdrawalURLRequest(nil)"
	}
	type Alias StatsGetBroadcastRevenueWithdrawalURLRequest
	return fmt.Sprintf("StatsGetBroadcastRevenueWithdrawalURLRequest%+v", Alias(*g))
}

// FillFrom fills StatsGetBroadcastRevenueWithdrawalURLRequest from given interface.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
	GetPassword() (value InputCheckPasswordSRPClass)
}) {
	g.Channel = from.GetChannel()
	g.Password = from.GetPassword()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGetBroadcastRevenueWithdrawalURLRequest) TypeID() uint32 {
	return StatsGetBroadcastRevenueWithdrawalURLRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGetBroadcastRevenueWithdrawalURLRequest) TypeName() string {
	return "stats.getBroadcastRevenueWithdrawalUrl"
}

// TypeInfo returns info about TL type.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stats.getBroadcastRevenueWithdrawalUrl",
		ID:   StatsGetBroadcastRevenueWithdrawalURLRequestTypeID,
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
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73 as nil")
	}
	b.PutID(StatsGetBroadcastRevenueWithdrawalURLRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73 as nil")
	}
	if g.Channel == nil {
		return fmt.Errorf("unable to encode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73: field channel is nil")
	}
	if err := g.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73: field channel: %w", err)
	}
	if g.Password == nil {
		return fmt.Errorf("unable to encode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73: field password is nil")
	}
	if err := g.Password.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73: field password: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73 to nil")
	}
	if err := b.ConsumeID(StatsGetBroadcastRevenueWithdrawalURLRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73: field channel: %w", err)
		}
		g.Channel = value
	}
	{
		value, err := DecodeInputCheckPasswordSRP(b)
		if err != nil {
			return fmt.Errorf("unable to decode stats.getBroadcastRevenueWithdrawalUrl#2a65ef73: field password: %w", err)
		}
		g.Password = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) GetChannel() (value InputChannelClass) {
	if g == nil {
		return
	}
	return g.Channel
}

// GetPassword returns value of Password field.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) GetPassword() (value InputCheckPasswordSRPClass) {
	if g == nil {
		return
	}
	return g.Password
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return g.Channel.AsNotEmpty()
}

// GetPasswordAsNotEmpty returns mapped value of Password field.
func (g *StatsGetBroadcastRevenueWithdrawalURLRequest) GetPasswordAsNotEmpty() (*InputCheckPasswordSRP, bool) {
	return g.Password.AsNotEmpty()
}

// StatsGetBroadcastRevenueWithdrawalURL invokes method stats.getBroadcastRevenueWithdrawalUrl#2a65ef73 returning error if any.
//
// See https://core.telegram.org/method/stats.getBroadcastRevenueWithdrawalUrl for reference.
func (c *Client) StatsGetBroadcastRevenueWithdrawalURL(ctx context.Context, request *StatsGetBroadcastRevenueWithdrawalURLRequest) (*StatsBroadcastRevenueWithdrawalURL, error) {
	var result StatsBroadcastRevenueWithdrawalURL

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
