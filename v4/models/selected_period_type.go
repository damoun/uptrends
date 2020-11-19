// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// SelectedPeriodType selected period type
//
// swagger:model SelectedPeriodType
type SelectedPeriodType string

const (

	// SelectedPeriodTypePresetPeriod captures enum value "PresetPeriod"
	SelectedPeriodTypePresetPeriod SelectedPeriodType = "PresetPeriod"

	// SelectedPeriodTypeSpecificDates captures enum value "SpecificDates"
	SelectedPeriodTypeSpecificDates SelectedPeriodType = "SpecificDates"
)

// for schema
var selectedPeriodTypeEnum []interface{}

func init() {
	var res []SelectedPeriodType
	if err := json.Unmarshal([]byte(`["PresetPeriod","SpecificDates"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		selectedPeriodTypeEnum = append(selectedPeriodTypeEnum, v)
	}
}

func (m SelectedPeriodType) validateSelectedPeriodTypeEnum(path, location string, value SelectedPeriodType) error {
	if err := validate.EnumCase(path, location, value, selectedPeriodTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this selected period type
func (m SelectedPeriodType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSelectedPeriodTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
