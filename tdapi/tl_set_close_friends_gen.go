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

// SetCloseFriendsRequest represents TL type `setCloseFriends#7ff8f69f`.
type SetCloseFriendsRequest struct {
	// User identifiers of close friends; the users must be contacts of the current user
	UserIDs []int64
}

// SetCloseFriendsRequestTypeID is TL type id of SetCloseFriendsRequest.
const SetCloseFriendsRequestTypeID = 0x7ff8f69f

// Ensuring interfaces in compile-time for SetCloseFriendsRequest.
var (
	_ bin.Encoder     = &SetCloseFriendsRequest{}
	_ bin.Decoder     = &SetCloseFriendsRequest{}
	_ bin.BareEncoder = &SetCloseFriendsRequest{}
	_ bin.BareDecoder = &SetCloseFriendsRequest{}
)

func (s *SetCloseFriendsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.UserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetCloseFriendsRequest) String() string {
	if s == nil {
		return "SetCloseFriendsRequest(nil)"
	}
	type Alias SetCloseFriendsRequest
	return fmt.Sprintf("SetCloseFriendsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetCloseFriendsRequest) TypeID() uint32 {
	return SetCloseFriendsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetCloseFriendsRequest) TypeName() string {
	return "setCloseFriends"
}

// TypeInfo returns info about TL type.
func (s *SetCloseFriendsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setCloseFriends",
		ID:   SetCloseFriendsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "UserIDs",
			SchemaName: "user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetCloseFriendsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setCloseFriends#7ff8f69f as nil")
	}
	b.PutID(SetCloseFriendsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetCloseFriendsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setCloseFriends#7ff8f69f as nil")
	}
	b.PutInt(len(s.UserIDs))
	for _, v := range s.UserIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetCloseFriendsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setCloseFriends#7ff8f69f to nil")
	}
	if err := b.ConsumeID(SetCloseFriendsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setCloseFriends#7ff8f69f: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetCloseFriendsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setCloseFriends#7ff8f69f to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode setCloseFriends#7ff8f69f: field user_ids: %w", err)
		}

		if headerLen > 0 {
			s.UserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setCloseFriends#7ff8f69f: field user_ids: %w", err)
			}
			s.UserIDs = append(s.UserIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetCloseFriendsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setCloseFriends#7ff8f69f as nil")
	}
	b.ObjStart()
	b.PutID("setCloseFriends")
	b.Comma()
	b.FieldStart("user_ids")
	b.ArrStart()
	for _, v := range s.UserIDs {
		b.PutInt53(v)
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetCloseFriendsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setCloseFriends#7ff8f69f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setCloseFriends"); err != nil {
				return fmt.Errorf("unable to decode setCloseFriends#7ff8f69f: %w", err)
			}
		case "user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode setCloseFriends#7ff8f69f: field user_ids: %w", err)
				}
				s.UserIDs = append(s.UserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode setCloseFriends#7ff8f69f: field user_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUserIDs returns value of UserIDs field.
func (s *SetCloseFriendsRequest) GetUserIDs() (value []int64) {
	if s == nil {
		return
	}
	return s.UserIDs
}

// SetCloseFriends invokes method setCloseFriends#7ff8f69f returning error if any.
func (c *Client) SetCloseFriends(ctx context.Context, userids []int64) error {
	var ok Ok

	request := &SetCloseFriendsRequest{
		UserIDs: userids,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
