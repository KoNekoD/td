// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// PhoneJoinGroupCallRequest represents TL type `phone.joinGroupCall#b132ff7b`.
// Join a group call
//
// See https://core.telegram.org/method/phone.joinGroupCall for reference.
type PhoneJoinGroupCallRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// If set, the user will be muted by default upon joining.
	Muted bool
	// If set, the user's video will be disabled by default upon joining.
	VideoStopped bool
	// The group call
	Call InputGroupCall
	// Join the group call, presenting yourself as the specified user/channel
	JoinAs InputPeerClass
	// The invitation hash from the invite link »¹, if provided allows speaking in a
	// livestream or muted group chat.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#video-chat-livestream-links
	//
	// Use SetInviteHash and GetInviteHash helpers.
	InviteHash string
	// WebRTC parameters
	Params DataJSON
}

// PhoneJoinGroupCallRequestTypeID is TL type id of PhoneJoinGroupCallRequest.
const PhoneJoinGroupCallRequestTypeID = 0xb132ff7b

// Ensuring interfaces in compile-time for PhoneJoinGroupCallRequest.
var (
	_ bin.Encoder     = &PhoneJoinGroupCallRequest{}
	_ bin.Decoder     = &PhoneJoinGroupCallRequest{}
	_ bin.BareEncoder = &PhoneJoinGroupCallRequest{}
	_ bin.BareDecoder = &PhoneJoinGroupCallRequest{}
)

func (j *PhoneJoinGroupCallRequest) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.Flags.Zero()) {
		return false
	}
	if !(j.Muted == false) {
		return false
	}
	if !(j.VideoStopped == false) {
		return false
	}
	if !(j.Call.Zero()) {
		return false
	}
	if !(j.JoinAs == nil) {
		return false
	}
	if !(j.InviteHash == "") {
		return false
	}
	if !(j.Params.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *PhoneJoinGroupCallRequest) String() string {
	if j == nil {
		return "PhoneJoinGroupCallRequest(nil)"
	}
	type Alias PhoneJoinGroupCallRequest
	return fmt.Sprintf("PhoneJoinGroupCallRequest%+v", Alias(*j))
}

// FillFrom fills PhoneJoinGroupCallRequest from given interface.
func (j *PhoneJoinGroupCallRequest) FillFrom(from interface {
	GetMuted() (value bool)
	GetVideoStopped() (value bool)
	GetCall() (value InputGroupCall)
	GetJoinAs() (value InputPeerClass)
	GetInviteHash() (value string, ok bool)
	GetParams() (value DataJSON)
}) {
	j.Muted = from.GetMuted()
	j.VideoStopped = from.GetVideoStopped()
	j.Call = from.GetCall()
	j.JoinAs = from.GetJoinAs()
	if val, ok := from.GetInviteHash(); ok {
		j.InviteHash = val
	}

	j.Params = from.GetParams()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhoneJoinGroupCallRequest) TypeID() uint32 {
	return PhoneJoinGroupCallRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*PhoneJoinGroupCallRequest) TypeName() string {
	return "phone.joinGroupCall"
}

// TypeInfo returns info about TL type.
func (j *PhoneJoinGroupCallRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "phone.joinGroupCall",
		ID:   PhoneJoinGroupCallRequestTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Muted",
			SchemaName: "muted",
			Null:       !j.Flags.Has(0),
		},
		{
			Name:       "VideoStopped",
			SchemaName: "video_stopped",
			Null:       !j.Flags.Has(2),
		},
		{
			Name:       "Call",
			SchemaName: "call",
		},
		{
			Name:       "JoinAs",
			SchemaName: "join_as",
		},
		{
			Name:       "InviteHash",
			SchemaName: "invite_hash",
			Null:       !j.Flags.Has(1),
		},
		{
			Name:       "Params",
			SchemaName: "params",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (j *PhoneJoinGroupCallRequest) SetFlags() {
	if !(j.Muted == false) {
		j.Flags.Set(0)
	}
	if !(j.VideoStopped == false) {
		j.Flags.Set(2)
	}
	if !(j.InviteHash == "") {
		j.Flags.Set(1)
	}
}

// Encode implements bin.Encoder.
func (j *PhoneJoinGroupCallRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode phone.joinGroupCall#b132ff7b as nil")
	}
	b.PutID(PhoneJoinGroupCallRequestTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *PhoneJoinGroupCallRequest) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode phone.joinGroupCall#b132ff7b as nil")
	}
	j.SetFlags()
	if err := j.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.joinGroupCall#b132ff7b: field flags: %w", err)
	}
	if err := j.Call.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.joinGroupCall#b132ff7b: field call: %w", err)
	}
	if j.JoinAs == nil {
		return fmt.Errorf("unable to encode phone.joinGroupCall#b132ff7b: field join_as is nil")
	}
	if err := j.JoinAs.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.joinGroupCall#b132ff7b: field join_as: %w", err)
	}
	if j.Flags.Has(1) {
		b.PutString(j.InviteHash)
	}
	if err := j.Params.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.joinGroupCall#b132ff7b: field params: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (j *PhoneJoinGroupCallRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode phone.joinGroupCall#b132ff7b to nil")
	}
	if err := b.ConsumeID(PhoneJoinGroupCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.joinGroupCall#b132ff7b: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *PhoneJoinGroupCallRequest) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode phone.joinGroupCall#b132ff7b to nil")
	}
	{
		if err := j.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.joinGroupCall#b132ff7b: field flags: %w", err)
		}
	}
	j.Muted = j.Flags.Has(0)
	j.VideoStopped = j.Flags.Has(2)
	{
		if err := j.Call.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.joinGroupCall#b132ff7b: field call: %w", err)
		}
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode phone.joinGroupCall#b132ff7b: field join_as: %w", err)
		}
		j.JoinAs = value
	}
	if j.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode phone.joinGroupCall#b132ff7b: field invite_hash: %w", err)
		}
		j.InviteHash = value
	}
	{
		if err := j.Params.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.joinGroupCall#b132ff7b: field params: %w", err)
		}
	}
	return nil
}

// SetMuted sets value of Muted conditional field.
func (j *PhoneJoinGroupCallRequest) SetMuted(value bool) {
	if value {
		j.Flags.Set(0)
		j.Muted = true
	} else {
		j.Flags.Unset(0)
		j.Muted = false
	}
}

// GetMuted returns value of Muted conditional field.
func (j *PhoneJoinGroupCallRequest) GetMuted() (value bool) {
	if j == nil {
		return
	}
	return j.Flags.Has(0)
}

// SetVideoStopped sets value of VideoStopped conditional field.
func (j *PhoneJoinGroupCallRequest) SetVideoStopped(value bool) {
	if value {
		j.Flags.Set(2)
		j.VideoStopped = true
	} else {
		j.Flags.Unset(2)
		j.VideoStopped = false
	}
}

// GetVideoStopped returns value of VideoStopped conditional field.
func (j *PhoneJoinGroupCallRequest) GetVideoStopped() (value bool) {
	if j == nil {
		return
	}
	return j.Flags.Has(2)
}

// GetCall returns value of Call field.
func (j *PhoneJoinGroupCallRequest) GetCall() (value InputGroupCall) {
	if j == nil {
		return
	}
	return j.Call
}

// GetJoinAs returns value of JoinAs field.
func (j *PhoneJoinGroupCallRequest) GetJoinAs() (value InputPeerClass) {
	if j == nil {
		return
	}
	return j.JoinAs
}

// SetInviteHash sets value of InviteHash conditional field.
func (j *PhoneJoinGroupCallRequest) SetInviteHash(value string) {
	j.Flags.Set(1)
	j.InviteHash = value
}

// GetInviteHash returns value of InviteHash conditional field and
// boolean which is true if field was set.
func (j *PhoneJoinGroupCallRequest) GetInviteHash() (value string, ok bool) {
	if j == nil {
		return
	}
	if !j.Flags.Has(1) {
		return value, false
	}
	return j.InviteHash, true
}

// GetParams returns value of Params field.
func (j *PhoneJoinGroupCallRequest) GetParams() (value DataJSON) {
	if j == nil {
		return
	}
	return j.Params
}

// PhoneJoinGroupCall invokes method phone.joinGroupCall#b132ff7b returning error if any.
// Join a group call
//
// Possible errors:
//
//	400 DATA_JSON_INVALID: The provided JSON data is invalid.
//	403 GROUPCALL_FORBIDDEN: The group call has already ended.
//	400 GROUPCALL_INVALID: The specified group call is invalid.
//	400 GROUPCALL_SSRC_DUPLICATE_MUCH: The app needs to retry joining the group call with a new SSRC value.
//	400 JOIN_AS_PEER_INVALID: The specified peer cannot be used to join a group call.
//
// See https://core.telegram.org/method/phone.joinGroupCall for reference.
func (c *Client) PhoneJoinGroupCall(ctx context.Context, request *PhoneJoinGroupCallRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
