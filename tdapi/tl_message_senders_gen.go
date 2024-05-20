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

// MessageSenders represents TL type `messageSenders#f6929bcc`.
type MessageSenders struct {
	// Approximate total number of messages senders found
	TotalCount int32
	// List of message senders
	Senders []MessageSenderClass
}

// MessageSendersTypeID is TL type id of MessageSenders.
const MessageSendersTypeID = 0xf6929bcc

// Ensuring interfaces in compile-time for MessageSenders.
var (
	_ bin.Encoder     = &MessageSenders{}
	_ bin.Decoder     = &MessageSenders{}
	_ bin.BareEncoder = &MessageSenders{}
	_ bin.BareDecoder = &MessageSenders{}
)

func (m *MessageSenders) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.TotalCount == 0) {
		return false
	}
	if !(m.Senders == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSenders) String() string {
	if m == nil {
		return "MessageSenders(nil)"
	}
	type Alias MessageSenders
	return fmt.Sprintf("MessageSenders%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSenders) TypeID() uint32 {
	return MessageSendersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSenders) TypeName() string {
	return "messageSenders"
}

// TypeInfo returns info about TL type.
func (m *MessageSenders) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSenders",
		ID:   MessageSendersTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "TotalCount",
			SchemaName: "total_count",
		},
		{
			Name:       "Senders",
			SchemaName: "senders",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSenders) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSenders#f6929bcc as nil")
	}
	b.PutID(MessageSendersTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSenders) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSenders#f6929bcc as nil")
	}
	b.PutInt32(m.TotalCount)
	b.PutInt(len(m.Senders))
	for idx, v := range m.Senders {
		if v == nil {
			return fmt.Errorf("unable to encode messageSenders#f6929bcc: field senders element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare messageSenders#f6929bcc: field senders element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSenders) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSenders#f6929bcc to nil")
	}
	if err := b.ConsumeID(MessageSendersTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSenders#f6929bcc: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSenders) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSenders#f6929bcc to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageSenders#f6929bcc: field total_count: %w", err)
		}
		m.TotalCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageSenders#f6929bcc: field senders: %w", err)
		}

		if headerLen > 0 {
			m.Senders = make([]MessageSenderClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode messageSenders#f6929bcc: field senders: %w", err)
			}
			m.Senders = append(m.Senders, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (m *MessageSenders) EncodeTDLibJSON(b tdjson.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSenders#f6929bcc as nil")
	}
	b.ObjStart()
	b.PutID("messageSenders")
	b.Comma()
	b.FieldStart("total_count")
	b.PutInt32(m.TotalCount)
	b.Comma()
	b.FieldStart("senders")
	b.ArrStart()
	for idx, v := range m.Senders {
		if v == nil {
			return fmt.Errorf("unable to encode messageSenders#f6929bcc: field senders element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode messageSenders#f6929bcc: field senders element with index %d: %w", idx, err)
		}
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
func (m *MessageSenders) DecodeTDLibJSON(b tdjson.Decoder) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSenders#f6929bcc to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("messageSenders"); err != nil {
				return fmt.Errorf("unable to decode messageSenders#f6929bcc: %w", err)
			}
		case "total_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode messageSenders#f6929bcc: field total_count: %w", err)
			}
			m.TotalCount = value
		case "senders":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONMessageSender(b)
				if err != nil {
					return fmt.Errorf("unable to decode messageSenders#f6929bcc: field senders: %w", err)
				}
				m.Senders = append(m.Senders, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode messageSenders#f6929bcc: field senders: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetTotalCount returns value of TotalCount field.
func (m *MessageSenders) GetTotalCount() (value int32) {
	if m == nil {
		return
	}
	return m.TotalCount
}

// GetSenders returns value of Senders field.
func (m *MessageSenders) GetSenders() (value []MessageSenderClass) {
	if m == nil {
		return
	}
	return m.Senders
}
