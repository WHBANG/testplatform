// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: distributed.proto

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
var _distributed_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on DistributedSpec with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DistributedSpec) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := DistributedFramework_name[int32(m.GetFramework())]; !ok {
		return DistributedSpecValidationError{
			field:  "Framework",
			reason: "value must be one of the defined enum values",
		}
	}

	for idx, item := range m.GetReplicas() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return DistributedSpecValidationError{
						field:  fmt.Sprintf("Replicas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	// no validation rules for EnableLogger

	return nil
}

// DistributedSpecValidationError is the validation error returned by
// DistributedSpec.Validate if the designated constraints aren't met.
type DistributedSpecValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DistributedSpecValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DistributedSpecValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DistributedSpecValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DistributedSpecValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DistributedSpecValidationError) ErrorName() string { return "DistributedSpecValidationError" }

// Error satisfies the builtin error interface
func (e DistributedSpecValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDistributedSpec.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DistributedSpecValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DistributedSpecValidationError{}

// Validate checks the field values on DistributedReplica with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DistributedReplica) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := DistributedReplicaType_name[int32(m.GetType())]; !ok {
		return DistributedReplicaValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
	}

	// no validation rules for Replicas

	{
		tmp := m.GetPackage()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return DistributedReplicaValidationError{
					field:  "Package",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// DistributedReplicaValidationError is the validation error returned by
// DistributedReplica.Validate if the designated constraints aren't met.
type DistributedReplicaValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DistributedReplicaValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DistributedReplicaValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DistributedReplicaValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DistributedReplicaValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DistributedReplicaValidationError) ErrorName() string {
	return "DistributedReplicaValidationError"
}

// Error satisfies the builtin error interface
func (e DistributedReplicaValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDistributedReplica.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DistributedReplicaValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DistributedReplicaValidationError{}

// Validate checks the field values on DistributedStatus with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DistributedStatus) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetReplicas() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return DistributedStatusValidationError{
						field:  fmt.Sprintf("Replicas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// DistributedStatusValidationError is the validation error returned by
// DistributedStatus.Validate if the designated constraints aren't met.
type DistributedStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DistributedStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DistributedStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DistributedStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DistributedStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DistributedStatusValidationError) ErrorName() string {
	return "DistributedStatusValidationError"
}

// Error satisfies the builtin error interface
func (e DistributedStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDistributedStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DistributedStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DistributedStatusValidationError{}

// Validate checks the field values on ReplicaStatus with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ReplicaStatus) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Type

	// no validation rules for Replicas

	{
		tmp := m.GetResource()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ReplicaStatusValidationError{
					field:  "Resource",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ReplicaStatusValidationError is the validation error returned by
// ReplicaStatus.Validate if the designated constraints aren't met.
type ReplicaStatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReplicaStatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReplicaStatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReplicaStatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReplicaStatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReplicaStatusValidationError) ErrorName() string { return "ReplicaStatusValidationError" }

// Error satisfies the builtin error interface
func (e ReplicaStatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReplicaStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReplicaStatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReplicaStatusValidationError{}