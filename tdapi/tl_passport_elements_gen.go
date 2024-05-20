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

// PassportElements represents TL type `passportElements#bd8eaceb`.
type PassportElements struct {
	// Telegram Passport elements
	Elements []PassportElementClass
}

// PassportElementsTypeID is TL type id of PassportElements.
const PassportElementsTypeID = 0xbd8eaceb

// Ensuring interfaces in compile-time for PassportElements.
var (
	_ bin.Encoder     = &PassportElements{}
	_ bin.Decoder     = &PassportElements{}
	_ bin.BareEncoder = &PassportElements{}
	_ bin.BareDecoder = &PassportElements{}
)

func (p *PassportElements) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Elements == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PassportElements) String() string {
	if p == nil {
		return "PassportElements(nil)"
	}
	type Alias PassportElements
	return fmt.Sprintf("PassportElements%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PassportElements) TypeID() uint32 {
	return PassportElementsTypeID
}

// TypeName returns name of type in TL schema.
func (*PassportElements) TypeName() string {
	return "passportElements"
}

// TypeInfo returns info about TL type.
func (p *PassportElements) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "passportElements",
		ID:   PassportElementsTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Elements",
			SchemaName: "elements",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PassportElements) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passportElements#bd8eaceb as nil")
	}
	b.PutID(PassportElementsTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PassportElements) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode passportElements#bd8eaceb as nil")
	}
	b.PutInt(len(p.Elements))
	for idx, v := range p.Elements {
		if v == nil {
			return fmt.Errorf("unable to encode passportElements#bd8eaceb: field elements element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare passportElements#bd8eaceb: field elements element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PassportElements) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passportElements#bd8eaceb to nil")
	}
	if err := b.ConsumeID(PassportElementsTypeID); err != nil {
		return fmt.Errorf("unable to decode passportElements#bd8eaceb: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PassportElements) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode passportElements#bd8eaceb to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode passportElements#bd8eaceb: field elements: %w", err)
		}

		if headerLen > 0 {
			p.Elements = make([]PassportElementClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePassportElement(b)
			if err != nil {
				return fmt.Errorf("unable to decode passportElements#bd8eaceb: field elements: %w", err)
			}
			p.Elements = append(p.Elements, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *PassportElements) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode passportElements#bd8eaceb as nil")
	}
	b.ObjStart()
	b.PutID("passportElements")
	b.Comma()
	b.FieldStart("elements")
	b.ArrStart()
	for idx, v := range p.Elements {
		if v == nil {
			return fmt.Errorf("unable to encode passportElements#bd8eaceb: field elements element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode passportElements#bd8eaceb: field elements element with index %d: %w", idx, err)
		}
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
func (p *PassportElements) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode passportElements#bd8eaceb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("passportElements"); err != nil {
				return fmt.Errorf("unable to decode passportElements#bd8eaceb: %w", err)
			}
		case "elements":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONPassportElement(b)
				if err != nil {
					return fmt.Errorf("unable to decode passportElements#bd8eaceb: field elements: %w", err)
				}
				p.Elements = append(p.Elements, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode passportElements#bd8eaceb: field elements: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetElements returns value of Elements field.
func (p *PassportElements) GetElements() (value []PassportElementClass) {
	if p == nil {
		return
	}
	return p.Elements
}
