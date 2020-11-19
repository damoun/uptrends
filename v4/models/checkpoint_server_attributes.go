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

// CheckpointServerAttributes Checkpoint server attributes
//
// swagger:model CheckpointServerAttributes
type CheckpointServerAttributes struct {

	// List of server's capabilities
	Capabilities []CapabilityFilterEnum `json:"Capabilities"`

	// The ipv4 address
	Ip4Address string `json:"Ip4Address,omitempty"`

	// The ipv6 address
	IPV6Address string `json:"IpV6Address,omitempty"`
}

// Validate validates this checkpoint server attributes
func (m *CheckpointServerAttributes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CheckpointServerAttributes) validateCapabilities(formats strfmt.Registry) error {

	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Capabilities); i++ {

		if err := m.Capabilities[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Capabilities" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CheckpointServerAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CheckpointServerAttributes) UnmarshalBinary(b []byte) error {
	var res CheckpointServerAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
