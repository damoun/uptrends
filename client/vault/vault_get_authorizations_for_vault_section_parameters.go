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

	strfmt "github.com/go-openapi/strfmt"
)

// NewVaultGetAuthorizationsForVaultSectionParams creates a new VaultGetAuthorizationsForVaultSectionParams object
// with the default values initialized.
func NewVaultGetAuthorizationsForVaultSectionParams() *VaultGetAuthorizationsForVaultSectionParams {
	var ()
	return &VaultGetAuthorizationsForVaultSectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVaultGetAuthorizationsForVaultSectionParamsWithTimeout creates a new VaultGetAuthorizationsForVaultSectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVaultGetAuthorizationsForVaultSectionParamsWithTimeout(timeout time.Duration) *VaultGetAuthorizationsForVaultSectionParams {
	var ()
	return &VaultGetAuthorizationsForVaultSectionParams{

		timeout: timeout,
	}
}

// NewVaultGetAuthorizationsForVaultSectionParamsWithContext creates a new VaultGetAuthorizationsForVaultSectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewVaultGetAuthorizationsForVaultSectionParamsWithContext(ctx context.Context) *VaultGetAuthorizationsForVaultSectionParams {
	var ()
	return &VaultGetAuthorizationsForVaultSectionParams{

		Context: ctx,
	}
}

// NewVaultGetAuthorizationsForVaultSectionParamsWithHTTPClient creates a new VaultGetAuthorizationsForVaultSectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVaultGetAuthorizationsForVaultSectionParamsWithHTTPClient(client *http.Client) *VaultGetAuthorizationsForVaultSectionParams {
	var ()
	return &VaultGetAuthorizationsForVaultSectionParams{
		HTTPClient: client,
	}
}

/*VaultGetAuthorizationsForVaultSectionParams contains all the parameters to send to the API endpoint
for the vault get authorizations for vault section operation typically these are written to a http.Request
*/
type VaultGetAuthorizationsForVaultSectionParams struct {

	/*VaultSectionGUID
	  The Guid of the vault section for which to return authorizations.

	*/
	VaultSectionGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the vault get authorizations for vault section params
func (o *VaultGetAuthorizationsForVaultSectionParams) WithTimeout(timeout time.Duration) *VaultGetAuthorizationsForVaultSectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vault get authorizations for vault section params
func (o *VaultGetAuthorizationsForVaultSectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vault get authorizations for vault section params
func (o *VaultGetAuthorizationsForVaultSectionParams) WithContext(ctx context.Context) *VaultGetAuthorizationsForVaultSectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vault get authorizations for vault section params
func (o *VaultGetAuthorizationsForVaultSectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vault get authorizations for vault section params
func (o *VaultGetAuthorizationsForVaultSectionParams) WithHTTPClient(client *http.Client) *VaultGetAuthorizationsForVaultSectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vault get authorizations for vault section params
func (o *VaultGetAuthorizationsForVaultSectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithVaultSectionGUID adds the vaultSectionGUID to the vault get authorizations for vault section params
func (o *VaultGetAuthorizationsForVaultSectionParams) WithVaultSectionGUID(vaultSectionGUID string) *VaultGetAuthorizationsForVaultSectionParams {
	o.SetVaultSectionGUID(vaultSectionGUID)
	return o
}

// SetVaultSectionGUID adds the vaultSectionGuid to the vault get authorizations for vault section params
func (o *VaultGetAuthorizationsForVaultSectionParams) SetVaultSectionGUID(vaultSectionGUID string) {
	o.VaultSectionGUID = vaultSectionGUID
}

// WriteToRequest writes these params to a swagger request
func (o *VaultGetAuthorizationsForVaultSectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param vaultSectionGuid
	if err := r.SetPathParam("vaultSectionGuid", o.VaultSectionGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
