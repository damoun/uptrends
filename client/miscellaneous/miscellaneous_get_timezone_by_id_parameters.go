// Code generated by go-swagger; DO NOT EDIT.

package miscellaneous

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

// NewMiscellaneousGetTimezoneByIDParams creates a new MiscellaneousGetTimezoneByIDParams object
// with the default values initialized.
func NewMiscellaneousGetTimezoneByIDParams() *MiscellaneousGetTimezoneByIDParams {
	var ()
	return &MiscellaneousGetTimezoneByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMiscellaneousGetTimezoneByIDParamsWithTimeout creates a new MiscellaneousGetTimezoneByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMiscellaneousGetTimezoneByIDParamsWithTimeout(timeout time.Duration) *MiscellaneousGetTimezoneByIDParams {
	var ()
	return &MiscellaneousGetTimezoneByIDParams{

		timeout: timeout,
	}
}

// NewMiscellaneousGetTimezoneByIDParamsWithContext creates a new MiscellaneousGetTimezoneByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewMiscellaneousGetTimezoneByIDParamsWithContext(ctx context.Context) *MiscellaneousGetTimezoneByIDParams {
	var ()
	return &MiscellaneousGetTimezoneByIDParams{

		Context: ctx,
	}
}

// NewMiscellaneousGetTimezoneByIDParamsWithHTTPClient creates a new MiscellaneousGetTimezoneByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMiscellaneousGetTimezoneByIDParamsWithHTTPClient(client *http.Client) *MiscellaneousGetTimezoneByIDParams {
	var ()
	return &MiscellaneousGetTimezoneByIDParams{
		HTTPClient: client,
	}
}

/*MiscellaneousGetTimezoneByIDParams contains all the parameters to send to the API endpoint
for the miscellaneous get timezone by Id operation typically these are written to a http.Request
*/
type MiscellaneousGetTimezoneByIDParams struct {

	/*TimezoneID*/
	TimezoneID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the miscellaneous get timezone by Id params
func (o *MiscellaneousGetTimezoneByIDParams) WithTimeout(timeout time.Duration) *MiscellaneousGetTimezoneByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the miscellaneous get timezone by Id params
func (o *MiscellaneousGetTimezoneByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the miscellaneous get timezone by Id params
func (o *MiscellaneousGetTimezoneByIDParams) WithContext(ctx context.Context) *MiscellaneousGetTimezoneByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the miscellaneous get timezone by Id params
func (o *MiscellaneousGetTimezoneByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the miscellaneous get timezone by Id params
func (o *MiscellaneousGetTimezoneByIDParams) WithHTTPClient(client *http.Client) *MiscellaneousGetTimezoneByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the miscellaneous get timezone by Id params
func (o *MiscellaneousGetTimezoneByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTimezoneID adds the timezoneID to the miscellaneous get timezone by Id params
func (o *MiscellaneousGetTimezoneByIDParams) WithTimezoneID(timezoneID int64) *MiscellaneousGetTimezoneByIDParams {
	o.SetTimezoneID(timezoneID)
	return o
}

// SetTimezoneID adds the timezoneId to the miscellaneous get timezone by Id params
func (o *MiscellaneousGetTimezoneByIDParams) SetTimezoneID(timezoneID int64) {
	o.TimezoneID = timezoneID
}

// WriteToRequest writes these params to a swagger request
func (o *MiscellaneousGetTimezoneByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param timezoneId
	if err := r.SetPathParam("timezoneId", swag.FormatInt64(o.TimezoneID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}