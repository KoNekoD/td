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

// BotCommandScopeDefault represents TL type `botCommandScopeDefault#2f6cb2ab`.
type BotCommandScopeDefault struct {
}

// BotCommandScopeDefaultTypeID is TL type id of BotCommandScopeDefault.
const BotCommandScopeDefaultTypeID = 0x2f6cb2ab

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeDefault) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeDefault.
var (
	_ bin.Encoder     = &BotCommandScopeDefault{}
	_ bin.Decoder     = &BotCommandScopeDefault{}
	_ bin.BareEncoder = &BotCommandScopeDefault{}
	_ bin.BareDecoder = &BotCommandScopeDefault{}

	_ BotCommandScopeClass = &BotCommandScopeDefault{}
)

func (b *BotCommandScopeDefault) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeDefault) String() string {
	if b == nil {
		return "BotCommandScopeDefault(nil)"
	}
	type Alias BotCommandScopeDefault
	return fmt.Sprintf("BotCommandScopeDefault%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeDefault) TypeID() uint32 {
	return BotCommandScopeDefaultTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeDefault) TypeName() string {
	return "botCommandScopeDefault"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeDefault) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeDefault",
		ID:   BotCommandScopeDefaultTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeDefault) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeDefault#2f6cb2ab as nil")
	}
	buf.PutID(BotCommandScopeDefaultTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeDefault) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeDefault#2f6cb2ab as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeDefault) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeDefault#2f6cb2ab to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeDefaultTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeDefault#2f6cb2ab: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeDefault) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeDefault#2f6cb2ab to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotCommandScopeDefault) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeDefault#2f6cb2ab as nil")
	}
	buf.ObjStart()
	buf.PutID("botCommandScopeDefault")
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotCommandScopeDefault) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeDefault#2f6cb2ab to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botCommandScopeDefault"); err != nil {
				return fmt.Errorf("unable to decode botCommandScopeDefault#2f6cb2ab: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// BotCommandScopeAllPrivateChats represents TL type `botCommandScopeAllPrivateChats#eb716739`.
type BotCommandScopeAllPrivateChats struct {
}

// BotCommandScopeAllPrivateChatsTypeID is TL type id of BotCommandScopeAllPrivateChats.
const BotCommandScopeAllPrivateChatsTypeID = 0xeb716739

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeAllPrivateChats) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeAllPrivateChats.
var (
	_ bin.Encoder     = &BotCommandScopeAllPrivateChats{}
	_ bin.Decoder     = &BotCommandScopeAllPrivateChats{}
	_ bin.BareEncoder = &BotCommandScopeAllPrivateChats{}
	_ bin.BareDecoder = &BotCommandScopeAllPrivateChats{}

	_ BotCommandScopeClass = &BotCommandScopeAllPrivateChats{}
)

func (b *BotCommandScopeAllPrivateChats) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeAllPrivateChats) String() string {
	if b == nil {
		return "BotCommandScopeAllPrivateChats(nil)"
	}
	type Alias BotCommandScopeAllPrivateChats
	return fmt.Sprintf("BotCommandScopeAllPrivateChats%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeAllPrivateChats) TypeID() uint32 {
	return BotCommandScopeAllPrivateChatsTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeAllPrivateChats) TypeName() string {
	return "botCommandScopeAllPrivateChats"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeAllPrivateChats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeAllPrivateChats",
		ID:   BotCommandScopeAllPrivateChatsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeAllPrivateChats) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeAllPrivateChats#eb716739 as nil")
	}
	buf.PutID(BotCommandScopeAllPrivateChatsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeAllPrivateChats) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeAllPrivateChats#eb716739 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeAllPrivateChats) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeAllPrivateChats#eb716739 to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeAllPrivateChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeAllPrivateChats#eb716739: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeAllPrivateChats) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeAllPrivateChats#eb716739 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotCommandScopeAllPrivateChats) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeAllPrivateChats#eb716739 as nil")
	}
	buf.ObjStart()
	buf.PutID("botCommandScopeAllPrivateChats")
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotCommandScopeAllPrivateChats) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeAllPrivateChats#eb716739 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botCommandScopeAllPrivateChats"); err != nil {
				return fmt.Errorf("unable to decode botCommandScopeAllPrivateChats#eb716739: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// BotCommandScopeAllGroupChats represents TL type `botCommandScopeAllGroupChats#c585c85e`.
type BotCommandScopeAllGroupChats struct {
}

// BotCommandScopeAllGroupChatsTypeID is TL type id of BotCommandScopeAllGroupChats.
const BotCommandScopeAllGroupChatsTypeID = 0xc585c85e

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeAllGroupChats) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeAllGroupChats.
var (
	_ bin.Encoder     = &BotCommandScopeAllGroupChats{}
	_ bin.Decoder     = &BotCommandScopeAllGroupChats{}
	_ bin.BareEncoder = &BotCommandScopeAllGroupChats{}
	_ bin.BareDecoder = &BotCommandScopeAllGroupChats{}

	_ BotCommandScopeClass = &BotCommandScopeAllGroupChats{}
)

func (b *BotCommandScopeAllGroupChats) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeAllGroupChats) String() string {
	if b == nil {
		return "BotCommandScopeAllGroupChats(nil)"
	}
	type Alias BotCommandScopeAllGroupChats
	return fmt.Sprintf("BotCommandScopeAllGroupChats%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeAllGroupChats) TypeID() uint32 {
	return BotCommandScopeAllGroupChatsTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeAllGroupChats) TypeName() string {
	return "botCommandScopeAllGroupChats"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeAllGroupChats) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeAllGroupChats",
		ID:   BotCommandScopeAllGroupChatsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeAllGroupChats) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeAllGroupChats#c585c85e as nil")
	}
	buf.PutID(BotCommandScopeAllGroupChatsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeAllGroupChats) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeAllGroupChats#c585c85e as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeAllGroupChats) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeAllGroupChats#c585c85e to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeAllGroupChatsTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeAllGroupChats#c585c85e: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeAllGroupChats) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeAllGroupChats#c585c85e to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotCommandScopeAllGroupChats) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeAllGroupChats#c585c85e as nil")
	}
	buf.ObjStart()
	buf.PutID("botCommandScopeAllGroupChats")
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotCommandScopeAllGroupChats) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeAllGroupChats#c585c85e to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botCommandScopeAllGroupChats"); err != nil {
				return fmt.Errorf("unable to decode botCommandScopeAllGroupChats#c585c85e: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// BotCommandScopeAllChatAdministrators represents TL type `botCommandScopeAllChatAdministrators#771c1551`.
type BotCommandScopeAllChatAdministrators struct {
}

// BotCommandScopeAllChatAdministratorsTypeID is TL type id of BotCommandScopeAllChatAdministrators.
const BotCommandScopeAllChatAdministratorsTypeID = 0x771c1551

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeAllChatAdministrators) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeAllChatAdministrators.
var (
	_ bin.Encoder     = &BotCommandScopeAllChatAdministrators{}
	_ bin.Decoder     = &BotCommandScopeAllChatAdministrators{}
	_ bin.BareEncoder = &BotCommandScopeAllChatAdministrators{}
	_ bin.BareDecoder = &BotCommandScopeAllChatAdministrators{}

	_ BotCommandScopeClass = &BotCommandScopeAllChatAdministrators{}
)

func (b *BotCommandScopeAllChatAdministrators) Zero() bool {
	if b == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeAllChatAdministrators) String() string {
	if b == nil {
		return "BotCommandScopeAllChatAdministrators(nil)"
	}
	type Alias BotCommandScopeAllChatAdministrators
	return fmt.Sprintf("BotCommandScopeAllChatAdministrators%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeAllChatAdministrators) TypeID() uint32 {
	return BotCommandScopeAllChatAdministratorsTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeAllChatAdministrators) TypeName() string {
	return "botCommandScopeAllChatAdministrators"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeAllChatAdministrators) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeAllChatAdministrators",
		ID:   BotCommandScopeAllChatAdministratorsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeAllChatAdministrators) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeAllChatAdministrators#771c1551 as nil")
	}
	buf.PutID(BotCommandScopeAllChatAdministratorsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeAllChatAdministrators) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeAllChatAdministrators#771c1551 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeAllChatAdministrators) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeAllChatAdministrators#771c1551 to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeAllChatAdministratorsTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeAllChatAdministrators#771c1551: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeAllChatAdministrators) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeAllChatAdministrators#771c1551 to nil")
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotCommandScopeAllChatAdministrators) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeAllChatAdministrators#771c1551 as nil")
	}
	buf.ObjStart()
	buf.PutID("botCommandScopeAllChatAdministrators")
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotCommandScopeAllChatAdministrators) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeAllChatAdministrators#771c1551 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botCommandScopeAllChatAdministrators"); err != nil {
				return fmt.Errorf("unable to decode botCommandScopeAllChatAdministrators#771c1551: %w", err)
			}
		default:
			return buf.Skip()
		}
		return nil
	})
}

// BotCommandScopeChat represents TL type `botCommandScopeChat#e65b22a5`.
type BotCommandScopeChat struct {
	// Chat identifier
	ChatID int64
}

// BotCommandScopeChatTypeID is TL type id of BotCommandScopeChat.
const BotCommandScopeChatTypeID = 0xe65b22a5

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeChat) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeChat.
var (
	_ bin.Encoder     = &BotCommandScopeChat{}
	_ bin.Decoder     = &BotCommandScopeChat{}
	_ bin.BareEncoder = &BotCommandScopeChat{}
	_ bin.BareDecoder = &BotCommandScopeChat{}

	_ BotCommandScopeClass = &BotCommandScopeChat{}
)

func (b *BotCommandScopeChat) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeChat) String() string {
	if b == nil {
		return "BotCommandScopeChat(nil)"
	}
	type Alias BotCommandScopeChat
	return fmt.Sprintf("BotCommandScopeChat%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeChat) TypeID() uint32 {
	return BotCommandScopeChatTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeChat) TypeName() string {
	return "botCommandScopeChat"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeChat) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeChat",
		ID:   BotCommandScopeChatTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeChat) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChat#e65b22a5 as nil")
	}
	buf.PutID(BotCommandScopeChatTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeChat) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChat#e65b22a5 as nil")
	}
	buf.PutInt53(b.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeChat) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChat#e65b22a5 to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeChatTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeChat#e65b22a5: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeChat) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChat#e65b22a5 to nil")
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode botCommandScopeChat#e65b22a5: field chat_id: %w", err)
		}
		b.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotCommandScopeChat) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChat#e65b22a5 as nil")
	}
	buf.ObjStart()
	buf.PutID("botCommandScopeChat")
	buf.Comma()
	buf.FieldStart("chat_id")
	buf.PutInt53(b.ChatID)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotCommandScopeChat) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChat#e65b22a5 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botCommandScopeChat"); err != nil {
				return fmt.Errorf("unable to decode botCommandScopeChat#e65b22a5: %w", err)
			}
		case "chat_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode botCommandScopeChat#e65b22a5: field chat_id: %w", err)
			}
			b.ChatID = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (b *BotCommandScopeChat) GetChatID() (value int64) {
	if b == nil {
		return
	}
	return b.ChatID
}

// BotCommandScopeChatAdministrators represents TL type `botCommandScopeChatAdministrators#42bcfe4e`.
type BotCommandScopeChatAdministrators struct {
	// Chat identifier
	ChatID int64
}

// BotCommandScopeChatAdministratorsTypeID is TL type id of BotCommandScopeChatAdministrators.
const BotCommandScopeChatAdministratorsTypeID = 0x42bcfe4e

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeChatAdministrators) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeChatAdministrators.
var (
	_ bin.Encoder     = &BotCommandScopeChatAdministrators{}
	_ bin.Decoder     = &BotCommandScopeChatAdministrators{}
	_ bin.BareEncoder = &BotCommandScopeChatAdministrators{}
	_ bin.BareDecoder = &BotCommandScopeChatAdministrators{}

	_ BotCommandScopeClass = &BotCommandScopeChatAdministrators{}
)

func (b *BotCommandScopeChatAdministrators) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ChatID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeChatAdministrators) String() string {
	if b == nil {
		return "BotCommandScopeChatAdministrators(nil)"
	}
	type Alias BotCommandScopeChatAdministrators
	return fmt.Sprintf("BotCommandScopeChatAdministrators%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeChatAdministrators) TypeID() uint32 {
	return BotCommandScopeChatAdministratorsTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeChatAdministrators) TypeName() string {
	return "botCommandScopeChatAdministrators"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeChatAdministrators) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeChatAdministrators",
		ID:   BotCommandScopeChatAdministratorsTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeChatAdministrators) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChatAdministrators#42bcfe4e as nil")
	}
	buf.PutID(BotCommandScopeChatAdministratorsTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeChatAdministrators) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChatAdministrators#42bcfe4e as nil")
	}
	buf.PutInt53(b.ChatID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeChatAdministrators) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChatAdministrators#42bcfe4e to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeChatAdministratorsTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeChatAdministrators#42bcfe4e: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeChatAdministrators) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChatAdministrators#42bcfe4e to nil")
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode botCommandScopeChatAdministrators#42bcfe4e: field chat_id: %w", err)
		}
		b.ChatID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotCommandScopeChatAdministrators) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChatAdministrators#42bcfe4e as nil")
	}
	buf.ObjStart()
	buf.PutID("botCommandScopeChatAdministrators")
	buf.Comma()
	buf.FieldStart("chat_id")
	buf.PutInt53(b.ChatID)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotCommandScopeChatAdministrators) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChatAdministrators#42bcfe4e to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botCommandScopeChatAdministrators"); err != nil {
				return fmt.Errorf("unable to decode botCommandScopeChatAdministrators#42bcfe4e: %w", err)
			}
		case "chat_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode botCommandScopeChatAdministrators#42bcfe4e: field chat_id: %w", err)
			}
			b.ChatID = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (b *BotCommandScopeChatAdministrators) GetChatID() (value int64) {
	if b == nil {
		return
	}
	return b.ChatID
}

// BotCommandScopeChatMember represents TL type `botCommandScopeChatMember#f36696f2`.
type BotCommandScopeChatMember struct {
	// Chat identifier
	ChatID int64
	// User identifier
	UserID int64
}

// BotCommandScopeChatMemberTypeID is TL type id of BotCommandScopeChatMember.
const BotCommandScopeChatMemberTypeID = 0xf36696f2

// construct implements constructor of BotCommandScopeClass.
func (b BotCommandScopeChatMember) construct() BotCommandScopeClass { return &b }

// Ensuring interfaces in compile-time for BotCommandScopeChatMember.
var (
	_ bin.Encoder     = &BotCommandScopeChatMember{}
	_ bin.Decoder     = &BotCommandScopeChatMember{}
	_ bin.BareEncoder = &BotCommandScopeChatMember{}
	_ bin.BareDecoder = &BotCommandScopeChatMember{}

	_ BotCommandScopeClass = &BotCommandScopeChatMember{}
)

func (b *BotCommandScopeChatMember) Zero() bool {
	if b == nil {
		return true
	}
	if !(b.ChatID == 0) {
		return false
	}
	if !(b.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (b *BotCommandScopeChatMember) String() string {
	if b == nil {
		return "BotCommandScopeChatMember(nil)"
	}
	type Alias BotCommandScopeChatMember
	return fmt.Sprintf("BotCommandScopeChatMember%+v", Alias(*b))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*BotCommandScopeChatMember) TypeID() uint32 {
	return BotCommandScopeChatMemberTypeID
}

// TypeName returns name of type in TL schema.
func (*BotCommandScopeChatMember) TypeName() string {
	return "botCommandScopeChatMember"
}

// TypeInfo returns info about TL type.
func (b *BotCommandScopeChatMember) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "botCommandScopeChatMember",
		ID:   BotCommandScopeChatMemberTypeID,
	}
	if b == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "UserID",
			SchemaName: "user_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (b *BotCommandScopeChatMember) Encode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChatMember#f36696f2 as nil")
	}
	buf.PutID(BotCommandScopeChatMemberTypeID)
	return b.EncodeBare(buf)
}

// EncodeBare implements bin.BareEncoder.
func (b *BotCommandScopeChatMember) EncodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChatMember#f36696f2 as nil")
	}
	buf.PutInt53(b.ChatID)
	buf.PutInt53(b.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (b *BotCommandScopeChatMember) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChatMember#f36696f2 to nil")
	}
	if err := buf.ConsumeID(BotCommandScopeChatMemberTypeID); err != nil {
		return fmt.Errorf("unable to decode botCommandScopeChatMember#f36696f2: %w", err)
	}
	return b.DecodeBare(buf)
}

// DecodeBare implements bin.BareDecoder.
func (b *BotCommandScopeChatMember) DecodeBare(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChatMember#f36696f2 to nil")
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode botCommandScopeChatMember#f36696f2: field chat_id: %w", err)
		}
		b.ChatID = value
	}
	{
		value, err := buf.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode botCommandScopeChatMember#f36696f2: field user_id: %w", err)
		}
		b.UserID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (b *BotCommandScopeChatMember) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil {
		return fmt.Errorf("can't encode botCommandScopeChatMember#f36696f2 as nil")
	}
	buf.ObjStart()
	buf.PutID("botCommandScopeChatMember")
	buf.Comma()
	buf.FieldStart("chat_id")
	buf.PutInt53(b.ChatID)
	buf.Comma()
	buf.FieldStart("user_id")
	buf.PutInt53(b.UserID)
	buf.Comma()
	buf.StripComma()
	buf.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (b *BotCommandScopeChatMember) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("can't decode botCommandScopeChatMember#f36696f2 to nil")
	}

	return buf.Obj(func(buf tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := buf.ConsumeID("botCommandScopeChatMember"); err != nil {
				return fmt.Errorf("unable to decode botCommandScopeChatMember#f36696f2: %w", err)
			}
		case "chat_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode botCommandScopeChatMember#f36696f2: field chat_id: %w", err)
			}
			b.ChatID = value
		case "user_id":
			value, err := buf.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode botCommandScopeChatMember#f36696f2: field user_id: %w", err)
			}
			b.UserID = value
		default:
			return buf.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (b *BotCommandScopeChatMember) GetChatID() (value int64) {
	if b == nil {
		return
	}
	return b.ChatID
}

// GetUserID returns value of UserID field.
func (b *BotCommandScopeChatMember) GetUserID() (value int64) {
	if b == nil {
		return
	}
	return b.UserID
}

// BotCommandScopeClassName is schema name of BotCommandScopeClass.
const BotCommandScopeClassName = "BotCommandScope"

// BotCommandScopeClass represents BotCommandScope generic type.
//
// Example:
//
//	g, err := tdapi.DecodeBotCommandScope(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.BotCommandScopeDefault: // botCommandScopeDefault#2f6cb2ab
//	case *tdapi.BotCommandScopeAllPrivateChats: // botCommandScopeAllPrivateChats#eb716739
//	case *tdapi.BotCommandScopeAllGroupChats: // botCommandScopeAllGroupChats#c585c85e
//	case *tdapi.BotCommandScopeAllChatAdministrators: // botCommandScopeAllChatAdministrators#771c1551
//	case *tdapi.BotCommandScopeChat: // botCommandScopeChat#e65b22a5
//	case *tdapi.BotCommandScopeChatAdministrators: // botCommandScopeChatAdministrators#42bcfe4e
//	case *tdapi.BotCommandScopeChatMember: // botCommandScopeChatMember#f36696f2
//	default: panic(v)
//	}
type BotCommandScopeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() BotCommandScopeClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	EncodeTDLibJSON(b tdjson.Encoder) error
	DecodeTDLibJSON(b tdjson.Decoder) error
}

// DecodeBotCommandScope implements binary de-serialization for BotCommandScopeClass.
func DecodeBotCommandScope(buf *bin.Buffer) (BotCommandScopeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case BotCommandScopeDefaultTypeID:
		// Decoding botCommandScopeDefault#2f6cb2ab.
		v := BotCommandScopeDefault{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopeAllPrivateChatsTypeID:
		// Decoding botCommandScopeAllPrivateChats#eb716739.
		v := BotCommandScopeAllPrivateChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopeAllGroupChatsTypeID:
		// Decoding botCommandScopeAllGroupChats#c585c85e.
		v := BotCommandScopeAllGroupChats{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopeAllChatAdministratorsTypeID:
		// Decoding botCommandScopeAllChatAdministrators#771c1551.
		v := BotCommandScopeAllChatAdministrators{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopeChatTypeID:
		// Decoding botCommandScopeChat#e65b22a5.
		v := BotCommandScopeChat{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopeChatAdministratorsTypeID:
		// Decoding botCommandScopeChatAdministrators#42bcfe4e.
		v := BotCommandScopeChatAdministrators{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case BotCommandScopeChatMemberTypeID:
		// Decoding botCommandScopeChatMember#f36696f2.
		v := BotCommandScopeChatMember{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONBotCommandScope implements binary de-serialization for BotCommandScopeClass.
func DecodeTDLibJSONBotCommandScope(buf tdjson.Decoder) (BotCommandScopeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "botCommandScopeDefault":
		// Decoding botCommandScopeDefault#2f6cb2ab.
		v := BotCommandScopeDefault{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case "botCommandScopeAllPrivateChats":
		// Decoding botCommandScopeAllPrivateChats#eb716739.
		v := BotCommandScopeAllPrivateChats{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case "botCommandScopeAllGroupChats":
		// Decoding botCommandScopeAllGroupChats#c585c85e.
		v := BotCommandScopeAllGroupChats{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case "botCommandScopeAllChatAdministrators":
		// Decoding botCommandScopeAllChatAdministrators#771c1551.
		v := BotCommandScopeAllChatAdministrators{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case "botCommandScopeChat":
		// Decoding botCommandScopeChat#e65b22a5.
		v := BotCommandScopeChat{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case "botCommandScopeChatAdministrators":
		// Decoding botCommandScopeChatAdministrators#42bcfe4e.
		v := BotCommandScopeChatAdministrators{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	case "botCommandScopeChatMember":
		// Decoding botCommandScopeChatMember#f36696f2.
		v := BotCommandScopeChatMember{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode BotCommandScopeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// BotCommandScope boxes the BotCommandScopeClass providing a helper.
type BotCommandScopeBox struct {
	BotCommandScope BotCommandScopeClass
}

// Decode implements bin.Decoder for BotCommandScopeBox.
func (b *BotCommandScopeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode BotCommandScopeBox to nil")
	}
	v, err := DecodeBotCommandScope(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BotCommandScope = v
	return nil
}

// Encode implements bin.Encode for BotCommandScopeBox.
func (b *BotCommandScopeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.BotCommandScope == nil {
		return fmt.Errorf("unable to encode BotCommandScopeClass as nil")
	}
	return b.BotCommandScope.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for BotCommandScopeBox.
func (b *BotCommandScopeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode BotCommandScopeBox to nil")
	}
	v, err := DecodeTDLibJSONBotCommandScope(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.BotCommandScope = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for BotCommandScopeBox.
func (b *BotCommandScopeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.BotCommandScope == nil {
		return fmt.Errorf("unable to encode BotCommandScopeClass as nil")
	}
	return b.BotCommandScope.EncodeTDLibJSON(buf)
}
