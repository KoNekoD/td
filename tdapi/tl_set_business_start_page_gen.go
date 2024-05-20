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

// SetBusinessStartPageRequest represents TL type `setBusinessStartPage#9eed499e`.
type SetBusinessStartPageRequest struct {
	// The new start page of the business; pass null to remove custom start page
	StartPage InputBusinessStartPage
}

// SetBusinessStartPageRequestTypeID is TL type id of SetBusinessStartPageRequest.
const SetBusinessStartPageRequestTypeID = 0x9eed499e

// Ensuring interfaces in compile-time for SetBusinessStartPageRequest.
var (
	_ bin.Encoder     = &SetBusinessStartPageRequest{}
	_ bin.Decoder     = &SetBusinessStartPageRequest{}
	_ bin.BareEncoder = &SetBusinessStartPageRequest{}
	_ bin.BareDecoder = &SetBusinessStartPageRequest{}
)

func (s *SetBusinessStartPageRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.StartPage.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBusinessStartPageRequest) String() string {
	if s == nil {
		return "SetBusinessStartPageRequest(nil)"
	}
	type Alias SetBusinessStartPageRequest
	return fmt.Sprintf("SetBusinessStartPageRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBusinessStartPageRequest) TypeID() uint32 {
	return SetBusinessStartPageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBusinessStartPageRequest) TypeName() string {
	return "setBusinessStartPage"
}

// TypeInfo returns info about TL type.
func (s *SetBusinessStartPageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBusinessStartPage",
		ID:   SetBusinessStartPageRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StartPage",
			SchemaName: "start_page",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBusinessStartPageRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBusinessStartPage#9eed499e as nil")
	}
	b.PutID(SetBusinessStartPageRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBusinessStartPageRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBusinessStartPage#9eed499e as nil")
	}
	if err := s.StartPage.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setBusinessStartPage#9eed499e: field start_page: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBusinessStartPageRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBusinessStartPage#9eed499e to nil")
	}
	if err := b.ConsumeID(SetBusinessStartPageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBusinessStartPage#9eed499e: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBusinessStartPageRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBusinessStartPage#9eed499e to nil")
	}
	{
		if err := s.StartPage.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setBusinessStartPage#9eed499e: field start_page: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetBusinessStartPageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setBusinessStartPage#9eed499e as nil")
	}
	b.ObjStart()
	b.PutID("setBusinessStartPage")
	b.Comma()
	b.FieldStart("start_page")
	if err := s.StartPage.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setBusinessStartPage#9eed499e: field start_page: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetBusinessStartPageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setBusinessStartPage#9eed499e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setBusinessStartPage"); err != nil {
				return fmt.Errorf("unable to decode setBusinessStartPage#9eed499e: %w", err)
			}
		case "start_page":
			if err := s.StartPage.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setBusinessStartPage#9eed499e: field start_page: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStartPage returns value of StartPage field.
func (s *SetBusinessStartPageRequest) GetStartPage() (value InputBusinessStartPage) {
	if s == nil {
		return
	}
	return s.StartPage
}

// SetBusinessStartPage invokes method setBusinessStartPage#9eed499e returning error if any.
func (c *Client) SetBusinessStartPage(ctx context.Context, startpage InputBusinessStartPage) error {
	var ok Ok

	request := &SetBusinessStartPageRequest{
		StartPage: startpage,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
