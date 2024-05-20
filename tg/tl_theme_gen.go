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

// Theme represents TL type `theme#a00e67d6`.
// Theme
//
// See https://core.telegram.org/constructor/theme for reference.
type Theme struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the current user is the creator of this theme
	Creator bool
	// Whether this is the default theme
	Default bool
	// Whether this theme is meant to be used as a chat theme¹
	//
	// Links:
	//  1) https://telegram.org/blog/chat-themes-interactive-emoji-read-receipts
	ForChat bool
	// Theme ID
	ID int64
	// Theme access hash
	AccessHash int64
	// Unique theme ID
	Slug string
	// Theme name
	Title string
	// Theme
	//
	// Use SetDocument and GetDocument helpers.
	Document DocumentClass
	// Theme settings
	//
	// Use SetSettings and GetSettings helpers.
	Settings []ThemeSettings
	// Theme emoji
	//
	// Use SetEmoticon and GetEmoticon helpers.
	Emoticon string
	// Installation count
	//
	// Use SetInstallsCount and GetInstallsCount helpers.
	InstallsCount int
}

// ThemeTypeID is TL type id of Theme.
const ThemeTypeID = 0xa00e67d6

// Ensuring interfaces in compile-time for Theme.
var (
	_ bin.Encoder     = &Theme{}
	_ bin.Decoder     = &Theme{}
	_ bin.BareEncoder = &Theme{}
	_ bin.BareDecoder = &Theme{}
)

func (t *Theme) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Flags.Zero()) {
		return false
	}
	if !(t.Creator == false) {
		return false
	}
	if !(t.Default == false) {
		return false
	}
	if !(t.ForChat == false) {
		return false
	}
	if !(t.ID == 0) {
		return false
	}
	if !(t.AccessHash == 0) {
		return false
	}
	if !(t.Slug == "") {
		return false
	}
	if !(t.Title == "") {
		return false
	}
	if !(t.Document == nil) {
		return false
	}
	if !(t.Settings == nil) {
		return false
	}
	if !(t.Emoticon == "") {
		return false
	}
	if !(t.InstallsCount == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *Theme) String() string {
	if t == nil {
		return "Theme(nil)"
	}
	type Alias Theme
	return fmt.Sprintf("Theme%+v", Alias(*t))
}

// FillFrom fills Theme from given interface.
func (t *Theme) FillFrom(from interface {
	GetCreator() (value bool)
	GetDefault() (value bool)
	GetForChat() (value bool)
	GetID() (value int64)
	GetAccessHash() (value int64)
	GetSlug() (value string)
	GetTitle() (value string)
	GetDocument() (value DocumentClass, ok bool)
	GetSettings() (value []ThemeSettings, ok bool)
	GetEmoticon() (value string, ok bool)
	GetInstallsCount() (value int, ok bool)
}) {
	t.Creator = from.GetCreator()
	t.Default = from.GetDefault()
	t.ForChat = from.GetForChat()
	t.ID = from.GetID()
	t.AccessHash = from.GetAccessHash()
	t.Slug = from.GetSlug()
	t.Title = from.GetTitle()
	if val, ok := from.GetDocument(); ok {
		t.Document = val
	}

	if val, ok := from.GetSettings(); ok {
		t.Settings = val
	}

	if val, ok := from.GetEmoticon(); ok {
		t.Emoticon = val
	}

	if val, ok := from.GetInstallsCount(); ok {
		t.InstallsCount = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Theme) TypeID() uint32 {
	return ThemeTypeID
}

// TypeName returns name of type in TL schema.
func (*Theme) TypeName() string {
	return "theme"
}

// TypeInfo returns info about TL type.
func (t *Theme) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "theme",
		ID:   ThemeTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Creator",
			SchemaName: "creator",
			Null:       !t.Flags.Has(0),
		},
		{
			Name:       "Default",
			SchemaName: "default",
			Null:       !t.Flags.Has(1),
		},
		{
			Name:       "ForChat",
			SchemaName: "for_chat",
			Null:       !t.Flags.Has(5),
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
			Name:       "Slug",
			SchemaName: "slug",
		},
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "Document",
			SchemaName: "document",
			Null:       !t.Flags.Has(2),
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
			Null:       !t.Flags.Has(3),
		},
		{
			Name:       "Emoticon",
			SchemaName: "emoticon",
			Null:       !t.Flags.Has(6),
		},
		{
			Name:       "InstallsCount",
			SchemaName: "installs_count",
			Null:       !t.Flags.Has(4),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (t *Theme) SetFlags() {
	if !(t.Creator == false) {
		t.Flags.Set(0)
	}
	if !(t.Default == false) {
		t.Flags.Set(1)
	}
	if !(t.ForChat == false) {
		t.Flags.Set(5)
	}
	if !(t.Document == nil) {
		t.Flags.Set(2)
	}
	if !(t.Settings == nil) {
		t.Flags.Set(3)
	}
	if !(t.Emoticon == "") {
		t.Flags.Set(6)
	}
	if !(t.InstallsCount == 0) {
		t.Flags.Set(4)
	}
}

// Encode implements bin.Encoder.
func (t *Theme) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode theme#a00e67d6 as nil")
	}
	b.PutID(ThemeTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *Theme) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode theme#a00e67d6 as nil")
	}
	t.SetFlags()
	if err := t.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode theme#a00e67d6: field flags: %w", err)
	}
	b.PutLong(t.ID)
	b.PutLong(t.AccessHash)
	b.PutString(t.Slug)
	b.PutString(t.Title)
	if t.Flags.Has(2) {
		if t.Document == nil {
			return fmt.Errorf("unable to encode theme#a00e67d6: field document is nil")
		}
		if err := t.Document.Encode(b); err != nil {
			return fmt.Errorf("unable to encode theme#a00e67d6: field document: %w", err)
		}
	}
	if t.Flags.Has(3) {
		b.PutVectorHeader(len(t.Settings))
		for idx, v := range t.Settings {
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode theme#a00e67d6: field settings element with index %d: %w", idx, err)
			}
		}
	}
	if t.Flags.Has(6) {
		b.PutString(t.Emoticon)
	}
	if t.Flags.Has(4) {
		b.PutInt(t.InstallsCount)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *Theme) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode theme#a00e67d6 to nil")
	}
	if err := b.ConsumeID(ThemeTypeID); err != nil {
		return fmt.Errorf("unable to decode theme#a00e67d6: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *Theme) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode theme#a00e67d6 to nil")
	}
	{
		if err := t.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode theme#a00e67d6: field flags: %w", err)
		}
	}
	t.Creator = t.Flags.Has(0)
	t.Default = t.Flags.Has(1)
	t.ForChat = t.Flags.Has(5)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode theme#a00e67d6: field id: %w", err)
		}
		t.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode theme#a00e67d6: field access_hash: %w", err)
		}
		t.AccessHash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode theme#a00e67d6: field slug: %w", err)
		}
		t.Slug = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode theme#a00e67d6: field title: %w", err)
		}
		t.Title = value
	}
	if t.Flags.Has(2) {
		value, err := DecodeDocument(b)
		if err != nil {
			return fmt.Errorf("unable to decode theme#a00e67d6: field document: %w", err)
		}
		t.Document = value
	}
	if t.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode theme#a00e67d6: field settings: %w", err)
		}

		if headerLen > 0 {
			t.Settings = make([]ThemeSettings, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ThemeSettings
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode theme#a00e67d6: field settings: %w", err)
			}
			t.Settings = append(t.Settings, value)
		}
	}
	if t.Flags.Has(6) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode theme#a00e67d6: field emoticon: %w", err)
		}
		t.Emoticon = value
	}
	if t.Flags.Has(4) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode theme#a00e67d6: field installs_count: %w", err)
		}
		t.InstallsCount = value
	}
	return nil
}

// SetCreator sets value of Creator conditional field.
func (t *Theme) SetCreator(value bool) {
	if value {
		t.Flags.Set(0)
		t.Creator = true
	} else {
		t.Flags.Unset(0)
		t.Creator = false
	}
}

// GetCreator returns value of Creator conditional field.
func (t *Theme) GetCreator() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(0)
}

// SetDefault sets value of Default conditional field.
func (t *Theme) SetDefault(value bool) {
	if value {
		t.Flags.Set(1)
		t.Default = true
	} else {
		t.Flags.Unset(1)
		t.Default = false
	}
}

// GetDefault returns value of Default conditional field.
func (t *Theme) GetDefault() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(1)
}

// SetForChat sets value of ForChat conditional field.
func (t *Theme) SetForChat(value bool) {
	if value {
		t.Flags.Set(5)
		t.ForChat = true
	} else {
		t.Flags.Unset(5)
		t.ForChat = false
	}
}

// GetForChat returns value of ForChat conditional field.
func (t *Theme) GetForChat() (value bool) {
	if t == nil {
		return
	}
	return t.Flags.Has(5)
}

// GetID returns value of ID field.
func (t *Theme) GetID() (value int64) {
	if t == nil {
		return
	}
	return t.ID
}

// GetAccessHash returns value of AccessHash field.
func (t *Theme) GetAccessHash() (value int64) {
	if t == nil {
		return
	}
	return t.AccessHash
}

// GetSlug returns value of Slug field.
func (t *Theme) GetSlug() (value string) {
	if t == nil {
		return
	}
	return t.Slug
}

// GetTitle returns value of Title field.
func (t *Theme) GetTitle() (value string) {
	if t == nil {
		return
	}
	return t.Title
}

// SetDocument sets value of Document conditional field.
func (t *Theme) SetDocument(value DocumentClass) {
	t.Flags.Set(2)
	t.Document = value
}

// GetDocument returns value of Document conditional field and
// boolean which is true if field was set.
func (t *Theme) GetDocument() (value DocumentClass, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(2) {
		return value, false
	}
	return t.Document, true
}

// SetSettings sets value of Settings conditional field.
func (t *Theme) SetSettings(value []ThemeSettings) {
	t.Flags.Set(3)
	t.Settings = value
}

// GetSettings returns value of Settings conditional field and
// boolean which is true if field was set.
func (t *Theme) GetSettings() (value []ThemeSettings, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(3) {
		return value, false
	}
	return t.Settings, true
}

// SetEmoticon sets value of Emoticon conditional field.
func (t *Theme) SetEmoticon(value string) {
	t.Flags.Set(6)
	t.Emoticon = value
}

// GetEmoticon returns value of Emoticon conditional field and
// boolean which is true if field was set.
func (t *Theme) GetEmoticon() (value string, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(6) {
		return value, false
	}
	return t.Emoticon, true
}

// SetInstallsCount sets value of InstallsCount conditional field.
func (t *Theme) SetInstallsCount(value int) {
	t.Flags.Set(4)
	t.InstallsCount = value
}

// GetInstallsCount returns value of InstallsCount conditional field and
// boolean which is true if field was set.
func (t *Theme) GetInstallsCount() (value int, ok bool) {
	if t == nil {
		return
	}
	if !t.Flags.Has(4) {
		return value, false
	}
	return t.InstallsCount, true
}

// GetDocumentAsNotEmpty returns mapped value of Document conditional field and
// boolean which is true if field was set.
func (t *Theme) GetDocumentAsNotEmpty() (*Document, bool) {
	if value, ok := t.GetDocument(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}
