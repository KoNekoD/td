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

// JoinChatByInviteLinkRequest represents TL type `joinChatByInviteLink#c16aab86`.
type JoinChatByInviteLinkRequest struct {
	// Invite link to use
	InviteLink string
}

// JoinChatByInviteLinkRequestTypeID is TL type id of JoinChatByInviteLinkRequest.
const JoinChatByInviteLinkRequestTypeID = 0xc16aab86

// Ensuring interfaces in compile-time for JoinChatByInviteLinkRequest.
var (
	_ bin.Encoder     = &JoinChatByInviteLinkRequest{}
	_ bin.Decoder     = &JoinChatByInviteLinkRequest{}
	_ bin.BareEncoder = &JoinChatByInviteLinkRequest{}
	_ bin.BareDecoder = &JoinChatByInviteLinkRequest{}
)

func (j *JoinChatByInviteLinkRequest) Zero() bool {
	if j == nil {
		return true
	}
	if !(j.InviteLink == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (j *JoinChatByInviteLinkRequest) String() string {
	if j == nil {
		return "JoinChatByInviteLinkRequest(nil)"
	}
	type Alias JoinChatByInviteLinkRequest
	return fmt.Sprintf("JoinChatByInviteLinkRequest%+v", Alias(*j))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*JoinChatByInviteLinkRequest) TypeID() uint32 {
	return JoinChatByInviteLinkRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*JoinChatByInviteLinkRequest) TypeName() string {
	return "joinChatByInviteLink"
}

// TypeInfo returns info about TL type.
func (j *JoinChatByInviteLinkRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "joinChatByInviteLink",
		ID:   JoinChatByInviteLinkRequestTypeID,
	}
	if j == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InviteLink",
			SchemaName: "invite_link",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (j *JoinChatByInviteLinkRequest) Encode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode joinChatByInviteLink#c16aab86 as nil")
	}
	b.PutID(JoinChatByInviteLinkRequestTypeID)
	return j.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (j *JoinChatByInviteLinkRequest) EncodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't encode joinChatByInviteLink#c16aab86 as nil")
	}
	b.PutString(j.InviteLink)
	return nil
}

// Decode implements bin.Decoder.
func (j *JoinChatByInviteLinkRequest) Decode(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode joinChatByInviteLink#c16aab86 to nil")
	}
	if err := b.ConsumeID(JoinChatByInviteLinkRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode joinChatByInviteLink#c16aab86: %w", err)
	}
	return j.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (j *JoinChatByInviteLinkRequest) DecodeBare(b *bin.Buffer) error {
	if j == nil {
		return fmt.Errorf("can't decode joinChatByInviteLink#c16aab86 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode joinChatByInviteLink#c16aab86: field invite_link: %w", err)
		}
		j.InviteLink = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (j *JoinChatByInviteLinkRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if j == nil {
		return fmt.Errorf("can't encode joinChatByInviteLink#c16aab86 as nil")
	}
	b.ObjStart()
	b.PutID("joinChatByInviteLink")
	b.Comma()
	b.FieldStart("invite_link")
	b.PutString(j.InviteLink)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (j *JoinChatByInviteLinkRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if j == nil {
		return fmt.Errorf("can't decode joinChatByInviteLink#c16aab86 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("joinChatByInviteLink"); err != nil {
				return fmt.Errorf("unable to decode joinChatByInviteLink#c16aab86: %w", err)
			}
		case "invite_link":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode joinChatByInviteLink#c16aab86: field invite_link: %w", err)
			}
			j.InviteLink = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInviteLink returns value of InviteLink field.
func (j *JoinChatByInviteLinkRequest) GetInviteLink() (value string) {
	if j == nil {
		return
	}
	return j.InviteLink
}

// JoinChatByInviteLink invokes method joinChatByInviteLink#c16aab86 returning error if any.
func (c *Client) JoinChatByInviteLink(ctx context.Context, invitelink string) (*Chat, error) {
	var result Chat

	request := &JoinChatByInviteLinkRequest{
		InviteLink: invitelink,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
