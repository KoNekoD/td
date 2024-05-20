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

// ReqPqRequest represents TL type `req_pq#60469778`.
type ReqPqRequest struct {
	// Nonce field of ReqPqRequest.
	Nonce bin.Int128
}

// ReqPqRequestTypeID is TL type id of ReqPqRequest.
const ReqPqRequestTypeID = 0x60469778

// Ensuring interfaces in compile-time for ReqPqRequest.
var (
	_ bin.Encoder     = &ReqPqRequest{}
	_ bin.Decoder     = &ReqPqRequest{}
	_ bin.BareEncoder = &ReqPqRequest{}
	_ bin.BareDecoder = &ReqPqRequest{}
)

func (r *ReqPqRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Nonce == bin.Int128{}) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReqPqRequest) String() string {
	if r == nil {
		return "ReqPqRequest(nil)"
	}
	type Alias ReqPqRequest
	return fmt.Sprintf("ReqPqRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReqPqRequest) TypeID() uint32 {
	return ReqPqRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReqPqRequest) TypeName() string {
	return "req_pq"
}

// TypeInfo returns info about TL type.
func (r *ReqPqRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "req_pq",
		ID:   ReqPqRequestTypeID,
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
func (r *ReqPqRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode req_pq#60469778 as nil")
	}
	b.PutID(ReqPqRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReqPqRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode req_pq#60469778 as nil")
	}
	b.PutInt128(r.Nonce)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReqPqRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode req_pq#60469778 to nil")
	}
	if err := b.ConsumeID(ReqPqRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode req_pq#60469778: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReqPqRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode req_pq#60469778 to nil")
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode req_pq#60469778: field nonce: %w", err)
		}
		r.Nonce = value
	}
	return nil
}

// GetNonce returns value of Nonce field.
func (r *ReqPqRequest) GetNonce() (value bin.Int128) {
	if r == nil {
		return
	}
	return r.Nonce
}
