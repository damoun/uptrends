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

// AlertDefinitionAddOperatorGroupToEscalationLevelReader is a Reader for the AlertDefinitionAddOperatorGroupToEscalationLevel structure.
type AlertDefinitionAddOperatorGroupToEscalationLevelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertDefinitionAddOperatorGroupToEscalationLevelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAlertDefinitionAddOperatorGroupToEscalationLevelCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertDefinitionAddOperatorGroupToEscalationLevelBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertDefinitionAddOperatorGroupToEscalationLevelForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertDefinitionAddOperatorGroupToEscalationLevelNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertDefinitionAddOperatorGroupToEscalationLevelCreated creates a AlertDefinitionAddOperatorGroupToEscalationLevelCreated with default headers values
func NewAlertDefinitionAddOperatorGroupToEscalationLevelCreated() *AlertDefinitionAddOperatorGroupToEscalationLevelCreated {
	return &AlertDefinitionAddOperatorGroupToEscalationLevelCreated{}
}

/*AlertDefinitionAddOperatorGroupToEscalationLevelCreated handles this case with default header values.

The request completed successfully.
*/
type AlertDefinitionAddOperatorGroupToEscalationLevelCreated struct {
	Payload *models.AlertDefinitionOperatorGroup
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelCreated) Error() string {
	return fmt.Sprintf("[POST /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members/OperatorGroup/{operatorGroupGuid}][%d] alertDefinitionAddOperatorGroupToEscalationLevelCreated  %+v", 201, o.Payload)
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelCreated) GetPayload() *models.AlertDefinitionOperatorGroup {
	return o.Payload
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertDefinitionOperatorGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionAddOperatorGroupToEscalationLevelBadRequest creates a AlertDefinitionAddOperatorGroupToEscalationLevelBadRequest with default headers values
func NewAlertDefinitionAddOperatorGroupToEscalationLevelBadRequest() *AlertDefinitionAddOperatorGroupToEscalationLevelBadRequest {
	return &AlertDefinitionAddOperatorGroupToEscalationLevelBadRequest{}
}

/*AlertDefinitionAddOperatorGroupToEscalationLevelBadRequest handles this case with default header values.

The request failed.
*/
type AlertDefinitionAddOperatorGroupToEscalationLevelBadRequest struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelBadRequest) Error() string {
	return fmt.Sprintf("[POST /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members/OperatorGroup/{operatorGroupGuid}][%d] alertDefinitionAddOperatorGroupToEscalationLevelBadRequest  %+v", 400, o.Payload)
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionAddOperatorGroupToEscalationLevelForbidden creates a AlertDefinitionAddOperatorGroupToEscalationLevelForbidden with default headers values
func NewAlertDefinitionAddOperatorGroupToEscalationLevelForbidden() *AlertDefinitionAddOperatorGroupToEscalationLevelForbidden {
	return &AlertDefinitionAddOperatorGroupToEscalationLevelForbidden{}
}

/*AlertDefinitionAddOperatorGroupToEscalationLevelForbidden handles this case with default header values.

One or more validation errors occurred.
*/
type AlertDefinitionAddOperatorGroupToEscalationLevelForbidden struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelForbidden) Error() string {
	return fmt.Sprintf("[POST /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members/OperatorGroup/{operatorGroupGuid}][%d] alertDefinitionAddOperatorGroupToEscalationLevelForbidden  %+v", 403, o.Payload)
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelForbidden) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionAddOperatorGroupToEscalationLevelNotFound creates a AlertDefinitionAddOperatorGroupToEscalationLevelNotFound with default headers values
func NewAlertDefinitionAddOperatorGroupToEscalationLevelNotFound() *AlertDefinitionAddOperatorGroupToEscalationLevelNotFound {
	return &AlertDefinitionAddOperatorGroupToEscalationLevelNotFound{}
}

/*AlertDefinitionAddOperatorGroupToEscalationLevelNotFound handles this case with default header values.

One or more of the specified entities could not be found.
*/
type AlertDefinitionAddOperatorGroupToEscalationLevelNotFound struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelNotFound) Error() string {
	return fmt.Sprintf("[POST /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members/OperatorGroup/{operatorGroupGuid}][%d] alertDefinitionAddOperatorGroupToEscalationLevelNotFound  %+v", 404, o.Payload)
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionAddOperatorGroupToEscalationLevelNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}