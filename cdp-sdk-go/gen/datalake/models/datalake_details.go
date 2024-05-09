// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatalakeDetails Details about a datalake
//
// swagger:model DatalakeDetails
type DatalakeDetails struct {

	// The AWS configuration.
	AwsConfiguration *AWSConfiguration `json:"awsConfiguration,omitempty"`

	// The Azure configuration.
	AzureConfiguration *AzureConfiguration `json:"azureConfiguration,omitempty"`

	// The cloud platform.
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// The base location for the cloud storage used for the datalake.
	CloudStorageBaseLocation string `json:"cloudStorageBaseLocation,omitempty"`

	// The Cloudera Manager details.
	ClouderaManager *ClouderaManagerDetails `json:"clouderaManager,omitempty"`

	// The date when the datalake was created.
	// Format: date-time
	CreationDate strfmt.DateTime `json:"creationDate,omitempty"`

	// The CRN of credentials.
	CredentialCrn string `json:"credentialCrn,omitempty"`

	// The CRN of the datalake.
	// Required: true
	Crn *string `json:"crn"`

	// The name of the datalake.
	// Required: true
	DatalakeName *string `json:"datalakeName"`

	// Whether Ranger RAZ is enabled for the datalake.
	EnableRangerRaz bool `json:"enableRangerRaz,omitempty"`

	// The exposed service api endpoints.
	Endpoints *Endpoints `json:"endpoints,omitempty"`

	// The CRN of the environment.
	EnvironmentCrn string `json:"environmentCrn,omitempty"`

	// The GCP configuration.
	GcpConfiguration *GCPConfiguration `json:"gcpConfiguration,omitempty"`

	// The instance details.
	InstanceGroups []*InstanceGroup `json:"instanceGroups"`

	// The product versions.
	ProductVersions []*ProductVersion `json:"productVersions"`

	// The region of the datalake.
	Region string `json:"region,omitempty"`

	// The shape of the datalake (either LIGHT_DUTY or MEDIUM_DUTY_HA).
	Shape DatalakeScaleType `json:"shape,omitempty"`

	// The status of the datalake.
	Status string `json:"status,omitempty"`

	// The reason for the status of the datalake.
	StatusReason string `json:"statusReason,omitempty"`

	// Datalake tags object containing the tag values defined for the datalake.
	Tags []*DatalakeResourceTag `json:"tags"`
}

// Validate validates this datalake details
func (m *DatalakeDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAwsConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzureConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClouderaManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatalakeName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcpConfiguration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProductVersions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShape(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatalakeDetails) validateAwsConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.AwsConfiguration) { // not required
		return nil
	}

	if m.AwsConfiguration != nil {
		if err := m.AwsConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *DatalakeDetails) validateAzureConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.AzureConfiguration) { // not required
		return nil
	}

	if m.AzureConfiguration != nil {
		if err := m.AzureConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *DatalakeDetails) validateClouderaManager(formats strfmt.Registry) error {
	if swag.IsZero(m.ClouderaManager) { // not required
		return nil
	}

	if m.ClouderaManager != nil {
		if err := m.ClouderaManager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clouderaManager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clouderaManager")
			}
			return err
		}
	}

	return nil
}

func (m *DatalakeDetails) validateCreationDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DatalakeDetails) validateCrn(formats strfmt.Registry) error {

	if err := validate.Required("crn", "body", m.Crn); err != nil {
		return err
	}

	return nil
}

func (m *DatalakeDetails) validateDatalakeName(formats strfmt.Registry) error {

	if err := validate.Required("datalakeName", "body", m.DatalakeName); err != nil {
		return err
	}

	return nil
}

func (m *DatalakeDetails) validateEndpoints(formats strfmt.Registry) error {
	if swag.IsZero(m.Endpoints) { // not required
		return nil
	}

	if m.Endpoints != nil {
		if err := m.Endpoints.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endpoints")
			}
			return err
		}
	}

	return nil
}

func (m *DatalakeDetails) validateGcpConfiguration(formats strfmt.Registry) error {
	if swag.IsZero(m.GcpConfiguration) { // not required
		return nil
	}

	if m.GcpConfiguration != nil {
		if err := m.GcpConfiguration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcpConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcpConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *DatalakeDetails) validateInstanceGroups(formats strfmt.Registry) error {
	if swag.IsZero(m.InstanceGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceGroups); i++ {
		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {
			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatalakeDetails) validateProductVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.ProductVersions) { // not required
		return nil
	}

	for i := 0; i < len(m.ProductVersions); i++ {
		if swag.IsZero(m.ProductVersions[i]) { // not required
			continue
		}

		if m.ProductVersions[i] != nil {
			if err := m.ProductVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productVersions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("productVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatalakeDetails) validateShape(formats strfmt.Registry) error {
	if swag.IsZero(m.Shape) { // not required
		return nil
	}

	if err := m.Shape.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shape")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shape")
		}
		return err
	}

	return nil
}

func (m *DatalakeDetails) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this datalake details based on the context it is used
func (m *DatalakeDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAwsConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAzureConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClouderaManager(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEndpoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGcpConfiguration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstanceGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProductVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShape(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatalakeDetails) contextValidateAwsConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.AwsConfiguration != nil {

		if swag.IsZero(m.AwsConfiguration) { // not required
			return nil
		}

		if err := m.AwsConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("awsConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("awsConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *DatalakeDetails) contextValidateAzureConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.AzureConfiguration != nil {

		if swag.IsZero(m.AzureConfiguration) { // not required
			return nil
		}

		if err := m.AzureConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azureConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("azureConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *DatalakeDetails) contextValidateClouderaManager(ctx context.Context, formats strfmt.Registry) error {

	if m.ClouderaManager != nil {

		if swag.IsZero(m.ClouderaManager) { // not required
			return nil
		}

		if err := m.ClouderaManager.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clouderaManager")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clouderaManager")
			}
			return err
		}
	}

	return nil
}

func (m *DatalakeDetails) contextValidateEndpoints(ctx context.Context, formats strfmt.Registry) error {

	if m.Endpoints != nil {

		if swag.IsZero(m.Endpoints) { // not required
			return nil
		}

		if err := m.Endpoints.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoints")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("endpoints")
			}
			return err
		}
	}

	return nil
}

func (m *DatalakeDetails) contextValidateGcpConfiguration(ctx context.Context, formats strfmt.Registry) error {

	if m.GcpConfiguration != nil {

		if swag.IsZero(m.GcpConfiguration) { // not required
			return nil
		}

		if err := m.GcpConfiguration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcpConfiguration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gcpConfiguration")
			}
			return err
		}
	}

	return nil
}

func (m *DatalakeDetails) contextValidateInstanceGroups(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.InstanceGroups); i++ {

		if m.InstanceGroups[i] != nil {

			if swag.IsZero(m.InstanceGroups[i]) { // not required
				return nil
			}

			if err := m.InstanceGroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatalakeDetails) contextValidateProductVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ProductVersions); i++ {

		if m.ProductVersions[i] != nil {

			if swag.IsZero(m.ProductVersions[i]) { // not required
				return nil
			}

			if err := m.ProductVersions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("productVersions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("productVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DatalakeDetails) contextValidateShape(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Shape) { // not required
		return nil
	}

	if err := m.Shape.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("shape")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("shape")
		}
		return err
	}

	return nil
}

func (m *DatalakeDetails) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {

			if swag.IsZero(m.Tags[i]) { // not required
				return nil
			}

			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatalakeDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatalakeDetails) UnmarshalBinary(b []byte) error {
	var res DatalakeDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
