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

// AnswerInlineQueryRequest represents TL type `answerInlineQuery#ac7cc5e6`.
type AnswerInlineQueryRequest struct {
	// Identifier of the inline query
	InlineQueryID int64
	// Pass true if results may be cached and returned only for the user that sent the query.
	// By default, results may be returned to any user who sends the same query
	IsPersonal bool
	// Button to be shown above inline query results; pass null if none
	Button InlineQueryResultsButton
	// The results of the query
	Results []InputInlineQueryResultClass
	// Allowed time to cache the results of the query, in seconds
	CacheTime int32
	// Offset for the next inline query; pass an empty string if there are no more results
	NextOffset string
}

// AnswerInlineQueryRequestTypeID is TL type id of AnswerInlineQueryRequest.
const AnswerInlineQueryRequestTypeID = 0xac7cc5e6

// Ensuring interfaces in compile-time for AnswerInlineQueryRequest.
var (
	_ bin.Encoder     = &AnswerInlineQueryRequest{}
	_ bin.Decoder     = &AnswerInlineQueryRequest{}
	_ bin.BareEncoder = &AnswerInlineQueryRequest{}
	_ bin.BareDecoder = &AnswerInlineQueryRequest{}
)

func (a *AnswerInlineQueryRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.InlineQueryID == 0) {
		return false
	}
	if !(a.IsPersonal == false) {
		return false
	}
	if !(a.Button.Zero()) {
		return false
	}
	if !(a.Results == nil) {
		return false
	}
	if !(a.CacheTime == 0) {
		return false
	}
	if !(a.NextOffset == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AnswerInlineQueryRequest) String() string {
	if a == nil {
		return "AnswerInlineQueryRequest(nil)"
	}
	type Alias AnswerInlineQueryRequest
	return fmt.Sprintf("AnswerInlineQueryRequest%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AnswerInlineQueryRequest) TypeID() uint32 {
	return AnswerInlineQueryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AnswerInlineQueryRequest) TypeName() string {
	return "answerInlineQuery"
}

// TypeInfo returns info about TL type.
func (a *AnswerInlineQueryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "answerInlineQuery",
		ID:   AnswerInlineQueryRequestTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "InlineQueryID",
			SchemaName: "inline_query_id",
		},
		{
			Name:       "IsPersonal",
			SchemaName: "is_personal",
		},
		{
			Name:       "Button",
			SchemaName: "button",
		},
		{
			Name:       "Results",
			SchemaName: "results",
		},
		{
			Name:       "CacheTime",
			SchemaName: "cache_time",
		},
		{
			Name:       "NextOffset",
			SchemaName: "next_offset",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AnswerInlineQueryRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode answerInlineQuery#ac7cc5e6 as nil")
	}
	b.PutID(AnswerInlineQueryRequestTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AnswerInlineQueryRequest) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode answerInlineQuery#ac7cc5e6 as nil")
	}
	b.PutLong(a.InlineQueryID)
	b.PutBool(a.IsPersonal)
	if err := a.Button.Encode(b); err != nil {
		return fmt.Errorf("unable to encode answerInlineQuery#ac7cc5e6: field button: %w", err)
	}
	b.PutInt(len(a.Results))
	for idx, v := range a.Results {
		if v == nil {
			return fmt.Errorf("unable to encode answerInlineQuery#ac7cc5e6: field results element with index %d is nil", idx)
		}
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare answerInlineQuery#ac7cc5e6: field results element with index %d: %w", idx, err)
		}
	}
	b.PutInt32(a.CacheTime)
	b.PutString(a.NextOffset)
	return nil
}

// Decode implements bin.Decoder.
func (a *AnswerInlineQueryRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode answerInlineQuery#ac7cc5e6 to nil")
	}
	if err := b.ConsumeID(AnswerInlineQueryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AnswerInlineQueryRequest) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode answerInlineQuery#ac7cc5e6 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field inline_query_id: %w", err)
		}
		a.InlineQueryID = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field is_personal: %w", err)
		}
		a.IsPersonal = value
	}
	{
		if err := a.Button.Decode(b); err != nil {
			return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field button: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field results: %w", err)
		}

		if headerLen > 0 {
			a.Results = make([]InputInlineQueryResultClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputInlineQueryResult(b)
			if err != nil {
				return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field results: %w", err)
			}
			a.Results = append(a.Results, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field cache_time: %w", err)
		}
		a.CacheTime = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field next_offset: %w", err)
		}
		a.NextOffset = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (a *AnswerInlineQueryRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if a == nil {
		return fmt.Errorf("can't encode answerInlineQuery#ac7cc5e6 as nil")
	}
	b.ObjStart()
	b.PutID("answerInlineQuery")
	b.Comma()
	b.FieldStart("inline_query_id")
	b.PutLong(a.InlineQueryID)
	b.Comma()
	b.FieldStart("is_personal")
	b.PutBool(a.IsPersonal)
	b.Comma()
	b.FieldStart("button")
	if err := a.Button.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode answerInlineQuery#ac7cc5e6: field button: %w", err)
	}
	b.Comma()
	b.FieldStart("results")
	b.ArrStart()
	for idx, v := range a.Results {
		if v == nil {
			return fmt.Errorf("unable to encode answerInlineQuery#ac7cc5e6: field results element with index %d is nil", idx)
		}
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode answerInlineQuery#ac7cc5e6: field results element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("cache_time")
	b.PutInt32(a.CacheTime)
	b.Comma()
	b.FieldStart("next_offset")
	b.PutString(a.NextOffset)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (a *AnswerInlineQueryRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if a == nil {
		return fmt.Errorf("can't decode answerInlineQuery#ac7cc5e6 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("answerInlineQuery"); err != nil {
				return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: %w", err)
			}
		case "inline_query_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field inline_query_id: %w", err)
			}
			a.InlineQueryID = value
		case "is_personal":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field is_personal: %w", err)
			}
			a.IsPersonal = value
		case "button":
			if err := a.Button.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field button: %w", err)
			}
		case "results":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := DecodeTDLibJSONInputInlineQueryResult(b)
				if err != nil {
					return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field results: %w", err)
				}
				a.Results = append(a.Results, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field results: %w", err)
			}
		case "cache_time":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field cache_time: %w", err)
			}
			a.CacheTime = value
		case "next_offset":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode answerInlineQuery#ac7cc5e6: field next_offset: %w", err)
			}
			a.NextOffset = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInlineQueryID returns value of InlineQueryID field.
func (a *AnswerInlineQueryRequest) GetInlineQueryID() (value int64) {
	if a == nil {
		return
	}
	return a.InlineQueryID
}

// GetIsPersonal returns value of IsPersonal field.
func (a *AnswerInlineQueryRequest) GetIsPersonal() (value bool) {
	if a == nil {
		return
	}
	return a.IsPersonal
}

// GetButton returns value of Button field.
func (a *AnswerInlineQueryRequest) GetButton() (value InlineQueryResultsButton) {
	if a == nil {
		return
	}
	return a.Button
}

// GetResults returns value of Results field.
func (a *AnswerInlineQueryRequest) GetResults() (value []InputInlineQueryResultClass) {
	if a == nil {
		return
	}
	return a.Results
}

// GetCacheTime returns value of CacheTime field.
func (a *AnswerInlineQueryRequest) GetCacheTime() (value int32) {
	if a == nil {
		return
	}
	return a.CacheTime
}

// GetNextOffset returns value of NextOffset field.
func (a *AnswerInlineQueryRequest) GetNextOffset() (value string) {
	if a == nil {
		return
	}
	return a.NextOffset
}

// AnswerInlineQuery invokes method answerInlineQuery#ac7cc5e6 returning error if any.
func (c *Client) AnswerInlineQuery(ctx context.Context, request *AnswerInlineQueryRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
