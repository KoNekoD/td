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

// MessagesSaveDefaultSendAsRequest represents TL type `messages.saveDefaultSendAs#ccfddf96`.
// Change the default peer that should be used when sending messages, reactions, poll
// votes to a specific group
//
// See https://core.telegram.org/method/messages.saveDefaultSendAs for reference.
type MessagesSaveDefaultSendAsRequest struct {
	// Group
	Peer InputPeerClass
	// The default peer that should be used when sending messages to the group
	SendAs InputPeerClass
}

// MessagesSaveDefaultSendAsRequestTypeID is TL type id of MessagesSaveDefaultSendAsRequest.
const MessagesSaveDefaultSendAsRequestTypeID = 0xccfddf96

// Ensuring interfaces in compile-time for MessagesSaveDefaultSendAsRequest.
var (
	_ bin.Encoder     = &MessagesSaveDefaultSendAsRequest{}
	_ bin.Decoder     = &MessagesSaveDefaultSendAsRequest{}
	_ bin.BareEncoder = &MessagesSaveDefaultSendAsRequest{}
	_ bin.BareDecoder = &MessagesSaveDefaultSendAsRequest{}
)

func (s *MessagesSaveDefaultSendAsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.SendAs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSaveDefaultSendAsRequest) String() string {
	if s == nil {
		return "MessagesSaveDefaultSendAsRequest(nil)"
	}
	type Alias MessagesSaveDefaultSendAsRequest
	return fmt.Sprintf("MessagesSaveDefaultSendAsRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSaveDefaultSendAsRequest from given interface.
func (s *MessagesSaveDefaultSendAsRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetSendAs() (value InputPeerClass)
}) {
	s.Peer = from.GetPeer()
	s.SendAs = from.GetSendAs()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSaveDefaultSendAsRequest) TypeID() uint32 {
	return MessagesSaveDefaultSendAsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSaveDefaultSendAsRequest) TypeName() string {
	return "messages.saveDefaultSendAs"
}

// TypeInfo returns info about TL type.
func (s *MessagesSaveDefaultSendAsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.saveDefaultSendAs",
		ID:   MessagesSaveDefaultSendAsRequestTypeID,
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
			Name:       "SendAs",
			SchemaName: "send_as",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *MessagesSaveDefaultSendAsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.saveDefaultSendAs#ccfddf96 as nil")
	}
	b.PutID(MessagesSaveDefaultSendAsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSaveDefaultSendAsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.saveDefaultSendAs#ccfddf96 as nil")
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.saveDefaultSendAs#ccfddf96: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.saveDefaultSendAs#ccfddf96: field peer: %w", err)
	}
	if s.SendAs == nil {
		return fmt.Errorf("unable to encode messages.saveDefaultSendAs#ccfddf96: field send_as is nil")
	}
	if err := s.SendAs.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.saveDefaultSendAs#ccfddf96: field send_as: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSaveDefaultSendAsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.saveDefaultSendAs#ccfddf96 to nil")
	}
	if err := b.ConsumeID(MessagesSaveDefaultSendAsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.saveDefaultSendAs#ccfddf96: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSaveDefaultSendAsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.saveDefaultSendAs#ccfddf96 to nil")
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDefaultSendAs#ccfddf96: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDefaultSendAs#ccfddf96: field send_as: %w", err)
		}
		s.SendAs = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (s *MessagesSaveDefaultSendAsRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetSendAs returns value of SendAs field.
func (s *MessagesSaveDefaultSendAsRequest) GetSendAs() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.SendAs
}

// MessagesSaveDefaultSendAs invokes method messages.saveDefaultSendAs#ccfddf96 returning error if any.
// Change the default peer that should be used when sending messages, reactions, poll
// votes to a specific group
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//	400 SEND_AS_PEER_INVALID: You can't send messages as the specified peer.
//
// See https://core.telegram.org/method/messages.saveDefaultSendAs for reference.
func (c *Client) MessagesSaveDefaultSendAs(ctx context.Context, request *MessagesSaveDefaultSendAsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
