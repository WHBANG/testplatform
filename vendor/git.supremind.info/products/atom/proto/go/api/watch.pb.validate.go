// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: watch.proto

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
var _watch_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on IsWatchingRes with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *IsWatchingRes) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Watching

	return nil
}

// IsWatchingResValidationError is the validation error returned by
// IsWatchingRes.Validate if the designated constraints aren't met.
type IsWatchingResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IsWatchingResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IsWatchingResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IsWatchingResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IsWatchingResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IsWatchingResValidationError) ErrorName() string { return "IsWatchingResValidationError" }

// Error satisfies the builtin error interface
func (e IsWatchingResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIsWatchingRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IsWatchingResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IsWatchingResValidationError{}

// Validate checks the field values on ListWatchingReq with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListWatchingReq) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Kind

	{
		tmp := m.GetPager()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ListWatchingReqValidationError{
					field:  "Pager",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ListWatchingReqValidationError is the validation error returned by
// ListWatchingReq.Validate if the designated constraints aren't met.
type ListWatchingReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWatchingReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWatchingReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWatchingReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWatchingReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWatchingReqValidationError) ErrorName() string { return "ListWatchingReqValidationError" }

// Error satisfies the builtin error interface
func (e ListWatchingReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWatchingReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWatchingReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWatchingReqValidationError{}

// Validate checks the field values on ListWatchingRes with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListWatchingRes) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetResources() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListWatchingResValidationError{
						field:  fmt.Sprintf("Resources[%v]", idx),
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
				return ListWatchingResValidationError{
					field:  "Pager",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ListWatchingResValidationError is the validation error returned by
// ListWatchingRes.Validate if the designated constraints aren't met.
type ListWatchingResValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListWatchingResValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListWatchingResValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListWatchingResValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListWatchingResValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListWatchingResValidationError) ErrorName() string { return "ListWatchingResValidationError" }

// Error satisfies the builtin error interface
func (e ListWatchingResValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListWatchingRes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListWatchingResValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListWatchingResValidationError{}
