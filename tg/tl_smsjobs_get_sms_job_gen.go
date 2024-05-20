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

// SMSJobsGetSMSJobRequest represents TL type `smsjobs.getSmsJob#778d902f`.
//
// See https://core.telegram.org/method/smsjobs.getSmsJob for reference.
type SMSJobsGetSMSJobRequest struct {
	// JobID field of SMSJobsGetSMSJobRequest.
	JobID string
}

// SMSJobsGetSMSJobRequestTypeID is TL type id of SMSJobsGetSMSJobRequest.
const SMSJobsGetSMSJobRequestTypeID = 0x778d902f

// Ensuring interfaces in compile-time for SMSJobsGetSMSJobRequest.
var (
	_ bin.Encoder     = &SMSJobsGetSMSJobRequest{}
	_ bin.Decoder     = &SMSJobsGetSMSJobRequest{}
	_ bin.BareEncoder = &SMSJobsGetSMSJobRequest{}
	_ bin.BareDecoder = &SMSJobsGetSMSJobRequest{}
)

func (g *SMSJobsGetSMSJobRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.JobID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *SMSJobsGetSMSJobRequest) String() string {
	if g == nil {
		return "SMSJobsGetSMSJobRequest(nil)"
	}
	type Alias SMSJobsGetSMSJobRequest
	return fmt.Sprintf("SMSJobsGetSMSJobRequest%+v", Alias(*g))
}

// FillFrom fills SMSJobsGetSMSJobRequest from given interface.
func (g *SMSJobsGetSMSJobRequest) FillFrom(from interface {
	GetJobID() (value string)
}) {
	g.JobID = from.GetJobID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SMSJobsGetSMSJobRequest) TypeID() uint32 {
	return SMSJobsGetSMSJobRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SMSJobsGetSMSJobRequest) TypeName() string {
	return "smsjobs.getSmsJob"
}

// TypeInfo returns info about TL type.
func (g *SMSJobsGetSMSJobRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "smsjobs.getSmsJob",
		ID:   SMSJobsGetSMSJobRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "JobID",
			SchemaName: "job_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *SMSJobsGetSMSJobRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode smsjobs.getSmsJob#778d902f as nil")
	}
	b.PutID(SMSJobsGetSMSJobRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *SMSJobsGetSMSJobRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode smsjobs.getSmsJob#778d902f as nil")
	}
	b.PutString(g.JobID)
	return nil
}

// Decode implements bin.Decoder.
func (g *SMSJobsGetSMSJobRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode smsjobs.getSmsJob#778d902f to nil")
	}
	if err := b.ConsumeID(SMSJobsGetSMSJobRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode smsjobs.getSmsJob#778d902f: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *SMSJobsGetSMSJobRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode smsjobs.getSmsJob#778d902f to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode smsjobs.getSmsJob#778d902f: field job_id: %w", err)
		}
		g.JobID = value
	}
	return nil
}

// GetJobID returns value of JobID field.
func (g *SMSJobsGetSMSJobRequest) GetJobID() (value string) {
	if g == nil {
		return
	}
	return g.JobID
}

// SMSJobsGetSMSJob invokes method smsjobs.getSmsJob#778d902f returning error if any.
//
// See https://core.telegram.org/method/smsjobs.getSmsJob for reference.
func (c *Client) SMSJobsGetSMSJob(ctx context.Context, jobid string) (*SMSJob, error) {
	var result SMSJob

	request := &SMSJobsGetSMSJobRequest{
		JobID: jobid,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
