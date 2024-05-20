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

// ReactionNotificationSettings represents TL type `reactionNotificationSettings#2bb0f654`.
type ReactionNotificationSettings struct {
	// Source of message reactions for which notifications are shown
	MessageReactionSource ReactionNotificationSourceClass
	// Source of story reactions for which notifications are shown
	StoryReactionSource ReactionNotificationSourceClass
	// Identifier of the notification sound to be played; 0 if sound is disabled
	SoundID int64
	// True, if reaction sender and emoji must be displayed in notifications
	ShowPreview bool
}

// ReactionNotificationSettingsTypeID is TL type id of ReactionNotificationSettings.
const ReactionNotificationSettingsTypeID = 0x2bb0f654

// Ensuring interfaces in compile-time for ReactionNotificationSettings.
var (
	_ bin.Encoder     = &ReactionNotificationSettings{}
	_ bin.Decoder     = &ReactionNotificationSettings{}
	_ bin.BareEncoder = &ReactionNotificationSettings{}
	_ bin.BareDecoder = &ReactionNotificationSettings{}
)

func (r *ReactionNotificationSettings) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.MessageReactionSource == nil) {
		return false
	}
	if !(r.StoryReactionSource == nil) {
		return false
	}
	if !(r.SoundID == 0) {
		return false
	}
	if !(r.ShowPreview == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReactionNotificationSettings) String() string {
	if r == nil {
		return "ReactionNotificationSettings(nil)"
	}
	type Alias ReactionNotificationSettings
	return fmt.Sprintf("ReactionNotificationSettings%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReactionNotificationSettings) TypeID() uint32 {
	return ReactionNotificationSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*ReactionNotificationSettings) TypeName() string {
	return "reactionNotificationSettings"
}

// TypeInfo returns info about TL type.
func (r *ReactionNotificationSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reactionNotificationSettings",
		ID:   ReactionNotificationSettingsTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessageReactionSource",
			SchemaName: "message_reaction_source",
		},
		{
			Name:       "StoryReactionSource",
			SchemaName: "story_reaction_source",
		},
		{
			Name:       "SoundID",
			SchemaName: "sound_id",
		},
		{
			Name:       "ShowPreview",
			SchemaName: "show_preview",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReactionNotificationSettings) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionNotificationSettings#2bb0f654 as nil")
	}
	b.PutID(ReactionNotificationSettingsTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReactionNotificationSettings) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionNotificationSettings#2bb0f654 as nil")
	}
	if r.MessageReactionSource == nil {
		return fmt.Errorf("unable to encode reactionNotificationSettings#2bb0f654: field message_reaction_source is nil")
	}
	if err := r.MessageReactionSource.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reactionNotificationSettings#2bb0f654: field message_reaction_source: %w", err)
	}
	if r.StoryReactionSource == nil {
		return fmt.Errorf("unable to encode reactionNotificationSettings#2bb0f654: field story_reaction_source is nil")
	}
	if err := r.StoryReactionSource.Encode(b); err != nil {
		return fmt.Errorf("unable to encode reactionNotificationSettings#2bb0f654: field story_reaction_source: %w", err)
	}
	b.PutLong(r.SoundID)
	b.PutBool(r.ShowPreview)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReactionNotificationSettings) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionNotificationSettings#2bb0f654 to nil")
	}
	if err := b.ConsumeID(ReactionNotificationSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode reactionNotificationSettings#2bb0f654: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReactionNotificationSettings) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionNotificationSettings#2bb0f654 to nil")
	}
	{
		value, err := DecodeReactionNotificationSource(b)
		if err != nil {
			return fmt.Errorf("unable to decode reactionNotificationSettings#2bb0f654: field message_reaction_source: %w", err)
		}
		r.MessageReactionSource = value
	}
	{
		value, err := DecodeReactionNotificationSource(b)
		if err != nil {
			return fmt.Errorf("unable to decode reactionNotificationSettings#2bb0f654: field story_reaction_source: %w", err)
		}
		r.StoryReactionSource = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode reactionNotificationSettings#2bb0f654: field sound_id: %w", err)
		}
		r.SoundID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode reactionNotificationSettings#2bb0f654: field show_preview: %w", err)
		}
		r.ShowPreview = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReactionNotificationSettings) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reactionNotificationSettings#2bb0f654 as nil")
	}
	b.ObjStart()
	b.PutID("reactionNotificationSettings")
	b.Comma()
	b.FieldStart("message_reaction_source")
	if r.MessageReactionSource == nil {
		return fmt.Errorf("unable to encode reactionNotificationSettings#2bb0f654: field message_reaction_source is nil")
	}
	if err := r.MessageReactionSource.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reactionNotificationSettings#2bb0f654: field message_reaction_source: %w", err)
	}
	b.Comma()
	b.FieldStart("story_reaction_source")
	if r.StoryReactionSource == nil {
		return fmt.Errorf("unable to encode reactionNotificationSettings#2bb0f654: field story_reaction_source is nil")
	}
	if err := r.StoryReactionSource.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode reactionNotificationSettings#2bb0f654: field story_reaction_source: %w", err)
	}
	b.Comma()
	b.FieldStart("sound_id")
	b.PutLong(r.SoundID)
	b.Comma()
	b.FieldStart("show_preview")
	b.PutBool(r.ShowPreview)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ReactionNotificationSettings) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reactionNotificationSettings#2bb0f654 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reactionNotificationSettings"); err != nil {
				return fmt.Errorf("unable to decode reactionNotificationSettings#2bb0f654: %w", err)
			}
		case "message_reaction_source":
			value, err := DecodeTDLibJSONReactionNotificationSource(b)
			if err != nil {
				return fmt.Errorf("unable to decode reactionNotificationSettings#2bb0f654: field message_reaction_source: %w", err)
			}
			r.MessageReactionSource = value
		case "story_reaction_source":
			value, err := DecodeTDLibJSONReactionNotificationSource(b)
			if err != nil {
				return fmt.Errorf("unable to decode reactionNotificationSettings#2bb0f654: field story_reaction_source: %w", err)
			}
			r.StoryReactionSource = value
		case "sound_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode reactionNotificationSettings#2bb0f654: field sound_id: %w", err)
			}
			r.SoundID = value
		case "show_preview":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode reactionNotificationSettings#2bb0f654: field show_preview: %w", err)
			}
			r.ShowPreview = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMessageReactionSource returns value of MessageReactionSource field.
func (r *ReactionNotificationSettings) GetMessageReactionSource() (value ReactionNotificationSourceClass) {
	if r == nil {
		return
	}
	return r.MessageReactionSource
}

// GetStoryReactionSource returns value of StoryReactionSource field.
func (r *ReactionNotificationSettings) GetStoryReactionSource() (value ReactionNotificationSourceClass) {
	if r == nil {
		return
	}
	return r.StoryReactionSource
}

// GetSoundID returns value of SoundID field.
func (r *ReactionNotificationSettings) GetSoundID() (value int64) {
	if r == nil {
		return
	}
	return r.SoundID
}

// GetShowPreview returns value of ShowPreview field.
func (r *ReactionNotificationSettings) GetShowPreview() (value bool) {
	if r == nil {
		return
	}
	return r.ShowPreview
}
