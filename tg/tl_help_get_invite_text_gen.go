// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// HelpGetInviteTextRequest represents TL type `help.getInviteText#4d392343`.
// Returns localized text of a text message with an invitation.
//
// See https://core.telegram.org/method/help.getInviteText for reference.
type HelpGetInviteTextRequest struct {
}

// HelpGetInviteTextRequestTypeID is TL type id of HelpGetInviteTextRequest.
const HelpGetInviteTextRequestTypeID = 0x4d392343

// Ensuring interfaces in compile-time for HelpGetInviteTextRequest.
var (
	_ bin.Encoder     = &HelpGetInviteTextRequest{}
	_ bin.Decoder     = &HelpGetInviteTextRequest{}
	_ bin.BareEncoder = &HelpGetInviteTextRequest{}
	_ bin.BareDecoder = &HelpGetInviteTextRequest{}
)

func (g *HelpGetInviteTextRequest) Zero() bool {
	if g == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (g *HelpGetInviteTextRequest) String() string {
	if g == nil {
		return "HelpGetInviteTextRequest(nil)"
	}
	type Alias HelpGetInviteTextRequest
	return fmt.Sprintf("HelpGetInviteTextRequest%+v", Alias(*g))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*HelpGetInviteTextRequest) TypeID() uint32 {
	return HelpGetInviteTextRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*HelpGetInviteTextRequest) TypeName() string {
	return "help.getInviteText"
}

// TypeInfo returns info about TL type.
func (g *HelpGetInviteTextRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "help.getInviteText",
		ID:   HelpGetInviteTextRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (g *HelpGetInviteTextRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getInviteText#4d392343 as nil")
	}
	b.PutID(HelpGetInviteTextRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *HelpGetInviteTextRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode help.getInviteText#4d392343 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *HelpGetInviteTextRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getInviteText#4d392343 to nil")
	}
	if err := b.ConsumeID(HelpGetInviteTextRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode help.getInviteText#4d392343: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *HelpGetInviteTextRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode help.getInviteText#4d392343 to nil")
	}
	return nil
}

// HelpGetInviteText invokes method help.getInviteText#4d392343 returning error if any.
// Returns localized text of a text message with an invitation.
//
// See https://core.telegram.org/method/help.getInviteText for reference.
func (c *Client) HelpGetInviteText(ctx context.Context) (*HelpInviteText, error) {
	var result HelpInviteText

	request := &HelpGetInviteTextRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
