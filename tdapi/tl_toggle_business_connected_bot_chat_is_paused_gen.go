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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// ToggleBusinessConnectedBotChatIsPausedRequest represents TL type `toggleBusinessConnectedBotChatIsPaused#4f364845`.
type ToggleBusinessConnectedBotChatIsPausedRequest struct {
	// Chat identifier
	ChatID int64
	// Pass true to pause the connected bot in the chat; pass false to resume the bot
	IsPaused bool
}

// ToggleBusinessConnectedBotChatIsPausedRequestTypeID is TL type id of ToggleBusinessConnectedBotChatIsPausedRequest.
const ToggleBusinessConnectedBotChatIsPausedRequestTypeID = 0x4f364845

// Ensuring interfaces in compile-time for ToggleBusinessConnectedBotChatIsPausedRequest.
var (
	_ bin.Encoder     = &ToggleBusinessConnectedBotChatIsPausedRequest{}
	_ bin.Decoder     = &ToggleBusinessConnectedBotChatIsPausedRequest{}
	_ bin.BareEncoder = &ToggleBusinessConnectedBotChatIsPausedRequest{}
	_ bin.BareDecoder = &ToggleBusinessConnectedBotChatIsPausedRequest{}
)

func (t *ToggleBusinessConnectedBotChatIsPausedRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.ChatID == 0) {
		return false
	}
	if !(t.IsPaused == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleBusinessConnectedBotChatIsPausedRequest) String() string {
	if t == nil {
		return "ToggleBusinessConnectedBotChatIsPausedRequest(nil)"
	}
	type Alias ToggleBusinessConnectedBotChatIsPausedRequest
	return fmt.Sprintf("ToggleBusinessConnectedBotChatIsPausedRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleBusinessConnectedBotChatIsPausedRequest) TypeID() uint32 {
	return ToggleBusinessConnectedBotChatIsPausedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleBusinessConnectedBotChatIsPausedRequest) TypeName() string {
	return "toggleBusinessConnectedBotChatIsPaused"
}

// TypeInfo returns info about TL type.
func (t *ToggleBusinessConnectedBotChatIsPausedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleBusinessConnectedBotChatIsPaused",
		ID:   ToggleBusinessConnectedBotChatIsPausedRequestTypeID,
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
			Name:       "IsPaused",
			SchemaName: "is_paused",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleBusinessConnectedBotChatIsPausedRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleBusinessConnectedBotChatIsPaused#4f364845 as nil")
	}
	b.PutID(ToggleBusinessConnectedBotChatIsPausedRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleBusinessConnectedBotChatIsPausedRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleBusinessConnectedBotChatIsPaused#4f364845 as nil")
	}
	b.PutInt53(t.ChatID)
	b.PutBool(t.IsPaused)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleBusinessConnectedBotChatIsPausedRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleBusinessConnectedBotChatIsPaused#4f364845 to nil")
	}
	if err := b.ConsumeID(ToggleBusinessConnectedBotChatIsPausedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleBusinessConnectedBotChatIsPaused#4f364845: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleBusinessConnectedBotChatIsPausedRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleBusinessConnectedBotChatIsPaused#4f364845 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleBusinessConnectedBotChatIsPaused#4f364845: field chat_id: %w", err)
		}
		t.ChatID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleBusinessConnectedBotChatIsPaused#4f364845: field is_paused: %w", err)
		}
		t.IsPaused = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleBusinessConnectedBotChatIsPausedRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleBusinessConnectedBotChatIsPaused#4f364845 as nil")
	}
	b.ObjStart()
	b.PutID("toggleBusinessConnectedBotChatIsPaused")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(t.ChatID)
	b.Comma()
	b.FieldStart("is_paused")
	b.PutBool(t.IsPaused)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleBusinessConnectedBotChatIsPausedRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleBusinessConnectedBotChatIsPaused#4f364845 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleBusinessConnectedBotChatIsPaused"); err != nil {
				return fmt.Errorf("unable to decode toggleBusinessConnectedBotChatIsPaused#4f364845: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleBusinessConnectedBotChatIsPaused#4f364845: field chat_id: %w", err)
			}
			t.ChatID = value
		case "is_paused":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleBusinessConnectedBotChatIsPaused#4f364845: field is_paused: %w", err)
			}
			t.IsPaused = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (t *ToggleBusinessConnectedBotChatIsPausedRequest) GetChatID() (value int64) {
	if t == nil {
		return
	}
	return t.ChatID
}

// GetIsPaused returns value of IsPaused field.
func (t *ToggleBusinessConnectedBotChatIsPausedRequest) GetIsPaused() (value bool) {
	if t == nil {
		return
	}
	return t.IsPaused
}

// ToggleBusinessConnectedBotChatIsPaused invokes method toggleBusinessConnectedBotChatIsPaused#4f364845 returning error if any.
func (c *Client) ToggleBusinessConnectedBotChatIsPaused(ctx context.Context, request *ToggleBusinessConnectedBotChatIsPausedRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}