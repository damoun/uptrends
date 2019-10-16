// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/models"
)

// MonitorDeleteMonitorReader is a Reader for the MonitorDeleteMonitor structure.
type MonitorDeleteMonitorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorDeleteMonitorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewMonitorDeleteMonitorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorDeleteMonitorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitorDeleteMonitorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMonitorDeleteMonitorNoContent creates a MonitorDeleteMonitorNoContent with default headers values
func NewMonitorDeleteMonitorNoContent() *MonitorDeleteMonitorNoContent {
	return &MonitorDeleteMonitorNoContent{}
}

/*MonitorDeleteMonitorNoContent handles this case with default header values.

The monitor has been removed. No content is returned.
*/
type MonitorDeleteMonitorNoContent struct {
}

func (o *MonitorDeleteMonitorNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Monitor/{monitorGuid}][%d] monitorDeleteMonitorNoContent ", 204)
}

func (o *MonitorDeleteMonitorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitorDeleteMonitorBadRequest creates a MonitorDeleteMonitorBadRequest with default headers values
func NewMonitorDeleteMonitorBadRequest() *MonitorDeleteMonitorBadRequest {
	return &MonitorDeleteMonitorBadRequest{}
}

/*MonitorDeleteMonitorBadRequest handles this case with default header values.

The request failed.
*/
type MonitorDeleteMonitorBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorDeleteMonitorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /Monitor/{monitorGuid}][%d] monitorDeleteMonitorBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorDeleteMonitorBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorDeleteMonitorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorDeleteMonitorNotFound creates a MonitorDeleteMonitorNotFound with default headers values
func NewMonitorDeleteMonitorNotFound() *MonitorDeleteMonitorNotFound {
	return &MonitorDeleteMonitorNotFound{}
}

/*MonitorDeleteMonitorNotFound handles this case with default header values.

The specified monitor does not exist.
*/
type MonitorDeleteMonitorNotFound struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorDeleteMonitorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /Monitor/{monitorGuid}][%d] monitorDeleteMonitorNotFound  %+v", 404, o.Payload)
}

func (o *MonitorDeleteMonitorNotFound) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorDeleteMonitorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
