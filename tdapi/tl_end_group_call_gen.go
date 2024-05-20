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

// EndGroupCallRequest represents TL type `endGroupCall#22294cb7`.
type EndGroupCallRequest struct {
	// Group call identifier
	GroupCallID int32
}

// EndGroupCallRequestTypeID is TL type id of EndGroupCallRequest.
const EndGroupCallRequestTypeID = 0x22294cb7

// Ensuring interfaces in compile-time for EndGroupCallRequest.
var (
	_ bin.Encoder     = &EndGroupCallRequest{}
	_ bin.Decoder     = &EndGroupCallRequest{}
	_ bin.BareEncoder = &EndGroupCallRequest{}
	_ bin.BareDecoder = &EndGroupCallRequest{}
)

func (e *EndGroupCallRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.GroupCallID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EndGroupCallRequest) String() string {
	if e == nil {
		return "EndGroupCallRequest(nil)"
	}
	type Alias EndGroupCallRequest
	return fmt.Sprintf("EndGroupCallRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EndGroupCallRequest) TypeID() uint32 {
	return EndGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EndGroupCallRequest) TypeName() string {
	return "endGroupCall"
}

// TypeInfo returns info about TL type.
func (e *EndGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "endGroupCall",
		ID:   EndGroupCallRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EndGroupCallRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode endGroupCall#22294cb7 as nil")
	}
	b.PutID(EndGroupCallRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EndGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode endGroupCall#22294cb7 as nil")
	}
	b.PutInt32(e.GroupCallID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EndGroupCallRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode endGroupCall#22294cb7 to nil")
	}
	if err := b.ConsumeID(EndGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode endGroupCall#22294cb7: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EndGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode endGroupCall#22294cb7 to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode endGroupCall#22294cb7: field group_call_id: %w", err)
		}
		e.GroupCallID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EndGroupCallRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode endGroupCall#22294cb7 as nil")
	}
	b.ObjStart()
	b.PutID("endGroupCall")
	b.Comma()
	b.FieldStart("group_call_id")
	b.PutInt32(e.GroupCallID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EndGroupCallRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode endGroupCall#22294cb7 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("endGroupCall"); err != nil {
				return fmt.Errorf("unable to decode endGroupCall#22294cb7: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode endGroupCall#22294cb7: field group_call_id: %w", err)
			}
			e.GroupCallID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (e *EndGroupCallRequest) GetGroupCallID() (value int32) {
	if e == nil {
		return
	}
	return e.GroupCallID
}

// EndGroupCall invokes method endGroupCall#22294cb7 returning error if any.
func (c *Client) EndGroupCall(ctx context.Context, groupcallid int32) error {
	var ok Ok

	request := &EndGroupCallRequest{
		GroupCallID: groupcallid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
