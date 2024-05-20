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

// GetBotInfoShortDescriptionRequest represents TL type `getBotInfoShortDescription#4a1c2614`.
type GetBotInfoShortDescriptionRequest struct {
	// Identifier of the target bot
	BotUserID int64
	// A two-letter ISO 639-1 language code or an empty string
	LanguageCode string
}

// GetBotInfoShortDescriptionRequestTypeID is TL type id of GetBotInfoShortDescriptionRequest.
const GetBotInfoShortDescriptionRequestTypeID = 0x4a1c2614

// Ensuring interfaces in compile-time for GetBotInfoShortDescriptionRequest.
var (
	_ bin.Encoder     = &GetBotInfoShortDescriptionRequest{}
	_ bin.Decoder     = &GetBotInfoShortDescriptionRequest{}
	_ bin.BareEncoder = &GetBotInfoShortDescriptionRequest{}
	_ bin.BareDecoder = &GetBotInfoShortDescriptionRequest{}
)

func (g *GetBotInfoShortDescriptionRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.BotUserID == 0) {
		return false
	}
	if !(g.LanguageCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetBotInfoShortDescriptionRequest) String() string {
	if g == nil {
		return "GetBotInfoShortDescriptionRequest(nil)"
	}
	type Alias GetBotInfoShortDescriptionRequest
	return fmt.Sprintf("GetBotInfoShortDescriptionRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetBotInfoShortDescriptionRequest) TypeID() uint32 {
	return GetBotInfoShortDescriptionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetBotInfoShortDescriptionRequest) TypeName() string {
	return "getBotInfoShortDescription"
}

// TypeInfo returns info about TL type.
func (g *GetBotInfoShortDescriptionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getBotInfoShortDescription",
		ID:   GetBotInfoShortDescriptionRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "LanguageCode",
			SchemaName: "language_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetBotInfoShortDescriptionRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBotInfoShortDescription#4a1c2614 as nil")
	}
	b.PutID(GetBotInfoShortDescriptionRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetBotInfoShortDescriptionRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getBotInfoShortDescription#4a1c2614 as nil")
	}
	b.PutInt53(g.BotUserID)
	b.PutString(g.LanguageCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetBotInfoShortDescriptionRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBotInfoShortDescription#4a1c2614 to nil")
	}
	if err := b.ConsumeID(GetBotInfoShortDescriptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getBotInfoShortDescription#4a1c2614: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetBotInfoShortDescriptionRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getBotInfoShortDescription#4a1c2614 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getBotInfoShortDescription#4a1c2614: field bot_user_id: %w", err)
		}
		g.BotUserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getBotInfoShortDescription#4a1c2614: field language_code: %w", err)
		}
		g.LanguageCode = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetBotInfoShortDescriptionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getBotInfoShortDescription#4a1c2614 as nil")
	}
	b.ObjStart()
	b.PutID("getBotInfoShortDescription")
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(g.BotUserID)
	b.Comma()
	b.FieldStart("language_code")
	b.PutString(g.LanguageCode)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetBotInfoShortDescriptionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getBotInfoShortDescription#4a1c2614 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getBotInfoShortDescription"); err != nil {
				return fmt.Errorf("unable to decode getBotInfoShortDescription#4a1c2614: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getBotInfoShortDescription#4a1c2614: field bot_user_id: %w", err)
			}
			g.BotUserID = value
		case "language_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getBotInfoShortDescription#4a1c2614: field language_code: %w", err)
			}
			g.LanguageCode = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (g *GetBotInfoShortDescriptionRequest) GetBotUserID() (value int64) {
	if g == nil {
		return
	}
	return g.BotUserID
}

// GetLanguageCode returns value of LanguageCode field.
func (g *GetBotInfoShortDescriptionRequest) GetLanguageCode() (value string) {
	if g == nil {
		return
	}
	return g.LanguageCode
}

// GetBotInfoShortDescription invokes method getBotInfoShortDescription#4a1c2614 returning error if any.
func (c *Client) GetBotInfoShortDescription(ctx context.Context, request *GetBotInfoShortDescriptionRequest) (*Text, error) {
	var result Text

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
