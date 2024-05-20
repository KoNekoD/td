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

// SearchFileDownloadsRequest represents TL type `searchFileDownloads#2a1e0856`.
type SearchFileDownloadsRequest struct {
	// Query to search for; may be empty to return all downloaded files
	Query string
	// Pass true to search only for active downloads, including paused
	OnlyActive bool
	// Pass true to search only for completed downloads
	OnlyCompleted bool
	// Offset of the first entry to return as received from the previous request; use empty
	// string to get the first chunk of results
	Offset string
	// The maximum number of files to be returned
	Limit int32
}

// SearchFileDownloadsRequestTypeID is TL type id of SearchFileDownloadsRequest.
const SearchFileDownloadsRequestTypeID = 0x2a1e0856

// Ensuring interfaces in compile-time for SearchFileDownloadsRequest.
var (
	_ bin.Encoder     = &SearchFileDownloadsRequest{}
	_ bin.Decoder     = &SearchFileDownloadsRequest{}
	_ bin.BareEncoder = &SearchFileDownloadsRequest{}
	_ bin.BareDecoder = &SearchFileDownloadsRequest{}
)

func (s *SearchFileDownloadsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Query == "") {
		return false
	}
	if !(s.OnlyActive == false) {
		return false
	}
	if !(s.OnlyCompleted == false) {
		return false
	}
	if !(s.Offset == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchFileDownloadsRequest) String() string {
	if s == nil {
		return "SearchFileDownloadsRequest(nil)"
	}
	type Alias SearchFileDownloadsRequest
	return fmt.Sprintf("SearchFileDownloadsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchFileDownloadsRequest) TypeID() uint32 {
	return SearchFileDownloadsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchFileDownloadsRequest) TypeName() string {
	return "searchFileDownloads"
}

// TypeInfo returns info about TL type.
func (s *SearchFileDownloadsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchFileDownloads",
		ID:   SearchFileDownloadsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Query",
			SchemaName: "query",
		},
		{
			Name:       "OnlyActive",
			SchemaName: "only_active",
		},
		{
			Name:       "OnlyCompleted",
			SchemaName: "only_completed",
		},
		{
			Name:       "Offset",
			SchemaName: "offset",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchFileDownloadsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchFileDownloads#2a1e0856 as nil")
	}
	b.PutID(SearchFileDownloadsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchFileDownloadsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchFileDownloads#2a1e0856 as nil")
	}
	b.PutString(s.Query)
	b.PutBool(s.OnlyActive)
	b.PutBool(s.OnlyCompleted)
	b.PutString(s.Offset)
	b.PutInt32(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchFileDownloadsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchFileDownloads#2a1e0856 to nil")
	}
	if err := b.ConsumeID(SearchFileDownloadsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchFileDownloadsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchFileDownloads#2a1e0856 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: field query: %w", err)
		}
		s.Query = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: field only_active: %w", err)
		}
		s.OnlyActive = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: field only_completed: %w", err)
		}
		s.OnlyCompleted = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: field offset: %w", err)
		}
		s.Offset = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SearchFileDownloadsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchFileDownloads#2a1e0856 as nil")
	}
	b.ObjStart()
	b.PutID("searchFileDownloads")
	b.Comma()
	b.FieldStart("query")
	b.PutString(s.Query)
	b.Comma()
	b.FieldStart("only_active")
	b.PutBool(s.OnlyActive)
	b.Comma()
	b.FieldStart("only_completed")
	b.PutBool(s.OnlyCompleted)
	b.Comma()
	b.FieldStart("offset")
	b.PutString(s.Offset)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SearchFileDownloadsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode searchFileDownloads#2a1e0856 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("searchFileDownloads"); err != nil {
				return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: %w", err)
			}
		case "query":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: field query: %w", err)
			}
			s.Query = value
		case "only_active":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: field only_active: %w", err)
			}
			s.OnlyActive = value
		case "only_completed":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: field only_completed: %w", err)
			}
			s.OnlyCompleted = value
		case "offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: field offset: %w", err)
			}
			s.Offset = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode searchFileDownloads#2a1e0856: field limit: %w", err)
			}
			s.Limit = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetQuery returns value of Query field.
func (s *SearchFileDownloadsRequest) GetQuery() (value string) {
	if s == nil {
		return
	}
	return s.Query
}

// GetOnlyActive returns value of OnlyActive field.
func (s *SearchFileDownloadsRequest) GetOnlyActive() (value bool) {
	if s == nil {
		return
	}
	return s.OnlyActive
}

// GetOnlyCompleted returns value of OnlyCompleted field.
func (s *SearchFileDownloadsRequest) GetOnlyCompleted() (value bool) {
	if s == nil {
		return
	}
	return s.OnlyCompleted
}

// GetOffset returns value of Offset field.
func (s *SearchFileDownloadsRequest) GetOffset() (value string) {
	if s == nil {
		return
	}
	return s.Offset
}

// GetLimit returns value of Limit field.
func (s *SearchFileDownloadsRequest) GetLimit() (value int32) {
	if s == nil {
		return
	}
	return s.Limit
}

// SearchFileDownloads invokes method searchFileDownloads#2a1e0856 returning error if any.
func (c *Client) SearchFileDownloads(ctx context.Context, request *SearchFileDownloadsRequest) (*FoundFileDownloads, error) {
	var result FoundFileDownloads

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
