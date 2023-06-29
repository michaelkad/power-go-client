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

// AuxiliaryVolumesForOnboarding auxiliary volumes for onboarding
//
// swagger:model AuxiliaryVolumesForOnboarding
type AuxiliaryVolumesForOnboarding struct {

	// auxiliary volumes
	// Required: true
	AuxiliaryVolumes []*AuxiliaryVolumeForOnboarding `json:"auxiliaryVolumes"`

	// CRN of source ServiceBroker instance from where auxiliary volumes need to be onboarded
	// Required: true
	SourceCRN *string `json:"sourceCRN"`
}

// Validate validates this auxiliary volumes for onboarding
func (m *AuxiliaryVolumesForOnboarding) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuxiliaryVolumes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceCRN(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuxiliaryVolumesForOnboarding) validateAuxiliaryVolumes(formats strfmt.Registry) error {

	if err := validate.Required("auxiliaryVolumes", "body", m.AuxiliaryVolumes); err != nil {
		return err
	}

	for i := 0; i < len(m.AuxiliaryVolumes); i++ {
		if swag.IsZero(m.AuxiliaryVolumes[i]) { // not required
			continue
		}

		if m.AuxiliaryVolumes[i] != nil {
			if err := m.AuxiliaryVolumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auxiliaryVolumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auxiliaryVolumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuxiliaryVolumesForOnboarding) validateSourceCRN(formats strfmt.Registry) error {

	if err := validate.Required("sourceCRN", "body", m.SourceCRN); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this auxiliary volumes for onboarding based on the context it is used
func (m *AuxiliaryVolumesForOnboarding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuxiliaryVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuxiliaryVolumesForOnboarding) contextValidateAuxiliaryVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuxiliaryVolumes); i++ {

		if m.AuxiliaryVolumes[i] != nil {

			if swag.IsZero(m.AuxiliaryVolumes[i]) { // not required
				return nil
			}

			if err := m.AuxiliaryVolumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auxiliaryVolumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auxiliaryVolumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuxiliaryVolumesForOnboarding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuxiliaryVolumesForOnboarding) UnmarshalBinary(b []byte) error {
	var res AuxiliaryVolumesForOnboarding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
