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

	"github.com/damoun/uptrends/v4/models"
)

// NewDashboardPartiallyUpdateDashboardParams creates a new DashboardPartiallyUpdateDashboardParams object
// with the default values initialized.
func NewDashboardPartiallyUpdateDashboardParams() *DashboardPartiallyUpdateDashboardParams {
	var ()
	return &DashboardPartiallyUpdateDashboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDashboardPartiallyUpdateDashboardParamsWithTimeout creates a new DashboardPartiallyUpdateDashboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDashboardPartiallyUpdateDashboardParamsWithTimeout(timeout time.Duration) *DashboardPartiallyUpdateDashboardParams {
	var ()
	return &DashboardPartiallyUpdateDashboardParams{

		timeout: timeout,
	}
}

// NewDashboardPartiallyUpdateDashboardParamsWithContext creates a new DashboardPartiallyUpdateDashboardParams object
// with the default values initialized, and the ability to set a context for a request
func NewDashboardPartiallyUpdateDashboardParamsWithContext(ctx context.Context) *DashboardPartiallyUpdateDashboardParams {
	var ()
	return &DashboardPartiallyUpdateDashboardParams{

		Context: ctx,
	}
}

// NewDashboardPartiallyUpdateDashboardParamsWithHTTPClient creates a new DashboardPartiallyUpdateDashboardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDashboardPartiallyUpdateDashboardParamsWithHTTPClient(client *http.Client) *DashboardPartiallyUpdateDashboardParams {
	var ()
	return &DashboardPartiallyUpdateDashboardParams{
		HTTPClient: client,
	}
}

/*DashboardPartiallyUpdateDashboardParams contains all the parameters to send to the API endpoint
for the dashboard partially update dashboard operation typically these are written to a http.Request
*/
type DashboardPartiallyUpdateDashboardParams struct {

	/*Dashboard
	  The details for the dashboard.

	*/
	Dashboard *models.Dashboard
	/*DashboardGUID
	  The guid of the specified dashboard.

	*/
	DashboardGUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dashboard partially update dashboard params
func (o *DashboardPartiallyUpdateDashboardParams) WithTimeout(timeout time.Duration) *DashboardPartiallyUpdateDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dashboard partially update dashboard params
func (o *DashboardPartiallyUpdateDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dashboard partially update dashboard params
func (o *DashboardPartiallyUpdateDashboardParams) WithContext(ctx context.Context) *DashboardPartiallyUpdateDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dashboard partially update dashboard params
func (o *DashboardPartiallyUpdateDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dashboard partially update dashboard params
func (o *DashboardPartiallyUpdateDashboardParams) WithHTTPClient(client *http.Client) *DashboardPartiallyUpdateDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dashboard partially update dashboard params
func (o *DashboardPartiallyUpdateDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboard adds the dashboard to the dashboard partially update dashboard params
func (o *DashboardPartiallyUpdateDashboardParams) WithDashboard(dashboard *models.Dashboard) *DashboardPartiallyUpdateDashboardParams {
	o.SetDashboard(dashboard)
	return o
}

// SetDashboard adds the dashboard to the dashboard partially update dashboard params
func (o *DashboardPartiallyUpdateDashboardParams) SetDashboard(dashboard *models.Dashboard) {
	o.Dashboard = dashboard
}

// WithDashboardGUID adds the dashboardGUID to the dashboard partially update dashboard params
func (o *DashboardPartiallyUpdateDashboardParams) WithDashboardGUID(dashboardGUID string) *DashboardPartiallyUpdateDashboardParams {
	o.SetDashboardGUID(dashboardGUID)
	return o
}

// SetDashboardGUID adds the dashboardGuid to the dashboard partially update dashboard params
func (o *DashboardPartiallyUpdateDashboardParams) SetDashboardGUID(dashboardGUID string) {
	o.DashboardGUID = dashboardGUID
}

// WriteToRequest writes these params to a swagger request
func (o *DashboardPartiallyUpdateDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Dashboard != nil {
		if err := r.SetBodyParam(o.Dashboard); err != nil {
			return err
		}
	}

	// path param dashboardGuid
	if err := r.SetPathParam("dashboardGuid", o.DashboardGUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}