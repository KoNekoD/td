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

// ChatNotificationSettings represents TL type `chatNotificationSettings#dddc12cf`.
type ChatNotificationSettings struct {
	// If true, the value for the relevant type of chat or the forum chat is used instead of
	// mute_for
	UseDefaultMuteFor bool
	// Time left before notifications will be unmuted, in seconds
	MuteFor int32
	// If true, the value for the relevant type of chat or the forum chat is used instead of
	// sound_id
	UseDefaultSound bool
	// Identifier of the notification sound to be played for messages; 0 if sound is disabled
	SoundID int64
	// If true, the value for the relevant type of chat or the forum chat is used instead of
	// show_preview
	UseDefaultShowPreview bool
	// True, if message content must be displayed in notifications
	ShowPreview bool
	// If true, the value for the relevant type of chat is used instead of mute_stories
	UseDefaultMuteStories bool
	// True, if story notifications are disabled for the chat
	MuteStories bool
	// If true, the value for the relevant type of chat is used instead of story_sound_id
	UseDefaultStorySound bool
	// Identifier of the notification sound to be played for stories; 0 if sound is disabled
	StorySoundID int64
	// If true, the value for the relevant type of chat is used instead of show_story_sender
	UseDefaultShowStorySender bool
	// True, if the sender of stories must be displayed in notifications
	ShowStorySender bool
	// If true, the value for the relevant type of chat or the forum chat is used instead of
	// disable_pinned_message_notifications
	UseDefaultDisablePinnedMessageNotifications bool
	// If true, notifications for incoming pinned messages will be created as for an ordinary
	// unread message
	DisablePinnedMessageNotifications bool
	// If true, the value for the relevant type of chat or the forum chat is used instead of
	// disable_mention_notifications
	UseDefaultDisableMentionNotifications bool
	// If true, notifications for messages with mentions will be created as for an ordinary
	// unread message
	DisableMentionNotifications bool
}

// ChatNotificationSettingsTypeID is TL type id of ChatNotificationSettings.
const ChatNotificationSettingsTypeID = 0xdddc12cf

// Ensuring interfaces in compile-time for ChatNotificationSettings.
var (
	_ bin.Encoder     = &ChatNotificationSettings{}
	_ bin.Decoder     = &ChatNotificationSettings{}
	_ bin.BareEncoder = &ChatNotificationSettings{}
	_ bin.BareDecoder = &ChatNotificationSettings{}
)

func (c *ChatNotificationSettings) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UseDefaultMuteFor == false) {
		return false
	}
	if !(c.MuteFor == 0) {
		return false
	}
	if !(c.UseDefaultSound == false) {
		return false
	}
	if !(c.SoundID == 0) {
		return false
	}
	if !(c.UseDefaultShowPreview == false) {
		return false
	}
	if !(c.ShowPreview == false) {
		return false
	}
	if !(c.UseDefaultMuteStories == false) {
		return false
	}
	if !(c.MuteStories == false) {
		return false
	}
	if !(c.UseDefaultStorySound == false) {
		return false
	}
	if !(c.StorySoundID == 0) {
		return false
	}
	if !(c.UseDefaultShowStorySender == false) {
		return false
	}
	if !(c.ShowStorySender == false) {
		return false
	}
	if !(c.UseDefaultDisablePinnedMessageNotifications == false) {
		return false
	}
	if !(c.DisablePinnedMessageNotifications == false) {
		return false
	}
	if !(c.UseDefaultDisableMentionNotifications == false) {
		return false
	}
	if !(c.DisableMentionNotifications == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatNotificationSettings) String() string {
	if c == nil {
		return "ChatNotificationSettings(nil)"
	}
	type Alias ChatNotificationSettings
	return fmt.Sprintf("ChatNotificationSettings%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatNotificationSettings) TypeID() uint32 {
	return ChatNotificationSettingsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatNotificationSettings) TypeName() string {
	return "chatNotificationSettings"
}

// TypeInfo returns info about TL type.
func (c *ChatNotificationSettings) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatNotificationSettings",
		ID:   ChatNotificationSettingsTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UseDefaultMuteFor",
			SchemaName: "use_default_mute_for",
		},
		{
			Name:       "MuteFor",
			SchemaName: "mute_for",
		},
		{
			Name:       "UseDefaultSound",
			SchemaName: "use_default_sound",
		},
		{
			Name:       "SoundID",
			SchemaName: "sound_id",
		},
		{
			Name:       "UseDefaultShowPreview",
			SchemaName: "use_default_show_preview",
		},
		{
			Name:       "ShowPreview",
			SchemaName: "show_preview",
		},
		{
			Name:       "UseDefaultMuteStories",
			SchemaName: "use_default_mute_stories",
		},
		{
			Name:       "MuteStories",
			SchemaName: "mute_stories",
		},
		{
			Name:       "UseDefaultStorySound",
			SchemaName: "use_default_story_sound",
		},
		{
			Name:       "StorySoundID",
			SchemaName: "story_sound_id",
		},
		{
			Name:       "UseDefaultShowStorySender",
			SchemaName: "use_default_show_story_sender",
		},
		{
			Name:       "ShowStorySender",
			SchemaName: "show_story_sender",
		},
		{
			Name:       "UseDefaultDisablePinnedMessageNotifications",
			SchemaName: "use_default_disable_pinned_message_notifications",
		},
		{
			Name:       "DisablePinnedMessageNotifications",
			SchemaName: "disable_pinned_message_notifications",
		},
		{
			Name:       "UseDefaultDisableMentionNotifications",
			SchemaName: "use_default_disable_mention_notifications",
		},
		{
			Name:       "DisableMentionNotifications",
			SchemaName: "disable_mention_notifications",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatNotificationSettings) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatNotificationSettings#dddc12cf as nil")
	}
	b.PutID(ChatNotificationSettingsTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatNotificationSettings) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatNotificationSettings#dddc12cf as nil")
	}
	b.PutBool(c.UseDefaultMuteFor)
	b.PutInt32(c.MuteFor)
	b.PutBool(c.UseDefaultSound)
	b.PutLong(c.SoundID)
	b.PutBool(c.UseDefaultShowPreview)
	b.PutBool(c.ShowPreview)
	b.PutBool(c.UseDefaultMuteStories)
	b.PutBool(c.MuteStories)
	b.PutBool(c.UseDefaultStorySound)
	b.PutLong(c.StorySoundID)
	b.PutBool(c.UseDefaultShowStorySender)
	b.PutBool(c.ShowStorySender)
	b.PutBool(c.UseDefaultDisablePinnedMessageNotifications)
	b.PutBool(c.DisablePinnedMessageNotifications)
	b.PutBool(c.UseDefaultDisableMentionNotifications)
	b.PutBool(c.DisableMentionNotifications)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatNotificationSettings) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatNotificationSettings#dddc12cf to nil")
	}
	if err := b.ConsumeID(ChatNotificationSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatNotificationSettings) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatNotificationSettings#dddc12cf to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_mute_for: %w", err)
		}
		c.UseDefaultMuteFor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field mute_for: %w", err)
		}
		c.MuteFor = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_sound: %w", err)
		}
		c.UseDefaultSound = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field sound_id: %w", err)
		}
		c.SoundID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_show_preview: %w", err)
		}
		c.UseDefaultShowPreview = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field show_preview: %w", err)
		}
		c.ShowPreview = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_mute_stories: %w", err)
		}
		c.UseDefaultMuteStories = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field mute_stories: %w", err)
		}
		c.MuteStories = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_story_sound: %w", err)
		}
		c.UseDefaultStorySound = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field story_sound_id: %w", err)
		}
		c.StorySoundID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_show_story_sender: %w", err)
		}
		c.UseDefaultShowStorySender = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field show_story_sender: %w", err)
		}
		c.ShowStorySender = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_disable_pinned_message_notifications: %w", err)
		}
		c.UseDefaultDisablePinnedMessageNotifications = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field disable_pinned_message_notifications: %w", err)
		}
		c.DisablePinnedMessageNotifications = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_disable_mention_notifications: %w", err)
		}
		c.UseDefaultDisableMentionNotifications = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field disable_mention_notifications: %w", err)
		}
		c.DisableMentionNotifications = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatNotificationSettings) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatNotificationSettings#dddc12cf as nil")
	}
	b.ObjStart()
	b.PutID("chatNotificationSettings")
	b.Comma()
	b.FieldStart("use_default_mute_for")
	b.PutBool(c.UseDefaultMuteFor)
	b.Comma()
	b.FieldStart("mute_for")
	b.PutInt32(c.MuteFor)
	b.Comma()
	b.FieldStart("use_default_sound")
	b.PutBool(c.UseDefaultSound)
	b.Comma()
	b.FieldStart("sound_id")
	b.PutLong(c.SoundID)
	b.Comma()
	b.FieldStart("use_default_show_preview")
	b.PutBool(c.UseDefaultShowPreview)
	b.Comma()
	b.FieldStart("show_preview")
	b.PutBool(c.ShowPreview)
	b.Comma()
	b.FieldStart("use_default_mute_stories")
	b.PutBool(c.UseDefaultMuteStories)
	b.Comma()
	b.FieldStart("mute_stories")
	b.PutBool(c.MuteStories)
	b.Comma()
	b.FieldStart("use_default_story_sound")
	b.PutBool(c.UseDefaultStorySound)
	b.Comma()
	b.FieldStart("story_sound_id")
	b.PutLong(c.StorySoundID)
	b.Comma()
	b.FieldStart("use_default_show_story_sender")
	b.PutBool(c.UseDefaultShowStorySender)
	b.Comma()
	b.FieldStart("show_story_sender")
	b.PutBool(c.ShowStorySender)
	b.Comma()
	b.FieldStart("use_default_disable_pinned_message_notifications")
	b.PutBool(c.UseDefaultDisablePinnedMessageNotifications)
	b.Comma()
	b.FieldStart("disable_pinned_message_notifications")
	b.PutBool(c.DisablePinnedMessageNotifications)
	b.Comma()
	b.FieldStart("use_default_disable_mention_notifications")
	b.PutBool(c.UseDefaultDisableMentionNotifications)
	b.Comma()
	b.FieldStart("disable_mention_notifications")
	b.PutBool(c.DisableMentionNotifications)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatNotificationSettings) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatNotificationSettings#dddc12cf to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatNotificationSettings"); err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: %w", err)
			}
		case "use_default_mute_for":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_mute_for: %w", err)
			}
			c.UseDefaultMuteFor = value
		case "mute_for":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field mute_for: %w", err)
			}
			c.MuteFor = value
		case "use_default_sound":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_sound: %w", err)
			}
			c.UseDefaultSound = value
		case "sound_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field sound_id: %w", err)
			}
			c.SoundID = value
		case "use_default_show_preview":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_show_preview: %w", err)
			}
			c.UseDefaultShowPreview = value
		case "show_preview":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field show_preview: %w", err)
			}
			c.ShowPreview = value
		case "use_default_mute_stories":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_mute_stories: %w", err)
			}
			c.UseDefaultMuteStories = value
		case "mute_stories":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field mute_stories: %w", err)
			}
			c.MuteStories = value
		case "use_default_story_sound":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_story_sound: %w", err)
			}
			c.UseDefaultStorySound = value
		case "story_sound_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field story_sound_id: %w", err)
			}
			c.StorySoundID = value
		case "use_default_show_story_sender":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_show_story_sender: %w", err)
			}
			c.UseDefaultShowStorySender = value
		case "show_story_sender":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field show_story_sender: %w", err)
			}
			c.ShowStorySender = value
		case "use_default_disable_pinned_message_notifications":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_disable_pinned_message_notifications: %w", err)
			}
			c.UseDefaultDisablePinnedMessageNotifications = value
		case "disable_pinned_message_notifications":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field disable_pinned_message_notifications: %w", err)
			}
			c.DisablePinnedMessageNotifications = value
		case "use_default_disable_mention_notifications":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field use_default_disable_mention_notifications: %w", err)
			}
			c.UseDefaultDisableMentionNotifications = value
		case "disable_mention_notifications":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatNotificationSettings#dddc12cf: field disable_mention_notifications: %w", err)
			}
			c.DisableMentionNotifications = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUseDefaultMuteFor returns value of UseDefaultMuteFor field.
func (c *ChatNotificationSettings) GetUseDefaultMuteFor() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultMuteFor
}

// GetMuteFor returns value of MuteFor field.
func (c *ChatNotificationSettings) GetMuteFor() (value int32) {
	if c == nil {
		return
	}
	return c.MuteFor
}

// GetUseDefaultSound returns value of UseDefaultSound field.
func (c *ChatNotificationSettings) GetUseDefaultSound() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultSound
}

// GetSoundID returns value of SoundID field.
func (c *ChatNotificationSettings) GetSoundID() (value int64) {
	if c == nil {
		return
	}
	return c.SoundID
}

// GetUseDefaultShowPreview returns value of UseDefaultShowPreview field.
func (c *ChatNotificationSettings) GetUseDefaultShowPreview() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultShowPreview
}

// GetShowPreview returns value of ShowPreview field.
func (c *ChatNotificationSettings) GetShowPreview() (value bool) {
	if c == nil {
		return
	}
	return c.ShowPreview
}

// GetUseDefaultMuteStories returns value of UseDefaultMuteStories field.
func (c *ChatNotificationSettings) GetUseDefaultMuteStories() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultMuteStories
}

// GetMuteStories returns value of MuteStories field.
func (c *ChatNotificationSettings) GetMuteStories() (value bool) {
	if c == nil {
		return
	}
	return c.MuteStories
}

// GetUseDefaultStorySound returns value of UseDefaultStorySound field.
func (c *ChatNotificationSettings) GetUseDefaultStorySound() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultStorySound
}

// GetStorySoundID returns value of StorySoundID field.
func (c *ChatNotificationSettings) GetStorySoundID() (value int64) {
	if c == nil {
		return
	}
	return c.StorySoundID
}

// GetUseDefaultShowStorySender returns value of UseDefaultShowStorySender field.
func (c *ChatNotificationSettings) GetUseDefaultShowStorySender() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultShowStorySender
}

// GetShowStorySender returns value of ShowStorySender field.
func (c *ChatNotificationSettings) GetShowStorySender() (value bool) {
	if c == nil {
		return
	}
	return c.ShowStorySender
}

// GetUseDefaultDisablePinnedMessageNotifications returns value of UseDefaultDisablePinnedMessageNotifications field.
func (c *ChatNotificationSettings) GetUseDefaultDisablePinnedMessageNotifications() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultDisablePinnedMessageNotifications
}

// GetDisablePinnedMessageNotifications returns value of DisablePinnedMessageNotifications field.
func (c *ChatNotificationSettings) GetDisablePinnedMessageNotifications() (value bool) {
	if c == nil {
		return
	}
	return c.DisablePinnedMessageNotifications
}

// GetUseDefaultDisableMentionNotifications returns value of UseDefaultDisableMentionNotifications field.
func (c *ChatNotificationSettings) GetUseDefaultDisableMentionNotifications() (value bool) {
	if c == nil {
		return
	}
	return c.UseDefaultDisableMentionNotifications
}

// GetDisableMentionNotifications returns value of DisableMentionNotifications field.
func (c *ChatNotificationSettings) GetDisableMentionNotifications() (value bool) {
	if c == nil {
		return
	}
	return c.DisableMentionNotifications
}
