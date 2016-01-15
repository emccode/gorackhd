package patch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd/models"
)

type PatchSkusIdentifierReader struct {
	formats strfmt.Registry
}

func (o *PatchSkusIdentifierReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchSkusIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPatchSkusIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchSkusIdentifierInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchSkusIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPatchSkusIdentifierOK creates a PatchSkusIdentifierOK with default headers values
func NewPatchSkusIdentifierOK() *PatchSkusIdentifierOK {
	return &PatchSkusIdentifierOK{}
}

/*PatchSkusIdentifierOK

sku to patch

*/
type PatchSkusIdentifierOK struct {
	Payload PatchSkusIdentifierOKBodyBody
}

func (o *PatchSkusIdentifierOK) Error() string {
	return fmt.Sprintf("[PATCH /skus/{identifier}][%d] patchSkusIdentifierOK  %+v", 200, o.Payload)
}

func (o *PatchSkusIdentifierOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPatchSkusIdentifierNotFound creates a PatchSkusIdentifierNotFound with default headers values
func NewPatchSkusIdentifierNotFound() *PatchSkusIdentifierNotFound {
	return &PatchSkusIdentifierNotFound{}
}

/*PatchSkusIdentifierNotFound

Not found, no sku with identifier.

*/
type PatchSkusIdentifierNotFound struct {
	Payload *models.Error
}

func (o *PatchSkusIdentifierNotFound) Error() string {
	return fmt.Sprintf("[PATCH /skus/{identifier}][%d] patchSkusIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *PatchSkusIdentifierNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPatchSkusIdentifierInternalServerError creates a PatchSkusIdentifierInternalServerError with default headers values
func NewPatchSkusIdentifierInternalServerError() *PatchSkusIdentifierInternalServerError {
	return &PatchSkusIdentifierInternalServerError{}
}

/*PatchSkusIdentifierInternalServerError

Patch failed.

*/
type PatchSkusIdentifierInternalServerError struct {
	Payload *models.Error
}

func (o *PatchSkusIdentifierInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /skus/{identifier}][%d] patchSkusIdentifierInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchSkusIdentifierInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPatchSkusIdentifierDefault creates a PatchSkusIdentifierDefault with default headers values
func NewPatchSkusIdentifierDefault(code int) *PatchSkusIdentifierDefault {
	return &PatchSkusIdentifierDefault{
		_statusCode: code,
	}
}

/*PatchSkusIdentifierDefault

Unexpected error
*/
type PatchSkusIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch skus identifier default response
func (o *PatchSkusIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *PatchSkusIdentifierDefault) Error() string {
	return fmt.Sprintf("[PATCH /skus/{identifier}][%d] PatchSkusIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *PatchSkusIdentifierDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*PatchSkusIdentifierOKBodyBody PatchSkusIdentifierOKBodyBody patch skus identifier o k body body

swagger:model PatchSkusIdentifierOKBodyBody
*/
type PatchSkusIdentifierOKBodyBody interface{}
