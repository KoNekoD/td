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

// ChatInviteAlready represents TL type `chatInviteAlready#5a686d7c`.
// The user has already joined this chat
//
// See https://core.telegram.org/constructor/chatInviteAlready for reference.
type ChatInviteAlready struct {
	// The chat connected to the invite
	Chat ChatClass
}

// ChatInviteAlreadyTypeID is TL type id of ChatInviteAlready.
const ChatInviteAlreadyTypeID = 0x5a686d7c

// construct implements constructor of ChatInviteClass.
func (c ChatInviteAlready) construct() ChatInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatInviteAlready.
var (
	_ bin.Encoder     = &ChatInviteAlready{}
	_ bin.Decoder     = &ChatInviteAlready{}
	_ bin.BareEncoder = &ChatInviteAlready{}
	_ bin.BareDecoder = &ChatInviteAlready{}

	_ ChatInviteClass = &ChatInviteAlready{}
)

func (c *ChatInviteAlready) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Chat == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInviteAlready) String() string {
	if c == nil {
		return "ChatInviteAlready(nil)"
	}
	type Alias ChatInviteAlready
	return fmt.Sprintf("ChatInviteAlready%+v", Alias(*c))
}

// FillFrom fills ChatInviteAlready from given interface.
func (c *ChatInviteAlready) FillFrom(from interface {
	GetChat() (value ChatClass)
}) {
	c.Chat = from.GetChat()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInviteAlready) TypeID() uint32 {
	return ChatInviteAlreadyTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInviteAlready) TypeName() string {
	return "chatInviteAlready"
}

// TypeInfo returns info about TL type.
func (c *ChatInviteAlready) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInviteAlready",
		ID:   ChatInviteAlreadyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Chat",
			SchemaName: "chat",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInviteAlready) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteAlready#5a686d7c as nil")
	}
	b.PutID(ChatInviteAlreadyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInviteAlready) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInviteAlready#5a686d7c as nil")
	}
	if c.Chat == nil {
		return fmt.Errorf("unable to encode chatInviteAlready#5a686d7c: field chat is nil")
	}
	if err := c.Chat.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInviteAlready#5a686d7c: field chat: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInviteAlready) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteAlready#5a686d7c to nil")
	}
	if err := b.ConsumeID(ChatInviteAlreadyTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInviteAlready#5a686d7c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInviteAlready) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInviteAlready#5a686d7c to nil")
	}
	{
		value, err := DecodeChat(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatInviteAlready#5a686d7c: field chat: %w", err)
		}
		c.Chat = value
	}
	return nil
}

// GetChat returns value of Chat field.
func (c *ChatInviteAlready) GetChat() (value ChatClass) {
	if c == nil {
		return
	}
	return c.Chat
}

// ChatInvite represents TL type `chatInvite#cde0ec40`.
// Chat invite info
//
// See https://core.telegram.org/constructor/chatInvite for reference.
type ChatInvite struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is a channel/supergroup¹ or a normal group²
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	//  2) https://core.telegram.org/api/channel
	Channel bool
	// Whether this is a channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Broadcast bool
	// Whether this is a public channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Public bool
	// Whether this is a supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Megagroup bool
	// Whether the join request »¹ must be first approved by an administrator
	//
	// Links:
	//  1) https://core.telegram.org/api/invites#join-requests
	RequestNeeded bool
	// Is this chat or channel verified by Telegram?
	Verified bool
	// This chat is probably a scam
	Scam bool
	// If set, this chat was reported by many users as a fake or scam: be careful when
	// interacting with it.
	Fake bool
	// Chat/supergroup/channel title
	Title string
	// Description of the group of channel
	//
	// Use SetAbout and GetAbout helpers.
	About string
	// Chat/supergroup/channel photo
	Photo PhotoClass
	// Participant count
	ParticipantsCount int
	// A few of the participants that are in the group
	//
	// Use SetParticipants and GetParticipants helpers.
	Participants []UserClass
	// Profile color palette ID¹
	//
	// Links:
	//  1) https://core.telegram.org/api/colors
	Color int
}

// ChatInviteTypeID is TL type id of ChatInvite.
const ChatInviteTypeID = 0xcde0ec40

// construct implements constructor of ChatInviteClass.
func (c ChatInvite) construct() ChatInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatInvite.
var (
	_ bin.Encoder     = &ChatInvite{}
	_ bin.Decoder     = &ChatInvite{}
	_ bin.BareEncoder = &ChatInvite{}
	_ bin.BareDecoder = &ChatInvite{}

	_ ChatInviteClass = &ChatInvite{}
)

func (c *ChatInvite) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.Channel == false) {
		return false
	}
	if !(c.Broadcast == false) {
		return false
	}
	if !(c.Public == false) {
		return false
	}
	if !(c.Megagroup == false) {
		return false
	}
	if !(c.RequestNeeded == false) {
		return false
	}
	if !(c.Verified == false) {
		return false
	}
	if !(c.Scam == false) {
		return false
	}
	if !(c.Fake == false) {
		return false
	}
	if !(c.Title == "") {
		return false
	}
	if !(c.About == "") {
		return false
	}
	if !(c.Photo == nil) {
		return false
	}
	if !(c.ParticipantsCount == 0) {
		return false
	}
	if !(c.Participants == nil) {
		return false
	}
	if !(c.Color == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInvite) String() string {
	if c == nil {
		return "ChatInvite(nil)"
	}
	type Alias ChatInvite
	return fmt.Sprintf("ChatInvite%+v", Alias(*c))
}

// FillFrom fills ChatInvite from given interface.
func (c *ChatInvite) FillFrom(from interface {
	GetChannel() (value bool)
	GetBroadcast() (value bool)
	GetPublic() (value bool)
	GetMegagroup() (value bool)
	GetRequestNeeded() (value bool)
	GetVerified() (value bool)
	GetScam() (value bool)
	GetFake() (value bool)
	GetTitle() (value string)
	GetAbout() (value string, ok bool)
	GetPhoto() (value PhotoClass)
	GetParticipantsCount() (value int)
	GetParticipants() (value []UserClass, ok bool)
	GetColor() (value int)
}) {
	c.Channel = from.GetChannel()
	c.Broadcast = from.GetBroadcast()
	c.Public = from.GetPublic()
	c.Megagroup = from.GetMegagroup()
	c.RequestNeeded = from.GetRequestNeeded()
	c.Verified = from.GetVerified()
	c.Scam = from.GetScam()
	c.Fake = from.GetFake()
	c.Title = from.GetTitle()
	if val, ok := from.GetAbout(); ok {
		c.About = val
	}

	c.Photo = from.GetPhoto()
	c.ParticipantsCount = from.GetParticipantsCount()
	if val, ok := from.GetParticipants(); ok {
		c.Participants = val
	}

	c.Color = from.GetColor()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInvite) TypeID() uint32 {
	return ChatInviteTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInvite) TypeName() string {
	return "chatInvite"
}

// TypeInfo returns info about TL type.
func (c *ChatInvite) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInvite",
		ID:   ChatInviteTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Channel",
			SchemaName: "channel",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "Broadcast",
			SchemaName: "broadcast",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "Public",
			SchemaName: "public",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "Megagroup",
			SchemaName: "megagroup",
			Null:       !c.Flags.Has(3),
		},
		{
			Name:       "RequestNeeded",
			SchemaName: "request_needed",
			Null:       !c.Flags.Has(6),
		},
		{
			Name:       "Verified",
			SchemaName: "verified",
			Null:       !c.Flags.Has(7),
		},
		{
			Name:       "Scam",
			SchemaName: "scam",
			Null:       !c.Flags.Has(8),
		},
		{
			Name:       "Fake",
			SchemaName: "fake",
			Null:       !c.Flags.Has(9),
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "About",
			SchemaName: "about",
			Null:       !c.Flags.Has(5),
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "ParticipantsCount",
			SchemaName: "participants_count",
		},
		{
			Name:       "Participants",
			SchemaName: "participants",
			Null:       !c.Flags.Has(4),
		},
		{
			Name:       "Color",
			SchemaName: "color",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *ChatInvite) SetFlags() {
	if !(c.Channel == false) {
		c.Flags.Set(0)
	}
	if !(c.Broadcast == false) {
		c.Flags.Set(1)
	}
	if !(c.Public == false) {
		c.Flags.Set(2)
	}
	if !(c.Megagroup == false) {
		c.Flags.Set(3)
	}
	if !(c.RequestNeeded == false) {
		c.Flags.Set(6)
	}
	if !(c.Verified == false) {
		c.Flags.Set(7)
	}
	if !(c.Scam == false) {
		c.Flags.Set(8)
	}
	if !(c.Fake == false) {
		c.Flags.Set(9)
	}
	if !(c.About == "") {
		c.Flags.Set(5)
	}
	if !(c.Participants == nil) {
		c.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (c *ChatInvite) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInvite#cde0ec40 as nil")
	}
	b.PutID(ChatInviteTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInvite) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInvite#cde0ec40 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInvite#cde0ec40: field flags: %w", err)
	}
	b.PutString(c.Title)
	if c.Flags.Has(5) {
		b.PutString(c.About)
	}
	if c.Photo == nil {
		return fmt.Errorf("unable to encode chatInvite#cde0ec40: field photo is nil")
	}
	if err := c.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInvite#cde0ec40: field photo: %w", err)
	}
	b.PutInt(c.ParticipantsCount)
	if c.Flags.Has(4) {
		b.PutVectorHeader(len(c.Participants))
		for idx, v := range c.Participants {
			if v == nil {
				return fmt.Errorf("unable to encode chatInvite#cde0ec40: field participants element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode chatInvite#cde0ec40: field participants element with index %d: %w", idx, err)
			}
		}
	}
	b.PutInt(c.Color)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInvite) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInvite#cde0ec40 to nil")
	}
	if err := b.ConsumeID(ChatInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInvite#cde0ec40: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInvite) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInvite#cde0ec40 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatInvite#cde0ec40: field flags: %w", err)
		}
	}
	c.Channel = c.Flags.Has(0)
	c.Broadcast = c.Flags.Has(1)
	c.Public = c.Flags.Has(2)
	c.Megagroup = c.Flags.Has(3)
	c.RequestNeeded = c.Flags.Has(6)
	c.Verified = c.Flags.Has(7)
	c.Scam = c.Flags.Has(8)
	c.Fake = c.Flags.Has(9)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#cde0ec40: field title: %w", err)
		}
		c.Title = value
	}
	if c.Flags.Has(5) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#cde0ec40: field about: %w", err)
		}
		c.About = value
	}
	{
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#cde0ec40: field photo: %w", err)
		}
		c.Photo = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#cde0ec40: field participants_count: %w", err)
		}
		c.ParticipantsCount = value
	}
	if c.Flags.Has(4) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#cde0ec40: field participants: %w", err)
		}

		if headerLen > 0 {
			c.Participants = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatInvite#cde0ec40: field participants: %w", err)
			}
			c.Participants = append(c.Participants, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvite#cde0ec40: field color: %w", err)
		}
		c.Color = value
	}
	return nil
}

// SetChannel sets value of Channel conditional field.
func (c *ChatInvite) SetChannel(value bool) {
	if value {
		c.Flags.Set(0)
		c.Channel = true
	} else {
		c.Flags.Unset(0)
		c.Channel = false
	}
}

// GetChannel returns value of Channel conditional field.
func (c *ChatInvite) GetChannel() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// SetBroadcast sets value of Broadcast conditional field.
func (c *ChatInvite) SetBroadcast(value bool) {
	if value {
		c.Flags.Set(1)
		c.Broadcast = true
	} else {
		c.Flags.Unset(1)
		c.Broadcast = false
	}
}

// GetBroadcast returns value of Broadcast conditional field.
func (c *ChatInvite) GetBroadcast() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(1)
}

// SetPublic sets value of Public conditional field.
func (c *ChatInvite) SetPublic(value bool) {
	if value {
		c.Flags.Set(2)
		c.Public = true
	} else {
		c.Flags.Unset(2)
		c.Public = false
	}
}

// GetPublic returns value of Public conditional field.
func (c *ChatInvite) GetPublic() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(2)
}

// SetMegagroup sets value of Megagroup conditional field.
func (c *ChatInvite) SetMegagroup(value bool) {
	if value {
		c.Flags.Set(3)
		c.Megagroup = true
	} else {
		c.Flags.Unset(3)
		c.Megagroup = false
	}
}

// GetMegagroup returns value of Megagroup conditional field.
func (c *ChatInvite) GetMegagroup() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(3)
}

// SetRequestNeeded sets value of RequestNeeded conditional field.
func (c *ChatInvite) SetRequestNeeded(value bool) {
	if value {
		c.Flags.Set(6)
		c.RequestNeeded = true
	} else {
		c.Flags.Unset(6)
		c.RequestNeeded = false
	}
}

// GetRequestNeeded returns value of RequestNeeded conditional field.
func (c *ChatInvite) GetRequestNeeded() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(6)
}

// SetVerified sets value of Verified conditional field.
func (c *ChatInvite) SetVerified(value bool) {
	if value {
		c.Flags.Set(7)
		c.Verified = true
	} else {
		c.Flags.Unset(7)
		c.Verified = false
	}
}

// GetVerified returns value of Verified conditional field.
func (c *ChatInvite) GetVerified() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(7)
}

// SetScam sets value of Scam conditional field.
func (c *ChatInvite) SetScam(value bool) {
	if value {
		c.Flags.Set(8)
		c.Scam = true
	} else {
		c.Flags.Unset(8)
		c.Scam = false
	}
}

// GetScam returns value of Scam conditional field.
func (c *ChatInvite) GetScam() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(8)
}

// SetFake sets value of Fake conditional field.
func (c *ChatInvite) SetFake(value bool) {
	if value {
		c.Flags.Set(9)
		c.Fake = true
	} else {
		c.Flags.Unset(9)
		c.Fake = false
	}
}

// GetFake returns value of Fake conditional field.
func (c *ChatInvite) GetFake() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(9)
}

// GetTitle returns value of Title field.
func (c *ChatInvite) GetTitle() (value string) {
	if c == nil {
		return
	}
	return c.Title
}

// SetAbout sets value of About conditional field.
func (c *ChatInvite) SetAbout(value string) {
	c.Flags.Set(5)
	c.About = value
}

// GetAbout returns value of About conditional field and
// boolean which is true if field was set.
func (c *ChatInvite) GetAbout() (value string, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(5) {
		return value, false
	}
	return c.About, true
}

// GetPhoto returns value of Photo field.
func (c *ChatInvite) GetPhoto() (value PhotoClass) {
	if c == nil {
		return
	}
	return c.Photo
}

// GetParticipantsCount returns value of ParticipantsCount field.
func (c *ChatInvite) GetParticipantsCount() (value int) {
	if c == nil {
		return
	}
	return c.ParticipantsCount
}

// SetParticipants sets value of Participants conditional field.
func (c *ChatInvite) SetParticipants(value []UserClass) {
	c.Flags.Set(4)
	c.Participants = value
}

// GetParticipants returns value of Participants conditional field and
// boolean which is true if field was set.
func (c *ChatInvite) GetParticipants() (value []UserClass, ok bool) {
	if c == nil {
		return
	}
	if !c.Flags.Has(4) {
		return value, false
	}
	return c.Participants, true
}

// GetColor returns value of Color field.
func (c *ChatInvite) GetColor() (value int) {
	if c == nil {
		return
	}
	return c.Color
}

// MapParticipants returns field Participants wrapped in UserClassArray helper.
func (c *ChatInvite) MapParticipants() (value UserClassArray, ok bool) {
	if !c.Flags.Has(4) {
		return value, false
	}
	return UserClassArray(c.Participants), true
}

// ChatInvitePeek represents TL type `chatInvitePeek#61695cb0`.
// A chat invitation that also allows peeking into the group to read messages without
// joining it.
//
// See https://core.telegram.org/constructor/chatInvitePeek for reference.
type ChatInvitePeek struct {
	// Chat information
	Chat ChatClass
	// Read-only anonymous access to this group will be revoked at this date
	Expires int
}

// ChatInvitePeekTypeID is TL type id of ChatInvitePeek.
const ChatInvitePeekTypeID = 0x61695cb0

// construct implements constructor of ChatInviteClass.
func (c ChatInvitePeek) construct() ChatInviteClass { return &c }

// Ensuring interfaces in compile-time for ChatInvitePeek.
var (
	_ bin.Encoder     = &ChatInvitePeek{}
	_ bin.Decoder     = &ChatInvitePeek{}
	_ bin.BareEncoder = &ChatInvitePeek{}
	_ bin.BareDecoder = &ChatInvitePeek{}

	_ ChatInviteClass = &ChatInvitePeek{}
)

func (c *ChatInvitePeek) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Chat == nil) {
		return false
	}
	if !(c.Expires == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatInvitePeek) String() string {
	if c == nil {
		return "ChatInvitePeek(nil)"
	}
	type Alias ChatInvitePeek
	return fmt.Sprintf("ChatInvitePeek%+v", Alias(*c))
}

// FillFrom fills ChatInvitePeek from given interface.
func (c *ChatInvitePeek) FillFrom(from interface {
	GetChat() (value ChatClass)
	GetExpires() (value int)
}) {
	c.Chat = from.GetChat()
	c.Expires = from.GetExpires()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatInvitePeek) TypeID() uint32 {
	return ChatInvitePeekTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatInvitePeek) TypeName() string {
	return "chatInvitePeek"
}

// TypeInfo returns info about TL type.
func (c *ChatInvitePeek) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatInvitePeek",
		ID:   ChatInvitePeekTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Chat",
			SchemaName: "chat",
		},
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatInvitePeek) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInvitePeek#61695cb0 as nil")
	}
	b.PutID(ChatInvitePeekTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatInvitePeek) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatInvitePeek#61695cb0 as nil")
	}
	if c.Chat == nil {
		return fmt.Errorf("unable to encode chatInvitePeek#61695cb0: field chat is nil")
	}
	if err := c.Chat.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatInvitePeek#61695cb0: field chat: %w", err)
	}
	b.PutInt(c.Expires)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatInvitePeek) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInvitePeek#61695cb0 to nil")
	}
	if err := b.ConsumeID(ChatInvitePeekTypeID); err != nil {
		return fmt.Errorf("unable to decode chatInvitePeek#61695cb0: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatInvitePeek) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatInvitePeek#61695cb0 to nil")
	}
	{
		value, err := DecodeChat(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatInvitePeek#61695cb0: field chat: %w", err)
		}
		c.Chat = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatInvitePeek#61695cb0: field expires: %w", err)
		}
		c.Expires = value
	}
	return nil
}

// GetChat returns value of Chat field.
func (c *ChatInvitePeek) GetChat() (value ChatClass) {
	if c == nil {
		return
	}
	return c.Chat
}

// GetExpires returns value of Expires field.
func (c *ChatInvitePeek) GetExpires() (value int) {
	if c == nil {
		return
	}
	return c.Expires
}

// ChatInviteClassName is schema name of ChatInviteClass.
const ChatInviteClassName = "ChatInvite"

// ChatInviteClass represents ChatInvite generic type.
//
// See https://core.telegram.org/type/ChatInvite for reference.
//
// Example:
//
//	g, err := tg.DecodeChatInvite(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ChatInviteAlready: // chatInviteAlready#5a686d7c
//	case *tg.ChatInvite: // chatInvite#cde0ec40
//	case *tg.ChatInvitePeek: // chatInvitePeek#61695cb0
//	default: panic(v)
//	}
type ChatInviteClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatInviteClass

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
}

// DecodeChatInvite implements binary de-serialization for ChatInviteClass.
func DecodeChatInvite(buf *bin.Buffer) (ChatInviteClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatInviteAlreadyTypeID:
		// Decoding chatInviteAlready#5a686d7c.
		v := ChatInviteAlready{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatInviteClass: %w", err)
		}
		return &v, nil
	case ChatInviteTypeID:
		// Decoding chatInvite#cde0ec40.
		v := ChatInvite{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatInviteClass: %w", err)
		}
		return &v, nil
	case ChatInvitePeekTypeID:
		// Decoding chatInvitePeek#61695cb0.
		v := ChatInvitePeek{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatInviteClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatInviteClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatInvite boxes the ChatInviteClass providing a helper.
type ChatInviteBox struct {
	ChatInvite ChatInviteClass
}

// Decode implements bin.Decoder for ChatInviteBox.
func (b *ChatInviteBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatInviteBox to nil")
	}
	v, err := DecodeChatInvite(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatInvite = v
	return nil
}

// Encode implements bin.Encode for ChatInviteBox.
func (b *ChatInviteBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatInvite == nil {
		return fmt.Errorf("unable to encode ChatInviteClass as nil")
	}
	return b.ChatInvite.Encode(buf)
}
