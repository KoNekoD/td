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

// TestCallVectorStringObjectRequest represents TL type `testCallVectorStringObject#96cd6de`.
type TestCallVectorStringObjectRequest struct {
	// Vector of objects to return
	X []TestString
}

// TestCallVectorStringObjectRequestTypeID is TL type id of TestCallVectorStringObjectRequest.
const TestCallVectorStringObjectRequestTypeID = 0x96cd6de

// Ensuring interfaces in compile-time for TestCallVectorStringObjectRequest.
var (
	_ bin.Encoder     = &TestCallVectorStringObjectRequest{}
	_ bin.Decoder     = &TestCallVectorStringObjectRequest{}
	_ bin.BareEncoder = &TestCallVectorStringObjectRequest{}
	_ bin.BareDecoder = &TestCallVectorStringObjectRequest{}
)

func (t *TestCallVectorStringObjectRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.X == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestCallVectorStringObjectRequest) String() string {
	if t == nil {
		return "TestCallVectorStringObjectRequest(nil)"
	}
	type Alias TestCallVectorStringObjectRequest
	return fmt.Sprintf("TestCallVectorStringObjectRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestCallVectorStringObjectRequest) TypeID() uint32 {
	return TestCallVectorStringObjectRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TestCallVectorStringObjectRequest) TypeName() string {
	return "testCallVectorStringObject"
}

// TypeInfo returns info about TL type.
func (t *TestCallVectorStringObjectRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testCallVectorStringObject",
		ID:   TestCallVectorStringObjectRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "X",
			SchemaName: "x",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TestCallVectorStringObjectRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallVectorStringObject#96cd6de as nil")
	}
	b.PutID(TestCallVectorStringObjectRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestCallVectorStringObjectRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallVectorStringObject#96cd6de as nil")
	}
	b.PutInt(len(t.X))
	for idx, v := range t.X {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare testCallVectorStringObject#96cd6de: field x element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TestCallVectorStringObjectRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallVectorStringObject#96cd6de to nil")
	}
	if err := b.ConsumeID(TestCallVectorStringObjectRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode testCallVectorStringObject#96cd6de: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestCallVectorStringObjectRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallVectorStringObject#96cd6de to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode testCallVectorStringObject#96cd6de: field x: %w", err)
		}

		if headerLen > 0 {
			t.X = make([]TestString, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value TestString
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare testCallVectorStringObject#96cd6de: field x: %w", err)
			}
			t.X = append(t.X, value)
		}
	}
	return nil
}

// EncodeTDLibJSON encodes t in TDLib API JSON format.
func (t *TestCallVectorStringObjectRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallVectorStringObject#96cd6de as nil")
	}
	b.ObjStart()
	b.PutID("testCallVectorStringObject")
	b.FieldStart("x")
	b.ArrStart()
	for idx, v := range t.X {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode testCallVectorStringObject#96cd6de: field x element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// GetX returns value of X field.
func (t *TestCallVectorStringObjectRequest) GetX() (value []TestString) {
	return t.X
}

// TestCallVectorStringObject invokes method testCallVectorStringObject#96cd6de returning error if any.
func (c *Client) TestCallVectorStringObject(ctx context.Context, x []TestString) (*TestVectorStringObject, error) {
	var result TestVectorStringObject

	request := &TestCallVectorStringObjectRequest{
		X: x,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
