// Code generated by go-swagger; DO NOT EDIT.

package ibms

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// IbmsGetByIDReader is a Reader for the IbmsGetByID structure.
type IbmsGetByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IbmsGetByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIbmsGetByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewIbmsGetByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIbmsGetByIDOK creates a IbmsGetByIDOK with default headers values
func NewIbmsGetByIDOK() *IbmsGetByIDOK {
	return &IbmsGetByIDOK{}
}

/*IbmsGetByIDOK handles this case with default header values.

Successfully retrieved the specified IBMS service
*/
type IbmsGetByIDOK struct {
	Payload IbmsGetByIDOKBody
}

func (o *IbmsGetByIDOK) Error() string {
	return fmt.Sprintf("[GET /ibms/{identifier}][%d] ibmsGetByIdOK  %+v", 200, o.Payload)
}

func (o *IbmsGetByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIbmsGetByIDDefault creates a IbmsGetByIDDefault with default headers values
func NewIbmsGetByIDDefault(code int) *IbmsGetByIDDefault {
	return &IbmsGetByIDDefault{
		_statusCode: code,
	}
}

/*IbmsGetByIDDefault handles this case with default header values.

Unexpected error
*/
type IbmsGetByIDDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the ibms get by Id default response
func (o *IbmsGetByIDDefault) Code() int {
	return o._statusCode
}

func (o *IbmsGetByIDDefault) Error() string {
	return fmt.Sprintf("[GET /ibms/{identifier}][%d] ibmsGetById default  %+v", o._statusCode, o.Payload)
}

func (o *IbmsGetByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*IbmsGetByIDOKBody ibms get by ID o k body
swagger:model IbmsGetByIDOKBody
*/

type IbmsGetByIDOKBody interface{}
