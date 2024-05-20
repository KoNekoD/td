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

// MessagesDeleteQuickReplyShortcutRequest represents TL type `messages.deleteQuickReplyShortcut#3cc04740`.
//
// See https://core.telegram.org/method/messages.deleteQuickReplyShortcut for reference.
type MessagesDeleteQuickReplyShortcutRequest struct {
	// ShortcutID field of MessagesDeleteQuickReplyShortcutRequest.
	ShortcutID int
}

// MessagesDeleteQuickReplyShortcutRequestTypeID is TL type id of MessagesDeleteQuickReplyShortcutRequest.
const MessagesDeleteQuickReplyShortcutRequestTypeID = 0x3cc04740

// Ensuring interfaces in compile-time for MessagesDeleteQuickReplyShortcutRequest.
var (
	_ bin.Encoder     = &MessagesDeleteQuickReplyShortcutRequest{}
	_ bin.Decoder     = &MessagesDeleteQuickReplyShortcutRequest{}
	_ bin.BareEncoder = &MessagesDeleteQuickReplyShortcutRequest{}
	_ bin.BareDecoder = &MessagesDeleteQuickReplyShortcutRequest{}
)

func (d *MessagesDeleteQuickReplyShortcutRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ShortcutID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteQuickReplyShortcutRequest) String() string {
	if d == nil {
		return "MessagesDeleteQuickReplyShortcutRequest(nil)"
	}
	type Alias MessagesDeleteQuickReplyShortcutRequest
	return fmt.Sprintf("MessagesDeleteQuickReplyShortcutRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeleteQuickReplyShortcutRequest from given interface.
func (d *MessagesDeleteQuickReplyShortcutRequest) FillFrom(from interface {
	GetShortcutID() (value int)
}) {
	d.ShortcutID = from.GetShortcutID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDeleteQuickReplyShortcutRequest) TypeID() uint32 {
	return MessagesDeleteQuickReplyShortcutRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDeleteQuickReplyShortcutRequest) TypeName() string {
	return "messages.deleteQuickReplyShortcut"
}

// TypeInfo returns info about TL type.
func (d *MessagesDeleteQuickReplyShortcutRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.deleteQuickReplyShortcut",
		ID:   MessagesDeleteQuickReplyShortcutRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShortcutID",
			SchemaName: "shortcut_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteQuickReplyShortcutRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteQuickReplyShortcut#3cc04740 as nil")
	}
	b.PutID(MessagesDeleteQuickReplyShortcutRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDeleteQuickReplyShortcutRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteQuickReplyShortcut#3cc04740 as nil")
	}
	b.PutInt(d.ShortcutID)
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteQuickReplyShortcutRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteQuickReplyShortcut#3cc04740 to nil")
	}
	if err := b.ConsumeID(MessagesDeleteQuickReplyShortcutRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteQuickReplyShortcut#3cc04740: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDeleteQuickReplyShortcutRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteQuickReplyShortcut#3cc04740 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteQuickReplyShortcut#3cc04740: field shortcut_id: %w", err)
		}
		d.ShortcutID = value
	}
	return nil
}

// GetShortcutID returns value of ShortcutID field.
func (d *MessagesDeleteQuickReplyShortcutRequest) GetShortcutID() (value int) {
	if d == nil {
		return
	}
	return d.ShortcutID
}

// MessagesDeleteQuickReplyShortcut invokes method messages.deleteQuickReplyShortcut#3cc04740 returning error if any.
//
// See https://core.telegram.org/method/messages.deleteQuickReplyShortcut for reference.
func (c *Client) MessagesDeleteQuickReplyShortcut(ctx context.Context, shortcutid int) (bool, error) {
	var result BoolBox

	request := &MessagesDeleteQuickReplyShortcutRequest{
		ShortcutID: shortcutid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
