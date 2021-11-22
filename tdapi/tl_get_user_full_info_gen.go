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

// GetUserFullInfoRequest represents TL type `getUserFullInfo#d8eebac1`.
type GetUserFullInfoRequest struct {
	// User identifier
	UserID int32
}

// GetUserFullInfoRequestTypeID is TL type id of GetUserFullInfoRequest.
const GetUserFullInfoRequestTypeID = 0xd8eebac1

// Ensuring interfaces in compile-time for GetUserFullInfoRequest.
var (
	_ bin.Encoder     = &GetUserFullInfoRequest{}
	_ bin.Decoder     = &GetUserFullInfoRequest{}
	_ bin.BareEncoder = &GetUserFullInfoRequest{}
	_ bin.BareDecoder = &GetUserFullInfoRequest{}
)

func (g *GetUserFullInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetUserFullInfoRequest) String() string {
	if g == nil {
		return "GetUserFullInfoRequest(nil)"
	}
	type Alias GetUserFullInfoRequest
	return fmt.Sprintf("GetUserFullInfoRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetUserFullInfoRequest) TypeID() uint32 {
	return GetUserFullInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetUserFullInfoRequest) TypeName() string {
	return "getUserFullInfo"
}

// TypeInfo returns info about TL type.
func (g *GetUserFullInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getUserFullInfo",
		ID:   GetUserFullInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetUserFullInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserFullInfo#d8eebac1 as nil")
	}
	b.PutID(GetUserFullInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetUserFullInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserFullInfo#d8eebac1 as nil")
	}
	b.PutInt32(g.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetUserFullInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserFullInfo#d8eebac1 to nil")
	}
	if err := b.ConsumeID(GetUserFullInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getUserFullInfo#d8eebac1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetUserFullInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUserFullInfo#d8eebac1 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getUserFullInfo#d8eebac1: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// EncodeTDLibJSON encodes g in TDLib API JSON format.
func (g *GetUserFullInfoRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getUserFullInfo#d8eebac1 as nil")
	}
	b.ObjStart()
	b.PutID("getUserFullInfo")
	b.FieldStart("user_id")
	b.PutInt32(g.UserID)
	b.ObjEnd()
	return nil
}

// GetUserID returns value of UserID field.
func (g *GetUserFullInfoRequest) GetUserID() (value int32) {
	return g.UserID
}

// GetUserFullInfo invokes method getUserFullInfo#d8eebac1 returning error if any.
func (c *Client) GetUserFullInfo(ctx context.Context, userid int32) (*UserFullInfo, error) {
	var result UserFullInfo

	request := &GetUserFullInfoRequest{
		UserID: userid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
