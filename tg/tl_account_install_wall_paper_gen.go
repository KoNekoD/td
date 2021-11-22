// Code generated by gotdgen, DO NOT EDIT.

package tg

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

// AccountInstallWallPaperRequest represents TL type `account.installWallPaper#feed5769`.
// Install wallpaper
//
// See https://core.telegram.org/method/account.installWallPaper for reference.
type AccountInstallWallPaperRequest struct {
	// Wallpaper to install
	Wallpaper InputWallPaperClass
	// Wallpaper settings
	Settings WallPaperSettings
}

// AccountInstallWallPaperRequestTypeID is TL type id of AccountInstallWallPaperRequest.
const AccountInstallWallPaperRequestTypeID = 0xfeed5769

// Ensuring interfaces in compile-time for AccountInstallWallPaperRequest.
var (
	_ bin.Encoder     = &AccountInstallWallPaperRequest{}
	_ bin.Decoder     = &AccountInstallWallPaperRequest{}
	_ bin.BareEncoder = &AccountInstallWallPaperRequest{}
	_ bin.BareDecoder = &AccountInstallWallPaperRequest{}
)

func (i *AccountInstallWallPaperRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Wallpaper == nil) {
		return false
	}
	if !(i.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *AccountInstallWallPaperRequest) String() string {
	if i == nil {
		return "AccountInstallWallPaperRequest(nil)"
	}
	type Alias AccountInstallWallPaperRequest
	return fmt.Sprintf("AccountInstallWallPaperRequest%+v", Alias(*i))
}

// FillFrom fills AccountInstallWallPaperRequest from given interface.
func (i *AccountInstallWallPaperRequest) FillFrom(from interface {
	GetWallpaper() (value InputWallPaperClass)
	GetSettings() (value WallPaperSettings)
}) {
	i.Wallpaper = from.GetWallpaper()
	i.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountInstallWallPaperRequest) TypeID() uint32 {
	return AccountInstallWallPaperRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountInstallWallPaperRequest) TypeName() string {
	return "account.installWallPaper"
}

// TypeInfo returns info about TL type.
func (i *AccountInstallWallPaperRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.installWallPaper",
		ID:   AccountInstallWallPaperRequestTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Wallpaper",
			SchemaName: "wallpaper",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *AccountInstallWallPaperRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.installWallPaper#feed5769 as nil")
	}
	b.PutID(AccountInstallWallPaperRequestTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *AccountInstallWallPaperRequest) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode account.installWallPaper#feed5769 as nil")
	}
	if i.Wallpaper == nil {
		return fmt.Errorf("unable to encode account.installWallPaper#feed5769: field wallpaper is nil")
	}
	if err := i.Wallpaper.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.installWallPaper#feed5769: field wallpaper: %w", err)
	}
	if err := i.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.installWallPaper#feed5769: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *AccountInstallWallPaperRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.installWallPaper#feed5769 to nil")
	}
	if err := b.ConsumeID(AccountInstallWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.installWallPaper#feed5769: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *AccountInstallWallPaperRequest) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode account.installWallPaper#feed5769 to nil")
	}
	{
		value, err := DecodeInputWallPaper(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.installWallPaper#feed5769: field wallpaper: %w", err)
		}
		i.Wallpaper = value
	}
	{
		if err := i.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.installWallPaper#feed5769: field settings: %w", err)
		}
	}
	return nil
}

// GetWallpaper returns value of Wallpaper field.
func (i *AccountInstallWallPaperRequest) GetWallpaper() (value InputWallPaperClass) {
	return i.Wallpaper
}

// GetSettings returns value of Settings field.
func (i *AccountInstallWallPaperRequest) GetSettings() (value WallPaperSettings) {
	return i.Settings
}

// AccountInstallWallPaper invokes method account.installWallPaper#feed5769 returning error if any.
// Install wallpaper
//
// See https://core.telegram.org/method/account.installWallPaper for reference.
func (c *Client) AccountInstallWallPaper(ctx context.Context, request *AccountInstallWallPaperRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
