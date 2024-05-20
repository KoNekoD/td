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

// AuthResetLoginEmailRequest represents TL type `auth.resetLoginEmail#7e960193`.
// Reset the login email »¹.
//
// Links:
//  1. https://core.telegram.org/api/auth#email-verification
//
// See https://core.telegram.org/method/auth.resetLoginEmail for reference.
type AuthResetLoginEmailRequest struct {
	// Phone number of the account
	PhoneNumber string
	// Phone code hash, obtained as described in the documentation »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/auth
	PhoneCodeHash string
}

// AuthResetLoginEmailRequestTypeID is TL type id of AuthResetLoginEmailRequest.
const AuthResetLoginEmailRequestTypeID = 0x7e960193

// Ensuring interfaces in compile-time for AuthResetLoginEmailRequest.
var (
	_ bin.Encoder     = &AuthResetLoginEmailRequest{}
	_ bin.Decoder     = &AuthResetLoginEmailRequest{}
	_ bin.BareEncoder = &AuthResetLoginEmailRequest{}
	_ bin.BareDecoder = &AuthResetLoginEmailRequest{}
)

func (r *AuthResetLoginEmailRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.PhoneNumber == "") {
		return false
	}
	if !(r.PhoneCodeHash == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AuthResetLoginEmailRequest) String() string {
	if r == nil {
		return "AuthResetLoginEmailRequest(nil)"
	}
	type Alias AuthResetLoginEmailRequest
	return fmt.Sprintf("AuthResetLoginEmailRequest%+v", Alias(*r))
}

// FillFrom fills AuthResetLoginEmailRequest from given interface.
func (r *AuthResetLoginEmailRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetPhoneCodeHash() (value string)
}) {
	r.PhoneNumber = from.GetPhoneNumber()
	r.PhoneCodeHash = from.GetPhoneCodeHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthResetLoginEmailRequest) TypeID() uint32 {
	return AuthResetLoginEmailRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthResetLoginEmailRequest) TypeName() string {
	return "auth.resetLoginEmail"
}

// TypeInfo returns info about TL type.
func (r *AuthResetLoginEmailRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.resetLoginEmail",
		ID:   AuthResetLoginEmailRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "PhoneCodeHash",
			SchemaName: "phone_code_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *AuthResetLoginEmailRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.resetLoginEmail#7e960193 as nil")
	}
	b.PutID(AuthResetLoginEmailRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AuthResetLoginEmailRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.resetLoginEmail#7e960193 as nil")
	}
	b.PutString(r.PhoneNumber)
	b.PutString(r.PhoneCodeHash)
	return nil
}

// Decode implements bin.Decoder.
func (r *AuthResetLoginEmailRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.resetLoginEmail#7e960193 to nil")
	}
	if err := b.ConsumeID(AuthResetLoginEmailRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.resetLoginEmail#7e960193: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AuthResetLoginEmailRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.resetLoginEmail#7e960193 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.resetLoginEmail#7e960193: field phone_number: %w", err)
		}
		r.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.resetLoginEmail#7e960193: field phone_code_hash: %w", err)
		}
		r.PhoneCodeHash = value
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (r *AuthResetLoginEmailRequest) GetPhoneNumber() (value string) {
	if r == nil {
		return
	}
	return r.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (r *AuthResetLoginEmailRequest) GetPhoneCodeHash() (value string) {
	if r == nil {
		return
	}
	return r.PhoneCodeHash
}

// AuthResetLoginEmail invokes method auth.resetLoginEmail#7e960193 returning error if any.
// Reset the login email »¹.
//
// Links:
//  1. https://core.telegram.org/api/auth#email-verification
//
// Possible errors:
//
//	400 PHONE_NUMBER_INVALID: The phone number is invalid.
//	400 TASK_ALREADY_EXISTS: An email reset was already requested.
//
// See https://core.telegram.org/method/auth.resetLoginEmail for reference.
func (c *Client) AuthResetLoginEmail(ctx context.Context, request *AuthResetLoginEmailRequest) (AuthSentCodeClass, error) {
	var result AuthSentCodeBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.SentCode, nil
}
