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

// ContactsImportedContacts represents TL type `contacts.importedContacts#77d01c3b`.
// Info on succesfully imported contacts.
//
// See https://core.telegram.org/constructor/contacts.importedContacts for reference.
type ContactsImportedContacts struct {
	// List of succesfully imported contacts
	Imported []ImportedContact
	// Popular contacts
	PopularInvites []PopularContact
	// List of contact ids that could not be imported due to system limitation and will need
	// to be imported at a later date.Parameter added in Layer 13¹
	//
	// Links:
	//  1) https://core.telegram.org/api/layers#layer-13
	RetryContacts []int64
	// List of users
	Users []UserClass
}

// ContactsImportedContactsTypeID is TL type id of ContactsImportedContacts.
const ContactsImportedContactsTypeID = 0x77d01c3b

// Ensuring interfaces in compile-time for ContactsImportedContacts.
var (
	_ bin.Encoder     = &ContactsImportedContacts{}
	_ bin.Decoder     = &ContactsImportedContacts{}
	_ bin.BareEncoder = &ContactsImportedContacts{}
	_ bin.BareDecoder = &ContactsImportedContacts{}
)

func (i *ContactsImportedContacts) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Imported == nil) {
		return false
	}
	if !(i.PopularInvites == nil) {
		return false
	}
	if !(i.RetryContacts == nil) {
		return false
	}
	if !(i.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *ContactsImportedContacts) String() string {
	if i == nil {
		return "ContactsImportedContacts(nil)"
	}
	type Alias ContactsImportedContacts
	return fmt.Sprintf("ContactsImportedContacts%+v", Alias(*i))
}

// FillFrom fills ContactsImportedContacts from given interface.
func (i *ContactsImportedContacts) FillFrom(from interface {
	GetImported() (value []ImportedContact)
	GetPopularInvites() (value []PopularContact)
	GetRetryContacts() (value []int64)
	GetUsers() (value []UserClass)
}) {
	i.Imported = from.GetImported()
	i.PopularInvites = from.GetPopularInvites()
	i.RetryContacts = from.GetRetryContacts()
	i.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsImportedContacts) TypeID() uint32 {
	return ContactsImportedContactsTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsImportedContacts) TypeName() string {
	return "contacts.importedContacts"
}

// TypeInfo returns info about TL type.
func (i *ContactsImportedContacts) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.importedContacts",
		ID:   ContactsImportedContactsTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Imported",
			SchemaName: "imported",
		},
		{
			Name:       "PopularInvites",
			SchemaName: "popular_invites",
		},
		{
			Name:       "RetryContacts",
			SchemaName: "retry_contacts",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *ContactsImportedContacts) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode contacts.importedContacts#77d01c3b as nil")
	}
	b.PutID(ContactsImportedContactsTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *ContactsImportedContacts) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode contacts.importedContacts#77d01c3b as nil")
	}
	b.PutVectorHeader(len(i.Imported))
	for idx, v := range i.Imported {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field imported element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(i.PopularInvites))
	for idx, v := range i.PopularInvites {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field popular_invites element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(i.RetryContacts))
	for _, v := range i.RetryContacts {
		b.PutLong(v)
	}
	b.PutVectorHeader(len(i.Users))
	for idx, v := range i.Users {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.importedContacts#77d01c3b: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *ContactsImportedContacts) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode contacts.importedContacts#77d01c3b to nil")
	}
	if err := b.ConsumeID(ContactsImportedContactsTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *ContactsImportedContacts) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode contacts.importedContacts#77d01c3b to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field imported: %w", err)
		}

		if headerLen > 0 {
			i.Imported = make([]ImportedContact, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ImportedContact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field imported: %w", err)
			}
			i.Imported = append(i.Imported, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field popular_invites: %w", err)
		}

		if headerLen > 0 {
			i.PopularInvites = make([]PopularContact, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value PopularContact
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field popular_invites: %w", err)
			}
			i.PopularInvites = append(i.PopularInvites, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field retry_contacts: %w", err)
		}

		if headerLen > 0 {
			i.RetryContacts = make([]int64, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field retry_contacts: %w", err)
			}
			i.RetryContacts = append(i.RetryContacts, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field users: %w", err)
		}

		if headerLen > 0 {
			i.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.importedContacts#77d01c3b: field users: %w", err)
			}
			i.Users = append(i.Users, value)
		}
	}
	return nil
}

// GetImported returns value of Imported field.
func (i *ContactsImportedContacts) GetImported() (value []ImportedContact) {
	return i.Imported
}

// GetPopularInvites returns value of PopularInvites field.
func (i *ContactsImportedContacts) GetPopularInvites() (value []PopularContact) {
	return i.PopularInvites
}

// GetRetryContacts returns value of RetryContacts field.
func (i *ContactsImportedContacts) GetRetryContacts() (value []int64) {
	return i.RetryContacts
}

// GetUsers returns value of Users field.
func (i *ContactsImportedContacts) GetUsers() (value []UserClass) {
	return i.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (i *ContactsImportedContacts) MapUsers() (value UserClassArray) {
	return UserClassArray(i.Users)
}
