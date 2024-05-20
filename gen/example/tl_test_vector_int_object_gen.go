// Code generated by gotdgen, DO NOT EDIT.

package td

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

// TestVectorIntObject represents TL type `testVectorIntObject#f152999b`.
//
// See https://localhost:80/doc/constructor/testVectorIntObject for reference.
type TestVectorIntObject struct {
	// Vector of objects
	Value []TestInt
}

// TestVectorIntObjectTypeID is TL type id of TestVectorIntObject.
const TestVectorIntObjectTypeID = 0xf152999b

// Ensuring interfaces in compile-time for TestVectorIntObject.
var (
	_ bin.Encoder     = &TestVectorIntObject{}
	_ bin.Decoder     = &TestVectorIntObject{}
	_ bin.BareEncoder = &TestVectorIntObject{}
	_ bin.BareDecoder = &TestVectorIntObject{}
)

func (t *TestVectorIntObject) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Value == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestVectorIntObject) String() string {
	if t == nil {
		return "TestVectorIntObject(nil)"
	}
	type Alias TestVectorIntObject
	return fmt.Sprintf("TestVectorIntObject%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestVectorIntObject) TypeID() uint32 {
	return TestVectorIntObjectTypeID
}

// TypeName returns name of type in TL schema.
func (*TestVectorIntObject) TypeName() string {
	return "testVectorIntObject"
}

// TypeInfo returns info about TL type.
func (t *TestVectorIntObject) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testVectorIntObject",
		ID:   TestVectorIntObjectTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestVectorIntObject) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorIntObject#f152999b as nil")
	}
	b.PutID(TestVectorIntObjectTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestVectorIntObject) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testVectorIntObject#f152999b as nil")
	}
	b.PutInt(len(t.Value))
	for idx, v := range t.Value {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare testVectorIntObject#f152999b: field value element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestVectorIntObject) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorIntObject#f152999b to nil")
	}
	if err := b.ConsumeID(TestVectorIntObjectTypeID); err != nil {
		return fmt.Errorf("unable to decode testVectorIntObject#f152999b: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestVectorIntObject) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testVectorIntObject#f152999b to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode testVectorIntObject#f152999b: field value: %w", err)
		}

		if headerLen > 0 {
			t.Value = make([]TestInt, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TestInt
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare testVectorIntObject#f152999b: field value: %w", err)
			}
			t.Value = append(t.Value, value)
		}
	}
	return nil
}

// GetValue returns value of Value field.
func (t *TestVectorIntObject) GetValue() (value []TestInt) {
	if t == nil {
		return
	}
	return t.Value
}
