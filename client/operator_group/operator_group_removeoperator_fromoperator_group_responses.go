// Code generated by go-swagger; DO NOT EDIT.

package operator_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/models"
)

// OperatorGroupRemoveoperatorFromoperatorGroupReader is a Reader for the OperatorGroupRemoveoperatorFromoperatorGroup structure.
type OperatorGroupRemoveoperatorFromoperatorGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperatorGroupRemoveoperatorFromoperatorGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewOperatorGroupRemoveoperatorFromoperatorGroupNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOperatorGroupRemoveoperatorFromoperatorGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewOperatorGroupRemoveoperatorFromoperatorGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOperatorGroupRemoveoperatorFromoperatorGroupNoContent creates a OperatorGroupRemoveoperatorFromoperatorGroupNoContent with default headers values
func NewOperatorGroupRemoveoperatorFromoperatorGroupNoContent() *OperatorGroupRemoveoperatorFromoperatorGroupNoContent {
	return &OperatorGroupRemoveoperatorFromoperatorGroupNoContent{}
}

/*OperatorGroupRemoveoperatorFromoperatorGroupNoContent handles this case with default header values.

Request completed successfully.
*/
type OperatorGroupRemoveoperatorFromoperatorGroupNoContent struct {
}

func (o *OperatorGroupRemoveoperatorFromoperatorGroupNoContent) Error() string {
	return fmt.Sprintf("[DELETE /OperatorGroup/{operatorGroupGuid}/Member/{operatorGuid}][%d] operatorGroupRemoveoperatorFromoperatorGroupNoContent ", 204)
}

func (o *OperatorGroupRemoveoperatorFromoperatorGroupNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOperatorGroupRemoveoperatorFromoperatorGroupBadRequest creates a OperatorGroupRemoveoperatorFromoperatorGroupBadRequest with default headers values
func NewOperatorGroupRemoveoperatorFromoperatorGroupBadRequest() *OperatorGroupRemoveoperatorFromoperatorGroupBadRequest {
	return &OperatorGroupRemoveoperatorFromoperatorGroupBadRequest{}
}

/*OperatorGroupRemoveoperatorFromoperatorGroupBadRequest handles this case with default header values.

The request failed.
or
A operator cannot be removed from the All operators group.
*/
type OperatorGroupRemoveoperatorFromoperatorGroupBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *OperatorGroupRemoveoperatorFromoperatorGroupBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /OperatorGroup/{operatorGroupGuid}/Member/{operatorGuid}][%d] operatorGroupRemoveoperatorFromoperatorGroupBadRequest  %+v", 400, o.Payload)
}

func (o *OperatorGroupRemoveoperatorFromoperatorGroupBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *OperatorGroupRemoveoperatorFromoperatorGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorGroupRemoveoperatorFromoperatorGroupNotFound creates a OperatorGroupRemoveoperatorFromoperatorGroupNotFound with default headers values
func NewOperatorGroupRemoveoperatorFromoperatorGroupNotFound() *OperatorGroupRemoveoperatorFromoperatorGroupNotFound {
	return &OperatorGroupRemoveoperatorFromoperatorGroupNotFound{}
}

/*OperatorGroupRemoveoperatorFromoperatorGroupNotFound handles this case with default header values.

The requested operator group was not found.
*/
type OperatorGroupRemoveoperatorFromoperatorGroupNotFound struct {
	Payload *models.APIMessageInfo
}

func (o *OperatorGroupRemoveoperatorFromoperatorGroupNotFound) Error() string {
	return fmt.Sprintf("[DELETE /OperatorGroup/{operatorGroupGuid}/Member/{operatorGuid}][%d] operatorGroupRemoveoperatorFromoperatorGroupNotFound  %+v", 404, o.Payload)
}

func (o *OperatorGroupRemoveoperatorFromoperatorGroupNotFound) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *OperatorGroupRemoveoperatorFromoperatorGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
