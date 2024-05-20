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

// GetLogVerbosityLevelRequest represents TL type `getLogVerbosityLevel#23689ae4`.
type GetLogVerbosityLevelRequest struct {
}

// GetLogVerbosityLevelRequestTypeID is TL type id of GetLogVerbosityLevelRequest.
const GetLogVerbosityLevelRequestTypeID = 0x23689ae4

// Ensuring interfaces in compile-time for GetLogVerbosityLevelRequest.
var (
	_ bin.Encoder     = &GetLogVerbosityLevelRequest{}
	_ bin.Decoder     = &GetLogVerbosityLevelRequest{}
	_ bin.BareEncoder = &GetLogVerbosityLevelRequest{}
	_ bin.BareDecoder = &GetLogVerbosityLevelRequest{}
)

func (g *GetLogVerbosityLevelRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetLogVerbosityLevelRequest) String() string {
	if g == nil {
		return "GetLogVerbosityLevelRequest(nil)"
	}
	type Alias GetLogVerbosityLevelRequest
	return fmt.Sprintf("GetLogVerbosityLevelRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetLogVerbosityLevelRequest) TypeID() uint32 {
	return GetLogVerbosityLevelRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetLogVerbosityLevelRequest) TypeName() string {
	return "getLogVerbosityLevel"
}

// TypeInfo returns info about TL type.
func (g *GetLogVerbosityLevelRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getLogVerbosityLevel",
		ID:   GetLogVerbosityLevelRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetLogVerbosityLevelRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogVerbosityLevel#23689ae4 as nil")
	}
	b.PutID(GetLogVerbosityLevelRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetLogVerbosityLevelRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogVerbosityLevel#23689ae4 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetLogVerbosityLevelRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogVerbosityLevel#23689ae4 to nil")
	}
	if err := b.ConsumeID(GetLogVerbosityLevelRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getLogVerbosityLevel#23689ae4: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetLogVerbosityLevelRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogVerbosityLevel#23689ae4 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetLogVerbosityLevelRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getLogVerbosityLevel#23689ae4 as nil")
	}
	b.ObjStart()
	b.PutID("getLogVerbosityLevel")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetLogVerbosityLevelRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getLogVerbosityLevel#23689ae4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getLogVerbosityLevel"); err != nil {
				return fmt.Errorf("unable to decode getLogVerbosityLevel#23689ae4: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLogVerbosityLevel invokes method getLogVerbosityLevel#23689ae4 returning error if any.
func (c *Client) GetLogVerbosityLevel(ctx context.Context) (*LogVerbosityLevel, error) {
	var result LogVerbosityLevel

	request := &GetLogVerbosityLevelRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
