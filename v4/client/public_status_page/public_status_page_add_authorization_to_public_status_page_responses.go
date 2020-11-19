// Code generated by go-swagger; DO NOT EDIT.

package public_status_page

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// PublicStatusPageAddAuthorizationToPublicStatusPageReader is a Reader for the PublicStatusPageAddAuthorizationToPublicStatusPage structure.
type PublicStatusPageAddAuthorizationToPublicStatusPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicStatusPageAddAuthorizationToPublicStatusPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicStatusPageAddAuthorizationToPublicStatusPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicStatusPageAddAuthorizationToPublicStatusPageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPublicStatusPageAddAuthorizationToPublicStatusPageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPublicStatusPageAddAuthorizationToPublicStatusPageOK creates a PublicStatusPageAddAuthorizationToPublicStatusPageOK with default headers values
func NewPublicStatusPageAddAuthorizationToPublicStatusPageOK() *PublicStatusPageAddAuthorizationToPublicStatusPageOK {
	return &PublicStatusPageAddAuthorizationToPublicStatusPageOK{}
}

/*PublicStatusPageAddAuthorizationToPublicStatusPageOK handles this case with default header values.

The request was successful.
*/
type PublicStatusPageAddAuthorizationToPublicStatusPageOK struct {
	Payload *models.PSPAuthorization
}

func (o *PublicStatusPageAddAuthorizationToPublicStatusPageOK) Error() string {
	return fmt.Sprintf("[POST /PublicStatusPage/{publicStatusPageGuid}/Authorization][%d] publicStatusPageAddAuthorizationToPublicStatusPageOK  %+v", 200, o.Payload)
}

func (o *PublicStatusPageAddAuthorizationToPublicStatusPageOK) GetPayload() *models.PSPAuthorization {
	return o.Payload
}

func (o *PublicStatusPageAddAuthorizationToPublicStatusPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PSPAuthorization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicStatusPageAddAuthorizationToPublicStatusPageBadRequest creates a PublicStatusPageAddAuthorizationToPublicStatusPageBadRequest with default headers values
func NewPublicStatusPageAddAuthorizationToPublicStatusPageBadRequest() *PublicStatusPageAddAuthorizationToPublicStatusPageBadRequest {
	return &PublicStatusPageAddAuthorizationToPublicStatusPageBadRequest{}
}

/*PublicStatusPageAddAuthorizationToPublicStatusPageBadRequest handles this case with default header values.

The request failed.
*/
type PublicStatusPageAddAuthorizationToPublicStatusPageBadRequest struct {
	Payload *models.MessageList
}

func (o *PublicStatusPageAddAuthorizationToPublicStatusPageBadRequest) Error() string {
	return fmt.Sprintf("[POST /PublicStatusPage/{publicStatusPageGuid}/Authorization][%d] publicStatusPageAddAuthorizationToPublicStatusPageBadRequest  %+v", 400, o.Payload)
}

func (o *PublicStatusPageAddAuthorizationToPublicStatusPageBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *PublicStatusPageAddAuthorizationToPublicStatusPageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicStatusPageAddAuthorizationToPublicStatusPageNotFound creates a PublicStatusPageAddAuthorizationToPublicStatusPageNotFound with default headers values
func NewPublicStatusPageAddAuthorizationToPublicStatusPageNotFound() *PublicStatusPageAddAuthorizationToPublicStatusPageNotFound {
	return &PublicStatusPageAddAuthorizationToPublicStatusPageNotFound{}
}

/*PublicStatusPageAddAuthorizationToPublicStatusPageNotFound handles this case with default header values.

The specified public status page does not exist.
or
The specified operator does not exist.
or
The specified operator group does not exist.
*/
type PublicStatusPageAddAuthorizationToPublicStatusPageNotFound struct {
	Payload *models.MessageList
}

func (o *PublicStatusPageAddAuthorizationToPublicStatusPageNotFound) Error() string {
	return fmt.Sprintf("[POST /PublicStatusPage/{publicStatusPageGuid}/Authorization][%d] publicStatusPageAddAuthorizationToPublicStatusPageNotFound  %+v", 404, o.Payload)
}

func (o *PublicStatusPageAddAuthorizationToPublicStatusPageNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *PublicStatusPageAddAuthorizationToPublicStatusPageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}