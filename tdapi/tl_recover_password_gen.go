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

// RecoverPasswordRequest represents TL type `recoverPassword#a5259973`.
type RecoverPasswordRequest struct {
	// Recovery code to check
	RecoveryCode string
	// New password of the user; may be empty to remove the password
	NewPassword string
	// New password hint; may be empty
	NewHint string
}

// RecoverPasswordRequestTypeID is TL type id of RecoverPasswordRequest.
const RecoverPasswordRequestTypeID = 0xa5259973

// Ensuring interfaces in compile-time for RecoverPasswordRequest.
var (
	_ bin.Encoder     = &RecoverPasswordRequest{}
	_ bin.Decoder     = &RecoverPasswordRequest{}
	_ bin.BareEncoder = &RecoverPasswordRequest{}
	_ bin.BareDecoder = &RecoverPasswordRequest{}
)

func (r *RecoverPasswordRequest) Zero() bool {
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
func (r *RecoverPasswordRequest) String() string {
	if r == nil {
		return "RecoverPasswordRequest(nil)"
	}
	type Alias RecoverPasswordRequest
	return fmt.Sprintf("RecoverPasswordRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*RecoverPasswordRequest) TypeID() uint32 {
	return RecoverPasswordRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*RecoverPasswordRequest) TypeName() string {
	return "recoverPassword"
}

// TypeInfo returns info about TL type.
func (r *RecoverPasswordRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "recoverPassword",
		ID:   RecoverPasswordRequestTypeID,
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
func (r *RecoverPasswordRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recoverPassword#a5259973 as nil")
	}
	b.PutID(RecoverPasswordRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *RecoverPasswordRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode recoverPassword#a5259973 as nil")
	}
	b.PutString(r.RecoveryCode)
	b.PutString(r.NewPassword)
	b.PutString(r.NewHint)
	return nil
}

// Decode implements bin.Decoder.
func (r *RecoverPasswordRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recoverPassword#a5259973 to nil")
	}
	if err := b.ConsumeID(RecoverPasswordRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode recoverPassword#a5259973: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *RecoverPasswordRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode recoverPassword#a5259973 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recoverPassword#a5259973: field recovery_code: %w", err)
		}
		r.RecoveryCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recoverPassword#a5259973: field new_password: %w", err)
		}
		r.NewPassword = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode recoverPassword#a5259973: field new_hint: %w", err)
		}
		r.NewHint = value
	}
	return nil
}

// EncodeTDLibJSON encodes r in TDLib API JSON format.
func (r *RecoverPasswordRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode recoverPassword#a5259973 as nil")
	}
	b.ObjStart()
	b.PutID("recoverPassword")
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
func (r *RecoverPasswordRequest) GetRecoveryCode() (value string) {
	return r.RecoveryCode
}

// GetNewPassword returns value of NewPassword field.
func (r *RecoverPasswordRequest) GetNewPassword() (value string) {
	return r.NewPassword
}

// GetNewHint returns value of NewHint field.
func (r *RecoverPasswordRequest) GetNewHint() (value string) {
	return r.NewHint
}

// RecoverPassword invokes method recoverPassword#a5259973 returning error if any.
func (c *Client) RecoverPassword(ctx context.Context, request *RecoverPasswordRequest) (*PasswordState, error) {
	var result PasswordState

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
