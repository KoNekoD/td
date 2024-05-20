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

// MessagesSetChatAvailableReactionsRequest represents TL type `messages.setChatAvailableReactions#5a150bd4`.
// Change the set of message reactions »¹ that can be used in a certain group,
// supergroup or channel
//
// Links:
//  1. https://core.telegram.org/api/reactions
//
// See https://core.telegram.org/method/messages.setChatAvailableReactions for reference.
type MessagesSetChatAvailableReactionsRequest struct {
	// Flags field of MessagesSetChatAvailableReactionsRequest.
	Flags bin.Fields
	// Group where to apply changes
	Peer InputPeerClass
	// Allowed reaction emojis
	AvailableReactions ChatReactionsClass
	// ReactionsLimit field of MessagesSetChatAvailableReactionsRequest.
	//
	// Use SetReactionsLimit and GetReactionsLimit helpers.
	ReactionsLimit int
}

// MessagesSetChatAvailableReactionsRequestTypeID is TL type id of MessagesSetChatAvailableReactionsRequest.
const MessagesSetChatAvailableReactionsRequestTypeID = 0x5a150bd4

// Ensuring interfaces in compile-time for MessagesSetChatAvailableReactionsRequest.
var (
	_ bin.Encoder     = &MessagesSetChatAvailableReactionsRequest{}
	_ bin.Decoder     = &MessagesSetChatAvailableReactionsRequest{}
	_ bin.BareEncoder = &MessagesSetChatAvailableReactionsRequest{}
	_ bin.BareDecoder = &MessagesSetChatAvailableReactionsRequest{}
)

func (s *MessagesSetChatAvailableReactionsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.AvailableReactions == nil) {
		return false
	}
	if !(s.ReactionsLimit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSetChatAvailableReactionsRequest) String() string {
	if s == nil {
		return "MessagesSetChatAvailableReactionsRequest(nil)"
	}
	type Alias MessagesSetChatAvailableReactionsRequest
	return fmt.Sprintf("MessagesSetChatAvailableReactionsRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSetChatAvailableReactionsRequest from given interface.
func (s *MessagesSetChatAvailableReactionsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetAvailableReactions() (value ChatReactionsClass)
	GetReactionsLimit() (value int, ok bool)
}) {
	s.Peer = from.GetPeer()
	s.AvailableReactions = from.GetAvailableReactions()
	if val, ok := from.GetReactionsLimit(); ok {
		s.ReactionsLimit = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSetChatAvailableReactionsRequest) TypeID() uint32 {
	return MessagesSetChatAvailableReactionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSetChatAvailableReactionsRequest) TypeName() string {
	return "messages.setChatAvailableReactions"
}

// TypeInfo returns info about TL type.
func (s *MessagesSetChatAvailableReactionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.setChatAvailableReactions",
		ID:   MessagesSetChatAvailableReactionsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "AvailableReactions",
			SchemaName: "available_reactions",
		},
		{
			Name:       "ReactionsLimit",
			SchemaName: "reactions_limit",
			Null:       !s.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSetChatAvailableReactionsRequest) SetFlags() {
	if !(s.ReactionsLimit == 0) {
		s.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSetChatAvailableReactionsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setChatAvailableReactions#5a150bd4 as nil")
	}
	b.PutID(MessagesSetChatAvailableReactionsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSetChatAvailableReactionsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.setChatAvailableReactions#5a150bd4 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#5a150bd4: field flags: %w", err)
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#5a150bd4: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#5a150bd4: field peer: %w", err)
	}
	if s.AvailableReactions == nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#5a150bd4: field available_reactions is nil")
	}
	if err := s.AvailableReactions.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.setChatAvailableReactions#5a150bd4: field available_reactions: %w", err)
	}
	if s.Flags.Has(0) {
		b.PutInt(s.ReactionsLimit)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSetChatAvailableReactionsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setChatAvailableReactions#5a150bd4 to nil")
	}
	if err := b.ConsumeID(MessagesSetChatAvailableReactionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.setChatAvailableReactions#5a150bd4: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSetChatAvailableReactionsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.setChatAvailableReactions#5a150bd4 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.setChatAvailableReactions#5a150bd4: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setChatAvailableReactions#5a150bd4: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := DecodeChatReactions(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.setChatAvailableReactions#5a150bd4: field available_reactions: %w", err)
		}
		s.AvailableReactions = value
	}
	if s.Flags.Has(0) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.setChatAvailableReactions#5a150bd4: field reactions_limit: %w", err)
		}
		s.ReactionsLimit = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSetChatAvailableReactionsRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetAvailableReactions returns value of AvailableReactions field.
func (s *MessagesSetChatAvailableReactionsRequest) GetAvailableReactions() (value ChatReactionsClass) {
	if s == nil {
		return
	}
	return s.AvailableReactions
}

// SetReactionsLimit sets value of ReactionsLimit conditional field.
func (s *MessagesSetChatAvailableReactionsRequest) SetReactionsLimit(value int) {
	s.Flags.Set(0)
	s.ReactionsLimit = value
}

// GetReactionsLimit returns value of ReactionsLimit conditional field and
// boolean which is true if field was set.
func (s *MessagesSetChatAvailableReactionsRequest) GetReactionsLimit() (value int, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.ReactionsLimit, true
}

// MessagesSetChatAvailableReactions invokes method messages.setChatAvailableReactions#5a150bd4 returning error if any.
// Change the set of message reactions »¹ that can be used in a certain group,
// supergroup or channel
//
// Links:
//  1. https://core.telegram.org/api/reactions
//
// Possible errors:
//
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//	400 CHAT_NOT_MODIFIED: No changes were made to chat information because the new information you passed is identical to the current information.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.setChatAvailableReactions for reference.
func (c *Client) MessagesSetChatAvailableReactions(ctx context.Context, request *MessagesSetChatAvailableReactionsRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
