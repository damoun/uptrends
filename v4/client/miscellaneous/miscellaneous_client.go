// Code generated by go-swagger; DO NOT EDIT.

package miscellaneous

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new miscellaneous API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for miscellaneous API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	MiscellaneousGetAllOutgoingPhoneNumbers(params *MiscellaneousGetAllOutgoingPhoneNumbersParams, authInfo runtime.ClientAuthInfoWriter) (*MiscellaneousGetAllOutgoingPhoneNumbersOK, error)

	MiscellaneousGetAllTimezones(params *MiscellaneousGetAllTimezonesParams, authInfo runtime.ClientAuthInfoWriter) (*MiscellaneousGetAllTimezonesOK, error)

	MiscellaneousGetTimezoneByID(params *MiscellaneousGetTimezoneByIDParams, authInfo runtime.ClientAuthInfoWriter) (*MiscellaneousGetTimezoneByIDOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  MiscellaneousGetAllOutgoingPhoneNumbers gets a list of all outgoing phone numbers available
*/
func (a *Client) MiscellaneousGetAllOutgoingPhoneNumbers(params *MiscellaneousGetAllOutgoingPhoneNumbersParams, authInfo runtime.ClientAuthInfoWriter) (*MiscellaneousGetAllOutgoingPhoneNumbersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMiscellaneousGetAllOutgoingPhoneNumbersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Miscellaneous_GetAllOutgoingPhoneNumbers",
		Method:             "GET",
		PathPattern:        "/OutgoingPhoneNumber",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MiscellaneousGetAllOutgoingPhoneNumbersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MiscellaneousGetAllOutgoingPhoneNumbersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Miscellaneous_GetAllOutgoingPhoneNumbers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MiscellaneousGetAllTimezones gets all timezones available
*/
func (a *Client) MiscellaneousGetAllTimezones(params *MiscellaneousGetAllTimezonesParams, authInfo runtime.ClientAuthInfoWriter) (*MiscellaneousGetAllTimezonesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMiscellaneousGetAllTimezonesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Miscellaneous_GetAllTimezones",
		Method:             "GET",
		PathPattern:        "/Timezone",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MiscellaneousGetAllTimezonesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MiscellaneousGetAllTimezonesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Miscellaneous_GetAllTimezones: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MiscellaneousGetTimezoneByID gets the timezone with the specified Id
*/
func (a *Client) MiscellaneousGetTimezoneByID(params *MiscellaneousGetTimezoneByIDParams, authInfo runtime.ClientAuthInfoWriter) (*MiscellaneousGetTimezoneByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMiscellaneousGetTimezoneByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Miscellaneous_GetTimezoneById",
		Method:             "GET",
		PathPattern:        "/Timezone/{timezoneId}",
		ProducesMediaTypes: []string{"application/json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MiscellaneousGetTimezoneByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MiscellaneousGetTimezoneByIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Miscellaneous_GetTimezoneById: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
