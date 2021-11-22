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

// AddNetworkStatisticsRequest represents TL type `addNetworkStatistics#4b63b3d9`.
type AddNetworkStatisticsRequest struct {
	// The network statistics entry with the data to be added to statistics
	Entry NetworkStatisticsEntryClass
}

// AddNetworkStatisticsRequestTypeID is TL type id of AddNetworkStatisticsRequest.
const AddNetworkStatisticsRequestTypeID = 0x4b63b3d9

// Ensuring interfaces in compile-time for AddNetworkStatisticsRequest.
var (
	_ bin.Encoder     = &AddNetworkStatisticsRequest{}
	_ bin.Decoder     = &AddNetworkStatisticsRequest{}
	_ bin.BareEncoder = &AddNetworkStatisticsRequest{}
	_ bin.BareDecoder = &AddNetworkStatisticsRequest{}
)

func (a *AddNetworkStatisticsRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Entry == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AddNetworkStatisticsRequest) String() string {
	if a == nil {
		return "AddNetworkStatisticsRequest(nil)"
	}
	type Alias AddNetworkStatisticsRequest
	return fmt.Sprintf("AddNetworkStatisticsRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AddNetworkStatisticsRequest) TypeID() uint32 {
	return AddNetworkStatisticsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AddNetworkStatisticsRequest) TypeName() string {
	return "addNetworkStatistics"
}

// TypeInfo returns info about TL type.
func (a *AddNetworkStatisticsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "addNetworkStatistics",
		ID:   AddNetworkStatisticsRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Entry",
			SchemaName: "entry",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AddNetworkStatisticsRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addNetworkStatistics#4b63b3d9 as nil")
	}
	b.PutID(AddNetworkStatisticsRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AddNetworkStatisticsRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode addNetworkStatistics#4b63b3d9 as nil")
	}
	if a.Entry == nil {
		return fmt.Errorf("unable to encode addNetworkStatistics#4b63b3d9: field entry is nil")
	}
	if err := a.Entry.Encode(b); err != nil {
		return fmt.Errorf("unable to encode addNetworkStatistics#4b63b3d9: field entry: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AddNetworkStatisticsRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addNetworkStatistics#4b63b3d9 to nil")
	}
	if err := b.ConsumeID(AddNetworkStatisticsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode addNetworkStatistics#4b63b3d9: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AddNetworkStatisticsRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode addNetworkStatistics#4b63b3d9 to nil")
	}
	{
		value, err := DecodeNetworkStatisticsEntry(b)
		if err != nil {
			return fmt.Errorf("unable to decode addNetworkStatistics#4b63b3d9: field entry: %w", err)
		}
		a.Entry = value
	}
	return nil
}

// EncodeTDLibJSON encodes a in TDLib API JSON format.
func (a *AddNetworkStatisticsRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode addNetworkStatistics#4b63b3d9 as nil")
	}
	b.ObjStart()
	b.PutID("addNetworkStatistics")
	b.FieldStart("entry")
	if a.Entry == nil {
		return fmt.Errorf("unable to encode addNetworkStatistics#4b63b3d9: field entry is nil")
	}
	if err := a.Entry.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode addNetworkStatistics#4b63b3d9: field entry: %w", err)
	}
	b.ObjEnd()
	return nil
}

// GetEntry returns value of Entry field.
func (a *AddNetworkStatisticsRequest) GetEntry() (value NetworkStatisticsEntryClass) {
	return a.Entry
}

// AddNetworkStatistics invokes method addNetworkStatistics#4b63b3d9 returning error if any.
func (c *Client) AddNetworkStatistics(ctx context.Context, entry NetworkStatisticsEntryClass) error {
	var ok Ok

	request := &AddNetworkStatisticsRequest{
		Entry: entry,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
