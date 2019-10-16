// Code generated by go-swagger; DO NOT EDIT.

package register

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/models"
)

// RegisterPostReader is a Reader for the RegisterPost structure.
type RegisterPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterPostOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRegisterPostOK creates a RegisterPostOK with default headers values
func NewRegisterPostOK() *RegisterPostOK {
	return &RegisterPostOK{}
}

/*RegisterPostOK handles this case with default header values.

If you get this response, a new API account was created successfully.
*/
type RegisterPostOK struct {
	Payload *models.RegistrationResponse
}

func (o *RegisterPostOK) Error() string {
	return fmt.Sprintf("[POST /Register][%d] registerPostOK  %+v", 200, o.Payload)
}

func (o *RegisterPostOK) GetPayload() *models.RegistrationResponse {
	return o.Payload
}

func (o *RegisterPostOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}