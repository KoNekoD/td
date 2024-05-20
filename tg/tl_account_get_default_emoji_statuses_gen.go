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

// AccountGetDefaultEmojiStatusesRequest represents TL type `account.getDefaultEmojiStatuses#d6753386`.
// Get a list of default suggested emoji statuses¹
//
// Links:
//  1. https://core.telegram.org/api/emoji-status
//
// See https://core.telegram.org/method/account.getDefaultEmojiStatuses for reference.
type AccountGetDefaultEmojiStatusesRequest struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int64
}

// AccountGetDefaultEmojiStatusesRequestTypeID is TL type id of AccountGetDefaultEmojiStatusesRequest.
const AccountGetDefaultEmojiStatusesRequestTypeID = 0xd6753386

// Ensuring interfaces in compile-time for AccountGetDefaultEmojiStatusesRequest.
var (
	_ bin.Encoder     = &AccountGetDefaultEmojiStatusesRequest{}
	_ bin.Decoder     = &AccountGetDefaultEmojiStatusesRequest{}
	_ bin.BareEncoder = &AccountGetDefaultEmojiStatusesRequest{}
	_ bin.BareDecoder = &AccountGetDefaultEmojiStatusesRequest{}
)

func (g *AccountGetDefaultEmojiStatusesRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Hash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *AccountGetDefaultEmojiStatusesRequest) String() string {
	if g == nil {
		return "AccountGetDefaultEmojiStatusesRequest(nil)"
	}
	type Alias AccountGetDefaultEmojiStatusesRequest
	return fmt.Sprintf("AccountGetDefaultEmojiStatusesRequest%+v", Alias(*g))
}

// FillFrom fills AccountGetDefaultEmojiStatusesRequest from given interface.
func (g *AccountGetDefaultEmojiStatusesRequest) FillFrom(from interface {
	GetHash() (value int64)
}) {
	g.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountGetDefaultEmojiStatusesRequest) TypeID() uint32 {
	return AccountGetDefaultEmojiStatusesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountGetDefaultEmojiStatusesRequest) TypeName() string {
	return "account.getDefaultEmojiStatuses"
}

// TypeInfo returns info about TL type.
func (g *AccountGetDefaultEmojiStatusesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.getDefaultEmojiStatuses",
		ID:   AccountGetDefaultEmojiStatusesRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *AccountGetDefaultEmojiStatusesRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getDefaultEmojiStatuses#d6753386 as nil")
	}
	b.PutID(AccountGetDefaultEmojiStatusesRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *AccountGetDefaultEmojiStatusesRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode account.getDefaultEmojiStatuses#d6753386 as nil")
	}
	b.PutLong(g.Hash)
	return nil
}

// Decode implements bin.Decoder.
func (g *AccountGetDefaultEmojiStatusesRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getDefaultEmojiStatuses#d6753386 to nil")
	}
	if err := b.ConsumeID(AccountGetDefaultEmojiStatusesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.getDefaultEmojiStatuses#d6753386: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *AccountGetDefaultEmojiStatusesRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode account.getDefaultEmojiStatuses#d6753386 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode account.getDefaultEmojiStatuses#d6753386: field hash: %w", err)
		}
		g.Hash = value
	}
	return nil
}

// GetHash returns value of Hash field.
func (g *AccountGetDefaultEmojiStatusesRequest) GetHash() (value int64) {
	if g == nil {
		return
	}
	return g.Hash
}

// AccountGetDefaultEmojiStatuses invokes method account.getDefaultEmojiStatuses#d6753386 returning error if any.
// Get a list of default suggested emoji statuses¹
//
// Links:
//  1. https://core.telegram.org/api/emoji-status
//
// See https://core.telegram.org/method/account.getDefaultEmojiStatuses for reference.
func (c *Client) AccountGetDefaultEmojiStatuses(ctx context.Context, hash int64) (AccountEmojiStatusesClass, error) {
	var result AccountEmojiStatusesBox

	request := &AccountGetDefaultEmojiStatusesRequest{
		Hash: hash,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.EmojiStatuses, nil
}
