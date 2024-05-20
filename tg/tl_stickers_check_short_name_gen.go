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

// StickersCheckShortNameRequest represents TL type `stickers.checkShortName#284b3639`.
// Check whether the given short name is available
//
// See https://core.telegram.org/method/stickers.checkShortName for reference.
type StickersCheckShortNameRequest struct {
	// Short name
	ShortName string
}

// StickersCheckShortNameRequestTypeID is TL type id of StickersCheckShortNameRequest.
const StickersCheckShortNameRequestTypeID = 0x284b3639

// Ensuring interfaces in compile-time for StickersCheckShortNameRequest.
var (
	_ bin.Encoder     = &StickersCheckShortNameRequest{}
	_ bin.Decoder     = &StickersCheckShortNameRequest{}
	_ bin.BareEncoder = &StickersCheckShortNameRequest{}
	_ bin.BareDecoder = &StickersCheckShortNameRequest{}
)

func (c *StickersCheckShortNameRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ShortName == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *StickersCheckShortNameRequest) String() string {
	if c == nil {
		return "StickersCheckShortNameRequest(nil)"
	}
	type Alias StickersCheckShortNameRequest
	return fmt.Sprintf("StickersCheckShortNameRequest%+v", Alias(*c))
}

// FillFrom fills StickersCheckShortNameRequest from given interface.
func (c *StickersCheckShortNameRequest) FillFrom(from interface {
	GetShortName() (value string)
}) {
	c.ShortName = from.GetShortName()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickersCheckShortNameRequest) TypeID() uint32 {
	return StickersCheckShortNameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StickersCheckShortNameRequest) TypeName() string {
	return "stickers.checkShortName"
}

// TypeInfo returns info about TL type.
func (c *StickersCheckShortNameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickers.checkShortName",
		ID:   StickersCheckShortNameRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ShortName",
			SchemaName: "short_name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *StickersCheckShortNameRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stickers.checkShortName#284b3639 as nil")
	}
	b.PutID(StickersCheckShortNameRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *StickersCheckShortNameRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode stickers.checkShortName#284b3639 as nil")
	}
	b.PutString(c.ShortName)
	return nil
}

// Decode implements bin.Decoder.
func (c *StickersCheckShortNameRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stickers.checkShortName#284b3639 to nil")
	}
	if err := b.ConsumeID(StickersCheckShortNameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stickers.checkShortName#284b3639: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *StickersCheckShortNameRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode stickers.checkShortName#284b3639 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickers.checkShortName#284b3639: field short_name: %w", err)
		}
		c.ShortName = value
	}
	return nil
}

// GetShortName returns value of ShortName field.
func (c *StickersCheckShortNameRequest) GetShortName() (value string) {
	if c == nil {
		return
	}
	return c.ShortName
}

// StickersCheckShortName invokes method stickers.checkShortName#284b3639 returning error if any.
// Check whether the given short name is available
//
// Possible errors:
//
//	400 SHORT_NAME_INVALID: The specified short name is invalid.
//	400 SHORT_NAME_OCCUPIED: The specified short name is already in use.
//
// See https://core.telegram.org/method/stickers.checkShortName for reference.
func (c *Client) StickersCheckShortName(ctx context.Context, shortname string) (bool, error) {
	var result BoolBox

	request := &StickersCheckShortNameRequest{
		ShortName: shortname,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
