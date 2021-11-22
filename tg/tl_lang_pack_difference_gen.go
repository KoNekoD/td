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

// LangPackDifference represents TL type `langPackDifference#f385c1f6`.
// Changes to the app's localization pack
//
// See https://core.telegram.org/constructor/langPackDifference for reference.
type LangPackDifference struct {
	// Language code
	LangCode string
	// Previous version number
	FromVersion int
	// New version number
	Version int
	// Localized strings
	Strings []LangPackStringClass
}

// LangPackDifferenceTypeID is TL type id of LangPackDifference.
const LangPackDifferenceTypeID = 0xf385c1f6

// Ensuring interfaces in compile-time for LangPackDifference.
var (
	_ bin.Encoder     = &LangPackDifference{}
	_ bin.Decoder     = &LangPackDifference{}
	_ bin.BareEncoder = &LangPackDifference{}
	_ bin.BareDecoder = &LangPackDifference{}
)

func (l *LangPackDifference) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.LangCode == "") {
		return false
	}
	if !(l.FromVersion == 0) {
		return false
	}
	if !(l.Version == 0) {
		return false
	}
	if !(l.Strings == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LangPackDifference) String() string {
	if l == nil {
		return "LangPackDifference(nil)"
	}
	type Alias LangPackDifference
	return fmt.Sprintf("LangPackDifference%+v", Alias(*l))
}

// FillFrom fills LangPackDifference from given interface.
func (l *LangPackDifference) FillFrom(from interface {
	GetLangCode() (value string)
	GetFromVersion() (value int)
	GetVersion() (value int)
	GetStrings() (value []LangPackStringClass)
}) {
	l.LangCode = from.GetLangCode()
	l.FromVersion = from.GetFromVersion()
	l.Version = from.GetVersion()
	l.Strings = from.GetStrings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LangPackDifference) TypeID() uint32 {
	return LangPackDifferenceTypeID
}

// TypeName returns name of type in TL schema.
func (*LangPackDifference) TypeName() string {
	return "langPackDifference"
}

// TypeInfo returns info about TL type.
func (l *LangPackDifference) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "langPackDifference",
		ID:   LangPackDifferenceTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
		{
			Name:       "FromVersion",
			SchemaName: "from_version",
		},
		{
			Name:       "Version",
			SchemaName: "version",
		},
		{
			Name:       "Strings",
			SchemaName: "strings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LangPackDifference) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode langPackDifference#f385c1f6 as nil")
	}
	b.PutID(LangPackDifferenceTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LangPackDifference) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode langPackDifference#f385c1f6 as nil")
	}
	b.PutString(l.LangCode)
	b.PutInt(l.FromVersion)
	b.PutInt(l.Version)
	b.PutVectorHeader(len(l.Strings))
	for idx, v := range l.Strings {
		if v == nil {
			return fmt.Errorf("unable to encode langPackDifference#f385c1f6: field strings element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode langPackDifference#f385c1f6: field strings element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *LangPackDifference) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode langPackDifference#f385c1f6 to nil")
	}
	if err := b.ConsumeID(LangPackDifferenceTypeID); err != nil {
		return fmt.Errorf("unable to decode langPackDifference#f385c1f6: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LangPackDifference) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode langPackDifference#f385c1f6 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langPackDifference#f385c1f6: field lang_code: %w", err)
		}
		l.LangCode = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode langPackDifference#f385c1f6: field from_version: %w", err)
		}
		l.FromVersion = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode langPackDifference#f385c1f6: field version: %w", err)
		}
		l.Version = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode langPackDifference#f385c1f6: field strings: %w", err)
		}

		if headerLen > 0 {
			l.Strings = make([]LangPackStringClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeLangPackString(b)
			if err != nil {
				return fmt.Errorf("unable to decode langPackDifference#f385c1f6: field strings: %w", err)
			}
			l.Strings = append(l.Strings, value)
		}
	}
	return nil
}

// GetLangCode returns value of LangCode field.
func (l *LangPackDifference) GetLangCode() (value string) {
	return l.LangCode
}

// GetFromVersion returns value of FromVersion field.
func (l *LangPackDifference) GetFromVersion() (value int) {
	return l.FromVersion
}

// GetVersion returns value of Version field.
func (l *LangPackDifference) GetVersion() (value int) {
	return l.Version
}

// GetStrings returns value of Strings field.
func (l *LangPackDifference) GetStrings() (value []LangPackStringClass) {
	return l.Strings
}

// MapStrings returns field Strings wrapped in LangPackStringClassArray helper.
func (l *LangPackDifference) MapStrings() (value LangPackStringClassArray) {
	return LangPackStringClassArray(l.Strings)
}
