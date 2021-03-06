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

// IntegrationServiceMap integration service map
//
// swagger:model IntegrationServiceMap
type IntegrationServiceMap struct {

	// integration service Guid
	// Required: true
	IntegrationServiceGUID *string `json:"IntegrationServiceGuid"`

	// monitor Guid
	// Required: true
	MonitorGUID *string `json:"MonitorGuid"`
}

// Validate validates this integration service map
func (m *IntegrationServiceMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegrationServiceGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitorGUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationServiceMap) validateIntegrationServiceGUID(formats strfmt.Registry) error {

	if err := validate.Required("IntegrationServiceGuid", "body", m.IntegrationServiceGUID); err != nil {
		return err
	}

	return nil
}

func (m *IntegrationServiceMap) validateMonitorGUID(formats strfmt.Registry) error {

	if err := validate.Required("MonitorGuid", "body", m.MonitorGUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationServiceMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationServiceMap) UnmarshalBinary(b []byte) error {
	var res IntegrationServiceMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
