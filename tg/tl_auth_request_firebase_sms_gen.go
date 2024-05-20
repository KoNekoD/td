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

// AuthRequestFirebaseSMSRequest represents TL type `auth.requestFirebaseSms#89464b50`.
// Request an SMS code via Firebase.
//
// See https://core.telegram.org/method/auth.requestFirebaseSms for reference.
type AuthRequestFirebaseSMSRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Phone number
	PhoneNumber string
	// Phone code hash returned by auth.sendCode¹
	//
	// Links:
	//  1) https://core.telegram.org/method/auth.sendCode
	PhoneCodeHash string
	// On Android, a JWS object obtained as described in the auth documentation »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/auth
	//
	// Use SetSafetyNetToken and GetSafetyNetToken helpers.
	SafetyNetToken string
	// Secret token received via an apple push notification
	//
	// Use SetIosPushSecret and GetIosPushSecret helpers.
	IosPushSecret string
}

// AuthRequestFirebaseSMSRequestTypeID is TL type id of AuthRequestFirebaseSMSRequest.
const AuthRequestFirebaseSMSRequestTypeID = 0x89464b50

// Ensuring interfaces in compile-time for AuthRequestFirebaseSMSRequest.
var (
	_ bin.Encoder     = &AuthRequestFirebaseSMSRequest{}
	_ bin.Decoder     = &AuthRequestFirebaseSMSRequest{}
	_ bin.BareEncoder = &AuthRequestFirebaseSMSRequest{}
	_ bin.BareDecoder = &AuthRequestFirebaseSMSRequest{}
)

func (r *AuthRequestFirebaseSMSRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.PhoneNumber == "") {
		return false
	}
	if !(r.PhoneCodeHash == "") {
		return false
	}
	if !(r.SafetyNetToken == "") {
		return false
	}
	if !(r.IosPushSecret == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AuthRequestFirebaseSMSRequest) String() string {
	if r == nil {
		return "AuthRequestFirebaseSMSRequest(nil)"
	}
	type Alias AuthRequestFirebaseSMSRequest
	return fmt.Sprintf("AuthRequestFirebaseSMSRequest%+v", Alias(*r))
}

// FillFrom fills AuthRequestFirebaseSMSRequest from given interface.
func (r *AuthRequestFirebaseSMSRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetPhoneCodeHash() (value string)
	GetSafetyNetToken() (value string, ok bool)
	GetIosPushSecret() (value string, ok bool)
}) {
	r.PhoneNumber = from.GetPhoneNumber()
	r.PhoneCodeHash = from.GetPhoneCodeHash()
	if val, ok := from.GetSafetyNetToken(); ok {
		r.SafetyNetToken = val
	}

	if val, ok := from.GetIosPushSecret(); ok {
		r.IosPushSecret = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthRequestFirebaseSMSRequest) TypeID() uint32 {
	return AuthRequestFirebaseSMSRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthRequestFirebaseSMSRequest) TypeName() string {
	return "auth.requestFirebaseSms"
}

// TypeInfo returns info about TL type.
func (r *AuthRequestFirebaseSMSRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.requestFirebaseSms",
		ID:   AuthRequestFirebaseSMSRequestTypeID,
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
		{
			Name:       "SafetyNetToken",
			SchemaName: "safety_net_token",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "IosPushSecret",
			SchemaName: "ios_push_secret",
			Null:       !r.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *AuthRequestFirebaseSMSRequest) SetFlags() {
	if !(r.SafetyNetToken == "") {
		r.Flags.Set(0)
	}
	if !(r.IosPushSecret == "") {
		r.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (r *AuthRequestFirebaseSMSRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.requestFirebaseSms#89464b50 as nil")
	}
	b.PutID(AuthRequestFirebaseSMSRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *AuthRequestFirebaseSMSRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode auth.requestFirebaseSms#89464b50 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.requestFirebaseSms#89464b50: field flags: %w", err)
	}
	b.PutString(r.PhoneNumber)
	b.PutString(r.PhoneCodeHash)
	if r.Flags.Has(0) {
		b.PutString(r.SafetyNetToken)
	}
	if r.Flags.Has(1) {
		b.PutString(r.IosPushSecret)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *AuthRequestFirebaseSMSRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.requestFirebaseSms#89464b50 to nil")
	}
	if err := b.ConsumeID(AuthRequestFirebaseSMSRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.requestFirebaseSms#89464b50: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *AuthRequestFirebaseSMSRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode auth.requestFirebaseSms#89464b50 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode auth.requestFirebaseSms#89464b50: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.requestFirebaseSms#89464b50: field phone_number: %w", err)
		}
		r.PhoneNumber = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.requestFirebaseSms#89464b50: field phone_code_hash: %w", err)
		}
		r.PhoneCodeHash = value
	}
	if r.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.requestFirebaseSms#89464b50: field safety_net_token: %w", err)
		}
		r.SafetyNetToken = value
	}
	if r.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode auth.requestFirebaseSms#89464b50: field ios_push_secret: %w", err)
		}
		r.IosPushSecret = value
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (r *AuthRequestFirebaseSMSRequest) GetPhoneNumber() (value string) {
	if r == nil {
		return
	}
	return r.PhoneNumber
}

// GetPhoneCodeHash returns value of PhoneCodeHash field.
func (r *AuthRequestFirebaseSMSRequest) GetPhoneCodeHash() (value string) {
	if r == nil {
		return
	}
	return r.PhoneCodeHash
}

// SetSafetyNetToken sets value of SafetyNetToken conditional field.
func (r *AuthRequestFirebaseSMSRequest) SetSafetyNetToken(value string) {
	r.Flags.Set(0)
	r.SafetyNetToken = value
}

// GetSafetyNetToken returns value of SafetyNetToken conditional field and
// boolean which is true if field was set.
func (r *AuthRequestFirebaseSMSRequest) GetSafetyNetToken() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.SafetyNetToken, true
}

// SetIosPushSecret sets value of IosPushSecret conditional field.
func (r *AuthRequestFirebaseSMSRequest) SetIosPushSecret(value string) {
	r.Flags.Set(1)
	r.IosPushSecret = value
}

// GetIosPushSecret returns value of IosPushSecret conditional field and
// boolean which is true if field was set.
func (r *AuthRequestFirebaseSMSRequest) GetIosPushSecret() (value string, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(1) {
		return value, false
	}
	return r.IosPushSecret, true
}

// AuthRequestFirebaseSMS invokes method auth.requestFirebaseSms#89464b50 returning error if any.
// Request an SMS code via Firebase.
//
// Possible errors:
//
//	400 PHONE_NUMBER_INVALID: The phone number is invalid.
//
// See https://core.telegram.org/method/auth.requestFirebaseSms for reference.
func (c *Client) AuthRequestFirebaseSMS(ctx context.Context, request *AuthRequestFirebaseSMSRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
