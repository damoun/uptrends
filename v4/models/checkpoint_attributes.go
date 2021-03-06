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

// CheckpointAttributes Checkpoint attributes
//
// swagger:model CheckpointAttributes
type CheckpointAttributes struct {

	// The name of the checkpoint
	CheckpointName string `json:"CheckpointName,omitempty"`

	// Location code of the checkpoint
	Code string `json:"Code,omitempty"`

	// Checkpoint has high availability
	// Required: true
	HasHighAvailability *bool `json:"HasHighAvailability"`

	// Ipv6 addresses of the checkpoint
	IPV6Addresses []*IPV6Address `json:"IpV6Addresses"`

	// Ipv4 addresses of the checkpoint
	IPV4Addresses []string `json:"Ipv4Addresses"`

	// Checkpoint is primary
	// Required: true
	IsPrimaryCheckpoint *bool `json:"IsPrimaryCheckpoint"`

	// Checkpoint supports IPv6
	// Required: true
	SupportsIPV6 *bool `json:"SupportsIpv6"`
}

// Validate validates this checkpoint attributes
func (m *CheckpointAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasHighAvailability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6Addresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsPrimaryCheckpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportsIPV6(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckpointAttributes) validateHasHighAvailability(formats strfmt.Registry) error {

	if err := validate.Required("HasHighAvailability", "body", m.HasHighAvailability); err != nil {
		return err
	}

	return nil
}

func (m *CheckpointAttributes) validateIPV6Addresses(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.IPV6Addresses); i++ {
		if swag.IsZero(m.IPV6Addresses[i]) { // not required
			continue
		}

		if m.IPV6Addresses[i] != nil {
			if err := m.IPV6Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("IpV6Addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CheckpointAttributes) validateIsPrimaryCheckpoint(formats strfmt.Registry) error {

	if err := validate.Required("IsPrimaryCheckpoint", "body", m.IsPrimaryCheckpoint); err != nil {
		return err
	}

	return nil
}

func (m *CheckpointAttributes) validateSupportsIPV6(formats strfmt.Registry) error {

	if err := validate.Required("SupportsIpv6", "body", m.SupportsIPV6); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckpointAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckpointAttributes) UnmarshalBinary(b []byte) error {
	var res CheckpointAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
