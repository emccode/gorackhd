package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTemplatesLibraryIdentifierParams creates a new GetTemplatesLibraryIdentifierParams object
// with the default values initialized.
func NewGetTemplatesLibraryIdentifierParams() *GetTemplatesLibraryIdentifierParams {
	var ()
	return &GetTemplatesLibraryIdentifierParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTemplatesLibraryIdentifierParamsWithTimeout creates a new GetTemplatesLibraryIdentifierParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTemplatesLibraryIdentifierParamsWithTimeout(timeout time.Duration) *GetTemplatesLibraryIdentifierParams {
	var ()
	return &GetTemplatesLibraryIdentifierParams{

		timeout: timeout,
	}
}

/*GetTemplatesLibraryIdentifierParams contains all the parameters to send to the API endpoint
for the get templates library identifier operation typically these are written to a http.Request
*/
type GetTemplatesLibraryIdentifierParams struct {

	/*Identifier
	  template identifier


	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the get templates library identifier params
func (o *GetTemplatesLibraryIdentifierParams) WithIdentifier(identifier string) *GetTemplatesLibraryIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetTemplatesLibraryIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
