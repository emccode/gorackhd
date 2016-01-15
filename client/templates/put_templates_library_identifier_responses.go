package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd/models"
)

type PutTemplatesLibraryIdentifierReader struct {
	formats strfmt.Registry
}

func (o *PutTemplatesLibraryIdentifierReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutTemplatesLibraryIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPutTemplatesLibraryIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutTemplatesLibraryIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPutTemplatesLibraryIdentifierOK creates a PutTemplatesLibraryIdentifierOK with default headers values
func NewPutTemplatesLibraryIdentifierOK() *PutTemplatesLibraryIdentifierOK {
	return &PutTemplatesLibraryIdentifierOK{}
}

/*PutTemplatesLibraryIdentifierOK

return template

*/
type PutTemplatesLibraryIdentifierOK struct {
	Payload PutTemplatesLibraryIdentifierOKBodyBody
}

func (o *PutTemplatesLibraryIdentifierOK) Error() string {
	return fmt.Sprintf("[PUT /templates/library/{identifier}][%d] putTemplatesLibraryIdentifierOK  %+v", 200, o.Payload)
}

func (o *PutTemplatesLibraryIdentifierOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPutTemplatesLibraryIdentifierNotFound creates a PutTemplatesLibraryIdentifierNotFound with default headers values
func NewPutTemplatesLibraryIdentifierNotFound() *PutTemplatesLibraryIdentifierNotFound {
	return &PutTemplatesLibraryIdentifierNotFound{}
}

/*PutTemplatesLibraryIdentifierNotFound

There is no template with specified identifier.

*/
type PutTemplatesLibraryIdentifierNotFound struct {
	Payload *models.Error
}

func (o *PutTemplatesLibraryIdentifierNotFound) Error() string {
	return fmt.Sprintf("[PUT /templates/library/{identifier}][%d] putTemplatesLibraryIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *PutTemplatesLibraryIdentifierNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPutTemplatesLibraryIdentifierDefault creates a PutTemplatesLibraryIdentifierDefault with default headers values
func NewPutTemplatesLibraryIdentifierDefault(code int) *PutTemplatesLibraryIdentifierDefault {
	return &PutTemplatesLibraryIdentifierDefault{
		_statusCode: code,
	}
}

/*PutTemplatesLibraryIdentifierDefault

Unexpected error
*/
type PutTemplatesLibraryIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put templates library identifier default response
func (o *PutTemplatesLibraryIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *PutTemplatesLibraryIdentifierDefault) Error() string {
	return fmt.Sprintf("[PUT /templates/library/{identifier}][%d] PutTemplatesLibraryIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *PutTemplatesLibraryIdentifierDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*PutTemplatesLibraryIdentifierOKBodyBody PutTemplatesLibraryIdentifierOKBodyBody put templates library identifier o k body body

swagger:model PutTemplatesLibraryIdentifierOKBodyBody
*/
type PutTemplatesLibraryIdentifierOKBodyBody interface{}
