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

// GetAnimatedEmojiRequest represents TL type `getAnimatedEmoji#3f844f76`.
type GetAnimatedEmojiRequest struct {
	// The emoji
	Emoji string
}

// GetAnimatedEmojiRequestTypeID is TL type id of GetAnimatedEmojiRequest.
const GetAnimatedEmojiRequestTypeID = 0x3f844f76

// Ensuring interfaces in compile-time for GetAnimatedEmojiRequest.
var (
	_ bin.Encoder     = &GetAnimatedEmojiRequest{}
	_ bin.Decoder     = &GetAnimatedEmojiRequest{}
	_ bin.BareEncoder = &GetAnimatedEmojiRequest{}
	_ bin.BareDecoder = &GetAnimatedEmojiRequest{}
)

func (g *GetAnimatedEmojiRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Emoji == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetAnimatedEmojiRequest) String() string {
	if g == nil {
		return "GetAnimatedEmojiRequest(nil)"
	}
	type Alias GetAnimatedEmojiRequest
	return fmt.Sprintf("GetAnimatedEmojiRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetAnimatedEmojiRequest) TypeID() uint32 {
	return GetAnimatedEmojiRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetAnimatedEmojiRequest) TypeName() string {
	return "getAnimatedEmoji"
}

// TypeInfo returns info about TL type.
func (g *GetAnimatedEmojiRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getAnimatedEmoji",
		ID:   GetAnimatedEmojiRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Emoji",
			SchemaName: "emoji",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetAnimatedEmojiRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getAnimatedEmoji#3f844f76 as nil")
	}
	b.PutID(GetAnimatedEmojiRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetAnimatedEmojiRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getAnimatedEmoji#3f844f76 as nil")
	}
	b.PutString(g.Emoji)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetAnimatedEmojiRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getAnimatedEmoji#3f844f76 to nil")
	}
	if err := b.ConsumeID(GetAnimatedEmojiRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getAnimatedEmoji#3f844f76: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetAnimatedEmojiRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getAnimatedEmoji#3f844f76 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getAnimatedEmoji#3f844f76: field emoji: %w", err)
		}
		g.Emoji = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetAnimatedEmojiRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getAnimatedEmoji#3f844f76 as nil")
	}
	b.ObjStart()
	b.PutID("getAnimatedEmoji")
	b.Comma()
	b.FieldStart("emoji")
	b.PutString(g.Emoji)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetAnimatedEmojiRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getAnimatedEmoji#3f844f76 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getAnimatedEmoji"); err != nil {
				return fmt.Errorf("unable to decode getAnimatedEmoji#3f844f76: %w", err)
			}
		case "emoji":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getAnimatedEmoji#3f844f76: field emoji: %w", err)
			}
			g.Emoji = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEmoji returns value of Emoji field.
func (g *GetAnimatedEmojiRequest) GetEmoji() (value string) {
	if g == nil {
		return
	}
	return g.Emoji
}

// GetAnimatedEmoji invokes method getAnimatedEmoji#3f844f76 returning error if any.
func (c *Client) GetAnimatedEmoji(ctx context.Context, emoji string) (*AnimatedEmoji, error) {
	var result AnimatedEmoji

	request := &GetAnimatedEmojiRequest{
		Emoji: emoji,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
