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

// PaymentsPaymentForm represents TL type `payments.paymentForm#1694761b`.
// Payment form
//
// See https://core.telegram.org/constructor/payments.paymentForm for reference.
type PaymentsPaymentForm struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the user can choose to save credentials.
	CanSaveCredentials bool
	// Indicates that the user can save payment credentials, but only after setting up a 2FA
	// password¹ (currently the account doesn't have a 2FA password²)
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	//  2) https://core.telegram.org/api/srp
	PasswordMissing bool
	// FormID field of PaymentsPaymentForm.
	FormID int64
	// Bot ID
	BotID int64
	// Invoice
	Invoice Invoice
	// Payment provider ID.
	ProviderID int64
	// Payment form URL
	URL string
	// Payment provider name.One of the following:- stripe
	//
	// Use SetNativeProvider and GetNativeProvider helpers.
	NativeProvider string
	// Contains information about the payment provider, if available, to support it natively
	// without the need for opening the URL.A JSON object that can contain the following
	// fields:- publishable_key: Stripe API publishable key- apple_pay_merchant_id: Apple Pay
	// merchant ID- android_pay_public_key: Android Pay public key- android_pay_bgcolor:
	// Android Pay form background color- android_pay_inverse: Whether to use the dark theme
	// in the Android Pay form- need_country: True, if the user country must be provided,-
	// need_zip: True, if the user ZIP/postal code must be provided,- need_cardholder_name:
	// True, if the cardholder name must be provided
	//
	// Use SetNativeParams and GetNativeParams helpers.
	NativeParams DataJSON
	// Saved server-side order information
	//
	// Use SetSavedInfo and GetSavedInfo helpers.
	SavedInfo PaymentRequestedInfo
	// Contains information about saved card credentials
	//
	// Use SetSavedCredentials and GetSavedCredentials helpers.
	SavedCredentials PaymentSavedCredentialsCard
	// Users
	Users []UserClass
}

// PaymentsPaymentFormTypeID is TL type id of PaymentsPaymentForm.
const PaymentsPaymentFormTypeID = 0x1694761b

// Ensuring interfaces in compile-time for PaymentsPaymentForm.
var (
	_ bin.Encoder     = &PaymentsPaymentForm{}
	_ bin.Decoder     = &PaymentsPaymentForm{}
	_ bin.BareEncoder = &PaymentsPaymentForm{}
	_ bin.BareDecoder = &PaymentsPaymentForm{}
)

func (p *PaymentsPaymentForm) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.CanSaveCredentials == false) {
		return false
	}
	if !(p.PasswordMissing == false) {
		return false
	}
	if !(p.FormID == 0) {
		return false
	}
	if !(p.BotID == 0) {
		return false
	}
	if !(p.Invoice.Zero()) {
		return false
	}
	if !(p.ProviderID == 0) {
		return false
	}
	if !(p.URL == "") {
		return false
	}
	if !(p.NativeProvider == "") {
		return false
	}
	if !(p.NativeParams.Zero()) {
		return false
	}
	if !(p.SavedInfo.Zero()) {
		return false
	}
	if !(p.SavedCredentials.Zero()) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentsPaymentForm) String() string {
	if p == nil {
		return "PaymentsPaymentForm(nil)"
	}
	type Alias PaymentsPaymentForm
	return fmt.Sprintf("PaymentsPaymentForm%+v", Alias(*p))
}

// FillFrom fills PaymentsPaymentForm from given interface.
func (p *PaymentsPaymentForm) FillFrom(from interface {
	GetCanSaveCredentials() (value bool)
	GetPasswordMissing() (value bool)
	GetFormID() (value int64)
	GetBotID() (value int64)
	GetInvoice() (value Invoice)
	GetProviderID() (value int64)
	GetURL() (value string)
	GetNativeProvider() (value string, ok bool)
	GetNativeParams() (value DataJSON, ok bool)
	GetSavedInfo() (value PaymentRequestedInfo, ok bool)
	GetSavedCredentials() (value PaymentSavedCredentialsCard, ok bool)
	GetUsers() (value []UserClass)
}) {
	p.CanSaveCredentials = from.GetCanSaveCredentials()
	p.PasswordMissing = from.GetPasswordMissing()
	p.FormID = from.GetFormID()
	p.BotID = from.GetBotID()
	p.Invoice = from.GetInvoice()
	p.ProviderID = from.GetProviderID()
	p.URL = from.GetURL()
	if val, ok := from.GetNativeProvider(); ok {
		p.NativeProvider = val
	}

	if val, ok := from.GetNativeParams(); ok {
		p.NativeParams = val
	}

	if val, ok := from.GetSavedInfo(); ok {
		p.SavedInfo = val
	}

	if val, ok := from.GetSavedCredentials(); ok {
		p.SavedCredentials = val
	}

	p.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentsPaymentForm) TypeID() uint32 {
	return PaymentsPaymentFormTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentsPaymentForm) TypeName() string {
	return "payments.paymentForm"
}

// TypeInfo returns info about TL type.
func (p *PaymentsPaymentForm) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "payments.paymentForm",
		ID:   PaymentsPaymentFormTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CanSaveCredentials",
			SchemaName: "can_save_credentials",
			Null:       !p.Flags.Has(2),
		},
		{
			Name:       "PasswordMissing",
			SchemaName: "password_missing",
			Null:       !p.Flags.Has(3),
		},
		{
			Name:       "FormID",
			SchemaName: "form_id",
		},
		{
			Name:       "BotID",
			SchemaName: "bot_id",
		},
		{
			Name:       "Invoice",
			SchemaName: "invoice",
		},
		{
			Name:       "ProviderID",
			SchemaName: "provider_id",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "NativeProvider",
			SchemaName: "native_provider",
			Null:       !p.Flags.Has(4),
		},
		{
			Name:       "NativeParams",
			SchemaName: "native_params",
			Null:       !p.Flags.Has(4),
		},
		{
			Name:       "SavedInfo",
			SchemaName: "saved_info",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "SavedCredentials",
			SchemaName: "saved_credentials",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentsPaymentForm) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentForm#1694761b as nil")
	}
	b.PutID(PaymentsPaymentFormTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentsPaymentForm) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode payments.paymentForm#1694761b as nil")
	}
	if !(p.CanSaveCredentials == false) {
		p.Flags.Set(2)
	}
	if !(p.PasswordMissing == false) {
		p.Flags.Set(3)
	}
	if !(p.NativeProvider == "") {
		p.Flags.Set(4)
	}
	if !(p.NativeParams.Zero()) {
		p.Flags.Set(4)
	}
	if !(p.SavedInfo.Zero()) {
		p.Flags.Set(0)
	}
	if !(p.SavedCredentials.Zero()) {
		p.Flags.Set(1)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.paymentForm#1694761b: field flags: %w", err)
	}
	b.PutLong(p.FormID)
	b.PutLong(p.BotID)
	if err := p.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode payments.paymentForm#1694761b: field invoice: %w", err)
	}
	b.PutLong(p.ProviderID)
	b.PutString(p.URL)
	if p.Flags.Has(4) {
		b.PutString(p.NativeProvider)
	}
	if p.Flags.Has(4) {
		if err := p.NativeParams.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentForm#1694761b: field native_params: %w", err)
		}
	}
	if p.Flags.Has(0) {
		if err := p.SavedInfo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentForm#1694761b: field saved_info: %w", err)
		}
	}
	if p.Flags.Has(1) {
		if err := p.SavedCredentials.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentForm#1694761b: field saved_credentials: %w", err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode payments.paymentForm#1694761b: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode payments.paymentForm#1694761b: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentsPaymentForm) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentForm#1694761b to nil")
	}
	if err := b.ConsumeID(PaymentsPaymentFormTypeID); err != nil {
		return fmt.Errorf("unable to decode payments.paymentForm#1694761b: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentsPaymentForm) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode payments.paymentForm#1694761b to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field flags: %w", err)
		}
	}
	p.CanSaveCredentials = p.Flags.Has(2)
	p.PasswordMissing = p.Flags.Has(3)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field form_id: %w", err)
		}
		p.FormID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field bot_id: %w", err)
		}
		p.BotID = value
	}
	{
		if err := p.Invoice.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field invoice: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field provider_id: %w", err)
		}
		p.ProviderID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field url: %w", err)
		}
		p.URL = value
	}
	if p.Flags.Has(4) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field native_provider: %w", err)
		}
		p.NativeProvider = value
	}
	if p.Flags.Has(4) {
		if err := p.NativeParams.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field native_params: %w", err)
		}
	}
	if p.Flags.Has(0) {
		if err := p.SavedInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field saved_info: %w", err)
		}
	}
	if p.Flags.Has(1) {
		if err := p.SavedCredentials.Decode(b); err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field saved_credentials: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field users: %w", err)
		}

		if headerLen > 0 {
			p.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode payments.paymentForm#1694761b: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// SetCanSaveCredentials sets value of CanSaveCredentials conditional field.
func (p *PaymentsPaymentForm) SetCanSaveCredentials(value bool) {
	if value {
		p.Flags.Set(2)
		p.CanSaveCredentials = true
	} else {
		p.Flags.Unset(2)
		p.CanSaveCredentials = false
	}
}

// GetCanSaveCredentials returns value of CanSaveCredentials conditional field.
func (p *PaymentsPaymentForm) GetCanSaveCredentials() (value bool) {
	return p.Flags.Has(2)
}

// SetPasswordMissing sets value of PasswordMissing conditional field.
func (p *PaymentsPaymentForm) SetPasswordMissing(value bool) {
	if value {
		p.Flags.Set(3)
		p.PasswordMissing = true
	} else {
		p.Flags.Unset(3)
		p.PasswordMissing = false
	}
}

// GetPasswordMissing returns value of PasswordMissing conditional field.
func (p *PaymentsPaymentForm) GetPasswordMissing() (value bool) {
	return p.Flags.Has(3)
}

// GetFormID returns value of FormID field.
func (p *PaymentsPaymentForm) GetFormID() (value int64) {
	return p.FormID
}

// GetBotID returns value of BotID field.
func (p *PaymentsPaymentForm) GetBotID() (value int64) {
	return p.BotID
}

// GetInvoice returns value of Invoice field.
func (p *PaymentsPaymentForm) GetInvoice() (value Invoice) {
	return p.Invoice
}

// GetProviderID returns value of ProviderID field.
func (p *PaymentsPaymentForm) GetProviderID() (value int64) {
	return p.ProviderID
}

// GetURL returns value of URL field.
func (p *PaymentsPaymentForm) GetURL() (value string) {
	return p.URL
}

// SetNativeProvider sets value of NativeProvider conditional field.
func (p *PaymentsPaymentForm) SetNativeProvider(value string) {
	p.Flags.Set(4)
	p.NativeProvider = value
}

// GetNativeProvider returns value of NativeProvider conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentForm) GetNativeProvider() (value string, ok bool) {
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.NativeProvider, true
}

// SetNativeParams sets value of NativeParams conditional field.
func (p *PaymentsPaymentForm) SetNativeParams(value DataJSON) {
	p.Flags.Set(4)
	p.NativeParams = value
}

// GetNativeParams returns value of NativeParams conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentForm) GetNativeParams() (value DataJSON, ok bool) {
	if !p.Flags.Has(4) {
		return value, false
	}
	return p.NativeParams, true
}

// SetSavedInfo sets value of SavedInfo conditional field.
func (p *PaymentsPaymentForm) SetSavedInfo(value PaymentRequestedInfo) {
	p.Flags.Set(0)
	p.SavedInfo = value
}

// GetSavedInfo returns value of SavedInfo conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentForm) GetSavedInfo() (value PaymentRequestedInfo, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.SavedInfo, true
}

// SetSavedCredentials sets value of SavedCredentials conditional field.
func (p *PaymentsPaymentForm) SetSavedCredentials(value PaymentSavedCredentialsCard) {
	p.Flags.Set(1)
	p.SavedCredentials = value
}

// GetSavedCredentials returns value of SavedCredentials conditional field and
// boolean which is true if field was set.
func (p *PaymentsPaymentForm) GetSavedCredentials() (value PaymentSavedCredentialsCard, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.SavedCredentials, true
}

// GetUsers returns value of Users field.
func (p *PaymentsPaymentForm) GetUsers() (value []UserClass) {
	return p.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (p *PaymentsPaymentForm) MapUsers() (value UserClassArray) {
	return UserClassArray(p.Users)
}
