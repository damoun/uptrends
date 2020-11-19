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

// OperatorGetAllOperatorsReader is a Reader for the OperatorGetAllOperators structure.
type OperatorGetAllOperatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperatorGetAllOperatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperatorGetAllOperatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOperatorGetAllOperatorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewOperatorGetAllOperatorsOK creates a OperatorGetAllOperatorsOK with default headers values
func NewOperatorGetAllOperatorsOK() *OperatorGetAllOperatorsOK {
	return &OperatorGetAllOperatorsOK{}
}

/*OperatorGetAllOperatorsOK handles this case with default header values.

The request completed successfully.
*/
type OperatorGetAllOperatorsOK struct {
	Payload []*models.Operator
}

func (o *OperatorGetAllOperatorsOK) Error() string {
	return fmt.Sprintf("[GET /Operator][%d] operatorGetAllOperatorsOK  %+v", 200, o.Payload)
}

func (o *OperatorGetAllOperatorsOK) GetPayload() []*models.Operator {
	return o.Payload
}

func (o *OperatorGetAllOperatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorGetAllOperatorsBadRequest creates a OperatorGetAllOperatorsBadRequest with default headers values
func NewOperatorGetAllOperatorsBadRequest() *OperatorGetAllOperatorsBadRequest {
	return &OperatorGetAllOperatorsBadRequest{}
}

/*OperatorGetAllOperatorsBadRequest handles this case with default header values.

The request failed.
*/
type OperatorGetAllOperatorsBadRequest struct {
	Payload *models.MessageList
}

func (o *OperatorGetAllOperatorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /Operator][%d] operatorGetAllOperatorsBadRequest  %+v", 400, o.Payload)
}

func (o *OperatorGetAllOperatorsBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *OperatorGetAllOperatorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
