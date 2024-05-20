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

// UsersUserFull represents TL type `users.userFull#3b6d152e`.
// Full user information
//
// See https://core.telegram.org/constructor/users.userFull for reference.
type UsersUserFull struct {
	// Full user information
	FullUser UserFull
	// Mentioned chats
	Chats []ChatClass
	// Mentioned users
	Users []UserClass
}

// UsersUserFullTypeID is TL type id of UsersUserFull.
const UsersUserFullTypeID = 0x3b6d152e

// Ensuring interfaces in compile-time for UsersUserFull.
var (
	_ bin.Encoder     = &UsersUserFull{}
	_ bin.Decoder     = &UsersUserFull{}
	_ bin.BareEncoder = &UsersUserFull{}
	_ bin.BareDecoder = &UsersUserFull{}
)

func (u *UsersUserFull) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.FullUser.Zero()) {
		return false
	}
	if !(u.Chats == nil) {
		return false
	}
	if !(u.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UsersUserFull) String() string {
	if u == nil {
		return "UsersUserFull(nil)"
	}
	type Alias UsersUserFull
	return fmt.Sprintf("UsersUserFull%+v", Alias(*u))
}

// FillFrom fills UsersUserFull from given interface.
func (u *UsersUserFull) FillFrom(from interface {
	GetFullUser() (value UserFull)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	u.FullUser = from.GetFullUser()
	u.Chats = from.GetChats()
	u.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UsersUserFull) TypeID() uint32 {
	return UsersUserFullTypeID
}

// TypeName returns name of type in TL schema.
func (*UsersUserFull) TypeName() string {
	return "users.userFull"
}

// TypeInfo returns info about TL type.
func (u *UsersUserFull) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "users.userFull",
		ID:   UsersUserFullTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FullUser",
			SchemaName: "full_user",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UsersUserFull) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode users.userFull#3b6d152e as nil")
	}
	b.PutID(UsersUserFullTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UsersUserFull) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode users.userFull#3b6d152e as nil")
	}
	if err := u.FullUser.Encode(b); err != nil {
		return fmt.Errorf("unable to encode users.userFull#3b6d152e: field full_user: %w", err)
	}
	b.PutVectorHeader(len(u.Chats))
	for idx, v := range u.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode users.userFull#3b6d152e: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode users.userFull#3b6d152e: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(u.Users))
	for idx, v := range u.Users {
		if v == nil {
			return fmt.Errorf("unable to encode users.userFull#3b6d152e: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode users.userFull#3b6d152e: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UsersUserFull) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode users.userFull#3b6d152e to nil")
	}
	if err := b.ConsumeID(UsersUserFullTypeID); err != nil {
		return fmt.Errorf("unable to decode users.userFull#3b6d152e: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UsersUserFull) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode users.userFull#3b6d152e to nil")
	}
	{
		if err := u.FullUser.Decode(b); err != nil {
			return fmt.Errorf("unable to decode users.userFull#3b6d152e: field full_user: %w", err)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode users.userFull#3b6d152e: field chats: %w", err)
		}

		if headerLen > 0 {
			u.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode users.userFull#3b6d152e: field chats: %w", err)
			}
			u.Chats = append(u.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode users.userFull#3b6d152e: field users: %w", err)
		}

		if headerLen > 0 {
			u.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode users.userFull#3b6d152e: field users: %w", err)
			}
			u.Users = append(u.Users, value)
		}
	}
	return nil
}

// GetFullUser returns value of FullUser field.
func (u *UsersUserFull) GetFullUser() (value UserFull) {
	if u == nil {
		return
	}
	return u.FullUser
}

// GetChats returns value of Chats field.
func (u *UsersUserFull) GetChats() (value []ChatClass) {
	if u == nil {
		return
	}
	return u.Chats
}

// GetUsers returns value of Users field.
func (u *UsersUserFull) GetUsers() (value []UserClass) {
	if u == nil {
		return
	}
	return u.Users
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (u *UsersUserFull) MapChats() (value ChatClassArray) {
	return ChatClassArray(u.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (u *UsersUserFull) MapUsers() (value UserClassArray) {
	return UserClassArray(u.Users)
}
