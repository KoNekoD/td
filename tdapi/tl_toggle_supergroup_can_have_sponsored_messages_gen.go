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

// ToggleSupergroupCanHaveSponsoredMessagesRequest represents TL type `toggleSupergroupCanHaveSponsoredMessages#be8abb72`.
type ToggleSupergroupCanHaveSponsoredMessagesRequest struct {
	// The identifier of the channel
	SupergroupID int64
	// The new value of can_have_sponsored_messages
	CanHaveSponsoredMessages bool
}

// ToggleSupergroupCanHaveSponsoredMessagesRequestTypeID is TL type id of ToggleSupergroupCanHaveSponsoredMessagesRequest.
const ToggleSupergroupCanHaveSponsoredMessagesRequestTypeID = 0xbe8abb72

// Ensuring interfaces in compile-time for ToggleSupergroupCanHaveSponsoredMessagesRequest.
var (
	_ bin.Encoder     = &ToggleSupergroupCanHaveSponsoredMessagesRequest{}
	_ bin.Decoder     = &ToggleSupergroupCanHaveSponsoredMessagesRequest{}
	_ bin.BareEncoder = &ToggleSupergroupCanHaveSponsoredMessagesRequest{}
	_ bin.BareDecoder = &ToggleSupergroupCanHaveSponsoredMessagesRequest{}
)

func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SupergroupID == 0) {
		return false
	}
	if !(t.CanHaveSponsoredMessages == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) String() string {
	if t == nil {
		return "ToggleSupergroupCanHaveSponsoredMessagesRequest(nil)"
	}
	type Alias ToggleSupergroupCanHaveSponsoredMessagesRequest
	return fmt.Sprintf("ToggleSupergroupCanHaveSponsoredMessagesRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleSupergroupCanHaveSponsoredMessagesRequest) TypeID() uint32 {
	return ToggleSupergroupCanHaveSponsoredMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleSupergroupCanHaveSponsoredMessagesRequest) TypeName() string {
	return "toggleSupergroupCanHaveSponsoredMessages"
}

// TypeInfo returns info about TL type.
func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleSupergroupCanHaveSponsoredMessages",
		ID:   ToggleSupergroupCanHaveSponsoredMessagesRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SupergroupID",
			SchemaName: "supergroup_id",
		},
		{
			Name:       "CanHaveSponsoredMessages",
			SchemaName: "can_have_sponsored_messages",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupCanHaveSponsoredMessages#be8abb72 as nil")
	}
	b.PutID(ToggleSupergroupCanHaveSponsoredMessagesRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupCanHaveSponsoredMessages#be8abb72 as nil")
	}
	b.PutInt53(t.SupergroupID)
	b.PutBool(t.CanHaveSponsoredMessages)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupCanHaveSponsoredMessages#be8abb72 to nil")
	}
	if err := b.ConsumeID(ToggleSupergroupCanHaveSponsoredMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleSupergroupCanHaveSponsoredMessages#be8abb72: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupCanHaveSponsoredMessages#be8abb72 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSupergroupCanHaveSponsoredMessages#be8abb72: field supergroup_id: %w", err)
		}
		t.SupergroupID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSupergroupCanHaveSponsoredMessages#be8abb72: field can_have_sponsored_messages: %w", err)
		}
		t.CanHaveSponsoredMessages = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupCanHaveSponsoredMessages#be8abb72 as nil")
	}
	b.ObjStart()
	b.PutID("toggleSupergroupCanHaveSponsoredMessages")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(t.SupergroupID)
	b.Comma()
	b.FieldStart("can_have_sponsored_messages")
	b.PutBool(t.CanHaveSponsoredMessages)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupCanHaveSponsoredMessages#be8abb72 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleSupergroupCanHaveSponsoredMessages"); err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupCanHaveSponsoredMessages#be8abb72: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupCanHaveSponsoredMessages#be8abb72: field supergroup_id: %w", err)
			}
			t.SupergroupID = value
		case "can_have_sponsored_messages":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupCanHaveSponsoredMessages#be8abb72: field can_have_sponsored_messages: %w", err)
			}
			t.CanHaveSponsoredMessages = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) GetSupergroupID() (value int64) {
	if t == nil {
		return
	}
	return t.SupergroupID
}

// GetCanHaveSponsoredMessages returns value of CanHaveSponsoredMessages field.
func (t *ToggleSupergroupCanHaveSponsoredMessagesRequest) GetCanHaveSponsoredMessages() (value bool) {
	if t == nil {
		return
	}
	return t.CanHaveSponsoredMessages
}

// ToggleSupergroupCanHaveSponsoredMessages invokes method toggleSupergroupCanHaveSponsoredMessages#be8abb72 returning error if any.
func (c *Client) ToggleSupergroupCanHaveSponsoredMessages(ctx context.Context, request *ToggleSupergroupCanHaveSponsoredMessagesRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
