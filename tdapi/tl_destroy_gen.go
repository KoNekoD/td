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

// DestroyRequest represents TL type `destroy#28d9534a`.
type DestroyRequest struct {
}

// DestroyRequestTypeID is TL type id of DestroyRequest.
const DestroyRequestTypeID = 0x28d9534a

// Ensuring interfaces in compile-time for DestroyRequest.
var (
	_ bin.Encoder     = &DestroyRequest{}
	_ bin.Decoder     = &DestroyRequest{}
	_ bin.BareEncoder = &DestroyRequest{}
	_ bin.BareDecoder = &DestroyRequest{}
)

func (d *DestroyRequest) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *DestroyRequest) String() string {
	if d == nil {
		return "DestroyRequest(nil)"
	}
	type Alias DestroyRequest
	return fmt.Sprintf("DestroyRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DestroyRequest) TypeID() uint32 {
	return DestroyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DestroyRequest) TypeName() string {
	return "destroy"
}

// TypeInfo returns info about TL type.
func (d *DestroyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "destroy",
		ID:   DestroyRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (d *DestroyRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy#28d9534a as nil")
	}
	b.PutID(DestroyRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DestroyRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy#28d9534a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DestroyRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy#28d9534a to nil")
	}
	if err := b.ConsumeID(DestroyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode destroy#28d9534a: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DestroyRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy#28d9534a to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DestroyRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode destroy#28d9534a as nil")
	}
	b.ObjStart()
	b.PutID("destroy")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DestroyRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode destroy#28d9534a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("destroy"); err != nil {
				return fmt.Errorf("unable to decode destroy#28d9534a: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// Destroy invokes method destroy#28d9534a returning error if any.
func (c *Client) Destroy(ctx context.Context) error {
	var ok Ok

	request := &DestroyRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
