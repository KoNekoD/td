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

// GetChatFolderRequest represents TL type `getChatFolder#5882a98`.
type GetChatFolderRequest struct {
	// Chat folder identifier
	ChatFolderID int32
}

// GetChatFolderRequestTypeID is TL type id of GetChatFolderRequest.
const GetChatFolderRequestTypeID = 0x5882a98

// Ensuring interfaces in compile-time for GetChatFolderRequest.
var (
	_ bin.Encoder     = &GetChatFolderRequest{}
	_ bin.Decoder     = &GetChatFolderRequest{}
	_ bin.BareEncoder = &GetChatFolderRequest{}
	_ bin.BareDecoder = &GetChatFolderRequest{}
)

func (g *GetChatFolderRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatFolderID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatFolderRequest) String() string {
	if g == nil {
		return "GetChatFolderRequest(nil)"
	}
	type Alias GetChatFolderRequest
	return fmt.Sprintf("GetChatFolderRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatFolderRequest) TypeID() uint32 {
	return GetChatFolderRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatFolderRequest) TypeName() string {
	return "getChatFolder"
}

// TypeInfo returns info about TL type.
func (g *GetChatFolderRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatFolder",
		ID:   GetChatFolderRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatFolderID",
			SchemaName: "chat_folder_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatFolderRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFolder#5882a98 as nil")
	}
	b.PutID(GetChatFolderRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatFolderRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFolder#5882a98 as nil")
	}
	b.PutInt32(g.ChatFolderID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatFolderRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFolder#5882a98 to nil")
	}
	if err := b.ConsumeID(GetChatFolderRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatFolder#5882a98: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatFolderRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFolder#5882a98 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatFolder#5882a98: field chat_folder_id: %w", err)
		}
		g.ChatFolderID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatFolderRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatFolder#5882a98 as nil")
	}
	b.ObjStart()
	b.PutID("getChatFolder")
	b.Comma()
	b.FieldStart("chat_folder_id")
	b.PutInt32(g.ChatFolderID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetChatFolderRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatFolder#5882a98 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatFolder"); err != nil {
				return fmt.Errorf("unable to decode getChatFolder#5882a98: %w", err)
			}
		case "chat_folder_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatFolder#5882a98: field chat_folder_id: %w", err)
			}
			g.ChatFolderID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatFolderID returns value of ChatFolderID field.
func (g *GetChatFolderRequest) GetChatFolderID() (value int32) {
	if g == nil {
		return
	}
	return g.ChatFolderID
}

// GetChatFolder invokes method getChatFolder#5882a98 returning error if any.
func (c *Client) GetChatFolder(ctx context.Context, chatfolderid int32) (*ChatFolder, error) {
	var result ChatFolder

	request := &GetChatFolderRequest{
		ChatFolderID: chatfolderid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
