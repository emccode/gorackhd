package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteSkusIdentifierParams creates a new DeleteSkusIdentifierParams object
// with the default values initialized.
func NewDeleteSkusIdentifierParams() *DeleteSkusIdentifierParams {
	return &DeleteSkusIdentifierParams{}
}

/*DeleteSkusIdentifierParams contains all the parameters to send to the API endpoint
for the delete skus identifier operation typically these are written to a http.Request
*/
type DeleteSkusIdentifierParams struct {

	/*Identifier
	  The sku objectid.


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the delete skus identifier params
func (o *DeleteSkusIdentifierParams) WithIdentifier(identifier string) *DeleteSkusIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSkusIdentifierParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
