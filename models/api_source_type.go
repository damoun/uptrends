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

// APISourceType Api source type
// swagger:model ApiSourceType
type APISourceType string

const (

	// APISourceTypeNone captures enum value "None"
	APISourceTypeNone APISourceType = "None"

	// APISourceTypeConstant captures enum value "Constant"
	APISourceTypeConstant APISourceType = "Constant"

	// APISourceTypeVariable captures enum value "Variable"
	APISourceTypeVariable APISourceType = "Variable"

	// APISourceTypeResponseStatusCode captures enum value "ResponseStatusCode"
	APISourceTypeResponseStatusCode APISourceType = "ResponseStatusCode"

	// APISourceTypeResponseStatusDescription captures enum value "ResponseStatusDescription"
	APISourceTypeResponseStatusDescription APISourceType = "ResponseStatusDescription"

	// APISourceTypeResponseBodyJSON captures enum value "ResponseBodyJson"
	APISourceTypeResponseBodyJSON APISourceType = "ResponseBodyJson"

	// APISourceTypeResponseCompleted captures enum value "ResponseCompleted"
	APISourceTypeResponseCompleted APISourceType = "ResponseCompleted"

	// APISourceTypeResponseBodyText captures enum value "ResponseBodyText"
	APISourceTypeResponseBodyText APISourceType = "ResponseBodyText"

	// APISourceTypeFail captures enum value "Fail"
	APISourceTypeFail APISourceType = "Fail"

	// APISourceTypeContentLengthCalculated captures enum value "ContentLengthCalculated"
	APISourceTypeContentLengthCalculated APISourceType = "ContentLengthCalculated"

	// APISourceTypeDuration captures enum value "Duration"
	APISourceTypeDuration APISourceType = "Duration"

	// APISourceTypeSum captures enum value "Sum"
	APISourceTypeSum APISourceType = "Sum"

	// APISourceTypeConcatenation captures enum value "Concatenation"
	APISourceTypeConcatenation APISourceType = "Concatenation"

	// APISourceTypeToBase64 captures enum value "ToBase64"
	APISourceTypeToBase64 APISourceType = "ToBase64"

	// APISourceTypeToSHA1Hex captures enum value "ToSHA1Hex"
	APISourceTypeToSHA1Hex APISourceType = "ToSHA1Hex"

	// APISourceTypeToMD5Hex captures enum value "ToMD5Hex"
	APISourceTypeToMD5Hex APISourceType = "ToMD5Hex"

	// APISourceTypeResponseBodyXML captures enum value "ResponseBodyXml"
	APISourceTypeResponseBodyXML APISourceType = "ResponseBodyXml"

	// APISourceTypeResponseHeader captures enum value "ResponseHeader"
	APISourceTypeResponseHeader APISourceType = "ResponseHeader"

	// APISourceTypeCookie captures enum value "Cookie"
	APISourceTypeCookie APISourceType = "Cookie"

	// APISourceTypeVariablesResolved captures enum value "VariablesResolved"
	APISourceTypeVariablesResolved APISourceType = "VariablesResolved"

	// APISourceTypeCumulativeDuration captures enum value "CumulativeDuration"
	APISourceTypeCumulativeDuration APISourceType = "CumulativeDuration"

	// APISourceTypeResponseHasException captures enum value "ResponseHasException"
	APISourceTypeResponseHasException APISourceType = "ResponseHasException"
)

// for schema
var apiSourceTypeEnum []interface{}

func init() {
	var res []APISourceType
	if err := json.Unmarshal([]byte(`["None","Constant","Variable","ResponseStatusCode","ResponseStatusDescription","ResponseBodyJson","ResponseCompleted","ResponseBodyText","Fail","ContentLengthCalculated","Duration","Sum","Concatenation","ToBase64","ToSHA1Hex","ToMD5Hex","ResponseBodyXml","ResponseHeader","Cookie","VariablesResolved","CumulativeDuration","ResponseHasException"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiSourceTypeEnum = append(apiSourceTypeEnum, v)
	}
}

func (m APISourceType) validateAPISourceTypeEnum(path, location string, value APISourceType) error {
	if err := validate.Enum(path, location, value, apiSourceTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this Api source type
func (m APISourceType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAPISourceTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
