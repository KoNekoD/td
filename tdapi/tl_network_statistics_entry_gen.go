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

// NetworkStatisticsEntryFile represents TL type `networkStatisticsEntryFile#b3b8f62`.
type NetworkStatisticsEntryFile struct {
	// Type of the file the data is part of
	FileType FileTypeClass
	// Type of the network the data was sent through. Call setNetworkType to maintain the
	// actual network type
	NetworkType NetworkTypeClass
	// Total number of bytes sent
	SentBytes int64
	// Total number of bytes received
	ReceivedBytes int64
}

// NetworkStatisticsEntryFileTypeID is TL type id of NetworkStatisticsEntryFile.
const NetworkStatisticsEntryFileTypeID = 0xb3b8f62

// construct implements constructor of NetworkStatisticsEntryClass.
func (n NetworkStatisticsEntryFile) construct() NetworkStatisticsEntryClass { return &n }

// Ensuring interfaces in compile-time for NetworkStatisticsEntryFile.
var (
	_ bin.Encoder     = &NetworkStatisticsEntryFile{}
	_ bin.Decoder     = &NetworkStatisticsEntryFile{}
	_ bin.BareEncoder = &NetworkStatisticsEntryFile{}
	_ bin.BareDecoder = &NetworkStatisticsEntryFile{}

	_ NetworkStatisticsEntryClass = &NetworkStatisticsEntryFile{}
)

func (n *NetworkStatisticsEntryFile) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.FileType == nil) {
		return false
	}
	if !(n.NetworkType == nil) {
		return false
	}
	if !(n.SentBytes == 0) {
		return false
	}
	if !(n.ReceivedBytes == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NetworkStatisticsEntryFile) String() string {
	if n == nil {
		return "NetworkStatisticsEntryFile(nil)"
	}
	type Alias NetworkStatisticsEntryFile
	return fmt.Sprintf("NetworkStatisticsEntryFile%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NetworkStatisticsEntryFile) TypeID() uint32 {
	return NetworkStatisticsEntryFileTypeID
}

// TypeName returns name of type in TL schema.
func (*NetworkStatisticsEntryFile) TypeName() string {
	return "networkStatisticsEntryFile"
}

// TypeInfo returns info about TL type.
func (n *NetworkStatisticsEntryFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "networkStatisticsEntryFile",
		ID:   NetworkStatisticsEntryFileTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "FileType",
			SchemaName: "file_type",
		},
		{
			Name:       "NetworkType",
			SchemaName: "network_type",
		},
		{
			Name:       "SentBytes",
			SchemaName: "sent_bytes",
		},
		{
			Name:       "ReceivedBytes",
			SchemaName: "received_bytes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NetworkStatisticsEntryFile) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryFile#b3b8f62 as nil")
	}
	b.PutID(NetworkStatisticsEntryFileTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NetworkStatisticsEntryFile) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryFile#b3b8f62 as nil")
	}
	if n.FileType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field file_type is nil")
	}
	if err := n.FileType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field file_type: %w", err)
	}
	if n.NetworkType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field network_type is nil")
	}
	if err := n.NetworkType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field network_type: %w", err)
	}
	b.PutLong(n.SentBytes)
	b.PutLong(n.ReceivedBytes)
	return nil
}

// Decode implements bin.Decoder.
func (n *NetworkStatisticsEntryFile) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatisticsEntryFile#b3b8f62 to nil")
	}
	if err := b.ConsumeID(NetworkStatisticsEntryFileTypeID); err != nil {
		return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NetworkStatisticsEntryFile) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatisticsEntryFile#b3b8f62 to nil")
	}
	{
		value, err := DecodeFileType(b)
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field file_type: %w", err)
		}
		n.FileType = value
	}
	{
		value, err := DecodeNetworkType(b)
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field network_type: %w", err)
		}
		n.NetworkType = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field sent_bytes: %w", err)
		}
		n.SentBytes = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryFile#b3b8f62: field received_bytes: %w", err)
		}
		n.ReceivedBytes = value
	}
	return nil
}

// EncodeTDLibJSON encodes n in TDLib API JSON format.
func (n *NetworkStatisticsEntryFile) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryFile#b3b8f62 as nil")
	}
	b.ObjStart()
	b.PutID("networkStatisticsEntryFile")
	b.FieldStart("file_type")
	if n.FileType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field file_type is nil")
	}
	if err := n.FileType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field file_type: %w", err)
	}
	b.FieldStart("network_type")
	if n.NetworkType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field network_type is nil")
	}
	if err := n.NetworkType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryFile#b3b8f62: field network_type: %w", err)
	}
	b.FieldStart("sent_bytes")
	b.PutLong(n.SentBytes)
	b.FieldStart("received_bytes")
	b.PutLong(n.ReceivedBytes)
	b.ObjEnd()
	return nil
}

// GetFileType returns value of FileType field.
func (n *NetworkStatisticsEntryFile) GetFileType() (value FileTypeClass) {
	return n.FileType
}

// GetNetworkType returns value of NetworkType field.
func (n *NetworkStatisticsEntryFile) GetNetworkType() (value NetworkTypeClass) {
	return n.NetworkType
}

// GetSentBytes returns value of SentBytes field.
func (n *NetworkStatisticsEntryFile) GetSentBytes() (value int64) {
	return n.SentBytes
}

// GetReceivedBytes returns value of ReceivedBytes field.
func (n *NetworkStatisticsEntryFile) GetReceivedBytes() (value int64) {
	return n.ReceivedBytes
}

// NetworkStatisticsEntryCall represents TL type `networkStatisticsEntryCall#2bedbbad`.
type NetworkStatisticsEntryCall struct {
	// Type of the network the data was sent through. Call setNetworkType to maintain the
	// actual network type
	NetworkType NetworkTypeClass
	// Total number of bytes sent
	SentBytes int64
	// Total number of bytes received
	ReceivedBytes int64
	// Total call duration, in seconds
	Duration float64
}

// NetworkStatisticsEntryCallTypeID is TL type id of NetworkStatisticsEntryCall.
const NetworkStatisticsEntryCallTypeID = 0x2bedbbad

// construct implements constructor of NetworkStatisticsEntryClass.
func (n NetworkStatisticsEntryCall) construct() NetworkStatisticsEntryClass { return &n }

// Ensuring interfaces in compile-time for NetworkStatisticsEntryCall.
var (
	_ bin.Encoder     = &NetworkStatisticsEntryCall{}
	_ bin.Decoder     = &NetworkStatisticsEntryCall{}
	_ bin.BareEncoder = &NetworkStatisticsEntryCall{}
	_ bin.BareDecoder = &NetworkStatisticsEntryCall{}

	_ NetworkStatisticsEntryClass = &NetworkStatisticsEntryCall{}
)

func (n *NetworkStatisticsEntryCall) Zero() bool {
	if n == nil {
		return true
	}
	if !(n.NetworkType == nil) {
		return false
	}
	if !(n.SentBytes == 0) {
		return false
	}
	if !(n.ReceivedBytes == 0) {
		return false
	}
	if !(n.Duration == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (n *NetworkStatisticsEntryCall) String() string {
	if n == nil {
		return "NetworkStatisticsEntryCall(nil)"
	}
	type Alias NetworkStatisticsEntryCall
	return fmt.Sprintf("NetworkStatisticsEntryCall%+v", Alias(*n))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*NetworkStatisticsEntryCall) TypeID() uint32 {
	return NetworkStatisticsEntryCallTypeID
}

// TypeName returns name of type in TL schema.
func (*NetworkStatisticsEntryCall) TypeName() string {
	return "networkStatisticsEntryCall"
}

// TypeInfo returns info about TL type.
func (n *NetworkStatisticsEntryCall) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "networkStatisticsEntryCall",
		ID:   NetworkStatisticsEntryCallTypeID,
	}
	if n == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "NetworkType",
			SchemaName: "network_type",
		},
		{
			Name:       "SentBytes",
			SchemaName: "sent_bytes",
		},
		{
			Name:       "ReceivedBytes",
			SchemaName: "received_bytes",
		},
		{
			Name:       "Duration",
			SchemaName: "duration",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (n *NetworkStatisticsEntryCall) Encode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryCall#2bedbbad as nil")
	}
	b.PutID(NetworkStatisticsEntryCallTypeID)
	return n.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (n *NetworkStatisticsEntryCall) EncodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryCall#2bedbbad as nil")
	}
	if n.NetworkType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryCall#2bedbbad: field network_type is nil")
	}
	if err := n.NetworkType.Encode(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryCall#2bedbbad: field network_type: %w", err)
	}
	b.PutLong(n.SentBytes)
	b.PutLong(n.ReceivedBytes)
	b.PutDouble(n.Duration)
	return nil
}

// Decode implements bin.Decoder.
func (n *NetworkStatisticsEntryCall) Decode(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatisticsEntryCall#2bedbbad to nil")
	}
	if err := b.ConsumeID(NetworkStatisticsEntryCallTypeID); err != nil {
		return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: %w", err)
	}
	return n.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (n *NetworkStatisticsEntryCall) DecodeBare(b *bin.Buffer) error {
	if n == nil {
		return fmt.Errorf("can't decode networkStatisticsEntryCall#2bedbbad to nil")
	}
	{
		value, err := DecodeNetworkType(b)
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field network_type: %w", err)
		}
		n.NetworkType = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field sent_bytes: %w", err)
		}
		n.SentBytes = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field received_bytes: %w", err)
		}
		n.ReceivedBytes = value
	}
	{
		value, err := b.Double()
		if err != nil {
			return fmt.Errorf("unable to decode networkStatisticsEntryCall#2bedbbad: field duration: %w", err)
		}
		n.Duration = value
	}
	return nil
}

// EncodeTDLibJSON encodes n in TDLib API JSON format.
func (n *NetworkStatisticsEntryCall) EncodeTDLibJSON(b *jsontd.Encoder) error {
	if n == nil {
		return fmt.Errorf("can't encode networkStatisticsEntryCall#2bedbbad as nil")
	}
	b.ObjStart()
	b.PutID("networkStatisticsEntryCall")
	b.FieldStart("network_type")
	if n.NetworkType == nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryCall#2bedbbad: field network_type is nil")
	}
	if err := n.NetworkType.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode networkStatisticsEntryCall#2bedbbad: field network_type: %w", err)
	}
	b.FieldStart("sent_bytes")
	b.PutLong(n.SentBytes)
	b.FieldStart("received_bytes")
	b.PutLong(n.ReceivedBytes)
	b.FieldStart("duration")
	b.PutDouble(n.Duration)
	b.ObjEnd()
	return nil
}

// GetNetworkType returns value of NetworkType field.
func (n *NetworkStatisticsEntryCall) GetNetworkType() (value NetworkTypeClass) {
	return n.NetworkType
}

// GetSentBytes returns value of SentBytes field.
func (n *NetworkStatisticsEntryCall) GetSentBytes() (value int64) {
	return n.SentBytes
}

// GetReceivedBytes returns value of ReceivedBytes field.
func (n *NetworkStatisticsEntryCall) GetReceivedBytes() (value int64) {
	return n.ReceivedBytes
}

// GetDuration returns value of Duration field.
func (n *NetworkStatisticsEntryCall) GetDuration() (value float64) {
	return n.Duration
}

// NetworkStatisticsEntryClass represents NetworkStatisticsEntry generic type.
//
// Example:
//  g, err := tdapi.DecodeNetworkStatisticsEntry(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tdapi.NetworkStatisticsEntryFile: // networkStatisticsEntryFile#b3b8f62
//  case *tdapi.NetworkStatisticsEntryCall: // networkStatisticsEntryCall#2bedbbad
//  default: panic(v)
//  }
type NetworkStatisticsEntryClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() NetworkStatisticsEntryClass

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
	EncodeTDLibJSON(b *jsontd.Encoder) error

	// Type of the network the data was sent through. Call setNetworkType to maintain the
	// actual network type
	GetNetworkType() (value NetworkTypeClass)
	// Total number of bytes sent
	GetSentBytes() (value int64)
	// Total number of bytes received
	GetReceivedBytes() (value int64)
}

// DecodeNetworkStatisticsEntry implements binary de-serialization for NetworkStatisticsEntryClass.
func DecodeNetworkStatisticsEntry(buf *bin.Buffer) (NetworkStatisticsEntryClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case NetworkStatisticsEntryFileTypeID:
		// Decoding networkStatisticsEntryFile#b3b8f62.
		v := NetworkStatisticsEntryFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkStatisticsEntryClass: %w", err)
		}
		return &v, nil
	case NetworkStatisticsEntryCallTypeID:
		// Decoding networkStatisticsEntryCall#2bedbbad.
		v := NetworkStatisticsEntryCall{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode NetworkStatisticsEntryClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode NetworkStatisticsEntryClass: %w", bin.NewUnexpectedID(id))
	}
}

// NetworkStatisticsEntry boxes the NetworkStatisticsEntryClass providing a helper.
type NetworkStatisticsEntryBox struct {
	NetworkStatisticsEntry NetworkStatisticsEntryClass
}

// Decode implements bin.Decoder for NetworkStatisticsEntryBox.
func (b *NetworkStatisticsEntryBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode NetworkStatisticsEntryBox to nil")
	}
	v, err := DecodeNetworkStatisticsEntry(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.NetworkStatisticsEntry = v
	return nil
}

// Encode implements bin.Encode for NetworkStatisticsEntryBox.
func (b *NetworkStatisticsEntryBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.NetworkStatisticsEntry == nil {
		return fmt.Errorf("unable to encode NetworkStatisticsEntryClass as nil")
	}
	return b.NetworkStatisticsEntry.Encode(buf)
}
