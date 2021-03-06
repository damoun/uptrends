// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ResourceObject18 resource object18
// swagger:model ResourceObject18
type ResourceObject18 struct {
	ResourceObject19

	// attributes
	Attributes *TransactionStep `json:"Attributes,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResourceObject18) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceObject19
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceObject19 = aO0

	// now for regular properties
	var propsResourceObject18 struct {
		Attributes *TransactionStep `json:"Attributes,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsResourceObject18); err != nil {
		return err
	}
	m.Attributes = propsResourceObject18.Attributes

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResourceObject18) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ResourceObject19)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsResourceObject18 struct {
		Attributes *TransactionStep `json:"Attributes,omitempty"`
	}
	propsResourceObject18.Attributes = m.Attributes

	jsonDataPropsResourceObject18, errResourceObject18 := swag.WriteJSON(propsResourceObject18)
	if errResourceObject18 != nil {
		return nil, errResourceObject18
	}
	_parts = append(_parts, jsonDataPropsResourceObject18)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resource object18
func (m *ResourceObject18) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ResourceObject19
	if err := m.ResourceObject19.Validate(formats); err != nil {
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

func (m *ResourceObject18) validateAttributes(formats strfmt.Registry) error {

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
func (m *ResourceObject18) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceObject18) UnmarshalBinary(b []byte) error {
	var res ResourceObject18
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
