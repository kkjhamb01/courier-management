// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: offering_model.proto

package offeringPb

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

	"github.com/golang/protobuf/ptypes"

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
	_ = ptypes.DynamicAny{}

	_ = commonPb.VehicleType(0)

	_ = commonPb.CourierStatus(0)
)

// define the regex for a UUID once up-front
var _offering_model_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Offer with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Offer) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetId()); err != nil {
		return OfferValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetCustomerId()); err != nil {
		return OfferValidationError{
			field:  "CustomerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if v, ok := interface{}(m.GetSource()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OfferValidationError{
				field:  "Source",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetDestinations() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OfferValidationError{
					field:  fmt.Sprintf("Destinations[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for VehicleType

	// no validation rules for RequiredWorkers

	// no validation rules for Price

	// no validation rules for Currency

	return nil
}

func (m *Offer) _validateUuid(uuid string) error {
	if matched := _offering_model_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// OfferValidationError is the validation error returned by Offer.Validate if
// the designated constraints aren't met.
type OfferValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OfferValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OfferValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OfferValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OfferValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OfferValidationError) ErrorName() string { return "OfferValidationError" }

// Error satisfies the builtin error interface
func (e OfferValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOffer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OfferValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OfferValidationError{}

// Validate checks the field values on Location with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Location) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PhoneNumber

	// no validation rules for AddressDetails

	// no validation rules for FullName

	// no validation rules for Lat

	// no validation rules for Lon

	// no validation rules for Order

	// no validation rules for CourierInstructions

	return nil
}

// LocationValidationError is the validation error returned by
// Location.Validate if the designated constraints aren't met.
type LocationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocationValidationError) ErrorName() string { return "LocationValidationError" }

// Error satisfies the builtin error interface
func (e LocationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocation.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocationValidationError{}

// Validate checks the field values on NewOfferEvent with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *NewOfferEvent) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetOffer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NewOfferEventValidationError{
				field:  "Offer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if err := m._validateUuid(m.GetCourierId()); err != nil {
		return NewOfferEventValidationError{
			field:  "CourierId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if v, ok := interface{}(m.GetCourierResponseTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NewOfferEventValidationError{
				field:  "CourierResponseTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RequesterName

	// no validation rules for RequesterPhone

	// no validation rules for Desc

	// no validation rules for DistanceMeters

	if v, ok := interface{}(m.GetDuration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NewOfferEventValidationError{
				field:  "Duration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

func (m *NewOfferEvent) _validateUuid(uuid string) error {
	if matched := _offering_model_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// NewOfferEventValidationError is the validation error returned by
// NewOfferEvent.Validate if the designated constraints aren't met.
type NewOfferEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewOfferEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewOfferEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewOfferEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewOfferEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewOfferEventValidationError) ErrorName() string { return "NewOfferEventValidationError" }

// Error satisfies the builtin error interface
func (e NewOfferEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewOfferEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewOfferEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewOfferEventValidationError{}

// Validate checks the field values on NewOfferSentToCouriersEvent with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *NewOfferSentToCouriersEvent) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetOffer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NewOfferSentToCouriersEventValidationError{
				field:  "Offer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// NewOfferSentToCouriersEventValidationError is the validation error returned
// by NewOfferSentToCouriersEvent.Validate if the designated constraints
// aren't met.
type NewOfferSentToCouriersEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NewOfferSentToCouriersEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NewOfferSentToCouriersEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NewOfferSentToCouriersEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NewOfferSentToCouriersEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NewOfferSentToCouriersEventValidationError) ErrorName() string {
	return "NewOfferSentToCouriersEventValidationError"
}

// Error satisfies the builtin error interface
func (e NewOfferSentToCouriersEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNewOfferSentToCouriersEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NewOfferSentToCouriersEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NewOfferSentToCouriersEventValidationError{}

// Validate checks the field values on OfferCancelledEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OfferCancelledEvent) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetOfferId()); err != nil {
		return OfferCancelledEventValidationError{
			field:  "OfferId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetCustomerId()); err != nil {
		return OfferCancelledEventValidationError{
			field:  "CustomerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetUpdatedBy()); err != nil {
		return OfferCancelledEventValidationError{
			field:  "UpdatedBy",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	// no validation rules for CancelReason

	// no validation rules for CancelledBy

	return nil
}

func (m *OfferCancelledEvent) _validateUuid(uuid string) error {
	if matched := _offering_model_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// OfferCancelledEventValidationError is the validation error returned by
// OfferCancelledEvent.Validate if the designated constraints aren't met.
type OfferCancelledEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OfferCancelledEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OfferCancelledEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OfferCancelledEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OfferCancelledEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OfferCancelledEventValidationError) ErrorName() string {
	return "OfferCancelledEventValidationError"
}

// Error satisfies the builtin error interface
func (e OfferCancelledEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOfferCancelledEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OfferCancelledEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OfferCancelledEventValidationError{}

// Validate checks the field values on OfferAcceptedEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OfferAcceptedEvent) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetOfferId()); err != nil {
		return OfferAcceptedEventValidationError{
			field:  "OfferId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetCustomerId()); err != nil {
		return OfferAcceptedEventValidationError{
			field:  "CustomerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetCourierId()); err != nil {
		return OfferAcceptedEventValidationError{
			field:  "CourierId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	// no validation rules for Desc

	return nil
}

func (m *OfferAcceptedEvent) _validateUuid(uuid string) error {
	if matched := _offering_model_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// OfferAcceptedEventValidationError is the validation error returned by
// OfferAcceptedEvent.Validate if the designated constraints aren't met.
type OfferAcceptedEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OfferAcceptedEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OfferAcceptedEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OfferAcceptedEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OfferAcceptedEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OfferAcceptedEventValidationError) ErrorName() string {
	return "OfferAcceptedEventValidationError"
}

// Error satisfies the builtin error interface
func (e OfferAcceptedEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOfferAcceptedEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OfferAcceptedEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OfferAcceptedEventValidationError{}

// Validate checks the field values on OfferRejectedEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OfferRejectedEvent) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetOfferId()); err != nil {
		return OfferRejectedEventValidationError{
			field:  "OfferId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetCourierId()); err != nil {
		return OfferRejectedEventValidationError{
			field:  "CourierId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	// no validation rules for Desc

	return nil
}

func (m *OfferRejectedEvent) _validateUuid(uuid string) error {
	if matched := _offering_model_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// OfferRejectedEventValidationError is the validation error returned by
// OfferRejectedEvent.Validate if the designated constraints aren't met.
type OfferRejectedEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OfferRejectedEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OfferRejectedEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OfferRejectedEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OfferRejectedEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OfferRejectedEventValidationError) ErrorName() string {
	return "OfferRejectedEventValidationError"
}

// Error satisfies the builtin error interface
func (e OfferRejectedEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOfferRejectedEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OfferRejectedEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OfferRejectedEventValidationError{}

// Validate checks the field values on MaxOfferRetriesEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MaxOfferRetriesEvent) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetOfferId()); err != nil {
		return MaxOfferRetriesEventValidationError{
			field:  "OfferId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetCustomerId()); err != nil {
		return MaxOfferRetriesEventValidationError{
			field:  "CustomerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	// no validation rules for Desc

	return nil
}

func (m *MaxOfferRetriesEvent) _validateUuid(uuid string) error {
	if matched := _offering_model_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// MaxOfferRetriesEventValidationError is the validation error returned by
// MaxOfferRetriesEvent.Validate if the designated constraints aren't met.
type MaxOfferRetriesEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MaxOfferRetriesEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MaxOfferRetriesEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MaxOfferRetriesEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MaxOfferRetriesEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MaxOfferRetriesEventValidationError) ErrorName() string {
	return "MaxOfferRetriesEventValidationError"
}

// Error satisfies the builtin error interface
func (e MaxOfferRetriesEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMaxOfferRetriesEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MaxOfferRetriesEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MaxOfferRetriesEventValidationError{}

// Validate checks the field values on RetryOfferRequestEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RetryOfferRequestEvent) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetOffer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RetryOfferRequestEventValidationError{
				field:  "Offer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Desc

	return nil
}

// RetryOfferRequestEventValidationError is the validation error returned by
// RetryOfferRequestEvent.Validate if the designated constraints aren't met.
type RetryOfferRequestEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RetryOfferRequestEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RetryOfferRequestEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RetryOfferRequestEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RetryOfferRequestEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RetryOfferRequestEventValidationError) ErrorName() string {
	return "RetryOfferRequestEventValidationError"
}

// Error satisfies the builtin error interface
func (e RetryOfferRequestEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRetryOfferRequestEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RetryOfferRequestEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RetryOfferRequestEventValidationError{}

// Validate checks the field values on OfferCreationFailedEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *OfferCreationFailedEvent) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetOffer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OfferCreationFailedEventValidationError{
				field:  "Offer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Msg

	return nil
}

// OfferCreationFailedEventValidationError is the validation error returned by
// OfferCreationFailedEvent.Validate if the designated constraints aren't met.
type OfferCreationFailedEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OfferCreationFailedEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OfferCreationFailedEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OfferCreationFailedEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OfferCreationFailedEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OfferCreationFailedEventValidationError) ErrorName() string {
	return "OfferCreationFailedEventValidationError"
}

// Error satisfies the builtin error interface
func (e OfferCreationFailedEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOfferCreationFailedEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OfferCreationFailedEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OfferCreationFailedEventValidationError{}

// Validate checks the field values on AcceptOfferFailedEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *AcceptOfferFailedEvent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OfferId

	// no validation rules for CustomerId

	// no validation rules for CourierId

	// no validation rules for Msg

	return nil
}

// AcceptOfferFailedEventValidationError is the validation error returned by
// AcceptOfferFailedEvent.Validate if the designated constraints aren't met.
type AcceptOfferFailedEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AcceptOfferFailedEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AcceptOfferFailedEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AcceptOfferFailedEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AcceptOfferFailedEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AcceptOfferFailedEventValidationError) ErrorName() string {
	return "AcceptOfferFailedEventValidationError"
}

// Error satisfies the builtin error interface
func (e AcceptOfferFailedEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAcceptOfferFailedEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AcceptOfferFailedEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AcceptOfferFailedEventValidationError{}

// Validate checks the field values on RejectOfferFailedEvent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RejectOfferFailedEvent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for OfferId

	// no validation rules for CustomerId

	// no validation rules for CourierId

	// no validation rules for Msg

	return nil
}

// RejectOfferFailedEventValidationError is the validation error returned by
// RejectOfferFailedEvent.Validate if the designated constraints aren't met.
type RejectOfferFailedEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RejectOfferFailedEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RejectOfferFailedEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RejectOfferFailedEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RejectOfferFailedEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RejectOfferFailedEventValidationError) ErrorName() string {
	return "RejectOfferFailedEventValidationError"
}

// Error satisfies the builtin error interface
func (e RejectOfferFailedEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRejectOfferFailedEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RejectOfferFailedEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RejectOfferFailedEventValidationError{}

// Validate checks the field values on CourierStatusLog with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CourierStatusLog) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CourierId

	// no validation rules for Status

	if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CourierStatusLogValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// CourierStatusLogValidationError is the validation error returned by
// CourierStatusLog.Validate if the designated constraints aren't met.
type CourierStatusLogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CourierStatusLogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CourierStatusLogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CourierStatusLogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CourierStatusLogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CourierStatusLogValidationError) ErrorName() string { return "CourierStatusLogValidationError" }

// Error satisfies the builtin error interface
func (e CourierStatusLogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCourierStatusLog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CourierStatusLogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CourierStatusLogValidationError{}
