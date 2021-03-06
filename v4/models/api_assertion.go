// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIAssertion Api assertion
//
// swagger:model ApiAssertion
type APIAssertion struct {

	// comparison
	// Required: true
	Comparison APIComparisonType `json:"Comparison"`

	// property
	Property string `json:"Property,omitempty"`

	// source
	// Required: true
	Source APIAssertionSourceType `json:"Source"`

	// target value
	TargetValue string `json:"TargetValue,omitempty"`
}

// Validate validates this Api assertion
func (m *APIAssertion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComparison(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIAssertion) validateComparison(formats strfmt.Registry) error {

	if err := m.Comparison.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Comparison")
		}
		return err
	}

	return nil
}

func (m *APIAssertion) validateSource(formats strfmt.Registry) error {

	if err := m.Source.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Source")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIAssertion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIAssertion) UnmarshalBinary(b []byte) error {
	var res APIAssertion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
