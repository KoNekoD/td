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

// ClearAllDraftMessagesRequest represents TL type `clearAllDraftMessages#fd3c74db`.
type ClearAllDraftMessagesRequest struct {
	// Pass true to keep local message drafts in secret chats
	ExcludeSecretChats bool
}

// ClearAllDraftMessagesRequestTypeID is TL type id of ClearAllDraftMessagesRequest.
const ClearAllDraftMessagesRequestTypeID = 0xfd3c74db

// Ensuring interfaces in compile-time for ClearAllDraftMessagesRequest.
var (
	_ bin.Encoder     = &ClearAllDraftMessagesRequest{}
	_ bin.Decoder     = &ClearAllDraftMessagesRequest{}
	_ bin.BareEncoder = &ClearAllDraftMessagesRequest{}
	_ bin.BareDecoder = &ClearAllDraftMessagesRequest{}
)

func (c *ClearAllDraftMessagesRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.ExcludeSecretChats == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ClearAllDraftMessagesRequest) String() string {
	if c == nil {
		return "ClearAllDraftMessagesRequest(nil)"
	}
	type Alias ClearAllDraftMessagesRequest
	return fmt.Sprintf("ClearAllDraftMessagesRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ClearAllDraftMessagesRequest) TypeID() uint32 {
	return ClearAllDraftMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ClearAllDraftMessagesRequest) TypeName() string {
	return "clearAllDraftMessages"
}

// TypeInfo returns info about TL type.
func (c *ClearAllDraftMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "clearAllDraftMessages",
		ID:   ClearAllDraftMessagesRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ExcludeSecretChats",
			SchemaName: "exclude_secret_chats",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ClearAllDraftMessagesRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode clearAllDraftMessages#fd3c74db as nil")
	}
	b.PutID(ClearAllDraftMessagesRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ClearAllDraftMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode clearAllDraftMessages#fd3c74db as nil")
	}
	b.PutBool(c.ExcludeSecretChats)
	return nil
}

// Decode implements bin.Decoder.
func (c *ClearAllDraftMessagesRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode clearAllDraftMessages#fd3c74db to nil")
	}
	if err := b.ConsumeID(ClearAllDraftMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode clearAllDraftMessages#fd3c74db: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ClearAllDraftMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode clearAllDraftMessages#fd3c74db to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode clearAllDraftMessages#fd3c74db: field exclude_secret_chats: %w", err)
		}
		c.ExcludeSecretChats = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ClearAllDraftMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode clearAllDraftMessages#fd3c74db as nil")
	}
	b.ObjStart()
	b.PutID("clearAllDraftMessages")
	b.Comma()
	b.FieldStart("exclude_secret_chats")
	b.PutBool(c.ExcludeSecretChats)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ClearAllDraftMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode clearAllDraftMessages#fd3c74db to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("clearAllDraftMessages"); err != nil {
				return fmt.Errorf("unable to decode clearAllDraftMessages#fd3c74db: %w", err)
			}
		case "exclude_secret_chats":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode clearAllDraftMessages#fd3c74db: field exclude_secret_chats: %w", err)
			}
			c.ExcludeSecretChats = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetExcludeSecretChats returns value of ExcludeSecretChats field.
func (c *ClearAllDraftMessagesRequest) GetExcludeSecretChats() (value bool) {
	if c == nil {
		return
	}
	return c.ExcludeSecretChats
}

// ClearAllDraftMessages invokes method clearAllDraftMessages#fd3c74db returning error if any.
func (c *Client) ClearAllDraftMessages(ctx context.Context, excludesecretchats bool) error {
	var ok Ok

	request := &ClearAllDraftMessagesRequest{
		ExcludeSecretChats: excludesecretchats,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
