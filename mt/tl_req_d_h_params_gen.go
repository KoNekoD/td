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

// ReqDHParamsRequest represents TL type `req_DH_params#d712e4be`.
type ReqDHParamsRequest struct {
	// Nonce field of ReqDHParamsRequest.
	Nonce bin.Int128
	// ServerNonce field of ReqDHParamsRequest.
	ServerNonce bin.Int128
	// P field of ReqDHParamsRequest.
	P []byte
	// Q field of ReqDHParamsRequest.
	Q []byte
	// PublicKeyFingerprint field of ReqDHParamsRequest.
	PublicKeyFingerprint int64
	// EncryptedData field of ReqDHParamsRequest.
	EncryptedData []byte
}

// ReqDHParamsRequestTypeID is TL type id of ReqDHParamsRequest.
const ReqDHParamsRequestTypeID = 0xd712e4be

// Ensuring interfaces in compile-time for ReqDHParamsRequest.
var (
	_ bin.Encoder     = &ReqDHParamsRequest{}
	_ bin.Decoder     = &ReqDHParamsRequest{}
	_ bin.BareEncoder = &ReqDHParamsRequest{}
	_ bin.BareDecoder = &ReqDHParamsRequest{}
)

func (r *ReqDHParamsRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Nonce == bin.Int128{}) {
		return false
	}
	if !(r.ServerNonce == bin.Int128{}) {
		return false
	}
	if !(r.P == nil) {
		return false
	}
	if !(r.Q == nil) {
		return false
	}
	if !(r.PublicKeyFingerprint == 0) {
		return false
	}
	if !(r.EncryptedData == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReqDHParamsRequest) String() string {
	if r == nil {
		return "ReqDHParamsRequest(nil)"
	}
	type Alias ReqDHParamsRequest
	return fmt.Sprintf("ReqDHParamsRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReqDHParamsRequest) TypeID() uint32 {
	return ReqDHParamsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReqDHParamsRequest) TypeName() string {
	return "req_DH_params"
}

// TypeInfo returns info about TL type.
func (r *ReqDHParamsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "req_DH_params",
		ID:   ReqDHParamsRequestTypeID,
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
		{
			Name:       "ServerNonce",
			SchemaName: "server_nonce",
		},
		{
			Name:       "P",
			SchemaName: "p",
		},
		{
			Name:       "Q",
			SchemaName: "q",
		},
		{
			Name:       "PublicKeyFingerprint",
			SchemaName: "public_key_fingerprint",
		},
		{
			Name:       "EncryptedData",
			SchemaName: "encrypted_data",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReqDHParamsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode req_DH_params#d712e4be as nil")
	}
	b.PutID(ReqDHParamsRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReqDHParamsRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode req_DH_params#d712e4be as nil")
	}
	b.PutInt128(r.Nonce)
	b.PutInt128(r.ServerNonce)
	b.PutBytes(r.P)
	b.PutBytes(r.Q)
	b.PutLong(r.PublicKeyFingerprint)
	b.PutBytes(r.EncryptedData)
	return nil
}

// Decode implements bin.Decoder.
func (r *ReqDHParamsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode req_DH_params#d712e4be to nil")
	}
	if err := b.ConsumeID(ReqDHParamsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode req_DH_params#d712e4be: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReqDHParamsRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode req_DH_params#d712e4be to nil")
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field nonce: %w", err)
		}
		r.Nonce = value
	}
	{
		value, err := b.Int128()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field server_nonce: %w", err)
		}
		r.ServerNonce = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field p: %w", err)
		}
		r.P = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field q: %w", err)
		}
		r.Q = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field public_key_fingerprint: %w", err)
		}
		r.PublicKeyFingerprint = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode req_DH_params#d712e4be: field encrypted_data: %w", err)
		}
		r.EncryptedData = value
	}
	return nil
}

// GetNonce returns value of Nonce field.
func (r *ReqDHParamsRequest) GetNonce() (value bin.Int128) {
	if r == nil {
		return
	}
	return r.Nonce
}

// GetServerNonce returns value of ServerNonce field.
func (r *ReqDHParamsRequest) GetServerNonce() (value bin.Int128) {
	if r == nil {
		return
	}
	return r.ServerNonce
}

// GetP returns value of P field.
func (r *ReqDHParamsRequest) GetP() (value []byte) {
	if r == nil {
		return
	}
	return r.P
}

// GetQ returns value of Q field.
func (r *ReqDHParamsRequest) GetQ() (value []byte) {
	if r == nil {
		return
	}
	return r.Q
}

// GetPublicKeyFingerprint returns value of PublicKeyFingerprint field.
func (r *ReqDHParamsRequest) GetPublicKeyFingerprint() (value int64) {
	if r == nil {
		return
	}
	return r.PublicKeyFingerprint
}

// GetEncryptedData returns value of EncryptedData field.
func (r *ReqDHParamsRequest) GetEncryptedData() (value []byte) {
	if r == nil {
		return
	}
	return r.EncryptedData
}
