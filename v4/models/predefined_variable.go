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

// PredefinedVariable predefined variable
//
// swagger:model PredefinedVariable
type PredefinedVariable struct {

	// key
	Key string `json:"Key,omitempty"`

	// sensitive
	// Required: true
	Sensitive *bool `json:"Sensitive"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this predefined variable
func (m *PredefinedVariable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSensitive(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PredefinedVariable) validateSensitive(formats strfmt.Registry) error {

	if err := validate.Required("Sensitive", "body", m.Sensitive); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PredefinedVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PredefinedVariable) UnmarshalBinary(b []byte) error {
	var res PredefinedVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}