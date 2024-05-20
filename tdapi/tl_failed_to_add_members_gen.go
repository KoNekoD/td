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

// FailedToAddMembers represents TL type `failedToAddMembers#98b4b04`.
type FailedToAddMembers struct {
	// Information about users that weren't added to the chat
	FailedToAddMembers []FailedToAddMember
}

// FailedToAddMembersTypeID is TL type id of FailedToAddMembers.
const FailedToAddMembersTypeID = 0x98b4b04

// Ensuring interfaces in compile-time for FailedToAddMembers.
var (
	_ bin.Encoder     = &FailedToAddMembers{}
	_ bin.Decoder     = &FailedToAddMembers{}
	_ bin.BareEncoder = &FailedToAddMembers{}
	_ bin.BareDecoder = &FailedToAddMembers{}
)

func (f *FailedToAddMembers) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.FailedToAddMembers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FailedToAddMembers) String() string {
	if f == nil {
		return "FailedToAddMembers(nil)"
	}
	type Alias FailedToAddMembers
	return fmt.Sprintf("FailedToAddMembers%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FailedToAddMembers) TypeID() uint32 {
	return FailedToAddMembersTypeID
}

// TypeName returns name of type in TL schema.
func (*FailedToAddMembers) TypeName() string {
	return "failedToAddMembers"
}

// TypeInfo returns info about TL type.
func (f *FailedToAddMembers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "failedToAddMembers",
		ID:   FailedToAddMembersTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FailedToAddMembers",
			SchemaName: "failed_to_add_members",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FailedToAddMembers) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode failedToAddMembers#98b4b04 as nil")
	}
	b.PutID(FailedToAddMembersTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FailedToAddMembers) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode failedToAddMembers#98b4b04 as nil")
	}
	b.PutInt(len(f.FailedToAddMembers))
	for idx, v := range f.FailedToAddMembers {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare failedToAddMembers#98b4b04: field failed_to_add_members element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *FailedToAddMembers) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode failedToAddMembers#98b4b04 to nil")
	}
	if err := b.ConsumeID(FailedToAddMembersTypeID); err != nil {
		return fmt.Errorf("unable to decode failedToAddMembers#98b4b04: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FailedToAddMembers) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode failedToAddMembers#98b4b04 to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode failedToAddMembers#98b4b04: field failed_to_add_members: %w", err)
		}

		if headerLen > 0 {
			f.FailedToAddMembers = make([]FailedToAddMember, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value FailedToAddMember
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare failedToAddMembers#98b4b04: field failed_to_add_members: %w", err)
			}
			f.FailedToAddMembers = append(f.FailedToAddMembers, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FailedToAddMembers) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode failedToAddMembers#98b4b04 as nil")
	}
	b.ObjStart()
	b.PutID("failedToAddMembers")
	b.Comma()
	b.FieldStart("failed_to_add_members")
	b.ArrStart()
	for idx, v := range f.FailedToAddMembers {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode failedToAddMembers#98b4b04: field failed_to_add_members element with index %d: %w", idx, err)
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
func (f *FailedToAddMembers) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode failedToAddMembers#98b4b04 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("failedToAddMembers"); err != nil {
				return fmt.Errorf("unable to decode failedToAddMembers#98b4b04: %w", err)
			}
		case "failed_to_add_members":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value FailedToAddMember
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode failedToAddMembers#98b4b04: field failed_to_add_members: %w", err)
				}
				f.FailedToAddMembers = append(f.FailedToAddMembers, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode failedToAddMembers#98b4b04: field failed_to_add_members: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetFailedToAddMembers returns value of FailedToAddMembers field.
func (f *FailedToAddMembers) GetFailedToAddMembers() (value []FailedToAddMember) {
	if f == nil {
		return
	}
	return f.FailedToAddMembers
}
