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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewMonitorCheckGetHTTPDetailsParams creates a new MonitorCheckGetHTTPDetailsParams object
// with the default values initialized.
func NewMonitorCheckGetHTTPDetailsParams() *MonitorCheckGetHTTPDetailsParams {
	var ()
	return &MonitorCheckGetHTTPDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorCheckGetHTTPDetailsParamsWithTimeout creates a new MonitorCheckGetHTTPDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorCheckGetHTTPDetailsParamsWithTimeout(timeout time.Duration) *MonitorCheckGetHTTPDetailsParams {
	var ()
	return &MonitorCheckGetHTTPDetailsParams{

		timeout: timeout,
	}
}

// NewMonitorCheckGetHTTPDetailsParamsWithContext creates a new MonitorCheckGetHTTPDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorCheckGetHTTPDetailsParamsWithContext(ctx context.Context) *MonitorCheckGetHTTPDetailsParams {
	var ()
	return &MonitorCheckGetHTTPDetailsParams{

		Context: ctx,
	}
}

// NewMonitorCheckGetHTTPDetailsParamsWithHTTPClient creates a new MonitorCheckGetHTTPDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorCheckGetHTTPDetailsParamsWithHTTPClient(client *http.Client) *MonitorCheckGetHTTPDetailsParams {
	var ()
	return &MonitorCheckGetHTTPDetailsParams{
		HTTPClient: client,
	}
}

/*MonitorCheckGetHTTPDetailsParams contains all the parameters to send to the API endpoint
for the monitor check get Http details operation typically these are written to a http.Request
*/
type MonitorCheckGetHTTPDetailsParams struct {

	/*MonitorCheckID
	  Monitor check Id to get the detailed data for.

	*/
	MonitorCheckID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the monitor check get Http details params
func (o *MonitorCheckGetHTTPDetailsParams) WithTimeout(timeout time.Duration) *MonitorCheckGetHTTPDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor check get Http details params
func (o *MonitorCheckGetHTTPDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor check get Http details params
func (o *MonitorCheckGetHTTPDetailsParams) WithContext(ctx context.Context) *MonitorCheckGetHTTPDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor check get Http details params
func (o *MonitorCheckGetHTTPDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor check get Http details params
func (o *MonitorCheckGetHTTPDetailsParams) WithHTTPClient(client *http.Client) *MonitorCheckGetHTTPDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor check get Http details params
func (o *MonitorCheckGetHTTPDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitorCheckID adds the monitorCheckID to the monitor check get Http details params
func (o *MonitorCheckGetHTTPDetailsParams) WithMonitorCheckID(monitorCheckID int64) *MonitorCheckGetHTTPDetailsParams {
	o.SetMonitorCheckID(monitorCheckID)
	return o
}

// SetMonitorCheckID adds the monitorCheckId to the monitor check get Http details params
func (o *MonitorCheckGetHTTPDetailsParams) SetMonitorCheckID(monitorCheckID int64) {
	o.MonitorCheckID = monitorCheckID
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorCheckGetHTTPDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
