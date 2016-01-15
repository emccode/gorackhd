package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd/models"
)

type GetDhcpReader struct {
	formats strfmt.Registry
}

func (o *GetDhcpReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDhcpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDhcpDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetDhcpOK creates a GetDhcpOK with default headers values
func NewGetDhcpOK() *GetDhcpOK {
	return &GetDhcpOK{}
}

/*GetDhcpOK

A list of all DHCP leases registered with this instance of the service.

*/
type GetDhcpOK struct {
	Payload []*models.Lease
}

func (o *GetDhcpOK) Error() string {
	return fmt.Sprintf("[GET /dhcp][%d] getDhcpOK  %+v", 200, o.Payload)
}

func (o *GetDhcpOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetDhcpDefault creates a GetDhcpDefault with default headers values
func NewGetDhcpDefault(code int) *GetDhcpDefault {
	return &GetDhcpDefault{
		_statusCode: code,
	}
}

/*GetDhcpDefault

NotFound error
*/
type GetDhcpDefault struct {
	_statusCode int
}

// Code gets the status code for the get dhcp default response
func (o *GetDhcpDefault) Code() int {
	return o._statusCode
}

func (o *GetDhcpDefault) Error() string {
	return fmt.Sprintf("[GET /dhcp][%d] GetDhcp default ", o._statusCode)
}

func (o *GetDhcpDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}