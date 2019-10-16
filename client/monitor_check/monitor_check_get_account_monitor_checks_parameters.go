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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewMonitorCheckGetAccountMonitorChecksParams creates a new MonitorCheckGetAccountMonitorChecksParams object
// with the default values initialized.
func NewMonitorCheckGetAccountMonitorChecksParams() *MonitorCheckGetAccountMonitorChecksParams {
	var (
		errorLevelDefault = string("NoError")
		sortingDefault    = string("Ascending")
		takeDefault       = int32(100)
	)
	return &MonitorCheckGetAccountMonitorChecksParams{
		ErrorLevel: &errorLevelDefault,
		Sorting:    &sortingDefault,
		Take:       &takeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorCheckGetAccountMonitorChecksParamsWithTimeout creates a new MonitorCheckGetAccountMonitorChecksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorCheckGetAccountMonitorChecksParamsWithTimeout(timeout time.Duration) *MonitorCheckGetAccountMonitorChecksParams {
	var (
		errorLevelDefault = string("NoError")
		sortingDefault    = string("Ascending")
		takeDefault       = int32(100)
	)
	return &MonitorCheckGetAccountMonitorChecksParams{
		ErrorLevel: &errorLevelDefault,
		Sorting:    &sortingDefault,
		Take:       &takeDefault,

		timeout: timeout,
	}
}

// NewMonitorCheckGetAccountMonitorChecksParamsWithContext creates a new MonitorCheckGetAccountMonitorChecksParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorCheckGetAccountMonitorChecksParamsWithContext(ctx context.Context) *MonitorCheckGetAccountMonitorChecksParams {
	var (
		errorLevelDefault = string("NoError")
		sortingDefault    = string("Ascending")
		takeDefault       = int32(100)
	)
	return &MonitorCheckGetAccountMonitorChecksParams{
		ErrorLevel: &errorLevelDefault,
		Sorting:    &sortingDefault,
		Take:       &takeDefault,

		Context: ctx,
	}
}

// NewMonitorCheckGetAccountMonitorChecksParamsWithHTTPClient creates a new MonitorCheckGetAccountMonitorChecksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorCheckGetAccountMonitorChecksParamsWithHTTPClient(client *http.Client) *MonitorCheckGetAccountMonitorChecksParams {
	var (
		errorLevelDefault = string("NoError")
		sortingDefault    = string("Ascending")
		takeDefault       = int32(100)
	)
	return &MonitorCheckGetAccountMonitorChecksParams{
		ErrorLevel: &errorLevelDefault,
		Sorting:    &sortingDefault,
		Take:       &takeDefault,
		HTTPClient: client,
	}
}

/*MonitorCheckGetAccountMonitorChecksParams contains all the parameters to send to the API endpoint
for the monitor check get account monitor checks operation typically these are written to a http.Request
*/
type MonitorCheckGetAccountMonitorChecksParams struct {

	/*Cursor
	  A cursor value that should be used for traversing the dataset.

	*/
	Cursor *string
	/*End
	  The end of a custom period (can't be longer than 90 days)

	*/
	End *strfmt.DateTime
	/*ErrorLevel
	  Error level filter that should be applied. (default = NoError and above)

	*/
	ErrorLevel *string
	/*Period
	  The requested time period. (default = Last24Hours)

	*/
	Period *string
	/*Sorting
	  Sorting direction based on monitor check timestamp. (default = Descending)

	*/
	Sorting *string
	/*Start
	  The start of a custom period (can't be used together with the period parameter)

	*/
	Start *strfmt.DateTime
	/*Take
	  The number of checks to return (default = 100, max = 100)

	*/
	Take *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) WithTimeout(timeout time.Duration) *MonitorCheckGetAccountMonitorChecksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) WithContext(ctx context.Context) *MonitorCheckGetAccountMonitorChecksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) WithHTTPClient(client *http.Client) *MonitorCheckGetAccountMonitorChecksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) WithCursor(cursor *string) *MonitorCheckGetAccountMonitorChecksParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithEnd adds the end to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) WithEnd(end *strfmt.DateTime) *MonitorCheckGetAccountMonitorChecksParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) SetEnd(end *strfmt.DateTime) {
	o.End = end
}

// WithErrorLevel adds the errorLevel to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) WithErrorLevel(errorLevel *string) *MonitorCheckGetAccountMonitorChecksParams {
	o.SetErrorLevel(errorLevel)
	return o
}

// SetErrorLevel adds the errorLevel to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) SetErrorLevel(errorLevel *string) {
	o.ErrorLevel = errorLevel
}

// WithPeriod adds the period to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) WithPeriod(period *string) *MonitorCheckGetAccountMonitorChecksParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) SetPeriod(period *string) {
	o.Period = period
}

// WithSorting adds the sorting to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) WithSorting(sorting *string) *MonitorCheckGetAccountMonitorChecksParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithStart adds the start to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) WithStart(start *strfmt.DateTime) *MonitorCheckGetAccountMonitorChecksParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) SetStart(start *strfmt.DateTime) {
	o.Start = start
}

// WithTake adds the take to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) WithTake(take *int32) *MonitorCheckGetAccountMonitorChecksParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the monitor check get account monitor checks params
func (o *MonitorCheckGetAccountMonitorChecksParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorCheckGetAccountMonitorChecksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	if o.End != nil {

		// query param end
		var qrEnd strfmt.DateTime
		if o.End != nil {
			qrEnd = *o.End
		}
		qEnd := qrEnd.String()
		if qEnd != "" {
			if err := r.SetQueryParam("end", qEnd); err != nil {
				return err
			}
		}

	}

	if o.ErrorLevel != nil {

		// query param errorLevel
		var qrErrorLevel string
		if o.ErrorLevel != nil {
			qrErrorLevel = *o.ErrorLevel
		}
		qErrorLevel := qrErrorLevel
		if qErrorLevel != "" {
			if err := r.SetQueryParam("errorLevel", qErrorLevel); err != nil {
				return err
			}
		}

	}

	if o.Period != nil {

		// query param period
		var qrPeriod string
		if o.Period != nil {
			qrPeriod = *o.Period
		}
		qPeriod := qrPeriod
		if qPeriod != "" {
			if err := r.SetQueryParam("period", qPeriod); err != nil {
				return err
			}
		}

	}

	if o.Sorting != nil {

		// query param sorting
		var qrSorting string
		if o.Sorting != nil {
			qrSorting = *o.Sorting
		}
		qSorting := qrSorting
		if qSorting != "" {
			if err := r.SetQueryParam("sorting", qSorting); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart strfmt.DateTime
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := qrStart.String()
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if o.Take != nil {

		// query param take
		var qrTake int32
		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt32(qrTake)
		if qTake != "" {
			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
