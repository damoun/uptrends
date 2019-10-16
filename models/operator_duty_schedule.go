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

// OperatorDutySchedule operator duty schedule
// swagger:model OperatorDutySchedule
type OperatorDutySchedule struct {

	// end date time
	// Format: date-time
	EndDateTime strfmt.DateTime `json:"EndDateTime,omitempty"`

	// end time
	EndTime string `json:"EndTime,omitempty"`

	// Id
	// Required: true
	ID *int32 `json:"Id"`

	// month day
	MonthDay int32 `json:"MonthDay,omitempty"`

	// schedule mode
	// Required: true
	ScheduleMode OperatorScheduleMode `json:"ScheduleMode"`

	// start date time
	// Format: date-time
	StartDateTime strfmt.DateTime `json:"StartDateTime,omitempty"`

	// start time
	StartTime string `json:"StartTime,omitempty"`

	// week day
	WeekDay DayOfWeek `json:"WeekDay,omitempty"`
}

// Validate validates this operator duty schedule
func (m *OperatorDutySchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekDay(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OperatorDutySchedule) validateEndDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("EndDateTime", "body", "date-time", m.EndDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OperatorDutySchedule) validateID(formats strfmt.Registry) error {

	if err := validate.Required("Id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *OperatorDutySchedule) validateScheduleMode(formats strfmt.Registry) error {

	if err := m.ScheduleMode.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("ScheduleMode")
		}
		return err
	}

	return nil
}

func (m *OperatorDutySchedule) validateStartDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("StartDateTime", "body", "date-time", m.StartDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OperatorDutySchedule) validateWeekDay(formats strfmt.Registry) error {

	if swag.IsZero(m.WeekDay) { // not required
		return nil
	}

	if err := m.WeekDay.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("WeekDay")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OperatorDutySchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OperatorDutySchedule) UnmarshalBinary(b []byte) error {
	var res OperatorDutySchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
