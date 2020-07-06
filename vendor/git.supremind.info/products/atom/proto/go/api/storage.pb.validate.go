// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: storage.proto

package api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _storage_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Storage with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Storage) Validate() error {
	if m == nil {
		return nil
	}

	// skipping validation for meta

	{
		tmp := m.GetSpec()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return StorageValidationError{
					field:  "Spec",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetStatus()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return StorageValidationError{
					field:  "Status",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// StorageValidationError is the validation error returned by Storage.Validate
// if the designated constraints aren't met.
type StorageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageValidationError) ErrorName() string { return "StorageValidationError" }

// Error satisfies the builtin error interface
func (e StorageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageValidationError{}

// Validate checks the field values on StorageSpec with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StorageSpec) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetCapacity() <= 0 {
		return StorageSpecValidationError{
			field:  "Capacity",
			reason: "value must be greater than 0",
		}
	}

	if _, ok := StorageMode_name[int32(m.GetMode())]; !ok {
		return StorageSpecValidationError{
			field:  "Mode",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// StorageSpecValidationError is the validation error returned by
// StorageSpec.Validate if the designated constraints aren't met.
type StorageSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageSpecValidationError) ErrorName() string { return "StorageSpecValidationError" }

// Error satisfies the builtin error interface
func (e StorageSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageSpecValidationError{}

// Validate checks the field values on StorageStatus with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *StorageStatus) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetUsedBy() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return StorageStatusValidationError{
						field:  fmt.Sprintf("UsedBy[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetAvailable()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return StorageStatusValidationError{
					field:  "Available",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// StorageStatusValidationError is the validation error returned by
// StorageStatus.Validate if the designated constraints aren't met.
type StorageStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageStatusValidationError) ErrorName() string { return "StorageStatusValidationError" }

// Error satisfies the builtin error interface
func (e StorageStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageStatusValidationError{}

// Validate checks the field values on CreateStorageReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateStorageReq) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetStorage()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return CreateStorageReqValidationError{
					field:  "Storage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// CreateStorageReqValidationError is the validation error returned by
// CreateStorageReq.Validate if the designated constraints aren't met.
type CreateStorageReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateStorageReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateStorageReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateStorageReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateStorageReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateStorageReqValidationError) ErrorName() string { return "CreateStorageReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateStorageReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateStorageReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateStorageReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateStorageReqValidationError{}

// Validate checks the field values on GetStorageReq with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetStorageReq) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return GetStorageReqValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetCreator()) < 1 {
		return GetStorageReqValidationError{
			field:  "Creator",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// GetStorageReqValidationError is the validation error returned by
// GetStorageReq.Validate if the designated constraints aren't met.
type GetStorageReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetStorageReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetStorageReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetStorageReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetStorageReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetStorageReqValidationError) ErrorName() string { return "GetStorageReqValidationError" }

// Error satisfies the builtin error interface
func (e GetStorageReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetStorageReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetStorageReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetStorageReqValidationError{}

// Validate checks the field values on ListStoragesReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListStoragesReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Creator

	// no validation rules for Mode

	{
		tmp := m.GetPager()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListStoragesReqValidationError{
					field:  "Pager",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ListStoragesReqValidationError is the validation error returned by
// ListStoragesReq.Validate if the designated constraints aren't met.
type ListStoragesReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStoragesReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStoragesReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStoragesReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStoragesReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStoragesReqValidationError) ErrorName() string { return "ListStoragesReqValidationError" }

// Error satisfies the builtin error interface
func (e ListStoragesReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStoragesReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStoragesReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStoragesReqValidationError{}

// Validate checks the field values on ListStoragesRes with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListStoragesRes) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetStorages() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListStoragesResValidationError{
						field:  fmt.Sprintf("Storages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetPager()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListStoragesResValidationError{
					field:  "Pager",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ListStoragesResValidationError is the validation error returned by
// ListStoragesRes.Validate if the designated constraints aren't met.
type ListStoragesResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListStoragesResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListStoragesResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListStoragesResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListStoragesResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListStoragesResValidationError) ErrorName() string { return "ListStoragesResValidationError" }

// Error satisfies the builtin error interface
func (e ListStoragesResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListStoragesRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListStoragesResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListStoragesResValidationError{}

// Validate checks the field values on RemoveStorageReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RemoveStorageReq) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetName()) < 1 {
		return RemoveStorageReqValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetCreator()) < 1 {
		return RemoveStorageReqValidationError{
			field:  "Creator",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// RemoveStorageReqValidationError is the validation error returned by
// RemoveStorageReq.Validate if the designated constraints aren't met.
type RemoveStorageReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveStorageReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveStorageReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveStorageReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveStorageReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveStorageReqValidationError) ErrorName() string { return "RemoveStorageReqValidationError" }

// Error satisfies the builtin error interface
func (e RemoveStorageReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveStorageReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveStorageReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveStorageReqValidationError{}

// Validate checks the field values on StorageStatus_UsedBy with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StorageStatus_UsedBy) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetResource()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return StorageStatus_UsedByValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for ReadOnly

	return nil
}

// StorageStatus_UsedByValidationError is the validation error returned by
// StorageStatus_UsedBy.Validate if the designated constraints aren't met.
type StorageStatus_UsedByValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageStatus_UsedByValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageStatus_UsedByValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageStatus_UsedByValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageStatus_UsedByValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageStatus_UsedByValidationError) ErrorName() string {
	return "StorageStatus_UsedByValidationError"
}

// Error satisfies the builtin error interface
func (e StorageStatus_UsedByValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageStatus_UsedBy.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageStatus_UsedByValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageStatus_UsedByValidationError{}

// Validate checks the field values on StorageStatus_Available with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StorageStatus_Available) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ReadOnly

	// no validation rules for ReadWrite

	return nil
}

// StorageStatus_AvailableValidationError is the validation error returned by
// StorageStatus_Available.Validate if the designated constraints aren't met.
type StorageStatus_AvailableValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StorageStatus_AvailableValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StorageStatus_AvailableValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StorageStatus_AvailableValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StorageStatus_AvailableValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StorageStatus_AvailableValidationError) ErrorName() string {
	return "StorageStatus_AvailableValidationError"
}

// Error satisfies the builtin error interface
func (e StorageStatus_AvailableValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStorageStatus_Available.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StorageStatus_AvailableValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StorageStatus_AvailableValidationError{}