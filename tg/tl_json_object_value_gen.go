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

// JSONObjectValue represents TL type `jsonObjectValue#c0de1bd9`.
// JSON key: value pair
//
// See https://core.telegram.org/constructor/jsonObjectValue for reference.
type JSONObjectValue struct {
	// Key
	Key string
	// Value
	Value JSONValueClass
}

// JSONObjectValueTypeID is TL type id of JSONObjectValue.
const JSONObjectValueTypeID = 0xc0de1bd9

// Ensuring interfaces in compile-time for JSONObjectValue.
var (
	_ bin.Encoder     = &JSONObjectValue{}
	_ bin.Decoder     = &JSONObjectValue{}
	_ bin.BareEncoder = &JSONObjectValue{}
	_ bin.BareDecoder = &JSONObjectValue{}
)

func (j *JSONObjectValue) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Key == "") {
		return false
	}
	if !(j.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JSONObjectValue) String() string {
	if j == nil {
		return "JSONObjectValue(nil)"
	}
	type Alias JSONObjectValue
	return fmt.Sprintf("JSONObjectValue%+v", Alias(*j))
}

// FillFrom fills JSONObjectValue from given interface.
func (j *JSONObjectValue) FillFrom(from interface {
	GetKey() (value string)
	GetValue() (value JSONValueClass)
}) {
	j.Key = from.GetKey()
	j.Value = from.GetValue()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JSONObjectValue) TypeID() uint32 {
	return JSONObjectValueTypeID
}

// TypeName returns name of type in TL schema.
func (*JSONObjectValue) TypeName() string {
	return "jsonObjectValue"
}

// TypeInfo returns info about TL type.
func (j *JSONObjectValue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "jsonObjectValue",
		ID:   JSONObjectValueTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Key",
			SchemaName: "key",
		},
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *JSONObjectValue) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonObjectValue#c0de1bd9 as nil")
	}
	b.PutID(JSONObjectValueTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JSONObjectValue) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode jsonObjectValue#c0de1bd9 as nil")
	}
	b.PutString(j.Key)
	if j.Value == nil {
		return fmt.Errorf("unable to encode jsonObjectValue#c0de1bd9: field value is nil")
	}
	if err := j.Value.Encode(b); err != nil {
		return fmt.Errorf("unable to encode jsonObjectValue#c0de1bd9: field value: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *JSONObjectValue) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonObjectValue#c0de1bd9 to nil")
	}
	if err := b.ConsumeID(JSONObjectValueTypeID); err != nil {
		return fmt.Errorf("unable to decode jsonObjectValue#c0de1bd9: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JSONObjectValue) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode jsonObjectValue#c0de1bd9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode jsonObjectValue#c0de1bd9: field key: %w", err)
		}
		j.Key = value
	}
	{
		value, err := DecodeJSONValue(b)
		if err != nil {
			return fmt.Errorf("unable to decode jsonObjectValue#c0de1bd9: field value: %w", err)
		}
		j.Value = value
	}
	return nil
}

// GetKey returns value of Key field.
func (j *JSONObjectValue) GetKey() (value string) {
	if j == nil {
		return
	}
	return j.Key
}

// GetValue returns value of Value field.
func (j *JSONObjectValue) GetValue() (value JSONValueClass) {
	if j == nil {
		return
	}
	return j.Value
}
