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

// TerminateSessionRequest represents TL type `terminateSession#e7b7c92c`.
type TerminateSessionRequest struct {
	// Session identifier
	SessionID int64
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
	if !(t.SessionID == 0) {
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
	b.PutLong(t.SessionID)
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
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode terminateSession#e7b7c92c: field session_id: %w", err)
		}
		t.SessionID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (t *TerminateSessionRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if t == nil {
		return fmt.Errorf("can't encode terminateSession#e7b7c92c as nil")
	}
	b.ObjStart()
	b.PutID("terminateSession")
	b.Comma()
	b.FieldStart("session_id")
	b.PutLong(t.SessionID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (t *TerminateSessionRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if t == nil {
		return fmt.Errorf("can't decode terminateSession#e7b7c92c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("terminateSession"); err != nil {
				return fmt.Errorf("unable to decode terminateSession#e7b7c92c: %w", err)
			}
		case "session_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode terminateSession#e7b7c92c: field session_id: %w", err)
			}
			t.SessionID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSessionID returns value of SessionID field.
func (t *TerminateSessionRequest) GetSessionID() (value int64) {
	if t == nil {
		return
	}
	return t.SessionID
}

// TerminateSession invokes method terminateSession#e7b7c92c returning error if any.
func (c *Client) TerminateSession(ctx context.Context, sessionid int64) error {
	var ok Ok

	request := &TerminateSessionRequest{
		SessionID: sessionid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
