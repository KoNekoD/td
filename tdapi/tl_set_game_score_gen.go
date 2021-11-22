// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// SetGameScoreRequest represents TL type `setGameScore#9699c683`.
type SetGameScoreRequest struct {
	// The chat to which the message with the game belongs
	ChatID int64
	// Identifier of the message
	MessageID int64
	// True, if the message should be edited
	EditMessage bool
	// User identifier
	UserID int32
	// The new score
	Score int32
	// Pass true to update the score even if it decreases. If the score is 0, the user will
	// be deleted from the high score table
	Force bool
}

// SetGameScoreRequestTypeID is TL type id of SetGameScoreRequest.
const SetGameScoreRequestTypeID = 0x9699c683

// Ensuring interfaces in compile-time for SetGameScoreRequest.
var (
	_ bin.Encoder     = &SetGameScoreRequest{}
	_ bin.Decoder     = &SetGameScoreRequest{}
	_ bin.BareEncoder = &SetGameScoreRequest{}
	_ bin.BareDecoder = &SetGameScoreRequest{}
)

func (s *SetGameScoreRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.MessageID == 0) {
		return false
	}
	if !(s.EditMessage == false) {
		return false
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Score == 0) {
		return false
	}
	if !(s.Force == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetGameScoreRequest) String() string {
	if s == nil {
		return "SetGameScoreRequest(nil)"
	}
	type Alias SetGameScoreRequest
	return fmt.Sprintf("SetGameScoreRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetGameScoreRequest) TypeID() uint32 {
	return SetGameScoreRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetGameScoreRequest) TypeName() string {
	return "setGameScore"
}

// TypeInfo returns info about TL type.
func (s *SetGameScoreRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setGameScore",
		ID:   SetGameScoreRequestTypeID,
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
			Name:       "MessageID",
			SchemaName: "message_id",
		},
		{
			Name:       "EditMessage",
			SchemaName: "edit_message",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Score",
			SchemaName: "score",
		},
		{
			Name:       "Force",
			SchemaName: "force",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetGameScoreRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setGameScore#9699c683 as nil")
	}
	b.PutID(SetGameScoreRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetGameScoreRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setGameScore#9699c683 as nil")
	}
	b.PutLong(s.ChatID)
	b.PutLong(s.MessageID)
	b.PutBool(s.EditMessage)
	b.PutInt32(s.UserID)
	b.PutInt32(s.Score)
	b.PutBool(s.Force)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetGameScoreRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setGameScore#9699c683 to nil")
	}
	if err := b.ConsumeID(SetGameScoreRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setGameScore#9699c683: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetGameScoreRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setGameScore#9699c683 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#9699c683: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#9699c683: field message_id: %w", err)
		}
		s.MessageID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#9699c683: field edit_message: %w", err)
		}
		s.EditMessage = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#9699c683: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#9699c683: field score: %w", err)
		}
		s.Score = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode setGameScore#9699c683: field force: %w", err)
		}
		s.Force = value
	}
	return nil
}

// EncodeTDLibJSON encodes s in TDLib API JSON format.
func (s *SetGameScoreRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setGameScore#9699c683 as nil")
	}
	b.ObjStart()
	b.PutID("setGameScore")
	b.FieldStart("chat_id")
	b.PutLong(s.ChatID)
	b.FieldStart("message_id")
	b.PutLong(s.MessageID)
	b.FieldStart("edit_message")
	b.PutBool(s.EditMessage)
	b.FieldStart("user_id")
	b.PutInt32(s.UserID)
	b.FieldStart("score")
	b.PutInt32(s.Score)
	b.FieldStart("force")
	b.PutBool(s.Force)
	b.ObjEnd()
	return nil
}

// GetChatID returns value of ChatID field.
func (s *SetGameScoreRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetMessageID returns value of MessageID field.
func (s *SetGameScoreRequest) GetMessageID() (value int64) {
	return s.MessageID
}

// GetEditMessage returns value of EditMessage field.
func (s *SetGameScoreRequest) GetEditMessage() (value bool) {
	return s.EditMessage
}

// GetUserID returns value of UserID field.
func (s *SetGameScoreRequest) GetUserID() (value int32) {
	return s.UserID
}

// GetScore returns value of Score field.
func (s *SetGameScoreRequest) GetScore() (value int32) {
	return s.Score
}

// GetForce returns value of Force field.
func (s *SetGameScoreRequest) GetForce() (value bool) {
	return s.Force
}

// SetGameScore invokes method setGameScore#9699c683 returning error if any.
func (c *Client) SetGameScore(ctx context.Context, request *SetGameScoreRequest) (*Message, error) {
	var result Message

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
