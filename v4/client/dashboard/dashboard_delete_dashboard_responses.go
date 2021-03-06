// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// DashboardDeleteDashboardReader is a Reader for the DashboardDeleteDashboard structure.
type DashboardDeleteDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DashboardDeleteDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDashboardDeleteDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDashboardDeleteDashboardBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDashboardDeleteDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDashboardDeleteDashboardNoContent creates a DashboardDeleteDashboardNoContent with default headers values
func NewDashboardDeleteDashboardNoContent() *DashboardDeleteDashboardNoContent {
	return &DashboardDeleteDashboardNoContent{}
}

/*DashboardDeleteDashboardNoContent handles this case with default header values.

The request completed successfully.
*/
type DashboardDeleteDashboardNoContent struct {
}

func (o *DashboardDeleteDashboardNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Dashboard/{dashboardGuid}][%d] dashboardDeleteDashboardNoContent ", 204)
}

func (o *DashboardDeleteDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDashboardDeleteDashboardBadRequest creates a DashboardDeleteDashboardBadRequest with default headers values
func NewDashboardDeleteDashboardBadRequest() *DashboardDeleteDashboardBadRequest {
	return &DashboardDeleteDashboardBadRequest{}
}

/*DashboardDeleteDashboardBadRequest handles this case with default header values.

The request failed.
*/
type DashboardDeleteDashboardBadRequest struct {
	Payload *models.MessageList
}

func (o *DashboardDeleteDashboardBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /Dashboard/{dashboardGuid}][%d] dashboardDeleteDashboardBadRequest  %+v", 400, o.Payload)
}

func (o *DashboardDeleteDashboardBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *DashboardDeleteDashboardBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardDeleteDashboardNotFound creates a DashboardDeleteDashboardNotFound with default headers values
func NewDashboardDeleteDashboardNotFound() *DashboardDeleteDashboardNotFound {
	return &DashboardDeleteDashboardNotFound{}
}

/*DashboardDeleteDashboardNotFound handles this case with default header values.

The specified dashboard could not be found.
*/
type DashboardDeleteDashboardNotFound struct {
	Payload *models.MessageList
}

func (o *DashboardDeleteDashboardNotFound) Error() string {
	return fmt.Sprintf("[DELETE /Dashboard/{dashboardGuid}][%d] dashboardDeleteDashboardNotFound  %+v", 404, o.Payload)
}

func (o *DashboardDeleteDashboardNotFound) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *DashboardDeleteDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
