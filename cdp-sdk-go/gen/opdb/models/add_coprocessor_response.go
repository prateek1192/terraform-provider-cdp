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

// AddCoprocessorResponse The response for adding a coprocessor of a table in a database.
//
// swagger:model AddCoprocessorResponse
type AddCoprocessorResponse struct {

	// The command ID of the submitted request.
	CommandID int64 `json:"commandId,omitempty"`

	// When the coprocessor request was submitted.
	CreationTime int64 `json:"creationTime,omitempty"`

	// The name or CRN of the database.
	Database string `json:"database,omitempty"`

	// The name or CRN of the environment.
	Environment string `json:"environment,omitempty"`

	// The current status when submitting the add coprocessor request.
	Status CoprocessorStatusType `json:"status,omitempty"`

	// Reason for the response if any.
	StatusReason string `json:"statusReason,omitempty"`
}

// Validate validates this add coprocessor response
func (m *AddCoprocessorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddCoprocessorResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this add coprocessor response based on the context it is used
func (m *AddCoprocessorResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddCoprocessorResponse) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddCoprocessorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddCoprocessorResponse) UnmarshalBinary(b []byte) error {
	var res AddCoprocessorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
