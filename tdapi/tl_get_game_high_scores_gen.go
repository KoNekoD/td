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

// GetGameHighScoresRequest represents TL type `getGameHighScores#f0459b`.
type GetGameHighScoresRequest struct {
	// The chat that contains the message with the game
	ChatID int64
	// Identifier of the message
	MessageID int64
	// User identifier
	UserID int64
}

// GetGameHighScoresRequestTypeID is TL type id of GetGameHighScoresRequest.
const GetGameHighScoresRequestTypeID = 0xf0459b

// Ensuring interfaces in compile-time for GetGameHighScoresRequest.
var (
	_ bin.Encoder     = &GetGameHighScoresRequest{}
	_ bin.Decoder     = &GetGameHighScoresRequest{}
	_ bin.BareEncoder = &GetGameHighScoresRequest{}
	_ bin.BareDecoder = &GetGameHighScoresRequest{}
)

func (g *GetGameHighScoresRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}
	if !(g.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetGameHighScoresRequest) String() string {
	if g == nil {
		return "GetGameHighScoresRequest(nil)"
	}
	type Alias GetGameHighScoresRequest
	return fmt.Sprintf("GetGameHighScoresRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetGameHighScoresRequest) TypeID() uint32 {
	return GetGameHighScoresRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetGameHighScoresRequest) TypeName() string {
	return "getGameHighScores"
}

// TypeInfo returns info about TL type.
func (g *GetGameHighScoresRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getGameHighScores",
		ID:   GetGameHighScoresRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetGameHighScoresRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGameHighScores#f0459b as nil")
	}
	b.PutID(GetGameHighScoresRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetGameHighScoresRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getGameHighScores#f0459b as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutInt53(g.MessageID)
	b.PutInt53(g.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetGameHighScoresRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGameHighScores#f0459b to nil")
	}
	if err := b.ConsumeID(GetGameHighScoresRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getGameHighScores#f0459b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetGameHighScoresRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getGameHighScores#f0459b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getGameHighScores#f0459b: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getGameHighScores#f0459b: field message_id: %w", err)
		}
		g.MessageID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getGameHighScores#f0459b: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetGameHighScoresRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getGameHighScores#f0459b as nil")
	}
	b.ObjStart()
	b.PutID("getGameHighScores")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(g.MessageID)
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(g.UserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetGameHighScoresRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getGameHighScores#f0459b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getGameHighScores"); err != nil {
				return fmt.Errorf("unable to decode getGameHighScores#f0459b: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getGameHighScores#f0459b: field chat_id: %w", err)
			}
			g.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getGameHighScores#f0459b: field message_id: %w", err)
			}
			g.MessageID = value
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getGameHighScores#f0459b: field user_id: %w", err)
			}
			g.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetGameHighScoresRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetGameHighScoresRequest) GetMessageID() (value int64) {
	if g == nil {
		return
	}
	return g.MessageID
}

// GetUserID returns value of UserID field.
func (g *GetGameHighScoresRequest) GetUserID() (value int64) {
	if g == nil {
		return
	}
	return g.UserID
}

// GetGameHighScores invokes method getGameHighScores#f0459b returning error if any.
func (c *Client) GetGameHighScores(ctx context.Context, request *GetGameHighScoresRequest) (*GameHighScores, error) {
	var result GameHighScores

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
