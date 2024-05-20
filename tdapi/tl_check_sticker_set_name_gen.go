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

// CheckStickerSetNameRequest represents TL type `checkStickerSetName#955808fe`.
type CheckStickerSetNameRequest struct {
	// Name to be checked
	Name string
}

// CheckStickerSetNameRequestTypeID is TL type id of CheckStickerSetNameRequest.
const CheckStickerSetNameRequestTypeID = 0x955808fe

// Ensuring interfaces in compile-time for CheckStickerSetNameRequest.
var (
	_ bin.Encoder     = &CheckStickerSetNameRequest{}
	_ bin.Decoder     = &CheckStickerSetNameRequest{}
	_ bin.BareEncoder = &CheckStickerSetNameRequest{}
	_ bin.BareDecoder = &CheckStickerSetNameRequest{}
)

func (c *CheckStickerSetNameRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckStickerSetNameRequest) String() string {
	if c == nil {
		return "CheckStickerSetNameRequest(nil)"
	}
	type Alias CheckStickerSetNameRequest
	return fmt.Sprintf("CheckStickerSetNameRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckStickerSetNameRequest) TypeID() uint32 {
	return CheckStickerSetNameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckStickerSetNameRequest) TypeName() string {
	return "checkStickerSetName"
}

// TypeInfo returns info about TL type.
func (c *CheckStickerSetNameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkStickerSetName",
		ID:   CheckStickerSetNameRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckStickerSetNameRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetName#955808fe as nil")
	}
	b.PutID(CheckStickerSetNameRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckStickerSetNameRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetName#955808fe as nil")
	}
	b.PutString(c.Name)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckStickerSetNameRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetName#955808fe to nil")
	}
	if err := b.ConsumeID(CheckStickerSetNameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkStickerSetName#955808fe: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckStickerSetNameRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetName#955808fe to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode checkStickerSetName#955808fe: field name: %w", err)
		}
		c.Name = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckStickerSetNameRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkStickerSetName#955808fe as nil")
	}
	b.ObjStart()
	b.PutID("checkStickerSetName")
	b.Comma()
	b.FieldStart("name")
	b.PutString(c.Name)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckStickerSetNameRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkStickerSetName#955808fe to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkStickerSetName"); err != nil {
				return fmt.Errorf("unable to decode checkStickerSetName#955808fe: %w", err)
			}
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode checkStickerSetName#955808fe: field name: %w", err)
			}
			c.Name = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetName returns value of Name field.
func (c *CheckStickerSetNameRequest) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

// CheckStickerSetName invokes method checkStickerSetName#955808fe returning error if any.
func (c *Client) CheckStickerSetName(ctx context.Context, name string) (CheckStickerSetNameResultClass, error) {
	var result CheckStickerSetNameResultBox

	request := &CheckStickerSetNameRequest{
		Name: name,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.CheckStickerSetNameResult, nil
}
