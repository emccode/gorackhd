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

type GetNodesIdentifierObmReader struct {
	formats strfmt.Registry
}

func (o *GetNodesIdentifierObmReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesIdentifierObmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetNodesIdentifierObmNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNodesIdentifierObmDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetNodesIdentifierObmOK creates a GetNodesIdentifierObmOK with default headers values
func NewGetNodesIdentifierObmOK() *GetNodesIdentifierObmOK {
	return &GetNodesIdentifierObmOK{}
}

/*GetNodesIdentifierObmOK

obm settings
*/
type GetNodesIdentifierObmOK struct {
	Payload GetNodesIdentifierObmOKBodyBody
}

func (o *GetNodesIdentifierObmOK) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/obm][%d] getNodesIdentifierObmOK  %+v", 200, o.Payload)
}

func (o *GetNodesIdentifierObmOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetNodesIdentifierObmNotFound creates a GetNodesIdentifierObmNotFound with default headers values
func NewGetNodesIdentifierObmNotFound() *GetNodesIdentifierObmNotFound {
	return &GetNodesIdentifierObmNotFound{}
}

/*GetNodesIdentifierObmNotFound

The node with the identifier was not found or has no obm settings.

*/
type GetNodesIdentifierObmNotFound struct {
	Payload *models.Error
}

func (o *GetNodesIdentifierObmNotFound) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/obm][%d] getNodesIdentifierObmNotFound  %+v", 404, o.Payload)
}

func (o *GetNodesIdentifierObmNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetNodesIdentifierObmDefault creates a GetNodesIdentifierObmDefault with default headers values
func NewGetNodesIdentifierObmDefault(code int) *GetNodesIdentifierObmDefault {
	return &GetNodesIdentifierObmDefault{
		_statusCode: code,
	}
}

/*GetNodesIdentifierObmDefault

Unexpected error
*/
type GetNodesIdentifierObmDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get nodes identifier obm default response
func (o *GetNodesIdentifierObmDefault) Code() int {
	return o._statusCode
}

func (o *GetNodesIdentifierObmDefault) Error() string {
	return fmt.Sprintf("[GET /nodes/{identifier}/obm][%d] GetNodesIdentifierObm default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodesIdentifierObmDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*GetNodesIdentifierObmOKBodyBody GetNodesIdentifierObmOKBodyBody get nodes identifier obm o k body body

swagger:model GetNodesIdentifierObmOKBodyBody
*/
type GetNodesIdentifierObmOKBodyBody interface{}
