// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// JoinChatRequest represents TL type `joinChat#137a1aa1`.
type JoinChatRequest struct {
	// Chat identifier
	ChatID int64
}

// JoinChatRequestTypeID is TL type id of JoinChatRequest.
const JoinChatRequestTypeID = 0x137a1aa1

// Ensuring interfaces in compile-time for JoinChatRequest.
var (
	_ bin.Encoder     = &JoinChatRequest{}
	_ bin.Decoder     = &JoinChatRequest{}
	_ bin.BareEncoder = &JoinChatRequest{}
	_ bin.BareDecoder = &JoinChatRequest{}
)

func (j *JoinChatRequest) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JoinChatRequest) String() string {
	if j == nil {
		return "JoinChatRequest(nil)"
	}
	type Alias JoinChatRequest
	return fmt.Sprintf("JoinChatRequest%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JoinChatRequest) TypeID() uint32 {
	return JoinChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*JoinChatRequest) TypeName() string {
	return "joinChat"
}

// TypeInfo returns info about TL type.
func (j *JoinChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "joinChat",
		ID:   JoinChatRequestTypeID,
	}
	if j == nil {
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
func (j *JoinChatRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode joinChat#137a1aa1 as nil")
	}
	b.PutID(JoinChatRequestTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JoinChatRequest) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode joinChat#137a1aa1 as nil")
	}
	b.PutLong(j.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (j *JoinChatRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode joinChat#137a1aa1 to nil")
	}
	if err := b.ConsumeID(JoinChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode joinChat#137a1aa1: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JoinChatRequest) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode joinChat#137a1aa1 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode joinChat#137a1aa1: field chat_id: %w", err)
		}
		j.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON encodes j in TDLib API JSON format.
func (j *JoinChatRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if j == nil {
		return fmt.Errorf("can't encode joinChat#137a1aa1 as nil")
	}
	b.ObjStart()
	b.PutID("joinChat")
	b.FieldStart("chat_id")
	b.PutLong(j.ChatID)
	b.ObjEnd()
	return nil
}

// GetChatID returns value of ChatID field.
func (j *JoinChatRequest) GetChatID() (value int64) {
	return j.ChatID
}

// JoinChat invokes method joinChat#137a1aa1 returning error if any.
func (c *Client) JoinChat(ctx context.Context, chatid int64) error {
	var ok Ok

	request := &JoinChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
