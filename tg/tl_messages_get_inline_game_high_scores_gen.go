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

// MessagesGetInlineGameHighScoresRequest represents TL type `messages.getInlineGameHighScores#f635e1b`.
// Get highscores of a game sent using an inline bot
//
// See https://core.telegram.org/method/messages.getInlineGameHighScores for reference.
type MessagesGetInlineGameHighScoresRequest struct {
	// ID of inline message
	ID InputBotInlineMessageIDClass
	// Get high scores of a certain user
	UserID InputUserClass
}

// MessagesGetInlineGameHighScoresRequestTypeID is TL type id of MessagesGetInlineGameHighScoresRequest.
const MessagesGetInlineGameHighScoresRequestTypeID = 0xf635e1b

// Ensuring interfaces in compile-time for MessagesGetInlineGameHighScoresRequest.
var (
	_ bin.Encoder     = &MessagesGetInlineGameHighScoresRequest{}
	_ bin.Decoder     = &MessagesGetInlineGameHighScoresRequest{}
	_ bin.BareEncoder = &MessagesGetInlineGameHighScoresRequest{}
	_ bin.BareDecoder = &MessagesGetInlineGameHighScoresRequest{}
)

func (g *MessagesGetInlineGameHighScoresRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == nil) {
		return false
	}
	if !(g.UserID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *MessagesGetInlineGameHighScoresRequest) String() string {
	if g == nil {
		return "MessagesGetInlineGameHighScoresRequest(nil)"
	}
	type Alias MessagesGetInlineGameHighScoresRequest
	return fmt.Sprintf("MessagesGetInlineGameHighScoresRequest%+v", Alias(*g))
}

// FillFrom fills MessagesGetInlineGameHighScoresRequest from given interface.
func (g *MessagesGetInlineGameHighScoresRequest) FillFrom(from interface {
	GetID() (value InputBotInlineMessageIDClass)
	GetUserID() (value InputUserClass)
}) {
	g.ID = from.GetID()
	g.UserID = from.GetUserID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesGetInlineGameHighScoresRequest) TypeID() uint32 {
	return MessagesGetInlineGameHighScoresRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesGetInlineGameHighScoresRequest) TypeName() string {
	return "messages.getInlineGameHighScores"
}

// TypeInfo returns info about TL type.
func (g *MessagesGetInlineGameHighScoresRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.getInlineGameHighScores",
		ID:   MessagesGetInlineGameHighScoresRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *MessagesGetInlineGameHighScoresRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getInlineGameHighScores#f635e1b as nil")
	}
	b.PutID(MessagesGetInlineGameHighScoresRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *MessagesGetInlineGameHighScoresRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode messages.getInlineGameHighScores#f635e1b as nil")
	}
	if g.ID == nil {
		return fmt.Errorf("unable to encode messages.getInlineGameHighScores#f635e1b: field id is nil")
	}
	if err := g.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getInlineGameHighScores#f635e1b: field id: %w", err)
	}
	if g.UserID == nil {
		return fmt.Errorf("unable to encode messages.getInlineGameHighScores#f635e1b: field user_id is nil")
	}
	if err := g.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.getInlineGameHighScores#f635e1b: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *MessagesGetInlineGameHighScoresRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getInlineGameHighScores#f635e1b to nil")
	}
	if err := b.ConsumeID(MessagesGetInlineGameHighScoresRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.getInlineGameHighScores#f635e1b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *MessagesGetInlineGameHighScoresRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode messages.getInlineGameHighScores#f635e1b to nil")
	}
	{
		value, err := DecodeInputBotInlineMessageID(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getInlineGameHighScores#f635e1b: field id: %w", err)
		}
		g.ID = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.getInlineGameHighScores#f635e1b: field user_id: %w", err)
		}
		g.UserID = value
	}
	return nil
}

// GetID returns value of ID field.
func (g *MessagesGetInlineGameHighScoresRequest) GetID() (value InputBotInlineMessageIDClass) {
	if g == nil {
		return
	}
	return g.ID
}

// GetUserID returns value of UserID field.
func (g *MessagesGetInlineGameHighScoresRequest) GetUserID() (value InputUserClass) {
	if g == nil {
		return
	}
	return g.UserID
}

// MessagesGetInlineGameHighScores invokes method messages.getInlineGameHighScores#f635e1b returning error if any.
// Get highscores of a game sent using an inline bot
//
// Possible errors:
//
//	400 MESSAGE_ID_INVALID: The provided message id is invalid.
//	400 USER_BOT_REQUIRED: This method can only be called by a bot.
//
// See https://core.telegram.org/method/messages.getInlineGameHighScores for reference.
// Can be used by bots.
func (c *Client) MessagesGetInlineGameHighScores(ctx context.Context, request *MessagesGetInlineGameHighScoresRequest) (*MessagesHighScores, error) {
	var result MessagesHighScores

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
