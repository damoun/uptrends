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

// SLADeleteSLAReader is a Reader for the SLADeleteSLA structure.
type SLADeleteSLAReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SLADeleteSLAReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSLADeleteSLANoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSLADeleteSLABadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSLADeleteSLAForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSLADeleteSLANotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSLADeleteSLANoContent creates a SLADeleteSLANoContent with default headers values
func NewSLADeleteSLANoContent() *SLADeleteSLANoContent {
	return &SLADeleteSLANoContent{}
}

/*SLADeleteSLANoContent handles this case with default header values.

The request completed successfully.
*/
type SLADeleteSLANoContent struct {
}

func (o *SLADeleteSLANoContent) Error() string {
	return fmt.Sprintf("[DELETE /Sla/{slaGuid}][%d] slaDeleteSlaNoContent ", 204)
}

func (o *SLADeleteSLANoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSLADeleteSLABadRequest creates a SLADeleteSLABadRequest with default headers values
func NewSLADeleteSLABadRequest() *SLADeleteSLABadRequest {
	return &SLADeleteSLABadRequest{}
}

/*SLADeleteSLABadRequest handles this case with default header values.

The request failed.
*/
type SLADeleteSLABadRequest struct {
	Payload *models.MessageList
}

func (o *SLADeleteSLABadRequest) Error() string {
	return fmt.Sprintf("[DELETE /Sla/{slaGuid}][%d] slaDeleteSlaBadRequest  %+v", 400, o.Payload)
}

func (o *SLADeleteSLABadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *SLADeleteSLABadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSLADeleteSLAForbidden creates a SLADeleteSLAForbidden with default headers values
func NewSLADeleteSLAForbidden() *SLADeleteSLAForbidden {
	return &SLADeleteSLAForbidden{}
}

/*SLADeleteSLAForbidden handles this case with default header values.

One or more validation errors occurred.
*/
type SLADeleteSLAForbidden struct {
	Payload *models.MessageList
}

func (o *SLADeleteSLAForbidden) Error() string {
	return fmt.Sprintf("[DELETE /Sla/{slaGuid}][%d] slaDeleteSlaForbidden  %+v", 403, o.Payload)
}

func (o *SLADeleteSLAForbidden) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *SLADeleteSLAForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSLADeleteSLANotFound creates a SLADeleteSLANotFound with default headers values
func NewSLADeleteSLANotFound() *SLADeleteSLANotFound {
	return &SLADeleteSLANotFound{}
}

/*SLADeleteSLANotFound handles this case with default header values.

The specified SLA does not exist.
*/
type SLADeleteSLANotFound struct {
	Payload *models.MessageList
}

func (o *SLADeleteSLANotFound) Error() string {
	return fmt.Sprintf("[DELETE /Sla/{slaGuid}][%d] slaDeleteSlaNotFound  %+v", 404, o.Payload)
}

func (o *SLADeleteSLANotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *SLADeleteSLANotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
