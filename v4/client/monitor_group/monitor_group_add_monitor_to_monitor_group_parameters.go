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

// NewMonitorGroupAddMonitorToMonitorGroupParams creates a new MonitorGroupAddMonitorToMonitorGroupParams object
// with the default values initialized.
func NewMonitorGroupAddMonitorToMonitorGroupParams() *MonitorGroupAddMonitorToMonitorGroupParams {
	var ()
	return &MonitorGroupAddMonitorToMonitorGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorGroupAddMonitorToMonitorGroupParamsWithTimeout creates a new MonitorGroupAddMonitorToMonitorGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorGroupAddMonitorToMonitorGroupParamsWithTimeout(timeout time.Duration) *MonitorGroupAddMonitorToMonitorGroupParams {
	var ()
	return &MonitorGroupAddMonitorToMonitorGroupParams{

		timeout: timeout,
	}
}

// NewMonitorGroupAddMonitorToMonitorGroupParamsWithContext creates a new MonitorGroupAddMonitorToMonitorGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorGroupAddMonitorToMonitorGroupParamsWithContext(ctx context.Context) *MonitorGroupAddMonitorToMonitorGroupParams {
	var ()
	return &MonitorGroupAddMonitorToMonitorGroupParams{

		Context: ctx,
	}
}

// NewMonitorGroupAddMonitorToMonitorGroupParamsWithHTTPClient creates a new MonitorGroupAddMonitorToMonitorGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorGroupAddMonitorToMonitorGroupParamsWithHTTPClient(client *http.Client) *MonitorGroupAddMonitorToMonitorGroupParams {
	var ()
	return &MonitorGroupAddMonitorToMonitorGroupParams{
		HTTPClient: client,
	}
}

/*MonitorGroupAddMonitorToMonitorGroupParams contains all the parameters to send to the API endpoint
for the monitor group add monitor to monitor group operation typically these are written to a http.Request
*/
type MonitorGroupAddMonitorToMonitorGroupParams struct {

	/*MonitorGroupGUID
	  The Guid of the monitor group to add the monitor to

	*/
	MonitorGroupGUID string
	/*MonitorGUID
	  The monitor Guid

	*/
	MonitorGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the monitor group add monitor to monitor group params
func (o *MonitorGroupAddMonitorToMonitorGroupParams) WithTimeout(timeout time.Duration) *MonitorGroupAddMonitorToMonitorGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor group add monitor to monitor group params
func (o *MonitorGroupAddMonitorToMonitorGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor group add monitor to monitor group params
func (o *MonitorGroupAddMonitorToMonitorGroupParams) WithContext(ctx context.Context) *MonitorGroupAddMonitorToMonitorGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor group add monitor to monitor group params
func (o *MonitorGroupAddMonitorToMonitorGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor group add monitor to monitor group params
func (o *MonitorGroupAddMonitorToMonitorGroupParams) WithHTTPClient(client *http.Client) *MonitorGroupAddMonitorToMonitorGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor group add monitor to monitor group params
func (o *MonitorGroupAddMonitorToMonitorGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitorGroupGUID adds the monitorGroupGUID to the monitor group add monitor to monitor group params
func (o *MonitorGroupAddMonitorToMonitorGroupParams) WithMonitorGroupGUID(monitorGroupGUID string) *MonitorGroupAddMonitorToMonitorGroupParams {
	o.SetMonitorGroupGUID(monitorGroupGUID)
	return o
}

// SetMonitorGroupGUID adds the monitorGroupGuid to the monitor group add monitor to monitor group params
func (o *MonitorGroupAddMonitorToMonitorGroupParams) SetMonitorGroupGUID(monitorGroupGUID string) {
	o.MonitorGroupGUID = monitorGroupGUID
}

// WithMonitorGUID adds the monitorGUID to the monitor group add monitor to monitor group params
func (o *MonitorGroupAddMonitorToMonitorGroupParams) WithMonitorGUID(monitorGUID string) *MonitorGroupAddMonitorToMonitorGroupParams {
	o.SetMonitorGUID(monitorGUID)
	return o
}

// SetMonitorGUID adds the monitorGuid to the monitor group add monitor to monitor group params
func (o *MonitorGroupAddMonitorToMonitorGroupParams) SetMonitorGUID(monitorGUID string) {
	o.MonitorGUID = monitorGUID
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorGroupAddMonitorToMonitorGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param monitorGroupGuid
	if err := r.SetPathParam("monitorGroupGuid", o.MonitorGroupGUID); err != nil {
		return err
	}

	// path param monitorGuid
	if err := r.SetPathParam("monitorGuid", o.MonitorGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
