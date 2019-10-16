// Code generated by go-swagger; DO NOT EDIT.

package operator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewOperatorDeleteDutyScheduleFromOperatorParams creates a new OperatorDeleteDutyScheduleFromOperatorParams object
// with the default values initialized.
func NewOperatorDeleteDutyScheduleFromOperatorParams() *OperatorDeleteDutyScheduleFromOperatorParams {
	var ()
	return &OperatorDeleteDutyScheduleFromOperatorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOperatorDeleteDutyScheduleFromOperatorParamsWithTimeout creates a new OperatorDeleteDutyScheduleFromOperatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperatorDeleteDutyScheduleFromOperatorParamsWithTimeout(timeout time.Duration) *OperatorDeleteDutyScheduleFromOperatorParams {
	var ()
	return &OperatorDeleteDutyScheduleFromOperatorParams{

		timeout: timeout,
	}
}

// NewOperatorDeleteDutyScheduleFromOperatorParamsWithContext creates a new OperatorDeleteDutyScheduleFromOperatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperatorDeleteDutyScheduleFromOperatorParamsWithContext(ctx context.Context) *OperatorDeleteDutyScheduleFromOperatorParams {
	var ()
	return &OperatorDeleteDutyScheduleFromOperatorParams{

		Context: ctx,
	}
}

// NewOperatorDeleteDutyScheduleFromOperatorParamsWithHTTPClient creates a new OperatorDeleteDutyScheduleFromOperatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperatorDeleteDutyScheduleFromOperatorParamsWithHTTPClient(client *http.Client) *OperatorDeleteDutyScheduleFromOperatorParams {
	var ()
	return &OperatorDeleteDutyScheduleFromOperatorParams{
		HTTPClient: client,
	}
}

/*OperatorDeleteDutyScheduleFromOperatorParams contains all the parameters to send to the API endpoint
for the operator delete duty schedule from operator operation typically these are written to a http.Request
*/
type OperatorDeleteDutyScheduleFromOperatorParams struct {

	/*DutyScheduleID*/
	DutyScheduleID int32
	/*OperatorGUID*/
	OperatorGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operator delete duty schedule from operator params
func (o *OperatorDeleteDutyScheduleFromOperatorParams) WithTimeout(timeout time.Duration) *OperatorDeleteDutyScheduleFromOperatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operator delete duty schedule from operator params
func (o *OperatorDeleteDutyScheduleFromOperatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operator delete duty schedule from operator params
func (o *OperatorDeleteDutyScheduleFromOperatorParams) WithContext(ctx context.Context) *OperatorDeleteDutyScheduleFromOperatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operator delete duty schedule from operator params
func (o *OperatorDeleteDutyScheduleFromOperatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operator delete duty schedule from operator params
func (o *OperatorDeleteDutyScheduleFromOperatorParams) WithHTTPClient(client *http.Client) *OperatorDeleteDutyScheduleFromOperatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operator delete duty schedule from operator params
func (o *OperatorDeleteDutyScheduleFromOperatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDutyScheduleID adds the dutyScheduleID to the operator delete duty schedule from operator params
func (o *OperatorDeleteDutyScheduleFromOperatorParams) WithDutyScheduleID(dutyScheduleID int32) *OperatorDeleteDutyScheduleFromOperatorParams {
	o.SetDutyScheduleID(dutyScheduleID)
	return o
}

// SetDutyScheduleID adds the dutyScheduleId to the operator delete duty schedule from operator params
func (o *OperatorDeleteDutyScheduleFromOperatorParams) SetDutyScheduleID(dutyScheduleID int32) {
	o.DutyScheduleID = dutyScheduleID
}

// WithOperatorGUID adds the operatorGUID to the operator delete duty schedule from operator params
func (o *OperatorDeleteDutyScheduleFromOperatorParams) WithOperatorGUID(operatorGUID string) *OperatorDeleteDutyScheduleFromOperatorParams {
	o.SetOperatorGUID(operatorGUID)
	return o
}

// SetOperatorGUID adds the operatorGuid to the operator delete duty schedule from operator params
func (o *OperatorDeleteDutyScheduleFromOperatorParams) SetOperatorGUID(operatorGUID string) {
	o.OperatorGUID = operatorGUID
}

// WriteToRequest writes these params to a swagger request
func (o *OperatorDeleteDutyScheduleFromOperatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dutyScheduleId
	if err := r.SetPathParam("dutyScheduleId", swag.FormatInt32(o.DutyScheduleID)); err != nil {
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
