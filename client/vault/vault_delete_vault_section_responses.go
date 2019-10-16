// Code generated by go-swagger; DO NOT EDIT.

package vault

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/models"
)

// VaultDeleteVaultSectionReader is a Reader for the VaultDeleteVaultSection structure.
type VaultDeleteVaultSectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VaultDeleteVaultSectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVaultDeleteVaultSectionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVaultDeleteVaultSectionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewVaultDeleteVaultSectionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVaultDeleteVaultSectionNoContent creates a VaultDeleteVaultSectionNoContent with default headers values
func NewVaultDeleteVaultSectionNoContent() *VaultDeleteVaultSectionNoContent {
	return &VaultDeleteVaultSectionNoContent{}
}

/*VaultDeleteVaultSectionNoContent handles this case with default header values.

Request completed successfully.
*/
type VaultDeleteVaultSectionNoContent struct {
	Payload *models.VaultSection
}

func (o *VaultDeleteVaultSectionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /VaultSection/{vaultSectionGuid}][%d] vaultDeleteVaultSectionNoContent  %+v", 204, o.Payload)
}

func (o *VaultDeleteVaultSectionNoContent) GetPayload() *models.VaultSection {
	return o.Payload
}

func (o *VaultDeleteVaultSectionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VaultSection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVaultDeleteVaultSectionBadRequest creates a VaultDeleteVaultSectionBadRequest with default headers values
func NewVaultDeleteVaultSectionBadRequest() *VaultDeleteVaultSectionBadRequest {
	return &VaultDeleteVaultSectionBadRequest{}
}

/*VaultDeleteVaultSectionBadRequest handles this case with default header values.

The request failed.
*/
type VaultDeleteVaultSectionBadRequest struct {
	Payload *models.APIMessageInfo
}

func (o *VaultDeleteVaultSectionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /VaultSection/{vaultSectionGuid}][%d] vaultDeleteVaultSectionBadRequest  %+v", 400, o.Payload)
}

func (o *VaultDeleteVaultSectionBadRequest) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *VaultDeleteVaultSectionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVaultDeleteVaultSectionNotFound creates a VaultDeleteVaultSectionNotFound with default headers values
func NewVaultDeleteVaultSectionNotFound() *VaultDeleteVaultSectionNotFound {
	return &VaultDeleteVaultSectionNotFound{}
}

/*VaultDeleteVaultSectionNotFound handles this case with default header values.

The requested vault section does not exist.
*/
type VaultDeleteVaultSectionNotFound struct {
	Payload *models.APIMessageInfo
}

func (o *VaultDeleteVaultSectionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /VaultSection/{vaultSectionGuid}][%d] vaultDeleteVaultSectionNotFound  %+v", 404, o.Payload)
}

func (o *VaultDeleteVaultSectionNotFound) GetPayload() *models.APIMessageInfo {
	return o.Payload
}

func (o *VaultDeleteVaultSectionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMessageInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
