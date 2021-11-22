// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesMarkDialogUnreadRequest represents TL type `messages.markDialogUnread#c286d98f`.
// Manually mark dialog as unread
//
// See https://core.telegram.org/method/messages.markDialogUnread for reference.
type MessagesMarkDialogUnreadRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Mark as unread/read
	Unread bool
	// Dialog
	Peer InputDialogPeerClass
}

// MessagesMarkDialogUnreadRequestTypeID is TL type id of MessagesMarkDialogUnreadRequest.
const MessagesMarkDialogUnreadRequestTypeID = 0xc286d98f

// Ensuring interfaces in compile-time for MessagesMarkDialogUnreadRequest.
var (
	_ bin.Encoder     = &MessagesMarkDialogUnreadRequest{}
	_ bin.Decoder     = &MessagesMarkDialogUnreadRequest{}
	_ bin.BareEncoder = &MessagesMarkDialogUnreadRequest{}
	_ bin.BareDecoder = &MessagesMarkDialogUnreadRequest{}
)

func (m *MessagesMarkDialogUnreadRequest) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Unread == false) {
		return false
	}
	if !(m.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMarkDialogUnreadRequest) String() string {
	if m == nil {
		return "MessagesMarkDialogUnreadRequest(nil)"
	}
	type Alias MessagesMarkDialogUnreadRequest
	return fmt.Sprintf("MessagesMarkDialogUnreadRequest%+v", Alias(*m))
}

// FillFrom fills MessagesMarkDialogUnreadRequest from given interface.
func (m *MessagesMarkDialogUnreadRequest) FillFrom(from interface {
	GetUnread() (value bool)
	GetPeer() (value InputDialogPeerClass)
}) {
	m.Unread = from.GetUnread()
	m.Peer = from.GetPeer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesMarkDialogUnreadRequest) TypeID() uint32 {
	return MessagesMarkDialogUnreadRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesMarkDialogUnreadRequest) TypeName() string {
	return "messages.markDialogUnread"
}

// TypeInfo returns info about TL type.
func (m *MessagesMarkDialogUnreadRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.markDialogUnread",
		ID:   MessagesMarkDialogUnreadRequestTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Unread",
			SchemaName: "unread",
			Null:       !m.Flags.Has(0),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagesMarkDialogUnreadRequest) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.markDialogUnread#c286d98f as nil")
	}
	b.PutID(MessagesMarkDialogUnreadRequestTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagesMarkDialogUnreadRequest) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.markDialogUnread#c286d98f as nil")
	}
	if !(m.Unread == false) {
		m.Flags.Set(0)
	}
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.markDialogUnread#c286d98f: field flags: %w", err)
	}
	if m.Peer == nil {
		return fmt.Errorf("unable to encode messages.markDialogUnread#c286d98f: field peer is nil")
	}
	if err := m.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.markDialogUnread#c286d98f: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagesMarkDialogUnreadRequest) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.markDialogUnread#c286d98f to nil")
	}
	if err := b.ConsumeID(MessagesMarkDialogUnreadRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.markDialogUnread#c286d98f: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagesMarkDialogUnreadRequest) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.markDialogUnread#c286d98f to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.markDialogUnread#c286d98f: field flags: %w", err)
		}
	}
	m.Unread = m.Flags.Has(0)
	{
		value, err := DecodeInputDialogPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.markDialogUnread#c286d98f: field peer: %w", err)
		}
		m.Peer = value
	}
	return nil
}

// SetUnread sets value of Unread conditional field.
func (m *MessagesMarkDialogUnreadRequest) SetUnread(value bool) {
	if value {
		m.Flags.Set(0)
		m.Unread = true
	} else {
		m.Flags.Unset(0)
		m.Unread = false
	}
}

// GetUnread returns value of Unread conditional field.
func (m *MessagesMarkDialogUnreadRequest) GetUnread() (value bool) {
	return m.Flags.Has(0)
}

// GetPeer returns value of Peer field.
func (m *MessagesMarkDialogUnreadRequest) GetPeer() (value InputDialogPeerClass) {
	return m.Peer
}

// MessagesMarkDialogUnread invokes method messages.markDialogUnread#c286d98f returning error if any.
// Manually mark dialog as unread
//
// See https://core.telegram.org/method/messages.markDialogUnread for reference.
func (c *Client) MessagesMarkDialogUnread(ctx context.Context, request *MessagesMarkDialogUnreadRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
