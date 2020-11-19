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

// PresetPeriodTypeWithExclusive preset period type with exclusive
//
// swagger:model PresetPeriodTypeWithExclusive
type PresetPeriodTypeWithExclusive string

const (

	// PresetPeriodTypeWithExclusiveCurrentDay captures enum value "CurrentDay"
	PresetPeriodTypeWithExclusiveCurrentDay PresetPeriodTypeWithExclusive = "CurrentDay"

	// PresetPeriodTypeWithExclusiveCurrentWeek captures enum value "CurrentWeek"
	PresetPeriodTypeWithExclusiveCurrentWeek PresetPeriodTypeWithExclusive = "CurrentWeek"

	// PresetPeriodTypeWithExclusiveCurrentMonth captures enum value "CurrentMonth"
	PresetPeriodTypeWithExclusiveCurrentMonth PresetPeriodTypeWithExclusive = "CurrentMonth"

	// PresetPeriodTypeWithExclusiveCurrentYear captures enum value "CurrentYear"
	PresetPeriodTypeWithExclusiveCurrentYear PresetPeriodTypeWithExclusive = "CurrentYear"

	// PresetPeriodTypeWithExclusivePreviousDay captures enum value "PreviousDay"
	PresetPeriodTypeWithExclusivePreviousDay PresetPeriodTypeWithExclusive = "PreviousDay"

	// PresetPeriodTypeWithExclusivePreviousWeek captures enum value "PreviousWeek"
	PresetPeriodTypeWithExclusivePreviousWeek PresetPeriodTypeWithExclusive = "PreviousWeek"

	// PresetPeriodTypeWithExclusivePreviousMonth captures enum value "PreviousMonth"
	PresetPeriodTypeWithExclusivePreviousMonth PresetPeriodTypeWithExclusive = "PreviousMonth"

	// PresetPeriodTypeWithExclusivePreviousYear captures enum value "PreviousYear"
	PresetPeriodTypeWithExclusivePreviousYear PresetPeriodTypeWithExclusive = "PreviousYear"

	// PresetPeriodTypeWithExclusiveLast24Hours captures enum value "Last24Hours"
	PresetPeriodTypeWithExclusiveLast24Hours PresetPeriodTypeWithExclusive = "Last24Hours"

	// PresetPeriodTypeWithExclusiveLast48Hours captures enum value "Last48Hours"
	PresetPeriodTypeWithExclusiveLast48Hours PresetPeriodTypeWithExclusive = "Last48Hours"

	// PresetPeriodTypeWithExclusiveLast7Days captures enum value "Last7Days"
	PresetPeriodTypeWithExclusiveLast7Days PresetPeriodTypeWithExclusive = "Last7Days"

	// PresetPeriodTypeWithExclusiveLast30Days captures enum value "Last30Days"
	PresetPeriodTypeWithExclusiveLast30Days PresetPeriodTypeWithExclusive = "Last30Days"

	// PresetPeriodTypeWithExclusiveLast60Days captures enum value "Last60Days"
	PresetPeriodTypeWithExclusiveLast60Days PresetPeriodTypeWithExclusive = "Last60Days"

	// PresetPeriodTypeWithExclusiveLast90Days captures enum value "Last90Days"
	PresetPeriodTypeWithExclusiveLast90Days PresetPeriodTypeWithExclusive = "Last90Days"

	// PresetPeriodTypeWithExclusiveLast365Days captures enum value "Last365Days"
	PresetPeriodTypeWithExclusiveLast365Days PresetPeriodTypeWithExclusive = "Last365Days"

	// PresetPeriodTypeWithExclusiveLast3Months captures enum value "Last3Months"
	PresetPeriodTypeWithExclusiveLast3Months PresetPeriodTypeWithExclusive = "Last3Months"

	// PresetPeriodTypeWithExclusiveLast6Months captures enum value "Last6Months"
	PresetPeriodTypeWithExclusiveLast6Months PresetPeriodTypeWithExclusive = "Last6Months"

	// PresetPeriodTypeWithExclusiveLast12Months captures enum value "Last12Months"
	PresetPeriodTypeWithExclusiveLast12Months PresetPeriodTypeWithExclusive = "Last12Months"

	// PresetPeriodTypeWithExclusiveLast24Months captures enum value "Last24Months"
	PresetPeriodTypeWithExclusiveLast24Months PresetPeriodTypeWithExclusive = "Last24Months"

	// PresetPeriodTypeWithExclusiveLast24HoursExcl captures enum value "Last24HoursExcl"
	PresetPeriodTypeWithExclusiveLast24HoursExcl PresetPeriodTypeWithExclusive = "Last24HoursExcl"

	// PresetPeriodTypeWithExclusiveLast48HoursExcl captures enum value "Last48HoursExcl"
	PresetPeriodTypeWithExclusiveLast48HoursExcl PresetPeriodTypeWithExclusive = "Last48HoursExcl"

	// PresetPeriodTypeWithExclusiveLast7DaysExcl captures enum value "Last7DaysExcl"
	PresetPeriodTypeWithExclusiveLast7DaysExcl PresetPeriodTypeWithExclusive = "Last7DaysExcl"

	// PresetPeriodTypeWithExclusiveLast30DaysExcl captures enum value "Last30DaysExcl"
	PresetPeriodTypeWithExclusiveLast30DaysExcl PresetPeriodTypeWithExclusive = "Last30DaysExcl"

	// PresetPeriodTypeWithExclusiveLast60DaysExcl captures enum value "Last60DaysExcl"
	PresetPeriodTypeWithExclusiveLast60DaysExcl PresetPeriodTypeWithExclusive = "Last60DaysExcl"

	// PresetPeriodTypeWithExclusiveLast90DaysExcl captures enum value "Last90DaysExcl"
	PresetPeriodTypeWithExclusiveLast90DaysExcl PresetPeriodTypeWithExclusive = "Last90DaysExcl"

	// PresetPeriodTypeWithExclusiveLast365DaysExcl captures enum value "Last365DaysExcl"
	PresetPeriodTypeWithExclusiveLast365DaysExcl PresetPeriodTypeWithExclusive = "Last365DaysExcl"

	// PresetPeriodTypeWithExclusiveLast3MonthsExcl captures enum value "Last3MonthsExcl"
	PresetPeriodTypeWithExclusiveLast3MonthsExcl PresetPeriodTypeWithExclusive = "Last3MonthsExcl"

	// PresetPeriodTypeWithExclusiveLast6MonthsExcl captures enum value "Last6MonthsExcl"
	PresetPeriodTypeWithExclusiveLast6MonthsExcl PresetPeriodTypeWithExclusive = "Last6MonthsExcl"

	// PresetPeriodTypeWithExclusiveLast12MonthsExcl captures enum value "Last12MonthsExcl"
	PresetPeriodTypeWithExclusiveLast12MonthsExcl PresetPeriodTypeWithExclusive = "Last12MonthsExcl"

	// PresetPeriodTypeWithExclusiveLast24MonthsExcl captures enum value "Last24MonthsExcl"
	PresetPeriodTypeWithExclusiveLast24MonthsExcl PresetPeriodTypeWithExclusive = "Last24MonthsExcl"

	// PresetPeriodTypeWithExclusiveAll captures enum value "All"
	PresetPeriodTypeWithExclusiveAll PresetPeriodTypeWithExclusive = "All"
)

// for schema
var presetPeriodTypeWithExclusiveEnum []interface{}

func init() {
	var res []PresetPeriodTypeWithExclusive
	if err := json.Unmarshal([]byte(`["CurrentDay","CurrentWeek","CurrentMonth","CurrentYear","PreviousDay","PreviousWeek","PreviousMonth","PreviousYear","Last24Hours","Last48Hours","Last7Days","Last30Days","Last60Days","Last90Days","Last365Days","Last3Months","Last6Months","Last12Months","Last24Months","Last24HoursExcl","Last48HoursExcl","Last7DaysExcl","Last30DaysExcl","Last60DaysExcl","Last90DaysExcl","Last365DaysExcl","Last3MonthsExcl","Last6MonthsExcl","Last12MonthsExcl","Last24MonthsExcl","All"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		presetPeriodTypeWithExclusiveEnum = append(presetPeriodTypeWithExclusiveEnum, v)
	}
}

func (m PresetPeriodTypeWithExclusive) validatePresetPeriodTypeWithExclusiveEnum(path, location string, value PresetPeriodTypeWithExclusive) error {
	if err := validate.EnumCase(path, location, value, presetPeriodTypeWithExclusiveEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this preset period type with exclusive
func (m PresetPeriodTypeWithExclusive) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePresetPeriodTypeWithExclusiveEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}