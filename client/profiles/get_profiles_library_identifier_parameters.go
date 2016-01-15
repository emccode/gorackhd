package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetProfilesLibraryIdentifierParams creates a new GetProfilesLibraryIdentifierParams object
// with the default values initialized.
func NewGetProfilesLibraryIdentifierParams() *GetProfilesLibraryIdentifierParams {
	return &GetProfilesLibraryIdentifierParams{}
}

/*GetProfilesLibraryIdentifierParams contains all the parameters to send to the API endpoint
for the get profiles library identifier operation typically these are written to a http.Request
*/
type GetProfilesLibraryIdentifierParams struct {

	/*Identifier
	  The profile name.


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get profiles library identifier params
func (o *GetProfilesLibraryIdentifierParams) WithIdentifier(identifier string) *GetProfilesLibraryIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetProfilesLibraryIdentifierParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
