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

// ToggleGroupCallParticipantIsMutedRequest represents TL type `toggleGroupCallParticipantIsMuted#b2081407`.
type ToggleGroupCallParticipantIsMutedRequest struct {
	// Group call identifier
	GroupCallID int32
	// Participant identifier
	ParticipantID MessageSenderClass
	// Pass true if the user must be muted and false otherwise
	IsMuted bool
}

// ToggleGroupCallParticipantIsMutedRequestTypeID is TL type id of ToggleGroupCallParticipantIsMutedRequest.
const ToggleGroupCallParticipantIsMutedRequestTypeID = 0xb2081407

// Ensuring interfaces in compile-time for ToggleGroupCallParticipantIsMutedRequest.
var (
	_ bin.Encoder     = &ToggleGroupCallParticipantIsMutedRequest{}
	_ bin.Decoder     = &ToggleGroupCallParticipantIsMutedRequest{}
	_ bin.BareEncoder = &ToggleGroupCallParticipantIsMutedRequest{}
	_ bin.BareDecoder = &ToggleGroupCallParticipantIsMutedRequest{}
)

func (t *ToggleGroupCallParticipantIsMutedRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.GroupCallID == 0) {
		return false
	}
	if !(t.ParticipantID == nil) {
		return false
	}
	if !(t.IsMuted == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleGroupCallParticipantIsMutedRequest) String() string {
	if t == nil {
		return "ToggleGroupCallParticipantIsMutedRequest(nil)"
	}
	type Alias ToggleGroupCallParticipantIsMutedRequest
	return fmt.Sprintf("ToggleGroupCallParticipantIsMutedRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleGroupCallParticipantIsMutedRequest) TypeID() uint32 {
	return ToggleGroupCallParticipantIsMutedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleGroupCallParticipantIsMutedRequest) TypeName() string {
	return "toggleGroupCallParticipantIsMuted"
}

// TypeInfo returns info about TL type.
func (t *ToggleGroupCallParticipantIsMutedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleGroupCallParticipantIsMuted",
		ID:   ToggleGroupCallParticipantIsMutedRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
		{
			Name:       "ParticipantID",
			SchemaName: "participant_id",
		},
		{
			Name:       "IsMuted",
			SchemaName: "is_muted",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleGroupCallParticipantIsMutedRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleGroupCallParticipantIsMuted#b2081407 as nil")
	}
	b.PutID(ToggleGroupCallParticipantIsMutedRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleGroupCallParticipantIsMutedRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleGroupCallParticipantIsMuted#b2081407 as nil")
	}
	b.PutInt32(t.GroupCallID)
	if t.ParticipantID == nil {
		return fmt.Errorf("unable to encode toggleGroupCallParticipantIsMuted#b2081407: field participant_id is nil")
	}
	if err := t.ParticipantID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode toggleGroupCallParticipantIsMuted#b2081407: field participant_id: %w", err)
	}
	b.PutBool(t.IsMuted)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleGroupCallParticipantIsMutedRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleGroupCallParticipantIsMuted#b2081407 to nil")
	}
	if err := b.ConsumeID(ToggleGroupCallParticipantIsMutedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleGroupCallParticipantIsMuted#b2081407: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleGroupCallParticipantIsMutedRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleGroupCallParticipantIsMuted#b2081407 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode toggleGroupCallParticipantIsMuted#b2081407: field group_call_id: %w", err)
		}
		t.GroupCallID = value
	}
	{
		value, err := DecodeMessageSender(b)
		if err != nil {
			return fmt.Errorf("unable to decode toggleGroupCallParticipantIsMuted#b2081407: field participant_id: %w", err)
		}
		t.ParticipantID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleGroupCallParticipantIsMuted#b2081407: field is_muted: %w", err)
		}
		t.IsMuted = value
	}
	return nil
}

// EncodeTDLibJSON encodes t in TDLib API JSON format.
func (t *ToggleGroupCallParticipantIsMutedRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleGroupCallParticipantIsMuted#b2081407 as nil")
	}
	b.ObjStart()
	b.PutID("toggleGroupCallParticipantIsMuted")
	b.FieldStart("group_call_id")
	b.PutInt32(t.GroupCallID)
	b.FieldStart("participant_id")
	if t.ParticipantID == nil {
		return fmt.Errorf("unable to encode toggleGroupCallParticipantIsMuted#b2081407: field participant_id is nil")
	}
	if err := t.ParticipantID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode toggleGroupCallParticipantIsMuted#b2081407: field participant_id: %w", err)
	}
	b.FieldStart("is_muted")
	b.PutBool(t.IsMuted)
	b.ObjEnd()
	return nil
}

// GetGroupCallID returns value of GroupCallID field.
func (t *ToggleGroupCallParticipantIsMutedRequest) GetGroupCallID() (value int32) {
	return t.GroupCallID
}

// GetParticipantID returns value of ParticipantID field.
func (t *ToggleGroupCallParticipantIsMutedRequest) GetParticipantID() (value MessageSenderClass) {
	return t.ParticipantID
}

// GetIsMuted returns value of IsMuted field.
func (t *ToggleGroupCallParticipantIsMutedRequest) GetIsMuted() (value bool) {
	return t.IsMuted
}

// ToggleGroupCallParticipantIsMuted invokes method toggleGroupCallParticipantIsMuted#b2081407 returning error if any.
func (c *Client) ToggleGroupCallParticipantIsMuted(ctx context.Context, request *ToggleGroupCallParticipantIsMutedRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
