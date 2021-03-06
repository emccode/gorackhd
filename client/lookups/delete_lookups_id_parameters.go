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

// NewDeleteLookupsIDParams creates a new DeleteLookupsIDParams object
// with the default values initialized.
func NewDeleteLookupsIDParams() *DeleteLookupsIDParams {
	var ()
	return &DeleteLookupsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLookupsIDParamsWithTimeout creates a new DeleteLookupsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLookupsIDParamsWithTimeout(timeout time.Duration) *DeleteLookupsIDParams {
	var ()
	return &DeleteLookupsIDParams{

		timeout: timeout,
	}
}

/*DeleteLookupsIDParams contains all the parameters to send to the API endpoint
for the delete lookups ID operation typically these are written to a http.Request
*/
type DeleteLookupsIDParams struct {

	/*ID
	  id to delete

	*/
	ID string

	timeout time.Duration
}

// WithID adds the id to the delete lookups ID params
func (o *DeleteLookupsIDParams) WithID(id string) *DeleteLookupsIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLookupsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
