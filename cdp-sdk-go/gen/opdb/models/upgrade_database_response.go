// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpgradeDatabaseResponse Response with the reason whether upgrade request is accepted or why it is not possible.
//
// swagger:model UpgradeDatabaseResponse
type UpgradeDatabaseResponse struct {

	// The reason whether upgrade request is accepted or why it is not possible.
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this upgrade database response
func (m *UpgradeDatabaseResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this upgrade database response based on context it is used
func (m *UpgradeDatabaseResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UpgradeDatabaseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpgradeDatabaseResponse) UnmarshalBinary(b []byte) error {
	var res UpgradeDatabaseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
