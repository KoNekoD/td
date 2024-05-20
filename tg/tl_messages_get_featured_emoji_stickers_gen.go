// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// MessagesGetFeaturedEmojiStickersRequest represents TL type `messages.getFeaturedEmojiStickers#ecf6736`.
// Gets featured custom emoji stickersets.
//
// See https://core.telegram.org/method/messages.getFeaturedEmojiStickers for reference.
type MessagesGetFeaturedEmojiStickersRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// MessagesGetFeaturedEmojiStickersRequestTypeID is TL type id of MessagesGetFeaturedEmojiStickersRequest.
const MessagesGetFeaturedEmojiStickersRequestTypeID = 0xecf6736

// Ensuring interfaces in compile-time for MessagesGetFeaturedEmojiStickersRequest.
var (
	_ bin.Encoder     = &MessagesGetFeaturedEmojiStickersRequest{}
	_ bin.Decoder     = &MessagesGetFeaturedEmojiStickersRequest{}
	_ bin.BareEncoder = &MessagesGetFeaturedEmojiStickersRequest{}
	_ bin.BareDecoder = &MessagesGetFeaturedEmojiStickersRequest{}
)

func (g *MessagesGetFeaturedEmojiStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetFeaturedEmojiStickersRequest) String() string {
	if g == nil {
		return "MessagesGetFeaturedEmojiStickersRequest(nil)"
	}
	type Alias MessagesGetFeaturedEmojiStickersRequest
	return fmt.Sprintf("MessagesGetFeaturedEmojiStickersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetFeaturedEmojiStickersRequest from given interface.
func (g *MessagesGetFeaturedEmojiStickersRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetFeaturedEmojiStickersRequest) TypeID() uint32 {
	return MessagesGetFeaturedEmojiStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetFeaturedEmojiStickersRequest) TypeName() string {
	return "messages.getFeaturedEmojiStickers"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetFeaturedEmojiStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getFeaturedEmojiStickers",
		ID:   MessagesGetFeaturedEmojiStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetFeaturedEmojiStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getFeaturedEmojiStickers#ecf6736 as nil")
	}
	b.PutID(MessagesGetFeaturedEmojiStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetFeaturedEmojiStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getFeaturedEmojiStickers#ecf6736 as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetFeaturedEmojiStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getFeaturedEmojiStickers#ecf6736 to nil")
	}
	if err := b.ConsumeID(MessagesGetFeaturedEmojiStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getFeaturedEmojiStickers#ecf6736: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetFeaturedEmojiStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getFeaturedEmojiStickers#ecf6736 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getFeaturedEmojiStickers#ecf6736: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetFeaturedEmojiStickersRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetFeaturedEmojiStickers invokes method messages.getFeaturedEmojiStickers#ecf6736 returning error if any.
// Gets featured custom emoji stickersets.
//
// See https://core.telegram.org/method/messages.getFeaturedEmojiStickers for reference.
func (c *Client) MessagesGetFeaturedEmojiStickers(ctx context.Context, hash int64) (MessagesFeaturedStickersClass, error) {
	var result MessagesFeaturedStickersBox

	request := &MessagesGetFeaturedEmojiStickersRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.FeaturedStickers, nil
}
