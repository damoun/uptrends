// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// MonitorUpdateMaintenancePeriodForMonitorReader is a Reader for the MonitorUpdateMaintenancePeriodForMonitor structure.
type MonitorUpdateMaintenancePeriodForMonitorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorUpdateMaintenancePeriodForMonitorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewMonitorUpdateMaintenancePeriodForMonitorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorUpdateMaintenancePeriodForMonitorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitorUpdateMaintenancePeriodForMonitorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMonitorUpdateMaintenancePeriodForMonitorNoContent creates a MonitorUpdateMaintenancePeriodForMonitorNoContent with default headers values
func NewMonitorUpdateMaintenancePeriodForMonitorNoContent() *MonitorUpdateMaintenancePeriodForMonitorNoContent {
	return &MonitorUpdateMaintenancePeriodForMonitorNoContent{}
}

/*MonitorUpdateMaintenancePeriodForMonitorNoContent handles this case with default header values.

The request completed successfully. No content is returned.
*/
type MonitorUpdateMaintenancePeriodForMonitorNoContent struct {
}

func (o *MonitorUpdateMaintenancePeriodForMonitorNoContent) Error() string {
	return fmt.Sprintf("[PUT /Monitor/{monitorGuid}/MaintenancePeriod/{maintenancePeriodId}][%d] monitorUpdateMaintenancePeriodForMonitorNoContent ", 204)
}

func (o *MonitorUpdateMaintenancePeriodForMonitorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitorUpdateMaintenancePeriodForMonitorBadRequest creates a MonitorUpdateMaintenancePeriodForMonitorBadRequest with default headers values
func NewMonitorUpdateMaintenancePeriodForMonitorBadRequest() *MonitorUpdateMaintenancePeriodForMonitorBadRequest {
	return &MonitorUpdateMaintenancePeriodForMonitorBadRequest{}
}

/*MonitorUpdateMaintenancePeriodForMonitorBadRequest handles this case with default header values.

The request failed.
*/
type MonitorUpdateMaintenancePeriodForMonitorBadRequest struct {
	Payload *models.MessageList
}

func (o *MonitorUpdateMaintenancePeriodForMonitorBadRequest) Error() string {
	return fmt.Sprintf("[PUT /Monitor/{monitorGuid}/MaintenancePeriod/{maintenancePeriodId}][%d] monitorUpdateMaintenancePeriodForMonitorBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorUpdateMaintenancePeriodForMonitorBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *MonitorUpdateMaintenancePeriodForMonitorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorUpdateMaintenancePeriodForMonitorNotFound creates a MonitorUpdateMaintenancePeriodForMonitorNotFound with default headers values
func NewMonitorUpdateMaintenancePeriodForMonitorNotFound() *MonitorUpdateMaintenancePeriodForMonitorNotFound {
	return &MonitorUpdateMaintenancePeriodForMonitorNotFound{}
}

/*MonitorUpdateMaintenancePeriodForMonitorNotFound handles this case with default header values.

The specified monitor or maintenanceperiod does not exist.
*/
type MonitorUpdateMaintenancePeriodForMonitorNotFound struct {
	Payload *models.MessageList
}

func (o *MonitorUpdateMaintenancePeriodForMonitorNotFound) Error() string {
	return fmt.Sprintf("[PUT /Monitor/{monitorGuid}/MaintenancePeriod/{maintenancePeriodId}][%d] monitorUpdateMaintenancePeriodForMonitorNotFound  %+v", 404, o.Payload)
}

func (o *MonitorUpdateMaintenancePeriodForMonitorNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *MonitorUpdateMaintenancePeriodForMonitorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
