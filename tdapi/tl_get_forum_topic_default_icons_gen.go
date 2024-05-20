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

// GetForumTopicDefaultIconsRequest represents TL type `getForumTopicDefaultIcons#583574dc`.
type GetForumTopicDefaultIconsRequest struct {
}

// GetForumTopicDefaultIconsRequestTypeID is TL type id of GetForumTopicDefaultIconsRequest.
const GetForumTopicDefaultIconsRequestTypeID = 0x583574dc

// Ensuring interfaces in compile-time for GetForumTopicDefaultIconsRequest.
var (
	_ bin.Encoder     = &GetForumTopicDefaultIconsRequest{}
	_ bin.Decoder     = &GetForumTopicDefaultIconsRequest{}
	_ bin.BareEncoder = &GetForumTopicDefaultIconsRequest{}
	_ bin.BareDecoder = &GetForumTopicDefaultIconsRequest{}
)

func (g *GetForumTopicDefaultIconsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetForumTopicDefaultIconsRequest) String() string {
	if g == nil {
		return "GetForumTopicDefaultIconsRequest(nil)"
	}
	type Alias GetForumTopicDefaultIconsRequest
	return fmt.Sprintf("GetForumTopicDefaultIconsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetForumTopicDefaultIconsRequest) TypeID() uint32 {
	return GetForumTopicDefaultIconsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetForumTopicDefaultIconsRequest) TypeName() string {
	return "getForumTopicDefaultIcons"
}

// TypeInfo returns info about TL type.
func (g *GetForumTopicDefaultIconsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getForumTopicDefaultIcons",
		ID:   GetForumTopicDefaultIconsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetForumTopicDefaultIconsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getForumTopicDefaultIcons#583574dc as nil")
	}
	b.PutID(GetForumTopicDefaultIconsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetForumTopicDefaultIconsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getForumTopicDefaultIcons#583574dc as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetForumTopicDefaultIconsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getForumTopicDefaultIcons#583574dc to nil")
	}
	if err := b.ConsumeID(GetForumTopicDefaultIconsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getForumTopicDefaultIcons#583574dc: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetForumTopicDefaultIconsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getForumTopicDefaultIcons#583574dc to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetForumTopicDefaultIconsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getForumTopicDefaultIcons#583574dc as nil")
	}
	b.ObjStart()
	b.PutID("getForumTopicDefaultIcons")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetForumTopicDefaultIconsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getForumTopicDefaultIcons#583574dc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getForumTopicDefaultIcons"); err != nil {
				return fmt.Errorf("unable to decode getForumTopicDefaultIcons#583574dc: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetForumTopicDefaultIcons invokes method getForumTopicDefaultIcons#583574dc returning error if any.
func (c *Client) GetForumTopicDefaultIcons(ctx context.Context) (*Stickers, error) {
	var result Stickers

	request := &GetForumTopicDefaultIconsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
