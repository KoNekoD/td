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

// SharePhoneNumberRequest represents TL type `sharePhoneNumber#4164e055`.
type SharePhoneNumberRequest struct {
	// Identifier of the user with whom to share the phone number. The user must be a mutual
	// contact
	UserID int64
}

// SharePhoneNumberRequestTypeID is TL type id of SharePhoneNumberRequest.
const SharePhoneNumberRequestTypeID = 0x4164e055

// Ensuring interfaces in compile-time for SharePhoneNumberRequest.
var (
	_ bin.Encoder     = &SharePhoneNumberRequest{}
	_ bin.Decoder     = &SharePhoneNumberRequest{}
	_ bin.BareEncoder = &SharePhoneNumberRequest{}
	_ bin.BareDecoder = &SharePhoneNumberRequest{}
)

func (s *SharePhoneNumberRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SharePhoneNumberRequest) String() string {
	if s == nil {
		return "SharePhoneNumberRequest(nil)"
	}
	type Alias SharePhoneNumberRequest
	return fmt.Sprintf("SharePhoneNumberRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SharePhoneNumberRequest) TypeID() uint32 {
	return SharePhoneNumberRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SharePhoneNumberRequest) TypeName() string {
	return "sharePhoneNumber"
}

// TypeInfo returns info about TL type.
func (s *SharePhoneNumberRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sharePhoneNumber",
		ID:   SharePhoneNumberRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SharePhoneNumberRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sharePhoneNumber#4164e055 as nil")
	}
	b.PutID(SharePhoneNumberRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SharePhoneNumberRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sharePhoneNumber#4164e055 as nil")
	}
	b.PutInt53(s.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (s *SharePhoneNumberRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sharePhoneNumber#4164e055 to nil")
	}
	if err := b.ConsumeID(SharePhoneNumberRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode sharePhoneNumber#4164e055: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SharePhoneNumberRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sharePhoneNumber#4164e055 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode sharePhoneNumber#4164e055: field user_id: %w", err)
		}
		s.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SharePhoneNumberRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sharePhoneNumber#4164e055 as nil")
	}
	b.ObjStart()
	b.PutID("sharePhoneNumber")
	b.Comma()
	b.FieldStart("user_id")
	b.PutInt53(s.UserID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SharePhoneNumberRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sharePhoneNumber#4164e055 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sharePhoneNumber"); err != nil {
				return fmt.Errorf("unable to decode sharePhoneNumber#4164e055: %w", err)
			}
		case "user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode sharePhoneNumber#4164e055: field user_id: %w", err)
			}
			s.UserID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserID returns value of UserID field.
func (s *SharePhoneNumberRequest) GetUserID() (value int64) {
	if s == nil {
		return
	}
	return s.UserID
}

// SharePhoneNumber invokes method sharePhoneNumber#4164e055 returning error if any.
func (c *Client) SharePhoneNumber(ctx context.Context, userid int64) error {
	var ok Ok

	request := &SharePhoneNumberRequest{
		UserID: userid,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
