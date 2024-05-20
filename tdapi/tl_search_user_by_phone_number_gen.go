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

// SearchUserByPhoneNumberRequest represents TL type `searchUserByPhoneNumber#a2e22b12`.
type SearchUserByPhoneNumberRequest struct {
	// Phone number to search for
	PhoneNumber string
}

// SearchUserByPhoneNumberRequestTypeID is TL type id of SearchUserByPhoneNumberRequest.
const SearchUserByPhoneNumberRequestTypeID = 0xa2e22b12

// Ensuring interfaces in compile-time for SearchUserByPhoneNumberRequest.
var (
	_ bin.Encoder     = &SearchUserByPhoneNumberRequest{}
	_ bin.Decoder     = &SearchUserByPhoneNumberRequest{}
	_ bin.BareEncoder = &SearchUserByPhoneNumberRequest{}
	_ bin.BareDecoder = &SearchUserByPhoneNumberRequest{}
)

func (s *SearchUserByPhoneNumberRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PhoneNumber == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchUserByPhoneNumberRequest) String() string {
	if s == nil {
		return "SearchUserByPhoneNumberRequest(nil)"
	}
	type Alias SearchUserByPhoneNumberRequest
	return fmt.Sprintf("SearchUserByPhoneNumberRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchUserByPhoneNumberRequest) TypeID() uint32 {
	return SearchUserByPhoneNumberRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchUserByPhoneNumberRequest) TypeName() string {
	return "searchUserByPhoneNumber"
}

// TypeInfo returns info about TL type.
func (s *SearchUserByPhoneNumberRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchUserByPhoneNumber",
		ID:   SearchUserByPhoneNumberRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchUserByPhoneNumberRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchUserByPhoneNumber#a2e22b12 as nil")
	}
	b.PutID(SearchUserByPhoneNumberRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchUserByPhoneNumberRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchUserByPhoneNumber#a2e22b12 as nil")
	}
	b.PutString(s.PhoneNumber)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchUserByPhoneNumberRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchUserByPhoneNumber#a2e22b12 to nil")
	}
	if err := b.ConsumeID(SearchUserByPhoneNumberRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchUserByPhoneNumber#a2e22b12: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchUserByPhoneNumberRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchUserByPhoneNumber#a2e22b12 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchUserByPhoneNumber#a2e22b12: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchUserByPhoneNumberRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchUserByPhoneNumber#a2e22b12 as nil")
	}
	b.ObjStart()
	b.PutID("searchUserByPhoneNumber")
	b.Comma()
	b.FieldStart("phone_number")
	b.PutString(s.PhoneNumber)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchUserByPhoneNumberRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchUserByPhoneNumber#a2e22b12 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchUserByPhoneNumber"); err != nil {
				return fmt.Errorf("unable to decode searchUserByPhoneNumber#a2e22b12: %w", err)
			}
		case "phone_number":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchUserByPhoneNumber#a2e22b12: field phone_number: %w", err)
			}
			s.PhoneNumber = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetPhoneNumber returns value of PhoneNumber field.
func (s *SearchUserByPhoneNumberRequest) GetPhoneNumber() (value string) {
	if s == nil {
		return
	}
	return s.PhoneNumber
}

// SearchUserByPhoneNumber invokes method searchUserByPhoneNumber#a2e22b12 returning error if any.
func (c *Client) SearchUserByPhoneNumber(ctx context.Context, phonenumber string) (*User, error) {
	var result User

	request := &SearchUserByPhoneNumberRequest{
		PhoneNumber: phonenumber,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
