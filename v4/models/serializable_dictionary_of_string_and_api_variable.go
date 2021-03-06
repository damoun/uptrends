// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SerializableDictionaryOfStringAndAPIVariable serializable dictionary of string and Api variable
// swagger:model SerializableDictionaryOfStringAndApiVariable
type SerializableDictionaryOfStringAndAPIVariable struct {

	// a o0
	// Required: true
	AO0 map[string]*APIVariable `json:"-"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SerializableDictionaryOfStringAndAPIVariable) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 map[string]*APIVariable
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.AO0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SerializableDictionaryOfStringAndAPIVariable) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.AO0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this serializable dictionary of string and Api variable
func (m *SerializableDictionaryOfStringAndAPIVariable) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with map[string]APIVariable

	for k := range m.AO0 {

		if err := validate.Required(k, "body", m.AO0[k]); err != nil {
			return err
		}
		if val, ok := m.AO0[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SerializableDictionaryOfStringAndAPIVariable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SerializableDictionaryOfStringAndAPIVariable) UnmarshalBinary(b []byte) error {
	var res SerializableDictionaryOfStringAndAPIVariable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
