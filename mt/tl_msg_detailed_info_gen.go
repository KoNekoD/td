// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// MsgDetailedInfo represents TL type `msg_detailed_info#276d3ec6`.
type MsgDetailedInfo struct {
	// MsgID field of MsgDetailedInfo.
	MsgID int64
	// AnswerMsgID field of MsgDetailedInfo.
	AnswerMsgID int64
	// Bytes field of MsgDetailedInfo.
	Bytes int
	// Status field of MsgDetailedInfo.
	Status int
}

// MsgDetailedInfoTypeID is TL type id of MsgDetailedInfo.
const MsgDetailedInfoTypeID = 0x276d3ec6

// construct implements constructor of MsgDetailedInfoClass.
func (m MsgDetailedInfo) construct() MsgDetailedInfoClass { return &m }

// Ensuring interfaces in compile-time for MsgDetailedInfo.
var (
	_ bin.Encoder     = &MsgDetailedInfo{}
	_ bin.Decoder     = &MsgDetailedInfo{}
	_ bin.BareEncoder = &MsgDetailedInfo{}
	_ bin.BareDecoder = &MsgDetailedInfo{}

	_ MsgDetailedInfoClass = &MsgDetailedInfo{}
)

func (m *MsgDetailedInfo) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.MsgID == 0) {
		return false
	}
	if !(m.AnswerMsgID == 0) {
		return false
	}
	if !(m.Bytes == 0) {
		return false
	}
	if !(m.Status == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgDetailedInfo) String() string {
	if m == nil {
		return "MsgDetailedInfo(nil)"
	}
	type Alias MsgDetailedInfo
	return fmt.Sprintf("MsgDetailedInfo%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MsgDetailedInfo) TypeID() uint32 {
	return MsgDetailedInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*MsgDetailedInfo) TypeName() string {
	return "msg_detailed_info"
}

// TypeInfo returns info about TL type.
func (m *MsgDetailedInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "msg_detailed_info",
		ID:   MsgDetailedInfoTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MsgID",
			SchemaName: "msg_id",
		},
		{
			Name:       "AnswerMsgID",
			SchemaName: "answer_msg_id",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
		{
			Name:       "Status",
			SchemaName: "status",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MsgDetailedInfo) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msg_detailed_info#276d3ec6 as nil")
	}
	b.PutID(MsgDetailedInfoTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MsgDetailedInfo) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msg_detailed_info#276d3ec6 as nil")
	}
	b.PutLong(m.MsgID)
	b.PutLong(m.AnswerMsgID)
	b.PutInt(m.Bytes)
	b.PutInt(m.Status)
	return nil
}

// Decode implements bin.Decoder.
func (m *MsgDetailedInfo) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msg_detailed_info#276d3ec6 to nil")
	}
	if err := b.ConsumeID(MsgDetailedInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode msg_detailed_info#276d3ec6: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MsgDetailedInfo) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msg_detailed_info#276d3ec6 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode msg_detailed_info#276d3ec6: field msg_id: %w", err)
		}
		m.MsgID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode msg_detailed_info#276d3ec6: field answer_msg_id: %w", err)
		}
		m.AnswerMsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode msg_detailed_info#276d3ec6: field bytes: %w", err)
		}
		m.Bytes = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode msg_detailed_info#276d3ec6: field status: %w", err)
		}
		m.Status = value
	}
	return nil
}

// GetMsgID returns value of MsgID field.
func (m *MsgDetailedInfo) GetMsgID() (value int64) {
	if m == nil {
		return
	}
	return m.MsgID
}

// GetAnswerMsgID returns value of AnswerMsgID field.
func (m *MsgDetailedInfo) GetAnswerMsgID() (value int64) {
	if m == nil {
		return
	}
	return m.AnswerMsgID
}

// GetBytes returns value of Bytes field.
func (m *MsgDetailedInfo) GetBytes() (value int) {
	if m == nil {
		return
	}
	return m.Bytes
}

// GetStatus returns value of Status field.
func (m *MsgDetailedInfo) GetStatus() (value int) {
	if m == nil {
		return
	}
	return m.Status
}

// MsgNewDetailedInfo represents TL type `msg_new_detailed_info#809db6df`.
type MsgNewDetailedInfo struct {
	// AnswerMsgID field of MsgNewDetailedInfo.
	AnswerMsgID int64
	// Bytes field of MsgNewDetailedInfo.
	Bytes int
	// Status field of MsgNewDetailedInfo.
	Status int
}

// MsgNewDetailedInfoTypeID is TL type id of MsgNewDetailedInfo.
const MsgNewDetailedInfoTypeID = 0x809db6df

// construct implements constructor of MsgDetailedInfoClass.
func (m MsgNewDetailedInfo) construct() MsgDetailedInfoClass { return &m }

// Ensuring interfaces in compile-time for MsgNewDetailedInfo.
var (
	_ bin.Encoder     = &MsgNewDetailedInfo{}
	_ bin.Decoder     = &MsgNewDetailedInfo{}
	_ bin.BareEncoder = &MsgNewDetailedInfo{}
	_ bin.BareDecoder = &MsgNewDetailedInfo{}

	_ MsgDetailedInfoClass = &MsgNewDetailedInfo{}
)

func (m *MsgNewDetailedInfo) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.AnswerMsgID == 0) {
		return false
	}
	if !(m.Bytes == 0) {
		return false
	}
	if !(m.Status == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MsgNewDetailedInfo) String() string {
	if m == nil {
		return "MsgNewDetailedInfo(nil)"
	}
	type Alias MsgNewDetailedInfo
	return fmt.Sprintf("MsgNewDetailedInfo%+v", Alias(*m))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MsgNewDetailedInfo) TypeID() uint32 {
	return MsgNewDetailedInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*MsgNewDetailedInfo) TypeName() string {
	return "msg_new_detailed_info"
}

// TypeInfo returns info about TL type.
func (m *MsgNewDetailedInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "msg_new_detailed_info",
		ID:   MsgNewDetailedInfoTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AnswerMsgID",
			SchemaName: "answer_msg_id",
		},
		{
			Name:       "Bytes",
			SchemaName: "bytes",
		},
		{
			Name:       "Status",
			SchemaName: "status",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MsgNewDetailedInfo) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msg_new_detailed_info#809db6df as nil")
	}
	b.PutID(MsgNewDetailedInfoTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MsgNewDetailedInfo) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode msg_new_detailed_info#809db6df as nil")
	}
	b.PutLong(m.AnswerMsgID)
	b.PutInt(m.Bytes)
	b.PutInt(m.Status)
	return nil
}

// Decode implements bin.Decoder.
func (m *MsgNewDetailedInfo) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msg_new_detailed_info#809db6df to nil")
	}
	if err := b.ConsumeID(MsgNewDetailedInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode msg_new_detailed_info#809db6df: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MsgNewDetailedInfo) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode msg_new_detailed_info#809db6df to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode msg_new_detailed_info#809db6df: field answer_msg_id: %w", err)
		}
		m.AnswerMsgID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode msg_new_detailed_info#809db6df: field bytes: %w", err)
		}
		m.Bytes = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode msg_new_detailed_info#809db6df: field status: %w", err)
		}
		m.Status = value
	}
	return nil
}

// GetAnswerMsgID returns value of AnswerMsgID field.
func (m *MsgNewDetailedInfo) GetAnswerMsgID() (value int64) {
	if m == nil {
		return
	}
	return m.AnswerMsgID
}

// GetBytes returns value of Bytes field.
func (m *MsgNewDetailedInfo) GetBytes() (value int) {
	if m == nil {
		return
	}
	return m.Bytes
}

// GetStatus returns value of Status field.
func (m *MsgNewDetailedInfo) GetStatus() (value int) {
	if m == nil {
		return
	}
	return m.Status
}

// MsgDetailedInfoClassName is schema name of MsgDetailedInfoClass.
const MsgDetailedInfoClassName = "MsgDetailedInfo"

// MsgDetailedInfoClass represents MsgDetailedInfo generic type.
//
// Example:
//
//	g, err := mt.DecodeMsgDetailedInfo(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *mt.MsgDetailedInfo: // msg_detailed_info#276d3ec6
//	case *mt.MsgNewDetailedInfo: // msg_new_detailed_info#809db6df
//	default: panic(v)
//	}
type MsgDetailedInfoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MsgDetailedInfoClass

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

	// AnswerMsgID field of MsgDetailedInfo.
	GetAnswerMsgID() (value int64)
	// Bytes field of MsgDetailedInfo.
	GetBytes() (value int)
	// Status field of MsgDetailedInfo.
	GetStatus() (value int)
}

// DecodeMsgDetailedInfo implements binary de-serialization for MsgDetailedInfoClass.
func DecodeMsgDetailedInfo(buf *bin.Buffer) (MsgDetailedInfoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MsgDetailedInfoTypeID:
		// Decoding msg_detailed_info#276d3ec6.
		v := MsgDetailedInfo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MsgDetailedInfoClass: %w", err)
		}
		return &v, nil
	case MsgNewDetailedInfoTypeID:
		// Decoding msg_new_detailed_info#809db6df.
		v := MsgNewDetailedInfo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MsgDetailedInfoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MsgDetailedInfoClass: %w", bin.NewUnexpectedID(id))
	}
}

// MsgDetailedInfo boxes the MsgDetailedInfoClass providing a helper.
type MsgDetailedInfoBox struct {
	MsgDetailedInfo MsgDetailedInfoClass
}

// Decode implements bin.Decoder for MsgDetailedInfoBox.
func (b *MsgDetailedInfoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MsgDetailedInfoBox to nil")
	}
	v, err := DecodeMsgDetailedInfo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.MsgDetailedInfo = v
	return nil
}

// Encode implements bin.Encode for MsgDetailedInfoBox.
func (b *MsgDetailedInfoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.MsgDetailedInfo == nil {
		return fmt.Errorf("unable to encode MsgDetailedInfoClass as nil")
	}
	return b.MsgDetailedInfo.Encode(buf)
}
