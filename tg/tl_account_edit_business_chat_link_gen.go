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

// AccountEditBusinessChatLinkRequest represents TL type `account.editBusinessChatLink#8c3410af`.
//
// See https://core.telegram.org/method/account.editBusinessChatLink for reference.
type AccountEditBusinessChatLinkRequest struct {
	// Slug field of AccountEditBusinessChatLinkRequest.
	Slug string
	// Link field of AccountEditBusinessChatLinkRequest.
	Link InputBusinessChatLink
}

// AccountEditBusinessChatLinkRequestTypeID is TL type id of AccountEditBusinessChatLinkRequest.
const AccountEditBusinessChatLinkRequestTypeID = 0x8c3410af

// Ensuring interfaces in compile-time for AccountEditBusinessChatLinkRequest.
var (
	_ bin.Encoder     = &AccountEditBusinessChatLinkRequest{}
	_ bin.Decoder     = &AccountEditBusinessChatLinkRequest{}
	_ bin.BareEncoder = &AccountEditBusinessChatLinkRequest{}
	_ bin.BareDecoder = &AccountEditBusinessChatLinkRequest{}
)

func (e *AccountEditBusinessChatLinkRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Slug == "") {
		return false
	}
	if !(e.Link.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *AccountEditBusinessChatLinkRequest) String() string {
	if e == nil {
		return "AccountEditBusinessChatLinkRequest(nil)"
	}
	type Alias AccountEditBusinessChatLinkRequest
	return fmt.Sprintf("AccountEditBusinessChatLinkRequest%+v", Alias(*e))
}

// FillFrom fills AccountEditBusinessChatLinkRequest from given interface.
func (e *AccountEditBusinessChatLinkRequest) FillFrom(from interface {
	GetSlug() (value string)
	GetLink() (value InputBusinessChatLink)
}) {
	e.Slug = from.GetSlug()
	e.Link = from.GetLink()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountEditBusinessChatLinkRequest) TypeID() uint32 {
	return AccountEditBusinessChatLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountEditBusinessChatLinkRequest) TypeName() string {
	return "account.editBusinessChatLink"
}

// TypeInfo returns info about TL type.
func (e *AccountEditBusinessChatLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.editBusinessChatLink",
		ID:   AccountEditBusinessChatLinkRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Slug",
			SchemaName: "slug",
		},
		{
			Name:       "Link",
			SchemaName: "link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *AccountEditBusinessChatLinkRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode account.editBusinessChatLink#8c3410af as nil")
	}
	b.PutID(AccountEditBusinessChatLinkRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *AccountEditBusinessChatLinkRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode account.editBusinessChatLink#8c3410af as nil")
	}
	b.PutString(e.Slug)
	if err := e.Link.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.editBusinessChatLink#8c3410af: field link: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *AccountEditBusinessChatLinkRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode account.editBusinessChatLink#8c3410af to nil")
	}
	if err := b.ConsumeID(AccountEditBusinessChatLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.editBusinessChatLink#8c3410af: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *AccountEditBusinessChatLinkRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode account.editBusinessChatLink#8c3410af to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.editBusinessChatLink#8c3410af: field slug: %w", err)
		}
		e.Slug = value
	}
	{
		if err := e.Link.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.editBusinessChatLink#8c3410af: field link: %w", err)
		}
	}
	return nil
}

// GetSlug returns value of Slug field.
func (e *AccountEditBusinessChatLinkRequest) GetSlug() (value string) {
	if e == nil {
		return
	}
	return e.Slug
}

// GetLink returns value of Link field.
func (e *AccountEditBusinessChatLinkRequest) GetLink() (value InputBusinessChatLink) {
	if e == nil {
		return
	}
	return e.Link
}

// AccountEditBusinessChatLink invokes method account.editBusinessChatLink#8c3410af returning error if any.
//
// See https://core.telegram.org/method/account.editBusinessChatLink for reference.
func (c *Client) AccountEditBusinessChatLink(ctx context.Context, request *AccountEditBusinessChatLinkRequest) (*BusinessChatLink, error) {
	var result BusinessChatLink

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
