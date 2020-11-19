// Code generated by go-swagger; DO NOT EDIT.

package alert_definition

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

// NewAlertDefinitionAddMonitorToAlertDefinitionParams creates a new AlertDefinitionAddMonitorToAlertDefinitionParams object
// with the default values initialized.
func NewAlertDefinitionAddMonitorToAlertDefinitionParams() *AlertDefinitionAddMonitorToAlertDefinitionParams {
	var ()
	return &AlertDefinitionAddMonitorToAlertDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAlertDefinitionAddMonitorToAlertDefinitionParamsWithTimeout creates a new AlertDefinitionAddMonitorToAlertDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAlertDefinitionAddMonitorToAlertDefinitionParamsWithTimeout(timeout time.Duration) *AlertDefinitionAddMonitorToAlertDefinitionParams {
	var ()
	return &AlertDefinitionAddMonitorToAlertDefinitionParams{

		timeout: timeout,
	}
}

// NewAlertDefinitionAddMonitorToAlertDefinitionParamsWithContext creates a new AlertDefinitionAddMonitorToAlertDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAlertDefinitionAddMonitorToAlertDefinitionParamsWithContext(ctx context.Context) *AlertDefinitionAddMonitorToAlertDefinitionParams {
	var ()
	return &AlertDefinitionAddMonitorToAlertDefinitionParams{

		Context: ctx,
	}
}

// NewAlertDefinitionAddMonitorToAlertDefinitionParamsWithHTTPClient creates a new AlertDefinitionAddMonitorToAlertDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAlertDefinitionAddMonitorToAlertDefinitionParamsWithHTTPClient(client *http.Client) *AlertDefinitionAddMonitorToAlertDefinitionParams {
	var ()
	return &AlertDefinitionAddMonitorToAlertDefinitionParams{
		HTTPClient: client,
	}
}

/*AlertDefinitionAddMonitorToAlertDefinitionParams contains all the parameters to send to the API endpoint
for the alert definition add monitor to alert definition operation typically these are written to a http.Request
*/
type AlertDefinitionAddMonitorToAlertDefinitionParams struct {

	/*AlertDefinitionGUID
	  The Guid of the alert definition to modify.

	*/
	AlertDefinitionGUID string
	/*MonitorGUID
	  The Guid of the monitor to add.

	*/
	MonitorGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alert definition add monitor to alert definition params
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) WithTimeout(timeout time.Duration) *AlertDefinitionAddMonitorToAlertDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alert definition add monitor to alert definition params
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alert definition add monitor to alert definition params
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) WithContext(ctx context.Context) *AlertDefinitionAddMonitorToAlertDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alert definition add monitor to alert definition params
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alert definition add monitor to alert definition params
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) WithHTTPClient(client *http.Client) *AlertDefinitionAddMonitorToAlertDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alert definition add monitor to alert definition params
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertDefinitionGUID adds the alertDefinitionGUID to the alert definition add monitor to alert definition params
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) WithAlertDefinitionGUID(alertDefinitionGUID string) *AlertDefinitionAddMonitorToAlertDefinitionParams {
	o.SetAlertDefinitionGUID(alertDefinitionGUID)
	return o
}

// SetAlertDefinitionGUID adds the alertDefinitionGuid to the alert definition add monitor to alert definition params
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) SetAlertDefinitionGUID(alertDefinitionGUID string) {
	o.AlertDefinitionGUID = alertDefinitionGUID
}

// WithMonitorGUID adds the monitorGUID to the alert definition add monitor to alert definition params
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) WithMonitorGUID(monitorGUID string) *AlertDefinitionAddMonitorToAlertDefinitionParams {
	o.SetMonitorGUID(monitorGUID)
	return o
}

// SetMonitorGUID adds the monitorGuid to the alert definition add monitor to alert definition params
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) SetMonitorGUID(monitorGUID string) {
	o.MonitorGUID = monitorGUID
}

// WriteToRequest writes these params to a swagger request
func (o *AlertDefinitionAddMonitorToAlertDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertDefinitionGuid
	if err := r.SetPathParam("alertDefinitionGuid", o.AlertDefinitionGUID); err != nil {
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
