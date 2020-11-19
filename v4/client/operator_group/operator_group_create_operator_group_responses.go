// Code generated by go-swagger; DO NOT EDIT.

package operator_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// OperatorGroupCreateOperatorGroupReader is a Reader for the OperatorGroupCreateOperatorGroup structure.
type OperatorGroupCreateOperatorGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperatorGroupCreateOperatorGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOperatorGroupCreateOperatorGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOperatorGroupCreateOperatorGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOperatorGroupCreateOperatorGroupCreated creates a OperatorGroupCreateOperatorGroupCreated with default headers values
func NewOperatorGroupCreateOperatorGroupCreated() *OperatorGroupCreateOperatorGroupCreated {
	return &OperatorGroupCreateOperatorGroupCreated{}
}

/*OperatorGroupCreateOperatorGroupCreated handles this case with default header values.

The request completed successfully.
*/
type OperatorGroupCreateOperatorGroupCreated struct {
	Payload *models.OperatorGroup
}

func (o *OperatorGroupCreateOperatorGroupCreated) Error() string {
	return fmt.Sprintf("[POST /OperatorGroup][%d] operatorGroupCreateOperatorGroupCreated  %+v", 201, o.Payload)
}

func (o *OperatorGroupCreateOperatorGroupCreated) GetPayload() *models.OperatorGroup {
	return o.Payload
}

func (o *OperatorGroupCreateOperatorGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OperatorGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorGroupCreateOperatorGroupBadRequest creates a OperatorGroupCreateOperatorGroupBadRequest with default headers values
func NewOperatorGroupCreateOperatorGroupBadRequest() *OperatorGroupCreateOperatorGroupBadRequest {
	return &OperatorGroupCreateOperatorGroupBadRequest{}
}

/*OperatorGroupCreateOperatorGroupBadRequest handles this case with default header values.

The request failed.
*/
type OperatorGroupCreateOperatorGroupBadRequest struct {
	Payload *models.MessageList
}

func (o *OperatorGroupCreateOperatorGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /OperatorGroup][%d] operatorGroupCreateOperatorGroupBadRequest  %+v", 400, o.Payload)
}

func (o *OperatorGroupCreateOperatorGroupBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *OperatorGroupCreateOperatorGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}