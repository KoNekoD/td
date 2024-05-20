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

// ExportedChatlistInvite represents TL type `exportedChatlistInvite#c5181ac`.
// Exported chat folder deep link »¹.
//
// Links:
//  1. https://core.telegram.org/api/links#chat-folder-links
//
// See https://core.telegram.org/constructor/exportedChatlistInvite for reference.
type ExportedChatlistInvite struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Name of the link
	Title string
	// The chat folder deep link »¹.
	//
	// Links:
	//  1) https://core.telegram.org/api/links#chat-folder-links
	URL string
	// Peers to import
	Peers []PeerClass
}

// ExportedChatlistInviteTypeID is TL type id of ExportedChatlistInvite.
const ExportedChatlistInviteTypeID = 0xc5181ac

// Ensuring interfaces in compile-time for ExportedChatlistInvite.
var (
	_ bin.Encoder     = &ExportedChatlistInvite{}
	_ bin.Decoder     = &ExportedChatlistInvite{}
	_ bin.BareEncoder = &ExportedChatlistInvite{}
	_ bin.BareDecoder = &ExportedChatlistInvite{}
)

func (e *ExportedChatlistInvite) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Flags.Zero()) {
		return false
	}
	if !(e.Title == "") {
		return false
	}
	if !(e.URL == "") {
		return false
	}
	if !(e.Peers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *ExportedChatlistInvite) String() string {
	if e == nil {
		return "ExportedChatlistInvite(nil)"
	}
	type Alias ExportedChatlistInvite
	return fmt.Sprintf("ExportedChatlistInvite%+v", Alias(*e))
}

// FillFrom fills ExportedChatlistInvite from given interface.
func (e *ExportedChatlistInvite) FillFrom(from interface {
	GetTitle() (value string)
	GetURL() (value string)
	GetPeers() (value []PeerClass)
}) {
	e.Title = from.GetTitle()
	e.URL = from.GetURL()
	e.Peers = from.GetPeers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ExportedChatlistInvite) TypeID() uint32 {
	return ExportedChatlistInviteTypeID
}

// TypeName returns name of type in TL schema.
func (*ExportedChatlistInvite) TypeName() string {
	return "exportedChatlistInvite"
}

// TypeInfo returns info about TL type.
func (e *ExportedChatlistInvite) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "exportedChatlistInvite",
		ID:   ExportedChatlistInviteTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Title",
			SchemaName: "title",
		},
		{
			Name:       "URL",
			SchemaName: "url",
		},
		{
			Name:       "Peers",
			SchemaName: "peers",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (e *ExportedChatlistInvite) SetFlags() {
}

// Encode implements bin.Encoder.
func (e *ExportedChatlistInvite) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode exportedChatlistInvite#c5181ac as nil")
	}
	b.PutID(ExportedChatlistInviteTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *ExportedChatlistInvite) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode exportedChatlistInvite#c5181ac as nil")
	}
	e.SetFlags()
	if err := e.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode exportedChatlistInvite#c5181ac: field flags: %w", err)
	}
	b.PutString(e.Title)
	b.PutString(e.URL)
	b.PutVectorHeader(len(e.Peers))
	for idx, v := range e.Peers {
		if v == nil {
			return fmt.Errorf("unable to encode exportedChatlistInvite#c5181ac: field peers element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode exportedChatlistInvite#c5181ac: field peers element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *ExportedChatlistInvite) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode exportedChatlistInvite#c5181ac to nil")
	}
	if err := b.ConsumeID(ExportedChatlistInviteTypeID); err != nil {
		return fmt.Errorf("unable to decode exportedChatlistInvite#c5181ac: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *ExportedChatlistInvite) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode exportedChatlistInvite#c5181ac to nil")
	}
	{
		if err := e.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode exportedChatlistInvite#c5181ac: field flags: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode exportedChatlistInvite#c5181ac: field title: %w", err)
		}
		e.Title = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode exportedChatlistInvite#c5181ac: field url: %w", err)
		}
		e.URL = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode exportedChatlistInvite#c5181ac: field peers: %w", err)
		}

		if headerLen > 0 {
			e.Peers = make([]PeerClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePeer(b)
			if err != nil {
				return fmt.Errorf("unable to decode exportedChatlistInvite#c5181ac: field peers: %w", err)
			}
			e.Peers = append(e.Peers, value)
		}
	}
	return nil
}

// GetTitle returns value of Title field.
func (e *ExportedChatlistInvite) GetTitle() (value string) {
	if e == nil {
		return
	}
	return e.Title
}

// GetURL returns value of URL field.
func (e *ExportedChatlistInvite) GetURL() (value string) {
	if e == nil {
		return
	}
	return e.URL
}

// GetPeers returns value of Peers field.
func (e *ExportedChatlistInvite) GetPeers() (value []PeerClass) {
	if e == nil {
		return
	}
	return e.Peers
}

// MapPeers returns field Peers wrapped in PeerClassArray helper.
func (e *ExportedChatlistInvite) MapPeers() (value PeerClassArray) {
	return PeerClassArray(e.Peers)
}
