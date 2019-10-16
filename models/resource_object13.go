// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ResourceObject13 resource object13
// swagger:model ResourceObject13
type ResourceObject13 struct {
	ResourceObject14
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResourceObject13) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceObject14
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceObject14 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResourceObject13) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ResourceObject14)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resource object13
func (m *ResourceObject13) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ResourceObject14
	if err := m.ResourceObject14.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceObject13) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceObject13) UnmarshalBinary(b []byte) error {
	var res ResourceObject13
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
