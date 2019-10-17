// Code generated by go-swagger; DO NOT EDIT.

package monitor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/v4/models"
)

// MonitorGetAllMaintenancePeriodsForMonitorReader is a Reader for the MonitorGetAllMaintenancePeriodsForMonitor structure.
type MonitorGetAllMaintenancePeriodsForMonitorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorGetAllMaintenancePeriodsForMonitorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMonitorGetAllMaintenancePeriodsForMonitorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorGetAllMaintenancePeriodsForMonitorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitorGetAllMaintenancePeriodsForMonitorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMonitorGetAllMaintenancePeriodsForMonitorOK creates a MonitorGetAllMaintenancePeriodsForMonitorOK with default headers values
func NewMonitorGetAllMaintenancePeriodsForMonitorOK() *MonitorGetAllMaintenancePeriodsForMonitorOK {
	return &MonitorGetAllMaintenancePeriodsForMonitorOK{}
}

/*MonitorGetAllMaintenancePeriodsForMonitorOK handles this case with default header values.

Request completed successfully.
*/
type MonitorGetAllMaintenancePeriodsForMonitorOK struct {
	Payload []*models.MaintenancePeriod
}

func (o *MonitorGetAllMaintenancePeriodsForMonitorOK) Error() string {
	return fmt.Sprintf("[GET /Monitor/{monitorGuid}/MaintenancePeriod][%d] monitorGetAllMaintenancePeriodsForMonitorOK  %+v", 200, o.Payload)
}

func (o *MonitorGetAllMaintenancePeriodsForMonitorOK) GetPayload() []*models.MaintenancePeriod {
	return o.Payload
}

func (o *MonitorGetAllMaintenancePeriodsForMonitorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorGetAllMaintenancePeriodsForMonitorBadRequest creates a MonitorGetAllMaintenancePeriodsForMonitorBadRequest with default headers values
func NewMonitorGetAllMaintenancePeriodsForMonitorBadRequest() *MonitorGetAllMaintenancePeriodsForMonitorBadRequest {
	return &MonitorGetAllMaintenancePeriodsForMonitorBadRequest{}
}

/*MonitorGetAllMaintenancePeriodsForMonitorBadRequest handles this case with default header values.

The request failed.
*/
type MonitorGetAllMaintenancePeriodsForMonitorBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorGetAllMaintenancePeriodsForMonitorBadRequest) Error() string {
	return fmt.Sprintf("[GET /Monitor/{monitorGuid}/MaintenancePeriod][%d] monitorGetAllMaintenancePeriodsForMonitorBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorGetAllMaintenancePeriodsForMonitorBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorGetAllMaintenancePeriodsForMonitorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorGetAllMaintenancePeriodsForMonitorNotFound creates a MonitorGetAllMaintenancePeriodsForMonitorNotFound with default headers values
func NewMonitorGetAllMaintenancePeriodsForMonitorNotFound() *MonitorGetAllMaintenancePeriodsForMonitorNotFound {
	return &MonitorGetAllMaintenancePeriodsForMonitorNotFound{}
}

/*MonitorGetAllMaintenancePeriodsForMonitorNotFound handles this case with default header values.

The specified monitor does not exist.
*/
type MonitorGetAllMaintenancePeriodsForMonitorNotFound struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorGetAllMaintenancePeriodsForMonitorNotFound) Error() string {
	return fmt.Sprintf("[GET /Monitor/{monitorGuid}/MaintenancePeriod][%d] monitorGetAllMaintenancePeriodsForMonitorNotFound  %+v", 404, o.Payload)
}

func (o *MonitorGetAllMaintenancePeriodsForMonitorNotFound) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorGetAllMaintenancePeriodsForMonitorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}