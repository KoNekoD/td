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

// GetInlineGameHighScoresRequest represents TL type `getInlineGameHighScores#e0396baa`.
type GetInlineGameHighScoresRequest struct {
	// Inline message identifier
	InlineMessageID string
	// User identifier
	UserID int64
}

// GetInlineGameHighScoresRequestTypeID is TL type id of GetInlineGameHighScoresRequest.
const GetInlineGameHighScoresRequestTypeID = 0xe0396baa

// Ensuring interfaces in compile-time for GetInlineGameHighScoresRequest.
var (
	_ bin.Encoder     = &GetInlineGameHighScoresRequest{}
	_ bin.Decoder     = &GetInlineGameHighScoresRequest{}
	_ bin.BareEncoder = &GetInlineGameHighScoresRequest{}
	_ bin.BareDecoder = &GetInlineGameHighScoresRequest{}
)

func (g *GetInlineGameHighScoresRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.InlineMessageID == "") {
		return false
	}
	if !(g.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetInlineGameHighScoresRequest) String() string {
	if g == nil {
		return "GetInlineGameHighScoresRequest(nil)"
	}
	type Alias GetInlineGameHighScoresRequest
	return fmt.Sprintf("GetInlineGameHighScoresRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetInlineGameHighScoresRequest) TypeID() uint32 {
	return GetInlineGameHighScoresRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetInlineGameHighScoresRequest) TypeName() string {
	return "getInlineGameHighScores"
}

// TypeInfo returns info about TL type.
func (g *GetInlineGameHighScoresRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getInlineGameHighScores",
		ID:   GetInlineGameHighScoresRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InlineMessageID",
			SchemaName: "inline_message_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetInlineGameHighScoresRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getInlineGameHighScores#e0396baa as nil")
	}
	b.PutID(GetInlineGameHighScoresRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetInlineGameHighScoresRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getInlineGameHighScores#e0396baa as nil")
	}
	b.PutString(g.InlineMessageID)
	b.PutInt53(g.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetInlineGameHighScoresRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getInlineGameHighScores#e0396baa to nil")
	}
	if err := b.ConsumeID(GetInlineGameHighScoresRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getInlineGameHighScores#e0396baa: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetInlineGameHighScoresRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getInlineGameHighScores#e0396baa to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getInlineGameHighScores#e0396baa: field inline_message_id: %w", err)
		}
		g.InlineMessageID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getInlineGameHighScores#e0396baa: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetInlineGameHighScoresRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getInlineGameHighScores#e0396baa as nil")
	}
	b.ObjStart()
	b.PutID("getInlineGameHighScores")
	b.Comma()
	b.FieldStart("inline_message_id")
	b.PutString(g.InlineMessageID)
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(g.UserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetInlineGameHighScoresRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getInlineGameHighScores#e0396baa to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getInlineGameHighScores"); err != nil {
				return fmt.Errorf("unable to decode getInlineGameHighScores#e0396baa: %w", err)
			}
		case "inline_message_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getInlineGameHighScores#e0396baa: field inline_message_id: %w", err)
			}
			g.InlineMessageID = value
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getInlineGameHighScores#e0396baa: field user_id: %w", err)
			}
			g.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInlineMessageID returns value of InlineMessageID field.
func (g *GetInlineGameHighScoresRequest) GetInlineMessageID() (value string) {
	if g == nil {
		return
	}
	return g.InlineMessageID
}

// GetUserID returns value of UserID field.
func (g *GetInlineGameHighScoresRequest) GetUserID() (value int64) {
	if g == nil {
		return
	}
	return g.UserID
}

// GetInlineGameHighScores invokes method getInlineGameHighScores#e0396baa returning error if any.
func (c *Client) GetInlineGameHighScores(ctx context.Context, request *GetInlineGameHighScoresRequest) (*GameHighScores, error) {
	var result GameHighScores

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
