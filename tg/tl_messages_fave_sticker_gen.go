// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesFaveStickerRequest represents TL type `messages.faveSticker#b9ffc55b`.
// Mark a sticker as favorite
//
// See https://core.telegram.org/method/messages.faveSticker for reference.
type MessagesFaveStickerRequest struct {
	// Sticker to mark as favorite
	ID InputDocumentClass
	// Unfavorite
	Unfave bool
}

// MessagesFaveStickerRequestTypeID is TL type id of MessagesFaveStickerRequest.
const MessagesFaveStickerRequestTypeID = 0xb9ffc55b

// Ensuring interfaces in compile-time for MessagesFaveStickerRequest.
var (
	_ bin.Encoder     = &MessagesFaveStickerRequest{}
	_ bin.Decoder     = &MessagesFaveStickerRequest{}
	_ bin.BareEncoder = &MessagesFaveStickerRequest{}
	_ bin.BareDecoder = &MessagesFaveStickerRequest{}
)

func (f *MessagesFaveStickerRequest) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.ID == nil) {
		return false
	}
	if !(f.Unfave == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *MessagesFaveStickerRequest) String() string {
	if f == nil {
		return "MessagesFaveStickerRequest(nil)"
	}
	type Alias MessagesFaveStickerRequest
	return fmt.Sprintf("MessagesFaveStickerRequest%+v", Alias(*f))
}

// FillFrom fills MessagesFaveStickerRequest from given interface.
func (f *MessagesFaveStickerRequest) FillFrom(from interface {
	GetID() (value InputDocumentClass)
	GetUnfave() (value bool)
}) {
	f.ID = from.GetID()
	f.Unfave = from.GetUnfave()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesFaveStickerRequest) TypeID() uint32 {
	return MessagesFaveStickerRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesFaveStickerRequest) TypeName() string {
	return "messages.faveSticker"
}

// TypeInfo returns info about TL type.
func (f *MessagesFaveStickerRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.faveSticker",
		ID:   MessagesFaveStickerRequestTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Unfave",
			SchemaName: "unfave",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *MessagesFaveStickerRequest) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.faveSticker#b9ffc55b as nil")
	}
	b.PutID(MessagesFaveStickerRequestTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *MessagesFaveStickerRequest) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode messages.faveSticker#b9ffc55b as nil")
	}
	if f.ID == nil {
		return fmt.Errorf("unable to encode messages.faveSticker#b9ffc55b: field id is nil")
	}
	if err := f.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.faveSticker#b9ffc55b: field id: %w", err)
	}
	b.PutBool(f.Unfave)
	return nil
}

// Decode implements bin.Decoder.
func (f *MessagesFaveStickerRequest) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.faveSticker#b9ffc55b to nil")
	}
	if err := b.ConsumeID(MessagesFaveStickerRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.faveSticker#b9ffc55b: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *MessagesFaveStickerRequest) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode messages.faveSticker#b9ffc55b to nil")
	}
	{
		value, err := DecodeInputDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.faveSticker#b9ffc55b: field id: %w", err)
		}
		f.ID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messages.faveSticker#b9ffc55b: field unfave: %w", err)
		}
		f.Unfave = value
	}
	return nil
}

// GetID returns value of ID field.
func (f *MessagesFaveStickerRequest) GetID() (value InputDocumentClass) {
	return f.ID
}

// GetUnfave returns value of Unfave field.
func (f *MessagesFaveStickerRequest) GetUnfave() (value bool) {
	return f.Unfave
}

// GetIDAsNotEmpty returns mapped value of ID field.
func (f *MessagesFaveStickerRequest) GetIDAsNotEmpty() (*InputDocument, bool) {
	return f.ID.AsNotEmpty()
}

// MessagesFaveSticker invokes method messages.faveSticker#b9ffc55b returning error if any.
// Mark a sticker as favorite
//
// Possible errors:
//  400 STICKER_ID_INVALID: The provided sticker ID is invalid
//
// See https://core.telegram.org/method/messages.faveSticker for reference.
func (c *Client) MessagesFaveSticker(ctx context.Context, request *MessagesFaveStickerRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
