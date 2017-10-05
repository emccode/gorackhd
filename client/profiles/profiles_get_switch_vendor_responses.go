// Code generated by go-swagger; DO NOT EDIT.

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

// ProfilesGetSwitchVendorReader is a Reader for the ProfilesGetSwitchVendor structure.
type ProfilesGetSwitchVendorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProfilesGetSwitchVendorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProfilesGetSwitchVendorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewProfilesGetSwitchVendorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewProfilesGetSwitchVendorDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProfilesGetSwitchVendorOK creates a ProfilesGetSwitchVendorOK with default headers values
func NewProfilesGetSwitchVendorOK() *ProfilesGetSwitchVendorOK {
	return &ProfilesGetSwitchVendorOK{}
}

/*ProfilesGetSwitchVendorOK handles this case with default header values.

Successfully returned the profile with switch vendor name.
*/
type ProfilesGetSwitchVendorOK struct {
	Payload ProfilesGetSwitchVendorOKBody
}

func (o *ProfilesGetSwitchVendorOK) Error() string {
	return fmt.Sprintf("[GET /profiles/switch/{vendor}][%d] profilesGetSwitchVendorOK  %+v", 200, o.Payload)
}

func (o *ProfilesGetSwitchVendorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfilesGetSwitchVendorNotFound creates a ProfilesGetSwitchVendorNotFound with default headers values
func NewProfilesGetSwitchVendorNotFound() *ProfilesGetSwitchVendorNotFound {
	return &ProfilesGetSwitchVendorNotFound{}
}

/*ProfilesGetSwitchVendorNotFound handles this case with default header values.

Profile not found
*/
type ProfilesGetSwitchVendorNotFound struct {
	Payload *models.Error
}

func (o *ProfilesGetSwitchVendorNotFound) Error() string {
	return fmt.Sprintf("[GET /profiles/switch/{vendor}][%d] profilesGetSwitchVendorNotFound  %+v", 404, o.Payload)
}

func (o *ProfilesGetSwitchVendorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProfilesGetSwitchVendorDefault creates a ProfilesGetSwitchVendorDefault with default headers values
func NewProfilesGetSwitchVendorDefault(code int) *ProfilesGetSwitchVendorDefault {
	return &ProfilesGetSwitchVendorDefault{
		_statusCode: code,
	}
}

/*ProfilesGetSwitchVendorDefault handles this case with default header values.

Unexpected error
*/
type ProfilesGetSwitchVendorDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the profiles get switch vendor default response
func (o *ProfilesGetSwitchVendorDefault) Code() int {
	return o._statusCode
}

func (o *ProfilesGetSwitchVendorDefault) Error() string {
	return fmt.Sprintf("[GET /profiles/switch/{vendor}][%d] profilesGetSwitchVendor default  %+v", o._statusCode, o.Payload)
}

func (o *ProfilesGetSwitchVendorDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ProfilesGetSwitchVendorOKBody profiles get switch vendor o k body
swagger:model ProfilesGetSwitchVendorOKBody
*/

type ProfilesGetSwitchVendorOKBody interface{}
