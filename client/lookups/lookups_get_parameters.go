package lookups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewLookupsGetParams creates a new LookupsGetParams object
// with the default values initialized.
func NewLookupsGetParams() *LookupsGetParams {
	var ()
	return &LookupsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLookupsGetParamsWithTimeout creates a new LookupsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLookupsGetParamsWithTimeout(timeout time.Duration) *LookupsGetParams {
	var ()
	return &LookupsGetParams{

		timeout: timeout,
	}
}

/*LookupsGetParams contains all the parameters to send to the API endpoint
for the lookups get operation typically these are written to a http.Request
*/
type LookupsGetParams struct {

	/*Q
	  Query string specifying properties to search for

	*/
	Q *string

	timeout time.Duration
}

// WithQ adds the q to the lookups get params
func (o *LookupsGetParams) WithQ(q *string) *LookupsGetParams {
	o.Q = q
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *LookupsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}