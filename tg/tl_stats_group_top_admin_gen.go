// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// StatsGroupTopAdmin represents TL type `statsGroupTopAdmin#d7584c87`.
// Information about an active admin in a supergroup
//
// See https://core.telegram.org/constructor/statsGroupTopAdmin for reference.
type StatsGroupTopAdmin struct {
	// User ID
	UserID int64
	// Number of deleted messages for statistics¹ period in consideration
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	Deleted int
	// Number of kicked users for statistics¹ period in consideration
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	Kicked int
	// Number of banned users for statistics¹ period in consideration
	//
	// Links:
	//  1) https://core.telegram.org/api/stats
	Banned int
}

// StatsGroupTopAdminTypeID is TL type id of StatsGroupTopAdmin.
const StatsGroupTopAdminTypeID = 0xd7584c87

// Ensuring interfaces in compile-time for StatsGroupTopAdmin.
var (
	_ bin.Encoder     = &StatsGroupTopAdmin{}
	_ bin.Decoder     = &StatsGroupTopAdmin{}
	_ bin.BareEncoder = &StatsGroupTopAdmin{}
	_ bin.BareDecoder = &StatsGroupTopAdmin{}
)

func (s *StatsGroupTopAdmin) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserID == 0) {
		return false
	}
	if !(s.Deleted == 0) {
		return false
	}
	if !(s.Kicked == 0) {
		return false
	}
	if !(s.Banned == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsGroupTopAdmin) String() string {
	if s == nil {
		return "StatsGroupTopAdmin(nil)"
	}
	type Alias StatsGroupTopAdmin
	return fmt.Sprintf("StatsGroupTopAdmin%+v", Alias(*s))
}

// FillFrom fills StatsGroupTopAdmin from given interface.
func (s *StatsGroupTopAdmin) FillFrom(from interface {
	GetUserID() (value int64)
	GetDeleted() (value int)
	GetKicked() (value int)
	GetBanned() (value int)
}) {
	s.UserID = from.GetUserID()
	s.Deleted = from.GetDeleted()
	s.Kicked = from.GetKicked()
	s.Banned = from.GetBanned()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StatsGroupTopAdmin) TypeID() uint32 {
	return StatsGroupTopAdminTypeID
}

// TypeName returns name of type in TL schema.
func (*StatsGroupTopAdmin) TypeName() string {
	return "statsGroupTopAdmin"
}

// TypeInfo returns info about TL type.
func (s *StatsGroupTopAdmin) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "statsGroupTopAdmin",
		ID:   StatsGroupTopAdminTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
		{
			Name:       "Deleted",
			SchemaName: "deleted",
		},
		{
			Name:       "Kicked",
			SchemaName: "kicked",
		},
		{
			Name:       "Banned",
			SchemaName: "banned",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StatsGroupTopAdmin) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopAdmin#d7584c87 as nil")
	}
	b.PutID(StatsGroupTopAdminTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StatsGroupTopAdmin) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsGroupTopAdmin#d7584c87 as nil")
	}
	b.PutLong(s.UserID)
	b.PutInt(s.Deleted)
	b.PutInt(s.Kicked)
	b.PutInt(s.Banned)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsGroupTopAdmin) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopAdmin#d7584c87 to nil")
	}
	if err := b.ConsumeID(StatsGroupTopAdminTypeID); err != nil {
		return fmt.Errorf("unable to decode statsGroupTopAdmin#d7584c87: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StatsGroupTopAdmin) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsGroupTopAdmin#d7584c87 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#d7584c87: field user_id: %w", err)
		}
		s.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#d7584c87: field deleted: %w", err)
		}
		s.Deleted = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#d7584c87: field kicked: %w", err)
		}
		s.Kicked = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode statsGroupTopAdmin#d7584c87: field banned: %w", err)
		}
		s.Banned = value
	}
	return nil
}

// GetUserID returns value of UserID field.
func (s *StatsGroupTopAdmin) GetUserID() (value int64) {
	return s.UserID
}

// GetDeleted returns value of Deleted field.
func (s *StatsGroupTopAdmin) GetDeleted() (value int) {
	return s.Deleted
}

// GetKicked returns value of Kicked field.
func (s *StatsGroupTopAdmin) GetKicked() (value int) {
	return s.Kicked
}

// GetBanned returns value of Banned field.
func (s *StatsGroupTopAdmin) GetBanned() (value int) {
	return s.Banned
}
