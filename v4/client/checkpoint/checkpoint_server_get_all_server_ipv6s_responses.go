// Code generated by go-swagger; DO NOT EDIT.

package checkpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/v4/models"
)

// CheckpointServerGetAllServerIpv6sReader is a Reader for the CheckpointServerGetAllServerIpv6s structure.
type CheckpointServerGetAllServerIpv6sReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckpointServerGetAllServerIpv6sReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckpointServerGetAllServerIpv6sOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckpointServerGetAllServerIpv6sBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCheckpointServerGetAllServerIpv6sOK creates a CheckpointServerGetAllServerIpv6sOK with default headers values
func NewCheckpointServerGetAllServerIpv6sOK() *CheckpointServerGetAllServerIpv6sOK {
	return &CheckpointServerGetAllServerIpv6sOK{}
}

/*CheckpointServerGetAllServerIpv6sOK handles this case with default header values.

CheckpointServerGetAllServerIpv6sOK checkpoint server get all server ipv6s o k
*/
type CheckpointServerGetAllServerIpv6sOK struct {
	Payload *models.ListStringResponse
}

func (o *CheckpointServerGetAllServerIpv6sOK) Error() string {
	return fmt.Sprintf("[GET /Checkpoint/Server/Ipv6][%d] checkpointServerGetAllServerIpv6sOK  %+v", 200, o.Payload)
}

func (o *CheckpointServerGetAllServerIpv6sOK) GetPayload() *models.ListStringResponse {
	return o.Payload
}

func (o *CheckpointServerGetAllServerIpv6sOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListStringResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckpointServerGetAllServerIpv6sBadRequest creates a CheckpointServerGetAllServerIpv6sBadRequest with default headers values
func NewCheckpointServerGetAllServerIpv6sBadRequest() *CheckpointServerGetAllServerIpv6sBadRequest {
	return &CheckpointServerGetAllServerIpv6sBadRequest{}
}

/*CheckpointServerGetAllServerIpv6sBadRequest handles this case with default header values.

The request failed.
*/
type CheckpointServerGetAllServerIpv6sBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *CheckpointServerGetAllServerIpv6sBadRequest) Error() string {
	return fmt.Sprintf("[GET /Checkpoint/Server/Ipv6][%d] checkpointServerGetAllServerIpv6sBadRequest  %+v", 400, o.Payload)
}

func (o *CheckpointServerGetAllServerIpv6sBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *CheckpointServerGetAllServerIpv6sBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
