// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RequestHeader request header
//
// swagger:model RequestHeader
type RequestHeader struct {

	// name
	Name string `json:"Name,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this request header
func (m *RequestHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RequestHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RequestHeader) UnmarshalBinary(b []byte) error {
	var res RequestHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
