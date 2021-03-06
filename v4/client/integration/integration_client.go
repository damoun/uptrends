// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new integration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for integration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	IntegrationGetAllIntegrations(params *IntegrationGetAllIntegrationsParams, authInfo runtime.ClientAuthInfoWriter) (*IntegrationGetAllIntegrationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  IntegrationGetAllIntegrations gets a list of all integrations
*/
func (a *Client) IntegrationGetAllIntegrations(params *IntegrationGetAllIntegrationsParams, authInfo runtime.ClientAuthInfoWriter) (*IntegrationGetAllIntegrationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIntegrationGetAllIntegrationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Integration_GetAllIntegrations",
		Method:             "GET",
		PathPattern:        "/Integration",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &IntegrationGetAllIntegrationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*IntegrationGetAllIntegrationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Integration_GetAllIntegrations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
