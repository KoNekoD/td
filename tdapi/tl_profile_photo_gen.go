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

// ProfilePhoto represents TL type `profilePhoto#c2dc3c5e`.
type ProfilePhoto struct {
	// Photo identifier; 0 for an empty photo. Can be used to find a photo in a list of user
	// profile photos
	ID int64
	// A small (160x160) user profile photo. The file can be downloaded only before the photo
	// is changed
	Small File
	// A big (640x640) user profile photo. The file can be downloaded only before the photo
	// is changed
	Big File
	// User profile photo minithumbnail; may be null
	Minithumbnail Minithumbnail
	// True, if the photo has animated variant
	HasAnimation bool
	// True, if the photo is visible only for the current user
	IsPersonal bool
}

// ProfilePhotoTypeID is TL type id of ProfilePhoto.
const ProfilePhotoTypeID = 0xc2dc3c5e

// Ensuring interfaces in compile-time for ProfilePhoto.
var (
	_ bin.Encoder     = &ProfilePhoto{}
	_ bin.Decoder     = &ProfilePhoto{}
	_ bin.BareEncoder = &ProfilePhoto{}
	_ bin.BareDecoder = &ProfilePhoto{}
)

func (p *ProfilePhoto) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.Small.Zero()) {
		return false
	}
	if !(p.Big.Zero()) {
		return false
	}
	if !(p.Minithumbnail.Zero()) {
		return false
	}
	if !(p.HasAnimation == false) {
		return false
	}
	if !(p.IsPersonal == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *ProfilePhoto) String() string {
	if p == nil {
		return "ProfilePhoto(nil)"
	}
	type Alias ProfilePhoto
	return fmt.Sprintf("ProfilePhoto%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ProfilePhoto) TypeID() uint32 {
	return ProfilePhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*ProfilePhoto) TypeName() string {
	return "profilePhoto"
}

// TypeInfo returns info about TL type.
func (p *ProfilePhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "profilePhoto",
		ID:   ProfilePhotoTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Small",
			SchemaName: "small",
		},
		{
			Name:       "Big",
			SchemaName: "big",
		},
		{
			Name:       "Minithumbnail",
			SchemaName: "minithumbnail",
		},
		{
			Name:       "HasAnimation",
			SchemaName: "has_animation",
		},
		{
			Name:       "IsPersonal",
			SchemaName: "is_personal",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *ProfilePhoto) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode profilePhoto#c2dc3c5e as nil")
	}
	b.PutID(ProfilePhotoTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *ProfilePhoto) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode profilePhoto#c2dc3c5e as nil")
	}
	b.PutLong(p.ID)
	if err := p.Small.Encode(b); err != nil {
		return fmt.Errorf("unable to encode profilePhoto#c2dc3c5e: field small: %w", err)
	}
	if err := p.Big.Encode(b); err != nil {
		return fmt.Errorf("unable to encode profilePhoto#c2dc3c5e: field big: %w", err)
	}
	if err := p.Minithumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode profilePhoto#c2dc3c5e: field minithumbnail: %w", err)
	}
	b.PutBool(p.HasAnimation)
	b.PutBool(p.IsPersonal)
	return nil
}

// Decode implements bin.Decoder.
func (p *ProfilePhoto) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode profilePhoto#c2dc3c5e to nil")
	}
	if err := b.ConsumeID(ProfilePhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *ProfilePhoto) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode profilePhoto#c2dc3c5e to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field id: %w", err)
		}
		p.ID = value
	}
	{
		if err := p.Small.Decode(b); err != nil {
			return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field small: %w", err)
		}
	}
	{
		if err := p.Big.Decode(b); err != nil {
			return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field big: %w", err)
		}
	}
	{
		if err := p.Minithumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field minithumbnail: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field has_animation: %w", err)
		}
		p.HasAnimation = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field is_personal: %w", err)
		}
		p.IsPersonal = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (p *ProfilePhoto) EncodeTDLibJSON(b tdjson.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode profilePhoto#c2dc3c5e as nil")
	}
	b.ObjStart()
	b.PutID("profilePhoto")
	b.Comma()
	b.FieldStart("id")
	b.PutLong(p.ID)
	b.Comma()
	b.FieldStart("small")
	if err := p.Small.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode profilePhoto#c2dc3c5e: field small: %w", err)
	}
	b.Comma()
	b.FieldStart("big")
	if err := p.Big.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode profilePhoto#c2dc3c5e: field big: %w", err)
	}
	b.Comma()
	b.FieldStart("minithumbnail")
	if err := p.Minithumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode profilePhoto#c2dc3c5e: field minithumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("has_animation")
	b.PutBool(p.HasAnimation)
	b.Comma()
	b.FieldStart("is_personal")
	b.PutBool(p.IsPersonal)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (p *ProfilePhoto) DecodeTDLibJSON(b tdjson.Decoder) error {
	if p == nil {
		return fmt.Errorf("can't decode profilePhoto#c2dc3c5e to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("profilePhoto"); err != nil {
				return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field id: %w", err)
			}
			p.ID = value
		case "small":
			if err := p.Small.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field small: %w", err)
			}
		case "big":
			if err := p.Big.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field big: %w", err)
			}
		case "minithumbnail":
			if err := p.Minithumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field minithumbnail: %w", err)
			}
		case "has_animation":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field has_animation: %w", err)
			}
			p.HasAnimation = value
		case "is_personal":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode profilePhoto#c2dc3c5e: field is_personal: %w", err)
			}
			p.IsPersonal = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (p *ProfilePhoto) GetID() (value int64) {
	if p == nil {
		return
	}
	return p.ID
}

// GetSmall returns value of Small field.
func (p *ProfilePhoto) GetSmall() (value File) {
	if p == nil {
		return
	}
	return p.Small
}

// GetBig returns value of Big field.
func (p *ProfilePhoto) GetBig() (value File) {
	if p == nil {
		return
	}
	return p.Big
}

// GetMinithumbnail returns value of Minithumbnail field.
func (p *ProfilePhoto) GetMinithumbnail() (value Minithumbnail) {
	if p == nil {
		return
	}
	return p.Minithumbnail
}

// GetHasAnimation returns value of HasAnimation field.
func (p *ProfilePhoto) GetHasAnimation() (value bool) {
	if p == nil {
		return
	}
	return p.HasAnimation
}

// GetIsPersonal returns value of IsPersonal field.
func (p *ProfilePhoto) GetIsPersonal() (value bool) {
	if p == nil {
		return
	}
	return p.IsPersonal
}
