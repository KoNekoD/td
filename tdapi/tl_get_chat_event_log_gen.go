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

// GetChatEventLogRequest represents TL type `getChatEventLog#f47b0a9b`.
type GetChatEventLogRequest struct {
	// Chat identifier
	ChatID int64
	// Search query by which to filter events
	Query string
	// Identifier of an event from which to return results. Use 0 to get results from the
	// latest events
	FromEventID int64
	// The maximum number of events to return; up to 100
	Limit int32
	// The types of events to return; pass null to get chat events of all types
	Filters ChatEventLogFilters
	// User identifiers by which to filter events. By default, events relating to all users
	// will be returned
	UserIDs []int64
}

// GetChatEventLogRequestTypeID is TL type id of GetChatEventLogRequest.
const GetChatEventLogRequestTypeID = 0xf47b0a9b

// Ensuring interfaces in compile-time for GetChatEventLogRequest.
var (
	_ bin.Encoder     = &GetChatEventLogRequest{}
	_ bin.Decoder     = &GetChatEventLogRequest{}
	_ bin.BareEncoder = &GetChatEventLogRequest{}
	_ bin.BareDecoder = &GetChatEventLogRequest{}
)

func (g *GetChatEventLogRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.Query == "") {
		return false
	}
	if !(g.FromEventID == 0) {
		return false
	}
	if !(g.Limit == 0) {
		return false
	}
	if !(g.Filters.Zero()) {
		return false
	}
	if !(g.UserIDs == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetChatEventLogRequest) String() string {
	if g == nil {
		return "GetChatEventLogRequest(nil)"
	}
	type Alias GetChatEventLogRequest
	return fmt.Sprintf("GetChatEventLogRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetChatEventLogRequest) TypeID() uint32 {
	return GetChatEventLogRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetChatEventLogRequest) TypeName() string {
	return "getChatEventLog"
}

// TypeInfo returns info about TL type.
func (g *GetChatEventLogRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getChatEventLog",
		ID:   GetChatEventLogRequestTypeID,
	}
	if g == nil {
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
			Name:       "FromEventID",
			SchemaName: "from_event_id",
		},
		{
			Name:       "Limit",
			SchemaName: "limit",
		},
		{
			Name:       "Filters",
			SchemaName: "filters",
		},
		{
			Name:       "UserIDs",
			SchemaName: "user_ids",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetChatEventLogRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatEventLog#f47b0a9b as nil")
	}
	b.PutID(GetChatEventLogRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetChatEventLogRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatEventLog#f47b0a9b as nil")
	}
	b.PutInt53(g.ChatID)
	b.PutString(g.Query)
	b.PutLong(g.FromEventID)
	b.PutInt32(g.Limit)
	if err := g.Filters.Encode(b); err != nil {
		return fmt.Errorf("unable to encode getChatEventLog#f47b0a9b: field filters: %w", err)
	}
	b.PutInt(len(g.UserIDs))
	for _, v := range g.UserIDs {
		b.PutInt53(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *GetChatEventLogRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatEventLog#f47b0a9b to nil")
	}
	if err := b.ConsumeID(GetChatEventLogRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetChatEventLogRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatEventLog#f47b0a9b to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field query: %w", err)
		}
		g.Query = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field from_event_id: %w", err)
		}
		g.FromEventID = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field limit: %w", err)
		}
		g.Limit = value
	}
	{
		if err := g.Filters.Decode(b); err != nil {
			return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field filters: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field user_ids: %w", err)
		}

		if headerLen > 0 {
			g.UserIDs = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field user_ids: %w", err)
			}
			g.UserIDs = append(g.UserIDs, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetChatEventLogRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getChatEventLog#f47b0a9b as nil")
	}
	b.ObjStart()
	b.PutID("getChatEventLog")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(g.ChatID)
	b.Comma()
	b.FieldStart("query")
	b.PutString(g.Query)
	b.Comma()
	b.FieldStart("from_event_id")
	b.PutLong(g.FromEventID)
	b.Comma()
	b.FieldStart("limit")
	b.PutInt32(g.Limit)
	b.Comma()
	b.FieldStart("filters")
	if err := g.Filters.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode getChatEventLog#f47b0a9b: field filters: %w", err)
	}
	b.Comma()
	b.FieldStart("user_ids")
	b.ArrStart()
	for _, v := range g.UserIDs {
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
func (g *GetChatEventLogRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getChatEventLog#f47b0a9b to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getChatEventLog"); err != nil {
				return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field chat_id: %w", err)
			}
			g.ChatID = value
		case "query":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field query: %w", err)
			}
			g.Query = value
		case "from_event_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field from_event_id: %w", err)
			}
			g.FromEventID = value
		case "limit":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field limit: %w", err)
			}
			g.Limit = value
		case "filters":
			if err := g.Filters.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field filters: %w", err)
			}
		case "user_ids":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.Int53()
				if err != nil {
					return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field user_ids: %w", err)
				}
				g.UserIDs = append(g.UserIDs, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode getChatEventLog#f47b0a9b: field user_ids: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (g *GetChatEventLogRequest) GetChatID() (value int64) {
	if g == nil {
		return
	}
	return g.ChatID
}

// GetQuery returns value of Query field.
func (g *GetChatEventLogRequest) GetQuery() (value string) {
	if g == nil {
		return
	}
	return g.Query
}

// GetFromEventID returns value of FromEventID field.
func (g *GetChatEventLogRequest) GetFromEventID() (value int64) {
	if g == nil {
		return
	}
	return g.FromEventID
}

// GetLimit returns value of Limit field.
func (g *GetChatEventLogRequest) GetLimit() (value int32) {
	if g == nil {
		return
	}
	return g.Limit
}

// GetFilters returns value of Filters field.
func (g *GetChatEventLogRequest) GetFilters() (value ChatEventLogFilters) {
	if g == nil {
		return
	}
	return g.Filters
}

// GetUserIDs returns value of UserIDs field.
func (g *GetChatEventLogRequest) GetUserIDs() (value []int64) {
	if g == nil {
		return
	}
	return g.UserIDs
}

// GetChatEventLog invokes method getChatEventLog#f47b0a9b returning error if any.
func (c *Client) GetChatEventLog(ctx context.Context, request *GetChatEventLogRequest) (*ChatEvents, error) {
	var result ChatEvents

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
