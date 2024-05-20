// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// PhoneNumberAuthenticationSettings represents TL type `phoneNumberAuthenticationSettings#3e1183db`.
type PhoneNumberAuthenticationSettings struct {
	// Pass true if the authentication code may be sent via a flash call to the specified
	// phone number
	AllowFlashCall bool
	// Pass true if the authentication code may be sent via a missed call to the specified
	// phone number
	AllowMissedCall bool
	// Pass true if the authenticated phone number is used on the current device
	IsCurrentPhoneNumber bool
	// Pass true if there is a SIM card in the current device, but it is not possible to
	// check whether phone number matches
	HasUnknownPhoneNumber bool
	// For official applications only. True, if the application can use Android SMS Retriever
	// API (requires Google Play Services >= 10.2) to automatically receive the
	// authentication code from the SMS. See https://developers.google
	// com/identity/sms-retriever/ for more details
	AllowSMSRetrieverAPI bool
	// For official Android and iOS applications only; pass null otherwise. Settings for
	// Firebase Authentication
	FirebaseAuthenticationSettings FirebaseAuthenticationSettingsClass
	// List of up to 20 authentication tokens, recently received in
	// updateOption("authentication_token") in previously logged out sessions
	AuthenticationTokens []string
}

// PhoneNumberAuthenticationSettingsTypeID is TL type id of PhoneNumberAuthenticationSettings.
const PhoneNumberAuthenticationSettingsTypeID = 0x3e1183db

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
	if !(p.AllowMissedCall == false) {
		return false
	}
	if !(p.IsCurrentPhoneNumber == false) {
		return false
	}
	if !(p.HasUnknownPhoneNumber == false) {
		return false
	}
	if !(p.AllowSMSRetrieverAPI == false) {
		return false
	}
	if !(p.FirebaseAuthenticationSettings == nil) {
		return false
	}
	if !(p.AuthenticationTokens == nil) {
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
			Name:       "AllowMissedCall",
			SchemaName: "allow_missed_call",
		},
		{
			Name:       "IsCurrentPhoneNumber",
			SchemaName: "is_current_phone_number",
		},
		{
			Name:       "HasUnknownPhoneNumber",
			SchemaName: "has_unknown_phone_number",
		},
		{
			Name:       "AllowSMSRetrieverAPI",
			SchemaName: "allow_sms_retriever_api",
		},
		{
			Name:       "FirebaseAuthenticationSettings",
			SchemaName: "firebase_authentication_settings",
		},
		{
			Name:       "AuthenticationTokens",
			SchemaName: "authentication_tokens",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhoneNumberAuthenticationSettings) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneNumberAuthenticationSettings#3e1183db as nil")
	}
	b.PutID(PhoneNumberAuthenticationSettingsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhoneNumberAuthenticationSettings) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneNumberAuthenticationSettings#3e1183db as nil")
	}
	b.PutBool(p.AllowFlashCall)
	b.PutBool(p.AllowMissedCall)
	b.PutBool(p.IsCurrentPhoneNumber)
	b.PutBool(p.HasUnknownPhoneNumber)
	b.PutBool(p.AllowSMSRetrieverAPI)
	if p.FirebaseAuthenticationSettings == nil {
		return fmt.Errorf("unable to encode phoneNumberAuthenticationSettings#3e1183db: field firebase_authentication_settings is nil")
	}
	if err := p.FirebaseAuthenticationSettings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phoneNumberAuthenticationSettings#3e1183db: field firebase_authentication_settings: %w", err)
	}
	b.PutInt(len(p.AuthenticationTokens))
	for _, v := range p.AuthenticationTokens {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PhoneNumberAuthenticationSettings) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneNumberAuthenticationSettings#3e1183db to nil")
	}
	if err := b.ConsumeID(PhoneNumberAuthenticationSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhoneNumberAuthenticationSettings) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneNumberAuthenticationSettings#3e1183db to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field allow_flash_call: %w", err)
		}
		p.AllowFlashCall = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field allow_missed_call: %w", err)
		}
		p.AllowMissedCall = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field is_current_phone_number: %w", err)
		}
		p.IsCurrentPhoneNumber = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field has_unknown_phone_number: %w", err)
		}
		p.HasUnknownPhoneNumber = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field allow_sms_retriever_api: %w", err)
		}
		p.AllowSMSRetrieverAPI = value
	}
	{
		value, err := DecodeFirebaseAuthenticationSettings(b)
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field firebase_authentication_settings: %w", err)
		}
		p.FirebaseAuthenticationSettings = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field authentication_tokens: %w", err)
		}

		if headerLen > 0 {
			p.AuthenticationTokens = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field authentication_tokens: %w", err)
			}
			p.AuthenticationTokens = append(p.AuthenticationTokens, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PhoneNumberAuthenticationSettings) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode phoneNumberAuthenticationSettings#3e1183db as nil")
	}
	b.ObjStart()
	b.PutID("phoneNumberAuthenticationSettings")
	b.Comma()
	b.FieldStart("allow_flash_call")
	b.PutBool(p.AllowFlashCall)
	b.Comma()
	b.FieldStart("allow_missed_call")
	b.PutBool(p.AllowMissedCall)
	b.Comma()
	b.FieldStart("is_current_phone_number")
	b.PutBool(p.IsCurrentPhoneNumber)
	b.Comma()
	b.FieldStart("has_unknown_phone_number")
	b.PutBool(p.HasUnknownPhoneNumber)
	b.Comma()
	b.FieldStart("allow_sms_retriever_api")
	b.PutBool(p.AllowSMSRetrieverAPI)
	b.Comma()
	b.FieldStart("firebase_authentication_settings")
	if p.FirebaseAuthenticationSettings == nil {
		return fmt.Errorf("unable to encode phoneNumberAuthenticationSettings#3e1183db: field firebase_authentication_settings is nil")
	}
	if err := p.FirebaseAuthenticationSettings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode phoneNumberAuthenticationSettings#3e1183db: field firebase_authentication_settings: %w", err)
	}
	b.Comma()
	b.FieldStart("authentication_tokens")
	b.ArrStart()
	for _, v := range p.AuthenticationTokens {
		b.PutString(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PhoneNumberAuthenticationSettings) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode phoneNumberAuthenticationSettings#3e1183db to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("phoneNumberAuthenticationSettings"); err != nil {
				return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: %w", err)
			}
		case "allow_flash_call":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field allow_flash_call: %w", err)
			}
			p.AllowFlashCall = value
		case "allow_missed_call":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field allow_missed_call: %w", err)
			}
			p.AllowMissedCall = value
		case "is_current_phone_number":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field is_current_phone_number: %w", err)
			}
			p.IsCurrentPhoneNumber = value
		case "has_unknown_phone_number":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field has_unknown_phone_number: %w", err)
			}
			p.HasUnknownPhoneNumber = value
		case "allow_sms_retriever_api":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field allow_sms_retriever_api: %w", err)
			}
			p.AllowSMSRetrieverAPI = value
		case "firebase_authentication_settings":
			value, err := DecodeTDLibJSONFirebaseAuthenticationSettings(b)
			if err != nil {
				return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field firebase_authentication_settings: %w", err)
			}
			p.FirebaseAuthenticationSettings = value
		case "authentication_tokens":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field authentication_tokens: %w", err)
				}
				p.AuthenticationTokens = append(p.AuthenticationTokens, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode phoneNumberAuthenticationSettings#3e1183db: field authentication_tokens: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAllowFlashCall returns value of AllowFlashCall field.
func (p *PhoneNumberAuthenticationSettings) GetAllowFlashCall() (value bool) {
	if p == nil {
		return
	}
	return p.AllowFlashCall
}

// GetAllowMissedCall returns value of AllowMissedCall field.
func (p *PhoneNumberAuthenticationSettings) GetAllowMissedCall() (value bool) {
	if p == nil {
		return
	}
	return p.AllowMissedCall
}

// GetIsCurrentPhoneNumber returns value of IsCurrentPhoneNumber field.
func (p *PhoneNumberAuthenticationSettings) GetIsCurrentPhoneNumber() (value bool) {
	if p == nil {
		return
	}
	return p.IsCurrentPhoneNumber
}

// GetHasUnknownPhoneNumber returns value of HasUnknownPhoneNumber field.
func (p *PhoneNumberAuthenticationSettings) GetHasUnknownPhoneNumber() (value bool) {
	if p == nil {
		return
	}
	return p.HasUnknownPhoneNumber
}

// GetAllowSMSRetrieverAPI returns value of AllowSMSRetrieverAPI field.
func (p *PhoneNumberAuthenticationSettings) GetAllowSMSRetrieverAPI() (value bool) {
	if p == nil {
		return
	}
	return p.AllowSMSRetrieverAPI
}

// GetFirebaseAuthenticationSettings returns value of FirebaseAuthenticationSettings field.
func (p *PhoneNumberAuthenticationSettings) GetFirebaseAuthenticationSettings() (value FirebaseAuthenticationSettingsClass) {
	if p == nil {
		return
	}
	return p.FirebaseAuthenticationSettings
}

// GetAuthenticationTokens returns value of AuthenticationTokens field.
func (p *PhoneNumberAuthenticationSettings) GetAuthenticationTokens() (value []string) {
	if p == nil {
		return
	}
	return p.AuthenticationTokens
}
