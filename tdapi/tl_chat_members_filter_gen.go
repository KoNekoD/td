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

// ChatMembersFilterContacts represents TL type `chatMembersFilterContacts#69c480a7`.
type ChatMembersFilterContacts struct {
}

// ChatMembersFilterContactsTypeID is TL type id of ChatMembersFilterContacts.
const ChatMembersFilterContactsTypeID = 0x69c480a7

// construct implements constructor of ChatMembersFilterClass.
func (c ChatMembersFilterContacts) construct() ChatMembersFilterClass { return &c }

// Ensuring interfaces in compile-time for ChatMembersFilterContacts.
var (
	_ bin.Encoder     = &ChatMembersFilterContacts{}
	_ bin.Decoder     = &ChatMembersFilterContacts{}
	_ bin.BareEncoder = &ChatMembersFilterContacts{}
	_ bin.BareDecoder = &ChatMembersFilterContacts{}

	_ ChatMembersFilterClass = &ChatMembersFilterContacts{}
)

func (c *ChatMembersFilterContacts) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMembersFilterContacts) String() string {
	if c == nil {
		return "ChatMembersFilterContacts(nil)"
	}
	type Alias ChatMembersFilterContacts
	return fmt.Sprintf("ChatMembersFilterContacts%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMembersFilterContacts) TypeID() uint32 {
	return ChatMembersFilterContactsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMembersFilterContacts) TypeName() string {
	return "chatMembersFilterContacts"
}

// TypeInfo returns info about TL type.
func (c *ChatMembersFilterContacts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMembersFilterContacts",
		ID:   ChatMembersFilterContactsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMembersFilterContacts) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterContacts#69c480a7 as nil")
	}
	b.PutID(ChatMembersFilterContactsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMembersFilterContacts) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterContacts#69c480a7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMembersFilterContacts) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterContacts#69c480a7 to nil")
	}
	if err := b.ConsumeID(ChatMembersFilterContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMembersFilterContacts#69c480a7: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMembersFilterContacts) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterContacts#69c480a7 to nil")
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatMembersFilterContacts) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterContacts#69c480a7 as nil")
	}
	b.ObjStart()
	b.PutID("chatMembersFilterContacts")
	b.ObjEnd()
	return nil
}

// ChatMembersFilterAdministrators represents TL type `chatMembersFilterAdministrators#b47cbc1c`.
type ChatMembersFilterAdministrators struct {
}

// ChatMembersFilterAdministratorsTypeID is TL type id of ChatMembersFilterAdministrators.
const ChatMembersFilterAdministratorsTypeID = 0xb47cbc1c

// construct implements constructor of ChatMembersFilterClass.
func (c ChatMembersFilterAdministrators) construct() ChatMembersFilterClass { return &c }

// Ensuring interfaces in compile-time for ChatMembersFilterAdministrators.
var (
	_ bin.Encoder     = &ChatMembersFilterAdministrators{}
	_ bin.Decoder     = &ChatMembersFilterAdministrators{}
	_ bin.BareEncoder = &ChatMembersFilterAdministrators{}
	_ bin.BareDecoder = &ChatMembersFilterAdministrators{}

	_ ChatMembersFilterClass = &ChatMembersFilterAdministrators{}
)

func (c *ChatMembersFilterAdministrators) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMembersFilterAdministrators) String() string {
	if c == nil {
		return "ChatMembersFilterAdministrators(nil)"
	}
	type Alias ChatMembersFilterAdministrators
	return fmt.Sprintf("ChatMembersFilterAdministrators%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMembersFilterAdministrators) TypeID() uint32 {
	return ChatMembersFilterAdministratorsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMembersFilterAdministrators) TypeName() string {
	return "chatMembersFilterAdministrators"
}

// TypeInfo returns info about TL type.
func (c *ChatMembersFilterAdministrators) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMembersFilterAdministrators",
		ID:   ChatMembersFilterAdministratorsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMembersFilterAdministrators) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterAdministrators#b47cbc1c as nil")
	}
	b.PutID(ChatMembersFilterAdministratorsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMembersFilterAdministrators) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterAdministrators#b47cbc1c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMembersFilterAdministrators) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterAdministrators#b47cbc1c to nil")
	}
	if err := b.ConsumeID(ChatMembersFilterAdministratorsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMembersFilterAdministrators#b47cbc1c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMembersFilterAdministrators) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterAdministrators#b47cbc1c to nil")
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatMembersFilterAdministrators) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterAdministrators#b47cbc1c as nil")
	}
	b.ObjStart()
	b.PutID("chatMembersFilterAdministrators")
	b.ObjEnd()
	return nil
}

// ChatMembersFilterMembers represents TL type `chatMembersFilterMembers#27f71596`.
type ChatMembersFilterMembers struct {
}

// ChatMembersFilterMembersTypeID is TL type id of ChatMembersFilterMembers.
const ChatMembersFilterMembersTypeID = 0x27f71596

// construct implements constructor of ChatMembersFilterClass.
func (c ChatMembersFilterMembers) construct() ChatMembersFilterClass { return &c }

// Ensuring interfaces in compile-time for ChatMembersFilterMembers.
var (
	_ bin.Encoder     = &ChatMembersFilterMembers{}
	_ bin.Decoder     = &ChatMembersFilterMembers{}
	_ bin.BareEncoder = &ChatMembersFilterMembers{}
	_ bin.BareDecoder = &ChatMembersFilterMembers{}

	_ ChatMembersFilterClass = &ChatMembersFilterMembers{}
)

func (c *ChatMembersFilterMembers) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMembersFilterMembers) String() string {
	if c == nil {
		return "ChatMembersFilterMembers(nil)"
	}
	type Alias ChatMembersFilterMembers
	return fmt.Sprintf("ChatMembersFilterMembers%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMembersFilterMembers) TypeID() uint32 {
	return ChatMembersFilterMembersTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMembersFilterMembers) TypeName() string {
	return "chatMembersFilterMembers"
}

// TypeInfo returns info about TL type.
func (c *ChatMembersFilterMembers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMembersFilterMembers",
		ID:   ChatMembersFilterMembersTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMembersFilterMembers) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterMembers#27f71596 as nil")
	}
	b.PutID(ChatMembersFilterMembersTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMembersFilterMembers) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterMembers#27f71596 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMembersFilterMembers) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterMembers#27f71596 to nil")
	}
	if err := b.ConsumeID(ChatMembersFilterMembersTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMembersFilterMembers#27f71596: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMembersFilterMembers) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterMembers#27f71596 to nil")
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatMembersFilterMembers) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterMembers#27f71596 as nil")
	}
	b.ObjStart()
	b.PutID("chatMembersFilterMembers")
	b.ObjEnd()
	return nil
}

// ChatMembersFilterMention represents TL type `chatMembersFilterMention#330bedf7`.
type ChatMembersFilterMention struct {
	// If non-zero, the identifier of the current message thread
	MessageThreadID int64
}

// ChatMembersFilterMentionTypeID is TL type id of ChatMembersFilterMention.
const ChatMembersFilterMentionTypeID = 0x330bedf7

// construct implements constructor of ChatMembersFilterClass.
func (c ChatMembersFilterMention) construct() ChatMembersFilterClass { return &c }

// Ensuring interfaces in compile-time for ChatMembersFilterMention.
var (
	_ bin.Encoder     = &ChatMembersFilterMention{}
	_ bin.Decoder     = &ChatMembersFilterMention{}
	_ bin.BareEncoder = &ChatMembersFilterMention{}
	_ bin.BareDecoder = &ChatMembersFilterMention{}

	_ ChatMembersFilterClass = &ChatMembersFilterMention{}
)

func (c *ChatMembersFilterMention) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.MessageThreadID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMembersFilterMention) String() string {
	if c == nil {
		return "ChatMembersFilterMention(nil)"
	}
	type Alias ChatMembersFilterMention
	return fmt.Sprintf("ChatMembersFilterMention%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMembersFilterMention) TypeID() uint32 {
	return ChatMembersFilterMentionTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMembersFilterMention) TypeName() string {
	return "chatMembersFilterMention"
}

// TypeInfo returns info about TL type.
func (c *ChatMembersFilterMention) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMembersFilterMention",
		ID:   ChatMembersFilterMentionTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMembersFilterMention) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterMention#330bedf7 as nil")
	}
	b.PutID(ChatMembersFilterMentionTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMembersFilterMention) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterMention#330bedf7 as nil")
	}
	b.PutLong(c.MessageThreadID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMembersFilterMention) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterMention#330bedf7 to nil")
	}
	if err := b.ConsumeID(ChatMembersFilterMentionTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMembersFilterMention#330bedf7: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMembersFilterMention) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterMention#330bedf7 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatMembersFilterMention#330bedf7: field message_thread_id: %w", err)
		}
		c.MessageThreadID = value
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatMembersFilterMention) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterMention#330bedf7 as nil")
	}
	b.ObjStart()
	b.PutID("chatMembersFilterMention")
	b.FieldStart("message_thread_id")
	b.PutLong(c.MessageThreadID)
	b.ObjEnd()
	return nil
}

// GetMessageThreadID returns value of MessageThreadID field.
func (c *ChatMembersFilterMention) GetMessageThreadID() (value int64) {
	return c.MessageThreadID
}

// ChatMembersFilterRestricted represents TL type `chatMembersFilterRestricted#4ae15abd`.
type ChatMembersFilterRestricted struct {
}

// ChatMembersFilterRestrictedTypeID is TL type id of ChatMembersFilterRestricted.
const ChatMembersFilterRestrictedTypeID = 0x4ae15abd

// construct implements constructor of ChatMembersFilterClass.
func (c ChatMembersFilterRestricted) construct() ChatMembersFilterClass { return &c }

// Ensuring interfaces in compile-time for ChatMembersFilterRestricted.
var (
	_ bin.Encoder     = &ChatMembersFilterRestricted{}
	_ bin.Decoder     = &ChatMembersFilterRestricted{}
	_ bin.BareEncoder = &ChatMembersFilterRestricted{}
	_ bin.BareDecoder = &ChatMembersFilterRestricted{}

	_ ChatMembersFilterClass = &ChatMembersFilterRestricted{}
)

func (c *ChatMembersFilterRestricted) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMembersFilterRestricted) String() string {
	if c == nil {
		return "ChatMembersFilterRestricted(nil)"
	}
	type Alias ChatMembersFilterRestricted
	return fmt.Sprintf("ChatMembersFilterRestricted%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMembersFilterRestricted) TypeID() uint32 {
	return ChatMembersFilterRestrictedTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMembersFilterRestricted) TypeName() string {
	return "chatMembersFilterRestricted"
}

// TypeInfo returns info about TL type.
func (c *ChatMembersFilterRestricted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMembersFilterRestricted",
		ID:   ChatMembersFilterRestrictedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMembersFilterRestricted) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterRestricted#4ae15abd as nil")
	}
	b.PutID(ChatMembersFilterRestrictedTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMembersFilterRestricted) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterRestricted#4ae15abd as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMembersFilterRestricted) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterRestricted#4ae15abd to nil")
	}
	if err := b.ConsumeID(ChatMembersFilterRestrictedTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMembersFilterRestricted#4ae15abd: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMembersFilterRestricted) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterRestricted#4ae15abd to nil")
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatMembersFilterRestricted) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterRestricted#4ae15abd as nil")
	}
	b.ObjStart()
	b.PutID("chatMembersFilterRestricted")
	b.ObjEnd()
	return nil
}

// ChatMembersFilterBanned represents TL type `chatMembersFilterBanned#90f34f48`.
type ChatMembersFilterBanned struct {
}

// ChatMembersFilterBannedTypeID is TL type id of ChatMembersFilterBanned.
const ChatMembersFilterBannedTypeID = 0x90f34f48

// construct implements constructor of ChatMembersFilterClass.
func (c ChatMembersFilterBanned) construct() ChatMembersFilterClass { return &c }

// Ensuring interfaces in compile-time for ChatMembersFilterBanned.
var (
	_ bin.Encoder     = &ChatMembersFilterBanned{}
	_ bin.Decoder     = &ChatMembersFilterBanned{}
	_ bin.BareEncoder = &ChatMembersFilterBanned{}
	_ bin.BareDecoder = &ChatMembersFilterBanned{}

	_ ChatMembersFilterClass = &ChatMembersFilterBanned{}
)

func (c *ChatMembersFilterBanned) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMembersFilterBanned) String() string {
	if c == nil {
		return "ChatMembersFilterBanned(nil)"
	}
	type Alias ChatMembersFilterBanned
	return fmt.Sprintf("ChatMembersFilterBanned%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMembersFilterBanned) TypeID() uint32 {
	return ChatMembersFilterBannedTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMembersFilterBanned) TypeName() string {
	return "chatMembersFilterBanned"
}

// TypeInfo returns info about TL type.
func (c *ChatMembersFilterBanned) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMembersFilterBanned",
		ID:   ChatMembersFilterBannedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMembersFilterBanned) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterBanned#90f34f48 as nil")
	}
	b.PutID(ChatMembersFilterBannedTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMembersFilterBanned) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterBanned#90f34f48 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMembersFilterBanned) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterBanned#90f34f48 to nil")
	}
	if err := b.ConsumeID(ChatMembersFilterBannedTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMembersFilterBanned#90f34f48: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMembersFilterBanned) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterBanned#90f34f48 to nil")
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatMembersFilterBanned) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterBanned#90f34f48 as nil")
	}
	b.ObjStart()
	b.PutID("chatMembersFilterBanned")
	b.ObjEnd()
	return nil
}

// ChatMembersFilterBots represents TL type `chatMembersFilterBots#ab355888`.
type ChatMembersFilterBots struct {
}

// ChatMembersFilterBotsTypeID is TL type id of ChatMembersFilterBots.
const ChatMembersFilterBotsTypeID = 0xab355888

// construct implements constructor of ChatMembersFilterClass.
func (c ChatMembersFilterBots) construct() ChatMembersFilterClass { return &c }

// Ensuring interfaces in compile-time for ChatMembersFilterBots.
var (
	_ bin.Encoder     = &ChatMembersFilterBots{}
	_ bin.Decoder     = &ChatMembersFilterBots{}
	_ bin.BareEncoder = &ChatMembersFilterBots{}
	_ bin.BareDecoder = &ChatMembersFilterBots{}

	_ ChatMembersFilterClass = &ChatMembersFilterBots{}
)

func (c *ChatMembersFilterBots) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMembersFilterBots) String() string {
	if c == nil {
		return "ChatMembersFilterBots(nil)"
	}
	type Alias ChatMembersFilterBots
	return fmt.Sprintf("ChatMembersFilterBots%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMembersFilterBots) TypeID() uint32 {
	return ChatMembersFilterBotsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMembersFilterBots) TypeName() string {
	return "chatMembersFilterBots"
}

// TypeInfo returns info about TL type.
func (c *ChatMembersFilterBots) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMembersFilterBots",
		ID:   ChatMembersFilterBotsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMembersFilterBots) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterBots#ab355888 as nil")
	}
	b.PutID(ChatMembersFilterBotsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMembersFilterBots) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterBots#ab355888 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMembersFilterBots) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterBots#ab355888 to nil")
	}
	if err := b.ConsumeID(ChatMembersFilterBotsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMembersFilterBots#ab355888: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMembersFilterBots) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMembersFilterBots#ab355888 to nil")
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatMembersFilterBots) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMembersFilterBots#ab355888 as nil")
	}
	b.ObjStart()
	b.PutID("chatMembersFilterBots")
	b.ObjEnd()
	return nil
}

// ChatMembersFilterClass represents ChatMembersFilter generic type.
//
// Example:
//  g, err := tdapi.DecodeChatMembersFilter(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.ChatMembersFilterContacts: // chatMembersFilterContacts#69c480a7
//  case *tdapi.ChatMembersFilterAdministrators: // chatMembersFilterAdministrators#b47cbc1c
//  case *tdapi.ChatMembersFilterMembers: // chatMembersFilterMembers#27f71596
//  case *tdapi.ChatMembersFilterMention: // chatMembersFilterMention#330bedf7
//  case *tdapi.ChatMembersFilterRestricted: // chatMembersFilterRestricted#4ae15abd
//  case *tdapi.ChatMembersFilterBanned: // chatMembersFilterBanned#90f34f48
//  case *tdapi.ChatMembersFilterBots: // chatMembersFilterBots#ab355888
//  default: panic(v)
//  }
type ChatMembersFilterClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatMembersFilterClass

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
	EncodeTDLibJSON(b *jsontd.Encoder) error
}

// DecodeChatMembersFilter implements binary de-serialization for ChatMembersFilterClass.
func DecodeChatMembersFilter(buf *bin.Buffer) (ChatMembersFilterClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatMembersFilterContactsTypeID:
		// Decoding chatMembersFilterContacts#69c480a7.
		v := ChatMembersFilterContacts{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMembersFilterClass: %w", err)
		}
		return &v, nil
	case ChatMembersFilterAdministratorsTypeID:
		// Decoding chatMembersFilterAdministrators#b47cbc1c.
		v := ChatMembersFilterAdministrators{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMembersFilterClass: %w", err)
		}
		return &v, nil
	case ChatMembersFilterMembersTypeID:
		// Decoding chatMembersFilterMembers#27f71596.
		v := ChatMembersFilterMembers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMembersFilterClass: %w", err)
		}
		return &v, nil
	case ChatMembersFilterMentionTypeID:
		// Decoding chatMembersFilterMention#330bedf7.
		v := ChatMembersFilterMention{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMembersFilterClass: %w", err)
		}
		return &v, nil
	case ChatMembersFilterRestrictedTypeID:
		// Decoding chatMembersFilterRestricted#4ae15abd.
		v := ChatMembersFilterRestricted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMembersFilterClass: %w", err)
		}
		return &v, nil
	case ChatMembersFilterBannedTypeID:
		// Decoding chatMembersFilterBanned#90f34f48.
		v := ChatMembersFilterBanned{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMembersFilterClass: %w", err)
		}
		return &v, nil
	case ChatMembersFilterBotsTypeID:
		// Decoding chatMembersFilterBots#ab355888.
		v := ChatMembersFilterBots{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMembersFilterClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatMembersFilterClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatMembersFilter boxes the ChatMembersFilterClass providing a helper.
type ChatMembersFilterBox struct {
	ChatMembersFilter ChatMembersFilterClass
}

// Decode implements bin.Decoder for ChatMembersFilterBox.
func (b *ChatMembersFilterBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatMembersFilterBox to nil")
	}
	v, err := DecodeChatMembersFilter(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatMembersFilter = v
	return nil
}

// Encode implements bin.Encode for ChatMembersFilterBox.
func (b *ChatMembersFilterBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatMembersFilter == nil {
		return fmt.Errorf("unable to encode ChatMembersFilterClass as nil")
	}
	return b.ChatMembersFilter.Encode(buf)
}
