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

// NewMonitorGroupGetMonitorGroupMembersParams creates a new MonitorGroupGetMonitorGroupMembersParams object
// with the default values initialized.
func NewMonitorGroupGetMonitorGroupMembersParams() *MonitorGroupGetMonitorGroupMembersParams {
	var ()
	return &MonitorGroupGetMonitorGroupMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorGroupGetMonitorGroupMembersParamsWithTimeout creates a new MonitorGroupGetMonitorGroupMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorGroupGetMonitorGroupMembersParamsWithTimeout(timeout time.Duration) *MonitorGroupGetMonitorGroupMembersParams {
	var ()
	return &MonitorGroupGetMonitorGroupMembersParams{

		timeout: timeout,
	}
}

// NewMonitorGroupGetMonitorGroupMembersParamsWithContext creates a new MonitorGroupGetMonitorGroupMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorGroupGetMonitorGroupMembersParamsWithContext(ctx context.Context) *MonitorGroupGetMonitorGroupMembersParams {
	var ()
	return &MonitorGroupGetMonitorGroupMembersParams{

		Context: ctx,
	}
}

// NewMonitorGroupGetMonitorGroupMembersParamsWithHTTPClient creates a new MonitorGroupGetMonitorGroupMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorGroupGetMonitorGroupMembersParamsWithHTTPClient(client *http.Client) *MonitorGroupGetMonitorGroupMembersParams {
	var ()
	return &MonitorGroupGetMonitorGroupMembersParams{
		HTTPClient: client,
	}
}

/*MonitorGroupGetMonitorGroupMembersParams contains all the parameters to send to the API endpoint
for the monitor group get monitor group members operation typically these are written to a http.Request
*/
type MonitorGroupGetMonitorGroupMembersParams struct {

	/*MonitorGroupGUID
	  The Guid of the monitor group to retrieve the members for

	*/
	MonitorGroupGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the monitor group get monitor group members params
func (o *MonitorGroupGetMonitorGroupMembersParams) WithTimeout(timeout time.Duration) *MonitorGroupGetMonitorGroupMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor group get monitor group members params
func (o *MonitorGroupGetMonitorGroupMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor group get monitor group members params
func (o *MonitorGroupGetMonitorGroupMembersParams) WithContext(ctx context.Context) *MonitorGroupGetMonitorGroupMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor group get monitor group members params
func (o *MonitorGroupGetMonitorGroupMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor group get monitor group members params
func (o *MonitorGroupGetMonitorGroupMembersParams) WithHTTPClient(client *http.Client) *MonitorGroupGetMonitorGroupMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor group get monitor group members params
func (o *MonitorGroupGetMonitorGroupMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitorGroupGUID adds the monitorGroupGUID to the monitor group get monitor group members params
func (o *MonitorGroupGetMonitorGroupMembersParams) WithMonitorGroupGUID(monitorGroupGUID string) *MonitorGroupGetMonitorGroupMembersParams {
	o.SetMonitorGroupGUID(monitorGroupGUID)
	return o
}

// SetMonitorGroupGUID adds the monitorGroupGuid to the monitor group get monitor group members params
func (o *MonitorGroupGetMonitorGroupMembersParams) SetMonitorGroupGUID(monitorGroupGUID string) {
	o.MonitorGroupGUID = monitorGroupGUID
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorGroupGetMonitorGroupMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
