package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/codedellemc/gorackhd/models"
)

// GetObmsLibraryIdentifierReader is a Reader for the GetObmsLibraryIdentifier structure.
type GetObmsLibraryIdentifierReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetObmsLibraryIdentifierReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetObmsLibraryIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetObmsLibraryIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetObmsLibraryIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetObmsLibraryIdentifierOK creates a GetObmsLibraryIdentifierOK with default headers values
func NewGetObmsLibraryIdentifierOK() *GetObmsLibraryIdentifierOK {
	return &GetObmsLibraryIdentifierOK{}
}

/*GetObmsLibraryIdentifierOK handles this case with default header values.

return OBM service

*/
type GetObmsLibraryIdentifierOK struct {
	Payload GetObmsLibraryIdentifierOKBodyBody
}

func (o *GetObmsLibraryIdentifierOK) Error() string {
	return fmt.Sprintf("[GET /obms/library/{identifier}][%d] getObmsLibraryIdentifierOK  %+v", 200, o.Payload)
}

func (o *GetObmsLibraryIdentifierOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetObmsLibraryIdentifierNotFound creates a GetObmsLibraryIdentifierNotFound with default headers values
func NewGetObmsLibraryIdentifierNotFound() *GetObmsLibraryIdentifierNotFound {
	return &GetObmsLibraryIdentifierNotFound{}
}

/*GetObmsLibraryIdentifierNotFound handles this case with default header values.

The obm service with the identifier was not found

*/
type GetObmsLibraryIdentifierNotFound struct {
	Payload *models.Error
}

func (o *GetObmsLibraryIdentifierNotFound) Error() string {
	return fmt.Sprintf("[GET /obms/library/{identifier}][%d] getObmsLibraryIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *GetObmsLibraryIdentifierNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetObmsLibraryIdentifierDefault creates a GetObmsLibraryIdentifierDefault with default headers values
func NewGetObmsLibraryIdentifierDefault(code int) *GetObmsLibraryIdentifierDefault {
	return &GetObmsLibraryIdentifierDefault{
		_statusCode: code,
	}
}

/*GetObmsLibraryIdentifierDefault handles this case with default header values.

Unexpected error
*/
type GetObmsLibraryIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get obms library identifier default response
func (o *GetObmsLibraryIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *GetObmsLibraryIdentifierDefault) Error() string {
	return fmt.Sprintf("[GET /obms/library/{identifier}][%d] GetObmsLibraryIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *GetObmsLibraryIdentifierDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetObmsLibraryIdentifierOKBodyBody get obms library identifier o k body body

swagger:model GetObmsLibraryIdentifierOKBodyBody
*/
type GetObmsLibraryIdentifierOKBodyBody interface{}
