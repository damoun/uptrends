// Code generated by go-swagger; DO NOT EDIT.

package monitor_check

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
	"github.com/go-openapi/swag"
)

// NewMonitorCheckGetScreenshotsParams creates a new MonitorCheckGetScreenshotsParams object
// with the default values initialized.
func NewMonitorCheckGetScreenshotsParams() *MonitorCheckGetScreenshotsParams {
	var ()
	return &MonitorCheckGetScreenshotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorCheckGetScreenshotsParamsWithTimeout creates a new MonitorCheckGetScreenshotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorCheckGetScreenshotsParamsWithTimeout(timeout time.Duration) *MonitorCheckGetScreenshotsParams {
	var ()
	return &MonitorCheckGetScreenshotsParams{

		timeout: timeout,
	}
}

// NewMonitorCheckGetScreenshotsParamsWithContext creates a new MonitorCheckGetScreenshotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorCheckGetScreenshotsParamsWithContext(ctx context.Context) *MonitorCheckGetScreenshotsParams {
	var ()
	return &MonitorCheckGetScreenshotsParams{

		Context: ctx,
	}
}

// NewMonitorCheckGetScreenshotsParamsWithHTTPClient creates a new MonitorCheckGetScreenshotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorCheckGetScreenshotsParamsWithHTTPClient(client *http.Client) *MonitorCheckGetScreenshotsParams {
	var ()
	return &MonitorCheckGetScreenshotsParams{
		HTTPClient: client,
	}
}

/*MonitorCheckGetScreenshotsParams contains all the parameters to send to the API endpoint
for the monitor check get screenshots operation typically these are written to a http.Request
*/
type MonitorCheckGetScreenshotsParams struct {

	/*MonitorCheckID
	  The monitor check Id to get the screenshot data for.

	*/
	MonitorCheckID int64
	/*ScreenshotID
	  The screenshot Id of the screenshot to get.

	*/
	ScreenshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the monitor check get screenshots params
func (o *MonitorCheckGetScreenshotsParams) WithTimeout(timeout time.Duration) *MonitorCheckGetScreenshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor check get screenshots params
func (o *MonitorCheckGetScreenshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor check get screenshots params
func (o *MonitorCheckGetScreenshotsParams) WithContext(ctx context.Context) *MonitorCheckGetScreenshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor check get screenshots params
func (o *MonitorCheckGetScreenshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor check get screenshots params
func (o *MonitorCheckGetScreenshotsParams) WithHTTPClient(client *http.Client) *MonitorCheckGetScreenshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor check get screenshots params
func (o *MonitorCheckGetScreenshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMonitorCheckID adds the monitorCheckID to the monitor check get screenshots params
func (o *MonitorCheckGetScreenshotsParams) WithMonitorCheckID(monitorCheckID int64) *MonitorCheckGetScreenshotsParams {
	o.SetMonitorCheckID(monitorCheckID)
	return o
}

// SetMonitorCheckID adds the monitorCheckId to the monitor check get screenshots params
func (o *MonitorCheckGetScreenshotsParams) SetMonitorCheckID(monitorCheckID int64) {
	o.MonitorCheckID = monitorCheckID
}

// WithScreenshotID adds the screenshotID to the monitor check get screenshots params
func (o *MonitorCheckGetScreenshotsParams) WithScreenshotID(screenshotID string) *MonitorCheckGetScreenshotsParams {
	o.SetScreenshotID(screenshotID)
	return o
}

// SetScreenshotID adds the screenshotId to the monitor check get screenshots params
func (o *MonitorCheckGetScreenshotsParams) SetScreenshotID(screenshotID string) {
	o.ScreenshotID = screenshotID
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorCheckGetScreenshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param monitorCheckId
	if err := r.SetPathParam("monitorCheckId", swag.FormatInt64(o.MonitorCheckID)); err != nil {
		return err
	}

	// path param screenshotId
	if err := r.SetPathParam("screenshotId", o.ScreenshotID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
