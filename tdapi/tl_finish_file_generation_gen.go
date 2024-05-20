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

// FinishFileGenerationRequest represents TL type `finishFileGeneration#c11d0c9d`.
type FinishFileGenerationRequest struct {
	// The identifier of the generation process
	GenerationID int64
	// If passed, the file generation has failed and must be terminated; pass null if the
	// file generation succeeded
	Error Error
}

// FinishFileGenerationRequestTypeID is TL type id of FinishFileGenerationRequest.
const FinishFileGenerationRequestTypeID = 0xc11d0c9d

// Ensuring interfaces in compile-time for FinishFileGenerationRequest.
var (
	_ bin.Encoder     = &FinishFileGenerationRequest{}
	_ bin.Decoder     = &FinishFileGenerationRequest{}
	_ bin.BareEncoder = &FinishFileGenerationRequest{}
	_ bin.BareDecoder = &FinishFileGenerationRequest{}
)

func (f *FinishFileGenerationRequest) Zero() bool {
	if f == nil {
		return true
	}
	if !(f.GenerationID == 0) {
		return false
	}
	if !(f.Error.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (f *FinishFileGenerationRequest) String() string {
	if f == nil {
		return "FinishFileGenerationRequest(nil)"
	}
	type Alias FinishFileGenerationRequest
	return fmt.Sprintf("FinishFileGenerationRequest%+v", Alias(*f))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*FinishFileGenerationRequest) TypeID() uint32 {
	return FinishFileGenerationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*FinishFileGenerationRequest) TypeName() string {
	return "finishFileGeneration"
}

// TypeInfo returns info about TL type.
func (f *FinishFileGenerationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "finishFileGeneration",
		ID:   FinishFileGenerationRequestTypeID,
	}
	if f == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "GenerationID",
			SchemaName: "generation_id",
		},
		{
			Name:       "Error",
			SchemaName: "error",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (f *FinishFileGenerationRequest) Encode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode finishFileGeneration#c11d0c9d as nil")
	}
	b.PutID(FinishFileGenerationRequestTypeID)
	return f.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (f *FinishFileGenerationRequest) EncodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't encode finishFileGeneration#c11d0c9d as nil")
	}
	b.PutLong(f.GenerationID)
	if err := f.Error.Encode(b); err != nil {
		return fmt.Errorf("unable to encode finishFileGeneration#c11d0c9d: field error: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (f *FinishFileGenerationRequest) Decode(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode finishFileGeneration#c11d0c9d to nil")
	}
	if err := b.ConsumeID(FinishFileGenerationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode finishFileGeneration#c11d0c9d: %w", err)
	}
	return f.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (f *FinishFileGenerationRequest) DecodeBare(b *bin.Buffer) error {
	if f == nil {
		return fmt.Errorf("can't decode finishFileGeneration#c11d0c9d to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode finishFileGeneration#c11d0c9d: field generation_id: %w", err)
		}
		f.GenerationID = value
	}
	{
		if err := f.Error.Decode(b); err != nil {
			return fmt.Errorf("unable to decode finishFileGeneration#c11d0c9d: field error: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (f *FinishFileGenerationRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if f == nil {
		return fmt.Errorf("can't encode finishFileGeneration#c11d0c9d as nil")
	}
	b.ObjStart()
	b.PutID("finishFileGeneration")
	b.Comma()
	b.FieldStart("generation_id")
	b.PutLong(f.GenerationID)
	b.Comma()
	b.FieldStart("error")
	if err := f.Error.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode finishFileGeneration#c11d0c9d: field error: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (f *FinishFileGenerationRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if f == nil {
		return fmt.Errorf("can't decode finishFileGeneration#c11d0c9d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("finishFileGeneration"); err != nil {
				return fmt.Errorf("unable to decode finishFileGeneration#c11d0c9d: %w", err)
			}
		case "generation_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode finishFileGeneration#c11d0c9d: field generation_id: %w", err)
			}
			f.GenerationID = value
		case "error":
			if err := f.Error.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode finishFileGeneration#c11d0c9d: field error: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetGenerationID returns value of GenerationID field.
func (f *FinishFileGenerationRequest) GetGenerationID() (value int64) {
	if f == nil {
		return
	}
	return f.GenerationID
}

// GetError returns value of Error field.
func (f *FinishFileGenerationRequest) GetError() (value Error) {
	if f == nil {
		return
	}
	return f.Error
}

// FinishFileGeneration invokes method finishFileGeneration#c11d0c9d returning error if any.
func (c *Client) FinishFileGeneration(ctx context.Context, request *FinishFileGenerationRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
