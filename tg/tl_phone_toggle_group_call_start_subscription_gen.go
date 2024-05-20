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

// PhoneToggleGroupCallStartSubscriptionRequest represents TL type `phone.toggleGroupCallStartSubscription#219c34e6`.
// Subscribe or unsubscribe to a scheduled group call
//
// See https://core.telegram.org/method/phone.toggleGroupCallStartSubscription for reference.
type PhoneToggleGroupCallStartSubscriptionRequest struct {
	// Scheduled group call
	Call InputGroupCall
	// Enable or disable subscription
	Subscribed bool
}

// PhoneToggleGroupCallStartSubscriptionRequestTypeID is TL type id of PhoneToggleGroupCallStartSubscriptionRequest.
const PhoneToggleGroupCallStartSubscriptionRequestTypeID = 0x219c34e6

// Ensuring interfaces in compile-time for PhoneToggleGroupCallStartSubscriptionRequest.
var (
	_ bin.Encoder     = &PhoneToggleGroupCallStartSubscriptionRequest{}
	_ bin.Decoder     = &PhoneToggleGroupCallStartSubscriptionRequest{}
	_ bin.BareEncoder = &PhoneToggleGroupCallStartSubscriptionRequest{}
	_ bin.BareDecoder = &PhoneToggleGroupCallStartSubscriptionRequest{}
)

func (t *PhoneToggleGroupCallStartSubscriptionRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.Call.Zero()) {
		return false
	}
	if !(t.Subscribed == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *PhoneToggleGroupCallStartSubscriptionRequest) String() string {
	if t == nil {
		return "PhoneToggleGroupCallStartSubscriptionRequest(nil)"
	}
	type Alias PhoneToggleGroupCallStartSubscriptionRequest
	return fmt.Sprintf("PhoneToggleGroupCallStartSubscriptionRequest%+v", Alias(*t))
}

// FillFrom fills PhoneToggleGroupCallStartSubscriptionRequest from given interface.
func (t *PhoneToggleGroupCallStartSubscriptionRequest) FillFrom(from interface {
	GetCall() (value InputGroupCall)
	GetSubscribed() (value bool)
}) {
	t.Call = from.GetCall()
	t.Subscribed = from.GetSubscribed()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneToggleGroupCallStartSubscriptionRequest) TypeID() uint32 {
	return PhoneToggleGroupCallStartSubscriptionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneToggleGroupCallStartSubscriptionRequest) TypeName() string {
	return "phone.toggleGroupCallStartSubscription"
}

// TypeInfo returns info about TL type.
func (t *PhoneToggleGroupCallStartSubscriptionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.toggleGroupCallStartSubscription",
		ID:   PhoneToggleGroupCallStartSubscriptionRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "Subscribed",
			SchemaName: "subscribed",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *PhoneToggleGroupCallStartSubscriptionRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode phone.toggleGroupCallStartSubscription#219c34e6 as nil")
	}
	b.PutID(PhoneToggleGroupCallStartSubscriptionRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *PhoneToggleGroupCallStartSubscriptionRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode phone.toggleGroupCallStartSubscription#219c34e6 as nil")
	}
	if err := t.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.toggleGroupCallStartSubscription#219c34e6: field call: %w", err)
	}
	b.PutBool(t.Subscribed)
	return nil
}

// Decode implements bin.Decoder.
func (t *PhoneToggleGroupCallStartSubscriptionRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode phone.toggleGroupCallStartSubscription#219c34e6 to nil")
	}
	if err := b.ConsumeID(PhoneToggleGroupCallStartSubscriptionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.toggleGroupCallStartSubscription#219c34e6: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *PhoneToggleGroupCallStartSubscriptionRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode phone.toggleGroupCallStartSubscription#219c34e6 to nil")
	}
	{
		if err := t.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallStartSubscription#219c34e6: field call: %w", err)
		}
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode phone.toggleGroupCallStartSubscription#219c34e6: field subscribed: %w", err)
		}
		t.Subscribed = value
	}
	return nil
}

// GetCall returns value of Call field.
func (t *PhoneToggleGroupCallStartSubscriptionRequest) GetCall() (value InputGroupCall) {
	if t == nil {
		return
	}
	return t.Call
}

// GetSubscribed returns value of Subscribed field.
func (t *PhoneToggleGroupCallStartSubscriptionRequest) GetSubscribed() (value bool) {
	if t == nil {
		return
	}
	return t.Subscribed
}

// PhoneToggleGroupCallStartSubscription invokes method phone.toggleGroupCallStartSubscription#219c34e6 returning error if any.
// Subscribe or unsubscribe to a scheduled group call
//
// Possible errors:
//
//	403 GROUPCALL_ALREADY_STARTED: The groupcall has already started, you can join directly using phone.joinGroupCall.
//
// See https://core.telegram.org/method/phone.toggleGroupCallStartSubscription for reference.
func (c *Client) PhoneToggleGroupCallStartSubscription(ctx context.Context, request *PhoneToggleGroupCallStartSubscriptionRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
