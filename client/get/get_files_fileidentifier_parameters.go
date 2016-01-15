package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetFilesFileidentifierParams creates a new GetFilesFileidentifierParams object
// with the default values initialized.
func NewGetFilesFileidentifierParams() *GetFilesFileidentifierParams {
	return &GetFilesFileidentifierParams{}
}

/*GetFilesFileidentifierParams contains all the parameters to send to the API endpoint
for the get files fileidentifier operation typically these are written to a http.Request
*/
type GetFilesFileidentifierParams struct {

	/*Fileidentifier
	  uuid of a file as provided when you originally stored it.

	*/
	Fileidentifier string
}

// WithFileidentifier adds the fileidentifier to the get files fileidentifier params
func (o *GetFilesFileidentifierParams) WithFileidentifier(fileidentifier string) *GetFilesFileidentifierParams {
	o.Fileidentifier = fileidentifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetFilesFileidentifierParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param fileidentifier
	if err := r.SetPathParam("fileidentifier", o.Fileidentifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
