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

// OperatorGetDutyScheduleForOperatorReader is a Reader for the OperatorGetDutyScheduleForOperator structure.
type OperatorGetDutyScheduleForOperatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperatorGetDutyScheduleForOperatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOperatorGetDutyScheduleForOperatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOperatorGetDutyScheduleForOperatorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOperatorGetDutyScheduleForOperatorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOperatorGetDutyScheduleForOperatorOK creates a OperatorGetDutyScheduleForOperatorOK with default headers values
func NewOperatorGetDutyScheduleForOperatorOK() *OperatorGetDutyScheduleForOperatorOK {
	return &OperatorGetDutyScheduleForOperatorOK{}
}

/*OperatorGetDutyScheduleForOperatorOK handles this case with default header values.

Request completed successfully.
*/
type OperatorGetDutyScheduleForOperatorOK struct {
	Payload []*models.OperatorDutySchedule
}

func (o *OperatorGetDutyScheduleForOperatorOK) Error() string {
	return fmt.Sprintf("[GET /Operator/{operatorGuid}/DutySchedule][%d] operatorGetDutyScheduleForOperatorOK  %+v", 200, o.Payload)
}

func (o *OperatorGetDutyScheduleForOperatorOK) GetPayload() []*models.OperatorDutySchedule {
	return o.Payload
}

func (o *OperatorGetDutyScheduleForOperatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorGetDutyScheduleForOperatorBadRequest creates a OperatorGetDutyScheduleForOperatorBadRequest with default headers values
func NewOperatorGetDutyScheduleForOperatorBadRequest() *OperatorGetDutyScheduleForOperatorBadRequest {
	return &OperatorGetDutyScheduleForOperatorBadRequest{}
}

/*OperatorGetDutyScheduleForOperatorBadRequest handles this case with default header values.

The request failed.
*/
type OperatorGetDutyScheduleForOperatorBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *OperatorGetDutyScheduleForOperatorBadRequest) Error() string {
	return fmt.Sprintf("[GET /Operator/{operatorGuid}/DutySchedule][%d] operatorGetDutyScheduleForOperatorBadRequest  %+v", 400, o.Payload)
}

func (o *OperatorGetDutyScheduleForOperatorBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *OperatorGetDutyScheduleForOperatorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorGetDutyScheduleForOperatorNotFound creates a OperatorGetDutyScheduleForOperatorNotFound with default headers values
func NewOperatorGetDutyScheduleForOperatorNotFound() *OperatorGetDutyScheduleForOperatorNotFound {
	return &OperatorGetDutyScheduleForOperatorNotFound{}
}

/*OperatorGetDutyScheduleForOperatorNotFound handles this case with default header values.

The specified operator was not found.
*/
type OperatorGetDutyScheduleForOperatorNotFound struct {
	Payload *models.APIMessageInfo
}

func (o *OperatorGetDutyScheduleForOperatorNotFound) Error() string {
	return fmt.Sprintf("[GET /Operator/{operatorGuid}/DutySchedule][%d] operatorGetDutyScheduleForOperatorNotFound  %+v", 404, o.Payload)
}

func (o *OperatorGetDutyScheduleForOperatorNotFound) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *OperatorGetDutyScheduleForOperatorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}