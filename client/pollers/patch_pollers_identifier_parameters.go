package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchPollersIdentifierParams creates a new PatchPollersIdentifierParams object
// with the default values initialized.
func NewPatchPollersIdentifierParams() *PatchPollersIdentifierParams {
	var ()
	return &PatchPollersIdentifierParams{}
}

/*PatchPollersIdentifierParams contains all the parameters to send to the API endpoint
for the patch pollers identifier operation typically these are written to a http.Request
*/
type PatchPollersIdentifierParams struct {

	/*Identifier
	  poller identifier


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the patch pollers identifier params
func (o *PatchPollersIdentifierParams) WithIdentifier(Identifier string) *PatchPollersIdentifierParams {
	o.Identifier = Identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPollersIdentifierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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