// Code generated by go-swagger; DO NOT EDIT.

package vault

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

// NewVaultGetAllVaultItemsParams creates a new VaultGetAllVaultItemsParams object
// with the default values initialized.
func NewVaultGetAllVaultItemsParams() *VaultGetAllVaultItemsParams {

	return &VaultGetAllVaultItemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVaultGetAllVaultItemsParamsWithTimeout creates a new VaultGetAllVaultItemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVaultGetAllVaultItemsParamsWithTimeout(timeout time.Duration) *VaultGetAllVaultItemsParams {

	return &VaultGetAllVaultItemsParams{

		timeout: timeout,
	}
}

// NewVaultGetAllVaultItemsParamsWithContext creates a new VaultGetAllVaultItemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewVaultGetAllVaultItemsParamsWithContext(ctx context.Context) *VaultGetAllVaultItemsParams {

	return &VaultGetAllVaultItemsParams{

		Context: ctx,
	}
}

// NewVaultGetAllVaultItemsParamsWithHTTPClient creates a new VaultGetAllVaultItemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVaultGetAllVaultItemsParamsWithHTTPClient(client *http.Client) *VaultGetAllVaultItemsParams {

	return &VaultGetAllVaultItemsParams{
		HTTPClient: client,
	}
}

/*VaultGetAllVaultItemsParams contains all the parameters to send to the API endpoint
for the vault get all vault items operation typically these are written to a http.Request
*/
type VaultGetAllVaultItemsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the vault get all vault items params
func (o *VaultGetAllVaultItemsParams) WithTimeout(timeout time.Duration) *VaultGetAllVaultItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vault get all vault items params
func (o *VaultGetAllVaultItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vault get all vault items params
func (o *VaultGetAllVaultItemsParams) WithContext(ctx context.Context) *VaultGetAllVaultItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vault get all vault items params
func (o *VaultGetAllVaultItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vault get all vault items params
func (o *VaultGetAllVaultItemsParams) WithHTTPClient(client *http.Client) *VaultGetAllVaultItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vault get all vault items params
func (o *VaultGetAllVaultItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *VaultGetAllVaultItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
