// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceObject8 resource object8
// swagger:model ResourceObject8
type ResourceObject8 struct {

	// Id
	// Required: true
	ID *int64 `json:"Id"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// relationships
	Relationships []*RelationObject `json:"Relationships"`

	// type
	Type string `json:"Type,omitempty"`
}

// Validate validates this resource object8
func (m *ResourceObject8) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceObject8) validateID(formats strfmt.Registry) error {

	if err := validate.Required("Id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *ResourceObject8) validateRelationships(formats strfmt.Registry) error {

	if swag.IsZero(m.Relationships) { // not required
		return nil
	}

	for i := 0; i < len(m.Relationships); i++ {
		if swag.IsZero(m.Relationships[i]) { // not required
			continue
		}

		if m.Relationships[i] != nil {
			if err := m.Relationships[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Relationships" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceObject8) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceObject8) UnmarshalBinary(b []byte) error {
	var res ResourceObject8
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
