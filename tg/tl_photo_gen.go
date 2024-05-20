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

// PhotoEmpty represents TL type `photoEmpty#2331b22d`.
// Empty constructor, non-existent photo
//
// See https://core.telegram.org/constructor/photoEmpty for reference.
type PhotoEmpty struct {
	// Photo identifier
	ID int64
}

// PhotoEmptyTypeID is TL type id of PhotoEmpty.
const PhotoEmptyTypeID = 0x2331b22d

// construct implements constructor of PhotoClass.
func (p PhotoEmpty) construct() PhotoClass { return &p }

// Ensuring interfaces in compile-time for PhotoEmpty.
var (
	_ bin.Encoder     = &PhotoEmpty{}
	_ bin.Decoder     = &PhotoEmpty{}
	_ bin.BareEncoder = &PhotoEmpty{}
	_ bin.BareDecoder = &PhotoEmpty{}

	_ PhotoClass = &PhotoEmpty{}
)

func (p *PhotoEmpty) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotoEmpty) String() string {
	if p == nil {
		return "PhotoEmpty(nil)"
	}
	type Alias PhotoEmpty
	return fmt.Sprintf("PhotoEmpty%+v", Alias(*p))
}

// FillFrom fills PhotoEmpty from given interface.
func (p *PhotoEmpty) FillFrom(from interface {
	GetID() (value int64)
}) {
	p.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotoEmpty) TypeID() uint32 {
	return PhotoEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotoEmpty) TypeName() string {
	return "photoEmpty"
}

// TypeInfo returns info about TL type.
func (p *PhotoEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photoEmpty",
		ID:   PhotoEmptyTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotoEmpty) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoEmpty#2331b22d as nil")
	}
	b.PutID(PhotoEmptyTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotoEmpty) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photoEmpty#2331b22d as nil")
	}
	b.PutLong(p.ID)
	return nil
}

// Decode implements bin.Decoder.
func (p *PhotoEmpty) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoEmpty#2331b22d to nil")
	}
	if err := b.ConsumeID(PhotoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode photoEmpty#2331b22d: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotoEmpty) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photoEmpty#2331b22d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode photoEmpty#2331b22d: field id: %w", err)
		}
		p.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (p *PhotoEmpty) GetID() (value int64) {
	if p == nil {
		return
	}
	return p.ID
}

// Photo represents TL type `photo#fb197a65`.
// Photo
//
// See https://core.telegram.org/constructor/photo for reference.
type Photo struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the photo has mask stickers attached to it
	HasStickers bool
	// ID
	ID int64
	// Access hash
	AccessHash int64
	// file reference¹
	//
	// Links:
	//  1) https://core.telegram.org/api/file_reference
	FileReference []byte
	// Date of upload
	Date int
	// Available sizes for download
	Sizes []PhotoSizeClass
	// For animated profiles¹, the MPEG4 videos
	//
	// Links:
	//  1) https://core.telegram.org/api/files#animated-profile-pictures
	//
	// Use SetVideoSizes and GetVideoSizes helpers.
	VideoSizes []VideoSizeClass
	// DC ID to use for download
	DCID int
}

// PhotoTypeID is TL type id of Photo.
const PhotoTypeID = 0xfb197a65

// construct implements constructor of PhotoClass.
func (p Photo) construct() PhotoClass { return &p }

// Ensuring interfaces in compile-time for Photo.
var (
	_ bin.Encoder     = &Photo{}
	_ bin.Decoder     = &Photo{}
	_ bin.BareEncoder = &Photo{}
	_ bin.BareDecoder = &Photo{}

	_ PhotoClass = &Photo{}
)

func (p *Photo) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.HasStickers == false) {
		return false
	}
	if !(p.ID == 0) {
		return false
	}
	if !(p.AccessHash == 0) {
		return false
	}
	if !(p.FileReference == nil) {
		return false
	}
	if !(p.Date == 0) {
		return false
	}
	if !(p.Sizes == nil) {
		return false
	}
	if !(p.VideoSizes == nil) {
		return false
	}
	if !(p.DCID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *Photo) String() string {
	if p == nil {
		return "Photo(nil)"
	}
	type Alias Photo
	return fmt.Sprintf("Photo%+v", Alias(*p))
}

// FillFrom fills Photo from given interface.
func (p *Photo) FillFrom(from interface {
	GetHasStickers() (value bool)
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetFileReference() (value []byte)
	GetDate() (value int)
	GetSizes() (value []PhotoSizeClass)
	GetVideoSizes() (value []VideoSizeClass, ok bool)
	GetDCID() (value int)
}) {
	p.HasStickers = from.GetHasStickers()
	p.ID = from.GetID()
	p.AccessHash = from.GetAccessHash()
	p.FileReference = from.GetFileReference()
	p.Date = from.GetDate()
	p.Sizes = from.GetSizes()
	if val, ok := from.GetVideoSizes(); ok {
		p.VideoSizes = val
	}

	p.DCID = from.GetDCID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Photo) TypeID() uint32 {
	return PhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*Photo) TypeName() string {
	return "photo"
}

// TypeInfo returns info about TL type.
func (p *Photo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photo",
		ID:   PhotoTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasStickers",
			SchemaName: "has_stickers",
			Null:       !p.Flags.Has(0),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
		{
			Name:       "FileReference",
			SchemaName: "file_reference",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "Sizes",
			SchemaName: "sizes",
		},
		{
			Name:       "VideoSizes",
			SchemaName: "video_sizes",
			Null:       !p.Flags.Has(1),
		},
		{
			Name:       "DCID",
			SchemaName: "dc_id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (p *Photo) SetFlags() {
	if !(p.HasStickers == false) {
		p.Flags.Set(0)
	}
	if !(p.VideoSizes == nil) {
		p.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (p *Photo) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photo#fb197a65 as nil")
	}
	b.PutID(PhotoTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *Photo) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photo#fb197a65 as nil")
	}
	p.SetFlags()
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photo#fb197a65: field flags: %w", err)
	}
	b.PutLong(p.ID)
	b.PutLong(p.AccessHash)
	b.PutBytes(p.FileReference)
	b.PutInt(p.Date)
	b.PutVectorHeader(len(p.Sizes))
	for idx, v := range p.Sizes {
		if v == nil {
			return fmt.Errorf("unable to encode photo#fb197a65: field sizes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photo#fb197a65: field sizes element with index %d: %w", idx, err)
		}
	}
	if p.Flags.Has(1) {
		b.PutVectorHeader(len(p.VideoSizes))
		for idx, v := range p.VideoSizes {
			if v == nil {
				return fmt.Errorf("unable to encode photo#fb197a65: field video_sizes element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode photo#fb197a65: field video_sizes element with index %d: %w", idx, err)
			}
		}
	}
	b.PutInt(p.DCID)
	return nil
}

// Decode implements bin.Decoder.
func (p *Photo) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photo#fb197a65 to nil")
	}
	if err := b.ConsumeID(PhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode photo#fb197a65: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *Photo) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photo#fb197a65 to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field flags: %w", err)
		}
	}
	p.HasStickers = p.Flags.Has(0)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field id: %w", err)
		}
		p.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field access_hash: %w", err)
		}
		p.AccessHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field file_reference: %w", err)
		}
		p.FileReference = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field date: %w", err)
		}
		p.Date = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field sizes: %w", err)
		}

		if headerLen > 0 {
			p.Sizes = make([]PhotoSizeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhotoSize(b)
			if err != nil {
				return fmt.Errorf("unable to decode photo#fb197a65: field sizes: %w", err)
			}
			p.Sizes = append(p.Sizes, value)
		}
	}
	if p.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field video_sizes: %w", err)
		}

		if headerLen > 0 {
			p.VideoSizes = make([]VideoSizeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeVideoSize(b)
			if err != nil {
				return fmt.Errorf("unable to decode photo#fb197a65: field video_sizes: %w", err)
			}
			p.VideoSizes = append(p.VideoSizes, value)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photo#fb197a65: field dc_id: %w", err)
		}
		p.DCID = value
	}
	return nil
}

// SetHasStickers sets value of HasStickers conditional field.
func (p *Photo) SetHasStickers(value bool) {
	if value {
		p.Flags.Set(0)
		p.HasStickers = true
	} else {
		p.Flags.Unset(0)
		p.HasStickers = false
	}
}

// GetHasStickers returns value of HasStickers conditional field.
func (p *Photo) GetHasStickers() (value bool) {
	if p == nil {
		return
	}
	return p.Flags.Has(0)
}

// GetID returns value of ID field.
func (p *Photo) GetID() (value int64) {
	if p == nil {
		return
	}
	return p.ID
}

// GetAccessHash returns value of AccessHash field.
func (p *Photo) GetAccessHash() (value int64) {
	if p == nil {
		return
	}
	return p.AccessHash
}

// GetFileReference returns value of FileReference field.
func (p *Photo) GetFileReference() (value []byte) {
	if p == nil {
		return
	}
	return p.FileReference
}

// GetDate returns value of Date field.
func (p *Photo) GetDate() (value int) {
	if p == nil {
		return
	}
	return p.Date
}

// GetSizes returns value of Sizes field.
func (p *Photo) GetSizes() (value []PhotoSizeClass) {
	if p == nil {
		return
	}
	return p.Sizes
}

// SetVideoSizes sets value of VideoSizes conditional field.
func (p *Photo) SetVideoSizes(value []VideoSizeClass) {
	p.Flags.Set(1)
	p.VideoSizes = value
}

// GetVideoSizes returns value of VideoSizes conditional field and
// boolean which is true if field was set.
func (p *Photo) GetVideoSizes() (value []VideoSizeClass, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.VideoSizes, true
}

// GetDCID returns value of DCID field.
func (p *Photo) GetDCID() (value int) {
	if p == nil {
		return
	}
	return p.DCID
}

// MapSizes returns field Sizes wrapped in PhotoSizeClassArray helper.
func (p *Photo) MapSizes() (value PhotoSizeClassArray) {
	return PhotoSizeClassArray(p.Sizes)
}

// MapVideoSizes returns field VideoSizes wrapped in VideoSizeClassArray helper.
func (p *Photo) MapVideoSizes() (value VideoSizeClassArray, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return VideoSizeClassArray(p.VideoSizes), true
}

// PhotoClassName is schema name of PhotoClass.
const PhotoClassName = "Photo"

// PhotoClass represents Photo generic type.
//
// See https://core.telegram.org/type/Photo for reference.
//
// Example:
//
//	g, err := tg.DecodePhoto(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.PhotoEmpty: // photoEmpty#2331b22d
//	case *tg.Photo: // photo#fb197a65
//	default: panic(v)
//	}
type PhotoClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PhotoClass

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

	// Photo identifier
	GetID() (value int64)

	// AsNotEmpty tries to map PhotoClass to Photo.
	AsNotEmpty() (*Photo, bool)
}

// AsInput tries to map Photo to InputPhoto.
func (p *Photo) AsInput() *InputPhoto {
	value := new(InputPhoto)
	value.ID = p.GetID()
	value.AccessHash = p.GetAccessHash()
	value.FileReference = p.GetFileReference()

	return value
}

// AsNotEmpty tries to map PhotoEmpty to Photo.
func (p *PhotoEmpty) AsNotEmpty() (*Photo, bool) {
	return nil, false
}

// AsNotEmpty tries to map Photo to Photo.
func (p *Photo) AsNotEmpty() (*Photo, bool) {
	return p, true
}

// DecodePhoto implements binary de-serialization for PhotoClass.
func DecodePhoto(buf *bin.Buffer) (PhotoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhotoEmptyTypeID:
		// Decoding photoEmpty#2331b22d.
		v := PhotoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoClass: %w", err)
		}
		return &v, nil
	case PhotoTypeID:
		// Decoding photo#fb197a65.
		v := Photo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhotoClass: %w", bin.NewUnexpectedID(id))
	}
}

// Photo boxes the PhotoClass providing a helper.
type PhotoBox struct {
	Photo PhotoClass
}

// Decode implements bin.Decoder for PhotoBox.
func (b *PhotoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhotoBox to nil")
	}
	v, err := DecodePhoto(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Photo = v
	return nil
}

// Encode implements bin.Encode for PhotoBox.
func (b *PhotoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Photo == nil {
		return fmt.Errorf("unable to encode PhotoClass as nil")
	}
	return b.Photo.Encode(buf)
}
