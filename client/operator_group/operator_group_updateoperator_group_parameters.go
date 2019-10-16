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

	models "github.com/damoun/uptrends/models"
)

// NewOperatorGroupUpdateoperatorGroupParams creates a new OperatorGroupUpdateoperatorGroupParams object
// with the default values initialized.
func NewOperatorGroupUpdateoperatorGroupParams() *OperatorGroupUpdateoperatorGroupParams {
	var ()
	return &OperatorGroupUpdateoperatorGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewOperatorGroupUpdateoperatorGroupParamsWithTimeout creates a new OperatorGroupUpdateoperatorGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewOperatorGroupUpdateoperatorGroupParamsWithTimeout(timeout time.Duration) *OperatorGroupUpdateoperatorGroupParams {
	var ()
	return &OperatorGroupUpdateoperatorGroupParams{

		timeout: timeout,
	}
}

// NewOperatorGroupUpdateoperatorGroupParamsWithContext creates a new OperatorGroupUpdateoperatorGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewOperatorGroupUpdateoperatorGroupParamsWithContext(ctx context.Context) *OperatorGroupUpdateoperatorGroupParams {
	var ()
	return &OperatorGroupUpdateoperatorGroupParams{

		Context: ctx,
	}
}

// NewOperatorGroupUpdateoperatorGroupParamsWithHTTPClient creates a new OperatorGroupUpdateoperatorGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewOperatorGroupUpdateoperatorGroupParamsWithHTTPClient(client *http.Client) *OperatorGroupUpdateoperatorGroupParams {
	var ()
	return &OperatorGroupUpdateoperatorGroupParams{
		HTTPClient: client,
	}
}

/*OperatorGroupUpdateoperatorGroupParams contains all the parameters to send to the API endpoint
for the operator group updateoperator group operation typically these are written to a http.Request
*/
type OperatorGroupUpdateoperatorGroupParams struct {

	/*Item
	  The operator group to be updated

	*/
	Item *models.OperatorGroup
	/*OperatorGroupGUID
	  The Guid of the operator group to be updated

	*/
	OperatorGroupGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the operator group updateoperator group params
func (o *OperatorGroupUpdateoperatorGroupParams) WithTimeout(timeout time.Duration) *OperatorGroupUpdateoperatorGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the operator group updateoperator group params
func (o *OperatorGroupUpdateoperatorGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the operator group updateoperator group params
func (o *OperatorGroupUpdateoperatorGroupParams) WithContext(ctx context.Context) *OperatorGroupUpdateoperatorGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the operator group updateoperator group params
func (o *OperatorGroupUpdateoperatorGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the operator group updateoperator group params
func (o *OperatorGroupUpdateoperatorGroupParams) WithHTTPClient(client *http.Client) *OperatorGroupUpdateoperatorGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the operator group updateoperator group params
func (o *OperatorGroupUpdateoperatorGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithItem adds the item to the operator group updateoperator group params
func (o *OperatorGroupUpdateoperatorGroupParams) WithItem(item *models.OperatorGroup) *OperatorGroupUpdateoperatorGroupParams {
	o.SetItem(item)
	return o
}

// SetItem adds the item to the operator group updateoperator group params
func (o *OperatorGroupUpdateoperatorGroupParams) SetItem(item *models.OperatorGroup) {
	o.Item = item
}

// WithOperatorGroupGUID adds the operatorGroupGUID to the operator group updateoperator group params
func (o *OperatorGroupUpdateoperatorGroupParams) WithOperatorGroupGUID(operatorGroupGUID string) *OperatorGroupUpdateoperatorGroupParams {
	o.SetOperatorGroupGUID(operatorGroupGUID)
	return o
}

// SetOperatorGroupGUID adds the operatorGroupGuid to the operator group updateoperator group params
func (o *OperatorGroupUpdateoperatorGroupParams) SetOperatorGroupGUID(operatorGroupGUID string) {
	o.OperatorGroupGUID = operatorGroupGUID
}

// WriteToRequest writes these params to a swagger request
func (o *OperatorGroupUpdateoperatorGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Item != nil {
		if err := r.SetBodyParam(o.Item); err != nil {
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