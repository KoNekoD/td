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

// GetContactsRequest represents TL type `getContacts#ab7f4470`.
type GetContactsRequest struct {
}

// GetContactsRequestTypeID is TL type id of GetContactsRequest.
const GetContactsRequestTypeID = 0xab7f4470

// Ensuring interfaces in compile-time for GetContactsRequest.
var (
	_ bin.Encoder     = &GetContactsRequest{}
	_ bin.Decoder     = &GetContactsRequest{}
	_ bin.BareEncoder = &GetContactsRequest{}
	_ bin.BareDecoder = &GetContactsRequest{}
)

func (g *GetContactsRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetContactsRequest) String() string {
	if g == nil {
		return "GetContactsRequest(nil)"
	}
	type Alias GetContactsRequest
	return fmt.Sprintf("GetContactsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetContactsRequest) TypeID() uint32 {
	return GetContactsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetContactsRequest) TypeName() string {
	return "getContacts"
}

// TypeInfo returns info about TL type.
func (g *GetContactsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getContacts",
		ID:   GetContactsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetContactsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getContacts#ab7f4470 as nil")
	}
	b.PutID(GetContactsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetContactsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getContacts#ab7f4470 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetContactsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getContacts#ab7f4470 to nil")
	}
	if err := b.ConsumeID(GetContactsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getContacts#ab7f4470: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetContactsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getContacts#ab7f4470 to nil")
	}
	return nil
}

// EncodeTDLibJSON encodes g in TDLib API JSON format.
func (g *GetContactsRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getContacts#ab7f4470 as nil")
	}
	b.ObjStart()
	b.PutID("getContacts")
	b.ObjEnd()
	return nil
}

// GetContacts invokes method getContacts#ab7f4470 returning error if any.
func (c *Client) GetContacts(ctx context.Context) (*Users, error) {
	var result Users

	request := &GetContactsRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
