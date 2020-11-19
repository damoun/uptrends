// Code generated by go-swagger; DO NOT EDIT.

package alert_definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// AlertDefinitionPatchAlertDefinitionReader is a Reader for the AlertDefinitionPatchAlertDefinition structure.
type AlertDefinitionPatchAlertDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertDefinitionPatchAlertDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAlertDefinitionPatchAlertDefinitionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertDefinitionPatchAlertDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertDefinitionPatchAlertDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertDefinitionPatchAlertDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertDefinitionPatchAlertDefinitionNoContent creates a AlertDefinitionPatchAlertDefinitionNoContent with default headers values
func NewAlertDefinitionPatchAlertDefinitionNoContent() *AlertDefinitionPatchAlertDefinitionNoContent {
	return &AlertDefinitionPatchAlertDefinitionNoContent{}
}

/*AlertDefinitionPatchAlertDefinitionNoContent handles this case with default header values.

The request completed successfully. No content is returned.
*/
type AlertDefinitionPatchAlertDefinitionNoContent struct {
}

func (o *AlertDefinitionPatchAlertDefinitionNoContent) Error() string {
	return fmt.Sprintf("[PATCH /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionPatchAlertDefinitionNoContent ", 204)
}

func (o *AlertDefinitionPatchAlertDefinitionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAlertDefinitionPatchAlertDefinitionBadRequest creates a AlertDefinitionPatchAlertDefinitionBadRequest with default headers values
func NewAlertDefinitionPatchAlertDefinitionBadRequest() *AlertDefinitionPatchAlertDefinitionBadRequest {
	return &AlertDefinitionPatchAlertDefinitionBadRequest{}
}

/*AlertDefinitionPatchAlertDefinitionBadRequest handles this case with default header values.

The request failed.
*/
type AlertDefinitionPatchAlertDefinitionBadRequest struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionPatchAlertDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionPatchAlertDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *AlertDefinitionPatchAlertDefinitionBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionPatchAlertDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionPatchAlertDefinitionForbidden creates a AlertDefinitionPatchAlertDefinitionForbidden with default headers values
func NewAlertDefinitionPatchAlertDefinitionForbidden() *AlertDefinitionPatchAlertDefinitionForbidden {
	return &AlertDefinitionPatchAlertDefinitionForbidden{}
}

/*AlertDefinitionPatchAlertDefinitionForbidden handles this case with default header values.

One or more validation errors occurred.
*/
type AlertDefinitionPatchAlertDefinitionForbidden struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionPatchAlertDefinitionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionPatchAlertDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *AlertDefinitionPatchAlertDefinitionForbidden) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionPatchAlertDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionPatchAlertDefinitionNotFound creates a AlertDefinitionPatchAlertDefinitionNotFound with default headers values
func NewAlertDefinitionPatchAlertDefinitionNotFound() *AlertDefinitionPatchAlertDefinitionNotFound {
	return &AlertDefinitionPatchAlertDefinitionNotFound{}
}

/*AlertDefinitionPatchAlertDefinitionNotFound handles this case with default header values.

The specified alert definition could not be found.
*/
type AlertDefinitionPatchAlertDefinitionNotFound struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionPatchAlertDefinitionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionPatchAlertDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *AlertDefinitionPatchAlertDefinitionNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionPatchAlertDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
