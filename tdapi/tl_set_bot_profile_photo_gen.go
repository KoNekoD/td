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

// SetBotProfilePhotoRequest represents TL type `setBotProfilePhoto#bd864b66`.
type SetBotProfilePhotoRequest struct {
	// Identifier of the target bot
	BotUserID int64
	// Profile photo to set; pass null to delete the chat photo
	Photo InputChatPhotoClass
}

// SetBotProfilePhotoRequestTypeID is TL type id of SetBotProfilePhotoRequest.
const SetBotProfilePhotoRequestTypeID = 0xbd864b66

// Ensuring interfaces in compile-time for SetBotProfilePhotoRequest.
var (
	_ bin.Encoder     = &SetBotProfilePhotoRequest{}
	_ bin.Decoder     = &SetBotProfilePhotoRequest{}
	_ bin.BareEncoder = &SetBotProfilePhotoRequest{}
	_ bin.BareDecoder = &SetBotProfilePhotoRequest{}
)

func (s *SetBotProfilePhotoRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.BotUserID == 0) {
		return false
	}
	if !(s.Photo == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBotProfilePhotoRequest) String() string {
	if s == nil {
		return "SetBotProfilePhotoRequest(nil)"
	}
	type Alias SetBotProfilePhotoRequest
	return fmt.Sprintf("SetBotProfilePhotoRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBotProfilePhotoRequest) TypeID() uint32 {
	return SetBotProfilePhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBotProfilePhotoRequest) TypeName() string {
	return "setBotProfilePhoto"
}

// TypeInfo returns info about TL type.
func (s *SetBotProfilePhotoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBotProfilePhoto",
		ID:   SetBotProfilePhotoRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBotProfilePhotoRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotProfilePhoto#bd864b66 as nil")
	}
	b.PutID(SetBotProfilePhotoRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBotProfilePhotoRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotProfilePhoto#bd864b66 as nil")
	}
	b.PutInt53(s.BotUserID)
	if s.Photo == nil {
		return fmt.Errorf("unable to encode setBotProfilePhoto#bd864b66: field photo is nil")
	}
	if err := s.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setBotProfilePhoto#bd864b66: field photo: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBotProfilePhotoRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotProfilePhoto#bd864b66 to nil")
	}
	if err := b.ConsumeID(SetBotProfilePhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBotProfilePhoto#bd864b66: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBotProfilePhotoRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotProfilePhoto#bd864b66 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setBotProfilePhoto#bd864b66: field bot_user_id: %w", err)
		}
		s.BotUserID = value
	}
	{
		value, err := DecodeInputChatPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode setBotProfilePhoto#bd864b66: field photo: %w", err)
		}
		s.Photo = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetBotProfilePhotoRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotProfilePhoto#bd864b66 as nil")
	}
	b.ObjStart()
	b.PutID("setBotProfilePhoto")
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(s.BotUserID)
	b.Comma()
	b.FieldStart("photo")
	if s.Photo == nil {
		return fmt.Errorf("unable to encode setBotProfilePhoto#bd864b66: field photo is nil")
	}
	if err := s.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setBotProfilePhoto#bd864b66: field photo: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetBotProfilePhotoRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotProfilePhoto#bd864b66 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setBotProfilePhoto"); err != nil {
				return fmt.Errorf("unable to decode setBotProfilePhoto#bd864b66: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setBotProfilePhoto#bd864b66: field bot_user_id: %w", err)
			}
			s.BotUserID = value
		case "photo":
			value, err := DecodeTDLibJSONInputChatPhoto(b)
			if err != nil {
				return fmt.Errorf("unable to decode setBotProfilePhoto#bd864b66: field photo: %w", err)
			}
			s.Photo = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (s *SetBotProfilePhotoRequest) GetBotUserID() (value int64) {
	if s == nil {
		return
	}
	return s.BotUserID
}

// GetPhoto returns value of Photo field.
func (s *SetBotProfilePhotoRequest) GetPhoto() (value InputChatPhotoClass) {
	if s == nil {
		return
	}
	return s.Photo
}

// SetBotProfilePhoto invokes method setBotProfilePhoto#bd864b66 returning error if any.
func (c *Client) SetBotProfilePhoto(ctx context.Context, request *SetBotProfilePhotoRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
