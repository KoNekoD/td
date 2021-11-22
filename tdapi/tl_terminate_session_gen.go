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

// TerminateSessionRequest represents TL type `terminateSession#e7b7c92c`.
type TerminateSessionRequest struct {
	// Session identifier
	SessionID Int64
}

// TerminateSessionRequestTypeID is TL type id of TerminateSessionRequest.
const TerminateSessionRequestTypeID = 0xe7b7c92c

// Ensuring interfaces in compile-time for TerminateSessionRequest.
var (
	_ bin.Encoder     = &TerminateSessionRequest{}
	_ bin.Decoder     = &TerminateSessionRequest{}
	_ bin.BareEncoder = &TerminateSessionRequest{}
	_ bin.BareDecoder = &TerminateSessionRequest{}
)

func (t *TerminateSessionRequest) Zero() bool {
	if t == nil {
		return true
	}
	if !(t.SessionID.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (t *TerminateSessionRequest) String() string {
	if t == nil {
		return "TerminateSessionRequest(nil)"
	}
	type Alias TerminateSessionRequest
	return fmt.Sprintf("TerminateSessionRequest%+v", Alias(*t))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TerminateSessionRequest) TypeID() uint32 {
	return TerminateSessionRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TerminateSessionRequest) TypeName() string {
	return "terminateSession"
}

// TypeInfo returns info about TL type.
func (t *TerminateSessionRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "terminateSession",
		ID:   TerminateSessionRequestTypeID,
	}
	if t == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SessionID",
			SchemaName: "session_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (t *TerminateSessionRequest) Encode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode terminateSession#e7b7c92c as nil")
	}
	b.PutID(TerminateSessionRequestTypeID)
	return t.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (t *TerminateSessionRequest) EncodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't encode terminateSession#e7b7c92c as nil")
	}
	if err := t.SessionID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode terminateSession#e7b7c92c: field session_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (t *TerminateSessionRequest) Decode(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode terminateSession#e7b7c92c to nil")
	}
	if err := b.ConsumeID(TerminateSessionRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode terminateSession#e7b7c92c: %w", err)
	}
	return t.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (t *TerminateSessionRequest) DecodeBare(b *bin.Buffer) error {
	if t == nil {
		return fmt.Errorf("can't decode terminateSession#e7b7c92c to nil")
	}
	{
		if err := t.SessionID.Decode(b); err != nil {
			return fmt.Errorf("unable to decode terminateSession#e7b7c92c: field session_id: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON encodes t in TDLib API JSON format.
func (t *TerminateSessionRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode terminateSession#e7b7c92c as nil")
	}
	b.ObjStart()
	b.PutID("terminateSession")
	b.FieldStart("session_id")
	if err := t.SessionID.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode terminateSession#e7b7c92c: field session_id: %w", err)
	}
	b.ObjEnd()
	return nil
}

// GetSessionID returns value of SessionID field.
func (t *TerminateSessionRequest) GetSessionID() (value Int64) {
	return t.SessionID
}

// TerminateSession invokes method terminateSession#e7b7c92c returning error if any.
func (c *Client) TerminateSession(ctx context.Context, sessionid Int64) error {
	var ok Ok

	request := &TerminateSessionRequest{
		SessionID: sessionid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
