package delete

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new delete API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for delete API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteDhcpLeaseMac gets list of all catalogs

Delete the lease for the mac specified and return information about deleted lease.

*/
func (a *Client) DeleteDhcpLeaseMac(params *DeleteDhcpLeaseMacParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteDhcpLeaseMacOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDhcpLeaseMacParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteDhcpLeaseMac",
		Method:             "DELETE",
		PathPattern:        "/dhcp/lease/{mac}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDhcpLeaseMacReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteDhcpLeaseMacOK), nil
}

/*
DeleteFilesFileidentifier deletes file based on uuid

Put file based on filename, returns the uuid of the stored file.

*/
func (a *Client) DeleteFilesFileidentifier(params *DeleteFilesFileidentifierParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteFilesFileidentifierNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteFilesFileidentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteFilesFileidentifier",
		Method:             "DELETE",
		PathPattern:        "/files/{fileidentifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteFilesFileidentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteFilesFileidentifierNoContent), nil
}

/*
DeleteNodesIdentifier deletes specified node

Delete specified node.

*/
func (a *Client) DeleteNodesIdentifier(params *DeleteNodesIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodesIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodesIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNodesIdentifier",
		Method:             "DELETE",
		PathPattern:        "/nodes/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNodesIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodesIdentifierOK), nil
}

/*
DeleteNodesIdentifierTagsTagname removes tag from specified node

Remove tag from specified node.

*/
func (a *Client) DeleteNodesIdentifierTagsTagname(params *DeleteNodesIdentifierTagsTagnameParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodesIdentifierTagsTagnameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodesIdentifierTagsTagnameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNodesIdentifierTagsTagname",
		Method:             "DELETE",
		PathPattern:        "/nodes/{identifier}/tags/{tagname}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNodesIdentifierTagsTagnameReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodesIdentifierTagsTagnameOK), nil
}

/*
DeleteNodesIdentifierWorkflowsActive cancels currently running workflows for specified node

Cancel currently running workflows for specified node

*/
func (a *Client) DeleteNodesIdentifierWorkflowsActive(params *DeleteNodesIdentifierWorkflowsActiveParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodesIdentifierWorkflowsActiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodesIdentifierWorkflowsActiveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNodesIdentifierWorkflowsActive",
		Method:             "DELETE",
		PathPattern:        "/nodes/{identifier}/workflows/active",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNodesIdentifierWorkflowsActiveReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodesIdentifierWorkflowsActiveOK), nil
}

/*
DeleteNodesMacaddressDhcpWhitelist removes a whitelist of specified mac address

Remove a whitelist of specified mac address

*/
func (a *Client) DeleteNodesMacaddressDhcpWhitelist(params *DeleteNodesMacaddressDhcpWhitelistParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteNodesMacaddressDhcpWhitelistNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNodesMacaddressDhcpWhitelistParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteNodesMacaddressDhcpWhitelist",
		Method:             "DELETE",
		PathPattern:        "/nodes/{macaddress}/dhcp/whitelist",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteNodesMacaddressDhcpWhitelistReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteNodesMacaddressDhcpWhitelistNoContent), nil
}

/*
DeletePollersIdentifier deletes the specified poller

delete the specified poller

*/
func (a *Client) DeletePollersIdentifier(params *DeletePollersIdentifierParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePollersIdentifierNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePollersIdentifierParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeletePollersIdentifier",
		Method:             "DELETE",
		PathPattern:        "/pollers/{identifier}",
		ProducesMediaTypes: []string{"application/json", "application/x-gzip"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePollersIdentifierReader{formats: a.formats},
		AuthInfo:           authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePollersIdentifierNoContent), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
