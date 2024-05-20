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

// VectorPathCommandLine represents TL type `vectorPathCommandLine#db663c8a`.
type VectorPathCommandLine struct {
	// The end point of the straight line
	EndPoint Point
}

// VectorPathCommandLineTypeID is TL type id of VectorPathCommandLine.
const VectorPathCommandLineTypeID = 0xdb663c8a

// construct implements constructor of VectorPathCommandClass.
func (v VectorPathCommandLine) construct() VectorPathCommandClass { return &v }

// Ensuring interfaces in compile-time for VectorPathCommandLine.
var (
	_ bin.Encoder     = &VectorPathCommandLine{}
	_ bin.Decoder     = &VectorPathCommandLine{}
	_ bin.BareEncoder = &VectorPathCommandLine{}
	_ bin.BareDecoder = &VectorPathCommandLine{}

	_ VectorPathCommandClass = &VectorPathCommandLine{}
)

func (v *VectorPathCommandLine) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.EndPoint.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *VectorPathCommandLine) String() string {
	if v == nil {
		return "VectorPathCommandLine(nil)"
	}
	type Alias VectorPathCommandLine
	return fmt.Sprintf("VectorPathCommandLine%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*VectorPathCommandLine) TypeID() uint32 {
	return VectorPathCommandLineTypeID
}

// TypeName returns name of type in TL schema.
func (*VectorPathCommandLine) TypeName() string {
	return "vectorPathCommandLine"
}

// TypeInfo returns info about TL type.
func (v *VectorPathCommandLine) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "vectorPathCommandLine",
		ID:   VectorPathCommandLineTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EndPoint",
			SchemaName: "end_point",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *VectorPathCommandLine) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode vectorPathCommandLine#db663c8a as nil")
	}
	b.PutID(VectorPathCommandLineTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *VectorPathCommandLine) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode vectorPathCommandLine#db663c8a as nil")
	}
	if err := v.EndPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode vectorPathCommandLine#db663c8a: field end_point: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *VectorPathCommandLine) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode vectorPathCommandLine#db663c8a to nil")
	}
	if err := b.ConsumeID(VectorPathCommandLineTypeID); err != nil {
		return fmt.Errorf("unable to decode vectorPathCommandLine#db663c8a: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *VectorPathCommandLine) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode vectorPathCommandLine#db663c8a to nil")
	}
	{
		if err := v.EndPoint.Decode(b); err != nil {
			return fmt.Errorf("unable to decode vectorPathCommandLine#db663c8a: field end_point: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *VectorPathCommandLine) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode vectorPathCommandLine#db663c8a as nil")
	}
	b.ObjStart()
	b.PutID("vectorPathCommandLine")
	b.Comma()
	b.FieldStart("end_point")
	if err := v.EndPoint.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode vectorPathCommandLine#db663c8a: field end_point: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *VectorPathCommandLine) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode vectorPathCommandLine#db663c8a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("vectorPathCommandLine"); err != nil {
				return fmt.Errorf("unable to decode vectorPathCommandLine#db663c8a: %w", err)
			}
		case "end_point":
			if err := v.EndPoint.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode vectorPathCommandLine#db663c8a: field end_point: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEndPoint returns value of EndPoint field.
func (v *VectorPathCommandLine) GetEndPoint() (value Point) {
	if v == nil {
		return
	}
	return v.EndPoint
}

// VectorPathCommandCubicBezierCurve represents TL type `vectorPathCommandCubicBezierCurve#494c3e3a`.
type VectorPathCommandCubicBezierCurve struct {
	// The start control point of the curve
	StartControlPoint Point
	// The end control point of the curve
	EndControlPoint Point
	// The end point of the curve
	EndPoint Point
}

// VectorPathCommandCubicBezierCurveTypeID is TL type id of VectorPathCommandCubicBezierCurve.
const VectorPathCommandCubicBezierCurveTypeID = 0x494c3e3a

// construct implements constructor of VectorPathCommandClass.
func (v VectorPathCommandCubicBezierCurve) construct() VectorPathCommandClass { return &v }

// Ensuring interfaces in compile-time for VectorPathCommandCubicBezierCurve.
var (
	_ bin.Encoder     = &VectorPathCommandCubicBezierCurve{}
	_ bin.Decoder     = &VectorPathCommandCubicBezierCurve{}
	_ bin.BareEncoder = &VectorPathCommandCubicBezierCurve{}
	_ bin.BareDecoder = &VectorPathCommandCubicBezierCurve{}

	_ VectorPathCommandClass = &VectorPathCommandCubicBezierCurve{}
)

func (v *VectorPathCommandCubicBezierCurve) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.StartControlPoint.Zero()) {
		return false
	}
	if !(v.EndControlPoint.Zero()) {
		return false
	}
	if !(v.EndPoint.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *VectorPathCommandCubicBezierCurve) String() string {
	if v == nil {
		return "VectorPathCommandCubicBezierCurve(nil)"
	}
	type Alias VectorPathCommandCubicBezierCurve
	return fmt.Sprintf("VectorPathCommandCubicBezierCurve%+v", Alias(*v))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*VectorPathCommandCubicBezierCurve) TypeID() uint32 {
	return VectorPathCommandCubicBezierCurveTypeID
}

// TypeName returns name of type in TL schema.
func (*VectorPathCommandCubicBezierCurve) TypeName() string {
	return "vectorPathCommandCubicBezierCurve"
}

// TypeInfo returns info about TL type.
func (v *VectorPathCommandCubicBezierCurve) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "vectorPathCommandCubicBezierCurve",
		ID:   VectorPathCommandCubicBezierCurveTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "StartControlPoint",
			SchemaName: "start_control_point",
		},
		{
			Name:       "EndControlPoint",
			SchemaName: "end_control_point",
		},
		{
			Name:       "EndPoint",
			SchemaName: "end_point",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *VectorPathCommandCubicBezierCurve) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode vectorPathCommandCubicBezierCurve#494c3e3a as nil")
	}
	b.PutID(VectorPathCommandCubicBezierCurveTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *VectorPathCommandCubicBezierCurve) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode vectorPathCommandCubicBezierCurve#494c3e3a as nil")
	}
	if err := v.StartControlPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode vectorPathCommandCubicBezierCurve#494c3e3a: field start_control_point: %w", err)
	}
	if err := v.EndControlPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode vectorPathCommandCubicBezierCurve#494c3e3a: field end_control_point: %w", err)
	}
	if err := v.EndPoint.Encode(b); err != nil {
		return fmt.Errorf("unable to encode vectorPathCommandCubicBezierCurve#494c3e3a: field end_point: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *VectorPathCommandCubicBezierCurve) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode vectorPathCommandCubicBezierCurve#494c3e3a to nil")
	}
	if err := b.ConsumeID(VectorPathCommandCubicBezierCurveTypeID); err != nil {
		return fmt.Errorf("unable to decode vectorPathCommandCubicBezierCurve#494c3e3a: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *VectorPathCommandCubicBezierCurve) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode vectorPathCommandCubicBezierCurve#494c3e3a to nil")
	}
	{
		if err := v.StartControlPoint.Decode(b); err != nil {
			return fmt.Errorf("unable to decode vectorPathCommandCubicBezierCurve#494c3e3a: field start_control_point: %w", err)
		}
	}
	{
		if err := v.EndControlPoint.Decode(b); err != nil {
			return fmt.Errorf("unable to decode vectorPathCommandCubicBezierCurve#494c3e3a: field end_control_point: %w", err)
		}
	}
	{
		if err := v.EndPoint.Decode(b); err != nil {
			return fmt.Errorf("unable to decode vectorPathCommandCubicBezierCurve#494c3e3a: field end_point: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (v *VectorPathCommandCubicBezierCurve) EncodeTDLibJSON(b tdjson.Encoder) error {
	if v == nil {
		return fmt.Errorf("can't encode vectorPathCommandCubicBezierCurve#494c3e3a as nil")
	}
	b.ObjStart()
	b.PutID("vectorPathCommandCubicBezierCurve")
	b.Comma()
	b.FieldStart("start_control_point")
	if err := v.StartControlPoint.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode vectorPathCommandCubicBezierCurve#494c3e3a: field start_control_point: %w", err)
	}
	b.Comma()
	b.FieldStart("end_control_point")
	if err := v.EndControlPoint.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode vectorPathCommandCubicBezierCurve#494c3e3a: field end_control_point: %w", err)
	}
	b.Comma()
	b.FieldStart("end_point")
	if err := v.EndPoint.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode vectorPathCommandCubicBezierCurve#494c3e3a: field end_point: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (v *VectorPathCommandCubicBezierCurve) DecodeTDLibJSON(b tdjson.Decoder) error {
	if v == nil {
		return fmt.Errorf("can't decode vectorPathCommandCubicBezierCurve#494c3e3a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("vectorPathCommandCubicBezierCurve"); err != nil {
				return fmt.Errorf("unable to decode vectorPathCommandCubicBezierCurve#494c3e3a: %w", err)
			}
		case "start_control_point":
			if err := v.StartControlPoint.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode vectorPathCommandCubicBezierCurve#494c3e3a: field start_control_point: %w", err)
			}
		case "end_control_point":
			if err := v.EndControlPoint.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode vectorPathCommandCubicBezierCurve#494c3e3a: field end_control_point: %w", err)
			}
		case "end_point":
			if err := v.EndPoint.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode vectorPathCommandCubicBezierCurve#494c3e3a: field end_point: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetStartControlPoint returns value of StartControlPoint field.
func (v *VectorPathCommandCubicBezierCurve) GetStartControlPoint() (value Point) {
	if v == nil {
		return
	}
	return v.StartControlPoint
}

// GetEndControlPoint returns value of EndControlPoint field.
func (v *VectorPathCommandCubicBezierCurve) GetEndControlPoint() (value Point) {
	if v == nil {
		return
	}
	return v.EndControlPoint
}

// GetEndPoint returns value of EndPoint field.
func (v *VectorPathCommandCubicBezierCurve) GetEndPoint() (value Point) {
	if v == nil {
		return
	}
	return v.EndPoint
}

// VectorPathCommandClassName is schema name of VectorPathCommandClass.
const VectorPathCommandClassName = "VectorPathCommand"

// VectorPathCommandClass represents VectorPathCommand generic type.
//
// Example:
//
//	g, err := tdapi.DecodeVectorPathCommand(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.VectorPathCommandLine: // vectorPathCommandLine#db663c8a
//	case *tdapi.VectorPathCommandCubicBezierCurve: // vectorPathCommandCubicBezierCurve#494c3e3a
//	default: panic(v)
//	}
type VectorPathCommandClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() VectorPathCommandClass

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

	// The end point of the straight line
	GetEndPoint() (value Point)
}

// DecodeVectorPathCommand implements binary de-serialization for VectorPathCommandClass.
func DecodeVectorPathCommand(buf *bin.Buffer) (VectorPathCommandClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case VectorPathCommandLineTypeID:
		// Decoding vectorPathCommandLine#db663c8a.
		v := VectorPathCommandLine{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode VectorPathCommandClass: %w", err)
		}
		return &v, nil
	case VectorPathCommandCubicBezierCurveTypeID:
		// Decoding vectorPathCommandCubicBezierCurve#494c3e3a.
		v := VectorPathCommandCubicBezierCurve{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode VectorPathCommandClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode VectorPathCommandClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONVectorPathCommand implements binary de-serialization for VectorPathCommandClass.
func DecodeTDLibJSONVectorPathCommand(buf tdjson.Decoder) (VectorPathCommandClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "vectorPathCommandLine":
		// Decoding vectorPathCommandLine#db663c8a.
		v := VectorPathCommandLine{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode VectorPathCommandClass: %w", err)
		}
		return &v, nil
	case "vectorPathCommandCubicBezierCurve":
		// Decoding vectorPathCommandCubicBezierCurve#494c3e3a.
		v := VectorPathCommandCubicBezierCurve{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode VectorPathCommandClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode VectorPathCommandClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// VectorPathCommand boxes the VectorPathCommandClass providing a helper.
type VectorPathCommandBox struct {
	VectorPathCommand VectorPathCommandClass
}

// Decode implements bin.Decoder for VectorPathCommandBox.
func (b *VectorPathCommandBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode VectorPathCommandBox to nil")
	}
	v, err := DecodeVectorPathCommand(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.VectorPathCommand = v
	return nil
}

// Encode implements bin.Encode for VectorPathCommandBox.
func (b *VectorPathCommandBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.VectorPathCommand == nil {
		return fmt.Errorf("unable to encode VectorPathCommandClass as nil")
	}
	return b.VectorPathCommand.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for VectorPathCommandBox.
func (b *VectorPathCommandBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode VectorPathCommandBox to nil")
	}
	v, err := DecodeTDLibJSONVectorPathCommand(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.VectorPathCommand = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for VectorPathCommandBox.
func (b *VectorPathCommandBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.VectorPathCommand == nil {
		return fmt.Errorf("unable to encode VectorPathCommandClass as nil")
	}
	return b.VectorPathCommand.EncodeTDLibJSON(buf)
}
