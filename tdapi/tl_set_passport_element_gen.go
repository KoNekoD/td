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

// SetPassportElementRequest represents TL type `setPassportElement#7b45d19c`.
type SetPassportElementRequest struct {
	// Input Telegram Passport element
	Element InputPassportElementClass
	// The 2-step verification password of the current user
	Password string
}

// SetPassportElementRequestTypeID is TL type id of SetPassportElementRequest.
const SetPassportElementRequestTypeID = 0x7b45d19c

// Ensuring interfaces in compile-time for SetPassportElementRequest.
var (
	_ bin.Encoder     = &SetPassportElementRequest{}
	_ bin.Decoder     = &SetPassportElementRequest{}
	_ bin.BareEncoder = &SetPassportElementRequest{}
	_ bin.BareDecoder = &SetPassportElementRequest{}
)

func (s *SetPassportElementRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Element == nil) {
		return false
	}
	if !(s.Password == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetPassportElementRequest) String() string {
	if s == nil {
		return "SetPassportElementRequest(nil)"
	}
	type Alias SetPassportElementRequest
	return fmt.Sprintf("SetPassportElementRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetPassportElementRequest) TypeID() uint32 {
	return SetPassportElementRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetPassportElementRequest) TypeName() string {
	return "setPassportElement"
}

// TypeInfo returns info about TL type.
func (s *SetPassportElementRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setPassportElement",
		ID:   SetPassportElementRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Element",
			SchemaName: "element",
		},
		{
			Name:       "Password",
			SchemaName: "password",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetPassportElementRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setPassportElement#7b45d19c as nil")
	}
	b.PutID(SetPassportElementRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetPassportElementRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setPassportElement#7b45d19c as nil")
	}
	if s.Element == nil {
		return fmt.Errorf("unable to encode setPassportElement#7b45d19c: field element is nil")
	}
	if err := s.Element.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setPassportElement#7b45d19c: field element: %w", err)
	}
	b.PutString(s.Password)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetPassportElementRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setPassportElement#7b45d19c to nil")
	}
	if err := b.ConsumeID(SetPassportElementRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setPassportElement#7b45d19c: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetPassportElementRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setPassportElement#7b45d19c to nil")
	}
	{
		value, err := DecodeInputPassportElement(b)
		if err != nil {
			return fmt.Errorf("unable to decode setPassportElement#7b45d19c: field element: %w", err)
		}
		s.Element = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setPassportElement#7b45d19c: field password: %w", err)
		}
		s.Password = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetPassportElementRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setPassportElement#7b45d19c as nil")
	}
	b.ObjStart()
	b.PutID("setPassportElement")
	b.Comma()
	b.FieldStart("element")
	if s.Element == nil {
		return fmt.Errorf("unable to encode setPassportElement#7b45d19c: field element is nil")
	}
	if err := s.Element.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setPassportElement#7b45d19c: field element: %w", err)
	}
	b.Comma()
	b.FieldStart("password")
	b.PutString(s.Password)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetPassportElementRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setPassportElement#7b45d19c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setPassportElement"); err != nil {
				return fmt.Errorf("unable to decode setPassportElement#7b45d19c: %w", err)
			}
		case "element":
			value, err := DecodeTDLibJSONInputPassportElement(b)
			if err != nil {
				return fmt.Errorf("unable to decode setPassportElement#7b45d19c: field element: %w", err)
			}
			s.Element = value
		case "password":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setPassportElement#7b45d19c: field password: %w", err)
			}
			s.Password = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetElement returns value of Element field.
func (s *SetPassportElementRequest) GetElement() (value InputPassportElementClass) {
	if s == nil {
		return
	}
	return s.Element
}

// GetPassword returns value of Password field.
func (s *SetPassportElementRequest) GetPassword() (value string) {
	if s == nil {
		return
	}
	return s.Password
}

// SetPassportElement invokes method setPassportElement#7b45d19c returning error if any.
func (c *Client) SetPassportElement(ctx context.Context, request *SetPassportElementRequest) (PassportElementClass, error) {
	var result PassportElementBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.PassportElement, nil
}
