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
)

// NewOperatorGroupRemoveoperatorFromoperatorGroupParams creates a new OperatorGroupRemoveoperatorFromoperatorGroupParams object
// with the default values initialized.
func NewOperatorGroupRemoveoperatorFromoperatorGroupParams() *OperatorGroupRemoveoperatorFromoperatorGroupParams {
	var ()
	return &OperatorGroupRemoveoperatorFromoperatorGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOperatorGroupRemoveoperatorFromoperatorGroupParamsWithTimeout creates a new OperatorGroupRemoveoperatorFromoperatorGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperatorGroupRemoveoperatorFromoperatorGroupParamsWithTimeout(timeout time.Duration) *OperatorGroupRemoveoperatorFromoperatorGroupParams {
	var ()
	return &OperatorGroupRemoveoperatorFromoperatorGroupParams{

		timeout: timeout,
	}
}

// NewOperatorGroupRemoveoperatorFromoperatorGroupParamsWithContext creates a new OperatorGroupRemoveoperatorFromoperatorGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperatorGroupRemoveoperatorFromoperatorGroupParamsWithContext(ctx context.Context) *OperatorGroupRemoveoperatorFromoperatorGroupParams {
	var ()
	return &OperatorGroupRemoveoperatorFromoperatorGroupParams{

		Context: ctx,
	}
}

// NewOperatorGroupRemoveoperatorFromoperatorGroupParamsWithHTTPClient creates a new OperatorGroupRemoveoperatorFromoperatorGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperatorGroupRemoveoperatorFromoperatorGroupParamsWithHTTPClient(client *http.Client) *OperatorGroupRemoveoperatorFromoperatorGroupParams {
	var ()
	return &OperatorGroupRemoveoperatorFromoperatorGroupParams{
		HTTPClient: client,
	}
}

/*OperatorGroupRemoveoperatorFromoperatorGroupParams contains all the parameters to send to the API endpoint
for the operator group removeoperator fromoperator group operation typically these are written to a http.Request
*/
type OperatorGroupRemoveoperatorFromoperatorGroupParams struct {

	/*OperatorGroupGUID
	  The Guid of the operator group to remove the operator from

	*/
	OperatorGroupGUID string
	/*OperatorGUID
	  The operator Guid to be removed

	*/
	OperatorGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operator group removeoperator fromoperator group params
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) WithTimeout(timeout time.Duration) *OperatorGroupRemoveoperatorFromoperatorGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operator group removeoperator fromoperator group params
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operator group removeoperator fromoperator group params
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) WithContext(ctx context.Context) *OperatorGroupRemoveoperatorFromoperatorGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operator group removeoperator fromoperator group params
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operator group removeoperator fromoperator group params
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) WithHTTPClient(client *http.Client) *OperatorGroupRemoveoperatorFromoperatorGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operator group removeoperator fromoperator group params
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperatorGroupGUID adds the operatorGroupGUID to the operator group removeoperator fromoperator group params
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) WithOperatorGroupGUID(operatorGroupGUID string) *OperatorGroupRemoveoperatorFromoperatorGroupParams {
	o.SetOperatorGroupGUID(operatorGroupGUID)
	return o
}

// SetOperatorGroupGUID adds the operatorGroupGuid to the operator group removeoperator fromoperator group params
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) SetOperatorGroupGUID(operatorGroupGUID string) {
	o.OperatorGroupGUID = operatorGroupGUID
}

// WithOperatorGUID adds the operatorGUID to the operator group removeoperator fromoperator group params
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) WithOperatorGUID(operatorGUID string) *OperatorGroupRemoveoperatorFromoperatorGroupParams {
	o.SetOperatorGUID(operatorGUID)
	return o
}

// SetOperatorGUID adds the operatorGuid to the operator group removeoperator fromoperator group params
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) SetOperatorGUID(operatorGUID string) {
	o.OperatorGUID = operatorGUID
}

// WriteToRequest writes these params to a swagger request
func (o *OperatorGroupRemoveoperatorFromoperatorGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param operatorGroupGuid
	if err := r.SetPathParam("operatorGroupGuid", o.OperatorGroupGUID); err != nil {
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
