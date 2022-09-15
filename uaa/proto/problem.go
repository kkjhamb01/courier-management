package proto

import (
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type Code uint32

const (
	OK Code = 0
	Canceled Code = 1
	Unknown Code = 2
	InvalidArgument Code = 3
	DeadlineExceeded Code = 4
	NotFound Code = 5
	AlreadyExists Code = 6
	PermissionDenied Code = 7
	ResourceExhausted Code = 8
	FailedPrecondition Code = 9
	Aborted Code = 10
	OutOfRange Code = 11
	Unimplemented Code = 12
	Internal Code = 13
	Unavailable Code = 14
	DataLoss Code = 15
	Unauthenticated Code = 16
	ExceedsMaximumRetry Code = 17
	InvalidPincode Code = 18
	InvalidCode Code = 19
	InvalidRegistrationNumber Code = 20
	NoDetailsHeldByDvla Code = 21
	TokenIsExpired Code = 22
	PromotionWillStartInTheFuture Code = 23
	PromotionIsExpired Code = 24
	PromotionNotAvailable Code = 25
	ZeroDiscount Code = 26
	DiscountIsGreaterThanTotalPayment Code = 27
	DeliveryInvalidStateTransition Code = 28
	RequestIsNotAvailableToAccept Code = 29
	RequestIsInPickedupState Code = 30
	InvalidLocationOrigin Code = 31
	InvalidDestinationOrder Code = 32
	LastDestinationIsNotDelivered Code = 33
	_maxCode = 34
)

var strToCode = map[string]Code{
	`"OK"`: OK,
	`"CANCELLED"`:/* [sic] */ Canceled,
	`"UNKNOWN"`:             Unknown,
	`"INVALID_ARGUMENT"`:    InvalidArgument,
	`"DEADLINE_EXCEEDED"`:   DeadlineExceeded,
	`"NOT_FOUND"`:           NotFound,
	`"ALREADY_EXISTS"`:      AlreadyExists,
	`"PERMISSION_DENIED"`:   PermissionDenied,
	`"RESOURCE_EXHAUSTED"`:  ResourceExhausted,
	`"FAILED_PRECONDITION"`: FailedPrecondition,
	`"ABORTED"`:             Aborted,
	`"OUT_OF_RANGE"`:        OutOfRange,
	`"UNIMPLEMENTED"`:       Unimplemented,
	`"INTERNAL"`:            Internal,
	`"UNAVAILABLE"`:         Unavailable,
	`"DATA_LOSS"`:           DataLoss,
	`"UNAUTHENTICATED"`:     Unauthenticated,
	`"EXCEEDS_MAXIMUM_RETRY"`: ExceedsMaximumRetry,
	`"INVALID_PINCODE"`: InvalidPincode,
	`"INVALID_CODE"`: InvalidCode,
	`"INVALID_REGISTRATION_NUMBER"`: InvalidRegistrationNumber,
	`"NO_DETAILS_HELD_BY_DVLA"`: NoDetailsHeldByDvla,
	`"TOKEN_IS_EXPIRED"`: TokenIsExpired,
	`"PROMOTION_WILL_START_IN_THE_FUTURE"`: PromotionWillStartInTheFuture,
	`"PROMOTION_IS_EXPIRED"`: PromotionIsExpired,
	`"PROMOTION_NOT_AVAILABLE"`: PromotionNotAvailable,
	`"ZERO_DISCOUNT"`: ZeroDiscount,
	`"DISCOUNT_IS_GREATER_THAN_TOTAL_PAYMENT"`: DiscountIsGreaterThanTotalPayment,
	`"INVALID_STATE_TRANSITION"`: DeliveryInvalidStateTransition,
	`"REQUEST_IS_NOT_AVAILABLE_TO_ACCEPT"`: RequestIsNotAvailableToAccept,
	`"REQUEST_IS_IN_PICKED_UP_STATE"`: RequestIsInPickedupState,
	`"INVALID_LOCATION_ORIGIN"`: InvalidLocationOrigin,
	`"INVALID_DESTINATION_ORDER"`: InvalidDestinationOrder,
	`"LAST_DESTINATION_IS_NOT_DELIVERED"`: LastDestinationIsNotDelivered,
}

var codeToStr = map[Code]string{
	OK	:	"OK",
	Canceled	:	"CANCELLED",
	Unknown	:	"UNKNOWN",
	InvalidArgument	:	"INVALID_ARGUMENT",
	DeadlineExceeded	:	"DEADLINE_EXCEEDED",
	NotFound	:	"NOT_FOUND",
	AlreadyExists	:	"ALREADY_EXISTS",
	PermissionDenied	:	"PERMISSION_DENIED",
	ResourceExhausted	:	"RESOURCE_EXHAUSTED",
	FailedPrecondition	:	"FAILED_PRECONDITION",
	Aborted	:	"ABORTED",
	OutOfRange	:	"OUT_OF_RANGE",
	Unimplemented	:	"UNIMPLEMENTED",
	Internal	:	"INTERNAL",
	Unavailable	:	"UNAVAILABLE",
	DataLoss	:	"DATA_LOSS",
	Unauthenticated	:	"UNAUTHENTICATED",
	ExceedsMaximumRetry	:	"EXCEEDS_MAXIMUM_RETRY",
	InvalidPincode	:	"INVALID_PINCODE",
	InvalidCode	:	"INVALID_CODE",
	InvalidRegistrationNumber	:	"INVALID_REGISTRATION_NUMBER",
	NoDetailsHeldByDvla	:	"NO_DETAILS_HELD_BY_DVLA",
	TokenIsExpired	:	"TOKEN_IS_EXPIRED",
	PromotionWillStartInTheFuture: `"PROMOTION_WILL_START_IN_THE_FUTURE"`,
	PromotionIsExpired: `"PROMOTION_IS_EXPIRED"`,
	PromotionNotAvailable: `"PROMOTION_NOT_AVAILABLE"`,
	ZeroDiscount: `"ZERO_DISCOUNT"`,
	DiscountIsGreaterThanTotalPayment: `"DISCOUNT_IS_GREATER_THAN_TOTAL_PAYMENT"`,
	DeliveryInvalidStateTransition: `"INVALID_STATE_TRANSITION"`,
	RequestIsNotAvailableToAccept: `"REQUEST_IS_NOT_AVAILABLE_TO_ACCEPT"`,
	RequestIsInPickedupState: `"REQUEST_IS_IN_PICKED_UP_STATE"`,
	InvalidLocationOrigin: `"INVALID_LOCATION_ORIGIN"`,
	InvalidDestinationOrder: `"INVALID_DESTINATION_ORDER"`,
	LastDestinationIsNotDelivered: `"LAST_DESTINATION_IS_NOT_DELIVERED"`,
}

func (c *Code) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	if c == nil {
		return fmt.Errorf("nil receiver passed to UnmarshalJSON")
	}

	if ci, err := strconv.ParseUint(string(b), 10, 32); err == nil {
		if ci >= _maxCode {
			return fmt.Errorf("invalid code: %q", ci)
		}

		*c = Code(ci)
		return nil
	}

	if jc, ok := strToCode[string(b)]; ok {
		*c = jc
		return nil
	}
	return fmt.Errorf("invalid code: %q", string(b))
}

func (c Code) Error(err error) error {
	if c == OK{
		return nil
	}
	// list of errors:
	// https://mariadb.com/kb/en/mariadb-error-codes/
	if sqlerr, ok := err.(*mysql.MySQLError); ok{
		switch sqlerr.Number {
		case 1406:
			return InvalidArgument.ErrorMsg(sqlerr.Message)
		case 1062:
			return AlreadyExists.ErrorMsg(sqlerr.Message)
		}
	}
	s, _ := json.Marshal(&problem{
		Code: c,
		Desc: codeToStr[c],
		Error: err.Error(),
	})
	return status.New(codes.Code(c), string(s)).Err()
}

func (c Code) ErrorMsg(msg string) error {
	if c == OK{
		return nil
	}
	s, _ := json.Marshal(&problem{
		Code: c,
		Desc: codeToStr[c],
		Error: msg,
	})
	return status.New(codes.Code(c), string(s)).Err()
}

func (c Code) ErrorNoMsg() error {
	if c == OK{
		return nil
	}
	s, _ := json.Marshal(&problem{
		Code: c,
		Desc: codeToStr[c],
	})
	return status.New(codes.Code(c), string(s)).Err()
}

type problem struct {
	Code Code
	Error string
	Desc string
}