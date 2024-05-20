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

// ChatAdminRights represents TL type `chatAdminRights#5fb224d5`.
// Represents the rights of an admin in a channel/supergroup¹.
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/constructor/chatAdminRights for reference.
type ChatAdminRights struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, allows the admin to modify the description of the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	ChangeInfo bool
	// If set, allows the admin to post messages in the channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	PostMessages bool
	// If set, allows the admin to also edit messages from other admins in the channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	EditMessages bool
	// If set, allows the admin to also delete messages from other admins in the channel¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	DeleteMessages bool
	// If set, allows the admin to ban users from the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	BanUsers bool
	// If set, allows the admin to invite users in the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	InviteUsers bool
	// If set, allows the admin to pin messages in the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	PinMessages bool
	// If set, allows the admin to add other admins with the same (or more limited)
	// permissions in the channel/supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	AddAdmins bool
	// Whether this admin is anonymous
	Anonymous bool
	// If set, allows the admin to change group call/livestream settings
	ManageCall bool
	// Set this flag if none of the other flags are set, but you still want the user to be an
	// admin: if this or any of the other flags are set, the admin can get the chat admin
	// log¹, get chat statistics², get message statistics in channels³, get channel
	// members, see anonymous administrators in supergroups and ignore slow mode.
	//
	// Links:
	//  1) https://core.telegram.org/api/recent-actions
	//  2) https://core.telegram.org/api/stats
	//  3) https://core.telegram.org/api/stats
	Other bool
	// If set, allows the admin to create, delete or modify forum topics »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/forum#forum-topics
	ManageTopics bool
	// If set, allows the admin to post stories¹ as the channel².
	//
	// Links:
	//  1) https://core.telegram.org/api/stories
	//  2) https://core.telegram.org/api/channel
	PostStories bool
	// If set, allows the admin to edit stories¹ posted by the other admins of the channel².
	//
	// Links:
	//  1) https://core.telegram.org/api/stories
	//  2) https://core.telegram.org/api/channel
	EditStories bool
	// If set, allows the admin to delete stories¹ posted by the other admins of the
	// channel².
	//
	// Links:
	//  1) https://core.telegram.org/api/stories
	//  2) https://core.telegram.org/api/channel
	DeleteStories bool
}

// ChatAdminRightsTypeID is TL type id of ChatAdminRights.
const ChatAdminRightsTypeID = 0x5fb224d5

// Ensuring interfaces in compile-time for ChatAdminRights.
var (
	_ bin.Encoder     = &ChatAdminRights{}
	_ bin.Decoder     = &ChatAdminRights{}
	_ bin.BareEncoder = &ChatAdminRights{}
	_ bin.BareDecoder = &ChatAdminRights{}
)

func (c *ChatAdminRights) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.ChangeInfo == false) {
		return false
	}
	if !(c.PostMessages == false) {
		return false
	}
	if !(c.EditMessages == false) {
		return false
	}
	if !(c.DeleteMessages == false) {
		return false
	}
	if !(c.BanUsers == false) {
		return false
	}
	if !(c.InviteUsers == false) {
		return false
	}
	if !(c.PinMessages == false) {
		return false
	}
	if !(c.AddAdmins == false) {
		return false
	}
	if !(c.Anonymous == false) {
		return false
	}
	if !(c.ManageCall == false) {
		return false
	}
	if !(c.Other == false) {
		return false
	}
	if !(c.ManageTopics == false) {
		return false
	}
	if !(c.PostStories == false) {
		return false
	}
	if !(c.EditStories == false) {
		return false
	}
	if !(c.DeleteStories == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatAdminRights) String() string {
	if c == nil {
		return "ChatAdminRights(nil)"
	}
	type Alias ChatAdminRights
	return fmt.Sprintf("ChatAdminRights%+v", Alias(*c))
}

// FillFrom fills ChatAdminRights from given interface.
func (c *ChatAdminRights) FillFrom(from interface {
	GetChangeInfo() (value bool)
	GetPostMessages() (value bool)
	GetEditMessages() (value bool)
	GetDeleteMessages() (value bool)
	GetBanUsers() (value bool)
	GetInviteUsers() (value bool)
	GetPinMessages() (value bool)
	GetAddAdmins() (value bool)
	GetAnonymous() (value bool)
	GetManageCall() (value bool)
	GetOther() (value bool)
	GetManageTopics() (value bool)
	GetPostStories() (value bool)
	GetEditStories() (value bool)
	GetDeleteStories() (value bool)
}) {
	c.ChangeInfo = from.GetChangeInfo()
	c.PostMessages = from.GetPostMessages()
	c.EditMessages = from.GetEditMessages()
	c.DeleteMessages = from.GetDeleteMessages()
	c.BanUsers = from.GetBanUsers()
	c.InviteUsers = from.GetInviteUsers()
	c.PinMessages = from.GetPinMessages()
	c.AddAdmins = from.GetAddAdmins()
	c.Anonymous = from.GetAnonymous()
	c.ManageCall = from.GetManageCall()
	c.Other = from.GetOther()
	c.ManageTopics = from.GetManageTopics()
	c.PostStories = from.GetPostStories()
	c.EditStories = from.GetEditStories()
	c.DeleteStories = from.GetDeleteStories()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatAdminRights) TypeID() uint32 {
	return ChatAdminRightsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatAdminRights) TypeName() string {
	return "chatAdminRights"
}

// TypeInfo returns info about TL type.
func (c *ChatAdminRights) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatAdminRights",
		ID:   ChatAdminRightsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChangeInfo",
			SchemaName: "change_info",
			Null:       !c.Flags.Has(0),
		},
		{
			Name:       "PostMessages",
			SchemaName: "post_messages",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "EditMessages",
			SchemaName: "edit_messages",
			Null:       !c.Flags.Has(2),
		},
		{
			Name:       "DeleteMessages",
			SchemaName: "delete_messages",
			Null:       !c.Flags.Has(3),
		},
		{
			Name:       "BanUsers",
			SchemaName: "ban_users",
			Null:       !c.Flags.Has(4),
		},
		{
			Name:       "InviteUsers",
			SchemaName: "invite_users",
			Null:       !c.Flags.Has(5),
		},
		{
			Name:       "PinMessages",
			SchemaName: "pin_messages",
			Null:       !c.Flags.Has(7),
		},
		{
			Name:       "AddAdmins",
			SchemaName: "add_admins",
			Null:       !c.Flags.Has(9),
		},
		{
			Name:       "Anonymous",
			SchemaName: "anonymous",
			Null:       !c.Flags.Has(10),
		},
		{
			Name:       "ManageCall",
			SchemaName: "manage_call",
			Null:       !c.Flags.Has(11),
		},
		{
			Name:       "Other",
			SchemaName: "other",
			Null:       !c.Flags.Has(12),
		},
		{
			Name:       "ManageTopics",
			SchemaName: "manage_topics",
			Null:       !c.Flags.Has(13),
		},
		{
			Name:       "PostStories",
			SchemaName: "post_stories",
			Null:       !c.Flags.Has(14),
		},
		{
			Name:       "EditStories",
			SchemaName: "edit_stories",
			Null:       !c.Flags.Has(15),
		},
		{
			Name:       "DeleteStories",
			SchemaName: "delete_stories",
			Null:       !c.Flags.Has(16),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *ChatAdminRights) SetFlags() {
	if !(c.ChangeInfo == false) {
		c.Flags.Set(0)
	}
	if !(c.PostMessages == false) {
		c.Flags.Set(1)
	}
	if !(c.EditMessages == false) {
		c.Flags.Set(2)
	}
	if !(c.DeleteMessages == false) {
		c.Flags.Set(3)
	}
	if !(c.BanUsers == false) {
		c.Flags.Set(4)
	}
	if !(c.InviteUsers == false) {
		c.Flags.Set(5)
	}
	if !(c.PinMessages == false) {
		c.Flags.Set(7)
	}
	if !(c.AddAdmins == false) {
		c.Flags.Set(9)
	}
	if !(c.Anonymous == false) {
		c.Flags.Set(10)
	}
	if !(c.ManageCall == false) {
		c.Flags.Set(11)
	}
	if !(c.Other == false) {
		c.Flags.Set(12)
	}
	if !(c.ManageTopics == false) {
		c.Flags.Set(13)
	}
	if !(c.PostStories == false) {
		c.Flags.Set(14)
	}
	if !(c.EditStories == false) {
		c.Flags.Set(15)
	}
	if !(c.DeleteStories == false) {
		c.Flags.Set(16)
	}
}

// Encode implements bin.Encoder.
func (c *ChatAdminRights) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdminRights#5fb224d5 as nil")
	}
	b.PutID(ChatAdminRightsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatAdminRights) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatAdminRights#5fb224d5 as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatAdminRights#5fb224d5: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatAdminRights) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdminRights#5fb224d5 to nil")
	}
	if err := b.ConsumeID(ChatAdminRightsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatAdminRights#5fb224d5: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatAdminRights) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatAdminRights#5fb224d5 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatAdminRights#5fb224d5: field flags: %w", err)
		}
	}
	c.ChangeInfo = c.Flags.Has(0)
	c.PostMessages = c.Flags.Has(1)
	c.EditMessages = c.Flags.Has(2)
	c.DeleteMessages = c.Flags.Has(3)
	c.BanUsers = c.Flags.Has(4)
	c.InviteUsers = c.Flags.Has(5)
	c.PinMessages = c.Flags.Has(7)
	c.AddAdmins = c.Flags.Has(9)
	c.Anonymous = c.Flags.Has(10)
	c.ManageCall = c.Flags.Has(11)
	c.Other = c.Flags.Has(12)
	c.ManageTopics = c.Flags.Has(13)
	c.PostStories = c.Flags.Has(14)
	c.EditStories = c.Flags.Has(15)
	c.DeleteStories = c.Flags.Has(16)
	return nil
}

// SetChangeInfo sets value of ChangeInfo conditional field.
func (c *ChatAdminRights) SetChangeInfo(value bool) {
	if value {
		c.Flags.Set(0)
		c.ChangeInfo = true
	} else {
		c.Flags.Unset(0)
		c.ChangeInfo = false
	}
}

// GetChangeInfo returns value of ChangeInfo conditional field.
func (c *ChatAdminRights) GetChangeInfo() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// SetPostMessages sets value of PostMessages conditional field.
func (c *ChatAdminRights) SetPostMessages(value bool) {
	if value {
		c.Flags.Set(1)
		c.PostMessages = true
	} else {
		c.Flags.Unset(1)
		c.PostMessages = false
	}
}

// GetPostMessages returns value of PostMessages conditional field.
func (c *ChatAdminRights) GetPostMessages() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(1)
}

// SetEditMessages sets value of EditMessages conditional field.
func (c *ChatAdminRights) SetEditMessages(value bool) {
	if value {
		c.Flags.Set(2)
		c.EditMessages = true
	} else {
		c.Flags.Unset(2)
		c.EditMessages = false
	}
}

// GetEditMessages returns value of EditMessages conditional field.
func (c *ChatAdminRights) GetEditMessages() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(2)
}

// SetDeleteMessages sets value of DeleteMessages conditional field.
func (c *ChatAdminRights) SetDeleteMessages(value bool) {
	if value {
		c.Flags.Set(3)
		c.DeleteMessages = true
	} else {
		c.Flags.Unset(3)
		c.DeleteMessages = false
	}
}

// GetDeleteMessages returns value of DeleteMessages conditional field.
func (c *ChatAdminRights) GetDeleteMessages() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(3)
}

// SetBanUsers sets value of BanUsers conditional field.
func (c *ChatAdminRights) SetBanUsers(value bool) {
	if value {
		c.Flags.Set(4)
		c.BanUsers = true
	} else {
		c.Flags.Unset(4)
		c.BanUsers = false
	}
}

// GetBanUsers returns value of BanUsers conditional field.
func (c *ChatAdminRights) GetBanUsers() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(4)
}

// SetInviteUsers sets value of InviteUsers conditional field.
func (c *ChatAdminRights) SetInviteUsers(value bool) {
	if value {
		c.Flags.Set(5)
		c.InviteUsers = true
	} else {
		c.Flags.Unset(5)
		c.InviteUsers = false
	}
}

// GetInviteUsers returns value of InviteUsers conditional field.
func (c *ChatAdminRights) GetInviteUsers() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(5)
}

// SetPinMessages sets value of PinMessages conditional field.
func (c *ChatAdminRights) SetPinMessages(value bool) {
	if value {
		c.Flags.Set(7)
		c.PinMessages = true
	} else {
		c.Flags.Unset(7)
		c.PinMessages = false
	}
}

// GetPinMessages returns value of PinMessages conditional field.
func (c *ChatAdminRights) GetPinMessages() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(7)
}

// SetAddAdmins sets value of AddAdmins conditional field.
func (c *ChatAdminRights) SetAddAdmins(value bool) {
	if value {
		c.Flags.Set(9)
		c.AddAdmins = true
	} else {
		c.Flags.Unset(9)
		c.AddAdmins = false
	}
}

// GetAddAdmins returns value of AddAdmins conditional field.
func (c *ChatAdminRights) GetAddAdmins() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(9)
}

// SetAnonymous sets value of Anonymous conditional field.
func (c *ChatAdminRights) SetAnonymous(value bool) {
	if value {
		c.Flags.Set(10)
		c.Anonymous = true
	} else {
		c.Flags.Unset(10)
		c.Anonymous = false
	}
}

// GetAnonymous returns value of Anonymous conditional field.
func (c *ChatAdminRights) GetAnonymous() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(10)
}

// SetManageCall sets value of ManageCall conditional field.
func (c *ChatAdminRights) SetManageCall(value bool) {
	if value {
		c.Flags.Set(11)
		c.ManageCall = true
	} else {
		c.Flags.Unset(11)
		c.ManageCall = false
	}
}

// GetManageCall returns value of ManageCall conditional field.
func (c *ChatAdminRights) GetManageCall() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(11)
}

// SetOther sets value of Other conditional field.
func (c *ChatAdminRights) SetOther(value bool) {
	if value {
		c.Flags.Set(12)
		c.Other = true
	} else {
		c.Flags.Unset(12)
		c.Other = false
	}
}

// GetOther returns value of Other conditional field.
func (c *ChatAdminRights) GetOther() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(12)
}

// SetManageTopics sets value of ManageTopics conditional field.
func (c *ChatAdminRights) SetManageTopics(value bool) {
	if value {
		c.Flags.Set(13)
		c.ManageTopics = true
	} else {
		c.Flags.Unset(13)
		c.ManageTopics = false
	}
}

// GetManageTopics returns value of ManageTopics conditional field.
func (c *ChatAdminRights) GetManageTopics() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(13)
}

// SetPostStories sets value of PostStories conditional field.
func (c *ChatAdminRights) SetPostStories(value bool) {
	if value {
		c.Flags.Set(14)
		c.PostStories = true
	} else {
		c.Flags.Unset(14)
		c.PostStories = false
	}
}

// GetPostStories returns value of PostStories conditional field.
func (c *ChatAdminRights) GetPostStories() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(14)
}

// SetEditStories sets value of EditStories conditional field.
func (c *ChatAdminRights) SetEditStories(value bool) {
	if value {
		c.Flags.Set(15)
		c.EditStories = true
	} else {
		c.Flags.Unset(15)
		c.EditStories = false
	}
}

// GetEditStories returns value of EditStories conditional field.
func (c *ChatAdminRights) GetEditStories() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(15)
}

// SetDeleteStories sets value of DeleteStories conditional field.
func (c *ChatAdminRights) SetDeleteStories(value bool) {
	if value {
		c.Flags.Set(16)
		c.DeleteStories = true
	} else {
		c.Flags.Unset(16)
		c.DeleteStories = false
	}
}

// GetDeleteStories returns value of DeleteStories conditional field.
func (c *ChatAdminRights) GetDeleteStories() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(16)
}
