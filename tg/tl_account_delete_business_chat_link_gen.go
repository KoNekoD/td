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

// AccountDeleteBusinessChatLinkRequest represents TL type `account.deleteBusinessChatLink#60073674`.
//
// See https://core.telegram.org/method/account.deleteBusinessChatLink for reference.
type AccountDeleteBusinessChatLinkRequest struct {
	// Slug field of AccountDeleteBusinessChatLinkRequest.
	Slug string
}

// AccountDeleteBusinessChatLinkRequestTypeID is TL type id of AccountDeleteBusinessChatLinkRequest.
const AccountDeleteBusinessChatLinkRequestTypeID = 0x60073674

// Ensuring interfaces in compile-time for AccountDeleteBusinessChatLinkRequest.
var (
	_ bin.Encoder     = &AccountDeleteBusinessChatLinkRequest{}
	_ bin.Decoder     = &AccountDeleteBusinessChatLinkRequest{}
	_ bin.BareEncoder = &AccountDeleteBusinessChatLinkRequest{}
	_ bin.BareDecoder = &AccountDeleteBusinessChatLinkRequest{}
)

func (d *AccountDeleteBusinessChatLinkRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Slug == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *AccountDeleteBusinessChatLinkRequest) String() string {
	if d == nil {
		return "AccountDeleteBusinessChatLinkRequest(nil)"
	}
	type Alias AccountDeleteBusinessChatLinkRequest
	return fmt.Sprintf("AccountDeleteBusinessChatLinkRequest%+v", Alias(*d))
}

// FillFrom fills AccountDeleteBusinessChatLinkRequest from given interface.
func (d *AccountDeleteBusinessChatLinkRequest) FillFrom(from interface {
	GetSlug() (value string)
}) {
	d.Slug = from.GetSlug()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountDeleteBusinessChatLinkRequest) TypeID() uint32 {
	return AccountDeleteBusinessChatLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountDeleteBusinessChatLinkRequest) TypeName() string {
	return "account.deleteBusinessChatLink"
}

// TypeInfo returns info about TL type.
func (d *AccountDeleteBusinessChatLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.deleteBusinessChatLink",
		ID:   AccountDeleteBusinessChatLinkRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Slug",
			SchemaName: "slug",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *AccountDeleteBusinessChatLinkRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode account.deleteBusinessChatLink#60073674 as nil")
	}
	b.PutID(AccountDeleteBusinessChatLinkRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *AccountDeleteBusinessChatLinkRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode account.deleteBusinessChatLink#60073674 as nil")
	}
	b.PutString(d.Slug)
	return nil
}

// Decode implements bin.Decoder.
func (d *AccountDeleteBusinessChatLinkRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode account.deleteBusinessChatLink#60073674 to nil")
	}
	if err := b.ConsumeID(AccountDeleteBusinessChatLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.deleteBusinessChatLink#60073674: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *AccountDeleteBusinessChatLinkRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode account.deleteBusinessChatLink#60073674 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.deleteBusinessChatLink#60073674: field slug: %w", err)
		}
		d.Slug = value
	}
	return nil
}

// GetSlug returns value of Slug field.
func (d *AccountDeleteBusinessChatLinkRequest) GetSlug() (value string) {
	if d == nil {
		return
	}
	return d.Slug
}

// AccountDeleteBusinessChatLink invokes method account.deleteBusinessChatLink#60073674 returning error if any.
//
// See https://core.telegram.org/method/account.deleteBusinessChatLink for reference.
func (c *Client) AccountDeleteBusinessChatLink(ctx context.Context, slug string) (bool, error) {
	var result BoolBox

	request := &AccountDeleteBusinessChatLinkRequest{
		Slug: slug,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
