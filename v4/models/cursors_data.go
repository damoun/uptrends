// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CursorsData Cursors can be used to navigate the dataset in a fixed manner
//
// swagger:model CursorsData
type CursorsData struct {

	// Cursor for next data set
	Next string `json:"Next,omitempty"`

	// Cursor for this data set (data might get updated, depending on your parameters)
	Self string `json:"Self,omitempty"`
}

// Validate validates this cursors data
func (m *CursorsData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CursorsData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CursorsData) UnmarshalBinary(b []byte) error {
	var res CursorsData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
