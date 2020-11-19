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
)

// NewAlertDefinitionGetEscalationLevelParams creates a new AlertDefinitionGetEscalationLevelParams object
// with the default values initialized.
func NewAlertDefinitionGetEscalationLevelParams() *AlertDefinitionGetEscalationLevelParams {
	var ()
	return &AlertDefinitionGetEscalationLevelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAlertDefinitionGetEscalationLevelParamsWithTimeout creates a new AlertDefinitionGetEscalationLevelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAlertDefinitionGetEscalationLevelParamsWithTimeout(timeout time.Duration) *AlertDefinitionGetEscalationLevelParams {
	var ()
	return &AlertDefinitionGetEscalationLevelParams{

		timeout: timeout,
	}
}

// NewAlertDefinitionGetEscalationLevelParamsWithContext creates a new AlertDefinitionGetEscalationLevelParams object
// with the default values initialized, and the ability to set a context for a request
func NewAlertDefinitionGetEscalationLevelParamsWithContext(ctx context.Context) *AlertDefinitionGetEscalationLevelParams {
	var ()
	return &AlertDefinitionGetEscalationLevelParams{

		Context: ctx,
	}
}

// NewAlertDefinitionGetEscalationLevelParamsWithHTTPClient creates a new AlertDefinitionGetEscalationLevelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAlertDefinitionGetEscalationLevelParamsWithHTTPClient(client *http.Client) *AlertDefinitionGetEscalationLevelParams {
	var ()
	return &AlertDefinitionGetEscalationLevelParams{
		HTTPClient: client,
	}
}

/*AlertDefinitionGetEscalationLevelParams contains all the parameters to send to the API endpoint
for the alert definition get escalation level operation typically these are written to a http.Request
*/
type AlertDefinitionGetEscalationLevelParams struct {

	/*AlertDefinitionGUID
	  The Guid of the alert definition.

	*/
	AlertDefinitionGUID string
	/*EscalationLevelID
	  The escalation level id.

	*/
	EscalationLevelID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alert definition get escalation level params
func (o *AlertDefinitionGetEscalationLevelParams) WithTimeout(timeout time.Duration) *AlertDefinitionGetEscalationLevelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alert definition get escalation level params
func (o *AlertDefinitionGetEscalationLevelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alert definition get escalation level params
func (o *AlertDefinitionGetEscalationLevelParams) WithContext(ctx context.Context) *AlertDefinitionGetEscalationLevelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alert definition get escalation level params
func (o *AlertDefinitionGetEscalationLevelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alert definition get escalation level params
func (o *AlertDefinitionGetEscalationLevelParams) WithHTTPClient(client *http.Client) *AlertDefinitionGetEscalationLevelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alert definition get escalation level params
func (o *AlertDefinitionGetEscalationLevelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertDefinitionGUID adds the alertDefinitionGUID to the alert definition get escalation level params
func (o *AlertDefinitionGetEscalationLevelParams) WithAlertDefinitionGUID(alertDefinitionGUID string) *AlertDefinitionGetEscalationLevelParams {
	o.SetAlertDefinitionGUID(alertDefinitionGUID)
	return o
}

// SetAlertDefinitionGUID adds the alertDefinitionGuid to the alert definition get escalation level params
func (o *AlertDefinitionGetEscalationLevelParams) SetAlertDefinitionGUID(alertDefinitionGUID string) {
	o.AlertDefinitionGUID = alertDefinitionGUID
}

// WithEscalationLevelID adds the escalationLevelID to the alert definition get escalation level params
func (o *AlertDefinitionGetEscalationLevelParams) WithEscalationLevelID(escalationLevelID int64) *AlertDefinitionGetEscalationLevelParams {
	o.SetEscalationLevelID(escalationLevelID)
	return o
}

// SetEscalationLevelID adds the escalationLevelId to the alert definition get escalation level params
func (o *AlertDefinitionGetEscalationLevelParams) SetEscalationLevelID(escalationLevelID int64) {
	o.EscalationLevelID = escalationLevelID
}

// WriteToRequest writes these params to a swagger request
func (o *AlertDefinitionGetEscalationLevelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
