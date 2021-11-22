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

// ContactsToggleTopPeersRequest represents TL type `contacts.toggleTopPeers#8514bdda`.
// Enable/disable top peers¹
//
// Links:
//  1) https://core.telegram.org/api/top-rating
//
// See https://core.telegram.org/method/contacts.toggleTopPeers for reference.
type ContactsToggleTopPeersRequest struct {
	// Enable/disable
	Enabled bool
}

// ContactsToggleTopPeersRequestTypeID is TL type id of ContactsToggleTopPeersRequest.
const ContactsToggleTopPeersRequestTypeID = 0x8514bdda

// Ensuring interfaces in compile-time for ContactsToggleTopPeersRequest.
var (
	_ bin.Encoder     = &ContactsToggleTopPeersRequest{}
	_ bin.Decoder     = &ContactsToggleTopPeersRequest{}
	_ bin.BareEncoder = &ContactsToggleTopPeersRequest{}
	_ bin.BareDecoder = &ContactsToggleTopPeersRequest{}
)

func (t *ContactsToggleTopPeersRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Enabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ContactsToggleTopPeersRequest) String() string {
	if t == nil {
		return "ContactsToggleTopPeersRequest(nil)"
	}
	type Alias ContactsToggleTopPeersRequest
	return fmt.Sprintf("ContactsToggleTopPeersRequest%+v", Alias(*t))
}

// FillFrom fills ContactsToggleTopPeersRequest from given interface.
func (t *ContactsToggleTopPeersRequest) FillFrom(from interface {
	GetEnabled() (value bool)
}) {
	t.Enabled = from.GetEnabled()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsToggleTopPeersRequest) TypeID() uint32 {
	return ContactsToggleTopPeersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsToggleTopPeersRequest) TypeName() string {
	return "contacts.toggleTopPeers"
}

// TypeInfo returns info about TL type.
func (t *ContactsToggleTopPeersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.toggleTopPeers",
		ID:   ContactsToggleTopPeersRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Enabled",
			SchemaName: "enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ContactsToggleTopPeersRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.toggleTopPeers#8514bdda as nil")
	}
	b.PutID(ContactsToggleTopPeersRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ContactsToggleTopPeersRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode contacts.toggleTopPeers#8514bdda as nil")
	}
	b.PutBool(t.Enabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ContactsToggleTopPeersRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.toggleTopPeers#8514bdda to nil")
	}
	if err := b.ConsumeID(ContactsToggleTopPeersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.toggleTopPeers#8514bdda: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ContactsToggleTopPeersRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode contacts.toggleTopPeers#8514bdda to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.toggleTopPeers#8514bdda: field enabled: %w", err)
		}
		t.Enabled = value
	}
	return nil
}

// GetEnabled returns value of Enabled field.
func (t *ContactsToggleTopPeersRequest) GetEnabled() (value bool) {
	return t.Enabled
}

// ContactsToggleTopPeers invokes method contacts.toggleTopPeers#8514bdda returning error if any.
// Enable/disable top peers¹
//
// Links:
//  1) https://core.telegram.org/api/top-rating
//
// See https://core.telegram.org/method/contacts.toggleTopPeers for reference.
func (c *Client) ContactsToggleTopPeers(ctx context.Context, enabled bool) (bool, error) {
	var result BoolBox

	request := &ContactsToggleTopPeersRequest{
		Enabled: enabled,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
