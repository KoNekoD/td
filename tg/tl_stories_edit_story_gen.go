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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// StoriesEditStoryRequest represents TL type `stories.editStory#2aae7a41`.
//
// See https://core.telegram.org/method/stories.editStory for reference.
type StoriesEditStoryRequest struct {
	// Flags field of StoriesEditStoryRequest.
	Flags bin.Fields
	// ID field of StoriesEditStoryRequest.
	ID int
	// Media field of StoriesEditStoryRequest.
	//
	// Use SetMedia and GetMedia helpers.
	Media InputMediaClass
	// Caption field of StoriesEditStoryRequest.
	//
	// Use SetCaption and GetCaption helpers.
	Caption string
	// Entities field of StoriesEditStoryRequest.
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// PrivacyRules field of StoriesEditStoryRequest.
	//
	// Use SetPrivacyRules and GetPrivacyRules helpers.
	PrivacyRules []InputPrivacyRuleClass
}

// StoriesEditStoryRequestTypeID is TL type id of StoriesEditStoryRequest.
const StoriesEditStoryRequestTypeID = 0x2aae7a41

// Ensuring interfaces in compile-time for StoriesEditStoryRequest.
var (
	_ bin.Encoder     = &StoriesEditStoryRequest{}
	_ bin.Decoder     = &StoriesEditStoryRequest{}
	_ bin.BareEncoder = &StoriesEditStoryRequest{}
	_ bin.BareDecoder = &StoriesEditStoryRequest{}
)

func (e *StoriesEditStoryRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.ID == 0) {
		return false
	}
	if !(e.Media == nil) {
		return false
	}
	if !(e.Caption == "") {
		return false
	}
	if !(e.Entities == nil) {
		return false
	}
	if !(e.PrivacyRules == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *StoriesEditStoryRequest) String() string {
	if e == nil {
		return "StoriesEditStoryRequest(nil)"
	}
	type Alias StoriesEditStoryRequest
	return fmt.Sprintf("StoriesEditStoryRequest%+v", Alias(*e))
}

// FillFrom fills StoriesEditStoryRequest from given interface.
func (e *StoriesEditStoryRequest) FillFrom(from interface {
	GetID() (value int)
	GetMedia() (value InputMediaClass, ok bool)
	GetCaption() (value string, ok bool)
	GetEntities() (value []MessageEntityClass, ok bool)
	GetPrivacyRules() (value []InputPrivacyRuleClass, ok bool)
}) {
	e.ID = from.GetID()
	if val, ok := from.GetMedia(); ok {
		e.Media = val
	}

	if val, ok := from.GetCaption(); ok {
		e.Caption = val
	}

	if val, ok := from.GetEntities(); ok {
		e.Entities = val
	}

	if val, ok := from.GetPrivacyRules(); ok {
		e.PrivacyRules = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoriesEditStoryRequest) TypeID() uint32 {
	return StoriesEditStoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StoriesEditStoryRequest) TypeName() string {
	return "stories.editStory"
}

// TypeInfo returns info about TL type.
func (e *StoriesEditStoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "stories.editStory",
		ID:   StoriesEditStoryRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Media",
			SchemaName: "media",
			Null:       !e.Flags.Has(0),
		},
		{
			Name:       "Caption",
			SchemaName: "caption",
			Null:       !e.Flags.Has(1),
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !e.Flags.Has(1),
		},
		{
			Name:       "PrivacyRules",
			SchemaName: "privacy_rules",
			Null:       !e.Flags.Has(2),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (e *StoriesEditStoryRequest) SetFlags() {
	if !(e.Media == nil) {
		e.Flags.Set(0)
	}
	if !(e.Caption == "") {
		e.Flags.Set(1)
	}
	if !(e.Entities == nil) {
		e.Flags.Set(1)
	}
	if !(e.PrivacyRules == nil) {
		e.Flags.Set(2)
	}
}

// Encode implements bin.Encoder.
func (e *StoriesEditStoryRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode stories.editStory#2aae7a41 as nil")
	}
	b.PutID(StoriesEditStoryRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *StoriesEditStoryRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode stories.editStory#2aae7a41 as nil")
	}
	e.SetFlags()
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode stories.editStory#2aae7a41: field flags: %w", err)
	}
	b.PutInt(e.ID)
	if e.Flags.Has(0) {
		if e.Media == nil {
			return fmt.Errorf("unable to encode stories.editStory#2aae7a41: field media is nil")
		}
		if err := e.Media.Encode(b); err != nil {
			return fmt.Errorf("unable to encode stories.editStory#2aae7a41: field media: %w", err)
		}
	}
	if e.Flags.Has(1) {
		b.PutString(e.Caption)
	}
	if e.Flags.Has(1) {
		b.PutVectorHeader(len(e.Entities))
		for idx, v := range e.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode stories.editStory#2aae7a41: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode stories.editStory#2aae7a41: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if e.Flags.Has(2) {
		b.PutVectorHeader(len(e.PrivacyRules))
		for idx, v := range e.PrivacyRules {
			if v == nil {
				return fmt.Errorf("unable to encode stories.editStory#2aae7a41: field privacy_rules element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode stories.editStory#2aae7a41: field privacy_rules element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *StoriesEditStoryRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode stories.editStory#2aae7a41 to nil")
	}
	if err := b.ConsumeID(StoriesEditStoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode stories.editStory#2aae7a41: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *StoriesEditStoryRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode stories.editStory#2aae7a41 to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode stories.editStory#2aae7a41: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode stories.editStory#2aae7a41: field id: %w", err)
		}
		e.ID = value
	}
	if e.Flags.Has(0) {
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode stories.editStory#2aae7a41: field media: %w", err)
		}
		e.Media = value
	}
	if e.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode stories.editStory#2aae7a41: field caption: %w", err)
		}
		e.Caption = value
	}
	if e.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.editStory#2aae7a41: field entities: %w", err)
		}

		if headerLen > 0 {
			e.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.editStory#2aae7a41: field entities: %w", err)
			}
			e.Entities = append(e.Entities, value)
		}
	}
	if e.Flags.Has(2) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode stories.editStory#2aae7a41: field privacy_rules: %w", err)
		}

		if headerLen > 0 {
			e.PrivacyRules = make([]InputPrivacyRuleClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPrivacyRule(b)
			if err != nil {
				return fmt.Errorf("unable to decode stories.editStory#2aae7a41: field privacy_rules: %w", err)
			}
			e.PrivacyRules = append(e.PrivacyRules, value)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (e *StoriesEditStoryRequest) GetID() (value int) {
	if e == nil {
		return
	}
	return e.ID
}

// SetMedia sets value of Media conditional field.
func (e *StoriesEditStoryRequest) SetMedia(value InputMediaClass) {
	e.Flags.Set(0)
	e.Media = value
}

// GetMedia returns value of Media conditional field and
// boolean which is true if field was set.
func (e *StoriesEditStoryRequest) GetMedia() (value InputMediaClass, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(0) {
		return value, false
	}
	return e.Media, true
}

// SetCaption sets value of Caption conditional field.
func (e *StoriesEditStoryRequest) SetCaption(value string) {
	e.Flags.Set(1)
	e.Caption = value
}

// GetCaption returns value of Caption conditional field and
// boolean which is true if field was set.
func (e *StoriesEditStoryRequest) GetCaption() (value string, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(1) {
		return value, false
	}
	return e.Caption, true
}

// SetEntities sets value of Entities conditional field.
func (e *StoriesEditStoryRequest) SetEntities(value []MessageEntityClass) {
	e.Flags.Set(1)
	e.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (e *StoriesEditStoryRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(1) {
		return value, false
	}
	return e.Entities, true
}

// SetPrivacyRules sets value of PrivacyRules conditional field.
func (e *StoriesEditStoryRequest) SetPrivacyRules(value []InputPrivacyRuleClass) {
	e.Flags.Set(2)
	e.PrivacyRules = value
}

// GetPrivacyRules returns value of PrivacyRules conditional field and
// boolean which is true if field was set.
func (e *StoriesEditStoryRequest) GetPrivacyRules() (value []InputPrivacyRuleClass, ok bool) {
	if e == nil {
		return
	}
	if !e.Flags.Has(2) {
		return value, false
	}
	return e.PrivacyRules, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (e *StoriesEditStoryRequest) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !e.Flags.Has(1) {
		return value, false
	}
	return MessageEntityClassArray(e.Entities), true
}

// MapPrivacyRules returns field PrivacyRules wrapped in InputPrivacyRuleClassArray helper.
func (e *StoriesEditStoryRequest) MapPrivacyRules() (value InputPrivacyRuleClassArray, ok bool) {
	if !e.Flags.Has(2) {
		return value, false
	}
	return InputPrivacyRuleClassArray(e.PrivacyRules), true
}

// StoriesEditStory invokes method stories.editStory#2aae7a41 returning error if any.
//
// See https://core.telegram.org/method/stories.editStory for reference.
func (c *Client) StoriesEditStory(ctx context.Context, request *StoriesEditStoryRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}