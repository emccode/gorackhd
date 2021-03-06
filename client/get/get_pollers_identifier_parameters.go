package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetPollersIdentifierParams creates a new GetPollersIdentifierParams object
// with the default values initialized.
func NewGetPollersIdentifierParams() *GetPollersIdentifierParams {
	var ()
	return &GetPollersIdentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPollersIdentifierParamsWithTimeout creates a new GetPollersIdentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPollersIdentifierParamsWithTimeout(timeout time.Duration) *GetPollersIdentifierParams {
	var ()
	return &GetPollersIdentifierParams{

		timeout: timeout,
	}
}

/*GetPollersIdentifierParams contains all the parameters to send to the API endpoint
for the get pollers identifier operation typically these are written to a http.Request
*/
type GetPollersIdentifierParams struct {

	/*Identifier
	  poller identifier


	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get pollers identifier params
func (o *GetPollersIdentifierParams) WithIdentifier(identifier string) *GetPollersIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPollersIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
