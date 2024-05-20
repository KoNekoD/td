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

// AccountGetAllSecureValuesRequest represents TL type `account.getAllSecureValues#b288bc7d`.
// Get all saved Telegram Passport¹ documents, for more info see the passport docs »²
//
// Links:
//  1. https://core.telegram.org/passport
//  2. https://core.telegram.org/passport/encryption#encryption
//
// See https://core.telegram.org/method/account.getAllSecureValues for reference.
type AccountGetAllSecureValuesRequest struct {
}

// AccountGetAllSecureValuesRequestTypeID is TL type id of AccountGetAllSecureValuesRequest.
const AccountGetAllSecureValuesRequestTypeID = 0xb288bc7d

// Ensuring interfaces in compile-time for AccountGetAllSecureValuesRequest.
var (
	_ bin.Encoder     = &AccountGetAllSecureValuesRequest{}
	_ bin.Decoder     = &AccountGetAllSecureValuesRequest{}
	_ bin.BareEncoder = &AccountGetAllSecureValuesRequest{}
	_ bin.BareDecoder = &AccountGetAllSecureValuesRequest{}
)

func (g *AccountGetAllSecureValuesRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetAllSecureValuesRequest) String() string {
	if g == nil {
		return "AccountGetAllSecureValuesRequest(nil)"
	}
	type Alias AccountGetAllSecureValuesRequest
	return fmt.Sprintf("AccountGetAllSecureValuesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetAllSecureValuesRequest) TypeID() uint32 {
	return AccountGetAllSecureValuesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetAllSecureValuesRequest) TypeName() string {
	return "account.getAllSecureValues"
}

// TypeInfo returns info about TL type.
func (g *AccountGetAllSecureValuesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getAllSecureValues",
		ID:   AccountGetAllSecureValuesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetAllSecureValuesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAllSecureValues#b288bc7d as nil")
	}
	b.PutID(AccountGetAllSecureValuesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetAllSecureValuesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getAllSecureValues#b288bc7d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetAllSecureValuesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAllSecureValues#b288bc7d to nil")
	}
	if err := b.ConsumeID(AccountGetAllSecureValuesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getAllSecureValues#b288bc7d: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetAllSecureValuesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getAllSecureValues#b288bc7d to nil")
	}
	return nil
}

// AccountGetAllSecureValues invokes method account.getAllSecureValues#b288bc7d returning error if any.
// Get all saved Telegram Passport¹ documents, for more info see the passport docs »²
//
// Links:
//  1. https://core.telegram.org/passport
//  2. https://core.telegram.org/passport/encryption#encryption
//
// See https://core.telegram.org/method/account.getAllSecureValues for reference.
func (c *Client) AccountGetAllSecureValues(ctx context.Context) ([]SecureValue, error) {
	var result SecureValueVector

	request := &AccountGetAllSecureValuesRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return []SecureValue(result.Elems), nil
}
