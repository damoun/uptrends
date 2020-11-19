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

// AlertDefinitionGetSpecifiedAlertDefinitionsReader is a Reader for the AlertDefinitionGetSpecifiedAlertDefinitions structure.
type AlertDefinitionGetSpecifiedAlertDefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertDefinitionGetSpecifiedAlertDefinitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertDefinitionGetSpecifiedAlertDefinitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertDefinitionGetSpecifiedAlertDefinitionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertDefinitionGetSpecifiedAlertDefinitionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertDefinitionGetSpecifiedAlertDefinitionsOK creates a AlertDefinitionGetSpecifiedAlertDefinitionsOK with default headers values
func NewAlertDefinitionGetSpecifiedAlertDefinitionsOK() *AlertDefinitionGetSpecifiedAlertDefinitionsOK {
	return &AlertDefinitionGetSpecifiedAlertDefinitionsOK{}
}

/*AlertDefinitionGetSpecifiedAlertDefinitionsOK handles this case with default header values.

The request completed successfully.
*/
type AlertDefinitionGetSpecifiedAlertDefinitionsOK struct {
	Payload *models.AlertDefinition
}

func (o *AlertDefinitionGetSpecifiedAlertDefinitionsOK) Error() string {
	return fmt.Sprintf("[GET /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionGetSpecifiedAlertDefinitionsOK  %+v", 200, o.Payload)
}

func (o *AlertDefinitionGetSpecifiedAlertDefinitionsOK) GetPayload() *models.AlertDefinition {
	return o.Payload
}

func (o *AlertDefinitionGetSpecifiedAlertDefinitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionGetSpecifiedAlertDefinitionsBadRequest creates a AlertDefinitionGetSpecifiedAlertDefinitionsBadRequest with default headers values
func NewAlertDefinitionGetSpecifiedAlertDefinitionsBadRequest() *AlertDefinitionGetSpecifiedAlertDefinitionsBadRequest {
	return &AlertDefinitionGetSpecifiedAlertDefinitionsBadRequest{}
}

/*AlertDefinitionGetSpecifiedAlertDefinitionsBadRequest handles this case with default header values.

The request failed.
*/
type AlertDefinitionGetSpecifiedAlertDefinitionsBadRequest struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionGetSpecifiedAlertDefinitionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionGetSpecifiedAlertDefinitionsBadRequest  %+v", 400, o.Payload)
}

func (o *AlertDefinitionGetSpecifiedAlertDefinitionsBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionGetSpecifiedAlertDefinitionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionGetSpecifiedAlertDefinitionsForbidden creates a AlertDefinitionGetSpecifiedAlertDefinitionsForbidden with default headers values
func NewAlertDefinitionGetSpecifiedAlertDefinitionsForbidden() *AlertDefinitionGetSpecifiedAlertDefinitionsForbidden {
	return &AlertDefinitionGetSpecifiedAlertDefinitionsForbidden{}
}

/*AlertDefinitionGetSpecifiedAlertDefinitionsForbidden handles this case with default header values.

One or more validation errors occurred.
*/
type AlertDefinitionGetSpecifiedAlertDefinitionsForbidden struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionGetSpecifiedAlertDefinitionsForbidden) Error() string {
	return fmt.Sprintf("[GET /AlertDefinition/{alertDefinitionGuid}][%d] alertDefinitionGetSpecifiedAlertDefinitionsForbidden  %+v", 403, o.Payload)
}

func (o *AlertDefinitionGetSpecifiedAlertDefinitionsForbidden) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionGetSpecifiedAlertDefinitionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}