// Code generated by go-swagger; DO NOT EDIT.

package monitor_check

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

// NewMonitorCheckGetMultistepDetailsParams creates a new MonitorCheckGetMultistepDetailsParams object
// with the default values initialized.
func NewMonitorCheckGetMultistepDetailsParams() *MonitorCheckGetMultistepDetailsParams {
	var ()
	return &MonitorCheckGetMultistepDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorCheckGetMultistepDetailsParamsWithTimeout creates a new MonitorCheckGetMultistepDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorCheckGetMultistepDetailsParamsWithTimeout(timeout time.Duration) *MonitorCheckGetMultistepDetailsParams {
	var ()
	return &MonitorCheckGetMultistepDetailsParams{

		timeout: timeout,
	}
}

// NewMonitorCheckGetMultistepDetailsParamsWithContext creates a new MonitorCheckGetMultistepDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorCheckGetMultistepDetailsParamsWithContext(ctx context.Context) *MonitorCheckGetMultistepDetailsParams {
	var ()
	return &MonitorCheckGetMultistepDetailsParams{

		Context: ctx,
	}
}

// NewMonitorCheckGetMultistepDetailsParamsWithHTTPClient creates a new MonitorCheckGetMultistepDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorCheckGetMultistepDetailsParamsWithHTTPClient(client *http.Client) *MonitorCheckGetMultistepDetailsParams {
	var ()
	return &MonitorCheckGetMultistepDetailsParams{
		HTTPClient: client,
	}
}

/*MonitorCheckGetMultistepDetailsParams contains all the parameters to send to the API endpoint
for the monitor check get multistep details operation typically these are written to a http.Request
*/
type MonitorCheckGetMultistepDetailsParams struct {

	/*MonitorCheckID
	  The monitor check Id to get the detailed data for.

	*/
	MonitorCheckID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the monitor check get multistep details params
func (o *MonitorCheckGetMultistepDetailsParams) WithTimeout(timeout time.Duration) *MonitorCheckGetMultistepDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor check get multistep details params
func (o *MonitorCheckGetMultistepDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor check get multistep details params
func (o *MonitorCheckGetMultistepDetailsParams) WithContext(ctx context.Context) *MonitorCheckGetMultistepDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor check get multistep details params
func (o *MonitorCheckGetMultistepDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor check get multistep details params
func (o *MonitorCheckGetMultistepDetailsParams) WithHTTPClient(client *http.Client) *MonitorCheckGetMultistepDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor check get multistep details params
func (o *MonitorCheckGetMultistepDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitorCheckID adds the monitorCheckID to the monitor check get multistep details params
func (o *MonitorCheckGetMultistepDetailsParams) WithMonitorCheckID(monitorCheckID int64) *MonitorCheckGetMultistepDetailsParams {
	o.SetMonitorCheckID(monitorCheckID)
	return o
}

// SetMonitorCheckID adds the monitorCheckId to the monitor check get multistep details params
func (o *MonitorCheckGetMultistepDetailsParams) SetMonitorCheckID(monitorCheckID int64) {
	o.MonitorCheckID = monitorCheckID
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorCheckGetMultistepDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param monitorCheckId
	if err := r.SetPathParam("monitorCheckId", swag.FormatInt64(o.MonitorCheckID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
