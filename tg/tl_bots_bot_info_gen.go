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

// BotsBotInfo represents TL type `bots.botInfo#e8a775b0`.
// Localized information about a bot.
//
// See https://core.telegram.org/constructor/bots.botInfo for reference.
type BotsBotInfo struct {
	// Bot name
	Name string
	// Bot about text
	About string
	// Bot description
	Description string
}

// BotsBotInfoTypeID is TL type id of BotsBotInfo.
const BotsBotInfoTypeID = 0xe8a775b0

// Ensuring interfaces in compile-time for BotsBotInfo.
var (
	_ bin.Encoder     = &BotsBotInfo{}
	_ bin.Decoder     = &BotsBotInfo{}
	_ bin.BareEncoder = &BotsBotInfo{}
	_ bin.BareDecoder = &BotsBotInfo{}
)

func (b *BotsBotInfo) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.Name == "") {
		return false
	}
	if !(b.About == "") {
		return false
	}
	if !(b.Description == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotsBotInfo) String() string {
	if b == nil {
		return "BotsBotInfo(nil)"
	}
	type Alias BotsBotInfo
	return fmt.Sprintf("BotsBotInfo%+v", Alias(*b))
}

// FillFrom fills BotsBotInfo from given interface.
func (b *BotsBotInfo) FillFrom(from interface {
	GetName() (value string)
	GetAbout() (value string)
	GetDescription() (value string)
}) {
	b.Name = from.GetName()
	b.About = from.GetAbout()
	b.Description = from.GetDescription()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotsBotInfo) TypeID() uint32 {
	return BotsBotInfoTypeID
}

// TypeName returns name of type in TL schema.
func (*BotsBotInfo) TypeName() string {
	return "bots.botInfo"
}

// TypeInfo returns info about TL type.
func (b *BotsBotInfo) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "bots.botInfo",
		ID:   BotsBotInfoTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "About",
			SchemaName: "about",
		},
		{
			Name:       "Description",
			SchemaName: "description",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotsBotInfo) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bots.botInfo#e8a775b0 as nil")
	}
	buf.PutID(BotsBotInfoTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotsBotInfo) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode bots.botInfo#e8a775b0 as nil")
	}
	buf.PutString(b.Name)
	buf.PutString(b.About)
	buf.PutString(b.Description)
	return nil
}

// Decode implements bin.Decoder.
func (b *BotsBotInfo) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bots.botInfo#e8a775b0 to nil")
	}
	if err := buf.ConsumeID(BotsBotInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode bots.botInfo#e8a775b0: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotsBotInfo) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode bots.botInfo#e8a775b0 to nil")
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.botInfo#e8a775b0: field name: %w", err)
		}
		b.Name = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.botInfo#e8a775b0: field about: %w", err)
		}
		b.About = value
	}
	{
		value, err := buf.String()
		if err != nil {
			return fmt.Errorf("unable to decode bots.botInfo#e8a775b0: field description: %w", err)
		}
		b.Description = value
	}
	return nil
}

// GetName returns value of Name field.
func (b *BotsBotInfo) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

// GetAbout returns value of About field.
func (b *BotsBotInfo) GetAbout() (value string) {
	if b == nil {
		return
	}
	return b.About
}

// GetDescription returns value of Description field.
func (b *BotsBotInfo) GetDescription() (value string) {
	if b == nil {
		return
	}
	return b.Description
}
