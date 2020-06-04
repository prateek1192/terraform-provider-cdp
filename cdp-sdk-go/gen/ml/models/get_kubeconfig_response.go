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

// GetKubeconfigResponse Response object for the GetKubeconfig method.
//
// swagger:model GetKubeconfigResponse
type GetKubeconfigResponse struct {

	// The list of users that have access.
	// Required: true
	Kubeconfig *string `json:"kubeconfig"`
}

// Validate validates this get kubeconfig response
func (m *GetKubeconfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKubeconfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetKubeconfigResponse) validateKubeconfig(formats strfmt.Registry) error {

	if err := validate.Required("kubeconfig", "body", m.Kubeconfig); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetKubeconfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetKubeconfigResponse) UnmarshalBinary(b []byte) error {
	var res GetKubeconfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
