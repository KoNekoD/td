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

// SendInlineQueryResultMessageRequest represents TL type `sendInlineQueryResultMessage#467c5478`.
type SendInlineQueryResultMessageRequest struct {
	// Target chat
	ChatID int64
	// If not 0, the message thread identifier in which the message will be sent
	MessageThreadID int64
	// Information about the message or story to be replied; pass null if none
	ReplyTo InputMessageReplyToClass
	// Options to be used to send the message; pass null to use default options
	Options MessageSendOptions
	// Identifier of the inline query
	QueryID int64
	// Identifier of the inline query result
	ResultID string
	// Pass true to hide the bot, via which the message is sent. Can be used only for bots
	// getOption("animation_search_bot_username"), getOption("photo_search_bot_username"),
	// and getOption("venue_search_bot_username")
	HideViaBot bool
}

// SendInlineQueryResultMessageRequestTypeID is TL type id of SendInlineQueryResultMessageRequest.
const SendInlineQueryResultMessageRequestTypeID = 0x467c5478

// Ensuring interfaces in compile-time for SendInlineQueryResultMessageRequest.
var (
	_ bin.Encoder     = &SendInlineQueryResultMessageRequest{}
	_ bin.Decoder     = &SendInlineQueryResultMessageRequest{}
	_ bin.BareEncoder = &SendInlineQueryResultMessageRequest{}
	_ bin.BareDecoder = &SendInlineQueryResultMessageRequest{}
)

func (s *SendInlineQueryResultMessageRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageThreadID == 0) {
		return false
	}
	if !(s.ReplyTo == nil) {
		return false
	}
	if !(s.Options.Zero()) {
		return false
	}
	if !(s.QueryID == 0) {
		return false
	}
	if !(s.ResultID == "") {
		return false
	}
	if !(s.HideViaBot == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SendInlineQueryResultMessageRequest) String() string {
	if s == nil {
		return "SendInlineQueryResultMessageRequest(nil)"
	}
	type Alias SendInlineQueryResultMessageRequest
	return fmt.Sprintf("SendInlineQueryResultMessageRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SendInlineQueryResultMessageRequest) TypeID() uint32 {
	return SendInlineQueryResultMessageRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SendInlineQueryResultMessageRequest) TypeName() string {
	return "sendInlineQueryResultMessage"
}

// TypeInfo returns info about TL type.
func (s *SendInlineQueryResultMessageRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sendInlineQueryResultMessage",
		ID:   SendInlineQueryResultMessageRequestTypeID,
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
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "ReplyTo",
			SchemaName: "reply_to",
		},
		{
			Name:       "Options",
			SchemaName: "options",
		},
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "ResultID",
			SchemaName: "result_id",
		},
		{
			Name:       "HideViaBot",
			SchemaName: "hide_via_bot",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SendInlineQueryResultMessageRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendInlineQueryResultMessage#467c5478 as nil")
	}
	b.PutID(SendInlineQueryResultMessageRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SendInlineQueryResultMessageRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sendInlineQueryResultMessage#467c5478 as nil")
	}
	b.PutInt53(s.ChatID)
	b.PutInt53(s.MessageThreadID)
	if s.ReplyTo == nil {
		return fmt.Errorf("unable to encode sendInlineQueryResultMessage#467c5478: field reply_to is nil")
	}
	if err := s.ReplyTo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendInlineQueryResultMessage#467c5478: field reply_to: %w", err)
	}
	if err := s.Options.Encode(b); err != nil {
		return fmt.Errorf("unable to encode sendInlineQueryResultMessage#467c5478: field options: %w", err)
	}
	b.PutLong(s.QueryID)
	b.PutString(s.ResultID)
	b.PutBool(s.HideViaBot)
	return nil
}

// Decode implements bin.Decoder.
func (s *SendInlineQueryResultMessageRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendInlineQueryResultMessage#467c5478 to nil")
	}
	if err := b.ConsumeID(SendInlineQueryResultMessageRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SendInlineQueryResultMessageRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sendInlineQueryResultMessage#467c5478 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field message_thread_id: %w", err)
		}
		s.MessageThreadID = value
	}
	{
		value, err := DecodeInputMessageReplyTo(b)
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field reply_to: %w", err)
		}
		s.ReplyTo = value
	}
	{
		if err := s.Options.Decode(b); err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field options: %w", err)
		}
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field query_id: %w", err)
		}
		s.QueryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field result_id: %w", err)
		}
		s.ResultID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field hide_via_bot: %w", err)
		}
		s.HideViaBot = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SendInlineQueryResultMessageRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sendInlineQueryResultMessage#467c5478 as nil")
	}
	b.ObjStart()
	b.PutID("sendInlineQueryResultMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(s.ChatID)
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(s.MessageThreadID)
	b.Comma()
	b.FieldStart("reply_to")
	if s.ReplyTo == nil {
		return fmt.Errorf("unable to encode sendInlineQueryResultMessage#467c5478: field reply_to is nil")
	}
	if err := s.ReplyTo.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendInlineQueryResultMessage#467c5478: field reply_to: %w", err)
	}
	b.Comma()
	b.FieldStart("options")
	if err := s.Options.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode sendInlineQueryResultMessage#467c5478: field options: %w", err)
	}
	b.Comma()
	b.FieldStart("query_id")
	b.PutLong(s.QueryID)
	b.Comma()
	b.FieldStart("result_id")
	b.PutString(s.ResultID)
	b.Comma()
	b.FieldStart("hide_via_bot")
	b.PutBool(s.HideViaBot)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SendInlineQueryResultMessageRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sendInlineQueryResultMessage#467c5478 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sendInlineQueryResultMessage"); err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field chat_id: %w", err)
			}
			s.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field message_thread_id: %w", err)
			}
			s.MessageThreadID = value
		case "reply_to":
			value, err := DecodeTDLibJSONInputMessageReplyTo(b)
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field reply_to: %w", err)
			}
			s.ReplyTo = value
		case "options":
			if err := s.Options.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field options: %w", err)
			}
		case "query_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field query_id: %w", err)
			}
			s.QueryID = value
		case "result_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field result_id: %w", err)
			}
			s.ResultID = value
		case "hide_via_bot":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode sendInlineQueryResultMessage#467c5478: field hide_via_bot: %w", err)
			}
			s.HideViaBot = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (s *SendInlineQueryResultMessageRequest) GetChatID() (value int64) {
	if s == nil {
		return
	}
	return s.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (s *SendInlineQueryResultMessageRequest) GetMessageThreadID() (value int64) {
	if s == nil {
		return
	}
	return s.MessageThreadID
}

// GetReplyTo returns value of ReplyTo field.
func (s *SendInlineQueryResultMessageRequest) GetReplyTo() (value InputMessageReplyToClass) {
	if s == nil {
		return
	}
	return s.ReplyTo
}

// GetOptions returns value of Options field.
func (s *SendInlineQueryResultMessageRequest) GetOptions() (value MessageSendOptions) {
	if s == nil {
		return
	}
	return s.Options
}

// GetQueryID returns value of QueryID field.
func (s *SendInlineQueryResultMessageRequest) GetQueryID() (value int64) {
	if s == nil {
		return
	}
	return s.QueryID
}

// GetResultID returns value of ResultID field.
func (s *SendInlineQueryResultMessageRequest) GetResultID() (value string) {
	if s == nil {
		return
	}
	return s.ResultID
}

// GetHideViaBot returns value of HideViaBot field.
func (s *SendInlineQueryResultMessageRequest) GetHideViaBot() (value bool) {
	if s == nil {
		return
	}
	return s.HideViaBot
}

// SendInlineQueryResultMessage invokes method sendInlineQueryResultMessage#467c5478 returning error if any.
func (c *Client) SendInlineQueryResultMessage(ctx context.Context, request *SendInlineQueryResultMessageRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
