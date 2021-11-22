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

// MessageSendingStatePending represents TL type `messageSendingStatePending#ada359c2`.
type MessageSendingStatePending struct {
}

// MessageSendingStatePendingTypeID is TL type id of MessageSendingStatePending.
const MessageSendingStatePendingTypeID = 0xada359c2

// construct implements constructor of MessageSendingStateClass.
func (m MessageSendingStatePending) construct() MessageSendingStateClass { return &m }

// Ensuring interfaces in compile-time for MessageSendingStatePending.
var (
	_ bin.Encoder     = &MessageSendingStatePending{}
	_ bin.Decoder     = &MessageSendingStatePending{}
	_ bin.BareEncoder = &MessageSendingStatePending{}
	_ bin.BareDecoder = &MessageSendingStatePending{}

	_ MessageSendingStateClass = &MessageSendingStatePending{}
)

func (m *MessageSendingStatePending) Zero() bool {
	if m == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSendingStatePending) String() string {
	if m == nil {
		return "MessageSendingStatePending(nil)"
	}
	type Alias MessageSendingStatePending
	return fmt.Sprintf("MessageSendingStatePending%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSendingStatePending) TypeID() uint32 {
	return MessageSendingStatePendingTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSendingStatePending) TypeName() string {
	return "messageSendingStatePending"
}

// TypeInfo returns info about TL type.
func (m *MessageSendingStatePending) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSendingStatePending",
		ID:   MessageSendingStatePendingTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSendingStatePending) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStatePending#ada359c2 as nil")
	}
	b.PutID(MessageSendingStatePendingTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSendingStatePending) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStatePending#ada359c2 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSendingStatePending) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendingStatePending#ada359c2 to nil")
	}
	if err := b.ConsumeID(MessageSendingStatePendingTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSendingStatePending#ada359c2: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSendingStatePending) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendingStatePending#ada359c2 to nil")
	}
	return nil
}

// EncodeTDLibJSON encodes m in TDLib API JSON format.
func (m *MessageSendingStatePending) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStatePending#ada359c2 as nil")
	}
	b.ObjStart()
	b.PutID("messageSendingStatePending")
	b.ObjEnd()
	return nil
}

// MessageSendingStateFailed represents TL type `messageSendingStateFailed#7a74d137`.
type MessageSendingStateFailed struct {
	// An error code; 0 if unknown
	ErrorCode int32
	// Error message
	ErrorMessage string
	// True, if the message can be re-sent
	CanRetry bool
	// Time left before the message can be re-sent, in seconds. No update is sent when this
	// field changes
	RetryAfter float64
}

// MessageSendingStateFailedTypeID is TL type id of MessageSendingStateFailed.
const MessageSendingStateFailedTypeID = 0x7a74d137

// construct implements constructor of MessageSendingStateClass.
func (m MessageSendingStateFailed) construct() MessageSendingStateClass { return &m }

// Ensuring interfaces in compile-time for MessageSendingStateFailed.
var (
	_ bin.Encoder     = &MessageSendingStateFailed{}
	_ bin.Decoder     = &MessageSendingStateFailed{}
	_ bin.BareEncoder = &MessageSendingStateFailed{}
	_ bin.BareDecoder = &MessageSendingStateFailed{}

	_ MessageSendingStateClass = &MessageSendingStateFailed{}
)

func (m *MessageSendingStateFailed) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.ErrorCode == 0) {
		return false
	}
	if !(m.ErrorMessage == "") {
		return false
	}
	if !(m.CanRetry == false) {
		return false
	}
	if !(m.RetryAfter == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessageSendingStateFailed) String() string {
	if m == nil {
		return "MessageSendingStateFailed(nil)"
	}
	type Alias MessageSendingStateFailed
	return fmt.Sprintf("MessageSendingStateFailed%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessageSendingStateFailed) TypeID() uint32 {
	return MessageSendingStateFailedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessageSendingStateFailed) TypeName() string {
	return "messageSendingStateFailed"
}

// TypeInfo returns info about TL type.
func (m *MessageSendingStateFailed) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messageSendingStateFailed",
		ID:   MessageSendingStateFailedTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ErrorCode",
			SchemaName: "error_code",
		},
		{
			Name:       "ErrorMessage",
			SchemaName: "error_message",
		},
		{
			Name:       "CanRetry",
			SchemaName: "can_retry",
		},
		{
			Name:       "RetryAfter",
			SchemaName: "retry_after",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessageSendingStateFailed) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStateFailed#7a74d137 as nil")
	}
	b.PutID(MessageSendingStateFailedTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessageSendingStateFailed) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStateFailed#7a74d137 as nil")
	}
	b.PutInt32(m.ErrorCode)
	b.PutString(m.ErrorMessage)
	b.PutBool(m.CanRetry)
	b.PutDouble(m.RetryAfter)
	return nil
}

// Decode implements bin.Decoder.
func (m *MessageSendingStateFailed) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendingStateFailed#7a74d137 to nil")
	}
	if err := b.ConsumeID(MessageSendingStateFailedTypeID); err != nil {
		return fmt.Errorf("unable to decode messageSendingStateFailed#7a74d137: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessageSendingStateFailed) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messageSendingStateFailed#7a74d137 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendingStateFailed#7a74d137: field error_code: %w", err)
		}
		m.ErrorCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendingStateFailed#7a74d137: field error_message: %w", err)
		}
		m.ErrorMessage = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendingStateFailed#7a74d137: field can_retry: %w", err)
		}
		m.CanRetry = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode messageSendingStateFailed#7a74d137: field retry_after: %w", err)
		}
		m.RetryAfter = value
	}
	return nil
}

// EncodeTDLibJSON encodes m in TDLib API JSON format.
func (m *MessageSendingStateFailed) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if m == nil {
		return fmt.Errorf("can't encode messageSendingStateFailed#7a74d137 as nil")
	}
	b.ObjStart()
	b.PutID("messageSendingStateFailed")
	b.FieldStart("error_code")
	b.PutInt32(m.ErrorCode)
	b.FieldStart("error_message")
	b.PutString(m.ErrorMessage)
	b.FieldStart("can_retry")
	b.PutBool(m.CanRetry)
	b.FieldStart("retry_after")
	b.PutDouble(m.RetryAfter)
	b.ObjEnd()
	return nil
}

// GetErrorCode returns value of ErrorCode field.
func (m *MessageSendingStateFailed) GetErrorCode() (value int32) {
	return m.ErrorCode
}

// GetErrorMessage returns value of ErrorMessage field.
func (m *MessageSendingStateFailed) GetErrorMessage() (value string) {
	return m.ErrorMessage
}

// GetCanRetry returns value of CanRetry field.
func (m *MessageSendingStateFailed) GetCanRetry() (value bool) {
	return m.CanRetry
}

// GetRetryAfter returns value of RetryAfter field.
func (m *MessageSendingStateFailed) GetRetryAfter() (value float64) {
	return m.RetryAfter
}

// MessageSendingStateClass represents MessageSendingState generic type.
//
// Example:
//  g, err := tdapi.DecodeMessageSendingState(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.MessageSendingStatePending: // messageSendingStatePending#ada359c2
//  case *tdapi.MessageSendingStateFailed: // messageSendingStateFailed#7a74d137
//  default: panic(v)
//  }
type MessageSendingStateClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessageSendingStateClass

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
	EncodeTDLibJSON(b *jsontd.Encoder) error
}

// DecodeMessageSendingState implements binary de-serialization for MessageSendingStateClass.
func DecodeMessageSendingState(buf *bin.Buffer) (MessageSendingStateClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessageSendingStatePendingTypeID:
		// Decoding messageSendingStatePending#ada359c2.
		v := MessageSendingStatePending{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSendingStateClass: %w", err)
		}
		return &v, nil
	case MessageSendingStateFailedTypeID:
		// Decoding messageSendingStateFailed#7a74d137.
		v := MessageSendingStateFailed{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessageSendingStateClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessageSendingStateClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessageSendingState boxes the MessageSendingStateClass providing a helper.
type MessageSendingStateBox struct {
	MessageSendingState MessageSendingStateClass
}

// Decode implements bin.Decoder for MessageSendingStateBox.
func (b *MessageSendingStateBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessageSendingStateBox to nil")
	}
	v, err := DecodeMessageSendingState(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MessageSendingState = v
	return nil
}

// Encode implements bin.Encode for MessageSendingStateBox.
func (b *MessageSendingStateBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MessageSendingState == nil {
		return fmt.Errorf("unable to encode MessageSendingStateClass as nil")
	}
	return b.MessageSendingState.Encode(buf)
}
