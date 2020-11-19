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

// VaultGetAllVaultSectionsReader is a Reader for the VaultGetAllVaultSections structure.
type VaultGetAllVaultSectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VaultGetAllVaultSectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVaultGetAllVaultSectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVaultGetAllVaultSectionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVaultGetAllVaultSectionsOK creates a VaultGetAllVaultSectionsOK with default headers values
func NewVaultGetAllVaultSectionsOK() *VaultGetAllVaultSectionsOK {
	return &VaultGetAllVaultSectionsOK{}
}

/*VaultGetAllVaultSectionsOK handles this case with default header values.

The request completed successfully.
*/
type VaultGetAllVaultSectionsOK struct {
	Payload []*models.VaultSection
}

func (o *VaultGetAllVaultSectionsOK) Error() string {
	return fmt.Sprintf("[GET /VaultSection][%d] vaultGetAllVaultSectionsOK  %+v", 200, o.Payload)
}

func (o *VaultGetAllVaultSectionsOK) GetPayload() []*models.VaultSection {
	return o.Payload
}

func (o *VaultGetAllVaultSectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVaultGetAllVaultSectionsBadRequest creates a VaultGetAllVaultSectionsBadRequest with default headers values
func NewVaultGetAllVaultSectionsBadRequest() *VaultGetAllVaultSectionsBadRequest {
	return &VaultGetAllVaultSectionsBadRequest{}
}

/*VaultGetAllVaultSectionsBadRequest handles this case with default header values.

The request failed.
*/
type VaultGetAllVaultSectionsBadRequest struct {
	Payload *models.MessageList
}

func (o *VaultGetAllVaultSectionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /VaultSection][%d] vaultGetAllVaultSectionsBadRequest  %+v", 400, o.Payload)
}

func (o *VaultGetAllVaultSectionsBadRequest) GetPayload() *models.MessageList {
	return o.Payload
}

func (o *VaultGetAllVaultSectionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MessageList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
