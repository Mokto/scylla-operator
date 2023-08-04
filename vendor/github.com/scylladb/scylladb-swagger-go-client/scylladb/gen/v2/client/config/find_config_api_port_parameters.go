// Code generated by go-swagger; DO NOT EDIT.

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewFindConfigAPIPortParams creates a new FindConfigAPIPortParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindConfigAPIPortParams() *FindConfigAPIPortParams {
	return &FindConfigAPIPortParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigAPIPortParamsWithTimeout creates a new FindConfigAPIPortParams object
// with the ability to set a timeout on a request.
func NewFindConfigAPIPortParamsWithTimeout(timeout time.Duration) *FindConfigAPIPortParams {
	return &FindConfigAPIPortParams{
		timeout: timeout,
	}
}

// NewFindConfigAPIPortParamsWithContext creates a new FindConfigAPIPortParams object
// with the ability to set a context for a request.
func NewFindConfigAPIPortParamsWithContext(ctx context.Context) *FindConfigAPIPortParams {
	return &FindConfigAPIPortParams{
		Context: ctx,
	}
}

// NewFindConfigAPIPortParamsWithHTTPClient creates a new FindConfigAPIPortParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindConfigAPIPortParamsWithHTTPClient(client *http.Client) *FindConfigAPIPortParams {
	return &FindConfigAPIPortParams{
		HTTPClient: client,
	}
}

/*
FindConfigAPIPortParams contains all the parameters to send to the API endpoint

	for the find config api port operation.

	Typically these are written to a http.Request.
*/
type FindConfigAPIPortParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find config api port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigAPIPortParams) WithDefaults() *FindConfigAPIPortParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find config api port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindConfigAPIPortParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find config api port params
func (o *FindConfigAPIPortParams) WithTimeout(timeout time.Duration) *FindConfigAPIPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config api port params
func (o *FindConfigAPIPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config api port params
func (o *FindConfigAPIPortParams) WithContext(ctx context.Context) *FindConfigAPIPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config api port params
func (o *FindConfigAPIPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config api port params
func (o *FindConfigAPIPortParams) WithHTTPClient(client *http.Client) *FindConfigAPIPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config api port params
func (o *FindConfigAPIPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigAPIPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
