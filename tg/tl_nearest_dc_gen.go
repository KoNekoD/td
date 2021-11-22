// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// NearestDC represents TL type `nearestDc#8e1a1775`.
// Nearest data centre, according to geo-ip.
//
// See https://core.telegram.org/constructor/nearestDc for reference.
type NearestDC struct {
	// Country code determined by geo-ip
	Country string
	// Number of current data centre
	ThisDC int
	// Number of nearest data centre
	NearestDC int
}

// NearestDCTypeID is TL type id of NearestDC.
const NearestDCTypeID = 0x8e1a1775

// Ensuring interfaces in compile-time for NearestDC.
var (
	_ bin.Encoder     = &NearestDC{}
	_ bin.Decoder     = &NearestDC{}
	_ bin.BareEncoder = &NearestDC{}
	_ bin.BareDecoder = &NearestDC{}
)

func (n *NearestDC) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.Country == "") {
		return false
	}
	if !(n.ThisDC == 0) {
		return false
	}
	if !(n.NearestDC == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NearestDC) String() string {
	if n == nil {
		return "NearestDC(nil)"
	}
	type Alias NearestDC
	return fmt.Sprintf("NearestDC%+v", Alias(*n))
}

// FillFrom fills NearestDC from given interface.
func (n *NearestDC) FillFrom(from interface {
	GetCountry() (value string)
	GetThisDC() (value int)
	GetNearestDC() (value int)
}) {
	n.Country = from.GetCountry()
	n.ThisDC = from.GetThisDC()
	n.NearestDC = from.GetNearestDC()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NearestDC) TypeID() uint32 {
	return NearestDCTypeID
}

// TypeName returns name of type in TL schema.
func (*NearestDC) TypeName() string {
	return "nearestDc"
}

// TypeInfo returns info about TL type.
func (n *NearestDC) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "nearestDc",
		ID:   NearestDCTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Country",
			SchemaName: "country",
		},
		{
			Name:       "ThisDC",
			SchemaName: "this_dc",
		},
		{
			Name:       "NearestDC",
			SchemaName: "nearest_dc",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NearestDC) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode nearestDc#8e1a1775 as nil")
	}
	b.PutID(NearestDCTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NearestDC) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode nearestDc#8e1a1775 as nil")
	}
	b.PutString(n.Country)
	b.PutInt(n.ThisDC)
	b.PutInt(n.NearestDC)
	return nil
}

// Decode implements bin.Decoder.
func (n *NearestDC) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode nearestDc#8e1a1775 to nil")
	}
	if err := b.ConsumeID(NearestDCTypeID); err != nil {
		return fmt.Errorf("unable to decode nearestDc#8e1a1775: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NearestDC) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode nearestDc#8e1a1775 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode nearestDc#8e1a1775: field country: %w", err)
		}
		n.Country = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode nearestDc#8e1a1775: field this_dc: %w", err)
		}
		n.ThisDC = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode nearestDc#8e1a1775: field nearest_dc: %w", err)
		}
		n.NearestDC = value
	}
	return nil
}

// GetCountry returns value of Country field.
func (n *NearestDC) GetCountry() (value string) {
	return n.Country
}

// GetThisDC returns value of ThisDC field.
func (n *NearestDC) GetThisDC() (value int) {
	return n.ThisDC
}

// GetNearestDC returns value of NearestDC field.
func (n *NearestDC) GetNearestDC() (value int) {
	return n.NearestDC
}
