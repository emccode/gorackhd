package whitelist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteNodesMacaddressDhcpWhitelistParams creates a new DeleteNodesMacaddressDhcpWhitelistParams object
// with the default values initialized.
func NewDeleteNodesMacaddressDhcpWhitelistParams() *DeleteNodesMacaddressDhcpWhitelistParams {
	return &DeleteNodesMacaddressDhcpWhitelistParams{}
}

/*DeleteNodesMacaddressDhcpWhitelistParams contains all the parameters to send to the API endpoint
for the delete nodes macaddress dhcp whitelist operation typically these are written to a http.Request
*/
type DeleteNodesMacaddressDhcpWhitelistParams struct {

	/*Macaddress
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Macaddress string
}

// WithMacaddress adds the macaddress to the delete nodes macaddress dhcp whitelist params
func (o *DeleteNodesMacaddressDhcpWhitelistParams) WithMacaddress(macaddress string) *DeleteNodesMacaddressDhcpWhitelistParams {
	o.Macaddress = macaddress
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNodesMacaddressDhcpWhitelistParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param macaddress
	if err := r.SetPathParam("macaddress", o.Macaddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
