// Code generated by go-swagger; DO NOT EDIT.

package vault

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// VaultCreateNewVaultSectionReader is a Reader for the VaultCreateNewVaultSection structure.
type VaultCreateNewVaultSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VaultCreateNewVaultSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVaultCreateNewVaultSectionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVaultCreateNewVaultSectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVaultCreateNewVaultSectionCreated creates a VaultCreateNewVaultSectionCreated with default headers values
func NewVaultCreateNewVaultSectionCreated() *VaultCreateNewVaultSectionCreated {
	return &VaultCreateNewVaultSectionCreated{}
}

/*VaultCreateNewVaultSectionCreated handles this case with default header values.

The request completed successfully.
*/
type VaultCreateNewVaultSectionCreated struct {
	Payload *models.VaultSection
}

func (o *VaultCreateNewVaultSectionCreated) Error() string {
	return fmt.Sprintf("[POST /VaultSection][%d] vaultCreateNewVaultSectionCreated  %+v", 201, o.Payload)
}

func (o *VaultCreateNewVaultSectionCreated) GetPayload() *models.VaultSection {
	return o.Payload
}

func (o *VaultCreateNewVaultSectionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VaultSection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVaultCreateNewVaultSectionBadRequest creates a VaultCreateNewVaultSectionBadRequest with default headers values
func NewVaultCreateNewVaultSectionBadRequest() *VaultCreateNewVaultSectionBadRequest {
	return &VaultCreateNewVaultSectionBadRequest{}
}

/*VaultCreateNewVaultSectionBadRequest handles this case with default header values.

The request failed.
*/
type VaultCreateNewVaultSectionBadRequest struct {
	Payload *models.MessageList
}

func (o *VaultCreateNewVaultSectionBadRequest) Error() string {
	return fmt.Sprintf("[POST /VaultSection][%d] vaultCreateNewVaultSectionBadRequest  %+v", 400, o.Payload)
}

func (o *VaultCreateNewVaultSectionBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *VaultCreateNewVaultSectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
