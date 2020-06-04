// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DatabaseAvailabilityType Represents the database availability type.
//
// swagger:model DatabaseAvailabilityType
type DatabaseAvailabilityType string

const (

	// DatabaseAvailabilityTypeHA captures enum value "HA"
	DatabaseAvailabilityTypeHA DatabaseAvailabilityType = "HA"

	// DatabaseAvailabilityTypeNONHA captures enum value "NON_HA"
	DatabaseAvailabilityTypeNONHA DatabaseAvailabilityType = "NON_HA"

	// DatabaseAvailabilityTypeNONE captures enum value "NONE"
	DatabaseAvailabilityTypeNONE DatabaseAvailabilityType = "NONE"
)

// for schema
var databaseAvailabilityTypeEnum []interface{}

func init() {
	var res []DatabaseAvailabilityType
	if err := json.Unmarshal([]byte(`["HA","NON_HA","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		databaseAvailabilityTypeEnum = append(databaseAvailabilityTypeEnum, v)
	}
}

func (m DatabaseAvailabilityType) validateDatabaseAvailabilityTypeEnum(path, location string, value DatabaseAvailabilityType) error {
	if err := validate.Enum(path, location, value, databaseAvailabilityTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this database availability type
func (m DatabaseAvailabilityType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDatabaseAvailabilityTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
