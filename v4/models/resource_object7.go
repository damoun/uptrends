// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ResourceObject7 resource object7
// swagger:model ResourceObject7
type ResourceObject7 struct {
	ResourceObject8

	// attributes
	Attributes *MonitorCheckAttributes `json:"Attributes,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResourceObject7) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceObject8
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceObject8 = aO0

	// now for regular properties
	var propsResourceObject7 struct {
		Attributes *MonitorCheckAttributes `json:"Attributes,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsResourceObject7); err != nil {
		return err
	}
	m.Attributes = propsResourceObject7.Attributes

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResourceObject7) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ResourceObject8)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsResourceObject7 struct {
		Attributes *MonitorCheckAttributes `json:"Attributes,omitempty"`
	}
	propsResourceObject7.Attributes = m.Attributes

	jsonDataPropsResourceObject7, errResourceObject7 := swag.WriteJSON(propsResourceObject7)
	if errResourceObject7 != nil {
		return nil, errResourceObject7
	}
	_parts = append(_parts, jsonDataPropsResourceObject7)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resource object7
func (m *ResourceObject7) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ResourceObject8
	if err := m.ResourceObject8.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceObject7) validateAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	if m.Attributes != nil {
		if err := m.Attributes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Attributes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceObject7) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceObject7) UnmarshalBinary(b []byte) error {
	var res ResourceObject7
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
