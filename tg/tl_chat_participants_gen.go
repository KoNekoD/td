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

// ChatParticipantsForbidden represents TL type `chatParticipantsForbidden#8763d3e1`.
// Info on members is unavailable
//
// See https://core.telegram.org/constructor/chatParticipantsForbidden for reference.
type ChatParticipantsForbidden struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Group ID
	ChatID int64
	// Info about the group membership of the current user
	//
	// Use SetSelfParticipant and GetSelfParticipant helpers.
	SelfParticipant ChatParticipantClass
}

// ChatParticipantsForbiddenTypeID is TL type id of ChatParticipantsForbidden.
const ChatParticipantsForbiddenTypeID = 0x8763d3e1

// construct implements constructor of ChatParticipantsClass.
func (c ChatParticipantsForbidden) construct() ChatParticipantsClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipantsForbidden.
var (
	_ bin.Encoder     = &ChatParticipantsForbidden{}
	_ bin.Decoder     = &ChatParticipantsForbidden{}
	_ bin.BareEncoder = &ChatParticipantsForbidden{}
	_ bin.BareDecoder = &ChatParticipantsForbidden{}

	_ ChatParticipantsClass = &ChatParticipantsForbidden{}
)

func (c *ChatParticipantsForbidden) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.ChatID == 0) {
		return false
	}
	if !(c.SelfParticipant == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipantsForbidden) String() string {
	if c == nil {
		return "ChatParticipantsForbidden(nil)"
	}
	type Alias ChatParticipantsForbidden
	return fmt.Sprintf("ChatParticipantsForbidden%+v", Alias(*c))
}

// FillFrom fills ChatParticipantsForbidden from given interface.
func (c *ChatParticipantsForbidden) FillFrom(from interface {
	GetChatID() (value int64)
	GetSelfParticipant() (value ChatParticipantClass, ok bool)
}) {
	c.ChatID = from.GetChatID()
	if val, ok := from.GetSelfParticipant(); ok {
		c.SelfParticipant = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatParticipantsForbidden) TypeID() uint32 {
	return ChatParticipantsForbiddenTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatParticipantsForbidden) TypeName() string {
	return "chatParticipantsForbidden"
}

// TypeInfo returns info about TL type.
func (c *ChatParticipantsForbidden) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatParticipantsForbidden",
		ID:   ChatParticipantsForbiddenTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "SelfParticipant",
			SchemaName: "self_participant",
			Null:       !c.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *ChatParticipantsForbidden) SetFlags() {
	if !(c.SelfParticipant == nil) {
		c.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (c *ChatParticipantsForbidden) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantsForbidden#8763d3e1 as nil")
	}
	b.PutID(ChatParticipantsForbiddenTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatParticipantsForbidden) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantsForbidden#8763d3e1 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatParticipantsForbidden#8763d3e1: field flags: %w", err)
	}
	b.PutLong(c.ChatID)
	if c.Flags.Has(0) {
		if c.SelfParticipant == nil {
			return fmt.Errorf("unable to encode chatParticipantsForbidden#8763d3e1: field self_participant is nil")
		}
		if err := c.SelfParticipant.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatParticipantsForbidden#8763d3e1: field self_participant: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatParticipantsForbidden) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantsForbidden#8763d3e1 to nil")
	}
	if err := b.ConsumeID(ChatParticipantsForbiddenTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipantsForbidden#8763d3e1: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatParticipantsForbidden) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantsForbidden#8763d3e1 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatParticipantsForbidden#8763d3e1: field flags: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantsForbidden#8763d3e1: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	if c.Flags.Has(0) {
		value, err := DecodeChatParticipant(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantsForbidden#8763d3e1: field self_participant: %w", err)
		}
		c.SelfParticipant = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (c *ChatParticipantsForbidden) GetChatID() (value int64) {
	if c == nil {
		return
	}
	return c.ChatID
}

// SetSelfParticipant sets value of SelfParticipant conditional field.
func (c *ChatParticipantsForbidden) SetSelfParticipant(value ChatParticipantClass) {
	c.Flags.Set(0)
	c.SelfParticipant = value
}

// GetSelfParticipant returns value of SelfParticipant conditional field and
// boolean which is true if field was set.
func (c *ChatParticipantsForbidden) GetSelfParticipant() (value ChatParticipantClass, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(0) {
		return value, false
	}
	return c.SelfParticipant, true
}

// ChatParticipants represents TL type `chatParticipants#3cbc93f8`.
// Group members.
//
// See https://core.telegram.org/constructor/chatParticipants for reference.
type ChatParticipants struct {
	// Group identifier
	ChatID int64
	// List of group members
	Participants []ChatParticipantClass
	// Group version number
	Version int
}

// ChatParticipantsTypeID is TL type id of ChatParticipants.
const ChatParticipantsTypeID = 0x3cbc93f8

// construct implements constructor of ChatParticipantsClass.
func (c ChatParticipants) construct() ChatParticipantsClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipants.
var (
	_ bin.Encoder     = &ChatParticipants{}
	_ bin.Decoder     = &ChatParticipants{}
	_ bin.BareEncoder = &ChatParticipants{}
	_ bin.BareDecoder = &ChatParticipants{}

	_ ChatParticipantsClass = &ChatParticipants{}
)

func (c *ChatParticipants) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ChatID == 0) {
		return false
	}
	if !(c.Participants == nil) {
		return false
	}
	if !(c.Version == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipants) String() string {
	if c == nil {
		return "ChatParticipants(nil)"
	}
	type Alias ChatParticipants
	return fmt.Sprintf("ChatParticipants%+v", Alias(*c))
}

// FillFrom fills ChatParticipants from given interface.
func (c *ChatParticipants) FillFrom(from interface {
	GetChatID() (value int64)
	GetParticipants() (value []ChatParticipantClass)
	GetVersion() (value int)
}) {
	c.ChatID = from.GetChatID()
	c.Participants = from.GetParticipants()
	c.Version = from.GetVersion()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatParticipants) TypeID() uint32 {
	return ChatParticipantsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatParticipants) TypeName() string {
	return "chatParticipants"
}

// TypeInfo returns info about TL type.
func (c *ChatParticipants) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatParticipants",
		ID:   ChatParticipantsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Participants",
			SchemaName: "participants",
		},
		{
			Name:       "Version",
			SchemaName: "version",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatParticipants) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipants#3cbc93f8 as nil")
	}
	b.PutID(ChatParticipantsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatParticipants) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipants#3cbc93f8 as nil")
	}
	b.PutLong(c.ChatID)
	b.PutVectorHeader(len(c.Participants))
	for idx, v := range c.Participants {
		if v == nil {
			return fmt.Errorf("unable to encode chatParticipants#3cbc93f8: field participants element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatParticipants#3cbc93f8: field participants element with index %d: %w", idx, err)
		}
	}
	b.PutInt(c.Version)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatParticipants) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipants#3cbc93f8 to nil")
	}
	if err := b.ConsumeID(ChatParticipantsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipants#3cbc93f8: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatParticipants) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipants#3cbc93f8 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipants#3cbc93f8: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipants#3cbc93f8: field participants: %w", err)
		}

		if headerLen > 0 {
			c.Participants = make([]ChatParticipantClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChatParticipant(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatParticipants#3cbc93f8: field participants: %w", err)
			}
			c.Participants = append(c.Participants, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipants#3cbc93f8: field version: %w", err)
		}
		c.Version = value
	}
	return nil
}

// GetChatID returns value of ChatID field.
func (c *ChatParticipants) GetChatID() (value int64) {
	if c == nil {
		return
	}
	return c.ChatID
}

// GetParticipants returns value of Participants field.
func (c *ChatParticipants) GetParticipants() (value []ChatParticipantClass) {
	if c == nil {
		return
	}
	return c.Participants
}

// GetVersion returns value of Version field.
func (c *ChatParticipants) GetVersion() (value int) {
	if c == nil {
		return
	}
	return c.Version
}

// MapParticipants returns field Participants wrapped in ChatParticipantClassArray helper.
func (c *ChatParticipants) MapParticipants() (value ChatParticipantClassArray) {
	return ChatParticipantClassArray(c.Participants)
}

// ChatParticipantsClassName is schema name of ChatParticipantsClass.
const ChatParticipantsClassName = "ChatParticipants"

// ChatParticipantsClass represents ChatParticipants generic type.
//
// See https://core.telegram.org/type/ChatParticipants for reference.
//
// Example:
//
//	g, err := tg.DecodeChatParticipants(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ChatParticipantsForbidden: // chatParticipantsForbidden#8763d3e1
//	case *tg.ChatParticipants: // chatParticipants#3cbc93f8
//	default: panic(v)
//	}
type ChatParticipantsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatParticipantsClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// Group ID
	GetChatID() (value int64)

	// AsNotForbidden tries to map ChatParticipantsClass to ChatParticipants.
	AsNotForbidden() (*ChatParticipants, bool)
}

// AsNotForbidden tries to map ChatParticipantsForbidden to ChatParticipants.
func (c *ChatParticipantsForbidden) AsNotForbidden() (*ChatParticipants, bool) {
	return nil, false
}

// AsNotForbidden tries to map ChatParticipants to ChatParticipants.
func (c *ChatParticipants) AsNotForbidden() (*ChatParticipants, bool) {
	return c, true
}

// DecodeChatParticipants implements binary de-serialization for ChatParticipantsClass.
func DecodeChatParticipants(buf *bin.Buffer) (ChatParticipantsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatParticipantsForbiddenTypeID:
		// Decoding chatParticipantsForbidden#8763d3e1.
		v := ChatParticipantsForbidden{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantsClass: %w", err)
		}
		return &v, nil
	case ChatParticipantsTypeID:
		// Decoding chatParticipants#3cbc93f8.
		v := ChatParticipants{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatParticipantsClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatParticipants boxes the ChatParticipantsClass providing a helper.
type ChatParticipantsBox struct {
	ChatParticipants ChatParticipantsClass
}

// Decode implements bin.Decoder for ChatParticipantsBox.
func (b *ChatParticipantsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatParticipantsBox to nil")
	}
	v, err := DecodeChatParticipants(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatParticipants = v
	return nil
}

// Encode implements bin.Encode for ChatParticipantsBox.
func (b *ChatParticipantsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatParticipants == nil {
		return fmt.Errorf("unable to encode ChatParticipantsClass as nil")
	}
	return b.ChatParticipants.Encode(buf)
}
