// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OperatorMember operator member
//
// swagger:model OperatorMember
type OperatorMember struct {

	// The unique identifier of this Operator
	// Required: true
	OperatorGroupGUID *string `json:"OperatorGroupGuid"`
}

// Validate validates this operator member
func (m *OperatorMember) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperatorGroupGUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperatorMember) validateOperatorGroupGUID(formats strfmt.Registry) error {

	if err := validate.Required("OperatorGroupGuid", "body", m.OperatorGroupGUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperatorMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperatorMember) UnmarshalBinary(b []byte) error {
	var res OperatorMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
