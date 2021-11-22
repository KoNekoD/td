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

// ChatLists represents TL type `chatLists#92c2d216`.
type ChatLists struct {
	// List of chat lists
	ChatLists []ChatListClass
}

// ChatListsTypeID is TL type id of ChatLists.
const ChatListsTypeID = 0x92c2d216

// Ensuring interfaces in compile-time for ChatLists.
var (
	_ bin.Encoder     = &ChatLists{}
	_ bin.Decoder     = &ChatLists{}
	_ bin.BareEncoder = &ChatLists{}
	_ bin.BareDecoder = &ChatLists{}
)

func (c *ChatLists) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ChatLists == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatLists) String() string {
	if c == nil {
		return "ChatLists(nil)"
	}
	type Alias ChatLists
	return fmt.Sprintf("ChatLists%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatLists) TypeID() uint32 {
	return ChatListsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatLists) TypeName() string {
	return "chatLists"
}

// TypeInfo returns info about TL type.
func (c *ChatLists) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatLists",
		ID:   ChatListsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatLists",
			SchemaName: "chat_lists",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatLists) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatLists#92c2d216 as nil")
	}
	b.PutID(ChatListsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatLists) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatLists#92c2d216 as nil")
	}
	b.PutInt(len(c.ChatLists))
	for idx, v := range c.ChatLists {
		if v == nil {
			return fmt.Errorf("unable to encode chatLists#92c2d216: field chat_lists element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatLists#92c2d216: field chat_lists element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatLists) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatLists#92c2d216 to nil")
	}
	if err := b.ConsumeID(ChatListsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatLists#92c2d216: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatLists) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatLists#92c2d216 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatLists#92c2d216: field chat_lists: %w", err)
		}

		if headerLen > 0 {
			c.ChatLists = make([]ChatListClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChatList(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatLists#92c2d216: field chat_lists: %w", err)
			}
			c.ChatLists = append(c.ChatLists, value)
		}
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatLists) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatLists#92c2d216 as nil")
	}
	b.ObjStart()
	b.PutID("chatLists")
	b.FieldStart("chat_lists")
	b.ArrStart()
	for idx, v := range c.ChatLists {
		if v == nil {
			return fmt.Errorf("unable to encode chatLists#92c2d216: field chat_lists element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatLists#92c2d216: field chat_lists element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// GetChatLists returns value of ChatLists field.
func (c *ChatLists) GetChatLists() (value []ChatListClass) {
	return c.ChatLists
}
