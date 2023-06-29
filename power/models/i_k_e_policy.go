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

// IKEPolicy IKE Policy object
//
// swagger:model IKEPolicy
type IKEPolicy struct {

	// authentication
	// Required: true
	Authentication *IKEPolicyAuthentication `json:"authentication"`

	// DH group of the IKE Policy
	// Example: 2
	// Required: true
	// Enum: [1 2 5 14 19 20 24]
	DhGroup *int64 `json:"dhGroup"`

	// encryption of the IKE Policy
	// Example: aes-256-cbc
	// Required: true
	// Enum: [aes-256-cbc aes-192-cbc aes-128-cbc aes-256-gcm aes-128-gcm 3des-cbc]
	Encryption *string `json:"encryption"`

	// unique identifier of the IKE Policy object
	// Example: 7edc8988-eb18-4b5c-a594-0d73d8254463
	// Required: true
	ID *string `json:"id"`

	// key lifetime
	// Required: true
	KeyLifetime *KeyLifetime `json:"keyLifetime"`

	// name of the IKE Policy
	// Example: ikePolicy1
	// Required: true
	Name *string `json:"name"`

	// version of the IKE Policy
	// Example: 2
	// Required: true
	// Enum: [1 2]
	Version *int64 `json:"version"`
}

// Validate validates this i k e policy
func (m *IKEPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *IKEPolicy) validateAuthentication(formats strfmt.Registry) error {

	if err := validate.Required("authentication", "body", m.Authentication); err != nil {
		return err
	}

	if err := validate.Required("authentication", "body", m.Authentication); err != nil {
		return err
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

var iKEPolicyTypeDhGroupPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,5,14,19,20,24]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iKEPolicyTypeDhGroupPropEnum = append(iKEPolicyTypeDhGroupPropEnum, v)
	}
}

// prop value enum
func (m *IKEPolicy) validateDhGroupEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, iKEPolicyTypeDhGroupPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IKEPolicy) validateDhGroup(formats strfmt.Registry) error {

	if err := validate.Required("dhGroup", "body", m.DhGroup); err != nil {
		return err
	}

	// value enum
	if err := m.validateDhGroupEnum("dhGroup", "body", *m.DhGroup); err != nil {
		return err
	}

	return nil
}

var iKEPolicyTypeEncryptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aes-256-cbc","aes-192-cbc","aes-128-cbc","aes-256-gcm","aes-128-gcm","3des-cbc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iKEPolicyTypeEncryptionPropEnum = append(iKEPolicyTypeEncryptionPropEnum, v)
	}
}

const (

	// IKEPolicyEncryptionAesDash256DashCbc captures enum value "aes-256-cbc"
	IKEPolicyEncryptionAesDash256DashCbc string = "aes-256-cbc"

	// IKEPolicyEncryptionAesDash192DashCbc captures enum value "aes-192-cbc"
	IKEPolicyEncryptionAesDash192DashCbc string = "aes-192-cbc"

	// IKEPolicyEncryptionAesDash128DashCbc captures enum value "aes-128-cbc"
	IKEPolicyEncryptionAesDash128DashCbc string = "aes-128-cbc"

	// IKEPolicyEncryptionAesDash256DashGcm captures enum value "aes-256-gcm"
	IKEPolicyEncryptionAesDash256DashGcm string = "aes-256-gcm"

	// IKEPolicyEncryptionAesDash128DashGcm captures enum value "aes-128-gcm"
	IKEPolicyEncryptionAesDash128DashGcm string = "aes-128-gcm"

	// IKEPolicyEncryptionNr3desDashCbc captures enum value "3des-cbc"
	IKEPolicyEncryptionNr3desDashCbc string = "3des-cbc"
)

// prop value enum
func (m *IKEPolicy) validateEncryptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, iKEPolicyTypeEncryptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IKEPolicy) validateEncryption(formats strfmt.Registry) error {

	if err := validate.Required("encryption", "body", m.Encryption); err != nil {
		return err
	}

	// value enum
	if err := m.validateEncryptionEnum("encryption", "body", *m.Encryption); err != nil {
		return err
	}

	return nil
}

func (m *IKEPolicy) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *IKEPolicy) validateKeyLifetime(formats strfmt.Registry) error {

	if err := validate.Required("keyLifetime", "body", m.KeyLifetime); err != nil {
		return err
	}

	if err := validate.Required("keyLifetime", "body", m.KeyLifetime); err != nil {
		return err
	}

	if m.KeyLifetime != nil {
		if err := m.KeyLifetime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyLifetime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyLifetime")
			}
			return err
		}
	}

	return nil
}

func (m *IKEPolicy) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var iKEPolicyTypeVersionPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iKEPolicyTypeVersionPropEnum = append(iKEPolicyTypeVersionPropEnum, v)
	}
}

// prop value enum
func (m *IKEPolicy) validateVersionEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, iKEPolicyTypeVersionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IKEPolicy) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	// value enum
	if err := m.validateVersionEnum("version", "body", *m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this i k e policy based on the context it is used
func (m *IKEPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyLifetime(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IKEPolicy) contextValidateAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if m.Authentication != nil {

		if err := m.Authentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *IKEPolicy) contextValidateKeyLifetime(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyLifetime != nil {

		if err := m.KeyLifetime.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyLifetime")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyLifetime")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IKEPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IKEPolicy) UnmarshalBinary(b []byte) error {
	var res IKEPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
