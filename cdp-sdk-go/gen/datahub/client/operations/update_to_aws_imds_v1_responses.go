// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/datahub/models"
)

// UpdateToAwsImdsV1Reader is a Reader for the UpdateToAwsImdsV1 structure.
type UpdateToAwsImdsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateToAwsImdsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateToAwsImdsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateToAwsImdsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateToAwsImdsV1OK creates a UpdateToAwsImdsV1OK with default headers values
func NewUpdateToAwsImdsV1OK() *UpdateToAwsImdsV1OK {
	return &UpdateToAwsImdsV1OK{}
}

/*
UpdateToAwsImdsV1OK describes a response with status code 200, with default header values.

Expected response to a valid request.
*/
type UpdateToAwsImdsV1OK struct {
	Payload models.UpdateToAwsImdsV1Response
}

// IsSuccess returns true when this update to aws imds v1 o k response has a 2xx status code
func (o *UpdateToAwsImdsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update to aws imds v1 o k response has a 3xx status code
func (o *UpdateToAwsImdsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update to aws imds v1 o k response has a 4xx status code
func (o *UpdateToAwsImdsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update to aws imds v1 o k response has a 5xx status code
func (o *UpdateToAwsImdsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this update to aws imds v1 o k response a status code equal to that given
func (o *UpdateToAwsImdsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update to aws imds v1 o k response
func (o *UpdateToAwsImdsV1OK) Code() int {
	return 200
}

func (o *UpdateToAwsImdsV1OK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/updateToAwsImdsV1][%d] updateToAwsImdsV1OK  %+v", 200, o.Payload)
}

func (o *UpdateToAwsImdsV1OK) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/updateToAwsImdsV1][%d] updateToAwsImdsV1OK  %+v", 200, o.Payload)
}

func (o *UpdateToAwsImdsV1OK) GetPayload() models.UpdateToAwsImdsV1Response {
	return o.Payload
}

func (o *UpdateToAwsImdsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateToAwsImdsV1Default creates a UpdateToAwsImdsV1Default with default headers values
func NewUpdateToAwsImdsV1Default(code int) *UpdateToAwsImdsV1Default {
	return &UpdateToAwsImdsV1Default{
		_statusCode: code,
	}
}

/*
UpdateToAwsImdsV1Default describes a response with status code -1, with default header values.

The default response on an error.
*/
type UpdateToAwsImdsV1Default struct {
	_statusCode int

	Payload *models.Error
}

// IsSuccess returns true when this update to aws imds v1 default response has a 2xx status code
func (o *UpdateToAwsImdsV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update to aws imds v1 default response has a 3xx status code
func (o *UpdateToAwsImdsV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update to aws imds v1 default response has a 4xx status code
func (o *UpdateToAwsImdsV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update to aws imds v1 default response has a 5xx status code
func (o *UpdateToAwsImdsV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update to aws imds v1 default response a status code equal to that given
func (o *UpdateToAwsImdsV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update to aws imds v1 default response
func (o *UpdateToAwsImdsV1Default) Code() int {
	return o._statusCode
}

func (o *UpdateToAwsImdsV1Default) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/updateToAwsImdsV1][%d] updateToAwsImdsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateToAwsImdsV1Default) String() string {
	return fmt.Sprintf("[POST /api/v1/datahub/updateToAwsImdsV1][%d] updateToAwsImdsV1 default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateToAwsImdsV1Default) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateToAwsImdsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
