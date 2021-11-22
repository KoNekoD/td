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

// PhoneNumberAuthenticationSettings represents TL type `phoneNumberAuthenticationSettings#ccc9aae9`.
type PhoneNumberAuthenticationSettings struct {
	// Pass true if the authentication code may be sent via flash call to the specified phone
	// number
	AllowFlashCall bool
	// Pass true if the authenticated phone number is used on the current device
	IsCurrentPhoneNumber bool
	// For official applications only. True, if the application can use Android SMS Retriever
	// API (requires Google Play Services >= 10.2) to automatically receive the
	// authentication code from the SMS. See https://developers.google
	// com/identity/sms-retriever/ for more details
	AllowSMSRetrieverAPI bool
}

// PhoneNumberAuthenticationSettingsTypeID is TL type id of PhoneNumberAuthenticationSettings.
const PhoneNumberAuthenticationSettingsTypeID = 0xccc9aae9

// Ensuring interfaces in compile-time for PhoneNumberAuthenticationSettings.
var (
	_ bin.Encoder     = &PhoneNumberAuthenticationSettings{}
	_ bin.Decoder     = &PhoneNumberAuthenticationSettings{}
	_ bin.BareEncoder = &PhoneNumberAuthenticationSettings{}
	_ bin.BareDecoder = &PhoneNumberAuthenticationSettings{}
)

func (p *PhoneNumberAuthenticationSettings) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.AllowFlashCall == false) {
		return false
	}
	if !(p.IsCurrentPhoneNumber == false) {
		return false
	}
	if !(p.AllowSMSRetrieverAPI == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhoneNumberAuthenticationSettings) String() string {
	if p == nil {
		return "PhoneNumberAuthenticationSettings(nil)"
	}
	type Alias PhoneNumberAuthenticationSettings
	return fmt.Sprintf("PhoneNumberAuthenticationSettings%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneNumberAuthenticationSettings) TypeID() uint32 {
	return PhoneNumberAuthenticationSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneNumberAuthenticationSettings) TypeName() string {
	return "phoneNumberAuthenticationSettings"
}

// TypeInfo returns info about TL type.
func (p *PhoneNumberAuthenticationSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phoneNumberAuthenticationSettings",
		ID:   PhoneNumberAuthenticationSettingsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AllowFlashCall",
			SchemaName: "allow_flash_call",
		},
		{
			Name:       "IsCurrentPhoneNumber",
			SchemaName: "is_current_phone_number",
		},
		{
			Name:       "AllowSMSRetrieverAPI",
			SchemaName: "allow_sms_retriever_api",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhoneNumberAuthenticationSettings) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneNumberAuthenticationSettings#ccc9aae9 as nil")
	}
	b.PutID(PhoneNumberAuthenticationSettingsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhoneNumberAuthenticationSettings) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneNumberAuthenticationSettings#ccc9aae9 as nil")
	}
	b.PutBool(p.AllowFlashCall)
	b.PutBool(p.IsCurrentPhoneNumber)
	b.PutBool(p.AllowSMSRetrieverAPI)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneNumberAuthenticationSettings) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneNumberAuthenticationSettings#ccc9aae9 to nil")
	}
	if err := b.ConsumeID(PhoneNumberAuthenticationSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#ccc9aae9: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhoneNumberAuthenticationSettings) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneNumberAuthenticationSettings#ccc9aae9 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#ccc9aae9: field allow_flash_call: %w", err)
		}
		p.AllowFlashCall = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#ccc9aae9: field is_current_phone_number: %w", err)
		}
		p.IsCurrentPhoneNumber = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#ccc9aae9: field allow_sms_retriever_api: %w", err)
		}
		p.AllowSMSRetrieverAPI = value
	}
	return nil
}

// EncodeTDLibJSON encodes p in TDLib API JSON format.
func (p *PhoneNumberAuthenticationSettings) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneNumberAuthenticationSettings#ccc9aae9 as nil")
	}
	b.ObjStart()
	b.PutID("phoneNumberAuthenticationSettings")
	b.FieldStart("allow_flash_call")
	b.PutBool(p.AllowFlashCall)
	b.FieldStart("is_current_phone_number")
	b.PutBool(p.IsCurrentPhoneNumber)
	b.FieldStart("allow_sms_retriever_api")
	b.PutBool(p.AllowSMSRetrieverAPI)
	b.ObjEnd()
	return nil
}

// GetAllowFlashCall returns value of AllowFlashCall field.
func (p *PhoneNumberAuthenticationSettings) GetAllowFlashCall() (value bool) {
	return p.AllowFlashCall
}

// GetIsCurrentPhoneNumber returns value of IsCurrentPhoneNumber field.
func (p *PhoneNumberAuthenticationSettings) GetIsCurrentPhoneNumber() (value bool) {
	return p.IsCurrentPhoneNumber
}

// GetAllowSMSRetrieverAPI returns value of AllowSMSRetrieverAPI field.
func (p *PhoneNumberAuthenticationSettings) GetAllowSMSRetrieverAPI() (value bool) {
	return p.AllowSMSRetrieverAPI
}
