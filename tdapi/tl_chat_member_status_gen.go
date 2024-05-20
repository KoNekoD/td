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

// ChatMemberStatusCreator represents TL type `chatMemberStatusCreator#f6764afe`.
type ChatMemberStatusCreator struct {
	// A custom title of the owner; 0-16 characters without emojis; applicable to supergroups
	// only
	CustomTitle string
	// True, if the creator isn't shown in the chat member list and sends messages
	// anonymously; applicable to supergroups only
	IsAnonymous bool
	// True, if the user is a member of the chat
	IsMember bool
}

// ChatMemberStatusCreatorTypeID is TL type id of ChatMemberStatusCreator.
const ChatMemberStatusCreatorTypeID = 0xf6764afe

// construct implements constructor of ChatMemberStatusClass.
func (c ChatMemberStatusCreator) construct() ChatMemberStatusClass { return &c }

// Ensuring interfaces in compile-time for ChatMemberStatusCreator.
var (
	_ bin.Encoder     = &ChatMemberStatusCreator{}
	_ bin.Decoder     = &ChatMemberStatusCreator{}
	_ bin.BareEncoder = &ChatMemberStatusCreator{}
	_ bin.BareDecoder = &ChatMemberStatusCreator{}

	_ ChatMemberStatusClass = &ChatMemberStatusCreator{}
)

func (c *ChatMemberStatusCreator) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.CustomTitle == "") {
		return false
	}
	if !(c.IsAnonymous == false) {
		return false
	}
	if !(c.IsMember == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMemberStatusCreator) String() string {
	if c == nil {
		return "ChatMemberStatusCreator(nil)"
	}
	type Alias ChatMemberStatusCreator
	return fmt.Sprintf("ChatMemberStatusCreator%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMemberStatusCreator) TypeID() uint32 {
	return ChatMemberStatusCreatorTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMemberStatusCreator) TypeName() string {
	return "chatMemberStatusCreator"
}

// TypeInfo returns info about TL type.
func (c *ChatMemberStatusCreator) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMemberStatusCreator",
		ID:   ChatMemberStatusCreatorTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CustomTitle",
			SchemaName: "custom_title",
		},
		{
			Name:       "IsAnonymous",
			SchemaName: "is_anonymous",
		},
		{
			Name:       "IsMember",
			SchemaName: "is_member",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMemberStatusCreator) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusCreator#f6764afe as nil")
	}
	b.PutID(ChatMemberStatusCreatorTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMemberStatusCreator) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusCreator#f6764afe as nil")
	}
	b.PutString(c.CustomTitle)
	b.PutBool(c.IsAnonymous)
	b.PutBool(c.IsMember)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMemberStatusCreator) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusCreator#f6764afe to nil")
	}
	if err := b.ConsumeID(ChatMemberStatusCreatorTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMemberStatusCreator#f6764afe: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMemberStatusCreator) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusCreator#f6764afe to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatMemberStatusCreator#f6764afe: field custom_title: %w", err)
		}
		c.CustomTitle = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatMemberStatusCreator#f6764afe: field is_anonymous: %w", err)
		}
		c.IsAnonymous = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatMemberStatusCreator#f6764afe: field is_member: %w", err)
		}
		c.IsMember = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatMemberStatusCreator) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusCreator#f6764afe as nil")
	}
	b.ObjStart()
	b.PutID("chatMemberStatusCreator")
	b.Comma()
	b.FieldStart("custom_title")
	b.PutString(c.CustomTitle)
	b.Comma()
	b.FieldStart("is_anonymous")
	b.PutBool(c.IsAnonymous)
	b.Comma()
	b.FieldStart("is_member")
	b.PutBool(c.IsMember)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatMemberStatusCreator) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusCreator#f6764afe to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatMemberStatusCreator"); err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusCreator#f6764afe: %w", err)
			}
		case "custom_title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusCreator#f6764afe: field custom_title: %w", err)
			}
			c.CustomTitle = value
		case "is_anonymous":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusCreator#f6764afe: field is_anonymous: %w", err)
			}
			c.IsAnonymous = value
		case "is_member":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusCreator#f6764afe: field is_member: %w", err)
			}
			c.IsMember = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCustomTitle returns value of CustomTitle field.
func (c *ChatMemberStatusCreator) GetCustomTitle() (value string) {
	if c == nil {
		return
	}
	return c.CustomTitle
}

// GetIsAnonymous returns value of IsAnonymous field.
func (c *ChatMemberStatusCreator) GetIsAnonymous() (value bool) {
	if c == nil {
		return
	}
	return c.IsAnonymous
}

// GetIsMember returns value of IsMember field.
func (c *ChatMemberStatusCreator) GetIsMember() (value bool) {
	if c == nil {
		return
	}
	return c.IsMember
}

// ChatMemberStatusAdministrator represents TL type `chatMemberStatusAdministrator#fbd3841d`.
type ChatMemberStatusAdministrator struct {
	// A custom title of the administrator; 0-16 characters without emojis; applicable to
	// supergroups only
	CustomTitle string
	// True, if the current user can edit the administrator privileges for the called user
	CanBeEdited bool
	// Rights of the administrator
	Rights ChatAdministratorRights
}

// ChatMemberStatusAdministratorTypeID is TL type id of ChatMemberStatusAdministrator.
const ChatMemberStatusAdministratorTypeID = 0xfbd3841d

// construct implements constructor of ChatMemberStatusClass.
func (c ChatMemberStatusAdministrator) construct() ChatMemberStatusClass { return &c }

// Ensuring interfaces in compile-time for ChatMemberStatusAdministrator.
var (
	_ bin.Encoder     = &ChatMemberStatusAdministrator{}
	_ bin.Decoder     = &ChatMemberStatusAdministrator{}
	_ bin.BareEncoder = &ChatMemberStatusAdministrator{}
	_ bin.BareDecoder = &ChatMemberStatusAdministrator{}

	_ ChatMemberStatusClass = &ChatMemberStatusAdministrator{}
)

func (c *ChatMemberStatusAdministrator) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.CustomTitle == "") {
		return false
	}
	if !(c.CanBeEdited == false) {
		return false
	}
	if !(c.Rights.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMemberStatusAdministrator) String() string {
	if c == nil {
		return "ChatMemberStatusAdministrator(nil)"
	}
	type Alias ChatMemberStatusAdministrator
	return fmt.Sprintf("ChatMemberStatusAdministrator%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMemberStatusAdministrator) TypeID() uint32 {
	return ChatMemberStatusAdministratorTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMemberStatusAdministrator) TypeName() string {
	return "chatMemberStatusAdministrator"
}

// TypeInfo returns info about TL type.
func (c *ChatMemberStatusAdministrator) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMemberStatusAdministrator",
		ID:   ChatMemberStatusAdministratorTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CustomTitle",
			SchemaName: "custom_title",
		},
		{
			Name:       "CanBeEdited",
			SchemaName: "can_be_edited",
		},
		{
			Name:       "Rights",
			SchemaName: "rights",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMemberStatusAdministrator) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusAdministrator#fbd3841d as nil")
	}
	b.PutID(ChatMemberStatusAdministratorTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMemberStatusAdministrator) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusAdministrator#fbd3841d as nil")
	}
	b.PutString(c.CustomTitle)
	b.PutBool(c.CanBeEdited)
	if err := c.Rights.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatMemberStatusAdministrator#fbd3841d: field rights: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMemberStatusAdministrator) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusAdministrator#fbd3841d to nil")
	}
	if err := b.ConsumeID(ChatMemberStatusAdministratorTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMemberStatusAdministrator#fbd3841d: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMemberStatusAdministrator) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusAdministrator#fbd3841d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode chatMemberStatusAdministrator#fbd3841d: field custom_title: %w", err)
		}
		c.CustomTitle = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatMemberStatusAdministrator#fbd3841d: field can_be_edited: %w", err)
		}
		c.CanBeEdited = value
	}
	{
		if err := c.Rights.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatMemberStatusAdministrator#fbd3841d: field rights: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatMemberStatusAdministrator) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusAdministrator#fbd3841d as nil")
	}
	b.ObjStart()
	b.PutID("chatMemberStatusAdministrator")
	b.Comma()
	b.FieldStart("custom_title")
	b.PutString(c.CustomTitle)
	b.Comma()
	b.FieldStart("can_be_edited")
	b.PutBool(c.CanBeEdited)
	b.Comma()
	b.FieldStart("rights")
	if err := c.Rights.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatMemberStatusAdministrator#fbd3841d: field rights: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatMemberStatusAdministrator) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusAdministrator#fbd3841d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatMemberStatusAdministrator"); err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusAdministrator#fbd3841d: %w", err)
			}
		case "custom_title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusAdministrator#fbd3841d: field custom_title: %w", err)
			}
			c.CustomTitle = value
		case "can_be_edited":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusAdministrator#fbd3841d: field can_be_edited: %w", err)
			}
			c.CanBeEdited = value
		case "rights":
			if err := c.Rights.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusAdministrator#fbd3841d: field rights: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetCustomTitle returns value of CustomTitle field.
func (c *ChatMemberStatusAdministrator) GetCustomTitle() (value string) {
	if c == nil {
		return
	}
	return c.CustomTitle
}

// GetCanBeEdited returns value of CanBeEdited field.
func (c *ChatMemberStatusAdministrator) GetCanBeEdited() (value bool) {
	if c == nil {
		return
	}
	return c.CanBeEdited
}

// GetRights returns value of Rights field.
func (c *ChatMemberStatusAdministrator) GetRights() (value ChatAdministratorRights) {
	if c == nil {
		return
	}
	return c.Rights
}

// ChatMemberStatusMember represents TL type `chatMemberStatusMember#32597455`.
type ChatMemberStatusMember struct {
}

// ChatMemberStatusMemberTypeID is TL type id of ChatMemberStatusMember.
const ChatMemberStatusMemberTypeID = 0x32597455

// construct implements constructor of ChatMemberStatusClass.
func (c ChatMemberStatusMember) construct() ChatMemberStatusClass { return &c }

// Ensuring interfaces in compile-time for ChatMemberStatusMember.
var (
	_ bin.Encoder     = &ChatMemberStatusMember{}
	_ bin.Decoder     = &ChatMemberStatusMember{}
	_ bin.BareEncoder = &ChatMemberStatusMember{}
	_ bin.BareDecoder = &ChatMemberStatusMember{}

	_ ChatMemberStatusClass = &ChatMemberStatusMember{}
)

func (c *ChatMemberStatusMember) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMemberStatusMember) String() string {
	if c == nil {
		return "ChatMemberStatusMember(nil)"
	}
	type Alias ChatMemberStatusMember
	return fmt.Sprintf("ChatMemberStatusMember%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMemberStatusMember) TypeID() uint32 {
	return ChatMemberStatusMemberTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMemberStatusMember) TypeName() string {
	return "chatMemberStatusMember"
}

// TypeInfo returns info about TL type.
func (c *ChatMemberStatusMember) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMemberStatusMember",
		ID:   ChatMemberStatusMemberTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMemberStatusMember) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusMember#32597455 as nil")
	}
	b.PutID(ChatMemberStatusMemberTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMemberStatusMember) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusMember#32597455 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMemberStatusMember) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusMember#32597455 to nil")
	}
	if err := b.ConsumeID(ChatMemberStatusMemberTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMemberStatusMember#32597455: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMemberStatusMember) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusMember#32597455 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatMemberStatusMember) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusMember#32597455 as nil")
	}
	b.ObjStart()
	b.PutID("chatMemberStatusMember")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatMemberStatusMember) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusMember#32597455 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatMemberStatusMember"); err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusMember#32597455: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ChatMemberStatusRestricted represents TL type `chatMemberStatusRestricted#630774a6`.
type ChatMemberStatusRestricted struct {
	// True, if the user is a member of the chat
	IsMember bool
	// Point in time (Unix timestamp) when restrictions will be lifted from the user; 0 if
	// never. If the user is restricted for more than 366 days or for less than 30 seconds
	// from the current time, the user is considered to be restricted forever
	RestrictedUntilDate int32
	// User permissions in the chat
	Permissions ChatPermissions
}

// ChatMemberStatusRestrictedTypeID is TL type id of ChatMemberStatusRestricted.
const ChatMemberStatusRestrictedTypeID = 0x630774a6

// construct implements constructor of ChatMemberStatusClass.
func (c ChatMemberStatusRestricted) construct() ChatMemberStatusClass { return &c }

// Ensuring interfaces in compile-time for ChatMemberStatusRestricted.
var (
	_ bin.Encoder     = &ChatMemberStatusRestricted{}
	_ bin.Decoder     = &ChatMemberStatusRestricted{}
	_ bin.BareEncoder = &ChatMemberStatusRestricted{}
	_ bin.BareDecoder = &ChatMemberStatusRestricted{}

	_ ChatMemberStatusClass = &ChatMemberStatusRestricted{}
)

func (c *ChatMemberStatusRestricted) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.IsMember == false) {
		return false
	}
	if !(c.RestrictedUntilDate == 0) {
		return false
	}
	if !(c.Permissions.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMemberStatusRestricted) String() string {
	if c == nil {
		return "ChatMemberStatusRestricted(nil)"
	}
	type Alias ChatMemberStatusRestricted
	return fmt.Sprintf("ChatMemberStatusRestricted%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMemberStatusRestricted) TypeID() uint32 {
	return ChatMemberStatusRestrictedTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMemberStatusRestricted) TypeName() string {
	return "chatMemberStatusRestricted"
}

// TypeInfo returns info about TL type.
func (c *ChatMemberStatusRestricted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMemberStatusRestricted",
		ID:   ChatMemberStatusRestrictedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "IsMember",
			SchemaName: "is_member",
		},
		{
			Name:       "RestrictedUntilDate",
			SchemaName: "restricted_until_date",
		},
		{
			Name:       "Permissions",
			SchemaName: "permissions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMemberStatusRestricted) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusRestricted#630774a6 as nil")
	}
	b.PutID(ChatMemberStatusRestrictedTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMemberStatusRestricted) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusRestricted#630774a6 as nil")
	}
	b.PutBool(c.IsMember)
	b.PutInt32(c.RestrictedUntilDate)
	if err := c.Permissions.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatMemberStatusRestricted#630774a6: field permissions: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMemberStatusRestricted) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusRestricted#630774a6 to nil")
	}
	if err := b.ConsumeID(ChatMemberStatusRestrictedTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMemberStatusRestricted#630774a6: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMemberStatusRestricted) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusRestricted#630774a6 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatMemberStatusRestricted#630774a6: field is_member: %w", err)
		}
		c.IsMember = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatMemberStatusRestricted#630774a6: field restricted_until_date: %w", err)
		}
		c.RestrictedUntilDate = value
	}
	{
		if err := c.Permissions.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatMemberStatusRestricted#630774a6: field permissions: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatMemberStatusRestricted) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusRestricted#630774a6 as nil")
	}
	b.ObjStart()
	b.PutID("chatMemberStatusRestricted")
	b.Comma()
	b.FieldStart("is_member")
	b.PutBool(c.IsMember)
	b.Comma()
	b.FieldStart("restricted_until_date")
	b.PutInt32(c.RestrictedUntilDate)
	b.Comma()
	b.FieldStart("permissions")
	if err := c.Permissions.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode chatMemberStatusRestricted#630774a6: field permissions: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatMemberStatusRestricted) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusRestricted#630774a6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatMemberStatusRestricted"); err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusRestricted#630774a6: %w", err)
			}
		case "is_member":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusRestricted#630774a6: field is_member: %w", err)
			}
			c.IsMember = value
		case "restricted_until_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusRestricted#630774a6: field restricted_until_date: %w", err)
			}
			c.RestrictedUntilDate = value
		case "permissions":
			if err := c.Permissions.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusRestricted#630774a6: field permissions: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetIsMember returns value of IsMember field.
func (c *ChatMemberStatusRestricted) GetIsMember() (value bool) {
	if c == nil {
		return
	}
	return c.IsMember
}

// GetRestrictedUntilDate returns value of RestrictedUntilDate field.
func (c *ChatMemberStatusRestricted) GetRestrictedUntilDate() (value int32) {
	if c == nil {
		return
	}
	return c.RestrictedUntilDate
}

// GetPermissions returns value of Permissions field.
func (c *ChatMemberStatusRestricted) GetPermissions() (value ChatPermissions) {
	if c == nil {
		return
	}
	return c.Permissions
}

// ChatMemberStatusLeft represents TL type `chatMemberStatusLeft#ffa74425`.
type ChatMemberStatusLeft struct {
}

// ChatMemberStatusLeftTypeID is TL type id of ChatMemberStatusLeft.
const ChatMemberStatusLeftTypeID = 0xffa74425

// construct implements constructor of ChatMemberStatusClass.
func (c ChatMemberStatusLeft) construct() ChatMemberStatusClass { return &c }

// Ensuring interfaces in compile-time for ChatMemberStatusLeft.
var (
	_ bin.Encoder     = &ChatMemberStatusLeft{}
	_ bin.Decoder     = &ChatMemberStatusLeft{}
	_ bin.BareEncoder = &ChatMemberStatusLeft{}
	_ bin.BareDecoder = &ChatMemberStatusLeft{}

	_ ChatMemberStatusClass = &ChatMemberStatusLeft{}
)

func (c *ChatMemberStatusLeft) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMemberStatusLeft) String() string {
	if c == nil {
		return "ChatMemberStatusLeft(nil)"
	}
	type Alias ChatMemberStatusLeft
	return fmt.Sprintf("ChatMemberStatusLeft%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMemberStatusLeft) TypeID() uint32 {
	return ChatMemberStatusLeftTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMemberStatusLeft) TypeName() string {
	return "chatMemberStatusLeft"
}

// TypeInfo returns info about TL type.
func (c *ChatMemberStatusLeft) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMemberStatusLeft",
		ID:   ChatMemberStatusLeftTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMemberStatusLeft) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusLeft#ffa74425 as nil")
	}
	b.PutID(ChatMemberStatusLeftTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMemberStatusLeft) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusLeft#ffa74425 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMemberStatusLeft) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusLeft#ffa74425 to nil")
	}
	if err := b.ConsumeID(ChatMemberStatusLeftTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMemberStatusLeft#ffa74425: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMemberStatusLeft) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusLeft#ffa74425 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatMemberStatusLeft) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusLeft#ffa74425 as nil")
	}
	b.ObjStart()
	b.PutID("chatMemberStatusLeft")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatMemberStatusLeft) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusLeft#ffa74425 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatMemberStatusLeft"); err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusLeft#ffa74425: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ChatMemberStatusBanned represents TL type `chatMemberStatusBanned#9d714eb6`.
type ChatMemberStatusBanned struct {
	// Point in time (Unix timestamp) when the user will be unbanned; 0 if never. If the user
	// is banned for more than 366 days or for less than 30 seconds from the current time,
	// the user is considered to be banned forever. Always 0 in basic groups
	BannedUntilDate int32
}

// ChatMemberStatusBannedTypeID is TL type id of ChatMemberStatusBanned.
const ChatMemberStatusBannedTypeID = 0x9d714eb6

// construct implements constructor of ChatMemberStatusClass.
func (c ChatMemberStatusBanned) construct() ChatMemberStatusClass { return &c }

// Ensuring interfaces in compile-time for ChatMemberStatusBanned.
var (
	_ bin.Encoder     = &ChatMemberStatusBanned{}
	_ bin.Decoder     = &ChatMemberStatusBanned{}
	_ bin.BareEncoder = &ChatMemberStatusBanned{}
	_ bin.BareDecoder = &ChatMemberStatusBanned{}

	_ ChatMemberStatusClass = &ChatMemberStatusBanned{}
)

func (c *ChatMemberStatusBanned) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.BannedUntilDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatMemberStatusBanned) String() string {
	if c == nil {
		return "ChatMemberStatusBanned(nil)"
	}
	type Alias ChatMemberStatusBanned
	return fmt.Sprintf("ChatMemberStatusBanned%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatMemberStatusBanned) TypeID() uint32 {
	return ChatMemberStatusBannedTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatMemberStatusBanned) TypeName() string {
	return "chatMemberStatusBanned"
}

// TypeInfo returns info about TL type.
func (c *ChatMemberStatusBanned) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatMemberStatusBanned",
		ID:   ChatMemberStatusBannedTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BannedUntilDate",
			SchemaName: "banned_until_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatMemberStatusBanned) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusBanned#9d714eb6 as nil")
	}
	b.PutID(ChatMemberStatusBannedTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatMemberStatusBanned) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusBanned#9d714eb6 as nil")
	}
	b.PutInt32(c.BannedUntilDate)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatMemberStatusBanned) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusBanned#9d714eb6 to nil")
	}
	if err := b.ConsumeID(ChatMemberStatusBannedTypeID); err != nil {
		return fmt.Errorf("unable to decode chatMemberStatusBanned#9d714eb6: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatMemberStatusBanned) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusBanned#9d714eb6 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatMemberStatusBanned#9d714eb6: field banned_until_date: %w", err)
		}
		c.BannedUntilDate = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatMemberStatusBanned) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatMemberStatusBanned#9d714eb6 as nil")
	}
	b.ObjStart()
	b.PutID("chatMemberStatusBanned")
	b.Comma()
	b.FieldStart("banned_until_date")
	b.PutInt32(c.BannedUntilDate)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatMemberStatusBanned) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatMemberStatusBanned#9d714eb6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatMemberStatusBanned"); err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusBanned#9d714eb6: %w", err)
			}
		case "banned_until_date":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatMemberStatusBanned#9d714eb6: field banned_until_date: %w", err)
			}
			c.BannedUntilDate = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBannedUntilDate returns value of BannedUntilDate field.
func (c *ChatMemberStatusBanned) GetBannedUntilDate() (value int32) {
	if c == nil {
		return
	}
	return c.BannedUntilDate
}

// ChatMemberStatusClassName is schema name of ChatMemberStatusClass.
const ChatMemberStatusClassName = "ChatMemberStatus"

// ChatMemberStatusClass represents ChatMemberStatus generic type.
//
// Example:
//
//	g, err := tdapi.DecodeChatMemberStatus(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.ChatMemberStatusCreator: // chatMemberStatusCreator#f6764afe
//	case *tdapi.ChatMemberStatusAdministrator: // chatMemberStatusAdministrator#fbd3841d
//	case *tdapi.ChatMemberStatusMember: // chatMemberStatusMember#32597455
//	case *tdapi.ChatMemberStatusRestricted: // chatMemberStatusRestricted#630774a6
//	case *tdapi.ChatMemberStatusLeft: // chatMemberStatusLeft#ffa74425
//	case *tdapi.ChatMemberStatusBanned: // chatMemberStatusBanned#9d714eb6
//	default: panic(v)
//	}
type ChatMemberStatusClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatMemberStatusClass

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

// DecodeChatMemberStatus implements binary de-serialization for ChatMemberStatusClass.
func DecodeChatMemberStatus(buf *bin.Buffer) (ChatMemberStatusClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatMemberStatusCreatorTypeID:
		// Decoding chatMemberStatusCreator#f6764afe.
		v := ChatMemberStatusCreator{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	case ChatMemberStatusAdministratorTypeID:
		// Decoding chatMemberStatusAdministrator#fbd3841d.
		v := ChatMemberStatusAdministrator{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	case ChatMemberStatusMemberTypeID:
		// Decoding chatMemberStatusMember#32597455.
		v := ChatMemberStatusMember{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	case ChatMemberStatusRestrictedTypeID:
		// Decoding chatMemberStatusRestricted#630774a6.
		v := ChatMemberStatusRestricted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	case ChatMemberStatusLeftTypeID:
		// Decoding chatMemberStatusLeft#ffa74425.
		v := ChatMemberStatusLeft{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	case ChatMemberStatusBannedTypeID:
		// Decoding chatMemberStatusBanned#9d714eb6.
		v := ChatMemberStatusBanned{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONChatMemberStatus implements binary de-serialization for ChatMemberStatusClass.
func DecodeTDLibJSONChatMemberStatus(buf tdjson.Decoder) (ChatMemberStatusClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "chatMemberStatusCreator":
		// Decoding chatMemberStatusCreator#f6764afe.
		v := ChatMemberStatusCreator{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	case "chatMemberStatusAdministrator":
		// Decoding chatMemberStatusAdministrator#fbd3841d.
		v := ChatMemberStatusAdministrator{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	case "chatMemberStatusMember":
		// Decoding chatMemberStatusMember#32597455.
		v := ChatMemberStatusMember{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	case "chatMemberStatusRestricted":
		// Decoding chatMemberStatusRestricted#630774a6.
		v := ChatMemberStatusRestricted{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	case "chatMemberStatusLeft":
		// Decoding chatMemberStatusLeft#ffa74425.
		v := ChatMemberStatusLeft{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	case "chatMemberStatusBanned":
		// Decoding chatMemberStatusBanned#9d714eb6.
		v := ChatMemberStatusBanned{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatMemberStatusClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// ChatMemberStatus boxes the ChatMemberStatusClass providing a helper.
type ChatMemberStatusBox struct {
	ChatMemberStatus ChatMemberStatusClass
}

// Decode implements bin.Decoder for ChatMemberStatusBox.
func (b *ChatMemberStatusBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatMemberStatusBox to nil")
	}
	v, err := DecodeChatMemberStatus(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatMemberStatus = v
	return nil
}

// Encode implements bin.Encode for ChatMemberStatusBox.
func (b *ChatMemberStatusBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatMemberStatus == nil {
		return fmt.Errorf("unable to encode ChatMemberStatusClass as nil")
	}
	return b.ChatMemberStatus.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for ChatMemberStatusBox.
func (b *ChatMemberStatusBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatMemberStatusBox to nil")
	}
	v, err := DecodeTDLibJSONChatMemberStatus(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatMemberStatus = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for ChatMemberStatusBox.
func (b *ChatMemberStatusBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.ChatMemberStatus == nil {
		return fmt.Errorf("unable to encode ChatMemberStatusClass as nil")
	}
	return b.ChatMemberStatus.EncodeTDLibJSON(buf)
}
