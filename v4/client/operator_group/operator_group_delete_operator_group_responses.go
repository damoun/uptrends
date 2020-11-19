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

// OperatorGroupDeleteOperatorGroupReader is a Reader for the OperatorGroupDeleteOperatorGroup structure.
type OperatorGroupDeleteOperatorGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperatorGroupDeleteOperatorGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOperatorGroupDeleteOperatorGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOperatorGroupDeleteOperatorGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOperatorGroupDeleteOperatorGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOperatorGroupDeleteOperatorGroupNoContent creates a OperatorGroupDeleteOperatorGroupNoContent with default headers values
func NewOperatorGroupDeleteOperatorGroupNoContent() *OperatorGroupDeleteOperatorGroupNoContent {
	return &OperatorGroupDeleteOperatorGroupNoContent{}
}

/*OperatorGroupDeleteOperatorGroupNoContent handles this case with default header values.

The request completed successfully. No content is returned.
*/
type OperatorGroupDeleteOperatorGroupNoContent struct {
}

func (o *OperatorGroupDeleteOperatorGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /OperatorGroup/{operatorGroupGuid}][%d] operatorGroupDeleteOperatorGroupNoContent ", 204)
}

func (o *OperatorGroupDeleteOperatorGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOperatorGroupDeleteOperatorGroupBadRequest creates a OperatorGroupDeleteOperatorGroupBadRequest with default headers values
func NewOperatorGroupDeleteOperatorGroupBadRequest() *OperatorGroupDeleteOperatorGroupBadRequest {
	return &OperatorGroupDeleteOperatorGroupBadRequest{}
}

/*OperatorGroupDeleteOperatorGroupBadRequest handles this case with default header values.

The request failed.
or
Deleting the Alloperators group is not allowed.
*/
type OperatorGroupDeleteOperatorGroupBadRequest struct {
	Payload *models.MessageList
}

func (o *OperatorGroupDeleteOperatorGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /OperatorGroup/{operatorGroupGuid}][%d] operatorGroupDeleteOperatorGroupBadRequest  %+v", 400, o.Payload)
}

func (o *OperatorGroupDeleteOperatorGroupBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *OperatorGroupDeleteOperatorGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorGroupDeleteOperatorGroupNotFound creates a OperatorGroupDeleteOperatorGroupNotFound with default headers values
func NewOperatorGroupDeleteOperatorGroupNotFound() *OperatorGroupDeleteOperatorGroupNotFound {
	return &OperatorGroupDeleteOperatorGroupNotFound{}
}

/*OperatorGroupDeleteOperatorGroupNotFound handles this case with default header values.

The requested operator group was not found.
*/
type OperatorGroupDeleteOperatorGroupNotFound struct {
	Payload *models.MessageList
}

func (o *OperatorGroupDeleteOperatorGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /OperatorGroup/{operatorGroupGuid}][%d] operatorGroupDeleteOperatorGroupNotFound  %+v", 404, o.Payload)
}

func (o *OperatorGroupDeleteOperatorGroupNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *OperatorGroupDeleteOperatorGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}