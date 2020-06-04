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

// ClusterDefinitionSummary Information about a cluster definition.
//
// swagger:model ClusterDefinitionSummary
type ClusterDefinitionSummary struct {

	// The cloud platform.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// The name of the cluster definition.
	// Required: true
	ClusterDefinitionName *string `json:"clusterDefinitionName"`

	// The CRN of the cluster definition.
	// Required: true
	Crn *string `json:"crn"`

	// The description of the cluster definition.
	Description string `json:"description,omitempty"`

	// The CRN of the environment.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// The node count of the cluster.
	NodeCount int32 `json:"nodeCount,omitempty"`

	// The product version.
	ProductVersion string `json:"productVersion,omitempty"`

	// The type of cluster definition.
	Type string `json:"type,omitempty"`
}

// Validate validates this cluster definition summary
func (m *ClusterDefinitionSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterDefinitionName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDefinitionSummary) validateClusterDefinitionName(formats strfmt.Registry) error {

	if err := validate.Required("clusterDefinitionName", "body", m.ClusterDefinitionName); err != nil {
		return err
	}

	return nil
}

func (m *ClusterDefinitionSummary) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterDefinitionSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDefinitionSummary) UnmarshalBinary(b []byte) error {
	var res ClusterDefinitionSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
