// Code generated by gotdgen, DO NOT EDIT.

package td

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

// Error represents TL type `error#14feebbc`.
//
// See https://localhost:80/doc/constructor/error for reference.
type Error struct {
	// Error code; subject to future changes. If the error code is 406, the error message
	// must not be processed in any way and must not be displayed to the user
	Code int32
	// Error message; subject to future changes
	Message string
	// Temporary field of Error.
	Temporary bool
}

// ErrorTypeID is TL type id of Error.
const ErrorTypeID = 0x14feebbc

// Ensuring interfaces in compile-time for Error.
var (
	_ bin.Encoder     = &Error{}
	_ bin.Decoder     = &Error{}
	_ bin.BareEncoder = &Error{}
	_ bin.BareDecoder = &Error{}
)

func (e *Error) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Code == 0) {
		return false
	}
	if !(e.Message == "") {
		return false
	}
	if !(e.Temporary == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *Error) String() string {
	if e == nil {
		return "Error(nil)"
	}
	type Alias Error
	return fmt.Sprintf("Error%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Error) TypeID() uint32 {
	return ErrorTypeID
}

// TypeName returns name of type in TL schema.
func (*Error) TypeName() string {
	return "error"
}

// TypeInfo returns info about TL type.
func (e *Error) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "error",
		ID:   ErrorTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Code",
			SchemaName: "code",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Temporary",
			SchemaName: "temporary",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *Error) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode error#14feebbc as nil")
	}
	b.PutID(ErrorTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *Error) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode error#14feebbc as nil")
	}
	b.PutInt32(e.Code)
	b.PutString(e.Message)
	b.PutBool(e.Temporary)
	return nil
}

// Decode implements bin.Decoder.
func (e *Error) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode error#14feebbc to nil")
	}
	if err := b.ConsumeID(ErrorTypeID); err != nil {
		return fmt.Errorf("unable to decode error#14feebbc: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *Error) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode error#14feebbc to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode error#14feebbc: field code: %w", err)
		}
		e.Code = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode error#14feebbc: field message: %w", err)
		}
		e.Message = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode error#14feebbc: field temporary: %w", err)
		}
		e.Temporary = value
	}
	return nil
}

// GetCode returns value of Code field.
func (e *Error) GetCode() (value int32) {
	if e == nil {
		return
	}
	return e.Code
}

// GetMessage returns value of Message field.
func (e *Error) GetMessage() (value string) {
	if e == nil {
		return
	}
	return e.Message
}

// GetTemporary returns value of Temporary field.
func (e *Error) GetTemporary() (value bool) {
	if e == nil {
		return
	}
	return e.Temporary
}