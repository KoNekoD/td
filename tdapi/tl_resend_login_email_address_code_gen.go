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

// ResendLoginEmailAddressCodeRequest represents TL type `resendLoginEmailAddressCode#11765215`.
type ResendLoginEmailAddressCodeRequest struct {
}

// ResendLoginEmailAddressCodeRequestTypeID is TL type id of ResendLoginEmailAddressCodeRequest.
const ResendLoginEmailAddressCodeRequestTypeID = 0x11765215

// Ensuring interfaces in compile-time for ResendLoginEmailAddressCodeRequest.
var (
	_ bin.Encoder     = &ResendLoginEmailAddressCodeRequest{}
	_ bin.Decoder     = &ResendLoginEmailAddressCodeRequest{}
	_ bin.BareEncoder = &ResendLoginEmailAddressCodeRequest{}
	_ bin.BareDecoder = &ResendLoginEmailAddressCodeRequest{}
)

func (r *ResendLoginEmailAddressCodeRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *ResendLoginEmailAddressCodeRequest) String() string {
	if r == nil {
		return "ResendLoginEmailAddressCodeRequest(nil)"
	}
	type Alias ResendLoginEmailAddressCodeRequest
	return fmt.Sprintf("ResendLoginEmailAddressCodeRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ResendLoginEmailAddressCodeRequest) TypeID() uint32 {
	return ResendLoginEmailAddressCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ResendLoginEmailAddressCodeRequest) TypeName() string {
	return "resendLoginEmailAddressCode"
}

// TypeInfo returns info about TL type.
func (r *ResendLoginEmailAddressCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "resendLoginEmailAddressCode",
		ID:   ResendLoginEmailAddressCodeRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (r *ResendLoginEmailAddressCodeRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendLoginEmailAddressCode#11765215 as nil")
	}
	b.PutID(ResendLoginEmailAddressCodeRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ResendLoginEmailAddressCodeRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode resendLoginEmailAddressCode#11765215 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ResendLoginEmailAddressCodeRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendLoginEmailAddressCode#11765215 to nil")
	}
	if err := b.ConsumeID(ResendLoginEmailAddressCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode resendLoginEmailAddressCode#11765215: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ResendLoginEmailAddressCodeRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode resendLoginEmailAddressCode#11765215 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ResendLoginEmailAddressCodeRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode resendLoginEmailAddressCode#11765215 as nil")
	}
	b.ObjStart()
	b.PutID("resendLoginEmailAddressCode")
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (r *ResendLoginEmailAddressCodeRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode resendLoginEmailAddressCode#11765215 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("resendLoginEmailAddressCode"); err != nil {
				return fmt.Errorf("unable to decode resendLoginEmailAddressCode#11765215: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// ResendLoginEmailAddressCode invokes method resendLoginEmailAddressCode#11765215 returning error if any.
func (c *Client) ResendLoginEmailAddressCode(ctx context.Context) (*EmailAddressAuthenticationCodeInfo, error) {
	var result EmailAddressAuthenticationCodeInfo

	request := &ResendLoginEmailAddressCodeRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
