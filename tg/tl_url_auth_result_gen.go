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

// URLAuthResultRequest represents TL type `urlAuthResultRequest#92d33a0e`.
// Details about the authorization request, for more info click here »¹
//
// Links:
//  1. https://core.telegram.org/api/url-authorization
//
// See https://core.telegram.org/constructor/urlAuthResultRequest for reference.
type URLAuthResultRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the bot would like to send messages to the user
	RequestWriteAccess bool
	// Username of a bot, which will be used for user authorization. If not specified, the
	// current bot's username will be assumed. The url's domain must be the same as the
	// domain linked with the bot. See Linking your domain to the bot¹ for more details.
	//
	// Links:
	//  1) https://core.telegram.org/widgets/login#linking-your-domain-to-the-bot
	Bot UserClass
	// The domain name of the website on which the user will log in.
	Domain string
}

// URLAuthResultRequestTypeID is TL type id of URLAuthResultRequest.
const URLAuthResultRequestTypeID = 0x92d33a0e

// construct implements constructor of URLAuthResultClass.
func (u URLAuthResultRequest) construct() URLAuthResultClass { return &u }

// Ensuring interfaces in compile-time for URLAuthResultRequest.
var (
	_ bin.Encoder     = &URLAuthResultRequest{}
	_ bin.Decoder     = &URLAuthResultRequest{}
	_ bin.BareEncoder = &URLAuthResultRequest{}
	_ bin.BareDecoder = &URLAuthResultRequest{}

	_ URLAuthResultClass = &URLAuthResultRequest{}
)

func (u *URLAuthResultRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.RequestWriteAccess == false) {
		return false
	}
	if !(u.Bot == nil) {
		return false
	}
	if !(u.Domain == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *URLAuthResultRequest) String() string {
	if u == nil {
		return "URLAuthResultRequest(nil)"
	}
	type Alias URLAuthResultRequest
	return fmt.Sprintf("URLAuthResultRequest%+v", Alias(*u))
}

// FillFrom fills URLAuthResultRequest from given interface.
func (u *URLAuthResultRequest) FillFrom(from interface {
	GetRequestWriteAccess() (value bool)
	GetBot() (value UserClass)
	GetDomain() (value string)
}) {
	u.RequestWriteAccess = from.GetRequestWriteAccess()
	u.Bot = from.GetBot()
	u.Domain = from.GetDomain()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*URLAuthResultRequest) TypeID() uint32 {
	return URLAuthResultRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*URLAuthResultRequest) TypeName() string {
	return "urlAuthResultRequest"
}

// TypeInfo returns info about TL type.
func (u *URLAuthResultRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "urlAuthResultRequest",
		ID:   URLAuthResultRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "RequestWriteAccess",
			SchemaName: "request_write_access",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "Bot",
			SchemaName: "bot",
		},
		{
			Name:       "Domain",
			SchemaName: "domain",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *URLAuthResultRequest) SetFlags() {
	if !(u.RequestWriteAccess == false) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *URLAuthResultRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode urlAuthResultRequest#92d33a0e as nil")
	}
	b.PutID(URLAuthResultRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *URLAuthResultRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode urlAuthResultRequest#92d33a0e as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode urlAuthResultRequest#92d33a0e: field flags: %w", err)
	}
	if u.Bot == nil {
		return fmt.Errorf("unable to encode urlAuthResultRequest#92d33a0e: field bot is nil")
	}
	if err := u.Bot.Encode(b); err != nil {
		return fmt.Errorf("unable to encode urlAuthResultRequest#92d33a0e: field bot: %w", err)
	}
	b.PutString(u.Domain)
	return nil
}

// Decode implements bin.Decoder.
func (u *URLAuthResultRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode urlAuthResultRequest#92d33a0e to nil")
	}
	if err := b.ConsumeID(URLAuthResultRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode urlAuthResultRequest#92d33a0e: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *URLAuthResultRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode urlAuthResultRequest#92d33a0e to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode urlAuthResultRequest#92d33a0e: field flags: %w", err)
		}
	}
	u.RequestWriteAccess = u.Flags.Has(0)
	{
		value, err := DecodeUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode urlAuthResultRequest#92d33a0e: field bot: %w", err)
		}
		u.Bot = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode urlAuthResultRequest#92d33a0e: field domain: %w", err)
		}
		u.Domain = value
	}
	return nil
}

// SetRequestWriteAccess sets value of RequestWriteAccess conditional field.
func (u *URLAuthResultRequest) SetRequestWriteAccess(value bool) {
	if value {
		u.Flags.Set(0)
		u.RequestWriteAccess = true
	} else {
		u.Flags.Unset(0)
		u.RequestWriteAccess = false
	}
}

// GetRequestWriteAccess returns value of RequestWriteAccess conditional field.
func (u *URLAuthResultRequest) GetRequestWriteAccess() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// GetBot returns value of Bot field.
func (u *URLAuthResultRequest) GetBot() (value UserClass) {
	if u == nil {
		return
	}
	return u.Bot
}

// GetDomain returns value of Domain field.
func (u *URLAuthResultRequest) GetDomain() (value string) {
	if u == nil {
		return
	}
	return u.Domain
}

// URLAuthResultAccepted represents TL type `urlAuthResultAccepted#8f8c0e4e`.
// Details about an accepted authorization request, for more info click here »¹
//
// Links:
//  1. https://core.telegram.org/api/url-authorization
//
// See https://core.telegram.org/constructor/urlAuthResultAccepted for reference.
type URLAuthResultAccepted struct {
	// The URL name of the website on which the user has logged in.
	URL string
}

// URLAuthResultAcceptedTypeID is TL type id of URLAuthResultAccepted.
const URLAuthResultAcceptedTypeID = 0x8f8c0e4e

// construct implements constructor of URLAuthResultClass.
func (u URLAuthResultAccepted) construct() URLAuthResultClass { return &u }

// Ensuring interfaces in compile-time for URLAuthResultAccepted.
var (
	_ bin.Encoder     = &URLAuthResultAccepted{}
	_ bin.Decoder     = &URLAuthResultAccepted{}
	_ bin.BareEncoder = &URLAuthResultAccepted{}
	_ bin.BareDecoder = &URLAuthResultAccepted{}

	_ URLAuthResultClass = &URLAuthResultAccepted{}
)

func (u *URLAuthResultAccepted) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *URLAuthResultAccepted) String() string {
	if u == nil {
		return "URLAuthResultAccepted(nil)"
	}
	type Alias URLAuthResultAccepted
	return fmt.Sprintf("URLAuthResultAccepted%+v", Alias(*u))
}

// FillFrom fills URLAuthResultAccepted from given interface.
func (u *URLAuthResultAccepted) FillFrom(from interface {
	GetURL() (value string)
}) {
	u.URL = from.GetURL()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*URLAuthResultAccepted) TypeID() uint32 {
	return URLAuthResultAcceptedTypeID
}

// TypeName returns name of type in TL schema.
func (*URLAuthResultAccepted) TypeName() string {
	return "urlAuthResultAccepted"
}

// TypeInfo returns info about TL type.
func (u *URLAuthResultAccepted) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "urlAuthResultAccepted",
		ID:   URLAuthResultAcceptedTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "URL",
			SchemaName: "url",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *URLAuthResultAccepted) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode urlAuthResultAccepted#8f8c0e4e as nil")
	}
	b.PutID(URLAuthResultAcceptedTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *URLAuthResultAccepted) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode urlAuthResultAccepted#8f8c0e4e as nil")
	}
	b.PutString(u.URL)
	return nil
}

// Decode implements bin.Decoder.
func (u *URLAuthResultAccepted) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode urlAuthResultAccepted#8f8c0e4e to nil")
	}
	if err := b.ConsumeID(URLAuthResultAcceptedTypeID); err != nil {
		return fmt.Errorf("unable to decode urlAuthResultAccepted#8f8c0e4e: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *URLAuthResultAccepted) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode urlAuthResultAccepted#8f8c0e4e to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode urlAuthResultAccepted#8f8c0e4e: field url: %w", err)
		}
		u.URL = value
	}
	return nil
}

// GetURL returns value of URL field.
func (u *URLAuthResultAccepted) GetURL() (value string) {
	if u == nil {
		return
	}
	return u.URL
}

// URLAuthResultDefault represents TL type `urlAuthResultDefault#a9d6db1f`.
// Details about an accepted authorization request, for more info click here »¹
//
// Links:
//  1. https://core.telegram.org/api/url-authorization
//
// See https://core.telegram.org/constructor/urlAuthResultDefault for reference.
type URLAuthResultDefault struct {
}

// URLAuthResultDefaultTypeID is TL type id of URLAuthResultDefault.
const URLAuthResultDefaultTypeID = 0xa9d6db1f

// construct implements constructor of URLAuthResultClass.
func (u URLAuthResultDefault) construct() URLAuthResultClass { return &u }

// Ensuring interfaces in compile-time for URLAuthResultDefault.
var (
	_ bin.Encoder     = &URLAuthResultDefault{}
	_ bin.Decoder     = &URLAuthResultDefault{}
	_ bin.BareEncoder = &URLAuthResultDefault{}
	_ bin.BareDecoder = &URLAuthResultDefault{}

	_ URLAuthResultClass = &URLAuthResultDefault{}
)

func (u *URLAuthResultDefault) Zero() bool {
	if u == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (u *URLAuthResultDefault) String() string {
	if u == nil {
		return "URLAuthResultDefault(nil)"
	}
	type Alias URLAuthResultDefault
	return fmt.Sprintf("URLAuthResultDefault%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*URLAuthResultDefault) TypeID() uint32 {
	return URLAuthResultDefaultTypeID
}

// TypeName returns name of type in TL schema.
func (*URLAuthResultDefault) TypeName() string {
	return "urlAuthResultDefault"
}

// TypeInfo returns info about TL type.
func (u *URLAuthResultDefault) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "urlAuthResultDefault",
		ID:   URLAuthResultDefaultTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (u *URLAuthResultDefault) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode urlAuthResultDefault#a9d6db1f as nil")
	}
	b.PutID(URLAuthResultDefaultTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *URLAuthResultDefault) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode urlAuthResultDefault#a9d6db1f as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *URLAuthResultDefault) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode urlAuthResultDefault#a9d6db1f to nil")
	}
	if err := b.ConsumeID(URLAuthResultDefaultTypeID); err != nil {
		return fmt.Errorf("unable to decode urlAuthResultDefault#a9d6db1f: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *URLAuthResultDefault) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode urlAuthResultDefault#a9d6db1f to nil")
	}
	return nil
}

// URLAuthResultClassName is schema name of URLAuthResultClass.
const URLAuthResultClassName = "UrlAuthResult"

// URLAuthResultClass represents UrlAuthResult generic type.
//
// See https://core.telegram.org/type/UrlAuthResult for reference.
//
// Example:
//
//	g, err := tg.DecodeURLAuthResult(buf)
//	if err != nil {
//	    panic(err)
//	}
//	switch v := g.(type) {
//	case *tg.URLAuthResultRequest: // urlAuthResultRequest#92d33a0e
//	case *tg.URLAuthResultAccepted: // urlAuthResultAccepted#8f8c0e4e
//	case *tg.URLAuthResultDefault: // urlAuthResultDefault#a9d6db1f
//	default: panic(v)
//	}
type URLAuthResultClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() URLAuthResultClass

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
}

// DecodeURLAuthResult implements binary de-serialization for URLAuthResultClass.
func DecodeURLAuthResult(buf *bin.Buffer) (URLAuthResultClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case URLAuthResultRequestTypeID:
		// Decoding urlAuthResultRequest#92d33a0e.
		v := URLAuthResultRequest{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode URLAuthResultClass: %w", err)
		}
		return &v, nil
	case URLAuthResultAcceptedTypeID:
		// Decoding urlAuthResultAccepted#8f8c0e4e.
		v := URLAuthResultAccepted{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode URLAuthResultClass: %w", err)
		}
		return &v, nil
	case URLAuthResultDefaultTypeID:
		// Decoding urlAuthResultDefault#a9d6db1f.
		v := URLAuthResultDefault{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode URLAuthResultClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode URLAuthResultClass: %w", bin.NewUnexpectedID(id))
	}
}

// URLAuthResult boxes the URLAuthResultClass providing a helper.
type URLAuthResultBox struct {
	UrlAuthResult URLAuthResultClass
}

// Decode implements bin.Decoder for URLAuthResultBox.
func (b *URLAuthResultBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode URLAuthResultBox to nil")
	}
	v, err := DecodeURLAuthResult(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.UrlAuthResult = v
	return nil
}

// Encode implements bin.Encode for URLAuthResultBox.
func (b *URLAuthResultBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.UrlAuthResult == nil {
		return fmt.Errorf("unable to encode URLAuthResultClass as nil")
	}
	return b.UrlAuthResult.Encode(buf)
}
