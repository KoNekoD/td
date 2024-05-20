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

// VideoSize represents TL type `videoSize#de33b094`.
// An animated profile picture¹ in MPEG4 format
//
// Links:
//  1. https://core.telegram.org/api/files#animated-profile-pictures
//
// See https://core.telegram.org/constructor/videoSize for reference.
type VideoSize struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// u for animated profile pictures, and v for trimmed and downscaled video previews
	Type string
	// Video width
	W int
	// Video height
	H int
	// File size
	Size int
	// Timestamp that should be shown as static preview to the user (seconds)
	//
	// Use SetVideoStartTs and GetVideoStartTs helpers.
	VideoStartTs float64
}

// VideoSizeTypeID is TL type id of VideoSize.
const VideoSizeTypeID = 0xde33b094

// construct implements constructor of VideoSizeClass.
func (v VideoSize) construct() VideoSizeClass { return &v }

// Ensuring interfaces in compile-time for VideoSize.
var (
	_ bin.Encoder     = &VideoSize{}
	_ bin.Decoder     = &VideoSize{}
	_ bin.BareEncoder = &VideoSize{}
	_ bin.BareDecoder = &VideoSize{}

	_ VideoSizeClass = &VideoSize{}
)

func (v *VideoSize) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Flags.Zero()) {
		return false
	}
	if !(v.Type == "") {
		return false
	}
	if !(v.W == 0) {
		return false
	}
	if !(v.H == 0) {
		return false
	}
	if !(v.Size == 0) {
		return false
	}
	if !(v.VideoStartTs == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *VideoSize) String() string {
	if v == nil {
		return "VideoSize(nil)"
	}
	type Alias VideoSize
	return fmt.Sprintf("VideoSize%+v", Alias(*v))
}

// FillFrom fills VideoSize from given interface.
func (v *VideoSize) FillFrom(from interface {
	GetType() (value string)
	GetW() (value int)
	GetH() (value int)
	GetSize() (value int)
	GetVideoStartTs() (value float64, ok bool)
}) {
	v.Type = from.GetType()
	v.W = from.GetW()
	v.H = from.GetH()
	v.Size = from.GetSize()
	if val, ok := from.GetVideoStartTs(); ok {
		v.VideoStartTs = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*VideoSize) TypeID() uint32 {
	return VideoSizeTypeID
}

// TypeName returns name of type in TL schema.
func (*VideoSize) TypeName() string {
	return "videoSize"
}

// TypeInfo returns info about TL type.
func (v *VideoSize) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "videoSize",
		ID:   VideoSizeTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "W",
			SchemaName: "w",
		},
		{
			Name:       "H",
			SchemaName: "h",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "VideoStartTs",
			SchemaName: "video_start_ts",
			Null:       !v.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (v *VideoSize) SetFlags() {
	if !(v.VideoStartTs == 0) {
		v.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (v *VideoSize) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoSize#de33b094 as nil")
	}
	b.PutID(VideoSizeTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *VideoSize) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoSize#de33b094 as nil")
	}
	v.SetFlags()
	if err := v.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode videoSize#de33b094: field flags: %w", err)
	}
	b.PutString(v.Type)
	b.PutInt(v.W)
	b.PutInt(v.H)
	b.PutInt(v.Size)
	if v.Flags.Has(0) {
		b.PutDouble(v.VideoStartTs)
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *VideoSize) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoSize#de33b094 to nil")
	}
	if err := b.ConsumeID(VideoSizeTypeID); err != nil {
		return fmt.Errorf("unable to decode videoSize#de33b094: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *VideoSize) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoSize#de33b094 to nil")
	}
	{
		if err := v.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field type: %w", err)
		}
		v.Type = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field w: %w", err)
		}
		v.W = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field h: %w", err)
		}
		v.H = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field size: %w", err)
		}
		v.Size = value
	}
	if v.Flags.Has(0) {
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode videoSize#de33b094: field video_start_ts: %w", err)
		}
		v.VideoStartTs = value
	}
	return nil
}

// GetType returns value of Type field.
func (v *VideoSize) GetType() (value string) {
	if v == nil {
		return
	}
	return v.Type
}

// GetW returns value of W field.
func (v *VideoSize) GetW() (value int) {
	if v == nil {
		return
	}
	return v.W
}

// GetH returns value of H field.
func (v *VideoSize) GetH() (value int) {
	if v == nil {
		return
	}
	return v.H
}

// GetSize returns value of Size field.
func (v *VideoSize) GetSize() (value int) {
	if v == nil {
		return
	}
	return v.Size
}

// SetVideoStartTs sets value of VideoStartTs conditional field.
func (v *VideoSize) SetVideoStartTs(value float64) {
	v.Flags.Set(0)
	v.VideoStartTs = value
}

// GetVideoStartTs returns value of VideoStartTs conditional field and
// boolean which is true if field was set.
func (v *VideoSize) GetVideoStartTs() (value float64, ok bool) {
	if v == nil {
		return
	}
	if !v.Flags.Has(0) {
		return value, false
	}
	return v.VideoStartTs, true
}

// VideoSizeEmojiMarkup represents TL type `videoSizeEmojiMarkup#f85c413c`.
// An animated profile picture¹ based on a custom emoji sticker².
//
// Links:
//  1. https://core.telegram.org/api/files#animated-profile-pictures
//  2. https://core.telegram.org/api/custom-emoji
//
// See https://core.telegram.org/constructor/videoSizeEmojiMarkup for reference.
type VideoSizeEmojiMarkup struct {
	// Custom emoji ID¹: the custom emoji sticker is shown at the center of the profile
	// picture and occupies at most 67% of it.
	//
	// Links:
	//  1) https://core.telegram.org/api/custom-emoji
	EmojiID int64
	// 1, 2, 3 or 4 RBG-24 colors used to generate a solid (1), gradient (2) or freeform
	// gradient (3, 4) background, similar to how fill wallpapers¹ are generated. The
	// rotation angle for gradient backgrounds is 0.
	//
	// Links:
	//  1) https://core.telegram.org/api/wallpapers#fill-types
	BackgroundColors []int
}

// VideoSizeEmojiMarkupTypeID is TL type id of VideoSizeEmojiMarkup.
const VideoSizeEmojiMarkupTypeID = 0xf85c413c

// construct implements constructor of VideoSizeClass.
func (v VideoSizeEmojiMarkup) construct() VideoSizeClass { return &v }

// Ensuring interfaces in compile-time for VideoSizeEmojiMarkup.
var (
	_ bin.Encoder     = &VideoSizeEmojiMarkup{}
	_ bin.Decoder     = &VideoSizeEmojiMarkup{}
	_ bin.BareEncoder = &VideoSizeEmojiMarkup{}
	_ bin.BareDecoder = &VideoSizeEmojiMarkup{}

	_ VideoSizeClass = &VideoSizeEmojiMarkup{}
)

func (v *VideoSizeEmojiMarkup) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.EmojiID == 0) {
		return false
	}
	if !(v.BackgroundColors == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *VideoSizeEmojiMarkup) String() string {
	if v == nil {
		return "VideoSizeEmojiMarkup(nil)"
	}
	type Alias VideoSizeEmojiMarkup
	return fmt.Sprintf("VideoSizeEmojiMarkup%+v", Alias(*v))
}

// FillFrom fills VideoSizeEmojiMarkup from given interface.
func (v *VideoSizeEmojiMarkup) FillFrom(from interface {
	GetEmojiID() (value int64)
	GetBackgroundColors() (value []int)
}) {
	v.EmojiID = from.GetEmojiID()
	v.BackgroundColors = from.GetBackgroundColors()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*VideoSizeEmojiMarkup) TypeID() uint32 {
	return VideoSizeEmojiMarkupTypeID
}

// TypeName returns name of type in TL schema.
func (*VideoSizeEmojiMarkup) TypeName() string {
	return "videoSizeEmojiMarkup"
}

// TypeInfo returns info about TL type.
func (v *VideoSizeEmojiMarkup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "videoSizeEmojiMarkup",
		ID:   VideoSizeEmojiMarkupTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EmojiID",
			SchemaName: "emoji_id",
		},
		{
			Name:       "BackgroundColors",
			SchemaName: "background_colors",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *VideoSizeEmojiMarkup) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoSizeEmojiMarkup#f85c413c as nil")
	}
	b.PutID(VideoSizeEmojiMarkupTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *VideoSizeEmojiMarkup) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoSizeEmojiMarkup#f85c413c as nil")
	}
	b.PutLong(v.EmojiID)
	b.PutVectorHeader(len(v.BackgroundColors))
	for _, v := range v.BackgroundColors {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *VideoSizeEmojiMarkup) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoSizeEmojiMarkup#f85c413c to nil")
	}
	if err := b.ConsumeID(VideoSizeEmojiMarkupTypeID); err != nil {
		return fmt.Errorf("unable to decode videoSizeEmojiMarkup#f85c413c: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *VideoSizeEmojiMarkup) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoSizeEmojiMarkup#f85c413c to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode videoSizeEmojiMarkup#f85c413c: field emoji_id: %w", err)
		}
		v.EmojiID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode videoSizeEmojiMarkup#f85c413c: field background_colors: %w", err)
		}

		if headerLen > 0 {
			v.BackgroundColors = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode videoSizeEmojiMarkup#f85c413c: field background_colors: %w", err)
			}
			v.BackgroundColors = append(v.BackgroundColors, value)
		}
	}
	return nil
}

// GetEmojiID returns value of EmojiID field.
func (v *VideoSizeEmojiMarkup) GetEmojiID() (value int64) {
	if v == nil {
		return
	}
	return v.EmojiID
}

// GetBackgroundColors returns value of BackgroundColors field.
func (v *VideoSizeEmojiMarkup) GetBackgroundColors() (value []int) {
	if v == nil {
		return
	}
	return v.BackgroundColors
}

// VideoSizeStickerMarkup represents TL type `videoSizeStickerMarkup#da082fe`.
// An animated profile picture¹ based on a sticker².
//
// Links:
//  1. https://core.telegram.org/api/files#animated-profile-pictures
//  2. https://core.telegram.org/api/stickers
//
// See https://core.telegram.org/constructor/videoSizeStickerMarkup for reference.
type VideoSizeStickerMarkup struct {
	// Stickerset
	Stickerset InputStickerSetClass
	// Sticker ID
	StickerID int64
	// 1, 2, 3 or 4 RBG-24 colors used to generate a solid (1), gradient (2) or freeform
	// gradient (3, 4) background, similar to how fill wallpapers¹ are generated. The
	// rotation angle for gradient backgrounds is 0.
	//
	// Links:
	//  1) https://core.telegram.org/api/wallpapers#fill-types
	BackgroundColors []int
}

// VideoSizeStickerMarkupTypeID is TL type id of VideoSizeStickerMarkup.
const VideoSizeStickerMarkupTypeID = 0xda082fe

// construct implements constructor of VideoSizeClass.
func (v VideoSizeStickerMarkup) construct() VideoSizeClass { return &v }

// Ensuring interfaces in compile-time for VideoSizeStickerMarkup.
var (
	_ bin.Encoder     = &VideoSizeStickerMarkup{}
	_ bin.Decoder     = &VideoSizeStickerMarkup{}
	_ bin.BareEncoder = &VideoSizeStickerMarkup{}
	_ bin.BareDecoder = &VideoSizeStickerMarkup{}

	_ VideoSizeClass = &VideoSizeStickerMarkup{}
)

func (v *VideoSizeStickerMarkup) Zero() bool {
	if v == nil {
		return true
	}
	if !(v.Stickerset == nil) {
		return false
	}
	if !(v.StickerID == 0) {
		return false
	}
	if !(v.BackgroundColors == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (v *VideoSizeStickerMarkup) String() string {
	if v == nil {
		return "VideoSizeStickerMarkup(nil)"
	}
	type Alias VideoSizeStickerMarkup
	return fmt.Sprintf("VideoSizeStickerMarkup%+v", Alias(*v))
}

// FillFrom fills VideoSizeStickerMarkup from given interface.
func (v *VideoSizeStickerMarkup) FillFrom(from interface {
	GetStickerset() (value InputStickerSetClass)
	GetStickerID() (value int64)
	GetBackgroundColors() (value []int)
}) {
	v.Stickerset = from.GetStickerset()
	v.StickerID = from.GetStickerID()
	v.BackgroundColors = from.GetBackgroundColors()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*VideoSizeStickerMarkup) TypeID() uint32 {
	return VideoSizeStickerMarkupTypeID
}

// TypeName returns name of type in TL schema.
func (*VideoSizeStickerMarkup) TypeName() string {
	return "videoSizeStickerMarkup"
}

// TypeInfo returns info about TL type.
func (v *VideoSizeStickerMarkup) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "videoSizeStickerMarkup",
		ID:   VideoSizeStickerMarkupTypeID,
	}
	if v == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Stickerset",
			SchemaName: "stickerset",
		},
		{
			Name:       "StickerID",
			SchemaName: "sticker_id",
		},
		{
			Name:       "BackgroundColors",
			SchemaName: "background_colors",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (v *VideoSizeStickerMarkup) Encode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoSizeStickerMarkup#da082fe as nil")
	}
	b.PutID(VideoSizeStickerMarkupTypeID)
	return v.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (v *VideoSizeStickerMarkup) EncodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't encode videoSizeStickerMarkup#da082fe as nil")
	}
	if v.Stickerset == nil {
		return fmt.Errorf("unable to encode videoSizeStickerMarkup#da082fe: field stickerset is nil")
	}
	if err := v.Stickerset.Encode(b); err != nil {
		return fmt.Errorf("unable to encode videoSizeStickerMarkup#da082fe: field stickerset: %w", err)
	}
	b.PutLong(v.StickerID)
	b.PutVectorHeader(len(v.BackgroundColors))
	for _, v := range v.BackgroundColors {
		b.PutInt(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (v *VideoSizeStickerMarkup) Decode(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoSizeStickerMarkup#da082fe to nil")
	}
	if err := b.ConsumeID(VideoSizeStickerMarkupTypeID); err != nil {
		return fmt.Errorf("unable to decode videoSizeStickerMarkup#da082fe: %w", err)
	}
	return v.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (v *VideoSizeStickerMarkup) DecodeBare(b *bin.Buffer) error {
	if v == nil {
		return fmt.Errorf("can't decode videoSizeStickerMarkup#da082fe to nil")
	}
	{
		value, err := DecodeInputStickerSet(b)
		if err != nil {
			return fmt.Errorf("unable to decode videoSizeStickerMarkup#da082fe: field stickerset: %w", err)
		}
		v.Stickerset = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode videoSizeStickerMarkup#da082fe: field sticker_id: %w", err)
		}
		v.StickerID = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode videoSizeStickerMarkup#da082fe: field background_colors: %w", err)
		}

		if headerLen > 0 {
			v.BackgroundColors = make([]int, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int()
			if err != nil {
				return fmt.Errorf("unable to decode videoSizeStickerMarkup#da082fe: field background_colors: %w", err)
			}
			v.BackgroundColors = append(v.BackgroundColors, value)
		}
	}
	return nil
}

// GetStickerset returns value of Stickerset field.
func (v *VideoSizeStickerMarkup) GetStickerset() (value InputStickerSetClass) {
	if v == nil {
		return
	}
	return v.Stickerset
}

// GetStickerID returns value of StickerID field.
func (v *VideoSizeStickerMarkup) GetStickerID() (value int64) {
	if v == nil {
		return
	}
	return v.StickerID
}

// GetBackgroundColors returns value of BackgroundColors field.
func (v *VideoSizeStickerMarkup) GetBackgroundColors() (value []int) {
	if v == nil {
		return
	}
	return v.BackgroundColors
}

// VideoSizeClassName is schema name of VideoSizeClass.
const VideoSizeClassName = "VideoSize"

// VideoSizeClass represents VideoSize generic type.
//
// See https://core.telegram.org/type/VideoSize for reference.
//
// Example:
//
//	g, err := tg.DecodeVideoSize(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.VideoSize: // videoSize#de33b094
//	case *tg.VideoSizeEmojiMarkup: // videoSizeEmojiMarkup#f85c413c
//	case *tg.VideoSizeStickerMarkup: // videoSizeStickerMarkup#da082fe
//	default: panic(v)
//	}
type VideoSizeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() VideoSizeClass

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

// DecodeVideoSize implements binary de-serialization for VideoSizeClass.
func DecodeVideoSize(buf *bin.Buffer) (VideoSizeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case VideoSizeTypeID:
		// Decoding videoSize#de33b094.
		v := VideoSize{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode VideoSizeClass: %w", err)
		}
		return &v, nil
	case VideoSizeEmojiMarkupTypeID:
		// Decoding videoSizeEmojiMarkup#f85c413c.
		v := VideoSizeEmojiMarkup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode VideoSizeClass: %w", err)
		}
		return &v, nil
	case VideoSizeStickerMarkupTypeID:
		// Decoding videoSizeStickerMarkup#da082fe.
		v := VideoSizeStickerMarkup{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode VideoSizeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode VideoSizeClass: %w", bin.NewUnexpectedID(id))
	}
}

// VideoSize boxes the VideoSizeClass providing a helper.
type VideoSizeBox struct {
	VideoSize VideoSizeClass
}

// Decode implements bin.Decoder for VideoSizeBox.
func (b *VideoSizeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode VideoSizeBox to nil")
	}
	v, err := DecodeVideoSize(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.VideoSize = v
	return nil
}

// Encode implements bin.Encode for VideoSizeBox.
func (b *VideoSizeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.VideoSize == nil {
		return fmt.Errorf("unable to encode VideoSizeClass as nil")
	}
	return b.VideoSize.Encode(buf)
}
