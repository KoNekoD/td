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

// SearchMessagesRequest represents TL type `searchMessages#f2938192`.
type SearchMessagesRequest struct {
	// Chat list in which to search messages; pass null to search in all chats regardless of
	// their chat list. Only Main and Archive chat lists are supported
	ChatList ChatListClass
	// Query to search for
	Query string
	// The date of the message starting from which the results should be fetched. Use 0 or
	// any date in the future to get results from the last message
	OffsetDate int32
	// The chat identifier of the last found message, or 0 for the first request
	OffsetChatID int64
	// The message identifier of the last found message, or 0 for the first request
	OffsetMessageID int64
	// The maximum number of messages to be returned; up to 100. For optimal performance, the
	// number of returned messages is chosen by TDLib and can be smaller than the specified
	// limit
	Limit int32
	// Filter for message content in the search results; searchMessagesFilterCall,
	// searchMessagesFilterMissedCall, searchMessagesFilterMention,
	// searchMessagesFilterUnreadMention, searchMessagesFilterFailedToSend and
	// searchMessagesFilterPinned are unsupported in this function
	Filter SearchMessagesFilterClass
	// If not 0, the minimum date of the messages to return
	MinDate int32
	// If not 0, the maximum date of the messages to return
	MaxDate int32
}

// SearchMessagesRequestTypeID is TL type id of SearchMessagesRequest.
const SearchMessagesRequestTypeID = 0xf2938192

// Ensuring interfaces in compile-time for SearchMessagesRequest.
var (
	_ bin.Encoder     = &SearchMessagesRequest{}
	_ bin.Decoder     = &SearchMessagesRequest{}
	_ bin.BareEncoder = &SearchMessagesRequest{}
	_ bin.BareDecoder = &SearchMessagesRequest{}
)

func (s *SearchMessagesRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.ChatList == nil) {
		return false
	}
	if !(s.Query == "") {
		return false
	}
	if !(s.OffsetDate == 0) {
		return false
	}
	if !(s.OffsetChatID == 0) {
		return false
	}
	if !(s.OffsetMessageID == 0) {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}
	if !(s.Filter == nil) {
		return false
	}
	if !(s.MinDate == 0) {
		return false
	}
	if !(s.MaxDate == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchMessagesRequest) String() string {
	if s == nil {
		return "SearchMessagesRequest(nil)"
	}
	type Alias SearchMessagesRequest
	return fmt.Sprintf("SearchMessagesRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchMessagesRequest) TypeID() uint32 {
	return SearchMessagesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchMessagesRequest) TypeName() string {
	return "searchMessages"
}

// TypeInfo returns info about TL type.
func (s *SearchMessagesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchMessages",
		ID:   SearchMessagesRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatList",
			SchemaName: "chat_list",
		},
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "OffsetDate",
			SchemaName: "offset_date",
		},
		{
			Name:       "OffsetChatID",
			SchemaName: "offset_chat_id",
		},
		{
			Name:       "OffsetMessageID",
			SchemaName: "offset_message_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Filter",
			SchemaName: "filter",
		},
		{
			Name:       "MinDate",
			SchemaName: "min_date",
		},
		{
			Name:       "MaxDate",
			SchemaName: "max_date",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchMessagesRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchMessages#f2938192 as nil")
	}
	b.PutID(SearchMessagesRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchMessagesRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchMessages#f2938192 as nil")
	}
	if s.ChatList == nil {
		return fmt.Errorf("unable to encode searchMessages#f2938192: field chat_list is nil")
	}
	if err := s.ChatList.Encode(b); err != nil {
		return fmt.Errorf("unable to encode searchMessages#f2938192: field chat_list: %w", err)
	}
	b.PutString(s.Query)
	b.PutInt32(s.OffsetDate)
	b.PutLong(s.OffsetChatID)
	b.PutLong(s.OffsetMessageID)
	b.PutInt32(s.Limit)
	if s.Filter == nil {
		return fmt.Errorf("unable to encode searchMessages#f2938192: field filter is nil")
	}
	if err := s.Filter.Encode(b); err != nil {
		return fmt.Errorf("unable to encode searchMessages#f2938192: field filter: %w", err)
	}
	b.PutInt32(s.MinDate)
	b.PutInt32(s.MaxDate)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchMessagesRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchMessages#f2938192 to nil")
	}
	if err := b.ConsumeID(SearchMessagesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchMessages#f2938192: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchMessagesRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchMessages#f2938192 to nil")
	}
	{
		value, err := DecodeChatList(b)
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#f2938192: field chat_list: %w", err)
		}
		s.ChatList = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#f2938192: field query: %w", err)
		}
		s.Query = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#f2938192: field offset_date: %w", err)
		}
		s.OffsetDate = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#f2938192: field offset_chat_id: %w", err)
		}
		s.OffsetChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#f2938192: field offset_message_id: %w", err)
		}
		s.OffsetMessageID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#f2938192: field limit: %w", err)
		}
		s.Limit = value
	}
	{
		value, err := DecodeSearchMessagesFilter(b)
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#f2938192: field filter: %w", err)
		}
		s.Filter = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#f2938192: field min_date: %w", err)
		}
		s.MinDate = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchMessages#f2938192: field max_date: %w", err)
		}
		s.MaxDate = value
	}
	return nil
}

// EncodeTDLibJSON encodes s in TDLib API JSON format.
func (s *SearchMessagesRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchMessages#f2938192 as nil")
	}
	b.ObjStart()
	b.PutID("searchMessages")
	b.FieldStart("chat_list")
	if s.ChatList == nil {
		return fmt.Errorf("unable to encode searchMessages#f2938192: field chat_list is nil")
	}
	if err := s.ChatList.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode searchMessages#f2938192: field chat_list: %w", err)
	}
	b.FieldStart("query")
	b.PutString(s.Query)
	b.FieldStart("offset_date")
	b.PutInt32(s.OffsetDate)
	b.FieldStart("offset_chat_id")
	b.PutLong(s.OffsetChatID)
	b.FieldStart("offset_message_id")
	b.PutLong(s.OffsetMessageID)
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.FieldStart("filter")
	if s.Filter == nil {
		return fmt.Errorf("unable to encode searchMessages#f2938192: field filter is nil")
	}
	if err := s.Filter.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode searchMessages#f2938192: field filter: %w", err)
	}
	b.FieldStart("min_date")
	b.PutInt32(s.MinDate)
	b.FieldStart("max_date")
	b.PutInt32(s.MaxDate)
	b.ObjEnd()
	return nil
}

// GetChatList returns value of ChatList field.
func (s *SearchMessagesRequest) GetChatList() (value ChatListClass) {
	return s.ChatList
}

// GetQuery returns value of Query field.
func (s *SearchMessagesRequest) GetQuery() (value string) {
	return s.Query
}

// GetOffsetDate returns value of OffsetDate field.
func (s *SearchMessagesRequest) GetOffsetDate() (value int32) {
	return s.OffsetDate
}

// GetOffsetChatID returns value of OffsetChatID field.
func (s *SearchMessagesRequest) GetOffsetChatID() (value int64) {
	return s.OffsetChatID
}

// GetOffsetMessageID returns value of OffsetMessageID field.
func (s *SearchMessagesRequest) GetOffsetMessageID() (value int64) {
	return s.OffsetMessageID
}

// GetLimit returns value of Limit field.
func (s *SearchMessagesRequest) GetLimit() (value int32) {
	return s.Limit
}

// GetFilter returns value of Filter field.
func (s *SearchMessagesRequest) GetFilter() (value SearchMessagesFilterClass) {
	return s.Filter
}

// GetMinDate returns value of MinDate field.
func (s *SearchMessagesRequest) GetMinDate() (value int32) {
	return s.MinDate
}

// GetMaxDate returns value of MaxDate field.
func (s *SearchMessagesRequest) GetMaxDate() (value int32) {
	return s.MaxDate
}

// SearchMessages invokes method searchMessages#f2938192 returning error if any.
func (c *Client) SearchMessages(ctx context.Context, request *SearchMessagesRequest) (*Messages, error) {
	var result Messages

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
