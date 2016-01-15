package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new profiles API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for profiles API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*Delete specific sku.


Delete specific sku.

*/
func (a *Client) DeleteSkusIdentifier(params *DeleteSkusIdentifierParams) (*DeleteSkusIdentifierNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSkusIdentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "DeleteSkusIdentifier",
		Params: params,
		Reader: &DeleteSkusIdentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSkusIdentifierNoContent), nil
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

/*put a single profile


put a single profile

*/
func (a *Client) PutProfilesLibraryIdentifier(params *PutProfilesLibraryIdentifierParams) (*PutProfilesLibraryIdentifierOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutProfilesLibraryIdentifierParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:     "PutProfilesLibraryIdentifier",
		Params: params,
		Reader: &PutProfilesLibraryIdentifierReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutProfilesLibraryIdentifierOK), nil
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