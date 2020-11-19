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

// AlertDefinitionPutAlertDefinitionReader is a Reader for the AlertDefinitionPutAlertDefinition structure.
type AlertDefinitionPutAlertDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertDefinitionPutAlertDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAlertDefinitionPutAlertDefinitionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertDefinitionPutAlertDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertDefinitionPutAlertDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertDefinitionPutAlertDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertDefinitionPutAlertDefinitionNoContent creates a AlertDefinitionPutAlertDefinitionNoContent with default headers values
func NewAlertDefinitionPutAlertDefinitionNoContent() *AlertDefinitionPutAlertDefinitionNoContent {
	return &AlertDefinitionPutAlertDefinitionNoContent{}
}

/*AlertDefinitionPutAlertDefinitionNoContent handles this case with default header values.

The request completed successfully. No content is returned.
*/
type AlertDefinitionPutAlertDefinitionNoContent struct {
}

func (o *AlertDefinitionPutAlertDefinitionNoContent) Error() string {
	return fmt.Sprintf("[PUT /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionPutAlertDefinitionNoContent ", 204)
}

func (o *AlertDefinitionPutAlertDefinitionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAlertDefinitionPutAlertDefinitionBadRequest creates a AlertDefinitionPutAlertDefinitionBadRequest with default headers values
func NewAlertDefinitionPutAlertDefinitionBadRequest() *AlertDefinitionPutAlertDefinitionBadRequest {
	return &AlertDefinitionPutAlertDefinitionBadRequest{}
}

/*AlertDefinitionPutAlertDefinitionBadRequest handles this case with default header values.

The request failed.
*/
type AlertDefinitionPutAlertDefinitionBadRequest struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionPutAlertDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionPutAlertDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *AlertDefinitionPutAlertDefinitionBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionPutAlertDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionPutAlertDefinitionForbidden creates a AlertDefinitionPutAlertDefinitionForbidden with default headers values
func NewAlertDefinitionPutAlertDefinitionForbidden() *AlertDefinitionPutAlertDefinitionForbidden {
	return &AlertDefinitionPutAlertDefinitionForbidden{}
}

/*AlertDefinitionPutAlertDefinitionForbidden handles this case with default header values.

One or more validation errors occurred.
*/
type AlertDefinitionPutAlertDefinitionForbidden struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionPutAlertDefinitionForbidden) Error() string {
	return fmt.Sprintf("[PUT /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionPutAlertDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *AlertDefinitionPutAlertDefinitionForbidden) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionPutAlertDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionPutAlertDefinitionNotFound creates a AlertDefinitionPutAlertDefinitionNotFound with default headers values
func NewAlertDefinitionPutAlertDefinitionNotFound() *AlertDefinitionPutAlertDefinitionNotFound {
	return &AlertDefinitionPutAlertDefinitionNotFound{}
}

/*AlertDefinitionPutAlertDefinitionNotFound handles this case with default header values.

The specified alert definition could not be found.
*/
type AlertDefinitionPutAlertDefinitionNotFound struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionPutAlertDefinitionNotFound) Error() string {
	return fmt.Sprintf("[PUT /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionPutAlertDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *AlertDefinitionPutAlertDefinitionNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionPutAlertDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
