// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserDefinedFunctionMapping user defined function mapping
//
// swagger:model UserDefinedFunctionMapping
type UserDefinedFunctionMapping struct {

	// key
	Key string `json:"Key,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this user defined function mapping
func (m *UserDefinedFunctionMapping) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserDefinedFunctionMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDefinedFunctionMapping) UnmarshalBinary(b []byte) error {
	var res UserDefinedFunctionMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
