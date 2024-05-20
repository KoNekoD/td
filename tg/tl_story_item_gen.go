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

// StoryItemDeleted represents TL type `storyItemDeleted#51e6ee4f`.
// Represents a previously active story, that was deleted
//
// See https://core.telegram.org/constructor/storyItemDeleted for reference.
type StoryItemDeleted struct {
	// Story ID
	ID int
}

// StoryItemDeletedTypeID is TL type id of StoryItemDeleted.
const StoryItemDeletedTypeID = 0x51e6ee4f

// construct implements constructor of StoryItemClass.
func (s StoryItemDeleted) construct() StoryItemClass { return &s }

// Ensuring interfaces in compile-time for StoryItemDeleted.
var (
	_ bin.Encoder     = &StoryItemDeleted{}
	_ bin.Decoder     = &StoryItemDeleted{}
	_ bin.BareEncoder = &StoryItemDeleted{}
	_ bin.BareDecoder = &StoryItemDeleted{}

	_ StoryItemClass = &StoryItemDeleted{}
)

func (s *StoryItemDeleted) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryItemDeleted) String() string {
	if s == nil {
		return "StoryItemDeleted(nil)"
	}
	type Alias StoryItemDeleted
	return fmt.Sprintf("StoryItemDeleted%+v", Alias(*s))
}

// FillFrom fills StoryItemDeleted from given interface.
func (s *StoryItemDeleted) FillFrom(from interface {
	GetID() (value int)
}) {
	s.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryItemDeleted) TypeID() uint32 {
	return StoryItemDeletedTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryItemDeleted) TypeName() string {
	return "storyItemDeleted"
}

// TypeInfo returns info about TL type.
func (s *StoryItemDeleted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyItemDeleted",
		ID:   StoryItemDeletedTypeID,
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
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryItemDeleted) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyItemDeleted#51e6ee4f as nil")
	}
	b.PutID(StoryItemDeletedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryItemDeleted) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyItemDeleted#51e6ee4f as nil")
	}
	b.PutInt(s.ID)
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryItemDeleted) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyItemDeleted#51e6ee4f to nil")
	}
	if err := b.ConsumeID(StoryItemDeletedTypeID); err != nil {
		return fmt.Errorf("unable to decode storyItemDeleted#51e6ee4f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryItemDeleted) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyItemDeleted#51e6ee4f to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyItemDeleted#51e6ee4f: field id: %w", err)
		}
		s.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (s *StoryItemDeleted) GetID() (value int) {
	if s == nil {
		return
	}
	return s.ID
}

// StoryItemSkipped represents TL type `storyItemSkipped#ffadc913`.
// Represents an active story, whose full information was omitted for space and
// performance reasons; use stories.getStoriesByID¹ to fetch full info about the skipped
// story when and if needed.
//
// Links:
//  1. https://core.telegram.org/method/stories.getStoriesByID
//
// See https://core.telegram.org/constructor/storyItemSkipped for reference.
type StoryItemSkipped struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this story can only be viewed by our close friends, see here »¹ for more
	// info
	//
	// Links:
	//  1) https://core.telegram.org/api/privacy
	CloseFriends bool
	// Story ID
	ID int
	// When was the story posted.
	Date int
	// When does the story expire.
	ExpireDate int
}

// StoryItemSkippedTypeID is TL type id of StoryItemSkipped.
const StoryItemSkippedTypeID = 0xffadc913

// construct implements constructor of StoryItemClass.
func (s StoryItemSkipped) construct() StoryItemClass { return &s }

// Ensuring interfaces in compile-time for StoryItemSkipped.
var (
	_ bin.Encoder     = &StoryItemSkipped{}
	_ bin.Decoder     = &StoryItemSkipped{}
	_ bin.BareEncoder = &StoryItemSkipped{}
	_ bin.BareDecoder = &StoryItemSkipped{}

	_ StoryItemClass = &StoryItemSkipped{}
)

func (s *StoryItemSkipped) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.CloseFriends == false) {
		return false
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.ExpireDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryItemSkipped) String() string {
	if s == nil {
		return "StoryItemSkipped(nil)"
	}
	type Alias StoryItemSkipped
	return fmt.Sprintf("StoryItemSkipped%+v", Alias(*s))
}

// FillFrom fills StoryItemSkipped from given interface.
func (s *StoryItemSkipped) FillFrom(from interface {
	GetCloseFriends() (value bool)
	GetID() (value int)
	GetDate() (value int)
	GetExpireDate() (value int)
}) {
	s.CloseFriends = from.GetCloseFriends()
	s.ID = from.GetID()
	s.Date = from.GetDate()
	s.ExpireDate = from.GetExpireDate()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryItemSkipped) TypeID() uint32 {
	return StoryItemSkippedTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryItemSkipped) TypeName() string {
	return "storyItemSkipped"
}

// TypeInfo returns info about TL type.
func (s *StoryItemSkipped) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyItemSkipped",
		ID:   StoryItemSkippedTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "CloseFriends",
			SchemaName: "close_friends",
			Null:       !s.Flags.Has(8),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "ExpireDate",
			SchemaName: "expire_date",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoryItemSkipped) SetFlags() {
	if !(s.CloseFriends == false) {
		s.Flags.Set(8)
	}
}

// Encode implements bin.Encoder.
func (s *StoryItemSkipped) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyItemSkipped#ffadc913 as nil")
	}
	b.PutID(StoryItemSkippedTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryItemSkipped) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyItemSkipped#ffadc913 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyItemSkipped#ffadc913: field flags: %w", err)
	}
	b.PutInt(s.ID)
	b.PutInt(s.Date)
	b.PutInt(s.ExpireDate)
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryItemSkipped) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyItemSkipped#ffadc913 to nil")
	}
	if err := b.ConsumeID(StoryItemSkippedTypeID); err != nil {
		return fmt.Errorf("unable to decode storyItemSkipped#ffadc913: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryItemSkipped) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyItemSkipped#ffadc913 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyItemSkipped#ffadc913: field flags: %w", err)
		}
	}
	s.CloseFriends = s.Flags.Has(8)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyItemSkipped#ffadc913: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyItemSkipped#ffadc913: field date: %w", err)
		}
		s.Date = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyItemSkipped#ffadc913: field expire_date: %w", err)
		}
		s.ExpireDate = value
	}
	return nil
}

// SetCloseFriends sets value of CloseFriends conditional field.
func (s *StoryItemSkipped) SetCloseFriends(value bool) {
	if value {
		s.Flags.Set(8)
		s.CloseFriends = true
	} else {
		s.Flags.Unset(8)
		s.CloseFriends = false
	}
}

// GetCloseFriends returns value of CloseFriends conditional field.
func (s *StoryItemSkipped) GetCloseFriends() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(8)
}

// GetID returns value of ID field.
func (s *StoryItemSkipped) GetID() (value int) {
	if s == nil {
		return
	}
	return s.ID
}

// GetDate returns value of Date field.
func (s *StoryItemSkipped) GetDate() (value int) {
	if s == nil {
		return
	}
	return s.Date
}

// GetExpireDate returns value of ExpireDate field.
func (s *StoryItemSkipped) GetExpireDate() (value int) {
	if s == nil {
		return
	}
	return s.ExpireDate
}

// StoryItem represents TL type `storyItem#79b26a24`.
// Represents a story¹.
//
// Links:
//  1. https://core.telegram.org/api/stories
//
// See https://core.telegram.org/constructor/storyItem for reference.
type StoryItem struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this story is pinned on the user's profile
	Pinned bool
	// Whether this story is public and can be viewed by everyone
	Public bool
	// Whether this story can only be viewed by our close friends, see here »¹ for more
	// info
	//
	// Links:
	//  1) https://core.telegram.org/api/privacy
	CloseFriends bool
	// Full information about this story was omitted for space and performance reasons; use
	// stories.getStoriesByID¹ to fetch full info about this story when and if needed.
	//
	// Links:
	//  1) https://core.telegram.org/method/stories.getStoriesByID
	Min bool
	// Whether this story is protected¹ and thus cannot be forwarded; clients should also
	// prevent users from saving attached media (i.e. videos should only be streamed, photos
	// should be kept in RAM, et cetera).
	//
	// Links:
	//  1) https://telegram.org/blog/protected-content-delete-by-date-and-more
	Noforwards bool
	// Indicates whether the story was edited.
	Edited bool
	// Whether this story can only be viewed by our contacts
	Contacts bool
	// Whether this story can only be viewed by a select list of our contacts
	SelectedContacts bool
	// indicates whether we sent this story.
	Out bool
	// ID of the story.
	ID int
	// When was the story posted.
	Date int
	// FromID field of StoryItem.
	//
	// Use SetFromID and GetFromID helpers.
	FromID PeerClass
	// For reposted stories »¹, contains info about the original story.
	//
	// Links:
	//  1) https://core.telegram.org/api/stories#reposting-stories
	//
	// Use SetFwdFrom and GetFwdFrom helpers.
	FwdFrom StoryFwdHeader
	// When does the story expire.
	ExpireDate int
	// Story caption.
	//
	// Use SetCaption and GetCaption helpers.
	Caption string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Story media.
	Media MessageMediaClass
	// List of media areas, see here »¹ for more info on media areas.
	//
	// Links:
	//  1) https://core.telegram.org/api/stories#media-areas
	//
	// Use SetMediaAreas and GetMediaAreas helpers.
	MediaAreas []MediaAreaClass
	// Privacy rules¹ indicating who can and can't view this story
	//
	// Links:
	//  1) https://core.telegram.org/api/privacy
	//
	// Use SetPrivacy and GetPrivacy helpers.
	Privacy []PrivacyRuleClass
	// View date and reaction information
	//
	// Use SetViews and GetViews helpers.
	Views StoryViews
	// The reaction we sent.
	//
	// Use SetSentReaction and GetSentReaction helpers.
	SentReaction ReactionClass
}

// StoryItemTypeID is TL type id of StoryItem.
const StoryItemTypeID = 0x79b26a24

// construct implements constructor of StoryItemClass.
func (s StoryItem) construct() StoryItemClass { return &s }

// Ensuring interfaces in compile-time for StoryItem.
var (
	_ bin.Encoder     = &StoryItem{}
	_ bin.Decoder     = &StoryItem{}
	_ bin.BareEncoder = &StoryItem{}
	_ bin.BareDecoder = &StoryItem{}

	_ StoryItemClass = &StoryItem{}
)

func (s *StoryItem) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Pinned == false) {
		return false
	}
	if !(s.Public == false) {
		return false
	}
	if !(s.CloseFriends == false) {
		return false
	}
	if !(s.Min == false) {
		return false
	}
	if !(s.Noforwards == false) {
		return false
	}
	if !(s.Edited == false) {
		return false
	}
	if !(s.Contacts == false) {
		return false
	}
	if !(s.SelectedContacts == false) {
		return false
	}
	if !(s.Out == false) {
		return false
	}
	if !(s.ID == 0) {
		return false
	}
	if !(s.Date == 0) {
		return false
	}
	if !(s.FromID == nil) {
		return false
	}
	if !(s.FwdFrom.Zero()) {
		return false
	}
	if !(s.ExpireDate == 0) {
		return false
	}
	if !(s.Caption == "") {
		return false
	}
	if !(s.Entities == nil) {
		return false
	}
	if !(s.Media == nil) {
		return false
	}
	if !(s.MediaAreas == nil) {
		return false
	}
	if !(s.Privacy == nil) {
		return false
	}
	if !(s.Views.Zero()) {
		return false
	}
	if !(s.SentReaction == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryItem) String() string {
	if s == nil {
		return "StoryItem(nil)"
	}
	type Alias StoryItem
	return fmt.Sprintf("StoryItem%+v", Alias(*s))
}

// FillFrom fills StoryItem from given interface.
func (s *StoryItem) FillFrom(from interface {
	GetPinned() (value bool)
	GetPublic() (value bool)
	GetCloseFriends() (value bool)
	GetMin() (value bool)
	GetNoforwards() (value bool)
	GetEdited() (value bool)
	GetContacts() (value bool)
	GetSelectedContacts() (value bool)
	GetOut() (value bool)
	GetID() (value int)
	GetDate() (value int)
	GetFromID() (value PeerClass, ok bool)
	GetFwdFrom() (value StoryFwdHeader, ok bool)
	GetExpireDate() (value int)
	GetCaption() (value string, ok bool)
	GetEntities() (value []MessageEntityClass, ok bool)
	GetMedia() (value MessageMediaClass)
	GetMediaAreas() (value []MediaAreaClass, ok bool)
	GetPrivacy() (value []PrivacyRuleClass, ok bool)
	GetViews() (value StoryViews, ok bool)
	GetSentReaction() (value ReactionClass, ok bool)
}) {
	s.Pinned = from.GetPinned()
	s.Public = from.GetPublic()
	s.CloseFriends = from.GetCloseFriends()
	s.Min = from.GetMin()
	s.Noforwards = from.GetNoforwards()
	s.Edited = from.GetEdited()
	s.Contacts = from.GetContacts()
	s.SelectedContacts = from.GetSelectedContacts()
	s.Out = from.GetOut()
	s.ID = from.GetID()
	s.Date = from.GetDate()
	if val, ok := from.GetFromID(); ok {
		s.FromID = val
	}

	if val, ok := from.GetFwdFrom(); ok {
		s.FwdFrom = val
	}

	s.ExpireDate = from.GetExpireDate()
	if val, ok := from.GetCaption(); ok {
		s.Caption = val
	}

	if val, ok := from.GetEntities(); ok {
		s.Entities = val
	}

	s.Media = from.GetMedia()
	if val, ok := from.GetMediaAreas(); ok {
		s.MediaAreas = val
	}

	if val, ok := from.GetPrivacy(); ok {
		s.Privacy = val
	}

	if val, ok := from.GetViews(); ok {
		s.Views = val
	}

	if val, ok := from.GetSentReaction(); ok {
		s.SentReaction = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryItem) TypeID() uint32 {
	return StoryItemTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryItem) TypeName() string {
	return "storyItem"
}

// TypeInfo returns info about TL type.
func (s *StoryItem) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyItem",
		ID:   StoryItemTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pinned",
			SchemaName: "pinned",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "Public",
			SchemaName: "public",
			Null:       !s.Flags.Has(7),
		},
		{
			Name:       "CloseFriends",
			SchemaName: "close_friends",
			Null:       !s.Flags.Has(8),
		},
		{
			Name:       "Min",
			SchemaName: "min",
			Null:       !s.Flags.Has(9),
		},
		{
			Name:       "Noforwards",
			SchemaName: "noforwards",
			Null:       !s.Flags.Has(10),
		},
		{
			Name:       "Edited",
			SchemaName: "edited",
			Null:       !s.Flags.Has(11),
		},
		{
			Name:       "Contacts",
			SchemaName: "contacts",
			Null:       !s.Flags.Has(12),
		},
		{
			Name:       "SelectedContacts",
			SchemaName: "selected_contacts",
			Null:       !s.Flags.Has(13),
		},
		{
			Name:       "Out",
			SchemaName: "out",
			Null:       !s.Flags.Has(16),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Date",
			SchemaName: "date",
		},
		{
			Name:       "FromID",
			SchemaName: "from_id",
			Null:       !s.Flags.Has(18),
		},
		{
			Name:       "FwdFrom",
			SchemaName: "fwd_from",
			Null:       !s.Flags.Has(17),
		},
		{
			Name:       "ExpireDate",
			SchemaName: "expire_date",
		},
		{
			Name:       "Caption",
			SchemaName: "caption",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "Media",
			SchemaName: "media",
		},
		{
			Name:       "MediaAreas",
			SchemaName: "media_areas",
			Null:       !s.Flags.Has(14),
		},
		{
			Name:       "Privacy",
			SchemaName: "privacy",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Views",
			SchemaName: "views",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "SentReaction",
			SchemaName: "sent_reaction",
			Null:       !s.Flags.Has(15),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *StoryItem) SetFlags() {
	if !(s.Pinned == false) {
		s.Flags.Set(5)
	}
	if !(s.Public == false) {
		s.Flags.Set(7)
	}
	if !(s.CloseFriends == false) {
		s.Flags.Set(8)
	}
	if !(s.Min == false) {
		s.Flags.Set(9)
	}
	if !(s.Noforwards == false) {
		s.Flags.Set(10)
	}
	if !(s.Edited == false) {
		s.Flags.Set(11)
	}
	if !(s.Contacts == false) {
		s.Flags.Set(12)
	}
	if !(s.SelectedContacts == false) {
		s.Flags.Set(13)
	}
	if !(s.Out == false) {
		s.Flags.Set(16)
	}
	if !(s.FromID == nil) {
		s.Flags.Set(18)
	}
	if !(s.FwdFrom.Zero()) {
		s.Flags.Set(17)
	}
	if !(s.Caption == "") {
		s.Flags.Set(0)
	}
	if !(s.Entities == nil) {
		s.Flags.Set(1)
	}
	if !(s.MediaAreas == nil) {
		s.Flags.Set(14)
	}
	if !(s.Privacy == nil) {
		s.Flags.Set(2)
	}
	if !(s.Views.Zero()) {
		s.Flags.Set(3)
	}
	if !(s.SentReaction == nil) {
		s.Flags.Set(15)
	}
}

// Encode implements bin.Encoder.
func (s *StoryItem) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyItem#79b26a24 as nil")
	}
	b.PutID(StoryItemTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryItem) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyItem#79b26a24 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyItem#79b26a24: field flags: %w", err)
	}
	b.PutInt(s.ID)
	b.PutInt(s.Date)
	if s.Flags.Has(18) {
		if s.FromID == nil {
			return fmt.Errorf("unable to encode storyItem#79b26a24: field from_id is nil")
		}
		if err := s.FromID.Encode(b); err != nil {
			return fmt.Errorf("unable to encode storyItem#79b26a24: field from_id: %w", err)
		}
	}
	if s.Flags.Has(17) {
		if err := s.FwdFrom.Encode(b); err != nil {
			return fmt.Errorf("unable to encode storyItem#79b26a24: field fwd_from: %w", err)
		}
	}
	b.PutInt(s.ExpireDate)
	if s.Flags.Has(0) {
		b.PutString(s.Caption)
	}
	if s.Flags.Has(1) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode storyItem#79b26a24: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode storyItem#79b26a24: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if s.Media == nil {
		return fmt.Errorf("unable to encode storyItem#79b26a24: field media is nil")
	}
	if err := s.Media.Encode(b); err != nil {
		return fmt.Errorf("unable to encode storyItem#79b26a24: field media: %w", err)
	}
	if s.Flags.Has(14) {
		b.PutVectorHeader(len(s.MediaAreas))
		for idx, v := range s.MediaAreas {
			if v == nil {
				return fmt.Errorf("unable to encode storyItem#79b26a24: field media_areas element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode storyItem#79b26a24: field media_areas element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(2) {
		b.PutVectorHeader(len(s.Privacy))
		for idx, v := range s.Privacy {
			if v == nil {
				return fmt.Errorf("unable to encode storyItem#79b26a24: field privacy element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode storyItem#79b26a24: field privacy element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(3) {
		if err := s.Views.Encode(b); err != nil {
			return fmt.Errorf("unable to encode storyItem#79b26a24: field views: %w", err)
		}
	}
	if s.Flags.Has(15) {
		if s.SentReaction == nil {
			return fmt.Errorf("unable to encode storyItem#79b26a24: field sent_reaction is nil")
		}
		if err := s.SentReaction.Encode(b); err != nil {
			return fmt.Errorf("unable to encode storyItem#79b26a24: field sent_reaction: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryItem) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyItem#79b26a24 to nil")
	}
	if err := b.ConsumeID(StoryItemTypeID); err != nil {
		return fmt.Errorf("unable to decode storyItem#79b26a24: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryItem) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyItem#79b26a24 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field flags: %w", err)
		}
	}
	s.Pinned = s.Flags.Has(5)
	s.Public = s.Flags.Has(7)
	s.CloseFriends = s.Flags.Has(8)
	s.Min = s.Flags.Has(9)
	s.Noforwards = s.Flags.Has(10)
	s.Edited = s.Flags.Has(11)
	s.Contacts = s.Flags.Has(12)
	s.SelectedContacts = s.Flags.Has(13)
	s.Out = s.Flags.Has(16)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field id: %w", err)
		}
		s.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field date: %w", err)
		}
		s.Date = value
	}
	if s.Flags.Has(18) {
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field from_id: %w", err)
		}
		s.FromID = value
	}
	if s.Flags.Has(17) {
		if err := s.FwdFrom.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field fwd_from: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field expire_date: %w", err)
		}
		s.ExpireDate = value
	}
	if s.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field caption: %w", err)
		}
		s.Caption = value
	}
	if s.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field entities: %w", err)
		}

		if headerLen > 0 {
			s.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode storyItem#79b26a24: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	{
		value, err := DecodeMessageMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field media: %w", err)
		}
		s.Media = value
	}
	if s.Flags.Has(14) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field media_areas: %w", err)
		}

		if headerLen > 0 {
			s.MediaAreas = make([]MediaAreaClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMediaArea(b)
			if err != nil {
				return fmt.Errorf("unable to decode storyItem#79b26a24: field media_areas: %w", err)
			}
			s.MediaAreas = append(s.MediaAreas, value)
		}
	}
	if s.Flags.Has(2) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field privacy: %w", err)
		}

		if headerLen > 0 {
			s.Privacy = make([]PrivacyRuleClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePrivacyRule(b)
			if err != nil {
				return fmt.Errorf("unable to decode storyItem#79b26a24: field privacy: %w", err)
			}
			s.Privacy = append(s.Privacy, value)
		}
	}
	if s.Flags.Has(3) {
		if err := s.Views.Decode(b); err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field views: %w", err)
		}
	}
	if s.Flags.Has(15) {
		value, err := DecodeReaction(b)
		if err != nil {
			return fmt.Errorf("unable to decode storyItem#79b26a24: field sent_reaction: %w", err)
		}
		s.SentReaction = value
	}
	return nil
}

// SetPinned sets value of Pinned conditional field.
func (s *StoryItem) SetPinned(value bool) {
	if value {
		s.Flags.Set(5)
		s.Pinned = true
	} else {
		s.Flags.Unset(5)
		s.Pinned = false
	}
}

// GetPinned returns value of Pinned conditional field.
func (s *StoryItem) GetPinned() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(5)
}

// SetPublic sets value of Public conditional field.
func (s *StoryItem) SetPublic(value bool) {
	if value {
		s.Flags.Set(7)
		s.Public = true
	} else {
		s.Flags.Unset(7)
		s.Public = false
	}
}

// GetPublic returns value of Public conditional field.
func (s *StoryItem) GetPublic() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(7)
}

// SetCloseFriends sets value of CloseFriends conditional field.
func (s *StoryItem) SetCloseFriends(value bool) {
	if value {
		s.Flags.Set(8)
		s.CloseFriends = true
	} else {
		s.Flags.Unset(8)
		s.CloseFriends = false
	}
}

// GetCloseFriends returns value of CloseFriends conditional field.
func (s *StoryItem) GetCloseFriends() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(8)
}

// SetMin sets value of Min conditional field.
func (s *StoryItem) SetMin(value bool) {
	if value {
		s.Flags.Set(9)
		s.Min = true
	} else {
		s.Flags.Unset(9)
		s.Min = false
	}
}

// GetMin returns value of Min conditional field.
func (s *StoryItem) GetMin() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(9)
}

// SetNoforwards sets value of Noforwards conditional field.
func (s *StoryItem) SetNoforwards(value bool) {
	if value {
		s.Flags.Set(10)
		s.Noforwards = true
	} else {
		s.Flags.Unset(10)
		s.Noforwards = false
	}
}

// GetNoforwards returns value of Noforwards conditional field.
func (s *StoryItem) GetNoforwards() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(10)
}

// SetEdited sets value of Edited conditional field.
func (s *StoryItem) SetEdited(value bool) {
	if value {
		s.Flags.Set(11)
		s.Edited = true
	} else {
		s.Flags.Unset(11)
		s.Edited = false
	}
}

// GetEdited returns value of Edited conditional field.
func (s *StoryItem) GetEdited() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(11)
}

// SetContacts sets value of Contacts conditional field.
func (s *StoryItem) SetContacts(value bool) {
	if value {
		s.Flags.Set(12)
		s.Contacts = true
	} else {
		s.Flags.Unset(12)
		s.Contacts = false
	}
}

// GetContacts returns value of Contacts conditional field.
func (s *StoryItem) GetContacts() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(12)
}

// SetSelectedContacts sets value of SelectedContacts conditional field.
func (s *StoryItem) SetSelectedContacts(value bool) {
	if value {
		s.Flags.Set(13)
		s.SelectedContacts = true
	} else {
		s.Flags.Unset(13)
		s.SelectedContacts = false
	}
}

// GetSelectedContacts returns value of SelectedContacts conditional field.
func (s *StoryItem) GetSelectedContacts() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(13)
}

// SetOut sets value of Out conditional field.
func (s *StoryItem) SetOut(value bool) {
	if value {
		s.Flags.Set(16)
		s.Out = true
	} else {
		s.Flags.Unset(16)
		s.Out = false
	}
}

// GetOut returns value of Out conditional field.
func (s *StoryItem) GetOut() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(16)
}

// GetID returns value of ID field.
func (s *StoryItem) GetID() (value int) {
	if s == nil {
		return
	}
	return s.ID
}

// GetDate returns value of Date field.
func (s *StoryItem) GetDate() (value int) {
	if s == nil {
		return
	}
	return s.Date
}

// SetFromID sets value of FromID conditional field.
func (s *StoryItem) SetFromID(value PeerClass) {
	s.Flags.Set(18)
	s.FromID = value
}

// GetFromID returns value of FromID conditional field and
// boolean which is true if field was set.
func (s *StoryItem) GetFromID() (value PeerClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(18) {
		return value, false
	}
	return s.FromID, true
}

// SetFwdFrom sets value of FwdFrom conditional field.
func (s *StoryItem) SetFwdFrom(value StoryFwdHeader) {
	s.Flags.Set(17)
	s.FwdFrom = value
}

// GetFwdFrom returns value of FwdFrom conditional field and
// boolean which is true if field was set.
func (s *StoryItem) GetFwdFrom() (value StoryFwdHeader, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(17) {
		return value, false
	}
	return s.FwdFrom, true
}

// GetExpireDate returns value of ExpireDate field.
func (s *StoryItem) GetExpireDate() (value int) {
	if s == nil {
		return
	}
	return s.ExpireDate
}

// SetCaption sets value of Caption conditional field.
func (s *StoryItem) SetCaption(value string) {
	s.Flags.Set(0)
	s.Caption = value
}

// GetCaption returns value of Caption conditional field and
// boolean which is true if field was set.
func (s *StoryItem) GetCaption() (value string, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Caption, true
}

// SetEntities sets value of Entities conditional field.
func (s *StoryItem) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(1)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *StoryItem) GetEntities() (value []MessageEntityClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.Entities, true
}

// GetMedia returns value of Media field.
func (s *StoryItem) GetMedia() (value MessageMediaClass) {
	if s == nil {
		return
	}
	return s.Media
}

// SetMediaAreas sets value of MediaAreas conditional field.
func (s *StoryItem) SetMediaAreas(value []MediaAreaClass) {
	s.Flags.Set(14)
	s.MediaAreas = value
}

// GetMediaAreas returns value of MediaAreas conditional field and
// boolean which is true if field was set.
func (s *StoryItem) GetMediaAreas() (value []MediaAreaClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(14) {
		return value, false
	}
	return s.MediaAreas, true
}

// SetPrivacy sets value of Privacy conditional field.
func (s *StoryItem) SetPrivacy(value []PrivacyRuleClass) {
	s.Flags.Set(2)
	s.Privacy = value
}

// GetPrivacy returns value of Privacy conditional field and
// boolean which is true if field was set.
func (s *StoryItem) GetPrivacy() (value []PrivacyRuleClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.Privacy, true
}

// SetViews sets value of Views conditional field.
func (s *StoryItem) SetViews(value StoryViews) {
	s.Flags.Set(3)
	s.Views = value
}

// GetViews returns value of Views conditional field and
// boolean which is true if field was set.
func (s *StoryItem) GetViews() (value StoryViews, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Views, true
}

// SetSentReaction sets value of SentReaction conditional field.
func (s *StoryItem) SetSentReaction(value ReactionClass) {
	s.Flags.Set(15)
	s.SentReaction = value
}

// GetSentReaction returns value of SentReaction conditional field and
// boolean which is true if field was set.
func (s *StoryItem) GetSentReaction() (value ReactionClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(15) {
		return value, false
	}
	return s.SentReaction, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (s *StoryItem) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return MessageEntityClassArray(s.Entities), true
}

// MapMediaAreas returns field MediaAreas wrapped in MediaAreaClassArray helper.
func (s *StoryItem) MapMediaAreas() (value MediaAreaClassArray, ok bool) {
	if !s.Flags.Has(14) {
		return value, false
	}
	return MediaAreaClassArray(s.MediaAreas), true
}

// MapPrivacy returns field Privacy wrapped in PrivacyRuleClassArray helper.
func (s *StoryItem) MapPrivacy() (value PrivacyRuleClassArray, ok bool) {
	if !s.Flags.Has(2) {
		return value, false
	}
	return PrivacyRuleClassArray(s.Privacy), true
}

// StoryItemClassName is schema name of StoryItemClass.
const StoryItemClassName = "StoryItem"

// StoryItemClass represents StoryItem generic type.
//
// See https://core.telegram.org/type/StoryItem for reference.
//
// Example:
//
//	g, err := tg.DecodeStoryItem(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.StoryItemDeleted: // storyItemDeleted#51e6ee4f
//	case *tg.StoryItemSkipped: // storyItemSkipped#ffadc913
//	case *tg.StoryItem: // storyItem#79b26a24
//	default: panic(v)
//	}
type StoryItemClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() StoryItemClass

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

	// Story ID
	GetID() (value int)
}

// DecodeStoryItem implements binary de-serialization for StoryItemClass.
func DecodeStoryItem(buf *bin.Buffer) (StoryItemClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case StoryItemDeletedTypeID:
		// Decoding storyItemDeleted#51e6ee4f.
		v := StoryItemDeleted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryItemClass: %w", err)
		}
		return &v, nil
	case StoryItemSkippedTypeID:
		// Decoding storyItemSkipped#ffadc913.
		v := StoryItemSkipped{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryItemClass: %w", err)
		}
		return &v, nil
	case StoryItemTypeID:
		// Decoding storyItem#79b26a24.
		v := StoryItem{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode StoryItemClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode StoryItemClass: %w", bin.NewUnexpectedID(id))
	}
}

// StoryItem boxes the StoryItemClass providing a helper.
type StoryItemBox struct {
	StoryItem StoryItemClass
}

// Decode implements bin.Decoder for StoryItemBox.
func (b *StoryItemBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode StoryItemBox to nil")
	}
	v, err := DecodeStoryItem(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.StoryItem = v
	return nil
}

// Encode implements bin.Encode for StoryItemBox.
func (b *StoryItemBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.StoryItem == nil {
		return fmt.Errorf("unable to encode StoryItemClass as nil")
	}
	return b.StoryItem.Encode(buf)
}
