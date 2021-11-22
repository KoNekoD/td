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

// SetAuthenticationPhoneNumberRequest represents TL type `setAuthenticationPhoneNumber#33c0d823`.
type SetAuthenticationPhoneNumberRequest struct {
	// The phone number of the user, in international format
	PhoneNumber string
	// Settings for the authentication of the user's phone number
	Settings PhoneNumberAuthenticationSettings
}

// SetAuthenticationPhoneNumberRequestTypeID is TL type id of SetAuthenticationPhoneNumberRequest.
const SetAuthenticationPhoneNumberRequestTypeID = 0x33c0d823

// Ensuring interfaces in compile-time for SetAuthenticationPhoneNumberRequest.
var (
	_ bin.Encoder     = &SetAuthenticationPhoneNumberRequest{}
	_ bin.Decoder     = &SetAuthenticationPhoneNumberRequest{}
	_ bin.BareEncoder = &SetAuthenticationPhoneNumberRequest{}
	_ bin.BareDecoder = &SetAuthenticationPhoneNumberRequest{}
)

func (s *SetAuthenticationPhoneNumberRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PhoneNumber == "") {
		return false
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetAuthenticationPhoneNumberRequest) String() string {
	if s == nil {
		return "SetAuthenticationPhoneNumberRequest(nil)"
	}
	type Alias SetAuthenticationPhoneNumberRequest
	return fmt.Sprintf("SetAuthenticationPhoneNumberRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetAuthenticationPhoneNumberRequest) TypeID() uint32 {
	return SetAuthenticationPhoneNumberRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetAuthenticationPhoneNumberRequest) TypeName() string {
	return "setAuthenticationPhoneNumber"
}

// TypeInfo returns info about TL type.
func (s *SetAuthenticationPhoneNumberRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setAuthenticationPhoneNumber",
		ID:   SetAuthenticationPhoneNumberRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetAuthenticationPhoneNumberRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setAuthenticationPhoneNumber#33c0d823 as nil")
	}
	b.PutID(SetAuthenticationPhoneNumberRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetAuthenticationPhoneNumberRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setAuthenticationPhoneNumber#33c0d823 as nil")
	}
	b.PutString(s.PhoneNumber)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setAuthenticationPhoneNumber#33c0d823: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetAuthenticationPhoneNumberRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setAuthenticationPhoneNumber#33c0d823 to nil")
	}
	if err := b.ConsumeID(SetAuthenticationPhoneNumberRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setAuthenticationPhoneNumber#33c0d823: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetAuthenticationPhoneNumberRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setAuthenticationPhoneNumber#33c0d823 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setAuthenticationPhoneNumber#33c0d823: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setAuthenticationPhoneNumber#33c0d823: field settings: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON encodes s in TDLib API JSON format.
func (s *SetAuthenticationPhoneNumberRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setAuthenticationPhoneNumber#33c0d823 as nil")
	}
	b.ObjStart()
	b.PutID("setAuthenticationPhoneNumber")
	b.FieldStart("phone_number")
	b.PutString(s.PhoneNumber)
	b.FieldStart("settings")
	if err := s.Settings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setAuthenticationPhoneNumber#33c0d823: field settings: %w", err)
	}
	b.ObjEnd()
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (s *SetAuthenticationPhoneNumberRequest) GetPhoneNumber() (value string) {
	return s.PhoneNumber
}

// GetSettings returns value of Settings field.
func (s *SetAuthenticationPhoneNumberRequest) GetSettings() (value PhoneNumberAuthenticationSettings) {
	return s.Settings
}

// SetAuthenticationPhoneNumber invokes method setAuthenticationPhoneNumber#33c0d823 returning error if any.
func (c *Client) SetAuthenticationPhoneNumber(ctx context.Context, request *SetAuthenticationPhoneNumberRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
