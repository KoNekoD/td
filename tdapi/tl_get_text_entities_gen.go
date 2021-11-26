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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// GetTextEntitiesRequest represents TL type `getTextEntities#eba543fb`.
type GetTextEntitiesRequest struct {
	// The text in which to look for entites
	Text string
}

// GetTextEntitiesRequestTypeID is TL type id of GetTextEntitiesRequest.
const GetTextEntitiesRequestTypeID = 0xeba543fb

// Ensuring interfaces in compile-time for GetTextEntitiesRequest.
var (
	_ bin.Encoder     = &GetTextEntitiesRequest{}
	_ bin.Decoder     = &GetTextEntitiesRequest{}
	_ bin.BareEncoder = &GetTextEntitiesRequest{}
	_ bin.BareDecoder = &GetTextEntitiesRequest{}
)

func (g *GetTextEntitiesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Text == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetTextEntitiesRequest) String() string {
	if g == nil {
		return "GetTextEntitiesRequest(nil)"
	}
	type Alias GetTextEntitiesRequest
	return fmt.Sprintf("GetTextEntitiesRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetTextEntitiesRequest) TypeID() uint32 {
	return GetTextEntitiesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetTextEntitiesRequest) TypeName() string {
	return "getTextEntities"
}

// TypeInfo returns info about TL type.
func (g *GetTextEntitiesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getTextEntities",
		ID:   GetTextEntitiesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetTextEntitiesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getTextEntities#eba543fb as nil")
	}
	b.PutID(GetTextEntitiesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetTextEntitiesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getTextEntities#eba543fb as nil")
	}
	b.PutString(g.Text)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetTextEntitiesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getTextEntities#eba543fb to nil")
	}
	if err := b.ConsumeID(GetTextEntitiesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getTextEntities#eba543fb: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetTextEntitiesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getTextEntities#eba543fb to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getTextEntities#eba543fb: field text: %w", err)
		}
		g.Text = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetTextEntitiesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getTextEntities#eba543fb as nil")
	}
	b.ObjStart()
	b.PutID("getTextEntities")
	b.FieldStart("text")
	b.PutString(g.Text)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetTextEntitiesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getTextEntities#eba543fb to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getTextEntities"); err != nil {
				return fmt.Errorf("unable to decode getTextEntities#eba543fb: %w", err)
			}
		case "text":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getTextEntities#eba543fb: field text: %w", err)
			}
			g.Text = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetText returns value of Text field.
func (g *GetTextEntitiesRequest) GetText() (value string) {
	return g.Text
}

// GetTextEntities invokes method getTextEntities#eba543fb returning error if any.
func (c *Client) GetTextEntities(ctx context.Context, text string) (*TextEntities, error) {
	var result TextEntities

	request := &GetTextEntitiesRequest{
		Text: text,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}