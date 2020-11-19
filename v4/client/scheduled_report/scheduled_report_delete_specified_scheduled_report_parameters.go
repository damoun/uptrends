// Code generated by go-swagger; DO NOT EDIT.

package scheduled_report

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

// NewScheduledReportDeleteSpecifiedScheduledReportParams creates a new ScheduledReportDeleteSpecifiedScheduledReportParams object
// with the default values initialized.
func NewScheduledReportDeleteSpecifiedScheduledReportParams() *ScheduledReportDeleteSpecifiedScheduledReportParams {
	var ()
	return &ScheduledReportDeleteSpecifiedScheduledReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewScheduledReportDeleteSpecifiedScheduledReportParamsWithTimeout creates a new ScheduledReportDeleteSpecifiedScheduledReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewScheduledReportDeleteSpecifiedScheduledReportParamsWithTimeout(timeout time.Duration) *ScheduledReportDeleteSpecifiedScheduledReportParams {
	var ()
	return &ScheduledReportDeleteSpecifiedScheduledReportParams{

		timeout: timeout,
	}
}

// NewScheduledReportDeleteSpecifiedScheduledReportParamsWithContext creates a new ScheduledReportDeleteSpecifiedScheduledReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewScheduledReportDeleteSpecifiedScheduledReportParamsWithContext(ctx context.Context) *ScheduledReportDeleteSpecifiedScheduledReportParams {
	var ()
	return &ScheduledReportDeleteSpecifiedScheduledReportParams{

		Context: ctx,
	}
}

// NewScheduledReportDeleteSpecifiedScheduledReportParamsWithHTTPClient creates a new ScheduledReportDeleteSpecifiedScheduledReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewScheduledReportDeleteSpecifiedScheduledReportParamsWithHTTPClient(client *http.Client) *ScheduledReportDeleteSpecifiedScheduledReportParams {
	var ()
	return &ScheduledReportDeleteSpecifiedScheduledReportParams{
		HTTPClient: client,
	}
}

/*ScheduledReportDeleteSpecifiedScheduledReportParams contains all the parameters to send to the API endpoint
for the scheduled report delete specified scheduled report operation typically these are written to a http.Request
*/
type ScheduledReportDeleteSpecifiedScheduledReportParams struct {

	/*ScheduledReportGUID
	  The guid of the specified scheduled report.

	*/
	ScheduledReportGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the scheduled report delete specified scheduled report params
func (o *ScheduledReportDeleteSpecifiedScheduledReportParams) WithTimeout(timeout time.Duration) *ScheduledReportDeleteSpecifiedScheduledReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scheduled report delete specified scheduled report params
func (o *ScheduledReportDeleteSpecifiedScheduledReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scheduled report delete specified scheduled report params
func (o *ScheduledReportDeleteSpecifiedScheduledReportParams) WithContext(ctx context.Context) *ScheduledReportDeleteSpecifiedScheduledReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scheduled report delete specified scheduled report params
func (o *ScheduledReportDeleteSpecifiedScheduledReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scheduled report delete specified scheduled report params
func (o *ScheduledReportDeleteSpecifiedScheduledReportParams) WithHTTPClient(client *http.Client) *ScheduledReportDeleteSpecifiedScheduledReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scheduled report delete specified scheduled report params
func (o *ScheduledReportDeleteSpecifiedScheduledReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithScheduledReportGUID adds the scheduledReportGUID to the scheduled report delete specified scheduled report params
func (o *ScheduledReportDeleteSpecifiedScheduledReportParams) WithScheduledReportGUID(scheduledReportGUID string) *ScheduledReportDeleteSpecifiedScheduledReportParams {
	o.SetScheduledReportGUID(scheduledReportGUID)
	return o
}

// SetScheduledReportGUID adds the scheduledReportGuid to the scheduled report delete specified scheduled report params
func (o *ScheduledReportDeleteSpecifiedScheduledReportParams) SetScheduledReportGUID(scheduledReportGUID string) {
	o.ScheduledReportGUID = scheduledReportGUID
}

// WriteToRequest writes these params to a swagger request
func (o *ScheduledReportDeleteSpecifiedScheduledReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param scheduledReportGuid
	if err := r.SetPathParam("scheduledReportGuid", o.ScheduledReportGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
