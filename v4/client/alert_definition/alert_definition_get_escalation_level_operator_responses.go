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

// AlertDefinitionGetEscalationLevelOperatorReader is a Reader for the AlertDefinitionGetEscalationLevelOperator structure.
type AlertDefinitionGetEscalationLevelOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertDefinitionGetEscalationLevelOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertDefinitionGetEscalationLevelOperatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertDefinitionGetEscalationLevelOperatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAlertDefinitionGetEscalationLevelOperatorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertDefinitionGetEscalationLevelOperatorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertDefinitionGetEscalationLevelOperatorOK creates a AlertDefinitionGetEscalationLevelOperatorOK with default headers values
func NewAlertDefinitionGetEscalationLevelOperatorOK() *AlertDefinitionGetEscalationLevelOperatorOK {
	return &AlertDefinitionGetEscalationLevelOperatorOK{}
}

/*AlertDefinitionGetEscalationLevelOperatorOK handles this case with default header values.

The request completed successfully.
*/
type AlertDefinitionGetEscalationLevelOperatorOK struct {
	Payload []*models.AlertEscalationLevelMember
}

func (o *AlertDefinitionGetEscalationLevelOperatorOK) Error() string {
	return fmt.Sprintf("[GET /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members][%d] alertDefinitionGetEscalationLevelOperatorOK  %+v", 200, o.Payload)
}

func (o *AlertDefinitionGetEscalationLevelOperatorOK) GetPayload() []*models.AlertEscalationLevelMember {
	return o.Payload
}

func (o *AlertDefinitionGetEscalationLevelOperatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionGetEscalationLevelOperatorBadRequest creates a AlertDefinitionGetEscalationLevelOperatorBadRequest with default headers values
func NewAlertDefinitionGetEscalationLevelOperatorBadRequest() *AlertDefinitionGetEscalationLevelOperatorBadRequest {
	return &AlertDefinitionGetEscalationLevelOperatorBadRequest{}
}

/*AlertDefinitionGetEscalationLevelOperatorBadRequest handles this case with default header values.

The request failed.
*/
type AlertDefinitionGetEscalationLevelOperatorBadRequest struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionGetEscalationLevelOperatorBadRequest) Error() string {
	return fmt.Sprintf("[GET /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members][%d] alertDefinitionGetEscalationLevelOperatorBadRequest  %+v", 400, o.Payload)
}

func (o *AlertDefinitionGetEscalationLevelOperatorBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionGetEscalationLevelOperatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionGetEscalationLevelOperatorForbidden creates a AlertDefinitionGetEscalationLevelOperatorForbidden with default headers values
func NewAlertDefinitionGetEscalationLevelOperatorForbidden() *AlertDefinitionGetEscalationLevelOperatorForbidden {
	return &AlertDefinitionGetEscalationLevelOperatorForbidden{}
}

/*AlertDefinitionGetEscalationLevelOperatorForbidden handles this case with default header values.

One or more validation errors occurred.
*/
type AlertDefinitionGetEscalationLevelOperatorForbidden struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionGetEscalationLevelOperatorForbidden) Error() string {
	return fmt.Sprintf("[GET /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members][%d] alertDefinitionGetEscalationLevelOperatorForbidden  %+v", 403, o.Payload)
}

func (o *AlertDefinitionGetEscalationLevelOperatorForbidden) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionGetEscalationLevelOperatorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertDefinitionGetEscalationLevelOperatorNotFound creates a AlertDefinitionGetEscalationLevelOperatorNotFound with default headers values
func NewAlertDefinitionGetEscalationLevelOperatorNotFound() *AlertDefinitionGetEscalationLevelOperatorNotFound {
	return &AlertDefinitionGetEscalationLevelOperatorNotFound{}
}

/*AlertDefinitionGetEscalationLevelOperatorNotFound handles this case with default header values.

One or more of the specified entities could not be found.
*/
type AlertDefinitionGetEscalationLevelOperatorNotFound struct {
	Payload *models.MessageList
}

func (o *AlertDefinitionGetEscalationLevelOperatorNotFound) Error() string {
	return fmt.Sprintf("[GET /AlertDefinition/{alertDefinitionGuid}/EscalationLevel/{escalationLevelId}/Members][%d] alertDefinitionGetEscalationLevelOperatorNotFound  %+v", 404, o.Payload)
}

func (o *AlertDefinitionGetEscalationLevelOperatorNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertDefinitionGetEscalationLevelOperatorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}