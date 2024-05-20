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

// Authorization represents TL type `authorization#ad01d61d`.
// Logged-in session
//
// See https://core.telegram.org/constructor/authorization for reference.
type Authorization struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether this is the current session
	Current bool
	// Whether the session is from an official app
	OfficialApp bool
	// Whether the session is still waiting for a 2FA password
	PasswordPending bool
	// Whether this session will accept encrypted chats
	EncryptedRequestsDisabled bool
	// Whether this session will accept phone calls
	CallRequestsDisabled bool
	// Whether the session is unconfirmed, see here »¹ for more info.
	//
	// Links:
	//  1) https://core.telegram.org/api/auth#confirming-login
	Unconfirmed bool
	// Identifier
	Hash int64
	// Device model
	DeviceModel string
	// Platform
	Platform string
	// System version
	SystemVersion string
	// API ID¹
	//
	// Links:
	//  1) https://core.telegram.org/api/obtaining_api_id
	APIID int
	// App name
	AppName string
	// App version
	AppVersion string
	// When was the session created
	DateCreated int
	// When was the session last active
	DateActive int
	// Last known IP
	IP string
	// Country determined from IP
	Country string
	// Region determined from IP
	Region string
}

// AuthorizationTypeID is TL type id of Authorization.
const AuthorizationTypeID = 0xad01d61d

// Ensuring interfaces in compile-time for Authorization.
var (
	_ bin.Encoder     = &Authorization{}
	_ bin.Decoder     = &Authorization{}
	_ bin.BareEncoder = &Authorization{}
	_ bin.BareDecoder = &Authorization{}
)

func (a *Authorization) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Flags.Zero()) {
		return false
	}
	if !(a.Current == false) {
		return false
	}
	if !(a.OfficialApp == false) {
		return false
	}
	if !(a.PasswordPending == false) {
		return false
	}
	if !(a.EncryptedRequestsDisabled == false) {
		return false
	}
	if !(a.CallRequestsDisabled == false) {
		return false
	}
	if !(a.Unconfirmed == false) {
		return false
	}
	if !(a.Hash == 0) {
		return false
	}
	if !(a.DeviceModel == "") {
		return false
	}
	if !(a.Platform == "") {
		return false
	}
	if !(a.SystemVersion == "") {
		return false
	}
	if !(a.APIID == 0) {
		return false
	}
	if !(a.AppName == "") {
		return false
	}
	if !(a.AppVersion == "") {
		return false
	}
	if !(a.DateCreated == 0) {
		return false
	}
	if !(a.DateActive == 0) {
		return false
	}
	if !(a.IP == "") {
		return false
	}
	if !(a.Country == "") {
		return false
	}
	if !(a.Region == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *Authorization) String() string {
	if a == nil {
		return "Authorization(nil)"
	}
	type Alias Authorization
	return fmt.Sprintf("Authorization%+v", Alias(*a))
}

// FillFrom fills Authorization from given interface.
func (a *Authorization) FillFrom(from interface {
	GetCurrent() (value bool)
	GetOfficialApp() (value bool)
	GetPasswordPending() (value bool)
	GetEncryptedRequestsDisabled() (value bool)
	GetCallRequestsDisabled() (value bool)
	GetUnconfirmed() (value bool)
	GetHash() (value int64)
	GetDeviceModel() (value string)
	GetPlatform() (value string)
	GetSystemVersion() (value string)
	GetAPIID() (value int)
	GetAppName() (value string)
	GetAppVersion() (value string)
	GetDateCreated() (value int)
	GetDateActive() (value int)
	GetIP() (value string)
	GetCountry() (value string)
	GetRegion() (value string)
}) {
	a.Current = from.GetCurrent()
	a.OfficialApp = from.GetOfficialApp()
	a.PasswordPending = from.GetPasswordPending()
	a.EncryptedRequestsDisabled = from.GetEncryptedRequestsDisabled()
	a.CallRequestsDisabled = from.GetCallRequestsDisabled()
	a.Unconfirmed = from.GetUnconfirmed()
	a.Hash = from.GetHash()
	a.DeviceModel = from.GetDeviceModel()
	a.Platform = from.GetPlatform()
	a.SystemVersion = from.GetSystemVersion()
	a.APIID = from.GetAPIID()
	a.AppName = from.GetAppName()
	a.AppVersion = from.GetAppVersion()
	a.DateCreated = from.GetDateCreated()
	a.DateActive = from.GetDateActive()
	a.IP = from.GetIP()
	a.Country = from.GetCountry()
	a.Region = from.GetRegion()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Authorization) TypeID() uint32 {
	return AuthorizationTypeID
}

// TypeName returns name of type in TL schema.
func (*Authorization) TypeName() string {
	return "authorization"
}

// TypeInfo returns info about TL type.
func (a *Authorization) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "authorization",
		ID:   AuthorizationTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Current",
			SchemaName: "current",
			Null:       !a.Flags.Has(0),
		},
		{
			Name:       "OfficialApp",
			SchemaName: "official_app",
			Null:       !a.Flags.Has(1),
		},
		{
			Name:       "PasswordPending",
			SchemaName: "password_pending",
			Null:       !a.Flags.Has(2),
		},
		{
			Name:       "EncryptedRequestsDisabled",
			SchemaName: "encrypted_requests_disabled",
			Null:       !a.Flags.Has(3),
		},
		{
			Name:       "CallRequestsDisabled",
			SchemaName: "call_requests_disabled",
			Null:       !a.Flags.Has(4),
		},
		{
			Name:       "Unconfirmed",
			SchemaName: "unconfirmed",
			Null:       !a.Flags.Has(5),
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "DeviceModel",
			SchemaName: "device_model",
		},
		{
			Name:       "Platform",
			SchemaName: "platform",
		},
		{
			Name:       "SystemVersion",
			SchemaName: "system_version",
		},
		{
			Name:       "APIID",
			SchemaName: "api_id",
		},
		{
			Name:       "AppName",
			SchemaName: "app_name",
		},
		{
			Name:       "AppVersion",
			SchemaName: "app_version",
		},
		{
			Name:       "DateCreated",
			SchemaName: "date_created",
		},
		{
			Name:       "DateActive",
			SchemaName: "date_active",
		},
		{
			Name:       "IP",
			SchemaName: "ip",
		},
		{
			Name:       "Country",
			SchemaName: "country",
		},
		{
			Name:       "Region",
			SchemaName: "region",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (a *Authorization) SetFlags() {
	if !(a.Current == false) {
		a.Flags.Set(0)
	}
	if !(a.OfficialApp == false) {
		a.Flags.Set(1)
	}
	if !(a.PasswordPending == false) {
		a.Flags.Set(2)
	}
	if !(a.EncryptedRequestsDisabled == false) {
		a.Flags.Set(3)
	}
	if !(a.CallRequestsDisabled == false) {
		a.Flags.Set(4)
	}
	if !(a.Unconfirmed == false) {
		a.Flags.Set(5)
	}
}

// Encode implements bin.Encoder.
func (a *Authorization) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authorization#ad01d61d as nil")
	}
	b.PutID(AuthorizationTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *Authorization) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode authorization#ad01d61d as nil")
	}
	a.SetFlags()
	if err := a.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode authorization#ad01d61d: field flags: %w", err)
	}
	b.PutLong(a.Hash)
	b.PutString(a.DeviceModel)
	b.PutString(a.Platform)
	b.PutString(a.SystemVersion)
	b.PutInt(a.APIID)
	b.PutString(a.AppName)
	b.PutString(a.AppVersion)
	b.PutInt(a.DateCreated)
	b.PutInt(a.DateActive)
	b.PutString(a.IP)
	b.PutString(a.Country)
	b.PutString(a.Region)
	return nil
}

// Decode implements bin.Decoder.
func (a *Authorization) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authorization#ad01d61d to nil")
	}
	if err := b.ConsumeID(AuthorizationTypeID); err != nil {
		return fmt.Errorf("unable to decode authorization#ad01d61d: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *Authorization) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode authorization#ad01d61d to nil")
	}
	{
		if err := a.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field flags: %w", err)
		}
	}
	a.Current = a.Flags.Has(0)
	a.OfficialApp = a.Flags.Has(1)
	a.PasswordPending = a.Flags.Has(2)
	a.EncryptedRequestsDisabled = a.Flags.Has(3)
	a.CallRequestsDisabled = a.Flags.Has(4)
	a.Unconfirmed = a.Flags.Has(5)
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field hash: %w", err)
		}
		a.Hash = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field device_model: %w", err)
		}
		a.DeviceModel = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field platform: %w", err)
		}
		a.Platform = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field system_version: %w", err)
		}
		a.SystemVersion = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field api_id: %w", err)
		}
		a.APIID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field app_name: %w", err)
		}
		a.AppName = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field app_version: %w", err)
		}
		a.AppVersion = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field date_created: %w", err)
		}
		a.DateCreated = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field date_active: %w", err)
		}
		a.DateActive = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field ip: %w", err)
		}
		a.IP = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field country: %w", err)
		}
		a.Country = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode authorization#ad01d61d: field region: %w", err)
		}
		a.Region = value
	}
	return nil
}

// SetCurrent sets value of Current conditional field.
func (a *Authorization) SetCurrent(value bool) {
	if value {
		a.Flags.Set(0)
		a.Current = true
	} else {
		a.Flags.Unset(0)
		a.Current = false
	}
}

// GetCurrent returns value of Current conditional field.
func (a *Authorization) GetCurrent() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(0)
}

// SetOfficialApp sets value of OfficialApp conditional field.
func (a *Authorization) SetOfficialApp(value bool) {
	if value {
		a.Flags.Set(1)
		a.OfficialApp = true
	} else {
		a.Flags.Unset(1)
		a.OfficialApp = false
	}
}

// GetOfficialApp returns value of OfficialApp conditional field.
func (a *Authorization) GetOfficialApp() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(1)
}

// SetPasswordPending sets value of PasswordPending conditional field.
func (a *Authorization) SetPasswordPending(value bool) {
	if value {
		a.Flags.Set(2)
		a.PasswordPending = true
	} else {
		a.Flags.Unset(2)
		a.PasswordPending = false
	}
}

// GetPasswordPending returns value of PasswordPending conditional field.
func (a *Authorization) GetPasswordPending() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(2)
}

// SetEncryptedRequestsDisabled sets value of EncryptedRequestsDisabled conditional field.
func (a *Authorization) SetEncryptedRequestsDisabled(value bool) {
	if value {
		a.Flags.Set(3)
		a.EncryptedRequestsDisabled = true
	} else {
		a.Flags.Unset(3)
		a.EncryptedRequestsDisabled = false
	}
}

// GetEncryptedRequestsDisabled returns value of EncryptedRequestsDisabled conditional field.
func (a *Authorization) GetEncryptedRequestsDisabled() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(3)
}

// SetCallRequestsDisabled sets value of CallRequestsDisabled conditional field.
func (a *Authorization) SetCallRequestsDisabled(value bool) {
	if value {
		a.Flags.Set(4)
		a.CallRequestsDisabled = true
	} else {
		a.Flags.Unset(4)
		a.CallRequestsDisabled = false
	}
}

// GetCallRequestsDisabled returns value of CallRequestsDisabled conditional field.
func (a *Authorization) GetCallRequestsDisabled() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(4)
}

// SetUnconfirmed sets value of Unconfirmed conditional field.
func (a *Authorization) SetUnconfirmed(value bool) {
	if value {
		a.Flags.Set(5)
		a.Unconfirmed = true
	} else {
		a.Flags.Unset(5)
		a.Unconfirmed = false
	}
}

// GetUnconfirmed returns value of Unconfirmed conditional field.
func (a *Authorization) GetUnconfirmed() (value bool) {
	if a == nil {
		return
	}
	return a.Flags.Has(5)
}

// GetHash returns value of Hash field.
func (a *Authorization) GetHash() (value int64) {
	if a == nil {
		return
	}
	return a.Hash
}

// GetDeviceModel returns value of DeviceModel field.
func (a *Authorization) GetDeviceModel() (value string) {
	if a == nil {
		return
	}
	return a.DeviceModel
}

// GetPlatform returns value of Platform field.
func (a *Authorization) GetPlatform() (value string) {
	if a == nil {
		return
	}
	return a.Platform
}

// GetSystemVersion returns value of SystemVersion field.
func (a *Authorization) GetSystemVersion() (value string) {
	if a == nil {
		return
	}
	return a.SystemVersion
}

// GetAPIID returns value of APIID field.
func (a *Authorization) GetAPIID() (value int) {
	if a == nil {
		return
	}
	return a.APIID
}

// GetAppName returns value of AppName field.
func (a *Authorization) GetAppName() (value string) {
	if a == nil {
		return
	}
	return a.AppName
}

// GetAppVersion returns value of AppVersion field.
func (a *Authorization) GetAppVersion() (value string) {
	if a == nil {
		return
	}
	return a.AppVersion
}

// GetDateCreated returns value of DateCreated field.
func (a *Authorization) GetDateCreated() (value int) {
	if a == nil {
		return
	}
	return a.DateCreated
}

// GetDateActive returns value of DateActive field.
func (a *Authorization) GetDateActive() (value int) {
	if a == nil {
		return
	}
	return a.DateActive
}

// GetIP returns value of IP field.
func (a *Authorization) GetIP() (value string) {
	if a == nil {
		return
	}
	return a.IP
}

// GetCountry returns value of Country field.
func (a *Authorization) GetCountry() (value string) {
	if a == nil {
		return
	}
	return a.Country
}

// GetRegion returns value of Region field.
func (a *Authorization) GetRegion() (value string) {
	if a == nil {
		return
	}
	return a.Region
}
