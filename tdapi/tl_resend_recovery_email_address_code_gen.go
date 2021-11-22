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

// ResendRecoveryEmailAddressCodeRequest represents TL type `resendRecoveryEmailAddressCode#19d66f1c`.
type ResendRecoveryEmailAddressCodeRequest struct {
}

// ResendRecoveryEmailAddressCodeRequestTypeID is TL type id of ResendRecoveryEmailAddressCodeRequest.
const ResendRecoveryEmailAddressCodeRequestTypeID = 0x19d66f1c

// Ensuring interfaces in compile-time for ResendRecoveryEmailAddressCodeRequest.
var (
	_ bin.Encoder     = &ResendRecoveryEmailAddressCodeRequest{}
	_ bin.Decoder     = &ResendRecoveryEmailAddressCodeRequest{}
	_ bin.BareEncoder = &ResendRecoveryEmailAddressCodeRequest{}
	_ bin.BareDecoder = &ResendRecoveryEmailAddressCodeRequest{}
)

func (r *ResendRecoveryEmailAddressCodeRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResendRecoveryEmailAddressCodeRequest) String() string {
	if r == nil {
		return "ResendRecoveryEmailAddressCodeRequest(nil)"
	}
	type Alias ResendRecoveryEmailAddressCodeRequest
	return fmt.Sprintf("ResendRecoveryEmailAddressCodeRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResendRecoveryEmailAddressCodeRequest) TypeID() uint32 {
	return ResendRecoveryEmailAddressCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ResendRecoveryEmailAddressCodeRequest) TypeName() string {
	return "resendRecoveryEmailAddressCode"
}

// TypeInfo returns info about TL type.
func (r *ResendRecoveryEmailAddressCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resendRecoveryEmailAddressCode",
		ID:   ResendRecoveryEmailAddressCodeRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResendRecoveryEmailAddressCodeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendRecoveryEmailAddressCode#19d66f1c as nil")
	}
	b.PutID(ResendRecoveryEmailAddressCodeRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResendRecoveryEmailAddressCodeRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendRecoveryEmailAddressCode#19d66f1c as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ResendRecoveryEmailAddressCodeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendRecoveryEmailAddressCode#19d66f1c to nil")
	}
	if err := b.ConsumeID(ResendRecoveryEmailAddressCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode resendRecoveryEmailAddressCode#19d66f1c: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResendRecoveryEmailAddressCodeRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendRecoveryEmailAddressCode#19d66f1c to nil")
	}
	return nil
}

// EncodeTDLibJSON encodes r in TDLib API JSON format.
func (r *ResendRecoveryEmailAddressCodeRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode resendRecoveryEmailAddressCode#19d66f1c as nil")
	}
	b.ObjStart()
	b.PutID("resendRecoveryEmailAddressCode")
	b.ObjEnd()
	return nil
}

// ResendRecoveryEmailAddressCode invokes method resendRecoveryEmailAddressCode#19d66f1c returning error if any.
func (c *Client) ResendRecoveryEmailAddressCode(ctx context.Context) (*PasswordState, error) {
	var result PasswordState

	request := &ResendRecoveryEmailAddressCodeRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
