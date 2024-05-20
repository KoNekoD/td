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

// FirebaseAuthenticationSettingsAndroid represents TL type `firebaseAuthenticationSettingsAndroid#966ef61c`.
type FirebaseAuthenticationSettingsAndroid struct {
}

// FirebaseAuthenticationSettingsAndroidTypeID is TL type id of FirebaseAuthenticationSettingsAndroid.
const FirebaseAuthenticationSettingsAndroidTypeID = 0x966ef61c

// construct implements constructor of FirebaseAuthenticationSettingsClass.
func (f FirebaseAuthenticationSettingsAndroid) construct() FirebaseAuthenticationSettingsClass {
	return &f
}

// Ensuring interfaces in compile-time for FirebaseAuthenticationSettingsAndroid.
var (
	_ bin.Encoder     = &FirebaseAuthenticationSettingsAndroid{}
	_ bin.Decoder     = &FirebaseAuthenticationSettingsAndroid{}
	_ bin.BareEncoder = &FirebaseAuthenticationSettingsAndroid{}
	_ bin.BareDecoder = &FirebaseAuthenticationSettingsAndroid{}

	_ FirebaseAuthenticationSettingsClass = &FirebaseAuthenticationSettingsAndroid{}
)

func (f *FirebaseAuthenticationSettingsAndroid) Zero() bool {
	if f == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (f *FirebaseAuthenticationSettingsAndroid) String() string {
	if f == nil {
		return "FirebaseAuthenticationSettingsAndroid(nil)"
	}
	type Alias FirebaseAuthenticationSettingsAndroid
	return fmt.Sprintf("FirebaseAuthenticationSettingsAndroid%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FirebaseAuthenticationSettingsAndroid) TypeID() uint32 {
	return FirebaseAuthenticationSettingsAndroidTypeID
}

// TypeName returns name of type in TL schema.
func (*FirebaseAuthenticationSettingsAndroid) TypeName() string {
	return "firebaseAuthenticationSettingsAndroid"
}

// TypeInfo returns info about TL type.
func (f *FirebaseAuthenticationSettingsAndroid) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "firebaseAuthenticationSettingsAndroid",
		ID:   FirebaseAuthenticationSettingsAndroidTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (f *FirebaseAuthenticationSettingsAndroid) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode firebaseAuthenticationSettingsAndroid#966ef61c as nil")
	}
	b.PutID(FirebaseAuthenticationSettingsAndroidTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FirebaseAuthenticationSettingsAndroid) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode firebaseAuthenticationSettingsAndroid#966ef61c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *FirebaseAuthenticationSettingsAndroid) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode firebaseAuthenticationSettingsAndroid#966ef61c to nil")
	}
	if err := b.ConsumeID(FirebaseAuthenticationSettingsAndroidTypeID); err != nil {
		return fmt.Errorf("unable to decode firebaseAuthenticationSettingsAndroid#966ef61c: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FirebaseAuthenticationSettingsAndroid) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode firebaseAuthenticationSettingsAndroid#966ef61c to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FirebaseAuthenticationSettingsAndroid) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode firebaseAuthenticationSettingsAndroid#966ef61c as nil")
	}
	b.ObjStart()
	b.PutID("firebaseAuthenticationSettingsAndroid")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *FirebaseAuthenticationSettingsAndroid) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode firebaseAuthenticationSettingsAndroid#966ef61c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("firebaseAuthenticationSettingsAndroid"); err != nil {
				return fmt.Errorf("unable to decode firebaseAuthenticationSettingsAndroid#966ef61c: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// FirebaseAuthenticationSettingsIos represents TL type `firebaseAuthenticationSettingsIos#d49a4c4`.
type FirebaseAuthenticationSettingsIos struct {
	// Device token from Apple Push Notification service
	DeviceToken string
	// True, if App Sandbox is enabled
	IsAppSandbox bool
}

// FirebaseAuthenticationSettingsIosTypeID is TL type id of FirebaseAuthenticationSettingsIos.
const FirebaseAuthenticationSettingsIosTypeID = 0xd49a4c4

// construct implements constructor of FirebaseAuthenticationSettingsClass.
func (f FirebaseAuthenticationSettingsIos) construct() FirebaseAuthenticationSettingsClass { return &f }

// Ensuring interfaces in compile-time for FirebaseAuthenticationSettingsIos.
var (
	_ bin.Encoder     = &FirebaseAuthenticationSettingsIos{}
	_ bin.Decoder     = &FirebaseAuthenticationSettingsIos{}
	_ bin.BareEncoder = &FirebaseAuthenticationSettingsIos{}
	_ bin.BareDecoder = &FirebaseAuthenticationSettingsIos{}

	_ FirebaseAuthenticationSettingsClass = &FirebaseAuthenticationSettingsIos{}
)

func (f *FirebaseAuthenticationSettingsIos) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.DeviceToken == "") {
		return false
	}
	if !(f.IsAppSandbox == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FirebaseAuthenticationSettingsIos) String() string {
	if f == nil {
		return "FirebaseAuthenticationSettingsIos(nil)"
	}
	type Alias FirebaseAuthenticationSettingsIos
	return fmt.Sprintf("FirebaseAuthenticationSettingsIos%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FirebaseAuthenticationSettingsIos) TypeID() uint32 {
	return FirebaseAuthenticationSettingsIosTypeID
}

// TypeName returns name of type in TL schema.
func (*FirebaseAuthenticationSettingsIos) TypeName() string {
	return "firebaseAuthenticationSettingsIos"
}

// TypeInfo returns info about TL type.
func (f *FirebaseAuthenticationSettingsIos) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "firebaseAuthenticationSettingsIos",
		ID:   FirebaseAuthenticationSettingsIosTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DeviceToken",
			SchemaName: "device_token",
		},
		{
			Name:       "IsAppSandbox",
			SchemaName: "is_app_sandbox",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FirebaseAuthenticationSettingsIos) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode firebaseAuthenticationSettingsIos#d49a4c4 as nil")
	}
	b.PutID(FirebaseAuthenticationSettingsIosTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FirebaseAuthenticationSettingsIos) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode firebaseAuthenticationSettingsIos#d49a4c4 as nil")
	}
	b.PutString(f.DeviceToken)
	b.PutBool(f.IsAppSandbox)
	return nil
}

// Decode implements bin.Decoder.
func (f *FirebaseAuthenticationSettingsIos) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode firebaseAuthenticationSettingsIos#d49a4c4 to nil")
	}
	if err := b.ConsumeID(FirebaseAuthenticationSettingsIosTypeID); err != nil {
		return fmt.Errorf("unable to decode firebaseAuthenticationSettingsIos#d49a4c4: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FirebaseAuthenticationSettingsIos) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode firebaseAuthenticationSettingsIos#d49a4c4 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode firebaseAuthenticationSettingsIos#d49a4c4: field device_token: %w", err)
		}
		f.DeviceToken = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode firebaseAuthenticationSettingsIos#d49a4c4: field is_app_sandbox: %w", err)
		}
		f.IsAppSandbox = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FirebaseAuthenticationSettingsIos) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode firebaseAuthenticationSettingsIos#d49a4c4 as nil")
	}
	b.ObjStart()
	b.PutID("firebaseAuthenticationSettingsIos")
	b.Comma()
	b.FieldStart("device_token")
	b.PutString(f.DeviceToken)
	b.Comma()
	b.FieldStart("is_app_sandbox")
	b.PutBool(f.IsAppSandbox)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *FirebaseAuthenticationSettingsIos) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode firebaseAuthenticationSettingsIos#d49a4c4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("firebaseAuthenticationSettingsIos"); err != nil {
				return fmt.Errorf("unable to decode firebaseAuthenticationSettingsIos#d49a4c4: %w", err)
			}
		case "device_token":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode firebaseAuthenticationSettingsIos#d49a4c4: field device_token: %w", err)
			}
			f.DeviceToken = value
		case "is_app_sandbox":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode firebaseAuthenticationSettingsIos#d49a4c4: field is_app_sandbox: %w", err)
			}
			f.IsAppSandbox = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDeviceToken returns value of DeviceToken field.
func (f *FirebaseAuthenticationSettingsIos) GetDeviceToken() (value string) {
	if f == nil {
		return
	}
	return f.DeviceToken
}

// GetIsAppSandbox returns value of IsAppSandbox field.
func (f *FirebaseAuthenticationSettingsIos) GetIsAppSandbox() (value bool) {
	if f == nil {
		return
	}
	return f.IsAppSandbox
}

// FirebaseAuthenticationSettingsClassName is schema name of FirebaseAuthenticationSettingsClass.
const FirebaseAuthenticationSettingsClassName = "FirebaseAuthenticationSettings"

// FirebaseAuthenticationSettingsClass represents FirebaseAuthenticationSettings generic type.
//
// Example:
//
//	g, err := tdapi.DecodeFirebaseAuthenticationSettings(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.FirebaseAuthenticationSettingsAndroid: // firebaseAuthenticationSettingsAndroid#966ef61c
//	case *tdapi.FirebaseAuthenticationSettingsIos: // firebaseAuthenticationSettingsIos#d49a4c4
//	default: panic(v)
//	}
type FirebaseAuthenticationSettingsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() FirebaseAuthenticationSettingsClass

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

// DecodeFirebaseAuthenticationSettings implements binary de-serialization for FirebaseAuthenticationSettingsClass.
func DecodeFirebaseAuthenticationSettings(buf *bin.Buffer) (FirebaseAuthenticationSettingsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case FirebaseAuthenticationSettingsAndroidTypeID:
		// Decoding firebaseAuthenticationSettingsAndroid#966ef61c.
		v := FirebaseAuthenticationSettingsAndroid{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode FirebaseAuthenticationSettingsClass: %w", err)
		}
		return &v, nil
	case FirebaseAuthenticationSettingsIosTypeID:
		// Decoding firebaseAuthenticationSettingsIos#d49a4c4.
		v := FirebaseAuthenticationSettingsIos{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode FirebaseAuthenticationSettingsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode FirebaseAuthenticationSettingsClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONFirebaseAuthenticationSettings implements binary de-serialization for FirebaseAuthenticationSettingsClass.
func DecodeTDLibJSONFirebaseAuthenticationSettings(buf tdjson.Decoder) (FirebaseAuthenticationSettingsClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "firebaseAuthenticationSettingsAndroid":
		// Decoding firebaseAuthenticationSettingsAndroid#966ef61c.
		v := FirebaseAuthenticationSettingsAndroid{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode FirebaseAuthenticationSettingsClass: %w", err)
		}
		return &v, nil
	case "firebaseAuthenticationSettingsIos":
		// Decoding firebaseAuthenticationSettingsIos#d49a4c4.
		v := FirebaseAuthenticationSettingsIos{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode FirebaseAuthenticationSettingsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode FirebaseAuthenticationSettingsClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// FirebaseAuthenticationSettings boxes the FirebaseAuthenticationSettingsClass providing a helper.
type FirebaseAuthenticationSettingsBox struct {
	FirebaseAuthenticationSettings FirebaseAuthenticationSettingsClass
}

// Decode implements bin.Decoder for FirebaseAuthenticationSettingsBox.
func (b *FirebaseAuthenticationSettingsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode FirebaseAuthenticationSettingsBox to nil")
	}
	v, err := DecodeFirebaseAuthenticationSettings(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FirebaseAuthenticationSettings = v
	return nil
}

// Encode implements bin.Encode for FirebaseAuthenticationSettingsBox.
func (b *FirebaseAuthenticationSettingsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.FirebaseAuthenticationSettings == nil {
		return fmt.Errorf("unable to encode FirebaseAuthenticationSettingsClass as nil")
	}
	return b.FirebaseAuthenticationSettings.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for FirebaseAuthenticationSettingsBox.
func (b *FirebaseAuthenticationSettingsBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode FirebaseAuthenticationSettingsBox to nil")
	}
	v, err := DecodeTDLibJSONFirebaseAuthenticationSettings(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.FirebaseAuthenticationSettings = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for FirebaseAuthenticationSettingsBox.
func (b *FirebaseAuthenticationSettingsBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.FirebaseAuthenticationSettings == nil {
		return fmt.Errorf("unable to encode FirebaseAuthenticationSettingsClass as nil")
	}
	return b.FirebaseAuthenticationSettings.EncodeTDLibJSON(buf)
}
