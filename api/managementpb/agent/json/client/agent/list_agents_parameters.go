// Code generated by go-swagger; DO NOT EDIT.

package agent

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

// NewListAgentsParams creates a new ListAgentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAgentsParams() *ListAgentsParams {
	return &ListAgentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAgentsParamsWithTimeout creates a new ListAgentsParams object
// with the ability to set a timeout on a request.
func NewListAgentsParamsWithTimeout(timeout time.Duration) *ListAgentsParams {
	return &ListAgentsParams{
		timeout: timeout,
	}
}

// NewListAgentsParamsWithContext creates a new ListAgentsParams object
// with the ability to set a context for a request.
func NewListAgentsParamsWithContext(ctx context.Context) *ListAgentsParams {
	return &ListAgentsParams{
		Context: ctx,
	}
}

// NewListAgentsParamsWithHTTPClient creates a new ListAgentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAgentsParamsWithHTTPClient(client *http.Client) *ListAgentsParams {
	return &ListAgentsParams{
		HTTPClient: client,
	}
}

/*
ListAgentsParams contains all the parameters to send to the API endpoint

	for the list agents operation.

	Typically these are written to a http.Request.
*/
type ListAgentsParams struct {
	/* Body.

	   Only one of the parameters below must be set.
	*/
	Body ListAgentsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list agents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAgentsParams) WithDefaults() *ListAgentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list agents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAgentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list agents params
func (o *ListAgentsParams) WithTimeout(timeout time.Duration) *ListAgentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list agents params
func (o *ListAgentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list agents params
func (o *ListAgentsParams) WithContext(ctx context.Context) *ListAgentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list agents params
func (o *ListAgentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list agents params
func (o *ListAgentsParams) WithHTTPClient(client *http.Client) *ListAgentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list agents params
func (o *ListAgentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list agents params
func (o *ListAgentsParams) WithBody(body ListAgentsBody) *ListAgentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list agents params
func (o *ListAgentsParams) SetBody(body ListAgentsBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListAgentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
