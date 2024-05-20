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

// PaymentProviderSmartGlocal represents TL type `paymentProviderSmartGlocal#ba047774`.
type PaymentProviderSmartGlocal struct {
	// Public payment token
	PublicToken string
	// URL for sending card tokenization requests
	TokenizeURL string
}

// PaymentProviderSmartGlocalTypeID is TL type id of PaymentProviderSmartGlocal.
const PaymentProviderSmartGlocalTypeID = 0xba047774

// construct implements constructor of PaymentProviderClass.
func (p PaymentProviderSmartGlocal) construct() PaymentProviderClass { return &p }

// Ensuring interfaces in compile-time for PaymentProviderSmartGlocal.
var (
	_ bin.Encoder     = &PaymentProviderSmartGlocal{}
	_ bin.Decoder     = &PaymentProviderSmartGlocal{}
	_ bin.BareEncoder = &PaymentProviderSmartGlocal{}
	_ bin.BareDecoder = &PaymentProviderSmartGlocal{}

	_ PaymentProviderClass = &PaymentProviderSmartGlocal{}
)

func (p *PaymentProviderSmartGlocal) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PublicToken == "") {
		return false
	}
	if !(p.TokenizeURL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentProviderSmartGlocal) String() string {
	if p == nil {
		return "PaymentProviderSmartGlocal(nil)"
	}
	type Alias PaymentProviderSmartGlocal
	return fmt.Sprintf("PaymentProviderSmartGlocal%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentProviderSmartGlocal) TypeID() uint32 {
	return PaymentProviderSmartGlocalTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentProviderSmartGlocal) TypeName() string {
	return "paymentProviderSmartGlocal"
}

// TypeInfo returns info about TL type.
func (p *PaymentProviderSmartGlocal) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentProviderSmartGlocal",
		ID:   PaymentProviderSmartGlocalTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PublicToken",
			SchemaName: "public_token",
		},
		{
			Name:       "TokenizeURL",
			SchemaName: "tokenize_url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentProviderSmartGlocal) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentProviderSmartGlocal#ba047774 as nil")
	}
	b.PutID(PaymentProviderSmartGlocalTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentProviderSmartGlocal) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentProviderSmartGlocal#ba047774 as nil")
	}
	b.PutString(p.PublicToken)
	b.PutString(p.TokenizeURL)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentProviderSmartGlocal) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentProviderSmartGlocal#ba047774 to nil")
	}
	if err := b.ConsumeID(PaymentProviderSmartGlocalTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentProviderSmartGlocal#ba047774: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentProviderSmartGlocal) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentProviderSmartGlocal#ba047774 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentProviderSmartGlocal#ba047774: field public_token: %w", err)
		}
		p.PublicToken = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentProviderSmartGlocal#ba047774: field tokenize_url: %w", err)
		}
		p.TokenizeURL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PaymentProviderSmartGlocal) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentProviderSmartGlocal#ba047774 as nil")
	}
	b.ObjStart()
	b.PutID("paymentProviderSmartGlocal")
	b.Comma()
	b.FieldStart("public_token")
	b.PutString(p.PublicToken)
	b.Comma()
	b.FieldStart("tokenize_url")
	b.PutString(p.TokenizeURL)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PaymentProviderSmartGlocal) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentProviderSmartGlocal#ba047774 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("paymentProviderSmartGlocal"); err != nil {
				return fmt.Errorf("unable to decode paymentProviderSmartGlocal#ba047774: %w", err)
			}
		case "public_token":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode paymentProviderSmartGlocal#ba047774: field public_token: %w", err)
			}
			p.PublicToken = value
		case "tokenize_url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode paymentProviderSmartGlocal#ba047774: field tokenize_url: %w", err)
			}
			p.TokenizeURL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPublicToken returns value of PublicToken field.
func (p *PaymentProviderSmartGlocal) GetPublicToken() (value string) {
	if p == nil {
		return
	}
	return p.PublicToken
}

// GetTokenizeURL returns value of TokenizeURL field.
func (p *PaymentProviderSmartGlocal) GetTokenizeURL() (value string) {
	if p == nil {
		return
	}
	return p.TokenizeURL
}

// PaymentProviderStripe represents TL type `paymentProviderStripe#1614e19b`.
type PaymentProviderStripe struct {
	// Stripe API publishable key
	PublishableKey string
	// True, if the user country must be provided
	NeedCountry bool
	// True, if the user ZIP/postal code must be provided
	NeedPostalCode bool
	// True, if the cardholder name must be provided
	NeedCardholderName bool
}

// PaymentProviderStripeTypeID is TL type id of PaymentProviderStripe.
const PaymentProviderStripeTypeID = 0x1614e19b

// construct implements constructor of PaymentProviderClass.
func (p PaymentProviderStripe) construct() PaymentProviderClass { return &p }

// Ensuring interfaces in compile-time for PaymentProviderStripe.
var (
	_ bin.Encoder     = &PaymentProviderStripe{}
	_ bin.Decoder     = &PaymentProviderStripe{}
	_ bin.BareEncoder = &PaymentProviderStripe{}
	_ bin.BareDecoder = &PaymentProviderStripe{}

	_ PaymentProviderClass = &PaymentProviderStripe{}
)

func (p *PaymentProviderStripe) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.PublishableKey == "") {
		return false
	}
	if !(p.NeedCountry == false) {
		return false
	}
	if !(p.NeedPostalCode == false) {
		return false
	}
	if !(p.NeedCardholderName == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentProviderStripe) String() string {
	if p == nil {
		return "PaymentProviderStripe(nil)"
	}
	type Alias PaymentProviderStripe
	return fmt.Sprintf("PaymentProviderStripe%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentProviderStripe) TypeID() uint32 {
	return PaymentProviderStripeTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentProviderStripe) TypeName() string {
	return "paymentProviderStripe"
}

// TypeInfo returns info about TL type.
func (p *PaymentProviderStripe) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentProviderStripe",
		ID:   PaymentProviderStripeTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PublishableKey",
			SchemaName: "publishable_key",
		},
		{
			Name:       "NeedCountry",
			SchemaName: "need_country",
		},
		{
			Name:       "NeedPostalCode",
			SchemaName: "need_postal_code",
		},
		{
			Name:       "NeedCardholderName",
			SchemaName: "need_cardholder_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentProviderStripe) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentProviderStripe#1614e19b as nil")
	}
	b.PutID(PaymentProviderStripeTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentProviderStripe) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentProviderStripe#1614e19b as nil")
	}
	b.PutString(p.PublishableKey)
	b.PutBool(p.NeedCountry)
	b.PutBool(p.NeedPostalCode)
	b.PutBool(p.NeedCardholderName)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentProviderStripe) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentProviderStripe#1614e19b to nil")
	}
	if err := b.ConsumeID(PaymentProviderStripeTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentProviderStripe#1614e19b: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentProviderStripe) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentProviderStripe#1614e19b to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentProviderStripe#1614e19b: field publishable_key: %w", err)
		}
		p.PublishableKey = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode paymentProviderStripe#1614e19b: field need_country: %w", err)
		}
		p.NeedCountry = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode paymentProviderStripe#1614e19b: field need_postal_code: %w", err)
		}
		p.NeedPostalCode = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode paymentProviderStripe#1614e19b: field need_cardholder_name: %w", err)
		}
		p.NeedCardholderName = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PaymentProviderStripe) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentProviderStripe#1614e19b as nil")
	}
	b.ObjStart()
	b.PutID("paymentProviderStripe")
	b.Comma()
	b.FieldStart("publishable_key")
	b.PutString(p.PublishableKey)
	b.Comma()
	b.FieldStart("need_country")
	b.PutBool(p.NeedCountry)
	b.Comma()
	b.FieldStart("need_postal_code")
	b.PutBool(p.NeedPostalCode)
	b.Comma()
	b.FieldStart("need_cardholder_name")
	b.PutBool(p.NeedCardholderName)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PaymentProviderStripe) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentProviderStripe#1614e19b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("paymentProviderStripe"); err != nil {
				return fmt.Errorf("unable to decode paymentProviderStripe#1614e19b: %w", err)
			}
		case "publishable_key":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode paymentProviderStripe#1614e19b: field publishable_key: %w", err)
			}
			p.PublishableKey = value
		case "need_country":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode paymentProviderStripe#1614e19b: field need_country: %w", err)
			}
			p.NeedCountry = value
		case "need_postal_code":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode paymentProviderStripe#1614e19b: field need_postal_code: %w", err)
			}
			p.NeedPostalCode = value
		case "need_cardholder_name":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode paymentProviderStripe#1614e19b: field need_cardholder_name: %w", err)
			}
			p.NeedCardholderName = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPublishableKey returns value of PublishableKey field.
func (p *PaymentProviderStripe) GetPublishableKey() (value string) {
	if p == nil {
		return
	}
	return p.PublishableKey
}

// GetNeedCountry returns value of NeedCountry field.
func (p *PaymentProviderStripe) GetNeedCountry() (value bool) {
	if p == nil {
		return
	}
	return p.NeedCountry
}

// GetNeedPostalCode returns value of NeedPostalCode field.
func (p *PaymentProviderStripe) GetNeedPostalCode() (value bool) {
	if p == nil {
		return
	}
	return p.NeedPostalCode
}

// GetNeedCardholderName returns value of NeedCardholderName field.
func (p *PaymentProviderStripe) GetNeedCardholderName() (value bool) {
	if p == nil {
		return
	}
	return p.NeedCardholderName
}

// PaymentProviderOther represents TL type `paymentProviderOther#b050e0e4`.
type PaymentProviderOther struct {
	// Payment form URL
	URL string
}

// PaymentProviderOtherTypeID is TL type id of PaymentProviderOther.
const PaymentProviderOtherTypeID = 0xb050e0e4

// construct implements constructor of PaymentProviderClass.
func (p PaymentProviderOther) construct() PaymentProviderClass { return &p }

// Ensuring interfaces in compile-time for PaymentProviderOther.
var (
	_ bin.Encoder     = &PaymentProviderOther{}
	_ bin.Decoder     = &PaymentProviderOther{}
	_ bin.BareEncoder = &PaymentProviderOther{}
	_ bin.BareDecoder = &PaymentProviderOther{}

	_ PaymentProviderClass = &PaymentProviderOther{}
)

func (p *PaymentProviderOther) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PaymentProviderOther) String() string {
	if p == nil {
		return "PaymentProviderOther(nil)"
	}
	type Alias PaymentProviderOther
	return fmt.Sprintf("PaymentProviderOther%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PaymentProviderOther) TypeID() uint32 {
	return PaymentProviderOtherTypeID
}

// TypeName returns name of type in TL schema.
func (*PaymentProviderOther) TypeName() string {
	return "paymentProviderOther"
}

// TypeInfo returns info about TL type.
func (p *PaymentProviderOther) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "paymentProviderOther",
		ID:   PaymentProviderOtherTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PaymentProviderOther) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentProviderOther#b050e0e4 as nil")
	}
	b.PutID(PaymentProviderOtherTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PaymentProviderOther) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentProviderOther#b050e0e4 as nil")
	}
	b.PutString(p.URL)
	return nil
}

// Decode implements bin.Decoder.
func (p *PaymentProviderOther) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentProviderOther#b050e0e4 to nil")
	}
	if err := b.ConsumeID(PaymentProviderOtherTypeID); err != nil {
		return fmt.Errorf("unable to decode paymentProviderOther#b050e0e4: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PaymentProviderOther) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentProviderOther#b050e0e4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode paymentProviderOther#b050e0e4: field url: %w", err)
		}
		p.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PaymentProviderOther) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode paymentProviderOther#b050e0e4 as nil")
	}
	b.ObjStart()
	b.PutID("paymentProviderOther")
	b.Comma()
	b.FieldStart("url")
	b.PutString(p.URL)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PaymentProviderOther) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode paymentProviderOther#b050e0e4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("paymentProviderOther"); err != nil {
				return fmt.Errorf("unable to decode paymentProviderOther#b050e0e4: %w", err)
			}
		case "url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode paymentProviderOther#b050e0e4: field url: %w", err)
			}
			p.URL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetURL returns value of URL field.
func (p *PaymentProviderOther) GetURL() (value string) {
	if p == nil {
		return
	}
	return p.URL
}

// PaymentProviderClassName is schema name of PaymentProviderClass.
const PaymentProviderClassName = "PaymentProvider"

// PaymentProviderClass represents PaymentProvider generic type.
//
// Example:
//
//	g, err := tdapi.DecodePaymentProvider(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.PaymentProviderSmartGlocal: // paymentProviderSmartGlocal#ba047774
//	case *tdapi.PaymentProviderStripe: // paymentProviderStripe#1614e19b
//	case *tdapi.PaymentProviderOther: // paymentProviderOther#b050e0e4
//	default: panic(v)
//	}
type PaymentProviderClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PaymentProviderClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodePaymentProvider implements binary de-serialization for PaymentProviderClass.
func DecodePaymentProvider(buf *bin.Buffer) (PaymentProviderClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PaymentProviderSmartGlocalTypeID:
		// Decoding paymentProviderSmartGlocal#ba047774.
		v := PaymentProviderSmartGlocal{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentProviderClass: %w", err)
		}
		return &v, nil
	case PaymentProviderStripeTypeID:
		// Decoding paymentProviderStripe#1614e19b.
		v := PaymentProviderStripe{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentProviderClass: %w", err)
		}
		return &v, nil
	case PaymentProviderOtherTypeID:
		// Decoding paymentProviderOther#b050e0e4.
		v := PaymentProviderOther{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentProviderClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PaymentProviderClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONPaymentProvider implements binary de-serialization for PaymentProviderClass.
func DecodeTDLibJSONPaymentProvider(buf tdjson.Decoder) (PaymentProviderClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "paymentProviderSmartGlocal":
		// Decoding paymentProviderSmartGlocal#ba047774.
		v := PaymentProviderSmartGlocal{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentProviderClass: %w", err)
		}
		return &v, nil
	case "paymentProviderStripe":
		// Decoding paymentProviderStripe#1614e19b.
		v := PaymentProviderStripe{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentProviderClass: %w", err)
		}
		return &v, nil
	case "paymentProviderOther":
		// Decoding paymentProviderOther#b050e0e4.
		v := PaymentProviderOther{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PaymentProviderClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PaymentProviderClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// PaymentProvider boxes the PaymentProviderClass providing a helper.
type PaymentProviderBox struct {
	PaymentProvider PaymentProviderClass
}

// Decode implements bin.Decoder for PaymentProviderBox.
func (b *PaymentProviderBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PaymentProviderBox to nil")
	}
	v, err := DecodePaymentProvider(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PaymentProvider = v
	return nil
}

// Encode implements bin.Encode for PaymentProviderBox.
func (b *PaymentProviderBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.PaymentProvider == nil {
		return fmt.Errorf("unable to encode PaymentProviderClass as nil")
	}
	return b.PaymentProvider.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for PaymentProviderBox.
func (b *PaymentProviderBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode PaymentProviderBox to nil")
	}
	v, err := DecodeTDLibJSONPaymentProvider(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.PaymentProvider = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for PaymentProviderBox.
func (b *PaymentProviderBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.PaymentProvider == nil {
		return fmt.Errorf("unable to encode PaymentProviderClass as nil")
	}
	return b.PaymentProvider.EncodeTDLibJSON(buf)
}
