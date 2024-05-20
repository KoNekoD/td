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

// AutoSaveException represents TL type `autoSaveException#81602d47`.
// Peer-specific media autosave settings
//
// See https://core.telegram.org/constructor/autoSaveException for reference.
type AutoSaveException struct {
	// The peer
	Peer PeerClass
	// Media autosave settings
	Settings AutoSaveSettings
}

// AutoSaveExceptionTypeID is TL type id of AutoSaveException.
const AutoSaveExceptionTypeID = 0x81602d47

// Ensuring interfaces in compile-time for AutoSaveException.
var (
	_ bin.Encoder     = &AutoSaveException{}
	_ bin.Decoder     = &AutoSaveException{}
	_ bin.BareEncoder = &AutoSaveException{}
	_ bin.BareDecoder = &AutoSaveException{}
)

func (a *AutoSaveException) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Peer == nil) {
		return false
	}
	if !(a.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *AutoSaveException) String() string {
	if a == nil {
		return "AutoSaveException(nil)"
	}
	type Alias AutoSaveException
	return fmt.Sprintf("AutoSaveException%+v", Alias(*a))
}

// FillFrom fills AutoSaveException from given interface.
func (a *AutoSaveException) FillFrom(from interface {
	GetPeer() (value PeerClass)
	GetSettings() (value AutoSaveSettings)
}) {
	a.Peer = from.GetPeer()
	a.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AutoSaveException) TypeID() uint32 {
	return AutoSaveExceptionTypeID
}

// TypeName returns name of type in TL schema.
func (*AutoSaveException) TypeName() string {
	return "autoSaveException"
}

// TypeInfo returns info about TL type.
func (a *AutoSaveException) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "autoSaveException",
		ID:   AutoSaveExceptionTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *AutoSaveException) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autoSaveException#81602d47 as nil")
	}
	b.PutID(AutoSaveExceptionTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *AutoSaveException) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode autoSaveException#81602d47 as nil")
	}
	if a.Peer == nil {
		return fmt.Errorf("unable to encode autoSaveException#81602d47: field peer is nil")
	}
	if err := a.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autoSaveException#81602d47: field peer: %w", err)
	}
	if err := a.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode autoSaveException#81602d47: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *AutoSaveException) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autoSaveException#81602d47 to nil")
	}
	if err := b.ConsumeID(AutoSaveExceptionTypeID); err != nil {
		return fmt.Errorf("unable to decode autoSaveException#81602d47: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *AutoSaveException) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode autoSaveException#81602d47 to nil")
	}
	{
		value, err := DecodePeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode autoSaveException#81602d47: field peer: %w", err)
		}
		a.Peer = value
	}
	{
		if err := a.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode autoSaveException#81602d47: field settings: %w", err)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (a *AutoSaveException) GetPeer() (value PeerClass) {
	if a == nil {
		return
	}
	return a.Peer
}

// GetSettings returns value of Settings field.
func (a *AutoSaveException) GetSettings() (value AutoSaveSettings) {
	if a == nil {
		return
	}
	return a.Settings
}
