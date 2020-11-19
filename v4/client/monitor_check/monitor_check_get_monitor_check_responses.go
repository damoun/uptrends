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

// MonitorCheckGetMonitorCheckReader is a Reader for the MonitorCheckGetMonitorCheck structure.
type MonitorCheckGetMonitorCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorCheckGetMonitorCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMonitorCheckGetMonitorCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorCheckGetMonitorCheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMonitorCheckGetMonitorCheckOK creates a MonitorCheckGetMonitorCheckOK with default headers values
func NewMonitorCheckGetMonitorCheckOK() *MonitorCheckGetMonitorCheckOK {
	return &MonitorCheckGetMonitorCheckOK{}
}

/*MonitorCheckGetMonitorCheckOK handles this case with default header values.

The request completed successfully.
*/
type MonitorCheckGetMonitorCheckOK struct {
	Payload *models.MonitorCheckResponse
}

func (o *MonitorCheckGetMonitorCheckOK) Error() string {
	return fmt.Sprintf("[GET /MonitorCheck/Monitor/{monitorGuid}][%d] monitorCheckGetMonitorCheckOK  %+v", 200, o.Payload)
}

func (o *MonitorCheckGetMonitorCheckOK) GetPayload() *models.MonitorCheckResponse {
	return o.Payload
}

func (o *MonitorCheckGetMonitorCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitorCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorCheckGetMonitorCheckBadRequest creates a MonitorCheckGetMonitorCheckBadRequest with default headers values
func NewMonitorCheckGetMonitorCheckBadRequest() *MonitorCheckGetMonitorCheckBadRequest {
	return &MonitorCheckGetMonitorCheckBadRequest{}
}

/*MonitorCheckGetMonitorCheckBadRequest handles this case with default header values.

The request failed.
*/
type MonitorCheckGetMonitorCheckBadRequest struct {
	Payload *models.MessageList
}

func (o *MonitorCheckGetMonitorCheckBadRequest) Error() string {
	return fmt.Sprintf("[GET /MonitorCheck/Monitor/{monitorGuid}][%d] monitorCheckGetMonitorCheckBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorCheckGetMonitorCheckBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *MonitorCheckGetMonitorCheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
