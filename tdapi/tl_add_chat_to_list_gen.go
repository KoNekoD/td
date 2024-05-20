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

// AddChatToListRequest represents TL type `addChatToList#fb334eb5`.
type AddChatToListRequest struct {
	// Chat identifier
	ChatID int64
	// The chat list. Use getChatListsToAddChat to get suitable chat lists
	ChatList ChatListClass
}

// AddChatToListRequestTypeID is TL type id of AddChatToListRequest.
const AddChatToListRequestTypeID = 0xfb334eb5

// Ensuring interfaces in compile-time for AddChatToListRequest.
var (
	_ bin.Encoder     = &AddChatToListRequest{}
	_ bin.Decoder     = &AddChatToListRequest{}
	_ bin.BareEncoder = &AddChatToListRequest{}
	_ bin.BareDecoder = &AddChatToListRequest{}
)

func (a *AddChatToListRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ChatID == 0) {
		return false
	}
	if !(a.ChatList == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddChatToListRequest) String() string {
	if a == nil {
		return "AddChatToListRequest(nil)"
	}
	type Alias AddChatToListRequest
	return fmt.Sprintf("AddChatToListRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddChatToListRequest) TypeID() uint32 {
	return AddChatToListRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddChatToListRequest) TypeName() string {
	return "addChatToList"
}

// TypeInfo returns info about TL type.
func (a *AddChatToListRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addChatToList",
		ID:   AddChatToListRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "ChatList",
			SchemaName: "chat_list",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddChatToListRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatToList#fb334eb5 as nil")
	}
	b.PutID(AddChatToListRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddChatToListRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatToList#fb334eb5 as nil")
	}
	b.PutInt53(a.ChatID)
	if a.ChatList == nil {
		return fmt.Errorf("unable to encode addChatToList#fb334eb5: field chat_list is nil")
	}
	if err := a.ChatList.Encode(b); err != nil {
		return fmt.Errorf("unable to encode addChatToList#fb334eb5: field chat_list: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AddChatToListRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatToList#fb334eb5 to nil")
	}
	if err := b.ConsumeID(AddChatToListRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addChatToList#fb334eb5: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddChatToListRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatToList#fb334eb5 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode addChatToList#fb334eb5: field chat_id: %w", err)
		}
		a.ChatID = value
	}
	{
		value, err := DecodeChatList(b)
		if err != nil {
			return fmt.Errorf("unable to decode addChatToList#fb334eb5: field chat_list: %w", err)
		}
		a.ChatList = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AddChatToListRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addChatToList#fb334eb5 as nil")
	}
	b.ObjStart()
	b.PutID("addChatToList")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(a.ChatID)
	b.Comma()
	b.FieldStart("chat_list")
	if a.ChatList == nil {
		return fmt.Errorf("unable to encode addChatToList#fb334eb5: field chat_list is nil")
	}
	if err := a.ChatList.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode addChatToList#fb334eb5: field chat_list: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AddChatToListRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode addChatToList#fb334eb5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("addChatToList"); err != nil {
				return fmt.Errorf("unable to decode addChatToList#fb334eb5: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode addChatToList#fb334eb5: field chat_id: %w", err)
			}
			a.ChatID = value
		case "chat_list":
			value, err := DecodeTDLibJSONChatList(b)
			if err != nil {
				return fmt.Errorf("unable to decode addChatToList#fb334eb5: field chat_list: %w", err)
			}
			a.ChatList = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (a *AddChatToListRequest) GetChatID() (value int64) {
	if a == nil {
		return
	}
	return a.ChatID
}

// GetChatList returns value of ChatList field.
func (a *AddChatToListRequest) GetChatList() (value ChatListClass) {
	if a == nil {
		return
	}
	return a.ChatList
}

// AddChatToList invokes method addChatToList#fb334eb5 returning error if any.
func (c *Client) AddChatToList(ctx context.Context, request *AddChatToListRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
