// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EscalationLevelIntegration escalation level integration
//
// swagger:model EscalationLevelIntegration
type EscalationLevelIntegration struct {

	// Extra email addresses
	ExtraEmailAddresses []string `json:"ExtraEmailAddresses"`

	// Specified Extra EmailAddresses For Patch request
	ExtraEmailAddressesSpecified bool `json:"ExtraEmailAddressesSpecified,omitempty"`

	// hash
	Hash string `json:"Hash,omitempty"`

	// The unique key of this Integration.
	IntegrationGUID string `json:"IntegrationGuid,omitempty"`

	// Indicates whether this Integration is active.
	IsActive bool `json:"IsActive,omitempty"`

	// StatusHub Service Mapping
	StatusHubServiceList []*IntegrationServiceMap `json:"StatusHubServiceList"`

	// Specified StatusHubServiceList For Patch request
	StatusHubServiceListSpecified bool `json:"StatusHubServiceListSpecified,omitempty"`
}

// Validate validates this escalation level integration
func (m *EscalationLevelIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatusHubServiceList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EscalationLevelIntegration) validateStatusHubServiceList(formats strfmt.Registry) error {

	if swag.IsZero(m.StatusHubServiceList) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusHubServiceList); i++ {
		if swag.IsZero(m.StatusHubServiceList[i]) { // not required
			continue
		}

		if m.StatusHubServiceList[i] != nil {
			if err := m.StatusHubServiceList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StatusHubServiceList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EscalationLevelIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EscalationLevelIntegration) UnmarshalBinary(b []byte) error {
	var res EscalationLevelIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
