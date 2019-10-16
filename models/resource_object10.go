// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ResourceObject10 resource object10
// swagger:model ResourceObject10
type ResourceObject10 struct {
	ResourceObject9
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResourceObject10) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceObject9
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceObject9 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResourceObject10) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ResourceObject9)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resource object10
func (m *ResourceObject10) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ResourceObject9
	if err := m.ResourceObject9.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceObject10) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceObject10) UnmarshalBinary(b []byte) error {
	var res ResourceObject10
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}