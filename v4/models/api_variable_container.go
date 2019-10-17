// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// APIVariableContainer Api variable container
// swagger:model ApiVariableContainer
type APIVariableContainer struct {
	SerializableDictionaryOfStringAndAPIVariable
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *APIVariableContainer) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 SerializableDictionaryOfStringAndAPIVariable
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.SerializableDictionaryOfStringAndAPIVariable = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m APIVariableContainer) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.SerializableDictionaryOfStringAndAPIVariable)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this Api variable container
func (m *APIVariableContainer) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with SerializableDictionaryOfStringAndAPIVariable
	if err := m.SerializableDictionaryOfStringAndAPIVariable.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *APIVariableContainer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIVariableContainer) UnmarshalBinary(b []byte) error {
	var res APIVariableContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}