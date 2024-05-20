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

// AccountUploadWallPaperRequest represents TL type `account.uploadWallPaper#e39a8f03`.
// Create and upload a new wallpaper¹
//
// Links:
//  1. https://core.telegram.org/api/wallpapers
//
// See https://core.telegram.org/method/account.uploadWallPaper for reference.
type AccountUploadWallPaperRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Set this flag when uploading wallpapers to be passed to messages.setChatWallPaper¹.
	//
	// Links:
	//  1) https://core.telegram.org/method/messages.setChatWallPaper
	ForChat bool
	// The JPG/PNG wallpaper
	File InputFileClass
	// MIME type of uploaded wallpaper
	MimeType string
	// Wallpaper settings
	Settings WallPaperSettings
}

// AccountUploadWallPaperRequestTypeID is TL type id of AccountUploadWallPaperRequest.
const AccountUploadWallPaperRequestTypeID = 0xe39a8f03

// Ensuring interfaces in compile-time for AccountUploadWallPaperRequest.
var (
	_ bin.Encoder     = &AccountUploadWallPaperRequest{}
	_ bin.Decoder     = &AccountUploadWallPaperRequest{}
	_ bin.BareEncoder = &AccountUploadWallPaperRequest{}
	_ bin.BareDecoder = &AccountUploadWallPaperRequest{}
)

func (u *AccountUploadWallPaperRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Flags.Zero()) {
		return false
	}
	if !(u.ForChat == false) {
		return false
	}
	if !(u.File == nil) {
		return false
	}
	if !(u.MimeType == "") {
		return false
	}
	if !(u.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUploadWallPaperRequest) String() string {
	if u == nil {
		return "AccountUploadWallPaperRequest(nil)"
	}
	type Alias AccountUploadWallPaperRequest
	return fmt.Sprintf("AccountUploadWallPaperRequest%+v", Alias(*u))
}

// FillFrom fills AccountUploadWallPaperRequest from given interface.
func (u *AccountUploadWallPaperRequest) FillFrom(from interface {
	GetForChat() (value bool)
	GetFile() (value InputFileClass)
	GetMimeType() (value string)
	GetSettings() (value WallPaperSettings)
}) {
	u.ForChat = from.GetForChat()
	u.File = from.GetFile()
	u.MimeType = from.GetMimeType()
	u.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUploadWallPaperRequest) TypeID() uint32 {
	return AccountUploadWallPaperRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUploadWallPaperRequest) TypeName() string {
	return "account.uploadWallPaper"
}

// TypeInfo returns info about TL type.
func (u *AccountUploadWallPaperRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.uploadWallPaper",
		ID:   AccountUploadWallPaperRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ForChat",
			SchemaName: "for_chat",
			Null:       !u.Flags.Has(0),
		},
		{
			Name:       "File",
			SchemaName: "file",
		},
		{
			Name:       "MimeType",
			SchemaName: "mime_type",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (u *AccountUploadWallPaperRequest) SetFlags() {
	if !(u.ForChat == false) {
		u.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (u *AccountUploadWallPaperRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.uploadWallPaper#e39a8f03 as nil")
	}
	b.PutID(AccountUploadWallPaperRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUploadWallPaperRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.uploadWallPaper#e39a8f03 as nil")
	}
	u.SetFlags()
	if err := u.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadWallPaper#e39a8f03: field flags: %w", err)
	}
	if u.File == nil {
		return fmt.Errorf("unable to encode account.uploadWallPaper#e39a8f03: field file is nil")
	}
	if err := u.File.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadWallPaper#e39a8f03: field file: %w", err)
	}
	b.PutString(u.MimeType)
	if err := u.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.uploadWallPaper#e39a8f03: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUploadWallPaperRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.uploadWallPaper#e39a8f03 to nil")
	}
	if err := b.ConsumeID(AccountUploadWallPaperRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.uploadWallPaper#e39a8f03: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUploadWallPaperRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.uploadWallPaper#e39a8f03 to nil")
	}
	{
		if err := u.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.uploadWallPaper#e39a8f03: field flags: %w", err)
		}
	}
	u.ForChat = u.Flags.Has(0)
	{
		value, err := DecodeInputFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadWallPaper#e39a8f03: field file: %w", err)
		}
		u.File = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.uploadWallPaper#e39a8f03: field mime_type: %w", err)
		}
		u.MimeType = value
	}
	{
		if err := u.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.uploadWallPaper#e39a8f03: field settings: %w", err)
		}
	}
	return nil
}

// SetForChat sets value of ForChat conditional field.
func (u *AccountUploadWallPaperRequest) SetForChat(value bool) {
	if value {
		u.Flags.Set(0)
		u.ForChat = true
	} else {
		u.Flags.Unset(0)
		u.ForChat = false
	}
}

// GetForChat returns value of ForChat conditional field.
func (u *AccountUploadWallPaperRequest) GetForChat() (value bool) {
	if u == nil {
		return
	}
	return u.Flags.Has(0)
}

// GetFile returns value of File field.
func (u *AccountUploadWallPaperRequest) GetFile() (value InputFileClass) {
	if u == nil {
		return
	}
	return u.File
}

// GetMimeType returns value of MimeType field.
func (u *AccountUploadWallPaperRequest) GetMimeType() (value string) {
	if u == nil {
		return
	}
	return u.MimeType
}

// GetSettings returns value of Settings field.
func (u *AccountUploadWallPaperRequest) GetSettings() (value WallPaperSettings) {
	if u == nil {
		return
	}
	return u.Settings
}

// AccountUploadWallPaper invokes method account.uploadWallPaper#e39a8f03 returning error if any.
// Create and upload a new wallpaper¹
//
// Links:
//  1. https://core.telegram.org/api/wallpapers
//
// Possible errors:
//
//	400 WALLPAPER_FILE_INVALID: The specified wallpaper file is invalid.
//	400 WALLPAPER_MIME_INVALID: The specified wallpaper MIME type is invalid.
//
// See https://core.telegram.org/method/account.uploadWallPaper for reference.
func (c *Client) AccountUploadWallPaper(ctx context.Context, request *AccountUploadWallPaperRequest) (WallPaperClass, error) {
	var result WallPaperBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.WallPaper, nil
}
