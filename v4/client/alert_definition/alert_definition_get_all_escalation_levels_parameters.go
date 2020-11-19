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

// NewAlertDefinitionGetAllEscalationLevelsParams creates a new AlertDefinitionGetAllEscalationLevelsParams object
// with the default values initialized.
func NewAlertDefinitionGetAllEscalationLevelsParams() *AlertDefinitionGetAllEscalationLevelsParams {
	var ()
	return &AlertDefinitionGetAllEscalationLevelsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAlertDefinitionGetAllEscalationLevelsParamsWithTimeout creates a new AlertDefinitionGetAllEscalationLevelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAlertDefinitionGetAllEscalationLevelsParamsWithTimeout(timeout time.Duration) *AlertDefinitionGetAllEscalationLevelsParams {
	var ()
	return &AlertDefinitionGetAllEscalationLevelsParams{

		timeout: timeout,
	}
}

// NewAlertDefinitionGetAllEscalationLevelsParamsWithContext creates a new AlertDefinitionGetAllEscalationLevelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAlertDefinitionGetAllEscalationLevelsParamsWithContext(ctx context.Context) *AlertDefinitionGetAllEscalationLevelsParams {
	var ()
	return &AlertDefinitionGetAllEscalationLevelsParams{

		Context: ctx,
	}
}

// NewAlertDefinitionGetAllEscalationLevelsParamsWithHTTPClient creates a new AlertDefinitionGetAllEscalationLevelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAlertDefinitionGetAllEscalationLevelsParamsWithHTTPClient(client *http.Client) *AlertDefinitionGetAllEscalationLevelsParams {
	var ()
	return &AlertDefinitionGetAllEscalationLevelsParams{
		HTTPClient: client,
	}
}

/*AlertDefinitionGetAllEscalationLevelsParams contains all the parameters to send to the API endpoint
for the alert definition get all escalation levels operation typically these are written to a http.Request
*/
type AlertDefinitionGetAllEscalationLevelsParams struct {

	/*AlertDefinitionGUID
	  The Guid of the alert definition for which to return all escalation levels.

	*/
	AlertDefinitionGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the alert definition get all escalation levels params
func (o *AlertDefinitionGetAllEscalationLevelsParams) WithTimeout(timeout time.Duration) *AlertDefinitionGetAllEscalationLevelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the alert definition get all escalation levels params
func (o *AlertDefinitionGetAllEscalationLevelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the alert definition get all escalation levels params
func (o *AlertDefinitionGetAllEscalationLevelsParams) WithContext(ctx context.Context) *AlertDefinitionGetAllEscalationLevelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the alert definition get all escalation levels params
func (o *AlertDefinitionGetAllEscalationLevelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the alert definition get all escalation levels params
func (o *AlertDefinitionGetAllEscalationLevelsParams) WithHTTPClient(client *http.Client) *AlertDefinitionGetAllEscalationLevelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the alert definition get all escalation levels params
func (o *AlertDefinitionGetAllEscalationLevelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlertDefinitionGUID adds the alertDefinitionGUID to the alert definition get all escalation levels params
func (o *AlertDefinitionGetAllEscalationLevelsParams) WithAlertDefinitionGUID(alertDefinitionGUID string) *AlertDefinitionGetAllEscalationLevelsParams {
	o.SetAlertDefinitionGUID(alertDefinitionGUID)
	return o
}

// SetAlertDefinitionGUID adds the alertDefinitionGuid to the alert definition get all escalation levels params
func (o *AlertDefinitionGetAllEscalationLevelsParams) SetAlertDefinitionGUID(alertDefinitionGUID string) {
	o.AlertDefinitionGUID = alertDefinitionGUID
}

// WriteToRequest writes these params to a swagger request
func (o *AlertDefinitionGetAllEscalationLevelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
