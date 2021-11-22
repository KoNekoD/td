// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// UserFullInfo represents TL type `userFullInfo#26be6d7c`.
type UserFullInfo struct {
	// User profile photo; may be null
	Photo ChatPhoto
	// True, if the user is blocked by the current user
	IsBlocked bool
	// True, if the user can be called
	CanBeCalled bool
	// True, if a video call can be created with the user
	SupportsVideoCalls bool
	// True, if the user can't be called due to their privacy settings
	HasPrivateCalls bool
	// True, if the current user needs to explicitly allow to share their phone number with
	// the user when the method addContact is used
	NeedPhoneNumberPrivacyException bool
	// A short user bio
	Bio string
	// For bots, the text that is shown on the bot's profile page and is sent together with
	// the link when users share the bot
	ShareText string
	// Contains full information about a user
	Description string
	// Number of group chats where both the other user and the current user are a member; 0
	// for the current user
	GroupInCommonCount int32
	// For bots, list of the bot commands
	Commands []BotCommand
}

// UserFullInfoTypeID is TL type id of UserFullInfo.
const UserFullInfoTypeID = 0x26be6d7c

// Ensuring interfaces in compile-time for UserFullInfo.
var (
	_ bin.Encoder     = &UserFullInfo{}
	_ bin.Decoder     = &UserFullInfo{}
	_ bin.BareEncoder = &UserFullInfo{}
	_ bin.BareDecoder = &UserFullInfo{}
)

func (u *UserFullInfo) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Photo.Zero()) {
		return false
	}
	if !(u.IsBlocked == false) {
		return false
	}
	if !(u.CanBeCalled == false) {
		return false
	}
	if !(u.SupportsVideoCalls == false) {
		return false
	}
	if !(u.HasPrivateCalls == false) {
		return false
	}
	if !(u.NeedPhoneNumberPrivacyException == false) {
		return false
	}
	if !(u.Bio == "") {
		return false
	}
	if !(u.ShareText == "") {
		return false
	}
	if !(u.Description == "") {
		return false
	}
	if !(u.GroupInCommonCount == 0) {
		return false
	}
	if !(u.Commands == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserFullInfo) String() string {
	if u == nil {
		return "UserFullInfo(nil)"
	}
	type Alias UserFullInfo
	return fmt.Sprintf("UserFullInfo%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserFullInfo) TypeID() uint32 {
	return UserFullInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*UserFullInfo) TypeName() string {
	return "userFullInfo"
}

// TypeInfo returns info about TL type.
func (u *UserFullInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userFullInfo",
		ID:   UserFullInfoTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "IsBlocked",
			SchemaName: "is_blocked",
		},
		{
			Name:       "CanBeCalled",
			SchemaName: "can_be_called",
		},
		{
			Name:       "SupportsVideoCalls",
			SchemaName: "supports_video_calls",
		},
		{
			Name:       "HasPrivateCalls",
			SchemaName: "has_private_calls",
		},
		{
			Name:       "NeedPhoneNumberPrivacyException",
			SchemaName: "need_phone_number_privacy_exception",
		},
		{
			Name:       "Bio",
			SchemaName: "bio",
		},
		{
			Name:       "ShareText",
			SchemaName: "share_text",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
		{
			Name:       "GroupInCommonCount",
			SchemaName: "group_in_common_count",
		},
		{
			Name:       "Commands",
			SchemaName: "commands",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserFullInfo) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userFullInfo#26be6d7c as nil")
	}
	b.PutID(UserFullInfoTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserFullInfo) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userFullInfo#26be6d7c as nil")
	}
	if err := u.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#26be6d7c: field photo: %w", err)
	}
	b.PutBool(u.IsBlocked)
	b.PutBool(u.CanBeCalled)
	b.PutBool(u.SupportsVideoCalls)
	b.PutBool(u.HasPrivateCalls)
	b.PutBool(u.NeedPhoneNumberPrivacyException)
	b.PutString(u.Bio)
	b.PutString(u.ShareText)
	b.PutString(u.Description)
	b.PutInt32(u.GroupInCommonCount)
	b.PutInt(len(u.Commands))
	for idx, v := range u.Commands {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare userFullInfo#26be6d7c: field commands element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UserFullInfo) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userFullInfo#26be6d7c to nil")
	}
	if err := b.ConsumeID(UserFullInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode userFullInfo#26be6d7c: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserFullInfo) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userFullInfo#26be6d7c to nil")
	}
	{
		if err := u.Photo.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field photo: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field is_blocked: %w", err)
		}
		u.IsBlocked = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field can_be_called: %w", err)
		}
		u.CanBeCalled = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field supports_video_calls: %w", err)
		}
		u.SupportsVideoCalls = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field has_private_calls: %w", err)
		}
		u.HasPrivateCalls = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field need_phone_number_privacy_exception: %w", err)
		}
		u.NeedPhoneNumberPrivacyException = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field bio: %w", err)
		}
		u.Bio = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field share_text: %w", err)
		}
		u.ShareText = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field description: %w", err)
		}
		u.Description = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field group_in_common_count: %w", err)
		}
		u.GroupInCommonCount = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userFullInfo#26be6d7c: field commands: %w", err)
		}

		if headerLen > 0 {
			u.Commands = make([]BotCommand, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value BotCommand
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare userFullInfo#26be6d7c: field commands: %w", err)
			}
			u.Commands = append(u.Commands, value)
		}
	}
	return nil
}

// EncodeTDLibJSON encodes u in TDLib API JSON format.
func (u *UserFullInfo) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if u == nil {
		return fmt.Errorf("can't encode userFullInfo#26be6d7c as nil")
	}
	b.ObjStart()
	b.PutID("userFullInfo")
	b.FieldStart("photo")
	if err := u.Photo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode userFullInfo#26be6d7c: field photo: %w", err)
	}
	b.FieldStart("is_blocked")
	b.PutBool(u.IsBlocked)
	b.FieldStart("can_be_called")
	b.PutBool(u.CanBeCalled)
	b.FieldStart("supports_video_calls")
	b.PutBool(u.SupportsVideoCalls)
	b.FieldStart("has_private_calls")
	b.PutBool(u.HasPrivateCalls)
	b.FieldStart("need_phone_number_privacy_exception")
	b.PutBool(u.NeedPhoneNumberPrivacyException)
	b.FieldStart("bio")
	b.PutString(u.Bio)
	b.FieldStart("share_text")
	b.PutString(u.ShareText)
	b.FieldStart("description")
	b.PutString(u.Description)
	b.FieldStart("group_in_common_count")
	b.PutInt32(u.GroupInCommonCount)
	b.FieldStart("commands")
	b.ArrStart()
	for idx, v := range u.Commands {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode userFullInfo#26be6d7c: field commands element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// GetPhoto returns value of Photo field.
func (u *UserFullInfo) GetPhoto() (value ChatPhoto) {
	return u.Photo
}

// GetIsBlocked returns value of IsBlocked field.
func (u *UserFullInfo) GetIsBlocked() (value bool) {
	return u.IsBlocked
}

// GetCanBeCalled returns value of CanBeCalled field.
func (u *UserFullInfo) GetCanBeCalled() (value bool) {
	return u.CanBeCalled
}

// GetSupportsVideoCalls returns value of SupportsVideoCalls field.
func (u *UserFullInfo) GetSupportsVideoCalls() (value bool) {
	return u.SupportsVideoCalls
}

// GetHasPrivateCalls returns value of HasPrivateCalls field.
func (u *UserFullInfo) GetHasPrivateCalls() (value bool) {
	return u.HasPrivateCalls
}

// GetNeedPhoneNumberPrivacyException returns value of NeedPhoneNumberPrivacyException field.
func (u *UserFullInfo) GetNeedPhoneNumberPrivacyException() (value bool) {
	return u.NeedPhoneNumberPrivacyException
}

// GetBio returns value of Bio field.
func (u *UserFullInfo) GetBio() (value string) {
	return u.Bio
}

// GetShareText returns value of ShareText field.
func (u *UserFullInfo) GetShareText() (value string) {
	return u.ShareText
}

// GetDescription returns value of Description field.
func (u *UserFullInfo) GetDescription() (value string) {
	return u.Description
}

// GetGroupInCommonCount returns value of GroupInCommonCount field.
func (u *UserFullInfo) GetGroupInCommonCount() (value int32) {
	return u.GroupInCommonCount
}

// GetCommands returns value of Commands field.
func (u *UserFullInfo) GetCommands() (value []BotCommand) {
	return u.Commands
}
