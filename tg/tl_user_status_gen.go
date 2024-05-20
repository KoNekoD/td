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

// UserStatusEmpty represents TL type `userStatusEmpty#9d05049`.
// User status has not been set yet.
//
// See https://core.telegram.org/constructor/userStatusEmpty for reference.
type UserStatusEmpty struct {
}

// UserStatusEmptyTypeID is TL type id of UserStatusEmpty.
const UserStatusEmptyTypeID = 0x9d05049

// construct implements constructor of UserStatusClass.
func (u UserStatusEmpty) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusEmpty.
var (
	_ bin.Encoder     = &UserStatusEmpty{}
	_ bin.Decoder     = &UserStatusEmpty{}
	_ bin.BareEncoder = &UserStatusEmpty{}
	_ bin.BareDecoder = &UserStatusEmpty{}

	_ UserStatusClass = &UserStatusEmpty{}
)

func (u *UserStatusEmpty) Zero() bool {
	if u == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserStatusEmpty) String() string {
	if u == nil {
		return "UserStatusEmpty(nil)"
	}
	type Alias UserStatusEmpty
	return fmt.Sprintf("UserStatusEmpty%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserStatusEmpty) TypeID() uint32 {
	return UserStatusEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*UserStatusEmpty) TypeName() string {
	return "userStatusEmpty"
}

// TypeInfo returns info about TL type.
func (u *UserStatusEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userStatusEmpty",
		ID:   UserStatusEmptyTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserStatusEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusEmpty#9d05049 as nil")
	}
	b.PutID(UserStatusEmptyTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserStatusEmpty) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusEmpty#9d05049 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusEmpty) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusEmpty#9d05049 to nil")
	}
	if err := b.ConsumeID(UserStatusEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusEmpty#9d05049: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserStatusEmpty) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusEmpty#9d05049 to nil")
	}
	return nil
}

// UserStatusOnline represents TL type `userStatusOnline#edb93949`.
// Online status of the user.
//
// See https://core.telegram.org/constructor/userStatusOnline for reference.
type UserStatusOnline struct {
	// Time to expiration of the current online status
	Expires int
}

// UserStatusOnlineTypeID is TL type id of UserStatusOnline.
const UserStatusOnlineTypeID = 0xedb93949

// construct implements constructor of UserStatusClass.
func (u UserStatusOnline) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusOnline.
var (
	_ bin.Encoder     = &UserStatusOnline{}
	_ bin.Decoder     = &UserStatusOnline{}
	_ bin.BareEncoder = &UserStatusOnline{}
	_ bin.BareDecoder = &UserStatusOnline{}

	_ UserStatusClass = &UserStatusOnline{}
)

func (u *UserStatusOnline) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Expires == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserStatusOnline) String() string {
	if u == nil {
		return "UserStatusOnline(nil)"
	}
	type Alias UserStatusOnline
	return fmt.Sprintf("UserStatusOnline%+v", Alias(*u))
}

// FillFrom fills UserStatusOnline from given interface.
func (u *UserStatusOnline) FillFrom(from interface {
	GetExpires() (value int)
}) {
	u.Expires = from.GetExpires()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserStatusOnline) TypeID() uint32 {
	return UserStatusOnlineTypeID
}

// TypeName returns name of type in TL schema.
func (*UserStatusOnline) TypeName() string {
	return "userStatusOnline"
}

// TypeInfo returns info about TL type.
func (u *UserStatusOnline) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userStatusOnline",
		ID:   UserStatusOnlineTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Expires",
			SchemaName: "expires",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserStatusOnline) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusOnline#edb93949 as nil")
	}
	b.PutID(UserStatusOnlineTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserStatusOnline) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusOnline#edb93949 as nil")
	}
	b.PutInt(u.Expires)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusOnline) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusOnline#edb93949 to nil")
	}
	if err := b.ConsumeID(UserStatusOnlineTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusOnline#edb93949: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserStatusOnline) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusOnline#edb93949 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userStatusOnline#edb93949: field expires: %w", err)
		}
		u.Expires = value
	}
	return nil
}

// GetExpires returns value of Expires field.
func (u *UserStatusOnline) GetExpires() (value int) {
	if u == nil {
		return
	}
	return u.Expires
}

// UserStatusOffline represents TL type `userStatusOffline#8c703f`.
// The user's offline status.
//
// See https://core.telegram.org/constructor/userStatusOffline for reference.
type UserStatusOffline struct {
	// Time the user was last seen online
	WasOnline int
}

// UserStatusOfflineTypeID is TL type id of UserStatusOffline.
const UserStatusOfflineTypeID = 0x8c703f

// construct implements constructor of UserStatusClass.
func (u UserStatusOffline) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusOffline.
var (
	_ bin.Encoder     = &UserStatusOffline{}
	_ bin.Decoder     = &UserStatusOffline{}
	_ bin.BareEncoder = &UserStatusOffline{}
	_ bin.BareDecoder = &UserStatusOffline{}

	_ UserStatusClass = &UserStatusOffline{}
)

func (u *UserStatusOffline) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.WasOnline == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserStatusOffline) String() string {
	if u == nil {
		return "UserStatusOffline(nil)"
	}
	type Alias UserStatusOffline
	return fmt.Sprintf("UserStatusOffline%+v", Alias(*u))
}

// FillFrom fills UserStatusOffline from given interface.
func (u *UserStatusOffline) FillFrom(from interface {
	GetWasOnline() (value int)
}) {
	u.WasOnline = from.GetWasOnline()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserStatusOffline) TypeID() uint32 {
	return UserStatusOfflineTypeID
}

// TypeName returns name of type in TL schema.
func (*UserStatusOffline) TypeName() string {
	return "userStatusOffline"
}

// TypeInfo returns info about TL type.
func (u *UserStatusOffline) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userStatusOffline",
		ID:   UserStatusOfflineTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "WasOnline",
			SchemaName: "was_online",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserStatusOffline) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusOffline#8c703f as nil")
	}
	b.PutID(UserStatusOfflineTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserStatusOffline) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusOffline#8c703f as nil")
	}
	b.PutInt(u.WasOnline)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusOffline) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusOffline#8c703f to nil")
	}
	if err := b.ConsumeID(UserStatusOfflineTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusOffline#8c703f: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserStatusOffline) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusOffline#8c703f to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userStatusOffline#8c703f: field was_online: %w", err)
		}
		u.WasOnline = value
	}
	return nil
}

// GetWasOnline returns value of WasOnline field.
func (u *UserStatusOffline) GetWasOnline() (value int) {
	if u == nil {
		return
	}
	return u.WasOnline
}

// UserStatusRecently represents TL type `userStatusRecently#7b197dc8`.
// Online status: last seen recently
//
// See https://core.telegram.org/constructor/userStatusRecently for reference.
type UserStatusRecently struct {
	// Flags field of UserStatusRecently.
	Flags bin.Fields
	// ByMe field of UserStatusRecently.
	ByMe bool
}

// UserStatusRecentlyTypeID is TL type id of UserStatusRecently.
const UserStatusRecentlyTypeID = 0x7b197dc8

// construct implements constructor of UserStatusClass.
func (u UserStatusRecently) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusRecently.
var (
	_ bin.Encoder     = &UserStatusRecently{}
	_ bin.Decoder     = &UserStatusRecently{}
	_ bin.BareEncoder = &UserStatusRecently{}
	_ bin.BareDecoder = &UserStatusRecently{}

	_ UserStatusClass = &UserStatusRecently{}
)

func (u *UserStatusRecently) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.ByMe == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserStatusRecently) String() string {
	if u == nil {
		return "UserStatusRecently(nil)"
	}
	type Alias UserStatusRecently
	return fmt.Sprintf("UserStatusRecently%+v", Alias(*u))
}

// FillFrom fills UserStatusRecently from given interface.
func (u *UserStatusRecently) FillFrom(from interface {
	GetByMe() (value bool)
}) {
	u.ByMe = from.GetByMe()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserStatusRecently) TypeID() uint32 {
	return UserStatusRecentlyTypeID
}

// TypeName returns name of type in TL schema.
func (*UserStatusRecently) TypeName() string {
	return "userStatusRecently"
}

// TypeInfo returns info about TL type.
func (u *UserStatusRecently) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userStatusRecently",
		ID:   UserStatusRecentlyTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ByMe",
			SchemaName: "by_me",
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *UserStatusRecently) SetFlags() {
	if !(u.ByMe == false) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *UserStatusRecently) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusRecently#7b197dc8 as nil")
	}
	b.PutID(UserStatusRecentlyTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserStatusRecently) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusRecently#7b197dc8 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userStatusRecently#7b197dc8: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusRecently) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusRecently#7b197dc8 to nil")
	}
	if err := b.ConsumeID(UserStatusRecentlyTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusRecently#7b197dc8: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserStatusRecently) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusRecently#7b197dc8 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userStatusRecently#7b197dc8: field flags: %w", err)
		}
	}
	u.ByMe = u.Flags.Has(0)
	return nil
}

// SetByMe sets value of ByMe conditional field.
func (u *UserStatusRecently) SetByMe(value bool) {
	if value {
		u.Flags.Set(0)
		u.ByMe = true
	} else {
		u.Flags.Unset(0)
		u.ByMe = false
	}
}

// GetByMe returns value of ByMe conditional field.
func (u *UserStatusRecently) GetByMe() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// UserStatusLastWeek represents TL type `userStatusLastWeek#541a1d1a`.
// Online status: last seen last week
//
// See https://core.telegram.org/constructor/userStatusLastWeek for reference.
type UserStatusLastWeek struct {
	// Flags field of UserStatusLastWeek.
	Flags bin.Fields
	// ByMe field of UserStatusLastWeek.
	ByMe bool
}

// UserStatusLastWeekTypeID is TL type id of UserStatusLastWeek.
const UserStatusLastWeekTypeID = 0x541a1d1a

// construct implements constructor of UserStatusClass.
func (u UserStatusLastWeek) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusLastWeek.
var (
	_ bin.Encoder     = &UserStatusLastWeek{}
	_ bin.Decoder     = &UserStatusLastWeek{}
	_ bin.BareEncoder = &UserStatusLastWeek{}
	_ bin.BareDecoder = &UserStatusLastWeek{}

	_ UserStatusClass = &UserStatusLastWeek{}
)

func (u *UserStatusLastWeek) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.ByMe == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserStatusLastWeek) String() string {
	if u == nil {
		return "UserStatusLastWeek(nil)"
	}
	type Alias UserStatusLastWeek
	return fmt.Sprintf("UserStatusLastWeek%+v", Alias(*u))
}

// FillFrom fills UserStatusLastWeek from given interface.
func (u *UserStatusLastWeek) FillFrom(from interface {
	GetByMe() (value bool)
}) {
	u.ByMe = from.GetByMe()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserStatusLastWeek) TypeID() uint32 {
	return UserStatusLastWeekTypeID
}

// TypeName returns name of type in TL schema.
func (*UserStatusLastWeek) TypeName() string {
	return "userStatusLastWeek"
}

// TypeInfo returns info about TL type.
func (u *UserStatusLastWeek) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userStatusLastWeek",
		ID:   UserStatusLastWeekTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ByMe",
			SchemaName: "by_me",
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *UserStatusLastWeek) SetFlags() {
	if !(u.ByMe == false) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *UserStatusLastWeek) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusLastWeek#541a1d1a as nil")
	}
	b.PutID(UserStatusLastWeekTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserStatusLastWeek) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusLastWeek#541a1d1a as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userStatusLastWeek#541a1d1a: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusLastWeek) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusLastWeek#541a1d1a to nil")
	}
	if err := b.ConsumeID(UserStatusLastWeekTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusLastWeek#541a1d1a: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserStatusLastWeek) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusLastWeek#541a1d1a to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userStatusLastWeek#541a1d1a: field flags: %w", err)
		}
	}
	u.ByMe = u.Flags.Has(0)
	return nil
}

// SetByMe sets value of ByMe conditional field.
func (u *UserStatusLastWeek) SetByMe(value bool) {
	if value {
		u.Flags.Set(0)
		u.ByMe = true
	} else {
		u.Flags.Unset(0)
		u.ByMe = false
	}
}

// GetByMe returns value of ByMe conditional field.
func (u *UserStatusLastWeek) GetByMe() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// UserStatusLastMonth represents TL type `userStatusLastMonth#65899777`.
// Online status: last seen last month
//
// See https://core.telegram.org/constructor/userStatusLastMonth for reference.
type UserStatusLastMonth struct {
	// Flags field of UserStatusLastMonth.
	Flags bin.Fields
	// ByMe field of UserStatusLastMonth.
	ByMe bool
}

// UserStatusLastMonthTypeID is TL type id of UserStatusLastMonth.
const UserStatusLastMonthTypeID = 0x65899777

// construct implements constructor of UserStatusClass.
func (u UserStatusLastMonth) construct() UserStatusClass { return &u }

// Ensuring interfaces in compile-time for UserStatusLastMonth.
var (
	_ bin.Encoder     = &UserStatusLastMonth{}
	_ bin.Decoder     = &UserStatusLastMonth{}
	_ bin.BareEncoder = &UserStatusLastMonth{}
	_ bin.BareDecoder = &UserStatusLastMonth{}

	_ UserStatusClass = &UserStatusLastMonth{}
)

func (u *UserStatusLastMonth) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.ByMe == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserStatusLastMonth) String() string {
	if u == nil {
		return "UserStatusLastMonth(nil)"
	}
	type Alias UserStatusLastMonth
	return fmt.Sprintf("UserStatusLastMonth%+v", Alias(*u))
}

// FillFrom fills UserStatusLastMonth from given interface.
func (u *UserStatusLastMonth) FillFrom(from interface {
	GetByMe() (value bool)
}) {
	u.ByMe = from.GetByMe()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserStatusLastMonth) TypeID() uint32 {
	return UserStatusLastMonthTypeID
}

// TypeName returns name of type in TL schema.
func (*UserStatusLastMonth) TypeName() string {
	return "userStatusLastMonth"
}

// TypeInfo returns info about TL type.
func (u *UserStatusLastMonth) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userStatusLastMonth",
		ID:   UserStatusLastMonthTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ByMe",
			SchemaName: "by_me",
			Null:       !u.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *UserStatusLastMonth) SetFlags() {
	if !(u.ByMe == false) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *UserStatusLastMonth) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusLastMonth#65899777 as nil")
	}
	b.PutID(UserStatusLastMonthTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserStatusLastMonth) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userStatusLastMonth#65899777 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userStatusLastMonth#65899777: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UserStatusLastMonth) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusLastMonth#65899777 to nil")
	}
	if err := b.ConsumeID(UserStatusLastMonthTypeID); err != nil {
		return fmt.Errorf("unable to decode userStatusLastMonth#65899777: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserStatusLastMonth) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userStatusLastMonth#65899777 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userStatusLastMonth#65899777: field flags: %w", err)
		}
	}
	u.ByMe = u.Flags.Has(0)
	return nil
}

// SetByMe sets value of ByMe conditional field.
func (u *UserStatusLastMonth) SetByMe(value bool) {
	if value {
		u.Flags.Set(0)
		u.ByMe = true
	} else {
		u.Flags.Unset(0)
		u.ByMe = false
	}
}

// GetByMe returns value of ByMe conditional field.
func (u *UserStatusLastMonth) GetByMe() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// UserStatusClassName is schema name of UserStatusClass.
const UserStatusClassName = "UserStatus"

// UserStatusClass represents UserStatus generic type.
//
// See https://core.telegram.org/type/UserStatus for reference.
//
// Example:
//
//	g, err := tg.DecodeUserStatus(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.UserStatusEmpty: // userStatusEmpty#9d05049
//	case *tg.UserStatusOnline: // userStatusOnline#edb93949
//	case *tg.UserStatusOffline: // userStatusOffline#8c703f
//	case *tg.UserStatusRecently: // userStatusRecently#7b197dc8
//	case *tg.UserStatusLastWeek: // userStatusLastWeek#541a1d1a
//	case *tg.UserStatusLastMonth: // userStatusLastMonth#65899777
//	default: panic(v)
//	}
type UserStatusClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() UserStatusClass

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

// DecodeUserStatus implements binary de-serialization for UserStatusClass.
func DecodeUserStatus(buf *bin.Buffer) (UserStatusClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UserStatusEmptyTypeID:
		// Decoding userStatusEmpty#9d05049.
		v := UserStatusEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	case UserStatusOnlineTypeID:
		// Decoding userStatusOnline#edb93949.
		v := UserStatusOnline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	case UserStatusOfflineTypeID:
		// Decoding userStatusOffline#8c703f.
		v := UserStatusOffline{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	case UserStatusRecentlyTypeID:
		// Decoding userStatusRecently#7b197dc8.
		v := UserStatusRecently{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	case UserStatusLastWeekTypeID:
		// Decoding userStatusLastWeek#541a1d1a.
		v := UserStatusLastWeek{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	case UserStatusLastMonthTypeID:
		// Decoding userStatusLastMonth#65899777.
		v := UserStatusLastMonth{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserStatusClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UserStatusClass: %w", bin.NewUnexpectedID(id))
	}
}

// UserStatus boxes the UserStatusClass providing a helper.
type UserStatusBox struct {
	UserStatus UserStatusClass
}

// Decode implements bin.Decoder for UserStatusBox.
func (b *UserStatusBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UserStatusBox to nil")
	}
	v, err := DecodeUserStatus(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.UserStatus = v
	return nil
}

// Encode implements bin.Encode for UserStatusBox.
func (b *UserStatusBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.UserStatus == nil {
		return fmt.Errorf("unable to encode UserStatusClass as nil")
	}
	return b.UserStatus.Encode(buf)
}
