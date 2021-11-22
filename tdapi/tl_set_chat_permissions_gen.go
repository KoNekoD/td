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

// SetChatPermissionsRequest represents TL type `setChatPermissions#7f7706fe`.
type SetChatPermissionsRequest struct {
	// Chat identifier
	ChatID int64
	// New non-administrator members permissions in the chat
	Permissions ChatPermissions
}

// SetChatPermissionsRequestTypeID is TL type id of SetChatPermissionsRequest.
const SetChatPermissionsRequestTypeID = 0x7f7706fe

// Ensuring interfaces in compile-time for SetChatPermissionsRequest.
var (
	_ bin.Encoder     = &SetChatPermissionsRequest{}
	_ bin.Decoder     = &SetChatPermissionsRequest{}
	_ bin.BareEncoder = &SetChatPermissionsRequest{}
	_ bin.BareDecoder = &SetChatPermissionsRequest{}
)

func (s *SetChatPermissionsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Permissions.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetChatPermissionsRequest) String() string {
	if s == nil {
		return "SetChatPermissionsRequest(nil)"
	}
	type Alias SetChatPermissionsRequest
	return fmt.Sprintf("SetChatPermissionsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetChatPermissionsRequest) TypeID() uint32 {
	return SetChatPermissionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetChatPermissionsRequest) TypeName() string {
	return "setChatPermissions"
}

// TypeInfo returns info about TL type.
func (s *SetChatPermissionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setChatPermissions",
		ID:   SetChatPermissionsRequestTypeID,
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
			Name:       "Permissions",
			SchemaName: "permissions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetChatPermissionsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatPermissions#7f7706fe as nil")
	}
	b.PutID(SetChatPermissionsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetChatPermissionsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatPermissions#7f7706fe as nil")
	}
	b.PutLong(s.ChatID)
	if err := s.Permissions.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setChatPermissions#7f7706fe: field permissions: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetChatPermissionsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatPermissions#7f7706fe to nil")
	}
	if err := b.ConsumeID(SetChatPermissionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setChatPermissions#7f7706fe: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetChatPermissionsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setChatPermissions#7f7706fe to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode setChatPermissions#7f7706fe: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		if err := s.Permissions.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setChatPermissions#7f7706fe: field permissions: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON encodes s in TDLib API JSON format.
func (s *SetChatPermissionsRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setChatPermissions#7f7706fe as nil")
	}
	b.ObjStart()
	b.PutID("setChatPermissions")
	b.FieldStart("chat_id")
	b.PutLong(s.ChatID)
	b.FieldStart("permissions")
	if err := s.Permissions.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setChatPermissions#7f7706fe: field permissions: %w", err)
	}
	b.ObjEnd()
	return nil
}

// GetChatID returns value of ChatID field.
func (s *SetChatPermissionsRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetPermissions returns value of Permissions field.
func (s *SetChatPermissionsRequest) GetPermissions() (value ChatPermissions) {
	return s.Permissions
}

// SetChatPermissions invokes method setChatPermissions#7f7706fe returning error if any.
func (c *Client) SetChatPermissions(ctx context.Context, request *SetChatPermissionsRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
