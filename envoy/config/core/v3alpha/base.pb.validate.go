// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3alpha/base.proto

package envoy_config_core_v3alpha

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
)

// define the regex for a UUID once up-front
var _base_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Locality with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Locality) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Region

	// no validation rules for Zone

	// no validation rules for SubZone

	return nil
}

// LocalityValidationError is the validation error returned by
// Locality.Validate if the designated constraints aren't met.
type LocalityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityValidationError) ErrorName() string { return "LocalityValidationError" }

// Error satisfies the builtin error interface
func (e LocalityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocality.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityValidationError{}

// Validate checks the field values on BuildVersion with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *BuildVersion) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetVersion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BuildVersionValidationError{
				field:  "Version",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BuildVersionValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// BuildVersionValidationError is the validation error returned by
// BuildVersion.Validate if the designated constraints aren't met.
type BuildVersionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BuildVersionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BuildVersionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BuildVersionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BuildVersionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BuildVersionValidationError) ErrorName() string { return "BuildVersionValidationError" }

// Error satisfies the builtin error interface
func (e BuildVersionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBuildVersion.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BuildVersionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BuildVersionValidationError{}

// Validate checks the field values on Extension with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Extension) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Category

	// no validation rules for TypeDescriptor

	if v, ok := interface{}(m.GetVersion()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExtensionValidationError{
				field:  "Version",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Disabled

	return nil
}

// ExtensionValidationError is the validation error returned by
// Extension.Validate if the designated constraints aren't met.
type ExtensionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtensionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtensionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtensionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtensionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtensionValidationError) ErrorName() string { return "ExtensionValidationError" }

// Error satisfies the builtin error interface
func (e ExtensionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtension.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtensionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtensionValidationError{}

// Validate checks the field values on Node with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Node) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Cluster

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLocality()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return NodeValidationError{
				field:  "Locality",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HiddenEnvoyDeprecatedBuildVersion

	// no validation rules for UserAgentName

	for idx, item := range m.GetExtensions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  fmt.Sprintf("Extensions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	switch m.UserAgentVersionType.(type) {

	case *Node_UserAgentVersion:
		// no validation rules for UserAgentVersion

	case *Node_UserAgentBuildVersion:

		if v, ok := interface{}(m.GetUserAgentBuildVersion()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return NodeValidationError{
					field:  "UserAgentBuildVersion",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// NodeValidationError is the validation error returned by Node.Validate if the
// designated constraints aren't met.
type NodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e NodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e NodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e NodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e NodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e NodeValidationError) ErrorName() string { return "NodeValidationError" }

// Error satisfies the builtin error interface
func (e NodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sNode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = NodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = NodeValidationError{}

// Validate checks the field values on Metadata with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Metadata) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetFilterMetadata() {
		_ = val

		// no validation rules for FilterMetadata[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MetadataValidationError{
					field:  fmt.Sprintf("FilterMetadata[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MetadataValidationError is the validation error returned by
// Metadata.Validate if the designated constraints aren't met.
type MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataValidationError) ErrorName() string { return "MetadataValidationError" }

// Error satisfies the builtin error interface
func (e MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataValidationError{}

// Validate checks the field values on RuntimeUInt32 with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RuntimeUInt32) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DefaultValue

	if len(m.GetRuntimeKey()) < 1 {
		return RuntimeUInt32ValidationError{
			field:  "RuntimeKey",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// RuntimeUInt32ValidationError is the validation error returned by
// RuntimeUInt32.Validate if the designated constraints aren't met.
type RuntimeUInt32ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeUInt32ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeUInt32ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeUInt32ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeUInt32ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeUInt32ValidationError) ErrorName() string { return "RuntimeUInt32ValidationError" }

// Error satisfies the builtin error interface
func (e RuntimeUInt32ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntimeUInt32.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeUInt32ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeUInt32ValidationError{}

// Validate checks the field values on RuntimeFeatureFlag with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RuntimeFeatureFlag) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDefaultValue() == nil {
		return RuntimeFeatureFlagValidationError{
			field:  "DefaultValue",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDefaultValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RuntimeFeatureFlagValidationError{
				field:  "DefaultValue",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetRuntimeKey()) < 1 {
		return RuntimeFeatureFlagValidationError{
			field:  "RuntimeKey",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// RuntimeFeatureFlagValidationError is the validation error returned by
// RuntimeFeatureFlag.Validate if the designated constraints aren't met.
type RuntimeFeatureFlagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeFeatureFlagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeFeatureFlagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeFeatureFlagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeFeatureFlagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeFeatureFlagValidationError) ErrorName() string {
	return "RuntimeFeatureFlagValidationError"
}

// Error satisfies the builtin error interface
func (e RuntimeFeatureFlagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntimeFeatureFlag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeFeatureFlagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeFeatureFlagValidationError{}

// Validate checks the field values on HeaderValue with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HeaderValue) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetKey()); l < 1 || l > 16384 {
		return HeaderValueValidationError{
			field:  "Key",
			reason: "value length must be between 1 and 16384 bytes, inclusive",
		}
	}

	if len(m.GetValue()) > 16384 {
		return HeaderValueValidationError{
			field:  "Value",
			reason: "value length must be at most 16384 bytes",
		}
	}

	return nil
}

// HeaderValueValidationError is the validation error returned by
// HeaderValue.Validate if the designated constraints aren't met.
type HeaderValueValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderValueValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderValueValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderValueValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderValueValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderValueValidationError) ErrorName() string { return "HeaderValueValidationError" }

// Error satisfies the builtin error interface
func (e HeaderValueValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderValue.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderValueValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderValueValidationError{}

// Validate checks the field values on HeaderValueOption with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HeaderValueOption) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetHeader() == nil {
		return HeaderValueOptionValidationError{
			field:  "Header",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetHeader()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeaderValueOptionValidationError{
				field:  "Header",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAppend()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HeaderValueOptionValidationError{
				field:  "Append",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// HeaderValueOptionValidationError is the validation error returned by
// HeaderValueOption.Validate if the designated constraints aren't met.
type HeaderValueOptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderValueOptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderValueOptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderValueOptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderValueOptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderValueOptionValidationError) ErrorName() string {
	return "HeaderValueOptionValidationError"
}

// Error satisfies the builtin error interface
func (e HeaderValueOptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderValueOption.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderValueOptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderValueOptionValidationError{}

// Validate checks the field values on HeaderMap with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *HeaderMap) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HeaderMapValidationError{
					field:  fmt.Sprintf("Headers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HeaderMapValidationError is the validation error returned by
// HeaderMap.Validate if the designated constraints aren't met.
type HeaderMapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HeaderMapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HeaderMapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HeaderMapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HeaderMapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HeaderMapValidationError) ErrorName() string { return "HeaderMapValidationError" }

// Error satisfies the builtin error interface
func (e HeaderMapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHeaderMap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HeaderMapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HeaderMapValidationError{}

// Validate checks the field values on DataSource with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *DataSource) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Specifier.(type) {

	case *DataSource_Filename:

		if len(m.GetFilename()) < 1 {
			return DataSourceValidationError{
				field:  "Filename",
				reason: "value length must be at least 1 bytes",
			}
		}

	case *DataSource_InlineBytes:

		if len(m.GetInlineBytes()) < 1 {
			return DataSourceValidationError{
				field:  "InlineBytes",
				reason: "value length must be at least 1 bytes",
			}
		}

	case *DataSource_InlineString:

		if len(m.GetInlineString()) < 1 {
			return DataSourceValidationError{
				field:  "InlineString",
				reason: "value length must be at least 1 bytes",
			}
		}

	default:
		return DataSourceValidationError{
			field:  "Specifier",
			reason: "value is required",
		}

	}

	return nil
}

// DataSourceValidationError is the validation error returned by
// DataSource.Validate if the designated constraints aren't met.
type DataSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataSourceValidationError) ErrorName() string { return "DataSourceValidationError" }

// Error satisfies the builtin error interface
func (e DataSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataSourceValidationError{}

// Validate checks the field values on RemoteDataSource with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *RemoteDataSource) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetHttpUri() == nil {
		return RemoteDataSourceValidationError{
			field:  "HttpUri",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetHttpUri()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RemoteDataSourceValidationError{
				field:  "HttpUri",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetSha256()) < 1 {
		return RemoteDataSourceValidationError{
			field:  "Sha256",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// RemoteDataSourceValidationError is the validation error returned by
// RemoteDataSource.Validate if the designated constraints aren't met.
type RemoteDataSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoteDataSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoteDataSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoteDataSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoteDataSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoteDataSourceValidationError) ErrorName() string { return "RemoteDataSourceValidationError" }

// Error satisfies the builtin error interface
func (e RemoteDataSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoteDataSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoteDataSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoteDataSourceValidationError{}

// Validate checks the field values on AsyncDataSource with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AsyncDataSource) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Specifier.(type) {

	case *AsyncDataSource_Local:

		if v, ok := interface{}(m.GetLocal()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AsyncDataSourceValidationError{
					field:  "Local",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *AsyncDataSource_Remote:

		if v, ok := interface{}(m.GetRemote()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AsyncDataSourceValidationError{
					field:  "Remote",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return AsyncDataSourceValidationError{
			field:  "Specifier",
			reason: "value is required",
		}

	}

	return nil
}

// AsyncDataSourceValidationError is the validation error returned by
// AsyncDataSource.Validate if the designated constraints aren't met.
type AsyncDataSourceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AsyncDataSourceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AsyncDataSourceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AsyncDataSourceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AsyncDataSourceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AsyncDataSourceValidationError) ErrorName() string { return "AsyncDataSourceValidationError" }

// Error satisfies the builtin error interface
func (e AsyncDataSourceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAsyncDataSource.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AsyncDataSourceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AsyncDataSourceValidationError{}

// Validate checks the field values on TransportSocket with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *TransportSocket) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return TransportSocketValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.ConfigType.(type) {

	case *TransportSocket_HiddenEnvoyDeprecatedConfig:

		if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TransportSocketValidationError{
					field:  "HiddenEnvoyDeprecatedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TransportSocket_TypedConfig:

		if v, ok := interface{}(m.GetTypedConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TransportSocketValidationError{
					field:  "TypedConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// TransportSocketValidationError is the validation error returned by
// TransportSocket.Validate if the designated constraints aren't met.
type TransportSocketValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TransportSocketValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TransportSocketValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TransportSocketValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TransportSocketValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TransportSocketValidationError) ErrorName() string { return "TransportSocketValidationError" }

// Error satisfies the builtin error interface
func (e TransportSocketValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTransportSocket.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TransportSocketValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TransportSocketValidationError{}

// Validate checks the field values on SocketOption with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SocketOption) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Description

	// no validation rules for Level

	// no validation rules for Name

	if _, ok := SocketOption_SocketState_name[int32(m.GetState())]; !ok {
		return SocketOptionValidationError{
			field:  "State",
			reason: "value must be one of the defined enum values",
		}
	}

	switch m.Value.(type) {

	case *SocketOption_IntValue:
		// no validation rules for IntValue

	case *SocketOption_BufValue:
		// no validation rules for BufValue

	default:
		return SocketOptionValidationError{
			field:  "Value",
			reason: "value is required",
		}

	}

	return nil
}

// SocketOptionValidationError is the validation error returned by
// SocketOption.Validate if the designated constraints aren't met.
type SocketOptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SocketOptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SocketOptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SocketOptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SocketOptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SocketOptionValidationError) ErrorName() string { return "SocketOptionValidationError" }

// Error satisfies the builtin error interface
func (e SocketOptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSocketOption.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SocketOptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SocketOptionValidationError{}

// Validate checks the field values on RuntimeFractionalPercent with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RuntimeFractionalPercent) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetDefaultValue() == nil {
		return RuntimeFractionalPercentValidationError{
			field:  "DefaultValue",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetDefaultValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RuntimeFractionalPercentValidationError{
				field:  "DefaultValue",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RuntimeKey

	return nil
}

// RuntimeFractionalPercentValidationError is the validation error returned by
// RuntimeFractionalPercent.Validate if the designated constraints aren't met.
type RuntimeFractionalPercentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeFractionalPercentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeFractionalPercentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeFractionalPercentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeFractionalPercentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeFractionalPercentValidationError) ErrorName() string {
	return "RuntimeFractionalPercentValidationError"
}

// Error satisfies the builtin error interface
func (e RuntimeFractionalPercentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntimeFractionalPercent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeFractionalPercentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeFractionalPercentValidationError{}

// Validate checks the field values on ControlPlane with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ControlPlane) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Identifier

	return nil
}

// ControlPlaneValidationError is the validation error returned by
// ControlPlane.Validate if the designated constraints aren't met.
type ControlPlaneValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ControlPlaneValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ControlPlaneValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ControlPlaneValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ControlPlaneValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ControlPlaneValidationError) ErrorName() string { return "ControlPlaneValidationError" }

// Error satisfies the builtin error interface
func (e ControlPlaneValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sControlPlane.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ControlPlaneValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ControlPlaneValidationError{}
