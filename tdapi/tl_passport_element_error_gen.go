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

// PassportElementError represents TL type `passportElementError#91059fc5`.
type PassportElementError struct {
	// Type of the Telegram Passport element which has the error
	Type PassportElementTypeClass
	// Error message
	Message string
	// Error source
	Source PassportElementErrorSourceClass
}

// PassportElementErrorTypeID is TL type id of PassportElementError.
const PassportElementErrorTypeID = 0x91059fc5

// Ensuring interfaces in compile-time for PassportElementError.
var (
	_ bin.Encoder     = &PassportElementError{}
	_ bin.Decoder     = &PassportElementError{}
	_ bin.BareEncoder = &PassportElementError{}
	_ bin.BareDecoder = &PassportElementError{}
)

func (p *PassportElementError) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Type == nil) {
		return false
	}
	if !(p.Message == "") {
		return false
	}
	if !(p.Source == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PassportElementError) String() string {
	if p == nil {
		return "PassportElementError(nil)"
	}
	type Alias PassportElementError
	return fmt.Sprintf("PassportElementError%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PassportElementError) TypeID() uint32 {
	return PassportElementErrorTypeID
}

// TypeName returns name of type in TL schema.
func (*PassportElementError) TypeName() string {
	return "passportElementError"
}

// TypeInfo returns info about TL type.
func (p *PassportElementError) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "passportElementError",
		ID:   PassportElementErrorTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Source",
			SchemaName: "source",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PassportElementError) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passportElementError#91059fc5 as nil")
	}
	b.PutID(PassportElementErrorTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PassportElementError) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passportElementError#91059fc5 as nil")
	}
	if p.Type == nil {
		return fmt.Errorf("unable to encode passportElementError#91059fc5: field type is nil")
	}
	if err := p.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode passportElementError#91059fc5: field type: %w", err)
	}
	b.PutString(p.Message)
	if p.Source == nil {
		return fmt.Errorf("unable to encode passportElementError#91059fc5: field source is nil")
	}
	if err := p.Source.Encode(b); err != nil {
		return fmt.Errorf("unable to encode passportElementError#91059fc5: field source: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PassportElementError) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passportElementError#91059fc5 to nil")
	}
	if err := b.ConsumeID(PassportElementErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode passportElementError#91059fc5: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PassportElementError) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passportElementError#91059fc5 to nil")
	}
	{
		value, err := DecodePassportElementType(b)
		if err != nil {
			return fmt.Errorf("unable to decode passportElementError#91059fc5: field type: %w", err)
		}
		p.Type = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode passportElementError#91059fc5: field message: %w", err)
		}
		p.Message = value
	}
	{
		value, err := DecodePassportElementErrorSource(b)
		if err != nil {
			return fmt.Errorf("unable to decode passportElementError#91059fc5: field source: %w", err)
		}
		p.Source = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PassportElementError) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode passportElementError#91059fc5 as nil")
	}
	b.ObjStart()
	b.PutID("passportElementError")
	b.Comma()
	b.FieldStart("type")
	if p.Type == nil {
		return fmt.Errorf("unable to encode passportElementError#91059fc5: field type is nil")
	}
	if err := p.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode passportElementError#91059fc5: field type: %w", err)
	}
	b.Comma()
	b.FieldStart("message")
	b.PutString(p.Message)
	b.Comma()
	b.FieldStart("source")
	if p.Source == nil {
		return fmt.Errorf("unable to encode passportElementError#91059fc5: field source is nil")
	}
	if err := p.Source.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode passportElementError#91059fc5: field source: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *PassportElementError) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode passportElementError#91059fc5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("passportElementError"); err != nil {
				return fmt.Errorf("unable to decode passportElementError#91059fc5: %w", err)
			}
		case "type":
			value, err := DecodeTDLibJSONPassportElementType(b)
			if err != nil {
				return fmt.Errorf("unable to decode passportElementError#91059fc5: field type: %w", err)
			}
			p.Type = value
		case "message":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode passportElementError#91059fc5: field message: %w", err)
			}
			p.Message = value
		case "source":
			value, err := DecodeTDLibJSONPassportElementErrorSource(b)
			if err != nil {
				return fmt.Errorf("unable to decode passportElementError#91059fc5: field source: %w", err)
			}
			p.Source = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetType returns value of Type field.
func (p *PassportElementError) GetType() (value PassportElementTypeClass) {
	if p == nil {
		return
	}
	return p.Type
}

// GetMessage returns value of Message field.
func (p *PassportElementError) GetMessage() (value string) {
	if p == nil {
		return
	}
	return p.Message
}

// GetSource returns value of Source field.
func (p *PassportElementError) GetSource() (value PassportElementErrorSourceClass) {
	if p == nil {
		return
	}
	return p.Source
}
