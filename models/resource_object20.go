// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ResourceObject20 resource object20
// swagger:model ResourceObject20
type ResourceObject20 struct {
	ResourceObject8

	// attributes
	Attributes *WaterfallInfo `json:"Attributes,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResourceObject20) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceObject8
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceObject8 = aO0

	// now for regular properties
	var propsResourceObject20 struct {
		Attributes *WaterfallInfo `json:"Attributes,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsResourceObject20); err != nil {
		return err
	}
	m.Attributes = propsResourceObject20.Attributes

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResourceObject20) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ResourceObject8)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsResourceObject20 struct {
		Attributes *WaterfallInfo `json:"Attributes,omitempty"`
	}
	propsResourceObject20.Attributes = m.Attributes

	jsonDataPropsResourceObject20, errResourceObject20 := swag.WriteJSON(propsResourceObject20)
	if errResourceObject20 != nil {
		return nil, errResourceObject20
	}
	_parts = append(_parts, jsonDataPropsResourceObject20)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resource object20
func (m *ResourceObject20) Validate(formats strfmt.Registry) error {
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

func (m *ResourceObject20) validateAttributes(formats strfmt.Registry) error {

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
func (m *ResourceObject20) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceObject20) UnmarshalBinary(b []byte) error {
	var res ResourceObject20
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}