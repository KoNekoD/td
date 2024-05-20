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

// BusinessAwayMessageScheduleAlways represents TL type `businessAwayMessageScheduleAlways#c9b9e2b9`.
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

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BusinessAwayMessageScheduleAlways) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleAlways#c9b9e2b9 as nil")
	}
	buf.ObjStart()
	buf.PutID("businessAwayMessageScheduleAlways")
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BusinessAwayMessageScheduleAlways) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleAlways#c9b9e2b9 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("businessAwayMessageScheduleAlways"); err != nil {
				return fmt.Errorf("unable to decode businessAwayMessageScheduleAlways#c9b9e2b9: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// BusinessAwayMessageScheduleOutsideOfOpeningHours represents TL type `businessAwayMessageScheduleOutsideOfOpeningHours#c643df16`.
type BusinessAwayMessageScheduleOutsideOfOpeningHours struct {
}

// BusinessAwayMessageScheduleOutsideOfOpeningHoursTypeID is TL type id of BusinessAwayMessageScheduleOutsideOfOpeningHours.
const BusinessAwayMessageScheduleOutsideOfOpeningHoursTypeID = 0xc643df16

// construct implements constructor of BusinessAwayMessageScheduleClass.
func (b BusinessAwayMessageScheduleOutsideOfOpeningHours) construct() BusinessAwayMessageScheduleClass {
	return &b
}

// Ensuring interfaces in compile-time for BusinessAwayMessageScheduleOutsideOfOpeningHours.
var (
	_ bin.Encoder     = &BusinessAwayMessageScheduleOutsideOfOpeningHours{}
	_ bin.Decoder     = &BusinessAwayMessageScheduleOutsideOfOpeningHours{}
	_ bin.BareEncoder = &BusinessAwayMessageScheduleOutsideOfOpeningHours{}
	_ bin.BareDecoder = &BusinessAwayMessageScheduleOutsideOfOpeningHours{}

	_ BusinessAwayMessageScheduleClass = &BusinessAwayMessageScheduleOutsideOfOpeningHours{}
)

func (b *BusinessAwayMessageScheduleOutsideOfOpeningHours) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BusinessAwayMessageScheduleOutsideOfOpeningHours) String() string {
	if b == nil {
		return "BusinessAwayMessageScheduleOutsideOfOpeningHours(nil)"
	}
	type Alias BusinessAwayMessageScheduleOutsideOfOpeningHours
	return fmt.Sprintf("BusinessAwayMessageScheduleOutsideOfOpeningHours%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BusinessAwayMessageScheduleOutsideOfOpeningHours) TypeID() uint32 {
	return BusinessAwayMessageScheduleOutsideOfOpeningHoursTypeID
}

// TypeName returns name of type in TL schema.
func (*BusinessAwayMessageScheduleOutsideOfOpeningHours) TypeName() string {
	return "businessAwayMessageScheduleOutsideOfOpeningHours"
}

// TypeInfo returns info about TL type.
func (b *BusinessAwayMessageScheduleOutsideOfOpeningHours) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "businessAwayMessageScheduleOutsideOfOpeningHours",
		ID:   BusinessAwayMessageScheduleOutsideOfOpeningHoursTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BusinessAwayMessageScheduleOutsideOfOpeningHours) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleOutsideOfOpeningHours#c643df16 as nil")
	}
	buf.PutID(BusinessAwayMessageScheduleOutsideOfOpeningHoursTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessAwayMessageScheduleOutsideOfOpeningHours) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleOutsideOfOpeningHours#c643df16 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessAwayMessageScheduleOutsideOfOpeningHours) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleOutsideOfOpeningHours#c643df16 to nil")
	}
	if err := buf.ConsumeID(BusinessAwayMessageScheduleOutsideOfOpeningHoursTypeID); err != nil {
		return fmt.Errorf("unable to decode businessAwayMessageScheduleOutsideOfOpeningHours#c643df16: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessAwayMessageScheduleOutsideOfOpeningHours) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleOutsideOfOpeningHours#c643df16 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BusinessAwayMessageScheduleOutsideOfOpeningHours) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleOutsideOfOpeningHours#c643df16 as nil")
	}
	buf.ObjStart()
	buf.PutID("businessAwayMessageScheduleOutsideOfOpeningHours")
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BusinessAwayMessageScheduleOutsideOfOpeningHours) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleOutsideOfOpeningHours#c643df16 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("businessAwayMessageScheduleOutsideOfOpeningHours"); err != nil {
				return fmt.Errorf("unable to decode businessAwayMessageScheduleOutsideOfOpeningHours#c643df16: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// BusinessAwayMessageScheduleCustom represents TL type `businessAwayMessageScheduleCustom#8ac04dd2`.
type BusinessAwayMessageScheduleCustom struct {
	// Point in time (Unix timestamp) when the away messages will start to be sent
	StartDate int32
	// Point in time (Unix timestamp) when the away messages will stop to be sent
	EndDate int32
}

// BusinessAwayMessageScheduleCustomTypeID is TL type id of BusinessAwayMessageScheduleCustom.
const BusinessAwayMessageScheduleCustomTypeID = 0x8ac04dd2

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
		return fmt.Errorf("can't encode businessAwayMessageScheduleCustom#8ac04dd2 as nil")
	}
	buf.PutID(BusinessAwayMessageScheduleCustomTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BusinessAwayMessageScheduleCustom) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleCustom#8ac04dd2 as nil")
	}
	buf.PutInt32(b.StartDate)
	buf.PutInt32(b.EndDate)
	return nil
}

// Decode implements bin.Decoder.
func (b *BusinessAwayMessageScheduleCustom) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleCustom#8ac04dd2 to nil")
	}
	if err := buf.ConsumeID(BusinessAwayMessageScheduleCustomTypeID); err != nil {
		return fmt.Errorf("unable to decode businessAwayMessageScheduleCustom#8ac04dd2: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BusinessAwayMessageScheduleCustom) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleCustom#8ac04dd2 to nil")
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode businessAwayMessageScheduleCustom#8ac04dd2: field start_date: %w", err)
		}
		b.StartDate = value
	}
	{
		value, err := buf.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode businessAwayMessageScheduleCustom#8ac04dd2: field end_date: %w", err)
		}
		b.EndDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BusinessAwayMessageScheduleCustom) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode businessAwayMessageScheduleCustom#8ac04dd2 as nil")
	}
	buf.ObjStart()
	buf.PutID("businessAwayMessageScheduleCustom")
	buf.Comma()
	buf.FieldStart("start_date")
	buf.PutInt32(b.StartDate)
	buf.Comma()
	buf.FieldStart("end_date")
	buf.PutInt32(b.EndDate)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BusinessAwayMessageScheduleCustom) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode businessAwayMessageScheduleCustom#8ac04dd2 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("businessAwayMessageScheduleCustom"); err != nil {
				return fmt.Errorf("unable to decode businessAwayMessageScheduleCustom#8ac04dd2: %w", err)
			}
		case "start_date":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode businessAwayMessageScheduleCustom#8ac04dd2: field start_date: %w", err)
			}
			b.StartDate = value
		case "end_date":
			value, err := buf.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode businessAwayMessageScheduleCustom#8ac04dd2: field end_date: %w", err)
			}
			b.EndDate = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetStartDate returns value of StartDate field.
func (b *BusinessAwayMessageScheduleCustom) GetStartDate() (value int32) {
	if b == nil {
		return
	}
	return b.StartDate
}

// GetEndDate returns value of EndDate field.
func (b *BusinessAwayMessageScheduleCustom) GetEndDate() (value int32) {
	if b == nil {
		return
	}
	return b.EndDate
}

// BusinessAwayMessageScheduleClassName is schema name of BusinessAwayMessageScheduleClass.
const BusinessAwayMessageScheduleClassName = "BusinessAwayMessageSchedule"

// BusinessAwayMessageScheduleClass represents BusinessAwayMessageSchedule generic type.
//
// Example:
//
//	g, err := tdapi.DecodeBusinessAwayMessageSchedule(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.BusinessAwayMessageScheduleAlways: // businessAwayMessageScheduleAlways#c9b9e2b9
//	case *tdapi.BusinessAwayMessageScheduleOutsideOfOpeningHours: // businessAwayMessageScheduleOutsideOfOpeningHours#c643df16
//	case *tdapi.BusinessAwayMessageScheduleCustom: // businessAwayMessageScheduleCustom#8ac04dd2
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

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
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
	case BusinessAwayMessageScheduleOutsideOfOpeningHoursTypeID:
		// Decoding businessAwayMessageScheduleOutsideOfOpeningHours#c643df16.
		v := BusinessAwayMessageScheduleOutsideOfOpeningHours{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", err)
		}
		return &v, nil
	case BusinessAwayMessageScheduleCustomTypeID:
		// Decoding businessAwayMessageScheduleCustom#8ac04dd2.
		v := BusinessAwayMessageScheduleCustom{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONBusinessAwayMessageSchedule implements binary de-serialization for BusinessAwayMessageScheduleClass.
func DecodeTDLibJSONBusinessAwayMessageSchedule(buf tdjson.Decoder) (BusinessAwayMessageScheduleClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "businessAwayMessageScheduleAlways":
		// Decoding businessAwayMessageScheduleAlways#c9b9e2b9.
		v := BusinessAwayMessageScheduleAlways{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", err)
		}
		return &v, nil
	case "businessAwayMessageScheduleOutsideOfOpeningHours":
		// Decoding businessAwayMessageScheduleOutsideOfOpeningHours#c643df16.
		v := BusinessAwayMessageScheduleOutsideOfOpeningHours{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", err)
		}
		return &v, nil
	case "businessAwayMessageScheduleCustom":
		// Decoding businessAwayMessageScheduleCustom#8ac04dd2.
		v := BusinessAwayMessageScheduleCustom{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BusinessAwayMessageScheduleClass: %w", tdjson.NewUnexpectedID(id))
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

// DecodeTDLibJSON implements bin.Decoder for BusinessAwayMessageScheduleBox.
func (b *BusinessAwayMessageScheduleBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode BusinessAwayMessageScheduleBox to nil")
	}
	v, err := DecodeTDLibJSONBusinessAwayMessageSchedule(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BusinessAwayMessageSchedule = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for BusinessAwayMessageScheduleBox.
func (b *BusinessAwayMessageScheduleBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.BusinessAwayMessageSchedule == nil {
		return fmt.Errorf("unable to encode BusinessAwayMessageScheduleClass as nil")
	}
	return b.BusinessAwayMessageSchedule.EncodeTDLibJSON(buf)
}
