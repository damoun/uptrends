// Code generated by go-swagger; DO NOT EDIT.

package alert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// AlertGetMonitorAlertsReader is a Reader for the AlertGetMonitorAlerts structure.
type AlertGetMonitorAlertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AlertGetMonitorAlertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAlertGetMonitorAlertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAlertGetMonitorAlertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAlertGetMonitorAlertsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAlertGetMonitorAlertsOK creates a AlertGetMonitorAlertsOK with default headers values
func NewAlertGetMonitorAlertsOK() *AlertGetMonitorAlertsOK {
	return &AlertGetMonitorAlertsOK{}
}

/*AlertGetMonitorAlertsOK handles this case with default header values.

The request completed successfully.
*/
type AlertGetMonitorAlertsOK struct {
	Payload *models.AlertResponse
}

func (o *AlertGetMonitorAlertsOK) Error() string {
	return fmt.Sprintf("[GET /Alert/Monitor/{monitorGuid}][%d] alertGetMonitorAlertsOK  %+v", 200, o.Payload)
}

func (o *AlertGetMonitorAlertsOK) GetPayload() *models.AlertResponse {
	return o.Payload
}

func (o *AlertGetMonitorAlertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertGetMonitorAlertsBadRequest creates a AlertGetMonitorAlertsBadRequest with default headers values
func NewAlertGetMonitorAlertsBadRequest() *AlertGetMonitorAlertsBadRequest {
	return &AlertGetMonitorAlertsBadRequest{}
}

/*AlertGetMonitorAlertsBadRequest handles this case with default header values.

The request failed.
*/
type AlertGetMonitorAlertsBadRequest struct {
	Payload *models.MessageList
}

func (o *AlertGetMonitorAlertsBadRequest) Error() string {
	return fmt.Sprintf("[GET /Alert/Monitor/{monitorGuid}][%d] alertGetMonitorAlertsBadRequest  %+v", 400, o.Payload)
}

func (o *AlertGetMonitorAlertsBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertGetMonitorAlertsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAlertGetMonitorAlertsNotFound creates a AlertGetMonitorAlertsNotFound with default headers values
func NewAlertGetMonitorAlertsNotFound() *AlertGetMonitorAlertsNotFound {
	return &AlertGetMonitorAlertsNotFound{}
}

/*AlertGetMonitorAlertsNotFound handles this case with default header values.

The specified Monitor for monitorGuid does not exist.
*/
type AlertGetMonitorAlertsNotFound struct {
	Payload *models.MessageList
}

func (o *AlertGetMonitorAlertsNotFound) Error() string {
	return fmt.Sprintf("[GET /Alert/Monitor/{monitorGuid}][%d] alertGetMonitorAlertsNotFound  %+v", 404, o.Payload)
}

func (o *AlertGetMonitorAlertsNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *AlertGetMonitorAlertsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
