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

// MessagesGetFeaturedStickersRequest represents TL type `messages.getFeaturedStickers#64780b14`.
// Get featured stickers
//
// See https://core.telegram.org/method/messages.getFeaturedStickers for reference.
type MessagesGetFeaturedStickersRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// MessagesGetFeaturedStickersRequestTypeID is TL type id of MessagesGetFeaturedStickersRequest.
const MessagesGetFeaturedStickersRequestTypeID = 0x64780b14

// Ensuring interfaces in compile-time for MessagesGetFeaturedStickersRequest.
var (
	_ bin.Encoder     = &MessagesGetFeaturedStickersRequest{}
	_ bin.Decoder     = &MessagesGetFeaturedStickersRequest{}
	_ bin.BareEncoder = &MessagesGetFeaturedStickersRequest{}
	_ bin.BareDecoder = &MessagesGetFeaturedStickersRequest{}
)

func (g *MessagesGetFeaturedStickersRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetFeaturedStickersRequest) String() string {
	if g == nil {
		return "MessagesGetFeaturedStickersRequest(nil)"
	}
	type Alias MessagesGetFeaturedStickersRequest
	return fmt.Sprintf("MessagesGetFeaturedStickersRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetFeaturedStickersRequest from given interface.
func (g *MessagesGetFeaturedStickersRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetFeaturedStickersRequest) TypeID() uint32 {
	return MessagesGetFeaturedStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetFeaturedStickersRequest) TypeName() string {
	return "messages.getFeaturedStickers"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetFeaturedStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getFeaturedStickers",
		ID:   MessagesGetFeaturedStickersRequestTypeID,
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
func (g *MessagesGetFeaturedStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getFeaturedStickers#64780b14 as nil")
	}
	b.PutID(MessagesGetFeaturedStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetFeaturedStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getFeaturedStickers#64780b14 as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetFeaturedStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getFeaturedStickers#64780b14 to nil")
	}
	if err := b.ConsumeID(MessagesGetFeaturedStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getFeaturedStickers#64780b14: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetFeaturedStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getFeaturedStickers#64780b14 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messages.getFeaturedStickers#64780b14: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *MessagesGetFeaturedStickersRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// MessagesGetFeaturedStickers invokes method messages.getFeaturedStickers#64780b14 returning error if any.
// Get featured stickers
//
// See https://core.telegram.org/method/messages.getFeaturedStickers for reference.
func (c *Client) MessagesGetFeaturedStickers(ctx context.Context, hash int64) (MessagesFeaturedStickersClass, error) {
	var result MessagesFeaturedStickersBox

	request := &MessagesGetFeaturedStickersRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.FeaturedStickers, nil
}
