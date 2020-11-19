// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MsaStep msa step
//
// swagger:model MsaStep
type MsaStep struct {

	// assertions
	Assertions []*APIAssertion `json:"Assertions"`

	// authentication
	Authentication *APIAuthenticationInfo `json:"Authentication,omitempty"`

	// body
	Body string `json:"Body,omitempty"`

	// client certificate vault item
	ClientCertificateVaultItem string `json:"ClientCertificateVaultItem,omitempty"`

	// delay
	// Required: true
	Delay *int32 `json:"Delay"`

	// disabled
	// Required: true
	Disabled *bool `json:"Disabled"`

	// ignore certificate errors
	// Required: true
	IgnoreCertificateErrors *bool `json:"IgnoreCertificateErrors"`

	// method
	Method string `json:"Method,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// request headers
	RequestHeaders []*APIHTTPHeaderValue `json:"RequestHeaders"`

	// step type
	// Required: true
	StepType MsaStepType `json:"StepType"`

	// Url
	URL string `json:"Url,omitempty"`

	// use fixed client certificate
	// Required: true
	UseFixedClientCertificate *bool `json:"UseFixedClientCertificate"`

	// variables
	Variables []*APIVariableDefinition `json:"Variables"`
}

// Validate validates this msa step
func (m *MsaStep) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssertions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgnoreCertificateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestHeaders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStepType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUseFixedClientCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariables(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MsaStep) validateAssertions(formats strfmt.Registry) error {

	if swag.IsZero(m.Assertions) { // not required
		return nil
	}

	for i := 0; i < len(m.Assertions); i++ {
		if swag.IsZero(m.Assertions[i]) { // not required
			continue
		}

		if m.Assertions[i] != nil {
			if err := m.Assertions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Assertions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsaStep) validateAuthentication(formats strfmt.Registry) error {

	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Authentication")
			}
			return err
		}
	}

	return nil
}

func (m *MsaStep) validateDelay(formats strfmt.Registry) error {

	if err := validate.Required("Delay", "body", m.Delay); err != nil {
		return err
	}

	return nil
}

func (m *MsaStep) validateDisabled(formats strfmt.Registry) error {

	if err := validate.Required("Disabled", "body", m.Disabled); err != nil {
		return err
	}

	return nil
}

func (m *MsaStep) validateIgnoreCertificateErrors(formats strfmt.Registry) error {

	if err := validate.Required("IgnoreCertificateErrors", "body", m.IgnoreCertificateErrors); err != nil {
		return err
	}

	return nil
}

func (m *MsaStep) validateRequestHeaders(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestHeaders) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestHeaders); i++ {
		if swag.IsZero(m.RequestHeaders[i]) { // not required
			continue
		}

		if m.RequestHeaders[i] != nil {
			if err := m.RequestHeaders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RequestHeaders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MsaStep) validateStepType(formats strfmt.Registry) error {

	if err := m.StepType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("StepType")
		}
		return err
	}

	return nil
}

func (m *MsaStep) validateUseFixedClientCertificate(formats strfmt.Registry) error {

	if err := validate.Required("UseFixedClientCertificate", "body", m.UseFixedClientCertificate); err != nil {
		return err
	}

	return nil
}

func (m *MsaStep) validateVariables(formats strfmt.Registry) error {

	if swag.IsZero(m.Variables) { // not required
		return nil
	}

	for i := 0; i < len(m.Variables); i++ {
		if swag.IsZero(m.Variables[i]) { // not required
			continue
		}

		if m.Variables[i] != nil {
			if err := m.Variables[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Variables" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MsaStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MsaStep) UnmarshalBinary(b []byte) error {
	var res MsaStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}