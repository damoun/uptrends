// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// HTTPAttributes Http details attributes
// swagger:model HttpAttributes
type HTTPAttributes struct {

	// The HTML code retrieved from the target
	ResponseBody string `json:"ResponseBody,omitempty"`

	// The HTTP response headers retrieved from the target
	ResponseHeaders string `json:"ResponseHeaders,omitempty"`
}

// Validate validates this Http attributes
func (m *HTTPAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HTTPAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HTTPAttributes) UnmarshalBinary(b []byte) error {
	var res HTTPAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}