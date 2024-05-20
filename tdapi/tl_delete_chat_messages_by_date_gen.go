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

// DeleteChatMessagesByDateRequest represents TL type `deleteChatMessagesByDate#9e44e0bf`.
type DeleteChatMessagesByDateRequest struct {
	// Chat identifier
	ChatID int64
	// The minimum date of the messages to delete
	MinDate int32
	// The maximum date of the messages to delete
	MaxDate int32
	// Pass true to delete chat messages for all users; private chats only
	Revoke bool
}

// DeleteChatMessagesByDateRequestTypeID is TL type id of DeleteChatMessagesByDateRequest.
const DeleteChatMessagesByDateRequestTypeID = 0x9e44e0bf

// Ensuring interfaces in compile-time for DeleteChatMessagesByDateRequest.
var (
	_ bin.Encoder     = &DeleteChatMessagesByDateRequest{}
	_ bin.Decoder     = &DeleteChatMessagesByDateRequest{}
	_ bin.BareEncoder = &DeleteChatMessagesByDateRequest{}
	_ bin.BareDecoder = &DeleteChatMessagesByDateRequest{}
)

func (d *DeleteChatMessagesByDateRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ChatID == 0) {
		return false
	}
	if !(d.MinDate == 0) {
		return false
	}
	if !(d.MaxDate == 0) {
		return false
	}
	if !(d.Revoke == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DeleteChatMessagesByDateRequest) String() string {
	if d == nil {
		return "DeleteChatMessagesByDateRequest(nil)"
	}
	type Alias DeleteChatMessagesByDateRequest
	return fmt.Sprintf("DeleteChatMessagesByDateRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DeleteChatMessagesByDateRequest) TypeID() uint32 {
	return DeleteChatMessagesByDateRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DeleteChatMessagesByDateRequest) TypeName() string {
	return "deleteChatMessagesByDate"
}

// TypeInfo returns info about TL type.
func (d *DeleteChatMessagesByDateRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "deleteChatMessagesByDate",
		ID:   DeleteChatMessagesByDateRequestTypeID,
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
			Name:       "MinDate",
			SchemaName: "min_date",
		},
		{
			Name:       "MaxDate",
			SchemaName: "max_date",
		},
		{
			Name:       "Revoke",
			SchemaName: "revoke",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *DeleteChatMessagesByDateRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatMessagesByDate#9e44e0bf as nil")
	}
	b.PutID(DeleteChatMessagesByDateRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DeleteChatMessagesByDateRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatMessagesByDate#9e44e0bf as nil")
	}
	b.PutInt53(d.ChatID)
	b.PutInt32(d.MinDate)
	b.PutInt32(d.MaxDate)
	b.PutBool(d.Revoke)
	return nil
}

// Decode implements bin.Decoder.
func (d *DeleteChatMessagesByDateRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatMessagesByDate#9e44e0bf to nil")
	}
	if err := b.ConsumeID(DeleteChatMessagesByDateRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode deleteChatMessagesByDate#9e44e0bf: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DeleteChatMessagesByDateRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatMessagesByDate#9e44e0bf to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChatMessagesByDate#9e44e0bf: field chat_id: %w", err)
		}
		d.ChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChatMessagesByDate#9e44e0bf: field min_date: %w", err)
		}
		d.MinDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChatMessagesByDate#9e44e0bf: field max_date: %w", err)
		}
		d.MaxDate = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode deleteChatMessagesByDate#9e44e0bf: field revoke: %w", err)
		}
		d.Revoke = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DeleteChatMessagesByDateRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode deleteChatMessagesByDate#9e44e0bf as nil")
	}
	b.ObjStart()
	b.PutID("deleteChatMessagesByDate")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(d.ChatID)
	b.Comma()
	b.FieldStart("min_date")
	b.PutInt32(d.MinDate)
	b.Comma()
	b.FieldStart("max_date")
	b.PutInt32(d.MaxDate)
	b.Comma()
	b.FieldStart("revoke")
	b.PutBool(d.Revoke)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DeleteChatMessagesByDateRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode deleteChatMessagesByDate#9e44e0bf to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("deleteChatMessagesByDate"); err != nil {
				return fmt.Errorf("unable to decode deleteChatMessagesByDate#9e44e0bf: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode deleteChatMessagesByDate#9e44e0bf: field chat_id: %w", err)
			}
			d.ChatID = value
		case "min_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode deleteChatMessagesByDate#9e44e0bf: field min_date: %w", err)
			}
			d.MinDate = value
		case "max_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode deleteChatMessagesByDate#9e44e0bf: field max_date: %w", err)
			}
			d.MaxDate = value
		case "revoke":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode deleteChatMessagesByDate#9e44e0bf: field revoke: %w", err)
			}
			d.Revoke = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (d *DeleteChatMessagesByDateRequest) GetChatID() (value int64) {
	if d == nil {
		return
	}
	return d.ChatID
}

// GetMinDate returns value of MinDate field.
func (d *DeleteChatMessagesByDateRequest) GetMinDate() (value int32) {
	if d == nil {
		return
	}
	return d.MinDate
}

// GetMaxDate returns value of MaxDate field.
func (d *DeleteChatMessagesByDateRequest) GetMaxDate() (value int32) {
	if d == nil {
		return
	}
	return d.MaxDate
}

// GetRevoke returns value of Revoke field.
func (d *DeleteChatMessagesByDateRequest) GetRevoke() (value bool) {
	if d == nil {
		return
	}
	return d.Revoke
}

// DeleteChatMessagesByDate invokes method deleteChatMessagesByDate#9e44e0bf returning error if any.
func (c *Client) DeleteChatMessagesByDate(ctx context.Context, request *DeleteChatMessagesByDateRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
