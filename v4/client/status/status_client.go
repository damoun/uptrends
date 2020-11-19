// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new status API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for status API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	StatusGetMonitorGroupStatus(params *StatusGetMonitorGroupStatusParams, authInfo runtime.ClientAuthInfoWriter) (*StatusGetMonitorGroupStatusOK, error)

	StatusGetMonitorStatus(params *StatusGetMonitorStatusParams, authInfo runtime.ClientAuthInfoWriter) (*StatusGetMonitorStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  StatusGetMonitorGroupStatus gets a list of all monitor group status data
*/
func (a *Client) StatusGetMonitorGroupStatus(params *StatusGetMonitorGroupStatusParams, authInfo runtime.ClientAuthInfoWriter) (*StatusGetMonitorGroupStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusGetMonitorGroupStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Status_GetMonitorGroupStatus",
		Method:             "GET",
		PathPattern:        "/Status/MonitorGroup/{monitorGroupGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatusGetMonitorGroupStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatusGetMonitorGroupStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Status_GetMonitorGroupStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StatusGetMonitorStatus gets all monitor status data
*/
func (a *Client) StatusGetMonitorStatus(params *StatusGetMonitorStatusParams, authInfo runtime.ClientAuthInfoWriter) (*StatusGetMonitorStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusGetMonitorStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Status_GetMonitorStatus",
		Method:             "GET",
		PathPattern:        "/Status/Monitor/{monitorGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StatusGetMonitorStatusReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StatusGetMonitorStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Status_GetMonitorStatus: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
