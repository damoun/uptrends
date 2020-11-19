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

// ExclusionPeriod exclusion period
//
// swagger:model ExclusionPeriod
type ExclusionPeriod struct {

	// description
	Description string `json:"Description,omitempty"`

	// exclusion period Id
	ExclusionPeriodID int32 `json:"ExclusionPeriodId,omitempty"`

	// from
	// Format: date-time
	From strfmt.DateTime `json:"From,omitempty"`

	// hash
	Hash string `json:"Hash,omitempty"`

	// until
	// Format: date-time
	Until strfmt.DateTime `json:"Until,omitempty"`
}

// Validate validates this exclusion period
func (m *ExclusionPeriod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExclusionPeriod) validateFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.From) { // not required
		return nil
	}

	if err := validate.FormatOf("From", "body", "date-time", m.From.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ExclusionPeriod) validateUntil(formats strfmt.Registry) error {

	if swag.IsZero(m.Until) { // not required
		return nil
	}

	if err := validate.FormatOf("Until", "body", "date-time", m.Until.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExclusionPeriod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExclusionPeriod) UnmarshalBinary(b []byte) error {
	var res ExclusionPeriod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
