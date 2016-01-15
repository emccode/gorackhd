package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new get API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for get API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*get list of all catalogs


The catalogs endpoint returns json data that represent the catalogs of
all hardware in the system.

*/
func (a *Client) GetCatalogs(params *GetCatalogsParams) (*GetCatalogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetCatalogs",
		Params: params,
		Reader: &GetCatalogsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCatalogsOK), nil
}

/*get list of all catalogs


The catalogs endpoint returns json data that represent the catalogs of
all hardware in the system.

*/
func (a *Client) GetCatalogsIdentifier(params *GetCatalogsIdentifierParams) (*GetCatalogsIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCatalogsIdentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetCatalogsIdentifier",
		Params: params,
		Reader: &GetCatalogsIdentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCatalogsIdentifierOK), nil
}

/*get server configuration


Get server configuration.

*/
func (a *Client) GetConfig(params *GetConfigParams) (*GetConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetConfig",
		Params: params,
		Reader: &GetConfigReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfigOK), nil
}

/*get DHCP lease table


Fetch the dhcp leases.

*/
func (a *Client) GetDhcp(params *GetDhcpParams) (*GetDhcpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDhcpParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetDhcp",
		Params: params,
		Reader: &GetDhcpReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDhcpOK), nil
}

/*fetch lease information for the mac specified


Fetch lease information for the mac specified.

*/
func (a *Client) GetDhcpLeaseMac(params *GetDhcpLeaseMacParams) (*GetDhcpLeaseMacOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDhcpLeaseMacParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetDhcpLeaseMac",
		Params: params,
		Reader: &GetDhcpLeaseMacReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDhcpLeaseMacOK), nil
}

/*get file based on uuid


Get file based on uuid.

*/
func (a *Client) GetFilesFileidentifier(params *GetFilesFileidentifierParams) (*GetFilesFileidentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetFilesFileidentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetFilesFileidentifier",
		Params: params,
		Reader: &GetFilesFileidentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetFilesFileidentifierOK), nil
}

/*find all


find all

*/
func (a *Client) GetLookups(params *GetLookupsParams) (*GetLookupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLookupsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetLookups",
		Params: params,
		Reader: &GetLookupsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLookupsOK), nil
}

/*lookup id


lookup id

*/
func (a *Client) GetLookupsID(params *GetLookupsIDParams) (*GetLookupsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLookupsIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetLookupsID",
		Params: params,
		Reader: &GetLookupsIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLookupsIDOK), nil
}

/*List of all nodes or if there are none an empty object


List of all nodes or if there are none an empty object

*/
func (a *Client) GetNodes(params *GetNodesParams) (*GetNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetNodes",
		Params: params,
		Reader: &GetNodesReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesOK), nil
}

/*List of all nodes or if there are none an empty object


List of all nodes or if there are none an empty object

*/
func (a *Client) GetNodesIdentifier(params *GetNodesIdentifierParams) (*GetNodesIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetNodesIdentifier",
		Params: params,
		Reader: &GetNodesIdentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierOK), nil
}

/*Fetch catalog of specified node


Fetch catalog of specified node

*/
func (a *Client) GetNodesIdentifierCatalogs(params *GetNodesIdentifierCatalogsParams) (*GetNodesIdentifierCatalogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierCatalogsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetNodesIdentifierCatalogs",
		Params: params,
		Reader: &GetNodesIdentifierCatalogsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierCatalogsOK), nil
}

/*Fetch catalog of specified node for given source


Fetch catalog of specified node for given source

*/
func (a *Client) GetNodesIdentifierCatalogsSource(params *GetNodesIdentifierCatalogsSourceParams) (*GetNodesIdentifierCatalogsSourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierCatalogsSourceParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetNodesIdentifierCatalogsSource",
		Params: params,
		Reader: &GetNodesIdentifierCatalogsSourceReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierCatalogsSourceOK), nil
}

/*get the obm settings associated with a node.


get the obm settings associated with a node.

*/
func (a *Client) GetNodesIdentifierObm(params *GetNodesIdentifierObmParams) (*GetNodesIdentifierObmOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierObmParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetNodesIdentifierObm",
		Params: params,
		Reader: &GetNodesIdentifierObmReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierObmOK), nil
}

/*Fetch status of identify light on node through OBM (if supported)


Fetch status of identify light on node through OBM (if supported)

*/
func (a *Client) GetNodesIdentifierObmIdentify(params *GetNodesIdentifierObmIdentifyParams) (*GetNodesIdentifierObmIdentifyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierObmIdentifyParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetNodesIdentifierObmIdentify",
		Params: params,
		Reader: &GetNodesIdentifierObmIdentifyReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierObmIdentifyOK), nil
}

/*Fetch all pollers for specified node


Fetch all pollers for specified node

*/
func (a *Client) GetNodesIdentifierPollers(params *GetNodesIdentifierPollersParams) (*GetNodesIdentifierPollersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierPollersParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetNodesIdentifierPollers",
		Params: params,
		Reader: &GetNodesIdentifierPollersReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierPollersOK), nil
}

/*Fetch all workflows for specified node


Fetch all workflows for specified node

*/
func (a *Client) GetNodesIdentifierWorkflows(params *GetNodesIdentifierWorkflowsParams) (*GetNodesIdentifierWorkflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierWorkflowsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetNodesIdentifierWorkflows",
		Params: params,
		Reader: &GetNodesIdentifierWorkflowsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierWorkflowsOK), nil
}

/*Fetch currently running workflows for specified node


Fetch currently running workflows for specified node

*/
func (a *Client) GetNodesIdentifierWorkflowsActive(params *GetNodesIdentifierWorkflowsActiveParams) (*GetNodesIdentifierWorkflowsActiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNodesIdentifierWorkflowsActiveParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetNodesIdentifierWorkflowsActive",
		Params: params,
		Reader: &GetNodesIdentifierWorkflowsActiveReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNodesIdentifierWorkflowsActiveOK), nil
}

/*get list of possible OBM services


get list of possible OBM services

*/
func (a *Client) GetObmsLibrary(params *GetObmsLibraryParams) (*GetObmsLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetObmsLibraryParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetObmsLibrary",
		Params: params,
		Reader: &GetObmsLibraryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetObmsLibraryOK), nil
}

/*get a single OBM service


get a single OBM service

*/
func (a *Client) GetObmsLibraryIdentifier(params *GetObmsLibraryIdentifierParams) (*GetObmsLibraryIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetObmsLibraryIdentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetObmsLibraryIdentifier",
		Params: params,
		Reader: &GetObmsLibraryIdentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetObmsLibraryIdentifierOK), nil
}

/*get list of all pollers


get list of all pollers

*/
func (a *Client) GetPollers(params *GetPollersParams) (*GetPollersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPollersParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetPollers",
		Params: params,
		Reader: &GetPollersReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPollersOK), nil
}

/*Get specifics of the specified poller


Get specifics of the specified poller

*/
func (a *Client) GetPollersIdentifier(params *GetPollersIdentifierParams) (*GetPollersIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPollersIdentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetPollersIdentifier",
		Params: params,
		Reader: &GetPollersIdentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPollersIdentifierOK), nil
}

/*Get data for the specific poller


Get data for the specific poller

*/
func (a *Client) GetPollersIdentifierData(params *GetPollersIdentifierDataParams) (*GetPollersIdentifierDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPollersIdentifierDataParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetPollersIdentifierData",
		Params: params,
		Reader: &GetPollersIdentifierDataReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPollersIdentifierDataOK), nil
}

/*get list of possible library pollers


get list of possible library pollers

*/
func (a *Client) GetPollersLibrary(params *GetPollersLibraryParams) (*GetPollersLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPollersLibraryParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetPollersLibrary",
		Params: params,
		Reader: &GetPollersLibraryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPollersLibraryOK), nil
}

/*get a single library poller


get a single library poller

*/
func (a *Client) GetPollersLibraryIdentifier(params *GetPollersLibraryIdentifierParams) (*GetPollersLibraryIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPollersLibraryIdentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetPollersLibraryIdentifier",
		Params: params,
		Reader: &GetPollersLibraryIdentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPollersLibraryIdentifierOK), nil
}

/*get list of possible profiles


get list of possible profiles

*/
func (a *Client) GetProfilesLibrary(params *GetProfilesLibraryParams) (*GetProfilesLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProfilesLibraryParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetProfilesLibrary",
		Params: params,
		Reader: &GetProfilesLibraryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProfilesLibraryOK), nil
}

/*get a single profile


get a single profile

*/
func (a *Client) GetProfilesLibraryIdentifier(params *GetProfilesLibraryIdentifierParams) (*GetProfilesLibraryIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetProfilesLibraryIdentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetProfilesLibraryIdentifier",
		Params: params,
		Reader: &GetProfilesLibraryIdentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetProfilesLibraryIdentifierOK), nil
}

/*get list of skus


get list of skus

*/
func (a *Client) GetSkus(params *GetSkusParams) (*GetSkusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSkusParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetSkus",
		Params: params,
		Reader: &GetSkusReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSkusOK), nil
}

/*get a single sku


get a single sku

*/
func (a *Client) GetSkusIdentifier(params *GetSkusIdentifierParams) (*GetSkusIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSkusIdentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetSkusIdentifier",
		Params: params,
		Reader: &GetSkusIdentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSkusIdentifierOK), nil
}

/*get nodes for specific sku


get nodes for specific sku

*/
func (a *Client) GetSkusIdentifierNodes(params *GetSkusIdentifierNodesParams) (*GetSkusIdentifierNodesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSkusIdentifierNodesParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetSkusIdentifierNodes",
		Params: params,
		Reader: &GetSkusIdentifierNodesReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSkusIdentifierNodesOK), nil
}

/*get list of possible templates


get list of possible templates

*/
func (a *Client) GetTemplatesLibrary(params *GetTemplatesLibraryParams) (*GetTemplatesLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTemplatesLibraryParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetTemplatesLibrary",
		Params: params,
		Reader: &GetTemplatesLibraryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTemplatesLibraryOK), nil
}

/*get a single template


get a single template

*/
func (a *Client) GetTemplatesLibraryIdentifier(params *GetTemplatesLibraryIdentifierParams) (*GetTemplatesLibraryIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTemplatesLibraryIdentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetTemplatesLibraryIdentifier",
		Params: params,
		Reader: &GetTemplatesLibraryIdentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTemplatesLibraryIdentifierOK), nil
}

/*get list of all versions of all packages plus our code


get list of all versions of all packages plus our code

*/
func (a *Client) GetVersions(params *GetVersionsParams) (*GetVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetVersions",
		Params: params,
		Reader: &GetVersionsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVersionsOK), nil
}

/*Fetch workflows


Fetch workflows

*/
func (a *Client) GetWorkflows(params *GetWorkflowsParams) (*GetWorkflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetWorkflows",
		Params: params,
		Reader: &GetWorkflowsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowsOK), nil
}

/*List all workflows available to run


List all workflows available to run

*/
func (a *Client) GetWorkflowsLibrary(params *GetWorkflowsLibraryParams) (*GetWorkflowsLibraryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsLibraryParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetWorkflowsLibrary",
		Params: params,
		Reader: &GetWorkflowsLibraryReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowsLibraryOK), nil
}

/*Fetch tasks from task library


Fetch tasks from task library

*/
func (a *Client) GetWorkflowsTasks(params *GetWorkflowsTasksParams) (*GetWorkflowsTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsTasksParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "GetWorkflowsTasks",
		Params: params,
		Reader: &GetWorkflowsTasksReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkflowsTasksOK), nil
}

/*create a poller


create a poller

*/
func (a *Client) PostPollers(params *PostPollersParams) (*PostPollersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPollersParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PostPollers",
		Params: params,
		Reader: &PostPollersReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPollersOK), nil
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
