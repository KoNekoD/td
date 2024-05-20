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

// ToggleForumTopicIsClosedRequest represents TL type `toggleForumTopicIsClosed#c7648af3`.
type ToggleForumTopicIsClosedRequest struct {
	// Identifier of the chat
	ChatID int64
	// Message thread identifier of the forum topic
	MessageThreadID int64
	// Pass true to close the topic; pass false to reopen it
	IsClosed bool
}

// ToggleForumTopicIsClosedRequestTypeID is TL type id of ToggleForumTopicIsClosedRequest.
const ToggleForumTopicIsClosedRequestTypeID = 0xc7648af3

// Ensuring interfaces in compile-time for ToggleForumTopicIsClosedRequest.
var (
	_ bin.Encoder     = &ToggleForumTopicIsClosedRequest{}
	_ bin.Decoder     = &ToggleForumTopicIsClosedRequest{}
	_ bin.BareEncoder = &ToggleForumTopicIsClosedRequest{}
	_ bin.BareDecoder = &ToggleForumTopicIsClosedRequest{}
)

func (t *ToggleForumTopicIsClosedRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ChatID == 0) {
		return false
	}
	if !(t.MessageThreadID == 0) {
		return false
	}
	if !(t.IsClosed == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleForumTopicIsClosedRequest) String() string {
	if t == nil {
		return "ToggleForumTopicIsClosedRequest(nil)"
	}
	type Alias ToggleForumTopicIsClosedRequest
	return fmt.Sprintf("ToggleForumTopicIsClosedRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleForumTopicIsClosedRequest) TypeID() uint32 {
	return ToggleForumTopicIsClosedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleForumTopicIsClosedRequest) TypeName() string {
	return "toggleForumTopicIsClosed"
}

// TypeInfo returns info about TL type.
func (t *ToggleForumTopicIsClosedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleForumTopicIsClosed",
		ID:   ToggleForumTopicIsClosedRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "IsClosed",
			SchemaName: "is_closed",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleForumTopicIsClosedRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleForumTopicIsClosed#c7648af3 as nil")
	}
	b.PutID(ToggleForumTopicIsClosedRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleForumTopicIsClosedRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleForumTopicIsClosed#c7648af3 as nil")
	}
	b.PutInt53(t.ChatID)
	b.PutInt53(t.MessageThreadID)
	b.PutBool(t.IsClosed)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleForumTopicIsClosedRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleForumTopicIsClosed#c7648af3 to nil")
	}
	if err := b.ConsumeID(ToggleForumTopicIsClosedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleForumTopicIsClosed#c7648af3: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleForumTopicIsClosedRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleForumTopicIsClosed#c7648af3 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleForumTopicIsClosed#c7648af3: field chat_id: %w", err)
		}
		t.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleForumTopicIsClosed#c7648af3: field message_thread_id: %w", err)
		}
		t.MessageThreadID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleForumTopicIsClosed#c7648af3: field is_closed: %w", err)
		}
		t.IsClosed = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleForumTopicIsClosedRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleForumTopicIsClosed#c7648af3 as nil")
	}
	b.ObjStart()
	b.PutID("toggleForumTopicIsClosed")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(t.ChatID)
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(t.MessageThreadID)
	b.Comma()
	b.FieldStart("is_closed")
	b.PutBool(t.IsClosed)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleForumTopicIsClosedRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleForumTopicIsClosed#c7648af3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleForumTopicIsClosed"); err != nil {
				return fmt.Errorf("unable to decode toggleForumTopicIsClosed#c7648af3: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleForumTopicIsClosed#c7648af3: field chat_id: %w", err)
			}
			t.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleForumTopicIsClosed#c7648af3: field message_thread_id: %w", err)
			}
			t.MessageThreadID = value
		case "is_closed":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleForumTopicIsClosed#c7648af3: field is_closed: %w", err)
			}
			t.IsClosed = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (t *ToggleForumTopicIsClosedRequest) GetChatID() (value int64) {
	if t == nil {
		return
	}
	return t.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (t *ToggleForumTopicIsClosedRequest) GetMessageThreadID() (value int64) {
	if t == nil {
		return
	}
	return t.MessageThreadID
}

// GetIsClosed returns value of IsClosed field.
func (t *ToggleForumTopicIsClosedRequest) GetIsClosed() (value bool) {
	if t == nil {
		return
	}
	return t.IsClosed
}

// ToggleForumTopicIsClosed invokes method toggleForumTopicIsClosed#c7648af3 returning error if any.
func (c *Client) ToggleForumTopicIsClosed(ctx context.Context, request *ToggleForumTopicIsClosedRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
