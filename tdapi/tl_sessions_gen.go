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

// Sessions represents TL type `sessions#997833aa`.
type Sessions struct {
	// List of sessions
	Sessions []Session
	// Number of days of inactivity before sessions will automatically be terminated; 1-366
	// days
	InactiveSessionTTLDays int32
}

// SessionsTypeID is TL type id of Sessions.
const SessionsTypeID = 0x997833aa

// Ensuring interfaces in compile-time for Sessions.
var (
	_ bin.Encoder     = &Sessions{}
	_ bin.Decoder     = &Sessions{}
	_ bin.BareEncoder = &Sessions{}
	_ bin.BareDecoder = &Sessions{}
)

func (s *Sessions) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Sessions == nil) {
		return false
	}
	if !(s.InactiveSessionTTLDays == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *Sessions) String() string {
	if s == nil {
		return "Sessions(nil)"
	}
	type Alias Sessions
	return fmt.Sprintf("Sessions%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Sessions) TypeID() uint32 {
	return SessionsTypeID
}

// TypeName returns name of type in TL schema.
func (*Sessions) TypeName() string {
	return "sessions"
}

// TypeInfo returns info about TL type.
func (s *Sessions) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "sessions",
		ID:   SessionsTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Sessions",
			SchemaName: "sessions",
		},
		{
			Name:       "InactiveSessionTTLDays",
			SchemaName: "inactive_session_ttl_days",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *Sessions) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sessions#997833aa as nil")
	}
	b.PutID(SessionsTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *Sessions) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode sessions#997833aa as nil")
	}
	b.PutInt(len(s.Sessions))
	for idx, v := range s.Sessions {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare sessions#997833aa: field sessions element with index %d: %w", idx, err)
		}
	}
	b.PutInt32(s.InactiveSessionTTLDays)
	return nil
}

// Decode implements bin.Decoder.
func (s *Sessions) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sessions#997833aa to nil")
	}
	if err := b.ConsumeID(SessionsTypeID); err != nil {
		return fmt.Errorf("unable to decode sessions#997833aa: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *Sessions) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode sessions#997833aa to nil")
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode sessions#997833aa: field sessions: %w", err)
		}

		if headerLen > 0 {
			s.Sessions = make([]Session, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value Session
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare sessions#997833aa: field sessions: %w", err)
			}
			s.Sessions = append(s.Sessions, value)
		}
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode sessions#997833aa: field inactive_session_ttl_days: %w", err)
		}
		s.InactiveSessionTTLDays = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *Sessions) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode sessions#997833aa as nil")
	}
	b.ObjStart()
	b.PutID("sessions")
	b.Comma()
	b.FieldStart("sessions")
	b.ArrStart()
	for idx, v := range s.Sessions {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode sessions#997833aa: field sessions element with index %d: %w", idx, err)
		}
		b.Comma()
	}
	b.StripComma()
	b.ArrEnd()
	b.Comma()
	b.FieldStart("inactive_session_ttl_days")
	b.PutInt32(s.InactiveSessionTTLDays)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *Sessions) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode sessions#997833aa to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("sessions"); err != nil {
				return fmt.Errorf("unable to decode sessions#997833aa: %w", err)
			}
		case "sessions":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value Session
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode sessions#997833aa: field sessions: %w", err)
				}
				s.Sessions = append(s.Sessions, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode sessions#997833aa: field sessions: %w", err)
			}
		case "inactive_session_ttl_days":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode sessions#997833aa: field inactive_session_ttl_days: %w", err)
			}
			s.InactiveSessionTTLDays = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetSessions returns value of Sessions field.
func (s *Sessions) GetSessions() (value []Session) {
	if s == nil {
		return
	}
	return s.Sessions
}

// GetInactiveSessionTTLDays returns value of InactiveSessionTTLDays field.
func (s *Sessions) GetInactiveSessionTTLDays() (value int32) {
	if s == nil {
		return
	}
	return s.InactiveSessionTTLDays
}
