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

// VolumeGroupDetails volume group details
//
// swagger:model VolumeGroupDetails
type VolumeGroupDetails struct {

	// Indicates whether the volume group is for auxiliary volumes or master volumes
	Auxiliary bool `json:"auxiliary,omitempty"`

	// The name of volume group at storage host level
	ConsistencyGroupName string `json:"consistencyGroupName,omitempty"`

	// The ID of the volume group
	// Required: true
	ID *string `json:"id"`

	// The name of the volume group
	// Required: true
	Name *string `json:"name"`

	// Indicates the replication site of the volume group
	ReplicationSites []string `json:"replicationSites"`

	// Replication status of volume group
	ReplicationStatus string `json:"replicationStatus,omitempty"`

	// Status of the volume group
	Status string `json:"status,omitempty"`

	// Status details of the volume group
	StatusDescription *StatusDescription `json:"statusDescription,omitempty"`

	// Indicates the storage pool of the volume group
	StoragePool string `json:"storagePool,omitempty"`

	// List of volume IDs,member of VolumeGroup
	VolumeIDs []string `json:"volumeIDs"`
}

// Validate validates this volume group details
func (m *VolumeGroupDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeGroupDetails) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *VolumeGroupDetails) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *VolumeGroupDetails) validateStatusDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusDescription) { // not required
		return nil
	}

	if m.StatusDescription != nil {
		if err := m.StatusDescription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusDescription")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusDescription")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this volume group details based on the context it is used
func (m *VolumeGroupDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStatusDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolumeGroupDetails) contextValidateStatusDescription(ctx context.Context, formats strfmt.Registry) error {

	if m.StatusDescription != nil {

		if swag.IsZero(m.StatusDescription) { // not required
			return nil
		}

		if err := m.StatusDescription.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statusDescription")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statusDescription")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolumeGroupDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeGroupDetails) UnmarshalBinary(b []byte) error {
	var res VolumeGroupDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
