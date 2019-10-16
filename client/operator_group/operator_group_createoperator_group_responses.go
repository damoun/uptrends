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

// OperatorGroupCreateoperatorGroupReader is a Reader for the OperatorGroupCreateoperatorGroup structure.
type OperatorGroupCreateoperatorGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OperatorGroupCreateoperatorGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewOperatorGroupCreateoperatorGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewOperatorGroupCreateoperatorGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOperatorGroupCreateoperatorGroupCreated creates a OperatorGroupCreateoperatorGroupCreated with default headers values
func NewOperatorGroupCreateoperatorGroupCreated() *OperatorGroupCreateoperatorGroupCreated {
	return &OperatorGroupCreateoperatorGroupCreated{}
}

/*OperatorGroupCreateoperatorGroupCreated handles this case with default header values.

Request completed successfully.
*/
type OperatorGroupCreateoperatorGroupCreated struct {
	Payload *models.OperatorGroup
}

func (o *OperatorGroupCreateoperatorGroupCreated) Error() string {
	return fmt.Sprintf("[POST /OperatorGroup][%d] operatorGroupCreateoperatorGroupCreated  %+v", 201, o.Payload)
}

func (o *OperatorGroupCreateoperatorGroupCreated) GetPayload() *models.OperatorGroup {
	return o.Payload
}

func (o *OperatorGroupCreateoperatorGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OperatorGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOperatorGroupCreateoperatorGroupBadRequest creates a OperatorGroupCreateoperatorGroupBadRequest with default headers values
func NewOperatorGroupCreateoperatorGroupBadRequest() *OperatorGroupCreateoperatorGroupBadRequest {
	return &OperatorGroupCreateoperatorGroupBadRequest{}
}

/*OperatorGroupCreateoperatorGroupBadRequest handles this case with default header values.

The request failed.
*/
type OperatorGroupCreateoperatorGroupBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *OperatorGroupCreateoperatorGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /OperatorGroup][%d] operatorGroupCreateoperatorGroupBadRequest  %+v", 400, o.Payload)
}

func (o *OperatorGroupCreateoperatorGroupBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *OperatorGroupCreateoperatorGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
