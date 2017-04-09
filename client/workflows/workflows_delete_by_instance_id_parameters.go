package workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewWorkflowsDeleteByInstanceIDParams creates a new WorkflowsDeleteByInstanceIDParams object
// with the default values initialized.
func NewWorkflowsDeleteByInstanceIDParams() *WorkflowsDeleteByInstanceIDParams {
	var ()
	return &WorkflowsDeleteByInstanceIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowsDeleteByInstanceIDParamsWithTimeout creates a new WorkflowsDeleteByInstanceIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWorkflowsDeleteByInstanceIDParamsWithTimeout(timeout time.Duration) *WorkflowsDeleteByInstanceIDParams {
	var ()
	return &WorkflowsDeleteByInstanceIDParams{

		timeout: timeout,
	}
}

/*WorkflowsDeleteByInstanceIDParams contains all the parameters to send to the API endpoint
for the workflows delete by instance Id operation typically these are written to a http.Request
*/
type WorkflowsDeleteByInstanceIDParams struct {

	/*Identifier
	  The workflow instance identifier

	*/
	Identifier string

	timeout time.Duration
}

// WithIdentifier adds the identifier to the workflows delete by instance Id params
func (o *WorkflowsDeleteByInstanceIDParams) WithIdentifier(identifier string) *WorkflowsDeleteByInstanceIDParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowsDeleteByInstanceIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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