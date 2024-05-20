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

// ReportChatSponsoredMessageOption represents TL type `reportChatSponsoredMessageOption#aabd054a`.
type ReportChatSponsoredMessageOption struct {
	// Unique identifier of the option
	ID []byte
	// Text of the option
	Text string
}

// ReportChatSponsoredMessageOptionTypeID is TL type id of ReportChatSponsoredMessageOption.
const ReportChatSponsoredMessageOptionTypeID = 0xaabd054a

// Ensuring interfaces in compile-time for ReportChatSponsoredMessageOption.
var (
	_ bin.Encoder     = &ReportChatSponsoredMessageOption{}
	_ bin.Decoder     = &ReportChatSponsoredMessageOption{}
	_ bin.BareEncoder = &ReportChatSponsoredMessageOption{}
	_ bin.BareDecoder = &ReportChatSponsoredMessageOption{}
)

func (r *ReportChatSponsoredMessageOption) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.ID == nil) {
		return false
	}
	if !(r.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReportChatSponsoredMessageOption) String() string {
	if r == nil {
		return "ReportChatSponsoredMessageOption(nil)"
	}
	type Alias ReportChatSponsoredMessageOption
	return fmt.Sprintf("ReportChatSponsoredMessageOption%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReportChatSponsoredMessageOption) TypeID() uint32 {
	return ReportChatSponsoredMessageOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*ReportChatSponsoredMessageOption) TypeName() string {
	return "reportChatSponsoredMessageOption"
}

// TypeInfo returns info about TL type.
func (r *ReportChatSponsoredMessageOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reportChatSponsoredMessageOption",
		ID:   ReportChatSponsoredMessageOptionTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReportChatSponsoredMessageOption) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageOption#aabd054a as nil")
	}
	b.PutID(ReportChatSponsoredMessageOptionTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReportChatSponsoredMessageOption) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageOption#aabd054a as nil")
	}
	b.PutBytes(r.ID)
	b.PutString(r.Text)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReportChatSponsoredMessageOption) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageOption#aabd054a to nil")
	}
	if err := b.ConsumeID(ReportChatSponsoredMessageOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode reportChatSponsoredMessageOption#aabd054a: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReportChatSponsoredMessageOption) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageOption#aabd054a to nil")
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatSponsoredMessageOption#aabd054a: field id: %w", err)
		}
		r.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode reportChatSponsoredMessageOption#aabd054a: field text: %w", err)
		}
		r.Text = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReportChatSponsoredMessageOption) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reportChatSponsoredMessageOption#aabd054a as nil")
	}
	b.ObjStart()
	b.PutID("reportChatSponsoredMessageOption")
	b.Comma()
	b.FieldStart("id")
	b.PutBytes(r.ID)
	b.Comma()
	b.FieldStart("text")
	b.PutString(r.Text)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReportChatSponsoredMessageOption) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reportChatSponsoredMessageOption#aabd054a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reportChatSponsoredMessageOption"); err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessageOption#aabd054a: %w", err)
			}
		case "id":
			value, err := b.Bytes()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessageOption#aabd054a: field id: %w", err)
			}
			r.ID = value
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reportChatSponsoredMessageOption#aabd054a: field text: %w", err)
			}
			r.Text = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (r *ReportChatSponsoredMessageOption) GetID() (value []byte) {
	if r == nil {
		return
	}
	return r.ID
}

// GetText returns value of Text field.
func (r *ReportChatSponsoredMessageOption) GetText() (value string) {
	if r == nil {
		return
	}
	return r.Text
}
