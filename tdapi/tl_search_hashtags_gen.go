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

// SearchHashtagsRequest represents TL type `searchHashtags#3e34a571`.
type SearchHashtagsRequest struct {
	// Hashtag prefix to search for
	Prefix string
	// The maximum number of hashtags to be returned
	Limit int32
}

// SearchHashtagsRequestTypeID is TL type id of SearchHashtagsRequest.
const SearchHashtagsRequestTypeID = 0x3e34a571

// Ensuring interfaces in compile-time for SearchHashtagsRequest.
var (
	_ bin.Encoder     = &SearchHashtagsRequest{}
	_ bin.Decoder     = &SearchHashtagsRequest{}
	_ bin.BareEncoder = &SearchHashtagsRequest{}
	_ bin.BareDecoder = &SearchHashtagsRequest{}
)

func (s *SearchHashtagsRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Prefix == "") {
		return false
	}
	if !(s.Limit == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SearchHashtagsRequest) String() string {
	if s == nil {
		return "SearchHashtagsRequest(nil)"
	}
	type Alias SearchHashtagsRequest
	return fmt.Sprintf("SearchHashtagsRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SearchHashtagsRequest) TypeID() uint32 {
	return SearchHashtagsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SearchHashtagsRequest) TypeName() string {
	return "searchHashtags"
}

// TypeInfo returns info about TL type.
func (s *SearchHashtagsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "searchHashtags",
		ID:   SearchHashtagsRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Prefix",
			SchemaName: "prefix",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SearchHashtagsRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchHashtags#3e34a571 as nil")
	}
	b.PutID(SearchHashtagsRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SearchHashtagsRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode searchHashtags#3e34a571 as nil")
	}
	b.PutString(s.Prefix)
	b.PutInt32(s.Limit)
	return nil
}

// Decode implements bin.Decoder.
func (s *SearchHashtagsRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchHashtags#3e34a571 to nil")
	}
	if err := b.ConsumeID(SearchHashtagsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode searchHashtags#3e34a571: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SearchHashtagsRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode searchHashtags#3e34a571 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode searchHashtags#3e34a571: field prefix: %w", err)
		}
		s.Prefix = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode searchHashtags#3e34a571: field limit: %w", err)
		}
		s.Limit = value
	}
	return nil
}

// EncodeTDLibJSON encodes s in TDLib API JSON format.
func (s *SearchHashtagsRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode searchHashtags#3e34a571 as nil")
	}
	b.ObjStart()
	b.PutID("searchHashtags")
	b.FieldStart("prefix")
	b.PutString(s.Prefix)
	b.FieldStart("limit")
	b.PutInt32(s.Limit)
	b.ObjEnd()
	return nil
}

// GetPrefix returns value of Prefix field.
func (s *SearchHashtagsRequest) GetPrefix() (value string) {
	return s.Prefix
}

// GetLimit returns value of Limit field.
func (s *SearchHashtagsRequest) GetLimit() (value int32) {
	return s.Limit
}

// SearchHashtags invokes method searchHashtags#3e34a571 returning error if any.
func (c *Client) SearchHashtags(ctx context.Context, request *SearchHashtagsRequest) (*Hashtags, error) {
	var result Hashtags

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
