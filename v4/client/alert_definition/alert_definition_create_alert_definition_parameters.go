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

// NewAlertDefinitionCreateAlertDefinitionParams creates a new AlertDefinitionCreateAlertDefinitionParams object
// with the default values initialized.
func NewAlertDefinitionCreateAlertDefinitionParams() *AlertDefinitionCreateAlertDefinitionParams {
	var ()
	return &AlertDefinitionCreateAlertDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAlertDefinitionCreateAlertDefinitionParamsWithTimeout creates a new AlertDefinitionCreateAlertDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAlertDefinitionCreateAlertDefinitionParamsWithTimeout(timeout time.Duration) *AlertDefinitionCreateAlertDefinitionParams {
	var ()
	return &AlertDefinitionCreateAlertDefinitionParams{

		timeout: timeout,
	}
}

// NewAlertDefinitionCreateAlertDefinitionParamsWithContext creates a new AlertDefinitionCreateAlertDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAlertDefinitionCreateAlertDefinitionParamsWithContext(ctx context.Context) *AlertDefinitionCreateAlertDefinitionParams {
	var ()
	return &AlertDefinitionCreateAlertDefinitionParams{

		Context: ctx,
	}
}

// NewAlertDefinitionCreateAlertDefinitionParamsWithHTTPClient creates a new AlertDefinitionCreateAlertDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAlertDefinitionCreateAlertDefinitionParamsWithHTTPClient(client *http.Client) *AlertDefinitionCreateAlertDefinitionParams {
	var ()
	return &AlertDefinitionCreateAlertDefinitionParams{
		HTTPClient: client,
	}
}

/*AlertDefinitionCreateAlertDefinitionParams contains all the parameters to send to the API endpoint
for the alert definition create alert definition operation typically these are written to a http.Request
*/
type AlertDefinitionCreateAlertDefinitionParams struct {

	/*AlertDefinition
	  The details of the alert definition to create.

	*/
	AlertDefinition *models.AlertDefinition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alert definition create alert definition params
func (o *AlertDefinitionCreateAlertDefinitionParams) WithTimeout(timeout time.Duration) *AlertDefinitionCreateAlertDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alert definition create alert definition params
func (o *AlertDefinitionCreateAlertDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alert definition create alert definition params
func (o *AlertDefinitionCreateAlertDefinitionParams) WithContext(ctx context.Context) *AlertDefinitionCreateAlertDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alert definition create alert definition params
func (o *AlertDefinitionCreateAlertDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alert definition create alert definition params
func (o *AlertDefinitionCreateAlertDefinitionParams) WithHTTPClient(client *http.Client) *AlertDefinitionCreateAlertDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alert definition create alert definition params
func (o *AlertDefinitionCreateAlertDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertDefinition adds the alertDefinition to the alert definition create alert definition params
func (o *AlertDefinitionCreateAlertDefinitionParams) WithAlertDefinition(alertDefinition *models.AlertDefinition) *AlertDefinitionCreateAlertDefinitionParams {
	o.SetAlertDefinition(alertDefinition)
	return o
}

// SetAlertDefinition adds the alertDefinition to the alert definition create alert definition params
func (o *AlertDefinitionCreateAlertDefinitionParams) SetAlertDefinition(alertDefinition *models.AlertDefinition) {
	o.AlertDefinition = alertDefinition
}

// WriteToRequest writes these params to a swagger request
func (o *AlertDefinitionCreateAlertDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AlertDefinition != nil {
		if err := r.SetBodyParam(o.AlertDefinition); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}