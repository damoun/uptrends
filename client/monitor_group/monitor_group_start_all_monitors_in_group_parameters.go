// Code generated by go-swagger; DO NOT EDIT.

package monitor_group

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

// NewMonitorGroupStartAllMonitorsInGroupParams creates a new MonitorGroupStartAllMonitorsInGroupParams object
// with the default values initialized.
func NewMonitorGroupStartAllMonitorsInGroupParams() *MonitorGroupStartAllMonitorsInGroupParams {
	var ()
	return &MonitorGroupStartAllMonitorsInGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorGroupStartAllMonitorsInGroupParamsWithTimeout creates a new MonitorGroupStartAllMonitorsInGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorGroupStartAllMonitorsInGroupParamsWithTimeout(timeout time.Duration) *MonitorGroupStartAllMonitorsInGroupParams {
	var ()
	return &MonitorGroupStartAllMonitorsInGroupParams{

		timeout: timeout,
	}
}

// NewMonitorGroupStartAllMonitorsInGroupParamsWithContext creates a new MonitorGroupStartAllMonitorsInGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorGroupStartAllMonitorsInGroupParamsWithContext(ctx context.Context) *MonitorGroupStartAllMonitorsInGroupParams {
	var ()
	return &MonitorGroupStartAllMonitorsInGroupParams{

		Context: ctx,
	}
}

// NewMonitorGroupStartAllMonitorsInGroupParamsWithHTTPClient creates a new MonitorGroupStartAllMonitorsInGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorGroupStartAllMonitorsInGroupParamsWithHTTPClient(client *http.Client) *MonitorGroupStartAllMonitorsInGroupParams {
	var ()
	return &MonitorGroupStartAllMonitorsInGroupParams{
		HTTPClient: client,
	}
}

/*MonitorGroupStartAllMonitorsInGroupParams contains all the parameters to send to the API endpoint
for the monitor group start all monitors in group operation typically these are written to a http.Request
*/
type MonitorGroupStartAllMonitorsInGroupParams struct {

	/*MonitorGroupGUID
	  The monitor group GUID

	*/
	MonitorGroupGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the monitor group start all monitors in group params
func (o *MonitorGroupStartAllMonitorsInGroupParams) WithTimeout(timeout time.Duration) *MonitorGroupStartAllMonitorsInGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor group start all monitors in group params
func (o *MonitorGroupStartAllMonitorsInGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor group start all monitors in group params
func (o *MonitorGroupStartAllMonitorsInGroupParams) WithContext(ctx context.Context) *MonitorGroupStartAllMonitorsInGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor group start all monitors in group params
func (o *MonitorGroupStartAllMonitorsInGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor group start all monitors in group params
func (o *MonitorGroupStartAllMonitorsInGroupParams) WithHTTPClient(client *http.Client) *MonitorGroupStartAllMonitorsInGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor group start all monitors in group params
func (o *MonitorGroupStartAllMonitorsInGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitorGroupGUID adds the monitorGroupGUID to the monitor group start all monitors in group params
func (o *MonitorGroupStartAllMonitorsInGroupParams) WithMonitorGroupGUID(monitorGroupGUID string) *MonitorGroupStartAllMonitorsInGroupParams {
	o.SetMonitorGroupGUID(monitorGroupGUID)
	return o
}

// SetMonitorGroupGUID adds the monitorGroupGuid to the monitor group start all monitors in group params
func (o *MonitorGroupStartAllMonitorsInGroupParams) SetMonitorGroupGUID(monitorGroupGUID string) {
	o.MonitorGroupGUID = monitorGroupGUID
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorGroupStartAllMonitorsInGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
