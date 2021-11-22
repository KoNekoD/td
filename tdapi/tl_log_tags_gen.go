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

// LogTags represents TL type `logTags#dc09ced4`.
type LogTags struct {
	// List of log tags
	Tags []string
}

// LogTagsTypeID is TL type id of LogTags.
const LogTagsTypeID = 0xdc09ced4

// Ensuring interfaces in compile-time for LogTags.
var (
	_ bin.Encoder     = &LogTags{}
	_ bin.Decoder     = &LogTags{}
	_ bin.BareEncoder = &LogTags{}
	_ bin.BareDecoder = &LogTags{}
)

func (l *LogTags) Zero() bool {
	if l == nil {
		return true
	}
	if !(l.Tags == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (l *LogTags) String() string {
	if l == nil {
		return "LogTags(nil)"
	}
	type Alias LogTags
	return fmt.Sprintf("LogTags%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LogTags) TypeID() uint32 {
	return LogTagsTypeID
}

// TypeName returns name of type in TL schema.
func (*LogTags) TypeName() string {
	return "logTags"
}

// TypeInfo returns info about TL type.
func (l *LogTags) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "logTags",
		ID:   LogTagsTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Tags",
			SchemaName: "tags",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (l *LogTags) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode logTags#dc09ced4 as nil")
	}
	b.PutID(LogTagsTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *LogTags) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode logTags#dc09ced4 as nil")
	}
	b.PutInt(len(l.Tags))
	for _, v := range l.Tags {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *LogTags) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode logTags#dc09ced4 to nil")
	}
	if err := b.ConsumeID(LogTagsTypeID); err != nil {
		return fmt.Errorf("unable to decode logTags#dc09ced4: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *LogTags) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode logTags#dc09ced4 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode logTags#dc09ced4: field tags: %w", err)
		}

		if headerLen > 0 {
			l.Tags = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode logTags#dc09ced4: field tags: %w", err)
			}
			l.Tags = append(l.Tags, value)
		}
	}
	return nil
}

// EncodeTDLibJSON encodes l in TDLib API JSON format.
func (l *LogTags) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if l == nil {
		return fmt.Errorf("can't encode logTags#dc09ced4 as nil")
	}
	b.ObjStart()
	b.PutID("logTags")
	b.FieldStart("tags")
	b.ArrStart()
	for _, v := range l.Tags {
		b.PutString(v)
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// GetTags returns value of Tags field.
func (l *LogTags) GetTags() (value []string) {
	return l.Tags
}
