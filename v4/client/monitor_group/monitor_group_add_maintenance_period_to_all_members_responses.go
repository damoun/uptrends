// Code generated by go-swagger; DO NOT EDIT.

package monitor_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// MonitorGroupAddMaintenancePeriodToAllMembersReader is a Reader for the MonitorGroupAddMaintenancePeriodToAllMembers structure.
type MonitorGroupAddMaintenancePeriodToAllMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorGroupAddMaintenancePeriodToAllMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewMonitorGroupAddMaintenancePeriodToAllMembersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorGroupAddMaintenancePeriodToAllMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMonitorGroupAddMaintenancePeriodToAllMembersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMonitorGroupAddMaintenancePeriodToAllMembersNoContent creates a MonitorGroupAddMaintenancePeriodToAllMembersNoContent with default headers values
func NewMonitorGroupAddMaintenancePeriodToAllMembersNoContent() *MonitorGroupAddMaintenancePeriodToAllMembersNoContent {
	return &MonitorGroupAddMaintenancePeriodToAllMembersNoContent{}
}

/*MonitorGroupAddMaintenancePeriodToAllMembersNoContent handles this case with default header values.

The request completed successfully. No content is returned.
*/
type MonitorGroupAddMaintenancePeriodToAllMembersNoContent struct {
}

func (o *MonitorGroupAddMaintenancePeriodToAllMembersNoContent) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/AddMaintenancePeriodToAllMembers][%d] monitorGroupAddMaintenancePeriodToAllMembersNoContent ", 204)
}

func (o *MonitorGroupAddMaintenancePeriodToAllMembersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewMonitorGroupAddMaintenancePeriodToAllMembersBadRequest creates a MonitorGroupAddMaintenancePeriodToAllMembersBadRequest with default headers values
func NewMonitorGroupAddMaintenancePeriodToAllMembersBadRequest() *MonitorGroupAddMaintenancePeriodToAllMembersBadRequest {
	return &MonitorGroupAddMaintenancePeriodToAllMembersBadRequest{}
}

/*MonitorGroupAddMaintenancePeriodToAllMembersBadRequest handles this case with default header values.

The request failed.
*/
type MonitorGroupAddMaintenancePeriodToAllMembersBadRequest struct {
	Payload *models.MessageList
}

func (o *MonitorGroupAddMaintenancePeriodToAllMembersBadRequest) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/AddMaintenancePeriodToAllMembers][%d] monitorGroupAddMaintenancePeriodToAllMembersBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorGroupAddMaintenancePeriodToAllMembersBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *MonitorGroupAddMaintenancePeriodToAllMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorGroupAddMaintenancePeriodToAllMembersNotFound creates a MonitorGroupAddMaintenancePeriodToAllMembersNotFound with default headers values
func NewMonitorGroupAddMaintenancePeriodToAllMembersNotFound() *MonitorGroupAddMaintenancePeriodToAllMembersNotFound {
	return &MonitorGroupAddMaintenancePeriodToAllMembersNotFound{}
}

/*MonitorGroupAddMaintenancePeriodToAllMembersNotFound handles this case with default header values.

The requested monitor group was not found.
*/
type MonitorGroupAddMaintenancePeriodToAllMembersNotFound struct {
	Payload *models.MessageList
}

func (o *MonitorGroupAddMaintenancePeriodToAllMembersNotFound) Error() string {
	return fmt.Sprintf("[POST /MonitorGroup/{monitorGroupGuid}/AddMaintenancePeriodToAllMembers][%d] monitorGroupAddMaintenancePeriodToAllMembersNotFound  %+v", 404, o.Payload)
}

func (o *MonitorGroupAddMaintenancePeriodToAllMembersNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *MonitorGroupAddMaintenancePeriodToAllMembersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
