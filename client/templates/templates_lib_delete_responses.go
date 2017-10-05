// Code generated by go-swagger; DO NOT EDIT.

package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// TemplatesLibDeleteReader is a Reader for the TemplatesLibDelete structure.
type TemplatesLibDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplatesLibDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewTemplatesLibDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewTemplatesLibDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewTemplatesLibDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTemplatesLibDeleteNoContent creates a TemplatesLibDeleteNoContent with default headers values
func NewTemplatesLibDeleteNoContent() *TemplatesLibDeleteNoContent {
	return &TemplatesLibDeleteNoContent{}
}

/*TemplatesLibDeleteNoContent handles this case with default header values.

Successfully deleted the specified template
*/
type TemplatesLibDeleteNoContent struct {
}

func (o *TemplatesLibDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /templates/library/{name}][%d] templatesLibDeleteNoContent ", 204)
}

func (o *TemplatesLibDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTemplatesLibDeleteNotFound creates a TemplatesLibDeleteNotFound with default headers values
func NewTemplatesLibDeleteNotFound() *TemplatesLibDeleteNotFound {
	return &TemplatesLibDeleteNotFound{}
}

/*TemplatesLibDeleteNotFound handles this case with default header values.

The template with specified identifier was not found
*/
type TemplatesLibDeleteNotFound struct {
	Payload *models.Error
}

func (o *TemplatesLibDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /templates/library/{name}][%d] templatesLibDeleteNotFound  %+v", 404, o.Payload)
}

func (o *TemplatesLibDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplatesLibDeleteDefault creates a TemplatesLibDeleteDefault with default headers values
func NewTemplatesLibDeleteDefault(code int) *TemplatesLibDeleteDefault {
	return &TemplatesLibDeleteDefault{
		_statusCode: code,
	}
}

/*TemplatesLibDeleteDefault handles this case with default header values.

Unexpected error
*/
type TemplatesLibDeleteDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the templates lib delete default response
func (o *TemplatesLibDeleteDefault) Code() int {
	return o._statusCode
}

func (o *TemplatesLibDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /templates/library/{name}][%d] templatesLibDelete default  %+v", o._statusCode, o.Payload)
}

func (o *TemplatesLibDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
