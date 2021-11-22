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

// HelpGetRecentMeURLsRequest represents TL type `help.getRecentMeUrls#3dc0f114`.
// Get recently used t.me links
//
// See https://core.telegram.org/method/help.getRecentMeUrls for reference.
type HelpGetRecentMeURLsRequest struct {
	// Referer
	Referer string
}

// HelpGetRecentMeURLsRequestTypeID is TL type id of HelpGetRecentMeURLsRequest.
const HelpGetRecentMeURLsRequestTypeID = 0x3dc0f114

// Ensuring interfaces in compile-time for HelpGetRecentMeURLsRequest.
var (
	_ bin.Encoder     = &HelpGetRecentMeURLsRequest{}
	_ bin.Decoder     = &HelpGetRecentMeURLsRequest{}
	_ bin.BareEncoder = &HelpGetRecentMeURLsRequest{}
	_ bin.BareDecoder = &HelpGetRecentMeURLsRequest{}
)

func (g *HelpGetRecentMeURLsRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.Referer == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetRecentMeURLsRequest) String() string {
	if g == nil {
		return "HelpGetRecentMeURLsRequest(nil)"
	}
	type Alias HelpGetRecentMeURLsRequest
	return fmt.Sprintf("HelpGetRecentMeURLsRequest%+v", Alias(*g))
}

// FillFrom fills HelpGetRecentMeURLsRequest from given interface.
func (g *HelpGetRecentMeURLsRequest) FillFrom(from interface {
	GetReferer() (value string)
}) {
	g.Referer = from.GetReferer()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetRecentMeURLsRequest) TypeID() uint32 {
	return HelpGetRecentMeURLsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetRecentMeURLsRequest) TypeName() string {
	return "help.getRecentMeUrls"
}

// TypeInfo returns info about TL type.
func (g *HelpGetRecentMeURLsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getRecentMeUrls",
		ID:   HelpGetRecentMeURLsRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Referer",
			SchemaName: "referer",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetRecentMeURLsRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getRecentMeUrls#3dc0f114 as nil")
	}
	b.PutID(HelpGetRecentMeURLsRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetRecentMeURLsRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getRecentMeUrls#3dc0f114 as nil")
	}
	b.PutString(g.Referer)
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetRecentMeURLsRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getRecentMeUrls#3dc0f114 to nil")
	}
	if err := b.ConsumeID(HelpGetRecentMeURLsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getRecentMeUrls#3dc0f114: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetRecentMeURLsRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getRecentMeUrls#3dc0f114 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.getRecentMeUrls#3dc0f114: field referer: %w", err)
		}
		g.Referer = value
	}
	return nil
}

// GetReferer returns value of Referer field.
func (g *HelpGetRecentMeURLsRequest) GetReferer() (value string) {
	return g.Referer
}

// HelpGetRecentMeURLs invokes method help.getRecentMeUrls#3dc0f114 returning error if any.
// Get recently used t.me links
//
// See https://core.telegram.org/method/help.getRecentMeUrls for reference.
func (c *Client) HelpGetRecentMeURLs(ctx context.Context, referer string) (*HelpRecentMeURLs, error) {
	var result HelpRecentMeURLs

	request := &HelpGetRecentMeURLsRequest{
		Referer: referer,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
