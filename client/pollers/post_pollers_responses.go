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

type PostPollersReader struct {
	formats strfmt.Registry
}

func (o *PostPollersReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPollersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPostPollersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostPollersOK creates a PostPollersOK with default headers values
func NewPostPollersOK() *PostPollersOK {
	return &PostPollersOK{}
}

/*PostPollersOK

poller created

*/
type PostPollersOK struct {
	Payload PostPollersOKBodyBody
}

func (o *PostPollersOK) Error() string {
	return fmt.Sprintf("[POST /pollers][%d] postPollersOK  %+v", 200, o.Payload)
}

func (o *PostPollersOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostPollersDefault creates a PostPollersDefault with default headers values
func NewPostPollersDefault(code int) *PostPollersDefault {
	return &PostPollersDefault{
		_statusCode: code,
	}
}

/*PostPollersDefault

Unexpected error
*/
type PostPollersDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post pollers default response
func (o *PostPollersDefault) Code() int {
	return o._statusCode
}

func (o *PostPollersDefault) Error() string {
	return fmt.Sprintf("[POST /pollers][%d] PostPollers default  %+v", o._statusCode, o.Payload)
}

func (o *PostPollersDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*PostPollersOKBodyBody PostPollersOKBodyBody post pollers o k body body

swagger:model PostPollersOKBodyBody
*/
type PostPollersOKBodyBody interface{}
