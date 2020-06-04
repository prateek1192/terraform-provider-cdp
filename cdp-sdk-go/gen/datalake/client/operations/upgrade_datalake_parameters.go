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

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datalake/models"
)

// NewUpgradeDatalakeParams creates a new UpgradeDatalakeParams object
// with the default values initialized.
func NewUpgradeDatalakeParams() *UpgradeDatalakeParams {
	var ()
	return &UpgradeDatalakeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpgradeDatalakeParamsWithTimeout creates a new UpgradeDatalakeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpgradeDatalakeParamsWithTimeout(timeout time.Duration) *UpgradeDatalakeParams {
	var ()
	return &UpgradeDatalakeParams{

		timeout: timeout,
	}
}

// NewUpgradeDatalakeParamsWithContext creates a new UpgradeDatalakeParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpgradeDatalakeParamsWithContext(ctx context.Context) *UpgradeDatalakeParams {
	var ()
	return &UpgradeDatalakeParams{

		Context: ctx,
	}
}

// NewUpgradeDatalakeParamsWithHTTPClient creates a new UpgradeDatalakeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpgradeDatalakeParamsWithHTTPClient(client *http.Client) *UpgradeDatalakeParams {
	var ()
	return &UpgradeDatalakeParams{
		HTTPClient: client,
	}
}

/*UpgradeDatalakeParams contains all the parameters to send to the API endpoint
for the upgrade datalake operation typically these are written to a http.Request
*/
type UpgradeDatalakeParams struct {

	/*Input*/
	Input *models.UpgradeDatalakeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upgrade datalake params
func (o *UpgradeDatalakeParams) WithTimeout(timeout time.Duration) *UpgradeDatalakeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upgrade datalake params
func (o *UpgradeDatalakeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upgrade datalake params
func (o *UpgradeDatalakeParams) WithContext(ctx context.Context) *UpgradeDatalakeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upgrade datalake params
func (o *UpgradeDatalakeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upgrade datalake params
func (o *UpgradeDatalakeParams) WithHTTPClient(client *http.Client) *UpgradeDatalakeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upgrade datalake params
func (o *UpgradeDatalakeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInput adds the input to the upgrade datalake params
func (o *UpgradeDatalakeParams) WithInput(input *models.UpgradeDatalakeRequest) *UpgradeDatalakeParams {
	o.SetInput(input)
	return o
}

// SetInput adds the input to the upgrade datalake params
func (o *UpgradeDatalakeParams) SetInput(input *models.UpgradeDatalakeRequest) {
	o.Input = input
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeDatalakeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
