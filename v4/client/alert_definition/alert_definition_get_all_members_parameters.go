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

// NewAlertDefinitionGetAllMembersParams creates a new AlertDefinitionGetAllMembersParams object
// with the default values initialized.
func NewAlertDefinitionGetAllMembersParams() *AlertDefinitionGetAllMembersParams {
	var ()
	return &AlertDefinitionGetAllMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAlertDefinitionGetAllMembersParamsWithTimeout creates a new AlertDefinitionGetAllMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAlertDefinitionGetAllMembersParamsWithTimeout(timeout time.Duration) *AlertDefinitionGetAllMembersParams {
	var ()
	return &AlertDefinitionGetAllMembersParams{

		timeout: timeout,
	}
}

// NewAlertDefinitionGetAllMembersParamsWithContext creates a new AlertDefinitionGetAllMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewAlertDefinitionGetAllMembersParamsWithContext(ctx context.Context) *AlertDefinitionGetAllMembersParams {
	var ()
	return &AlertDefinitionGetAllMembersParams{

		Context: ctx,
	}
}

// NewAlertDefinitionGetAllMembersParamsWithHTTPClient creates a new AlertDefinitionGetAllMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAlertDefinitionGetAllMembersParamsWithHTTPClient(client *http.Client) *AlertDefinitionGetAllMembersParams {
	var ()
	return &AlertDefinitionGetAllMembersParams{
		HTTPClient: client,
	}
}

/*AlertDefinitionGetAllMembersParams contains all the parameters to send to the API endpoint
for the alert definition get all members operation typically these are written to a http.Request
*/
type AlertDefinitionGetAllMembersParams struct {

	/*AlertDefinitionGUID
	  The Guid of the alert definition for which to return the members.

	*/
	AlertDefinitionGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alert definition get all members params
func (o *AlertDefinitionGetAllMembersParams) WithTimeout(timeout time.Duration) *AlertDefinitionGetAllMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alert definition get all members params
func (o *AlertDefinitionGetAllMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alert definition get all members params
func (o *AlertDefinitionGetAllMembersParams) WithContext(ctx context.Context) *AlertDefinitionGetAllMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alert definition get all members params
func (o *AlertDefinitionGetAllMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alert definition get all members params
func (o *AlertDefinitionGetAllMembersParams) WithHTTPClient(client *http.Client) *AlertDefinitionGetAllMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alert definition get all members params
func (o *AlertDefinitionGetAllMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertDefinitionGUID adds the alertDefinitionGUID to the alert definition get all members params
func (o *AlertDefinitionGetAllMembersParams) WithAlertDefinitionGUID(alertDefinitionGUID string) *AlertDefinitionGetAllMembersParams {
	o.SetAlertDefinitionGUID(alertDefinitionGUID)
	return o
}

// SetAlertDefinitionGUID adds the alertDefinitionGuid to the alert definition get all members params
func (o *AlertDefinitionGetAllMembersParams) SetAlertDefinitionGUID(alertDefinitionGUID string) {
	o.AlertDefinitionGUID = alertDefinitionGUID
}

// WriteToRequest writes these params to a swagger request
func (o *AlertDefinitionGetAllMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertDefinitionGuid
	if err := r.SetPathParam("alertDefinitionGuid", o.AlertDefinitionGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
