// Code generated by go-swagger; DO NOT EDIT.

package operator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// OperatorUpdateOperatorWithPatchReader is a Reader for the OperatorUpdateOperatorWithPatch structure.
type OperatorUpdateOperatorWithPatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperatorUpdateOperatorWithPatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOperatorUpdateOperatorWithPatchNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOperatorUpdateOperatorWithPatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOperatorUpdateOperatorWithPatchForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOperatorUpdateOperatorWithPatchNoContent creates a OperatorUpdateOperatorWithPatchNoContent with default headers values
func NewOperatorUpdateOperatorWithPatchNoContent() *OperatorUpdateOperatorWithPatchNoContent {
	return &OperatorUpdateOperatorWithPatchNoContent{}
}

/*OperatorUpdateOperatorWithPatchNoContent handles this case with default header values.

The request completed successfully. No content is returned.
*/
type OperatorUpdateOperatorWithPatchNoContent struct {
}

func (o *OperatorUpdateOperatorWithPatchNoContent) Error() string {
	return fmt.Sprintf("[PATCH /Operator/{operatorGuid}][%d] operatorUpdateOperatorWithPatchNoContent ", 204)
}

func (o *OperatorUpdateOperatorWithPatchNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOperatorUpdateOperatorWithPatchBadRequest creates a OperatorUpdateOperatorWithPatchBadRequest with default headers values
func NewOperatorUpdateOperatorWithPatchBadRequest() *OperatorUpdateOperatorWithPatchBadRequest {
	return &OperatorUpdateOperatorWithPatchBadRequest{}
}

/*OperatorUpdateOperatorWithPatchBadRequest handles this case with default header values.

The request failed.
*/
type OperatorUpdateOperatorWithPatchBadRequest struct {
	Payload *models.MessageList
}

func (o *OperatorUpdateOperatorWithPatchBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /Operator/{operatorGuid}][%d] operatorUpdateOperatorWithPatchBadRequest  %+v", 400, o.Payload)
}

func (o *OperatorUpdateOperatorWithPatchBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *OperatorUpdateOperatorWithPatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorUpdateOperatorWithPatchForbidden creates a OperatorUpdateOperatorWithPatchForbidden with default headers values
func NewOperatorUpdateOperatorWithPatchForbidden() *OperatorUpdateOperatorWithPatchForbidden {
	return &OperatorUpdateOperatorWithPatchForbidden{}
}

/*OperatorUpdateOperatorWithPatchForbidden handles this case with default header values.

One or more validation errors occurred.
*/
type OperatorUpdateOperatorWithPatchForbidden struct {
	Payload *models.MessageList
}

func (o *OperatorUpdateOperatorWithPatchForbidden) Error() string {
	return fmt.Sprintf("[PATCH /Operator/{operatorGuid}][%d] operatorUpdateOperatorWithPatchForbidden  %+v", 403, o.Payload)
}

func (o *OperatorUpdateOperatorWithPatchForbidden) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *OperatorUpdateOperatorWithPatchForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
