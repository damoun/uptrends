// Code generated by go-swagger; DO NOT EDIT.

package register

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new register API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for register API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
RegisterPost creates a new API account

This method requires that you specify the username and password of an Uptrends operator login as authentication. When registration is successful, please save the UserName and Password fields from the response: these are your new API credentials.
*/
func (a *Client) RegisterPost(params *RegisterPostParams, authInfo runtime.ClientAuthInfoWriter) (*RegisterPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Register_Post",
		Method:             "POST",
		PathPattern:        "/Register",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RegisterPostReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RegisterPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Register_Post: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}