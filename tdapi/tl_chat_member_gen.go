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

// ChatMember represents TL type `chatMember#1265e1e8`.
type ChatMember struct {
	// Identifier of the chat member. Currently, other chats can be only Left or Banned. Only
	// supergroups and channels can have other chats as Left or Banned members and these
	// chats must be supergroups or channels
	MemberID MessageSenderClass
	// Identifier of a user that invited/promoted/banned this member in the chat; 0 if
	// unknown
	InviterUserID int32
	// Point in time (Unix timestamp) when the user joined the chat
	JoinedChatDate int32
	// Status of the member in the chat
	Status ChatMemberStatusClass
}

// ChatMemberTypeID is TL type id of ChatMember.
const ChatMemberTypeID = 0x1265e1e8

// Ensuring interfaces in compile-time for ChatMember.
var (
	_ bin.Encoder     = &ChatMember{}
	_ bin.Decoder     = &ChatMember{}
	_ bin.BareEncoder = &ChatMember{}
	_ bin.BareDecoder = &ChatMember{}
)

func (c *ChatMember) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.MemberID == nil) {
		return false
	}
	if !(c.InviterUserID == 0) {
		return false
	}
	if !(c.JoinedChatDate == 0) {
		return false
	}
	if !(c.Status == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMember) String() string {
	if c == nil {
		return "ChatMember(nil)"
	}
	type Alias ChatMember
	return fmt.Sprintf("ChatMember%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMember) TypeID() uint32 {
	return ChatMemberTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMember) TypeName() string {
	return "chatMember"
}

// TypeInfo returns info about TL type.
func (c *ChatMember) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMember",
		ID:   ChatMemberTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MemberID",
			SchemaName: "member_id",
		},
		{
			Name:       "InviterUserID",
			SchemaName: "inviter_user_id",
		},
		{
			Name:       "JoinedChatDate",
			SchemaName: "joined_chat_date",
		},
		{
			Name:       "Status",
			SchemaName: "status",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMember) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMember#1265e1e8 as nil")
	}
	b.PutID(ChatMemberTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMember) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMember#1265e1e8 as nil")
	}
	if c.MemberID == nil {
		return fmt.Errorf("unable to encode chatMember#1265e1e8: field member_id is nil")
	}
	if err := c.MemberID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatMember#1265e1e8: field member_id: %w", err)
	}
	b.PutInt32(c.InviterUserID)
	b.PutInt32(c.JoinedChatDate)
	if c.Status == nil {
		return fmt.Errorf("unable to encode chatMember#1265e1e8: field status is nil")
	}
	if err := c.Status.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatMember#1265e1e8: field status: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMember) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMember#1265e1e8 to nil")
	}
	if err := b.ConsumeID(ChatMemberTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMember#1265e1e8: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMember) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMember#1265e1e8 to nil")
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatMember#1265e1e8: field member_id: %w", err)
		}
		c.MemberID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatMember#1265e1e8: field inviter_user_id: %w", err)
		}
		c.InviterUserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatMember#1265e1e8: field joined_chat_date: %w", err)
		}
		c.JoinedChatDate = value
	}
	{
		value, err := DecodeChatMemberStatus(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatMember#1265e1e8: field status: %w", err)
		}
		c.Status = value
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatMember) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMember#1265e1e8 as nil")
	}
	b.ObjStart()
	b.PutID("chatMember")
	b.FieldStart("member_id")
	if c.MemberID == nil {
		return fmt.Errorf("unable to encode chatMember#1265e1e8: field member_id is nil")
	}
	if err := c.MemberID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatMember#1265e1e8: field member_id: %w", err)
	}
	b.FieldStart("inviter_user_id")
	b.PutInt32(c.InviterUserID)
	b.FieldStart("joined_chat_date")
	b.PutInt32(c.JoinedChatDate)
	b.FieldStart("status")
	if c.Status == nil {
		return fmt.Errorf("unable to encode chatMember#1265e1e8: field status is nil")
	}
	if err := c.Status.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatMember#1265e1e8: field status: %w", err)
	}
	b.ObjEnd()
	return nil
}

// GetMemberID returns value of MemberID field.
func (c *ChatMember) GetMemberID() (value MessageSenderClass) {
	return c.MemberID
}

// GetInviterUserID returns value of InviterUserID field.
func (c *ChatMember) GetInviterUserID() (value int32) {
	return c.InviterUserID
}

// GetJoinedChatDate returns value of JoinedChatDate field.
func (c *ChatMember) GetJoinedChatDate() (value int32) {
	return c.JoinedChatDate
}

// GetStatus returns value of Status field.
func (c *ChatMember) GetStatus() (value ChatMemberStatusClass) {
	return c.Status
}
