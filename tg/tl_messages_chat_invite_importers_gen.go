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

// MessagesChatInviteImporters represents TL type `messages.chatInviteImporters#81b6b00a`.
// Info about the users that joined the chat using a specific chat invite
//
// See https://core.telegram.org/constructor/messages.chatInviteImporters for reference.
type MessagesChatInviteImporters struct {
	// Number of users that joined
	Count int
	// The users that joined
	Importers []ChatInviteImporter
	// The users that joined
	Users []UserClass
}

// MessagesChatInviteImportersTypeID is TL type id of MessagesChatInviteImporters.
const MessagesChatInviteImportersTypeID = 0x81b6b00a

// Ensuring interfaces in compile-time for MessagesChatInviteImporters.
var (
	_ bin.Encoder     = &MessagesChatInviteImporters{}
	_ bin.Decoder     = &MessagesChatInviteImporters{}
	_ bin.BareEncoder = &MessagesChatInviteImporters{}
	_ bin.BareDecoder = &MessagesChatInviteImporters{}
)

func (c *MessagesChatInviteImporters) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Count == 0) {
		return false
	}
	if !(c.Importers == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *MessagesChatInviteImporters) String() string {
	if c == nil {
		return "MessagesChatInviteImporters(nil)"
	}
	type Alias MessagesChatInviteImporters
	return fmt.Sprintf("MessagesChatInviteImporters%+v", Alias(*c))
}

// FillFrom fills MessagesChatInviteImporters from given interface.
func (c *MessagesChatInviteImporters) FillFrom(from interface {
	GetCount() (value int)
	GetImporters() (value []ChatInviteImporter)
	GetUsers() (value []UserClass)
}) {
	c.Count = from.GetCount()
	c.Importers = from.GetImporters()
	c.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesChatInviteImporters) TypeID() uint32 {
	return MessagesChatInviteImportersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesChatInviteImporters) TypeName() string {
	return "messages.chatInviteImporters"
}

// TypeInfo returns info about TL type.
func (c *MessagesChatInviteImporters) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.chatInviteImporters",
		ID:   MessagesChatInviteImportersTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Importers",
			SchemaName: "importers",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *MessagesChatInviteImporters) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.chatInviteImporters#81b6b00a as nil")
	}
	b.PutID(MessagesChatInviteImportersTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *MessagesChatInviteImporters) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode messages.chatInviteImporters#81b6b00a as nil")
	}
	b.PutInt(c.Count)
	b.PutVectorHeader(len(c.Importers))
	for idx, v := range c.Importers {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.chatInviteImporters#81b6b00a: field importers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.chatInviteImporters#81b6b00a: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.chatInviteImporters#81b6b00a: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *MessagesChatInviteImporters) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.chatInviteImporters#81b6b00a to nil")
	}
	if err := b.ConsumeID(MessagesChatInviteImportersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.chatInviteImporters#81b6b00a: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *MessagesChatInviteImporters) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode messages.chatInviteImporters#81b6b00a to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatInviteImporters#81b6b00a: field count: %w", err)
		}
		c.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatInviteImporters#81b6b00a: field importers: %w", err)
		}

		if headerLen > 0 {
			c.Importers = make([]ChatInviteImporter, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatInviteImporter
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.chatInviteImporters#81b6b00a: field importers: %w", err)
			}
			c.Importers = append(c.Importers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.chatInviteImporters#81b6b00a: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.chatInviteImporters#81b6b00a: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (c *MessagesChatInviteImporters) GetCount() (value int) {
	if c == nil {
		return
	}
	return c.Count
}

// GetImporters returns value of Importers field.
func (c *MessagesChatInviteImporters) GetImporters() (value []ChatInviteImporter) {
	if c == nil {
		return
	}
	return c.Importers
}

// GetUsers returns value of Users field.
func (c *MessagesChatInviteImporters) GetUsers() (value []UserClass) {
	if c == nil {
		return
	}
	return c.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (c *MessagesChatInviteImporters) MapUsers() (value UserClassArray) {
	return UserClassArray(c.Users)
}
