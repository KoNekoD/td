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

// DisableProxyRequest represents TL type `disableProxy#82d31782`.
type DisableProxyRequest struct {
}

// DisableProxyRequestTypeID is TL type id of DisableProxyRequest.
const DisableProxyRequestTypeID = 0x82d31782

// Ensuring interfaces in compile-time for DisableProxyRequest.
var (
	_ bin.Encoder     = &DisableProxyRequest{}
	_ bin.Decoder     = &DisableProxyRequest{}
	_ bin.BareEncoder = &DisableProxyRequest{}
	_ bin.BareDecoder = &DisableProxyRequest{}
)

func (d *DisableProxyRequest) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *DisableProxyRequest) String() string {
	if d == nil {
		return "DisableProxyRequest(nil)"
	}
	type Alias DisableProxyRequest
	return fmt.Sprintf("DisableProxyRequest%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DisableProxyRequest) TypeID() uint32 {
	return DisableProxyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*DisableProxyRequest) TypeName() string {
	return "disableProxy"
}

// TypeInfo returns info about TL type.
func (d *DisableProxyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "disableProxy",
		ID:   DisableProxyRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (d *DisableProxyRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode disableProxy#82d31782 as nil")
	}
	b.PutID(DisableProxyRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DisableProxyRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode disableProxy#82d31782 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DisableProxyRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode disableProxy#82d31782 to nil")
	}
	if err := b.ConsumeID(DisableProxyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode disableProxy#82d31782: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DisableProxyRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode disableProxy#82d31782 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (d *DisableProxyRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if d == nil {
		return fmt.Errorf("can't encode disableProxy#82d31782 as nil")
	}
	b.ObjStart()
	b.PutID("disableProxy")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (d *DisableProxyRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if d == nil {
		return fmt.Errorf("can't decode disableProxy#82d31782 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("disableProxy"); err != nil {
				return fmt.Errorf("unable to decode disableProxy#82d31782: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// DisableProxy invokes method disableProxy#82d31782 returning error if any.
func (c *Client) DisableProxy(ctx context.Context) error {
	var ok Ok

	request := &DisableProxyRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
