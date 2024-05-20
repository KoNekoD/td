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

// CreateNewSecretChatRequest represents TL type `createNewSecretChat#db012265`.
type CreateNewSecretChatRequest struct {
	// Identifier of the target user
	UserID int64
}

// CreateNewSecretChatRequestTypeID is TL type id of CreateNewSecretChatRequest.
const CreateNewSecretChatRequestTypeID = 0xdb012265

// Ensuring interfaces in compile-time for CreateNewSecretChatRequest.
var (
	_ bin.Encoder     = &CreateNewSecretChatRequest{}
	_ bin.Decoder     = &CreateNewSecretChatRequest{}
	_ bin.BareEncoder = &CreateNewSecretChatRequest{}
	_ bin.BareDecoder = &CreateNewSecretChatRequest{}
)

func (c *CreateNewSecretChatRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CreateNewSecretChatRequest) String() string {
	if c == nil {
		return "CreateNewSecretChatRequest(nil)"
	}
	type Alias CreateNewSecretChatRequest
	return fmt.Sprintf("CreateNewSecretChatRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CreateNewSecretChatRequest) TypeID() uint32 {
	return CreateNewSecretChatRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CreateNewSecretChatRequest) TypeName() string {
	return "createNewSecretChat"
}

// TypeInfo returns info about TL type.
func (c *CreateNewSecretChatRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "createNewSecretChat",
		ID:   CreateNewSecretChatRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CreateNewSecretChatRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createNewSecretChat#db012265 as nil")
	}
	b.PutID(CreateNewSecretChatRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CreateNewSecretChatRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode createNewSecretChat#db012265 as nil")
	}
	b.PutInt53(c.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (c *CreateNewSecretChatRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createNewSecretChat#db012265 to nil")
	}
	if err := b.ConsumeID(CreateNewSecretChatRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode createNewSecretChat#db012265: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CreateNewSecretChatRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode createNewSecretChat#db012265 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode createNewSecretChat#db012265: field user_id: %w", err)
		}
		c.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CreateNewSecretChatRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode createNewSecretChat#db012265 as nil")
	}
	b.ObjStart()
	b.PutID("createNewSecretChat")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(c.UserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CreateNewSecretChatRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode createNewSecretChat#db012265 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("createNewSecretChat"); err != nil {
				return fmt.Errorf("unable to decode createNewSecretChat#db012265: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode createNewSecretChat#db012265: field user_id: %w", err)
			}
			c.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (c *CreateNewSecretChatRequest) GetUserID() (value int64) {
	if c == nil {
		return
	}
	return c.UserID
}

// CreateNewSecretChat invokes method createNewSecretChat#db012265 returning error if any.
func (c *Client) CreateNewSecretChat(ctx context.Context, userid int64) (*Chat, error) {
	var result Chat

	request := &CreateNewSecretChatRequest{
		UserID: userid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
