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

// UpgradeBasicGroupChatToSupergroupChatRequest represents TL type `upgradeBasicGroupChatToSupergroupChat#11e915ba`.
type UpgradeBasicGroupChatToSupergroupChatRequest struct {
	// Identifier of the chat to upgrade
	ChatID int64
}

// UpgradeBasicGroupChatToSupergroupChatRequestTypeID is TL type id of UpgradeBasicGroupChatToSupergroupChatRequest.
const UpgradeBasicGroupChatToSupergroupChatRequestTypeID = 0x11e915ba

// Ensuring interfaces in compile-time for UpgradeBasicGroupChatToSupergroupChatRequest.
var (
	_ bin.Encoder     = &UpgradeBasicGroupChatToSupergroupChatRequest{}
	_ bin.Decoder     = &UpgradeBasicGroupChatToSupergroupChatRequest{}
	_ bin.BareEncoder = &UpgradeBasicGroupChatToSupergroupChatRequest{}
	_ bin.BareDecoder = &UpgradeBasicGroupChatToSupergroupChatRequest{}
)

func (u *UpgradeBasicGroupChatToSupergroupChatRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UpgradeBasicGroupChatToSupergroupChatRequest) String() string {
	if u == nil {
		return "UpgradeBasicGroupChatToSupergroupChatRequest(nil)"
	}
	type Alias UpgradeBasicGroupChatToSupergroupChatRequest
	return fmt.Sprintf("UpgradeBasicGroupChatToSupergroupChatRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UpgradeBasicGroupChatToSupergroupChatRequest) TypeID() uint32 {
	return UpgradeBasicGroupChatToSupergroupChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UpgradeBasicGroupChatToSupergroupChatRequest) TypeName() string {
	return "upgradeBasicGroupChatToSupergroupChat"
}

// TypeInfo returns info about TL type.
func (u *UpgradeBasicGroupChatToSupergroupChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "upgradeBasicGroupChatToSupergroupChat",
		ID:   UpgradeBasicGroupChatToSupergroupChatRequestTypeID,
	}
	if u == nil {
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
func (u *UpgradeBasicGroupChatToSupergroupChatRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradeBasicGroupChatToSupergroupChat#11e915ba as nil")
	}
	b.PutID(UpgradeBasicGroupChatToSupergroupChatRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UpgradeBasicGroupChatToSupergroupChatRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradeBasicGroupChatToSupergroupChat#11e915ba as nil")
	}
	b.PutInt53(u.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UpgradeBasicGroupChatToSupergroupChatRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradeBasicGroupChatToSupergroupChat#11e915ba to nil")
	}
	if err := b.ConsumeID(UpgradeBasicGroupChatToSupergroupChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode upgradeBasicGroupChatToSupergroupChat#11e915ba: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UpgradeBasicGroupChatToSupergroupChatRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradeBasicGroupChatToSupergroupChat#11e915ba to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode upgradeBasicGroupChatToSupergroupChat#11e915ba: field chat_id: %w", err)
		}
		u.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (u *UpgradeBasicGroupChatToSupergroupChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode upgradeBasicGroupChatToSupergroupChat#11e915ba as nil")
	}
	b.ObjStart()
	b.PutID("upgradeBasicGroupChatToSupergroupChat")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(u.ChatID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (u *UpgradeBasicGroupChatToSupergroupChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if u == nil {
		return fmt.Errorf("can't decode upgradeBasicGroupChatToSupergroupChat#11e915ba to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("upgradeBasicGroupChatToSupergroupChat"); err != nil {
				return fmt.Errorf("unable to decode upgradeBasicGroupChatToSupergroupChat#11e915ba: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode upgradeBasicGroupChatToSupergroupChat#11e915ba: field chat_id: %w", err)
			}
			u.ChatID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (u *UpgradeBasicGroupChatToSupergroupChatRequest) GetChatID() (value int64) {
	if u == nil {
		return
	}
	return u.ChatID
}

// UpgradeBasicGroupChatToSupergroupChat invokes method upgradeBasicGroupChatToSupergroupChat#11e915ba returning error if any.
func (c *Client) UpgradeBasicGroupChatToSupergroupChat(ctx context.Context, chatid int64) (*Chat, error) {
	var result Chat

	request := &UpgradeBasicGroupChatToSupergroupChatRequest{
		ChatID: chatid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
