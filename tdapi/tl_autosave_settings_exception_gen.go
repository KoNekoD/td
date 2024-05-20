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

// AutosaveSettingsException represents TL type `autosaveSettingsException#586bf5c8`.
type AutosaveSettingsException struct {
	// Chat identifier
	ChatID int64
	// Autosave settings for the chat
	Settings ScopeAutosaveSettings
}

// AutosaveSettingsExceptionTypeID is TL type id of AutosaveSettingsException.
const AutosaveSettingsExceptionTypeID = 0x586bf5c8

// Ensuring interfaces in compile-time for AutosaveSettingsException.
var (
	_ bin.Encoder     = &AutosaveSettingsException{}
	_ bin.Decoder     = &AutosaveSettingsException{}
	_ bin.BareEncoder = &AutosaveSettingsException{}
	_ bin.BareDecoder = &AutosaveSettingsException{}
)

func (a *AutosaveSettingsException) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.ChatID == 0) {
		return false
	}
	if !(a.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AutosaveSettingsException) String() string {
	if a == nil {
		return "AutosaveSettingsException(nil)"
	}
	type Alias AutosaveSettingsException
	return fmt.Sprintf("AutosaveSettingsException%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AutosaveSettingsException) TypeID() uint32 {
	return AutosaveSettingsExceptionTypeID
}

// TypeName returns name of type in TL schema.
func (*AutosaveSettingsException) TypeName() string {
	return "autosaveSettingsException"
}

// TypeInfo returns info about TL type.
func (a *AutosaveSettingsException) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "autosaveSettingsException",
		ID:   AutosaveSettingsExceptionTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AutosaveSettingsException) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autosaveSettingsException#586bf5c8 as nil")
	}
	b.PutID(AutosaveSettingsExceptionTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AutosaveSettingsException) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autosaveSettingsException#586bf5c8 as nil")
	}
	b.PutInt53(a.ChatID)
	if err := a.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autosaveSettingsException#586bf5c8: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AutosaveSettingsException) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autosaveSettingsException#586bf5c8 to nil")
	}
	if err := b.ConsumeID(AutosaveSettingsExceptionTypeID); err != nil {
		return fmt.Errorf("unable to decode autosaveSettingsException#586bf5c8: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AutosaveSettingsException) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autosaveSettingsException#586bf5c8 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode autosaveSettingsException#586bf5c8: field chat_id: %w", err)
		}
		a.ChatID = value
	}
	{
		if err := a.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode autosaveSettingsException#586bf5c8: field settings: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AutosaveSettingsException) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode autosaveSettingsException#586bf5c8 as nil")
	}
	b.ObjStart()
	b.PutID("autosaveSettingsException")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(a.ChatID)
	b.Comma()
	b.FieldStart("settings")
	if err := a.Settings.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode autosaveSettingsException#586bf5c8: field settings: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AutosaveSettingsException) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode autosaveSettingsException#586bf5c8 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("autosaveSettingsException"); err != nil {
				return fmt.Errorf("unable to decode autosaveSettingsException#586bf5c8: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode autosaveSettingsException#586bf5c8: field chat_id: %w", err)
			}
			a.ChatID = value
		case "settings":
			if err := a.Settings.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode autosaveSettingsException#586bf5c8: field settings: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (a *AutosaveSettingsException) GetChatID() (value int64) {
	if a == nil {
		return
	}
	return a.ChatID
}

// GetSettings returns value of Settings field.
func (a *AutosaveSettingsException) GetSettings() (value ScopeAutosaveSettings) {
	if a == nil {
		return
	}
	return a.Settings
}
