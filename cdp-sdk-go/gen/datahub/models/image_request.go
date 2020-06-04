// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ImageRequest The details of the image used for cluster instances.
//
// swagger:model ImageRequest
type ImageRequest struct {

	// The image catalog name.
	// Required: true
	CatalogName *string `json:"catalogName"`

	// The ID of the image used for cluster instances. This is generated by the cloud provider to uniquely identify the image.
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this image request
func (m *ImageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalogName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImageRequest) validateCatalogName(formats strfmt.Registry) error {

	if err := validate.Required("catalogName", "body", m.CatalogName); err != nil {
		return err
	}

	return nil
}

func (m *ImageRequest) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ImageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImageRequest) UnmarshalBinary(b []byte) error {
	var res ImageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
