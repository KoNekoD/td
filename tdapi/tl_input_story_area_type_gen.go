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

// InputStoryAreaTypeLocation represents TL type `inputStoryAreaTypeLocation#6849d762`.
type InputStoryAreaTypeLocation struct {
	// The location
	Location Location
}

// InputStoryAreaTypeLocationTypeID is TL type id of InputStoryAreaTypeLocation.
const InputStoryAreaTypeLocationTypeID = 0x6849d762

// construct implements constructor of InputStoryAreaTypeClass.
func (i InputStoryAreaTypeLocation) construct() InputStoryAreaTypeClass { return &i }

// Ensuring interfaces in compile-time for InputStoryAreaTypeLocation.
var (
	_ bin.Encoder     = &InputStoryAreaTypeLocation{}
	_ bin.Decoder     = &InputStoryAreaTypeLocation{}
	_ bin.BareEncoder = &InputStoryAreaTypeLocation{}
	_ bin.BareDecoder = &InputStoryAreaTypeLocation{}

	_ InputStoryAreaTypeClass = &InputStoryAreaTypeLocation{}
)

func (i *InputStoryAreaTypeLocation) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Location.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStoryAreaTypeLocation) String() string {
	if i == nil {
		return "InputStoryAreaTypeLocation(nil)"
	}
	type Alias InputStoryAreaTypeLocation
	return fmt.Sprintf("InputStoryAreaTypeLocation%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStoryAreaTypeLocation) TypeID() uint32 {
	return InputStoryAreaTypeLocationTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStoryAreaTypeLocation) TypeName() string {
	return "inputStoryAreaTypeLocation"
}

// TypeInfo returns info about TL type.
func (i *InputStoryAreaTypeLocation) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStoryAreaTypeLocation",
		ID:   InputStoryAreaTypeLocationTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Location",
			SchemaName: "location",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStoryAreaTypeLocation) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeLocation#6849d762 as nil")
	}
	b.PutID(InputStoryAreaTypeLocationTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStoryAreaTypeLocation) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeLocation#6849d762 as nil")
	}
	if err := i.Location.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStoryAreaTypeLocation#6849d762: field location: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStoryAreaTypeLocation) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeLocation#6849d762 to nil")
	}
	if err := b.ConsumeID(InputStoryAreaTypeLocationTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStoryAreaTypeLocation#6849d762: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStoryAreaTypeLocation) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeLocation#6849d762 to nil")
	}
	{
		if err := i.Location.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputStoryAreaTypeLocation#6849d762: field location: %w", err)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputStoryAreaTypeLocation) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeLocation#6849d762 as nil")
	}
	b.ObjStart()
	b.PutID("inputStoryAreaTypeLocation")
	b.Comma()
	b.FieldStart("location")
	if err := i.Location.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputStoryAreaTypeLocation#6849d762: field location: %w", err)
	}
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputStoryAreaTypeLocation) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeLocation#6849d762 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputStoryAreaTypeLocation"); err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeLocation#6849d762: %w", err)
			}
		case "location":
			if err := i.Location.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeLocation#6849d762: field location: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetLocation returns value of Location field.
func (i *InputStoryAreaTypeLocation) GetLocation() (value Location) {
	if i == nil {
		return
	}
	return i.Location
}

// InputStoryAreaTypeFoundVenue represents TL type `inputStoryAreaTypeFoundVenue#accda496`.
type InputStoryAreaTypeFoundVenue struct {
	// Identifier of the inline query, used to found the venue
	QueryID int64
	// Identifier of the inline query result
	ResultID string
}

// InputStoryAreaTypeFoundVenueTypeID is TL type id of InputStoryAreaTypeFoundVenue.
const InputStoryAreaTypeFoundVenueTypeID = 0xaccda496

// construct implements constructor of InputStoryAreaTypeClass.
func (i InputStoryAreaTypeFoundVenue) construct() InputStoryAreaTypeClass { return &i }

// Ensuring interfaces in compile-time for InputStoryAreaTypeFoundVenue.
var (
	_ bin.Encoder     = &InputStoryAreaTypeFoundVenue{}
	_ bin.Decoder     = &InputStoryAreaTypeFoundVenue{}
	_ bin.BareEncoder = &InputStoryAreaTypeFoundVenue{}
	_ bin.BareDecoder = &InputStoryAreaTypeFoundVenue{}

	_ InputStoryAreaTypeClass = &InputStoryAreaTypeFoundVenue{}
)

func (i *InputStoryAreaTypeFoundVenue) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.QueryID == 0) {
		return false
	}
	if !(i.ResultID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStoryAreaTypeFoundVenue) String() string {
	if i == nil {
		return "InputStoryAreaTypeFoundVenue(nil)"
	}
	type Alias InputStoryAreaTypeFoundVenue
	return fmt.Sprintf("InputStoryAreaTypeFoundVenue%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStoryAreaTypeFoundVenue) TypeID() uint32 {
	return InputStoryAreaTypeFoundVenueTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStoryAreaTypeFoundVenue) TypeName() string {
	return "inputStoryAreaTypeFoundVenue"
}

// TypeInfo returns info about TL type.
func (i *InputStoryAreaTypeFoundVenue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStoryAreaTypeFoundVenue",
		ID:   InputStoryAreaTypeFoundVenueTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "QueryID",
			SchemaName: "query_id",
		},
		{
			Name:       "ResultID",
			SchemaName: "result_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStoryAreaTypeFoundVenue) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeFoundVenue#accda496 as nil")
	}
	b.PutID(InputStoryAreaTypeFoundVenueTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStoryAreaTypeFoundVenue) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeFoundVenue#accda496 as nil")
	}
	b.PutLong(i.QueryID)
	b.PutString(i.ResultID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStoryAreaTypeFoundVenue) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeFoundVenue#accda496 to nil")
	}
	if err := b.ConsumeID(InputStoryAreaTypeFoundVenueTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStoryAreaTypeFoundVenue#accda496: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStoryAreaTypeFoundVenue) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeFoundVenue#accda496 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryAreaTypeFoundVenue#accda496: field query_id: %w", err)
		}
		i.QueryID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryAreaTypeFoundVenue#accda496: field result_id: %w", err)
		}
		i.ResultID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputStoryAreaTypeFoundVenue) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeFoundVenue#accda496 as nil")
	}
	b.ObjStart()
	b.PutID("inputStoryAreaTypeFoundVenue")
	b.Comma()
	b.FieldStart("query_id")
	b.PutLong(i.QueryID)
	b.Comma()
	b.FieldStart("result_id")
	b.PutString(i.ResultID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputStoryAreaTypeFoundVenue) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeFoundVenue#accda496 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputStoryAreaTypeFoundVenue"); err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeFoundVenue#accda496: %w", err)
			}
		case "query_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeFoundVenue#accda496: field query_id: %w", err)
			}
			i.QueryID = value
		case "result_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeFoundVenue#accda496: field result_id: %w", err)
			}
			i.ResultID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetQueryID returns value of QueryID field.
func (i *InputStoryAreaTypeFoundVenue) GetQueryID() (value int64) {
	if i == nil {
		return
	}
	return i.QueryID
}

// GetResultID returns value of ResultID field.
func (i *InputStoryAreaTypeFoundVenue) GetResultID() (value string) {
	if i == nil {
		return
	}
	return i.ResultID
}

// InputStoryAreaTypePreviousVenue represents TL type `inputStoryAreaTypePreviousVenue#6e124e0c`.
type InputStoryAreaTypePreviousVenue struct {
	// Provider of the venue
	VenueProvider string
	// Identifier of the venue in the provider database
	VenueID string
}

// InputStoryAreaTypePreviousVenueTypeID is TL type id of InputStoryAreaTypePreviousVenue.
const InputStoryAreaTypePreviousVenueTypeID = 0x6e124e0c

// construct implements constructor of InputStoryAreaTypeClass.
func (i InputStoryAreaTypePreviousVenue) construct() InputStoryAreaTypeClass { return &i }

// Ensuring interfaces in compile-time for InputStoryAreaTypePreviousVenue.
var (
	_ bin.Encoder     = &InputStoryAreaTypePreviousVenue{}
	_ bin.Decoder     = &InputStoryAreaTypePreviousVenue{}
	_ bin.BareEncoder = &InputStoryAreaTypePreviousVenue{}
	_ bin.BareDecoder = &InputStoryAreaTypePreviousVenue{}

	_ InputStoryAreaTypeClass = &InputStoryAreaTypePreviousVenue{}
)

func (i *InputStoryAreaTypePreviousVenue) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.VenueProvider == "") {
		return false
	}
	if !(i.VenueID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStoryAreaTypePreviousVenue) String() string {
	if i == nil {
		return "InputStoryAreaTypePreviousVenue(nil)"
	}
	type Alias InputStoryAreaTypePreviousVenue
	return fmt.Sprintf("InputStoryAreaTypePreviousVenue%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStoryAreaTypePreviousVenue) TypeID() uint32 {
	return InputStoryAreaTypePreviousVenueTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStoryAreaTypePreviousVenue) TypeName() string {
	return "inputStoryAreaTypePreviousVenue"
}

// TypeInfo returns info about TL type.
func (i *InputStoryAreaTypePreviousVenue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStoryAreaTypePreviousVenue",
		ID:   InputStoryAreaTypePreviousVenueTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "VenueProvider",
			SchemaName: "venue_provider",
		},
		{
			Name:       "VenueID",
			SchemaName: "venue_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStoryAreaTypePreviousVenue) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypePreviousVenue#6e124e0c as nil")
	}
	b.PutID(InputStoryAreaTypePreviousVenueTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStoryAreaTypePreviousVenue) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypePreviousVenue#6e124e0c as nil")
	}
	b.PutString(i.VenueProvider)
	b.PutString(i.VenueID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStoryAreaTypePreviousVenue) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypePreviousVenue#6e124e0c to nil")
	}
	if err := b.ConsumeID(InputStoryAreaTypePreviousVenueTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStoryAreaTypePreviousVenue#6e124e0c: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStoryAreaTypePreviousVenue) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypePreviousVenue#6e124e0c to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryAreaTypePreviousVenue#6e124e0c: field venue_provider: %w", err)
		}
		i.VenueProvider = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryAreaTypePreviousVenue#6e124e0c: field venue_id: %w", err)
		}
		i.VenueID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputStoryAreaTypePreviousVenue) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypePreviousVenue#6e124e0c as nil")
	}
	b.ObjStart()
	b.PutID("inputStoryAreaTypePreviousVenue")
	b.Comma()
	b.FieldStart("venue_provider")
	b.PutString(i.VenueProvider)
	b.Comma()
	b.FieldStart("venue_id")
	b.PutString(i.VenueID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputStoryAreaTypePreviousVenue) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypePreviousVenue#6e124e0c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputStoryAreaTypePreviousVenue"); err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypePreviousVenue#6e124e0c: %w", err)
			}
		case "venue_provider":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypePreviousVenue#6e124e0c: field venue_provider: %w", err)
			}
			i.VenueProvider = value
		case "venue_id":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypePreviousVenue#6e124e0c: field venue_id: %w", err)
			}
			i.VenueID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetVenueProvider returns value of VenueProvider field.
func (i *InputStoryAreaTypePreviousVenue) GetVenueProvider() (value string) {
	if i == nil {
		return
	}
	return i.VenueProvider
}

// GetVenueID returns value of VenueID field.
func (i *InputStoryAreaTypePreviousVenue) GetVenueID() (value string) {
	if i == nil {
		return
	}
	return i.VenueID
}

// InputStoryAreaTypeSuggestedReaction represents TL type `inputStoryAreaTypeSuggestedReaction#7d4751d3`.
type InputStoryAreaTypeSuggestedReaction struct {
	// Type of the reaction
	ReactionType ReactionTypeClass
	// True, if reaction has a dark background
	IsDark bool
	// True, if reaction corner is flipped
	IsFlipped bool
}

// InputStoryAreaTypeSuggestedReactionTypeID is TL type id of InputStoryAreaTypeSuggestedReaction.
const InputStoryAreaTypeSuggestedReactionTypeID = 0x7d4751d3

// construct implements constructor of InputStoryAreaTypeClass.
func (i InputStoryAreaTypeSuggestedReaction) construct() InputStoryAreaTypeClass { return &i }

// Ensuring interfaces in compile-time for InputStoryAreaTypeSuggestedReaction.
var (
	_ bin.Encoder     = &InputStoryAreaTypeSuggestedReaction{}
	_ bin.Decoder     = &InputStoryAreaTypeSuggestedReaction{}
	_ bin.BareEncoder = &InputStoryAreaTypeSuggestedReaction{}
	_ bin.BareDecoder = &InputStoryAreaTypeSuggestedReaction{}

	_ InputStoryAreaTypeClass = &InputStoryAreaTypeSuggestedReaction{}
)

func (i *InputStoryAreaTypeSuggestedReaction) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ReactionType == nil) {
		return false
	}
	if !(i.IsDark == false) {
		return false
	}
	if !(i.IsFlipped == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStoryAreaTypeSuggestedReaction) String() string {
	if i == nil {
		return "InputStoryAreaTypeSuggestedReaction(nil)"
	}
	type Alias InputStoryAreaTypeSuggestedReaction
	return fmt.Sprintf("InputStoryAreaTypeSuggestedReaction%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStoryAreaTypeSuggestedReaction) TypeID() uint32 {
	return InputStoryAreaTypeSuggestedReactionTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStoryAreaTypeSuggestedReaction) TypeName() string {
	return "inputStoryAreaTypeSuggestedReaction"
}

// TypeInfo returns info about TL type.
func (i *InputStoryAreaTypeSuggestedReaction) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStoryAreaTypeSuggestedReaction",
		ID:   InputStoryAreaTypeSuggestedReactionTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ReactionType",
			SchemaName: "reaction_type",
		},
		{
			Name:       "IsDark",
			SchemaName: "is_dark",
		},
		{
			Name:       "IsFlipped",
			SchemaName: "is_flipped",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStoryAreaTypeSuggestedReaction) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeSuggestedReaction#7d4751d3 as nil")
	}
	b.PutID(InputStoryAreaTypeSuggestedReactionTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStoryAreaTypeSuggestedReaction) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeSuggestedReaction#7d4751d3 as nil")
	}
	if i.ReactionType == nil {
		return fmt.Errorf("unable to encode inputStoryAreaTypeSuggestedReaction#7d4751d3: field reaction_type is nil")
	}
	if err := i.ReactionType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputStoryAreaTypeSuggestedReaction#7d4751d3: field reaction_type: %w", err)
	}
	b.PutBool(i.IsDark)
	b.PutBool(i.IsFlipped)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStoryAreaTypeSuggestedReaction) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeSuggestedReaction#7d4751d3 to nil")
	}
	if err := b.ConsumeID(InputStoryAreaTypeSuggestedReactionTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStoryAreaTypeSuggestedReaction#7d4751d3: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStoryAreaTypeSuggestedReaction) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeSuggestedReaction#7d4751d3 to nil")
	}
	{
		value, err := DecodeReactionType(b)
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryAreaTypeSuggestedReaction#7d4751d3: field reaction_type: %w", err)
		}
		i.ReactionType = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryAreaTypeSuggestedReaction#7d4751d3: field is_dark: %w", err)
		}
		i.IsDark = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryAreaTypeSuggestedReaction#7d4751d3: field is_flipped: %w", err)
		}
		i.IsFlipped = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputStoryAreaTypeSuggestedReaction) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeSuggestedReaction#7d4751d3 as nil")
	}
	b.ObjStart()
	b.PutID("inputStoryAreaTypeSuggestedReaction")
	b.Comma()
	b.FieldStart("reaction_type")
	if i.ReactionType == nil {
		return fmt.Errorf("unable to encode inputStoryAreaTypeSuggestedReaction#7d4751d3: field reaction_type is nil")
	}
	if err := i.ReactionType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode inputStoryAreaTypeSuggestedReaction#7d4751d3: field reaction_type: %w", err)
	}
	b.Comma()
	b.FieldStart("is_dark")
	b.PutBool(i.IsDark)
	b.Comma()
	b.FieldStart("is_flipped")
	b.PutBool(i.IsFlipped)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputStoryAreaTypeSuggestedReaction) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeSuggestedReaction#7d4751d3 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputStoryAreaTypeSuggestedReaction"); err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeSuggestedReaction#7d4751d3: %w", err)
			}
		case "reaction_type":
			value, err := DecodeTDLibJSONReactionType(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeSuggestedReaction#7d4751d3: field reaction_type: %w", err)
			}
			i.ReactionType = value
		case "is_dark":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeSuggestedReaction#7d4751d3: field is_dark: %w", err)
			}
			i.IsDark = value
		case "is_flipped":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeSuggestedReaction#7d4751d3: field is_flipped: %w", err)
			}
			i.IsFlipped = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetReactionType returns value of ReactionType field.
func (i *InputStoryAreaTypeSuggestedReaction) GetReactionType() (value ReactionTypeClass) {
	if i == nil {
		return
	}
	return i.ReactionType
}

// GetIsDark returns value of IsDark field.
func (i *InputStoryAreaTypeSuggestedReaction) GetIsDark() (value bool) {
	if i == nil {
		return
	}
	return i.IsDark
}

// GetIsFlipped returns value of IsFlipped field.
func (i *InputStoryAreaTypeSuggestedReaction) GetIsFlipped() (value bool) {
	if i == nil {
		return
	}
	return i.IsFlipped
}

// InputStoryAreaTypeMessage represents TL type `inputStoryAreaTypeMessage#f01be457`.
type InputStoryAreaTypeMessage struct {
	// Identifier of the chat with the message. Currently, the chat must be a supergroup or a
	// channel chat
	ChatID int64
	// Identifier of the message. Only successfully sent non-scheduled messages can be
	// specified
	MessageID int64
}

// InputStoryAreaTypeMessageTypeID is TL type id of InputStoryAreaTypeMessage.
const InputStoryAreaTypeMessageTypeID = 0xf01be457

// construct implements constructor of InputStoryAreaTypeClass.
func (i InputStoryAreaTypeMessage) construct() InputStoryAreaTypeClass { return &i }

// Ensuring interfaces in compile-time for InputStoryAreaTypeMessage.
var (
	_ bin.Encoder     = &InputStoryAreaTypeMessage{}
	_ bin.Decoder     = &InputStoryAreaTypeMessage{}
	_ bin.BareEncoder = &InputStoryAreaTypeMessage{}
	_ bin.BareDecoder = &InputStoryAreaTypeMessage{}

	_ InputStoryAreaTypeClass = &InputStoryAreaTypeMessage{}
)

func (i *InputStoryAreaTypeMessage) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ChatID == 0) {
		return false
	}
	if !(i.MessageID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputStoryAreaTypeMessage) String() string {
	if i == nil {
		return "InputStoryAreaTypeMessage(nil)"
	}
	type Alias InputStoryAreaTypeMessage
	return fmt.Sprintf("InputStoryAreaTypeMessage%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputStoryAreaTypeMessage) TypeID() uint32 {
	return InputStoryAreaTypeMessageTypeID
}

// TypeName returns name of type in TL schema.
func (*InputStoryAreaTypeMessage) TypeName() string {
	return "inputStoryAreaTypeMessage"
}

// TypeInfo returns info about TL type.
func (i *InputStoryAreaTypeMessage) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputStoryAreaTypeMessage",
		ID:   InputStoryAreaTypeMessageTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageID",
			SchemaName: "message_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputStoryAreaTypeMessage) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeMessage#f01be457 as nil")
	}
	b.PutID(InputStoryAreaTypeMessageTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputStoryAreaTypeMessage) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeMessage#f01be457 as nil")
	}
	b.PutInt53(i.ChatID)
	b.PutInt53(i.MessageID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputStoryAreaTypeMessage) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeMessage#f01be457 to nil")
	}
	if err := b.ConsumeID(InputStoryAreaTypeMessageTypeID); err != nil {
		return fmt.Errorf("unable to decode inputStoryAreaTypeMessage#f01be457: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputStoryAreaTypeMessage) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeMessage#f01be457 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryAreaTypeMessage#f01be457: field chat_id: %w", err)
		}
		i.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode inputStoryAreaTypeMessage#f01be457: field message_id: %w", err)
		}
		i.MessageID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (i *InputStoryAreaTypeMessage) EncodeTDLibJSON(b tdjson.Encoder) error {
	if i == nil {
		return fmt.Errorf("can't encode inputStoryAreaTypeMessage#f01be457 as nil")
	}
	b.ObjStart()
	b.PutID("inputStoryAreaTypeMessage")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(i.ChatID)
	b.Comma()
	b.FieldStart("message_id")
	b.PutInt53(i.MessageID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (i *InputStoryAreaTypeMessage) DecodeTDLibJSON(b tdjson.Decoder) error {
	if i == nil {
		return fmt.Errorf("can't decode inputStoryAreaTypeMessage#f01be457 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("inputStoryAreaTypeMessage"); err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeMessage#f01be457: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeMessage#f01be457: field chat_id: %w", err)
			}
			i.ChatID = value
		case "message_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode inputStoryAreaTypeMessage#f01be457: field message_id: %w", err)
			}
			i.MessageID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (i *InputStoryAreaTypeMessage) GetChatID() (value int64) {
	if i == nil {
		return
	}
	return i.ChatID
}

// GetMessageID returns value of MessageID field.
func (i *InputStoryAreaTypeMessage) GetMessageID() (value int64) {
	if i == nil {
		return
	}
	return i.MessageID
}

// InputStoryAreaTypeClassName is schema name of InputStoryAreaTypeClass.
const InputStoryAreaTypeClassName = "InputStoryAreaType"

// InputStoryAreaTypeClass represents InputStoryAreaType generic type.
//
// Example:
//
//	g, err := tdapi.DecodeInputStoryAreaType(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tdapi.InputStoryAreaTypeLocation: // inputStoryAreaTypeLocation#6849d762
//	case *tdapi.InputStoryAreaTypeFoundVenue: // inputStoryAreaTypeFoundVenue#accda496
//	case *tdapi.InputStoryAreaTypePreviousVenue: // inputStoryAreaTypePreviousVenue#6e124e0c
//	case *tdapi.InputStoryAreaTypeSuggestedReaction: // inputStoryAreaTypeSuggestedReaction#7d4751d3
//	case *tdapi.InputStoryAreaTypeMessage: // inputStoryAreaTypeMessage#f01be457
//	default: panic(v)
//	}
type InputStoryAreaTypeClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputStoryAreaTypeClass

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

// DecodeInputStoryAreaType implements binary de-serialization for InputStoryAreaTypeClass.
func DecodeInputStoryAreaType(buf *bin.Buffer) (InputStoryAreaTypeClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputStoryAreaTypeLocationTypeID:
		// Decoding inputStoryAreaTypeLocation#6849d762.
		v := InputStoryAreaTypeLocation{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", err)
		}
		return &v, nil
	case InputStoryAreaTypeFoundVenueTypeID:
		// Decoding inputStoryAreaTypeFoundVenue#accda496.
		v := InputStoryAreaTypeFoundVenue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", err)
		}
		return &v, nil
	case InputStoryAreaTypePreviousVenueTypeID:
		// Decoding inputStoryAreaTypePreviousVenue#6e124e0c.
		v := InputStoryAreaTypePreviousVenue{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", err)
		}
		return &v, nil
	case InputStoryAreaTypeSuggestedReactionTypeID:
		// Decoding inputStoryAreaTypeSuggestedReaction#7d4751d3.
		v := InputStoryAreaTypeSuggestedReaction{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", err)
		}
		return &v, nil
	case InputStoryAreaTypeMessageTypeID:
		// Decoding inputStoryAreaTypeMessage#f01be457.
		v := InputStoryAreaTypeMessage{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", bin.NewUnexpectedID(id))
	}
}

// DecodeTDLibJSONInputStoryAreaType implements binary de-serialization for InputStoryAreaTypeClass.
func DecodeTDLibJSONInputStoryAreaType(buf tdjson.Decoder) (InputStoryAreaTypeClass, error) {
	id, err := buf.FindTypeID()
	if err != nil {
		return nil, err
	}
	switch id {
	case "inputStoryAreaTypeLocation":
		// Decoding inputStoryAreaTypeLocation#6849d762.
		v := InputStoryAreaTypeLocation{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", err)
		}
		return &v, nil
	case "inputStoryAreaTypeFoundVenue":
		// Decoding inputStoryAreaTypeFoundVenue#accda496.
		v := InputStoryAreaTypeFoundVenue{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", err)
		}
		return &v, nil
	case "inputStoryAreaTypePreviousVenue":
		// Decoding inputStoryAreaTypePreviousVenue#6e124e0c.
		v := InputStoryAreaTypePreviousVenue{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", err)
		}
		return &v, nil
	case "inputStoryAreaTypeSuggestedReaction":
		// Decoding inputStoryAreaTypeSuggestedReaction#7d4751d3.
		v := InputStoryAreaTypeSuggestedReaction{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", err)
		}
		return &v, nil
	case "inputStoryAreaTypeMessage":
		// Decoding inputStoryAreaTypeMessage#f01be457.
		v := InputStoryAreaTypeMessage{}
		if err := v.DecodeTDLibJSON(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputStoryAreaTypeClass: %w", tdjson.NewUnexpectedID(id))
	}
}

// InputStoryAreaType boxes the InputStoryAreaTypeClass providing a helper.
type InputStoryAreaTypeBox struct {
	InputStoryAreaType InputStoryAreaTypeClass
}

// Decode implements bin.Decoder for InputStoryAreaTypeBox.
func (b *InputStoryAreaTypeBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStoryAreaTypeBox to nil")
	}
	v, err := DecodeInputStoryAreaType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStoryAreaType = v
	return nil
}

// Encode implements bin.Encode for InputStoryAreaTypeBox.
func (b *InputStoryAreaTypeBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputStoryAreaType == nil {
		return fmt.Errorf("unable to encode InputStoryAreaTypeClass as nil")
	}
	return b.InputStoryAreaType.Encode(buf)
}

// DecodeTDLibJSON implements bin.Decoder for InputStoryAreaTypeBox.
func (b *InputStoryAreaTypeBox) DecodeTDLibJSON(buf tdjson.Decoder) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputStoryAreaTypeBox to nil")
	}
	v, err := DecodeTDLibJSONInputStoryAreaType(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputStoryAreaType = v
	return nil
}

// EncodeTDLibJSON implements bin.Encode for InputStoryAreaTypeBox.
func (b *InputStoryAreaTypeBox) EncodeTDLibJSON(buf tdjson.Encoder) error {
	if b == nil || b.InputStoryAreaType == nil {
		return fmt.Errorf("unable to encode InputStoryAreaTypeClass as nil")
	}
	return b.InputStoryAreaType.EncodeTDLibJSON(buf)
}
