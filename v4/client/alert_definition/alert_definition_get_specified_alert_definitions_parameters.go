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

// NewAlertDefinitionGetSpecifiedAlertDefinitionsParams creates a new AlertDefinitionGetSpecifiedAlertDefinitionsParams object
// with the default values initialized.
func NewAlertDefinitionGetSpecifiedAlertDefinitionsParams() *AlertDefinitionGetSpecifiedAlertDefinitionsParams {
	var ()
	return &AlertDefinitionGetSpecifiedAlertDefinitionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAlertDefinitionGetSpecifiedAlertDefinitionsParamsWithTimeout creates a new AlertDefinitionGetSpecifiedAlertDefinitionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAlertDefinitionGetSpecifiedAlertDefinitionsParamsWithTimeout(timeout time.Duration) *AlertDefinitionGetSpecifiedAlertDefinitionsParams {
	var ()
	return &AlertDefinitionGetSpecifiedAlertDefinitionsParams{

		timeout: timeout,
	}
}

// NewAlertDefinitionGetSpecifiedAlertDefinitionsParamsWithContext creates a new AlertDefinitionGetSpecifiedAlertDefinitionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAlertDefinitionGetSpecifiedAlertDefinitionsParamsWithContext(ctx context.Context) *AlertDefinitionGetSpecifiedAlertDefinitionsParams {
	var ()
	return &AlertDefinitionGetSpecifiedAlertDefinitionsParams{

		Context: ctx,
	}
}

// NewAlertDefinitionGetSpecifiedAlertDefinitionsParamsWithHTTPClient creates a new AlertDefinitionGetSpecifiedAlertDefinitionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAlertDefinitionGetSpecifiedAlertDefinitionsParamsWithHTTPClient(client *http.Client) *AlertDefinitionGetSpecifiedAlertDefinitionsParams {
	var ()
	return &AlertDefinitionGetSpecifiedAlertDefinitionsParams{
		HTTPClient: client,
	}
}

/*AlertDefinitionGetSpecifiedAlertDefinitionsParams contains all the parameters to send to the API endpoint
for the alert definition get specified alert definitions operation typically these are written to a http.Request
*/
type AlertDefinitionGetSpecifiedAlertDefinitionsParams struct {

	/*AlertDefinitionGUID
	  The Guid of the alert definition.

	*/
	AlertDefinitionGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alert definition get specified alert definitions params
func (o *AlertDefinitionGetSpecifiedAlertDefinitionsParams) WithTimeout(timeout time.Duration) *AlertDefinitionGetSpecifiedAlertDefinitionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alert definition get specified alert definitions params
func (o *AlertDefinitionGetSpecifiedAlertDefinitionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alert definition get specified alert definitions params
func (o *AlertDefinitionGetSpecifiedAlertDefinitionsParams) WithContext(ctx context.Context) *AlertDefinitionGetSpecifiedAlertDefinitionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alert definition get specified alert definitions params
func (o *AlertDefinitionGetSpecifiedAlertDefinitionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alert definition get specified alert definitions params
func (o *AlertDefinitionGetSpecifiedAlertDefinitionsParams) WithHTTPClient(client *http.Client) *AlertDefinitionGetSpecifiedAlertDefinitionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alert definition get specified alert definitions params
func (o *AlertDefinitionGetSpecifiedAlertDefinitionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertDefinitionGUID adds the alertDefinitionGUID to the alert definition get specified alert definitions params
func (o *AlertDefinitionGetSpecifiedAlertDefinitionsParams) WithAlertDefinitionGUID(alertDefinitionGUID string) *AlertDefinitionGetSpecifiedAlertDefinitionsParams {
	o.SetAlertDefinitionGUID(alertDefinitionGUID)
	return o
}

// SetAlertDefinitionGUID adds the alertDefinitionGuid to the alert definition get specified alert definitions params
func (o *AlertDefinitionGetSpecifiedAlertDefinitionsParams) SetAlertDefinitionGUID(alertDefinitionGUID string) {
	o.AlertDefinitionGUID = alertDefinitionGUID
}

// WriteToRequest writes these params to a swagger request
func (o *AlertDefinitionGetSpecifiedAlertDefinitionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
