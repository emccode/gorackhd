// Code generated by go-swagger; DO NOT EDIT.

package profiles

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

// NewProfilesPutLibByNameParams creates a new ProfilesPutLibByNameParams object
// with the default values initialized.
func NewProfilesPutLibByNameParams() *ProfilesPutLibByNameParams {
	var ()
	return &ProfilesPutLibByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewProfilesPutLibByNameParamsWithTimeout creates a new ProfilesPutLibByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewProfilesPutLibByNameParamsWithTimeout(timeout time.Duration) *ProfilesPutLibByNameParams {
	var ()
	return &ProfilesPutLibByNameParams{

		timeout: timeout,
	}
}

// NewProfilesPutLibByNameParamsWithContext creates a new ProfilesPutLibByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewProfilesPutLibByNameParamsWithContext(ctx context.Context) *ProfilesPutLibByNameParams {
	var ()
	return &ProfilesPutLibByNameParams{

		Context: ctx,
	}
}

// NewProfilesPutLibByNameParamsWithHTTPClient creates a new ProfilesPutLibByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewProfilesPutLibByNameParamsWithHTTPClient(client *http.Client) *ProfilesPutLibByNameParams {
	var ()
	return &ProfilesPutLibByNameParams{
		HTTPClient: client,
	}
}

/*ProfilesPutLibByNameParams contains all the parameters to send to the API endpoint
for the profiles put lib by name operation typically these are written to a http.Request
*/
type ProfilesPutLibByNameParams struct {

	/*Name
	  The profile name

	*/
	Name string
	/*Scope
	  The profile scope

	*/
	Scope *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the profiles put lib by name params
func (o *ProfilesPutLibByNameParams) WithTimeout(timeout time.Duration) *ProfilesPutLibByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the profiles put lib by name params
func (o *ProfilesPutLibByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the profiles put lib by name params
func (o *ProfilesPutLibByNameParams) WithContext(ctx context.Context) *ProfilesPutLibByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the profiles put lib by name params
func (o *ProfilesPutLibByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the profiles put lib by name params
func (o *ProfilesPutLibByNameParams) WithHTTPClient(client *http.Client) *ProfilesPutLibByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the profiles put lib by name params
func (o *ProfilesPutLibByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the profiles put lib by name params
func (o *ProfilesPutLibByNameParams) WithName(name string) *ProfilesPutLibByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the profiles put lib by name params
func (o *ProfilesPutLibByNameParams) SetName(name string) {
	o.Name = name
}

// WithScope adds the scope to the profiles put lib by name params
func (o *ProfilesPutLibByNameParams) WithScope(scope *string) *ProfilesPutLibByNameParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the profiles put lib by name params
func (o *ProfilesPutLibByNameParams) SetScope(scope *string) {
	o.Scope = scope
}

// WriteToRequest writes these params to a swagger request
func (o *ProfilesPutLibByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Scope != nil {

		// query param scope
		var qrScope string
		if o.Scope != nil {
			qrScope = *o.Scope
		}
		qScope := qrScope
		if qScope != "" {
			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
