// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ResourceObject6 Represents the values/results of a single check done by a monitor
// swagger:model ResourceObject6
type ResourceObject6 struct {
	ResourceObject7

	// type
	Type string `json:"Type,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResourceObject6) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceObject7
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceObject7 = aO0

	// now for regular properties
	var propsResourceObject6 struct {
		Type string `json:"Type,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsResourceObject6); err != nil {
		return err
	}
	m.Type = propsResourceObject6.Type

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResourceObject6) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ResourceObject7)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsResourceObject6 struct {
		Type string `json:"Type,omitempty"`
	}
	propsResourceObject6.Type = m.Type

	jsonDataPropsResourceObject6, errResourceObject6 := swag.WriteJSON(propsResourceObject6)
	if errResourceObject6 != nil {
		return nil, errResourceObject6
	}
	_parts = append(_parts, jsonDataPropsResourceObject6)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resource object6
func (m *ResourceObject6) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ResourceObject7
	if err := m.ResourceObject7.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceObject6) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceObject6) UnmarshalBinary(b []byte) error {
	var res ResourceObject6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
