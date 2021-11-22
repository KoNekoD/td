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

// InputWebDocument represents TL type `inputWebDocument#9bed434d`.
// The document
//
// See https://core.telegram.org/constructor/inputWebDocument for reference.
type InputWebDocument struct {
	// Remote document URL to be downloaded using the appropriate method¹
	//
	// Links:
	//  1) https://core.telegram.org/api/files
	URL string
	// Remote file size
	Size int
	// Mime type
	MimeType string
	// Attributes for media types
	Attributes []DocumentAttributeClass
}

// InputWebDocumentTypeID is TL type id of InputWebDocument.
const InputWebDocumentTypeID = 0x9bed434d

// Ensuring interfaces in compile-time for InputWebDocument.
var (
	_ bin.Encoder     = &InputWebDocument{}
	_ bin.Decoder     = &InputWebDocument{}
	_ bin.BareEncoder = &InputWebDocument{}
	_ bin.BareDecoder = &InputWebDocument{}
)

func (i *InputWebDocument) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.URL == "") {
		return false
	}
	if !(i.Size == 0) {
		return false
	}
	if !(i.MimeType == "") {
		return false
	}
	if !(i.Attributes == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputWebDocument) String() string {
	if i == nil {
		return "InputWebDocument(nil)"
	}
	type Alias InputWebDocument
	return fmt.Sprintf("InputWebDocument%+v", Alias(*i))
}

// FillFrom fills InputWebDocument from given interface.
func (i *InputWebDocument) FillFrom(from interface {
	GetURL() (value string)
	GetSize() (value int)
	GetMimeType() (value string)
	GetAttributes() (value []DocumentAttributeClass)
}) {
	i.URL = from.GetURL()
	i.Size = from.GetSize()
	i.MimeType = from.GetMimeType()
	i.Attributes = from.GetAttributes()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputWebDocument) TypeID() uint32 {
	return InputWebDocumentTypeID
}

// TypeName returns name of type in TL schema.
func (*InputWebDocument) TypeName() string {
	return "inputWebDocument"
}

// TypeInfo returns info about TL type.
func (i *InputWebDocument) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputWebDocument",
		ID:   InputWebDocumentTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Size",
			SchemaName: "size",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "Attributes",
			SchemaName: "attributes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputWebDocument) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWebDocument#9bed434d as nil")
	}
	b.PutID(InputWebDocumentTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputWebDocument) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputWebDocument#9bed434d as nil")
	}
	b.PutString(i.URL)
	b.PutInt(i.Size)
	b.PutString(i.MimeType)
	b.PutVectorHeader(len(i.Attributes))
	for idx, v := range i.Attributes {
		if v == nil {
			return fmt.Errorf("unable to encode inputWebDocument#9bed434d: field attributes element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode inputWebDocument#9bed434d: field attributes element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputWebDocument) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWebDocument#9bed434d to nil")
	}
	if err := b.ConsumeID(InputWebDocumentTypeID); err != nil {
		return fmt.Errorf("unable to decode inputWebDocument#9bed434d: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputWebDocument) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputWebDocument#9bed434d to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field url: %w", err)
		}
		i.URL = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field size: %w", err)
		}
		i.Size = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field mime_type: %w", err)
		}
		i.MimeType = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field attributes: %w", err)
		}

		if headerLen > 0 {
			i.Attributes = make([]DocumentAttributeClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeDocumentAttribute(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputWebDocument#9bed434d: field attributes: %w", err)
			}
			i.Attributes = append(i.Attributes, value)
		}
	}
	return nil
}

// GetURL returns value of URL field.
func (i *InputWebDocument) GetURL() (value string) {
	return i.URL
}

// GetSize returns value of Size field.
func (i *InputWebDocument) GetSize() (value int) {
	return i.Size
}

// GetMimeType returns value of MimeType field.
func (i *InputWebDocument) GetMimeType() (value string) {
	return i.MimeType
}

// GetAttributes returns value of Attributes field.
func (i *InputWebDocument) GetAttributes() (value []DocumentAttributeClass) {
	return i.Attributes
}

// MapAttributes returns field Attributes wrapped in DocumentAttributeClassArray helper.
func (i *InputWebDocument) MapAttributes() (value DocumentAttributeClassArray) {
	return DocumentAttributeClassArray(i.Attributes)
}
