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

// ChatRevenueTransactionTypeEarnings represents TL type `chatRevenueTransactionTypeEarnings#e81ca488`.
type ChatRevenueTransactionTypeEarnings struct {
	// Point in time (Unix timestamp) when the earnings started
	StartDate int32
	// Point in time (Unix timestamp) when the earnings ended
	EndDate int32
}

// ChatRevenueTransactionTypeEarningsTypeID is TL type id of ChatRevenueTransactionTypeEarnings.
const ChatRevenueTransactionTypeEarningsTypeID = 0xe81ca488

// construct implements constructor of ChatRevenueTransactionTypeClass.
func (c ChatRevenueTransactionTypeEarnings) construct() ChatRevenueTransactionTypeClass { return &c }

// Ensuring interfaces in compile-time for ChatRevenueTransactionTypeEarnings.
var (
	_ bin.Encoder     = &ChatRevenueTransactionTypeEarnings{}
	_ bin.Decoder     = &ChatRevenueTransactionTypeEarnings{}
	_ bin.BareEncoder = &ChatRevenueTransactionTypeEarnings{}
	_ bin.BareDecoder = &ChatRevenueTransactionTypeEarnings{}

	_ ChatRevenueTransactionTypeClass = &ChatRevenueTransactionTypeEarnings{}
)

func (c *ChatRevenueTransactionTypeEarnings) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.StartDate == 0) {
		return false
	}
	if !(c.EndDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatRevenueTransactionTypeEarnings) String() string {
	if c == nil {
		return "ChatRevenueTransactionTypeEarnings(nil)"
	}
	type Alias ChatRevenueTransactionTypeEarnings
	return fmt.Sprintf("ChatRevenueTransactionTypeEarnings%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatRevenueTransactionTypeEarnings) TypeID() uint32 {
	return ChatRevenueTransactionTypeEarningsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatRevenueTransactionTypeEarnings) TypeName() string {
	return "chatRevenueTransactionTypeEarnings"
}

// TypeInfo returns info about TL type.
func (c *ChatRevenueTransactionTypeEarnings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatRevenueTransactionTypeEarnings",
		ID:   ChatRevenueTransactionTypeEarningsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StartDate",
			SchemaName: "start_date",
		},
		{
			Name:       "EndDate",
			SchemaName: "end_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatRevenueTransactionTypeEarnings) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactionTypeEarnings#e81ca488 as nil")
	}
	b.PutID(ChatRevenueTransactionTypeEarningsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatRevenueTransactionTypeEarnings) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactionTypeEarnings#e81ca488 as nil")
	}
	b.PutInt32(c.StartDate)
	b.PutInt32(c.EndDate)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatRevenueTransactionTypeEarnings) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactionTypeEarnings#e81ca488 to nil")
	}
	if err := b.ConsumeID(ChatRevenueTransactionTypeEarningsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatRevenueTransactionTypeEarnings#e81ca488: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatRevenueTransactionTypeEarnings) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactionTypeEarnings#e81ca488 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueTransactionTypeEarnings#e81ca488: field start_date: %w", err)
		}
		c.StartDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueTransactionTypeEarnings#e81ca488: field end_date: %w", err)
		}
		c.EndDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatRevenueTransactionTypeEarnings) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactionTypeEarnings#e81ca488 as nil")
	}
	b.ObjStart()
	b.PutID("chatRevenueTransactionTypeEarnings")
	b.Comma()
	b.FieldStart("start_date")
	b.PutInt32(c.StartDate)
	b.Comma()
	b.FieldStart("end_date")
	b.PutInt32(c.EndDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatRevenueTransactionTypeEarnings) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactionTypeEarnings#e81ca488 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatRevenueTransactionTypeEarnings"); err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactionTypeEarnings#e81ca488: %w", err)
			}
		case "start_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactionTypeEarnings#e81ca488: field start_date: %w", err)
			}
			c.StartDate = value
		case "end_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactionTypeEarnings#e81ca488: field end_date: %w", err)
			}
			c.EndDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStartDate returns value of StartDate field.
func (c *ChatRevenueTransactionTypeEarnings) GetStartDate() (value int32) {
	if c == nil {
		return
	}
	return c.StartDate
}

// GetEndDate returns value of EndDate field.
func (c *ChatRevenueTransactionTypeEarnings) GetEndDate() (value int32) {
	if c == nil {
		return
	}
	return c.EndDate
}

// ChatRevenueTransactionTypeWithdrawal represents TL type `chatRevenueTransactionTypeWithdrawal#6c72a9de`.
type ChatRevenueTransactionTypeWithdrawal struct {
	// Point in time (Unix timestamp) when the earnings withdrawal started
	WithdrawalDate int32
	// Name of the payment provider
	Provider string
	// State of the withdrawal
	State ChatRevenueWithdrawalStateClass
}

// ChatRevenueTransactionTypeWithdrawalTypeID is TL type id of ChatRevenueTransactionTypeWithdrawal.
const ChatRevenueTransactionTypeWithdrawalTypeID = 0x6c72a9de

// construct implements constructor of ChatRevenueTransactionTypeClass.
func (c ChatRevenueTransactionTypeWithdrawal) construct() ChatRevenueTransactionTypeClass { return &c }

// Ensuring interfaces in compile-time for ChatRevenueTransactionTypeWithdrawal.
var (
	_ bin.Encoder     = &ChatRevenueTransactionTypeWithdrawal{}
	_ bin.Decoder     = &ChatRevenueTransactionTypeWithdrawal{}
	_ bin.BareEncoder = &ChatRevenueTransactionTypeWithdrawal{}
	_ bin.BareDecoder = &ChatRevenueTransactionTypeWithdrawal{}

	_ ChatRevenueTransactionTypeClass = &ChatRevenueTransactionTypeWithdrawal{}
)

func (c *ChatRevenueTransactionTypeWithdrawal) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.WithdrawalDate == 0) {
		return false
	}
	if !(c.Provider == "") {
		return false
	}
	if !(c.State == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatRevenueTransactionTypeWithdrawal) String() string {
	if c == nil {
		return "ChatRevenueTransactionTypeWithdrawal(nil)"
	}
	type Alias ChatRevenueTransactionTypeWithdrawal
	return fmt.Sprintf("ChatRevenueTransactionTypeWithdrawal%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatRevenueTransactionTypeWithdrawal) TypeID() uint32 {
	return ChatRevenueTransactionTypeWithdrawalTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatRevenueTransactionTypeWithdrawal) TypeName() string {
	return "chatRevenueTransactionTypeWithdrawal"
}

// TypeInfo returns info about TL type.
func (c *ChatRevenueTransactionTypeWithdrawal) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatRevenueTransactionTypeWithdrawal",
		ID:   ChatRevenueTransactionTypeWithdrawalTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "WithdrawalDate",
			SchemaName: "withdrawal_date",
		},
		{
			Name:       "Provider",
			SchemaName: "provider",
		},
		{
			Name:       "State",
			SchemaName: "state",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatRevenueTransactionTypeWithdrawal) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactionTypeWithdrawal#6c72a9de as nil")
	}
	b.PutID(ChatRevenueTransactionTypeWithdrawalTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatRevenueTransactionTypeWithdrawal) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactionTypeWithdrawal#6c72a9de as nil")
	}
	b.PutInt32(c.WithdrawalDate)
	b.PutString(c.Provider)
	if c.State == nil {
		return fmt.Errorf("unable to encode chatRevenueTransactionTypeWithdrawal#6c72a9de: field state is nil")
	}
	if err := c.State.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatRevenueTransactionTypeWithdrawal#6c72a9de: field state: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatRevenueTransactionTypeWithdrawal) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactionTypeWithdrawal#6c72a9de to nil")
	}
	if err := b.ConsumeID(ChatRevenueTransactionTypeWithdrawalTypeID); err != nil {
		return fmt.Errorf("unable to decode chatRevenueTransactionTypeWithdrawal#6c72a9de: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatRevenueTransactionTypeWithdrawal) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactionTypeWithdrawal#6c72a9de to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueTransactionTypeWithdrawal#6c72a9de: field withdrawal_date: %w", err)
		}
		c.WithdrawalDate = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueTransactionTypeWithdrawal#6c72a9de: field provider: %w", err)
		}
		c.Provider = value
	}
	{
		value, err := DecodeChatRevenueWithdrawalState(b)
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueTransactionTypeWithdrawal#6c72a9de: field state: %w", err)
		}
		c.State = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatRevenueTransactionTypeWithdrawal) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactionTypeWithdrawal#6c72a9de as nil")
	}
	b.ObjStart()
	b.PutID("chatRevenueTransactionTypeWithdrawal")
	b.Comma()
	b.FieldStart("withdrawal_date")
	b.PutInt32(c.WithdrawalDate)
	b.Comma()
	b.FieldStart("provider")
	b.PutString(c.Provider)
	b.Comma()
	b.FieldStart("state")
	if c.State == nil {
		return fmt.Errorf("unable to encode chatRevenueTransactionTypeWithdrawal#6c72a9de: field state is nil")
	}
	if err := c.State.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatRevenueTransactionTypeWithdrawal#6c72a9de: field state: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatRevenueTransactionTypeWithdrawal) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactionTypeWithdrawal#6c72a9de to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatRevenueTransactionTypeWithdrawal"); err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactionTypeWithdrawal#6c72a9de: %w", err)
			}
		case "withdrawal_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactionTypeWithdrawal#6c72a9de: field withdrawal_date: %w", err)
			}
			c.WithdrawalDate = value
		case "provider":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactionTypeWithdrawal#6c72a9de: field provider: %w", err)
			}
			c.Provider = value
		case "state":
			value, err := DecodeTDLibJSONChatRevenueWithdrawalState(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactionTypeWithdrawal#6c72a9de: field state: %w", err)
			}
			c.State = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetWithdrawalDate returns value of WithdrawalDate field.
func (c *ChatRevenueTransactionTypeWithdrawal) GetWithdrawalDate() (value int32) {
	if c == nil {
		return
	}
	return c.WithdrawalDate
}

// GetProvider returns value of Provider field.
func (c *ChatRevenueTransactionTypeWithdrawal) GetProvider() (value string) {
	if c == nil {
		return
	}
	return c.Provider
}

// GetState returns value of State field.
func (c *ChatRevenueTransactionTypeWithdrawal) GetState() (value ChatRevenueWithdrawalStateClass) {
	if c == nil {
		return
	}
	return c.State
}

// ChatRevenueTransactionTypeRefund represents TL type `chatRevenueTransactionTypeRefund#1206b847`.
type ChatRevenueTransactionTypeRefund struct {
	// Point in time (Unix timestamp) when the transaction was refunded
	RefundDate int32
	// Name of the payment provider
	Provider string
}

// ChatRevenueTransactionTypeRefundTypeID is TL type id of ChatRevenueTransactionTypeRefund.
const ChatRevenueTransactionTypeRefundTypeID = 0x1206b847

// construct implements constructor of ChatRevenueTransactionTypeClass.
func (c ChatRevenueTransactionTypeRefund) construct() ChatRevenueTransactionTypeClass { return &c }

// Ensuring interfaces in compile-time for ChatRevenueTransactionTypeRefund.
var (
	_ bin.Encoder     = &ChatRevenueTransactionTypeRefund{}
	_ bin.Decoder     = &ChatRevenueTransactionTypeRefund{}
	_ bin.BareEncoder = &ChatRevenueTransactionTypeRefund{}
	_ bin.BareDecoder = &ChatRevenueTransactionTypeRefund{}

	_ ChatRevenueTransactionTypeClass = &ChatRevenueTransactionTypeRefund{}
)

func (c *ChatRevenueTransactionTypeRefund) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.RefundDate == 0) {
		return false
	}
	if !(c.Provider == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatRevenueTransactionTypeRefund) String() string {
	if c == nil {
		return "ChatRevenueTransactionTypeRefund(nil)"
	}
	type Alias ChatRevenueTransactionTypeRefund
	return fmt.Sprintf("ChatRevenueTransactionTypeRefund%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatRevenueTransactionTypeRefund) TypeID() uint32 {
	return ChatRevenueTransactionTypeRefundTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatRevenueTransactionTypeRefund) TypeName() string {
	return "chatRevenueTransactionTypeRefund"
}

// TypeInfo returns info about TL type.
func (c *ChatRevenueTransactionTypeRefund) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatRevenueTransactionTypeRefund",
		ID:   ChatRevenueTransactionTypeRefundTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RefundDate",
			SchemaName: "refund_date",
		},
		{
			Name:       "Provider",
			SchemaName: "provider",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatRevenueTransactionTypeRefund) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactionTypeRefund#1206b847 as nil")
	}
	b.PutID(ChatRevenueTransactionTypeRefundTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatRevenueTransactionTypeRefund) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactionTypeRefund#1206b847 as nil")
	}
	b.PutInt32(c.RefundDate)
	b.PutString(c.Provider)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatRevenueTransactionTypeRefund) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactionTypeRefund#1206b847 to nil")
	}
	if err := b.ConsumeID(ChatRevenueTransactionTypeRefundTypeID); err != nil {
		return fmt.Errorf("unable to decode chatRevenueTransactionTypeRefund#1206b847: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatRevenueTransactionTypeRefund) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactionTypeRefund#1206b847 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueTransactionTypeRefund#1206b847: field refund_date: %w", err)
		}
		c.RefundDate = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatRevenueTransactionTypeRefund#1206b847: field provider: %w", err)
		}
		c.Provider = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatRevenueTransactionTypeRefund) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatRevenueTransactionTypeRefund#1206b847 as nil")
	}
	b.ObjStart()
	b.PutID("chatRevenueTransactionTypeRefund")
	b.Comma()
	b.FieldStart("refund_date")
	b.PutInt32(c.RefundDate)
	b.Comma()
	b.FieldStart("provider")
	b.PutString(c.Provider)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatRevenueTransactionTypeRefund) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatRevenueTransactionTypeRefund#1206b847 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatRevenueTransactionTypeRefund"); err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactionTypeRefund#1206b847: %w", err)
			}
		case "refund_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactionTypeRefund#1206b847: field refund_date: %w", err)
			}
			c.RefundDate = value
		case "provider":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatRevenueTransactionTypeRefund#1206b847: field provider: %w", err)
			}
			c.Provider = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRefundDate returns value of RefundDate field.
func (c *ChatRevenueTransactionTypeRefund) GetRefundDate() (value int32) {
	if c == nil {
		return
	}
	return c.RefundDate
}

// GetProvider returns value of Provider field.
func (c *ChatRevenueTransactionTypeRefund) GetProvider() (value string) {
	if c == nil {
		return
	}
	return c.Provider
}

// ChatRevenueTransactionTypeClassName is schema name of ChatRevenueTransactionTypeClass.
const ChatRevenueTransactionTypeClassName = "ChatRevenueTransactionType"

// ChatRevenueTransactionTypeClass represents ChatRevenueTransactionType generic type.
//
// Example:
//
//	g, err := tdapi.DecodeChatRevenueTransactionType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.ChatRevenueTransactionTypeEarnings: // chatRevenueTransactionTypeEarnings#e81ca488
//	case *tdapi.ChatRevenueTransactionTypeWithdrawal: // chatRevenueTransactionTypeWithdrawal#6c72a9de
//	case *tdapi.ChatRevenueTransactionTypeRefund: // chatRevenueTransactionTypeRefund#1206b847
//	default: panic(v)
//	}
type ChatRevenueTransactionTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatRevenueTransactionTypeClass

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

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeChatRevenueTransactionType implements binary de-serialization for ChatRevenueTransactionTypeClass.
func DecodeChatRevenueTransactionType(buf *bin.Buffer) (ChatRevenueTransactionTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatRevenueTransactionTypeEarningsTypeID:
		// Decoding chatRevenueTransactionTypeEarnings#e81ca488.
		v := ChatRevenueTransactionTypeEarnings{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatRevenueTransactionTypeClass: %w", err)
		}
		return &v, nil
	case ChatRevenueTransactionTypeWithdrawalTypeID:
		// Decoding chatRevenueTransactionTypeWithdrawal#6c72a9de.
		v := ChatRevenueTransactionTypeWithdrawal{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatRevenueTransactionTypeClass: %w", err)
		}
		return &v, nil
	case ChatRevenueTransactionTypeRefundTypeID:
		// Decoding chatRevenueTransactionTypeRefund#1206b847.
		v := ChatRevenueTransactionTypeRefund{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatRevenueTransactionTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatRevenueTransactionTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONChatRevenueTransactionType implements binary de-serialization for ChatRevenueTransactionTypeClass.
func DecodeTDLibJSONChatRevenueTransactionType(buf tdjson.Decoder) (ChatRevenueTransactionTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "chatRevenueTransactionTypeEarnings":
		// Decoding chatRevenueTransactionTypeEarnings#e81ca488.
		v := ChatRevenueTransactionTypeEarnings{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatRevenueTransactionTypeClass: %w", err)
		}
		return &v, nil
	case "chatRevenueTransactionTypeWithdrawal":
		// Decoding chatRevenueTransactionTypeWithdrawal#6c72a9de.
		v := ChatRevenueTransactionTypeWithdrawal{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatRevenueTransactionTypeClass: %w", err)
		}
		return &v, nil
	case "chatRevenueTransactionTypeRefund":
		// Decoding chatRevenueTransactionTypeRefund#1206b847.
		v := ChatRevenueTransactionTypeRefund{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatRevenueTransactionTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatRevenueTransactionTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ChatRevenueTransactionType boxes the ChatRevenueTransactionTypeClass providing a helper.
type ChatRevenueTransactionTypeBox struct {
	ChatRevenueTransactionType ChatRevenueTransactionTypeClass
}

// Decode implements bin.Decoder for ChatRevenueTransactionTypeBox.
func (b *ChatRevenueTransactionTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatRevenueTransactionTypeBox to nil")
	}
	v, err := DecodeChatRevenueTransactionType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatRevenueTransactionType = v
	return nil
}

// Encode implements bin.Encode for ChatRevenueTransactionTypeBox.
func (b *ChatRevenueTransactionTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatRevenueTransactionType == nil {
		return fmt.Errorf("unable to encode ChatRevenueTransactionTypeClass as nil")
	}
	return b.ChatRevenueTransactionType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ChatRevenueTransactionTypeBox.
func (b *ChatRevenueTransactionTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatRevenueTransactionTypeBox to nil")
	}
	v, err := DecodeTDLibJSONChatRevenueTransactionType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatRevenueTransactionType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ChatRevenueTransactionTypeBox.
func (b *ChatRevenueTransactionTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ChatRevenueTransactionType == nil {
		return fmt.Errorf("unable to encode ChatRevenueTransactionTypeClass as nil")
	}
	return b.ChatRevenueTransactionType.EncodeTDLibJSON(buf)
}
