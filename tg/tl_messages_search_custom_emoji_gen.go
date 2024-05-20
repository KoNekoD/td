// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesSearchCustomEmojiRequest represents TL type `messages.searchCustomEmoji#2c11c0d7`.
// Look for custom emojis¹ associated to a UTF8 emoji
//
// Links:
//  1. https://core.telegram.org/api/custom-emoji
//
// See https://core.telegram.org/method/messages.searchCustomEmoji for reference.
type MessagesSearchCustomEmojiRequest struct {
	// The emoji
	Emoticon string
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// MessagesSearchCustomEmojiRequestTypeID is TL type id of MessagesSearchCustomEmojiRequest.
const MessagesSearchCustomEmojiRequestTypeID = 0x2c11c0d7

// Ensuring interfaces in compile-time for MessagesSearchCustomEmojiRequest.
var (
	_ bin.Encoder     = &MessagesSearchCustomEmojiRequest{}
	_ bin.Decoder     = &MessagesSearchCustomEmojiRequest{}
	_ bin.BareEncoder = &MessagesSearchCustomEmojiRequest{}
	_ bin.BareDecoder = &MessagesSearchCustomEmojiRequest{}
)

func (s *MessagesSearchCustomEmojiRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Emoticon == "") {
		return false
	}
	if !(s.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSearchCustomEmojiRequest) String() string {
	if s == nil {
		return "MessagesSearchCustomEmojiRequest(nil)"
	}
	type Alias MessagesSearchCustomEmojiRequest
	return fmt.Sprintf("MessagesSearchCustomEmojiRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSearchCustomEmojiRequest from given interface.
func (s *MessagesSearchCustomEmojiRequest) FillFrom(from interface {
	GetEmoticon() (value string)
	GetHash() (value int64)
}) {
	s.Emoticon = from.GetEmoticon()
	s.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSearchCustomEmojiRequest) TypeID() uint32 {
	return MessagesSearchCustomEmojiRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSearchCustomEmojiRequest) TypeName() string {
	return "messages.searchCustomEmoji"
}

// TypeInfo returns info about TL type.
func (s *MessagesSearchCustomEmojiRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.searchCustomEmoji",
		ID:   MessagesSearchCustomEmojiRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Emoticon",
			SchemaName: "emoticon",
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSearchCustomEmojiRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchCustomEmoji#2c11c0d7 as nil")
	}
	b.PutID(MessagesSearchCustomEmojiRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSearchCustomEmojiRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.searchCustomEmoji#2c11c0d7 as nil")
	}
	b.PutString(s.Emoticon)
	b.PutLong(s.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSearchCustomEmojiRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchCustomEmoji#2c11c0d7 to nil")
	}
	if err := b.ConsumeID(MessagesSearchCustomEmojiRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.searchCustomEmoji#2c11c0d7: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSearchCustomEmojiRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.searchCustomEmoji#2c11c0d7 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchCustomEmoji#2c11c0d7: field emoticon: %w", err)
		}
		s.Emoticon = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.searchCustomEmoji#2c11c0d7: field hash: %w", err)
		}
		s.Hash = value
	}
	return nil
}

// GetEmoticon returns value of Emoticon field.
func (s *MessagesSearchCustomEmojiRequest) GetEmoticon() (value string) {
	if s == nil {
		return
	}
	return s.Emoticon
}

// GetHash returns value of Hash field.
func (s *MessagesSearchCustomEmojiRequest) GetHash() (value int64) {
	if s == nil {
		return
	}
	return s.Hash
}

// MessagesSearchCustomEmoji invokes method messages.searchCustomEmoji#2c11c0d7 returning error if any.
// Look for custom emojis¹ associated to a UTF8 emoji
//
// Links:
//  1. https://core.telegram.org/api/custom-emoji
//
// Possible errors:
//
//	400 EMOTICON_EMPTY: The emoji is empty.
//
// See https://core.telegram.org/method/messages.searchCustomEmoji for reference.
// Can be used by bots.
func (c *Client) MessagesSearchCustomEmoji(ctx context.Context, request *MessagesSearchCustomEmojiRequest) (EmojiListClass, error) {
	var result EmojiListBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.EmojiList, nil
}
