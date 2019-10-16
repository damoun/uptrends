// Code generated by go-swagger; DO NOT EDIT.

package monitor_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/models"
)

// MonitorCheckGetAccountMonitorChecksReader is a Reader for the MonitorCheckGetAccountMonitorChecks structure.
type MonitorCheckGetAccountMonitorChecksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorCheckGetAccountMonitorChecksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMonitorCheckGetAccountMonitorChecksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorCheckGetAccountMonitorChecksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMonitorCheckGetAccountMonitorChecksOK creates a MonitorCheckGetAccountMonitorChecksOK with default headers values
func NewMonitorCheckGetAccountMonitorChecksOK() *MonitorCheckGetAccountMonitorChecksOK {
	return &MonitorCheckGetAccountMonitorChecksOK{}
}

/*MonitorCheckGetAccountMonitorChecksOK handles this case with default header values.

Monitor checks for the entire account.
*/
type MonitorCheckGetAccountMonitorChecksOK struct {
	Payload *models.MonitorCheckResponse
}

func (o *MonitorCheckGetAccountMonitorChecksOK) Error() string {
	return fmt.Sprintf("[GET /MonitorCheck][%d] monitorCheckGetAccountMonitorChecksOK  %+v", 200, o.Payload)
}

func (o *MonitorCheckGetAccountMonitorChecksOK) GetPayload() *models.MonitorCheckResponse {
	return o.Payload
}

func (o *MonitorCheckGetAccountMonitorChecksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MonitorCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorCheckGetAccountMonitorChecksBadRequest creates a MonitorCheckGetAccountMonitorChecksBadRequest with default headers values
func NewMonitorCheckGetAccountMonitorChecksBadRequest() *MonitorCheckGetAccountMonitorChecksBadRequest {
	return &MonitorCheckGetAccountMonitorChecksBadRequest{}
}

/*MonitorCheckGetAccountMonitorChecksBadRequest handles this case with default header values.

The request failed.
*/
type MonitorCheckGetAccountMonitorChecksBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorCheckGetAccountMonitorChecksBadRequest) Error() string {
	return fmt.Sprintf("[GET /MonitorCheck][%d] monitorCheckGetAccountMonitorChecksBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorCheckGetAccountMonitorChecksBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorCheckGetAccountMonitorChecksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
