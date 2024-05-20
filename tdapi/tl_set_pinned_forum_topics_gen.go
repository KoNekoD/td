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

// SetPinnedForumTopicsRequest represents TL type `setPinnedForumTopics#ddf26f21`.
type SetPinnedForumTopicsRequest struct {
	// Chat identifier
	ChatID int64
	// The new list of pinned forum topics
	MessageThreadIDs []int64
}

// SetPinnedForumTopicsRequestTypeID is TL type id of SetPinnedForumTopicsRequest.
const SetPinnedForumTopicsRequestTypeID = 0xddf26f21

// Ensuring interfaces in compile-time for SetPinnedForumTopicsRequest.
var (
	_ bin.Encoder     = &SetPinnedForumTopicsRequest{}
	_ bin.Decoder     = &SetPinnedForumTopicsRequest{}
	_ bin.BareEncoder = &SetPinnedForumTopicsRequest{}
	_ bin.BareDecoder = &SetPinnedForumTopicsRequest{}
)

func (s *SetPinnedForumTopicsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageThreadIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetPinnedForumTopicsRequest) String() string {
	if s == nil {
		return "SetPinnedForumTopicsRequest(nil)"
	}
	type Alias SetPinnedForumTopicsRequest
	return fmt.Sprintf("SetPinnedForumTopicsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetPinnedForumTopicsRequest) TypeID() uint32 {
	return SetPinnedForumTopicsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetPinnedForumTopicsRequest) TypeName() string {
	return "setPinnedForumTopics"
}

// TypeInfo returns info about TL type.
func (s *SetPinnedForumTopicsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setPinnedForumTopics",
		ID:   SetPinnedForumTopicsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadIDs",
			SchemaName: "message_thread_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetPinnedForumTopicsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setPinnedForumTopics#ddf26f21 as nil")
	}
	b.PutID(SetPinnedForumTopicsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetPinnedForumTopicsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setPinnedForumTopics#ddf26f21 as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt(len(s.MessageThreadIDs))
	for _, v := range s.MessageThreadIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetPinnedForumTopicsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setPinnedForumTopics#ddf26f21 to nil")
	}
	if err := b.ConsumeID(SetPinnedForumTopicsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setPinnedForumTopics#ddf26f21: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetPinnedForumTopicsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setPinnedForumTopics#ddf26f21 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setPinnedForumTopics#ddf26f21: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode setPinnedForumTopics#ddf26f21: field message_thread_ids: %w", err)
		}

		if headerLen > 0 {
			s.MessageThreadIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setPinnedForumTopics#ddf26f21: field message_thread_ids: %w", err)
			}
			s.MessageThreadIDs = append(s.MessageThreadIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetPinnedForumTopicsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setPinnedForumTopics#ddf26f21 as nil")
	}
	b.ObjStart()
	b.PutID("setPinnedForumTopics")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("message_thread_ids")
	b.ArrStart()
	for _, v := range s.MessageThreadIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetPinnedForumTopicsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setPinnedForumTopics#ddf26f21 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setPinnedForumTopics"); err != nil {
				return fmt.Errorf("unable to decode setPinnedForumTopics#ddf26f21: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setPinnedForumTopics#ddf26f21: field chat_id: %w", err)
			}
			s.ChatID = value
		case "message_thread_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode setPinnedForumTopics#ddf26f21: field message_thread_ids: %w", err)
				}
				s.MessageThreadIDs = append(s.MessageThreadIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode setPinnedForumTopics#ddf26f21: field message_thread_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SetPinnedForumTopicsRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetMessageThreadIDs returns value of MessageThreadIDs field.
func (s *SetPinnedForumTopicsRequest) GetMessageThreadIDs() (value []int64) {
	if s == nil {
		return
	}
	return s.MessageThreadIDs
}

// SetPinnedForumTopics invokes method setPinnedForumTopics#ddf26f21 returning error if any.
func (c *Client) SetPinnedForumTopics(ctx context.Context, request *SetPinnedForumTopicsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
