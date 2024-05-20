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

// MessagesHideChatJoinRequestRequest represents TL type `messages.hideChatJoinRequest#7fe7e815`.
// Dismiss or approve a chat join request¹ related to a specific chat or channel.
//
// Links:
//  1. https://core.telegram.org/api/invites#join-requests
//
// See https://core.telegram.org/method/messages.hideChatJoinRequest for reference.
type MessagesHideChatJoinRequestRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to dismiss or approve the chat join request »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/invites#join-requests
	Approved bool
	// The chat or channel
	Peer InputPeerClass
	// The user whose join request »¹ should be dismissed or approved
	//
	// Links:
	//  1) https://core.telegram.org/api/invites#join-requests
	UserID InputUserClass
}

// MessagesHideChatJoinRequestRequestTypeID is TL type id of MessagesHideChatJoinRequestRequest.
const MessagesHideChatJoinRequestRequestTypeID = 0x7fe7e815

// Ensuring interfaces in compile-time for MessagesHideChatJoinRequestRequest.
var (
	_ bin.Encoder     = &MessagesHideChatJoinRequestRequest{}
	_ bin.Decoder     = &MessagesHideChatJoinRequestRequest{}
	_ bin.BareEncoder = &MessagesHideChatJoinRequestRequest{}
	_ bin.BareDecoder = &MessagesHideChatJoinRequestRequest{}
)

func (h *MessagesHideChatJoinRequestRequest) Zero() bool {
	if h == nil {
		return true
	}
	if !(h.Flags.Zero()) {
		return false
	}
	if !(h.Approved == false) {
		return false
	}
	if !(h.Peer == nil) {
		return false
	}
	if !(h.UserID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (h *MessagesHideChatJoinRequestRequest) String() string {
	if h == nil {
		return "MessagesHideChatJoinRequestRequest(nil)"
	}
	type Alias MessagesHideChatJoinRequestRequest
	return fmt.Sprintf("MessagesHideChatJoinRequestRequest%+v", Alias(*h))
}

// FillFrom fills MessagesHideChatJoinRequestRequest from given interface.
func (h *MessagesHideChatJoinRequestRequest) FillFrom(from interface {
	GetApproved() (value bool)
	GetPeer() (value InputPeerClass)
	GetUserID() (value InputUserClass)
}) {
	h.Approved = from.GetApproved()
	h.Peer = from.GetPeer()
	h.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesHideChatJoinRequestRequest) TypeID() uint32 {
	return MessagesHideChatJoinRequestRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesHideChatJoinRequestRequest) TypeName() string {
	return "messages.hideChatJoinRequest"
}

// TypeInfo returns info about TL type.
func (h *MessagesHideChatJoinRequestRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.hideChatJoinRequest",
		ID:   MessagesHideChatJoinRequestRequestTypeID,
	}
	if h == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Approved",
			SchemaName: "approved",
			Null:       !h.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (h *MessagesHideChatJoinRequestRequest) SetFlags() {
	if !(h.Approved == false) {
		h.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (h *MessagesHideChatJoinRequestRequest) Encode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.hideChatJoinRequest#7fe7e815 as nil")
	}
	b.PutID(MessagesHideChatJoinRequestRequestTypeID)
	return h.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (h *MessagesHideChatJoinRequestRequest) EncodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't encode messages.hideChatJoinRequest#7fe7e815 as nil")
	}
	h.SetFlags()
	if err := h.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.hideChatJoinRequest#7fe7e815: field flags: %w", err)
	}
	if h.Peer == nil {
		return fmt.Errorf("unable to encode messages.hideChatJoinRequest#7fe7e815: field peer is nil")
	}
	if err := h.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.hideChatJoinRequest#7fe7e815: field peer: %w", err)
	}
	if h.UserID == nil {
		return fmt.Errorf("unable to encode messages.hideChatJoinRequest#7fe7e815: field user_id is nil")
	}
	if err := h.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.hideChatJoinRequest#7fe7e815: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (h *MessagesHideChatJoinRequestRequest) Decode(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.hideChatJoinRequest#7fe7e815 to nil")
	}
	if err := b.ConsumeID(MessagesHideChatJoinRequestRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.hideChatJoinRequest#7fe7e815: %w", err)
	}
	return h.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (h *MessagesHideChatJoinRequestRequest) DecodeBare(b *bin.Buffer) error {
	if h == nil {
		return fmt.Errorf("can't decode messages.hideChatJoinRequest#7fe7e815 to nil")
	}
	{
		if err := h.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.hideChatJoinRequest#7fe7e815: field flags: %w", err)
		}
	}
	h.Approved = h.Flags.Has(0)
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.hideChatJoinRequest#7fe7e815: field peer: %w", err)
		}
		h.Peer = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.hideChatJoinRequest#7fe7e815: field user_id: %w", err)
		}
		h.UserID = value
	}
	return nil
}

// SetApproved sets value of Approved conditional field.
func (h *MessagesHideChatJoinRequestRequest) SetApproved(value bool) {
	if value {
		h.Flags.Set(0)
		h.Approved = true
	} else {
		h.Flags.Unset(0)
		h.Approved = false
	}
}

// GetApproved returns value of Approved conditional field.
func (h *MessagesHideChatJoinRequestRequest) GetApproved() (value bool) {
	if h == nil {
		return
	}
	return h.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (h *MessagesHideChatJoinRequestRequest) GetPeer() (value InputPeerClass) {
	if h == nil {
		return
	}
	return h.Peer
}

// GetUserID returns value of UserID field.
func (h *MessagesHideChatJoinRequestRequest) GetUserID() (value InputUserClass) {
	if h == nil {
		return
	}
	return h.UserID
}

// MessagesHideChatJoinRequest invokes method messages.hideChatJoinRequest#7fe7e815 returning error if any.
// Dismiss or approve a chat join request¹ related to a specific chat or channel.
//
// Links:
//  1. https://core.telegram.org/api/invites#join-requests
//
// Possible errors:
//
//	400 CHANNELS_TOO_MUCH: You have joined too many channels/supergroups.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	403 CHAT_WRITE_FORBIDDEN: You can't write in this chat.
//	400 HIDE_REQUESTER_MISSING: The join request was missing or was already handled.
//	400 INPUT_USER_DEACTIVATED: The specified user was deleted.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 USER_ALREADY_PARTICIPANT: The user is already in the group.
//	403 USER_CHANNELS_TOO_MUCH: One of the users you tried to add is already in too many channels/supergroups.
//	400 USER_ID_INVALID: The provided user ID is invalid.
//
// See https://core.telegram.org/method/messages.hideChatJoinRequest for reference.
// Can be used by bots.
func (c *Client) MessagesHideChatJoinRequest(ctx context.Context, request *MessagesHideChatJoinRequestRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
