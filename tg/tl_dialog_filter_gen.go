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

// DialogFilter represents TL type `dialogFilter#5fb5523b`.
// Dialog filter AKA folder¹
//
// Links:
//  1. https://core.telegram.org/api/folders
//
// See https://core.telegram.org/constructor/dialogFilter for reference.
type DialogFilter struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to include all contacts in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Contacts bool
	// Whether to include all non-contacts in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	NonContacts bool
	// Whether to include all groups in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Groups bool
	// Whether to include all channels in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Broadcasts bool
	// Whether to include all bots in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Bots bool
	// Whether to exclude muted chats from this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ExcludeMuted bool
	// Whether to exclude read chats from this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ExcludeRead bool
	// Whether to exclude archived chats from this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ExcludeArchived bool
	// Folder¹ ID
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ID int
	// Folder¹ name
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	Title string
	// Emoji to use as icon for the folder.
	//
	// Use SetEmoticon and GetEmoticon helpers.
	Emoticon string
	// Color field of DialogFilter.
	//
	// Use SetColor and GetColor helpers.
	Color int
	// Pinned chats, folders¹ can have unlimited pinned chats
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	PinnedPeers []InputPeerClass
	// Include the following chats in this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	IncludePeers []InputPeerClass
	// Exclude the following chats from this folder¹
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	ExcludePeers []InputPeerClass
}

// DialogFilterTypeID is TL type id of DialogFilter.
const DialogFilterTypeID = 0x5fb5523b

// construct implements constructor of DialogFilterClass.
func (d DialogFilter) construct() DialogFilterClass { return &d }

// Ensuring interfaces in compile-time for DialogFilter.
var (
	_ bin.Encoder     = &DialogFilter{}
	_ bin.Decoder     = &DialogFilter{}
	_ bin.BareEncoder = &DialogFilter{}
	_ bin.BareDecoder = &DialogFilter{}

	_ DialogFilterClass = &DialogFilter{}
)

func (d *DialogFilter) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.Contacts == false) {
		return false
	}
	if !(d.NonContacts == false) {
		return false
	}
	if !(d.Groups == false) {
		return false
	}
	if !(d.Broadcasts == false) {
		return false
	}
	if !(d.Bots == false) {
		return false
	}
	if !(d.ExcludeMuted == false) {
		return false
	}
	if !(d.ExcludeRead == false) {
		return false
	}
	if !(d.ExcludeArchived == false) {
		return false
	}
	if !(d.ID == 0) {
		return false
	}
	if !(d.Title == "") {
		return false
	}
	if !(d.Emoticon == "") {
		return false
	}
	if !(d.Color == 0) {
		return false
	}
	if !(d.PinnedPeers == nil) {
		return false
	}
	if !(d.IncludePeers == nil) {
		return false
	}
	if !(d.ExcludePeers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DialogFilter) String() string {
	if d == nil {
		return "DialogFilter(nil)"
	}
	type Alias DialogFilter
	return fmt.Sprintf("DialogFilter%+v", Alias(*d))
}

// FillFrom fills DialogFilter from given interface.
func (d *DialogFilter) FillFrom(from interface {
	GetContacts() (value bool)
	GetNonContacts() (value bool)
	GetGroups() (value bool)
	GetBroadcasts() (value bool)
	GetBots() (value bool)
	GetExcludeMuted() (value bool)
	GetExcludeRead() (value bool)
	GetExcludeArchived() (value bool)
	GetID() (value int)
	GetTitle() (value string)
	GetEmoticon() (value string, ok bool)
	GetColor() (value int, ok bool)
	GetPinnedPeers() (value []InputPeerClass)
	GetIncludePeers() (value []InputPeerClass)
	GetExcludePeers() (value []InputPeerClass)
}) {
	d.Contacts = from.GetContacts()
	d.NonContacts = from.GetNonContacts()
	d.Groups = from.GetGroups()
	d.Broadcasts = from.GetBroadcasts()
	d.Bots = from.GetBots()
	d.ExcludeMuted = from.GetExcludeMuted()
	d.ExcludeRead = from.GetExcludeRead()
	d.ExcludeArchived = from.GetExcludeArchived()
	d.ID = from.GetID()
	d.Title = from.GetTitle()
	if val, ok := from.GetEmoticon(); ok {
		d.Emoticon = val
	}

	if val, ok := from.GetColor(); ok {
		d.Color = val
	}

	d.PinnedPeers = from.GetPinnedPeers()
	d.IncludePeers = from.GetIncludePeers()
	d.ExcludePeers = from.GetExcludePeers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DialogFilter) TypeID() uint32 {
	return DialogFilterTypeID
}

// TypeName returns name of type in TL schema.
func (*DialogFilter) TypeName() string {
	return "dialogFilter"
}

// TypeInfo returns info about TL type.
func (d *DialogFilter) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dialogFilter",
		ID:   DialogFilterTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Contacts",
			SchemaName: "contacts",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "NonContacts",
			SchemaName: "non_contacts",
			Null:       !d.Flags.Has(1),
		},
		{
			Name:       "Groups",
			SchemaName: "groups",
			Null:       !d.Flags.Has(2),
		},
		{
			Name:       "Broadcasts",
			SchemaName: "broadcasts",
			Null:       !d.Flags.Has(3),
		},
		{
			Name:       "Bots",
			SchemaName: "bots",
			Null:       !d.Flags.Has(4),
		},
		{
			Name:       "ExcludeMuted",
			SchemaName: "exclude_muted",
			Null:       !d.Flags.Has(11),
		},
		{
			Name:       "ExcludeRead",
			SchemaName: "exclude_read",
			Null:       !d.Flags.Has(12),
		},
		{
			Name:       "ExcludeArchived",
			SchemaName: "exclude_archived",
			Null:       !d.Flags.Has(13),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Emoticon",
			SchemaName: "emoticon",
			Null:       !d.Flags.Has(25),
		},
		{
			Name:       "Color",
			SchemaName: "color",
			Null:       !d.Flags.Has(27),
		},
		{
			Name:       "PinnedPeers",
			SchemaName: "pinned_peers",
		},
		{
			Name:       "IncludePeers",
			SchemaName: "include_peers",
		},
		{
			Name:       "ExcludePeers",
			SchemaName: "exclude_peers",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (d *DialogFilter) SetFlags() {
	if !(d.Contacts == false) {
		d.Flags.Set(0)
	}
	if !(d.NonContacts == false) {
		d.Flags.Set(1)
	}
	if !(d.Groups == false) {
		d.Flags.Set(2)
	}
	if !(d.Broadcasts == false) {
		d.Flags.Set(3)
	}
	if !(d.Bots == false) {
		d.Flags.Set(4)
	}
	if !(d.ExcludeMuted == false) {
		d.Flags.Set(11)
	}
	if !(d.ExcludeRead == false) {
		d.Flags.Set(12)
	}
	if !(d.ExcludeArchived == false) {
		d.Flags.Set(13)
	}
	if !(d.Emoticon == "") {
		d.Flags.Set(25)
	}
	if !(d.Color == 0) {
		d.Flags.Set(27)
	}
}

// Encode implements bin.Encoder.
func (d *DialogFilter) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogFilter#5fb5523b as nil")
	}
	b.PutID(DialogFilterTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DialogFilter) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogFilter#5fb5523b as nil")
	}
	d.SetFlags()
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialogFilter#5fb5523b: field flags: %w", err)
	}
	b.PutInt(d.ID)
	b.PutString(d.Title)
	if d.Flags.Has(25) {
		b.PutString(d.Emoticon)
	}
	if d.Flags.Has(27) {
		b.PutInt(d.Color)
	}
	b.PutVectorHeader(len(d.PinnedPeers))
	for idx, v := range d.PinnedPeers {
		if v == nil {
			return fmt.Errorf("unable to encode dialogFilter#5fb5523b: field pinned_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode dialogFilter#5fb5523b: field pinned_peers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.IncludePeers))
	for idx, v := range d.IncludePeers {
		if v == nil {
			return fmt.Errorf("unable to encode dialogFilter#5fb5523b: field include_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode dialogFilter#5fb5523b: field include_peers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.ExcludePeers))
	for idx, v := range d.ExcludePeers {
		if v == nil {
			return fmt.Errorf("unable to encode dialogFilter#5fb5523b: field exclude_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode dialogFilter#5fb5523b: field exclude_peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DialogFilter) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogFilter#5fb5523b to nil")
	}
	if err := b.ConsumeID(DialogFilterTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogFilter#5fb5523b: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DialogFilter) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogFilter#5fb5523b to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field flags: %w", err)
		}
	}
	d.Contacts = d.Flags.Has(0)
	d.NonContacts = d.Flags.Has(1)
	d.Groups = d.Flags.Has(2)
	d.Broadcasts = d.Flags.Has(3)
	d.Bots = d.Flags.Has(4)
	d.ExcludeMuted = d.Flags.Has(11)
	d.ExcludeRead = d.Flags.Has(12)
	d.ExcludeArchived = d.Flags.Has(13)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field id: %w", err)
		}
		d.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field title: %w", err)
		}
		d.Title = value
	}
	if d.Flags.Has(25) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field emoticon: %w", err)
		}
		d.Emoticon = value
	}
	if d.Flags.Has(27) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field color: %w", err)
		}
		d.Color = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field pinned_peers: %w", err)
		}

		if headerLen > 0 {
			d.PinnedPeers = make([]InputPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field pinned_peers: %w", err)
			}
			d.PinnedPeers = append(d.PinnedPeers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field include_peers: %w", err)
		}

		if headerLen > 0 {
			d.IncludePeers = make([]InputPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field include_peers: %w", err)
			}
			d.IncludePeers = append(d.IncludePeers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field exclude_peers: %w", err)
		}

		if headerLen > 0 {
			d.ExcludePeers = make([]InputPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode dialogFilter#5fb5523b: field exclude_peers: %w", err)
			}
			d.ExcludePeers = append(d.ExcludePeers, value)
		}
	}
	return nil
}

// SetContacts sets value of Contacts conditional field.
func (d *DialogFilter) SetContacts(value bool) {
	if value {
		d.Flags.Set(0)
		d.Contacts = true
	} else {
		d.Flags.Unset(0)
		d.Contacts = false
	}
}

// GetContacts returns value of Contacts conditional field.
func (d *DialogFilter) GetContacts() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(0)
}

// SetNonContacts sets value of NonContacts conditional field.
func (d *DialogFilter) SetNonContacts(value bool) {
	if value {
		d.Flags.Set(1)
		d.NonContacts = true
	} else {
		d.Flags.Unset(1)
		d.NonContacts = false
	}
}

// GetNonContacts returns value of NonContacts conditional field.
func (d *DialogFilter) GetNonContacts() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(1)
}

// SetGroups sets value of Groups conditional field.
func (d *DialogFilter) SetGroups(value bool) {
	if value {
		d.Flags.Set(2)
		d.Groups = true
	} else {
		d.Flags.Unset(2)
		d.Groups = false
	}
}

// GetGroups returns value of Groups conditional field.
func (d *DialogFilter) GetGroups() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(2)
}

// SetBroadcasts sets value of Broadcasts conditional field.
func (d *DialogFilter) SetBroadcasts(value bool) {
	if value {
		d.Flags.Set(3)
		d.Broadcasts = true
	} else {
		d.Flags.Unset(3)
		d.Broadcasts = false
	}
}

// GetBroadcasts returns value of Broadcasts conditional field.
func (d *DialogFilter) GetBroadcasts() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(3)
}

// SetBots sets value of Bots conditional field.
func (d *DialogFilter) SetBots(value bool) {
	if value {
		d.Flags.Set(4)
		d.Bots = true
	} else {
		d.Flags.Unset(4)
		d.Bots = false
	}
}

// GetBots returns value of Bots conditional field.
func (d *DialogFilter) GetBots() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(4)
}

// SetExcludeMuted sets value of ExcludeMuted conditional field.
func (d *DialogFilter) SetExcludeMuted(value bool) {
	if value {
		d.Flags.Set(11)
		d.ExcludeMuted = true
	} else {
		d.Flags.Unset(11)
		d.ExcludeMuted = false
	}
}

// GetExcludeMuted returns value of ExcludeMuted conditional field.
func (d *DialogFilter) GetExcludeMuted() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(11)
}

// SetExcludeRead sets value of ExcludeRead conditional field.
func (d *DialogFilter) SetExcludeRead(value bool) {
	if value {
		d.Flags.Set(12)
		d.ExcludeRead = true
	} else {
		d.Flags.Unset(12)
		d.ExcludeRead = false
	}
}

// GetExcludeRead returns value of ExcludeRead conditional field.
func (d *DialogFilter) GetExcludeRead() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(12)
}

// SetExcludeArchived sets value of ExcludeArchived conditional field.
func (d *DialogFilter) SetExcludeArchived(value bool) {
	if value {
		d.Flags.Set(13)
		d.ExcludeArchived = true
	} else {
		d.Flags.Unset(13)
		d.ExcludeArchived = false
	}
}

// GetExcludeArchived returns value of ExcludeArchived conditional field.
func (d *DialogFilter) GetExcludeArchived() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(13)
}

// GetID returns value of ID field.
func (d *DialogFilter) GetID() (value int) {
	if d == nil {
		return
	}
	return d.ID
}

// GetTitle returns value of Title field.
func (d *DialogFilter) GetTitle() (value string) {
	if d == nil {
		return
	}
	return d.Title
}

// SetEmoticon sets value of Emoticon conditional field.
func (d *DialogFilter) SetEmoticon(value string) {
	d.Flags.Set(25)
	d.Emoticon = value
}

// GetEmoticon returns value of Emoticon conditional field and
// boolean which is true if field was set.
func (d *DialogFilter) GetEmoticon() (value string, ok bool) {
	if d == nil {
		return
	}
	if !d.Flags.Has(25) {
		return value, false
	}
	return d.Emoticon, true
}

// SetColor sets value of Color conditional field.
func (d *DialogFilter) SetColor(value int) {
	d.Flags.Set(27)
	d.Color = value
}

// GetColor returns value of Color conditional field and
// boolean which is true if field was set.
func (d *DialogFilter) GetColor() (value int, ok bool) {
	if d == nil {
		return
	}
	if !d.Flags.Has(27) {
		return value, false
	}
	return d.Color, true
}

// GetPinnedPeers returns value of PinnedPeers field.
func (d *DialogFilter) GetPinnedPeers() (value []InputPeerClass) {
	if d == nil {
		return
	}
	return d.PinnedPeers
}

// GetIncludePeers returns value of IncludePeers field.
func (d *DialogFilter) GetIncludePeers() (value []InputPeerClass) {
	if d == nil {
		return
	}
	return d.IncludePeers
}

// GetExcludePeers returns value of ExcludePeers field.
func (d *DialogFilter) GetExcludePeers() (value []InputPeerClass) {
	if d == nil {
		return
	}
	return d.ExcludePeers
}

// MapPinnedPeers returns field PinnedPeers wrapped in InputPeerClassArray helper.
func (d *DialogFilter) MapPinnedPeers() (value InputPeerClassArray) {
	return InputPeerClassArray(d.PinnedPeers)
}

// MapIncludePeers returns field IncludePeers wrapped in InputPeerClassArray helper.
func (d *DialogFilter) MapIncludePeers() (value InputPeerClassArray) {
	return InputPeerClassArray(d.IncludePeers)
}

// MapExcludePeers returns field ExcludePeers wrapped in InputPeerClassArray helper.
func (d *DialogFilter) MapExcludePeers() (value InputPeerClassArray) {
	return InputPeerClassArray(d.ExcludePeers)
}

// DialogFilterDefault represents TL type `dialogFilterDefault#363293ae`.
// Used only when reordering folders to indicate the default (all chats) folder.
//
// See https://core.telegram.org/constructor/dialogFilterDefault for reference.
type DialogFilterDefault struct {
}

// DialogFilterDefaultTypeID is TL type id of DialogFilterDefault.
const DialogFilterDefaultTypeID = 0x363293ae

// construct implements constructor of DialogFilterClass.
func (d DialogFilterDefault) construct() DialogFilterClass { return &d }

// Ensuring interfaces in compile-time for DialogFilterDefault.
var (
	_ bin.Encoder     = &DialogFilterDefault{}
	_ bin.Decoder     = &DialogFilterDefault{}
	_ bin.BareEncoder = &DialogFilterDefault{}
	_ bin.BareDecoder = &DialogFilterDefault{}

	_ DialogFilterClass = &DialogFilterDefault{}
)

func (d *DialogFilterDefault) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *DialogFilterDefault) String() string {
	if d == nil {
		return "DialogFilterDefault(nil)"
	}
	type Alias DialogFilterDefault
	return fmt.Sprintf("DialogFilterDefault%+v", Alias(*d))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DialogFilterDefault) TypeID() uint32 {
	return DialogFilterDefaultTypeID
}

// TypeName returns name of type in TL schema.
func (*DialogFilterDefault) TypeName() string {
	return "dialogFilterDefault"
}

// TypeInfo returns info about TL type.
func (d *DialogFilterDefault) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dialogFilterDefault",
		ID:   DialogFilterDefaultTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (d *DialogFilterDefault) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogFilterDefault#363293ae as nil")
	}
	b.PutID(DialogFilterDefaultTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DialogFilterDefault) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogFilterDefault#363293ae as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DialogFilterDefault) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogFilterDefault#363293ae to nil")
	}
	if err := b.ConsumeID(DialogFilterDefaultTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogFilterDefault#363293ae: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DialogFilterDefault) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogFilterDefault#363293ae to nil")
	}
	return nil
}

// DialogFilterChatlist represents TL type `dialogFilterChatlist#9fe28ea4`.
// A folder imported using a chat folder deep link »¹.
//
// Links:
//  1. https://core.telegram.org/api/links#chat-folder-links
//
// See https://core.telegram.org/constructor/dialogFilterChatlist for reference.
type DialogFilterChatlist struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the current user has created some chat folder deep links »¹ to share the
	// folder as well.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#chat-folder-links
	HasMyInvites bool
	// ID of the folder
	ID int
	// Name of the folder
	Title string
	// Emoji to use as icon for the folder.
	//
	// Use SetEmoticon and GetEmoticon helpers.
	Emoticon string
	// Color field of DialogFilterChatlist.
	//
	// Use SetColor and GetColor helpers.
	Color int
	// Pinned chats, folders¹ can have unlimited pinned chats
	//
	// Links:
	//  1) https://core.telegram.org/api/folders
	PinnedPeers []InputPeerClass
	// Chats to include in the folder
	IncludePeers []InputPeerClass
}

// DialogFilterChatlistTypeID is TL type id of DialogFilterChatlist.
const DialogFilterChatlistTypeID = 0x9fe28ea4

// construct implements constructor of DialogFilterClass.
func (d DialogFilterChatlist) construct() DialogFilterClass { return &d }

// Ensuring interfaces in compile-time for DialogFilterChatlist.
var (
	_ bin.Encoder     = &DialogFilterChatlist{}
	_ bin.Decoder     = &DialogFilterChatlist{}
	_ bin.BareEncoder = &DialogFilterChatlist{}
	_ bin.BareDecoder = &DialogFilterChatlist{}

	_ DialogFilterClass = &DialogFilterChatlist{}
)

func (d *DialogFilterChatlist) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.HasMyInvites == false) {
		return false
	}
	if !(d.ID == 0) {
		return false
	}
	if !(d.Title == "") {
		return false
	}
	if !(d.Emoticon == "") {
		return false
	}
	if !(d.Color == 0) {
		return false
	}
	if !(d.PinnedPeers == nil) {
		return false
	}
	if !(d.IncludePeers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *DialogFilterChatlist) String() string {
	if d == nil {
		return "DialogFilterChatlist(nil)"
	}
	type Alias DialogFilterChatlist
	return fmt.Sprintf("DialogFilterChatlist%+v", Alias(*d))
}

// FillFrom fills DialogFilterChatlist from given interface.
func (d *DialogFilterChatlist) FillFrom(from interface {
	GetHasMyInvites() (value bool)
	GetID() (value int)
	GetTitle() (value string)
	GetEmoticon() (value string, ok bool)
	GetColor() (value int, ok bool)
	GetPinnedPeers() (value []InputPeerClass)
	GetIncludePeers() (value []InputPeerClass)
}) {
	d.HasMyInvites = from.GetHasMyInvites()
	d.ID = from.GetID()
	d.Title = from.GetTitle()
	if val, ok := from.GetEmoticon(); ok {
		d.Emoticon = val
	}

	if val, ok := from.GetColor(); ok {
		d.Color = val
	}

	d.PinnedPeers = from.GetPinnedPeers()
	d.IncludePeers = from.GetIncludePeers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*DialogFilterChatlist) TypeID() uint32 {
	return DialogFilterChatlistTypeID
}

// TypeName returns name of type in TL schema.
func (*DialogFilterChatlist) TypeName() string {
	return "dialogFilterChatlist"
}

// TypeInfo returns info about TL type.
func (d *DialogFilterChatlist) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "dialogFilterChatlist",
		ID:   DialogFilterChatlistTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "HasMyInvites",
			SchemaName: "has_my_invites",
			Null:       !d.Flags.Has(26),
		},
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Emoticon",
			SchemaName: "emoticon",
			Null:       !d.Flags.Has(25),
		},
		{
			Name:       "Color",
			SchemaName: "color",
			Null:       !d.Flags.Has(27),
		},
		{
			Name:       "PinnedPeers",
			SchemaName: "pinned_peers",
		},
		{
			Name:       "IncludePeers",
			SchemaName: "include_peers",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (d *DialogFilterChatlist) SetFlags() {
	if !(d.HasMyInvites == false) {
		d.Flags.Set(26)
	}
	if !(d.Emoticon == "") {
		d.Flags.Set(25)
	}
	if !(d.Color == 0) {
		d.Flags.Set(27)
	}
}

// Encode implements bin.Encoder.
func (d *DialogFilterChatlist) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogFilterChatlist#9fe28ea4 as nil")
	}
	b.PutID(DialogFilterChatlistTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *DialogFilterChatlist) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode dialogFilterChatlist#9fe28ea4 as nil")
	}
	d.SetFlags()
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode dialogFilterChatlist#9fe28ea4: field flags: %w", err)
	}
	b.PutInt(d.ID)
	b.PutString(d.Title)
	if d.Flags.Has(25) {
		b.PutString(d.Emoticon)
	}
	if d.Flags.Has(27) {
		b.PutInt(d.Color)
	}
	b.PutVectorHeader(len(d.PinnedPeers))
	for idx, v := range d.PinnedPeers {
		if v == nil {
			return fmt.Errorf("unable to encode dialogFilterChatlist#9fe28ea4: field pinned_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode dialogFilterChatlist#9fe28ea4: field pinned_peers element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(d.IncludePeers))
	for idx, v := range d.IncludePeers {
		if v == nil {
			return fmt.Errorf("unable to encode dialogFilterChatlist#9fe28ea4: field include_peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode dialogFilterChatlist#9fe28ea4: field include_peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *DialogFilterChatlist) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogFilterChatlist#9fe28ea4 to nil")
	}
	if err := b.ConsumeID(DialogFilterChatlistTypeID); err != nil {
		return fmt.Errorf("unable to decode dialogFilterChatlist#9fe28ea4: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *DialogFilterChatlist) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode dialogFilterChatlist#9fe28ea4 to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode dialogFilterChatlist#9fe28ea4: field flags: %w", err)
		}
	}
	d.HasMyInvites = d.Flags.Has(26)
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterChatlist#9fe28ea4: field id: %w", err)
		}
		d.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterChatlist#9fe28ea4: field title: %w", err)
		}
		d.Title = value
	}
	if d.Flags.Has(25) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterChatlist#9fe28ea4: field emoticon: %w", err)
		}
		d.Emoticon = value
	}
	if d.Flags.Has(27) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterChatlist#9fe28ea4: field color: %w", err)
		}
		d.Color = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterChatlist#9fe28ea4: field pinned_peers: %w", err)
		}

		if headerLen > 0 {
			d.PinnedPeers = make([]InputPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode dialogFilterChatlist#9fe28ea4: field pinned_peers: %w", err)
			}
			d.PinnedPeers = append(d.PinnedPeers, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode dialogFilterChatlist#9fe28ea4: field include_peers: %w", err)
		}

		if headerLen > 0 {
			d.IncludePeers = make([]InputPeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode dialogFilterChatlist#9fe28ea4: field include_peers: %w", err)
			}
			d.IncludePeers = append(d.IncludePeers, value)
		}
	}
	return nil
}

// SetHasMyInvites sets value of HasMyInvites conditional field.
func (d *DialogFilterChatlist) SetHasMyInvites(value bool) {
	if value {
		d.Flags.Set(26)
		d.HasMyInvites = true
	} else {
		d.Flags.Unset(26)
		d.HasMyInvites = false
	}
}

// GetHasMyInvites returns value of HasMyInvites conditional field.
func (d *DialogFilterChatlist) GetHasMyInvites() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(26)
}

// GetID returns value of ID field.
func (d *DialogFilterChatlist) GetID() (value int) {
	if d == nil {
		return
	}
	return d.ID
}

// GetTitle returns value of Title field.
func (d *DialogFilterChatlist) GetTitle() (value string) {
	if d == nil {
		return
	}
	return d.Title
}

// SetEmoticon sets value of Emoticon conditional field.
func (d *DialogFilterChatlist) SetEmoticon(value string) {
	d.Flags.Set(25)
	d.Emoticon = value
}

// GetEmoticon returns value of Emoticon conditional field and
// boolean which is true if field was set.
func (d *DialogFilterChatlist) GetEmoticon() (value string, ok bool) {
	if d == nil {
		return
	}
	if !d.Flags.Has(25) {
		return value, false
	}
	return d.Emoticon, true
}

// SetColor sets value of Color conditional field.
func (d *DialogFilterChatlist) SetColor(value int) {
	d.Flags.Set(27)
	d.Color = value
}

// GetColor returns value of Color conditional field and
// boolean which is true if field was set.
func (d *DialogFilterChatlist) GetColor() (value int, ok bool) {
	if d == nil {
		return
	}
	if !d.Flags.Has(27) {
		return value, false
	}
	return d.Color, true
}

// GetPinnedPeers returns value of PinnedPeers field.
func (d *DialogFilterChatlist) GetPinnedPeers() (value []InputPeerClass) {
	if d == nil {
		return
	}
	return d.PinnedPeers
}

// GetIncludePeers returns value of IncludePeers field.
func (d *DialogFilterChatlist) GetIncludePeers() (value []InputPeerClass) {
	if d == nil {
		return
	}
	return d.IncludePeers
}

// MapPinnedPeers returns field PinnedPeers wrapped in InputPeerClassArray helper.
func (d *DialogFilterChatlist) MapPinnedPeers() (value InputPeerClassArray) {
	return InputPeerClassArray(d.PinnedPeers)
}

// MapIncludePeers returns field IncludePeers wrapped in InputPeerClassArray helper.
func (d *DialogFilterChatlist) MapIncludePeers() (value InputPeerClassArray) {
	return InputPeerClassArray(d.IncludePeers)
}

// DialogFilterClassName is schema name of DialogFilterClass.
const DialogFilterClassName = "DialogFilter"

// DialogFilterClass represents DialogFilter generic type.
//
// See https://core.telegram.org/type/DialogFilter for reference.
//
// Example:
//
//	g, err := tg.DecodeDialogFilter(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.DialogFilter: // dialogFilter#5fb5523b
//	case *tg.DialogFilterDefault: // dialogFilterDefault#363293ae
//	case *tg.DialogFilterChatlist: // dialogFilterChatlist#9fe28ea4
//	default: panic(v)
//	}
type DialogFilterClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() DialogFilterClass

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

// AsInputChatlist tries to map DialogFilter to InputChatlistDialogFilter.
func (d *DialogFilter) AsInputChatlist() *InputChatlistDialogFilter {
	value := new(InputChatlistDialogFilter)
	value.FilterID = d.GetID()

	return value
}

// DecodeDialogFilter implements binary de-serialization for DialogFilterClass.
func DecodeDialogFilter(buf *bin.Buffer) (DialogFilterClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case DialogFilterTypeID:
		// Decoding dialogFilter#5fb5523b.
		v := DialogFilter{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DialogFilterClass: %w", err)
		}
		return &v, nil
	case DialogFilterDefaultTypeID:
		// Decoding dialogFilterDefault#363293ae.
		v := DialogFilterDefault{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DialogFilterClass: %w", err)
		}
		return &v, nil
	case DialogFilterChatlistTypeID:
		// Decoding dialogFilterChatlist#9fe28ea4.
		v := DialogFilterChatlist{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode DialogFilterClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode DialogFilterClass: %w", bin.NewUnexpectedID(id))
	}
}

// DialogFilter boxes the DialogFilterClass providing a helper.
type DialogFilterBox struct {
	DialogFilter DialogFilterClass
}

// Decode implements bin.Decoder for DialogFilterBox.
func (b *DialogFilterBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode DialogFilterBox to nil")
	}
	v, err := DecodeDialogFilter(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DialogFilter = v
	return nil
}

// Encode implements bin.Encode for DialogFilterBox.
func (b *DialogFilterBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DialogFilter == nil {
		return fmt.Errorf("unable to encode DialogFilterClass as nil")
	}
	return b.DialogFilter.Encode(buf)
}
