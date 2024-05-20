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

// ReplaceStickerInSetRequest represents TL type `replaceStickerInSet#e7c82e19`.
type ReplaceStickerInSetRequest struct {
	// Sticker set owner; ignored for regular users
	UserID int64
	// Sticker set name. The sticker set must be owned by the current user
	Name string
	// Sticker to remove from the set
	OldSticker InputFileClass
	// Sticker to add to the set
	NewSticker InputSticker
}

// ReplaceStickerInSetRequestTypeID is TL type id of ReplaceStickerInSetRequest.
const ReplaceStickerInSetRequestTypeID = 0xe7c82e19

// Ensuring interfaces in compile-time for ReplaceStickerInSetRequest.
var (
	_ bin.Encoder     = &ReplaceStickerInSetRequest{}
	_ bin.Decoder     = &ReplaceStickerInSetRequest{}
	_ bin.BareEncoder = &ReplaceStickerInSetRequest{}
	_ bin.BareDecoder = &ReplaceStickerInSetRequest{}
)

func (r *ReplaceStickerInSetRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.UserID == 0) {
		return false
	}
	if !(r.Name == "") {
		return false
	}
	if !(r.OldSticker == nil) {
		return false
	}
	if !(r.NewSticker.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReplaceStickerInSetRequest) String() string {
	if r == nil {
		return "ReplaceStickerInSetRequest(nil)"
	}
	type Alias ReplaceStickerInSetRequest
	return fmt.Sprintf("ReplaceStickerInSetRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReplaceStickerInSetRequest) TypeID() uint32 {
	return ReplaceStickerInSetRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReplaceStickerInSetRequest) TypeName() string {
	return "replaceStickerInSet"
}

// TypeInfo returns info about TL type.
func (r *ReplaceStickerInSetRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "replaceStickerInSet",
		ID:   ReplaceStickerInSetRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "OldSticker",
			SchemaName: "old_sticker",
		},
		{
			Name:       "NewSticker",
			SchemaName: "new_sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReplaceStickerInSetRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replaceStickerInSet#e7c82e19 as nil")
	}
	b.PutID(ReplaceStickerInSetRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReplaceStickerInSetRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode replaceStickerInSet#e7c82e19 as nil")
	}
	b.PutInt53(r.UserID)
	b.PutString(r.Name)
	if r.OldSticker == nil {
		return fmt.Errorf("unable to encode replaceStickerInSet#e7c82e19: field old_sticker is nil")
	}
	if err := r.OldSticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replaceStickerInSet#e7c82e19: field old_sticker: %w", err)
	}
	if err := r.NewSticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode replaceStickerInSet#e7c82e19: field new_sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReplaceStickerInSetRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replaceStickerInSet#e7c82e19 to nil")
	}
	if err := b.ConsumeID(ReplaceStickerInSetRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode replaceStickerInSet#e7c82e19: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReplaceStickerInSetRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode replaceStickerInSet#e7c82e19 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode replaceStickerInSet#e7c82e19: field user_id: %w", err)
		}
		r.UserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode replaceStickerInSet#e7c82e19: field name: %w", err)
		}
		r.Name = value
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode replaceStickerInSet#e7c82e19: field old_sticker: %w", err)
		}
		r.OldSticker = value
	}
	{
		if err := r.NewSticker.Decode(b); err != nil {
			return fmt.Errorf("unable to decode replaceStickerInSet#e7c82e19: field new_sticker: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReplaceStickerInSetRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode replaceStickerInSet#e7c82e19 as nil")
	}
	b.ObjStart()
	b.PutID("replaceStickerInSet")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(r.UserID)
	b.Comma()
	b.FieldStart("name")
	b.PutString(r.Name)
	b.Comma()
	b.FieldStart("old_sticker")
	if r.OldSticker == nil {
		return fmt.Errorf("unable to encode replaceStickerInSet#e7c82e19: field old_sticker is nil")
	}
	if err := r.OldSticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode replaceStickerInSet#e7c82e19: field old_sticker: %w", err)
	}
	b.Comma()
	b.FieldStart("new_sticker")
	if err := r.NewSticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode replaceStickerInSet#e7c82e19: field new_sticker: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReplaceStickerInSetRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode replaceStickerInSet#e7c82e19 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("replaceStickerInSet"); err != nil {
				return fmt.Errorf("unable to decode replaceStickerInSet#e7c82e19: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode replaceStickerInSet#e7c82e19: field user_id: %w", err)
			}
			r.UserID = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode replaceStickerInSet#e7c82e19: field name: %w", err)
			}
			r.Name = value
		case "old_sticker":
			value, err := DecodeTDLibJSONInputFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode replaceStickerInSet#e7c82e19: field old_sticker: %w", err)
			}
			r.OldSticker = value
		case "new_sticker":
			if err := r.NewSticker.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode replaceStickerInSet#e7c82e19: field new_sticker: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (r *ReplaceStickerInSetRequest) GetUserID() (value int64) {
	if r == nil {
		return
	}
	return r.UserID
}

// GetName returns value of Name field.
func (r *ReplaceStickerInSetRequest) GetName() (value string) {
	if r == nil {
		return
	}
	return r.Name
}

// GetOldSticker returns value of OldSticker field.
func (r *ReplaceStickerInSetRequest) GetOldSticker() (value InputFileClass) {
	if r == nil {
		return
	}
	return r.OldSticker
}

// GetNewSticker returns value of NewSticker field.
func (r *ReplaceStickerInSetRequest) GetNewSticker() (value InputSticker) {
	if r == nil {
		return
	}
	return r.NewSticker
}

// ReplaceStickerInSet invokes method replaceStickerInSet#e7c82e19 returning error if any.
func (c *Client) ReplaceStickerInSet(ctx context.Context, request *ReplaceStickerInSetRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
