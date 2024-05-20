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

// StoryFullID represents TL type `storyFullId#701d31f5`.
type StoryFullID struct {
	// Identifier of the chat that posted the story
	SenderChatID int64
	// Unique story identifier among stories of the given sender
	StoryID int32
}

// StoryFullIDTypeID is TL type id of StoryFullID.
const StoryFullIDTypeID = 0x701d31f5

// Ensuring interfaces in compile-time for StoryFullID.
var (
	_ bin.Encoder     = &StoryFullID{}
	_ bin.Decoder     = &StoryFullID{}
	_ bin.BareEncoder = &StoryFullID{}
	_ bin.BareDecoder = &StoryFullID{}
)

func (s *StoryFullID) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.SenderChatID == 0) {
		return false
	}
	if !(s.StoryID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StoryFullID) String() string {
	if s == nil {
		return "StoryFullID(nil)"
	}
	type Alias StoryFullID
	return fmt.Sprintf("StoryFullID%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StoryFullID) TypeID() uint32 {
	return StoryFullIDTypeID
}

// TypeName returns name of type in TL schema.
func (*StoryFullID) TypeName() string {
	return "storyFullId"
}

// TypeInfo returns info about TL type.
func (s *StoryFullID) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "storyFullId",
		ID:   StoryFullIDTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SenderChatID",
			SchemaName: "sender_chat_id",
		},
		{
			Name:       "StoryID",
			SchemaName: "story_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StoryFullID) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyFullId#701d31f5 as nil")
	}
	b.PutID(StoryFullIDTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StoryFullID) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode storyFullId#701d31f5 as nil")
	}
	b.PutInt53(s.SenderChatID)
	b.PutInt32(s.StoryID)
	return nil
}

// Decode implements bin.Decoder.
func (s *StoryFullID) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyFullId#701d31f5 to nil")
	}
	if err := b.ConsumeID(StoryFullIDTypeID); err != nil {
		return fmt.Errorf("unable to decode storyFullId#701d31f5: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StoryFullID) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode storyFullId#701d31f5 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode storyFullId#701d31f5: field sender_chat_id: %w", err)
		}
		s.SenderChatID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode storyFullId#701d31f5: field story_id: %w", err)
		}
		s.StoryID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StoryFullID) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode storyFullId#701d31f5 as nil")
	}
	b.ObjStart()
	b.PutID("storyFullId")
	b.Comma()
	b.FieldStart("sender_chat_id")
	b.PutInt53(s.SenderChatID)
	b.Comma()
	b.FieldStart("story_id")
	b.PutInt32(s.StoryID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StoryFullID) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode storyFullId#701d31f5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("storyFullId"); err != nil {
				return fmt.Errorf("unable to decode storyFullId#701d31f5: %w", err)
			}
		case "sender_chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode storyFullId#701d31f5: field sender_chat_id: %w", err)
			}
			s.SenderChatID = value
		case "story_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode storyFullId#701d31f5: field story_id: %w", err)
			}
			s.StoryID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSenderChatID returns value of SenderChatID field.
func (s *StoryFullID) GetSenderChatID() (value int64) {
	if s == nil {
		return
	}
	return s.SenderChatID
}

// GetStoryID returns value of StoryID field.
func (s *StoryFullID) GetStoryID() (value int32) {
	if s == nil {
		return
	}
	return s.StoryID
}
