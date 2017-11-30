package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTagsParams creates a new GetTagsParams object
// with the default values initialized.
func NewGetTagsParams() *GetTagsParams {

	return &GetTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTagsParamsWithTimeout creates a new GetTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTagsParamsWithTimeout(timeout time.Duration) *GetTagsParams {

	return &GetTagsParams{

		timeout: timeout,
	}
}

/*GetTagsParams contains all the parameters to send to the API endpoint
for the get tags operation typically these are written to a http.Request
*/
type GetTagsParams struct {
	timeout time.Duration
}

// WriteToRequest writes these params to a swagger request
func (o *GetTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}