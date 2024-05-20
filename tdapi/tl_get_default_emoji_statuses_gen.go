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

// GetDefaultEmojiStatusesRequest represents TL type `getDefaultEmojiStatuses#24e45ec3`.
type GetDefaultEmojiStatusesRequest struct {
}

// GetDefaultEmojiStatusesRequestTypeID is TL type id of GetDefaultEmojiStatusesRequest.
const GetDefaultEmojiStatusesRequestTypeID = 0x24e45ec3

// Ensuring interfaces in compile-time for GetDefaultEmojiStatusesRequest.
var (
	_ bin.Encoder     = &GetDefaultEmojiStatusesRequest{}
	_ bin.Decoder     = &GetDefaultEmojiStatusesRequest{}
	_ bin.BareEncoder = &GetDefaultEmojiStatusesRequest{}
	_ bin.BareDecoder = &GetDefaultEmojiStatusesRequest{}
)

func (g *GetDefaultEmojiStatusesRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetDefaultEmojiStatusesRequest) String() string {
	if g == nil {
		return "GetDefaultEmojiStatusesRequest(nil)"
	}
	type Alias GetDefaultEmojiStatusesRequest
	return fmt.Sprintf("GetDefaultEmojiStatusesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetDefaultEmojiStatusesRequest) TypeID() uint32 {
	return GetDefaultEmojiStatusesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetDefaultEmojiStatusesRequest) TypeName() string {
	return "getDefaultEmojiStatuses"
}

// TypeInfo returns info about TL type.
func (g *GetDefaultEmojiStatusesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getDefaultEmojiStatuses",
		ID:   GetDefaultEmojiStatusesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetDefaultEmojiStatusesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDefaultEmojiStatuses#24e45ec3 as nil")
	}
	b.PutID(GetDefaultEmojiStatusesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetDefaultEmojiStatusesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getDefaultEmojiStatuses#24e45ec3 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetDefaultEmojiStatusesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDefaultEmojiStatuses#24e45ec3 to nil")
	}
	if err := b.ConsumeID(GetDefaultEmojiStatusesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getDefaultEmojiStatuses#24e45ec3: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetDefaultEmojiStatusesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getDefaultEmojiStatuses#24e45ec3 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetDefaultEmojiStatusesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getDefaultEmojiStatuses#24e45ec3 as nil")
	}
	b.ObjStart()
	b.PutID("getDefaultEmojiStatuses")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetDefaultEmojiStatusesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getDefaultEmojiStatuses#24e45ec3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getDefaultEmojiStatuses"); err != nil {
				return fmt.Errorf("unable to decode getDefaultEmojiStatuses#24e45ec3: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDefaultEmojiStatuses invokes method getDefaultEmojiStatuses#24e45ec3 returning error if any.
func (c *Client) GetDefaultEmojiStatuses(ctx context.Context) (*EmojiStatuses, error) {
	var result EmojiStatuses

	request := &GetDefaultEmojiStatusesRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
