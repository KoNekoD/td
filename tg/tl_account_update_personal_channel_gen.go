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

// AccountUpdatePersonalChannelRequest represents TL type `account.updatePersonalChannel#d94305e0`.
//
// See https://core.telegram.org/method/account.updatePersonalChannel for reference.
type AccountUpdatePersonalChannelRequest struct {
	// Channel field of AccountUpdatePersonalChannelRequest.
	Channel InputChannelClass
}

// AccountUpdatePersonalChannelRequestTypeID is TL type id of AccountUpdatePersonalChannelRequest.
const AccountUpdatePersonalChannelRequestTypeID = 0xd94305e0

// Ensuring interfaces in compile-time for AccountUpdatePersonalChannelRequest.
var (
	_ bin.Encoder     = &AccountUpdatePersonalChannelRequest{}
	_ bin.Decoder     = &AccountUpdatePersonalChannelRequest{}
	_ bin.BareEncoder = &AccountUpdatePersonalChannelRequest{}
	_ bin.BareDecoder = &AccountUpdatePersonalChannelRequest{}
)

func (u *AccountUpdatePersonalChannelRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Channel == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdatePersonalChannelRequest) String() string {
	if u == nil {
		return "AccountUpdatePersonalChannelRequest(nil)"
	}
	type Alias AccountUpdatePersonalChannelRequest
	return fmt.Sprintf("AccountUpdatePersonalChannelRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdatePersonalChannelRequest from given interface.
func (u *AccountUpdatePersonalChannelRequest) FillFrom(from interface {
	GetChannel() (value InputChannelClass)
}) {
	u.Channel = from.GetChannel()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdatePersonalChannelRequest) TypeID() uint32 {
	return AccountUpdatePersonalChannelRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdatePersonalChannelRequest) TypeName() string {
	return "account.updatePersonalChannel"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdatePersonalChannelRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updatePersonalChannel",
		ID:   AccountUpdatePersonalChannelRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUpdatePersonalChannelRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updatePersonalChannel#d94305e0 as nil")
	}
	b.PutID(AccountUpdatePersonalChannelRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdatePersonalChannelRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updatePersonalChannel#d94305e0 as nil")
	}
	if u.Channel == nil {
		return fmt.Errorf("unable to encode account.updatePersonalChannel#d94305e0: field channel is nil")
	}
	if err := u.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updatePersonalChannel#d94305e0: field channel: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdatePersonalChannelRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updatePersonalChannel#d94305e0 to nil")
	}
	if err := b.ConsumeID(AccountUpdatePersonalChannelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updatePersonalChannel#d94305e0: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdatePersonalChannelRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updatePersonalChannel#d94305e0 to nil")
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.updatePersonalChannel#d94305e0: field channel: %w", err)
		}
		u.Channel = value
	}
	return nil
}

// GetChannel returns value of Channel field.
func (u *AccountUpdatePersonalChannelRequest) GetChannel() (value InputChannelClass) {
	if u == nil {
		return
	}
	return u.Channel
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (u *AccountUpdatePersonalChannelRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return u.Channel.AsNotEmpty()
}

// AccountUpdatePersonalChannel invokes method account.updatePersonalChannel#d94305e0 returning error if any.
//
// See https://core.telegram.org/method/account.updatePersonalChannel for reference.
func (c *Client) AccountUpdatePersonalChannel(ctx context.Context, channel InputChannelClass) (bool, error) {
	var result BoolBox

	request := &AccountUpdatePersonalChannelRequest{
		Channel: channel,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
