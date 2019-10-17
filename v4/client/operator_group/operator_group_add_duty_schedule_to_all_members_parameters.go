// Code generated by go-swagger; DO NOT EDIT.

package operator_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/damoun/uptrends/v4/models"
)

// NewOperatorGroupAddDutyScheduleToAllMembersParams creates a new OperatorGroupAddDutyScheduleToAllMembersParams object
// with the default values initialized.
func NewOperatorGroupAddDutyScheduleToAllMembersParams() *OperatorGroupAddDutyScheduleToAllMembersParams {
	var ()
	return &OperatorGroupAddDutyScheduleToAllMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOperatorGroupAddDutyScheduleToAllMembersParamsWithTimeout creates a new OperatorGroupAddDutyScheduleToAllMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperatorGroupAddDutyScheduleToAllMembersParamsWithTimeout(timeout time.Duration) *OperatorGroupAddDutyScheduleToAllMembersParams {
	var ()
	return &OperatorGroupAddDutyScheduleToAllMembersParams{

		timeout: timeout,
	}
}

// NewOperatorGroupAddDutyScheduleToAllMembersParamsWithContext creates a new OperatorGroupAddDutyScheduleToAllMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperatorGroupAddDutyScheduleToAllMembersParamsWithContext(ctx context.Context) *OperatorGroupAddDutyScheduleToAllMembersParams {
	var ()
	return &OperatorGroupAddDutyScheduleToAllMembersParams{

		Context: ctx,
	}
}

// NewOperatorGroupAddDutyScheduleToAllMembersParamsWithHTTPClient creates a new OperatorGroupAddDutyScheduleToAllMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperatorGroupAddDutyScheduleToAllMembersParamsWithHTTPClient(client *http.Client) *OperatorGroupAddDutyScheduleToAllMembersParams {
	var ()
	return &OperatorGroupAddDutyScheduleToAllMembersParams{
		HTTPClient: client,
	}
}

/*OperatorGroupAddDutyScheduleToAllMembersParams contains all the parameters to send to the API endpoint
for the operator group add duty schedule to all members operation typically these are written to a http.Request
*/
type OperatorGroupAddDutyScheduleToAllMembersParams struct {

	/*DutySchedule*/
	DutySchedule *models.OperatorDutySchedule
	/*OperatorGroupGUID*/
	OperatorGroupGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operator group add duty schedule to all members params
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) WithTimeout(timeout time.Duration) *OperatorGroupAddDutyScheduleToAllMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operator group add duty schedule to all members params
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operator group add duty schedule to all members params
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) WithContext(ctx context.Context) *OperatorGroupAddDutyScheduleToAllMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operator group add duty schedule to all members params
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operator group add duty schedule to all members params
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) WithHTTPClient(client *http.Client) *OperatorGroupAddDutyScheduleToAllMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operator group add duty schedule to all members params
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDutySchedule adds the dutySchedule to the operator group add duty schedule to all members params
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) WithDutySchedule(dutySchedule *models.OperatorDutySchedule) *OperatorGroupAddDutyScheduleToAllMembersParams {
	o.SetDutySchedule(dutySchedule)
	return o
}

// SetDutySchedule adds the dutySchedule to the operator group add duty schedule to all members params
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) SetDutySchedule(dutySchedule *models.OperatorDutySchedule) {
	o.DutySchedule = dutySchedule
}

// WithOperatorGroupGUID adds the operatorGroupGUID to the operator group add duty schedule to all members params
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) WithOperatorGroupGUID(operatorGroupGUID string) *OperatorGroupAddDutyScheduleToAllMembersParams {
	o.SetOperatorGroupGUID(operatorGroupGUID)
	return o
}

// SetOperatorGroupGUID adds the operatorGroupGuid to the operator group add duty schedule to all members params
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) SetOperatorGroupGUID(operatorGroupGUID string) {
	o.OperatorGroupGUID = operatorGroupGUID
}

// WriteToRequest writes these params to a swagger request
func (o *OperatorGroupAddDutyScheduleToAllMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DutySchedule != nil {
		if err := r.SetBodyParam(o.DutySchedule); err != nil {
			return err
		}
	}

	// path param operatorGroupGuid
	if err := r.SetPathParam("operatorGroupGuid", o.OperatorGroupGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}