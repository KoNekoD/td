// Code generated by gotdgen, DO NOT EDIT.

package mt

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

// ReqPqMultiRequest represents TL type `req_pq_multi#be7e8ef1`.
type ReqPqMultiRequest struct {
	// Nonce field of ReqPqMultiRequest.
	Nonce bin.Int128
}

// ReqPqMultiRequestTypeID is TL type id of ReqPqMultiRequest.
const ReqPqMultiRequestTypeID = 0xbe7e8ef1

// Ensuring interfaces in compile-time for ReqPqMultiRequest.
var (
	_ bin.Encoder     = &ReqPqMultiRequest{}
	_ bin.Decoder     = &ReqPqMultiRequest{}
	_ bin.BareEncoder = &ReqPqMultiRequest{}
	_ bin.BareDecoder = &ReqPqMultiRequest{}
)

func (r *ReqPqMultiRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Nonce == bin.Int128{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReqPqMultiRequest) String() string {
	if r == nil {
		return "ReqPqMultiRequest(nil)"
	}
	type Alias ReqPqMultiRequest
	return fmt.Sprintf("ReqPqMultiRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReqPqMultiRequest) TypeID() uint32 {
	return ReqPqMultiRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReqPqMultiRequest) TypeName() string {
	return "req_pq_multi"
}

// TypeInfo returns info about TL type.
func (r *ReqPqMultiRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "req_pq_multi",
		ID:   ReqPqMultiRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Nonce",
			SchemaName: "nonce",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReqPqMultiRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode req_pq_multi#be7e8ef1 as nil")
	}
	b.PutID(ReqPqMultiRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReqPqMultiRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode req_pq_multi#be7e8ef1 as nil")
	}
	b.PutInt128(r.Nonce)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReqPqMultiRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode req_pq_multi#be7e8ef1 to nil")
	}
	if err := b.ConsumeID(ReqPqMultiRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode req_pq_multi#be7e8ef1: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReqPqMultiRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode req_pq_multi#be7e8ef1 to nil")
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode req_pq_multi#be7e8ef1: field nonce: %w", err)
		}
		r.Nonce = value
	}
	return nil
}

// GetNonce returns value of Nonce field.
func (r *ReqPqMultiRequest) GetNonce() (value bin.Int128) {
	if r == nil {
		return
	}
	return r.Nonce
}
