// Code generated by go-swagger; DO NOT EDIT.

package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPollersDataGetParams creates a new PollersDataGetParams object
// with the default values initialized.
func NewPollersDataGetParams() *PollersDataGetParams {
	var ()
	return &PollersDataGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPollersDataGetParamsWithTimeout creates a new PollersDataGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPollersDataGetParamsWithTimeout(timeout time.Duration) *PollersDataGetParams {
	var ()
	return &PollersDataGetParams{

		timeout: timeout,
	}
}

// NewPollersDataGetParamsWithContext creates a new PollersDataGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPollersDataGetParamsWithContext(ctx context.Context) *PollersDataGetParams {
	var ()
	return &PollersDataGetParams{

		Context: ctx,
	}
}

// NewPollersDataGetParamsWithHTTPClient creates a new PollersDataGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPollersDataGetParamsWithHTTPClient(client *http.Client) *PollersDataGetParams {
	var ()
	return &PollersDataGetParams{
		HTTPClient: client,
	}
}

/*PollersDataGetParams contains all the parameters to send to the API endpoint
for the pollers data get operation typically these are written to a http.Request
*/
type PollersDataGetParams struct {

	/*Identifier
	  The identifier of the poller

	*/
	Identifier string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pollers data get params
func (o *PollersDataGetParams) WithTimeout(timeout time.Duration) *PollersDataGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pollers data get params
func (o *PollersDataGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pollers data get params
func (o *PollersDataGetParams) WithContext(ctx context.Context) *PollersDataGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pollers data get params
func (o *PollersDataGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pollers data get params
func (o *PollersDataGetParams) WithHTTPClient(client *http.Client) *PollersDataGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pollers data get params
func (o *PollersDataGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIdentifier adds the identifier to the pollers data get params
func (o *PollersDataGetParams) WithIdentifier(identifier string) *PollersDataGetParams {
	o.SetIdentifier(identifier)
	return o
}

// SetIdentifier adds the identifier to the pollers data get params
func (o *PollersDataGetParams) SetIdentifier(identifier string) {
	o.Identifier = identifier
}

// WriteToRequest writes these params to a swagger request
func (o *PollersDataGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
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
