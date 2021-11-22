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

// GetUserRequest represents TL type `getUser#fd29e51f`.
type GetUserRequest struct {
	// User identifier
	UserID int32
}

// GetUserRequestTypeID is TL type id of GetUserRequest.
const GetUserRequestTypeID = 0xfd29e51f

// Ensuring interfaces in compile-time for GetUserRequest.
var (
	_ bin.Encoder     = &GetUserRequest{}
	_ bin.Decoder     = &GetUserRequest{}
	_ bin.BareEncoder = &GetUserRequest{}
	_ bin.BareDecoder = &GetUserRequest{}
)

func (g *GetUserRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetUserRequest) String() string {
	if g == nil {
		return "GetUserRequest(nil)"
	}
	type Alias GetUserRequest
	return fmt.Sprintf("GetUserRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetUserRequest) TypeID() uint32 {
	return GetUserRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetUserRequest) TypeName() string {
	return "getUser"
}

// TypeInfo returns info about TL type.
func (g *GetUserRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getUser",
		ID:   GetUserRequestTypeID,
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
func (g *GetUserRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUser#fd29e51f as nil")
	}
	b.PutID(GetUserRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetUserRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getUser#fd29e51f as nil")
	}
	b.PutInt32(g.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetUserRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUser#fd29e51f to nil")
	}
	if err := b.ConsumeID(GetUserRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getUser#fd29e51f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetUserRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getUser#fd29e51f to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getUser#fd29e51f: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// EncodeTDLibJSON encodes g in TDLib API JSON format.
func (g *GetUserRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getUser#fd29e51f as nil")
	}
	b.ObjStart()
	b.PutID("getUser")
	b.FieldStart("user_id")
	b.PutInt32(g.UserID)
	b.ObjEnd()
	return nil
}

// GetUserID returns value of UserID field.
func (g *GetUserRequest) GetUserID() (value int32) {
	return g.UserID
}

// GetUser invokes method getUser#fd29e51f returning error if any.
func (c *Client) GetUser(ctx context.Context, userid int32) (*User, error) {
	var result User

	request := &GetUserRequest{
		UserID: userid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
