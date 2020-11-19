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

// AuthorizationTypeWithoutContext authorization type without context
//
// swagger:model AuthorizationTypeWithoutContext
type AuthorizationTypeWithoutContext string

const (

	// AuthorizationTypeWithoutContextAccountAccess captures enum value "AccountAccess"
	AuthorizationTypeWithoutContextAccountAccess AuthorizationTypeWithoutContext = "AccountAccess"
)

// for schema
var authorizationTypeWithoutContextEnum []interface{}

func init() {
	var res []AuthorizationTypeWithoutContext
	if err := json.Unmarshal([]byte(`["AccountAccess"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authorizationTypeWithoutContextEnum = append(authorizationTypeWithoutContextEnum, v)
	}
}

func (m AuthorizationTypeWithoutContext) validateAuthorizationTypeWithoutContextEnum(path, location string, value AuthorizationTypeWithoutContext) error {
	if err := validate.EnumCase(path, location, value, authorizationTypeWithoutContextEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this authorization type without context
func (m AuthorizationTypeWithoutContext) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAuthorizationTypeWithoutContextEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}