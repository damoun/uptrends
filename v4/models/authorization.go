// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Authorization authorization
// swagger:model Authorization
type Authorization struct {

	// authorization Id
	// Required: true
	AuthorizationID *string `json:"AuthorizationId"`

	// authorization type
	// Required: true
	AuthorizationType AuthorizationTypeEnum `json:"AuthorizationType"`

	// context Id
	// Required: true
	ContextID *string `json:"ContextId"`

	// operator group Guid
	OperatorGroupGUID string `json:"OperatorGroupGuid,omitempty"`

	// operator Guid
	OperatorGUID string `json:"OperatorGuid,omitempty"`
}

// Validate validates this authorization
func (m *Authorization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizationType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContextID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Authorization) validateAuthorizationID(formats strfmt.Registry) error {

	if err := validate.Required("AuthorizationId", "body", m.AuthorizationID); err != nil {
		return err
	}

	return nil
}

func (m *Authorization) validateAuthorizationType(formats strfmt.Registry) error {

	if err := m.AuthorizationType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("AuthorizationType")
		}
		return err
	}

	return nil
}

func (m *Authorization) validateContextID(formats strfmt.Registry) error {

	if err := validate.Required("ContextId", "body", m.ContextID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Authorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Authorization) UnmarshalBinary(b []byte) error {
	var res Authorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
