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

// SelectedPeriod selected period
//
// swagger:model SelectedPeriod
type SelectedPeriod struct {

	// The end of a custom period
	// Format: date-time
	End strfmt.DateTime `json:"End,omitempty"`

	// The requested time period.
	PresetPeriod struct {
		PresetPeriodType
	} `json:"PresetPeriod,omitempty"`

	// The type of period
	SelectedPeriodType struct {
		SelectedPeriodType
	} `json:"SelectedPeriodType,omitempty"`

	// The start of a custom period (can't be used together with the SelectedPeriodType parameter)
	// Format: date-time
	Start strfmt.DateTime `json:"Start,omitempty"`
}

// Validate validates this selected period
func (m *SelectedPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePresetPeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectedPeriodType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SelectedPeriod) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("End", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SelectedPeriod) validatePresetPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.PresetPeriod) { // not required
		return nil
	}

	return nil
}

func (m *SelectedPeriod) validateSelectedPeriodType(formats strfmt.Registry) error {

	if swag.IsZero(m.SelectedPeriodType) { // not required
		return nil
	}

	return nil
}

func (m *SelectedPeriod) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("Start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SelectedPeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SelectedPeriod) UnmarshalBinary(b []byte) error {
	var res SelectedPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
