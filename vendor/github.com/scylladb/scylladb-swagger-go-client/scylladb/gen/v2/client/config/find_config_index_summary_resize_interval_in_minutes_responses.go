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

// FindConfigIndexSummaryResizeIntervalInMinutesReader is a Reader for the FindConfigIndexSummaryResizeIntervalInMinutes structure.
type FindConfigIndexSummaryResizeIntervalInMinutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindConfigIndexSummaryResizeIntervalInMinutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindConfigIndexSummaryResizeIntervalInMinutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindConfigIndexSummaryResizeIntervalInMinutesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindConfigIndexSummaryResizeIntervalInMinutesOK creates a FindConfigIndexSummaryResizeIntervalInMinutesOK with default headers values
func NewFindConfigIndexSummaryResizeIntervalInMinutesOK() *FindConfigIndexSummaryResizeIntervalInMinutesOK {
	return &FindConfigIndexSummaryResizeIntervalInMinutesOK{}
}

/*
FindConfigIndexSummaryResizeIntervalInMinutesOK handles this case with default header values.

Config value
*/
type FindConfigIndexSummaryResizeIntervalInMinutesOK struct {
	Payload int64
}

func (o *FindConfigIndexSummaryResizeIntervalInMinutesOK) GetPayload() int64 {
	return o.Payload
}

func (o *FindConfigIndexSummaryResizeIntervalInMinutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindConfigIndexSummaryResizeIntervalInMinutesDefault creates a FindConfigIndexSummaryResizeIntervalInMinutesDefault with default headers values
func NewFindConfigIndexSummaryResizeIntervalInMinutesDefault(code int) *FindConfigIndexSummaryResizeIntervalInMinutesDefault {
	return &FindConfigIndexSummaryResizeIntervalInMinutesDefault{
		_statusCode: code,
	}
}

/*
FindConfigIndexSummaryResizeIntervalInMinutesDefault handles this case with default header values.

unexpected error
*/
type FindConfigIndexSummaryResizeIntervalInMinutesDefault struct {
	_statusCode int

	Payload *models.ErrorModel
}

// Code gets the status code for the find config index summary resize interval in minutes default response
func (o *FindConfigIndexSummaryResizeIntervalInMinutesDefault) Code() int {
	return o._statusCode
}

func (o *FindConfigIndexSummaryResizeIntervalInMinutesDefault) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *FindConfigIndexSummaryResizeIntervalInMinutesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (o *FindConfigIndexSummaryResizeIntervalInMinutesDefault) Error() string {
	return fmt.Sprintf("agent [HTTP %d] %s", o._statusCode, strings.TrimRight(o.Payload.Message, "."))
}
