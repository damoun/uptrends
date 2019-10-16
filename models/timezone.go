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

// Timezone timezone
// swagger:model Timezone
type Timezone struct {

	// daylight saving offset
	DaylightSavingOffset int64 `json:"DaylightSavingOffset,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// has daylight saving
	// Required: true
	HasDaylightSaving *bool `json:"HasDaylightSaving"`

	// offset from utc
	// Required: true
	OffsetFromUtc *int64 `json:"OffsetFromUtc"`

	// timezone Id
	// Required: true
	TimezoneID *int64 `json:"TimezoneId"`
}

// Validate validates this timezone
func (m *Timezone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasDaylightSaving(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOffsetFromUtc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimezoneID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Timezone) validateHasDaylightSaving(formats strfmt.Registry) error {

	if err := validate.Required("HasDaylightSaving", "body", m.HasDaylightSaving); err != nil {
		return err
	}

	return nil
}

func (m *Timezone) validateOffsetFromUtc(formats strfmt.Registry) error {

	if err := validate.Required("OffsetFromUtc", "body", m.OffsetFromUtc); err != nil {
		return err
	}

	return nil
}

func (m *Timezone) validateTimezoneID(formats strfmt.Registry) error {

	if err := validate.Required("TimezoneId", "body", m.TimezoneID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Timezone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Timezone) UnmarshalBinary(b []byte) error {
	var res Timezone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}