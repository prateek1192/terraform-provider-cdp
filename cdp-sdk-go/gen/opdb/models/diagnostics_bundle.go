// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DiagnosticsBundle Describe a diagnostics bundle
//
// swagger:model DiagnosticsBundle
type DiagnosticsBundle struct {

	// The object on which diagnostics was collected
	Crn string `json:"crn,omitempty"`

	// The end time, if the command is finished.
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// Identifier for each bundle collection
	ID string `json:"id,omitempty"`

	// If this is a download, a link to the download location of the bundle
	Result string `json:"result,omitempty"`

	// Start time
	// Format: date-time
	StartTime strfmt.DateTime `json:"startTime,omitempty"`

	// A status of the diagnostics collection process
	Status string `json:"status,omitempty"`
}

// Validate validates this diagnostics bundle
func (m *DiagnosticsBundle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DiagnosticsBundle) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DiagnosticsBundle) validateStartTime(formats strfmt.Registry) error {
	if swag.IsZero(m.StartTime) { // not required
		return nil
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this diagnostics bundle based on context it is used
func (m *DiagnosticsBundle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DiagnosticsBundle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DiagnosticsBundle) UnmarshalBinary(b []byte) error {
	var res DiagnosticsBundle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
