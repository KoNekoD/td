// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// Int represents TL type `int#a8509bda`.
//
// See https://core.telegram.org/constructor/int for reference.
type Int struct {
}

// IntTypeID is TL type id of Int.
const IntTypeID = 0xa8509bda

// Ensuring interfaces in compile-time for Int.
var (
	_ bin.Encoder     = &Int{}
	_ bin.Decoder     = &Int{}
	_ bin.BareEncoder = &Int{}
	_ bin.BareDecoder = &Int{}
)

func (i *Int) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *Int) String() string {
	if i == nil {
		return "Int(nil)"
	}
	type Alias Int
	return fmt.Sprintf("Int%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Int) TypeID() uint32 {
	return IntTypeID
}

// TypeName returns name of type in TL schema.
func (*Int) TypeName() string {
	return "int"
}

// TypeInfo returns info about TL type.
func (i *Int) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "int",
		ID:   IntTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *Int) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode int#a8509bda as nil")
	}
	b.PutID(IntTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *Int) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode int#a8509bda as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *Int) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode int#a8509bda to nil")
	}
	if err := b.ConsumeID(IntTypeID); err != nil {
		return fmt.Errorf("unable to decode int#a8509bda: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *Int) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode int#a8509bda to nil")
	}
	return nil
}
