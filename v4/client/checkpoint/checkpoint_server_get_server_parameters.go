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
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCheckpointServerGetServerParams creates a new CheckpointServerGetServerParams object
// with the default values initialized.
func NewCheckpointServerGetServerParams() *CheckpointServerGetServerParams {
	var ()
	return &CheckpointServerGetServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckpointServerGetServerParamsWithTimeout creates a new CheckpointServerGetServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckpointServerGetServerParamsWithTimeout(timeout time.Duration) *CheckpointServerGetServerParams {
	var ()
	return &CheckpointServerGetServerParams{

		timeout: timeout,
	}
}

// NewCheckpointServerGetServerParamsWithContext creates a new CheckpointServerGetServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckpointServerGetServerParamsWithContext(ctx context.Context) *CheckpointServerGetServerParams {
	var ()
	return &CheckpointServerGetServerParams{

		Context: ctx,
	}
}

// NewCheckpointServerGetServerParamsWithHTTPClient creates a new CheckpointServerGetServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckpointServerGetServerParamsWithHTTPClient(client *http.Client) *CheckpointServerGetServerParams {
	var ()
	return &CheckpointServerGetServerParams{
		HTTPClient: client,
	}
}

/*CheckpointServerGetServerParams contains all the parameters to send to the API endpoint
for the checkpoint server get server operation typically these are written to a http.Request
*/
type CheckpointServerGetServerParams struct {

	/*ServerID
	  The Id of the requested server.

	*/
	ServerID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the checkpoint server get server params
func (o *CheckpointServerGetServerParams) WithTimeout(timeout time.Duration) *CheckpointServerGetServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the checkpoint server get server params
func (o *CheckpointServerGetServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the checkpoint server get server params
func (o *CheckpointServerGetServerParams) WithContext(ctx context.Context) *CheckpointServerGetServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the checkpoint server get server params
func (o *CheckpointServerGetServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the checkpoint server get server params
func (o *CheckpointServerGetServerParams) WithHTTPClient(client *http.Client) *CheckpointServerGetServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the checkpoint server get server params
func (o *CheckpointServerGetServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServerID adds the serverID to the checkpoint server get server params
func (o *CheckpointServerGetServerParams) WithServerID(serverID int64) *CheckpointServerGetServerParams {
	o.SetServerID(serverID)
	return o
}

// SetServerID adds the serverId to the checkpoint server get server params
func (o *CheckpointServerGetServerParams) SetServerID(serverID int64) {
	o.ServerID = serverID
}

// WriteToRequest writes these params to a swagger request
func (o *CheckpointServerGetServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param serverId
	if err := r.SetPathParam("serverId", swag.FormatInt64(o.ServerID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
