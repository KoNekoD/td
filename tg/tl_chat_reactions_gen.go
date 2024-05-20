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

// ChatReactionsNone represents TL type `chatReactionsNone#eafc32bc`.
// No reactions are allowed
//
// See https://core.telegram.org/constructor/chatReactionsNone for reference.
type ChatReactionsNone struct {
}

// ChatReactionsNoneTypeID is TL type id of ChatReactionsNone.
const ChatReactionsNoneTypeID = 0xeafc32bc

// construct implements constructor of ChatReactionsClass.
func (c ChatReactionsNone) construct() ChatReactionsClass { return &c }

// Ensuring interfaces in compile-time for ChatReactionsNone.
var (
	_ bin.Encoder     = &ChatReactionsNone{}
	_ bin.Decoder     = &ChatReactionsNone{}
	_ bin.BareEncoder = &ChatReactionsNone{}
	_ bin.BareDecoder = &ChatReactionsNone{}

	_ ChatReactionsClass = &ChatReactionsNone{}
)

func (c *ChatReactionsNone) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatReactionsNone) String() string {
	if c == nil {
		return "ChatReactionsNone(nil)"
	}
	type Alias ChatReactionsNone
	return fmt.Sprintf("ChatReactionsNone%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatReactionsNone) TypeID() uint32 {
	return ChatReactionsNoneTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatReactionsNone) TypeName() string {
	return "chatReactionsNone"
}

// TypeInfo returns info about TL type.
func (c *ChatReactionsNone) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatReactionsNone",
		ID:   ChatReactionsNoneTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatReactionsNone) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatReactionsNone#eafc32bc as nil")
	}
	b.PutID(ChatReactionsNoneTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatReactionsNone) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatReactionsNone#eafc32bc as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatReactionsNone) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatReactionsNone#eafc32bc to nil")
	}
	if err := b.ConsumeID(ChatReactionsNoneTypeID); err != nil {
		return fmt.Errorf("unable to decode chatReactionsNone#eafc32bc: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatReactionsNone) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatReactionsNone#eafc32bc to nil")
	}
	return nil
}

// ChatReactionsAll represents TL type `chatReactionsAll#52928bca`.
// All reactions or all non-custom reactions are allowed
//
// See https://core.telegram.org/constructor/chatReactionsAll for reference.
type ChatReactionsAll struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to allow custom reactions
	AllowCustom bool
}

// ChatReactionsAllTypeID is TL type id of ChatReactionsAll.
const ChatReactionsAllTypeID = 0x52928bca

// construct implements constructor of ChatReactionsClass.
func (c ChatReactionsAll) construct() ChatReactionsClass { return &c }

// Ensuring interfaces in compile-time for ChatReactionsAll.
var (
	_ bin.Encoder     = &ChatReactionsAll{}
	_ bin.Decoder     = &ChatReactionsAll{}
	_ bin.BareEncoder = &ChatReactionsAll{}
	_ bin.BareDecoder = &ChatReactionsAll{}

	_ ChatReactionsClass = &ChatReactionsAll{}
)

func (c *ChatReactionsAll) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.AllowCustom == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatReactionsAll) String() string {
	if c == nil {
		return "ChatReactionsAll(nil)"
	}
	type Alias ChatReactionsAll
	return fmt.Sprintf("ChatReactionsAll%+v", Alias(*c))
}

// FillFrom fills ChatReactionsAll from given interface.
func (c *ChatReactionsAll) FillFrom(from interface {
	GetAllowCustom() (value bool)
}) {
	c.AllowCustom = from.GetAllowCustom()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatReactionsAll) TypeID() uint32 {
	return ChatReactionsAllTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatReactionsAll) TypeName() string {
	return "chatReactionsAll"
}

// TypeInfo returns info about TL type.
func (c *ChatReactionsAll) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatReactionsAll",
		ID:   ChatReactionsAllTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "AllowCustom",
			SchemaName: "allow_custom",
			Null:       !c.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (c *ChatReactionsAll) SetFlags() {
	if !(c.AllowCustom == false) {
		c.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (c *ChatReactionsAll) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatReactionsAll#52928bca as nil")
	}
	b.PutID(ChatReactionsAllTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatReactionsAll) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatReactionsAll#52928bca as nil")
	}
	c.SetFlags()
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode chatReactionsAll#52928bca: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatReactionsAll) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatReactionsAll#52928bca to nil")
	}
	if err := b.ConsumeID(ChatReactionsAllTypeID); err != nil {
		return fmt.Errorf("unable to decode chatReactionsAll#52928bca: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatReactionsAll) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatReactionsAll#52928bca to nil")
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode chatReactionsAll#52928bca: field flags: %w", err)
		}
	}
	c.AllowCustom = c.Flags.Has(0)
	return nil
}

// SetAllowCustom sets value of AllowCustom conditional field.
func (c *ChatReactionsAll) SetAllowCustom(value bool) {
	if value {
		c.Flags.Set(0)
		c.AllowCustom = true
	} else {
		c.Flags.Unset(0)
		c.AllowCustom = false
	}
}

// GetAllowCustom returns value of AllowCustom conditional field.
func (c *ChatReactionsAll) GetAllowCustom() (value bool) {
	if c == nil {
		return
	}
	return c.Flags.Has(0)
}

// ChatReactionsSome represents TL type `chatReactionsSome#661d4037`.
// Some reactions are allowed
//
// See https://core.telegram.org/constructor/chatReactionsSome for reference.
type ChatReactionsSome struct {
	// Allowed set of reactions: the reactions_in_chat_max¹ configuration field indicates
	// the maximum number of reactions that can be specified in this field.
	//
	// Links:
	//  1) https://core.telegram.org/api/config#reactions-in-chat-max
	Reactions []ReactionClass
}

// ChatReactionsSomeTypeID is TL type id of ChatReactionsSome.
const ChatReactionsSomeTypeID = 0x661d4037

// construct implements constructor of ChatReactionsClass.
func (c ChatReactionsSome) construct() ChatReactionsClass { return &c }

// Ensuring interfaces in compile-time for ChatReactionsSome.
var (
	_ bin.Encoder     = &ChatReactionsSome{}
	_ bin.Decoder     = &ChatReactionsSome{}
	_ bin.BareEncoder = &ChatReactionsSome{}
	_ bin.BareDecoder = &ChatReactionsSome{}

	_ ChatReactionsClass = &ChatReactionsSome{}
)

func (c *ChatReactionsSome) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Reactions == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatReactionsSome) String() string {
	if c == nil {
		return "ChatReactionsSome(nil)"
	}
	type Alias ChatReactionsSome
	return fmt.Sprintf("ChatReactionsSome%+v", Alias(*c))
}

// FillFrom fills ChatReactionsSome from given interface.
func (c *ChatReactionsSome) FillFrom(from interface {
	GetReactions() (value []ReactionClass)
}) {
	c.Reactions = from.GetReactions()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatReactionsSome) TypeID() uint32 {
	return ChatReactionsSomeTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatReactionsSome) TypeName() string {
	return "chatReactionsSome"
}

// TypeInfo returns info about TL type.
func (c *ChatReactionsSome) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatReactionsSome",
		ID:   ChatReactionsSomeTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Reactions",
			SchemaName: "reactions",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatReactionsSome) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatReactionsSome#661d4037 as nil")
	}
	b.PutID(ChatReactionsSomeTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatReactionsSome) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatReactionsSome#661d4037 as nil")
	}
	b.PutVectorHeader(len(c.Reactions))
	for idx, v := range c.Reactions {
		if v == nil {
			return fmt.Errorf("unable to encode chatReactionsSome#661d4037: field reactions element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode chatReactionsSome#661d4037: field reactions element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatReactionsSome) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatReactionsSome#661d4037 to nil")
	}
	if err := b.ConsumeID(ChatReactionsSomeTypeID); err != nil {
		return fmt.Errorf("unable to decode chatReactionsSome#661d4037: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatReactionsSome) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatReactionsSome#661d4037 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode chatReactionsSome#661d4037: field reactions: %w", err)
		}

		if headerLen > 0 {
			c.Reactions = make([]ReactionClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeReaction(b)
			if err != nil {
				return fmt.Errorf("unable to decode chatReactionsSome#661d4037: field reactions: %w", err)
			}
			c.Reactions = append(c.Reactions, value)
		}
	}
	return nil
}

// GetReactions returns value of Reactions field.
func (c *ChatReactionsSome) GetReactions() (value []ReactionClass) {
	if c == nil {
		return
	}
	return c.Reactions
}

// MapReactions returns field Reactions wrapped in ReactionClassArray helper.
func (c *ChatReactionsSome) MapReactions() (value ReactionClassArray) {
	return ReactionClassArray(c.Reactions)
}

// ChatReactionsClassName is schema name of ChatReactionsClass.
const ChatReactionsClassName = "ChatReactions"

// ChatReactionsClass represents ChatReactions generic type.
//
// See https://core.telegram.org/type/ChatReactions for reference.
//
// Example:
//
//	g, err := tg.DecodeChatReactions(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.ChatReactionsNone: // chatReactionsNone#eafc32bc
//	case *tg.ChatReactionsAll: // chatReactionsAll#52928bca
//	case *tg.ChatReactionsSome: // chatReactionsSome#661d4037
//	default: panic(v)
//	}
type ChatReactionsClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ChatReactionsClass

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
}

// DecodeChatReactions implements binary de-serialization for ChatReactionsClass.
func DecodeChatReactions(buf *bin.Buffer) (ChatReactionsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatReactionsNoneTypeID:
		// Decoding chatReactionsNone#eafc32bc.
		v := ChatReactionsNone{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatReactionsClass: %w", err)
		}
		return &v, nil
	case ChatReactionsAllTypeID:
		// Decoding chatReactionsAll#52928bca.
		v := ChatReactionsAll{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatReactionsClass: %w", err)
		}
		return &v, nil
	case ChatReactionsSomeTypeID:
		// Decoding chatReactionsSome#661d4037.
		v := ChatReactionsSome{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatReactionsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatReactionsClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatReactions boxes the ChatReactionsClass providing a helper.
type ChatReactionsBox struct {
	ChatReactions ChatReactionsClass
}

// Decode implements bin.Decoder for ChatReactionsBox.
func (b *ChatReactionsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatReactionsBox to nil")
	}
	v, err := DecodeChatReactions(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatReactions = v
	return nil
}

// Encode implements bin.Encode for ChatReactionsBox.
func (b *ChatReactionsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatReactions == nil {
		return fmt.Errorf("unable to encode ChatReactionsClass as nil")
	}
	return b.ChatReactions.Encode(buf)
}
