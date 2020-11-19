// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dashboard API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dashboard API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DashboardCloneDashboard(params *DashboardCloneDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardCloneDashboardCreated, error)

	DashboardDeleteDashboard(params *DashboardDeleteDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardDeleteDashboardNoContent, error)

	DashboardGetAllDashboards(params *DashboardGetAllDashboardsParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardGetAllDashboardsOK, error)

	DashboardGetOneDashboard(params *DashboardGetOneDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardGetOneDashboardOK, error)

	DashboardPartiallyUpdateDashboard(params *DashboardPartiallyUpdateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardPartiallyUpdateDashboardNoContent, error)

	DashboardUpdateDashboard(params *DashboardUpdateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardUpdateDashboardNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DashboardCloneDashboard clones the specified dashboard
*/
func (a *Client) DashboardCloneDashboard(params *DashboardCloneDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardCloneDashboardCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDashboardCloneDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Dashboard_CloneDashboard",
		Method:             "POST",
		PathPattern:        "/Dashboard/{dashboardGuid}/Clone",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DashboardCloneDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DashboardCloneDashboardCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Dashboard_CloneDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DashboardDeleteDashboard deletes the specified dashboard
*/
func (a *Client) DashboardDeleteDashboard(params *DashboardDeleteDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardDeleteDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDashboardDeleteDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Dashboard_DeleteDashboard",
		Method:             "DELETE",
		PathPattern:        "/Dashboard/{dashboardGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DashboardDeleteDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DashboardDeleteDashboardNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Dashboard_DeleteDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DashboardGetAllDashboards returns data for all dashboards
*/
func (a *Client) DashboardGetAllDashboards(params *DashboardGetAllDashboardsParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardGetAllDashboardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDashboardGetAllDashboardsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Dashboard_GetAllDashboards",
		Method:             "GET",
		PathPattern:        "/Dashboard",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DashboardGetAllDashboardsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DashboardGetAllDashboardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Dashboard_GetAllDashboards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DashboardGetOneDashboard returns data for the specified dashboard
*/
func (a *Client) DashboardGetOneDashboard(params *DashboardGetOneDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardGetOneDashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDashboardGetOneDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Dashboard_GetOneDashboard",
		Method:             "GET",
		PathPattern:        "/Dashboard/{dashboardGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DashboardGetOneDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DashboardGetOneDashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Dashboard_GetOneDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DashboardPartiallyUpdateDashboard partiallies update the specified dashboard
*/
func (a *Client) DashboardPartiallyUpdateDashboard(params *DashboardPartiallyUpdateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardPartiallyUpdateDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDashboardPartiallyUpdateDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Dashboard_PartiallyUpdateDashboard",
		Method:             "PATCH",
		PathPattern:        "/Dashboard/{dashboardGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DashboardPartiallyUpdateDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DashboardPartiallyUpdateDashboardNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Dashboard_PartiallyUpdateDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DashboardUpdateDashboard updates the specified dashboard
*/
func (a *Client) DashboardUpdateDashboard(params *DashboardUpdateDashboardParams, authInfo runtime.ClientAuthInfoWriter) (*DashboardUpdateDashboardNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDashboardUpdateDashboardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Dashboard_UpdateDashboard",
		Method:             "PUT",
		PathPattern:        "/Dashboard/{dashboardGuid}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DashboardUpdateDashboardReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DashboardUpdateDashboardNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Dashboard_UpdateDashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
