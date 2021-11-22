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

// PollOption represents TL type `pollOption#57d9d5a5`.
type PollOption struct {
	// Option text; 1-100 characters
	Text string
	// Number of voters for this option, available only for closed or voted polls
	VoterCount int32
	// The percentage of votes for this option; 0-100
	VotePercentage int32
	// True, if the option was chosen by the user
	IsChosen bool
	// True, if the option is being chosen by a pending setPollAnswer request
	IsBeingChosen bool
}

// PollOptionTypeID is TL type id of PollOption.
const PollOptionTypeID = 0x57d9d5a5

// Ensuring interfaces in compile-time for PollOption.
var (
	_ bin.Encoder     = &PollOption{}
	_ bin.Decoder     = &PollOption{}
	_ bin.BareEncoder = &PollOption{}
	_ bin.BareDecoder = &PollOption{}
)

func (p *PollOption) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Text == "") {
		return false
	}
	if !(p.VoterCount == 0) {
		return false
	}
	if !(p.VotePercentage == 0) {
		return false
	}
	if !(p.IsChosen == false) {
		return false
	}
	if !(p.IsBeingChosen == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PollOption) String() string {
	if p == nil {
		return "PollOption(nil)"
	}
	type Alias PollOption
	return fmt.Sprintf("PollOption%+v", Alias(*p))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PollOption) TypeID() uint32 {
	return PollOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*PollOption) TypeName() string {
	return "pollOption"
}

// TypeInfo returns info about TL type.
func (p *PollOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "pollOption",
		ID:   PollOptionTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Text",
			SchemaName: "text",
		},
		{
			Name:       "VoterCount",
			SchemaName: "voter_count",
		},
		{
			Name:       "VotePercentage",
			SchemaName: "vote_percentage",
		},
		{
			Name:       "IsChosen",
			SchemaName: "is_chosen",
		},
		{
			Name:       "IsBeingChosen",
			SchemaName: "is_being_chosen",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PollOption) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollOption#57d9d5a5 as nil")
	}
	b.PutID(PollOptionTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PollOption) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pollOption#57d9d5a5 as nil")
	}
	b.PutString(p.Text)
	b.PutInt32(p.VoterCount)
	b.PutInt32(p.VotePercentage)
	b.PutBool(p.IsChosen)
	b.PutBool(p.IsBeingChosen)
	return nil
}

// Decode implements bin.Decoder.
func (p *PollOption) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollOption#57d9d5a5 to nil")
	}
	if err := b.ConsumeID(PollOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode pollOption#57d9d5a5: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PollOption) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pollOption#57d9d5a5 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode pollOption#57d9d5a5: field text: %w", err)
		}
		p.Text = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode pollOption#57d9d5a5: field voter_count: %w", err)
		}
		p.VoterCount = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode pollOption#57d9d5a5: field vote_percentage: %w", err)
		}
		p.VotePercentage = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode pollOption#57d9d5a5: field is_chosen: %w", err)
		}
		p.IsChosen = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode pollOption#57d9d5a5: field is_being_chosen: %w", err)
		}
		p.IsBeingChosen = value
	}
	return nil
}

// EncodeTDLibJSON encodes p in TDLib API JSON format.
func (p *PollOption) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if p == nil {
		return fmt.Errorf("can't encode pollOption#57d9d5a5 as nil")
	}
	b.ObjStart()
	b.PutID("pollOption")
	b.FieldStart("text")
	b.PutString(p.Text)
	b.FieldStart("voter_count")
	b.PutInt32(p.VoterCount)
	b.FieldStart("vote_percentage")
	b.PutInt32(p.VotePercentage)
	b.FieldStart("is_chosen")
	b.PutBool(p.IsChosen)
	b.FieldStart("is_being_chosen")
	b.PutBool(p.IsBeingChosen)
	b.ObjEnd()
	return nil
}

// GetText returns value of Text field.
func (p *PollOption) GetText() (value string) {
	return p.Text
}

// GetVoterCount returns value of VoterCount field.
func (p *PollOption) GetVoterCount() (value int32) {
	return p.VoterCount
}

// GetVotePercentage returns value of VotePercentage field.
func (p *PollOption) GetVotePercentage() (value int32) {
	return p.VotePercentage
}

// GetIsChosen returns value of IsChosen field.
func (p *PollOption) GetIsChosen() (value bool) {
	return p.IsChosen
}

// GetIsBeingChosen returns value of IsBeingChosen field.
func (p *PollOption) GetIsBeingChosen() (value bool) {
	return p.IsBeingChosen
}
