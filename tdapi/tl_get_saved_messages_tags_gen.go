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

// GetSavedMessagesTagsRequest represents TL type `getSavedMessagesTags#8cd667a9`.
type GetSavedMessagesTagsRequest struct {
	// Identifier of Saved Messages topic which tags will be returned; pass 0 to get all
	// Saved Messages tags
	SavedMessagesTopicID int64
}

// GetSavedMessagesTagsRequestTypeID is TL type id of GetSavedMessagesTagsRequest.
const GetSavedMessagesTagsRequestTypeID = 0x8cd667a9

// Ensuring interfaces in compile-time for GetSavedMessagesTagsRequest.
var (
	_ bin.Encoder     = &GetSavedMessagesTagsRequest{}
	_ bin.Decoder     = &GetSavedMessagesTagsRequest{}
	_ bin.BareEncoder = &GetSavedMessagesTagsRequest{}
	_ bin.BareDecoder = &GetSavedMessagesTagsRequest{}
)

func (g *GetSavedMessagesTagsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.SavedMessagesTopicID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetSavedMessagesTagsRequest) String() string {
	if g == nil {
		return "GetSavedMessagesTagsRequest(nil)"
	}
	type Alias GetSavedMessagesTagsRequest
	return fmt.Sprintf("GetSavedMessagesTagsRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetSavedMessagesTagsRequest) TypeID() uint32 {
	return GetSavedMessagesTagsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetSavedMessagesTagsRequest) TypeName() string {
	return "getSavedMessagesTags"
}

// TypeInfo returns info about TL type.
func (g *GetSavedMessagesTagsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getSavedMessagesTags",
		ID:   GetSavedMessagesTagsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "SavedMessagesTopicID",
			SchemaName: "saved_messages_topic_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetSavedMessagesTagsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedMessagesTags#8cd667a9 as nil")
	}
	b.PutID(GetSavedMessagesTagsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetSavedMessagesTagsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedMessagesTags#8cd667a9 as nil")
	}
	b.PutInt53(g.SavedMessagesTopicID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetSavedMessagesTagsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedMessagesTags#8cd667a9 to nil")
	}
	if err := b.ConsumeID(GetSavedMessagesTagsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getSavedMessagesTags#8cd667a9: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetSavedMessagesTagsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedMessagesTags#8cd667a9 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode getSavedMessagesTags#8cd667a9: field saved_messages_topic_id: %w", err)
		}
		g.SavedMessagesTopicID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (g *GetSavedMessagesTagsRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getSavedMessagesTags#8cd667a9 as nil")
	}
	b.ObjStart()
	b.PutID("getSavedMessagesTags")
	b.Comma()
	b.FieldStart("saved_messages_topic_id")
	b.PutInt53(g.SavedMessagesTopicID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (g *GetSavedMessagesTagsRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if g == nil {
		return fmt.Errorf("can't decode getSavedMessagesTags#8cd667a9 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("getSavedMessagesTags"); err != nil {
				return fmt.Errorf("unable to decode getSavedMessagesTags#8cd667a9: %w", err)
			}
		case "saved_messages_topic_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode getSavedMessagesTags#8cd667a9: field saved_messages_topic_id: %w", err)
			}
			g.SavedMessagesTopicID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSavedMessagesTopicID returns value of SavedMessagesTopicID field.
func (g *GetSavedMessagesTagsRequest) GetSavedMessagesTopicID() (value int64) {
	if g == nil {
		return
	}
	return g.SavedMessagesTopicID
}

// GetSavedMessagesTags invokes method getSavedMessagesTags#8cd667a9 returning error if any.
func (c *Client) GetSavedMessagesTags(ctx context.Context, savedmessagestopicid int64) (*SavedMessagesTags, error) {
	var result SavedMessagesTags

	request := &GetSavedMessagesTagsRequest{
		SavedMessagesTopicID: savedmessagestopicid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
