package profiles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// ProfilesGetMetadataReader is a Reader for the ProfilesGetMetadata structure.
type ProfilesGetMetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *ProfilesGetMetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProfilesGetMetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewProfilesGetMetadataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewProfilesGetMetadataOK creates a ProfilesGetMetadataOK with default headers values
func NewProfilesGetMetadataOK() *ProfilesGetMetadataOK {
	return &ProfilesGetMetadataOK{}
}

/*ProfilesGetMetadataOK handles this case with default header values.

Successfully retrieved the list of profile metadata
*/
type ProfilesGetMetadataOK struct {
	Payload ProfilesGetMetadataOKBodyBody
}

func (o *ProfilesGetMetadataOK) Error() string {
	return fmt.Sprintf("[GET /profiles/metadata][%d] profilesGetMetadataOK  %+v", 200, o.Payload)
}

func (o *ProfilesGetMetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfilesGetMetadataDefault creates a ProfilesGetMetadataDefault with default headers values
func NewProfilesGetMetadataDefault(code int) *ProfilesGetMetadataDefault {
	return &ProfilesGetMetadataDefault{
		_statusCode: code,
	}
}

/*ProfilesGetMetadataDefault handles this case with default header values.

Unexpected error
*/
type ProfilesGetMetadataDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the profiles get metadata default response
func (o *ProfilesGetMetadataDefault) Code() int {
	return o._statusCode
}

func (o *ProfilesGetMetadataDefault) Error() string {
	return fmt.Sprintf("[GET /profiles/metadata][%d] profilesGetMetadata default  %+v", o._statusCode, o.Payload)
}

func (o *ProfilesGetMetadataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProfilesGetMetadataOKBodyBody profiles get metadata o k body body

swagger:model ProfilesGetMetadataOKBodyBody
*/
type ProfilesGetMetadataOKBodyBody interface{}