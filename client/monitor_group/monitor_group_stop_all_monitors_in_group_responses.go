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

// MonitorGroupStopAllMonitorsInGroupReader is a Reader for the MonitorGroupStopAllMonitorsInGroup structure.
type MonitorGroupStopAllMonitorsInGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorGroupStopAllMonitorsInGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewMonitorGroupStopAllMonitorsInGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorGroupStopAllMonitorsInGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitorGroupStopAllMonitorsInGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMonitorGroupStopAllMonitorsInGroupNoContent creates a MonitorGroupStopAllMonitorsInGroupNoContent with default headers values
func NewMonitorGroupStopAllMonitorsInGroupNoContent() *MonitorGroupStopAllMonitorsInGroupNoContent {
	return &MonitorGroupStopAllMonitorsInGroupNoContent{}
}

/*MonitorGroupStopAllMonitorsInGroupNoContent handles this case with default header values.

All monitors in the group have been stopped.
*/
type MonitorGroupStopAllMonitorsInGroupNoContent struct {
}

func (o *MonitorGroupStopAllMonitorsInGroupNoContent) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/StopAllMonitors][%d] monitorGroupStopAllMonitorsInGroupNoContent ", 204)
}

func (o *MonitorGroupStopAllMonitorsInGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitorGroupStopAllMonitorsInGroupBadRequest creates a MonitorGroupStopAllMonitorsInGroupBadRequest with default headers values
func NewMonitorGroupStopAllMonitorsInGroupBadRequest() *MonitorGroupStopAllMonitorsInGroupBadRequest {
	return &MonitorGroupStopAllMonitorsInGroupBadRequest{}
}

/*MonitorGroupStopAllMonitorsInGroupBadRequest handles this case with default header values.

The request failed.
*/
type MonitorGroupStopAllMonitorsInGroupBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorGroupStopAllMonitorsInGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/StopAllMonitors][%d] monitorGroupStopAllMonitorsInGroupBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorGroupStopAllMonitorsInGroupBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorGroupStopAllMonitorsInGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorGroupStopAllMonitorsInGroupNotFound creates a MonitorGroupStopAllMonitorsInGroupNotFound with default headers values
func NewMonitorGroupStopAllMonitorsInGroupNotFound() *MonitorGroupStopAllMonitorsInGroupNotFound {
	return &MonitorGroupStopAllMonitorsInGroupNotFound{}
}

/*MonitorGroupStopAllMonitorsInGroupNotFound handles this case with default header values.

The requested monitor group was not found.
*/
type MonitorGroupStopAllMonitorsInGroupNotFound struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorGroupStopAllMonitorsInGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/StopAllMonitors][%d] monitorGroupStopAllMonitorsInGroupNotFound  %+v", 404, o.Payload)
}

func (o *MonitorGroupStopAllMonitorsInGroupNotFound) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorGroupStopAllMonitorsInGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}