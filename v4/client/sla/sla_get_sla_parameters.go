// Code generated by go-swagger; DO NOT EDIT.

package sla

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

// NewSLAGetSLAParams creates a new SLAGetSLAParams object
// with the default values initialized.
func NewSLAGetSLAParams() *SLAGetSLAParams {
	var ()
	return &SLAGetSLAParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSLAGetSLAParamsWithTimeout creates a new SLAGetSLAParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSLAGetSLAParamsWithTimeout(timeout time.Duration) *SLAGetSLAParams {
	var ()
	return &SLAGetSLAParams{

		timeout: timeout,
	}
}

// NewSLAGetSLAParamsWithContext creates a new SLAGetSLAParams object
// with the default values initialized, and the ability to set a context for a request
func NewSLAGetSLAParamsWithContext(ctx context.Context) *SLAGetSLAParams {
	var ()
	return &SLAGetSLAParams{

		Context: ctx,
	}
}

// NewSLAGetSLAParamsWithHTTPClient creates a new SLAGetSLAParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSLAGetSLAParamsWithHTTPClient(client *http.Client) *SLAGetSLAParams {
	var ()
	return &SLAGetSLAParams{
		HTTPClient: client,
	}
}

/*SLAGetSLAParams contains all the parameters to send to the API endpoint
for the Sla get Sla operation typically these are written to a http.Request
*/
type SLAGetSLAParams struct {

	/*SLAGUID
	  The Guid of the SLA definition.

	*/
	SLAGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sla get Sla params
func (o *SLAGetSLAParams) WithTimeout(timeout time.Duration) *SLAGetSLAParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sla get Sla params
func (o *SLAGetSLAParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sla get Sla params
func (o *SLAGetSLAParams) WithContext(ctx context.Context) *SLAGetSLAParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sla get Sla params
func (o *SLAGetSLAParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sla get Sla params
func (o *SLAGetSLAParams) WithHTTPClient(client *http.Client) *SLAGetSLAParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sla get Sla params
func (o *SLAGetSLAParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSLAGUID adds the sLAGUID to the Sla get Sla params
func (o *SLAGetSLAParams) WithSLAGUID(sLAGUID string) *SLAGetSLAParams {
	o.SetSLAGUID(sLAGUID)
	return o
}

// SetSLAGUID adds the slaGuid to the Sla get Sla params
func (o *SLAGetSLAParams) SetSLAGUID(sLAGUID string) {
	o.SLAGUID = sLAGUID
}

// WriteToRequest writes these params to a swagger request
func (o *SLAGetSLAParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param slaGuid
	if err := r.SetPathParam("slaGuid", o.SLAGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
