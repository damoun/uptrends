// Code generated by go-swagger; DO NOT EDIT.

package checkpoint

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

// NewCheckpointServerGetAllServerIpv6sParams creates a new CheckpointServerGetAllServerIpv6sParams object
// with the default values initialized.
func NewCheckpointServerGetAllServerIpv6sParams() *CheckpointServerGetAllServerIpv6sParams {

	return &CheckpointServerGetAllServerIpv6sParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckpointServerGetAllServerIpv6sParamsWithTimeout creates a new CheckpointServerGetAllServerIpv6sParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckpointServerGetAllServerIpv6sParamsWithTimeout(timeout time.Duration) *CheckpointServerGetAllServerIpv6sParams {

	return &CheckpointServerGetAllServerIpv6sParams{

		timeout: timeout,
	}
}

// NewCheckpointServerGetAllServerIpv6sParamsWithContext creates a new CheckpointServerGetAllServerIpv6sParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckpointServerGetAllServerIpv6sParamsWithContext(ctx context.Context) *CheckpointServerGetAllServerIpv6sParams {

	return &CheckpointServerGetAllServerIpv6sParams{

		Context: ctx,
	}
}

// NewCheckpointServerGetAllServerIpv6sParamsWithHTTPClient creates a new CheckpointServerGetAllServerIpv6sParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckpointServerGetAllServerIpv6sParamsWithHTTPClient(client *http.Client) *CheckpointServerGetAllServerIpv6sParams {

	return &CheckpointServerGetAllServerIpv6sParams{
		HTTPClient: client,
	}
}

/*CheckpointServerGetAllServerIpv6sParams contains all the parameters to send to the API endpoint
for the checkpoint server get all server ipv6s operation typically these are written to a http.Request
*/
type CheckpointServerGetAllServerIpv6sParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkpoint server get all server ipv6s params
func (o *CheckpointServerGetAllServerIpv6sParams) WithTimeout(timeout time.Duration) *CheckpointServerGetAllServerIpv6sParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkpoint server get all server ipv6s params
func (o *CheckpointServerGetAllServerIpv6sParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkpoint server get all server ipv6s params
func (o *CheckpointServerGetAllServerIpv6sParams) WithContext(ctx context.Context) *CheckpointServerGetAllServerIpv6sParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkpoint server get all server ipv6s params
func (o *CheckpointServerGetAllServerIpv6sParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkpoint server get all server ipv6s params
func (o *CheckpointServerGetAllServerIpv6sParams) WithHTTPClient(client *http.Client) *CheckpointServerGetAllServerIpv6sParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkpoint server get all server ipv6s params
func (o *CheckpointServerGetAllServerIpv6sParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CheckpointServerGetAllServerIpv6sParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
