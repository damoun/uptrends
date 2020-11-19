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

// NewAlertDefinitionAddIntegrationToEscalationLevelParams creates a new AlertDefinitionAddIntegrationToEscalationLevelParams object
// with the default values initialized.
func NewAlertDefinitionAddIntegrationToEscalationLevelParams() *AlertDefinitionAddIntegrationToEscalationLevelParams {
	var ()
	return &AlertDefinitionAddIntegrationToEscalationLevelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAlertDefinitionAddIntegrationToEscalationLevelParamsWithTimeout creates a new AlertDefinitionAddIntegrationToEscalationLevelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAlertDefinitionAddIntegrationToEscalationLevelParamsWithTimeout(timeout time.Duration) *AlertDefinitionAddIntegrationToEscalationLevelParams {
	var ()
	return &AlertDefinitionAddIntegrationToEscalationLevelParams{

		timeout: timeout,
	}
}

// NewAlertDefinitionAddIntegrationToEscalationLevelParamsWithContext creates a new AlertDefinitionAddIntegrationToEscalationLevelParams object
// with the default values initialized, and the ability to set a context for a request
func NewAlertDefinitionAddIntegrationToEscalationLevelParamsWithContext(ctx context.Context) *AlertDefinitionAddIntegrationToEscalationLevelParams {
	var ()
	return &AlertDefinitionAddIntegrationToEscalationLevelParams{

		Context: ctx,
	}
}

// NewAlertDefinitionAddIntegrationToEscalationLevelParamsWithHTTPClient creates a new AlertDefinitionAddIntegrationToEscalationLevelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAlertDefinitionAddIntegrationToEscalationLevelParamsWithHTTPClient(client *http.Client) *AlertDefinitionAddIntegrationToEscalationLevelParams {
	var ()
	return &AlertDefinitionAddIntegrationToEscalationLevelParams{
		HTTPClient: client,
	}
}

/*AlertDefinitionAddIntegrationToEscalationLevelParams contains all the parameters to send to the API endpoint
for the alert definition add integration to escalation level operation typically these are written to a http.Request
*/
type AlertDefinitionAddIntegrationToEscalationLevelParams struct {

	/*AlertDefinitionGUID
	  The Guid of the alert definition.

	*/
	AlertDefinitionGUID string
	/*EscalationLevelID
	  The escalation level id.

	*/
	EscalationLevelID int64
	/*EscalationLevelIntegration
	  The integration to add.

	*/
	EscalationLevelIntegration *models.EscalationLevelIntegration

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) WithTimeout(timeout time.Duration) *AlertDefinitionAddIntegrationToEscalationLevelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) WithContext(ctx context.Context) *AlertDefinitionAddIntegrationToEscalationLevelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) WithHTTPClient(client *http.Client) *AlertDefinitionAddIntegrationToEscalationLevelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertDefinitionGUID adds the alertDefinitionGUID to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) WithAlertDefinitionGUID(alertDefinitionGUID string) *AlertDefinitionAddIntegrationToEscalationLevelParams {
	o.SetAlertDefinitionGUID(alertDefinitionGUID)
	return o
}

// SetAlertDefinitionGUID adds the alertDefinitionGuid to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) SetAlertDefinitionGUID(alertDefinitionGUID string) {
	o.AlertDefinitionGUID = alertDefinitionGUID
}

// WithEscalationLevelID adds the escalationLevelID to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) WithEscalationLevelID(escalationLevelID int64) *AlertDefinitionAddIntegrationToEscalationLevelParams {
	o.SetEscalationLevelID(escalationLevelID)
	return o
}

// SetEscalationLevelID adds the escalationLevelId to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) SetEscalationLevelID(escalationLevelID int64) {
	o.EscalationLevelID = escalationLevelID
}

// WithEscalationLevelIntegration adds the escalationLevelIntegration to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) WithEscalationLevelIntegration(escalationLevelIntegration *models.EscalationLevelIntegration) *AlertDefinitionAddIntegrationToEscalationLevelParams {
	o.SetEscalationLevelIntegration(escalationLevelIntegration)
	return o
}

// SetEscalationLevelIntegration adds the escalationLevelIntegration to the alert definition add integration to escalation level params
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) SetEscalationLevelIntegration(escalationLevelIntegration *models.EscalationLevelIntegration) {
	o.EscalationLevelIntegration = escalationLevelIntegration
}

// WriteToRequest writes these params to a swagger request
func (o *AlertDefinitionAddIntegrationToEscalationLevelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}