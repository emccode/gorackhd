// Code generated by go-swagger; DO NOT EDIT.

package obms

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

// NewObmsDefinitionsGetByNameParams creates a new ObmsDefinitionsGetByNameParams object
// with the default values initialized.
func NewObmsDefinitionsGetByNameParams() *ObmsDefinitionsGetByNameParams {
	var ()
	return &ObmsDefinitionsGetByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewObmsDefinitionsGetByNameParamsWithTimeout creates a new ObmsDefinitionsGetByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewObmsDefinitionsGetByNameParamsWithTimeout(timeout time.Duration) *ObmsDefinitionsGetByNameParams {
	var ()
	return &ObmsDefinitionsGetByNameParams{

		timeout: timeout,
	}
}

// NewObmsDefinitionsGetByNameParamsWithContext creates a new ObmsDefinitionsGetByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewObmsDefinitionsGetByNameParamsWithContext(ctx context.Context) *ObmsDefinitionsGetByNameParams {
	var ()
	return &ObmsDefinitionsGetByNameParams{

		Context: ctx,
	}
}

// NewObmsDefinitionsGetByNameParamsWithHTTPClient creates a new ObmsDefinitionsGetByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewObmsDefinitionsGetByNameParamsWithHTTPClient(client *http.Client) *ObmsDefinitionsGetByNameParams {
	var ()
	return &ObmsDefinitionsGetByNameParams{
		HTTPClient: client,
	}
}

/*ObmsDefinitionsGetByNameParams contains all the parameters to send to the API endpoint
for the obms definitions get by name operation typically these are written to a http.Request
*/
type ObmsDefinitionsGetByNameParams struct {

	/*Name
	  The OBM service name

	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the obms definitions get by name params
func (o *ObmsDefinitionsGetByNameParams) WithTimeout(timeout time.Duration) *ObmsDefinitionsGetByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the obms definitions get by name params
func (o *ObmsDefinitionsGetByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the obms definitions get by name params
func (o *ObmsDefinitionsGetByNameParams) WithContext(ctx context.Context) *ObmsDefinitionsGetByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the obms definitions get by name params
func (o *ObmsDefinitionsGetByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the obms definitions get by name params
func (o *ObmsDefinitionsGetByNameParams) WithHTTPClient(client *http.Client) *ObmsDefinitionsGetByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the obms definitions get by name params
func (o *ObmsDefinitionsGetByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the obms definitions get by name params
func (o *ObmsDefinitionsGetByNameParams) WithName(name string) *ObmsDefinitionsGetByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the obms definitions get by name params
func (o *ObmsDefinitionsGetByNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ObmsDefinitionsGetByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
