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

// MonitorGroupAddMonitorToMonitorGroupReader is a Reader for the MonitorGroupAddMonitorToMonitorGroup structure.
type MonitorGroupAddMonitorToMonitorGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorGroupAddMonitorToMonitorGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewMonitorGroupAddMonitorToMonitorGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorGroupAddMonitorToMonitorGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitorGroupAddMonitorToMonitorGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMonitorGroupAddMonitorToMonitorGroupCreated creates a MonitorGroupAddMonitorToMonitorGroupCreated with default headers values
func NewMonitorGroupAddMonitorToMonitorGroupCreated() *MonitorGroupAddMonitorToMonitorGroupCreated {
	return &MonitorGroupAddMonitorToMonitorGroupCreated{}
}

/*MonitorGroupAddMonitorToMonitorGroupCreated handles this case with default header values.

The monitor has been added to the group successfully.
*/
type MonitorGroupAddMonitorToMonitorGroupCreated struct {
}

func (o *MonitorGroupAddMonitorToMonitorGroupCreated) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/Members/{monitorGuid}][%d] monitorGroupAddMonitorToMonitorGroupCreated ", 201)
}

func (o *MonitorGroupAddMonitorToMonitorGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitorGroupAddMonitorToMonitorGroupBadRequest creates a MonitorGroupAddMonitorToMonitorGroupBadRequest with default headers values
func NewMonitorGroupAddMonitorToMonitorGroupBadRequest() *MonitorGroupAddMonitorToMonitorGroupBadRequest {
	return &MonitorGroupAddMonitorToMonitorGroupBadRequest{}
}

/*MonitorGroupAddMonitorToMonitorGroupBadRequest handles this case with default header values.

The request failed.
*/
type MonitorGroupAddMonitorToMonitorGroupBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorGroupAddMonitorToMonitorGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/Members/{monitorGuid}][%d] monitorGroupAddMonitorToMonitorGroupBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorGroupAddMonitorToMonitorGroupBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorGroupAddMonitorToMonitorGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorGroupAddMonitorToMonitorGroupNotFound creates a MonitorGroupAddMonitorToMonitorGroupNotFound with default headers values
func NewMonitorGroupAddMonitorToMonitorGroupNotFound() *MonitorGroupAddMonitorToMonitorGroupNotFound {
	return &MonitorGroupAddMonitorToMonitorGroupNotFound{}
}

/*MonitorGroupAddMonitorToMonitorGroupNotFound handles this case with default header values.

The requested monitor group was not found.
*/
type MonitorGroupAddMonitorToMonitorGroupNotFound struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorGroupAddMonitorToMonitorGroupNotFound) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/Members/{monitorGuid}][%d] monitorGroupAddMonitorToMonitorGroupNotFound  %+v", 404, o.Payload)
}

func (o *MonitorGroupAddMonitorToMonitorGroupNotFound) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorGroupAddMonitorToMonitorGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
