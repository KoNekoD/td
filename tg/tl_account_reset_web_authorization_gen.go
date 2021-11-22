// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/jsontd"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
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
	_ = jsontd.Encoder{}
)

// AccountResetWebAuthorizationRequest represents TL type `account.resetWebAuthorization#2d01b9ef`.
// Log out an active web telegram login¹ session
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.resetWebAuthorization for reference.
type AccountResetWebAuthorizationRequest struct {
	// Session¹ hash
	//
	// Links:
	//  1) https://core.telegram.org/constructor/webAuthorization
	Hash int64
}

// AccountResetWebAuthorizationRequestTypeID is TL type id of AccountResetWebAuthorizationRequest.
const AccountResetWebAuthorizationRequestTypeID = 0x2d01b9ef

// Ensuring interfaces in compile-time for AccountResetWebAuthorizationRequest.
var (
	_ bin.Encoder     = &AccountResetWebAuthorizationRequest{}
	_ bin.Decoder     = &AccountResetWebAuthorizationRequest{}
	_ bin.BareEncoder = &AccountResetWebAuthorizationRequest{}
	_ bin.BareDecoder = &AccountResetWebAuthorizationRequest{}
)

func (r *AccountResetWebAuthorizationRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetWebAuthorizationRequest) String() string {
	if r == nil {
		return "AccountResetWebAuthorizationRequest(nil)"
	}
	type Alias AccountResetWebAuthorizationRequest
	return fmt.Sprintf("AccountResetWebAuthorizationRequest%+v", Alias(*r))
}

// FillFrom fills AccountResetWebAuthorizationRequest from given interface.
func (r *AccountResetWebAuthorizationRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	r.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountResetWebAuthorizationRequest) TypeID() uint32 {
	return AccountResetWebAuthorizationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountResetWebAuthorizationRequest) TypeName() string {
	return "account.resetWebAuthorization"
}

// TypeInfo returns info about TL type.
func (r *AccountResetWebAuthorizationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.resetWebAuthorization",
		ID:   AccountResetWebAuthorizationRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *AccountResetWebAuthorizationRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetWebAuthorization#2d01b9ef as nil")
	}
	b.PutID(AccountResetWebAuthorizationRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AccountResetWebAuthorizationRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetWebAuthorization#2d01b9ef as nil")
	}
	b.PutLong(r.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetWebAuthorizationRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetWebAuthorization#2d01b9ef to nil")
	}
	if err := b.ConsumeID(AccountResetWebAuthorizationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetWebAuthorization#2d01b9ef: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AccountResetWebAuthorizationRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetWebAuthorization#2d01b9ef to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.resetWebAuthorization#2d01b9ef: field hash: %w", err)
		}
		r.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (r *AccountResetWebAuthorizationRequest) GetHash() (value int64) {
	return r.Hash
}

// AccountResetWebAuthorization invokes method account.resetWebAuthorization#2d01b9ef returning error if any.
// Log out an active web telegram login¹ session
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.resetWebAuthorization for reference.
func (c *Client) AccountResetWebAuthorization(ctx context.Context, hash int64) (bool, error) {
	var result BoolBox

	request := &AccountResetWebAuthorizationRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
