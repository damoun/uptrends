// Code generated by go-swagger; DO NOT EDIT.

package sla

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// SLADeleteExclusionPeriodReader is a Reader for the SLADeleteExclusionPeriod structure.
type SLADeleteExclusionPeriodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SLADeleteExclusionPeriodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSLADeleteExclusionPeriodNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSLADeleteExclusionPeriodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSLADeleteExclusionPeriodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSLADeleteExclusionPeriodNoContent creates a SLADeleteExclusionPeriodNoContent with default headers values
func NewSLADeleteExclusionPeriodNoContent() *SLADeleteExclusionPeriodNoContent {
	return &SLADeleteExclusionPeriodNoContent{}
}

/*SLADeleteExclusionPeriodNoContent handles this case with default header values.

The request completed successfully.
*/
type SLADeleteExclusionPeriodNoContent struct {
}

func (o *SLADeleteExclusionPeriodNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId}][%d] slaDeleteExclusionPeriodNoContent ", 204)
}

func (o *SLADeleteExclusionPeriodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSLADeleteExclusionPeriodBadRequest creates a SLADeleteExclusionPeriodBadRequest with default headers values
func NewSLADeleteExclusionPeriodBadRequest() *SLADeleteExclusionPeriodBadRequest {
	return &SLADeleteExclusionPeriodBadRequest{}
}

/*SLADeleteExclusionPeriodBadRequest handles this case with default header values.

The request failed.
*/
type SLADeleteExclusionPeriodBadRequest struct {
	Payload *models.MessageList
}

func (o *SLADeleteExclusionPeriodBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId}][%d] slaDeleteExclusionPeriodBadRequest  %+v", 400, o.Payload)
}

func (o *SLADeleteExclusionPeriodBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *SLADeleteExclusionPeriodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSLADeleteExclusionPeriodNotFound creates a SLADeleteExclusionPeriodNotFound with default headers values
func NewSLADeleteExclusionPeriodNotFound() *SLADeleteExclusionPeriodNotFound {
	return &SLADeleteExclusionPeriodNotFound{}
}

/*SLADeleteExclusionPeriodNotFound handles this case with default header values.

The specified SLA does not exist.
or
The specified exclusion period does not exist.
*/
type SLADeleteExclusionPeriodNotFound struct {
	Payload *models.MessageList
}

func (o *SLADeleteExclusionPeriodNotFound) Error() string {
	return fmt.Sprintf("[DELETE /Sla/{slaGuid}/ExclusionPeriod/{exclusionPeriodId}][%d] slaDeleteExclusionPeriodNotFound  %+v", 404, o.Payload)
}

func (o *SLADeleteExclusionPeriodNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *SLADeleteExclusionPeriodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
