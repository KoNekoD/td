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

// GetLogTagsRequest represents TL type `getLogTags#f0d569da`.
type GetLogTagsRequest struct {
}

// GetLogTagsRequestTypeID is TL type id of GetLogTagsRequest.
const GetLogTagsRequestTypeID = 0xf0d569da

// Ensuring interfaces in compile-time for GetLogTagsRequest.
var (
	_ bin.Encoder     = &GetLogTagsRequest{}
	_ bin.Decoder     = &GetLogTagsRequest{}
	_ bin.BareEncoder = &GetLogTagsRequest{}
	_ bin.BareDecoder = &GetLogTagsRequest{}
)

func (g *GetLogTagsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetLogTagsRequest) String() string {
	if g == nil {
		return "GetLogTagsRequest(nil)"
	}
	type Alias GetLogTagsRequest
	return fmt.Sprintf("GetLogTagsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetLogTagsRequest) TypeID() uint32 {
	return GetLogTagsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetLogTagsRequest) TypeName() string {
	return "getLogTags"
}

// TypeInfo returns info about TL type.
func (g *GetLogTagsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getLogTags",
		ID:   GetLogTagsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetLogTagsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogTags#f0d569da as nil")
	}
	b.PutID(GetLogTagsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetLogTagsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogTags#f0d569da as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetLogTagsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogTags#f0d569da to nil")
	}
	if err := b.ConsumeID(GetLogTagsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getLogTags#f0d569da: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetLogTagsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogTags#f0d569da to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetLogTagsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogTags#f0d569da as nil")
	}
	b.ObjStart()
	b.PutID("getLogTags")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetLogTagsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogTags#f0d569da to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getLogTags"); err != nil {
				return fmt.Errorf("unable to decode getLogTags#f0d569da: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLogTags invokes method getLogTags#f0d569da returning error if any.
func (c *Client) GetLogTags(ctx context.Context) (*LogTags, error) {
	var result LogTags

	request := &GetLogTagsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
