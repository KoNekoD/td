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

// LanguagePackInfo represents TL type `languagePackInfo#20514f5a`.
type LanguagePackInfo struct {
	// Unique language pack identifier
	ID string
	// Identifier of a base language pack; may be empty. If a string is missed in the
	// language pack, then it must be fetched from base language pack. Unsupported in custom
	// language packs
	BaseLanguagePackID string
	// Language name
	Name string
	// Name of the language in that language
	NativeName string
	// A language code to be used to apply plural forms. See https://www.unicode
	// org/cldr/charts/latest/supplemental/language_plural_rules.html for more information
	PluralCode string
	// True, if the language pack is official
	IsOfficial bool
	// True, if the language pack strings are RTL
	IsRtl bool
	// True, if the language pack is a beta language pack
	IsBeta bool
	// True, if the language pack is installed by the current user
	IsInstalled bool
	// Total number of non-deleted strings from the language pack
	TotalStringCount int32
	// Total number of translated strings from the language pack
	TranslatedStringCount int32
	// Total number of non-deleted strings from the language pack available locally
	LocalStringCount int32
	// Link to language translation interface; empty for custom local language packs
	TranslationURL string
}

// LanguagePackInfoTypeID is TL type id of LanguagePackInfo.
const LanguagePackInfoTypeID = 0x20514f5a

// Ensuring interfaces in compile-time for LanguagePackInfo.
var (
	_ bin.Encoder     = &LanguagePackInfo{}
	_ bin.Decoder     = &LanguagePackInfo{}
	_ bin.BareEncoder = &LanguagePackInfo{}
	_ bin.BareDecoder = &LanguagePackInfo{}
)

func (l *LanguagePackInfo) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.ID == "") {
		return false
	}
	if !(l.BaseLanguagePackID == "") {
		return false
	}
	if !(l.Name == "") {
		return false
	}
	if !(l.NativeName == "") {
		return false
	}
	if !(l.PluralCode == "") {
		return false
	}
	if !(l.IsOfficial == false) {
		return false
	}
	if !(l.IsRtl == false) {
		return false
	}
	if !(l.IsBeta == false) {
		return false
	}
	if !(l.IsInstalled == false) {
		return false
	}
	if !(l.TotalStringCount == 0) {
		return false
	}
	if !(l.TranslatedStringCount == 0) {
		return false
	}
	if !(l.LocalStringCount == 0) {
		return false
	}
	if !(l.TranslationURL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LanguagePackInfo) String() string {
	if l == nil {
		return "LanguagePackInfo(nil)"
	}
	type Alias LanguagePackInfo
	return fmt.Sprintf("LanguagePackInfo%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LanguagePackInfo) TypeID() uint32 {
	return LanguagePackInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*LanguagePackInfo) TypeName() string {
	return "languagePackInfo"
}

// TypeInfo returns info about TL type.
func (l *LanguagePackInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "languagePackInfo",
		ID:   LanguagePackInfoTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "BaseLanguagePackID",
			SchemaName: "base_language_pack_id",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "NativeName",
			SchemaName: "native_name",
		},
		{
			Name:       "PluralCode",
			SchemaName: "plural_code",
		},
		{
			Name:       "IsOfficial",
			SchemaName: "is_official",
		},
		{
			Name:       "IsRtl",
			SchemaName: "is_rtl",
		},
		{
			Name:       "IsBeta",
			SchemaName: "is_beta",
		},
		{
			Name:       "IsInstalled",
			SchemaName: "is_installed",
		},
		{
			Name:       "TotalStringCount",
			SchemaName: "total_string_count",
		},
		{
			Name:       "TranslatedStringCount",
			SchemaName: "translated_string_count",
		},
		{
			Name:       "LocalStringCount",
			SchemaName: "local_string_count",
		},
		{
			Name:       "TranslationURL",
			SchemaName: "translation_url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LanguagePackInfo) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackInfo#20514f5a as nil")
	}
	b.PutID(LanguagePackInfoTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LanguagePackInfo) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackInfo#20514f5a as nil")
	}
	b.PutString(l.ID)
	b.PutString(l.BaseLanguagePackID)
	b.PutString(l.Name)
	b.PutString(l.NativeName)
	b.PutString(l.PluralCode)
	b.PutBool(l.IsOfficial)
	b.PutBool(l.IsRtl)
	b.PutBool(l.IsBeta)
	b.PutBool(l.IsInstalled)
	b.PutInt32(l.TotalStringCount)
	b.PutInt32(l.TranslatedStringCount)
	b.PutInt32(l.LocalStringCount)
	b.PutString(l.TranslationURL)
	return nil
}

// Decode implements bin.Decoder.
func (l *LanguagePackInfo) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackInfo#20514f5a to nil")
	}
	if err := b.ConsumeID(LanguagePackInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode languagePackInfo#20514f5a: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LanguagePackInfo) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackInfo#20514f5a to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field id: %w", err)
		}
		l.ID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field base_language_pack_id: %w", err)
		}
		l.BaseLanguagePackID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field name: %w", err)
		}
		l.Name = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field native_name: %w", err)
		}
		l.NativeName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field plural_code: %w", err)
		}
		l.PluralCode = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field is_official: %w", err)
		}
		l.IsOfficial = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field is_rtl: %w", err)
		}
		l.IsRtl = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field is_beta: %w", err)
		}
		l.IsBeta = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field is_installed: %w", err)
		}
		l.IsInstalled = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field total_string_count: %w", err)
		}
		l.TotalStringCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field translated_string_count: %w", err)
		}
		l.TranslatedStringCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field local_string_count: %w", err)
		}
		l.LocalStringCount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field translation_url: %w", err)
		}
		l.TranslationURL = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (l *LanguagePackInfo) EncodeTDLibJSON(b tdjson.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode languagePackInfo#20514f5a as nil")
	}
	b.ObjStart()
	b.PutID("languagePackInfo")
	b.Comma()
	b.FieldStart("id")
	b.PutString(l.ID)
	b.Comma()
	b.FieldStart("base_language_pack_id")
	b.PutString(l.BaseLanguagePackID)
	b.Comma()
	b.FieldStart("name")
	b.PutString(l.Name)
	b.Comma()
	b.FieldStart("native_name")
	b.PutString(l.NativeName)
	b.Comma()
	b.FieldStart("plural_code")
	b.PutString(l.PluralCode)
	b.Comma()
	b.FieldStart("is_official")
	b.PutBool(l.IsOfficial)
	b.Comma()
	b.FieldStart("is_rtl")
	b.PutBool(l.IsRtl)
	b.Comma()
	b.FieldStart("is_beta")
	b.PutBool(l.IsBeta)
	b.Comma()
	b.FieldStart("is_installed")
	b.PutBool(l.IsInstalled)
	b.Comma()
	b.FieldStart("total_string_count")
	b.PutInt32(l.TotalStringCount)
	b.Comma()
	b.FieldStart("translated_string_count")
	b.PutInt32(l.TranslatedStringCount)
	b.Comma()
	b.FieldStart("local_string_count")
	b.PutInt32(l.LocalStringCount)
	b.Comma()
	b.FieldStart("translation_url")
	b.PutString(l.TranslationURL)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (l *LanguagePackInfo) DecodeTDLibJSON(b tdjson.Decoder) error {
	if l == nil {
		return fmt.Errorf("can't decode languagePackInfo#20514f5a to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("languagePackInfo"); err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: %w", err)
			}
		case "id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field id: %w", err)
			}
			l.ID = value
		case "base_language_pack_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field base_language_pack_id: %w", err)
			}
			l.BaseLanguagePackID = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field name: %w", err)
			}
			l.Name = value
		case "native_name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field native_name: %w", err)
			}
			l.NativeName = value
		case "plural_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field plural_code: %w", err)
			}
			l.PluralCode = value
		case "is_official":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field is_official: %w", err)
			}
			l.IsOfficial = value
		case "is_rtl":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field is_rtl: %w", err)
			}
			l.IsRtl = value
		case "is_beta":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field is_beta: %w", err)
			}
			l.IsBeta = value
		case "is_installed":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field is_installed: %w", err)
			}
			l.IsInstalled = value
		case "total_string_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field total_string_count: %w", err)
			}
			l.TotalStringCount = value
		case "translated_string_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field translated_string_count: %w", err)
			}
			l.TranslatedStringCount = value
		case "local_string_count":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field local_string_count: %w", err)
			}
			l.LocalStringCount = value
		case "translation_url":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode languagePackInfo#20514f5a: field translation_url: %w", err)
			}
			l.TranslationURL = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetID returns value of ID field.
func (l *LanguagePackInfo) GetID() (value string) {
	if l == nil {
		return
	}
	return l.ID
}

// GetBaseLanguagePackID returns value of BaseLanguagePackID field.
func (l *LanguagePackInfo) GetBaseLanguagePackID() (value string) {
	if l == nil {
		return
	}
	return l.BaseLanguagePackID
}

// GetName returns value of Name field.
func (l *LanguagePackInfo) GetName() (value string) {
	if l == nil {
		return
	}
	return l.Name
}

// GetNativeName returns value of NativeName field.
func (l *LanguagePackInfo) GetNativeName() (value string) {
	if l == nil {
		return
	}
	return l.NativeName
}

// GetPluralCode returns value of PluralCode field.
func (l *LanguagePackInfo) GetPluralCode() (value string) {
	if l == nil {
		return
	}
	return l.PluralCode
}

// GetIsOfficial returns value of IsOfficial field.
func (l *LanguagePackInfo) GetIsOfficial() (value bool) {
	if l == nil {
		return
	}
	return l.IsOfficial
}

// GetIsRtl returns value of IsRtl field.
func (l *LanguagePackInfo) GetIsRtl() (value bool) {
	if l == nil {
		return
	}
	return l.IsRtl
}

// GetIsBeta returns value of IsBeta field.
func (l *LanguagePackInfo) GetIsBeta() (value bool) {
	if l == nil {
		return
	}
	return l.IsBeta
}

// GetIsInstalled returns value of IsInstalled field.
func (l *LanguagePackInfo) GetIsInstalled() (value bool) {
	if l == nil {
		return
	}
	return l.IsInstalled
}

// GetTotalStringCount returns value of TotalStringCount field.
func (l *LanguagePackInfo) GetTotalStringCount() (value int32) {
	if l == nil {
		return
	}
	return l.TotalStringCount
}

// GetTranslatedStringCount returns value of TranslatedStringCount field.
func (l *LanguagePackInfo) GetTranslatedStringCount() (value int32) {
	if l == nil {
		return
	}
	return l.TranslatedStringCount
}

// GetLocalStringCount returns value of LocalStringCount field.
func (l *LanguagePackInfo) GetLocalStringCount() (value int32) {
	if l == nil {
		return
	}
	return l.LocalStringCount
}

// GetTranslationURL returns value of TranslationURL field.
func (l *LanguagePackInfo) GetTranslationURL() (value string) {
	if l == nil {
		return
	}
	return l.TranslationURL
}
