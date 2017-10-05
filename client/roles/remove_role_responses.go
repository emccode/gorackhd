// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// RemoveRoleReader is a Reader for the RemoveRole structure.
type RemoveRoleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveRoleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRemoveRoleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRemoveRoleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRemoveRoleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewRemoveRoleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveRoleNoContent creates a RemoveRoleNoContent with default headers values
func NewRemoveRoleNoContent() *RemoveRoleNoContent {
	return &RemoveRoleNoContent{}
}

/*RemoveRoleNoContent handles this case with default header values.

Successfully deleted the specified role
*/
type RemoveRoleNoContent struct {
}

func (o *RemoveRoleNoContent) Error() string {
	return fmt.Sprintf("[DELETE /roles/{name}][%d] removeRoleNoContent ", 204)
}

func (o *RemoveRoleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveRoleUnauthorized creates a RemoveRoleUnauthorized with default headers values
func NewRemoveRoleUnauthorized() *RemoveRoleUnauthorized {
	return &RemoveRoleUnauthorized{}
}

/*RemoveRoleUnauthorized handles this case with default header values.

Unauthorized
*/
type RemoveRoleUnauthorized struct {
	Payload RemoveRoleUnauthorizedBody
}

func (o *RemoveRoleUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /roles/{name}][%d] removeRoleUnauthorized  %+v", 401, o.Payload)
}

func (o *RemoveRoleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveRoleForbidden creates a RemoveRoleForbidden with default headers values
func NewRemoveRoleForbidden() *RemoveRoleForbidden {
	return &RemoveRoleForbidden{}
}

/*RemoveRoleForbidden handles this case with default header values.

Forbidden
*/
type RemoveRoleForbidden struct {
	Payload RemoveRoleForbiddenBody
}

func (o *RemoveRoleForbidden) Error() string {
	return fmt.Sprintf("[DELETE /roles/{name}][%d] removeRoleForbidden  %+v", 403, o.Payload)
}

func (o *RemoveRoleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveRoleDefault creates a RemoveRoleDefault with default headers values
func NewRemoveRoleDefault(code int) *RemoveRoleDefault {
	return &RemoveRoleDefault{
		_statusCode: code,
	}
}

/*RemoveRoleDefault handles this case with default header values.

Unexpected error
*/
type RemoveRoleDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the remove role default response
func (o *RemoveRoleDefault) Code() int {
	return o._statusCode
}

func (o *RemoveRoleDefault) Error() string {
	return fmt.Sprintf("[DELETE /roles/{name}][%d] removeRole default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveRoleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveRoleForbiddenBody remove role forbidden body
swagger:model RemoveRoleForbiddenBody
*/

type RemoveRoleForbiddenBody interface{}

/*RemoveRoleUnauthorizedBody remove role unauthorized body
swagger:model RemoveRoleUnauthorizedBody
*/

type RemoveRoleUnauthorizedBody interface{}
