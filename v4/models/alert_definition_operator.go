// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertDefinitionOperator alert definition operator
//
// swagger:model AlertDefinitionOperator
type AlertDefinitionOperator struct {

	// alert definition
	// Required: true
	AlertDefinition *string `json:"AlertDefinition"`

	// escalationlevel
	// Required: true
	Escalationlevel *int64 `json:"Escalationlevel"`

	// operator
	// Required: true
	Operator *string `json:"Operator"`
}

// Validate validates this alert definition operator
func (m *AlertDefinitionOperator) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertDefinition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEscalationlevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertDefinitionOperator) validateAlertDefinition(formats strfmt.Registry) error {

	if err := validate.Required("AlertDefinition", "body", m.AlertDefinition); err != nil {
		return err
	}

	return nil
}

func (m *AlertDefinitionOperator) validateEscalationlevel(formats strfmt.Registry) error {

	if err := validate.Required("Escalationlevel", "body", m.Escalationlevel); err != nil {
		return err
	}

	return nil
}

func (m *AlertDefinitionOperator) validateOperator(formats strfmt.Registry) error {

	if err := validate.Required("Operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertDefinitionOperator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertDefinitionOperator) UnmarshalBinary(b []byte) error {
	var res AlertDefinitionOperator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
