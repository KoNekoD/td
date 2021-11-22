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

// TermsOfService represents TL type `termsOfService#2c12b185`.
type TermsOfService struct {
	// Text of the terms of service
	Text FormattedText
	// The minimum age of a user to be able to accept the terms; 0 if any
	MinUserAge int32
	// True, if a blocking popup with terms of service must be shown to the user
	ShowPopup bool
}

// TermsOfServiceTypeID is TL type id of TermsOfService.
const TermsOfServiceTypeID = 0x2c12b185

// Ensuring interfaces in compile-time for TermsOfService.
var (
	_ bin.Encoder     = &TermsOfService{}
	_ bin.Decoder     = &TermsOfService{}
	_ bin.BareEncoder = &TermsOfService{}
	_ bin.BareDecoder = &TermsOfService{}
)

func (t *TermsOfService) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Text.Zero()) {
		return false
	}
	if !(t.MinUserAge == 0) {
		return false
	}
	if !(t.ShowPopup == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TermsOfService) String() string {
	if t == nil {
		return "TermsOfService(nil)"
	}
	type Alias TermsOfService
	return fmt.Sprintf("TermsOfService%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TermsOfService) TypeID() uint32 {
	return TermsOfServiceTypeID
}

// TypeName returns name of type in TL schema.
func (*TermsOfService) TypeName() string {
	return "termsOfService"
}

// TypeInfo returns info about TL type.
func (t *TermsOfService) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "termsOfService",
		ID:   TermsOfServiceTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "MinUserAge",
			SchemaName: "min_user_age",
		},
		{
			Name:       "ShowPopup",
			SchemaName: "show_popup",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TermsOfService) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode termsOfService#2c12b185 as nil")
	}
	b.PutID(TermsOfServiceTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TermsOfService) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode termsOfService#2c12b185 as nil")
	}
	if err := t.Text.Encode(b); err != nil {
		return fmt.Errorf("unable to encode termsOfService#2c12b185: field text: %w", err)
	}
	b.PutInt32(t.MinUserAge)
	b.PutBool(t.ShowPopup)
	return nil
}

// Decode implements bin.Decoder.
func (t *TermsOfService) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode termsOfService#2c12b185 to nil")
	}
	if err := b.ConsumeID(TermsOfServiceTypeID); err != nil {
		return fmt.Errorf("unable to decode termsOfService#2c12b185: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TermsOfService) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode termsOfService#2c12b185 to nil")
	}
	{
		if err := t.Text.Decode(b); err != nil {
			return fmt.Errorf("unable to decode termsOfService#2c12b185: field text: %w", err)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode termsOfService#2c12b185: field min_user_age: %w", err)
		}
		t.MinUserAge = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode termsOfService#2c12b185: field show_popup: %w", err)
		}
		t.ShowPopup = value
	}
	return nil
}

// EncodeTDLibJSON encodes t in TDLib API JSON format.
func (t *TermsOfService) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode termsOfService#2c12b185 as nil")
	}
	b.ObjStart()
	b.PutID("termsOfService")
	b.FieldStart("text")
	if err := t.Text.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode termsOfService#2c12b185: field text: %w", err)
	}
	b.FieldStart("min_user_age")
	b.PutInt32(t.MinUserAge)
	b.FieldStart("show_popup")
	b.PutBool(t.ShowPopup)
	b.ObjEnd()
	return nil
}

// GetText returns value of Text field.
func (t *TermsOfService) GetText() (value FormattedText) {
	return t.Text
}

// GetMinUserAge returns value of MinUserAge field.
func (t *TermsOfService) GetMinUserAge() (value int32) {
	return t.MinUserAge
}

// GetShowPopup returns value of ShowPopup field.
func (t *TermsOfService) GetShowPopup() (value bool) {
	return t.ShowPopup
}
