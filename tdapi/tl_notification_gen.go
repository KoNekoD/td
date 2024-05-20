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

// Notification represents TL type `notification#2f0343d0`.
type Notification struct {
	// Unique persistent identifier of this notification
	ID int32
	// Notification date
	Date int32
	// True, if the notification was explicitly sent without sound
	IsSilent bool
	// Notification type
	Type NotificationTypeClass
}

// NotificationTypeID is TL type id of Notification.
const NotificationTypeID = 0x2f0343d0

// Ensuring interfaces in compile-time for Notification.
var (
	_ bin.Encoder     = &Notification{}
	_ bin.Decoder     = &Notification{}
	_ bin.BareEncoder = &Notification{}
	_ bin.BareDecoder = &Notification{}
)

func (n *Notification) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.ID == 0) {
		return false
	}
	if !(n.Date == 0) {
		return false
	}
	if !(n.IsSilent == false) {
		return false
	}
	if !(n.Type == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *Notification) String() string {
	if n == nil {
		return "Notification(nil)"
	}
	type Alias Notification
	return fmt.Sprintf("Notification%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Notification) TypeID() uint32 {
	return NotificationTypeID
}

// TypeName returns name of type in TL schema.
func (*Notification) TypeName() string {
	return "notification"
}

// TypeInfo returns info about TL type.
func (n *Notification) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "notification",
		ID:   NotificationTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "IsSilent",
			SchemaName: "is_silent",
		},
		{
			Name:       "Type",
			SchemaName: "type",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *Notification) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notification#2f0343d0 as nil")
	}
	b.PutID(NotificationTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *Notification) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode notification#2f0343d0 as nil")
	}
	b.PutInt32(n.ID)
	b.PutInt32(n.Date)
	b.PutBool(n.IsSilent)
	if n.Type == nil {
		return fmt.Errorf("unable to encode notification#2f0343d0: field type is nil")
	}
	if err := n.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode notification#2f0343d0: field type: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (n *Notification) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notification#2f0343d0 to nil")
	}
	if err := b.ConsumeID(NotificationTypeID); err != nil {
		return fmt.Errorf("unable to decode notification#2f0343d0: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *Notification) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode notification#2f0343d0 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode notification#2f0343d0: field id: %w", err)
		}
		n.ID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode notification#2f0343d0: field date: %w", err)
		}
		n.Date = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode notification#2f0343d0: field is_silent: %w", err)
		}
		n.IsSilent = value
	}
	{
		value, err := DecodeNotificationType(b)
		if err != nil {
			return fmt.Errorf("unable to decode notification#2f0343d0: field type: %w", err)
		}
		n.Type = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (n *Notification) EncodeTDLibJSON(b tdjson.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode notification#2f0343d0 as nil")
	}
	b.ObjStart()
	b.PutID("notification")
	b.Comma()
	b.FieldStart("id")
	b.PutInt32(n.ID)
	b.Comma()
	b.FieldStart("date")
	b.PutInt32(n.Date)
	b.Comma()
	b.FieldStart("is_silent")
	b.PutBool(n.IsSilent)
	b.Comma()
	b.FieldStart("type")
	if n.Type == nil {
		return fmt.Errorf("unable to encode notification#2f0343d0: field type is nil")
	}
	if err := n.Type.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode notification#2f0343d0: field type: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (n *Notification) DecodeTDLibJSON(b tdjson.Decoder) error {
	if n == nil {
		return fmt.Errorf("can't decode notification#2f0343d0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("notification"); err != nil {
				return fmt.Errorf("unable to decode notification#2f0343d0: %w", err)
			}
		case "id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode notification#2f0343d0: field id: %w", err)
			}
			n.ID = value
		case "date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode notification#2f0343d0: field date: %w", err)
			}
			n.Date = value
		case "is_silent":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode notification#2f0343d0: field is_silent: %w", err)
			}
			n.IsSilent = value
		case "type":
			value, err := DecodeTDLibJSONNotificationType(b)
			if err != nil {
				return fmt.Errorf("unable to decode notification#2f0343d0: field type: %w", err)
			}
			n.Type = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (n *Notification) GetID() (value int32) {
	if n == nil {
		return
	}
	return n.ID
}

// GetDate returns value of Date field.
func (n *Notification) GetDate() (value int32) {
	if n == nil {
		return
	}
	return n.Date
}

// GetIsSilent returns value of IsSilent field.
func (n *Notification) GetIsSilent() (value bool) {
	if n == nil {
		return
	}
	return n.IsSilent
}

// GetType returns value of Type field.
func (n *Notification) GetType() (value NotificationTypeClass) {
	if n == nil {
		return
	}
	return n.Type
}
