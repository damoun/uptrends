// Code generated by go-swagger; DO NOT EDIT.

package integration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewIntegrationGetAllIntegrationsParams creates a new IntegrationGetAllIntegrationsParams object
// with the default values initialized.
func NewIntegrationGetAllIntegrationsParams() *IntegrationGetAllIntegrationsParams {

	return &IntegrationGetAllIntegrationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIntegrationGetAllIntegrationsParamsWithTimeout creates a new IntegrationGetAllIntegrationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIntegrationGetAllIntegrationsParamsWithTimeout(timeout time.Duration) *IntegrationGetAllIntegrationsParams {

	return &IntegrationGetAllIntegrationsParams{

		timeout: timeout,
	}
}

// NewIntegrationGetAllIntegrationsParamsWithContext creates a new IntegrationGetAllIntegrationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewIntegrationGetAllIntegrationsParamsWithContext(ctx context.Context) *IntegrationGetAllIntegrationsParams {

	return &IntegrationGetAllIntegrationsParams{

		Context: ctx,
	}
}

// NewIntegrationGetAllIntegrationsParamsWithHTTPClient creates a new IntegrationGetAllIntegrationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIntegrationGetAllIntegrationsParamsWithHTTPClient(client *http.Client) *IntegrationGetAllIntegrationsParams {

	return &IntegrationGetAllIntegrationsParams{
		HTTPClient: client,
	}
}

/*IntegrationGetAllIntegrationsParams contains all the parameters to send to the API endpoint
for the integration get all integrations operation typically these are written to a http.Request
*/
type IntegrationGetAllIntegrationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the integration get all integrations params
func (o *IntegrationGetAllIntegrationsParams) WithTimeout(timeout time.Duration) *IntegrationGetAllIntegrationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the integration get all integrations params
func (o *IntegrationGetAllIntegrationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the integration get all integrations params
func (o *IntegrationGetAllIntegrationsParams) WithContext(ctx context.Context) *IntegrationGetAllIntegrationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the integration get all integrations params
func (o *IntegrationGetAllIntegrationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the integration get all integrations params
func (o *IntegrationGetAllIntegrationsParams) WithHTTPClient(client *http.Client) *IntegrationGetAllIntegrationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the integration get all integrations params
func (o *IntegrationGetAllIntegrationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *IntegrationGetAllIntegrationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}