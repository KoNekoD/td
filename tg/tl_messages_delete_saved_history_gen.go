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

// MessagesDeleteSavedHistoryRequest represents TL type `messages.deleteSavedHistory#6e98102b`.
// Deletes messages forwarded from a specific peer to saved messages »¹.
//
// Links:
//  1. https://core.telegram.org/api/saved-messages
//
// See https://core.telegram.org/method/messages.deleteSavedHistory for reference.
type MessagesDeleteSavedHistoryRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Peer, whose messages will be deleted from saved messages »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/saved-messages
	Peer InputPeerClass
	// Maximum ID of message to delete
	MaxID int
	// Delete all messages newer than this UNIX timestamp
	//
	// Use SetMinDate and GetMinDate helpers.
	MinDate int
	// Delete all messages older than this UNIX timestamp
	//
	// Use SetMaxDate and GetMaxDate helpers.
	MaxDate int
}

// MessagesDeleteSavedHistoryRequestTypeID is TL type id of MessagesDeleteSavedHistoryRequest.
const MessagesDeleteSavedHistoryRequestTypeID = 0x6e98102b

// Ensuring interfaces in compile-time for MessagesDeleteSavedHistoryRequest.
var (
	_ bin.Encoder     = &MessagesDeleteSavedHistoryRequest{}
	_ bin.Decoder     = &MessagesDeleteSavedHistoryRequest{}
	_ bin.BareEncoder = &MessagesDeleteSavedHistoryRequest{}
	_ bin.BareDecoder = &MessagesDeleteSavedHistoryRequest{}
)

func (d *MessagesDeleteSavedHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.Peer == nil) {
		return false
	}
	if !(d.MaxID == 0) {
		return false
	}
	if !(d.MinDate == 0) {
		return false
	}
	if !(d.MaxDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *MessagesDeleteSavedHistoryRequest) String() string {
	if d == nil {
		return "MessagesDeleteSavedHistoryRequest(nil)"
	}
	type Alias MessagesDeleteSavedHistoryRequest
	return fmt.Sprintf("MessagesDeleteSavedHistoryRequest%+v", Alias(*d))
}

// FillFrom fills MessagesDeleteSavedHistoryRequest from given interface.
func (d *MessagesDeleteSavedHistoryRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetMaxID() (value int)
	GetMinDate() (value int, ok bool)
	GetMaxDate() (value int, ok bool)
}) {
	d.Peer = from.GetPeer()
	d.MaxID = from.GetMaxID()
	if val, ok := from.GetMinDate(); ok {
		d.MinDate = val
	}

	if val, ok := from.GetMaxDate(); ok {
		d.MaxDate = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesDeleteSavedHistoryRequest) TypeID() uint32 {
	return MessagesDeleteSavedHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesDeleteSavedHistoryRequest) TypeName() string {
	return "messages.deleteSavedHistory"
}

// TypeInfo returns info about TL type.
func (d *MessagesDeleteSavedHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.deleteSavedHistory",
		ID:   MessagesDeleteSavedHistoryRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
		{
			Name:       "MinDate",
			SchemaName: "min_date",
			Null:       !d.Flags.Has(2),
		},
		{
			Name:       "MaxDate",
			SchemaName: "max_date",
			Null:       !d.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (d *MessagesDeleteSavedHistoryRequest) SetFlags() {
	if !(d.MinDate == 0) {
		d.Flags.Set(2)
	}
	if !(d.MaxDate == 0) {
		d.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (d *MessagesDeleteSavedHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteSavedHistory#6e98102b as nil")
	}
	b.PutID(MessagesDeleteSavedHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *MessagesDeleteSavedHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode messages.deleteSavedHistory#6e98102b as nil")
	}
	d.SetFlags()
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteSavedHistory#6e98102b: field flags: %w", err)
	}
	if d.Peer == nil {
		return fmt.Errorf("unable to encode messages.deleteSavedHistory#6e98102b: field peer is nil")
	}
	if err := d.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.deleteSavedHistory#6e98102b: field peer: %w", err)
	}
	b.PutInt(d.MaxID)
	if d.Flags.Has(2) {
		b.PutInt(d.MinDate)
	}
	if d.Flags.Has(3) {
		b.PutInt(d.MaxDate)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *MessagesDeleteSavedHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteSavedHistory#6e98102b to nil")
	}
	if err := b.ConsumeID(MessagesDeleteSavedHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.deleteSavedHistory#6e98102b: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *MessagesDeleteSavedHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode messages.deleteSavedHistory#6e98102b to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.deleteSavedHistory#6e98102b: field flags: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteSavedHistory#6e98102b: field peer: %w", err)
		}
		d.Peer = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteSavedHistory#6e98102b: field max_id: %w", err)
		}
		d.MaxID = value
	}
	if d.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteSavedHistory#6e98102b: field min_date: %w", err)
		}
		d.MinDate = value
	}
	if d.Flags.Has(3) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.deleteSavedHistory#6e98102b: field max_date: %w", err)
		}
		d.MaxDate = value
	}
	return nil
}

// GetPeer returns value of Peer field.
func (d *MessagesDeleteSavedHistoryRequest) GetPeer() (value InputPeerClass) {
	if d == nil {
		return
	}
	return d.Peer
}

// GetMaxID returns value of MaxID field.
func (d *MessagesDeleteSavedHistoryRequest) GetMaxID() (value int) {
	if d == nil {
		return
	}
	return d.MaxID
}

// SetMinDate sets value of MinDate conditional field.
func (d *MessagesDeleteSavedHistoryRequest) SetMinDate(value int) {
	d.Flags.Set(2)
	d.MinDate = value
}

// GetMinDate returns value of MinDate conditional field and
// boolean which is true if field was set.
func (d *MessagesDeleteSavedHistoryRequest) GetMinDate() (value int, ok bool) {
	if d == nil {
		return
	}
	if !d.Flags.Has(2) {
		return value, false
	}
	return d.MinDate, true
}

// SetMaxDate sets value of MaxDate conditional field.
func (d *MessagesDeleteSavedHistoryRequest) SetMaxDate(value int) {
	d.Flags.Set(3)
	d.MaxDate = value
}

// GetMaxDate returns value of MaxDate conditional field and
// boolean which is true if field was set.
func (d *MessagesDeleteSavedHistoryRequest) GetMaxDate() (value int, ok bool) {
	if d == nil {
		return
	}
	if !d.Flags.Has(3) {
		return value, false
	}
	return d.MaxDate, true
}

// MessagesDeleteSavedHistory invokes method messages.deleteSavedHistory#6e98102b returning error if any.
// Deletes messages forwarded from a specific peer to saved messages »¹.
//
// Links:
//  1. https://core.telegram.org/api/saved-messages
//
// Possible errors:
//
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.deleteSavedHistory for reference.
func (c *Client) MessagesDeleteSavedHistory(ctx context.Context, request *MessagesDeleteSavedHistoryRequest) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
