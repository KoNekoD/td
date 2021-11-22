// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesSponsoredMessages represents TL type `messages.sponsoredMessages#65a4c7d5`.
//
// See https://core.telegram.org/constructor/messages.sponsoredMessages for reference.
type MessagesSponsoredMessages struct {
	// Messages field of MessagesSponsoredMessages.
	Messages []SponsoredMessage
	// Chats field of MessagesSponsoredMessages.
	Chats []ChatClass
	// Users field of MessagesSponsoredMessages.
	Users []UserClass
}

// MessagesSponsoredMessagesTypeID is TL type id of MessagesSponsoredMessages.
const MessagesSponsoredMessagesTypeID = 0x65a4c7d5

// Ensuring interfaces in compile-time for MessagesSponsoredMessages.
var (
	_ bin.Encoder     = &MessagesSponsoredMessages{}
	_ bin.Decoder     = &MessagesSponsoredMessages{}
	_ bin.BareEncoder = &MessagesSponsoredMessages{}
	_ bin.BareDecoder = &MessagesSponsoredMessages{}
)

func (s *MessagesSponsoredMessages) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Messages == nil) {
		return false
	}
	if !(s.Chats == nil) {
		return false
	}
	if !(s.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSponsoredMessages) String() string {
	if s == nil {
		return "MessagesSponsoredMessages(nil)"
	}
	type Alias MessagesSponsoredMessages
	return fmt.Sprintf("MessagesSponsoredMessages%+v", Alias(*s))
}

// FillFrom fills MessagesSponsoredMessages from given interface.
func (s *MessagesSponsoredMessages) FillFrom(from interface {
	GetMessages() (value []SponsoredMessage)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	s.Messages = from.GetMessages()
	s.Chats = from.GetChats()
	s.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSponsoredMessages) TypeID() uint32 {
	return MessagesSponsoredMessagesTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSponsoredMessages) TypeName() string {
	return "messages.sponsoredMessages"
}

// TypeInfo returns info about TL type.
func (s *MessagesSponsoredMessages) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.sponsoredMessages",
		ID:   MessagesSponsoredMessagesTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSponsoredMessages) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sponsoredMessages#65a4c7d5 as nil")
	}
	b.PutID(MessagesSponsoredMessagesTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSponsoredMessages) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.sponsoredMessages#65a4c7d5 as nil")
	}
	b.PutVectorHeader(len(s.Messages))
	for idx, v := range s.Messages {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sponsoredMessages#65a4c7d5: field messages element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Chats))
	for idx, v := range s.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode messages.sponsoredMessages#65a4c7d5: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sponsoredMessages#65a4c7d5: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(s.Users))
	for idx, v := range s.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.sponsoredMessages#65a4c7d5: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.sponsoredMessages#65a4c7d5: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSponsoredMessages) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sponsoredMessages#65a4c7d5 to nil")
	}
	if err := b.ConsumeID(MessagesSponsoredMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.sponsoredMessages#65a4c7d5: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSponsoredMessages) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.sponsoredMessages#65a4c7d5 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sponsoredMessages#65a4c7d5: field messages: %w", err)
		}

		if headerLen > 0 {
			s.Messages = make([]SponsoredMessage, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value SponsoredMessage
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.sponsoredMessages#65a4c7d5: field messages: %w", err)
			}
			s.Messages = append(s.Messages, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sponsoredMessages#65a4c7d5: field chats: %w", err)
		}

		if headerLen > 0 {
			s.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.sponsoredMessages#65a4c7d5: field chats: %w", err)
			}
			s.Chats = append(s.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.sponsoredMessages#65a4c7d5: field users: %w", err)
		}

		if headerLen > 0 {
			s.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.sponsoredMessages#65a4c7d5: field users: %w", err)
			}
			s.Users = append(s.Users, value)
		}
	}
	return nil
}

// GetMessages returns value of Messages field.
func (s *MessagesSponsoredMessages) GetMessages() (value []SponsoredMessage) {
	return s.Messages
}

// GetChats returns value of Chats field.
func (s *MessagesSponsoredMessages) GetChats() (value []ChatClass) {
	return s.Chats
}

// GetUsers returns value of Users field.
func (s *MessagesSponsoredMessages) GetUsers() (value []UserClass) {
	return s.Users
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (s *MessagesSponsoredMessages) MapChats() (value ChatClassArray) {
	return ChatClassArray(s.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (s *MessagesSponsoredMessages) MapUsers() (value UserClassArray) {
	return UserClassArray(s.Users)
}
