// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/jsontd"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
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
	_ = jsontd.Encoder{}
)

// RemoveRecentStickerRequest represents TL type `removeRecentSticker#4a4d440d`.
type RemoveRecentStickerRequest struct {
	// Pass true to remove the sticker from the list of stickers recently attached to photo
	// or video files; pass false to remove the sticker from the list of recently sent
	// stickers
	IsAttached bool
	// Sticker file to delete
	Sticker InputFileClass
}

// RemoveRecentStickerRequestTypeID is TL type id of RemoveRecentStickerRequest.
const RemoveRecentStickerRequestTypeID = 0x4a4d440d

// Ensuring interfaces in compile-time for RemoveRecentStickerRequest.
var (
	_ bin.Encoder     = &RemoveRecentStickerRequest{}
	_ bin.Decoder     = &RemoveRecentStickerRequest{}
	_ bin.BareEncoder = &RemoveRecentStickerRequest{}
	_ bin.BareDecoder = &RemoveRecentStickerRequest{}
)

func (r *RemoveRecentStickerRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.IsAttached == false) {
		return false
	}
	if !(r.Sticker == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RemoveRecentStickerRequest) String() string {
	if r == nil {
		return "RemoveRecentStickerRequest(nil)"
	}
	type Alias RemoveRecentStickerRequest
	return fmt.Sprintf("RemoveRecentStickerRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RemoveRecentStickerRequest) TypeID() uint32 {
	return RemoveRecentStickerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RemoveRecentStickerRequest) TypeName() string {
	return "removeRecentSticker"
}

// TypeInfo returns info about TL type.
func (r *RemoveRecentStickerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "removeRecentSticker",
		ID:   RemoveRecentStickerRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsAttached",
			SchemaName: "is_attached",
		},
		{
			Name:       "Sticker",
			SchemaName: "sticker",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RemoveRecentStickerRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeRecentSticker#4a4d440d as nil")
	}
	b.PutID(RemoveRecentStickerRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RemoveRecentStickerRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode removeRecentSticker#4a4d440d as nil")
	}
	b.PutBool(r.IsAttached)
	if r.Sticker == nil {
		return fmt.Errorf("unable to encode removeRecentSticker#4a4d440d: field sticker is nil")
	}
	if err := r.Sticker.Encode(b); err != nil {
		return fmt.Errorf("unable to encode removeRecentSticker#4a4d440d: field sticker: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RemoveRecentStickerRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeRecentSticker#4a4d440d to nil")
	}
	if err := b.ConsumeID(RemoveRecentStickerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode removeRecentSticker#4a4d440d: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RemoveRecentStickerRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode removeRecentSticker#4a4d440d to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode removeRecentSticker#4a4d440d: field is_attached: %w", err)
		}
		r.IsAttached = value
	}
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode removeRecentSticker#4a4d440d: field sticker: %w", err)
		}
		r.Sticker = value
	}
	return nil
}

// EncodeTDLibJSON encodes r in TDLib API JSON format.
func (r *RemoveRecentStickerRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode removeRecentSticker#4a4d440d as nil")
	}
	b.ObjStart()
	b.PutID("removeRecentSticker")
	b.FieldStart("is_attached")
	b.PutBool(r.IsAttached)
	b.FieldStart("sticker")
	if r.Sticker == nil {
		return fmt.Errorf("unable to encode removeRecentSticker#4a4d440d: field sticker is nil")
	}
	if err := r.Sticker.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode removeRecentSticker#4a4d440d: field sticker: %w", err)
	}
	b.ObjEnd()
	return nil
}

// GetIsAttached returns value of IsAttached field.
func (r *RemoveRecentStickerRequest) GetIsAttached() (value bool) {
	return r.IsAttached
}

// GetSticker returns value of Sticker field.
func (r *RemoveRecentStickerRequest) GetSticker() (value InputFileClass) {
	return r.Sticker
}

// RemoveRecentSticker invokes method removeRecentSticker#4a4d440d returning error if any.
func (c *Client) RemoveRecentSticker(ctx context.Context, request *RemoveRecentStickerRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
