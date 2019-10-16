// Code generated by go-swagger; DO NOT EDIT.

package miscellaneous

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/models"
)

// MiscellaneousGetTimezoneByIDReader is a Reader for the MiscellaneousGetTimezoneByID structure.
type MiscellaneousGetTimezoneByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MiscellaneousGetTimezoneByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMiscellaneousGetTimezoneByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMiscellaneousGetTimezoneByIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMiscellaneousGetTimezoneByIDOK creates a MiscellaneousGetTimezoneByIDOK with default headers values
func NewMiscellaneousGetTimezoneByIDOK() *MiscellaneousGetTimezoneByIDOK {
	return &MiscellaneousGetTimezoneByIDOK{}
}

/*MiscellaneousGetTimezoneByIDOK handles this case with default header values.

Request completed successfully.
*/
type MiscellaneousGetTimezoneByIDOK struct {
	Payload *models.Timezone
}

func (o *MiscellaneousGetTimezoneByIDOK) Error() string {
	return fmt.Sprintf("[GET /Timezone/{timezoneId}][%d] miscellaneousGetTimezoneByIdOK  %+v", 200, o.Payload)
}

func (o *MiscellaneousGetTimezoneByIDOK) GetPayload() *models.Timezone {
	return o.Payload
}

func (o *MiscellaneousGetTimezoneByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Timezone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMiscellaneousGetTimezoneByIDBadRequest creates a MiscellaneousGetTimezoneByIDBadRequest with default headers values
func NewMiscellaneousGetTimezoneByIDBadRequest() *MiscellaneousGetTimezoneByIDBadRequest {
	return &MiscellaneousGetTimezoneByIDBadRequest{}
}

/*MiscellaneousGetTimezoneByIDBadRequest handles this case with default header values.

The request failed.
*/
type MiscellaneousGetTimezoneByIDBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *MiscellaneousGetTimezoneByIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /Timezone/{timezoneId}][%d] miscellaneousGetTimezoneByIdBadRequest  %+v", 400, o.Payload)
}

func (o *MiscellaneousGetTimezoneByIDBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *MiscellaneousGetTimezoneByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
