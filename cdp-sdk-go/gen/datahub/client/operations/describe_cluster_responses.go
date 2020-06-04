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

// DescribeClusterReader is a Reader for the DescribeCluster structure.
type DescribeClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDescribeClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDescribeClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDescribeClusterOK creates a DescribeClusterOK with default headers values
func NewDescribeClusterOK() *DescribeClusterOK {
	return &DescribeClusterOK{}
}

/*DescribeClusterOK handles this case with default header values.

Expected response to a valid request.
*/
type DescribeClusterOK struct {
	Payload *models.DescribeClusterResponse
}

func (o *DescribeClusterOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/describeCluster][%d] describeClusterOK  %+v", 200, o.Payload)
}

func (o *DescribeClusterOK) GetPayload() *models.DescribeClusterResponse {
	return o.Payload
}

func (o *DescribeClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DescribeClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDescribeClusterDefault creates a DescribeClusterDefault with default headers values
func NewDescribeClusterDefault(code int) *DescribeClusterDefault {
	return &DescribeClusterDefault{
		_statusCode: code,
	}
}

/*DescribeClusterDefault handles this case with default header values.

The default response on an error.
*/
type DescribeClusterDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the describe cluster default response
func (o *DescribeClusterDefault) Code() int {
	return o._statusCode
}

func (o *DescribeClusterDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/datahub/describeCluster][%d] describeCluster default  %+v", o._statusCode, o.Payload)
}

func (o *DescribeClusterDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *DescribeClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
