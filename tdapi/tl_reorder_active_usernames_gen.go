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

// ReorderActiveUsernamesRequest represents TL type `reorderActiveUsernames#c659414f`.
type ReorderActiveUsernamesRequest struct {
	// The new order of active usernames. All currently active usernames must be specified
	Usernames []string
}

// ReorderActiveUsernamesRequestTypeID is TL type id of ReorderActiveUsernamesRequest.
const ReorderActiveUsernamesRequestTypeID = 0xc659414f

// Ensuring interfaces in compile-time for ReorderActiveUsernamesRequest.
var (
	_ bin.Encoder     = &ReorderActiveUsernamesRequest{}
	_ bin.Decoder     = &ReorderActiveUsernamesRequest{}
	_ bin.BareEncoder = &ReorderActiveUsernamesRequest{}
	_ bin.BareDecoder = &ReorderActiveUsernamesRequest{}
)

func (r *ReorderActiveUsernamesRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Usernames == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *ReorderActiveUsernamesRequest) String() string {
	if r == nil {
		return "ReorderActiveUsernamesRequest(nil)"
	}
	type Alias ReorderActiveUsernamesRequest
	return fmt.Sprintf("ReorderActiveUsernamesRequest%+v", Alias(*r))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ReorderActiveUsernamesRequest) TypeID() uint32 {
	return ReorderActiveUsernamesRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ReorderActiveUsernamesRequest) TypeName() string {
	return "reorderActiveUsernames"
}

// TypeInfo returns info about TL type.
func (r *ReorderActiveUsernamesRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "reorderActiveUsernames",
		ID:   ReorderActiveUsernamesRequestTypeID,
	}
	if r == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Usernames",
			SchemaName: "usernames",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (r *ReorderActiveUsernamesRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reorderActiveUsernames#c659414f as nil")
	}
	b.PutID(ReorderActiveUsernamesRequestTypeID)
	return r.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (r *ReorderActiveUsernamesRequest) EncodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode reorderActiveUsernames#c659414f as nil")
	}
	b.PutInt(len(r.Usernames))
	for _, v := range r.Usernames {
		b.PutString(v)
	}
	return nil
}

// Decode implements bin.Decoder.
func (r *ReorderActiveUsernamesRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reorderActiveUsernames#c659414f to nil")
	}
	if err := b.ConsumeID(ReorderActiveUsernamesRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode reorderActiveUsernames#c659414f: %w", err)
	}
	return r.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (r *ReorderActiveUsernamesRequest) DecodeBare(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode reorderActiveUsernames#c659414f to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode reorderActiveUsernames#c659414f: field usernames: %w", err)
		}

		if headerLen > 0 {
			r.Usernames = make([]string, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode reorderActiveUsernames#c659414f: field usernames: %w", err)
			}
			r.Usernames = append(r.Usernames, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (r *ReorderActiveUsernamesRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if r == nil {
		return fmt.Errorf("can't encode reorderActiveUsernames#c659414f as nil")
	}
	b.ObjStart()
	b.PutID("reorderActiveUsernames")
	b.Comma()
	b.FieldStart("usernames")
	b.ArrStart()
	for _, v := range r.Usernames {
		b.PutString(v)
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
func (r *ReorderActiveUsernamesRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if r == nil {
		return fmt.Errorf("can't decode reorderActiveUsernames#c659414f to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("reorderActiveUsernames"); err != nil {
				return fmt.Errorf("unable to decode reorderActiveUsernames#c659414f: %w", err)
			}
		case "usernames":
			if err := b.Arr(func(b tdjson.Decoder) error {
				value, err := b.String()
				if err != nil {
					return fmt.Errorf("unable to decode reorderActiveUsernames#c659414f: field usernames: %w", err)
				}
				r.Usernames = append(r.Usernames, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode reorderActiveUsernames#c659414f: field usernames: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetUsernames returns value of Usernames field.
func (r *ReorderActiveUsernamesRequest) GetUsernames() (value []string) {
	if r == nil {
		return
	}
	return r.Usernames
}

// ReorderActiveUsernames invokes method reorderActiveUsernames#c659414f returning error if any.
func (c *Client) ReorderActiveUsernames(ctx context.Context, usernames []string) error {
	var ok Ok

	request := &ReorderActiveUsernamesRequest{
		Usernames: usernames,
	}
	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
