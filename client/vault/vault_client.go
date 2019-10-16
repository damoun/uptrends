// Code generated by go-swagger; DO NOT EDIT.

package vault

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new vault API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for vault API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
VaultCreateAuthorizationForVaultSection creates a new authorization for the specified vault section

The AuthorizationId attribute should be omitted in the request body. The AuthorizationId of the newly created authorization will be returned in the response. In the ContextID attribute, fill in the VaultSectionGuid that identifies the vault section for which to create the new authorization. Valid values for the AuthorizationType field are "ViewVaultSection" and "ChangeVaultSection". An authorization should be granted to either an individual operator, or an operator group. Therefore, either specify the OperatorGuid attribute or the OperatorGroupGuid attribute.
*/
func (a *Client) VaultCreateAuthorizationForVaultSection(params *VaultCreateAuthorizationForVaultSectionParams, authInfo runtime.ClientAuthInfoWriter) (*VaultCreateAuthorizationForVaultSectionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultCreateAuthorizationForVaultSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_CreateAuthorizationForVaultSection",
		Method:             "POST",
		PathPattern:        "/VaultSection/{vaultSectionGuid}/Authorization",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultCreateAuthorizationForVaultSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultCreateAuthorizationForVaultSectionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_CreateAuthorizationForVaultSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultCreateNewVaultItem creates a new vault item

The VaultItemGuid field should be omitted
*/
func (a *Client) VaultCreateNewVaultItem(params *VaultCreateNewVaultItemParams, authInfo runtime.ClientAuthInfoWriter) (*VaultCreateNewVaultItemCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultCreateNewVaultItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_CreateNewVaultItem",
		Method:             "POST",
		PathPattern:        "/VaultItem",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultCreateNewVaultItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultCreateNewVaultItemCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_CreateNewVaultItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultCreateNewVaultSection creates a new vault section

When a new vault section is created, the user that created the section is granted View and Edit authorizations to that section. The VaultSectionGuid attribute should be omitted in the request body. The Guid of the newly created section will be returned in the response.
*/
func (a *Client) VaultCreateNewVaultSection(params *VaultCreateNewVaultSectionParams, authInfo runtime.ClientAuthInfoWriter) (*VaultCreateNewVaultSectionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultCreateNewVaultSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_CreateNewVaultSection",
		Method:             "POST",
		PathPattern:        "/VaultSection",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultCreateNewVaultSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultCreateNewVaultSectionCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_CreateNewVaultSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultDeleteAuthorizationForVaultSection deletes the specified authorization for the specified vault section
*/
func (a *Client) VaultDeleteAuthorizationForVaultSection(params *VaultDeleteAuthorizationForVaultSectionParams, authInfo runtime.ClientAuthInfoWriter) (*VaultDeleteAuthorizationForVaultSectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultDeleteAuthorizationForVaultSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_DeleteAuthorizationForVaultSection",
		Method:             "DELETE",
		PathPattern:        "/VaultSection/{vaultSectionGuid}/Authorization/{authorizationGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultDeleteAuthorizationForVaultSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultDeleteAuthorizationForVaultSectionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_DeleteAuthorizationForVaultSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultDeleteVaultItem deletes the specified vault item
*/
func (a *Client) VaultDeleteVaultItem(params *VaultDeleteVaultItemParams, authInfo runtime.ClientAuthInfoWriter) (*VaultDeleteVaultItemNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultDeleteVaultItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_DeleteVaultItem",
		Method:             "DELETE",
		PathPattern:        "/VaultItem/{vaultItemGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultDeleteVaultItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultDeleteVaultItemNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_DeleteVaultItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultDeleteVaultSection deletes the specified vault section
*/
func (a *Client) VaultDeleteVaultSection(params *VaultDeleteVaultSectionParams, authInfo runtime.ClientAuthInfoWriter) (*VaultDeleteVaultSectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultDeleteVaultSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_DeleteVaultSection",
		Method:             "DELETE",
		PathPattern:        "/VaultSection/{vaultSectionGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultDeleteVaultSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultDeleteVaultSectionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_DeleteVaultSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultGetAllVaultItems returns all vault items
*/
func (a *Client) VaultGetAllVaultItems(params *VaultGetAllVaultItemsParams, authInfo runtime.ClientAuthInfoWriter) (*VaultGetAllVaultItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultGetAllVaultItemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_GetAllVaultItems",
		Method:             "GET",
		PathPattern:        "/VaultItem",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultGetAllVaultItemsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultGetAllVaultItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_GetAllVaultItems: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultGetAllVaultSections returns all vault sections
*/
func (a *Client) VaultGetAllVaultSections(params *VaultGetAllVaultSectionsParams, authInfo runtime.ClientAuthInfoWriter) (*VaultGetAllVaultSectionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultGetAllVaultSectionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_GetAllVaultSections",
		Method:             "GET",
		PathPattern:        "/VaultSection",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultGetAllVaultSectionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultGetAllVaultSectionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_GetAllVaultSections: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultGetAuthorizationsForVaultSection returns all authorizations for the specified vault section
*/
func (a *Client) VaultGetAuthorizationsForVaultSection(params *VaultGetAuthorizationsForVaultSectionParams, authInfo runtime.ClientAuthInfoWriter) (*VaultGetAuthorizationsForVaultSectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultGetAuthorizationsForVaultSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_GetAuthorizationsForVaultSection",
		Method:             "GET",
		PathPattern:        "/VaultSection/{vaultSectionGuid}/Authorization",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultGetAuthorizationsForVaultSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultGetAuthorizationsForVaultSectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_GetAuthorizationsForVaultSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultGetVaultItem returns the specified vault item
*/
func (a *Client) VaultGetVaultItem(params *VaultGetVaultItemParams, authInfo runtime.ClientAuthInfoWriter) (*VaultGetVaultItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultGetVaultItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_GetVaultItem",
		Method:             "GET",
		PathPattern:        "/VaultItem/{vaultItemGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultGetVaultItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultGetVaultItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_GetVaultItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultGetVaultSection returns the specified vault section
*/
func (a *Client) VaultGetVaultSection(params *VaultGetVaultSectionParams, authInfo runtime.ClientAuthInfoWriter) (*VaultGetVaultSectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultGetVaultSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_GetVaultSection",
		Method:             "GET",
		PathPattern:        "/VaultSection/{vaultSectionGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultGetVaultSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultGetVaultSectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_GetVaultSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultUpdateVaultItem updates the specified vault item

Only complete definitions are accepted. Fields not specified will be NULLed.
*/
func (a *Client) VaultUpdateVaultItem(params *VaultUpdateVaultItemParams, authInfo runtime.ClientAuthInfoWriter) (*VaultUpdateVaultItemNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultUpdateVaultItemParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_UpdateVaultItem",
		Method:             "PUT",
		PathPattern:        "/VaultItem/{vaultItemGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultUpdateVaultItemReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultUpdateVaultItemNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_UpdateVaultItem: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VaultUpdateVaultSection updates the specified vault section
*/
func (a *Client) VaultUpdateVaultSection(params *VaultUpdateVaultSectionParams, authInfo runtime.ClientAuthInfoWriter) (*VaultUpdateVaultSectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVaultUpdateVaultSectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Vault_UpdateVaultSection",
		Method:             "PUT",
		PathPattern:        "/VaultSection/{vaultSectionGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &VaultUpdateVaultSectionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VaultUpdateVaultSectionNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Vault_UpdateVaultSection: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}