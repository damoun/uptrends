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

// NewStatusGetMonitorStatusParams creates a new StatusGetMonitorStatusParams object
// with the default values initialized.
func NewStatusGetMonitorStatusParams() *StatusGetMonitorStatusParams {
	var ()
	return &StatusGetMonitorStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStatusGetMonitorStatusParamsWithTimeout creates a new StatusGetMonitorStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStatusGetMonitorStatusParamsWithTimeout(timeout time.Duration) *StatusGetMonitorStatusParams {
	var ()
	return &StatusGetMonitorStatusParams{

		timeout: timeout,
	}
}

// NewStatusGetMonitorStatusParamsWithContext creates a new StatusGetMonitorStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewStatusGetMonitorStatusParamsWithContext(ctx context.Context) *StatusGetMonitorStatusParams {
	var ()
	return &StatusGetMonitorStatusParams{

		Context: ctx,
	}
}

// NewStatusGetMonitorStatusParamsWithHTTPClient creates a new StatusGetMonitorStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStatusGetMonitorStatusParamsWithHTTPClient(client *http.Client) *StatusGetMonitorStatusParams {
	var ()
	return &StatusGetMonitorStatusParams{
		HTTPClient: client,
	}
}

/*StatusGetMonitorStatusParams contains all the parameters to send to the API endpoint
for the status get monitor status operation typically these are written to a http.Request
*/
type StatusGetMonitorStatusParams struct {

	/*MonitorGUID
	  The Guid of the monitor.

	*/
	MonitorGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the status get monitor status params
func (o *StatusGetMonitorStatusParams) WithTimeout(timeout time.Duration) *StatusGetMonitorStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the status get monitor status params
func (o *StatusGetMonitorStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the status get monitor status params
func (o *StatusGetMonitorStatusParams) WithContext(ctx context.Context) *StatusGetMonitorStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the status get monitor status params
func (o *StatusGetMonitorStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the status get monitor status params
func (o *StatusGetMonitorStatusParams) WithHTTPClient(client *http.Client) *StatusGetMonitorStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the status get monitor status params
func (o *StatusGetMonitorStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitorGUID adds the monitorGUID to the status get monitor status params
func (o *StatusGetMonitorStatusParams) WithMonitorGUID(monitorGUID string) *StatusGetMonitorStatusParams {
	o.SetMonitorGUID(monitorGUID)
	return o
}

// SetMonitorGUID adds the monitorGuid to the status get monitor status params
func (o *StatusGetMonitorStatusParams) SetMonitorGUID(monitorGUID string) {
	o.MonitorGUID = monitorGUID
}

// WriteToRequest writes these params to a swagger request
func (o *StatusGetMonitorStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param monitorGuid
	if err := r.SetPathParam("monitorGuid", o.MonitorGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
