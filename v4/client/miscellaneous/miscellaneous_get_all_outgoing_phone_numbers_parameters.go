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
	"github.com/go-openapi/strfmt"
)

// NewMiscellaneousGetAllOutgoingPhoneNumbersParams creates a new MiscellaneousGetAllOutgoingPhoneNumbersParams object
// with the default values initialized.
func NewMiscellaneousGetAllOutgoingPhoneNumbersParams() *MiscellaneousGetAllOutgoingPhoneNumbersParams {

	return &MiscellaneousGetAllOutgoingPhoneNumbersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMiscellaneousGetAllOutgoingPhoneNumbersParamsWithTimeout creates a new MiscellaneousGetAllOutgoingPhoneNumbersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMiscellaneousGetAllOutgoingPhoneNumbersParamsWithTimeout(timeout time.Duration) *MiscellaneousGetAllOutgoingPhoneNumbersParams {

	return &MiscellaneousGetAllOutgoingPhoneNumbersParams{

		timeout: timeout,
	}
}

// NewMiscellaneousGetAllOutgoingPhoneNumbersParamsWithContext creates a new MiscellaneousGetAllOutgoingPhoneNumbersParams object
// with the default values initialized, and the ability to set a context for a request
func NewMiscellaneousGetAllOutgoingPhoneNumbersParamsWithContext(ctx context.Context) *MiscellaneousGetAllOutgoingPhoneNumbersParams {

	return &MiscellaneousGetAllOutgoingPhoneNumbersParams{

		Context: ctx,
	}
}

// NewMiscellaneousGetAllOutgoingPhoneNumbersParamsWithHTTPClient creates a new MiscellaneousGetAllOutgoingPhoneNumbersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMiscellaneousGetAllOutgoingPhoneNumbersParamsWithHTTPClient(client *http.Client) *MiscellaneousGetAllOutgoingPhoneNumbersParams {

	return &MiscellaneousGetAllOutgoingPhoneNumbersParams{
		HTTPClient: client,
	}
}

/*MiscellaneousGetAllOutgoingPhoneNumbersParams contains all the parameters to send to the API endpoint
for the miscellaneous get all outgoing phone numbers operation typically these are written to a http.Request
*/
type MiscellaneousGetAllOutgoingPhoneNumbersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the miscellaneous get all outgoing phone numbers params
func (o *MiscellaneousGetAllOutgoingPhoneNumbersParams) WithTimeout(timeout time.Duration) *MiscellaneousGetAllOutgoingPhoneNumbersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the miscellaneous get all outgoing phone numbers params
func (o *MiscellaneousGetAllOutgoingPhoneNumbersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the miscellaneous get all outgoing phone numbers params
func (o *MiscellaneousGetAllOutgoingPhoneNumbersParams) WithContext(ctx context.Context) *MiscellaneousGetAllOutgoingPhoneNumbersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the miscellaneous get all outgoing phone numbers params
func (o *MiscellaneousGetAllOutgoingPhoneNumbersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the miscellaneous get all outgoing phone numbers params
func (o *MiscellaneousGetAllOutgoingPhoneNumbersParams) WithHTTPClient(client *http.Client) *MiscellaneousGetAllOutgoingPhoneNumbersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the miscellaneous get all outgoing phone numbers params
func (o *MiscellaneousGetAllOutgoingPhoneNumbersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *MiscellaneousGetAllOutgoingPhoneNumbersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
