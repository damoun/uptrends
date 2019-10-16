// Code generated by go-swagger; DO NOT EDIT.

package monitor

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

// NewMonitorCleanupMaintenancePeriodsParams creates a new MonitorCleanupMaintenancePeriodsParams object
// with the default values initialized.
func NewMonitorCleanupMaintenancePeriodsParams() *MonitorCleanupMaintenancePeriodsParams {
	var ()
	return &MonitorCleanupMaintenancePeriodsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorCleanupMaintenancePeriodsParamsWithTimeout creates a new MonitorCleanupMaintenancePeriodsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorCleanupMaintenancePeriodsParamsWithTimeout(timeout time.Duration) *MonitorCleanupMaintenancePeriodsParams {
	var ()
	return &MonitorCleanupMaintenancePeriodsParams{

		timeout: timeout,
	}
}

// NewMonitorCleanupMaintenancePeriodsParamsWithContext creates a new MonitorCleanupMaintenancePeriodsParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorCleanupMaintenancePeriodsParamsWithContext(ctx context.Context) *MonitorCleanupMaintenancePeriodsParams {
	var ()
	return &MonitorCleanupMaintenancePeriodsParams{

		Context: ctx,
	}
}

// NewMonitorCleanupMaintenancePeriodsParamsWithHTTPClient creates a new MonitorCleanupMaintenancePeriodsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorCleanupMaintenancePeriodsParamsWithHTTPClient(client *http.Client) *MonitorCleanupMaintenancePeriodsParams {
	var ()
	return &MonitorCleanupMaintenancePeriodsParams{
		HTTPClient: client,
	}
}

/*MonitorCleanupMaintenancePeriodsParams contains all the parameters to send to the API endpoint
for the monitor cleanup maintenance periods operation typically these are written to a http.Request
*/
type MonitorCleanupMaintenancePeriodsParams struct {

	/*BeforeDate
	  A string representing the date, formatted as "yyyy-MM-dd"

	*/
	BeforeDate strfmt.DateTime
	/*MonitorGUID*/
	MonitorGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the monitor cleanup maintenance periods params
func (o *MonitorCleanupMaintenancePeriodsParams) WithTimeout(timeout time.Duration) *MonitorCleanupMaintenancePeriodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor cleanup maintenance periods params
func (o *MonitorCleanupMaintenancePeriodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor cleanup maintenance periods params
func (o *MonitorCleanupMaintenancePeriodsParams) WithContext(ctx context.Context) *MonitorCleanupMaintenancePeriodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor cleanup maintenance periods params
func (o *MonitorCleanupMaintenancePeriodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor cleanup maintenance periods params
func (o *MonitorCleanupMaintenancePeriodsParams) WithHTTPClient(client *http.Client) *MonitorCleanupMaintenancePeriodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor cleanup maintenance periods params
func (o *MonitorCleanupMaintenancePeriodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBeforeDate adds the beforeDate to the monitor cleanup maintenance periods params
func (o *MonitorCleanupMaintenancePeriodsParams) WithBeforeDate(beforeDate strfmt.DateTime) *MonitorCleanupMaintenancePeriodsParams {
	o.SetBeforeDate(beforeDate)
	return o
}

// SetBeforeDate adds the beforeDate to the monitor cleanup maintenance periods params
func (o *MonitorCleanupMaintenancePeriodsParams) SetBeforeDate(beforeDate strfmt.DateTime) {
	o.BeforeDate = beforeDate
}

// WithMonitorGUID adds the monitorGUID to the monitor cleanup maintenance periods params
func (o *MonitorCleanupMaintenancePeriodsParams) WithMonitorGUID(monitorGUID string) *MonitorCleanupMaintenancePeriodsParams {
	o.SetMonitorGUID(monitorGUID)
	return o
}

// SetMonitorGUID adds the monitorGuid to the monitor cleanup maintenance periods params
func (o *MonitorCleanupMaintenancePeriodsParams) SetMonitorGUID(monitorGUID string) {
	o.MonitorGUID = monitorGUID
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorCleanupMaintenancePeriodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param beforeDate
	if err := r.SetPathParam("beforeDate", o.BeforeDate.String()); err != nil {
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
