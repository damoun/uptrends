// Code generated by go-swagger; DO NOT EDIT.

package checkpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// CheckpointRegionGetAllCheckpointRegionsReader is a Reader for the CheckpointRegionGetAllCheckpointRegions structure.
type CheckpointRegionGetAllCheckpointRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckpointRegionGetAllCheckpointRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckpointRegionGetAllCheckpointRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCheckpointRegionGetAllCheckpointRegionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckpointRegionGetAllCheckpointRegionsOK creates a CheckpointRegionGetAllCheckpointRegionsOK with default headers values
func NewCheckpointRegionGetAllCheckpointRegionsOK() *CheckpointRegionGetAllCheckpointRegionsOK {
	return &CheckpointRegionGetAllCheckpointRegionsOK{}
}

/*CheckpointRegionGetAllCheckpointRegionsOK handles this case with default header values.

The request completed successfully.
*/
type CheckpointRegionGetAllCheckpointRegionsOK struct {
	Payload []*models.CheckpointRegion
}

func (o *CheckpointRegionGetAllCheckpointRegionsOK) Error() string {
	return fmt.Sprintf("[GET /CheckpointRegion][%d] checkpointRegionGetAllCheckpointRegionsOK  %+v", 200, o.Payload)
}

func (o *CheckpointRegionGetAllCheckpointRegionsOK) GetPayload() []*models.CheckpointRegion {
	return o.Payload
}

func (o *CheckpointRegionGetAllCheckpointRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckpointRegionGetAllCheckpointRegionsBadRequest creates a CheckpointRegionGetAllCheckpointRegionsBadRequest with default headers values
func NewCheckpointRegionGetAllCheckpointRegionsBadRequest() *CheckpointRegionGetAllCheckpointRegionsBadRequest {
	return &CheckpointRegionGetAllCheckpointRegionsBadRequest{}
}

/*CheckpointRegionGetAllCheckpointRegionsBadRequest handles this case with default header values.

The request failed.
*/
type CheckpointRegionGetAllCheckpointRegionsBadRequest struct {
	Payload *models.MessageList
}

func (o *CheckpointRegionGetAllCheckpointRegionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /CheckpointRegion][%d] checkpointRegionGetAllCheckpointRegionsBadRequest  %+v", 400, o.Payload)
}

func (o *CheckpointRegionGetAllCheckpointRegionsBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *CheckpointRegionGetAllCheckpointRegionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
