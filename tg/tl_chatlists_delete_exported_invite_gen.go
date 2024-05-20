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

// ChatlistsDeleteExportedInviteRequest represents TL type `chatlists.deleteExportedInvite#719c5c5e`.
// Delete a previously created chat folder deep link »¹.
//
// Links:
//  1. https://core.telegram.org/api/links#chat-folder-links
//
// See https://core.telegram.org/method/chatlists.deleteExportedInvite for reference.
type ChatlistsDeleteExportedInviteRequest struct {
	// The related folder
	Chatlist InputChatlistDialogFilter
	// slug obtained from the chat folder deep link »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#chat-folder-links
	Slug string
}

// ChatlistsDeleteExportedInviteRequestTypeID is TL type id of ChatlistsDeleteExportedInviteRequest.
const ChatlistsDeleteExportedInviteRequestTypeID = 0x719c5c5e

// Ensuring interfaces in compile-time for ChatlistsDeleteExportedInviteRequest.
var (
	_ bin.Encoder     = &ChatlistsDeleteExportedInviteRequest{}
	_ bin.Decoder     = &ChatlistsDeleteExportedInviteRequest{}
	_ bin.BareEncoder = &ChatlistsDeleteExportedInviteRequest{}
	_ bin.BareDecoder = &ChatlistsDeleteExportedInviteRequest{}
)

func (d *ChatlistsDeleteExportedInviteRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Chatlist.Zero()) {
		return false
	}
	if !(d.Slug == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChatlistsDeleteExportedInviteRequest) String() string {
	if d == nil {
		return "ChatlistsDeleteExportedInviteRequest(nil)"
	}
	type Alias ChatlistsDeleteExportedInviteRequest
	return fmt.Sprintf("ChatlistsDeleteExportedInviteRequest%+v", Alias(*d))
}

// FillFrom fills ChatlistsDeleteExportedInviteRequest from given interface.
func (d *ChatlistsDeleteExportedInviteRequest) FillFrom(from interface {
	GetChatlist() (value InputChatlistDialogFilter)
	GetSlug() (value string)
}) {
	d.Chatlist = from.GetChatlist()
	d.Slug = from.GetSlug()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatlistsDeleteExportedInviteRequest) TypeID() uint32 {
	return ChatlistsDeleteExportedInviteRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatlistsDeleteExportedInviteRequest) TypeName() string {
	return "chatlists.deleteExportedInvite"
}

// TypeInfo returns info about TL type.
func (d *ChatlistsDeleteExportedInviteRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatlists.deleteExportedInvite",
		ID:   ChatlistsDeleteExportedInviteRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Chatlist",
			SchemaName: "chatlist",
		},
		{
			Name:       "Slug",
			SchemaName: "slug",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *ChatlistsDeleteExportedInviteRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode chatlists.deleteExportedInvite#719c5c5e as nil")
	}
	b.PutID(ChatlistsDeleteExportedInviteRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *ChatlistsDeleteExportedInviteRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode chatlists.deleteExportedInvite#719c5c5e as nil")
	}
	if err := d.Chatlist.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatlists.deleteExportedInvite#719c5c5e: field chatlist: %w", err)
	}
	b.PutString(d.Slug)
	return nil
}

// Decode implements bin.Decoder.
func (d *ChatlistsDeleteExportedInviteRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode chatlists.deleteExportedInvite#719c5c5e to nil")
	}
	if err := b.ConsumeID(ChatlistsDeleteExportedInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode chatlists.deleteExportedInvite#719c5c5e: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *ChatlistsDeleteExportedInviteRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode chatlists.deleteExportedInvite#719c5c5e to nil")
	}
	{
		if err := d.Chatlist.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatlists.deleteExportedInvite#719c5c5e: field chatlist: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatlists.deleteExportedInvite#719c5c5e: field slug: %w", err)
		}
		d.Slug = value
	}
	return nil
}

// GetChatlist returns value of Chatlist field.
func (d *ChatlistsDeleteExportedInviteRequest) GetChatlist() (value InputChatlistDialogFilter) {
	if d == nil {
		return
	}
	return d.Chatlist
}

// GetSlug returns value of Slug field.
func (d *ChatlistsDeleteExportedInviteRequest) GetSlug() (value string) {
	if d == nil {
		return
	}
	return d.Slug
}

// ChatlistsDeleteExportedInvite invokes method chatlists.deleteExportedInvite#719c5c5e returning error if any.
// Delete a previously created chat folder deep link »¹.
//
// Links:
//  1. https://core.telegram.org/api/links#chat-folder-links
//
// Possible errors:
//
//	400 FILTER_ID_INVALID: The specified filter ID is invalid.
//	400 FILTER_NOT_SUPPORTED: The specified filter cannot be used in this context.
//
// See https://core.telegram.org/method/chatlists.deleteExportedInvite for reference.
// Can be used by bots.
func (c *Client) ChatlistsDeleteExportedInvite(ctx context.Context, request *ChatlistsDeleteExportedInviteRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
