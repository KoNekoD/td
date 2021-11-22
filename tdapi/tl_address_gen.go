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

// Address represents TL type `address#86304f3a`.
type Address struct {
	// A two-letter ISO 3166-1 alpha-2 country code
	CountryCode string
	// State, if applicable
	State string
	// City
	City string
	// First line of the address
	StreetLine1 string
	// Second line of the address
	StreetLine2 string
	// Address postal code
	PostalCode string
}

// AddressTypeID is TL type id of Address.
const AddressTypeID = 0x86304f3a

// Ensuring interfaces in compile-time for Address.
var (
	_ bin.Encoder     = &Address{}
	_ bin.Decoder     = &Address{}
	_ bin.BareEncoder = &Address{}
	_ bin.BareDecoder = &Address{}
)

func (a *Address) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.CountryCode == "") {
		return false
	}
	if !(a.State == "") {
		return false
	}
	if !(a.City == "") {
		return false
	}
	if !(a.StreetLine1 == "") {
		return false
	}
	if !(a.StreetLine2 == "") {
		return false
	}
	if !(a.PostalCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *Address) String() string {
	if a == nil {
		return "Address(nil)"
	}
	type Alias Address
	return fmt.Sprintf("Address%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Address) TypeID() uint32 {
	return AddressTypeID
}

// TypeName returns name of type in TL schema.
func (*Address) TypeName() string {
	return "address"
}

// TypeInfo returns info about TL type.
func (a *Address) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "address",
		ID:   AddressTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CountryCode",
			SchemaName: "country_code",
		},
		{
			Name:       "State",
			SchemaName: "state",
		},
		{
			Name:       "City",
			SchemaName: "city",
		},
		{
			Name:       "StreetLine1",
			SchemaName: "street_line1",
		},
		{
			Name:       "StreetLine2",
			SchemaName: "street_line2",
		},
		{
			Name:       "PostalCode",
			SchemaName: "postal_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *Address) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode address#86304f3a as nil")
	}
	b.PutID(AddressTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *Address) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode address#86304f3a as nil")
	}
	b.PutString(a.CountryCode)
	b.PutString(a.State)
	b.PutString(a.City)
	b.PutString(a.StreetLine1)
	b.PutString(a.StreetLine2)
	b.PutString(a.PostalCode)
	return nil
}

// Decode implements bin.Decoder.
func (a *Address) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode address#86304f3a to nil")
	}
	if err := b.ConsumeID(AddressTypeID); err != nil {
		return fmt.Errorf("unable to decode address#86304f3a: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *Address) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode address#86304f3a to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode address#86304f3a: field country_code: %w", err)
		}
		a.CountryCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode address#86304f3a: field state: %w", err)
		}
		a.State = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode address#86304f3a: field city: %w", err)
		}
		a.City = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode address#86304f3a: field street_line1: %w", err)
		}
		a.StreetLine1 = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode address#86304f3a: field street_line2: %w", err)
		}
		a.StreetLine2 = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode address#86304f3a: field postal_code: %w", err)
		}
		a.PostalCode = value
	}
	return nil
}

// EncodeTDLibJSON encodes a in TDLib API JSON format.
func (a *Address) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode address#86304f3a as nil")
	}
	b.ObjStart()
	b.PutID("address")
	b.FieldStart("country_code")
	b.PutString(a.CountryCode)
	b.FieldStart("state")
	b.PutString(a.State)
	b.FieldStart("city")
	b.PutString(a.City)
	b.FieldStart("street_line1")
	b.PutString(a.StreetLine1)
	b.FieldStart("street_line2")
	b.PutString(a.StreetLine2)
	b.FieldStart("postal_code")
	b.PutString(a.PostalCode)
	b.ObjEnd()
	return nil
}

// GetCountryCode returns value of CountryCode field.
func (a *Address) GetCountryCode() (value string) {
	return a.CountryCode
}

// GetState returns value of State field.
func (a *Address) GetState() (value string) {
	return a.State
}

// GetCity returns value of City field.
func (a *Address) GetCity() (value string) {
	return a.City
}

// GetStreetLine1 returns value of StreetLine1 field.
func (a *Address) GetStreetLine1() (value string) {
	return a.StreetLine1
}

// GetStreetLine2 returns value of StreetLine2 field.
func (a *Address) GetStreetLine2() (value string) {
	return a.StreetLine2
}

// GetPostalCode returns value of PostalCode field.
func (a *Address) GetPostalCode() (value string) {
	return a.PostalCode
}
