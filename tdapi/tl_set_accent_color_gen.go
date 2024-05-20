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

// SetAccentColorRequest represents TL type `setAccentColor#6389cb39`.
type SetAccentColorRequest struct {
	// Identifier of the accent color to use
	AccentColorID int32
	// Identifier of a custom emoji to be shown on the reply header and link preview
	// background; 0 if none
	BackgroundCustomEmojiID int64
}

// SetAccentColorRequestTypeID is TL type id of SetAccentColorRequest.
const SetAccentColorRequestTypeID = 0x6389cb39

// Ensuring interfaces in compile-time for SetAccentColorRequest.
var (
	_ bin.Encoder     = &SetAccentColorRequest{}
	_ bin.Decoder     = &SetAccentColorRequest{}
	_ bin.BareEncoder = &SetAccentColorRequest{}
	_ bin.BareDecoder = &SetAccentColorRequest{}
)

func (s *SetAccentColorRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.AccentColorID == 0) {
		return false
	}
	if !(s.BackgroundCustomEmojiID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetAccentColorRequest) String() string {
	if s == nil {
		return "SetAccentColorRequest(nil)"
	}
	type Alias SetAccentColorRequest
	return fmt.Sprintf("SetAccentColorRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetAccentColorRequest) TypeID() uint32 {
	return SetAccentColorRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetAccentColorRequest) TypeName() string {
	return "setAccentColor"
}

// TypeInfo returns info about TL type.
func (s *SetAccentColorRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setAccentColor",
		ID:   SetAccentColorRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AccentColorID",
			SchemaName: "accent_color_id",
		},
		{
			Name:       "BackgroundCustomEmojiID",
			SchemaName: "background_custom_emoji_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetAccentColorRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setAccentColor#6389cb39 as nil")
	}
	b.PutID(SetAccentColorRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetAccentColorRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setAccentColor#6389cb39 as nil")
	}
	b.PutInt32(s.AccentColorID)
	b.PutLong(s.BackgroundCustomEmojiID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetAccentColorRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setAccentColor#6389cb39 to nil")
	}
	if err := b.ConsumeID(SetAccentColorRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setAccentColor#6389cb39: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetAccentColorRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setAccentColor#6389cb39 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setAccentColor#6389cb39: field accent_color_id: %w", err)
		}
		s.AccentColorID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode setAccentColor#6389cb39: field background_custom_emoji_id: %w", err)
		}
		s.BackgroundCustomEmojiID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetAccentColorRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setAccentColor#6389cb39 as nil")
	}
	b.ObjStart()
	b.PutID("setAccentColor")
	b.Comma()
	b.FieldStart("accent_color_id")
	b.PutInt32(s.AccentColorID)
	b.Comma()
	b.FieldStart("background_custom_emoji_id")
	b.PutLong(s.BackgroundCustomEmojiID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetAccentColorRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setAccentColor#6389cb39 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setAccentColor"); err != nil {
				return fmt.Errorf("unable to decode setAccentColor#6389cb39: %w", err)
			}
		case "accent_color_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode setAccentColor#6389cb39: field accent_color_id: %w", err)
			}
			s.AccentColorID = value
		case "background_custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode setAccentColor#6389cb39: field background_custom_emoji_id: %w", err)
			}
			s.BackgroundCustomEmojiID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetAccentColorID returns value of AccentColorID field.
func (s *SetAccentColorRequest) GetAccentColorID() (value int32) {
	if s == nil {
		return
	}
	return s.AccentColorID
}

// GetBackgroundCustomEmojiID returns value of BackgroundCustomEmojiID field.
func (s *SetAccentColorRequest) GetBackgroundCustomEmojiID() (value int64) {
	if s == nil {
		return
	}
	return s.BackgroundCustomEmojiID
}

// SetAccentColor invokes method setAccentColor#6389cb39 returning error if any.
func (c *Client) SetAccentColor(ctx context.Context, request *SetAccentColorRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
