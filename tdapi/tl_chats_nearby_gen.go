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

// ChatsNearby represents TL type `chatsNearby#cc744cff`.
type ChatsNearby struct {
	// List of users nearby
	UsersNearby []ChatNearby
	// List of location-based supergroups nearby
	SupergroupsNearby []ChatNearby
}

// ChatsNearbyTypeID is TL type id of ChatsNearby.
const ChatsNearbyTypeID = 0xcc744cff

// Ensuring interfaces in compile-time for ChatsNearby.
var (
	_ bin.Encoder     = &ChatsNearby{}
	_ bin.Decoder     = &ChatsNearby{}
	_ bin.BareEncoder = &ChatsNearby{}
	_ bin.BareDecoder = &ChatsNearby{}
)

func (c *ChatsNearby) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UsersNearby == nil) {
		return false
	}
	if !(c.SupergroupsNearby == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatsNearby) String() string {
	if c == nil {
		return "ChatsNearby(nil)"
	}
	type Alias ChatsNearby
	return fmt.Sprintf("ChatsNearby%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatsNearby) TypeID() uint32 {
	return ChatsNearbyTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatsNearby) TypeName() string {
	return "chatsNearby"
}

// TypeInfo returns info about TL type.
func (c *ChatsNearby) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatsNearby",
		ID:   ChatsNearbyTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UsersNearby",
			SchemaName: "users_nearby",
		},
		{
			Name:       "SupergroupsNearby",
			SchemaName: "supergroups_nearby",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatsNearby) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatsNearby#cc744cff as nil")
	}
	b.PutID(ChatsNearbyTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatsNearby) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatsNearby#cc744cff as nil")
	}
	b.PutInt(len(c.UsersNearby))
	for idx, v := range c.UsersNearby {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatsNearby#cc744cff: field users_nearby element with index %d: %w", idx, err)
		}
	}
	b.PutInt(len(c.SupergroupsNearby))
	for idx, v := range c.SupergroupsNearby {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare chatsNearby#cc744cff: field supergroups_nearby element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatsNearby) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatsNearby#cc744cff to nil")
	}
	if err := b.ConsumeID(ChatsNearbyTypeID); err != nil {
		return fmt.Errorf("unable to decode chatsNearby#cc744cff: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatsNearby) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatsNearby#cc744cff to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatsNearby#cc744cff: field users_nearby: %w", err)
		}

		if headerLen > 0 {
			c.UsersNearby = make([]ChatNearby, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatNearby
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatsNearby#cc744cff: field users_nearby: %w", err)
			}
			c.UsersNearby = append(c.UsersNearby, value)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatsNearby#cc744cff: field supergroups_nearby: %w", err)
		}

		if headerLen > 0 {
			c.SupergroupsNearby = make([]ChatNearby, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatNearby
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare chatsNearby#cc744cff: field supergroups_nearby: %w", err)
			}
			c.SupergroupsNearby = append(c.SupergroupsNearby, value)
		}
	}
	return nil
}

// EncodeTDLibJSON encodes c in TDLib API JSON format.
func (c *ChatsNearby) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatsNearby#cc744cff as nil")
	}
	b.ObjStart()
	b.PutID("chatsNearby")
	b.FieldStart("users_nearby")
	b.ArrStart()
	for idx, v := range c.UsersNearby {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatsNearby#cc744cff: field users_nearby element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.FieldStart("supergroups_nearby")
	b.ArrStart()
	for idx, v := range c.SupergroupsNearby {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode chatsNearby#cc744cff: field supergroups_nearby element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// GetUsersNearby returns value of UsersNearby field.
func (c *ChatsNearby) GetUsersNearby() (value []ChatNearby) {
	return c.UsersNearby
}

// GetSupergroupsNearby returns value of SupergroupsNearby field.
func (c *ChatsNearby) GetSupergroupsNearby() (value []ChatNearby) {
	return c.SupergroupsNearby
}
