// Code generated by go-swagger; DO NOT EDIT.

package tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// NewPostWorkflowByIDParams creates a new PostWorkflowByIDParams object
// with the default values initialized.
func NewPostWorkflowByIDParams() *PostWorkflowByIDParams {
	var ()
	return &PostWorkflowByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkflowByIDParamsWithTimeout creates a new PostWorkflowByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWorkflowByIDParamsWithTimeout(timeout time.Duration) *PostWorkflowByIDParams {
	var ()
	return &PostWorkflowByIDParams{

		timeout: timeout,
	}
}

// NewPostWorkflowByIDParamsWithContext creates a new PostWorkflowByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWorkflowByIDParamsWithContext(ctx context.Context) *PostWorkflowByIDParams {
	var ()
	return &PostWorkflowByIDParams{

		Context: ctx,
	}
}

// NewPostWorkflowByIDParamsWithHTTPClient creates a new PostWorkflowByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWorkflowByIDParamsWithHTTPClient(client *http.Client) *PostWorkflowByIDParams {
	var ()
	return &PostWorkflowByIDParams{
		HTTPClient: client,
	}
}

/*PostWorkflowByIDParams contains all the parameters to send to the API endpoint
for the post workflow by Id operation typically these are written to a http.Request
*/
type PostWorkflowByIDParams struct {

	/*Body
	  The workflow options to post

	*/
	Body *models.PostNodeWorkflow
	/*Name
	  Query string specifying the optional workflow injectable name

	*/
	Name *string
	/*TagName
	  The tag identifier

	*/
	TagName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post workflow by Id params
func (o *PostWorkflowByIDParams) WithTimeout(timeout time.Duration) *PostWorkflowByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workflow by Id params
func (o *PostWorkflowByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workflow by Id params
func (o *PostWorkflowByIDParams) WithContext(ctx context.Context) *PostWorkflowByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workflow by Id params
func (o *PostWorkflowByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workflow by Id params
func (o *PostWorkflowByIDParams) WithHTTPClient(client *http.Client) *PostWorkflowByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workflow by Id params
func (o *PostWorkflowByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workflow by Id params
func (o *PostWorkflowByIDParams) WithBody(body *models.PostNodeWorkflow) *PostWorkflowByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workflow by Id params
func (o *PostWorkflowByIDParams) SetBody(body *models.PostNodeWorkflow) {
	o.Body = body
}

// WithName adds the name to the post workflow by Id params
func (o *PostWorkflowByIDParams) WithName(name *string) *PostWorkflowByIDParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post workflow by Id params
func (o *PostWorkflowByIDParams) SetName(name *string) {
	o.Name = name
}

// WithTagName adds the tagName to the post workflow by Id params
func (o *PostWorkflowByIDParams) WithTagName(tagName string) *PostWorkflowByIDParams {
	o.SetTagName(tagName)
	return o
}

// SetTagName adds the tagName to the post workflow by Id params
func (o *PostWorkflowByIDParams) SetTagName(tagName string) {
	o.TagName = tagName
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkflowByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.PostNodeWorkflow)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	// path param tagName
	if err := r.SetPathParam("tagName", o.TagName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
