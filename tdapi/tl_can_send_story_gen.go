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

// CanSendStoryRequest represents TL type `canSendStory#b6e0216b`.
type CanSendStoryRequest struct {
	// Chat identifier
	ChatID int64
}

// CanSendStoryRequestTypeID is TL type id of CanSendStoryRequest.
const CanSendStoryRequestTypeID = 0xb6e0216b

// Ensuring interfaces in compile-time for CanSendStoryRequest.
var (
	_ bin.Encoder     = &CanSendStoryRequest{}
	_ bin.Decoder     = &CanSendStoryRequest{}
	_ bin.BareEncoder = &CanSendStoryRequest{}
	_ bin.BareDecoder = &CanSendStoryRequest{}
)

func (c *CanSendStoryRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CanSendStoryRequest) String() string {
	if c == nil {
		return "CanSendStoryRequest(nil)"
	}
	type Alias CanSendStoryRequest
	return fmt.Sprintf("CanSendStoryRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CanSendStoryRequest) TypeID() uint32 {
	return CanSendStoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CanSendStoryRequest) TypeName() string {
	return "canSendStory"
}

// TypeInfo returns info about TL type.
func (c *CanSendStoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "canSendStory",
		ID:   CanSendStoryRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CanSendStoryRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode canSendStory#b6e0216b as nil")
	}
	b.PutID(CanSendStoryRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CanSendStoryRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode canSendStory#b6e0216b as nil")
	}
	b.PutInt53(c.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (c *CanSendStoryRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode canSendStory#b6e0216b to nil")
	}
	if err := b.ConsumeID(CanSendStoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode canSendStory#b6e0216b: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CanSendStoryRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode canSendStory#b6e0216b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode canSendStory#b6e0216b: field chat_id: %w", err)
		}
		c.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CanSendStoryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode canSendStory#b6e0216b as nil")
	}
	b.ObjStart()
	b.PutID("canSendStory")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(c.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CanSendStoryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode canSendStory#b6e0216b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("canSendStory"); err != nil {
				return fmt.Errorf("unable to decode canSendStory#b6e0216b: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode canSendStory#b6e0216b: field chat_id: %w", err)
			}
			c.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (c *CanSendStoryRequest) GetChatID() (value int64) {
	if c == nil {
		return
	}
	return c.ChatID
}

// CanSendStory invokes method canSendStory#b6e0216b returning error if any.
func (c *Client) CanSendStory(ctx context.Context, chatid int64) (CanSendStoryResultClass, error) {
	var result CanSendStoryResultBox

	request := &CanSendStoryRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.CanSendStoryResult, nil
}
