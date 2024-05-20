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

// Audio represents TL type `audio#a3769600`.
type Audio struct {
	// Duration of the audio, in seconds; as defined by the sender
	Duration int32
	// Title of the audio; as defined by the sender
	Title string
	// Performer of the audio; as defined by the sender
	Performer string
	// Original name of the file; as defined by the sender
	FileName string
	// The MIME type of the file; as defined by the sender
	MimeType string
	// The minithumbnail of the album cover; may be null
	AlbumCoverMinithumbnail Minithumbnail
	// The thumbnail of the album cover in JPEG format; as defined by the sender. The full
	// size thumbnail is supposed to be extracted from the downloaded audio file; may be null
	AlbumCoverThumbnail Thumbnail
	// Album cover variants to use if the downloaded audio file contains no album cover.
	// Provided thumbnail dimensions are approximate
	ExternalAlbumCovers []Thumbnail
	// File containing the audio
	Audio File
}

// AudioTypeID is TL type id of Audio.
const AudioTypeID = 0xa3769600

// Ensuring interfaces in compile-time for Audio.
var (
	_ bin.Encoder     = &Audio{}
	_ bin.Decoder     = &Audio{}
	_ bin.BareEncoder = &Audio{}
	_ bin.BareDecoder = &Audio{}
)

func (a *Audio) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Duration == 0) {
		return false
	}
	if !(a.Title == "") {
		return false
	}
	if !(a.Performer == "") {
		return false
	}
	if !(a.FileName == "") {
		return false
	}
	if !(a.MimeType == "") {
		return false
	}
	if !(a.AlbumCoverMinithumbnail.Zero()) {
		return false
	}
	if !(a.AlbumCoverThumbnail.Zero()) {
		return false
	}
	if !(a.ExternalAlbumCovers == nil) {
		return false
	}
	if !(a.Audio.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *Audio) String() string {
	if a == nil {
		return "Audio(nil)"
	}
	type Alias Audio
	return fmt.Sprintf("Audio%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Audio) TypeID() uint32 {
	return AudioTypeID
}

// TypeName returns name of type in TL schema.
func (*Audio) TypeName() string {
	return "audio"
}

// TypeInfo returns info about TL type.
func (a *Audio) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "audio",
		ID:   AudioTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Performer",
			SchemaName: "performer",
		},
		{
			Name:       "FileName",
			SchemaName: "file_name",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "AlbumCoverMinithumbnail",
			SchemaName: "album_cover_minithumbnail",
		},
		{
			Name:       "AlbumCoverThumbnail",
			SchemaName: "album_cover_thumbnail",
		},
		{
			Name:       "ExternalAlbumCovers",
			SchemaName: "external_album_covers",
		},
		{
			Name:       "Audio",
			SchemaName: "audio",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *Audio) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode audio#a3769600 as nil")
	}
	b.PutID(AudioTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *Audio) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode audio#a3769600 as nil")
	}
	b.PutInt32(a.Duration)
	b.PutString(a.Title)
	b.PutString(a.Performer)
	b.PutString(a.FileName)
	b.PutString(a.MimeType)
	if err := a.AlbumCoverMinithumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode audio#a3769600: field album_cover_minithumbnail: %w", err)
	}
	if err := a.AlbumCoverThumbnail.Encode(b); err != nil {
		return fmt.Errorf("unable to encode audio#a3769600: field album_cover_thumbnail: %w", err)
	}
	b.PutInt(len(a.ExternalAlbumCovers))
	for idx, v := range a.ExternalAlbumCovers {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare audio#a3769600: field external_album_covers element with index %d: %w", idx, err)
		}
	}
	if err := a.Audio.Encode(b); err != nil {
		return fmt.Errorf("unable to encode audio#a3769600: field audio: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *Audio) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode audio#a3769600 to nil")
	}
	if err := b.ConsumeID(AudioTypeID); err != nil {
		return fmt.Errorf("unable to decode audio#a3769600: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *Audio) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode audio#a3769600 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode audio#a3769600: field duration: %w", err)
		}
		a.Duration = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode audio#a3769600: field title: %w", err)
		}
		a.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode audio#a3769600: field performer: %w", err)
		}
		a.Performer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode audio#a3769600: field file_name: %w", err)
		}
		a.FileName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode audio#a3769600: field mime_type: %w", err)
		}
		a.MimeType = value
	}
	{
		if err := a.AlbumCoverMinithumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode audio#a3769600: field album_cover_minithumbnail: %w", err)
		}
	}
	{
		if err := a.AlbumCoverThumbnail.Decode(b); err != nil {
			return fmt.Errorf("unable to decode audio#a3769600: field album_cover_thumbnail: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode audio#a3769600: field external_album_covers: %w", err)
		}

		if headerLen > 0 {
			a.ExternalAlbumCovers = make([]Thumbnail, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Thumbnail
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare audio#a3769600: field external_album_covers: %w", err)
			}
			a.ExternalAlbumCovers = append(a.ExternalAlbumCovers, value)
		}
	}
	{
		if err := a.Audio.Decode(b); err != nil {
			return fmt.Errorf("unable to decode audio#a3769600: field audio: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *Audio) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode audio#a3769600 as nil")
	}
	b.ObjStart()
	b.PutID("audio")
	b.Comma()
	b.FieldStart("duration")
	b.PutInt32(a.Duration)
	b.Comma()
	b.FieldStart("title")
	b.PutString(a.Title)
	b.Comma()
	b.FieldStart("performer")
	b.PutString(a.Performer)
	b.Comma()
	b.FieldStart("file_name")
	b.PutString(a.FileName)
	b.Comma()
	b.FieldStart("mime_type")
	b.PutString(a.MimeType)
	b.Comma()
	b.FieldStart("album_cover_minithumbnail")
	if err := a.AlbumCoverMinithumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode audio#a3769600: field album_cover_minithumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("album_cover_thumbnail")
	if err := a.AlbumCoverThumbnail.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode audio#a3769600: field album_cover_thumbnail: %w", err)
	}
	b.Comma()
	b.FieldStart("external_album_covers")
	b.ArrStart()
	for idx, v := range a.ExternalAlbumCovers {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode audio#a3769600: field external_album_covers element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("audio")
	if err := a.Audio.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode audio#a3769600: field audio: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *Audio) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode audio#a3769600 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("audio"); err != nil {
				return fmt.Errorf("unable to decode audio#a3769600: %w", err)
			}
		case "duration":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode audio#a3769600: field duration: %w", err)
			}
			a.Duration = value
		case "title":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode audio#a3769600: field title: %w", err)
			}
			a.Title = value
		case "performer":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode audio#a3769600: field performer: %w", err)
			}
			a.Performer = value
		case "file_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode audio#a3769600: field file_name: %w", err)
			}
			a.FileName = value
		case "mime_type":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode audio#a3769600: field mime_type: %w", err)
			}
			a.MimeType = value
		case "album_cover_minithumbnail":
			if err := a.AlbumCoverMinithumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode audio#a3769600: field album_cover_minithumbnail: %w", err)
			}
		case "album_cover_thumbnail":
			if err := a.AlbumCoverThumbnail.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode audio#a3769600: field album_cover_thumbnail: %w", err)
			}
		case "external_album_covers":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value Thumbnail
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode audio#a3769600: field external_album_covers: %w", err)
				}
				a.ExternalAlbumCovers = append(a.ExternalAlbumCovers, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode audio#a3769600: field external_album_covers: %w", err)
			}
		case "audio":
			if err := a.Audio.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode audio#a3769600: field audio: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetDuration returns value of Duration field.
func (a *Audio) GetDuration() (value int32) {
	if a == nil {
		return
	}
	return a.Duration
}

// GetTitle returns value of Title field.
func (a *Audio) GetTitle() (value string) {
	if a == nil {
		return
	}
	return a.Title
}

// GetPerformer returns value of Performer field.
func (a *Audio) GetPerformer() (value string) {
	if a == nil {
		return
	}
	return a.Performer
}

// GetFileName returns value of FileName field.
func (a *Audio) GetFileName() (value string) {
	if a == nil {
		return
	}
	return a.FileName
}

// GetMimeType returns value of MimeType field.
func (a *Audio) GetMimeType() (value string) {
	if a == nil {
		return
	}
	return a.MimeType
}

// GetAlbumCoverMinithumbnail returns value of AlbumCoverMinithumbnail field.
func (a *Audio) GetAlbumCoverMinithumbnail() (value Minithumbnail) {
	if a == nil {
		return
	}
	return a.AlbumCoverMinithumbnail
}

// GetAlbumCoverThumbnail returns value of AlbumCoverThumbnail field.
func (a *Audio) GetAlbumCoverThumbnail() (value Thumbnail) {
	if a == nil {
		return
	}
	return a.AlbumCoverThumbnail
}

// GetExternalAlbumCovers returns value of ExternalAlbumCovers field.
func (a *Audio) GetExternalAlbumCovers() (value []Thumbnail) {
	if a == nil {
		return
	}
	return a.ExternalAlbumCovers
}

// GetAudio returns value of Audio field.
func (a *Audio) GetAudio() (value File) {
	if a == nil {
		return
	}
	return a.Audio
}
