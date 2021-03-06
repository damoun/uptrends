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
	"github.com/go-openapi/strfmt"
)

// NewMonitorGroupStopAllMonitorsInGroupParams creates a new MonitorGroupStopAllMonitorsInGroupParams object
// with the default values initialized.
func NewMonitorGroupStopAllMonitorsInGroupParams() *MonitorGroupStopAllMonitorsInGroupParams {
	var ()
	return &MonitorGroupStopAllMonitorsInGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorGroupStopAllMonitorsInGroupParamsWithTimeout creates a new MonitorGroupStopAllMonitorsInGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorGroupStopAllMonitorsInGroupParamsWithTimeout(timeout time.Duration) *MonitorGroupStopAllMonitorsInGroupParams {
	var ()
	return &MonitorGroupStopAllMonitorsInGroupParams{

		timeout: timeout,
	}
}

// NewMonitorGroupStopAllMonitorsInGroupParamsWithContext creates a new MonitorGroupStopAllMonitorsInGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorGroupStopAllMonitorsInGroupParamsWithContext(ctx context.Context) *MonitorGroupStopAllMonitorsInGroupParams {
	var ()
	return &MonitorGroupStopAllMonitorsInGroupParams{

		Context: ctx,
	}
}

// NewMonitorGroupStopAllMonitorsInGroupParamsWithHTTPClient creates a new MonitorGroupStopAllMonitorsInGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorGroupStopAllMonitorsInGroupParamsWithHTTPClient(client *http.Client) *MonitorGroupStopAllMonitorsInGroupParams {
	var ()
	return &MonitorGroupStopAllMonitorsInGroupParams{
		HTTPClient: client,
	}
}

/*MonitorGroupStopAllMonitorsInGroupParams contains all the parameters to send to the API endpoint
for the monitor group stop all monitors in group operation typically these are written to a http.Request
*/
type MonitorGroupStopAllMonitorsInGroupParams struct {

	/*MonitorGroupGUID
	  The monitor group GUID

	*/
	MonitorGroupGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the monitor group stop all monitors in group params
func (o *MonitorGroupStopAllMonitorsInGroupParams) WithTimeout(timeout time.Duration) *MonitorGroupStopAllMonitorsInGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor group stop all monitors in group params
func (o *MonitorGroupStopAllMonitorsInGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor group stop all monitors in group params
func (o *MonitorGroupStopAllMonitorsInGroupParams) WithContext(ctx context.Context) *MonitorGroupStopAllMonitorsInGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor group stop all monitors in group params
func (o *MonitorGroupStopAllMonitorsInGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor group stop all monitors in group params
func (o *MonitorGroupStopAllMonitorsInGroupParams) WithHTTPClient(client *http.Client) *MonitorGroupStopAllMonitorsInGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor group stop all monitors in group params
func (o *MonitorGroupStopAllMonitorsInGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitorGroupGUID adds the monitorGroupGUID to the monitor group stop all monitors in group params
func (o *MonitorGroupStopAllMonitorsInGroupParams) WithMonitorGroupGUID(monitorGroupGUID string) *MonitorGroupStopAllMonitorsInGroupParams {
	o.SetMonitorGroupGUID(monitorGroupGUID)
	return o
}

// SetMonitorGroupGUID adds the monitorGroupGuid to the monitor group stop all monitors in group params
func (o *MonitorGroupStopAllMonitorsInGroupParams) SetMonitorGroupGUID(monitorGroupGUID string) {
	o.MonitorGroupGUID = monitorGroupGUID
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorGroupStopAllMonitorsInGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
