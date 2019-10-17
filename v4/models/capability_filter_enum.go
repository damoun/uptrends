// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// CapabilityFilterEnum capability filter enum
// swagger:model CapabilityFilterEnum
type CapabilityFilterEnum string

const (

	// CapabilityFilterEnumIPV6 captures enum value "IPv6"
	CapabilityFilterEnumIPV6 CapabilityFilterEnum = "IPv6"

	// CapabilityFilterEnumNativeIPV6 captures enum value "NativeIPv6"
	CapabilityFilterEnumNativeIPV6 CapabilityFilterEnum = "NativeIPv6"

	// CapabilityFilterEnumPrimaryServer captures enum value "PrimaryServer"
	CapabilityFilterEnumPrimaryServer CapabilityFilterEnum = "PrimaryServer"
)

// for schema
var capabilityFilterEnumEnum []interface{}

func init() {
	var res []CapabilityFilterEnum
	if err := json.Unmarshal([]byte(`["IPv6","NativeIPv6","PrimaryServer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		capabilityFilterEnumEnum = append(capabilityFilterEnumEnum, v)
	}
}

func (m CapabilityFilterEnum) validateCapabilityFilterEnumEnum(path, location string, value CapabilityFilterEnum) error {
	if err := validate.Enum(path, location, value, capabilityFilterEnumEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this capability filter enum
func (m CapabilityFilterEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCapabilityFilterEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}