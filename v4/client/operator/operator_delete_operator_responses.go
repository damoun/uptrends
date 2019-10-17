// Code generated by go-swagger; DO NOT EDIT.

package operator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/v4/models"
)

// OperatorDeleteOperatorReader is a Reader for the OperatorDeleteOperator structure.
type OperatorDeleteOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperatorDeleteOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOperatorDeleteOperatorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOperatorDeleteOperatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOperatorDeleteOperatorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOperatorDeleteOperatorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOperatorDeleteOperatorNoContent creates a OperatorDeleteOperatorNoContent with default headers values
func NewOperatorDeleteOperatorNoContent() *OperatorDeleteOperatorNoContent {
	return &OperatorDeleteOperatorNoContent{}
}

/*OperatorDeleteOperatorNoContent handles this case with default header values.

Request completed successfully
*/
type OperatorDeleteOperatorNoContent struct {
}

func (o *OperatorDeleteOperatorNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Operator/{operatorGuid}][%d] operatorDeleteOperatorNoContent ", 204)
}

func (o *OperatorDeleteOperatorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOperatorDeleteOperatorBadRequest creates a OperatorDeleteOperatorBadRequest with default headers values
func NewOperatorDeleteOperatorBadRequest() *OperatorDeleteOperatorBadRequest {
	return &OperatorDeleteOperatorBadRequest{}
}

/*OperatorDeleteOperatorBadRequest handles this case with default header values.

The request failed.
*/
type OperatorDeleteOperatorBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *OperatorDeleteOperatorBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /Operator/{operatorGuid}][%d] operatorDeleteOperatorBadRequest  %+v", 400, o.Payload)
}

func (o *OperatorDeleteOperatorBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *OperatorDeleteOperatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorDeleteOperatorForbidden creates a OperatorDeleteOperatorForbidden with default headers values
func NewOperatorDeleteOperatorForbidden() *OperatorDeleteOperatorForbidden {
	return &OperatorDeleteOperatorForbidden{}
}

/*OperatorDeleteOperatorForbidden handles this case with default header values.

Deleting the current operator is not allowed.
*/
type OperatorDeleteOperatorForbidden struct {
	Payload *models.APIMessageInfo
}

func (o *OperatorDeleteOperatorForbidden) Error() string {
	return fmt.Sprintf("[DELETE /Operator/{operatorGuid}][%d] operatorDeleteOperatorForbidden  %+v", 403, o.Payload)
}

func (o *OperatorDeleteOperatorForbidden) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *OperatorDeleteOperatorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorDeleteOperatorNotFound creates a OperatorDeleteOperatorNotFound with default headers values
func NewOperatorDeleteOperatorNotFound() *OperatorDeleteOperatorNotFound {
	return &OperatorDeleteOperatorNotFound{}
}

/*OperatorDeleteOperatorNotFound handles this case with default header values.

The specified operator was not found.
*/
type OperatorDeleteOperatorNotFound struct {
	Payload *models.APIMessageInfo
}

func (o *OperatorDeleteOperatorNotFound) Error() string {
	return fmt.Sprintf("[DELETE /Operator/{operatorGuid}][%d] operatorDeleteOperatorNotFound  %+v", 404, o.Payload)
}

func (o *OperatorDeleteOperatorNotFound) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *OperatorDeleteOperatorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}