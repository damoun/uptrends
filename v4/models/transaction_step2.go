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

// TransactionStep2 transaction step2
//
// swagger:model TransactionStep2
type TransactionStep2 struct {

	// has screenshot
	// Required: true
	HasScreenshot *bool `json:"HasScreenshot"`

	// has waterfall
	// Required: true
	HasWaterfall *bool `json:"HasWaterfall"`

	// name
	Name string `json:"Name,omitempty"`

	// sub steps
	SubSteps []*TransactionSubStep `json:"SubSteps"`
}

// Validate validates this transaction step2
func (m *TransactionStep2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasScreenshot(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasWaterfall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubSteps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransactionStep2) validateHasScreenshot(formats strfmt.Registry) error {

	if err := validate.Required("HasScreenshot", "body", m.HasScreenshot); err != nil {
		return err
	}

	return nil
}

func (m *TransactionStep2) validateHasWaterfall(formats strfmt.Registry) error {

	if err := validate.Required("HasWaterfall", "body", m.HasWaterfall); err != nil {
		return err
	}

	return nil
}

func (m *TransactionStep2) validateSubSteps(formats strfmt.Registry) error {

	if swag.IsZero(m.SubSteps) { // not required
		return nil
	}

	for i := 0; i < len(m.SubSteps); i++ {
		if swag.IsZero(m.SubSteps[i]) { // not required
			continue
		}

		if m.SubSteps[i] != nil {
			if err := m.SubSteps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SubSteps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransactionStep2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransactionStep2) UnmarshalBinary(b []byte) error {
	var res TransactionStep2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
