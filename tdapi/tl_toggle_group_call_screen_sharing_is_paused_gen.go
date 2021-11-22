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

// ToggleGroupCallScreenSharingIsPausedRequest represents TL type `toggleGroupCallScreenSharingIsPaused#a07b5360`.
type ToggleGroupCallScreenSharingIsPausedRequest struct {
	// Group call identifier
	GroupCallID int32
	// True if screen sharing is paused
	IsPaused bool
}

// ToggleGroupCallScreenSharingIsPausedRequestTypeID is TL type id of ToggleGroupCallScreenSharingIsPausedRequest.
const ToggleGroupCallScreenSharingIsPausedRequestTypeID = 0xa07b5360

// Ensuring interfaces in compile-time for ToggleGroupCallScreenSharingIsPausedRequest.
var (
	_ bin.Encoder     = &ToggleGroupCallScreenSharingIsPausedRequest{}
	_ bin.Decoder     = &ToggleGroupCallScreenSharingIsPausedRequest{}
	_ bin.BareEncoder = &ToggleGroupCallScreenSharingIsPausedRequest{}
	_ bin.BareDecoder = &ToggleGroupCallScreenSharingIsPausedRequest{}
)

func (t *ToggleGroupCallScreenSharingIsPausedRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.GroupCallID == 0) {
		return false
	}
	if !(t.IsPaused == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *ToggleGroupCallScreenSharingIsPausedRequest) String() string {
	if t == nil {
		return "ToggleGroupCallScreenSharingIsPausedRequest(nil)"
	}
	type Alias ToggleGroupCallScreenSharingIsPausedRequest
	return fmt.Sprintf("ToggleGroupCallScreenSharingIsPausedRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ToggleGroupCallScreenSharingIsPausedRequest) TypeID() uint32 {
	return ToggleGroupCallScreenSharingIsPausedRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ToggleGroupCallScreenSharingIsPausedRequest) TypeName() string {
	return "toggleGroupCallScreenSharingIsPaused"
}

// TypeInfo returns info about TL type.
func (t *ToggleGroupCallScreenSharingIsPausedRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "toggleGroupCallScreenSharingIsPaused",
		ID:   ToggleGroupCallScreenSharingIsPausedRequestTypeID,
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
			Name:       "IsPaused",
			SchemaName: "is_paused",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *ToggleGroupCallScreenSharingIsPausedRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleGroupCallScreenSharingIsPaused#a07b5360 as nil")
	}
	b.PutID(ToggleGroupCallScreenSharingIsPausedRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *ToggleGroupCallScreenSharingIsPausedRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleGroupCallScreenSharingIsPaused#a07b5360 as nil")
	}
	b.PutInt32(t.GroupCallID)
	b.PutBool(t.IsPaused)
	return nil
}

// Decode implements bin.Decoder.
func (t *ToggleGroupCallScreenSharingIsPausedRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleGroupCallScreenSharingIsPaused#a07b5360 to nil")
	}
	if err := b.ConsumeID(ToggleGroupCallScreenSharingIsPausedRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode toggleGroupCallScreenSharingIsPaused#a07b5360: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *ToggleGroupCallScreenSharingIsPausedRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode toggleGroupCallScreenSharingIsPaused#a07b5360 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode toggleGroupCallScreenSharingIsPaused#a07b5360: field group_call_id: %w", err)
		}
		t.GroupCallID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode toggleGroupCallScreenSharingIsPaused#a07b5360: field is_paused: %w", err)
		}
		t.IsPaused = value
	}
	return nil
}

// EncodeTDLibJSON encodes t in TDLib API JSON format.
func (t *ToggleGroupCallScreenSharingIsPausedRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode toggleGroupCallScreenSharingIsPaused#a07b5360 as nil")
	}
	b.ObjStart()
	b.PutID("toggleGroupCallScreenSharingIsPaused")
	b.FieldStart("group_call_id")
	b.PutInt32(t.GroupCallID)
	b.FieldStart("is_paused")
	b.PutBool(t.IsPaused)
	b.ObjEnd()
	return nil
}

// GetGroupCallID returns value of GroupCallID field.
func (t *ToggleGroupCallScreenSharingIsPausedRequest) GetGroupCallID() (value int32) {
	return t.GroupCallID
}

// GetIsPaused returns value of IsPaused field.
func (t *ToggleGroupCallScreenSharingIsPausedRequest) GetIsPaused() (value bool) {
	return t.IsPaused
}

// ToggleGroupCallScreenSharingIsPaused invokes method toggleGroupCallScreenSharingIsPaused#a07b5360 returning error if any.
func (c *Client) ToggleGroupCallScreenSharingIsPaused(ctx context.Context, request *ToggleGroupCallScreenSharingIsPausedRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
