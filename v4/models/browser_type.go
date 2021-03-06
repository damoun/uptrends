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

// BrowserType browser type
//
// swagger:model BrowserType
type BrowserType string

const (

	// BrowserTypeIE captures enum value "IE"
	BrowserTypeIE BrowserType = "IE"

	// BrowserTypeFirefox captures enum value "Firefox"
	BrowserTypeFirefox BrowserType = "Firefox"

	// BrowserTypeChrome captures enum value "Chrome"
	BrowserTypeChrome BrowserType = "Chrome"

	// BrowserTypeSafari captures enum value "Safari"
	BrowserTypeSafari BrowserType = "Safari"

	// BrowserTypePhantomJS captures enum value "PhantomJS"
	BrowserTypePhantomJS BrowserType = "PhantomJS"

	// BrowserTypePhantomJS20 captures enum value "PhantomJS20"
	BrowserTypePhantomJS20 BrowserType = "PhantomJS20"
)

// for schema
var browserTypeEnum []interface{}

func init() {
	var res []BrowserType
	if err := json.Unmarshal([]byte(`["IE","Firefox","Chrome","Safari","PhantomJS","PhantomJS20"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		browserTypeEnum = append(browserTypeEnum, v)
	}
}

func (m BrowserType) validateBrowserTypeEnum(path, location string, value BrowserType) error {
	if err := validate.EnumCase(path, location, value, browserTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this browser type
func (m BrowserType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBrowserTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
