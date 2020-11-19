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
	"github.com/go-openapi/strfmt"

	"github.com/damoun/uptrends/v4/models"
)

// NewOperatorUpdateOperatorWithPatchParams creates a new OperatorUpdateOperatorWithPatchParams object
// with the default values initialized.
func NewOperatorUpdateOperatorWithPatchParams() *OperatorUpdateOperatorWithPatchParams {
	var ()
	return &OperatorUpdateOperatorWithPatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOperatorUpdateOperatorWithPatchParamsWithTimeout creates a new OperatorUpdateOperatorWithPatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperatorUpdateOperatorWithPatchParamsWithTimeout(timeout time.Duration) *OperatorUpdateOperatorWithPatchParams {
	var ()
	return &OperatorUpdateOperatorWithPatchParams{

		timeout: timeout,
	}
}

// NewOperatorUpdateOperatorWithPatchParamsWithContext creates a new OperatorUpdateOperatorWithPatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperatorUpdateOperatorWithPatchParamsWithContext(ctx context.Context) *OperatorUpdateOperatorWithPatchParams {
	var ()
	return &OperatorUpdateOperatorWithPatchParams{

		Context: ctx,
	}
}

// NewOperatorUpdateOperatorWithPatchParamsWithHTTPClient creates a new OperatorUpdateOperatorWithPatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperatorUpdateOperatorWithPatchParamsWithHTTPClient(client *http.Client) *OperatorUpdateOperatorWithPatchParams {
	var ()
	return &OperatorUpdateOperatorWithPatchParams{
		HTTPClient: client,
	}
}

/*OperatorUpdateOperatorWithPatchParams contains all the parameters to send to the API endpoint
for the operator update operator with patch operation typically these are written to a http.Request
*/
type OperatorUpdateOperatorWithPatchParams struct {

	/*OperatorGUID
	  The Guid of the operator to update

	*/
	OperatorGUID string
	/*UptrendsOperator
	  The updated details of the operator

	*/
	UptrendsOperator *models.Operator

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operator update operator with patch params
func (o *OperatorUpdateOperatorWithPatchParams) WithTimeout(timeout time.Duration) *OperatorUpdateOperatorWithPatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operator update operator with patch params
func (o *OperatorUpdateOperatorWithPatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operator update operator with patch params
func (o *OperatorUpdateOperatorWithPatchParams) WithContext(ctx context.Context) *OperatorUpdateOperatorWithPatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operator update operator with patch params
func (o *OperatorUpdateOperatorWithPatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operator update operator with patch params
func (o *OperatorUpdateOperatorWithPatchParams) WithHTTPClient(client *http.Client) *OperatorUpdateOperatorWithPatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operator update operator with patch params
func (o *OperatorUpdateOperatorWithPatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperatorGUID adds the operatorGUID to the operator update operator with patch params
func (o *OperatorUpdateOperatorWithPatchParams) WithOperatorGUID(operatorGUID string) *OperatorUpdateOperatorWithPatchParams {
	o.SetOperatorGUID(operatorGUID)
	return o
}

// SetOperatorGUID adds the operatorGuid to the operator update operator with patch params
func (o *OperatorUpdateOperatorWithPatchParams) SetOperatorGUID(operatorGUID string) {
	o.OperatorGUID = operatorGUID
}

// WithUptrendsOperator adds the uptrendsOperator to the operator update operator with patch params
func (o *OperatorUpdateOperatorWithPatchParams) WithUptrendsOperator(uptrendsOperator *models.Operator) *OperatorUpdateOperatorWithPatchParams {
	o.SetUptrendsOperator(uptrendsOperator)
	return o
}

// SetUptrendsOperator adds the uptrendsOperator to the operator update operator with patch params
func (o *OperatorUpdateOperatorWithPatchParams) SetUptrendsOperator(uptrendsOperator *models.Operator) {
	o.UptrendsOperator = uptrendsOperator
}

// WriteToRequest writes these params to a swagger request
func (o *OperatorUpdateOperatorWithPatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param operatorGuid
	if err := r.SetPathParam("operatorGuid", o.OperatorGUID); err != nil {
		return err
	}

	if o.UptrendsOperator != nil {
		if err := r.SetBodyParam(o.UptrendsOperator); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}