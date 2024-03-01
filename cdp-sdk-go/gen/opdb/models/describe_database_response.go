// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DescribeDatabaseResponse Details of the database.
//
// swagger:model DescribeDatabaseResponse
type DescribeDatabaseResponse struct {

	// The details of the database.
	DatabaseDetails *DatabaseDetails `json:"databaseDetails,omitempty"`
}

// Validate validates this describe database response
func (m *DescribeDatabaseResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeDatabaseResponse) validateDatabaseDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.DatabaseDetails) { // not required
		return nil
	}

	if m.DatabaseDetails != nil {
		if err := m.DatabaseDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databaseDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("databaseDetails")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this describe database response based on the context it is used
func (m *DescribeDatabaseResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDatabaseDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DescribeDatabaseResponse) contextValidateDatabaseDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.DatabaseDetails != nil {

		if swag.IsZero(m.DatabaseDetails) { // not required
			return nil
		}

		if err := m.DatabaseDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databaseDetails")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("databaseDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DescribeDatabaseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DescribeDatabaseResponse) UnmarshalBinary(b []byte) error {
	var res DescribeDatabaseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
