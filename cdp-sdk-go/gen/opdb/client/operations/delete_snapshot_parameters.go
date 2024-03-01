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

// NewDeleteSnapshotParams creates a new DeleteSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSnapshotParams() *DeleteSnapshotParams {
	return &DeleteSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnapshotParamsWithTimeout creates a new DeleteSnapshotParams object
// with the ability to set a timeout on a request.
func NewDeleteSnapshotParamsWithTimeout(timeout time.Duration) *DeleteSnapshotParams {
	return &DeleteSnapshotParams{
		timeout: timeout,
	}
}

// NewDeleteSnapshotParamsWithContext creates a new DeleteSnapshotParams object
// with the ability to set a context for a request.
func NewDeleteSnapshotParamsWithContext(ctx context.Context) *DeleteSnapshotParams {
	return &DeleteSnapshotParams{
		Context: ctx,
	}
}

// NewDeleteSnapshotParamsWithHTTPClient creates a new DeleteSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSnapshotParamsWithHTTPClient(client *http.Client) *DeleteSnapshotParams {
	return &DeleteSnapshotParams{
		HTTPClient: client,
	}
}

/*
DeleteSnapshotParams contains all the parameters to send to the API endpoint

	for the delete snapshot operation.

	Typically these are written to a http.Request.
*/
type DeleteSnapshotParams struct {

	// Input.
	Input *models.DeleteSnapshotRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSnapshotParams) WithDefaults() *DeleteSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete snapshot params
func (o *DeleteSnapshotParams) WithTimeout(timeout time.Duration) *DeleteSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete snapshot params
func (o *DeleteSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete snapshot params
func (o *DeleteSnapshotParams) WithContext(ctx context.Context) *DeleteSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete snapshot params
func (o *DeleteSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete snapshot params
func (o *DeleteSnapshotParams) WithHTTPClient(client *http.Client) *DeleteSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete snapshot params
func (o *DeleteSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the delete snapshot params
func (o *DeleteSnapshotParams) WithInput(input *models.DeleteSnapshotRequest) *DeleteSnapshotParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the delete snapshot params
func (o *DeleteSnapshotParams) SetInput(input *models.DeleteSnapshotRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
