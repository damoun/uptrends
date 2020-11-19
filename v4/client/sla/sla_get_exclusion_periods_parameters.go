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

// NewSLAGetExclusionPeriodsParams creates a new SLAGetExclusionPeriodsParams object
// with the default values initialized.
func NewSLAGetExclusionPeriodsParams() *SLAGetExclusionPeriodsParams {
	var ()
	return &SLAGetExclusionPeriodsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSLAGetExclusionPeriodsParamsWithTimeout creates a new SLAGetExclusionPeriodsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSLAGetExclusionPeriodsParamsWithTimeout(timeout time.Duration) *SLAGetExclusionPeriodsParams {
	var ()
	return &SLAGetExclusionPeriodsParams{

		timeout: timeout,
	}
}

// NewSLAGetExclusionPeriodsParamsWithContext creates a new SLAGetExclusionPeriodsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSLAGetExclusionPeriodsParamsWithContext(ctx context.Context) *SLAGetExclusionPeriodsParams {
	var ()
	return &SLAGetExclusionPeriodsParams{

		Context: ctx,
	}
}

// NewSLAGetExclusionPeriodsParamsWithHTTPClient creates a new SLAGetExclusionPeriodsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSLAGetExclusionPeriodsParamsWithHTTPClient(client *http.Client) *SLAGetExclusionPeriodsParams {
	var ()
	return &SLAGetExclusionPeriodsParams{
		HTTPClient: client,
	}
}

/*SLAGetExclusionPeriodsParams contains all the parameters to send to the API endpoint
for the Sla get exclusion periods operation typically these are written to a http.Request
*/
type SLAGetExclusionPeriodsParams struct {

	/*SLAGUID
	  The Guid of the SLA definition.

	*/
	SLAGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Sla get exclusion periods params
func (o *SLAGetExclusionPeriodsParams) WithTimeout(timeout time.Duration) *SLAGetExclusionPeriodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Sla get exclusion periods params
func (o *SLAGetExclusionPeriodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Sla get exclusion periods params
func (o *SLAGetExclusionPeriodsParams) WithContext(ctx context.Context) *SLAGetExclusionPeriodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Sla get exclusion periods params
func (o *SLAGetExclusionPeriodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Sla get exclusion periods params
func (o *SLAGetExclusionPeriodsParams) WithHTTPClient(client *http.Client) *SLAGetExclusionPeriodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Sla get exclusion periods params
func (o *SLAGetExclusionPeriodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSLAGUID adds the sLAGUID to the Sla get exclusion periods params
func (o *SLAGetExclusionPeriodsParams) WithSLAGUID(sLAGUID string) *SLAGetExclusionPeriodsParams {
	o.SetSLAGUID(sLAGUID)
	return o
}

// SetSLAGUID adds the slaGuid to the Sla get exclusion periods params
func (o *SLAGetExclusionPeriodsParams) SetSLAGUID(sLAGUID string) {
	o.SLAGUID = sLAGUID
}

// WriteToRequest writes these params to a swagger request
func (o *SLAGetExclusionPeriodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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