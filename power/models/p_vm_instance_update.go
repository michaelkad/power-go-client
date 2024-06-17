// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PVMInstanceUpdate p VM instance update
//
// swagger:model PVMInstanceUpdate
type PVMInstanceUpdate struct {

	// The VTL license repository capacity TB value
	LicenseRepositoryCapacity int64 `json:"licenseRepositoryCapacity,omitempty"`

	// Amount of memory allocated (in GB)
	Memory float64 `json:"memory,omitempty"`

	// (deprecated - replaced by pinPolicy) Indicates if the server is allowed to migrate between hosts
	Migratable *bool `json:"migratable,omitempty"`

	// pin policy
	PinPolicy PinPolicy `json:"pinPolicy,omitempty"`

	// Processor type (dedicated, shared, capped)
	// Enum: ["dedicated","shared","capped"]
	ProcType string `json:"procType,omitempty"`

	// Number of processors allocated
	Processors float64 `json:"processors,omitempty"`

	// If an SAP pvm-instance, the SAP profile ID to switch to (only while shutdown)
	SapProfileID string `json:"sapProfileID,omitempty"`

	// Name of the server to create
	ServerName string `json:"serverName,omitempty"`

	// The pvm instance Software Licenses
	SoftwareLicenses *SoftwareLicenses `json:"softwareLicenses,omitempty"`

	// Indicates if all volumes attached to the server must reside in the same storage pool; If set to false then volumes from any storage type and pool can be attached to the PVMInstance; Impacts PVMInstance snapshot, capture, and clone, for capture and clone - only data volumes that are of the same storage type and in the same storage pool of the PVMInstance's boot volume can be included; for snapshot - all data volumes to be included in the snapshot must reside in the same storage type and pool. Once set to false, cannot be set back to true unless all volumes attached reside in the same storage type and pool.
	StoragePoolAffinity *bool `json:"storagePoolAffinity,omitempty"`

	// The pvm instance virtual CPU information
	VirtualCores *VirtualCores `json:"virtualCores,omitempty"`
}

// Validate validates this p VM instance update
func (m *PVMInstanceUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePinPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProcType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftwareLicenses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualCores(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceUpdate) validatePinPolicy(formats strfmt.Registry) error {
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

var pVmInstanceUpdateTypeProcTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dedicated","shared","capped"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		pVmInstanceUpdateTypeProcTypePropEnum = append(pVmInstanceUpdateTypeProcTypePropEnum, v)
	}
}

const (

	// PVMInstanceUpdateProcTypeDedicated captures enum value "dedicated"
	PVMInstanceUpdateProcTypeDedicated string = "dedicated"

	// PVMInstanceUpdateProcTypeShared captures enum value "shared"
	PVMInstanceUpdateProcTypeShared string = "shared"

	// PVMInstanceUpdateProcTypeCapped captures enum value "capped"
	PVMInstanceUpdateProcTypeCapped string = "capped"
)

// prop value enum
func (m *PVMInstanceUpdate) validateProcTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, pVmInstanceUpdateTypeProcTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PVMInstanceUpdate) validateProcType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProcType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProcTypeEnum("procType", "body", m.ProcType); err != nil {
		return err
	}

	return nil
}

func (m *PVMInstanceUpdate) validateSoftwareLicenses(formats strfmt.Registry) error {
	if swag.IsZero(m.SoftwareLicenses) { // not required
		return nil
	}

	if m.SoftwareLicenses != nil {
		if err := m.SoftwareLicenses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareLicenses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softwareLicenses")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceUpdate) validateVirtualCores(formats strfmt.Registry) error {
	if swag.IsZero(m.VirtualCores) { // not required
		return nil
	}

	if m.VirtualCores != nil {
		if err := m.VirtualCores.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualCores")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualCores")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this p VM instance update based on the context it is used
func (m *PVMInstanceUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePinPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSoftwareLicenses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualCores(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstanceUpdate) contextValidatePinPolicy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PVMInstanceUpdate) contextValidateSoftwareLicenses(ctx context.Context, formats strfmt.Registry) error {

	if m.SoftwareLicenses != nil {

		if swag.IsZero(m.SoftwareLicenses) { // not required
			return nil
		}

		if err := m.SoftwareLicenses.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("softwareLicenses")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("softwareLicenses")
			}
			return err
		}
	}

	return nil
}

func (m *PVMInstanceUpdate) contextValidateVirtualCores(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualCores != nil {

		if swag.IsZero(m.VirtualCores) { // not required
			return nil
		}

		if err := m.VirtualCores.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtualCores")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtualCores")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstanceUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstanceUpdate) UnmarshalBinary(b []byte) error {
	var res PVMInstanceUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
