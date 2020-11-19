// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LinksData Links with reference to self and next data set
//
// swagger:model LinksData
type LinksData struct {

	// Url for next data set
	Next string `json:"Next,omitempty"`

	// Url for this data set (data might get updated, depending on your parameters)
	Self string `json:"Self,omitempty"`
}

// Validate validates this links data
func (m *LinksData) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LinksData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LinksData) UnmarshalBinary(b []byte) error {
	var res LinksData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
