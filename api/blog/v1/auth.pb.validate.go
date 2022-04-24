// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/blog/v1/auth.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on RegisterReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterReqMultiError, or
// nil if none found.
func (m *RegisterReq) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	if len(errors) > 0 {
		return RegisterReqMultiError(errors)
	}

	return nil
}

// RegisterReqMultiError is an error wrapping multiple validation errors
// returned by RegisterReq.ValidateAll() if the designated constraints aren't met.
type RegisterReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterReqMultiError) AllErrors() []error { return m }

// RegisterReqValidationError is the validation error returned by
// RegisterReq.Validate if the designated constraints aren't met.
type RegisterReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterReqValidationError) ErrorName() string { return "RegisterReqValidationError" }

// Error satisfies the builtin error interface
func (e RegisterReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterReqValidationError{}

// Validate checks the field values on RegisterReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterReplyMultiError, or
// nil if none found.
func (m *RegisterReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return RegisterReplyMultiError(errors)
	}

	return nil
}

// RegisterReplyMultiError is an error wrapping multiple validation errors
// returned by RegisterReply.ValidateAll() if the designated constraints
// aren't met.
type RegisterReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterReplyMultiError) AllErrors() []error { return m }

// RegisterReplyValidationError is the validation error returned by
// RegisterReply.Validate if the designated constraints aren't met.
type RegisterReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterReplyValidationError) ErrorName() string { return "RegisterReplyValidationError" }

// Error satisfies the builtin error interface
func (e RegisterReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterReplyValidationError{}

// Validate checks the field values on LoginReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReqMultiError, or nil
// if none found.
func (m *LoginReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Password

	if len(errors) > 0 {
		return LoginReqMultiError(errors)
	}

	return nil
}

// LoginReqMultiError is an error wrapping multiple validation errors returned
// by LoginReq.ValidateAll() if the designated constraints aren't met.
type LoginReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReqMultiError) AllErrors() []error { return m }

// LoginReqValidationError is the validation error returned by
// LoginReq.Validate if the designated constraints aren't met.
type LoginReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReqValidationError) ErrorName() string { return "LoginReqValidationError" }

// Error satisfies the builtin error interface
func (e LoginReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReqValidationError{}

// Validate checks the field values on LoginReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReplyMultiError, or
// nil if none found.
func (m *LoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return LoginReplyMultiError(errors)
	}

	return nil
}

// LoginReplyMultiError is an error wrapping multiple validation errors
// returned by LoginReply.ValidateAll() if the designated constraints aren't met.
type LoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReplyMultiError) AllErrors() []error { return m }

// LoginReplyValidationError is the validation error returned by
// LoginReply.Validate if the designated constraints aren't met.
type LoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReplyValidationError) ErrorName() string { return "LoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReplyValidationError{}

// Validate checks the field values on LogoutReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LogoutReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogoutReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LogoutReqMultiError, or nil
// if none found.
func (m *LogoutReq) ValidateAll() error {
	return m.validate(true)
}

func (m *LogoutReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LogoutReqMultiError(errors)
	}

	return nil
}

// LogoutReqMultiError is an error wrapping multiple validation errors returned
// by LogoutReq.ValidateAll() if the designated constraints aren't met.
type LogoutReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogoutReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogoutReqMultiError) AllErrors() []error { return m }

// LogoutReqValidationError is the validation error returned by
// LogoutReq.Validate if the designated constraints aren't met.
type LogoutReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogoutReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogoutReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogoutReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogoutReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogoutReqValidationError) ErrorName() string { return "LogoutReqValidationError" }

// Error satisfies the builtin error interface
func (e LogoutReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogoutReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogoutReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogoutReqValidationError{}

// Validate checks the field values on LogoutReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LogoutReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LogoutReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LogoutReplyMultiError, or
// nil if none found.
func (m *LogoutReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LogoutReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return LogoutReplyMultiError(errors)
	}

	return nil
}

// LogoutReplyMultiError is an error wrapping multiple validation errors
// returned by LogoutReply.ValidateAll() if the designated constraints aren't met.
type LogoutReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LogoutReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LogoutReplyMultiError) AllErrors() []error { return m }

// LogoutReplyValidationError is the validation error returned by
// LogoutReply.Validate if the designated constraints aren't met.
type LogoutReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LogoutReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LogoutReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LogoutReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LogoutReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LogoutReplyValidationError) ErrorName() string { return "LogoutReplyValidationError" }

// Error satisfies the builtin error interface
func (e LogoutReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLogoutReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LogoutReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LogoutReplyValidationError{}
