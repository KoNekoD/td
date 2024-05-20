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

// MessagesAffectedFoundMessages represents TL type `messages.affectedFoundMessages#ef8d3e6c`.
// Messages found and affected by changes
//
// See https://core.telegram.org/constructor/messages.affectedFoundMessages for reference.
type MessagesAffectedFoundMessages struct {
	// Event count after generation¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	Pts int
	// Number of events that were generated¹
	//
	// Links:
	//  1) https://core.telegram.org/api/updates
	PtsCount int
	// If bigger than zero, the request must be repeated to remove more messages
	Offset int
	// Affected message IDs
	Messages []int
}

// MessagesAffectedFoundMessagesTypeID is TL type id of MessagesAffectedFoundMessages.
const MessagesAffectedFoundMessagesTypeID = 0xef8d3e6c

// Ensuring interfaces in compile-time for MessagesAffectedFoundMessages.
var (
	_ bin.Encoder     = &MessagesAffectedFoundMessages{}
	_ bin.Decoder     = &MessagesAffectedFoundMessages{}
	_ bin.BareEncoder = &MessagesAffectedFoundMessages{}
	_ bin.BareDecoder = &MessagesAffectedFoundMessages{}
)

func (a *MessagesAffectedFoundMessages) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Pts == 0) {
		return false
	}
	if !(a.PtsCount == 0) {
		return false
	}
	if !(a.Offset == 0) {
		return false
	}
	if !(a.Messages == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *MessagesAffectedFoundMessages) String() string {
	if a == nil {
		return "MessagesAffectedFoundMessages(nil)"
	}
	type Alias MessagesAffectedFoundMessages
	return fmt.Sprintf("MessagesAffectedFoundMessages%+v", Alias(*a))
}

// FillFrom fills MessagesAffectedFoundMessages from given interface.
func (a *MessagesAffectedFoundMessages) FillFrom(from interface {
	GetPts() (value int)
	GetPtsCount() (value int)
	GetOffset() (value int)
	GetMessages() (value []int)
}) {
	a.Pts = from.GetPts()
	a.PtsCount = from.GetPtsCount()
	a.Offset = from.GetOffset()
	a.Messages = from.GetMessages()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesAffectedFoundMessages) TypeID() uint32 {
	return MessagesAffectedFoundMessagesTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesAffectedFoundMessages) TypeName() string {
	return "messages.affectedFoundMessages"
}

// TypeInfo returns info about TL type.
func (a *MessagesAffectedFoundMessages) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.affectedFoundMessages",
		ID:   MessagesAffectedFoundMessagesTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pts",
			SchemaName: "pts",
		},
		{
			Name:       "PtsCount",
			SchemaName: "pts_count",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Messages",
			SchemaName: "messages",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *MessagesAffectedFoundMessages) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.affectedFoundMessages#ef8d3e6c as nil")
	}
	b.PutID(MessagesAffectedFoundMessagesTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *MessagesAffectedFoundMessages) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.affectedFoundMessages#ef8d3e6c as nil")
	}
	b.PutInt(a.Pts)
	b.PutInt(a.PtsCount)
	b.PutInt(a.Offset)
	b.PutVectorHeader(len(a.Messages))
	for _, v := range a.Messages {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAffectedFoundMessages) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.affectedFoundMessages#ef8d3e6c to nil")
	}
	if err := b.ConsumeID(MessagesAffectedFoundMessagesTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.affectedFoundMessages#ef8d3e6c: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *MessagesAffectedFoundMessages) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.affectedFoundMessages#ef8d3e6c to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedFoundMessages#ef8d3e6c: field pts: %w", err)
		}
		a.Pts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedFoundMessages#ef8d3e6c: field pts_count: %w", err)
		}
		a.PtsCount = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedFoundMessages#ef8d3e6c: field offset: %w", err)
		}
		a.Offset = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.affectedFoundMessages#ef8d3e6c: field messages: %w", err)
		}

		if headerLen > 0 {
			a.Messages = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode messages.affectedFoundMessages#ef8d3e6c: field messages: %w", err)
			}
			a.Messages = append(a.Messages, value)
		}
	}
	return nil
}

// GetPts returns value of Pts field.
func (a *MessagesAffectedFoundMessages) GetPts() (value int) {
	if a == nil {
		return
	}
	return a.Pts
}

// GetPtsCount returns value of PtsCount field.
func (a *MessagesAffectedFoundMessages) GetPtsCount() (value int) {
	if a == nil {
		return
	}
	return a.PtsCount
}

// GetOffset returns value of Offset field.
func (a *MessagesAffectedFoundMessages) GetOffset() (value int) {
	if a == nil {
		return
	}
	return a.Offset
}

// GetMessages returns value of Messages field.
func (a *MessagesAffectedFoundMessages) GetMessages() (value []int) {
	if a == nil {
		return
	}
	return a.Messages
}
