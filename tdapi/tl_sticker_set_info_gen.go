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

// StickerSetInfo represents TL type `stickerSetInfo#3cd2cd82`.
type StickerSetInfo struct {
	// Identifier of the sticker set
	ID int64
	// Title of the sticker set
	Title string
	// Name of the sticker set
	Name string
	// Sticker set thumbnail in WEBP, TGS, or WEBM format with width and height 100; may be
	// null. The file can be downloaded only before the thumbnail is changed
	Thumbnail Thumbnail
	// Sticker set thumbnail's outline represented as a list of closed vector paths; may be
	// empty. The coordinate system origin is in the upper-left corner
	ThumbnailOutline []ClosedVectorPath
	// True, if the sticker set is owned by the current user
	IsOwned bool
	// True, if the sticker set has been installed by the current user
	IsInstalled bool
	// True, if the sticker set has been archived. A sticker set can't be installed and
	// archived simultaneously
	IsArchived bool
	// True, if the sticker set is official
	IsOfficial bool
	// Type of the stickers in the set
	StickerType StickerTypeClass
	// True, if stickers in the sticker set are custom emoji that must be repainted; for
	// custom emoji sticker sets only
	NeedsRepainting bool
	// True, if stickers in the sticker set are custom emoji that can be used as chat emoji
	// status; for custom emoji sticker sets only
	IsAllowedAsChatEmojiStatus bool
	// True for already viewed trending sticker sets
	IsViewed bool
	// Total number of stickers in the set
	Size int32
	// Up to the first 5 stickers from the set, depending on the context. If the application
	// needs more stickers the full sticker set needs to be requested
	Covers []Sticker
}

// StickerSetInfoTypeID is TL type id of StickerSetInfo.
const StickerSetInfoTypeID = 0x3cd2cd82

// Ensuring interfaces in compile-time for StickerSetInfo.
var (
	_ bin.Encoder     = &StickerSetInfo{}
	_ bin.Decoder     = &StickerSetInfo{}
	_ bin.BareEncoder = &StickerSetInfo{}
	_ bin.BareDecoder = &StickerSetInfo{}
)

func (s *StickerSetInfo) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.Title == "") {
		return false
	}
	if !(s.Name == "") {
		return false
	}
	if !(s.Thumbnail.Zero()) {
		return false
	}
	if !(s.ThumbnailOutline == nil) {
		return false
	}
	if !(s.IsOwned == false) {
		return false
	}
	if !(s.IsInstalled == false) {
		return false
	}
	if !(s.IsArchived == false) {
		return false
	}
	if !(s.IsOfficial == false) {
		return false
	}
	if !(s.StickerType == nil) {
		return false
	}
	if !(s.NeedsRepainting == false) {
		return false
	}
	if !(s.IsAllowedAsChatEmojiStatus == false) {
		return false
	}
	if !(s.IsViewed == false) {
		return false
	}
	if !(s.Size == 0) {
		return false
	}
	if !(s.Covers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StickerSetInfo) String() string {
	if s == nil {
		return "StickerSetInfo(nil)"
	}
	type Alias StickerSetInfo
	return fmt.Sprintf("StickerSetInfo%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StickerSetInfo) TypeID() uint32 {
	return StickerSetInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*StickerSetInfo) TypeName() string {
	return "stickerSetInfo"
}

// TypeInfo returns info about TL type.
func (s *StickerSetInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stickerSetInfo",
		ID:   StickerSetInfoTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "Thumbnail",
			SchemaName: "thumbnail",
		},
		{
			Name:       "ThumbnailOutline",
			SchemaName: "thumbnail_outline",
		},
		{
			Name:       "IsOwned",
			SchemaName: "is_owned",
		},
		{
			Name:       "IsInstalled",
			SchemaName: "is_installed",
		},
		{
			Name:       "IsArchived",
			SchemaName: "is_archived",
		},
		{
			Name:       "IsOfficial",
			SchemaName: "is_official",
		},
		{
			Name:       "StickerType",
			SchemaName: "sticker_type",
		},
		{
			Name:       "NeedsRepainting",
			SchemaName: "needs_repainting",
		},
		{
			Name:       "IsAllowedAsChatEmojiStatus",
			SchemaName: "is_allowed_as_chat_emoji_status",
		},
		{
			Name:       "IsViewed",
			SchemaName: "is_viewed",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "Covers",
			SchemaName: "covers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StickerSetInfo) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerSetInfo#3cd2cd82 as nil")
	}
	b.PutID(StickerSetInfoTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StickerSetInfo) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerSetInfo#3cd2cd82 as nil")
	}
	b.PutLong(s.ID)
	b.PutString(s.Title)
	b.PutString(s.Name)
	if err := s.Thumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickerSetInfo#3cd2cd82: field thumbnail: %w", err)
	}
	b.PutInt(len(s.ThumbnailOutline))
	for idx, v := range s.ThumbnailOutline {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare stickerSetInfo#3cd2cd82: field thumbnail_outline element with index %d: %w", idx, err)
		}
	}
	b.PutBool(s.IsOwned)
	b.PutBool(s.IsInstalled)
	b.PutBool(s.IsArchived)
	b.PutBool(s.IsOfficial)
	if s.StickerType == nil {
		return fmt.Errorf("unable to encode stickerSetInfo#3cd2cd82: field sticker_type is nil")
	}
	if err := s.StickerType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stickerSetInfo#3cd2cd82: field sticker_type: %w", err)
	}
	b.PutBool(s.NeedsRepainting)
	b.PutBool(s.IsAllowedAsChatEmojiStatus)
	b.PutBool(s.IsViewed)
	b.PutInt32(s.Size)
	b.PutInt(len(s.Covers))
	for idx, v := range s.Covers {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare stickerSetInfo#3cd2cd82: field covers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StickerSetInfo) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerSetInfo#3cd2cd82 to nil")
	}
	if err := b.ConsumeID(StickerSetInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StickerSetInfo) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerSetInfo#3cd2cd82 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field title: %w", err)
		}
		s.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field name: %w", err)
		}
		s.Name = value
	}
	{
		if err := s.Thumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field thumbnail: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field thumbnail_outline: %w", err)
		}

		if headerLen > 0 {
			s.ThumbnailOutline = make([]ClosedVectorPath, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ClosedVectorPath
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare stickerSetInfo#3cd2cd82: field thumbnail_outline: %w", err)
			}
			s.ThumbnailOutline = append(s.ThumbnailOutline, value)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_owned: %w", err)
		}
		s.IsOwned = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_installed: %w", err)
		}
		s.IsInstalled = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_archived: %w", err)
		}
		s.IsArchived = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_official: %w", err)
		}
		s.IsOfficial = value
	}
	{
		value, err := DecodeStickerType(b)
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field sticker_type: %w", err)
		}
		s.StickerType = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field needs_repainting: %w", err)
		}
		s.NeedsRepainting = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_allowed_as_chat_emoji_status: %w", err)
		}
		s.IsAllowedAsChatEmojiStatus = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_viewed: %w", err)
		}
		s.IsViewed = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field size: %w", err)
		}
		s.Size = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field covers: %w", err)
		}

		if headerLen > 0 {
			s.Covers = make([]Sticker, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Sticker
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare stickerSetInfo#3cd2cd82: field covers: %w", err)
			}
			s.Covers = append(s.Covers, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StickerSetInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode stickerSetInfo#3cd2cd82 as nil")
	}
	b.ObjStart()
	b.PutID("stickerSetInfo")
	b.Comma()
	b.FieldStart("id")
	b.PutLong(s.ID)
	b.Comma()
	b.FieldStart("title")
	b.PutString(s.Title)
	b.Comma()
	b.FieldStart("name")
	b.PutString(s.Name)
	b.Comma()
	b.FieldStart("thumbnail")
	if err := s.Thumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode stickerSetInfo#3cd2cd82: field thumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("thumbnail_outline")
	b.ArrStart()
	for idx, v := range s.ThumbnailOutline {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode stickerSetInfo#3cd2cd82: field thumbnail_outline element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("is_owned")
	b.PutBool(s.IsOwned)
	b.Comma()
	b.FieldStart("is_installed")
	b.PutBool(s.IsInstalled)
	b.Comma()
	b.FieldStart("is_archived")
	b.PutBool(s.IsArchived)
	b.Comma()
	b.FieldStart("is_official")
	b.PutBool(s.IsOfficial)
	b.Comma()
	b.FieldStart("sticker_type")
	if s.StickerType == nil {
		return fmt.Errorf("unable to encode stickerSetInfo#3cd2cd82: field sticker_type is nil")
	}
	if err := s.StickerType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode stickerSetInfo#3cd2cd82: field sticker_type: %w", err)
	}
	b.Comma()
	b.FieldStart("needs_repainting")
	b.PutBool(s.NeedsRepainting)
	b.Comma()
	b.FieldStart("is_allowed_as_chat_emoji_status")
	b.PutBool(s.IsAllowedAsChatEmojiStatus)
	b.Comma()
	b.FieldStart("is_viewed")
	b.PutBool(s.IsViewed)
	b.Comma()
	b.FieldStart("size")
	b.PutInt32(s.Size)
	b.Comma()
	b.FieldStart("covers")
	b.ArrStart()
	for idx, v := range s.Covers {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode stickerSetInfo#3cd2cd82: field covers element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StickerSetInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode stickerSetInfo#3cd2cd82 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("stickerSetInfo"); err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: %w", err)
			}
		case "id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field id: %w", err)
			}
			s.ID = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field title: %w", err)
			}
			s.Title = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field name: %w", err)
			}
			s.Name = value
		case "thumbnail":
			if err := s.Thumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field thumbnail: %w", err)
			}
		case "thumbnail_outline":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value ClosedVectorPath
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field thumbnail_outline: %w", err)
				}
				s.ThumbnailOutline = append(s.ThumbnailOutline, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field thumbnail_outline: %w", err)
			}
		case "is_owned":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_owned: %w", err)
			}
			s.IsOwned = value
		case "is_installed":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_installed: %w", err)
			}
			s.IsInstalled = value
		case "is_archived":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_archived: %w", err)
			}
			s.IsArchived = value
		case "is_official":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_official: %w", err)
			}
			s.IsOfficial = value
		case "sticker_type":
			value, err := DecodeTDLibJSONStickerType(b)
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field sticker_type: %w", err)
			}
			s.StickerType = value
		case "needs_repainting":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field needs_repainting: %w", err)
			}
			s.NeedsRepainting = value
		case "is_allowed_as_chat_emoji_status":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_allowed_as_chat_emoji_status: %w", err)
			}
			s.IsAllowedAsChatEmojiStatus = value
		case "is_viewed":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field is_viewed: %w", err)
			}
			s.IsViewed = value
		case "size":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field size: %w", err)
			}
			s.Size = value
		case "covers":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value Sticker
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field covers: %w", err)
				}
				s.Covers = append(s.Covers, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode stickerSetInfo#3cd2cd82: field covers: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (s *StickerSetInfo) GetID() (value int64) {
	if s == nil {
		return
	}
	return s.ID
}

// GetTitle returns value of Title field.
func (s *StickerSetInfo) GetTitle() (value string) {
	if s == nil {
		return
	}
	return s.Title
}

// GetName returns value of Name field.
func (s *StickerSetInfo) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

// GetThumbnail returns value of Thumbnail field.
func (s *StickerSetInfo) GetThumbnail() (value Thumbnail) {
	if s == nil {
		return
	}
	return s.Thumbnail
}

// GetThumbnailOutline returns value of ThumbnailOutline field.
func (s *StickerSetInfo) GetThumbnailOutline() (value []ClosedVectorPath) {
	if s == nil {
		return
	}
	return s.ThumbnailOutline
}

// GetIsOwned returns value of IsOwned field.
func (s *StickerSetInfo) GetIsOwned() (value bool) {
	if s == nil {
		return
	}
	return s.IsOwned
}

// GetIsInstalled returns value of IsInstalled field.
func (s *StickerSetInfo) GetIsInstalled() (value bool) {
	if s == nil {
		return
	}
	return s.IsInstalled
}

// GetIsArchived returns value of IsArchived field.
func (s *StickerSetInfo) GetIsArchived() (value bool) {
	if s == nil {
		return
	}
	return s.IsArchived
}

// GetIsOfficial returns value of IsOfficial field.
func (s *StickerSetInfo) GetIsOfficial() (value bool) {
	if s == nil {
		return
	}
	return s.IsOfficial
}

// GetStickerType returns value of StickerType field.
func (s *StickerSetInfo) GetStickerType() (value StickerTypeClass) {
	if s == nil {
		return
	}
	return s.StickerType
}

// GetNeedsRepainting returns value of NeedsRepainting field.
func (s *StickerSetInfo) GetNeedsRepainting() (value bool) {
	if s == nil {
		return
	}
	return s.NeedsRepainting
}

// GetIsAllowedAsChatEmojiStatus returns value of IsAllowedAsChatEmojiStatus field.
func (s *StickerSetInfo) GetIsAllowedAsChatEmojiStatus() (value bool) {
	if s == nil {
		return
	}
	return s.IsAllowedAsChatEmojiStatus
}

// GetIsViewed returns value of IsViewed field.
func (s *StickerSetInfo) GetIsViewed() (value bool) {
	if s == nil {
		return
	}
	return s.IsViewed
}

// GetSize returns value of Size field.
func (s *StickerSetInfo) GetSize() (value int32) {
	if s == nil {
		return
	}
	return s.Size
}

// GetCovers returns value of Covers field.
func (s *StickerSetInfo) GetCovers() (value []Sticker) {
	if s == nil {
		return
	}
	return s.Covers
}
