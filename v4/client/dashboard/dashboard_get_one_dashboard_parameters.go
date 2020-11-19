// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewDashboardGetOneDashboardParams creates a new DashboardGetOneDashboardParams object
// with the default values initialized.
func NewDashboardGetOneDashboardParams() *DashboardGetOneDashboardParams {
	var ()
	return &DashboardGetOneDashboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDashboardGetOneDashboardParamsWithTimeout creates a new DashboardGetOneDashboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDashboardGetOneDashboardParamsWithTimeout(timeout time.Duration) *DashboardGetOneDashboardParams {
	var ()
	return &DashboardGetOneDashboardParams{

		timeout: timeout,
	}
}

// NewDashboardGetOneDashboardParamsWithContext creates a new DashboardGetOneDashboardParams object
// with the default values initialized, and the ability to set a context for a request
func NewDashboardGetOneDashboardParamsWithContext(ctx context.Context) *DashboardGetOneDashboardParams {
	var ()
	return &DashboardGetOneDashboardParams{

		Context: ctx,
	}
}

// NewDashboardGetOneDashboardParamsWithHTTPClient creates a new DashboardGetOneDashboardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDashboardGetOneDashboardParamsWithHTTPClient(client *http.Client) *DashboardGetOneDashboardParams {
	var ()
	return &DashboardGetOneDashboardParams{
		HTTPClient: client,
	}
}

/*DashboardGetOneDashboardParams contains all the parameters to send to the API endpoint
for the dashboard get one dashboard operation typically these are written to a http.Request
*/
type DashboardGetOneDashboardParams struct {

	/*DashboardGUID
	  The guid of the specified dashboard.

	*/
	DashboardGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dashboard get one dashboard params
func (o *DashboardGetOneDashboardParams) WithTimeout(timeout time.Duration) *DashboardGetOneDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dashboard get one dashboard params
func (o *DashboardGetOneDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dashboard get one dashboard params
func (o *DashboardGetOneDashboardParams) WithContext(ctx context.Context) *DashboardGetOneDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dashboard get one dashboard params
func (o *DashboardGetOneDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dashboard get one dashboard params
func (o *DashboardGetOneDashboardParams) WithHTTPClient(client *http.Client) *DashboardGetOneDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dashboard get one dashboard params
func (o *DashboardGetOneDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboardGUID adds the dashboardGUID to the dashboard get one dashboard params
func (o *DashboardGetOneDashboardParams) WithDashboardGUID(dashboardGUID string) *DashboardGetOneDashboardParams {
	o.SetDashboardGUID(dashboardGUID)
	return o
}

// SetDashboardGUID adds the dashboardGuid to the dashboard get one dashboard params
func (o *DashboardGetOneDashboardParams) SetDashboardGUID(dashboardGUID string) {
	o.DashboardGUID = dashboardGUID
}

// WriteToRequest writes these params to a swagger request
func (o *DashboardGetOneDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param dashboardGuid
	if err := r.SetPathParam("dashboardGuid", o.DashboardGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
