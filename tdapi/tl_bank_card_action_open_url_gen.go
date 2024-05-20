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

// BankCardActionOpenURL represents TL type `bankCardActionOpenUrl#f44a5885`.
type BankCardActionOpenURL struct {
	// Action text
	Text string
	// The URL to be opened
	URL string
}

// BankCardActionOpenURLTypeID is TL type id of BankCardActionOpenURL.
const BankCardActionOpenURLTypeID = 0xf44a5885

// Ensuring interfaces in compile-time for BankCardActionOpenURL.
var (
	_ bin.Encoder     = &BankCardActionOpenURL{}
	_ bin.Decoder     = &BankCardActionOpenURL{}
	_ bin.BareEncoder = &BankCardActionOpenURL{}
	_ bin.BareDecoder = &BankCardActionOpenURL{}
)

func (b *BankCardActionOpenURL) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Text == "") {
		return false
	}
	if !(b.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BankCardActionOpenURL) String() string {
	if b == nil {
		return "BankCardActionOpenURL(nil)"
	}
	type Alias BankCardActionOpenURL
	return fmt.Sprintf("BankCardActionOpenURL%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BankCardActionOpenURL) TypeID() uint32 {
	return BankCardActionOpenURLTypeID
}

// TypeName returns name of type in TL schema.
func (*BankCardActionOpenURL) TypeName() string {
	return "bankCardActionOpenUrl"
}

// TypeInfo returns info about TL type.
func (b *BankCardActionOpenURL) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bankCardActionOpenUrl",
		ID:   BankCardActionOpenURLTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BankCardActionOpenURL) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bankCardActionOpenUrl#f44a5885 as nil")
	}
	buf.PutID(BankCardActionOpenURLTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BankCardActionOpenURL) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bankCardActionOpenUrl#f44a5885 as nil")
	}
	buf.PutString(b.Text)
	buf.PutString(b.URL)
	return nil
}

// Decode implements bin.Decoder.
func (b *BankCardActionOpenURL) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bankCardActionOpenUrl#f44a5885 to nil")
	}
	if err := buf.ConsumeID(BankCardActionOpenURLTypeID); err != nil {
		return fmt.Errorf("unable to decode bankCardActionOpenUrl#f44a5885: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BankCardActionOpenURL) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bankCardActionOpenUrl#f44a5885 to nil")
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode bankCardActionOpenUrl#f44a5885: field text: %w", err)
		}
		b.Text = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode bankCardActionOpenUrl#f44a5885: field url: %w", err)
		}
		b.URL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BankCardActionOpenURL) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode bankCardActionOpenUrl#f44a5885 as nil")
	}
	buf.ObjStart()
	buf.PutID("bankCardActionOpenUrl")
	buf.Comma()
	buf.FieldStart("text")
	buf.PutString(b.Text)
	buf.Comma()
	buf.FieldStart("url")
	buf.PutString(b.URL)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BankCardActionOpenURL) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode bankCardActionOpenUrl#f44a5885 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("bankCardActionOpenUrl"); err != nil {
				return fmt.Errorf("unable to decode bankCardActionOpenUrl#f44a5885: %w", err)
			}
		case "text":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode bankCardActionOpenUrl#f44a5885: field text: %w", err)
			}
			b.Text = value
		case "url":
			value, err := buf.String()
			if err != nil {
				return fmt.Errorf("unable to decode bankCardActionOpenUrl#f44a5885: field url: %w", err)
			}
			b.URL = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (b *BankCardActionOpenURL) GetText() (value string) {
	if b == nil {
		return
	}
	return b.Text
}

// GetURL returns value of URL field.
func (b *BankCardActionOpenURL) GetURL() (value string) {
	if b == nil {
		return
	}
	return b.URL
}
