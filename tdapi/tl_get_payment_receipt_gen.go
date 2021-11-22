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

// GetPaymentReceiptRequest represents TL type `getPaymentReceipt#3c6cb956`.
type GetPaymentReceiptRequest struct {
	// Chat identifier of the PaymentSuccessful message
	ChatID int64
	// Message identifier
	MessageID int64
}

// GetPaymentReceiptRequestTypeID is TL type id of GetPaymentReceiptRequest.
const GetPaymentReceiptRequestTypeID = 0x3c6cb956

// Ensuring interfaces in compile-time for GetPaymentReceiptRequest.
var (
	_ bin.Encoder     = &GetPaymentReceiptRequest{}
	_ bin.Decoder     = &GetPaymentReceiptRequest{}
	_ bin.BareEncoder = &GetPaymentReceiptRequest{}
	_ bin.BareDecoder = &GetPaymentReceiptRequest{}
)

func (g *GetPaymentReceiptRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ChatID == 0) {
		return false
	}
	if !(g.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *GetPaymentReceiptRequest) String() string {
	if g == nil {
		return "GetPaymentReceiptRequest(nil)"
	}
	type Alias GetPaymentReceiptRequest
	return fmt.Sprintf("GetPaymentReceiptRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*GetPaymentReceiptRequest) TypeID() uint32 {
	return GetPaymentReceiptRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*GetPaymentReceiptRequest) TypeName() string {
	return "getPaymentReceipt"
}

// TypeInfo returns info about TL type.
func (g *GetPaymentReceiptRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "getPaymentReceipt",
		ID:   GetPaymentReceiptRequestTypeID,
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
			Name:       "MessageID",
			SchemaName: "message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *GetPaymentReceiptRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPaymentReceipt#3c6cb956 as nil")
	}
	b.PutID(GetPaymentReceiptRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *GetPaymentReceiptRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode getPaymentReceipt#3c6cb956 as nil")
	}
	b.PutLong(g.ChatID)
	b.PutLong(g.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (g *GetPaymentReceiptRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPaymentReceipt#3c6cb956 to nil")
	}
	if err := b.ConsumeID(GetPaymentReceiptRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode getPaymentReceipt#3c6cb956: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *GetPaymentReceiptRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode getPaymentReceipt#3c6cb956 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getPaymentReceipt#3c6cb956: field chat_id: %w", err)
		}
		g.ChatID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode getPaymentReceipt#3c6cb956: field message_id: %w", err)
		}
		g.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON encodes g in TDLib API JSON format.
func (g *GetPaymentReceiptRequest) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if g == nil {
		return fmt.Errorf("can't encode getPaymentReceipt#3c6cb956 as nil")
	}
	b.ObjStart()
	b.PutID("getPaymentReceipt")
	b.FieldStart("chat_id")
	b.PutLong(g.ChatID)
	b.FieldStart("message_id")
	b.PutLong(g.MessageID)
	b.ObjEnd()
	return nil
}

// GetChatID returns value of ChatID field.
func (g *GetPaymentReceiptRequest) GetChatID() (value int64) {
	return g.ChatID
}

// GetMessageID returns value of MessageID field.
func (g *GetPaymentReceiptRequest) GetMessageID() (value int64) {
	return g.MessageID
}

// GetPaymentReceipt invokes method getPaymentReceipt#3c6cb956 returning error if any.
func (c *Client) GetPaymentReceipt(ctx context.Context, request *GetPaymentReceiptRequest) (*PaymentReceipt, error) {
	var result PaymentReceipt

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
