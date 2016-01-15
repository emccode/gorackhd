package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new post API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for post API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*find all


put all

*/
func (a *Client) PostLookups(params *PostLookupsParams) (*PostLookupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLookupsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PostLookups",
		Params: params,
		Reader: &PostLookupsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLookupsOK), nil
}

/*post id


post id

*/
func (a *Client) PostLookupsID(params *PostLookupsIDParams) (*PostLookupsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostLookupsIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PostLookupsID",
		Params: params,
		Reader: &PostLookupsIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLookupsIDOK), nil
}

/*post


post

*/
func (a *Client) PostNodes(params *PostNodesParams) (*PostNodesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PostNodes",
		Params: params,
		Reader: &PostNodesReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesCreated), nil
}

/*set the obm settings associated with a node.


set he obm settings associated with a node.

*/
func (a *Client) PostNodesIdentifierObm(params *PostNodesIdentifierObmParams) (*PostNodesIdentifierObmCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesIdentifierObmParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PostNodesIdentifierObm",
		Params: params,
		Reader: &PostNodesIdentifierObmReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesIdentifierObmCreated), nil
}

/*Enable or disable identify light on node through OBM (if supported)


Enable or disable identify light on node through OBM (if supported)

*/
func (a *Client) PostNodesIdentifierObmIdentify(params *PostNodesIdentifierObmIdentifyParams) (*PostNodesIdentifierObmIdentifyCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesIdentifierObmIdentifyParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PostNodesIdentifierObmIdentify",
		Params: params,
		Reader: &PostNodesIdentifierObmIdentifyReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesIdentifierObmIdentifyCreated), nil
}

/*create workflow for specified node


create workflow for specified node

*/
func (a *Client) PostNodesIdentifierWorkflows(params *PostNodesIdentifierWorkflowsParams) (*PostNodesIdentifierWorkflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesIdentifierWorkflowsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PostNodesIdentifierWorkflows",
		Params: params,
		Reader: &PostNodesIdentifierWorkflowsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesIdentifierWorkflowsOK), nil
}

/*Add a whitelist of specified mac address


Add a whitelist of specified mac address

*/
func (a *Client) PostNodesMacaddressDhcpWhitelist(params *PostNodesMacaddressDhcpWhitelistParams) (*PostNodesMacaddressDhcpWhitelistCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNodesMacaddressDhcpWhitelistParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PostNodesMacaddressDhcpWhitelist",
		Params: params,
		Reader: &PostNodesMacaddressDhcpWhitelistReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNodesMacaddressDhcpWhitelistCreated), nil
}

/*create a sku


create a sku

*/
func (a *Client) PostSkus(params *PostSkusParams) (*PostSkusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostSkusParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PostSkus",
		Params: params,
		Reader: &PostSkusReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSkusOK), nil
}

/*put file based on filename


Put file based on filename, returns the uuid of the stored file.

*/
func (a *Client) PutFilesFileidentifier(params *PutFilesFileidentifierParams) (*PutFilesFileidentifierCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutFilesFileidentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PutFilesFileidentifier",
		Params: params,
		Reader: &PutFilesFileidentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutFilesFileidentifierCreated), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}