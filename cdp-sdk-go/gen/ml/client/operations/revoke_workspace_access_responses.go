// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cloudera/terraform-provider-cdp/cdp-sdk-go/gen/ml/models"
)

// RevokeWorkspaceAccessReader is a Reader for the RevokeWorkspaceAccess structure.
type RevokeWorkspaceAccessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeWorkspaceAccessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeWorkspaceAccessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRevokeWorkspaceAccessDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRevokeWorkspaceAccessOK creates a RevokeWorkspaceAccessOK with default headers values
func NewRevokeWorkspaceAccessOK() *RevokeWorkspaceAccessOK {
	return &RevokeWorkspaceAccessOK{}
}

/*RevokeWorkspaceAccessOK handles this case with default header values.

Expected response to a valid request.
*/
type RevokeWorkspaceAccessOK struct {
	Payload models.RevokeWorkspaceAccessResponse
}

func (o *RevokeWorkspaceAccessOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/revokeWorkspaceAccess][%d] revokeWorkspaceAccessOK  %+v", 200, o.Payload)
}

func (o *RevokeWorkspaceAccessOK) GetPayload() models.RevokeWorkspaceAccessResponse {
	return o.Payload
}

func (o *RevokeWorkspaceAccessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeWorkspaceAccessDefault creates a RevokeWorkspaceAccessDefault with default headers values
func NewRevokeWorkspaceAccessDefault(code int) *RevokeWorkspaceAccessDefault {
	return &RevokeWorkspaceAccessDefault{
		_statusCode: code,
	}
}

/*RevokeWorkspaceAccessDefault handles this case with default header values.

The default response on an error.
*/
type RevokeWorkspaceAccessDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the revoke workspace access default response
func (o *RevokeWorkspaceAccessDefault) Code() int {
	return o._statusCode
}

func (o *RevokeWorkspaceAccessDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/revokeWorkspaceAccess][%d] revokeWorkspaceAccess default  %+v", o._statusCode, o.Payload)
}

func (o *RevokeWorkspaceAccessDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *RevokeWorkspaceAccessDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
