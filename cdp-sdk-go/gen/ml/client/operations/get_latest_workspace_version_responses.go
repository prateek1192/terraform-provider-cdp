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

// GetLatestWorkspaceVersionReader is a Reader for the GetLatestWorkspaceVersion structure.
type GetLatestWorkspaceVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestWorkspaceVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestWorkspaceVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetLatestWorkspaceVersionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLatestWorkspaceVersionOK creates a GetLatestWorkspaceVersionOK with default headers values
func NewGetLatestWorkspaceVersionOK() *GetLatestWorkspaceVersionOK {
	return &GetLatestWorkspaceVersionOK{}
}

/*GetLatestWorkspaceVersionOK handles this case with default header values.

Expected response to a valid request.
*/
type GetLatestWorkspaceVersionOK struct {
	Payload *models.GetLatestWorkspaceVersionResponse
}

func (o *GetLatestWorkspaceVersionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/getLatestWorkspaceVersion][%d] getLatestWorkspaceVersionOK  %+v", 200, o.Payload)
}

func (o *GetLatestWorkspaceVersionOK) GetPayload() *models.GetLatestWorkspaceVersionResponse {
	return o.Payload
}

func (o *GetLatestWorkspaceVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetLatestWorkspaceVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestWorkspaceVersionDefault creates a GetLatestWorkspaceVersionDefault with default headers values
func NewGetLatestWorkspaceVersionDefault(code int) *GetLatestWorkspaceVersionDefault {
	return &GetLatestWorkspaceVersionDefault{
		_statusCode: code,
	}
}

/*GetLatestWorkspaceVersionDefault handles this case with default header values.

The default response on an error.
*/
type GetLatestWorkspaceVersionDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get latest workspace version default response
func (o *GetLatestWorkspaceVersionDefault) Code() int {
	return o._statusCode
}

func (o *GetLatestWorkspaceVersionDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/ml/getLatestWorkspaceVersion][%d] getLatestWorkspaceVersion default  %+v", o._statusCode, o.Payload)
}

func (o *GetLatestWorkspaceVersionDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetLatestWorkspaceVersionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
