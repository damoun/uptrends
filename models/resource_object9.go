// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ResourceObject9 resource object9
// swagger:model ResourceObject9
type ResourceObject9 struct {
	ResourceObject3

	// attributes
	Attributes *CheckpointServerAttributes `json:"Attributes,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResourceObject9) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceObject3
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceObject3 = aO0

	// now for regular properties
	var propsResourceObject9 struct {
		Attributes *CheckpointServerAttributes `json:"Attributes,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsResourceObject9); err != nil {
		return err
	}
	m.Attributes = propsResourceObject9.Attributes

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResourceObject9) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ResourceObject3)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsResourceObject9 struct {
		Attributes *CheckpointServerAttributes `json:"Attributes,omitempty"`
	}
	propsResourceObject9.Attributes = m.Attributes

	jsonDataPropsResourceObject9, errResourceObject9 := swag.WriteJSON(propsResourceObject9)
	if errResourceObject9 != nil {
		return nil, errResourceObject9
	}
	_parts = append(_parts, jsonDataPropsResourceObject9)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resource object9
func (m *ResourceObject9) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ResourceObject3
	if err := m.ResourceObject3.Validate(formats); err != nil {
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

func (m *ResourceObject9) validateAttributes(formats strfmt.Registry) error {

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
func (m *ResourceObject9) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceObject9) UnmarshalBinary(b []byte) error {
	var res ResourceObject9
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
