// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// ChannelsExportMessageLinkRequest represents TL type `channels.exportMessageLink#e63fadeb`.
// Get link and embed info of a message in a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.exportMessageLink for reference.
type ChannelsExportMessageLinkRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to include other grouped media (for albums)
	Grouped bool
	// Whether to also include a thread ID, if available, inside of the link
	Thread bool
	// Channel
	Channel InputChannelClass
	// Message ID
	ID int
}

// ChannelsExportMessageLinkRequestTypeID is TL type id of ChannelsExportMessageLinkRequest.
const ChannelsExportMessageLinkRequestTypeID = 0xe63fadeb

// Ensuring interfaces in compile-time for ChannelsExportMessageLinkRequest.
var (
	_ bin.Encoder     = &ChannelsExportMessageLinkRequest{}
	_ bin.Decoder     = &ChannelsExportMessageLinkRequest{}
	_ bin.BareEncoder = &ChannelsExportMessageLinkRequest{}
	_ bin.BareDecoder = &ChannelsExportMessageLinkRequest{}
)

func (e *ChannelsExportMessageLinkRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.Grouped == false) {
		return false
	}
	if !(e.Thread == false) {
		return false
	}
	if !(e.Channel == nil) {
		return false
	}
	if !(e.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ChannelsExportMessageLinkRequest) String() string {
	if e == nil {
		return "ChannelsExportMessageLinkRequest(nil)"
	}
	type Alias ChannelsExportMessageLinkRequest
	return fmt.Sprintf("ChannelsExportMessageLinkRequest%+v", Alias(*e))
}

// FillFrom fills ChannelsExportMessageLinkRequest from given interface.
func (e *ChannelsExportMessageLinkRequest) FillFrom(from interface {
	GetGrouped() (value bool)
	GetThread() (value bool)
	GetChannel() (value InputChannelClass)
	GetID() (value int)
}) {
	e.Grouped = from.GetGrouped()
	e.Thread = from.GetThread()
	e.Channel = from.GetChannel()
	e.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsExportMessageLinkRequest) TypeID() uint32 {
	return ChannelsExportMessageLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsExportMessageLinkRequest) TypeName() string {
	return "channels.exportMessageLink"
}

// TypeInfo returns info about TL type.
func (e *ChannelsExportMessageLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.exportMessageLink",
		ID:   ChannelsExportMessageLinkRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Grouped",
			SchemaName: "grouped",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "Thread",
			SchemaName: "thread",
			Null:       !e.Flags.Has(1),
		},
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *ChannelsExportMessageLinkRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.exportMessageLink#e63fadeb as nil")
	}
	b.PutID(ChannelsExportMessageLinkRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ChannelsExportMessageLinkRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode channels.exportMessageLink#e63fadeb as nil")
	}
	if !(e.Grouped == false) {
		e.Flags.Set(0)
	}
	if !(e.Thread == false) {
		e.Flags.Set(1)
	}
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.exportMessageLink#e63fadeb: field flags: %w", err)
	}
	if e.Channel == nil {
		return fmt.Errorf("unable to encode channels.exportMessageLink#e63fadeb: field channel is nil")
	}
	if err := e.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.exportMessageLink#e63fadeb: field channel: %w", err)
	}
	b.PutInt(e.ID)
	return nil
}

// Decode implements bin.Decoder.
func (e *ChannelsExportMessageLinkRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.exportMessageLink#e63fadeb to nil")
	}
	if err := b.ConsumeID(ChannelsExportMessageLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ChannelsExportMessageLinkRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode channels.exportMessageLink#e63fadeb to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: field flags: %w", err)
		}
	}
	e.Grouped = e.Flags.Has(0)
	e.Thread = e.Flags.Has(1)
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: field channel: %w", err)
		}
		e.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.exportMessageLink#e63fadeb: field id: %w", err)
		}
		e.ID = value
	}
	return nil
}

// SetGrouped sets value of Grouped conditional field.
func (e *ChannelsExportMessageLinkRequest) SetGrouped(value bool) {
	if value {
		e.Flags.Set(0)
		e.Grouped = true
	} else {
		e.Flags.Unset(0)
		e.Grouped = false
	}
}

// GetGrouped returns value of Grouped conditional field.
func (e *ChannelsExportMessageLinkRequest) GetGrouped() (value bool) {
	return e.Flags.Has(0)
}

// SetThread sets value of Thread conditional field.
func (e *ChannelsExportMessageLinkRequest) SetThread(value bool) {
	if value {
		e.Flags.Set(1)
		e.Thread = true
	} else {
		e.Flags.Unset(1)
		e.Thread = false
	}
}

// GetThread returns value of Thread conditional field.
func (e *ChannelsExportMessageLinkRequest) GetThread() (value bool) {
	return e.Flags.Has(1)
}

// GetChannel returns value of Channel field.
func (e *ChannelsExportMessageLinkRequest) GetChannel() (value InputChannelClass) {
	return e.Channel
}

// GetID returns value of ID field.
func (e *ChannelsExportMessageLinkRequest) GetID() (value int) {
	return e.ID
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (e *ChannelsExportMessageLinkRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return e.Channel.AsNotEmpty()
}

// ChannelsExportMessageLink invokes method channels.exportMessageLink#e63fadeb returning error if any.
// Get link and embed info of a message in a channel/supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 MESSAGE_ID_INVALID: The provided message id is invalid
//  400 MSG_ID_INVALID: Invalid message ID provided
//
// See https://core.telegram.org/method/channels.exportMessageLink for reference.
func (c *Client) ChannelsExportMessageLink(ctx context.Context, request *ChannelsExportMessageLinkRequest) (*ExportedMessageLink, error) {
	var result ExportedMessageLink

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
