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

// ScheduleMode schedule mode
//
// swagger:model ScheduleMode
type ScheduleMode string

const (

	// ScheduleModeOneTime captures enum value "OneTime"
	ScheduleModeOneTime ScheduleMode = "OneTime"

	// ScheduleModeDaily captures enum value "Daily"
	ScheduleModeDaily ScheduleMode = "Daily"

	// ScheduleModeWeekly captures enum value "Weekly"
	ScheduleModeWeekly ScheduleMode = "Weekly"

	// ScheduleModeMonthly captures enum value "Monthly"
	ScheduleModeMonthly ScheduleMode = "Monthly"
)

// for schema
var scheduleModeEnum []interface{}

func init() {
	var res []ScheduleMode
	if err := json.Unmarshal([]byte(`["OneTime","Daily","Weekly","Monthly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleModeEnum = append(scheduleModeEnum, v)
	}
}

func (m ScheduleMode) validateScheduleModeEnum(path, location string, value ScheduleMode) error {
	if err := validate.EnumCase(path, location, value, scheduleModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this schedule mode
func (m ScheduleMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateScheduleModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
