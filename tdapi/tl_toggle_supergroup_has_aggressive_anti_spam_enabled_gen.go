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

// ToggleSupergroupHasAggressiveAntiSpamEnabledRequest represents TL type `toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f`.
type ToggleSupergroupHasAggressiveAntiSpamEnabledRequest struct {
	// The identifier of the supergroup, which isn't a broadcast group
	SupergroupID int64
	// The new value of has_aggressive_anti_spam_enabled
	HasAggressiveAntiSpamEnabled bool
}

// ToggleSupergroupHasAggressiveAntiSpamEnabledRequestTypeID is TL type id of ToggleSupergroupHasAggressiveAntiSpamEnabledRequest.
const ToggleSupergroupHasAggressiveAntiSpamEnabledRequestTypeID = 0x683ef70f

// Ensuring interfaces in compile-time for ToggleSupergroupHasAggressiveAntiSpamEnabledRequest.
var (
	_ bin.Encoder     = &ToggleSupergroupHasAggressiveAntiSpamEnabledRequest{}
	_ bin.Decoder     = &ToggleSupergroupHasAggressiveAntiSpamEnabledRequest{}
	_ bin.BareEncoder = &ToggleSupergroupHasAggressiveAntiSpamEnabledRequest{}
	_ bin.BareDecoder = &ToggleSupergroupHasAggressiveAntiSpamEnabledRequest{}
)

func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SupergroupID == 0) {
		return false
	}
	if !(t.HasAggressiveAntiSpamEnabled == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) String() string {
	if t == nil {
		return "ToggleSupergroupHasAggressiveAntiSpamEnabledRequest(nil)"
	}
	type Alias ToggleSupergroupHasAggressiveAntiSpamEnabledRequest
	return fmt.Sprintf("ToggleSupergroupHasAggressiveAntiSpamEnabledRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) TypeID() uint32 {
	return ToggleSupergroupHasAggressiveAntiSpamEnabledRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) TypeName() string {
	return "toggleSupergroupHasAggressiveAntiSpamEnabled"
}

// TypeInfo returns info about TL type.
func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleSupergroupHasAggressiveAntiSpamEnabled",
		ID:   ToggleSupergroupHasAggressiveAntiSpamEnabledRequestTypeID,
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
			Name:       "HasAggressiveAntiSpamEnabled",
			SchemaName: "has_aggressive_anti_spam_enabled",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f as nil")
	}
	b.PutID(ToggleSupergroupHasAggressiveAntiSpamEnabledRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f as nil")
	}
	b.PutInt53(t.SupergroupID)
	b.PutBool(t.HasAggressiveAntiSpamEnabled)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f to nil")
	}
	if err := b.ConsumeID(ToggleSupergroupHasAggressiveAntiSpamEnabledRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f: field supergroup_id: %w", err)
		}
		t.SupergroupID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f: field has_aggressive_anti_spam_enabled: %w", err)
		}
		t.HasAggressiveAntiSpamEnabled = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f as nil")
	}
	b.ObjStart()
	b.PutID("toggleSupergroupHasAggressiveAntiSpamEnabled")
	b.Comma()
	b.FieldStart("supergroup_id")
	b.PutInt53(t.SupergroupID)
	b.Comma()
	b.FieldStart("has_aggressive_anti_spam_enabled")
	b.PutBool(t.HasAggressiveAntiSpamEnabled)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("toggleSupergroupHasAggressiveAntiSpamEnabled"); err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f: %w", err)
			}
		case "supergroup_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f: field supergroup_id: %w", err)
			}
			t.SupergroupID = value
		case "has_aggressive_anti_spam_enabled":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f: field has_aggressive_anti_spam_enabled: %w", err)
			}
			t.HasAggressiveAntiSpamEnabled = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSupergroupID returns value of SupergroupID field.
func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) GetSupergroupID() (value int64) {
	if t == nil {
		return
	}
	return t.SupergroupID
}

// GetHasAggressiveAntiSpamEnabled returns value of HasAggressiveAntiSpamEnabled field.
func (t *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) GetHasAggressiveAntiSpamEnabled() (value bool) {
	if t == nil {
		return
	}
	return t.HasAggressiveAntiSpamEnabled
}

// ToggleSupergroupHasAggressiveAntiSpamEnabled invokes method toggleSupergroupHasAggressiveAntiSpamEnabled#683ef70f returning error if any.
func (c *Client) ToggleSupergroupHasAggressiveAntiSpamEnabled(ctx context.Context, request *ToggleSupergroupHasAggressiveAntiSpamEnabledRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
