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

// ChatBoostLinkInfo represents TL type `chatBoostLinkInfo#dc123884`.
type ChatBoostLinkInfo struct {
	// True, if the link will work for non-members of the chat
	IsPublic bool
	// Identifier of the chat to which the link points; 0 if the chat isn't found
	ChatID int64
}

// ChatBoostLinkInfoTypeID is TL type id of ChatBoostLinkInfo.
const ChatBoostLinkInfoTypeID = 0xdc123884

// Ensuring interfaces in compile-time for ChatBoostLinkInfo.
var (
	_ bin.Encoder     = &ChatBoostLinkInfo{}
	_ bin.Decoder     = &ChatBoostLinkInfo{}
	_ bin.BareEncoder = &ChatBoostLinkInfo{}
	_ bin.BareDecoder = &ChatBoostLinkInfo{}
)

func (c *ChatBoostLinkInfo) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.IsPublic == false) {
		return false
	}
	if !(c.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatBoostLinkInfo) String() string {
	if c == nil {
		return "ChatBoostLinkInfo(nil)"
	}
	type Alias ChatBoostLinkInfo
	return fmt.Sprintf("ChatBoostLinkInfo%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatBoostLinkInfo) TypeID() uint32 {
	return ChatBoostLinkInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatBoostLinkInfo) TypeName() string {
	return "chatBoostLinkInfo"
}

// TypeInfo returns info about TL type.
func (c *ChatBoostLinkInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatBoostLinkInfo",
		ID:   ChatBoostLinkInfoTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsPublic",
			SchemaName: "is_public",
		},
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatBoostLinkInfo) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoostLinkInfo#dc123884 as nil")
	}
	b.PutID(ChatBoostLinkInfoTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatBoostLinkInfo) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoostLinkInfo#dc123884 as nil")
	}
	b.PutBool(c.IsPublic)
	b.PutInt53(c.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatBoostLinkInfo) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoostLinkInfo#dc123884 to nil")
	}
	if err := b.ConsumeID(ChatBoostLinkInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode chatBoostLinkInfo#dc123884: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatBoostLinkInfo) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoostLinkInfo#dc123884 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLinkInfo#dc123884: field is_public: %w", err)
		}
		c.IsPublic = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode chatBoostLinkInfo#dc123884: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatBoostLinkInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatBoostLinkInfo#dc123884 as nil")
	}
	b.ObjStart()
	b.PutID("chatBoostLinkInfo")
	b.Comma()
	b.FieldStart("is_public")
	b.PutBool(c.IsPublic)
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(c.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatBoostLinkInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatBoostLinkInfo#dc123884 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatBoostLinkInfo"); err != nil {
				return fmt.Errorf("unable to decode chatBoostLinkInfo#dc123884: %w", err)
			}
		case "is_public":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLinkInfo#dc123884: field is_public: %w", err)
			}
			c.IsPublic = value
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode chatBoostLinkInfo#dc123884: field chat_id: %w", err)
			}
			c.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsPublic returns value of IsPublic field.
func (c *ChatBoostLinkInfo) GetIsPublic() (value bool) {
	if c == nil {
		return
	}
	return c.IsPublic
}

// GetChatID returns value of ChatID field.
func (c *ChatBoostLinkInfo) GetChatID() (value int64) {
	if c == nil {
		return
	}
	return c.ChatID
}
