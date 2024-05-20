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

// TestCallStringRequest represents TL type `testCallString#98b74a2f`.
type TestCallStringRequest struct {
	// String to return
	X string
}

// TestCallStringRequestTypeID is TL type id of TestCallStringRequest.
const TestCallStringRequestTypeID = 0x98b74a2f

// Ensuring interfaces in compile-time for TestCallStringRequest.
var (
	_ bin.Encoder     = &TestCallStringRequest{}
	_ bin.Decoder     = &TestCallStringRequest{}
	_ bin.BareEncoder = &TestCallStringRequest{}
	_ bin.BareDecoder = &TestCallStringRequest{}
)

func (t *TestCallStringRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.X == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TestCallStringRequest) String() string {
	if t == nil {
		return "TestCallStringRequest(nil)"
	}
	type Alias TestCallStringRequest
	return fmt.Sprintf("TestCallStringRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestCallStringRequest) TypeID() uint32 {
	return TestCallStringRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TestCallStringRequest) TypeName() string {
	return "testCallString"
}

// TypeInfo returns info about TL type.
func (t *TestCallStringRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "testCallString",
		ID:   TestCallStringRequestTypeID,
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
func (t *TestCallStringRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallString#98b74a2f as nil")
	}
	b.PutID(TestCallStringRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TestCallStringRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallString#98b74a2f as nil")
	}
	b.PutString(t.X)
	return nil
}

// Decode implements bin.Decoder.
func (t *TestCallStringRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallString#98b74a2f to nil")
	}
	if err := b.ConsumeID(TestCallStringRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode testCallString#98b74a2f: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TestCallStringRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallString#98b74a2f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode testCallString#98b74a2f: field x: %w", err)
		}
		t.X = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TestCallStringRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode testCallString#98b74a2f as nil")
	}
	b.ObjStart()
	b.PutID("testCallString")
	b.Comma()
	b.FieldStart("x")
	b.PutString(t.X)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TestCallStringRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode testCallString#98b74a2f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("testCallString"); err != nil {
				return fmt.Errorf("unable to decode testCallString#98b74a2f: %w", err)
			}
		case "x":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode testCallString#98b74a2f: field x: %w", err)
			}
			t.X = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetX returns value of X field.
func (t *TestCallStringRequest) GetX() (value string) {
	if t == nil {
		return
	}
	return t.X
}

// TestCallString invokes method testCallString#98b74a2f returning error if any.
func (c *Client) TestCallString(ctx context.Context, x string) (*TestString, error) {
	var result TestString

	request := &TestCallStringRequest{
		X: x,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
