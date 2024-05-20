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

// ClearAutosaveSettingsExceptionsRequest represents TL type `clearAutosaveSettingsExceptions#57ec63f2`.
type ClearAutosaveSettingsExceptionsRequest struct {
}

// ClearAutosaveSettingsExceptionsRequestTypeID is TL type id of ClearAutosaveSettingsExceptionsRequest.
const ClearAutosaveSettingsExceptionsRequestTypeID = 0x57ec63f2

// Ensuring interfaces in compile-time for ClearAutosaveSettingsExceptionsRequest.
var (
	_ bin.Encoder     = &ClearAutosaveSettingsExceptionsRequest{}
	_ bin.Decoder     = &ClearAutosaveSettingsExceptionsRequest{}
	_ bin.BareEncoder = &ClearAutosaveSettingsExceptionsRequest{}
	_ bin.BareDecoder = &ClearAutosaveSettingsExceptionsRequest{}
)

func (c *ClearAutosaveSettingsExceptionsRequest) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ClearAutosaveSettingsExceptionsRequest) String() string {
	if c == nil {
		return "ClearAutosaveSettingsExceptionsRequest(nil)"
	}
	type Alias ClearAutosaveSettingsExceptionsRequest
	return fmt.Sprintf("ClearAutosaveSettingsExceptionsRequest%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ClearAutosaveSettingsExceptionsRequest) TypeID() uint32 {
	return ClearAutosaveSettingsExceptionsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ClearAutosaveSettingsExceptionsRequest) TypeName() string {
	return "clearAutosaveSettingsExceptions"
}

// TypeInfo returns info about TL type.
func (c *ClearAutosaveSettingsExceptionsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "clearAutosaveSettingsExceptions",
		ID:   ClearAutosaveSettingsExceptionsRequestTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ClearAutosaveSettingsExceptionsRequest) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode clearAutosaveSettingsExceptions#57ec63f2 as nil")
	}
	b.PutID(ClearAutosaveSettingsExceptionsRequestTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ClearAutosaveSettingsExceptionsRequest) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode clearAutosaveSettingsExceptions#57ec63f2 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ClearAutosaveSettingsExceptionsRequest) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode clearAutosaveSettingsExceptions#57ec63f2 to nil")
	}
	if err := b.ConsumeID(ClearAutosaveSettingsExceptionsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode clearAutosaveSettingsExceptions#57ec63f2: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ClearAutosaveSettingsExceptionsRequest) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode clearAutosaveSettingsExceptions#57ec63f2 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ClearAutosaveSettingsExceptionsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode clearAutosaveSettingsExceptions#57ec63f2 as nil")
	}
	b.ObjStart()
	b.PutID("clearAutosaveSettingsExceptions")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ClearAutosaveSettingsExceptionsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode clearAutosaveSettingsExceptions#57ec63f2 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("clearAutosaveSettingsExceptions"); err != nil {
				return fmt.Errorf("unable to decode clearAutosaveSettingsExceptions#57ec63f2: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ClearAutosaveSettingsExceptions invokes method clearAutosaveSettingsExceptions#57ec63f2 returning error if any.
func (c *Client) ClearAutosaveSettingsExceptions(ctx context.Context) error {
	var ok Ok

	request := &ClearAutosaveSettingsExceptionsRequest{}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
