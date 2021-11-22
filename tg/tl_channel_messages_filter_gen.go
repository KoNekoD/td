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

// ChannelMessagesFilterEmpty represents TL type `channelMessagesFilterEmpty#94d42ee7`.
// No filter
//
// See https://core.telegram.org/constructor/channelMessagesFilterEmpty for reference.
type ChannelMessagesFilterEmpty struct {
}

// ChannelMessagesFilterEmptyTypeID is TL type id of ChannelMessagesFilterEmpty.
const ChannelMessagesFilterEmptyTypeID = 0x94d42ee7

// construct implements constructor of ChannelMessagesFilterClass.
func (c ChannelMessagesFilterEmpty) construct() ChannelMessagesFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelMessagesFilterEmpty.
var (
	_ bin.Encoder     = &ChannelMessagesFilterEmpty{}
	_ bin.Decoder     = &ChannelMessagesFilterEmpty{}
	_ bin.BareEncoder = &ChannelMessagesFilterEmpty{}
	_ bin.BareDecoder = &ChannelMessagesFilterEmpty{}

	_ ChannelMessagesFilterClass = &ChannelMessagesFilterEmpty{}
)

func (c *ChannelMessagesFilterEmpty) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelMessagesFilterEmpty) String() string {
	if c == nil {
		return "ChannelMessagesFilterEmpty(nil)"
	}
	type Alias ChannelMessagesFilterEmpty
	return fmt.Sprintf("ChannelMessagesFilterEmpty%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelMessagesFilterEmpty) TypeID() uint32 {
	return ChannelMessagesFilterEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelMessagesFilterEmpty) TypeName() string {
	return "channelMessagesFilterEmpty"
}

// TypeInfo returns info about TL type.
func (c *ChannelMessagesFilterEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelMessagesFilterEmpty",
		ID:   ChannelMessagesFilterEmptyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelMessagesFilterEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelMessagesFilterEmpty#94d42ee7 as nil")
	}
	b.PutID(ChannelMessagesFilterEmptyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelMessagesFilterEmpty) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelMessagesFilterEmpty#94d42ee7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelMessagesFilterEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelMessagesFilterEmpty#94d42ee7 to nil")
	}
	if err := b.ConsumeID(ChannelMessagesFilterEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode channelMessagesFilterEmpty#94d42ee7: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelMessagesFilterEmpty) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelMessagesFilterEmpty#94d42ee7 to nil")
	}
	return nil
}

// ChannelMessagesFilter represents TL type `channelMessagesFilter#cd77d957`.
// Filter for getting only certain types of channel messages
//
// See https://core.telegram.org/constructor/channelMessagesFilter for reference.
type ChannelMessagesFilter struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to exclude new messages from the search
	ExcludeNewMessages bool
	// A range of messages to fetch
	Ranges []MessageRange
}

// ChannelMessagesFilterTypeID is TL type id of ChannelMessagesFilter.
const ChannelMessagesFilterTypeID = 0xcd77d957

// construct implements constructor of ChannelMessagesFilterClass.
func (c ChannelMessagesFilter) construct() ChannelMessagesFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelMessagesFilter.
var (
	_ bin.Encoder     = &ChannelMessagesFilter{}
	_ bin.Decoder     = &ChannelMessagesFilter{}
	_ bin.BareEncoder = &ChannelMessagesFilter{}
	_ bin.BareDecoder = &ChannelMessagesFilter{}

	_ ChannelMessagesFilterClass = &ChannelMessagesFilter{}
)

func (c *ChannelMessagesFilter) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.ExcludeNewMessages == false) {
		return false
	}
	if !(c.Ranges == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelMessagesFilter) String() string {
	if c == nil {
		return "ChannelMessagesFilter(nil)"
	}
	type Alias ChannelMessagesFilter
	return fmt.Sprintf("ChannelMessagesFilter%+v", Alias(*c))
}

// FillFrom fills ChannelMessagesFilter from given interface.
func (c *ChannelMessagesFilter) FillFrom(from interface {
	GetExcludeNewMessages() (value bool)
	GetRanges() (value []MessageRange)
}) {
	c.ExcludeNewMessages = from.GetExcludeNewMessages()
	c.Ranges = from.GetRanges()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelMessagesFilter) TypeID() uint32 {
	return ChannelMessagesFilterTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelMessagesFilter) TypeName() string {
	return "channelMessagesFilter"
}

// TypeInfo returns info about TL type.
func (c *ChannelMessagesFilter) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channelMessagesFilter",
		ID:   ChannelMessagesFilterTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ExcludeNewMessages",
			SchemaName: "exclude_new_messages",
			Null:       !c.Flags.Has(1),
		},
		{
			Name:       "Ranges",
			SchemaName: "ranges",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelMessagesFilter) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelMessagesFilter#cd77d957 as nil")
	}
	b.PutID(ChannelMessagesFilterTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelMessagesFilter) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelMessagesFilter#cd77d957 as nil")
	}
	if !(c.ExcludeNewMessages == false) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelMessagesFilter#cd77d957: field flags: %w", err)
	}
	b.PutVectorHeader(len(c.Ranges))
	for idx, v := range c.Ranges {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channelMessagesFilter#cd77d957: field ranges element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelMessagesFilter) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelMessagesFilter#cd77d957 to nil")
	}
	if err := b.ConsumeID(ChannelMessagesFilterTypeID); err != nil {
		return fmt.Errorf("unable to decode channelMessagesFilter#cd77d957: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelMessagesFilter) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelMessagesFilter#cd77d957 to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channelMessagesFilter#cd77d957: field flags: %w", err)
		}
	}
	c.ExcludeNewMessages = c.Flags.Has(1)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channelMessagesFilter#cd77d957: field ranges: %w", err)
		}

		if headerLen > 0 {
			c.Ranges = make([]MessageRange, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessageRange
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode channelMessagesFilter#cd77d957: field ranges: %w", err)
			}
			c.Ranges = append(c.Ranges, value)
		}
	}
	return nil
}

// SetExcludeNewMessages sets value of ExcludeNewMessages conditional field.
func (c *ChannelMessagesFilter) SetExcludeNewMessages(value bool) {
	if value {
		c.Flags.Set(1)
		c.ExcludeNewMessages = true
	} else {
		c.Flags.Unset(1)
		c.ExcludeNewMessages = false
	}
}

// GetExcludeNewMessages returns value of ExcludeNewMessages conditional field.
func (c *ChannelMessagesFilter) GetExcludeNewMessages() (value bool) {
	return c.Flags.Has(1)
}

// GetRanges returns value of Ranges field.
func (c *ChannelMessagesFilter) GetRanges() (value []MessageRange) {
	return c.Ranges
}

// ChannelMessagesFilterClass represents ChannelMessagesFilter generic type.
//
// See https://core.telegram.org/type/ChannelMessagesFilter for reference.
//
// Example:
//  g, err := tg.DecodeChannelMessagesFilter(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.ChannelMessagesFilterEmpty: // channelMessagesFilterEmpty#94d42ee7
//  case *tg.ChannelMessagesFilter: // channelMessagesFilter#cd77d957
//  default: panic(v)
//  }
type ChannelMessagesFilterClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChannelMessagesFilterClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsNotEmpty tries to map ChannelMessagesFilterClass to ChannelMessagesFilter.
	AsNotEmpty() (*ChannelMessagesFilter, bool)
}

// AsNotEmpty tries to map ChannelMessagesFilterEmpty to ChannelMessagesFilter.
func (c *ChannelMessagesFilterEmpty) AsNotEmpty() (*ChannelMessagesFilter, bool) {
	return nil, false
}

// AsNotEmpty tries to map ChannelMessagesFilter to ChannelMessagesFilter.
func (c *ChannelMessagesFilter) AsNotEmpty() (*ChannelMessagesFilter, bool) {
	return c, true
}

// DecodeChannelMessagesFilter implements binary de-serialization for ChannelMessagesFilterClass.
func DecodeChannelMessagesFilter(buf *bin.Buffer) (ChannelMessagesFilterClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChannelMessagesFilterEmptyTypeID:
		// Decoding channelMessagesFilterEmpty#94d42ee7.
		v := ChannelMessagesFilterEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelMessagesFilterClass: %w", err)
		}
		return &v, nil
	case ChannelMessagesFilterTypeID:
		// Decoding channelMessagesFilter#cd77d957.
		v := ChannelMessagesFilter{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelMessagesFilterClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChannelMessagesFilterClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChannelMessagesFilter boxes the ChannelMessagesFilterClass providing a helper.
type ChannelMessagesFilterBox struct {
	ChannelMessagesFilter ChannelMessagesFilterClass
}

// Decode implements bin.Decoder for ChannelMessagesFilterBox.
func (b *ChannelMessagesFilterBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChannelMessagesFilterBox to nil")
	}
	v, err := DecodeChannelMessagesFilter(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChannelMessagesFilter = v
	return nil
}

// Encode implements bin.Encode for ChannelMessagesFilterBox.
func (b *ChannelMessagesFilterBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChannelMessagesFilter == nil {
		return fmt.Errorf("unable to encode ChannelMessagesFilterClass as nil")
	}
	return b.ChannelMessagesFilter.Encode(buf)
}
