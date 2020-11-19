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

// PublicStatusPageGetPublicStatusPagesReader is a Reader for the PublicStatusPageGetPublicStatusPages structure.
type PublicStatusPageGetPublicStatusPagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PublicStatusPageGetPublicStatusPagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPublicStatusPageGetPublicStatusPagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPublicStatusPageGetPublicStatusPagesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPublicStatusPageGetPublicStatusPagesOK creates a PublicStatusPageGetPublicStatusPagesOK with default headers values
func NewPublicStatusPageGetPublicStatusPagesOK() *PublicStatusPageGetPublicStatusPagesOK {
	return &PublicStatusPageGetPublicStatusPagesOK{}
}

/*PublicStatusPageGetPublicStatusPagesOK handles this case with default header values.

The request completed successfully.
*/
type PublicStatusPageGetPublicStatusPagesOK struct {
	Payload []*models.PublicStatusPage
}

func (o *PublicStatusPageGetPublicStatusPagesOK) Error() string {
	return fmt.Sprintf("[GET /PublicStatusPage][%d] publicStatusPageGetPublicStatusPagesOK  %+v", 200, o.Payload)
}

func (o *PublicStatusPageGetPublicStatusPagesOK) GetPayload() []*models.PublicStatusPage {
	return o.Payload
}

func (o *PublicStatusPageGetPublicStatusPagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPublicStatusPageGetPublicStatusPagesBadRequest creates a PublicStatusPageGetPublicStatusPagesBadRequest with default headers values
func NewPublicStatusPageGetPublicStatusPagesBadRequest() *PublicStatusPageGetPublicStatusPagesBadRequest {
	return &PublicStatusPageGetPublicStatusPagesBadRequest{}
}

/*PublicStatusPageGetPublicStatusPagesBadRequest handles this case with default header values.

The request failed.
*/
type PublicStatusPageGetPublicStatusPagesBadRequest struct {
	Payload *models.MessageList
}

func (o *PublicStatusPageGetPublicStatusPagesBadRequest) Error() string {
	return fmt.Sprintf("[GET /PublicStatusPage][%d] publicStatusPageGetPublicStatusPagesBadRequest  %+v", 400, o.Payload)
}

func (o *PublicStatusPageGetPublicStatusPagesBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *PublicStatusPageGetPublicStatusPagesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
