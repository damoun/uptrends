// Code generated by go-swagger; DO NOT EDIT.

package monitor_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/models"
)

// MonitorGroupStartAllMonitorsInGroupReader is a Reader for the MonitorGroupStartAllMonitorsInGroup structure.
type MonitorGroupStartAllMonitorsInGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorGroupStartAllMonitorsInGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewMonitorGroupStartAllMonitorsInGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorGroupStartAllMonitorsInGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitorGroupStartAllMonitorsInGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMonitorGroupStartAllMonitorsInGroupNoContent creates a MonitorGroupStartAllMonitorsInGroupNoContent with default headers values
func NewMonitorGroupStartAllMonitorsInGroupNoContent() *MonitorGroupStartAllMonitorsInGroupNoContent {
	return &MonitorGroupStartAllMonitorsInGroupNoContent{}
}

/*MonitorGroupStartAllMonitorsInGroupNoContent handles this case with default header values.

All monitors in the group have been started.
*/
type MonitorGroupStartAllMonitorsInGroupNoContent struct {
}

func (o *MonitorGroupStartAllMonitorsInGroupNoContent) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/StartAllMonitors][%d] monitorGroupStartAllMonitorsInGroupNoContent ", 204)
}

func (o *MonitorGroupStartAllMonitorsInGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitorGroupStartAllMonitorsInGroupBadRequest creates a MonitorGroupStartAllMonitorsInGroupBadRequest with default headers values
func NewMonitorGroupStartAllMonitorsInGroupBadRequest() *MonitorGroupStartAllMonitorsInGroupBadRequest {
	return &MonitorGroupStartAllMonitorsInGroupBadRequest{}
}

/*MonitorGroupStartAllMonitorsInGroupBadRequest handles this case with default header values.

The request failed.
*/
type MonitorGroupStartAllMonitorsInGroupBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorGroupStartAllMonitorsInGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/StartAllMonitors][%d] monitorGroupStartAllMonitorsInGroupBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorGroupStartAllMonitorsInGroupBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorGroupStartAllMonitorsInGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorGroupStartAllMonitorsInGroupNotFound creates a MonitorGroupStartAllMonitorsInGroupNotFound with default headers values
func NewMonitorGroupStartAllMonitorsInGroupNotFound() *MonitorGroupStartAllMonitorsInGroupNotFound {
	return &MonitorGroupStartAllMonitorsInGroupNotFound{}
}

/*MonitorGroupStartAllMonitorsInGroupNotFound handles this case with default header values.

The requested monitor group was not found.
*/
type MonitorGroupStartAllMonitorsInGroupNotFound struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorGroupStartAllMonitorsInGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/StartAllMonitors][%d] monitorGroupStartAllMonitorsInGroupNotFound  %+v", 404, o.Payload)
}

func (o *MonitorGroupStartAllMonitorsInGroupNotFound) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorGroupStartAllMonitorsInGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
