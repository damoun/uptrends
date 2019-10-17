// Code generated by go-swagger; DO NOT EDIT.

package monitor_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/v4/models"
)

// MonitorGroupUpdateMonitorGroupReader is a Reader for the MonitorGroupUpdateMonitorGroup structure.
type MonitorGroupUpdateMonitorGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorGroupUpdateMonitorGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewMonitorGroupUpdateMonitorGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorGroupUpdateMonitorGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitorGroupUpdateMonitorGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMonitorGroupUpdateMonitorGroupNoContent creates a MonitorGroupUpdateMonitorGroupNoContent with default headers values
func NewMonitorGroupUpdateMonitorGroupNoContent() *MonitorGroupUpdateMonitorGroupNoContent {
	return &MonitorGroupUpdateMonitorGroupNoContent{}
}

/*MonitorGroupUpdateMonitorGroupNoContent handles this case with default header values.

Request completed successfully.
*/
type MonitorGroupUpdateMonitorGroupNoContent struct {
}

func (o *MonitorGroupUpdateMonitorGroupNoContent) Error() string {
	return fmt.Sprintf("[PUT /MonitorGroup/{monitorGroupGuid}][%d] monitorGroupUpdateMonitorGroupNoContent ", 204)
}

func (o *MonitorGroupUpdateMonitorGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitorGroupUpdateMonitorGroupBadRequest creates a MonitorGroupUpdateMonitorGroupBadRequest with default headers values
func NewMonitorGroupUpdateMonitorGroupBadRequest() *MonitorGroupUpdateMonitorGroupBadRequest {
	return &MonitorGroupUpdateMonitorGroupBadRequest{}
}

/*MonitorGroupUpdateMonitorGroupBadRequest handles this case with default header values.

The request failed.
or
The All Monitors group cannot be changed or deleted.
or
The All Monitors group already exists.
*/
type MonitorGroupUpdateMonitorGroupBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorGroupUpdateMonitorGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /MonitorGroup/{monitorGroupGuid}][%d] monitorGroupUpdateMonitorGroupBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorGroupUpdateMonitorGroupBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorGroupUpdateMonitorGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorGroupUpdateMonitorGroupNotFound creates a MonitorGroupUpdateMonitorGroupNotFound with default headers values
func NewMonitorGroupUpdateMonitorGroupNotFound() *MonitorGroupUpdateMonitorGroupNotFound {
	return &MonitorGroupUpdateMonitorGroupNotFound{}
}

/*MonitorGroupUpdateMonitorGroupNotFound handles this case with default header values.

The requested monitor group was not found.
*/
type MonitorGroupUpdateMonitorGroupNotFound struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorGroupUpdateMonitorGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /MonitorGroup/{monitorGroupGuid}][%d] monitorGroupUpdateMonitorGroupNotFound  %+v", 404, o.Payload)
}

func (o *MonitorGroupUpdateMonitorGroupNotFound) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorGroupUpdateMonitorGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}