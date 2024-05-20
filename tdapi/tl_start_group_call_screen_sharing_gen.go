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

// StartGroupCallScreenSharingRequest represents TL type `startGroupCallScreenSharing#cb4e312d`.
type StartGroupCallScreenSharingRequest struct {
	// Group call identifier
	GroupCallID int32
	// Screen sharing audio channel synchronization source identifier; received from tgcalls
	AudioSourceID int32
	// Group call join payload; received from tgcalls
	Payload string
}

// StartGroupCallScreenSharingRequestTypeID is TL type id of StartGroupCallScreenSharingRequest.
const StartGroupCallScreenSharingRequestTypeID = 0xcb4e312d

// Ensuring interfaces in compile-time for StartGroupCallScreenSharingRequest.
var (
	_ bin.Encoder     = &StartGroupCallScreenSharingRequest{}
	_ bin.Decoder     = &StartGroupCallScreenSharingRequest{}
	_ bin.BareEncoder = &StartGroupCallScreenSharingRequest{}
	_ bin.BareDecoder = &StartGroupCallScreenSharingRequest{}
)

func (s *StartGroupCallScreenSharingRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.GroupCallID == 0) {
		return false
	}
	if !(s.AudioSourceID == 0) {
		return false
	}
	if !(s.Payload == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StartGroupCallScreenSharingRequest) String() string {
	if s == nil {
		return "StartGroupCallScreenSharingRequest(nil)"
	}
	type Alias StartGroupCallScreenSharingRequest
	return fmt.Sprintf("StartGroupCallScreenSharingRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*StartGroupCallScreenSharingRequest) TypeID() uint32 {
	return StartGroupCallScreenSharingRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*StartGroupCallScreenSharingRequest) TypeName() string {
	return "startGroupCallScreenSharing"
}

// TypeInfo returns info about TL type.
func (s *StartGroupCallScreenSharingRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "startGroupCallScreenSharing",
		ID:   StartGroupCallScreenSharingRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GroupCallID",
			SchemaName: "group_call_id",
		},
		{
			Name:       "AudioSourceID",
			SchemaName: "audio_source_id",
		},
		{
			Name:       "Payload",
			SchemaName: "payload",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *StartGroupCallScreenSharingRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode startGroupCallScreenSharing#cb4e312d as nil")
	}
	b.PutID(StartGroupCallScreenSharingRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *StartGroupCallScreenSharingRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode startGroupCallScreenSharing#cb4e312d as nil")
	}
	b.PutInt32(s.GroupCallID)
	b.PutInt32(s.AudioSourceID)
	b.PutString(s.Payload)
	return nil
}

// Decode implements bin.Decoder.
func (s *StartGroupCallScreenSharingRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode startGroupCallScreenSharing#cb4e312d to nil")
	}
	if err := b.ConsumeID(StartGroupCallScreenSharingRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode startGroupCallScreenSharing#cb4e312d: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *StartGroupCallScreenSharingRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode startGroupCallScreenSharing#cb4e312d to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode startGroupCallScreenSharing#cb4e312d: field group_call_id: %w", err)
		}
		s.GroupCallID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode startGroupCallScreenSharing#cb4e312d: field audio_source_id: %w", err)
		}
		s.AudioSourceID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode startGroupCallScreenSharing#cb4e312d: field payload: %w", err)
		}
		s.Payload = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *StartGroupCallScreenSharingRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode startGroupCallScreenSharing#cb4e312d as nil")
	}
	b.ObjStart()
	b.PutID("startGroupCallScreenSharing")
	b.Comma()
	b.FieldStart("group_call_id")
	b.PutInt32(s.GroupCallID)
	b.Comma()
	b.FieldStart("audio_source_id")
	b.PutInt32(s.AudioSourceID)
	b.Comma()
	b.FieldStart("payload")
	b.PutString(s.Payload)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *StartGroupCallScreenSharingRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode startGroupCallScreenSharing#cb4e312d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("startGroupCallScreenSharing"); err != nil {
				return fmt.Errorf("unable to decode startGroupCallScreenSharing#cb4e312d: %w", err)
			}
		case "group_call_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode startGroupCallScreenSharing#cb4e312d: field group_call_id: %w", err)
			}
			s.GroupCallID = value
		case "audio_source_id":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode startGroupCallScreenSharing#cb4e312d: field audio_source_id: %w", err)
			}
			s.AudioSourceID = value
		case "payload":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode startGroupCallScreenSharing#cb4e312d: field payload: %w", err)
			}
			s.Payload = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGroupCallID returns value of GroupCallID field.
func (s *StartGroupCallScreenSharingRequest) GetGroupCallID() (value int32) {
	if s == nil {
		return
	}
	return s.GroupCallID
}

// GetAudioSourceID returns value of AudioSourceID field.
func (s *StartGroupCallScreenSharingRequest) GetAudioSourceID() (value int32) {
	if s == nil {
		return
	}
	return s.AudioSourceID
}

// GetPayload returns value of Payload field.
func (s *StartGroupCallScreenSharingRequest) GetPayload() (value string) {
	if s == nil {
		return
	}
	return s.Payload
}

// StartGroupCallScreenSharing invokes method startGroupCallScreenSharing#cb4e312d returning error if any.
func (c *Client) StartGroupCallScreenSharing(ctx context.Context, request *StartGroupCallScreenSharingRequest) (*Text, error) {
	var result Text

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
