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

// DashboardGetAllDashboardsReader is a Reader for the DashboardGetAllDashboards structure.
type DashboardGetAllDashboardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DashboardGetAllDashboardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDashboardGetAllDashboardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDashboardGetAllDashboardsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDashboardGetAllDashboardsOK creates a DashboardGetAllDashboardsOK with default headers values
func NewDashboardGetAllDashboardsOK() *DashboardGetAllDashboardsOK {
	return &DashboardGetAllDashboardsOK{}
}

/*DashboardGetAllDashboardsOK handles this case with default header values.

The request completed successfully.
*/
type DashboardGetAllDashboardsOK struct {
	Payload []*models.Dashboard
}

func (o *DashboardGetAllDashboardsOK) Error() string {
	return fmt.Sprintf("[GET /Dashboard][%d] dashboardGetAllDashboardsOK  %+v", 200, o.Payload)
}

func (o *DashboardGetAllDashboardsOK) GetPayload() []*models.Dashboard {
	return o.Payload
}

func (o *DashboardGetAllDashboardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDashboardGetAllDashboardsBadRequest creates a DashboardGetAllDashboardsBadRequest with default headers values
func NewDashboardGetAllDashboardsBadRequest() *DashboardGetAllDashboardsBadRequest {
	return &DashboardGetAllDashboardsBadRequest{}
}

/*DashboardGetAllDashboardsBadRequest handles this case with default header values.

The request failed.
*/
type DashboardGetAllDashboardsBadRequest struct {
	Payload *models.MessageList
}

func (o *DashboardGetAllDashboardsBadRequest) Error() string {
	return fmt.Sprintf("[GET /Dashboard][%d] dashboardGetAllDashboardsBadRequest  %+v", 400, o.Payload)
}

func (o *DashboardGetAllDashboardsBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *DashboardGetAllDashboardsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
