// Code generated by protoc-gen-validate
// source: envoy/api/v2/cluster/circuit_breaker.proto
// DO NOT EDIT!!!

package cluster

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

// Validate checks the field values on CircuitBreakers with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CircuitBreakers) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetThresholds() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CircuitBreakersValidationError{
					Field:  fmt.Sprintf("Thresholds[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// CircuitBreakersValidationError is the validation error returned by
// CircuitBreakers.Validate if the designated constraints aren't met.
type CircuitBreakersValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e CircuitBreakersValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakers.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = CircuitBreakersValidationError{}

// Validate checks the field values on CircuitBreakers_Thresholds with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CircuitBreakers_Thresholds) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Priority

	if v, ok := interface{}(m.GetMaxConnections()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CircuitBreakers_ThresholdsValidationError{
				Field:  "MaxConnections",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMaxPendingRequests()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CircuitBreakers_ThresholdsValidationError{
				Field:  "MaxPendingRequests",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMaxRequests()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CircuitBreakers_ThresholdsValidationError{
				Field:  "MaxRequests",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMaxRetries()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CircuitBreakers_ThresholdsValidationError{
				Field:  "MaxRetries",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// CircuitBreakers_ThresholdsValidationError is the validation error returned
// by CircuitBreakers_Thresholds.Validate if the designated constraints aren't met.
type CircuitBreakers_ThresholdsValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e CircuitBreakers_ThresholdsValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCircuitBreakers_Thresholds.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = CircuitBreakers_ThresholdsValidationError{}