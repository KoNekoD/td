// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// AuthLoginToken represents TL type `auth.loginToken#629f1980`.
// Login token (for QR code login¹)
//
// Links:
//  1) https://core.telegram.org/api/qr-login
//
// See https://core.telegram.org/constructor/auth.loginToken for reference.
type AuthLoginToken struct {
	// Expiry date of QR code
	Expires int
	// Token to render in QR code
	Token []byte
}

// AuthLoginTokenTypeID is TL type id of AuthLoginToken.
const AuthLoginTokenTypeID = 0x629f1980

// construct implements constructor of AuthLoginTokenClass.
func (l AuthLoginToken) construct() AuthLoginTokenClass { return &l }

// Ensuring interfaces in compile-time for AuthLoginToken.
var (
	_ bin.Encoder     = &AuthLoginToken{}
	_ bin.Decoder     = &AuthLoginToken{}
	_ bin.BareEncoder = &AuthLoginToken{}
	_ bin.BareDecoder = &AuthLoginToken{}

	_ AuthLoginTokenClass = &AuthLoginToken{}
)

func (l *AuthLoginToken) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Expires == 0) {
		return false
	}
	if !(l.Token == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *AuthLoginToken) String() string {
	if l == nil {
		return "AuthLoginToken(nil)"
	}
	type Alias AuthLoginToken
	return fmt.Sprintf("AuthLoginToken%+v", Alias(*l))
}

// FillFrom fills AuthLoginToken from given interface.
func (l *AuthLoginToken) FillFrom(from interface {
	GetExpires() (value int)
	GetToken() (value []byte)
}) {
	l.Expires = from.GetExpires()
	l.Token = from.GetToken()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthLoginToken) TypeID() uint32 {
	return AuthLoginTokenTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthLoginToken) TypeName() string {
	return "auth.loginToken"
}

// TypeInfo returns info about TL type.
func (l *AuthLoginToken) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.loginToken",
		ID:   AuthLoginTokenTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
		{
			Name:       "Token",
			SchemaName: "token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *AuthLoginToken) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.loginToken#629f1980 as nil")
	}
	b.PutID(AuthLoginTokenTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *AuthLoginToken) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.loginToken#629f1980 as nil")
	}
	b.PutInt(l.Expires)
	b.PutBytes(l.Token)
	return nil
}

// Decode implements bin.Decoder.
func (l *AuthLoginToken) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.loginToken#629f1980 to nil")
	}
	if err := b.ConsumeID(AuthLoginTokenTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.loginToken#629f1980: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *AuthLoginToken) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.loginToken#629f1980 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.loginToken#629f1980: field expires: %w", err)
		}
		l.Expires = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.loginToken#629f1980: field token: %w", err)
		}
		l.Token = value
	}
	return nil
}

// GetExpires returns value of Expires field.
func (l *AuthLoginToken) GetExpires() (value int) {
	return l.Expires
}

// GetToken returns value of Token field.
func (l *AuthLoginToken) GetToken() (value []byte) {
	return l.Token
}

// AuthLoginTokenMigrateTo represents TL type `auth.loginTokenMigrateTo#68e9916`.
// Repeat the query to the specified DC
//
// See https://core.telegram.org/constructor/auth.loginTokenMigrateTo for reference.
type AuthLoginTokenMigrateTo struct {
	// DC ID
	DCID int
	// Token to use for login
	Token []byte
}

// AuthLoginTokenMigrateToTypeID is TL type id of AuthLoginTokenMigrateTo.
const AuthLoginTokenMigrateToTypeID = 0x68e9916

// construct implements constructor of AuthLoginTokenClass.
func (l AuthLoginTokenMigrateTo) construct() AuthLoginTokenClass { return &l }

// Ensuring interfaces in compile-time for AuthLoginTokenMigrateTo.
var (
	_ bin.Encoder     = &AuthLoginTokenMigrateTo{}
	_ bin.Decoder     = &AuthLoginTokenMigrateTo{}
	_ bin.BareEncoder = &AuthLoginTokenMigrateTo{}
	_ bin.BareDecoder = &AuthLoginTokenMigrateTo{}

	_ AuthLoginTokenClass = &AuthLoginTokenMigrateTo{}
)

func (l *AuthLoginTokenMigrateTo) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.DCID == 0) {
		return false
	}
	if !(l.Token == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *AuthLoginTokenMigrateTo) String() string {
	if l == nil {
		return "AuthLoginTokenMigrateTo(nil)"
	}
	type Alias AuthLoginTokenMigrateTo
	return fmt.Sprintf("AuthLoginTokenMigrateTo%+v", Alias(*l))
}

// FillFrom fills AuthLoginTokenMigrateTo from given interface.
func (l *AuthLoginTokenMigrateTo) FillFrom(from interface {
	GetDCID() (value int)
	GetToken() (value []byte)
}) {
	l.DCID = from.GetDCID()
	l.Token = from.GetToken()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthLoginTokenMigrateTo) TypeID() uint32 {
	return AuthLoginTokenMigrateToTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthLoginTokenMigrateTo) TypeName() string {
	return "auth.loginTokenMigrateTo"
}

// TypeInfo returns info about TL type.
func (l *AuthLoginTokenMigrateTo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.loginTokenMigrateTo",
		ID:   AuthLoginTokenMigrateToTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
		{
			Name:       "Token",
			SchemaName: "token",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *AuthLoginTokenMigrateTo) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.loginTokenMigrateTo#68e9916 as nil")
	}
	b.PutID(AuthLoginTokenMigrateToTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *AuthLoginTokenMigrateTo) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.loginTokenMigrateTo#68e9916 as nil")
	}
	b.PutInt(l.DCID)
	b.PutBytes(l.Token)
	return nil
}

// Decode implements bin.Decoder.
func (l *AuthLoginTokenMigrateTo) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.loginTokenMigrateTo#68e9916 to nil")
	}
	if err := b.ConsumeID(AuthLoginTokenMigrateToTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.loginTokenMigrateTo#68e9916: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *AuthLoginTokenMigrateTo) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.loginTokenMigrateTo#68e9916 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode auth.loginTokenMigrateTo#68e9916: field dc_id: %w", err)
		}
		l.DCID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode auth.loginTokenMigrateTo#68e9916: field token: %w", err)
		}
		l.Token = value
	}
	return nil
}

// GetDCID returns value of DCID field.
func (l *AuthLoginTokenMigrateTo) GetDCID() (value int) {
	return l.DCID
}

// GetToken returns value of Token field.
func (l *AuthLoginTokenMigrateTo) GetToken() (value []byte) {
	return l.Token
}

// AuthLoginTokenSuccess represents TL type `auth.loginTokenSuccess#390d5c5e`.
// Login via token (QR code) succeded!
//
// See https://core.telegram.org/constructor/auth.loginTokenSuccess for reference.
type AuthLoginTokenSuccess struct {
	// Authorization info
	Authorization AuthAuthorizationClass
}

// AuthLoginTokenSuccessTypeID is TL type id of AuthLoginTokenSuccess.
const AuthLoginTokenSuccessTypeID = 0x390d5c5e

// construct implements constructor of AuthLoginTokenClass.
func (l AuthLoginTokenSuccess) construct() AuthLoginTokenClass { return &l }

// Ensuring interfaces in compile-time for AuthLoginTokenSuccess.
var (
	_ bin.Encoder     = &AuthLoginTokenSuccess{}
	_ bin.Decoder     = &AuthLoginTokenSuccess{}
	_ bin.BareEncoder = &AuthLoginTokenSuccess{}
	_ bin.BareDecoder = &AuthLoginTokenSuccess{}

	_ AuthLoginTokenClass = &AuthLoginTokenSuccess{}
)

func (l *AuthLoginTokenSuccess) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Authorization == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *AuthLoginTokenSuccess) String() string {
	if l == nil {
		return "AuthLoginTokenSuccess(nil)"
	}
	type Alias AuthLoginTokenSuccess
	return fmt.Sprintf("AuthLoginTokenSuccess%+v", Alias(*l))
}

// FillFrom fills AuthLoginTokenSuccess from given interface.
func (l *AuthLoginTokenSuccess) FillFrom(from interface {
	GetAuthorization() (value AuthAuthorizationClass)
}) {
	l.Authorization = from.GetAuthorization()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AuthLoginTokenSuccess) TypeID() uint32 {
	return AuthLoginTokenSuccessTypeID
}

// TypeName returns name of type in TL schema.
func (*AuthLoginTokenSuccess) TypeName() string {
	return "auth.loginTokenSuccess"
}

// TypeInfo returns info about TL type.
func (l *AuthLoginTokenSuccess) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "auth.loginTokenSuccess",
		ID:   AuthLoginTokenSuccessTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Authorization",
			SchemaName: "authorization",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *AuthLoginTokenSuccess) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.loginTokenSuccess#390d5c5e as nil")
	}
	b.PutID(AuthLoginTokenSuccessTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *AuthLoginTokenSuccess) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode auth.loginTokenSuccess#390d5c5e as nil")
	}
	if l.Authorization == nil {
		return fmt.Errorf("unable to encode auth.loginTokenSuccess#390d5c5e: field authorization is nil")
	}
	if err := l.Authorization.Encode(b); err != nil {
		return fmt.Errorf("unable to encode auth.loginTokenSuccess#390d5c5e: field authorization: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *AuthLoginTokenSuccess) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.loginTokenSuccess#390d5c5e to nil")
	}
	if err := b.ConsumeID(AuthLoginTokenSuccessTypeID); err != nil {
		return fmt.Errorf("unable to decode auth.loginTokenSuccess#390d5c5e: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *AuthLoginTokenSuccess) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode auth.loginTokenSuccess#390d5c5e to nil")
	}
	{
		value, err := DecodeAuthAuthorization(b)
		if err != nil {
			return fmt.Errorf("unable to decode auth.loginTokenSuccess#390d5c5e: field authorization: %w", err)
		}
		l.Authorization = value
	}
	return nil
}

// GetAuthorization returns value of Authorization field.
func (l *AuthLoginTokenSuccess) GetAuthorization() (value AuthAuthorizationClass) {
	return l.Authorization
}

// AuthLoginTokenClass represents auth.LoginToken generic type.
//
// See https://core.telegram.org/type/auth.LoginToken for reference.
//
// Example:
//  g, err := tg.DecodeAuthLoginToken(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.AuthLoginToken: // auth.loginToken#629f1980
//  case *tg.AuthLoginTokenMigrateTo: // auth.loginTokenMigrateTo#68e9916
//  case *tg.AuthLoginTokenSuccess: // auth.loginTokenSuccess#390d5c5e
//  default: panic(v)
//  }
type AuthLoginTokenClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() AuthLoginTokenClass

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

// DecodeAuthLoginToken implements binary de-serialization for AuthLoginTokenClass.
func DecodeAuthLoginToken(buf *bin.Buffer) (AuthLoginTokenClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case AuthLoginTokenTypeID:
		// Decoding auth.loginToken#629f1980.
		v := AuthLoginToken{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthLoginTokenClass: %w", err)
		}
		return &v, nil
	case AuthLoginTokenMigrateToTypeID:
		// Decoding auth.loginTokenMigrateTo#68e9916.
		v := AuthLoginTokenMigrateTo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthLoginTokenClass: %w", err)
		}
		return &v, nil
	case AuthLoginTokenSuccessTypeID:
		// Decoding auth.loginTokenSuccess#390d5c5e.
		v := AuthLoginTokenSuccess{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode AuthLoginTokenClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode AuthLoginTokenClass: %w", bin.NewUnexpectedID(id))
	}
}

// AuthLoginToken boxes the AuthLoginTokenClass providing a helper.
type AuthLoginTokenBox struct {
	LoginToken AuthLoginTokenClass
}

// Decode implements bin.Decoder for AuthLoginTokenBox.
func (b *AuthLoginTokenBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode AuthLoginTokenBox to nil")
	}
	v, err := DecodeAuthLoginToken(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.LoginToken = v
	return nil
}

// Encode implements bin.Encode for AuthLoginTokenBox.
func (b *AuthLoginTokenBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.LoginToken == nil {
		return fmt.Errorf("unable to encode AuthLoginTokenClass as nil")
	}
	return b.LoginToken.Encode(buf)
}
