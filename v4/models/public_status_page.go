// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PublicStatusPage public status page
//
// swagger:model PublicStatusPage
type PublicStatusPage struct {

	// customization info
	CustomizationInfo *CustomizationInfo `json:"CustomizationInfo,omitempty"`

	// is published
	IsPublished bool `json:"IsPublished,omitempty"`

	// monitor group guids
	MonitorGroupGuids []string `json:"MonitorGroupGuids"`

	// monitor guids
	MonitorGuids []string `json:"MonitorGuids"`

	// name
	Name string `json:"Name,omitempty"`

	// preset period type
	PresetPeriodType PresetPeriodType `json:"PresetPeriodType,omitempty"`

	// public dashboard Guid
	PublicDashboardGUID string `json:"PublicDashboardGuid,omitempty"`

	// Sla Guid
	SLAGUID string `json:"SlaGuid,omitempty"`

	// Sla Guid specified
	SLAGUIDSpecified bool `json:"SlaGuidSpecified,omitempty"`
}

// Validate validates this public status page
func (m *PublicStatusPage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomizationInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresetPeriodType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PublicStatusPage) validateCustomizationInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomizationInfo) { // not required
		return nil
	}

	if m.CustomizationInfo != nil {
		if err := m.CustomizationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("CustomizationInfo")
			}
			return err
		}
	}

	return nil
}

func (m *PublicStatusPage) validatePresetPeriodType(formats strfmt.Registry) error {

	if swag.IsZero(m.PresetPeriodType) { // not required
		return nil
	}

	if err := m.PresetPeriodType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("PresetPeriodType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PublicStatusPage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PublicStatusPage) UnmarshalBinary(b []byte) error {
	var res PublicStatusPage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
