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

// NetworkSecurityGroupMember network security group member
//
// swagger:model NetworkSecurityGroupMember
type NetworkSecurityGroupMember struct {

	// The ID of the member in a Network Security Group
	// Required: true
	ID *string `json:"id"`

	// The mac address of a Network Interface included if the type is network-interface
	MacAddress string `json:"macAddress,omitempty"`

	// If ipv4-address type, then IPv4 address or if network-interface type, then network interface ID
	// Required: true
	Target *string `json:"target"`

	// The type of member
	// Required: true
	// Enum: ["ipv4-address","network-interface"]
	Type *string `json:"type"`
}

// Validate validates this network security group member
func (m *NetworkSecurityGroupMember) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkSecurityGroupMember) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *NetworkSecurityGroupMember) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

var networkSecurityGroupMemberTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ipv4-address","network-interface"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		networkSecurityGroupMemberTypeTypePropEnum = append(networkSecurityGroupMemberTypeTypePropEnum, v)
	}
}

const (

	// NetworkSecurityGroupMemberTypeIPV4DashAddress captures enum value "ipv4-address"
	NetworkSecurityGroupMemberTypeIPV4DashAddress string = "ipv4-address"

	// NetworkSecurityGroupMemberTypeNetworkDashInterface captures enum value "network-interface"
	NetworkSecurityGroupMemberTypeNetworkDashInterface string = "network-interface"
)

// prop value enum
func (m *NetworkSecurityGroupMember) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, networkSecurityGroupMemberTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NetworkSecurityGroupMember) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this network security group member based on context it is used
func (m *NetworkSecurityGroupMember) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkSecurityGroupMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkSecurityGroupMember) UnmarshalBinary(b []byte) error {
	var res NetworkSecurityGroupMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}