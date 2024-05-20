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

// BusinessAwayMessageScheduleAlways represents TL type `businessAwayMessageScheduleAlways#c9b9e2b9`.
//
// See https://core.telegram.org/constructor/businessAwayMessageScheduleAlways for reference.
type BusinessAwayMessageScheduleAlways struct {
}

// BusinessAwayMessageScheduleAlwaysTypeID is TL type id of BusinessAwayMessageScheduleAlways.
const BusinessAwayMessageScheduleAlwaysTypeID = 0xc9b9e2b9

// construct implements constructor of BusinessAwayMessageScheduleClass.
func (b BusinessAwayMessageScheduleAlways) construct() BusinessAwayMessageScheduleClass { return &b }

// Ensuring interfaces in compile-time for BusinessAwayMessageScheduleAlways.
var (
	_ bin.Encoder     = &BusinessAwayMessageScheduleAlways{}
	_ bin.Decoder     = &BusinessAwayMessageScheduleAlways{}
	_ bin.BareEncoder = &BusinessAwayMessageScheduleAlways{}
	_ bin.BareDecoder = &BusinessAwayMessageScheduleAlways{}

	_ BusinessAwayMessageScheduleClass = &BusinessAwayMessageScheduleAlways{}
)

func (b *BusinessAwayMessageScheduleAlways) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessAwayMessageScheduleAlways) String() string {
	if b == nil {
		return "BusinessAwayMessageScheduleAlways(nil)"
	}
	type Alias BusinessAwayMessageScheduleAlways
	return fmt.Sprintf("BusinessAwayMessageScheduleAlways%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessAwayMessageScheduleAlways) TypeID() uint32 {
	return BusinessAwayMessageScheduleAlwaysTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessAwayMessageScheduleAlways) TypeName() string {
	return "businessAwayMessageScheduleAlways"
}

// TypeInfo returns info about TL type.
func (b *BusinessAwayMessageScheduleAlways) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessAwayMessageScheduleAlways",
		ID:   BusinessAwayMessageScheduleAlwaysTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BusinessAwayMessageScheduleAlways) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleAlways#c9b9e2b9 as nil")
	}
	buf.PutID(BusinessAwayMessageScheduleAlwaysTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessAwayMessageScheduleAlways) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleAlways#c9b9e2b9 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessAwayMessageScheduleAlways) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleAlways#c9b9e2b9 to nil")
	}
	if err := buf.ConsumeID(BusinessAwayMessageScheduleAlwaysTypeID); err != nil {
		return fmt.Errorf("unable to decode businessAwayMessageScheduleAlways#c9b9e2b9: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessAwayMessageScheduleAlways) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleAlways#c9b9e2b9 to nil")
	}
	return nil
}

// BusinessAwayMessageScheduleOutsideWorkHours represents TL type `businessAwayMessageScheduleOutsideWorkHours#c3f2f501`.
//
// See https://core.telegram.org/constructor/businessAwayMessageScheduleOutsideWorkHours for reference.
type BusinessAwayMessageScheduleOutsideWorkHours struct {
}

// BusinessAwayMessageScheduleOutsideWorkHoursTypeID is TL type id of BusinessAwayMessageScheduleOutsideWorkHours.
const BusinessAwayMessageScheduleOutsideWorkHoursTypeID = 0xc3f2f501

// construct implements constructor of BusinessAwayMessageScheduleClass.
func (b BusinessAwayMessageScheduleOutsideWorkHours) construct() BusinessAwayMessageScheduleClass {
	return &b
}

// Ensuring interfaces in compile-time for BusinessAwayMessageScheduleOutsideWorkHours.
var (
	_ bin.Encoder     = &BusinessAwayMessageScheduleOutsideWorkHours{}
	_ bin.Decoder     = &BusinessAwayMessageScheduleOutsideWorkHours{}
	_ bin.BareEncoder = &BusinessAwayMessageScheduleOutsideWorkHours{}
	_ bin.BareDecoder = &BusinessAwayMessageScheduleOutsideWorkHours{}

	_ BusinessAwayMessageScheduleClass = &BusinessAwayMessageScheduleOutsideWorkHours{}
)

func (b *BusinessAwayMessageScheduleOutsideWorkHours) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessAwayMessageScheduleOutsideWorkHours) String() string {
	if b == nil {
		return "BusinessAwayMessageScheduleOutsideWorkHours(nil)"
	}
	type Alias BusinessAwayMessageScheduleOutsideWorkHours
	return fmt.Sprintf("BusinessAwayMessageScheduleOutsideWorkHours%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessAwayMessageScheduleOutsideWorkHours) TypeID() uint32 {
	return BusinessAwayMessageScheduleOutsideWorkHoursTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessAwayMessageScheduleOutsideWorkHours) TypeName() string {
	return "businessAwayMessageScheduleOutsideWorkHours"
}

// TypeInfo returns info about TL type.
func (b *BusinessAwayMessageScheduleOutsideWorkHours) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessAwayMessageScheduleOutsideWorkHours",
		ID:   BusinessAwayMessageScheduleOutsideWorkHoursTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BusinessAwayMessageScheduleOutsideWorkHours) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleOutsideWorkHours#c3f2f501 as nil")
	}
	buf.PutID(BusinessAwayMessageScheduleOutsideWorkHoursTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessAwayMessageScheduleOutsideWorkHours) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleOutsideWorkHours#c3f2f501 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessAwayMessageScheduleOutsideWorkHours) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleOutsideWorkHours#c3f2f501 to nil")
	}
	if err := buf.ConsumeID(BusinessAwayMessageScheduleOutsideWorkHoursTypeID); err != nil {
		return fmt.Errorf("unable to decode businessAwayMessageScheduleOutsideWorkHours#c3f2f501: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessAwayMessageScheduleOutsideWorkHours) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleOutsideWorkHours#c3f2f501 to nil")
	}
	return nil
}

// BusinessAwayMessageScheduleCustom represents TL type `businessAwayMessageScheduleCustom#cc4d9ecc`.
//
// See https://core.telegram.org/constructor/businessAwayMessageScheduleCustom for reference.
type BusinessAwayMessageScheduleCustom struct {
	// StartDate field of BusinessAwayMessageScheduleCustom.
	StartDate int
	// EndDate field of BusinessAwayMessageScheduleCustom.
	EndDate int
}

// BusinessAwayMessageScheduleCustomTypeID is TL type id of BusinessAwayMessageScheduleCustom.
const BusinessAwayMessageScheduleCustomTypeID = 0xcc4d9ecc

// construct implements constructor of BusinessAwayMessageScheduleClass.
func (b BusinessAwayMessageScheduleCustom) construct() BusinessAwayMessageScheduleClass { return &b }

// Ensuring interfaces in compile-time for BusinessAwayMessageScheduleCustom.
var (
	_ bin.Encoder     = &BusinessAwayMessageScheduleCustom{}
	_ bin.Decoder     = &BusinessAwayMessageScheduleCustom{}
	_ bin.BareEncoder = &BusinessAwayMessageScheduleCustom{}
	_ bin.BareDecoder = &BusinessAwayMessageScheduleCustom{}

	_ BusinessAwayMessageScheduleClass = &BusinessAwayMessageScheduleCustom{}
)

func (b *BusinessAwayMessageScheduleCustom) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.StartDate == 0) {
		return false
	}
	if !(b.EndDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessAwayMessageScheduleCustom) String() string {
	if b == nil {
		return "BusinessAwayMessageScheduleCustom(nil)"
	}
	type Alias BusinessAwayMessageScheduleCustom
	return fmt.Sprintf("BusinessAwayMessageScheduleCustom%+v", Alias(*b))
}

// FillFrom fills BusinessAwayMessageScheduleCustom from given interface.
func (b *BusinessAwayMessageScheduleCustom) FillFrom(from interface {
	GetStartDate() (value int)
	GetEndDate() (value int)
}) {
	b.StartDate = from.GetStartDate()
	b.EndDate = from.GetEndDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessAwayMessageScheduleCustom) TypeID() uint32 {
	return BusinessAwayMessageScheduleCustomTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessAwayMessageScheduleCustom) TypeName() string {
	return "businessAwayMessageScheduleCustom"
}

// TypeInfo returns info about TL type.
func (b *BusinessAwayMessageScheduleCustom) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessAwayMessageScheduleCustom",
		ID:   BusinessAwayMessageScheduleCustomTypeID,
	}
	if b == nil {
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
func (b *BusinessAwayMessageScheduleCustom) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleCustom#cc4d9ecc as nil")
	}
	buf.PutID(BusinessAwayMessageScheduleCustomTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessAwayMessageScheduleCustom) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleCustom#cc4d9ecc as nil")
	}
	buf.PutInt(b.StartDate)
	buf.PutInt(b.EndDate)
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessAwayMessageScheduleCustom) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleCustom#cc4d9ecc to nil")
	}
	if err := buf.ConsumeID(BusinessAwayMessageScheduleCustomTypeID); err != nil {
		return fmt.Errorf("unable to decode businessAwayMessageScheduleCustom#cc4d9ecc: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessAwayMessageScheduleCustom) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleCustom#cc4d9ecc to nil")
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode businessAwayMessageScheduleCustom#cc4d9ecc: field start_date: %w", err)
		}
		b.StartDate = value
	}
	{
		value, err := buf.Int()
		if err != nil {
			return fmt.Errorf("unable to decode businessAwayMessageScheduleCustom#cc4d9ecc: field end_date: %w", err)
		}
		b.EndDate = value
	}
	return nil
}

// GetStartDate returns value of StartDate field.
func (b *BusinessAwayMessageScheduleCustom) GetStartDate() (value int) {
	if b == nil {
		return
	}
	return b.StartDate
}

// GetEndDate returns value of EndDate field.
func (b *BusinessAwayMessageScheduleCustom) GetEndDate() (value int) {
	if b == nil {
		return
	}
	return b.EndDate
}

// BusinessAwayMessageScheduleClassName is schema name of BusinessAwayMessageScheduleClass.
const BusinessAwayMessageScheduleClassName = "BusinessAwayMessageSchedule"

// BusinessAwayMessageScheduleClass represents BusinessAwayMessageSchedule generic type.
//
// See https://core.telegram.org/type/BusinessAwayMessageSchedule for reference.
//
// Example:
//
//	g, err := tg.DecodeBusinessAwayMessageSchedule(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.BusinessAwayMessageScheduleAlways: // businessAwayMessageScheduleAlways#c9b9e2b9
//	case *tg.BusinessAwayMessageScheduleOutsideWorkHours: // businessAwayMessageScheduleOutsideWorkHours#c3f2f501
//	case *tg.BusinessAwayMessageScheduleCustom: // businessAwayMessageScheduleCustom#cc4d9ecc
//	default: panic(v)
//	}
type BusinessAwayMessageScheduleClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() BusinessAwayMessageScheduleClass

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
}

// DecodeBusinessAwayMessageSchedule implements binary de-serialization for BusinessAwayMessageScheduleClass.
func DecodeBusinessAwayMessageSchedule(buf *bin.Buffer) (BusinessAwayMessageScheduleClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BusinessAwayMessageScheduleAlwaysTypeID:
		// Decoding businessAwayMessageScheduleAlways#c9b9e2b9.
		v := BusinessAwayMessageScheduleAlways{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", err)
		}
		return &v, nil
	case BusinessAwayMessageScheduleOutsideWorkHoursTypeID:
		// Decoding businessAwayMessageScheduleOutsideWorkHours#c3f2f501.
		v := BusinessAwayMessageScheduleOutsideWorkHours{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", err)
		}
		return &v, nil
	case BusinessAwayMessageScheduleCustomTypeID:
		// Decoding businessAwayMessageScheduleCustom#cc4d9ecc.
		v := BusinessAwayMessageScheduleCustom{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", bin.NewUnexpectedID(id))
	}
}

// BusinessAwayMessageSchedule boxes the BusinessAwayMessageScheduleClass providing a helper.
type BusinessAwayMessageScheduleBox struct {
	BusinessAwayMessageSchedule BusinessAwayMessageScheduleClass
}

// Decode implements bin.Decoder for BusinessAwayMessageScheduleBox.
func (b *BusinessAwayMessageScheduleBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BusinessAwayMessageScheduleBox to nil")
	}
	v, err := DecodeBusinessAwayMessageSchedule(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BusinessAwayMessageSchedule = v
	return nil
}

// Encode implements bin.Encode for BusinessAwayMessageScheduleBox.
func (b *BusinessAwayMessageScheduleBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BusinessAwayMessageSchedule == nil {
		return fmt.Errorf("unable to encode BusinessAwayMessageScheduleClass as nil")
	}
	return b.BusinessAwayMessageSchedule.Encode(buf)
}
