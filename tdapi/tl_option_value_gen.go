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

// OptionValueBoolean represents TL type `optionValueBoolean#3c35f1e`.
type OptionValueBoolean struct {
	// The value of the option
	Value bool
}

// OptionValueBooleanTypeID is TL type id of OptionValueBoolean.
const OptionValueBooleanTypeID = 0x3c35f1e

// construct implements constructor of OptionValueClass.
func (o OptionValueBoolean) construct() OptionValueClass { return &o }

// Ensuring interfaces in compile-time for OptionValueBoolean.
var (
	_ bin.Encoder     = &OptionValueBoolean{}
	_ bin.Decoder     = &OptionValueBoolean{}
	_ bin.BareEncoder = &OptionValueBoolean{}
	_ bin.BareDecoder = &OptionValueBoolean{}

	_ OptionValueClass = &OptionValueBoolean{}
)

func (o *OptionValueBoolean) Zero() bool {
	if o == nil {
		return true
	}
	if !(o.Value == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (o *OptionValueBoolean) String() string {
	if o == nil {
		return "OptionValueBoolean(nil)"
	}
	type Alias OptionValueBoolean
	return fmt.Sprintf("OptionValueBoolean%+v", Alias(*o))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*OptionValueBoolean) TypeID() uint32 {
	return OptionValueBooleanTypeID
}

// TypeName returns name of type in TL schema.
func (*OptionValueBoolean) TypeName() string {
	return "optionValueBoolean"
}

// TypeInfo returns info about TL type.
func (o *OptionValueBoolean) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "optionValueBoolean",
		ID:   OptionValueBooleanTypeID,
	}
	if o == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (o *OptionValueBoolean) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueBoolean#3c35f1e as nil")
	}
	b.PutID(OptionValueBooleanTypeID)
	return o.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (o *OptionValueBoolean) EncodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueBoolean#3c35f1e as nil")
	}
	b.PutBool(o.Value)
	return nil
}

// Decode implements bin.Decoder.
func (o *OptionValueBoolean) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueBoolean#3c35f1e to nil")
	}
	if err := b.ConsumeID(OptionValueBooleanTypeID); err != nil {
		return fmt.Errorf("unable to decode optionValueBoolean#3c35f1e: %w", err)
	}
	return o.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (o *OptionValueBoolean) DecodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueBoolean#3c35f1e to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode optionValueBoolean#3c35f1e: field value: %w", err)
		}
		o.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (o *OptionValueBoolean) EncodeTDLibJSON(b tdjson.Encoder) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueBoolean#3c35f1e as nil")
	}
	b.ObjStart()
	b.PutID("optionValueBoolean")
	b.Comma()
	b.FieldStart("value")
	b.PutBool(o.Value)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (o *OptionValueBoolean) DecodeTDLibJSON(b tdjson.Decoder) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueBoolean#3c35f1e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("optionValueBoolean"); err != nil {
				return fmt.Errorf("unable to decode optionValueBoolean#3c35f1e: %w", err)
			}
		case "value":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode optionValueBoolean#3c35f1e: field value: %w", err)
			}
			o.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValue returns value of Value field.
func (o *OptionValueBoolean) GetValue() (value bool) {
	if o == nil {
		return
	}
	return o.Value
}

// OptionValueEmpty represents TL type `optionValueEmpty#36c62493`.
type OptionValueEmpty struct {
}

// OptionValueEmptyTypeID is TL type id of OptionValueEmpty.
const OptionValueEmptyTypeID = 0x36c62493

// construct implements constructor of OptionValueClass.
func (o OptionValueEmpty) construct() OptionValueClass { return &o }

// Ensuring interfaces in compile-time for OptionValueEmpty.
var (
	_ bin.Encoder     = &OptionValueEmpty{}
	_ bin.Decoder     = &OptionValueEmpty{}
	_ bin.BareEncoder = &OptionValueEmpty{}
	_ bin.BareDecoder = &OptionValueEmpty{}

	_ OptionValueClass = &OptionValueEmpty{}
)

func (o *OptionValueEmpty) Zero() bool {
	if o == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (o *OptionValueEmpty) String() string {
	if o == nil {
		return "OptionValueEmpty(nil)"
	}
	type Alias OptionValueEmpty
	return fmt.Sprintf("OptionValueEmpty%+v", Alias(*o))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*OptionValueEmpty) TypeID() uint32 {
	return OptionValueEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*OptionValueEmpty) TypeName() string {
	return "optionValueEmpty"
}

// TypeInfo returns info about TL type.
func (o *OptionValueEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "optionValueEmpty",
		ID:   OptionValueEmptyTypeID,
	}
	if o == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (o *OptionValueEmpty) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueEmpty#36c62493 as nil")
	}
	b.PutID(OptionValueEmptyTypeID)
	return o.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (o *OptionValueEmpty) EncodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueEmpty#36c62493 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (o *OptionValueEmpty) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueEmpty#36c62493 to nil")
	}
	if err := b.ConsumeID(OptionValueEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode optionValueEmpty#36c62493: %w", err)
	}
	return o.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (o *OptionValueEmpty) DecodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueEmpty#36c62493 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (o *OptionValueEmpty) EncodeTDLibJSON(b tdjson.Encoder) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueEmpty#36c62493 as nil")
	}
	b.ObjStart()
	b.PutID("optionValueEmpty")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (o *OptionValueEmpty) DecodeTDLibJSON(b tdjson.Decoder) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueEmpty#36c62493 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("optionValueEmpty"); err != nil {
				return fmt.Errorf("unable to decode optionValueEmpty#36c62493: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// OptionValueInteger represents TL type `optionValueInteger#f4dcc2e4`.
type OptionValueInteger struct {
	// The value of the option
	Value int64
}

// OptionValueIntegerTypeID is TL type id of OptionValueInteger.
const OptionValueIntegerTypeID = 0xf4dcc2e4

// construct implements constructor of OptionValueClass.
func (o OptionValueInteger) construct() OptionValueClass { return &o }

// Ensuring interfaces in compile-time for OptionValueInteger.
var (
	_ bin.Encoder     = &OptionValueInteger{}
	_ bin.Decoder     = &OptionValueInteger{}
	_ bin.BareEncoder = &OptionValueInteger{}
	_ bin.BareDecoder = &OptionValueInteger{}

	_ OptionValueClass = &OptionValueInteger{}
)

func (o *OptionValueInteger) Zero() bool {
	if o == nil {
		return true
	}
	if !(o.Value == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (o *OptionValueInteger) String() string {
	if o == nil {
		return "OptionValueInteger(nil)"
	}
	type Alias OptionValueInteger
	return fmt.Sprintf("OptionValueInteger%+v", Alias(*o))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*OptionValueInteger) TypeID() uint32 {
	return OptionValueIntegerTypeID
}

// TypeName returns name of type in TL schema.
func (*OptionValueInteger) TypeName() string {
	return "optionValueInteger"
}

// TypeInfo returns info about TL type.
func (o *OptionValueInteger) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "optionValueInteger",
		ID:   OptionValueIntegerTypeID,
	}
	if o == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (o *OptionValueInteger) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueInteger#f4dcc2e4 as nil")
	}
	b.PutID(OptionValueIntegerTypeID)
	return o.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (o *OptionValueInteger) EncodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueInteger#f4dcc2e4 as nil")
	}
	b.PutLong(o.Value)
	return nil
}

// Decode implements bin.Decoder.
func (o *OptionValueInteger) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueInteger#f4dcc2e4 to nil")
	}
	if err := b.ConsumeID(OptionValueIntegerTypeID); err != nil {
		return fmt.Errorf("unable to decode optionValueInteger#f4dcc2e4: %w", err)
	}
	return o.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (o *OptionValueInteger) DecodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueInteger#f4dcc2e4 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode optionValueInteger#f4dcc2e4: field value: %w", err)
		}
		o.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (o *OptionValueInteger) EncodeTDLibJSON(b tdjson.Encoder) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueInteger#f4dcc2e4 as nil")
	}
	b.ObjStart()
	b.PutID("optionValueInteger")
	b.Comma()
	b.FieldStart("value")
	b.PutLong(o.Value)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (o *OptionValueInteger) DecodeTDLibJSON(b tdjson.Decoder) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueInteger#f4dcc2e4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("optionValueInteger"); err != nil {
				return fmt.Errorf("unable to decode optionValueInteger#f4dcc2e4: %w", err)
			}
		case "value":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode optionValueInteger#f4dcc2e4: field value: %w", err)
			}
			o.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValue returns value of Value field.
func (o *OptionValueInteger) GetValue() (value int64) {
	if o == nil {
		return
	}
	return o.Value
}

// OptionValueString represents TL type `optionValueString#2d136e94`.
type OptionValueString struct {
	// The value of the option
	Value string
}

// OptionValueStringTypeID is TL type id of OptionValueString.
const OptionValueStringTypeID = 0x2d136e94

// construct implements constructor of OptionValueClass.
func (o OptionValueString) construct() OptionValueClass { return &o }

// Ensuring interfaces in compile-time for OptionValueString.
var (
	_ bin.Encoder     = &OptionValueString{}
	_ bin.Decoder     = &OptionValueString{}
	_ bin.BareEncoder = &OptionValueString{}
	_ bin.BareDecoder = &OptionValueString{}

	_ OptionValueClass = &OptionValueString{}
)

func (o *OptionValueString) Zero() bool {
	if o == nil {
		return true
	}
	if !(o.Value == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (o *OptionValueString) String() string {
	if o == nil {
		return "OptionValueString(nil)"
	}
	type Alias OptionValueString
	return fmt.Sprintf("OptionValueString%+v", Alias(*o))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*OptionValueString) TypeID() uint32 {
	return OptionValueStringTypeID
}

// TypeName returns name of type in TL schema.
func (*OptionValueString) TypeName() string {
	return "optionValueString"
}

// TypeInfo returns info about TL type.
func (o *OptionValueString) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "optionValueString",
		ID:   OptionValueStringTypeID,
	}
	if o == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Value",
			SchemaName: "value",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (o *OptionValueString) Encode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueString#2d136e94 as nil")
	}
	b.PutID(OptionValueStringTypeID)
	return o.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (o *OptionValueString) EncodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueString#2d136e94 as nil")
	}
	b.PutString(o.Value)
	return nil
}

// Decode implements bin.Decoder.
func (o *OptionValueString) Decode(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueString#2d136e94 to nil")
	}
	if err := b.ConsumeID(OptionValueStringTypeID); err != nil {
		return fmt.Errorf("unable to decode optionValueString#2d136e94: %w", err)
	}
	return o.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (o *OptionValueString) DecodeBare(b *bin.Buffer) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueString#2d136e94 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode optionValueString#2d136e94: field value: %w", err)
		}
		o.Value = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (o *OptionValueString) EncodeTDLibJSON(b tdjson.Encoder) error {
	if o == nil {
		return fmt.Errorf("can't encode optionValueString#2d136e94 as nil")
	}
	b.ObjStart()
	b.PutID("optionValueString")
	b.Comma()
	b.FieldStart("value")
	b.PutString(o.Value)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (o *OptionValueString) DecodeTDLibJSON(b tdjson.Decoder) error {
	if o == nil {
		return fmt.Errorf("can't decode optionValueString#2d136e94 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("optionValueString"); err != nil {
				return fmt.Errorf("unable to decode optionValueString#2d136e94: %w", err)
			}
		case "value":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode optionValueString#2d136e94: field value: %w", err)
			}
			o.Value = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetValue returns value of Value field.
func (o *OptionValueString) GetValue() (value string) {
	if o == nil {
		return
	}
	return o.Value
}

// OptionValueClassName is schema name of OptionValueClass.
const OptionValueClassName = "OptionValue"

// OptionValueClass represents OptionValue generic type.
//
// Example:
//
//	g, err := tdapi.DecodeOptionValue(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.OptionValueBoolean: // optionValueBoolean#3c35f1e
//	case *tdapi.OptionValueEmpty: // optionValueEmpty#36c62493
//	case *tdapi.OptionValueInteger: // optionValueInteger#f4dcc2e4
//	case *tdapi.OptionValueString: // optionValueString#2d136e94
//	default: panic(v)
//	}
type OptionValueClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() OptionValueClass

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

// DecodeOptionValue implements binary de-serialization for OptionValueClass.
func DecodeOptionValue(buf *bin.Buffer) (OptionValueClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case OptionValueBooleanTypeID:
		// Decoding optionValueBoolean#3c35f1e.
		v := OptionValueBoolean{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode OptionValueClass: %w", err)
		}
		return &v, nil
	case OptionValueEmptyTypeID:
		// Decoding optionValueEmpty#36c62493.
		v := OptionValueEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode OptionValueClass: %w", err)
		}
		return &v, nil
	case OptionValueIntegerTypeID:
		// Decoding optionValueInteger#f4dcc2e4.
		v := OptionValueInteger{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode OptionValueClass: %w", err)
		}
		return &v, nil
	case OptionValueStringTypeID:
		// Decoding optionValueString#2d136e94.
		v := OptionValueString{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode OptionValueClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode OptionValueClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONOptionValue implements binary de-serialization for OptionValueClass.
func DecodeTDLibJSONOptionValue(buf tdjson.Decoder) (OptionValueClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "optionValueBoolean":
		// Decoding optionValueBoolean#3c35f1e.
		v := OptionValueBoolean{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode OptionValueClass: %w", err)
		}
		return &v, nil
	case "optionValueEmpty":
		// Decoding optionValueEmpty#36c62493.
		v := OptionValueEmpty{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode OptionValueClass: %w", err)
		}
		return &v, nil
	case "optionValueInteger":
		// Decoding optionValueInteger#f4dcc2e4.
		v := OptionValueInteger{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode OptionValueClass: %w", err)
		}
		return &v, nil
	case "optionValueString":
		// Decoding optionValueString#2d136e94.
		v := OptionValueString{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode OptionValueClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode OptionValueClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// OptionValue boxes the OptionValueClass providing a helper.
type OptionValueBox struct {
	OptionValue OptionValueClass
}

// Decode implements bin.Decoder for OptionValueBox.
func (b *OptionValueBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode OptionValueBox to nil")
	}
	v, err := DecodeOptionValue(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.OptionValue = v
	return nil
}

// Encode implements bin.Encode for OptionValueBox.
func (b *OptionValueBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.OptionValue == nil {
		return fmt.Errorf("unable to encode OptionValueClass as nil")
	}
	return b.OptionValue.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for OptionValueBox.
func (b *OptionValueBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode OptionValueBox to nil")
	}
	v, err := DecodeTDLibJSONOptionValue(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.OptionValue = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for OptionValueBox.
func (b *OptionValueBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.OptionValue == nil {
		return fmt.Errorf("unable to encode OptionValueClass as nil")
	}
	return b.OptionValue.EncodeTDLibJSON(buf)
}
