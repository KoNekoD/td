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

// UserProfilePhotoEmpty represents TL type `userProfilePhotoEmpty#4f11bae1`.
// Profile photo has not been set, or was hidden.
//
// See https://core.telegram.org/constructor/userProfilePhotoEmpty for reference.
type UserProfilePhotoEmpty struct {
}

// UserProfilePhotoEmptyTypeID is TL type id of UserProfilePhotoEmpty.
const UserProfilePhotoEmptyTypeID = 0x4f11bae1

// construct implements constructor of UserProfilePhotoClass.
func (u UserProfilePhotoEmpty) construct() UserProfilePhotoClass { return &u }

// Ensuring interfaces in compile-time for UserProfilePhotoEmpty.
var (
	_ bin.Encoder     = &UserProfilePhotoEmpty{}
	_ bin.Decoder     = &UserProfilePhotoEmpty{}
	_ bin.BareEncoder = &UserProfilePhotoEmpty{}
	_ bin.BareDecoder = &UserProfilePhotoEmpty{}

	_ UserProfilePhotoClass = &UserProfilePhotoEmpty{}
)

func (u *UserProfilePhotoEmpty) Zero() bool {
	if u == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserProfilePhotoEmpty) String() string {
	if u == nil {
		return "UserProfilePhotoEmpty(nil)"
	}
	type Alias UserProfilePhotoEmpty
	return fmt.Sprintf("UserProfilePhotoEmpty%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserProfilePhotoEmpty) TypeID() uint32 {
	return UserProfilePhotoEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*UserProfilePhotoEmpty) TypeName() string {
	return "userProfilePhotoEmpty"
}

// TypeInfo returns info about TL type.
func (u *UserProfilePhotoEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userProfilePhotoEmpty",
		ID:   UserProfilePhotoEmptyTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserProfilePhotoEmpty) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userProfilePhotoEmpty#4f11bae1 as nil")
	}
	b.PutID(UserProfilePhotoEmptyTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserProfilePhotoEmpty) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userProfilePhotoEmpty#4f11bae1 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *UserProfilePhotoEmpty) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userProfilePhotoEmpty#4f11bae1 to nil")
	}
	if err := b.ConsumeID(UserProfilePhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode userProfilePhotoEmpty#4f11bae1: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserProfilePhotoEmpty) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userProfilePhotoEmpty#4f11bae1 to nil")
	}
	return nil
}

// UserProfilePhoto represents TL type `userProfilePhoto#82d1f706`.
// User profile photo.
//
// See https://core.telegram.org/constructor/userProfilePhoto for reference.
type UserProfilePhoto struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether an animated profile picture¹ is available for this user
	//
	// Links:
	//  1) https://core.telegram.org/api/files#animated-profile-pictures
	HasVideo bool
	// Identifier of the respective photoParameter added in Layer 2¹
	//
	// Links:
	//  1) https://core.telegram.org/api/layers#layer-2
	PhotoID int64
	// StrippedThumb field of UserProfilePhoto.
	//
	// Use SetStrippedThumb and GetStrippedThumb helpers.
	StrippedThumb []byte
	// DC ID where the photo is stored
	DCID int
}

// UserProfilePhotoTypeID is TL type id of UserProfilePhoto.
const UserProfilePhotoTypeID = 0x82d1f706

// construct implements constructor of UserProfilePhotoClass.
func (u UserProfilePhoto) construct() UserProfilePhotoClass { return &u }

// Ensuring interfaces in compile-time for UserProfilePhoto.
var (
	_ bin.Encoder     = &UserProfilePhoto{}
	_ bin.Decoder     = &UserProfilePhoto{}
	_ bin.BareEncoder = &UserProfilePhoto{}
	_ bin.BareDecoder = &UserProfilePhoto{}

	_ UserProfilePhotoClass = &UserProfilePhoto{}
)

func (u *UserProfilePhoto) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.HasVideo == false) {
		return false
	}
	if !(u.PhotoID == 0) {
		return false
	}
	if !(u.StrippedThumb == nil) {
		return false
	}
	if !(u.DCID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *UserProfilePhoto) String() string {
	if u == nil {
		return "UserProfilePhoto(nil)"
	}
	type Alias UserProfilePhoto
	return fmt.Sprintf("UserProfilePhoto%+v", Alias(*u))
}

// FillFrom fills UserProfilePhoto from given interface.
func (u *UserProfilePhoto) FillFrom(from interface {
	GetHasVideo() (value bool)
	GetPhotoID() (value int64)
	GetStrippedThumb() (value []byte, ok bool)
	GetDCID() (value int)
}) {
	u.HasVideo = from.GetHasVideo()
	u.PhotoID = from.GetPhotoID()
	if val, ok := from.GetStrippedThumb(); ok {
		u.StrippedThumb = val
	}

	u.DCID = from.GetDCID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserProfilePhoto) TypeID() uint32 {
	return UserProfilePhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*UserProfilePhoto) TypeName() string {
	return "userProfilePhoto"
}

// TypeInfo returns info about TL type.
func (u *UserProfilePhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "userProfilePhoto",
		ID:   UserProfilePhotoTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasVideo",
			SchemaName: "has_video",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "PhotoID",
			SchemaName: "photo_id",
		},
		{
			Name:       "StrippedThumb",
			SchemaName: "stripped_thumb",
			Null:       !u.Flags.Has(1),
		},
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *UserProfilePhoto) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userProfilePhoto#82d1f706 as nil")
	}
	b.PutID(UserProfilePhotoTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *UserProfilePhoto) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode userProfilePhoto#82d1f706 as nil")
	}
	if !(u.HasVideo == false) {
		u.Flags.Set(0)
	}
	if !(u.StrippedThumb == nil) {
		u.Flags.Set(1)
	}
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode userProfilePhoto#82d1f706: field flags: %w", err)
	}
	b.PutLong(u.PhotoID)
	if u.Flags.Has(1) {
		b.PutBytes(u.StrippedThumb)
	}
	b.PutInt(u.DCID)
	return nil
}

// Decode implements bin.Decoder.
func (u *UserProfilePhoto) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userProfilePhoto#82d1f706 to nil")
	}
	if err := b.ConsumeID(UserProfilePhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode userProfilePhoto#82d1f706: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *UserProfilePhoto) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode userProfilePhoto#82d1f706 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#82d1f706: field flags: %w", err)
		}
	}
	u.HasVideo = u.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#82d1f706: field photo_id: %w", err)
		}
		u.PhotoID = value
	}
	if u.Flags.Has(1) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#82d1f706: field stripped_thumb: %w", err)
		}
		u.StrippedThumb = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode userProfilePhoto#82d1f706: field dc_id: %w", err)
		}
		u.DCID = value
	}
	return nil
}

// SetHasVideo sets value of HasVideo conditional field.
func (u *UserProfilePhoto) SetHasVideo(value bool) {
	if value {
		u.Flags.Set(0)
		u.HasVideo = true
	} else {
		u.Flags.Unset(0)
		u.HasVideo = false
	}
}

// GetHasVideo returns value of HasVideo conditional field.
func (u *UserProfilePhoto) GetHasVideo() (value bool) {
	return u.Flags.Has(0)
}

// GetPhotoID returns value of PhotoID field.
func (u *UserProfilePhoto) GetPhotoID() (value int64) {
	return u.PhotoID
}

// SetStrippedThumb sets value of StrippedThumb conditional field.
func (u *UserProfilePhoto) SetStrippedThumb(value []byte) {
	u.Flags.Set(1)
	u.StrippedThumb = value
}

// GetStrippedThumb returns value of StrippedThumb conditional field and
// boolean which is true if field was set.
func (u *UserProfilePhoto) GetStrippedThumb() (value []byte, ok bool) {
	if !u.Flags.Has(1) {
		return value, false
	}
	return u.StrippedThumb, true
}

// GetDCID returns value of DCID field.
func (u *UserProfilePhoto) GetDCID() (value int) {
	return u.DCID
}

// UserProfilePhotoClass represents UserProfilePhoto generic type.
//
// See https://core.telegram.org/type/UserProfilePhoto for reference.
//
// Example:
//  g, err := tg.DecodeUserProfilePhoto(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.UserProfilePhotoEmpty: // userProfilePhotoEmpty#4f11bae1
//  case *tg.UserProfilePhoto: // userProfilePhoto#82d1f706
//  default: panic(v)
//  }
type UserProfilePhotoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() UserProfilePhotoClass

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

	// AsNotEmpty tries to map UserProfilePhotoClass to UserProfilePhoto.
	AsNotEmpty() (*UserProfilePhoto, bool)
}

// AsNotEmpty tries to map UserProfilePhotoEmpty to UserProfilePhoto.
func (u *UserProfilePhotoEmpty) AsNotEmpty() (*UserProfilePhoto, bool) {
	return nil, false
}

// AsNotEmpty tries to map UserProfilePhoto to UserProfilePhoto.
func (u *UserProfilePhoto) AsNotEmpty() (*UserProfilePhoto, bool) {
	return u, true
}

// DecodeUserProfilePhoto implements binary de-serialization for UserProfilePhotoClass.
func DecodeUserProfilePhoto(buf *bin.Buffer) (UserProfilePhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UserProfilePhotoEmptyTypeID:
		// Decoding userProfilePhotoEmpty#4f11bae1.
		v := UserProfilePhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserProfilePhotoClass: %w", err)
		}
		return &v, nil
	case UserProfilePhotoTypeID:
		// Decoding userProfilePhoto#82d1f706.
		v := UserProfilePhoto{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserProfilePhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UserProfilePhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// UserProfilePhoto boxes the UserProfilePhotoClass providing a helper.
type UserProfilePhotoBox struct {
	UserProfilePhoto UserProfilePhotoClass
}

// Decode implements bin.Decoder for UserProfilePhotoBox.
func (b *UserProfilePhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UserProfilePhotoBox to nil")
	}
	v, err := DecodeUserProfilePhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.UserProfilePhoto = v
	return nil
}

// Encode implements bin.Encode for UserProfilePhotoBox.
func (b *UserProfilePhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.UserProfilePhoto == nil {
		return fmt.Errorf("unable to encode UserProfilePhotoClass as nil")
	}
	return b.UserProfilePhoto.Encode(buf)
}
