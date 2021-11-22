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

// MessageReplyInfo represents TL type `messageReplyInfo#d1b3673b`.
type MessageReplyInfo struct {
	// Number of times the message was directly or indirectly replied
	ReplyCount int32
	// Recent repliers to the message; available in channels with a discussion supergroup
	RecentRepliers []MessageSenderClass
	// Identifier of the last read incoming reply to the message
	LastReadInboxMessageID int64
	// Identifier of the last read outgoing reply to the message
	LastReadOutboxMessageID int64
	// Identifier of the last reply to the message
	LastMessageID int64
}

// MessageReplyInfoTypeID is TL type id of MessageReplyInfo.
const MessageReplyInfoTypeID = 0xd1b3673b

// Ensuring interfaces in compile-time for MessageReplyInfo.
var (
	_ bin.Encoder     = &MessageReplyInfo{}
	_ bin.Decoder     = &MessageReplyInfo{}
	_ bin.BareEncoder = &MessageReplyInfo{}
	_ bin.BareDecoder = &MessageReplyInfo{}
)

func (m *MessageReplyInfo) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ReplyCount == 0) {
		return false
	}
	if !(m.RecentRepliers == nil) {
		return false
	}
	if !(m.LastReadInboxMessageID == 0) {
		return false
	}
	if !(m.LastReadOutboxMessageID == 0) {
		return false
	}
	if !(m.LastMessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageReplyInfo) String() string {
	if m == nil {
		return "MessageReplyInfo(nil)"
	}
	type Alias MessageReplyInfo
	return fmt.Sprintf("MessageReplyInfo%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageReplyInfo) TypeID() uint32 {
	return MessageReplyInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageReplyInfo) TypeName() string {
	return "messageReplyInfo"
}

// TypeInfo returns info about TL type.
func (m *MessageReplyInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageReplyInfo",
		ID:   MessageReplyInfoTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ReplyCount",
			SchemaName: "reply_count",
		},
		{
			Name:       "RecentRepliers",
			SchemaName: "recent_repliers",
		},
		{
			Name:       "LastReadInboxMessageID",
			SchemaName: "last_read_inbox_message_id",
		},
		{
			Name:       "LastReadOutboxMessageID",
			SchemaName: "last_read_outbox_message_id",
		},
		{
			Name:       "LastMessageID",
			SchemaName: "last_message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageReplyInfo) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyInfo#d1b3673b as nil")
	}
	b.PutID(MessageReplyInfoTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageReplyInfo) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyInfo#d1b3673b as nil")
	}
	b.PutInt32(m.ReplyCount)
	b.PutInt(len(m.RecentRepliers))
	for idx, v := range m.RecentRepliers {
		if v == nil {
			return fmt.Errorf("unable to encode messageReplyInfo#d1b3673b: field recent_repliers element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare messageReplyInfo#d1b3673b: field recent_repliers element with index %d: %w", idx, err)
		}
	}
	b.PutLong(m.LastReadInboxMessageID)
	b.PutLong(m.LastReadOutboxMessageID)
	b.PutLong(m.LastMessageID)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageReplyInfo) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyInfo#d1b3673b to nil")
	}
	if err := b.ConsumeID(MessageReplyInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode messageReplyInfo#d1b3673b: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageReplyInfo) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageReplyInfo#d1b3673b to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyInfo#d1b3673b: field reply_count: %w", err)
		}
		m.ReplyCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyInfo#d1b3673b: field recent_repliers: %w", err)
		}

		if headerLen > 0 {
			m.RecentRepliers = make([]MessageSenderClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageSender(b)
			if err != nil {
				return fmt.Errorf("unable to decode messageReplyInfo#d1b3673b: field recent_repliers: %w", err)
			}
			m.RecentRepliers = append(m.RecentRepliers, value)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyInfo#d1b3673b: field last_read_inbox_message_id: %w", err)
		}
		m.LastReadInboxMessageID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyInfo#d1b3673b: field last_read_outbox_message_id: %w", err)
		}
		m.LastReadOutboxMessageID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode messageReplyInfo#d1b3673b: field last_message_id: %w", err)
		}
		m.LastMessageID = value
	}
	return nil
}

// EncodeTDLibJSON encodes m in TDLib API JSON format.
func (m *MessageReplyInfo) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageReplyInfo#d1b3673b as nil")
	}
	b.ObjStart()
	b.PutID("messageReplyInfo")
	b.FieldStart("reply_count")
	b.PutInt32(m.ReplyCount)
	b.FieldStart("recent_repliers")
	b.ArrStart()
	for idx, v := range m.RecentRepliers {
		if v == nil {
			return fmt.Errorf("unable to encode messageReplyInfo#d1b3673b: field recent_repliers element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode messageReplyInfo#d1b3673b: field recent_repliers element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.FieldStart("last_read_inbox_message_id")
	b.PutLong(m.LastReadInboxMessageID)
	b.FieldStart("last_read_outbox_message_id")
	b.PutLong(m.LastReadOutboxMessageID)
	b.FieldStart("last_message_id")
	b.PutLong(m.LastMessageID)
	b.ObjEnd()
	return nil
}

// GetReplyCount returns value of ReplyCount field.
func (m *MessageReplyInfo) GetReplyCount() (value int32) {
	return m.ReplyCount
}

// GetRecentRepliers returns value of RecentRepliers field.
func (m *MessageReplyInfo) GetRecentRepliers() (value []MessageSenderClass) {
	return m.RecentRepliers
}

// GetLastReadInboxMessageID returns value of LastReadInboxMessageID field.
func (m *MessageReplyInfo) GetLastReadInboxMessageID() (value int64) {
	return m.LastReadInboxMessageID
}

// GetLastReadOutboxMessageID returns value of LastReadOutboxMessageID field.
func (m *MessageReplyInfo) GetLastReadOutboxMessageID() (value int64) {
	return m.LastReadOutboxMessageID
}

// GetLastMessageID returns value of LastMessageID field.
func (m *MessageReplyInfo) GetLastMessageID() (value int64) {
	return m.LastMessageID
}
