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

// PaymentReceipt represents TL type `paymentReceipt#7a9a2ea4`.
type PaymentReceipt struct {
	// Product title
	Title string
	// Contains information about a successful payment
	Description string
	// Product photo; may be null
	Photo Photo
	// Point in time (Unix timestamp) when the payment was made
	Date int32
	// User identifier of the seller bot
	SellerBotUserID int32
	// User identifier of the payment provider bot
	PaymentsProviderUserID int32
	// Contains information about the invoice
	Invoice Invoice
	// Order information; may be null
	OrderInfo OrderInfo
	// Chosen shipping option; may be null
	ShippingOption ShippingOption
	// Title of the saved credentials chosen by the buyer
	CredentialsTitle string
	// The amount of tip chosen by the buyer in the smallest units of the currency
	TipAmount int64
}

// PaymentReceiptTypeID is TL type id of PaymentReceipt.
const PaymentReceiptTypeID = 0x7a9a2ea4

// Ensuring interfaces in compile-time for PaymentReceipt.
var (
	_ bin.Encoder     = &PaymentReceipt{}
	_ bin.Decoder     = &PaymentReceipt{}
	_ bin.BareEncoder = &PaymentReceipt{}
	_ bin.BareDecoder = &PaymentReceipt{}
)

func (p *PaymentReceipt) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Title == "") {
		return false
	}
	if !(p.Description == "") {
		return false
	}
	if !(p.Photo.Zero()) {
		return false
	}
	if !(p.Date == 0) {
		return false
	}
	if !(p.SellerBotUserID == 0) {
		return false
	}
	if !(p.PaymentsProviderUserID == 0) {
		return false
	}
	if !(p.Invoice.Zero()) {
		return false
	}
	if !(p.OrderInfo.Zero()) {
		return false
	}
	if !(p.ShippingOption.Zero()) {
		return false
	}
	if !(p.CredentialsTitle == "") {
		return false
	}
	if !(p.TipAmount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentReceipt) String() string {
	if p == nil {
		return "PaymentReceipt(nil)"
	}
	type Alias PaymentReceipt
	return fmt.Sprintf("PaymentReceipt%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentReceipt) TypeID() uint32 {
	return PaymentReceiptTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentReceipt) TypeName() string {
	return "paymentReceipt"
}

// TypeInfo returns info about TL type.
func (p *PaymentReceipt) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentReceipt",
		ID:   PaymentReceiptTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "SellerBotUserID",
			SchemaName: "seller_bot_user_id",
		},
		{
			Name:       "PaymentsProviderUserID",
			SchemaName: "payments_provider_user_id",
		},
		{
			Name:       "Invoice",
			SchemaName: "invoice",
		},
		{
			Name:       "OrderInfo",
			SchemaName: "order_info",
		},
		{
			Name:       "ShippingOption",
			SchemaName: "shipping_option",
		},
		{
			Name:       "CredentialsTitle",
			SchemaName: "credentials_title",
		},
		{
			Name:       "TipAmount",
			SchemaName: "tip_amount",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentReceipt) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentReceipt#7a9a2ea4 as nil")
	}
	b.PutID(PaymentReceiptTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentReceipt) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentReceipt#7a9a2ea4 as nil")
	}
	b.PutString(p.Title)
	b.PutString(p.Description)
	if err := p.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#7a9a2ea4: field photo: %w", err)
	}
	b.PutInt32(p.Date)
	b.PutInt32(p.SellerBotUserID)
	b.PutInt32(p.PaymentsProviderUserID)
	if err := p.Invoice.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#7a9a2ea4: field invoice: %w", err)
	}
	if err := p.OrderInfo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#7a9a2ea4: field order_info: %w", err)
	}
	if err := p.ShippingOption.Encode(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#7a9a2ea4: field shipping_option: %w", err)
	}
	b.PutString(p.CredentialsTitle)
	b.PutLong(p.TipAmount)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentReceipt) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentReceipt#7a9a2ea4 to nil")
	}
	if err := b.ConsumeID(PaymentReceiptTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentReceipt) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentReceipt#7a9a2ea4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field title: %w", err)
		}
		p.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field description: %w", err)
		}
		p.Description = value
	}
	{
		if err := p.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field photo: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field date: %w", err)
		}
		p.Date = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field seller_bot_user_id: %w", err)
		}
		p.SellerBotUserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field payments_provider_user_id: %w", err)
		}
		p.PaymentsProviderUserID = value
	}
	{
		if err := p.Invoice.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field invoice: %w", err)
		}
	}
	{
		if err := p.OrderInfo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field order_info: %w", err)
		}
	}
	{
		if err := p.ShippingOption.Decode(b); err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field shipping_option: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field credentials_title: %w", err)
		}
		p.CredentialsTitle = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode paymentReceipt#7a9a2ea4: field tip_amount: %w", err)
		}
		p.TipAmount = value
	}
	return nil
}

// EncodeTDLibJSON encodes p in TDLib API JSON format.
func (p *PaymentReceipt) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentReceipt#7a9a2ea4 as nil")
	}
	b.ObjStart()
	b.PutID("paymentReceipt")
	b.FieldStart("title")
	b.PutString(p.Title)
	b.FieldStart("description")
	b.PutString(p.Description)
	b.FieldStart("photo")
	if err := p.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#7a9a2ea4: field photo: %w", err)
	}
	b.FieldStart("date")
	b.PutInt32(p.Date)
	b.FieldStart("seller_bot_user_id")
	b.PutInt32(p.SellerBotUserID)
	b.FieldStart("payments_provider_user_id")
	b.PutInt32(p.PaymentsProviderUserID)
	b.FieldStart("invoice")
	if err := p.Invoice.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#7a9a2ea4: field invoice: %w", err)
	}
	b.FieldStart("order_info")
	if err := p.OrderInfo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#7a9a2ea4: field order_info: %w", err)
	}
	b.FieldStart("shipping_option")
	if err := p.ShippingOption.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode paymentReceipt#7a9a2ea4: field shipping_option: %w", err)
	}
	b.FieldStart("credentials_title")
	b.PutString(p.CredentialsTitle)
	b.FieldStart("tip_amount")
	b.PutLong(p.TipAmount)
	b.ObjEnd()
	return nil
}

// GetTitle returns value of Title field.
func (p *PaymentReceipt) GetTitle() (value string) {
	return p.Title
}

// GetDescription returns value of Description field.
func (p *PaymentReceipt) GetDescription() (value string) {
	return p.Description
}

// GetPhoto returns value of Photo field.
func (p *PaymentReceipt) GetPhoto() (value Photo) {
	return p.Photo
}

// GetDate returns value of Date field.
func (p *PaymentReceipt) GetDate() (value int32) {
	return p.Date
}

// GetSellerBotUserID returns value of SellerBotUserID field.
func (p *PaymentReceipt) GetSellerBotUserID() (value int32) {
	return p.SellerBotUserID
}

// GetPaymentsProviderUserID returns value of PaymentsProviderUserID field.
func (p *PaymentReceipt) GetPaymentsProviderUserID() (value int32) {
	return p.PaymentsProviderUserID
}

// GetInvoice returns value of Invoice field.
func (p *PaymentReceipt) GetInvoice() (value Invoice) {
	return p.Invoice
}

// GetOrderInfo returns value of OrderInfo field.
func (p *PaymentReceipt) GetOrderInfo() (value OrderInfo) {
	return p.OrderInfo
}

// GetShippingOption returns value of ShippingOption field.
func (p *PaymentReceipt) GetShippingOption() (value ShippingOption) {
	return p.ShippingOption
}

// GetCredentialsTitle returns value of CredentialsTitle field.
func (p *PaymentReceipt) GetCredentialsTitle() (value string) {
	return p.CredentialsTitle
}

// GetTipAmount returns value of TipAmount field.
func (p *PaymentReceipt) GetTipAmount() (value int64) {
	return p.TipAmount
}
