// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkspaceSummary A Cloudera Machine Learning workspace which includes the deployed configuration details.
//
// swagger:model WorkspaceSummary
type WorkspaceSummary struct {

	// The cloud platform of the environment that was used to create this workspace.
	// Required: true
	CloudPlatform *string `json:"cloudPlatform"`

	// Creation date of workspace.
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// The CRN of the creator of the workspace.
	// Required: true
	CreatorCrn *string `json:"creatorCrn"`

	// The CRN of the workspace.
	// Required: true
	Crn *string `json:"crn"`

	// CRN of the environment.
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// The name of the workspace's environment.
	// Required: true
	EnvironmentName *string `json:"environmentName"`

	// Failure message from the most recent failure that has occurred during workspace provisioning.
	FailureMessage string `json:"failureMessage,omitempty"`

	// A filesystem ID referencing the filesystem that was created on the cloud provider environment that this workspace uses.
	// Required: true
	FilesystemID *string `json:"filesystemID"`

	// The health info information of the workspace.
	HealthInfoLists []*HealthInfo `json:"healthInfoLists"`

	// Indicates if HTTPs communication was enabled on this workspace when provisioned.
	// Required: true
	HTTPSEnabled *bool `json:"httpsEnabled"`

	// The name of the workspace.
	// Required: true
	InstanceName *string `json:"instanceName"`

	// The workspace's current status.
	// Required: true
	InstanceStatus *string `json:"instanceStatus"`

	// URL of the workspace's user interface.
	// Required: true
	InstanceURL *string `json:"instanceUrl"`

	// The Kubernetes cluster name.
	// Required: true
	K8sClusterName *string `json:"k8sClusterName"`

	// The whitelist of ips for loadBalancer.
	LoadBalancerIPWhitelists []string `json:"loadBalancerIPWhitelists"`

	// If usage monitoring is enabled or not on this workspace.
	// Required: true
	MonitoringEnabled *bool `json:"monitoringEnabled"`

	// The version of Cloudera Machine Learning that was installed on the workspace.
	// Required: true
	Version *string `json:"version"`
}

// Validate validates this workspace summary
func (m *WorkspaceSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatorCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilesystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthInfoLists(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHTTPSEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateK8sClusterName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMonitoringEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkspaceSummary) validateCloudPlatform(formats strfmt.Registry) error {

	if err := validate.Required("cloudPlatform", "body", m.CloudPlatform); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateCreationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateCreatorCrn(formats strfmt.Registry) error {

	if err := validate.Required("creatorCrn", "body", m.CreatorCrn); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateEnvironmentName(formats strfmt.Registry) error {

	if err := validate.Required("environmentName", "body", m.EnvironmentName); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateFilesystemID(formats strfmt.Registry) error {

	if err := validate.Required("filesystemID", "body", m.FilesystemID); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateHealthInfoLists(formats strfmt.Registry) error {

	if swag.IsZero(m.HealthInfoLists) { // not required
		return nil
	}

	for i := 0; i < len(m.HealthInfoLists); i++ {
		if swag.IsZero(m.HealthInfoLists[i]) { // not required
			continue
		}

		if m.HealthInfoLists[i] != nil {
			if err := m.HealthInfoLists[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("healthInfoLists" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkspaceSummary) validateHTTPSEnabled(formats strfmt.Registry) error {

	if err := validate.Required("httpsEnabled", "body", m.HTTPSEnabled); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateInstanceName(formats strfmt.Registry) error {

	if err := validate.Required("instanceName", "body", m.InstanceName); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateInstanceStatus(formats strfmt.Registry) error {

	if err := validate.Required("instanceStatus", "body", m.InstanceStatus); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateInstanceURL(formats strfmt.Registry) error {

	if err := validate.Required("instanceUrl", "body", m.InstanceURL); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateK8sClusterName(formats strfmt.Registry) error {

	if err := validate.Required("k8sClusterName", "body", m.K8sClusterName); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateMonitoringEnabled(formats strfmt.Registry) error {

	if err := validate.Required("monitoringEnabled", "body", m.MonitoringEnabled); err != nil {
		return err
	}

	return nil
}

func (m *WorkspaceSummary) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkspaceSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspaceSummary) UnmarshalBinary(b []byte) error {
	var res WorkspaceSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
