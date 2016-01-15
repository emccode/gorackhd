package obms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new obms API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for obms API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
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
