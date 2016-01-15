package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd/models"
)

type GetPollersLibraryIdentifierReader struct {
	formats strfmt.Registry
}

func (o *GetPollersLibraryIdentifierReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPollersLibraryIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetPollersLibraryIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetPollersLibraryIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetPollersLibraryIdentifierOK creates a GetPollersLibraryIdentifierOK with default headers values
func NewGetPollersLibraryIdentifierOK() *GetPollersLibraryIdentifierOK {
	return &GetPollersLibraryIdentifierOK{}
}

/*GetPollersLibraryIdentifierOK

single library poller

*/
type GetPollersLibraryIdentifierOK struct {
	Payload GetPollersLibraryIdentifierOKBodyBody
}

func (o *GetPollersLibraryIdentifierOK) Error() string {
	return fmt.Sprintf("[GET /pollers/library/{identifier}][%d] getPollersLibraryIdentifierOK  %+v", 200, o.Payload)
}

func (o *GetPollersLibraryIdentifierOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetPollersLibraryIdentifierNotFound creates a GetPollersLibraryIdentifierNotFound with default headers values
func NewGetPollersLibraryIdentifierNotFound() *GetPollersLibraryIdentifierNotFound {
	return &GetPollersLibraryIdentifierNotFound{}
}

/*GetPollersLibraryIdentifierNotFound

There is no library poller with specified identifier.

*/
type GetPollersLibraryIdentifierNotFound struct {
	Payload *models.Error
}

func (o *GetPollersLibraryIdentifierNotFound) Error() string {
	return fmt.Sprintf("[GET /pollers/library/{identifier}][%d] getPollersLibraryIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *GetPollersLibraryIdentifierNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetPollersLibraryIdentifierDefault creates a GetPollersLibraryIdentifierDefault with default headers values
func NewGetPollersLibraryIdentifierDefault(code int) *GetPollersLibraryIdentifierDefault {
	return &GetPollersLibraryIdentifierDefault{
		_statusCode: code,
	}
}

/*GetPollersLibraryIdentifierDefault

Unexpected error
*/
type GetPollersLibraryIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get pollers library identifier default response
func (o *GetPollersLibraryIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *GetPollersLibraryIdentifierDefault) Error() string {
	return fmt.Sprintf("[GET /pollers/library/{identifier}][%d] GetPollersLibraryIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *GetPollersLibraryIdentifierDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*GetPollersLibraryIdentifierOKBodyBody GetPollersLibraryIdentifierOKBodyBody get pollers library identifier o k body body

swagger:model GetPollersLibraryIdentifierOKBodyBody
*/
type GetPollersLibraryIdentifierOKBodyBody interface{}
