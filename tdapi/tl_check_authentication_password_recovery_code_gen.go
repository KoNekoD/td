// Code generated by gotdgen, DO NOT EDIT.

package tdapi

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

// CheckAuthenticationPasswordRecoveryCodeRequest represents TL type `checkAuthenticationPasswordRecoveryCode#dc0a3be5`.
type CheckAuthenticationPasswordRecoveryCodeRequest struct {
	// Recovery code to check
	RecoveryCode string
}

// CheckAuthenticationPasswordRecoveryCodeRequestTypeID is TL type id of CheckAuthenticationPasswordRecoveryCodeRequest.
const CheckAuthenticationPasswordRecoveryCodeRequestTypeID = 0xdc0a3be5

// Ensuring interfaces in compile-time for CheckAuthenticationPasswordRecoveryCodeRequest.
var (
	_ bin.Encoder     = &CheckAuthenticationPasswordRecoveryCodeRequest{}
	_ bin.Decoder     = &CheckAuthenticationPasswordRecoveryCodeRequest{}
	_ bin.BareEncoder = &CheckAuthenticationPasswordRecoveryCodeRequest{}
	_ bin.BareDecoder = &CheckAuthenticationPasswordRecoveryCodeRequest{}
)

func (c *CheckAuthenticationPasswordRecoveryCodeRequest) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.RecoveryCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *CheckAuthenticationPasswordRecoveryCodeRequest) String() string {
	if c == nil {
		return "CheckAuthenticationPasswordRecoveryCodeRequest(nil)"
	}
	type Alias CheckAuthenticationPasswordRecoveryCodeRequest
	return fmt.Sprintf("CheckAuthenticationPasswordRecoveryCodeRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*CheckAuthenticationPasswordRecoveryCodeRequest) TypeID() uint32 {
	return CheckAuthenticationPasswordRecoveryCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*CheckAuthenticationPasswordRecoveryCodeRequest) TypeName() string {
	return "checkAuthenticationPasswordRecoveryCode"
}

// TypeInfo returns info about TL type.
func (c *CheckAuthenticationPasswordRecoveryCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "checkAuthenticationPasswordRecoveryCode",
		ID:   CheckAuthenticationPasswordRecoveryCodeRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RecoveryCode",
			SchemaName: "recovery_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *CheckAuthenticationPasswordRecoveryCodeRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationPasswordRecoveryCode#dc0a3be5 as nil")
	}
	b.PutID(CheckAuthenticationPasswordRecoveryCodeRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *CheckAuthenticationPasswordRecoveryCodeRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationPasswordRecoveryCode#dc0a3be5 as nil")
	}
	b.PutString(c.RecoveryCode)
	return nil
}

// Decode implements bin.Decoder.
func (c *CheckAuthenticationPasswordRecoveryCodeRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationPasswordRecoveryCode#dc0a3be5 to nil")
	}
	if err := b.ConsumeID(CheckAuthenticationPasswordRecoveryCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode checkAuthenticationPasswordRecoveryCode#dc0a3be5: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *CheckAuthenticationPasswordRecoveryCodeRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationPasswordRecoveryCode#dc0a3be5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode checkAuthenticationPasswordRecoveryCode#dc0a3be5: field recovery_code: %w", err)
		}
		c.RecoveryCode = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *CheckAuthenticationPasswordRecoveryCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode checkAuthenticationPasswordRecoveryCode#dc0a3be5 as nil")
	}
	b.ObjStart()
	b.PutID("checkAuthenticationPasswordRecoveryCode")
	b.Comma()
	b.FieldStart("recovery_code")
	b.PutString(c.RecoveryCode)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *CheckAuthenticationPasswordRecoveryCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode checkAuthenticationPasswordRecoveryCode#dc0a3be5 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("checkAuthenticationPasswordRecoveryCode"); err != nil {
				return fmt.Errorf("unable to decode checkAuthenticationPasswordRecoveryCode#dc0a3be5: %w", err)
			}
		case "recovery_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode checkAuthenticationPasswordRecoveryCode#dc0a3be5: field recovery_code: %w", err)
			}
			c.RecoveryCode = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetRecoveryCode returns value of RecoveryCode field.
func (c *CheckAuthenticationPasswordRecoveryCodeRequest) GetRecoveryCode() (value string) {
	if c == nil {
		return
	}
	return c.RecoveryCode
}

// CheckAuthenticationPasswordRecoveryCode invokes method checkAuthenticationPasswordRecoveryCode#dc0a3be5 returning error if any.
func (c *Client) CheckAuthenticationPasswordRecoveryCode(ctx context.Context, recoverycode string) error {
	var ok Ok

	request := &CheckAuthenticationPasswordRecoveryCodeRequest{
		RecoveryCode: recoverycode,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
