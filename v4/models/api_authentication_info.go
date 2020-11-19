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

// APIAuthenticationInfo Api authentication info
//
// swagger:model ApiAuthenticationInfo
type APIAuthenticationInfo struct {

	// authentication type
	// Required: true
	AuthenticationType APIHTTPAuthenticationType `json:"AuthenticationType"`

	// Id
	// Required: true
	ID *string `json:"Id"`

	// password
	Password string `json:"Password,omitempty"`

	// password specified
	// Required: true
	PasswordSpecified *bool `json:"PasswordSpecified"`

	// user name
	UserName string `json:"UserName,omitempty"`
}

// Validate validates this Api authentication info
func (m *APIAuthenticationInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordSpecified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIAuthenticationInfo) validateAuthenticationType(formats strfmt.Registry) error {

	if err := m.AuthenticationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AuthenticationType")
		}
		return err
	}

	return nil
}

func (m *APIAuthenticationInfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("Id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *APIAuthenticationInfo) validatePasswordSpecified(formats strfmt.Registry) error {

	if err := validate.Required("PasswordSpecified", "body", m.PasswordSpecified); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIAuthenticationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIAuthenticationInfo) UnmarshalBinary(b []byte) error {
	var res APIAuthenticationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
