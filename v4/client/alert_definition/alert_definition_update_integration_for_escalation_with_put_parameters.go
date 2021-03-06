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
	"github.com/go-openapi/swag"

	"github.com/damoun/uptrends/v4/models"
)

// NewAlertDefinitionUpdateIntegrationForEscalationWithPutParams creates a new AlertDefinitionUpdateIntegrationForEscalationWithPutParams object
// with the default values initialized.
func NewAlertDefinitionUpdateIntegrationForEscalationWithPutParams() *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	var ()
	return &AlertDefinitionUpdateIntegrationForEscalationWithPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAlertDefinitionUpdateIntegrationForEscalationWithPutParamsWithTimeout creates a new AlertDefinitionUpdateIntegrationForEscalationWithPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAlertDefinitionUpdateIntegrationForEscalationWithPutParamsWithTimeout(timeout time.Duration) *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	var ()
	return &AlertDefinitionUpdateIntegrationForEscalationWithPutParams{

		timeout: timeout,
	}
}

// NewAlertDefinitionUpdateIntegrationForEscalationWithPutParamsWithContext creates a new AlertDefinitionUpdateIntegrationForEscalationWithPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewAlertDefinitionUpdateIntegrationForEscalationWithPutParamsWithContext(ctx context.Context) *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	var ()
	return &AlertDefinitionUpdateIntegrationForEscalationWithPutParams{

		Context: ctx,
	}
}

// NewAlertDefinitionUpdateIntegrationForEscalationWithPutParamsWithHTTPClient creates a new AlertDefinitionUpdateIntegrationForEscalationWithPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAlertDefinitionUpdateIntegrationForEscalationWithPutParamsWithHTTPClient(client *http.Client) *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	var ()
	return &AlertDefinitionUpdateIntegrationForEscalationWithPutParams{
		HTTPClient: client,
	}
}

/*AlertDefinitionUpdateIntegrationForEscalationWithPutParams contains all the parameters to send to the API endpoint
for the alert definition update integration for escalation with put operation typically these are written to a http.Request
*/
type AlertDefinitionUpdateIntegrationForEscalationWithPutParams struct {

	/*AlertDefinitionGUID
	  The Guid of the alert definition.

	*/
	AlertDefinitionGUID string
	/*EscalationLevelID
	  The escalation level id.

	*/
	EscalationLevelID int64
	/*EscalationLevelIntegration
	  The definition for the integration that should be updated.

	*/
	EscalationLevelIntegration *models.EscalationLevelIntegration
	/*IntegrationGUID
	  The Guid of the integration to update.

	*/
	IntegrationGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) WithTimeout(timeout time.Duration) *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) WithContext(ctx context.Context) *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) WithHTTPClient(client *http.Client) *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertDefinitionGUID adds the alertDefinitionGUID to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) WithAlertDefinitionGUID(alertDefinitionGUID string) *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	o.SetAlertDefinitionGUID(alertDefinitionGUID)
	return o
}

// SetAlertDefinitionGUID adds the alertDefinitionGuid to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) SetAlertDefinitionGUID(alertDefinitionGUID string) {
	o.AlertDefinitionGUID = alertDefinitionGUID
}

// WithEscalationLevelID adds the escalationLevelID to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) WithEscalationLevelID(escalationLevelID int64) *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	o.SetEscalationLevelID(escalationLevelID)
	return o
}

// SetEscalationLevelID adds the escalationLevelId to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) SetEscalationLevelID(escalationLevelID int64) {
	o.EscalationLevelID = escalationLevelID
}

// WithEscalationLevelIntegration adds the escalationLevelIntegration to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) WithEscalationLevelIntegration(escalationLevelIntegration *models.EscalationLevelIntegration) *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	o.SetEscalationLevelIntegration(escalationLevelIntegration)
	return o
}

// SetEscalationLevelIntegration adds the escalationLevelIntegration to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) SetEscalationLevelIntegration(escalationLevelIntegration *models.EscalationLevelIntegration) {
	o.EscalationLevelIntegration = escalationLevelIntegration
}

// WithIntegrationGUID adds the integrationGUID to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) WithIntegrationGUID(integrationGUID string) *AlertDefinitionUpdateIntegrationForEscalationWithPutParams {
	o.SetIntegrationGUID(integrationGUID)
	return o
}

// SetIntegrationGUID adds the integrationGuid to the alert definition update integration for escalation with put params
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) SetIntegrationGUID(integrationGUID string) {
	o.IntegrationGUID = integrationGUID
}

// WriteToRequest writes these params to a swagger request
func (o *AlertDefinitionUpdateIntegrationForEscalationWithPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alertDefinitionGuid
	if err := r.SetPathParam("alertDefinitionGuid", o.AlertDefinitionGUID); err != nil {
		return err
	}

	// path param escalationLevelId
	if err := r.SetPathParam("escalationLevelId", swag.FormatInt64(o.EscalationLevelID)); err != nil {
		return err
	}

	if o.EscalationLevelIntegration != nil {
		if err := r.SetBodyParam(o.EscalationLevelIntegration); err != nil {
			return err
		}
	}

	// path param integrationGuid
	if err := r.SetPathParam("integrationGuid", o.IntegrationGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
