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

// RequestPeerTypeUser represents TL type `requestPeerTypeUser#5f3b8a00`.
// Choose a user.
//
// See https://core.telegram.org/constructor/requestPeerTypeUser for reference.
type RequestPeerTypeUser struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to allow choosing only bots.
	//
	// Use SetBot and GetBot helpers.
	Bot bool
	// Whether to allow choosing only Premium¹ users.
	//
	// Links:
	//  1) https://core.telegram.org/api/premium
	//
	// Use SetPremium and GetPremium helpers.
	Premium bool
}

// RequestPeerTypeUserTypeID is TL type id of RequestPeerTypeUser.
const RequestPeerTypeUserTypeID = 0x5f3b8a00

// construct implements constructor of RequestPeerTypeClass.
func (r RequestPeerTypeUser) construct() RequestPeerTypeClass { return &r }

// Ensuring interfaces in compile-time for RequestPeerTypeUser.
var (
	_ bin.Encoder     = &RequestPeerTypeUser{}
	_ bin.Decoder     = &RequestPeerTypeUser{}
	_ bin.BareEncoder = &RequestPeerTypeUser{}
	_ bin.BareDecoder = &RequestPeerTypeUser{}

	_ RequestPeerTypeClass = &RequestPeerTypeUser{}
)

func (r *RequestPeerTypeUser) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Bot == false) {
		return false
	}
	if !(r.Premium == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RequestPeerTypeUser) String() string {
	if r == nil {
		return "RequestPeerTypeUser(nil)"
	}
	type Alias RequestPeerTypeUser
	return fmt.Sprintf("RequestPeerTypeUser%+v", Alias(*r))
}

// FillFrom fills RequestPeerTypeUser from given interface.
func (r *RequestPeerTypeUser) FillFrom(from interface {
	GetBot() (value bool, ok bool)
	GetPremium() (value bool, ok bool)
}) {
	if val, ok := from.GetBot(); ok {
		r.Bot = val
	}

	if val, ok := from.GetPremium(); ok {
		r.Premium = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RequestPeerTypeUser) TypeID() uint32 {
	return RequestPeerTypeUserTypeID
}

// TypeName returns name of type in TL schema.
func (*RequestPeerTypeUser) TypeName() string {
	return "requestPeerTypeUser"
}

// TypeInfo returns info about TL type.
func (r *RequestPeerTypeUser) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "requestPeerTypeUser",
		ID:   RequestPeerTypeUserTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Bot",
			SchemaName: "bot",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "Premium",
			SchemaName: "premium",
			Null:       !r.Flags.Has(1),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *RequestPeerTypeUser) SetFlags() {
	if !(r.Bot == false) {
		r.Flags.Set(0)
	}
	if !(r.Premium == false) {
		r.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (r *RequestPeerTypeUser) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode requestPeerTypeUser#5f3b8a00 as nil")
	}
	b.PutID(RequestPeerTypeUserTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RequestPeerTypeUser) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode requestPeerTypeUser#5f3b8a00 as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode requestPeerTypeUser#5f3b8a00: field flags: %w", err)
	}
	if r.Flags.Has(0) {
		b.PutBool(r.Bot)
	}
	if r.Flags.Has(1) {
		b.PutBool(r.Premium)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RequestPeerTypeUser) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode requestPeerTypeUser#5f3b8a00 to nil")
	}
	if err := b.ConsumeID(RequestPeerTypeUserTypeID); err != nil {
		return fmt.Errorf("unable to decode requestPeerTypeUser#5f3b8a00: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RequestPeerTypeUser) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode requestPeerTypeUser#5f3b8a00 to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeUser#5f3b8a00: field flags: %w", err)
		}
	}
	if r.Flags.Has(0) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeUser#5f3b8a00: field bot: %w", err)
		}
		r.Bot = value
	}
	if r.Flags.Has(1) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeUser#5f3b8a00: field premium: %w", err)
		}
		r.Premium = value
	}
	return nil
}

// SetBot sets value of Bot conditional field.
func (r *RequestPeerTypeUser) SetBot(value bool) {
	r.Flags.Set(0)
	r.Bot = value
}

// GetBot returns value of Bot conditional field and
// boolean which is true if field was set.
func (r *RequestPeerTypeUser) GetBot() (value bool, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(0) {
		return value, false
	}
	return r.Bot, true
}

// SetPremium sets value of Premium conditional field.
func (r *RequestPeerTypeUser) SetPremium(value bool) {
	r.Flags.Set(1)
	r.Premium = value
}

// GetPremium returns value of Premium conditional field and
// boolean which is true if field was set.
func (r *RequestPeerTypeUser) GetPremium() (value bool, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(1) {
		return value, false
	}
	return r.Premium, true
}

// RequestPeerTypeChat represents TL type `requestPeerTypeChat#c9f06e1b`.
// Choose a chat or supergroup
//
// See https://core.telegram.org/constructor/requestPeerTypeChat for reference.
type RequestPeerTypeChat struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to allow only choosing chats or supergroups that were created by the current
	// user.
	Creator bool
	// Whether to allow only choosing chats or supergroups where the bot is a participant.
	BotParticipant bool
	// If specified, allows only choosing channels with or without a username, according to
	// the value of Bool¹.
	//
	// Links:
	//  1) https://core.telegram.org/type/Bool
	//
	// Use SetHasUsername and GetHasUsername helpers.
	HasUsername bool
	// If specified, allows only choosing chats or supergroups that are or aren't forums¹,
	// according to the value of Bool².
	//
	// Links:
	//  1) https://core.telegram.org/api/forum
	//  2) https://core.telegram.org/type/Bool
	//
	// Use SetForum and GetForum helpers.
	Forum bool
	// If specified, allows only choosing chats or supergroups where the current user is an
	// admin with at least the specified admin rights.
	//
	// Use SetUserAdminRights and GetUserAdminRights helpers.
	UserAdminRights ChatAdminRights
	// If specified, allows only choosing chats or supergroups where the bot is an admin with
	// at least the specified admin rights.
	//
	// Use SetBotAdminRights and GetBotAdminRights helpers.
	BotAdminRights ChatAdminRights
}

// RequestPeerTypeChatTypeID is TL type id of RequestPeerTypeChat.
const RequestPeerTypeChatTypeID = 0xc9f06e1b

// construct implements constructor of RequestPeerTypeClass.
func (r RequestPeerTypeChat) construct() RequestPeerTypeClass { return &r }

// Ensuring interfaces in compile-time for RequestPeerTypeChat.
var (
	_ bin.Encoder     = &RequestPeerTypeChat{}
	_ bin.Decoder     = &RequestPeerTypeChat{}
	_ bin.BareEncoder = &RequestPeerTypeChat{}
	_ bin.BareDecoder = &RequestPeerTypeChat{}

	_ RequestPeerTypeClass = &RequestPeerTypeChat{}
)

func (r *RequestPeerTypeChat) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Creator == false) {
		return false
	}
	if !(r.BotParticipant == false) {
		return false
	}
	if !(r.HasUsername == false) {
		return false
	}
	if !(r.Forum == false) {
		return false
	}
	if !(r.UserAdminRights.Zero()) {
		return false
	}
	if !(r.BotAdminRights.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RequestPeerTypeChat) String() string {
	if r == nil {
		return "RequestPeerTypeChat(nil)"
	}
	type Alias RequestPeerTypeChat
	return fmt.Sprintf("RequestPeerTypeChat%+v", Alias(*r))
}

// FillFrom fills RequestPeerTypeChat from given interface.
func (r *RequestPeerTypeChat) FillFrom(from interface {
	GetCreator() (value bool)
	GetBotParticipant() (value bool)
	GetHasUsername() (value bool, ok bool)
	GetForum() (value bool, ok bool)
	GetUserAdminRights() (value ChatAdminRights, ok bool)
	GetBotAdminRights() (value ChatAdminRights, ok bool)
}) {
	r.Creator = from.GetCreator()
	r.BotParticipant = from.GetBotParticipant()
	if val, ok := from.GetHasUsername(); ok {
		r.HasUsername = val
	}

	if val, ok := from.GetForum(); ok {
		r.Forum = val
	}

	if val, ok := from.GetUserAdminRights(); ok {
		r.UserAdminRights = val
	}

	if val, ok := from.GetBotAdminRights(); ok {
		r.BotAdminRights = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RequestPeerTypeChat) TypeID() uint32 {
	return RequestPeerTypeChatTypeID
}

// TypeName returns name of type in TL schema.
func (*RequestPeerTypeChat) TypeName() string {
	return "requestPeerTypeChat"
}

// TypeInfo returns info about TL type.
func (r *RequestPeerTypeChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "requestPeerTypeChat",
		ID:   RequestPeerTypeChatTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Creator",
			SchemaName: "creator",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "BotParticipant",
			SchemaName: "bot_participant",
			Null:       !r.Flags.Has(5),
		},
		{
			Name:       "HasUsername",
			SchemaName: "has_username",
			Null:       !r.Flags.Has(3),
		},
		{
			Name:       "Forum",
			SchemaName: "forum",
			Null:       !r.Flags.Has(4),
		},
		{
			Name:       "UserAdminRights",
			SchemaName: "user_admin_rights",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "BotAdminRights",
			SchemaName: "bot_admin_rights",
			Null:       !r.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *RequestPeerTypeChat) SetFlags() {
	if !(r.Creator == false) {
		r.Flags.Set(0)
	}
	if !(r.BotParticipant == false) {
		r.Flags.Set(5)
	}
	if !(r.HasUsername == false) {
		r.Flags.Set(3)
	}
	if !(r.Forum == false) {
		r.Flags.Set(4)
	}
	if !(r.UserAdminRights.Zero()) {
		r.Flags.Set(1)
	}
	if !(r.BotAdminRights.Zero()) {
		r.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (r *RequestPeerTypeChat) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode requestPeerTypeChat#c9f06e1b as nil")
	}
	b.PutID(RequestPeerTypeChatTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RequestPeerTypeChat) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode requestPeerTypeChat#c9f06e1b as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode requestPeerTypeChat#c9f06e1b: field flags: %w", err)
	}
	if r.Flags.Has(3) {
		b.PutBool(r.HasUsername)
	}
	if r.Flags.Has(4) {
		b.PutBool(r.Forum)
	}
	if r.Flags.Has(1) {
		if err := r.UserAdminRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode requestPeerTypeChat#c9f06e1b: field user_admin_rights: %w", err)
		}
	}
	if r.Flags.Has(2) {
		if err := r.BotAdminRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode requestPeerTypeChat#c9f06e1b: field bot_admin_rights: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RequestPeerTypeChat) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode requestPeerTypeChat#c9f06e1b to nil")
	}
	if err := b.ConsumeID(RequestPeerTypeChatTypeID); err != nil {
		return fmt.Errorf("unable to decode requestPeerTypeChat#c9f06e1b: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RequestPeerTypeChat) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode requestPeerTypeChat#c9f06e1b to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeChat#c9f06e1b: field flags: %w", err)
		}
	}
	r.Creator = r.Flags.Has(0)
	r.BotParticipant = r.Flags.Has(5)
	if r.Flags.Has(3) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeChat#c9f06e1b: field has_username: %w", err)
		}
		r.HasUsername = value
	}
	if r.Flags.Has(4) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeChat#c9f06e1b: field forum: %w", err)
		}
		r.Forum = value
	}
	if r.Flags.Has(1) {
		if err := r.UserAdminRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeChat#c9f06e1b: field user_admin_rights: %w", err)
		}
	}
	if r.Flags.Has(2) {
		if err := r.BotAdminRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeChat#c9f06e1b: field bot_admin_rights: %w", err)
		}
	}
	return nil
}

// SetCreator sets value of Creator conditional field.
func (r *RequestPeerTypeChat) SetCreator(value bool) {
	if value {
		r.Flags.Set(0)
		r.Creator = true
	} else {
		r.Flags.Unset(0)
		r.Creator = false
	}
}

// GetCreator returns value of Creator conditional field.
func (r *RequestPeerTypeChat) GetCreator() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(0)
}

// SetBotParticipant sets value of BotParticipant conditional field.
func (r *RequestPeerTypeChat) SetBotParticipant(value bool) {
	if value {
		r.Flags.Set(5)
		r.BotParticipant = true
	} else {
		r.Flags.Unset(5)
		r.BotParticipant = false
	}
}

// GetBotParticipant returns value of BotParticipant conditional field.
func (r *RequestPeerTypeChat) GetBotParticipant() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(5)
}

// SetHasUsername sets value of HasUsername conditional field.
func (r *RequestPeerTypeChat) SetHasUsername(value bool) {
	r.Flags.Set(3)
	r.HasUsername = value
}

// GetHasUsername returns value of HasUsername conditional field and
// boolean which is true if field was set.
func (r *RequestPeerTypeChat) GetHasUsername() (value bool, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(3) {
		return value, false
	}
	return r.HasUsername, true
}

// SetForum sets value of Forum conditional field.
func (r *RequestPeerTypeChat) SetForum(value bool) {
	r.Flags.Set(4)
	r.Forum = value
}

// GetForum returns value of Forum conditional field and
// boolean which is true if field was set.
func (r *RequestPeerTypeChat) GetForum() (value bool, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(4) {
		return value, false
	}
	return r.Forum, true
}

// SetUserAdminRights sets value of UserAdminRights conditional field.
func (r *RequestPeerTypeChat) SetUserAdminRights(value ChatAdminRights) {
	r.Flags.Set(1)
	r.UserAdminRights = value
}

// GetUserAdminRights returns value of UserAdminRights conditional field and
// boolean which is true if field was set.
func (r *RequestPeerTypeChat) GetUserAdminRights() (value ChatAdminRights, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(1) {
		return value, false
	}
	return r.UserAdminRights, true
}

// SetBotAdminRights sets value of BotAdminRights conditional field.
func (r *RequestPeerTypeChat) SetBotAdminRights(value ChatAdminRights) {
	r.Flags.Set(2)
	r.BotAdminRights = value
}

// GetBotAdminRights returns value of BotAdminRights conditional field and
// boolean which is true if field was set.
func (r *RequestPeerTypeChat) GetBotAdminRights() (value ChatAdminRights, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(2) {
		return value, false
	}
	return r.BotAdminRights, true
}

// RequestPeerTypeBroadcast represents TL type `requestPeerTypeBroadcast#339bef6c`.
// Choose a channel
//
// See https://core.telegram.org/constructor/requestPeerTypeBroadcast for reference.
type RequestPeerTypeBroadcast struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to allow only choosing channels that were created by the current user.
	Creator bool
	// If specified, allows only choosing channels with or without a username, according to
	// the value of Bool¹.
	//
	// Links:
	//  1) https://core.telegram.org/type/Bool
	//
	// Use SetHasUsername and GetHasUsername helpers.
	HasUsername bool
	// If specified, allows only choosing channels where the current user is an admin with at
	// least the specified admin rights.
	//
	// Use SetUserAdminRights and GetUserAdminRights helpers.
	UserAdminRights ChatAdminRights
	// If specified, allows only choosing channels where the bot is an admin with at least
	// the specified admin rights.
	//
	// Use SetBotAdminRights and GetBotAdminRights helpers.
	BotAdminRights ChatAdminRights
}

// RequestPeerTypeBroadcastTypeID is TL type id of RequestPeerTypeBroadcast.
const RequestPeerTypeBroadcastTypeID = 0x339bef6c

// construct implements constructor of RequestPeerTypeClass.
func (r RequestPeerTypeBroadcast) construct() RequestPeerTypeClass { return &r }

// Ensuring interfaces in compile-time for RequestPeerTypeBroadcast.
var (
	_ bin.Encoder     = &RequestPeerTypeBroadcast{}
	_ bin.Decoder     = &RequestPeerTypeBroadcast{}
	_ bin.BareEncoder = &RequestPeerTypeBroadcast{}
	_ bin.BareDecoder = &RequestPeerTypeBroadcast{}

	_ RequestPeerTypeClass = &RequestPeerTypeBroadcast{}
)

func (r *RequestPeerTypeBroadcast) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Flags.Zero()) {
		return false
	}
	if !(r.Creator == false) {
		return false
	}
	if !(r.HasUsername == false) {
		return false
	}
	if !(r.UserAdminRights.Zero()) {
		return false
	}
	if !(r.BotAdminRights.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RequestPeerTypeBroadcast) String() string {
	if r == nil {
		return "RequestPeerTypeBroadcast(nil)"
	}
	type Alias RequestPeerTypeBroadcast
	return fmt.Sprintf("RequestPeerTypeBroadcast%+v", Alias(*r))
}

// FillFrom fills RequestPeerTypeBroadcast from given interface.
func (r *RequestPeerTypeBroadcast) FillFrom(from interface {
	GetCreator() (value bool)
	GetHasUsername() (value bool, ok bool)
	GetUserAdminRights() (value ChatAdminRights, ok bool)
	GetBotAdminRights() (value ChatAdminRights, ok bool)
}) {
	r.Creator = from.GetCreator()
	if val, ok := from.GetHasUsername(); ok {
		r.HasUsername = val
	}

	if val, ok := from.GetUserAdminRights(); ok {
		r.UserAdminRights = val
	}

	if val, ok := from.GetBotAdminRights(); ok {
		r.BotAdminRights = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RequestPeerTypeBroadcast) TypeID() uint32 {
	return RequestPeerTypeBroadcastTypeID
}

// TypeName returns name of type in TL schema.
func (*RequestPeerTypeBroadcast) TypeName() string {
	return "requestPeerTypeBroadcast"
}

// TypeInfo returns info about TL type.
func (r *RequestPeerTypeBroadcast) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "requestPeerTypeBroadcast",
		ID:   RequestPeerTypeBroadcastTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Creator",
			SchemaName: "creator",
			Null:       !r.Flags.Has(0),
		},
		{
			Name:       "HasUsername",
			SchemaName: "has_username",
			Null:       !r.Flags.Has(3),
		},
		{
			Name:       "UserAdminRights",
			SchemaName: "user_admin_rights",
			Null:       !r.Flags.Has(1),
		},
		{
			Name:       "BotAdminRights",
			SchemaName: "bot_admin_rights",
			Null:       !r.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (r *RequestPeerTypeBroadcast) SetFlags() {
	if !(r.Creator == false) {
		r.Flags.Set(0)
	}
	if !(r.HasUsername == false) {
		r.Flags.Set(3)
	}
	if !(r.UserAdminRights.Zero()) {
		r.Flags.Set(1)
	}
	if !(r.BotAdminRights.Zero()) {
		r.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (r *RequestPeerTypeBroadcast) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode requestPeerTypeBroadcast#339bef6c as nil")
	}
	b.PutID(RequestPeerTypeBroadcastTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RequestPeerTypeBroadcast) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode requestPeerTypeBroadcast#339bef6c as nil")
	}
	r.SetFlags()
	if err := r.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode requestPeerTypeBroadcast#339bef6c: field flags: %w", err)
	}
	if r.Flags.Has(3) {
		b.PutBool(r.HasUsername)
	}
	if r.Flags.Has(1) {
		if err := r.UserAdminRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode requestPeerTypeBroadcast#339bef6c: field user_admin_rights: %w", err)
		}
	}
	if r.Flags.Has(2) {
		if err := r.BotAdminRights.Encode(b); err != nil {
			return fmt.Errorf("unable to encode requestPeerTypeBroadcast#339bef6c: field bot_admin_rights: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *RequestPeerTypeBroadcast) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode requestPeerTypeBroadcast#339bef6c to nil")
	}
	if err := b.ConsumeID(RequestPeerTypeBroadcastTypeID); err != nil {
		return fmt.Errorf("unable to decode requestPeerTypeBroadcast#339bef6c: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RequestPeerTypeBroadcast) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode requestPeerTypeBroadcast#339bef6c to nil")
	}
	{
		if err := r.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeBroadcast#339bef6c: field flags: %w", err)
		}
	}
	r.Creator = r.Flags.Has(0)
	if r.Flags.Has(3) {
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeBroadcast#339bef6c: field has_username: %w", err)
		}
		r.HasUsername = value
	}
	if r.Flags.Has(1) {
		if err := r.UserAdminRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeBroadcast#339bef6c: field user_admin_rights: %w", err)
		}
	}
	if r.Flags.Has(2) {
		if err := r.BotAdminRights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode requestPeerTypeBroadcast#339bef6c: field bot_admin_rights: %w", err)
		}
	}
	return nil
}

// SetCreator sets value of Creator conditional field.
func (r *RequestPeerTypeBroadcast) SetCreator(value bool) {
	if value {
		r.Flags.Set(0)
		r.Creator = true
	} else {
		r.Flags.Unset(0)
		r.Creator = false
	}
}

// GetCreator returns value of Creator conditional field.
func (r *RequestPeerTypeBroadcast) GetCreator() (value bool) {
	if r == nil {
		return
	}
	return r.Flags.Has(0)
}

// SetHasUsername sets value of HasUsername conditional field.
func (r *RequestPeerTypeBroadcast) SetHasUsername(value bool) {
	r.Flags.Set(3)
	r.HasUsername = value
}

// GetHasUsername returns value of HasUsername conditional field and
// boolean which is true if field was set.
func (r *RequestPeerTypeBroadcast) GetHasUsername() (value bool, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(3) {
		return value, false
	}
	return r.HasUsername, true
}

// SetUserAdminRights sets value of UserAdminRights conditional field.
func (r *RequestPeerTypeBroadcast) SetUserAdminRights(value ChatAdminRights) {
	r.Flags.Set(1)
	r.UserAdminRights = value
}

// GetUserAdminRights returns value of UserAdminRights conditional field and
// boolean which is true if field was set.
func (r *RequestPeerTypeBroadcast) GetUserAdminRights() (value ChatAdminRights, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(1) {
		return value, false
	}
	return r.UserAdminRights, true
}

// SetBotAdminRights sets value of BotAdminRights conditional field.
func (r *RequestPeerTypeBroadcast) SetBotAdminRights(value ChatAdminRights) {
	r.Flags.Set(2)
	r.BotAdminRights = value
}

// GetBotAdminRights returns value of BotAdminRights conditional field and
// boolean which is true if field was set.
func (r *RequestPeerTypeBroadcast) GetBotAdminRights() (value ChatAdminRights, ok bool) {
	if r == nil {
		return
	}
	if !r.Flags.Has(2) {
		return value, false
	}
	return r.BotAdminRights, true
}

// RequestPeerTypeClassName is schema name of RequestPeerTypeClass.
const RequestPeerTypeClassName = "RequestPeerType"

// RequestPeerTypeClass represents RequestPeerType generic type.
//
// See https://core.telegram.org/type/RequestPeerType for reference.
//
// Example:
//
//	g, err := tg.DecodeRequestPeerType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.RequestPeerTypeUser: // requestPeerTypeUser#5f3b8a00
//	case *tg.RequestPeerTypeChat: // requestPeerTypeChat#c9f06e1b
//	case *tg.RequestPeerTypeBroadcast: // requestPeerTypeBroadcast#339bef6c
//	default: panic(v)
//	}
type RequestPeerTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() RequestPeerTypeClass

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

// DecodeRequestPeerType implements binary de-serialization for RequestPeerTypeClass.
func DecodeRequestPeerType(buf *bin.Buffer) (RequestPeerTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case RequestPeerTypeUserTypeID:
		// Decoding requestPeerTypeUser#5f3b8a00.
		v := RequestPeerTypeUser{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RequestPeerTypeClass: %w", err)
		}
		return &v, nil
	case RequestPeerTypeChatTypeID:
		// Decoding requestPeerTypeChat#c9f06e1b.
		v := RequestPeerTypeChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RequestPeerTypeClass: %w", err)
		}
		return &v, nil
	case RequestPeerTypeBroadcastTypeID:
		// Decoding requestPeerTypeBroadcast#339bef6c.
		v := RequestPeerTypeBroadcast{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode RequestPeerTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode RequestPeerTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// RequestPeerType boxes the RequestPeerTypeClass providing a helper.
type RequestPeerTypeBox struct {
	RequestPeerType RequestPeerTypeClass
}

// Decode implements bin.Decoder for RequestPeerTypeBox.
func (b *RequestPeerTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode RequestPeerTypeBox to nil")
	}
	v, err := DecodeRequestPeerType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.RequestPeerType = v
	return nil
}

// Encode implements bin.Encode for RequestPeerTypeBox.
func (b *RequestPeerTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.RequestPeerType == nil {
		return fmt.Errorf("unable to encode RequestPeerTypeClass as nil")
	}
	return b.RequestPeerType.Encode(buf)
}
