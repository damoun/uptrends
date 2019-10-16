// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// SortDirection sort direction
// swagger:model SortDirection
type SortDirection string

const (

	// SortDirectionAscending captures enum value "Ascending"
	SortDirectionAscending SortDirection = "Ascending"

	// SortDirectionDescending captures enum value "Descending"
	SortDirectionDescending SortDirection = "Descending"
)

// for schema
var sortDirectionEnum []interface{}

func init() {
	var res []SortDirection
	if err := json.Unmarshal([]byte(`["Ascending","Descending"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		sortDirectionEnum = append(sortDirectionEnum, v)
	}
}

func (m SortDirection) validateSortDirectionEnum(path, location string, value SortDirection) error {
	if err := validate.Enum(path, location, value, sortDirectionEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this sort direction
func (m SortDirection) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateSortDirectionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}