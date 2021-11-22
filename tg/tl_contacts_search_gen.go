// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// ContactsSearchRequest represents TL type `contacts.search#11f812d8`.
// Returns users found by username substring.
//
// See https://core.telegram.org/method/contacts.search for reference.
type ContactsSearchRequest struct {
	// Target substring
	Q string
	// Maximum number of users to be returned
	Limit int
}

// ContactsSearchRequestTypeID is TL type id of ContactsSearchRequest.
const ContactsSearchRequestTypeID = 0x11f812d8

// Ensuring interfaces in compile-time for ContactsSearchRequest.
var (
	_ bin.Encoder     = &ContactsSearchRequest{}
	_ bin.Decoder     = &ContactsSearchRequest{}
	_ bin.BareEncoder = &ContactsSearchRequest{}
	_ bin.BareDecoder = &ContactsSearchRequest{}
)

func (s *ContactsSearchRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Q == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *ContactsSearchRequest) String() string {
	if s == nil {
		return "ContactsSearchRequest(nil)"
	}
	type Alias ContactsSearchRequest
	return fmt.Sprintf("ContactsSearchRequest%+v", Alias(*s))
}

// FillFrom fills ContactsSearchRequest from given interface.
func (s *ContactsSearchRequest) FillFrom(from interface {
	GetQ() (value string)
	GetLimit() (value int)
}) {
	s.Q = from.GetQ()
	s.Limit = from.GetLimit()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsSearchRequest) TypeID() uint32 {
	return ContactsSearchRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsSearchRequest) TypeName() string {
	return "contacts.search"
}

// TypeInfo returns info about TL type.
func (s *ContactsSearchRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.search",
		ID:   ContactsSearchRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Q",
			SchemaName: "q",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *ContactsSearchRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode contacts.search#11f812d8 as nil")
	}
	b.PutID(ContactsSearchRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *ContactsSearchRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode contacts.search#11f812d8 as nil")
	}
	b.PutString(s.Q)
	b.PutInt(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *ContactsSearchRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode contacts.search#11f812d8 to nil")
	}
	if err := b.ConsumeID(ContactsSearchRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.search#11f812d8: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *ContactsSearchRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode contacts.search#11f812d8 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.search#11f812d8: field q: %w", err)
		}
		s.Q = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.search#11f812d8: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// GetQ returns value of Q field.
func (s *ContactsSearchRequest) GetQ() (value string) {
	return s.Q
}

// GetLimit returns value of Limit field.
func (s *ContactsSearchRequest) GetLimit() (value int) {
	return s.Limit
}

// ContactsSearch invokes method contacts.search#11f812d8 returning error if any.
// Returns users found by username substring.
//
// Possible errors:
//  400 QUERY_TOO_SHORT: The query string is too short
//  400 SEARCH_QUERY_EMPTY: The search query is empty
//
// See https://core.telegram.org/method/contacts.search for reference.
func (c *Client) ContactsSearch(ctx context.Context, request *ContactsSearchRequest) (*ContactsFound, error) {
	var result ContactsFound

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
