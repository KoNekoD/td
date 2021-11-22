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

// SearchChatMembersRequest represents TL type `searchChatMembers#e56d46c5`.
type SearchChatMembersRequest struct {
	// Chat identifier
	ChatID int64
	// Query to search for
	Query string
	// The maximum number of users to be returned
	Limit int32
	// The type of users to return. By default, chatMembersFilterMembers
	Filter ChatMembersFilterClass
}

// SearchChatMembersRequestTypeID is TL type id of SearchChatMembersRequest.
const SearchChatMembersRequestTypeID = 0xe56d46c5

// Ensuring interfaces in compile-time for SearchChatMembersRequest.
var (
	_ bin.Encoder     = &SearchChatMembersRequest{}
	_ bin.Decoder     = &SearchChatMembersRequest{}
	_ bin.BareEncoder = &SearchChatMembersRequest{}
	_ bin.BareDecoder = &SearchChatMembersRequest{}
)

func (s *SearchChatMembersRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatID == 0) {
		return false
	}
	if !(s.Query == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}
	if !(s.Filter == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchChatMembersRequest) String() string {
	if s == nil {
		return "SearchChatMembersRequest(nil)"
	}
	type Alias SearchChatMembersRequest
	return fmt.Sprintf("SearchChatMembersRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchChatMembersRequest) TypeID() uint32 {
	return SearchChatMembersRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchChatMembersRequest) TypeName() string {
	return "searchChatMembers"
}

// TypeInfo returns info about TL type.
func (s *SearchChatMembersRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchChatMembers",
		ID:   SearchChatMembersRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchChatMembersRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchChatMembers#e56d46c5 as nil")
	}
	b.PutID(SearchChatMembersRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchChatMembersRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchChatMembers#e56d46c5 as nil")
	}
	b.PutLong(s.ChatID)
	b.PutString(s.Query)
	b.PutInt32(s.Limit)
	if s.Filter == nil {
		return fmt.Errorf("unable to encode searchChatMembers#e56d46c5: field filter is nil")
	}
	if err := s.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode searchChatMembers#e56d46c5: field filter: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchChatMembersRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchChatMembers#e56d46c5 to nil")
	}
	if err := b.ConsumeID(SearchChatMembersRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchChatMembers#e56d46c5: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchChatMembersRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchChatMembers#e56d46c5 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode searchChatMembers#e56d46c5: field chat_id: %w", err)
		}
		s.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchChatMembers#e56d46c5: field query: %w", err)
		}
		s.Query = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchChatMembers#e56d46c5: field limit: %w", err)
		}
		s.Limit = value
	}
	{
		value, err := DecodeChatMembersFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode searchChatMembers#e56d46c5: field filter: %w", err)
		}
		s.Filter = value
	}
	return nil
}

// EncodeTDLibJSON encodes s in TDLib API JSON format.
func (s *SearchChatMembersRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchChatMembers#e56d46c5 as nil")
	}
	b.ObjStart()
	b.PutID("searchChatMembers")
	b.FieldStart("chat_id")
	b.PutLong(s.ChatID)
	b.FieldStart("query")
	b.PutString(s.Query)
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.FieldStart("filter")
	if s.Filter == nil {
		return fmt.Errorf("unable to encode searchChatMembers#e56d46c5: field filter is nil")
	}
	if err := s.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode searchChatMembers#e56d46c5: field filter: %w", err)
	}
	b.ObjEnd()
	return nil
}

// GetChatID returns value of ChatID field.
func (s *SearchChatMembersRequest) GetChatID() (value int64) {
	return s.ChatID
}

// GetQuery returns value of Query field.
func (s *SearchChatMembersRequest) GetQuery() (value string) {
	return s.Query
}

// GetLimit returns value of Limit field.
func (s *SearchChatMembersRequest) GetLimit() (value int32) {
	return s.Limit
}

// GetFilter returns value of Filter field.
func (s *SearchChatMembersRequest) GetFilter() (value ChatMembersFilterClass) {
	return s.Filter
}

// SearchChatMembers invokes method searchChatMembers#e56d46c5 returning error if any.
func (c *Client) SearchChatMembers(ctx context.Context, request *SearchChatMembersRequest) (*ChatMembers, error) {
	var result ChatMembers

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
