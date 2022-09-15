// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pricing_message.proto

package pricingPb

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

	"google.golang.org/protobuf/types/known/anypb"

	commonPb "gitlab.artin.ai/backend/courier-management/grpc/common/go"
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

	_ = commonPb.VehicleType(0)

	_ = commonPb.VehicleType(0)
)

// Validate checks the field values on CalculateCourierPriceRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CalculateCourierPriceRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for VehicleType

	// no validation rules for RequiredWorkers

	if m.GetSource() == nil {
		return CalculateCourierPriceRequestValidationError{
			field:  "Source",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CalculateCourierPriceRequestValidationError{
				field:  "Source",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetDestinations()) < 1 {
		return CalculateCourierPriceRequestValidationError{
			field:  "Destinations",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetDestinations() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CalculateCourierPriceRequestValidationError{
					field:  fmt.Sprintf("Destinations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CalculateCourierPriceRequestValidationError is the validation error returned
// by CalculateCourierPriceRequest.Validate if the designated constraints
// aren't met.
type CalculateCourierPriceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalculateCourierPriceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalculateCourierPriceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalculateCourierPriceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalculateCourierPriceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalculateCourierPriceRequestValidationError) ErrorName() string {
	return "CalculateCourierPriceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CalculateCourierPriceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalculateCourierPriceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalculateCourierPriceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalculateCourierPriceRequestValidationError{}

// Validate checks the field values on CalculateCourierPriceResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CalculateCourierPriceResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EstimatedDuration

	// no validation rules for EstimatedDistance

	// no validation rules for Amount

	// no validation rules for Currency

	return nil
}

// CalculateCourierPriceResponseValidationError is the validation error
// returned by CalculateCourierPriceResponse.Validate if the designated
// constraints aren't met.
type CalculateCourierPriceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CalculateCourierPriceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CalculateCourierPriceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CalculateCourierPriceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CalculateCourierPriceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CalculateCourierPriceResponseValidationError) ErrorName() string {
	return "CalculateCourierPriceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CalculateCourierPriceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCalculateCourierPriceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CalculateCourierPriceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CalculateCourierPriceResponseValidationError{}

// Validate checks the field values on ReviewCourierPriceRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReviewCourierPriceRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequiredWorkers

	if m.GetSource() == nil {
		return ReviewCourierPriceRequestValidationError{
			field:  "Source",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ReviewCourierPriceRequestValidationError{
				field:  "Source",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetDestinations()) < 1 {
		return ReviewCourierPriceRequestValidationError{
			field:  "Destinations",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetDestinations() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReviewCourierPriceRequestValidationError{
					field:  fmt.Sprintf("Destinations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ReviewCourierPriceRequestValidationError is the validation error returned by
// ReviewCourierPriceRequest.Validate if the designated constraints aren't met.
type ReviewCourierPriceRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReviewCourierPriceRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReviewCourierPriceRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReviewCourierPriceRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReviewCourierPriceRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReviewCourierPriceRequestValidationError) ErrorName() string {
	return "ReviewCourierPriceRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ReviewCourierPriceRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReviewCourierPriceRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReviewCourierPriceRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReviewCourierPriceRequestValidationError{}

// Validate checks the field values on ReviewCourierPriceResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ReviewCourierPriceResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EstimatedDuration

	// no validation rules for EstimatedDistance

	for idx, item := range m.GetPrices() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ReviewCourierPriceResponseValidationError{
					field:  fmt.Sprintf("Prices[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ReviewCourierPriceResponseValidationError is the validation error returned
// by ReviewCourierPriceResponse.Validate if the designated constraints aren't met.
type ReviewCourierPriceResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReviewCourierPriceResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReviewCourierPriceResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReviewCourierPriceResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReviewCourierPriceResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReviewCourierPriceResponseValidationError) ErrorName() string {
	return "ReviewCourierPriceResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ReviewCourierPriceResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReviewCourierPriceResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReviewCourierPriceResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReviewCourierPriceResponseValidationError{}

// Validate checks the field values on ReviewCourierPriceResponse_Price with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ReviewCourierPriceResponse_Price) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for VehicleType

	// no validation rules for Amount

	// no validation rules for Currency

	return nil
}

// ReviewCourierPriceResponse_PriceValidationError is the validation error
// returned by ReviewCourierPriceResponse_Price.Validate if the designated
// constraints aren't met.
type ReviewCourierPriceResponse_PriceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ReviewCourierPriceResponse_PriceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ReviewCourierPriceResponse_PriceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ReviewCourierPriceResponse_PriceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ReviewCourierPriceResponse_PriceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ReviewCourierPriceResponse_PriceValidationError) ErrorName() string {
	return "ReviewCourierPriceResponse_PriceValidationError"
}

// Error satisfies the builtin error interface
func (e ReviewCourierPriceResponse_PriceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sReviewCourierPriceResponse_Price.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ReviewCourierPriceResponse_PriceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ReviewCourierPriceResponse_PriceValidationError{}