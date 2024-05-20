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

// EmailAddressAuthenticationCode represents TL type `emailAddressAuthenticationCode#c4cc19c2`.
type EmailAddressAuthenticationCode struct {
	// The code
	Code string
}

// EmailAddressAuthenticationCodeTypeID is TL type id of EmailAddressAuthenticationCode.
const EmailAddressAuthenticationCodeTypeID = 0xc4cc19c2

// construct implements constructor of EmailAddressAuthenticationClass.
func (e EmailAddressAuthenticationCode) construct() EmailAddressAuthenticationClass { return &e }

// Ensuring interfaces in compile-time for EmailAddressAuthenticationCode.
var (
	_ bin.Encoder     = &EmailAddressAuthenticationCode{}
	_ bin.Decoder     = &EmailAddressAuthenticationCode{}
	_ bin.BareEncoder = &EmailAddressAuthenticationCode{}
	_ bin.BareDecoder = &EmailAddressAuthenticationCode{}

	_ EmailAddressAuthenticationClass = &EmailAddressAuthenticationCode{}
)

func (e *EmailAddressAuthenticationCode) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Code == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmailAddressAuthenticationCode) String() string {
	if e == nil {
		return "EmailAddressAuthenticationCode(nil)"
	}
	type Alias EmailAddressAuthenticationCode
	return fmt.Sprintf("EmailAddressAuthenticationCode%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmailAddressAuthenticationCode) TypeID() uint32 {
	return EmailAddressAuthenticationCodeTypeID
}

// TypeName returns name of type in TL schema.
func (*EmailAddressAuthenticationCode) TypeName() string {
	return "emailAddressAuthenticationCode"
}

// TypeInfo returns info about TL type.
func (e *EmailAddressAuthenticationCode) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emailAddressAuthenticationCode",
		ID:   EmailAddressAuthenticationCodeTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Code",
			SchemaName: "code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmailAddressAuthenticationCode) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emailAddressAuthenticationCode#c4cc19c2 as nil")
	}
	b.PutID(EmailAddressAuthenticationCodeTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmailAddressAuthenticationCode) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emailAddressAuthenticationCode#c4cc19c2 as nil")
	}
	b.PutString(e.Code)
	return nil
}

// Decode implements bin.Decoder.
func (e *EmailAddressAuthenticationCode) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emailAddressAuthenticationCode#c4cc19c2 to nil")
	}
	if err := b.ConsumeID(EmailAddressAuthenticationCodeTypeID); err != nil {
		return fmt.Errorf("unable to decode emailAddressAuthenticationCode#c4cc19c2: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmailAddressAuthenticationCode) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emailAddressAuthenticationCode#c4cc19c2 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emailAddressAuthenticationCode#c4cc19c2: field code: %w", err)
		}
		e.Code = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EmailAddressAuthenticationCode) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode emailAddressAuthenticationCode#c4cc19c2 as nil")
	}
	b.ObjStart()
	b.PutID("emailAddressAuthenticationCode")
	b.Comma()
	b.FieldStart("code")
	b.PutString(e.Code)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EmailAddressAuthenticationCode) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode emailAddressAuthenticationCode#c4cc19c2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("emailAddressAuthenticationCode"); err != nil {
				return fmt.Errorf("unable to decode emailAddressAuthenticationCode#c4cc19c2: %w", err)
			}
		case "code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emailAddressAuthenticationCode#c4cc19c2: field code: %w", err)
			}
			e.Code = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCode returns value of Code field.
func (e *EmailAddressAuthenticationCode) GetCode() (value string) {
	if e == nil {
		return
	}
	return e.Code
}

// EmailAddressAuthenticationAppleID represents TL type `emailAddressAuthenticationAppleId#25c94869`.
type EmailAddressAuthenticationAppleID struct {
	// The token
	Token string
}

// EmailAddressAuthenticationAppleIDTypeID is TL type id of EmailAddressAuthenticationAppleID.
const EmailAddressAuthenticationAppleIDTypeID = 0x25c94869

// construct implements constructor of EmailAddressAuthenticationClass.
func (e EmailAddressAuthenticationAppleID) construct() EmailAddressAuthenticationClass { return &e }

// Ensuring interfaces in compile-time for EmailAddressAuthenticationAppleID.
var (
	_ bin.Encoder     = &EmailAddressAuthenticationAppleID{}
	_ bin.Decoder     = &EmailAddressAuthenticationAppleID{}
	_ bin.BareEncoder = &EmailAddressAuthenticationAppleID{}
	_ bin.BareDecoder = &EmailAddressAuthenticationAppleID{}

	_ EmailAddressAuthenticationClass = &EmailAddressAuthenticationAppleID{}
)

func (e *EmailAddressAuthenticationAppleID) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Token == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmailAddressAuthenticationAppleID) String() string {
	if e == nil {
		return "EmailAddressAuthenticationAppleID(nil)"
	}
	type Alias EmailAddressAuthenticationAppleID
	return fmt.Sprintf("EmailAddressAuthenticationAppleID%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmailAddressAuthenticationAppleID) TypeID() uint32 {
	return EmailAddressAuthenticationAppleIDTypeID
}

// TypeName returns name of type in TL schema.
func (*EmailAddressAuthenticationAppleID) TypeName() string {
	return "emailAddressAuthenticationAppleId"
}

// TypeInfo returns info about TL type.
func (e *EmailAddressAuthenticationAppleID) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emailAddressAuthenticationAppleId",
		ID:   EmailAddressAuthenticationAppleIDTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Token",
			SchemaName: "token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmailAddressAuthenticationAppleID) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emailAddressAuthenticationAppleId#25c94869 as nil")
	}
	b.PutID(EmailAddressAuthenticationAppleIDTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmailAddressAuthenticationAppleID) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emailAddressAuthenticationAppleId#25c94869 as nil")
	}
	b.PutString(e.Token)
	return nil
}

// Decode implements bin.Decoder.
func (e *EmailAddressAuthenticationAppleID) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emailAddressAuthenticationAppleId#25c94869 to nil")
	}
	if err := b.ConsumeID(EmailAddressAuthenticationAppleIDTypeID); err != nil {
		return fmt.Errorf("unable to decode emailAddressAuthenticationAppleId#25c94869: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmailAddressAuthenticationAppleID) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emailAddressAuthenticationAppleId#25c94869 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emailAddressAuthenticationAppleId#25c94869: field token: %w", err)
		}
		e.Token = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EmailAddressAuthenticationAppleID) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode emailAddressAuthenticationAppleId#25c94869 as nil")
	}
	b.ObjStart()
	b.PutID("emailAddressAuthenticationAppleId")
	b.Comma()
	b.FieldStart("token")
	b.PutString(e.Token)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EmailAddressAuthenticationAppleID) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode emailAddressAuthenticationAppleId#25c94869 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("emailAddressAuthenticationAppleId"); err != nil {
				return fmt.Errorf("unable to decode emailAddressAuthenticationAppleId#25c94869: %w", err)
			}
		case "token":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emailAddressAuthenticationAppleId#25c94869: field token: %w", err)
			}
			e.Token = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetToken returns value of Token field.
func (e *EmailAddressAuthenticationAppleID) GetToken() (value string) {
	if e == nil {
		return
	}
	return e.Token
}

// EmailAddressAuthenticationGoogleID represents TL type `emailAddressAuthenticationGoogleId#fedbe742`.
type EmailAddressAuthenticationGoogleID struct {
	// The token
	Token string
}

// EmailAddressAuthenticationGoogleIDTypeID is TL type id of EmailAddressAuthenticationGoogleID.
const EmailAddressAuthenticationGoogleIDTypeID = 0xfedbe742

// construct implements constructor of EmailAddressAuthenticationClass.
func (e EmailAddressAuthenticationGoogleID) construct() EmailAddressAuthenticationClass { return &e }

// Ensuring interfaces in compile-time for EmailAddressAuthenticationGoogleID.
var (
	_ bin.Encoder     = &EmailAddressAuthenticationGoogleID{}
	_ bin.Decoder     = &EmailAddressAuthenticationGoogleID{}
	_ bin.BareEncoder = &EmailAddressAuthenticationGoogleID{}
	_ bin.BareDecoder = &EmailAddressAuthenticationGoogleID{}

	_ EmailAddressAuthenticationClass = &EmailAddressAuthenticationGoogleID{}
)

func (e *EmailAddressAuthenticationGoogleID) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Token == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmailAddressAuthenticationGoogleID) String() string {
	if e == nil {
		return "EmailAddressAuthenticationGoogleID(nil)"
	}
	type Alias EmailAddressAuthenticationGoogleID
	return fmt.Sprintf("EmailAddressAuthenticationGoogleID%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmailAddressAuthenticationGoogleID) TypeID() uint32 {
	return EmailAddressAuthenticationGoogleIDTypeID
}

// TypeName returns name of type in TL schema.
func (*EmailAddressAuthenticationGoogleID) TypeName() string {
	return "emailAddressAuthenticationGoogleId"
}

// TypeInfo returns info about TL type.
func (e *EmailAddressAuthenticationGoogleID) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emailAddressAuthenticationGoogleId",
		ID:   EmailAddressAuthenticationGoogleIDTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Token",
			SchemaName: "token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmailAddressAuthenticationGoogleID) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emailAddressAuthenticationGoogleId#fedbe742 as nil")
	}
	b.PutID(EmailAddressAuthenticationGoogleIDTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmailAddressAuthenticationGoogleID) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emailAddressAuthenticationGoogleId#fedbe742 as nil")
	}
	b.PutString(e.Token)
	return nil
}

// Decode implements bin.Decoder.
func (e *EmailAddressAuthenticationGoogleID) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emailAddressAuthenticationGoogleId#fedbe742 to nil")
	}
	if err := b.ConsumeID(EmailAddressAuthenticationGoogleIDTypeID); err != nil {
		return fmt.Errorf("unable to decode emailAddressAuthenticationGoogleId#fedbe742: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmailAddressAuthenticationGoogleID) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emailAddressAuthenticationGoogleId#fedbe742 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode emailAddressAuthenticationGoogleId#fedbe742: field token: %w", err)
		}
		e.Token = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EmailAddressAuthenticationGoogleID) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode emailAddressAuthenticationGoogleId#fedbe742 as nil")
	}
	b.ObjStart()
	b.PutID("emailAddressAuthenticationGoogleId")
	b.Comma()
	b.FieldStart("token")
	b.PutString(e.Token)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EmailAddressAuthenticationGoogleID) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode emailAddressAuthenticationGoogleId#fedbe742 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("emailAddressAuthenticationGoogleId"); err != nil {
				return fmt.Errorf("unable to decode emailAddressAuthenticationGoogleId#fedbe742: %w", err)
			}
		case "token":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode emailAddressAuthenticationGoogleId#fedbe742: field token: %w", err)
			}
			e.Token = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetToken returns value of Token field.
func (e *EmailAddressAuthenticationGoogleID) GetToken() (value string) {
	if e == nil {
		return
	}
	return e.Token
}

// EmailAddressAuthenticationClassName is schema name of EmailAddressAuthenticationClass.
const EmailAddressAuthenticationClassName = "EmailAddressAuthentication"

// EmailAddressAuthenticationClass represents EmailAddressAuthentication generic type.
//
// Example:
//
//	g, err := tdapi.DecodeEmailAddressAuthentication(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.EmailAddressAuthenticationCode: // emailAddressAuthenticationCode#c4cc19c2
//	case *tdapi.EmailAddressAuthenticationAppleID: // emailAddressAuthenticationAppleId#25c94869
//	case *tdapi.EmailAddressAuthenticationGoogleID: // emailAddressAuthenticationGoogleId#fedbe742
//	default: panic(v)
//	}
type EmailAddressAuthenticationClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() EmailAddressAuthenticationClass

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

// DecodeEmailAddressAuthentication implements binary de-serialization for EmailAddressAuthenticationClass.
func DecodeEmailAddressAuthentication(buf *bin.Buffer) (EmailAddressAuthenticationClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case EmailAddressAuthenticationCodeTypeID:
		// Decoding emailAddressAuthenticationCode#c4cc19c2.
		v := EmailAddressAuthenticationCode{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmailAddressAuthenticationClass: %w", err)
		}
		return &v, nil
	case EmailAddressAuthenticationAppleIDTypeID:
		// Decoding emailAddressAuthenticationAppleId#25c94869.
		v := EmailAddressAuthenticationAppleID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmailAddressAuthenticationClass: %w", err)
		}
		return &v, nil
	case EmailAddressAuthenticationGoogleIDTypeID:
		// Decoding emailAddressAuthenticationGoogleId#fedbe742.
		v := EmailAddressAuthenticationGoogleID{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmailAddressAuthenticationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EmailAddressAuthenticationClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONEmailAddressAuthentication implements binary de-serialization for EmailAddressAuthenticationClass.
func DecodeTDLibJSONEmailAddressAuthentication(buf tdjson.Decoder) (EmailAddressAuthenticationClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "emailAddressAuthenticationCode":
		// Decoding emailAddressAuthenticationCode#c4cc19c2.
		v := EmailAddressAuthenticationCode{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmailAddressAuthenticationClass: %w", err)
		}
		return &v, nil
	case "emailAddressAuthenticationAppleId":
		// Decoding emailAddressAuthenticationAppleId#25c94869.
		v := EmailAddressAuthenticationAppleID{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmailAddressAuthenticationClass: %w", err)
		}
		return &v, nil
	case "emailAddressAuthenticationGoogleId":
		// Decoding emailAddressAuthenticationGoogleId#fedbe742.
		v := EmailAddressAuthenticationGoogleID{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode EmailAddressAuthenticationClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode EmailAddressAuthenticationClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// EmailAddressAuthentication boxes the EmailAddressAuthenticationClass providing a helper.
type EmailAddressAuthenticationBox struct {
	EmailAddressAuthentication EmailAddressAuthenticationClass
}

// Decode implements bin.Decoder for EmailAddressAuthenticationBox.
func (b *EmailAddressAuthenticationBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode EmailAddressAuthenticationBox to nil")
	}
	v, err := DecodeEmailAddressAuthentication(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EmailAddressAuthentication = v
	return nil
}

// Encode implements bin.Encode for EmailAddressAuthenticationBox.
func (b *EmailAddressAuthenticationBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.EmailAddressAuthentication == nil {
		return fmt.Errorf("unable to encode EmailAddressAuthenticationClass as nil")
	}
	return b.EmailAddressAuthentication.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for EmailAddressAuthenticationBox.
func (b *EmailAddressAuthenticationBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode EmailAddressAuthenticationBox to nil")
	}
	v, err := DecodeTDLibJSONEmailAddressAuthentication(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.EmailAddressAuthentication = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for EmailAddressAuthenticationBox.
func (b *EmailAddressAuthenticationBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.EmailAddressAuthentication == nil {
		return fmt.Errorf("unable to encode EmailAddressAuthenticationClass as nil")
	}
	return b.EmailAddressAuthentication.EncodeTDLibJSON(buf)
}
