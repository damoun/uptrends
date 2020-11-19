// Code generated by go-swagger; DO NOT EDIT.

package monitor_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// MonitorCheckGetConcurrentMonitorPartialChecksReader is a Reader for the MonitorCheckGetConcurrentMonitorPartialChecks structure.
type MonitorCheckGetConcurrentMonitorPartialChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorCheckGetConcurrentMonitorPartialChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMonitorCheckGetConcurrentMonitorPartialChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorCheckGetConcurrentMonitorPartialChecksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMonitorCheckGetConcurrentMonitorPartialChecksOK creates a MonitorCheckGetConcurrentMonitorPartialChecksOK with default headers values
func NewMonitorCheckGetConcurrentMonitorPartialChecksOK() *MonitorCheckGetConcurrentMonitorPartialChecksOK {
	return &MonitorCheckGetConcurrentMonitorPartialChecksOK{}
}

/*MonitorCheckGetConcurrentMonitorPartialChecksOK handles this case with default header values.

The request completed successfully.
*/
type MonitorCheckGetConcurrentMonitorPartialChecksOK struct {
	Payload *models.MonitorCheckResponse
}

func (o *MonitorCheckGetConcurrentMonitorPartialChecksOK) Error() string {
	return fmt.Sprintf("[GET /MonitorCheck/{monitorCheckId}/Concurrent][%d] monitorCheckGetConcurrentMonitorPartialChecksOK  %+v", 200, o.Payload)
}

func (o *MonitorCheckGetConcurrentMonitorPartialChecksOK) GetPayload() *models.MonitorCheckResponse {
	return o.Payload
}

func (o *MonitorCheckGetConcurrentMonitorPartialChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitorCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorCheckGetConcurrentMonitorPartialChecksBadRequest creates a MonitorCheckGetConcurrentMonitorPartialChecksBadRequest with default headers values
func NewMonitorCheckGetConcurrentMonitorPartialChecksBadRequest() *MonitorCheckGetConcurrentMonitorPartialChecksBadRequest {
	return &MonitorCheckGetConcurrentMonitorPartialChecksBadRequest{}
}

/*MonitorCheckGetConcurrentMonitorPartialChecksBadRequest handles this case with default header values.

The request failed.
*/
type MonitorCheckGetConcurrentMonitorPartialChecksBadRequest struct {
	Payload *models.MessageList
}

func (o *MonitorCheckGetConcurrentMonitorPartialChecksBadRequest) Error() string {
	return fmt.Sprintf("[GET /MonitorCheck/{monitorCheckId}/Concurrent][%d] monitorCheckGetConcurrentMonitorPartialChecksBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorCheckGetConcurrentMonitorPartialChecksBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *MonitorCheckGetConcurrentMonitorPartialChecksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
