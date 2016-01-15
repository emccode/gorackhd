package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd/models"
)

type GetConfigReader struct {
	formats strfmt.Registry
}

func (o *GetConfigReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetConfigOK creates a GetConfigOK with default headers values
func NewGetConfigOK() *GetConfigOK {
	return &GetConfigOK{}
}

/*GetConfigOK

An array of configuration objects
*/
type GetConfigOK struct {
	Payload GetConfigOKBodyBody
}

func (o *GetConfigOK) Error() string {
	return fmt.Sprintf("[GET /config][%d] getConfigOK  %+v", 200, o.Payload)
}

func (o *GetConfigOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetConfigDefault creates a GetConfigDefault with default headers values
func NewGetConfigDefault(code int) *GetConfigDefault {
	return &GetConfigDefault{
		_statusCode: code,
	}
}

/*GetConfigDefault

Unexpected error
*/
type GetConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get config default response
func (o *GetConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetConfigDefault) Error() string {
	return fmt.Sprintf("[GET /config][%d] GetConfig default  %+v", o._statusCode, o.Payload)
}

func (o *GetConfigDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*GetConfigOKBodyBody GetConfigOKBodyBody get config o k body body

swagger:model GetConfigOKBodyBody
*/
type GetConfigOKBodyBody interface{}
