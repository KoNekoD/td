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

// AccountInstallThemeRequest represents TL type `account.installTheme#c727bb3b`.
// Install a theme
//
// See https://core.telegram.org/method/account.installTheme for reference.
type AccountInstallThemeRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to install the dark version
	Dark bool
	// Theme to install
	//
	// Use SetTheme and GetTheme helpers.
	Theme InputThemeClass
	// Theme format, a string that identifies the theming engines supported by the client
	//
	// Use SetFormat and GetFormat helpers.
	Format string
	// Indicates a basic theme provided by all clients
	//
	// Use SetBaseTheme and GetBaseTheme helpers.
	BaseTheme BaseThemeClass
}

// AccountInstallThemeRequestTypeID is TL type id of AccountInstallThemeRequest.
const AccountInstallThemeRequestTypeID = 0xc727bb3b

// Ensuring interfaces in compile-time for AccountInstallThemeRequest.
var (
	_ bin.Encoder     = &AccountInstallThemeRequest{}
	_ bin.Decoder     = &AccountInstallThemeRequest{}
	_ bin.BareEncoder = &AccountInstallThemeRequest{}
	_ bin.BareDecoder = &AccountInstallThemeRequest{}
)

func (i *AccountInstallThemeRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Dark == false) {
		return false
	}
	if !(i.Theme == nil) {
		return false
	}
	if !(i.Format == "") {
		return false
	}
	if !(i.BaseTheme == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AccountInstallThemeRequest) String() string {
	if i == nil {
		return "AccountInstallThemeRequest(nil)"
	}
	type Alias AccountInstallThemeRequest
	return fmt.Sprintf("AccountInstallThemeRequest%+v", Alias(*i))
}

// FillFrom fills AccountInstallThemeRequest from given interface.
func (i *AccountInstallThemeRequest) FillFrom(from interface {
	GetDark() (value bool)
	GetTheme() (value InputThemeClass, ok bool)
	GetFormat() (value string, ok bool)
	GetBaseTheme() (value BaseThemeClass, ok bool)
}) {
	i.Dark = from.GetDark()
	if val, ok := from.GetTheme(); ok {
		i.Theme = val
	}

	if val, ok := from.GetFormat(); ok {
		i.Format = val
	}

	if val, ok := from.GetBaseTheme(); ok {
		i.BaseTheme = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountInstallThemeRequest) TypeID() uint32 {
	return AccountInstallThemeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountInstallThemeRequest) TypeName() string {
	return "account.installTheme"
}

// TypeInfo returns info about TL type.
func (i *AccountInstallThemeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.installTheme",
		ID:   AccountInstallThemeRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Dark",
			SchemaName: "dark",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "Theme",
			SchemaName: "theme",
			Null:       !i.Flags.Has(1),
		},
		{
			Name:       "Format",
			SchemaName: "format",
			Null:       !i.Flags.Has(2),
		},
		{
			Name:       "BaseTheme",
			SchemaName: "base_theme",
			Null:       !i.Flags.Has(3),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *AccountInstallThemeRequest) SetFlags() {
	if !(i.Dark == false) {
		i.Flags.Set(0)
	}
	if !(i.Theme == nil) {
		i.Flags.Set(1)
	}
	if !(i.Format == "") {
		i.Flags.Set(2)
	}
	if !(i.BaseTheme == nil) {
		i.Flags.Set(3)
	}
}

// Encode implements bin.Encoder.
func (i *AccountInstallThemeRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.installTheme#c727bb3b as nil")
	}
	b.PutID(AccountInstallThemeRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *AccountInstallThemeRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.installTheme#c727bb3b as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.installTheme#c727bb3b: field flags: %w", err)
	}
	if i.Flags.Has(1) {
		if i.Theme == nil {
			return fmt.Errorf("unable to encode account.installTheme#c727bb3b: field theme is nil")
		}
		if err := i.Theme.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.installTheme#c727bb3b: field theme: %w", err)
		}
	}
	if i.Flags.Has(2) {
		b.PutString(i.Format)
	}
	if i.Flags.Has(3) {
		if i.BaseTheme == nil {
			return fmt.Errorf("unable to encode account.installTheme#c727bb3b: field base_theme is nil")
		}
		if err := i.BaseTheme.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.installTheme#c727bb3b: field base_theme: %w", err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *AccountInstallThemeRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.installTheme#c727bb3b to nil")
	}
	if err := b.ConsumeID(AccountInstallThemeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.installTheme#c727bb3b: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *AccountInstallThemeRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.installTheme#c727bb3b to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.installTheme#c727bb3b: field flags: %w", err)
		}
	}
	i.Dark = i.Flags.Has(0)
	if i.Flags.Has(1) {
		value, err := DecodeInputTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.installTheme#c727bb3b: field theme: %w", err)
		}
		i.Theme = value
	}
	if i.Flags.Has(2) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.installTheme#c727bb3b: field format: %w", err)
		}
		i.Format = value
	}
	if i.Flags.Has(3) {
		value, err := DecodeBaseTheme(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.installTheme#c727bb3b: field base_theme: %w", err)
		}
		i.BaseTheme = value
	}
	return nil
}

// SetDark sets value of Dark conditional field.
func (i *AccountInstallThemeRequest) SetDark(value bool) {
	if value {
		i.Flags.Set(0)
		i.Dark = true
	} else {
		i.Flags.Unset(0)
		i.Dark = false
	}
}

// GetDark returns value of Dark conditional field.
func (i *AccountInstallThemeRequest) GetDark() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(0)
}

// SetTheme sets value of Theme conditional field.
func (i *AccountInstallThemeRequest) SetTheme(value InputThemeClass) {
	i.Flags.Set(1)
	i.Theme = value
}

// GetTheme returns value of Theme conditional field and
// boolean which is true if field was set.
func (i *AccountInstallThemeRequest) GetTheme() (value InputThemeClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(1) {
		return value, false
	}
	return i.Theme, true
}

// SetFormat sets value of Format conditional field.
func (i *AccountInstallThemeRequest) SetFormat(value string) {
	i.Flags.Set(2)
	i.Format = value
}

// GetFormat returns value of Format conditional field and
// boolean which is true if field was set.
func (i *AccountInstallThemeRequest) GetFormat() (value string, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(2) {
		return value, false
	}
	return i.Format, true
}

// SetBaseTheme sets value of BaseTheme conditional field.
func (i *AccountInstallThemeRequest) SetBaseTheme(value BaseThemeClass) {
	i.Flags.Set(3)
	i.BaseTheme = value
}

// GetBaseTheme returns value of BaseTheme conditional field and
// boolean which is true if field was set.
func (i *AccountInstallThemeRequest) GetBaseTheme() (value BaseThemeClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(3) {
		return value, false
	}
	return i.BaseTheme, true
}

// AccountInstallTheme invokes method account.installTheme#c727bb3b returning error if any.
// Install a theme
//
// See https://core.telegram.org/method/account.installTheme for reference.
func (c *Client) AccountInstallTheme(ctx context.Context, request *AccountInstallThemeRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
