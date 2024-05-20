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

// FragmentGetCollectibleInfoRequest represents TL type `fragment.getCollectibleInfo#be1e85ba`.
//
// See https://core.telegram.org/method/fragment.getCollectibleInfo for reference.
type FragmentGetCollectibleInfoRequest struct {
	// Collectible field of FragmentGetCollectibleInfoRequest.
	Collectible InputCollectibleClass
}

// FragmentGetCollectibleInfoRequestTypeID is TL type id of FragmentGetCollectibleInfoRequest.
const FragmentGetCollectibleInfoRequestTypeID = 0xbe1e85ba

// Ensuring interfaces in compile-time for FragmentGetCollectibleInfoRequest.
var (
	_ bin.Encoder     = &FragmentGetCollectibleInfoRequest{}
	_ bin.Decoder     = &FragmentGetCollectibleInfoRequest{}
	_ bin.BareEncoder = &FragmentGetCollectibleInfoRequest{}
	_ bin.BareDecoder = &FragmentGetCollectibleInfoRequest{}
)

func (g *FragmentGetCollectibleInfoRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Collectible == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *FragmentGetCollectibleInfoRequest) String() string {
	if g == nil {
		return "FragmentGetCollectibleInfoRequest(nil)"
	}
	type Alias FragmentGetCollectibleInfoRequest
	return fmt.Sprintf("FragmentGetCollectibleInfoRequest%+v", Alias(*g))
}

// FillFrom fills FragmentGetCollectibleInfoRequest from given interface.
func (g *FragmentGetCollectibleInfoRequest) FillFrom(from interface {
	GetCollectible() (value InputCollectibleClass)
}) {
	g.Collectible = from.GetCollectible()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FragmentGetCollectibleInfoRequest) TypeID() uint32 {
	return FragmentGetCollectibleInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*FragmentGetCollectibleInfoRequest) TypeName() string {
	return "fragment.getCollectibleInfo"
}

// TypeInfo returns info about TL type.
func (g *FragmentGetCollectibleInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "fragment.getCollectibleInfo",
		ID:   FragmentGetCollectibleInfoRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Collectible",
			SchemaName: "collectible",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *FragmentGetCollectibleInfoRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode fragment.getCollectibleInfo#be1e85ba as nil")
	}
	b.PutID(FragmentGetCollectibleInfoRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *FragmentGetCollectibleInfoRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode fragment.getCollectibleInfo#be1e85ba as nil")
	}
	if g.Collectible == nil {
		return fmt.Errorf("unable to encode fragment.getCollectibleInfo#be1e85ba: field collectible is nil")
	}
	if err := g.Collectible.Encode(b); err != nil {
		return fmt.Errorf("unable to encode fragment.getCollectibleInfo#be1e85ba: field collectible: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *FragmentGetCollectibleInfoRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode fragment.getCollectibleInfo#be1e85ba to nil")
	}
	if err := b.ConsumeID(FragmentGetCollectibleInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode fragment.getCollectibleInfo#be1e85ba: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *FragmentGetCollectibleInfoRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode fragment.getCollectibleInfo#be1e85ba to nil")
	}
	{
		value, err := DecodeInputCollectible(b)
		if err != nil {
			return fmt.Errorf("unable to decode fragment.getCollectibleInfo#be1e85ba: field collectible: %w", err)
		}
		g.Collectible = value
	}
	return nil
}

// GetCollectible returns value of Collectible field.
func (g *FragmentGetCollectibleInfoRequest) GetCollectible() (value InputCollectibleClass) {
	if g == nil {
		return
	}
	return g.Collectible
}

// FragmentGetCollectibleInfo invokes method fragment.getCollectibleInfo#be1e85ba returning error if any.
//
// See https://core.telegram.org/method/fragment.getCollectibleInfo for reference.
func (c *Client) FragmentGetCollectibleInfo(ctx context.Context, collectible InputCollectibleClass) (*FragmentCollectibleInfo, error) {
	var result FragmentCollectibleInfo

	request := &FragmentGetCollectibleInfoRequest{
		Collectible: collectible,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
