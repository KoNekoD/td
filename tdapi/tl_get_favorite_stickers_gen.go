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

// GetFavoriteStickersRequest represents TL type `getFavoriteStickers#ebcbcf40`.
type GetFavoriteStickersRequest struct {
}

// GetFavoriteStickersRequestTypeID is TL type id of GetFavoriteStickersRequest.
const GetFavoriteStickersRequestTypeID = 0xebcbcf40

// Ensuring interfaces in compile-time for GetFavoriteStickersRequest.
var (
	_ bin.Encoder     = &GetFavoriteStickersRequest{}
	_ bin.Decoder     = &GetFavoriteStickersRequest{}
	_ bin.BareEncoder = &GetFavoriteStickersRequest{}
	_ bin.BareDecoder = &GetFavoriteStickersRequest{}
)

func (g *GetFavoriteStickersRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetFavoriteStickersRequest) String() string {
	if g == nil {
		return "GetFavoriteStickersRequest(nil)"
	}
	type Alias GetFavoriteStickersRequest
	return fmt.Sprintf("GetFavoriteStickersRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetFavoriteStickersRequest) TypeID() uint32 {
	return GetFavoriteStickersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetFavoriteStickersRequest) TypeName() string {
	return "getFavoriteStickers"
}

// TypeInfo returns info about TL type.
func (g *GetFavoriteStickersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getFavoriteStickers",
		ID:   GetFavoriteStickersRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetFavoriteStickersRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getFavoriteStickers#ebcbcf40 as nil")
	}
	b.PutID(GetFavoriteStickersRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetFavoriteStickersRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getFavoriteStickers#ebcbcf40 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetFavoriteStickersRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getFavoriteStickers#ebcbcf40 to nil")
	}
	if err := b.ConsumeID(GetFavoriteStickersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getFavoriteStickers#ebcbcf40: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetFavoriteStickersRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getFavoriteStickers#ebcbcf40 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetFavoriteStickersRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getFavoriteStickers#ebcbcf40 as nil")
	}
	b.ObjStart()
	b.PutID("getFavoriteStickers")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetFavoriteStickersRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getFavoriteStickers#ebcbcf40 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getFavoriteStickers"); err != nil {
				return fmt.Errorf("unable to decode getFavoriteStickers#ebcbcf40: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFavoriteStickers invokes method getFavoriteStickers#ebcbcf40 returning error if any.
func (c *Client) GetFavoriteStickers(ctx context.Context) (*Stickers, error) {
	var result Stickers

	request := &GetFavoriteStickersRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
