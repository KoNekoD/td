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

// RecoverAuthenticationPasswordRequest represents TL type `recoverAuthenticationPassword#f8311523`.
type RecoverAuthenticationPasswordRequest struct {
	// Recovery code to check
	RecoveryCode string
	// New password of the user; may be empty to remove the password
	NewPassword string
	// New password hint; may be empty
	NewHint string
}

// RecoverAuthenticationPasswordRequestTypeID is TL type id of RecoverAuthenticationPasswordRequest.
const RecoverAuthenticationPasswordRequestTypeID = 0xf8311523

// Ensuring interfaces in compile-time for RecoverAuthenticationPasswordRequest.
var (
	_ bin.Encoder     = &RecoverAuthenticationPasswordRequest{}
	_ bin.Decoder     = &RecoverAuthenticationPasswordRequest{}
	_ bin.BareEncoder = &RecoverAuthenticationPasswordRequest{}
	_ bin.BareDecoder = &RecoverAuthenticationPasswordRequest{}
)

func (r *RecoverAuthenticationPasswordRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.RecoveryCode == "") {
		return false
	}
	if !(r.NewPassword == "") {
		return false
	}
	if !(r.NewHint == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *RecoverAuthenticationPasswordRequest) String() string {
	if r == nil {
		return "RecoverAuthenticationPasswordRequest(nil)"
	}
	type Alias RecoverAuthenticationPasswordRequest
	return fmt.Sprintf("RecoverAuthenticationPasswordRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecoverAuthenticationPasswordRequest) TypeID() uint32 {
	return RecoverAuthenticationPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RecoverAuthenticationPasswordRequest) TypeName() string {
	return "recoverAuthenticationPassword"
}

// TypeInfo returns info about TL type.
func (r *RecoverAuthenticationPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recoverAuthenticationPassword",
		ID:   RecoverAuthenticationPasswordRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RecoveryCode",
			SchemaName: "recovery_code",
		},
		{
			Name:       "NewPassword",
			SchemaName: "new_password",
		},
		{
			Name:       "NewHint",
			SchemaName: "new_hint",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *RecoverAuthenticationPasswordRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recoverAuthenticationPassword#f8311523 as nil")
	}
	b.PutID(RecoverAuthenticationPasswordRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecoverAuthenticationPasswordRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recoverAuthenticationPassword#f8311523 as nil")
	}
	b.PutString(r.RecoveryCode)
	b.PutString(r.NewPassword)
	b.PutString(r.NewHint)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecoverAuthenticationPasswordRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recoverAuthenticationPassword#f8311523 to nil")
	}
	if err := b.ConsumeID(RecoverAuthenticationPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode recoverAuthenticationPassword#f8311523: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecoverAuthenticationPasswordRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recoverAuthenticationPassword#f8311523 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recoverAuthenticationPassword#f8311523: field recovery_code: %w", err)
		}
		r.RecoveryCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recoverAuthenticationPassword#f8311523: field new_password: %w", err)
		}
		r.NewPassword = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recoverAuthenticationPassword#f8311523: field new_hint: %w", err)
		}
		r.NewHint = value
	}
	return nil
}

// EncodeTDLibJSON encodes r in TDLib API JSON format.
func (r *RecoverAuthenticationPasswordRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode recoverAuthenticationPassword#f8311523 as nil")
	}
	b.ObjStart()
	b.PutID("recoverAuthenticationPassword")
	b.FieldStart("recovery_code")
	b.PutString(r.RecoveryCode)
	b.FieldStart("new_password")
	b.PutString(r.NewPassword)
	b.FieldStart("new_hint")
	b.PutString(r.NewHint)
	b.ObjEnd()
	return nil
}

// GetRecoveryCode returns value of RecoveryCode field.
func (r *RecoverAuthenticationPasswordRequest) GetRecoveryCode() (value string) {
	return r.RecoveryCode
}

// GetNewPassword returns value of NewPassword field.
func (r *RecoverAuthenticationPasswordRequest) GetNewPassword() (value string) {
	return r.NewPassword
}

// GetNewHint returns value of NewHint field.
func (r *RecoverAuthenticationPasswordRequest) GetNewHint() (value string) {
	return r.NewHint
}

// RecoverAuthenticationPassword invokes method recoverAuthenticationPassword#f8311523 returning error if any.
func (c *Client) RecoverAuthenticationPassword(ctx context.Context, request *RecoverAuthenticationPasswordRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
