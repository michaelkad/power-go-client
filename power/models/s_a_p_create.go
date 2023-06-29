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

// SAPCreate s a p create
//
// swagger:model SAPCreate
type SAPCreate struct {

	// Custom SAP Deployment Type Information (For Internal Use Only)
	DeploymentType string `json:"deploymentType,omitempty"`

	// Image ID of the sap image to use for the server
	// Required: true
	ImageID *string `json:"imageID"`

	// instances
	Instances *PVMInstanceMultiCreate `json:"instances,omitempty"`

	// Name of the sap pvm-instance
	// Required: true
	Name *string `json:"name"`

	// The pvm instance networks information
	// Required: true
	Networks []*PVMInstanceAddNetwork `json:"networks"`

	// pin policy
	PinPolicy PinPolicy `json:"pinPolicy,omitempty"`

	// The placement group for the server
	PlacementGroup string `json:"placementGroup,omitempty"`

	// SAP Profile ID for the amount of cores and memory
	// Required: true
	ProfileID *string `json:"profileID"`

	// The name of the SSH Key to provide to the server for authenticating
	SSHKeyName string `json:"sshKeyName,omitempty"`

	// The storage affinity data; ignored if storagePool is provided; Only valid when you deploy one of the IBM supplied stock images. Storage type and pool for a custom image (an imported image or an image that is created from a PVMInstance capture) defaults to the storage type and pool the image was created in
	StorageAffinity *StorageAffinity `json:"storageAffinity,omitempty"`

	// Storage Pool for server deployment; if provided then storageAffinity and storageType will be ignored; Only valid when you deploy one of the IBM supplied stock images. Storage type and pool for a custom image (an imported image or an image that is created from a PVMInstance capture) defaults to the storage type and pool the image was created in
	StoragePool string `json:"storagePool,omitempty"`

	// Storage type for server deployment; will be ignored if storagePool or storageAffinity is provided; Only valid when you deploy one of the IBM supplied stock images. Storage type and pool for a custom image (an imported image or an image that is created from a PVMInstance capture) defaults to the storage type and pool the image was created in
	StorageType string `json:"storageType,omitempty"`

	// System type used to host the instance. Only e880, e980, e1080 are supported
	SysType string `json:"sysType,omitempty"`

	// Cloud init user defined data
	UserData string `json:"userData,omitempty"`

	// List of Volume IDs to attach to the pvm-instance on creation
	VolumeIDs []string `json:"volumeIDs"`
}

// Validate validates this s a p create
func (m *SAPCreate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateImageID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePinPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageAffinity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SAPCreate) validateImageID(formats strfmt.Registry) error {

	if err := validate.Required("imageID", "body", m.ImageID); err != nil {
		return err
	}

	return nil
}

func (m *SAPCreate) validateInstances(formats strfmt.Registry) error {
	if swag.IsZero(m.Instances) { // not required
		return nil
	}

	if m.Instances != nil {
		if err := m.Instances.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instances")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instances")
			}
			return err
		}
	}

	return nil
}

func (m *SAPCreate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SAPCreate) validateNetworks(formats strfmt.Registry) error {

	if err := validate.Required("networks", "body", m.Networks); err != nil {
		return err
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SAPCreate) validatePinPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.PinPolicy) { // not required
		return nil
	}

	if err := m.PinPolicy.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pinPolicy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("pinPolicy")
		}
		return err
	}

	return nil
}

func (m *SAPCreate) validateProfileID(formats strfmt.Registry) error {

	if err := validate.Required("profileID", "body", m.ProfileID); err != nil {
		return err
	}

	return nil
}

func (m *SAPCreate) validateStorageAffinity(formats strfmt.Registry) error {
	if swag.IsZero(m.StorageAffinity) { // not required
		return nil
	}

	if m.StorageAffinity != nil {
		if err := m.StorageAffinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageAffinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageAffinity")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this s a p create based on the context it is used
func (m *SAPCreate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePinPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStorageAffinity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SAPCreate) contextValidateInstances(ctx context.Context, formats strfmt.Registry) error {

	if m.Instances != nil {

		if swag.IsZero(m.Instances) { // not required
			return nil
		}

		if err := m.Instances.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instances")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instances")
			}
			return err
		}
	}

	return nil
}

func (m *SAPCreate) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {

			if swag.IsZero(m.Networks[i]) { // not required
				return nil
			}

			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SAPCreate) contextValidatePinPolicy(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.PinPolicy) { // not required
		return nil
	}

	if err := m.PinPolicy.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("pinPolicy")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("pinPolicy")
		}
		return err
	}

	return nil
}

func (m *SAPCreate) contextValidateStorageAffinity(ctx context.Context, formats strfmt.Registry) error {

	if m.StorageAffinity != nil {

		if swag.IsZero(m.StorageAffinity) { // not required
			return nil
		}

		if err := m.StorageAffinity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("storageAffinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("storageAffinity")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SAPCreate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SAPCreate) UnmarshalBinary(b []byte) error {
	var res SAPCreate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
