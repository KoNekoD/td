// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// DeleteChatHistoryRequest represents TL type `deleteChatHistory#a841d09f`.
type DeleteChatHistoryRequest struct {
	// Chat identifier
	ChatID int64
	// Pass true to remove the chat from all chat lists
	RemoveFromChatList bool
	// Pass true to delete chat history for all users
	Revoke bool
}

// DeleteChatHistoryRequestTypeID is TL type id of DeleteChatHistoryRequest.
const DeleteChatHistoryRequestTypeID = 0xa841d09f

// Ensuring interfaces in compile-time for DeleteChatHistoryRequest.
var (
	_ bin.Encoder     = &DeleteChatHistoryRequest{}
	_ bin.Decoder     = &DeleteChatHistoryRequest{}
	_ bin.BareEncoder = &DeleteChatHistoryRequest{}
	_ bin.BareDecoder = &DeleteChatHistoryRequest{}
)

func (d *DeleteChatHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}
	if !(d.RemoveFromChatList == false) {
		return false
	}
	if !(d.Revoke == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteChatHistoryRequest) String() string {
	if d == nil {
		return "DeleteChatHistoryRequest(nil)"
	}
	type Alias DeleteChatHistoryRequest
	return fmt.Sprintf("DeleteChatHistoryRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteChatHistoryRequest) TypeID() uint32 {
	return DeleteChatHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteChatHistoryRequest) TypeName() string {
	return "deleteChatHistory"
}

// TypeInfo returns info about TL type.
func (d *DeleteChatHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteChatHistory",
		ID:   DeleteChatHistoryRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "RemoveFromChatList",
			SchemaName: "remove_from_chat_list",
		},
		{
			Name:       "Revoke",
			SchemaName: "revoke",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteChatHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatHistory#a841d09f as nil")
	}
	b.PutID(DeleteChatHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteChatHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatHistory#a841d09f as nil")
	}
	b.PutInt53(d.ChatID)
	b.PutBool(d.RemoveFromChatList)
	b.PutBool(d.Revoke)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteChatHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatHistory#a841d09f to nil")
	}
	if err := b.ConsumeID(DeleteChatHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteChatHistory#a841d09f: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteChatHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatHistory#a841d09f to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChatHistory#a841d09f: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChatHistory#a841d09f: field remove_from_chat_list: %w", err)
		}
		d.RemoveFromChatList = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChatHistory#a841d09f: field revoke: %w", err)
		}
		d.Revoke = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteChatHistoryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatHistory#a841d09f as nil")
	}
	b.ObjStart()
	b.PutID("deleteChatHistory")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(d.ChatID)
	b.Comma()
	b.FieldStart("remove_from_chat_list")
	b.PutBool(d.RemoveFromChatList)
	b.Comma()
	b.FieldStart("revoke")
	b.PutBool(d.Revoke)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteChatHistoryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatHistory#a841d09f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteChatHistory"); err != nil {
				return fmt.Errorf("unable to decode deleteChatHistory#a841d09f: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteChatHistory#a841d09f: field chat_id: %w", err)
			}
			d.ChatID = value
		case "remove_from_chat_list":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode deleteChatHistory#a841d09f: field remove_from_chat_list: %w", err)
			}
			d.RemoveFromChatList = value
		case "revoke":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode deleteChatHistory#a841d09f: field revoke: %w", err)
			}
			d.Revoke = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (d *DeleteChatHistoryRequest) GetChatID() (value int64) {
	if d == nil {
		return
	}
	return d.ChatID
}

// GetRemoveFromChatList returns value of RemoveFromChatList field.
func (d *DeleteChatHistoryRequest) GetRemoveFromChatList() (value bool) {
	if d == nil {
		return
	}
	return d.RemoveFromChatList
}

// GetRevoke returns value of Revoke field.
func (d *DeleteChatHistoryRequest) GetRevoke() (value bool) {
	if d == nil {
		return
	}
	return d.Revoke
}

// DeleteChatHistory invokes method deleteChatHistory#a841d09f returning error if any.
func (c *Client) DeleteChatHistory(ctx context.Context, request *DeleteChatHistoryRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
