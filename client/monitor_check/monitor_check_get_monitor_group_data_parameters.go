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

// NewMonitorCheckGetMonitorGroupDataParams creates a new MonitorCheckGetMonitorGroupDataParams object
// with the default values initialized.
func NewMonitorCheckGetMonitorGroupDataParams() *MonitorCheckGetMonitorGroupDataParams {
	var (
		errorLevelDefault = string("NoError")
		sortingDefault    = string("Ascending")
		takeDefault       = int32(100)
	)
	return &MonitorCheckGetMonitorGroupDataParams{
		ErrorLevel: &errorLevelDefault,
		Sorting:    &sortingDefault,
		Take:       &takeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewMonitorCheckGetMonitorGroupDataParamsWithTimeout creates a new MonitorCheckGetMonitorGroupDataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMonitorCheckGetMonitorGroupDataParamsWithTimeout(timeout time.Duration) *MonitorCheckGetMonitorGroupDataParams {
	var (
		errorLevelDefault = string("NoError")
		sortingDefault    = string("Ascending")
		takeDefault       = int32(100)
	)
	return &MonitorCheckGetMonitorGroupDataParams{
		ErrorLevel: &errorLevelDefault,
		Sorting:    &sortingDefault,
		Take:       &takeDefault,

		timeout: timeout,
	}
}

// NewMonitorCheckGetMonitorGroupDataParamsWithContext creates a new MonitorCheckGetMonitorGroupDataParams object
// with the default values initialized, and the ability to set a context for a request
func NewMonitorCheckGetMonitorGroupDataParamsWithContext(ctx context.Context) *MonitorCheckGetMonitorGroupDataParams {
	var (
		errorLevelDefault = string("NoError")
		sortingDefault    = string("Ascending")
		takeDefault       = int32(100)
	)
	return &MonitorCheckGetMonitorGroupDataParams{
		ErrorLevel: &errorLevelDefault,
		Sorting:    &sortingDefault,
		Take:       &takeDefault,

		Context: ctx,
	}
}

// NewMonitorCheckGetMonitorGroupDataParamsWithHTTPClient creates a new MonitorCheckGetMonitorGroupDataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMonitorCheckGetMonitorGroupDataParamsWithHTTPClient(client *http.Client) *MonitorCheckGetMonitorGroupDataParams {
	var (
		errorLevelDefault = string("NoError")
		sortingDefault    = string("Ascending")
		takeDefault       = int32(100)
	)
	return &MonitorCheckGetMonitorGroupDataParams{
		ErrorLevel: &errorLevelDefault,
		Sorting:    &sortingDefault,
		Take:       &takeDefault,
		HTTPClient: client,
	}
}

/*MonitorCheckGetMonitorGroupDataParams contains all the parameters to send to the API endpoint
for the monitor check get monitor group data operation typically these are written to a http.Request
*/
type MonitorCheckGetMonitorGroupDataParams struct {

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
	/*MonitorGroupGUID
	  The Guid of the specified monitor group.

	*/
	MonitorGroupGUID string
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

// WithTimeout adds the timeout to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithTimeout(timeout time.Duration) *MonitorCheckGetMonitorGroupDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithContext(ctx context.Context) *MonitorCheckGetMonitorGroupDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithHTTPClient(client *http.Client) *MonitorCheckGetMonitorGroupDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCursor adds the cursor to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithCursor(cursor *string) *MonitorCheckGetMonitorGroupDataParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithEnd adds the end to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithEnd(end *strfmt.DateTime) *MonitorCheckGetMonitorGroupDataParams {
	o.SetEnd(end)
	return o
}

// SetEnd adds the end to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetEnd(end *strfmt.DateTime) {
	o.End = end
}

// WithErrorLevel adds the errorLevel to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithErrorLevel(errorLevel *string) *MonitorCheckGetMonitorGroupDataParams {
	o.SetErrorLevel(errorLevel)
	return o
}

// SetErrorLevel adds the errorLevel to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetErrorLevel(errorLevel *string) {
	o.ErrorLevel = errorLevel
}

// WithMonitorGroupGUID adds the monitorGroupGUID to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithMonitorGroupGUID(monitorGroupGUID string) *MonitorCheckGetMonitorGroupDataParams {
	o.SetMonitorGroupGUID(monitorGroupGUID)
	return o
}

// SetMonitorGroupGUID adds the monitorGroupGuid to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetMonitorGroupGUID(monitorGroupGUID string) {
	o.MonitorGroupGUID = monitorGroupGUID
}

// WithPeriod adds the period to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithPeriod(period *string) *MonitorCheckGetMonitorGroupDataParams {
	o.SetPeriod(period)
	return o
}

// SetPeriod adds the period to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetPeriod(period *string) {
	o.Period = period
}

// WithSorting adds the sorting to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithSorting(sorting *string) *MonitorCheckGetMonitorGroupDataParams {
	o.SetSorting(sorting)
	return o
}

// SetSorting adds the sorting to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetSorting(sorting *string) {
	o.Sorting = sorting
}

// WithStart adds the start to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithStart(start *strfmt.DateTime) *MonitorCheckGetMonitorGroupDataParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetStart(start *strfmt.DateTime) {
	o.Start = start
}

// WithTake adds the take to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) WithTake(take *int32) *MonitorCheckGetMonitorGroupDataParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the monitor check get monitor group data params
func (o *MonitorCheckGetMonitorGroupDataParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *MonitorCheckGetMonitorGroupDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param monitorGroupGuid
	if err := r.SetPathParam("monitorGroupGuid", o.MonitorGroupGUID); err != nil {
		return err
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
