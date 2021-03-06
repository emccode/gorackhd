package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostLookupsIDParams creates a new PostLookupsIDParams object
// with the default values initialized.
func NewPostLookupsIDParams() *PostLookupsIDParams {
	var ()
	return &PostLookupsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLookupsIDParamsWithTimeout creates a new PostLookupsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLookupsIDParamsWithTimeout(timeout time.Duration) *PostLookupsIDParams {
	var ()
	return &PostLookupsIDParams{

		timeout: timeout,
	}
}

/*PostLookupsIDParams contains all the parameters to send to the API endpoint
for the post lookups ID operation typically these are written to a http.Request
*/
type PostLookupsIDParams struct {

	/*Content
	  foo

	*/
	Content interface{}
	/*ID
	  id of thing to lookup

	*/
	ID string

	timeout time.Duration
}

// WithContent adds the content to the post lookups ID params
func (o *PostLookupsIDParams) WithContent(content interface{}) *PostLookupsIDParams {
	o.Content = content
	return o
}

// WithID adds the id to the post lookups ID params
func (o *PostLookupsIDParams) WithID(id string) *PostLookupsIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostLookupsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Content); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
