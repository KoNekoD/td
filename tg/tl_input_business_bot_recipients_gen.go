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

// InputBusinessBotRecipients represents TL type `inputBusinessBotRecipients#c4e5921e`.
//
// See https://core.telegram.org/constructor/inputBusinessBotRecipients for reference.
type InputBusinessBotRecipients struct {
	// Flags field of InputBusinessBotRecipients.
	Flags bin.Fields
	// ExistingChats field of InputBusinessBotRecipients.
	ExistingChats bool
	// NewChats field of InputBusinessBotRecipients.
	NewChats bool
	// Contacts field of InputBusinessBotRecipients.
	Contacts bool
	// NonContacts field of InputBusinessBotRecipients.
	NonContacts bool
	// ExcludeSelected field of InputBusinessBotRecipients.
	ExcludeSelected bool
	// Users field of InputBusinessBotRecipients.
	//
	// Use SetUsers and GetUsers helpers.
	Users []InputUserClass
	// ExcludeUsers field of InputBusinessBotRecipients.
	//
	// Use SetExcludeUsers and GetExcludeUsers helpers.
	ExcludeUsers []InputUserClass
}

// InputBusinessBotRecipientsTypeID is TL type id of InputBusinessBotRecipients.
const InputBusinessBotRecipientsTypeID = 0xc4e5921e

// Ensuring interfaces in compile-time for InputBusinessBotRecipients.
var (
	_ bin.Encoder     = &InputBusinessBotRecipients{}
	_ bin.Decoder     = &InputBusinessBotRecipients{}
	_ bin.BareEncoder = &InputBusinessBotRecipients{}
	_ bin.BareDecoder = &InputBusinessBotRecipients{}
)

func (i *InputBusinessBotRecipients) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.ExistingChats == false) {
		return false
	}
	if !(i.NewChats == false) {
		return false
	}
	if !(i.Contacts == false) {
		return false
	}
	if !(i.NonContacts == false) {
		return false
	}
	if !(i.ExcludeSelected == false) {
		return false
	}
	if !(i.Users == nil) {
		return false
	}
	if !(i.ExcludeUsers == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputBusinessBotRecipients) String() string {
	if i == nil {
		return "InputBusinessBotRecipients(nil)"
	}
	type Alias InputBusinessBotRecipients
	return fmt.Sprintf("InputBusinessBotRecipients%+v", Alias(*i))
}

// FillFrom fills InputBusinessBotRecipients from given interface.
func (i *InputBusinessBotRecipients) FillFrom(from interface {
	GetExistingChats() (value bool)
	GetNewChats() (value bool)
	GetContacts() (value bool)
	GetNonContacts() (value bool)
	GetExcludeSelected() (value bool)
	GetUsers() (value []InputUserClass, ok bool)
	GetExcludeUsers() (value []InputUserClass, ok bool)
}) {
	i.ExistingChats = from.GetExistingChats()
	i.NewChats = from.GetNewChats()
	i.Contacts = from.GetContacts()
	i.NonContacts = from.GetNonContacts()
	i.ExcludeSelected = from.GetExcludeSelected()
	if val, ok := from.GetUsers(); ok {
		i.Users = val
	}

	if val, ok := from.GetExcludeUsers(); ok {
		i.ExcludeUsers = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputBusinessBotRecipients) TypeID() uint32 {
	return InputBusinessBotRecipientsTypeID
}

// TypeName returns name of type in TL schema.
func (*InputBusinessBotRecipients) TypeName() string {
	return "inputBusinessBotRecipients"
}

// TypeInfo returns info about TL type.
func (i *InputBusinessBotRecipients) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputBusinessBotRecipients",
		ID:   InputBusinessBotRecipientsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ExistingChats",
			SchemaName: "existing_chats",
			Null:       !i.Flags.Has(0),
		},
		{
			Name:       "NewChats",
			SchemaName: "new_chats",
			Null:       !i.Flags.Has(1),
		},
		{
			Name:       "Contacts",
			SchemaName: "contacts",
			Null:       !i.Flags.Has(2),
		},
		{
			Name:       "NonContacts",
			SchemaName: "non_contacts",
			Null:       !i.Flags.Has(3),
		},
		{
			Name:       "ExcludeSelected",
			SchemaName: "exclude_selected",
			Null:       !i.Flags.Has(5),
		},
		{
			Name:       "Users",
			SchemaName: "users",
			Null:       !i.Flags.Has(4),
		},
		{
			Name:       "ExcludeUsers",
			SchemaName: "exclude_users",
			Null:       !i.Flags.Has(6),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (i *InputBusinessBotRecipients) SetFlags() {
	if !(i.ExistingChats == false) {
		i.Flags.Set(0)
	}
	if !(i.NewChats == false) {
		i.Flags.Set(1)
	}
	if !(i.Contacts == false) {
		i.Flags.Set(2)
	}
	if !(i.NonContacts == false) {
		i.Flags.Set(3)
	}
	if !(i.ExcludeSelected == false) {
		i.Flags.Set(5)
	}
	if !(i.Users == nil) {
		i.Flags.Set(4)
	}
	if !(i.ExcludeUsers == nil) {
		i.Flags.Set(6)
	}
}

// Encode implements bin.Encoder.
func (i *InputBusinessBotRecipients) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBusinessBotRecipients#c4e5921e as nil")
	}
	b.PutID(InputBusinessBotRecipientsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputBusinessBotRecipients) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputBusinessBotRecipients#c4e5921e as nil")
	}
	i.SetFlags()
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputBusinessBotRecipients#c4e5921e: field flags: %w", err)
	}
	if i.Flags.Has(4) {
		b.PutVectorHeader(len(i.Users))
		for idx, v := range i.Users {
			if v == nil {
				return fmt.Errorf("unable to encode inputBusinessBotRecipients#c4e5921e: field users element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputBusinessBotRecipients#c4e5921e: field users element with index %d: %w", idx, err)
			}
		}
	}
	if i.Flags.Has(6) {
		b.PutVectorHeader(len(i.ExcludeUsers))
		for idx, v := range i.ExcludeUsers {
			if v == nil {
				return fmt.Errorf("unable to encode inputBusinessBotRecipients#c4e5921e: field exclude_users element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode inputBusinessBotRecipients#c4e5921e: field exclude_users element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputBusinessBotRecipients) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBusinessBotRecipients#c4e5921e to nil")
	}
	if err := b.ConsumeID(InputBusinessBotRecipientsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputBusinessBotRecipients#c4e5921e: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputBusinessBotRecipients) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputBusinessBotRecipients#c4e5921e to nil")
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputBusinessBotRecipients#c4e5921e: field flags: %w", err)
		}
	}
	i.ExistingChats = i.Flags.Has(0)
	i.NewChats = i.Flags.Has(1)
	i.Contacts = i.Flags.Has(2)
	i.NonContacts = i.Flags.Has(3)
	i.ExcludeSelected = i.Flags.Has(5)
	if i.Flags.Has(4) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputBusinessBotRecipients#c4e5921e: field users: %w", err)
		}

		if headerLen > 0 {
			i.Users = make([]InputUserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputBusinessBotRecipients#c4e5921e: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	if i.Flags.Has(6) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode inputBusinessBotRecipients#c4e5921e: field exclude_users: %w", err)
		}

		if headerLen > 0 {
			i.ExcludeUsers = make([]InputUserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode inputBusinessBotRecipients#c4e5921e: field exclude_users: %w", err)
			}
			i.ExcludeUsers = append(i.ExcludeUsers, value)
		}
	}
	return nil
}

// SetExistingChats sets value of ExistingChats conditional field.
func (i *InputBusinessBotRecipients) SetExistingChats(value bool) {
	if value {
		i.Flags.Set(0)
		i.ExistingChats = true
	} else {
		i.Flags.Unset(0)
		i.ExistingChats = false
	}
}

// GetExistingChats returns value of ExistingChats conditional field.
func (i *InputBusinessBotRecipients) GetExistingChats() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(0)
}

// SetNewChats sets value of NewChats conditional field.
func (i *InputBusinessBotRecipients) SetNewChats(value bool) {
	if value {
		i.Flags.Set(1)
		i.NewChats = true
	} else {
		i.Flags.Unset(1)
		i.NewChats = false
	}
}

// GetNewChats returns value of NewChats conditional field.
func (i *InputBusinessBotRecipients) GetNewChats() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(1)
}

// SetContacts sets value of Contacts conditional field.
func (i *InputBusinessBotRecipients) SetContacts(value bool) {
	if value {
		i.Flags.Set(2)
		i.Contacts = true
	} else {
		i.Flags.Unset(2)
		i.Contacts = false
	}
}

// GetContacts returns value of Contacts conditional field.
func (i *InputBusinessBotRecipients) GetContacts() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(2)
}

// SetNonContacts sets value of NonContacts conditional field.
func (i *InputBusinessBotRecipients) SetNonContacts(value bool) {
	if value {
		i.Flags.Set(3)
		i.NonContacts = true
	} else {
		i.Flags.Unset(3)
		i.NonContacts = false
	}
}

// GetNonContacts returns value of NonContacts conditional field.
func (i *InputBusinessBotRecipients) GetNonContacts() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(3)
}

// SetExcludeSelected sets value of ExcludeSelected conditional field.
func (i *InputBusinessBotRecipients) SetExcludeSelected(value bool) {
	if value {
		i.Flags.Set(5)
		i.ExcludeSelected = true
	} else {
		i.Flags.Unset(5)
		i.ExcludeSelected = false
	}
}

// GetExcludeSelected returns value of ExcludeSelected conditional field.
func (i *InputBusinessBotRecipients) GetExcludeSelected() (value bool) {
	if i == nil {
		return
	}
	return i.Flags.Has(5)
}

// SetUsers sets value of Users conditional field.
func (i *InputBusinessBotRecipients) SetUsers(value []InputUserClass) {
	i.Flags.Set(4)
	i.Users = value
}

// GetUsers returns value of Users conditional field and
// boolean which is true if field was set.
func (i *InputBusinessBotRecipients) GetUsers() (value []InputUserClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(4) {
		return value, false
	}
	return i.Users, true
}

// SetExcludeUsers sets value of ExcludeUsers conditional field.
func (i *InputBusinessBotRecipients) SetExcludeUsers(value []InputUserClass) {
	i.Flags.Set(6)
	i.ExcludeUsers = value
}

// GetExcludeUsers returns value of ExcludeUsers conditional field and
// boolean which is true if field was set.
func (i *InputBusinessBotRecipients) GetExcludeUsers() (value []InputUserClass, ok bool) {
	if i == nil {
		return
	}
	if !i.Flags.Has(6) {
		return value, false
	}
	return i.ExcludeUsers, true
}

// MapUsers returns field Users wrapped in InputUserClassArray helper.
func (i *InputBusinessBotRecipients) MapUsers() (value InputUserClassArray, ok bool) {
	if !i.Flags.Has(4) {
		return value, false
	}
	return InputUserClassArray(i.Users), true
}

// MapExcludeUsers returns field ExcludeUsers wrapped in InputUserClassArray helper.
func (i *InputBusinessBotRecipients) MapExcludeUsers() (value InputUserClassArray, ok bool) {
	if !i.Flags.Has(6) {
		return value, false
	}
	return InputUserClassArray(i.ExcludeUsers), true
}
