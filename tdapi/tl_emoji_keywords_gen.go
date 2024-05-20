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

// EmojiKeywords represents TL type `emojiKeywords#de023f4`.
type EmojiKeywords struct {
	// List of emoji with their keywords
	EmojiKeywords []EmojiKeyword
}

// EmojiKeywordsTypeID is TL type id of EmojiKeywords.
const EmojiKeywordsTypeID = 0xde023f4

// Ensuring interfaces in compile-time for EmojiKeywords.
var (
	_ bin.Encoder     = &EmojiKeywords{}
	_ bin.Decoder     = &EmojiKeywords{}
	_ bin.BareEncoder = &EmojiKeywords{}
	_ bin.BareDecoder = &EmojiKeywords{}
)

func (e *EmojiKeywords) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.EmojiKeywords == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EmojiKeywords) String() string {
	if e == nil {
		return "EmojiKeywords(nil)"
	}
	type Alias EmojiKeywords
	return fmt.Sprintf("EmojiKeywords%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EmojiKeywords) TypeID() uint32 {
	return EmojiKeywordsTypeID
}

// TypeName returns name of type in TL schema.
func (*EmojiKeywords) TypeName() string {
	return "emojiKeywords"
}

// TypeInfo returns info about TL type.
func (e *EmojiKeywords) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "emojiKeywords",
		ID:   EmojiKeywordsTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "EmojiKeywords",
			SchemaName: "emoji_keywords",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EmojiKeywords) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeywords#de023f4 as nil")
	}
	b.PutID(EmojiKeywordsTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EmojiKeywords) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeywords#de023f4 as nil")
	}
	b.PutInt(len(e.EmojiKeywords))
	for idx, v := range e.EmojiKeywords {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare emojiKeywords#de023f4: field emoji_keywords element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *EmojiKeywords) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeywords#de023f4 to nil")
	}
	if err := b.ConsumeID(EmojiKeywordsTypeID); err != nil {
		return fmt.Errorf("unable to decode emojiKeywords#de023f4: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EmojiKeywords) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeywords#de023f4 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode emojiKeywords#de023f4: field emoji_keywords: %w", err)
		}

		if headerLen > 0 {
			e.EmojiKeywords = make([]EmojiKeyword, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value EmojiKeyword
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare emojiKeywords#de023f4: field emoji_keywords: %w", err)
			}
			e.EmojiKeywords = append(e.EmojiKeywords, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EmojiKeywords) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode emojiKeywords#de023f4 as nil")
	}
	b.ObjStart()
	b.PutID("emojiKeywords")
	b.Comma()
	b.FieldStart("emoji_keywords")
	b.ArrStart()
	for idx, v := range e.EmojiKeywords {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode emojiKeywords#de023f4: field emoji_keywords element with index %d: %w", idx, err)
		}
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
func (e *EmojiKeywords) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode emojiKeywords#de023f4 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("emojiKeywords"); err != nil {
				return fmt.Errorf("unable to decode emojiKeywords#de023f4: %w", err)
			}
		case "emoji_keywords":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value EmojiKeyword
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode emojiKeywords#de023f4: field emoji_keywords: %w", err)
				}
				e.EmojiKeywords = append(e.EmojiKeywords, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode emojiKeywords#de023f4: field emoji_keywords: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetEmojiKeywords returns value of EmojiKeywords field.
func (e *EmojiKeywords) GetEmojiKeywords() (value []EmojiKeyword) {
	if e == nil {
		return
	}
	return e.EmojiKeywords
}
