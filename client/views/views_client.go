package views

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new views API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for views API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ViewsDelete deletes the specified view

Delete a view with the specified name.
*/
func (a *Client) ViewsDelete(params *ViewsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*ViewsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "viewsDelete",
		Method:             "DELETE",
		PathPattern:        "/views/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ViewsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ViewsDeleteNoContent), nil
}

/*
ViewsGet gets all views

Retrieve a list of all views. Views are used to render the output of various system resources, such as nodes, pollers, and OBM settings.

*/
func (a *Client) ViewsGet(params *ViewsGetParams, authInfo runtime.ClientAuthInfoWriter) (*ViewsGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewsGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "viewsGet",
		Method:             "GET",
		PathPattern:        "/views",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ViewsGetReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ViewsGetOK), nil
}

/*
ViewsGetByID gets the specified view

Get the view with the specified file name.
*/
func (a *Client) ViewsGetByID(params *ViewsGetByIDParams, authInfo runtime.ClientAuthInfoWriter) (*ViewsGetByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewsGetByIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "viewsGetById",
		Method:             "GET",
		PathPattern:        "/views/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ViewsGetByIDReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ViewsGetByIDOK), nil
}

/*
ViewsPut puts the specified view

Create or update a view with the specified name.
*/
func (a *Client) ViewsPut(params *ViewsPutParams, authInfo runtime.ClientAuthInfoWriter) (*ViewsPutCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewViewsPutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "viewsPut",
		Method:             "PUT",
		PathPattern:        "/views/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/octet-stream", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ViewsPutReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ViewsPutCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}