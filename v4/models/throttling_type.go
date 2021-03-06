// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ThrottlingType throttling type
//
// swagger:model ThrottlingType
type ThrottlingType string

const (

	// ThrottlingTypeInactive captures enum value "Inactive"
	ThrottlingTypeInactive ThrottlingType = "Inactive"

	// ThrottlingTypeBrowser captures enum value "Browser"
	ThrottlingTypeBrowser ThrottlingType = "Browser"

	// ThrottlingTypeSimulated captures enum value "Simulated"
	ThrottlingTypeSimulated ThrottlingType = "Simulated"
)

// for schema
var throttlingTypeEnum []interface{}

func init() {
	var res []ThrottlingType
	if err := json.Unmarshal([]byte(`["Inactive","Browser","Simulated"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		throttlingTypeEnum = append(throttlingTypeEnum, v)
	}
}

func (m ThrottlingType) validateThrottlingTypeEnum(path, location string, value ThrottlingType) error {
	if err := validate.EnumCase(path, location, value, throttlingTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this throttling type
func (m ThrottlingType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateThrottlingTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
