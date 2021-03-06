// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OperatorGroup operator group
//
// swagger:model OperatorGroup
type OperatorGroup struct {

	// The descriptive name of this operator group
	Description string `json:"Description,omitempty"`

	// is administrators group
	IsAdministratorsGroup bool `json:"IsAdministratorsGroup,omitempty"`

	// Indicates whether this is the default group for all operators
	IsEveryone bool `json:"IsEveryone,omitempty"`

	// The unique identifier of this Operator group
	OperatorGroupGUID string `json:"OperatorGroupGuid,omitempty"`
}

// Validate validates this operator group
func (m *OperatorGroup) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperatorGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperatorGroup) UnmarshalBinary(b []byte) error {
	var res OperatorGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
