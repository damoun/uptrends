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

// MonitorCheckGetTransactionDetailsReader is a Reader for the MonitorCheckGetTransactionDetails structure.
type MonitorCheckGetTransactionDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MonitorCheckGetTransactionDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMonitorCheckGetTransactionDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMonitorCheckGetTransactionDetailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMonitorCheckGetTransactionDetailsOK creates a MonitorCheckGetTransactionDetailsOK with default headers values
func NewMonitorCheckGetTransactionDetailsOK() *MonitorCheckGetTransactionDetailsOK {
	return &MonitorCheckGetTransactionDetailsOK{}
}

/*MonitorCheckGetTransactionDetailsOK handles this case with default header values.

MonitorCheckGetTransactionDetailsOK monitor check get transaction details o k
*/
type MonitorCheckGetTransactionDetailsOK struct {
	Payload *models.TransactionDetailsResponse
}

func (o *MonitorCheckGetTransactionDetailsOK) Error() string {
	return fmt.Sprintf("[GET /MonitorCheck/{monitorCheckId}/Transaction][%d] monitorCheckGetTransactionDetailsOK  %+v", 200, o.Payload)
}

func (o *MonitorCheckGetTransactionDetailsOK) GetPayload() *models.TransactionDetailsResponse {
	return o.Payload
}

func (o *MonitorCheckGetTransactionDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TransactionDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMonitorCheckGetTransactionDetailsBadRequest creates a MonitorCheckGetTransactionDetailsBadRequest with default headers values
func NewMonitorCheckGetTransactionDetailsBadRequest() *MonitorCheckGetTransactionDetailsBadRequest {
	return &MonitorCheckGetTransactionDetailsBadRequest{}
}

/*MonitorCheckGetTransactionDetailsBadRequest handles this case with default header values.

The request failed.
*/
type MonitorCheckGetTransactionDetailsBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *MonitorCheckGetTransactionDetailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /MonitorCheck/{monitorCheckId}/Transaction][%d] monitorCheckGetTransactionDetailsBadRequest  %+v", 400, o.Payload)
}

func (o *MonitorCheckGetTransactionDetailsBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MonitorCheckGetTransactionDetailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
