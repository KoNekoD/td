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

// MessagesSaveDraftRequest represents TL type `messages.saveDraft#7ff3b806`.
// Save a message draft¹ associated to a chat.
//
// Links:
//  1. https://core.telegram.org/api/drafts
//
// See https://core.telegram.org/method/messages.saveDraft for reference.
type MessagesSaveDraftRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Disable generation of the webpage preview
	NoWebpage bool
	// If set, any eventual webpage preview will be shown on top of the message instead of at
	// the bottom.
	InvertMedia bool
	// If set, indicates that the message should be sent in reply to the specified message or
	// story.
	//
	// Use SetReplyTo and GetReplyTo helpers.
	ReplyTo InputReplyToClass
	// Destination of the message that should be sent
	Peer InputPeerClass
	// The draft
	Message string
	// Message entities¹ for styled text
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
	// Attached media
	//
	// Use SetMedia and GetMedia helpers.
	Media InputMediaClass
}

// MessagesSaveDraftRequestTypeID is TL type id of MessagesSaveDraftRequest.
const MessagesSaveDraftRequestTypeID = 0x7ff3b806

// Ensuring interfaces in compile-time for MessagesSaveDraftRequest.
var (
	_ bin.Encoder     = &MessagesSaveDraftRequest{}
	_ bin.Decoder     = &MessagesSaveDraftRequest{}
	_ bin.BareEncoder = &MessagesSaveDraftRequest{}
	_ bin.BareDecoder = &MessagesSaveDraftRequest{}
)

func (s *MessagesSaveDraftRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.NoWebpage == false) {
		return false
	}
	if !(s.InvertMedia == false) {
		return false
	}
	if !(s.ReplyTo == nil) {
		return false
	}
	if !(s.Peer == nil) {
		return false
	}
	if !(s.Message == "") {
		return false
	}
	if !(s.Entities == nil) {
		return false
	}
	if !(s.Media == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *MessagesSaveDraftRequest) String() string {
	if s == nil {
		return "MessagesSaveDraftRequest(nil)"
	}
	type Alias MessagesSaveDraftRequest
	return fmt.Sprintf("MessagesSaveDraftRequest%+v", Alias(*s))
}

// FillFrom fills MessagesSaveDraftRequest from given interface.
func (s *MessagesSaveDraftRequest) FillFrom(from interface {
	GetNoWebpage() (value bool)
	GetInvertMedia() (value bool)
	GetReplyTo() (value InputReplyToClass, ok bool)
	GetPeer() (value InputPeerClass)
	GetMessage() (value string)
	GetEntities() (value []MessageEntityClass, ok bool)
	GetMedia() (value InputMediaClass, ok bool)
}) {
	s.NoWebpage = from.GetNoWebpage()
	s.InvertMedia = from.GetInvertMedia()
	if val, ok := from.GetReplyTo(); ok {
		s.ReplyTo = val
	}

	s.Peer = from.GetPeer()
	s.Message = from.GetMessage()
	if val, ok := from.GetEntities(); ok {
		s.Entities = val
	}

	if val, ok := from.GetMedia(); ok {
		s.Media = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesSaveDraftRequest) TypeID() uint32 {
	return MessagesSaveDraftRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesSaveDraftRequest) TypeName() string {
	return "messages.saveDraft"
}

// TypeInfo returns info about TL type.
func (s *MessagesSaveDraftRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.saveDraft",
		ID:   MessagesSaveDraftRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NoWebpage",
			SchemaName: "no_webpage",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "InvertMedia",
			SchemaName: "invert_media",
			Null:       !s.Flags.Has(6),
		},
		{
			Name:       "ReplyTo",
			SchemaName: "reply_to",
			Null:       !s.Flags.Has(4),
		},
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Message",
			SchemaName: "message",
		},
		{
			Name:       "Entities",
			SchemaName: "entities",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "Media",
			SchemaName: "media",
			Null:       !s.Flags.Has(5),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (s *MessagesSaveDraftRequest) SetFlags() {
	if !(s.NoWebpage == false) {
		s.Flags.Set(1)
	}
	if !(s.InvertMedia == false) {
		s.Flags.Set(6)
	}
	if !(s.ReplyTo == nil) {
		s.Flags.Set(4)
	}
	if !(s.Entities == nil) {
		s.Flags.Set(3)
	}
	if !(s.Media == nil) {
		s.Flags.Set(5)
	}
}

// Encode implements bin.Encoder.
func (s *MessagesSaveDraftRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.saveDraft#7ff3b806 as nil")
	}
	b.PutID(MessagesSaveDraftRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *MessagesSaveDraftRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode messages.saveDraft#7ff3b806 as nil")
	}
	s.SetFlags()
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.saveDraft#7ff3b806: field flags: %w", err)
	}
	if s.Flags.Has(4) {
		if s.ReplyTo == nil {
			return fmt.Errorf("unable to encode messages.saveDraft#7ff3b806: field reply_to is nil")
		}
		if err := s.ReplyTo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.saveDraft#7ff3b806: field reply_to: %w", err)
		}
	}
	if s.Peer == nil {
		return fmt.Errorf("unable to encode messages.saveDraft#7ff3b806: field peer is nil")
	}
	if err := s.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.saveDraft#7ff3b806: field peer: %w", err)
	}
	b.PutString(s.Message)
	if s.Flags.Has(3) {
		b.PutVectorHeader(len(s.Entities))
		for idx, v := range s.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode messages.saveDraft#7ff3b806: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode messages.saveDraft#7ff3b806: field entities element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(5) {
		if s.Media == nil {
			return fmt.Errorf("unable to encode messages.saveDraft#7ff3b806: field media is nil")
		}
		if err := s.Media.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.saveDraft#7ff3b806: field media: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *MessagesSaveDraftRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.saveDraft#7ff3b806 to nil")
	}
	if err := b.ConsumeID(MessagesSaveDraftRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.saveDraft#7ff3b806: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *MessagesSaveDraftRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode messages.saveDraft#7ff3b806 to nil")
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#7ff3b806: field flags: %w", err)
		}
	}
	s.NoWebpage = s.Flags.Has(1)
	s.InvertMedia = s.Flags.Has(6)
	if s.Flags.Has(4) {
		value, err := DecodeInputReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#7ff3b806: field reply_to: %w", err)
		}
		s.ReplyTo = value
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#7ff3b806: field peer: %w", err)
		}
		s.Peer = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#7ff3b806: field message: %w", err)
		}
		s.Message = value
	}
	if s.Flags.Has(3) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#7ff3b806: field entities: %w", err)
		}

		if headerLen > 0 {
			s.Entities = make([]MessageEntityClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.saveDraft#7ff3b806: field entities: %w", err)
			}
			s.Entities = append(s.Entities, value)
		}
	}
	if s.Flags.Has(5) {
		value, err := DecodeInputMedia(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.saveDraft#7ff3b806: field media: %w", err)
		}
		s.Media = value
	}
	return nil
}

// SetNoWebpage sets value of NoWebpage conditional field.
func (s *MessagesSaveDraftRequest) SetNoWebpage(value bool) {
	if value {
		s.Flags.Set(1)
		s.NoWebpage = true
	} else {
		s.Flags.Unset(1)
		s.NoWebpage = false
	}
}

// GetNoWebpage returns value of NoWebpage conditional field.
func (s *MessagesSaveDraftRequest) GetNoWebpage() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(1)
}

// SetInvertMedia sets value of InvertMedia conditional field.
func (s *MessagesSaveDraftRequest) SetInvertMedia(value bool) {
	if value {
		s.Flags.Set(6)
		s.InvertMedia = true
	} else {
		s.Flags.Unset(6)
		s.InvertMedia = false
	}
}

// GetInvertMedia returns value of InvertMedia conditional field.
func (s *MessagesSaveDraftRequest) GetInvertMedia() (value bool) {
	if s == nil {
		return
	}
	return s.Flags.Has(6)
}

// SetReplyTo sets value of ReplyTo conditional field.
func (s *MessagesSaveDraftRequest) SetReplyTo(value InputReplyToClass) {
	s.Flags.Set(4)
	s.ReplyTo = value
}

// GetReplyTo returns value of ReplyTo conditional field and
// boolean which is true if field was set.
func (s *MessagesSaveDraftRequest) GetReplyTo() (value InputReplyToClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(4) {
		return value, false
	}
	return s.ReplyTo, true
}

// GetPeer returns value of Peer field.
func (s *MessagesSaveDraftRequest) GetPeer() (value InputPeerClass) {
	if s == nil {
		return
	}
	return s.Peer
}

// GetMessage returns value of Message field.
func (s *MessagesSaveDraftRequest) GetMessage() (value string) {
	if s == nil {
		return
	}
	return s.Message
}

// SetEntities sets value of Entities conditional field.
func (s *MessagesSaveDraftRequest) SetEntities(value []MessageEntityClass) {
	s.Flags.Set(3)
	s.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (s *MessagesSaveDraftRequest) GetEntities() (value []MessageEntityClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Entities, true
}

// SetMedia sets value of Media conditional field.
func (s *MessagesSaveDraftRequest) SetMedia(value InputMediaClass) {
	s.Flags.Set(5)
	s.Media = value
}

// GetMedia returns value of Media conditional field and
// boolean which is true if field was set.
func (s *MessagesSaveDraftRequest) GetMedia() (value InputMediaClass, ok bool) {
	if s == nil {
		return
	}
	if !s.Flags.Has(5) {
		return value, false
	}
	return s.Media, true
}

// MapEntities returns field Entities wrapped in MessageEntityClassArray helper.
func (s *MessagesSaveDraftRequest) MapEntities() (value MessageEntityClassArray, ok bool) {
	if !s.Flags.Has(3) {
		return value, false
	}
	return MessageEntityClassArray(s.Entities), true
}

// MessagesSaveDraft invokes method messages.saveDraft#7ff3b806 returning error if any.
// Save a message draft¹ associated to a chat.
//
// Links:
//  1. https://core.telegram.org/api/drafts
//
// Possible errors:
//
//	400 ENTITY_BOUNDS_INVALID: A specified entity offset or length is invalid, see here » for info on how to properly compute the entity offset/length.
//	400 MSG_ID_INVALID: Invalid message ID provided.
//	400 PEER_ID_INVALID: The provided peer id is invalid.
//
// See https://core.telegram.org/method/messages.saveDraft for reference.
func (c *Client) MessagesSaveDraft(ctx context.Context, request *MessagesSaveDraftRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
