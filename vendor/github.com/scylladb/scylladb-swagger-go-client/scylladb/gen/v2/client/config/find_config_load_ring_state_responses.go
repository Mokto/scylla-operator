// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strings"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/scylladb/scylladb-swagger-go-client/scylladb/gen/v2/models"
)

// FindConfigLoadRingStateReader is a Reader for the FindConfigLoadRingState structure.
type FindConfigLoadRingStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigLoadRingStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigLoadRingStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigLoadRingStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigLoadRingStateOK creates a FindConfigLoadRingStateOK with default headers values
func NewFindConfigLoadRingStateOK() *FindConfigLoadRingStateOK {
	return &FindConfigLoadRingStateOK{}
}

/*
FindConfigLoadRingStateOK handles this case with default header values.

Config value
*/
type FindConfigLoadRingStateOK struct {
	Payload bool
}

func (o *FindConfigLoadRingStateOK) GetPayload() bool {
	return o.Payload
}

func (o *FindConfigLoadRingStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigLoadRingStateDefault creates a FindConfigLoadRingStateDefault with default headers values
func NewFindConfigLoadRingStateDefault(code int) *FindConfigLoadRingStateDefault {
	return &FindConfigLoadRingStateDefault{
		_statusCode: code,
	}
}

/*
FindConfigLoadRingStateDefault handles this case with default header values.

unexpected error
*/
type FindConfigLoadRingStateDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config load ring state default response
func (o *FindConfigLoadRingStateDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigLoadRingStateDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigLoadRingStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigLoadRingStateDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
