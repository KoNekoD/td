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

// ChatPermissions represents TL type `chatPermissions#5e73d8df`.
type ChatPermissions struct {
	// True, if the user can send text messages, contacts, locations, and venues
	CanSendMessages bool
	// True, if the user can send audio files, documents, photos, videos, video notes, and
	// voice notes. Implies can_send_messages permissions
	CanSendMediaMessages bool
	// True, if the user can send polls. Implies can_send_messages permissions
	CanSendPolls bool
	// True, if the user can send animations, games, stickers, and dice and use inline bots.
	// Implies can_send_messages permissions
	CanSendOtherMessages bool
	// True, if the user may add a web page preview to their messages. Implies
	// can_send_messages permissions
	CanAddWebPagePreviews bool
	// True, if the user can change the chat title, photo, and other settings
	CanChangeInfo bool
	// True, if the user can invite new users to the chat
	CanInviteUsers bool
	// True, if the user can pin messages
	CanPinMessages bool
}

// ChatPermissionsTypeID is TL type id of ChatPermissions.
const ChatPermissionsTypeID = 0x5e73d8df

// Ensuring interfaces in compile-time for ChatPermissions.
var (
	_ bin.Encoder     = &ChatPermissions{}
	_ bin.Decoder     = &ChatPermissions{}
	_ bin.BareEncoder = &ChatPermissions{}
	_ bin.BareDecoder = &ChatPermissions{}
)

func (c *ChatPermissions) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.CanSendMessages == false) {
		return false
	}
	if !(c.CanSendMediaMessages == false) {
		return false
	}
	if !(c.CanSendPolls == false) {
		return false
	}
	if !(c.CanSendOtherMessages == false) {
		return false
	}
	if !(c.CanAddWebPagePreviews == false) {
		return false
	}
	if !(c.CanChangeInfo == false) {
		return false
	}
	if !(c.CanInviteUsers == false) {
		return false
	}
	if !(c.CanPinMessages == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatPermissions) String() string {
	if c == nil {
		return "ChatPermissions(nil)"
	}
	type Alias ChatPermissions
	return fmt.Sprintf("ChatPermissions%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatPermissions) TypeID() uint32 {
	return ChatPermissionsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatPermissions) TypeName() string {
	return "chatPermissions"
}

// TypeInfo returns info about TL type.
func (c *ChatPermissions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatPermissions",
		ID:   ChatPermissionsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CanSendMessages",
			SchemaName: "can_send_messages",
		},
		{
			Name:       "CanSendMediaMessages",
			SchemaName: "can_send_media_messages",
		},
		{
			Name:       "CanSendPolls",
			SchemaName: "can_send_polls",
		},
		{
			Name:       "CanSendOtherMessages",
			SchemaName: "can_send_other_messages",
		},
		{
			Name:       "CanAddWebPagePreviews",
			SchemaName: "can_add_web_page_previews",
		},
		{
			Name:       "CanChangeInfo",
			SchemaName: "can_change_info",
		},
		{
			Name:       "CanInviteUsers",
			SchemaName: "can_invite_users",
		},
		{
			Name:       "CanPinMessages",
			SchemaName: "can_pin_messages",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatPermissions) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPermissions#5e73d8df as nil")
	}
	b.PutID(ChatPermissionsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatPermissions) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPermissions#5e73d8df as nil")
	}
	b.PutBool(c.CanSendMessages)
	b.PutBool(c.CanSendMediaMessages)
	b.PutBool(c.CanSendPolls)
	b.PutBool(c.CanSendOtherMessages)
	b.PutBool(c.CanAddWebPagePreviews)
	b.PutBool(c.CanChangeInfo)
	b.PutBool(c.CanInviteUsers)
	b.PutBool(c.CanPinMessages)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatPermissions) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPermissions#5e73d8df to nil")
	}
	if err := b.ConsumeID(ChatPermissionsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatPermissions#5e73d8df: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatPermissions) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatPermissions#5e73d8df to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatPermissions#5e73d8df: field can_send_messages: %w", err)
		}
		c.CanSendMessages = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatPermissions#5e73d8df: field can_send_media_messages: %w", err)
		}
		c.CanSendMediaMessages = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatPermissions#5e73d8df: field can_send_polls: %w", err)
		}
		c.CanSendPolls = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatPermissions#5e73d8df: field can_send_other_messages: %w", err)
		}
		c.CanSendOtherMessages = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatPermissions#5e73d8df: field can_add_web_page_previews: %w", err)
		}
		c.CanAddWebPagePreviews = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatPermissions#5e73d8df: field can_change_info: %w", err)
		}
		c.CanChangeInfo = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatPermissions#5e73d8df: field can_invite_users: %w", err)
		}
		c.CanInviteUsers = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatPermissions#5e73d8df: field can_pin_messages: %w", err)
		}
		c.CanPinMessages = value
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatPermissions) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatPermissions#5e73d8df as nil")
	}
	b.ObjStart()
	b.PutID("chatPermissions")
	b.FieldStart("can_send_messages")
	b.PutBool(c.CanSendMessages)
	b.FieldStart("can_send_media_messages")
	b.PutBool(c.CanSendMediaMessages)
	b.FieldStart("can_send_polls")
	b.PutBool(c.CanSendPolls)
	b.FieldStart("can_send_other_messages")
	b.PutBool(c.CanSendOtherMessages)
	b.FieldStart("can_add_web_page_previews")
	b.PutBool(c.CanAddWebPagePreviews)
	b.FieldStart("can_change_info")
	b.PutBool(c.CanChangeInfo)
	b.FieldStart("can_invite_users")
	b.PutBool(c.CanInviteUsers)
	b.FieldStart("can_pin_messages")
	b.PutBool(c.CanPinMessages)
	b.ObjEnd()
	return nil
}

// GetCanSendMessages returns value of CanSendMessages field.
func (c *ChatPermissions) GetCanSendMessages() (value bool) {
	return c.CanSendMessages
}

// GetCanSendMediaMessages returns value of CanSendMediaMessages field.
func (c *ChatPermissions) GetCanSendMediaMessages() (value bool) {
	return c.CanSendMediaMessages
}

// GetCanSendPolls returns value of CanSendPolls field.
func (c *ChatPermissions) GetCanSendPolls() (value bool) {
	return c.CanSendPolls
}

// GetCanSendOtherMessages returns value of CanSendOtherMessages field.
func (c *ChatPermissions) GetCanSendOtherMessages() (value bool) {
	return c.CanSendOtherMessages
}

// GetCanAddWebPagePreviews returns value of CanAddWebPagePreviews field.
func (c *ChatPermissions) GetCanAddWebPagePreviews() (value bool) {
	return c.CanAddWebPagePreviews
}

// GetCanChangeInfo returns value of CanChangeInfo field.
func (c *ChatPermissions) GetCanChangeInfo() (value bool) {
	return c.CanChangeInfo
}

// GetCanInviteUsers returns value of CanInviteUsers field.
func (c *ChatPermissions) GetCanInviteUsers() (value bool) {
	return c.CanInviteUsers
}

// GetCanPinMessages returns value of CanPinMessages field.
func (c *ChatPermissions) GetCanPinMessages() (value bool) {
	return c.CanPinMessages
}
