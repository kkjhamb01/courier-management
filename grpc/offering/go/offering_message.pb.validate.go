// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: offering_message.proto

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
)

// define the regex for a UUID once up-front
var _offering_message_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on SetCourierLiveLocationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SetCourierLiveLocationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLocation() == nil {
		return SetCourierLiveLocationRequestValidationError{
			field:  "Location",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetCourierLiveLocationRequestValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetTime() == nil {
		return SetCourierLiveLocationRequestValidationError{
			field:  "Time",
			reason: "value is required",
		}
	}

	return nil
}

// SetCourierLiveLocationRequestValidationError is the validation error
// returned by SetCourierLiveLocationRequest.Validate if the designated
// constraints aren't met.
type SetCourierLiveLocationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetCourierLiveLocationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetCourierLiveLocationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetCourierLiveLocationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetCourierLiveLocationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetCourierLiveLocationRequestValidationError) ErrorName() string {
	return "SetCourierLiveLocationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetCourierLiveLocationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetCourierLiveLocationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetCourierLiveLocationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetCourierLiveLocationRequestValidationError{}

// Validate checks the field values on SetCourierLiveLocationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SetCourierLiveLocationResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Successful

	// no validation rules for Message

	return nil
}

// SetCourierLiveLocationResponseValidationError is the validation error
// returned by SetCourierLiveLocationResponse.Validate if the designated
// constraints aren't met.
type SetCourierLiveLocationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetCourierLiveLocationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetCourierLiveLocationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetCourierLiveLocationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetCourierLiveLocationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetCourierLiveLocationResponseValidationError) ErrorName() string {
	return "SetCourierLiveLocationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SetCourierLiveLocationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetCourierLiveLocationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetCourierLiveLocationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetCourierLiveLocationResponseValidationError{}

// Validate checks the field values on GetCourierLiveLocationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetCourierLiveLocationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetCourierId()); err != nil {
		return GetCourierLiveLocationRequestValidationError{
			field:  "CourierId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if m.GetIntervalSeconds() <= 0 {
		return GetCourierLiveLocationRequestValidationError{
			field:  "IntervalSeconds",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

func (m *GetCourierLiveLocationRequest) _validateUuid(uuid string) error {
	if matched := _offering_message_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetCourierLiveLocationRequestValidationError is the validation error
// returned by GetCourierLiveLocationRequest.Validate if the designated
// constraints aren't met.
type GetCourierLiveLocationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCourierLiveLocationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCourierLiveLocationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCourierLiveLocationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCourierLiveLocationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCourierLiveLocationRequestValidationError) ErrorName() string {
	return "GetCourierLiveLocationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCourierLiveLocationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCourierLiveLocationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCourierLiveLocationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCourierLiveLocationRequestValidationError{}

// Validate checks the field values on GetCourierLiveLocationResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetCourierLiveLocationResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCourierLiveLocationResponseValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCourierLiveLocationResponseValidationError{
				field:  "Time",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetCourierLiveLocationResponseValidationError is the validation error
// returned by GetCourierLiveLocationResponse.Validate if the designated
// constraints aren't met.
type GetCourierLiveLocationResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCourierLiveLocationResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCourierLiveLocationResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCourierLiveLocationResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCourierLiveLocationResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCourierLiveLocationResponseValidationError) ErrorName() string {
	return "GetCourierLiveLocationResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCourierLiveLocationResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCourierLiveLocationResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCourierLiveLocationResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCourierLiveLocationResponseValidationError{}

// Validate checks the field values on CourierSubscriptionOnOfferResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *CourierSubscriptionOnOfferResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ResponseType

	switch m.Event.(type) {

	case *CourierSubscriptionOnOfferResponse_NewOfferEvent:

		if v, ok := interface{}(m.GetNewOfferEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CourierSubscriptionOnOfferResponseValidationError{
					field:  "NewOfferEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CourierSubscriptionOnOfferResponse_CancelOfferEvent:

		if v, ok := interface{}(m.GetCancelOfferEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CourierSubscriptionOnOfferResponseValidationError{
					field:  "CancelOfferEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CourierSubscriptionOnOfferResponse_AcceptOfferEvent:

		if v, ok := interface{}(m.GetAcceptOfferEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CourierSubscriptionOnOfferResponseValidationError{
					field:  "AcceptOfferEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CourierSubscriptionOnOfferResponse_AcceptOfferFailedEvent:

		if v, ok := interface{}(m.GetAcceptOfferFailedEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CourierSubscriptionOnOfferResponseValidationError{
					field:  "AcceptOfferFailedEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CourierSubscriptionOnOfferResponse_RejectOfferFailedEvent:

		if v, ok := interface{}(m.GetRejectOfferFailedEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CourierSubscriptionOnOfferResponseValidationError{
					field:  "RejectOfferFailedEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CourierSubscriptionOnOfferResponseValidationError is the validation error
// returned by CourierSubscriptionOnOfferResponse.Validate if the designated
// constraints aren't met.
type CourierSubscriptionOnOfferResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CourierSubscriptionOnOfferResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CourierSubscriptionOnOfferResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CourierSubscriptionOnOfferResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CourierSubscriptionOnOfferResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CourierSubscriptionOnOfferResponseValidationError) ErrorName() string {
	return "CourierSubscriptionOnOfferResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CourierSubscriptionOnOfferResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCourierSubscriptionOnOfferResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CourierSubscriptionOnOfferResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CourierSubscriptionOnOfferResponseValidationError{}

// Validate checks the field values on CustomerSubscriptionOnOfferResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *CustomerSubscriptionOnOfferResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ResponseType

	switch m.Event.(type) {

	case *CustomerSubscriptionOnOfferResponse_RetryOfferEvent:

		if v, ok := interface{}(m.GetRetryOfferEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomerSubscriptionOnOfferResponseValidationError{
					field:  "RetryOfferEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CustomerSubscriptionOnOfferResponse_MaxOfferRetries:

		if v, ok := interface{}(m.GetMaxOfferRetries()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomerSubscriptionOnOfferResponseValidationError{
					field:  "MaxOfferRetries",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CustomerSubscriptionOnOfferResponse_OfferAcceptedEvent:

		if v, ok := interface{}(m.GetOfferAcceptedEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomerSubscriptionOnOfferResponseValidationError{
					field:  "OfferAcceptedEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CustomerSubscriptionOnOfferResponse_OfferCancelledEvent:

		if v, ok := interface{}(m.GetOfferCancelledEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomerSubscriptionOnOfferResponseValidationError{
					field:  "OfferCancelledEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CustomerSubscriptionOnOfferResponse_OfferSentToCouriers:

		if v, ok := interface{}(m.GetOfferSentToCouriers()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomerSubscriptionOnOfferResponseValidationError{
					field:  "OfferSentToCouriers",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *CustomerSubscriptionOnOfferResponse_OfferCreationFailedEvent:

		if v, ok := interface{}(m.GetOfferCreationFailedEvent()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CustomerSubscriptionOnOfferResponseValidationError{
					field:  "OfferCreationFailedEvent",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// CustomerSubscriptionOnOfferResponseValidationError is the validation error
// returned by CustomerSubscriptionOnOfferResponse.Validate if the designated
// constraints aren't met.
type CustomerSubscriptionOnOfferResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomerSubscriptionOnOfferResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomerSubscriptionOnOfferResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomerSubscriptionOnOfferResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomerSubscriptionOnOfferResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomerSubscriptionOnOfferResponseValidationError) ErrorName() string {
	return "CustomerSubscriptionOnOfferResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CustomerSubscriptionOnOfferResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomerSubscriptionOnOfferResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomerSubscriptionOnOfferResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomerSubscriptionOnOfferResponseValidationError{}

// Validate checks the field values on SetCourierLocationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SetCourierLocationRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLocation() == nil {
		return SetCourierLocationRequestValidationError{
			field:  "Location",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SetCourierLocationRequestValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SetCourierLocationRequestValidationError is the validation error returned by
// SetCourierLocationRequest.Validate if the designated constraints aren't met.
type SetCourierLocationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SetCourierLocationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SetCourierLocationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SetCourierLocationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SetCourierLocationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SetCourierLocationRequestValidationError) ErrorName() string {
	return "SetCourierLocationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SetCourierLocationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSetCourierLocationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SetCourierLocationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SetCourierLocationRequestValidationError{}

// Validate checks the field values on GetNearbyCouriersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetNearbyCouriersRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetLocation() == nil {
		return GetNearbyCouriersRequestValidationError{
			field:  "Location",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetNearbyCouriersRequestValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetRadiusMeter() <= 0 {
		return GetNearbyCouriersRequestValidationError{
			field:  "RadiusMeter",
			reason: "value must be greater than 0",
		}
	}

	if _, ok := commonPb.VehicleType_name[int32(m.GetVehicleType())]; !ok {
		return GetNearbyCouriersRequestValidationError{
			field:  "VehicleType",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// GetNearbyCouriersRequestValidationError is the validation error returned by
// GetNearbyCouriersRequest.Validate if the designated constraints aren't met.
type GetNearbyCouriersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNearbyCouriersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNearbyCouriersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNearbyCouriersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNearbyCouriersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNearbyCouriersRequestValidationError) ErrorName() string {
	return "GetNearbyCouriersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetNearbyCouriersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNearbyCouriersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNearbyCouriersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNearbyCouriersRequestValidationError{}

// Validate checks the field values on GetNearbyCouriersResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetNearbyCouriersResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetCouriers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetNearbyCouriersResponseValidationError{
					field:  fmt.Sprintf("Couriers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetNearbyCouriersResponseValidationError is the validation error returned by
// GetNearbyCouriersResponse.Validate if the designated constraints aren't met.
type GetNearbyCouriersResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetNearbyCouriersResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetNearbyCouriersResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetNearbyCouriersResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetNearbyCouriersResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetNearbyCouriersResponseValidationError) ErrorName() string {
	return "GetNearbyCouriersResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetNearbyCouriersResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetNearbyCouriersResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetNearbyCouriersResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetNearbyCouriersResponseValidationError{}

// Validate checks the field values on HadCustomerRideWithCourierRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *HadCustomerRideWithCourierRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetCourierId()); err != nil {
		return HadCustomerRideWithCourierRequestValidationError{
			field:  "CourierId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetCustomerId()); err != nil {
		return HadCustomerRideWithCourierRequestValidationError{
			field:  "CustomerId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if err := m._validateUuid(m.GetOfferId()); err != nil {
		return HadCustomerRideWithCourierRequestValidationError{
			field:  "OfferId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *HadCustomerRideWithCourierRequest) _validateUuid(uuid string) error {
	if matched := _offering_message_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// HadCustomerRideWithCourierRequestValidationError is the validation error
// returned by HadCustomerRideWithCourierRequest.Validate if the designated
// constraints aren't met.
type HadCustomerRideWithCourierRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HadCustomerRideWithCourierRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HadCustomerRideWithCourierRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HadCustomerRideWithCourierRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HadCustomerRideWithCourierRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HadCustomerRideWithCourierRequestValidationError) ErrorName() string {
	return "HadCustomerRideWithCourierRequestValidationError"
}

// Error satisfies the builtin error interface
func (e HadCustomerRideWithCourierRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHadCustomerRideWithCourierRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HadCustomerRideWithCourierRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HadCustomerRideWithCourierRequestValidationError{}

// Validate checks the field values on HadCustomerRideWithCourierResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *HadCustomerRideWithCourierResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for HadRide

	return nil
}

// HadCustomerRideWithCourierResponseValidationError is the validation error
// returned by HadCustomerRideWithCourierResponse.Validate if the designated
// constraints aren't met.
type HadCustomerRideWithCourierResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HadCustomerRideWithCourierResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HadCustomerRideWithCourierResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HadCustomerRideWithCourierResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HadCustomerRideWithCourierResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HadCustomerRideWithCourierResponseValidationError) ErrorName() string {
	return "HadCustomerRideWithCourierResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HadCustomerRideWithCourierResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHadCustomerRideWithCourierResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HadCustomerRideWithCourierResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HadCustomerRideWithCourierResponseValidationError{}

// Validate checks the field values on GetOfferCourierAndCustomerRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *GetOfferCourierAndCustomerRequest) Validate() error {
	if m == nil {
		return nil
	}

	if err := m._validateUuid(m.GetOfferId()); err != nil {
		return GetOfferCourierAndCustomerRequestValidationError{
			field:  "OfferId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	return nil
}

func (m *GetOfferCourierAndCustomerRequest) _validateUuid(uuid string) error {
	if matched := _offering_message_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetOfferCourierAndCustomerRequestValidationError is the validation error
// returned by GetOfferCourierAndCustomerRequest.Validate if the designated
// constraints aren't met.
type GetOfferCourierAndCustomerRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOfferCourierAndCustomerRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOfferCourierAndCustomerRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOfferCourierAndCustomerRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOfferCourierAndCustomerRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOfferCourierAndCustomerRequestValidationError) ErrorName() string {
	return "GetOfferCourierAndCustomerRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetOfferCourierAndCustomerRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOfferCourierAndCustomerRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOfferCourierAndCustomerRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOfferCourierAndCustomerRequestValidationError{}

// Validate checks the field values on GetOfferCourierAndCustomerResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *GetOfferCourierAndCustomerResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CourierId

	// no validation rules for CustomerId

	return nil
}

// GetOfferCourierAndCustomerResponseValidationError is the validation error
// returned by GetOfferCourierAndCustomerResponse.Validate if the designated
// constraints aren't met.
type GetOfferCourierAndCustomerResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOfferCourierAndCustomerResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOfferCourierAndCustomerResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOfferCourierAndCustomerResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOfferCourierAndCustomerResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOfferCourierAndCustomerResponseValidationError) ErrorName() string {
	return "GetOfferCourierAndCustomerResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetOfferCourierAndCustomerResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOfferCourierAndCustomerResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOfferCourierAndCustomerResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOfferCourierAndCustomerResponseValidationError{}
