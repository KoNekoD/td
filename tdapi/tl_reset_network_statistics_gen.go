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

// ResetNetworkStatisticsRequest represents TL type `resetNetworkStatistics#6222dd86`.
type ResetNetworkStatisticsRequest struct {
}

// ResetNetworkStatisticsRequestTypeID is TL type id of ResetNetworkStatisticsRequest.
const ResetNetworkStatisticsRequestTypeID = 0x6222dd86

// Ensuring interfaces in compile-time for ResetNetworkStatisticsRequest.
var (
	_ bin.Encoder     = &ResetNetworkStatisticsRequest{}
	_ bin.Decoder     = &ResetNetworkStatisticsRequest{}
	_ bin.BareEncoder = &ResetNetworkStatisticsRequest{}
	_ bin.BareDecoder = &ResetNetworkStatisticsRequest{}
)

func (r *ResetNetworkStatisticsRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResetNetworkStatisticsRequest) String() string {
	if r == nil {
		return "ResetNetworkStatisticsRequest(nil)"
	}
	type Alias ResetNetworkStatisticsRequest
	return fmt.Sprintf("ResetNetworkStatisticsRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResetNetworkStatisticsRequest) TypeID() uint32 {
	return ResetNetworkStatisticsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ResetNetworkStatisticsRequest) TypeName() string {
	return "resetNetworkStatistics"
}

// TypeInfo returns info about TL type.
func (r *ResetNetworkStatisticsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resetNetworkStatistics",
		ID:   ResetNetworkStatisticsRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResetNetworkStatisticsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resetNetworkStatistics#6222dd86 as nil")
	}
	b.PutID(ResetNetworkStatisticsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResetNetworkStatisticsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resetNetworkStatistics#6222dd86 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ResetNetworkStatisticsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resetNetworkStatistics#6222dd86 to nil")
	}
	if err := b.ConsumeID(ResetNetworkStatisticsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode resetNetworkStatistics#6222dd86: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResetNetworkStatisticsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resetNetworkStatistics#6222dd86 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ResetNetworkStatisticsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode resetNetworkStatistics#6222dd86 as nil")
	}
	b.ObjStart()
	b.PutID("resetNetworkStatistics")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ResetNetworkStatisticsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode resetNetworkStatistics#6222dd86 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("resetNetworkStatistics"); err != nil {
				return fmt.Errorf("unable to decode resetNetworkStatistics#6222dd86: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ResetNetworkStatistics invokes method resetNetworkStatistics#6222dd86 returning error if any.
func (c *Client) ResetNetworkStatistics(ctx context.Context) error {
	var ok Ok

	request := &ResetNetworkStatisticsRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
