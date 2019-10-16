// Code generated by go-swagger; DO NOT EDIT.

package operator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/models"
)

// OperatorUpdateOperatorReader is a Reader for the OperatorUpdateOperator structure.
type OperatorUpdateOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperatorUpdateOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOperatorUpdateOperatorNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOperatorUpdateOperatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewOperatorUpdateOperatorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOperatorUpdateOperatorNoContent creates a OperatorUpdateOperatorNoContent with default headers values
func NewOperatorUpdateOperatorNoContent() *OperatorUpdateOperatorNoContent {
	return &OperatorUpdateOperatorNoContent{}
}

/*OperatorUpdateOperatorNoContent handles this case with default header values.

Request completed successfully.
*/
type OperatorUpdateOperatorNoContent struct {
}

func (o *OperatorUpdateOperatorNoContent) Error() string {
	return fmt.Sprintf("[PUT /Operator/{operatorGuid}][%d] operatorUpdateOperatorNoContent ", 204)
}

func (o *OperatorUpdateOperatorNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOperatorUpdateOperatorBadRequest creates a OperatorUpdateOperatorBadRequest with default headers values
func NewOperatorUpdateOperatorBadRequest() *OperatorUpdateOperatorBadRequest {
	return &OperatorUpdateOperatorBadRequest{}
}

/*OperatorUpdateOperatorBadRequest handles this case with default header values.

The request failed.
*/
type OperatorUpdateOperatorBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *OperatorUpdateOperatorBadRequest) Error() string {
	return fmt.Sprintf("[PUT /Operator/{operatorGuid}][%d] operatorUpdateOperatorBadRequest  %+v", 400, o.Payload)
}

func (o *OperatorUpdateOperatorBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *OperatorUpdateOperatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorUpdateOperatorForbidden creates a OperatorUpdateOperatorForbidden with default headers values
func NewOperatorUpdateOperatorForbidden() *OperatorUpdateOperatorForbidden {
	return &OperatorUpdateOperatorForbidden{}
}

/*OperatorUpdateOperatorForbidden handles this case with default header values.

One or more validation errors occurred.
*/
type OperatorUpdateOperatorForbidden struct {
	Payload []*models.APIMessageInfo
}

func (o *OperatorUpdateOperatorForbidden) Error() string {
	return fmt.Sprintf("[PUT /Operator/{operatorGuid}][%d] operatorUpdateOperatorForbidden  %+v", 403, o.Payload)
}

func (o *OperatorUpdateOperatorForbidden) GetPayload() []*models.APIMessageInfo {
	return o.Payload
}

func (o *OperatorUpdateOperatorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
