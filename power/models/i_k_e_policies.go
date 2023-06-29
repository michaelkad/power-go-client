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

// IKEPolicies i k e policies
//
// swagger:model IKEPolicies
type IKEPolicies struct {

	// IKE Policies array
	// Required: true
	IkePolicies []*IKEPolicy `json:"ikePolicies"`
}

// Validate validates this i k e policies
func (m *IKEPolicies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIkePolicies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IKEPolicies) validateIkePolicies(formats strfmt.Registry) error {

	if err := validate.Required("ikePolicies", "body", m.IkePolicies); err != nil {
		return err
	}

	for i := 0; i < len(m.IkePolicies); i++ {
		if swag.IsZero(m.IkePolicies[i]) { // not required
			continue
		}

		if m.IkePolicies[i] != nil {
			if err := m.IkePolicies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ikePolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ikePolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this i k e policies based on the context it is used
func (m *IKEPolicies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIkePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IKEPolicies) contextValidateIkePolicies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IkePolicies); i++ {

		if m.IkePolicies[i] != nil {

			if swag.IsZero(m.IkePolicies[i]) { // not required
				return nil
			}

			if err := m.IkePolicies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ikePolicies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ikePolicies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IKEPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IKEPolicies) UnmarshalBinary(b []byte) error {
	var res IKEPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
