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

// GetSupportNameRequest represents TL type `getSupportName#4d9e1562`.
type GetSupportNameRequest struct {
}

// GetSupportNameRequestTypeID is TL type id of GetSupportNameRequest.
const GetSupportNameRequestTypeID = 0x4d9e1562

// Ensuring interfaces in compile-time for GetSupportNameRequest.
var (
	_ bin.Encoder     = &GetSupportNameRequest{}
	_ bin.Decoder     = &GetSupportNameRequest{}
	_ bin.BareEncoder = &GetSupportNameRequest{}
	_ bin.BareDecoder = &GetSupportNameRequest{}
)

func (g *GetSupportNameRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSupportNameRequest) String() string {
	if g == nil {
		return "GetSupportNameRequest(nil)"
	}
	type Alias GetSupportNameRequest
	return fmt.Sprintf("GetSupportNameRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSupportNameRequest) TypeID() uint32 {
	return GetSupportNameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSupportNameRequest) TypeName() string {
	return "getSupportName"
}

// TypeInfo returns info about TL type.
func (g *GetSupportNameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSupportName",
		ID:   GetSupportNameRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSupportNameRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupportName#4d9e1562 as nil")
	}
	b.PutID(GetSupportNameRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSupportNameRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupportName#4d9e1562 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSupportNameRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupportName#4d9e1562 to nil")
	}
	if err := b.ConsumeID(GetSupportNameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSupportName#4d9e1562: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSupportNameRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupportName#4d9e1562 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSupportNameRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSupportName#4d9e1562 as nil")
	}
	b.ObjStart()
	b.PutID("getSupportName")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSupportNameRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSupportName#4d9e1562 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSupportName"); err != nil {
				return fmt.Errorf("unable to decode getSupportName#4d9e1562: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupportName invokes method getSupportName#4d9e1562 returning error if any.
func (c *Client) GetSupportName(ctx context.Context) (*Text, error) {
	var result Text

	request := &GetSupportNameRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
