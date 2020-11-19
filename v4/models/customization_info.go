// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomizationInfo customization info
//
// swagger:model CustomizationInfo
type CustomizationInfo struct {

	// background color
	BackgroundColor string `json:"BackgroundColor,omitempty"`

	// comment text
	CommentText string `json:"CommentText,omitempty"`

	// comment title
	CommentTitle string `json:"CommentTitle,omitempty"`

	// footer text
	FooterText string `json:"FooterText,omitempty"`

	// main color
	MainColor string `json:"MainColor,omitempty"`

	// sort columns new to old
	// Required: true
	SortColumnsNewToOld *bool `json:"SortColumnsNewToOld"`

	// sort rows property
	// Required: true
	SortRowsProperty SortOrderEnum `json:"SortRowsProperty"`

	// text color
	TextColor string `json:"TextColor,omitempty"`

	// title text
	TitleText string `json:"TitleText,omitempty"`
}

// Validate validates this customization info
func (m *CustomizationInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSortColumnsNewToOld(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSortRowsProperty(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomizationInfo) validateSortColumnsNewToOld(formats strfmt.Registry) error {

	if err := validate.Required("SortColumnsNewToOld", "body", m.SortColumnsNewToOld); err != nil {
		return err
	}

	return nil
}

func (m *CustomizationInfo) validateSortRowsProperty(formats strfmt.Registry) error {

	if err := m.SortRowsProperty.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("SortRowsProperty")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomizationInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomizationInfo) UnmarshalBinary(b []byte) error {
	var res CustomizationInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
