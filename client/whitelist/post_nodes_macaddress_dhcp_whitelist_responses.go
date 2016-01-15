package whitelist

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd/models"
)

type PostNodesMacaddressDhcpWhitelistReader struct {
	formats strfmt.Registry
}

func (o *PostNodesMacaddressDhcpWhitelistReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostNodesMacaddressDhcpWhitelistCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPostNodesMacaddressDhcpWhitelistNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostNodesMacaddressDhcpWhitelistDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostNodesMacaddressDhcpWhitelistCreated creates a PostNodesMacaddressDhcpWhitelistCreated with default headers values
func NewPostNodesMacaddressDhcpWhitelistCreated() *PostNodesMacaddressDhcpWhitelistCreated {
	return &PostNodesMacaddressDhcpWhitelistCreated{}
}

/*PostNodesMacaddressDhcpWhitelistCreated

the add was successful and it returns the whitelist

*/
type PostNodesMacaddressDhcpWhitelistCreated struct {
	Payload PostNodesMacaddressDhcpWhitelistCreatedBodyBody
}

func (o *PostNodesMacaddressDhcpWhitelistCreated) Error() string {
	return fmt.Sprintf("[POST /nodes/{macaddress}/dhcp/whitelist][%d] postNodesMacaddressDhcpWhitelistCreated  %+v", 201, o.Payload)
}

func (o *PostNodesMacaddressDhcpWhitelistCreated) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostNodesMacaddressDhcpWhitelistNotFound creates a PostNodesMacaddressDhcpWhitelistNotFound with default headers values
func NewPostNodesMacaddressDhcpWhitelistNotFound() *PostNodesMacaddressDhcpWhitelistNotFound {
	return &PostNodesMacaddressDhcpWhitelistNotFound{}
}

/*PostNodesMacaddressDhcpWhitelistNotFound

The node with the identifier was not found

*/
type PostNodesMacaddressDhcpWhitelistNotFound struct {
	Payload *models.Error
}

func (o *PostNodesMacaddressDhcpWhitelistNotFound) Error() string {
	return fmt.Sprintf("[POST /nodes/{macaddress}/dhcp/whitelist][%d] postNodesMacaddressDhcpWhitelistNotFound  %+v", 404, o.Payload)
}

func (o *PostNodesMacaddressDhcpWhitelistNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostNodesMacaddressDhcpWhitelistDefault creates a PostNodesMacaddressDhcpWhitelistDefault with default headers values
func NewPostNodesMacaddressDhcpWhitelistDefault(code int) *PostNodesMacaddressDhcpWhitelistDefault {
	return &PostNodesMacaddressDhcpWhitelistDefault{
		_statusCode: code,
	}
}

/*PostNodesMacaddressDhcpWhitelistDefault

Unexpected error
*/
type PostNodesMacaddressDhcpWhitelistDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post nodes macaddress dhcp whitelist default response
func (o *PostNodesMacaddressDhcpWhitelistDefault) Code() int {
	return o._statusCode
}

func (o *PostNodesMacaddressDhcpWhitelistDefault) Error() string {
	return fmt.Sprintf("[POST /nodes/{macaddress}/dhcp/whitelist][%d] PostNodesMacaddressDhcpWhitelist default  %+v", o._statusCode, o.Payload)
}

func (o *PostNodesMacaddressDhcpWhitelistDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*PostNodesMacaddressDhcpWhitelistCreatedBodyBody PostNodesMacaddressDhcpWhitelistCreatedBodyBody post nodes macaddress dhcp whitelist created body body

swagger:model PostNodesMacaddressDhcpWhitelistCreatedBodyBody
*/
type PostNodesMacaddressDhcpWhitelistCreatedBodyBody interface{}
