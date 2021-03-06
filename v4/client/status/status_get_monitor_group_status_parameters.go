// Code generated by go-swagger; DO NOT EDIT.

package status

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

// NewStatusGetMonitorGroupStatusParams creates a new StatusGetMonitorGroupStatusParams object
// with the default values initialized.
func NewStatusGetMonitorGroupStatusParams() *StatusGetMonitorGroupStatusParams {
	var ()
	return &StatusGetMonitorGroupStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStatusGetMonitorGroupStatusParamsWithTimeout creates a new StatusGetMonitorGroupStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStatusGetMonitorGroupStatusParamsWithTimeout(timeout time.Duration) *StatusGetMonitorGroupStatusParams {
	var ()
	return &StatusGetMonitorGroupStatusParams{

		timeout: timeout,
	}
}

// NewStatusGetMonitorGroupStatusParamsWithContext creates a new StatusGetMonitorGroupStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewStatusGetMonitorGroupStatusParamsWithContext(ctx context.Context) *StatusGetMonitorGroupStatusParams {
	var ()
	return &StatusGetMonitorGroupStatusParams{

		Context: ctx,
	}
}

// NewStatusGetMonitorGroupStatusParamsWithHTTPClient creates a new StatusGetMonitorGroupStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStatusGetMonitorGroupStatusParamsWithHTTPClient(client *http.Client) *StatusGetMonitorGroupStatusParams {
	var ()
	return &StatusGetMonitorGroupStatusParams{
		HTTPClient: client,
	}
}

/*StatusGetMonitorGroupStatusParams contains all the parameters to send to the API endpoint
for the status get monitor group status operation typically these are written to a http.Request
*/
type StatusGetMonitorGroupStatusParams struct {

	/*MonitorGroupGUID
	  The Guid of the monitor group.

	*/
	MonitorGroupGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the status get monitor group status params
func (o *StatusGetMonitorGroupStatusParams) WithTimeout(timeout time.Duration) *StatusGetMonitorGroupStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status get monitor group status params
func (o *StatusGetMonitorGroupStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status get monitor group status params
func (o *StatusGetMonitorGroupStatusParams) WithContext(ctx context.Context) *StatusGetMonitorGroupStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status get monitor group status params
func (o *StatusGetMonitorGroupStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status get monitor group status params
func (o *StatusGetMonitorGroupStatusParams) WithHTTPClient(client *http.Client) *StatusGetMonitorGroupStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status get monitor group status params
func (o *StatusGetMonitorGroupStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitorGroupGUID adds the monitorGroupGUID to the status get monitor group status params
func (o *StatusGetMonitorGroupStatusParams) WithMonitorGroupGUID(monitorGroupGUID string) *StatusGetMonitorGroupStatusParams {
	o.SetMonitorGroupGUID(monitorGroupGUID)
	return o
}

// SetMonitorGroupGUID adds the monitorGroupGuid to the status get monitor group status params
func (o *StatusGetMonitorGroupStatusParams) SetMonitorGroupGUID(monitorGroupGUID string) {
	o.MonitorGroupGUID = monitorGroupGUID
}

// WriteToRequest writes these params to a swagger request
func (o *StatusGetMonitorGroupStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param monitorGroupGuid
	if err := r.SetPathParam("monitorGroupGuid", o.MonitorGroupGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
