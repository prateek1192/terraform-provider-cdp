// Code generated by go-swagger; DO NOT EDIT.

package operations

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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/opdb/models"
)

// NewDropDatabaseParams creates a new DropDatabaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDropDatabaseParams() *DropDatabaseParams {
	return &DropDatabaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDropDatabaseParamsWithTimeout creates a new DropDatabaseParams object
// with the ability to set a timeout on a request.
func NewDropDatabaseParamsWithTimeout(timeout time.Duration) *DropDatabaseParams {
	return &DropDatabaseParams{
		timeout: timeout,
	}
}

// NewDropDatabaseParamsWithContext creates a new DropDatabaseParams object
// with the ability to set a context for a request.
func NewDropDatabaseParamsWithContext(ctx context.Context) *DropDatabaseParams {
	return &DropDatabaseParams{
		Context: ctx,
	}
}

// NewDropDatabaseParamsWithHTTPClient creates a new DropDatabaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewDropDatabaseParamsWithHTTPClient(client *http.Client) *DropDatabaseParams {
	return &DropDatabaseParams{
		HTTPClient: client,
	}
}

/*
DropDatabaseParams contains all the parameters to send to the API endpoint

	for the drop database operation.

	Typically these are written to a http.Request.
*/
type DropDatabaseParams struct {

	// Input.
	Input *models.DropDatabaseRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the drop database params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DropDatabaseParams) WithDefaults() *DropDatabaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the drop database params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DropDatabaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the drop database params
func (o *DropDatabaseParams) WithTimeout(timeout time.Duration) *DropDatabaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the drop database params
func (o *DropDatabaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the drop database params
func (o *DropDatabaseParams) WithContext(ctx context.Context) *DropDatabaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the drop database params
func (o *DropDatabaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the drop database params
func (o *DropDatabaseParams) WithHTTPClient(client *http.Client) *DropDatabaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the drop database params
func (o *DropDatabaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the drop database params
func (o *DropDatabaseParams) WithInput(input *models.DropDatabaseRequest) *DropDatabaseParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the drop database params
func (o *DropDatabaseParams) SetInput(input *models.DropDatabaseRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DropDatabaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Input != nil {
		if err := r.SetBodyParam(o.Input); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
