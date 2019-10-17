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

// NewMonitorCheckGetSingleMonitorCheckParams creates a new MonitorCheckGetSingleMonitorCheckParams object
// with the default values initialized.
func NewMonitorCheckGetSingleMonitorCheckParams() *MonitorCheckGetSingleMonitorCheckParams {
	var ()
	return &MonitorCheckGetSingleMonitorCheckParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorCheckGetSingleMonitorCheckParamsWithTimeout creates a new MonitorCheckGetSingleMonitorCheckParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorCheckGetSingleMonitorCheckParamsWithTimeout(timeout time.Duration) *MonitorCheckGetSingleMonitorCheckParams {
	var ()
	return &MonitorCheckGetSingleMonitorCheckParams{

		timeout: timeout,
	}
}

// NewMonitorCheckGetSingleMonitorCheckParamsWithContext creates a new MonitorCheckGetSingleMonitorCheckParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorCheckGetSingleMonitorCheckParamsWithContext(ctx context.Context) *MonitorCheckGetSingleMonitorCheckParams {
	var ()
	return &MonitorCheckGetSingleMonitorCheckParams{

		Context: ctx,
	}
}

// NewMonitorCheckGetSingleMonitorCheckParamsWithHTTPClient creates a new MonitorCheckGetSingleMonitorCheckParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorCheckGetSingleMonitorCheckParamsWithHTTPClient(client *http.Client) *MonitorCheckGetSingleMonitorCheckParams {
	var ()
	return &MonitorCheckGetSingleMonitorCheckParams{
		HTTPClient: client,
	}
}

/*MonitorCheckGetSingleMonitorCheckParams contains all the parameters to send to the API endpoint
for the monitor check get single monitor check operation typically these are written to a http.Request
*/
type MonitorCheckGetSingleMonitorCheckParams struct {

	/*MonitorCheckID*/
	MonitorCheckID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the monitor check get single monitor check params
func (o *MonitorCheckGetSingleMonitorCheckParams) WithTimeout(timeout time.Duration) *MonitorCheckGetSingleMonitorCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor check get single monitor check params
func (o *MonitorCheckGetSingleMonitorCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor check get single monitor check params
func (o *MonitorCheckGetSingleMonitorCheckParams) WithContext(ctx context.Context) *MonitorCheckGetSingleMonitorCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor check get single monitor check params
func (o *MonitorCheckGetSingleMonitorCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor check get single monitor check params
func (o *MonitorCheckGetSingleMonitorCheckParams) WithHTTPClient(client *http.Client) *MonitorCheckGetSingleMonitorCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor check get single monitor check params
func (o *MonitorCheckGetSingleMonitorCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitorCheckID adds the monitorCheckID to the monitor check get single monitor check params
func (o *MonitorCheckGetSingleMonitorCheckParams) WithMonitorCheckID(monitorCheckID int64) *MonitorCheckGetSingleMonitorCheckParams {
	o.SetMonitorCheckID(monitorCheckID)
	return o
}

// SetMonitorCheckID adds the monitorCheckId to the monitor check get single monitor check params
func (o *MonitorCheckGetSingleMonitorCheckParams) SetMonitorCheckID(monitorCheckID int64) {
	o.MonitorCheckID = monitorCheckID
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorCheckGetSingleMonitorCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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