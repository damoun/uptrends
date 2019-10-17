// Code generated by go-swagger; DO NOT EDIT.

package vault

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/v4/models"
)

// VaultCreateNewVaultItemReader is a Reader for the VaultCreateNewVaultItem structure.
type VaultCreateNewVaultItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VaultCreateNewVaultItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVaultCreateNewVaultItemCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVaultCreateNewVaultItemBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVaultCreateNewVaultItemCreated creates a VaultCreateNewVaultItemCreated with default headers values
func NewVaultCreateNewVaultItemCreated() *VaultCreateNewVaultItemCreated {
	return &VaultCreateNewVaultItemCreated{}
}

/*VaultCreateNewVaultItemCreated handles this case with default header values.

Request completed successfully.
*/
type VaultCreateNewVaultItemCreated struct {
	Payload *models.VaultItem
}

func (o *VaultCreateNewVaultItemCreated) Error() string {
	return fmt.Sprintf("[POST /VaultItem][%d] vaultCreateNewVaultItemCreated  %+v", 201, o.Payload)
}

func (o *VaultCreateNewVaultItemCreated) GetPayload() *models.VaultItem {
	return o.Payload
}

func (o *VaultCreateNewVaultItemCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VaultItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVaultCreateNewVaultItemBadRequest creates a VaultCreateNewVaultItemBadRequest with default headers values
func NewVaultCreateNewVaultItemBadRequest() *VaultCreateNewVaultItemBadRequest {
	return &VaultCreateNewVaultItemBadRequest{}
}

/*VaultCreateNewVaultItemBadRequest handles this case with default header values.

The request failed.
*/
type VaultCreateNewVaultItemBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *VaultCreateNewVaultItemBadRequest) Error() string {
	return fmt.Sprintf("[POST /VaultItem][%d] vaultCreateNewVaultItemBadRequest  %+v", 400, o.Payload)
}

func (o *VaultCreateNewVaultItemBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *VaultCreateNewVaultItemBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}