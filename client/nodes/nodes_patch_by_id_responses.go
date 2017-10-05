// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// NodesPatchByIDReader is a Reader for the NodesPatchByID structure.
type NodesPatchByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodesPatchByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewNodesPatchByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewNodesPatchByIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewNodesPatchByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNodesPatchByIDOK creates a NodesPatchByIDOK with default headers values
func NewNodesPatchByIDOK() *NodesPatchByIDOK {
	return &NodesPatchByIDOK{}
}

/*NodesPatchByIDOK handles this case with default header values.

Successfully modified the specified node
*/
type NodesPatchByIDOK struct {
	Payload *models.PartialNode
}

func (o *NodesPatchByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /nodes/{identifier}][%d] nodesPatchByIdOK  %+v", 200, o.Payload)
}

func (o *NodesPatchByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PartialNode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesPatchByIDNotFound creates a NodesPatchByIDNotFound with default headers values
func NewNodesPatchByIDNotFound() *NodesPatchByIDNotFound {
	return &NodesPatchByIDNotFound{}
}

/*NodesPatchByIDNotFound handles this case with default header values.

The specified node was not found
*/
type NodesPatchByIDNotFound struct {
	Payload *models.Error
}

func (o *NodesPatchByIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /nodes/{identifier}][%d] nodesPatchByIdNotFound  %+v", 404, o.Payload)
}

func (o *NodesPatchByIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodesPatchByIDDefault creates a NodesPatchByIDDefault with default headers values
func NewNodesPatchByIDDefault(code int) *NodesPatchByIDDefault {
	return &NodesPatchByIDDefault{
		_statusCode: code,
	}
}

/*NodesPatchByIDDefault handles this case with default header values.

Unexpected error
*/
type NodesPatchByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the nodes patch by Id default response
func (o *NodesPatchByIDDefault) Code() int {
	return o._statusCode
}

func (o *NodesPatchByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /nodes/{identifier}][%d] nodesPatchById default  %+v", o._statusCode, o.Payload)
}

func (o *NodesPatchByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
