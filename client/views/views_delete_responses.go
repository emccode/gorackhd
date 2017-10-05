// Code generated by go-swagger; DO NOT EDIT.

package views

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// ViewsDeleteReader is a Reader for the ViewsDelete structure.
type ViewsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ViewsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewViewsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewViewsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewViewsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewViewsDeleteNoContent creates a ViewsDeleteNoContent with default headers values
func NewViewsDeleteNoContent() *ViewsDeleteNoContent {
	return &ViewsDeleteNoContent{}
}

/*ViewsDeleteNoContent handles this case with default header values.

Successfully deleted the specified view
*/
type ViewsDeleteNoContent struct {
	Payload ViewsDeleteNoContentBody
}

func (o *ViewsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /views/{identifier}][%d] viewsDeleteNoContent  %+v", 204, o.Payload)
}

func (o *ViewsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewsDeleteNotFound creates a ViewsDeleteNotFound with default headers values
func NewViewsDeleteNotFound() *ViewsDeleteNotFound {
	return &ViewsDeleteNotFound{}
}

/*ViewsDeleteNotFound handles this case with default header values.

The view with specified name was not found
*/
type ViewsDeleteNotFound struct {
	Payload *models.Error
}

func (o *ViewsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /views/{identifier}][%d] viewsDeleteNotFound  %+v", 404, o.Payload)
}

func (o *ViewsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewViewsDeleteDefault creates a ViewsDeleteDefault with default headers values
func NewViewsDeleteDefault(code int) *ViewsDeleteDefault {
	return &ViewsDeleteDefault{
		_statusCode: code,
	}
}

/*ViewsDeleteDefault handles this case with default header values.

Unexpected error
*/
type ViewsDeleteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the views delete default response
func (o *ViewsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ViewsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /views/{identifier}][%d] viewsDelete default  %+v", o._statusCode, o.Payload)
}

func (o *ViewsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ViewsDeleteNoContentBody views delete no content body
swagger:model ViewsDeleteNoContentBody
*/

type ViewsDeleteNoContentBody interface{}
