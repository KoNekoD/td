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

// PhoneStartScheduledGroupCallRequest represents TL type `phone.startScheduledGroupCall#5680e342`.
//
// See https://core.telegram.org/method/phone.startScheduledGroupCall for reference.
type PhoneStartScheduledGroupCallRequest struct {
	// Call field of PhoneStartScheduledGroupCallRequest.
	Call InputGroupCall
}

// PhoneStartScheduledGroupCallRequestTypeID is TL type id of PhoneStartScheduledGroupCallRequest.
const PhoneStartScheduledGroupCallRequestTypeID = 0x5680e342

// Ensuring interfaces in compile-time for PhoneStartScheduledGroupCallRequest.
var (
	_ bin.Encoder     = &PhoneStartScheduledGroupCallRequest{}
	_ bin.Decoder     = &PhoneStartScheduledGroupCallRequest{}
	_ bin.BareEncoder = &PhoneStartScheduledGroupCallRequest{}
	_ bin.BareDecoder = &PhoneStartScheduledGroupCallRequest{}
)

func (s *PhoneStartScheduledGroupCallRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Call.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *PhoneStartScheduledGroupCallRequest) String() string {
	if s == nil {
		return "PhoneStartScheduledGroupCallRequest(nil)"
	}
	type Alias PhoneStartScheduledGroupCallRequest
	return fmt.Sprintf("PhoneStartScheduledGroupCallRequest%+v", Alias(*s))
}

// FillFrom fills PhoneStartScheduledGroupCallRequest from given interface.
func (s *PhoneStartScheduledGroupCallRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
}) {
	s.Call = from.GetCall()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneStartScheduledGroupCallRequest) TypeID() uint32 {
	return PhoneStartScheduledGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneStartScheduledGroupCallRequest) TypeName() string {
	return "phone.startScheduledGroupCall"
}

// TypeInfo returns info about TL type.
func (s *PhoneStartScheduledGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.startScheduledGroupCall",
		ID:   PhoneStartScheduledGroupCallRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *PhoneStartScheduledGroupCallRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.startScheduledGroupCall#5680e342 as nil")
	}
	b.PutID(PhoneStartScheduledGroupCallRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *PhoneStartScheduledGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode phone.startScheduledGroupCall#5680e342 as nil")
	}
	if err := s.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.startScheduledGroupCall#5680e342: field call: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *PhoneStartScheduledGroupCallRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.startScheduledGroupCall#5680e342 to nil")
	}
	if err := b.ConsumeID(PhoneStartScheduledGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.startScheduledGroupCall#5680e342: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *PhoneStartScheduledGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode phone.startScheduledGroupCall#5680e342 to nil")
	}
	{
		if err := s.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.startScheduledGroupCall#5680e342: field call: %w", err)
		}
	}
	return nil
}

// GetCall returns value of Call field.
func (s *PhoneStartScheduledGroupCallRequest) GetCall() (value InputGroupCall) {
	return s.Call
}

// PhoneStartScheduledGroupCall invokes method phone.startScheduledGroupCall#5680e342 returning error if any.
//
// See https://core.telegram.org/method/phone.startScheduledGroupCall for reference.
func (c *Client) PhoneStartScheduledGroupCall(ctx context.Context, call InputGroupCall) (UpdatesClass, error) {
	var result UpdatesBox

	request := &PhoneStartScheduledGroupCallRequest{
		Call: call,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
