package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostNodesIdentifierWorkflowsParams creates a new PostNodesIdentifierWorkflowsParams object
// with the default values initialized.
func NewPostNodesIdentifierWorkflowsParams() *PostNodesIdentifierWorkflowsParams {
	var ()
	return &PostNodesIdentifierWorkflowsParams{}
}

/*PostNodesIdentifierWorkflowsParams contains all the parameters to send to the API endpoint
for the post nodes identifier workflows operation typically these are written to a http.Request
*/
type PostNodesIdentifierWorkflowsParams struct {

	/*Body*/
	Body interface{}
	/*Identifier
	  The node unique identifier


	*/
	Identifier string
	/*Name
	  The injectable Graph name


	*/
	Name string
}

// WithBody adds the body to the post nodes identifier workflows params
func (o *PostNodesIdentifierWorkflowsParams) WithBody(Body interface{}) *PostNodesIdentifierWorkflowsParams {
	o.Body = Body
	return o
}

// WithIdentifier adds the identifier to the post nodes identifier workflows params
func (o *PostNodesIdentifierWorkflowsParams) WithIdentifier(Identifier string) *PostNodesIdentifierWorkflowsParams {
	o.Identifier = Identifier
	return o
}

// WithName adds the name to the post nodes identifier workflows params
func (o *PostNodesIdentifierWorkflowsParams) WithName(Name string) *PostNodesIdentifierWorkflowsParams {
	o.Name = Name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostNodesIdentifierWorkflowsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	// query param name
	qrName := o.Name
	qName := qrName
	if qName != "" {
		if err := r.SetQueryParam("name", qName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}