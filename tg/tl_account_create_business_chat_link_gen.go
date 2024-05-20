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

// AccountCreateBusinessChatLinkRequest represents TL type `account.createBusinessChatLink#8851e68e`.
//
// See https://core.telegram.org/method/account.createBusinessChatLink for reference.
type AccountCreateBusinessChatLinkRequest struct {
	// Link field of AccountCreateBusinessChatLinkRequest.
	Link InputBusinessChatLink
}

// AccountCreateBusinessChatLinkRequestTypeID is TL type id of AccountCreateBusinessChatLinkRequest.
const AccountCreateBusinessChatLinkRequestTypeID = 0x8851e68e

// Ensuring interfaces in compile-time for AccountCreateBusinessChatLinkRequest.
var (
	_ bin.Encoder     = &AccountCreateBusinessChatLinkRequest{}
	_ bin.Decoder     = &AccountCreateBusinessChatLinkRequest{}
	_ bin.BareEncoder = &AccountCreateBusinessChatLinkRequest{}
	_ bin.BareDecoder = &AccountCreateBusinessChatLinkRequest{}
)

func (c *AccountCreateBusinessChatLinkRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Link.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *AccountCreateBusinessChatLinkRequest) String() string {
	if c == nil {
		return "AccountCreateBusinessChatLinkRequest(nil)"
	}
	type Alias AccountCreateBusinessChatLinkRequest
	return fmt.Sprintf("AccountCreateBusinessChatLinkRequest%+v", Alias(*c))
}

// FillFrom fills AccountCreateBusinessChatLinkRequest from given interface.
func (c *AccountCreateBusinessChatLinkRequest) FillFrom(from interface {
	GetLink() (value InputBusinessChatLink)
}) {
	c.Link = from.GetLink()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountCreateBusinessChatLinkRequest) TypeID() uint32 {
	return AccountCreateBusinessChatLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountCreateBusinessChatLinkRequest) TypeName() string {
	return "account.createBusinessChatLink"
}

// TypeInfo returns info about TL type.
func (c *AccountCreateBusinessChatLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.createBusinessChatLink",
		ID:   AccountCreateBusinessChatLinkRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Link",
			SchemaName: "link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *AccountCreateBusinessChatLinkRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.createBusinessChatLink#8851e68e as nil")
	}
	b.PutID(AccountCreateBusinessChatLinkRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *AccountCreateBusinessChatLinkRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode account.createBusinessChatLink#8851e68e as nil")
	}
	if err := c.Link.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.createBusinessChatLink#8851e68e: field link: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *AccountCreateBusinessChatLinkRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.createBusinessChatLink#8851e68e to nil")
	}
	if err := b.ConsumeID(AccountCreateBusinessChatLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.createBusinessChatLink#8851e68e: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *AccountCreateBusinessChatLinkRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode account.createBusinessChatLink#8851e68e to nil")
	}
	{
		if err := c.Link.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.createBusinessChatLink#8851e68e: field link: %w", err)
		}
	}
	return nil
}

// GetLink returns value of Link field.
func (c *AccountCreateBusinessChatLinkRequest) GetLink() (value InputBusinessChatLink) {
	if c == nil {
		return
	}
	return c.Link
}

// AccountCreateBusinessChatLink invokes method account.createBusinessChatLink#8851e68e returning error if any.
//
// See https://core.telegram.org/method/account.createBusinessChatLink for reference.
func (c *Client) AccountCreateBusinessChatLink(ctx context.Context, link InputBusinessChatLink) (*BusinessChatLink, error) {
	var result BusinessChatLink

	request := &AccountCreateBusinessChatLinkRequest{
		Link: link,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
