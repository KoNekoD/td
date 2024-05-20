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

// CallProtocol represents TL type `callProtocol#a9a037e`.
type CallProtocol struct {
	// True, if UDP peer-to-peer connections are supported
	UDPP2P bool
	// True, if connection through UDP reflectors is supported
	UDPReflector bool
	// The minimum supported API layer; use 65
	MinLayer int32
	// The maximum supported API layer; use 92
	MaxLayer int32
	// List of supported tgcalls versions
	LibraryVersions []string
}

// CallProtocolTypeID is TL type id of CallProtocol.
const CallProtocolTypeID = 0xa9a037e

// Ensuring interfaces in compile-time for CallProtocol.
var (
	_ bin.Encoder     = &CallProtocol{}
	_ bin.Decoder     = &CallProtocol{}
	_ bin.BareEncoder = &CallProtocol{}
	_ bin.BareDecoder = &CallProtocol{}
)

func (c *CallProtocol) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UDPP2P == false) {
		return false
	}
	if !(c.UDPReflector == false) {
		return false
	}
	if !(c.MinLayer == 0) {
		return false
	}
	if !(c.MaxLayer == 0) {
		return false
	}
	if !(c.LibraryVersions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CallProtocol) String() string {
	if c == nil {
		return "CallProtocol(nil)"
	}
	type Alias CallProtocol
	return fmt.Sprintf("CallProtocol%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CallProtocol) TypeID() uint32 {
	return CallProtocolTypeID
}

// TypeName returns name of type in TL schema.
func (*CallProtocol) TypeName() string {
	return "callProtocol"
}

// TypeInfo returns info about TL type.
func (c *CallProtocol) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "callProtocol",
		ID:   CallProtocolTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UDPP2P",
			SchemaName: "udp_p2p",
		},
		{
			Name:       "UDPReflector",
			SchemaName: "udp_reflector",
		},
		{
			Name:       "MinLayer",
			SchemaName: "min_layer",
		},
		{
			Name:       "MaxLayer",
			SchemaName: "max_layer",
		},
		{
			Name:       "LibraryVersions",
			SchemaName: "library_versions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CallProtocol) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callProtocol#a9a037e as nil")
	}
	b.PutID(CallProtocolTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CallProtocol) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode callProtocol#a9a037e as nil")
	}
	b.PutBool(c.UDPP2P)
	b.PutBool(c.UDPReflector)
	b.PutInt32(c.MinLayer)
	b.PutInt32(c.MaxLayer)
	b.PutInt(len(c.LibraryVersions))
	for _, v := range c.LibraryVersions {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *CallProtocol) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callProtocol#a9a037e to nil")
	}
	if err := b.ConsumeID(CallProtocolTypeID); err != nil {
		return fmt.Errorf("unable to decode callProtocol#a9a037e: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CallProtocol) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode callProtocol#a9a037e to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode callProtocol#a9a037e: field udp_p2p: %w", err)
		}
		c.UDPP2P = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode callProtocol#a9a037e: field udp_reflector: %w", err)
		}
		c.UDPReflector = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode callProtocol#a9a037e: field min_layer: %w", err)
		}
		c.MinLayer = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode callProtocol#a9a037e: field max_layer: %w", err)
		}
		c.MaxLayer = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode callProtocol#a9a037e: field library_versions: %w", err)
		}

		if headerLen > 0 {
			c.LibraryVersions = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode callProtocol#a9a037e: field library_versions: %w", err)
			}
			c.LibraryVersions = append(c.LibraryVersions, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CallProtocol) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode callProtocol#a9a037e as nil")
	}
	b.ObjStart()
	b.PutID("callProtocol")
	b.Comma()
	b.FieldStart("udp_p2p")
	b.PutBool(c.UDPP2P)
	b.Comma()
	b.FieldStart("udp_reflector")
	b.PutBool(c.UDPReflector)
	b.Comma()
	b.FieldStart("min_layer")
	b.PutInt32(c.MinLayer)
	b.Comma()
	b.FieldStart("max_layer")
	b.PutInt32(c.MaxLayer)
	b.Comma()
	b.FieldStart("library_versions")
	b.ArrStart()
	for _, v := range c.LibraryVersions {
		b.PutString(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CallProtocol) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode callProtocol#a9a037e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("callProtocol"); err != nil {
				return fmt.Errorf("unable to decode callProtocol#a9a037e: %w", err)
			}
		case "udp_p2p":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode callProtocol#a9a037e: field udp_p2p: %w", err)
			}
			c.UDPP2P = value
		case "udp_reflector":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode callProtocol#a9a037e: field udp_reflector: %w", err)
			}
			c.UDPReflector = value
		case "min_layer":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode callProtocol#a9a037e: field min_layer: %w", err)
			}
			c.MinLayer = value
		case "max_layer":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode callProtocol#a9a037e: field max_layer: %w", err)
			}
			c.MaxLayer = value
		case "library_versions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode callProtocol#a9a037e: field library_versions: %w", err)
				}
				c.LibraryVersions = append(c.LibraryVersions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode callProtocol#a9a037e: field library_versions: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUDPP2P returns value of UDPP2P field.
func (c *CallProtocol) GetUDPP2P() (value bool) {
	if c == nil {
		return
	}
	return c.UDPP2P
}

// GetUDPReflector returns value of UDPReflector field.
func (c *CallProtocol) GetUDPReflector() (value bool) {
	if c == nil {
		return
	}
	return c.UDPReflector
}

// GetMinLayer returns value of MinLayer field.
func (c *CallProtocol) GetMinLayer() (value int32) {
	if c == nil {
		return
	}
	return c.MinLayer
}

// GetMaxLayer returns value of MaxLayer field.
func (c *CallProtocol) GetMaxLayer() (value int32) {
	if c == nil {
		return
	}
	return c.MaxLayer
}

// GetLibraryVersions returns value of LibraryVersions field.
func (c *CallProtocol) GetLibraryVersions() (value []string) {
	if c == nil {
		return
	}
	return c.LibraryVersions
}
