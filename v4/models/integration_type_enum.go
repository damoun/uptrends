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

// IntegrationTypeEnum integration type enum
//
// swagger:model IntegrationTypeEnum
type IntegrationTypeEnum string

const (

	// IntegrationTypeEnumSlack captures enum value "Slack"
	IntegrationTypeEnumSlack IntegrationTypeEnum = "Slack"

	// IntegrationTypeEnumPagerDuty captures enum value "PagerDuty"
	IntegrationTypeEnumPagerDuty IntegrationTypeEnum = "PagerDuty"

	// IntegrationTypeEnumSms captures enum value "Sms"
	IntegrationTypeEnumSms IntegrationTypeEnum = "Sms"

	// IntegrationTypeEnumEmail captures enum value "Email"
	IntegrationTypeEnumEmail IntegrationTypeEnum = "Email"

	// IntegrationTypeEnumPhone captures enum value "Phone"
	IntegrationTypeEnumPhone IntegrationTypeEnum = "Phone"

	// IntegrationTypeEnumStatushub captures enum value "Statushub"
	IntegrationTypeEnumStatushub IntegrationTypeEnum = "Statushub"

	// IntegrationTypeEnumGenericWebhook captures enum value "GenericWebhook"
	IntegrationTypeEnumGenericWebhook IntegrationTypeEnum = "GenericWebhook"
)

// for schema
var integrationTypeEnumEnum []interface{}

func init() {
	var res []IntegrationTypeEnum
	if err := json.Unmarshal([]byte(`["Slack","PagerDuty","Sms","Email","Phone","Statushub","GenericWebhook"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		integrationTypeEnumEnum = append(integrationTypeEnumEnum, v)
	}
}

func (m IntegrationTypeEnum) validateIntegrationTypeEnumEnum(path, location string, value IntegrationTypeEnum) error {
	if err := validate.EnumCase(path, location, value, integrationTypeEnumEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this integration type enum
func (m IntegrationTypeEnum) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIntegrationTypeEnumEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
