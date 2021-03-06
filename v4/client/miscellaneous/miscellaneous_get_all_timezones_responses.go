// Code generated by go-swagger; DO NOT EDIT.

package miscellaneous

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// MiscellaneousGetAllTimezonesReader is a Reader for the MiscellaneousGetAllTimezones structure.
type MiscellaneousGetAllTimezonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MiscellaneousGetAllTimezonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMiscellaneousGetAllTimezonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMiscellaneousGetAllTimezonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewMiscellaneousGetAllTimezonesOK creates a MiscellaneousGetAllTimezonesOK with default headers values
func NewMiscellaneousGetAllTimezonesOK() *MiscellaneousGetAllTimezonesOK {
	return &MiscellaneousGetAllTimezonesOK{}
}

/*MiscellaneousGetAllTimezonesOK handles this case with default header values.

The request completed successfully.
*/
type MiscellaneousGetAllTimezonesOK struct {
	Payload []*models.Timezone
}

func (o *MiscellaneousGetAllTimezonesOK) Error() string {
	return fmt.Sprintf("[GET /Timezone][%d] miscellaneousGetAllTimezonesOK  %+v", 200, o.Payload)
}

func (o *MiscellaneousGetAllTimezonesOK) GetPayload() []*models.Timezone {
	return o.Payload
}

func (o *MiscellaneousGetAllTimezonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMiscellaneousGetAllTimezonesBadRequest creates a MiscellaneousGetAllTimezonesBadRequest with default headers values
func NewMiscellaneousGetAllTimezonesBadRequest() *MiscellaneousGetAllTimezonesBadRequest {
	return &MiscellaneousGetAllTimezonesBadRequest{}
}

/*MiscellaneousGetAllTimezonesBadRequest handles this case with default header values.

The request failed.
*/
type MiscellaneousGetAllTimezonesBadRequest struct {
	Payload *models.MessageList
}

func (o *MiscellaneousGetAllTimezonesBadRequest) Error() string {
	return fmt.Sprintf("[GET /Timezone][%d] miscellaneousGetAllTimezonesBadRequest  %+v", 400, o.Payload)
}

func (o *MiscellaneousGetAllTimezonesBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *MiscellaneousGetAllTimezonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
