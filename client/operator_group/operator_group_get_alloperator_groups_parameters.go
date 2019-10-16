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

// NewOperatorGroupGetAlloperatorGroupsParams creates a new OperatorGroupGetAlloperatorGroupsParams object
// with the default values initialized.
func NewOperatorGroupGetAlloperatorGroupsParams() *OperatorGroupGetAlloperatorGroupsParams {

	return &OperatorGroupGetAlloperatorGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOperatorGroupGetAlloperatorGroupsParamsWithTimeout creates a new OperatorGroupGetAlloperatorGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperatorGroupGetAlloperatorGroupsParamsWithTimeout(timeout time.Duration) *OperatorGroupGetAlloperatorGroupsParams {

	return &OperatorGroupGetAlloperatorGroupsParams{

		timeout: timeout,
	}
}

// NewOperatorGroupGetAlloperatorGroupsParamsWithContext creates a new OperatorGroupGetAlloperatorGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperatorGroupGetAlloperatorGroupsParamsWithContext(ctx context.Context) *OperatorGroupGetAlloperatorGroupsParams {

	return &OperatorGroupGetAlloperatorGroupsParams{

		Context: ctx,
	}
}

// NewOperatorGroupGetAlloperatorGroupsParamsWithHTTPClient creates a new OperatorGroupGetAlloperatorGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperatorGroupGetAlloperatorGroupsParamsWithHTTPClient(client *http.Client) *OperatorGroupGetAlloperatorGroupsParams {

	return &OperatorGroupGetAlloperatorGroupsParams{
		HTTPClient: client,
	}
}

/*OperatorGroupGetAlloperatorGroupsParams contains all the parameters to send to the API endpoint
for the operator group get alloperator groups operation typically these are written to a http.Request
*/
type OperatorGroupGetAlloperatorGroupsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operator group get alloperator groups params
func (o *OperatorGroupGetAlloperatorGroupsParams) WithTimeout(timeout time.Duration) *OperatorGroupGetAlloperatorGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operator group get alloperator groups params
func (o *OperatorGroupGetAlloperatorGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operator group get alloperator groups params
func (o *OperatorGroupGetAlloperatorGroupsParams) WithContext(ctx context.Context) *OperatorGroupGetAlloperatorGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operator group get alloperator groups params
func (o *OperatorGroupGetAlloperatorGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operator group get alloperator groups params
func (o *OperatorGroupGetAlloperatorGroupsParams) WithHTTPClient(client *http.Client) *OperatorGroupGetAlloperatorGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operator group get alloperator groups params
func (o *OperatorGroupGetAlloperatorGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OperatorGroupGetAlloperatorGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
