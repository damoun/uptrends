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

// EscalationMode escalation mode
//
// swagger:model EscalationMode
type EscalationMode string

const (

	// EscalationModeAlertOnErrorCount captures enum value "AlertOnErrorCount"
	EscalationModeAlertOnErrorCount EscalationMode = "AlertOnErrorCount"

	// EscalationModeAlertOnErrorDuration captures enum value "AlertOnErrorDuration"
	EscalationModeAlertOnErrorDuration EscalationMode = "AlertOnErrorDuration"
)

// for schema
var escalationModeEnum []interface{}

func init() {
	var res []EscalationMode
	if err := json.Unmarshal([]byte(`["AlertOnErrorCount","AlertOnErrorDuration"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		escalationModeEnum = append(escalationModeEnum, v)
	}
}

func (m EscalationMode) validateEscalationModeEnum(path, location string, value EscalationMode) error {
	if err := validate.EnumCase(path, location, value, escalationModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this escalation mode
func (m EscalationMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEscalationModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
