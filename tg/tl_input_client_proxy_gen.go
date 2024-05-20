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

// InputClientProxy represents TL type `inputClientProxy#75588b3f`.
// Info about an MTProxy¹ used to connect.
//
// Links:
//  1. https://core.telegram.org/mtproto/mtproto-transports#transport-obfuscation
//
// See https://core.telegram.org/constructor/inputClientProxy for reference.
type InputClientProxy struct {
	// Proxy address
	Address string
	// Proxy port
	Port int
}

// InputClientProxyTypeID is TL type id of InputClientProxy.
const InputClientProxyTypeID = 0x75588b3f

// Ensuring interfaces in compile-time for InputClientProxy.
var (
	_ bin.Encoder     = &InputClientProxy{}
	_ bin.Decoder     = &InputClientProxy{}
	_ bin.BareEncoder = &InputClientProxy{}
	_ bin.BareDecoder = &InputClientProxy{}
)

func (i *InputClientProxy) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Address == "") {
		return false
	}
	if !(i.Port == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputClientProxy) String() string {
	if i == nil {
		return "InputClientProxy(nil)"
	}
	type Alias InputClientProxy
	return fmt.Sprintf("InputClientProxy%+v", Alias(*i))
}

// FillFrom fills InputClientProxy from given interface.
func (i *InputClientProxy) FillFrom(from interface {
	GetAddress() (value string)
	GetPort() (value int)
}) {
	i.Address = from.GetAddress()
	i.Port = from.GetPort()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputClientProxy) TypeID() uint32 {
	return InputClientProxyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputClientProxy) TypeName() string {
	return "inputClientProxy"
}

// TypeInfo returns info about TL type.
func (i *InputClientProxy) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputClientProxy",
		ID:   InputClientProxyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Address",
			SchemaName: "address",
		},
		{
			Name:       "Port",
			SchemaName: "port",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputClientProxy) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputClientProxy#75588b3f as nil")
	}
	b.PutID(InputClientProxyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputClientProxy) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputClientProxy#75588b3f as nil")
	}
	b.PutString(i.Address)
	b.PutInt(i.Port)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputClientProxy) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputClientProxy#75588b3f to nil")
	}
	if err := b.ConsumeID(InputClientProxyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputClientProxy#75588b3f: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputClientProxy) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputClientProxy#75588b3f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputClientProxy#75588b3f: field address: %w", err)
		}
		i.Address = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputClientProxy#75588b3f: field port: %w", err)
		}
		i.Port = value
	}
	return nil
}

// GetAddress returns value of Address field.
func (i *InputClientProxy) GetAddress() (value string) {
	if i == nil {
		return
	}
	return i.Address
}

// GetPort returns value of Port field.
func (i *InputClientProxy) GetPort() (value int) {
	if i == nil {
		return
	}
	return i.Port
}
