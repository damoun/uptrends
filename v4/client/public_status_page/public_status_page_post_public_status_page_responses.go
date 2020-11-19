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

// PublicStatusPagePostPublicStatusPageReader is a Reader for the PublicStatusPagePostPublicStatusPage structure.
type PublicStatusPagePostPublicStatusPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicStatusPagePostPublicStatusPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPublicStatusPagePostPublicStatusPageCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicStatusPagePostPublicStatusPageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPublicStatusPagePostPublicStatusPageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPublicStatusPagePostPublicStatusPageCreated creates a PublicStatusPagePostPublicStatusPageCreated with default headers values
func NewPublicStatusPagePostPublicStatusPageCreated() *PublicStatusPagePostPublicStatusPageCreated {
	return &PublicStatusPagePostPublicStatusPageCreated{}
}

/*PublicStatusPagePostPublicStatusPageCreated handles this case with default header values.

The request completed successfully.
*/
type PublicStatusPagePostPublicStatusPageCreated struct {
	Payload *models.PublicStatusPage
}

func (o *PublicStatusPagePostPublicStatusPageCreated) Error() string {
	return fmt.Sprintf("[POST /PublicStatusPage][%d] publicStatusPagePostPublicStatusPageCreated  %+v", 201, o.Payload)
}

func (o *PublicStatusPagePostPublicStatusPageCreated) GetPayload() *models.PublicStatusPage {
	return o.Payload
}

func (o *PublicStatusPagePostPublicStatusPageCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PublicStatusPage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicStatusPagePostPublicStatusPageBadRequest creates a PublicStatusPagePostPublicStatusPageBadRequest with default headers values
func NewPublicStatusPagePostPublicStatusPageBadRequest() *PublicStatusPagePostPublicStatusPageBadRequest {
	return &PublicStatusPagePostPublicStatusPageBadRequest{}
}

/*PublicStatusPagePostPublicStatusPageBadRequest handles this case with default header values.

The request failed.
*/
type PublicStatusPagePostPublicStatusPageBadRequest struct {
	Payload *models.MessageList
}

func (o *PublicStatusPagePostPublicStatusPageBadRequest) Error() string {
	return fmt.Sprintf("[POST /PublicStatusPage][%d] publicStatusPagePostPublicStatusPageBadRequest  %+v", 400, o.Payload)
}

func (o *PublicStatusPagePostPublicStatusPageBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *PublicStatusPagePostPublicStatusPageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicStatusPagePostPublicStatusPageNotFound creates a PublicStatusPagePostPublicStatusPageNotFound with default headers values
func NewPublicStatusPagePostPublicStatusPageNotFound() *PublicStatusPagePostPublicStatusPageNotFound {
	return &PublicStatusPagePostPublicStatusPageNotFound{}
}

/*PublicStatusPagePostPublicStatusPageNotFound handles this case with default header values.

The specified public status page does not exist.
*/
type PublicStatusPagePostPublicStatusPageNotFound struct {
	Payload *models.MessageList
}

func (o *PublicStatusPagePostPublicStatusPageNotFound) Error() string {
	return fmt.Sprintf("[POST /PublicStatusPage][%d] publicStatusPagePostPublicStatusPageNotFound  %+v", 404, o.Payload)
}

func (o *PublicStatusPagePostPublicStatusPageNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *PublicStatusPagePostPublicStatusPageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
