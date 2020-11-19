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

// NewAlertDefinitionRemoveOperatorFromEscalationLevelParams creates a new AlertDefinitionRemoveOperatorFromEscalationLevelParams object
// with the default values initialized.
func NewAlertDefinitionRemoveOperatorFromEscalationLevelParams() *AlertDefinitionRemoveOperatorFromEscalationLevelParams {
	var ()
	return &AlertDefinitionRemoveOperatorFromEscalationLevelParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAlertDefinitionRemoveOperatorFromEscalationLevelParamsWithTimeout creates a new AlertDefinitionRemoveOperatorFromEscalationLevelParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAlertDefinitionRemoveOperatorFromEscalationLevelParamsWithTimeout(timeout time.Duration) *AlertDefinitionRemoveOperatorFromEscalationLevelParams {
	var ()
	return &AlertDefinitionRemoveOperatorFromEscalationLevelParams{

		timeout: timeout,
	}
}

// NewAlertDefinitionRemoveOperatorFromEscalationLevelParamsWithContext creates a new AlertDefinitionRemoveOperatorFromEscalationLevelParams object
// with the default values initialized, and the ability to set a context for a request
func NewAlertDefinitionRemoveOperatorFromEscalationLevelParamsWithContext(ctx context.Context) *AlertDefinitionRemoveOperatorFromEscalationLevelParams {
	var ()
	return &AlertDefinitionRemoveOperatorFromEscalationLevelParams{

		Context: ctx,
	}
}

// NewAlertDefinitionRemoveOperatorFromEscalationLevelParamsWithHTTPClient creates a new AlertDefinitionRemoveOperatorFromEscalationLevelParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAlertDefinitionRemoveOperatorFromEscalationLevelParamsWithHTTPClient(client *http.Client) *AlertDefinitionRemoveOperatorFromEscalationLevelParams {
	var ()
	return &AlertDefinitionRemoveOperatorFromEscalationLevelParams{
		HTTPClient: client,
	}
}

/*AlertDefinitionRemoveOperatorFromEscalationLevelParams contains all the parameters to send to the API endpoint
for the alert definition remove operator from escalation level operation typically these are written to a http.Request
*/
type AlertDefinitionRemoveOperatorFromEscalationLevelParams struct {

	/*AlertDefinitionGUID
	  The Guid of the alert definition.

	*/
	AlertDefinitionGUID string
	/*EscalationLevelID
	  The escalation level id.

	*/
	EscalationLevelID int64
	/*OperatorGUID
	  The Guid of the operator to remove.

	*/
	OperatorGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) WithTimeout(timeout time.Duration) *AlertDefinitionRemoveOperatorFromEscalationLevelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) WithContext(ctx context.Context) *AlertDefinitionRemoveOperatorFromEscalationLevelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) WithHTTPClient(client *http.Client) *AlertDefinitionRemoveOperatorFromEscalationLevelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertDefinitionGUID adds the alertDefinitionGUID to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) WithAlertDefinitionGUID(alertDefinitionGUID string) *AlertDefinitionRemoveOperatorFromEscalationLevelParams {
	o.SetAlertDefinitionGUID(alertDefinitionGUID)
	return o
}

// SetAlertDefinitionGUID adds the alertDefinitionGuid to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) SetAlertDefinitionGUID(alertDefinitionGUID string) {
	o.AlertDefinitionGUID = alertDefinitionGUID
}

// WithEscalationLevelID adds the escalationLevelID to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) WithEscalationLevelID(escalationLevelID int64) *AlertDefinitionRemoveOperatorFromEscalationLevelParams {
	o.SetEscalationLevelID(escalationLevelID)
	return o
}

// SetEscalationLevelID adds the escalationLevelId to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) SetEscalationLevelID(escalationLevelID int64) {
	o.EscalationLevelID = escalationLevelID
}

// WithOperatorGUID adds the operatorGUID to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) WithOperatorGUID(operatorGUID string) *AlertDefinitionRemoveOperatorFromEscalationLevelParams {
	o.SetOperatorGUID(operatorGUID)
	return o
}

// SetOperatorGUID adds the operatorGuid to the alert definition remove operator from escalation level params
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) SetOperatorGUID(operatorGUID string) {
	o.OperatorGUID = operatorGUID
}

// WriteToRequest writes these params to a swagger request
func (o *AlertDefinitionRemoveOperatorFromEscalationLevelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param operatorGuid
	if err := r.SetPathParam("operatorGuid", o.OperatorGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
