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

// LastErrorLevel last error level
//
// swagger:model LastErrorLevel
type LastErrorLevel string

const (

	// LastErrorLevelNoError captures enum value "NoError"
	LastErrorLevelNoError LastErrorLevel = "NoError"

	// LastErrorLevelUnconfirmed captures enum value "Unconfirmed"
	LastErrorLevelUnconfirmed LastErrorLevel = "Unconfirmed"

	// LastErrorLevelConfirmed captures enum value "Confirmed"
	LastErrorLevelConfirmed LastErrorLevel = "Confirmed"

	// LastErrorLevelInactive captures enum value "Inactive"
	LastErrorLevelInactive LastErrorLevel = "Inactive"

	// LastErrorLevelInconclusive captures enum value "Inconclusive"
	LastErrorLevelInconclusive LastErrorLevel = "Inconclusive"
)

// for schema
var lastErrorLevelEnum []interface{}

func init() {
	var res []LastErrorLevel
	if err := json.Unmarshal([]byte(`["NoError","Unconfirmed","Confirmed","Inactive","Inconclusive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lastErrorLevelEnum = append(lastErrorLevelEnum, v)
	}
}

func (m LastErrorLevel) validateLastErrorLevelEnum(path, location string, value LastErrorLevel) error {
	if err := validate.EnumCase(path, location, value, lastErrorLevelEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this last error level
func (m LastErrorLevel) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLastErrorLevelEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
