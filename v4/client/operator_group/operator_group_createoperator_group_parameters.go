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

// NewOperatorGroupCreateoperatorGroupParams creates a new OperatorGroupCreateoperatorGroupParams object
// with the default values initialized.
func NewOperatorGroupCreateoperatorGroupParams() *OperatorGroupCreateoperatorGroupParams {
	var ()
	return &OperatorGroupCreateoperatorGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOperatorGroupCreateoperatorGroupParamsWithTimeout creates a new OperatorGroupCreateoperatorGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperatorGroupCreateoperatorGroupParamsWithTimeout(timeout time.Duration) *OperatorGroupCreateoperatorGroupParams {
	var ()
	return &OperatorGroupCreateoperatorGroupParams{

		timeout: timeout,
	}
}

// NewOperatorGroupCreateoperatorGroupParamsWithContext creates a new OperatorGroupCreateoperatorGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperatorGroupCreateoperatorGroupParamsWithContext(ctx context.Context) *OperatorGroupCreateoperatorGroupParams {
	var ()
	return &OperatorGroupCreateoperatorGroupParams{

		Context: ctx,
	}
}

// NewOperatorGroupCreateoperatorGroupParamsWithHTTPClient creates a new OperatorGroupCreateoperatorGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperatorGroupCreateoperatorGroupParamsWithHTTPClient(client *http.Client) *OperatorGroupCreateoperatorGroupParams {
	var ()
	return &OperatorGroupCreateoperatorGroupParams{
		HTTPClient: client,
	}
}

/*OperatorGroupCreateoperatorGroupParams contains all the parameters to send to the API endpoint
for the operator group createoperator group operation typically these are written to a http.Request
*/
type OperatorGroupCreateoperatorGroupParams struct {

	/*OperatorGroup
	  The operatorGroup object to be created

	*/
	OperatorGroup *models.OperatorGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operator group createoperator group params
func (o *OperatorGroupCreateoperatorGroupParams) WithTimeout(timeout time.Duration) *OperatorGroupCreateoperatorGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operator group createoperator group params
func (o *OperatorGroupCreateoperatorGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operator group createoperator group params
func (o *OperatorGroupCreateoperatorGroupParams) WithContext(ctx context.Context) *OperatorGroupCreateoperatorGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operator group createoperator group params
func (o *OperatorGroupCreateoperatorGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operator group createoperator group params
func (o *OperatorGroupCreateoperatorGroupParams) WithHTTPClient(client *http.Client) *OperatorGroupCreateoperatorGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operator group createoperator group params
func (o *OperatorGroupCreateoperatorGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperatorGroup adds the operatorGroup to the operator group createoperator group params
func (o *OperatorGroupCreateoperatorGroupParams) WithOperatorGroup(operatorGroup *models.OperatorGroup) *OperatorGroupCreateoperatorGroupParams {
	o.SetOperatorGroup(operatorGroup)
	return o
}

// SetOperatorGroup adds the operatorGroup to the operator group createoperator group params
func (o *OperatorGroupCreateoperatorGroupParams) SetOperatorGroup(operatorGroup *models.OperatorGroup) {
	o.OperatorGroup = operatorGroup
}

// WriteToRequest writes these params to a swagger request
func (o *OperatorGroupCreateoperatorGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.OperatorGroup != nil {
		if err := r.SetBodyParam(o.OperatorGroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
