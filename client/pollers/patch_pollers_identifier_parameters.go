package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewPatchPollersIdentifierParams creates a new PatchPollersIdentifierParams object
// with the default values initialized.
func NewPatchPollersIdentifierParams() *PatchPollersIdentifierParams {
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
func (o *PatchPollersIdentifierParams) WithIdentifier(identifier string) *PatchPollersIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PatchPollersIdentifierParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
