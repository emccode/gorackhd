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

// NewPostLookupsParams creates a new PostLookupsParams object
// with the default values initialized.
func NewPostLookupsParams() *PostLookupsParams {
	var ()
	return &PostLookupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLookupsParamsWithTimeout creates a new PostLookupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLookupsParamsWithTimeout(timeout time.Duration) *PostLookupsParams {
	var ()
	return &PostLookupsParams{

		timeout: timeout,
	}
}

/*PostLookupsParams contains all the parameters to send to the API endpoint
for the post lookups operation typically these are written to a http.Request
*/
type PostLookupsParams struct {

	/*Content
	  foo

	*/
	Content interface{}

	timeout time.Duration
}

// WithContent adds the content to the post lookups params
func (o *PostLookupsParams) WithContent(content interface{}) *PostLookupsParams {
	o.Content = content
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostLookupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Content); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
