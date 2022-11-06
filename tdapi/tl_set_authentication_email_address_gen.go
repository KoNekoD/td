// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// SetAuthenticationEmailAddressRequest represents TL type `setAuthenticationEmailAddress#69b2c502`.
type SetAuthenticationEmailAddressRequest struct {
	// The email address of the user
	EmailAddress string
}

// SetAuthenticationEmailAddressRequestTypeID is TL type id of SetAuthenticationEmailAddressRequest.
const SetAuthenticationEmailAddressRequestTypeID = 0x69b2c502

// Ensuring interfaces in compile-time for SetAuthenticationEmailAddressRequest.
var (
	_ bin.Encoder     = &SetAuthenticationEmailAddressRequest{}
	_ bin.Decoder     = &SetAuthenticationEmailAddressRequest{}
	_ bin.BareEncoder = &SetAuthenticationEmailAddressRequest{}
	_ bin.BareDecoder = &SetAuthenticationEmailAddressRequest{}
)

func (s *SetAuthenticationEmailAddressRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.EmailAddress == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetAuthenticationEmailAddressRequest) String() string {
	if s == nil {
		return "SetAuthenticationEmailAddressRequest(nil)"
	}
	type Alias SetAuthenticationEmailAddressRequest
	return fmt.Sprintf("SetAuthenticationEmailAddressRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetAuthenticationEmailAddressRequest) TypeID() uint32 {
	return SetAuthenticationEmailAddressRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetAuthenticationEmailAddressRequest) TypeName() string {
	return "setAuthenticationEmailAddress"
}

// TypeInfo returns info about TL type.
func (s *SetAuthenticationEmailAddressRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setAuthenticationEmailAddress",
		ID:   SetAuthenticationEmailAddressRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EmailAddress",
			SchemaName: "email_address",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetAuthenticationEmailAddressRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setAuthenticationEmailAddress#69b2c502 as nil")
	}
	b.PutID(SetAuthenticationEmailAddressRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetAuthenticationEmailAddressRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setAuthenticationEmailAddress#69b2c502 as nil")
	}
	b.PutString(s.EmailAddress)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetAuthenticationEmailAddressRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setAuthenticationEmailAddress#69b2c502 to nil")
	}
	if err := b.ConsumeID(SetAuthenticationEmailAddressRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setAuthenticationEmailAddress#69b2c502: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetAuthenticationEmailAddressRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setAuthenticationEmailAddress#69b2c502 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setAuthenticationEmailAddress#69b2c502: field email_address: %w", err)
		}
		s.EmailAddress = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetAuthenticationEmailAddressRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setAuthenticationEmailAddress#69b2c502 as nil")
	}
	b.ObjStart()
	b.PutID("setAuthenticationEmailAddress")
	b.Comma()
	b.FieldStart("email_address")
	b.PutString(s.EmailAddress)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetAuthenticationEmailAddressRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setAuthenticationEmailAddress#69b2c502 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setAuthenticationEmailAddress"); err != nil {
				return fmt.Errorf("unable to decode setAuthenticationEmailAddress#69b2c502: %w", err)
			}
		case "email_address":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setAuthenticationEmailAddress#69b2c502: field email_address: %w", err)
			}
			s.EmailAddress = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEmailAddress returns value of EmailAddress field.
func (s *SetAuthenticationEmailAddressRequest) GetEmailAddress() (value string) {
	if s == nil {
		return
	}
	return s.EmailAddress
}

// SetAuthenticationEmailAddress invokes method setAuthenticationEmailAddress#69b2c502 returning error if any.
func (c *Client) SetAuthenticationEmailAddress(ctx context.Context, emailaddress string) error {
	var ok Ok

	request := &SetAuthenticationEmailAddressRequest{
		EmailAddress: emailaddress,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}