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

// BusinessChatLinkInfo represents TL type `businessChatLinkInfo#cc7334af`.
type BusinessChatLinkInfo struct {
	// Identifier of the private chat that created the link
	ChatID int64
	// Message draft text that must be added to the input field
	Text FormattedText
}

// BusinessChatLinkInfoTypeID is TL type id of BusinessChatLinkInfo.
const BusinessChatLinkInfoTypeID = 0xcc7334af

// Ensuring interfaces in compile-time for BusinessChatLinkInfo.
var (
	_ bin.Encoder     = &BusinessChatLinkInfo{}
	_ bin.Decoder     = &BusinessChatLinkInfo{}
	_ bin.BareEncoder = &BusinessChatLinkInfo{}
	_ bin.BareDecoder = &BusinessChatLinkInfo{}
)

func (b *BusinessChatLinkInfo) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ChatID == 0) {
		return false
	}
	if !(b.Text.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessChatLinkInfo) String() string {
	if b == nil {
		return "BusinessChatLinkInfo(nil)"
	}
	type Alias BusinessChatLinkInfo
	return fmt.Sprintf("BusinessChatLinkInfo%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessChatLinkInfo) TypeID() uint32 {
	return BusinessChatLinkInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessChatLinkInfo) TypeName() string {
	return "businessChatLinkInfo"
}

// TypeInfo returns info about TL type.
func (b *BusinessChatLinkInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessChatLinkInfo",
		ID:   BusinessChatLinkInfoTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BusinessChatLinkInfo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessChatLinkInfo#cc7334af as nil")
	}
	buf.PutID(BusinessChatLinkInfoTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessChatLinkInfo) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessChatLinkInfo#cc7334af as nil")
	}
	buf.PutInt53(b.ChatID)
	if err := b.Text.Encode(buf); err != nil {
		return fmt.Errorf("unable to encode businessChatLinkInfo#cc7334af: field text: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessChatLinkInfo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessChatLinkInfo#cc7334af to nil")
	}
	if err := buf.ConsumeID(BusinessChatLinkInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode businessChatLinkInfo#cc7334af: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessChatLinkInfo) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessChatLinkInfo#cc7334af to nil")
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode businessChatLinkInfo#cc7334af: field chat_id: %w", err)
		}
		b.ChatID = value
	}
	{
		if err := b.Text.Decode(buf); err != nil {
			return fmt.Errorf("unable to decode businessChatLinkInfo#cc7334af: field text: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BusinessChatLinkInfo) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode businessChatLinkInfo#cc7334af as nil")
	}
	buf.ObjStart()
	buf.PutID("businessChatLinkInfo")
	buf.Comma()
	buf.FieldStart("chat_id")
	buf.PutInt53(b.ChatID)
	buf.Comma()
	buf.FieldStart("text")
	if err := b.Text.EncodeTDLibJSON(buf); err != nil {
		return fmt.Errorf("unable to encode businessChatLinkInfo#cc7334af: field text: %w", err)
	}
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BusinessChatLinkInfo) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode businessChatLinkInfo#cc7334af to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("businessChatLinkInfo"); err != nil {
				return fmt.Errorf("unable to decode businessChatLinkInfo#cc7334af: %w", err)
			}
		case "chat_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode businessChatLinkInfo#cc7334af: field chat_id: %w", err)
			}
			b.ChatID = value
		case "text":
			if err := b.Text.DecodeTDLibJSON(buf); err != nil {
				return fmt.Errorf("unable to decode businessChatLinkInfo#cc7334af: field text: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (b *BusinessChatLinkInfo) GetChatID() (value int64) {
	if b == nil {
		return
	}
	return b.ChatID
}

// GetText returns value of Text field.
func (b *BusinessChatLinkInfo) GetText() (value FormattedText) {
	if b == nil {
		return
	}
	return b.Text
}
