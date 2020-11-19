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

// SLAPostExclusionPeriodReader is a Reader for the SLAPostExclusionPeriod structure.
type SLAPostExclusionPeriodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SLAPostExclusionPeriodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSLAPostExclusionPeriodCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSLAPostExclusionPeriodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSLAPostExclusionPeriodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSLAPostExclusionPeriodCreated creates a SLAPostExclusionPeriodCreated with default headers values
func NewSLAPostExclusionPeriodCreated() *SLAPostExclusionPeriodCreated {
	return &SLAPostExclusionPeriodCreated{}
}

/*SLAPostExclusionPeriodCreated handles this case with default header values.

The request completed successfully.
*/
type SLAPostExclusionPeriodCreated struct {
	Payload *models.ExclusionPeriod
}

func (o *SLAPostExclusionPeriodCreated) Error() string {
	return fmt.Sprintf("[POST /Sla/{slaGuid}/ExclusionPeriod][%d] slaPostExclusionPeriodCreated  %+v", 201, o.Payload)
}

func (o *SLAPostExclusionPeriodCreated) GetPayload() *models.ExclusionPeriod {
	return o.Payload
}

func (o *SLAPostExclusionPeriodCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExclusionPeriod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSLAPostExclusionPeriodBadRequest creates a SLAPostExclusionPeriodBadRequest with default headers values
func NewSLAPostExclusionPeriodBadRequest() *SLAPostExclusionPeriodBadRequest {
	return &SLAPostExclusionPeriodBadRequest{}
}

/*SLAPostExclusionPeriodBadRequest handles this case with default header values.

The request failed.
*/
type SLAPostExclusionPeriodBadRequest struct {
	Payload *models.MessageList
}

func (o *SLAPostExclusionPeriodBadRequest) Error() string {
	return fmt.Sprintf("[POST /Sla/{slaGuid}/ExclusionPeriod][%d] slaPostExclusionPeriodBadRequest  %+v", 400, o.Payload)
}

func (o *SLAPostExclusionPeriodBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *SLAPostExclusionPeriodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSLAPostExclusionPeriodNotFound creates a SLAPostExclusionPeriodNotFound with default headers values
func NewSLAPostExclusionPeriodNotFound() *SLAPostExclusionPeriodNotFound {
	return &SLAPostExclusionPeriodNotFound{}
}

/*SLAPostExclusionPeriodNotFound handles this case with default header values.

The specified SLA does not exist.
*/
type SLAPostExclusionPeriodNotFound struct {
	Payload *models.MessageList
}

func (o *SLAPostExclusionPeriodNotFound) Error() string {
	return fmt.Sprintf("[POST /Sla/{slaGuid}/ExclusionPeriod][%d] slaPostExclusionPeriodNotFound  %+v", 404, o.Payload)
}

func (o *SLAPostExclusionPeriodNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *SLAPostExclusionPeriodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
