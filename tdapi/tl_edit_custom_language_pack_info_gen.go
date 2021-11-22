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

// EditCustomLanguagePackInfoRequest represents TL type `editCustomLanguagePackInfo#4eb91099`.
type EditCustomLanguagePackInfoRequest struct {
	// New information about the custom local language pack
	Info LanguagePackInfo
}

// EditCustomLanguagePackInfoRequestTypeID is TL type id of EditCustomLanguagePackInfoRequest.
const EditCustomLanguagePackInfoRequestTypeID = 0x4eb91099

// Ensuring interfaces in compile-time for EditCustomLanguagePackInfoRequest.
var (
	_ bin.Encoder     = &EditCustomLanguagePackInfoRequest{}
	_ bin.Decoder     = &EditCustomLanguagePackInfoRequest{}
	_ bin.BareEncoder = &EditCustomLanguagePackInfoRequest{}
	_ bin.BareDecoder = &EditCustomLanguagePackInfoRequest{}
)

func (e *EditCustomLanguagePackInfoRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Info.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditCustomLanguagePackInfoRequest) String() string {
	if e == nil {
		return "EditCustomLanguagePackInfoRequest(nil)"
	}
	type Alias EditCustomLanguagePackInfoRequest
	return fmt.Sprintf("EditCustomLanguagePackInfoRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditCustomLanguagePackInfoRequest) TypeID() uint32 {
	return EditCustomLanguagePackInfoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditCustomLanguagePackInfoRequest) TypeName() string {
	return "editCustomLanguagePackInfo"
}

// TypeInfo returns info about TL type.
func (e *EditCustomLanguagePackInfoRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editCustomLanguagePackInfo",
		ID:   EditCustomLanguagePackInfoRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Info",
			SchemaName: "info",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditCustomLanguagePackInfoRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editCustomLanguagePackInfo#4eb91099 as nil")
	}
	b.PutID(EditCustomLanguagePackInfoRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditCustomLanguagePackInfoRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editCustomLanguagePackInfo#4eb91099 as nil")
	}
	if err := e.Info.Encode(b); err != nil {
		return fmt.Errorf("unable to encode editCustomLanguagePackInfo#4eb91099: field info: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EditCustomLanguagePackInfoRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editCustomLanguagePackInfo#4eb91099 to nil")
	}
	if err := b.ConsumeID(EditCustomLanguagePackInfoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editCustomLanguagePackInfo#4eb91099: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditCustomLanguagePackInfoRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editCustomLanguagePackInfo#4eb91099 to nil")
	}
	{
		if err := e.Info.Decode(b); err != nil {
			return fmt.Errorf("unable to decode editCustomLanguagePackInfo#4eb91099: field info: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON encodes e in TDLib API JSON format.
func (e *EditCustomLanguagePackInfoRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editCustomLanguagePackInfo#4eb91099 as nil")
	}
	b.ObjStart()
	b.PutID("editCustomLanguagePackInfo")
	b.FieldStart("info")
	if err := e.Info.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode editCustomLanguagePackInfo#4eb91099: field info: %w", err)
	}
	b.ObjEnd()
	return nil
}

// GetInfo returns value of Info field.
func (e *EditCustomLanguagePackInfoRequest) GetInfo() (value LanguagePackInfo) {
	return e.Info
}

// EditCustomLanguagePackInfo invokes method editCustomLanguagePackInfo#4eb91099 returning error if any.
func (c *Client) EditCustomLanguagePackInfo(ctx context.Context, info LanguagePackInfo) error {
	var ok Ok

	request := &EditCustomLanguagePackInfoRequest{
		Info: info,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
