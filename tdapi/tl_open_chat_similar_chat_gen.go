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

// OpenChatSimilarChatRequest represents TL type `openChatSimilarChat#8fa6f413`.
type OpenChatSimilarChatRequest struct {
	// Identifier of the original chat, which similar chats were requested
	ChatID int64
	// Identifier of the opened chat
	OpenedChatID int64
}

// OpenChatSimilarChatRequestTypeID is TL type id of OpenChatSimilarChatRequest.
const OpenChatSimilarChatRequestTypeID = 0x8fa6f413

// Ensuring interfaces in compile-time for OpenChatSimilarChatRequest.
var (
	_ bin.Encoder     = &OpenChatSimilarChatRequest{}
	_ bin.Decoder     = &OpenChatSimilarChatRequest{}
	_ bin.BareEncoder = &OpenChatSimilarChatRequest{}
	_ bin.BareDecoder = &OpenChatSimilarChatRequest{}
)

func (o *OpenChatSimilarChatRequest) Zero() bool {
	if o == nil {
		return true
	}
	if !(o.ChatID == 0) {
		return false
	}
	if !(o.OpenedChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (o *OpenChatSimilarChatRequest) String() string {
	if o == nil {
		return "OpenChatSimilarChatRequest(nil)"
	}
	type Alias OpenChatSimilarChatRequest
	return fmt.Sprintf("OpenChatSimilarChatRequest%+v", Alias(*o))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*OpenChatSimilarChatRequest) TypeID() uint32 {
	return OpenChatSimilarChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*OpenChatSimilarChatRequest) TypeName() string {
	return "openChatSimilarChat"
}

// TypeInfo returns info about TL type.
func (o *OpenChatSimilarChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "openChatSimilarChat",
		ID:   OpenChatSimilarChatRequestTypeID,
	}
	if o == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "OpenedChatID",
			SchemaName: "opened_chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (o *OpenChatSimilarChatRequest) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode openChatSimilarChat#8fa6f413 as nil")
	}
	b.PutID(OpenChatSimilarChatRequestTypeID)
	return o.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (o *OpenChatSimilarChatRequest) EncodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode openChatSimilarChat#8fa6f413 as nil")
	}
	b.PutInt53(o.ChatID)
	b.PutInt53(o.OpenedChatID)
	return nil
}

// Decode implements bin.Decoder.
func (o *OpenChatSimilarChatRequest) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode openChatSimilarChat#8fa6f413 to nil")
	}
	if err := b.ConsumeID(OpenChatSimilarChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode openChatSimilarChat#8fa6f413: %w", err)
	}
	return o.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (o *OpenChatSimilarChatRequest) DecodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode openChatSimilarChat#8fa6f413 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode openChatSimilarChat#8fa6f413: field chat_id: %w", err)
		}
		o.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode openChatSimilarChat#8fa6f413: field opened_chat_id: %w", err)
		}
		o.OpenedChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (o *OpenChatSimilarChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if o == nil {
		return fmt.Errorf("can't encode openChatSimilarChat#8fa6f413 as nil")
	}
	b.ObjStart()
	b.PutID("openChatSimilarChat")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(o.ChatID)
	b.Comma()
	b.FieldStart("opened_chat_id")
	b.PutInt53(o.OpenedChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (o *OpenChatSimilarChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if o == nil {
		return fmt.Errorf("can't decode openChatSimilarChat#8fa6f413 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("openChatSimilarChat"); err != nil {
				return fmt.Errorf("unable to decode openChatSimilarChat#8fa6f413: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode openChatSimilarChat#8fa6f413: field chat_id: %w", err)
			}
			o.ChatID = value
		case "opened_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode openChatSimilarChat#8fa6f413: field opened_chat_id: %w", err)
			}
			o.OpenedChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (o *OpenChatSimilarChatRequest) GetChatID() (value int64) {
	if o == nil {
		return
	}
	return o.ChatID
}

// GetOpenedChatID returns value of OpenedChatID field.
func (o *OpenChatSimilarChatRequest) GetOpenedChatID() (value int64) {
	if o == nil {
		return
	}
	return o.OpenedChatID
}

// OpenChatSimilarChat invokes method openChatSimilarChat#8fa6f413 returning error if any.
func (c *Client) OpenChatSimilarChat(ctx context.Context, request *OpenChatSimilarChatRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
