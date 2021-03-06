// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OperatorQuota operator quota
//
// swagger:model OperatorQuota
type OperatorQuota struct {

	// operators
	Operators int32 `json:"Operators,omitempty"`

	// operators in use
	OperatorsInUse int32 `json:"OperatorsInUse,omitempty"`
}

// Validate validates this operator quota
func (m *OperatorQuota) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OperatorQuota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperatorQuota) UnmarshalBinary(b []byte) error {
	var res OperatorQuota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
