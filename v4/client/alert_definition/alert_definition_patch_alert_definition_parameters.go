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

	"github.com/damoun/uptrends/v4/models"
)

// NewAlertDefinitionPatchAlertDefinitionParams creates a new AlertDefinitionPatchAlertDefinitionParams object
// with the default values initialized.
func NewAlertDefinitionPatchAlertDefinitionParams() *AlertDefinitionPatchAlertDefinitionParams {
	var ()
	return &AlertDefinitionPatchAlertDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAlertDefinitionPatchAlertDefinitionParamsWithTimeout creates a new AlertDefinitionPatchAlertDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAlertDefinitionPatchAlertDefinitionParamsWithTimeout(timeout time.Duration) *AlertDefinitionPatchAlertDefinitionParams {
	var ()
	return &AlertDefinitionPatchAlertDefinitionParams{

		timeout: timeout,
	}
}

// NewAlertDefinitionPatchAlertDefinitionParamsWithContext creates a new AlertDefinitionPatchAlertDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAlertDefinitionPatchAlertDefinitionParamsWithContext(ctx context.Context) *AlertDefinitionPatchAlertDefinitionParams {
	var ()
	return &AlertDefinitionPatchAlertDefinitionParams{

		Context: ctx,
	}
}

// NewAlertDefinitionPatchAlertDefinitionParamsWithHTTPClient creates a new AlertDefinitionPatchAlertDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAlertDefinitionPatchAlertDefinitionParamsWithHTTPClient(client *http.Client) *AlertDefinitionPatchAlertDefinitionParams {
	var ()
	return &AlertDefinitionPatchAlertDefinitionParams{
		HTTPClient: client,
	}
}

/*AlertDefinitionPatchAlertDefinitionParams contains all the parameters to send to the API endpoint
for the alert definition patch alert definition operation typically these are written to a http.Request
*/
type AlertDefinitionPatchAlertDefinitionParams struct {

	/*AlertDefinition
	  The partial definition for the alert definition that should be updated.

	*/
	AlertDefinition *models.AlertDefinition
	/*AlertDefinitionGUID
	  The Guid of the alert definition that should be updated.

	*/
	AlertDefinitionGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alert definition patch alert definition params
func (o *AlertDefinitionPatchAlertDefinitionParams) WithTimeout(timeout time.Duration) *AlertDefinitionPatchAlertDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alert definition patch alert definition params
func (o *AlertDefinitionPatchAlertDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alert definition patch alert definition params
func (o *AlertDefinitionPatchAlertDefinitionParams) WithContext(ctx context.Context) *AlertDefinitionPatchAlertDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alert definition patch alert definition params
func (o *AlertDefinitionPatchAlertDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alert definition patch alert definition params
func (o *AlertDefinitionPatchAlertDefinitionParams) WithHTTPClient(client *http.Client) *AlertDefinitionPatchAlertDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alert definition patch alert definition params
func (o *AlertDefinitionPatchAlertDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertDefinition adds the alertDefinition to the alert definition patch alert definition params
func (o *AlertDefinitionPatchAlertDefinitionParams) WithAlertDefinition(alertDefinition *models.AlertDefinition) *AlertDefinitionPatchAlertDefinitionParams {
	o.SetAlertDefinition(alertDefinition)
	return o
}

// SetAlertDefinition adds the alertDefinition to the alert definition patch alert definition params
func (o *AlertDefinitionPatchAlertDefinitionParams) SetAlertDefinition(alertDefinition *models.AlertDefinition) {
	o.AlertDefinition = alertDefinition
}

// WithAlertDefinitionGUID adds the alertDefinitionGUID to the alert definition patch alert definition params
func (o *AlertDefinitionPatchAlertDefinitionParams) WithAlertDefinitionGUID(alertDefinitionGUID string) *AlertDefinitionPatchAlertDefinitionParams {
	o.SetAlertDefinitionGUID(alertDefinitionGUID)
	return o
}

// SetAlertDefinitionGUID adds the alertDefinitionGuid to the alert definition patch alert definition params
func (o *AlertDefinitionPatchAlertDefinitionParams) SetAlertDefinitionGUID(alertDefinitionGUID string) {
	o.AlertDefinitionGUID = alertDefinitionGUID
}

// WriteToRequest writes these params to a swagger request
func (o *AlertDefinitionPatchAlertDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AlertDefinition != nil {
		if err := r.SetBodyParam(o.AlertDefinition); err != nil {
			return err
		}
	}

	// path param alertDefinitionGuid
	if err := r.SetPathParam("alertDefinitionGuid", o.AlertDefinitionGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
