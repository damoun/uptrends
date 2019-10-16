// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StepTimingInfo step timing info
// swagger:model StepTimingInfo
type StepTimingInfo struct {

	// delay milliseconds
	// Required: true
	DelayMilliseconds *int64 `json:"DelayMilliseconds"`

	// description
	Description string `json:"Description,omitempty"`

	// elapsed milliseconds
	// Required: true
	ElapsedMilliseconds *int64 `json:"ElapsedMilliseconds"`

	// end utc
	// Required: true
	// Format: date-time
	EndUtc *strfmt.DateTime `json:"EndUtc"`

	// If true, this TimingInfo should be counted as part of the sum of its siblings. If false, the TimingInfo should be discarded (e.g. for PreDelays we don't want to count).
	// Required: true
	IsValid *bool `json:"IsValid"`

	// start utc
	// Required: true
	// Format: date-time
	StartUtc *strfmt.DateTime `json:"StartUtc"`

	// sub timing infos
	SubTimingInfos []*StepTimingInfo `json:"SubTimingInfos"`
}

// Validate validates this step timing info
func (m *StepTimingInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelayMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateElapsedMilliseconds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndUtc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsValid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartUtc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubTimingInfos(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StepTimingInfo) validateDelayMilliseconds(formats strfmt.Registry) error {

	if err := validate.Required("DelayMilliseconds", "body", m.DelayMilliseconds); err != nil {
		return err
	}

	return nil
}

func (m *StepTimingInfo) validateElapsedMilliseconds(formats strfmt.Registry) error {

	if err := validate.Required("ElapsedMilliseconds", "body", m.ElapsedMilliseconds); err != nil {
		return err
	}

	return nil
}

func (m *StepTimingInfo) validateEndUtc(formats strfmt.Registry) error {

	if err := validate.Required("EndUtc", "body", m.EndUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("EndUtc", "body", "date-time", m.EndUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StepTimingInfo) validateIsValid(formats strfmt.Registry) error {

	if err := validate.Required("IsValid", "body", m.IsValid); err != nil {
		return err
	}

	return nil
}

func (m *StepTimingInfo) validateStartUtc(formats strfmt.Registry) error {

	if err := validate.Required("StartUtc", "body", m.StartUtc); err != nil {
		return err
	}

	if err := validate.FormatOf("StartUtc", "body", "date-time", m.StartUtc.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *StepTimingInfo) validateSubTimingInfos(formats strfmt.Registry) error {

	if swag.IsZero(m.SubTimingInfos) { // not required
		return nil
	}

	for i := 0; i < len(m.SubTimingInfos); i++ {
		if swag.IsZero(m.SubTimingInfos[i]) { // not required
			continue
		}

		if m.SubTimingInfos[i] != nil {
			if err := m.SubTimingInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SubTimingInfos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StepTimingInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StepTimingInfo) UnmarshalBinary(b []byte) error {
	var res StepTimingInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
