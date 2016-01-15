package delete

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteDhcpLeaseMacParams creates a new DeleteDhcpLeaseMacParams object
// with the default values initialized.
func NewDeleteDhcpLeaseMacParams() *DeleteDhcpLeaseMacParams {
	return &DeleteDhcpLeaseMacParams{}
}

/*DeleteDhcpLeaseMacParams contains all the parameters to send to the API endpoint
for the delete dhcp lease mac operation typically these are written to a http.Request
*/
type DeleteDhcpLeaseMacParams struct {

	/*Mac
	  identifier of a mac address

	*/
	Mac string
}

// WithMac adds the mac to the delete dhcp lease mac params
func (o *DeleteDhcpLeaseMacParams) WithMac(mac string) *DeleteDhcpLeaseMacParams {
	o.Mac = mac
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDhcpLeaseMacParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param mac
	if err := r.SetPathParam("mac", o.Mac); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}