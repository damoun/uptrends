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

// NewDashboardUpdateDashboardParams creates a new DashboardUpdateDashboardParams object
// with the default values initialized.
func NewDashboardUpdateDashboardParams() *DashboardUpdateDashboardParams {
	var ()
	return &DashboardUpdateDashboardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDashboardUpdateDashboardParamsWithTimeout creates a new DashboardUpdateDashboardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDashboardUpdateDashboardParamsWithTimeout(timeout time.Duration) *DashboardUpdateDashboardParams {
	var ()
	return &DashboardUpdateDashboardParams{

		timeout: timeout,
	}
}

// NewDashboardUpdateDashboardParamsWithContext creates a new DashboardUpdateDashboardParams object
// with the default values initialized, and the ability to set a context for a request
func NewDashboardUpdateDashboardParamsWithContext(ctx context.Context) *DashboardUpdateDashboardParams {
	var ()
	return &DashboardUpdateDashboardParams{

		Context: ctx,
	}
}

// NewDashboardUpdateDashboardParamsWithHTTPClient creates a new DashboardUpdateDashboardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDashboardUpdateDashboardParamsWithHTTPClient(client *http.Client) *DashboardUpdateDashboardParams {
	var ()
	return &DashboardUpdateDashboardParams{
		HTTPClient: client,
	}
}

/*DashboardUpdateDashboardParams contains all the parameters to send to the API endpoint
for the dashboard update dashboard operation typically these are written to a http.Request
*/
type DashboardUpdateDashboardParams struct {

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

// WithTimeout adds the timeout to the dashboard update dashboard params
func (o *DashboardUpdateDashboardParams) WithTimeout(timeout time.Duration) *DashboardUpdateDashboardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dashboard update dashboard params
func (o *DashboardUpdateDashboardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dashboard update dashboard params
func (o *DashboardUpdateDashboardParams) WithContext(ctx context.Context) *DashboardUpdateDashboardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dashboard update dashboard params
func (o *DashboardUpdateDashboardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dashboard update dashboard params
func (o *DashboardUpdateDashboardParams) WithHTTPClient(client *http.Client) *DashboardUpdateDashboardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dashboard update dashboard params
func (o *DashboardUpdateDashboardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDashboard adds the dashboard to the dashboard update dashboard params
func (o *DashboardUpdateDashboardParams) WithDashboard(dashboard *models.Dashboard) *DashboardUpdateDashboardParams {
	o.SetDashboard(dashboard)
	return o
}

// SetDashboard adds the dashboard to the dashboard update dashboard params
func (o *DashboardUpdateDashboardParams) SetDashboard(dashboard *models.Dashboard) {
	o.Dashboard = dashboard
}

// WithDashboardGUID adds the dashboardGUID to the dashboard update dashboard params
func (o *DashboardUpdateDashboardParams) WithDashboardGUID(dashboardGUID string) *DashboardUpdateDashboardParams {
	o.SetDashboardGUID(dashboardGUID)
	return o
}

// SetDashboardGUID adds the dashboardGuid to the dashboard update dashboard params
func (o *DashboardUpdateDashboardParams) SetDashboardGUID(dashboardGUID string) {
	o.DashboardGUID = dashboardGUID
}

// WriteToRequest writes these params to a swagger request
func (o *DashboardUpdateDashboardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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