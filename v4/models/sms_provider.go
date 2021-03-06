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

// SmsProvider sms provider
//
// swagger:model SmsProvider
type SmsProvider string

const (

	// SmsProviderUseAccountSetting captures enum value "UseAccountSetting"
	SmsProviderUseAccountSetting SmsProvider = "UseAccountSetting"

	// SmsProviderSmsProviderEurope captures enum value "SmsProviderEurope"
	SmsProviderSmsProviderEurope SmsProvider = "SmsProviderEurope"

	// SmsProviderSmsProviderInternational captures enum value "SmsProviderInternational"
	SmsProviderSmsProviderInternational SmsProvider = "SmsProviderInternational"

	// SmsProviderSmsProviderEurope2 captures enum value "SmsProviderEurope2"
	SmsProviderSmsProviderEurope2 SmsProvider = "SmsProviderEurope2"

	// SmsProviderSmsProviderUSA captures enum value "SmsProviderUSA"
	SmsProviderSmsProviderUSA SmsProvider = "SmsProviderUSA"
)

// for schema
var smsProviderEnum []interface{}

func init() {
	var res []SmsProvider
	if err := json.Unmarshal([]byte(`["UseAccountSetting","SmsProviderEurope","SmsProviderInternational","SmsProviderEurope2","SmsProviderUSA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smsProviderEnum = append(smsProviderEnum, v)
	}
}

func (m SmsProvider) validateSmsProviderEnum(path, location string, value SmsProvider) error {
	if err := validate.EnumCase(path, location, value, smsProviderEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this sms provider
func (m SmsProvider) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSmsProviderEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
