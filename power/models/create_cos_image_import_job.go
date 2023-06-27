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

// CreateCosImageImportJob create cos image import job
//
// swagger:model CreateCosImageImportJob
type CreateCosImageImportJob struct {

	// Cloud Object Storage access key; required for buckets with private access
	AccessKey string `json:"accessKey,omitempty"`

	// indicates if the bucket has public or private access public access require no authentication keys private access requires hmac authentication keys (access,secret)
	// Enum: [public private]
	BucketAccess *string `json:"bucketAccess,omitempty"`

	// Cloud Object Storage bucket name; bucket-name[/optional/folder]
	// Required: true
	BucketName *string `json:"bucketName"`

	// Cloud Object Storage image filename
	// Required: true
	ImageFilename *string `json:"imageFilename"`

	// Name for the image that will be loaded into the boot image catalog
	// Required: true
	ImageName *string `json:"imageName"`

	// Image OS Type, required if importing a raw image; raw images can only be imported using the command line interface
	// Enum: [aix ibmi rhel sles]
	OsType string `json:"osType,omitempty"`

	// Cloud Object Storage region
	// Required: true
	Region *string `json:"region"`

	// Cloud Object Storage secret key; required for buckets with private access
	SecretKey string `json:"secretKey,omitempty"`

	// Storage affinity data used for storage pool selection
	StorageAffinity *StorageAffinity `json:"storageAffinity,omitempty"`

	// Storage pool where the image will be loaded, if provided then storageType and storageAffinity will be ignored
	StoragePool string `json:"storagePool,omitempty"`

	// Type of storage; will be ignored if storagePool or storageAffinity is provided. If only using storageType for storage selection then the storage pool with the most available space will be selected
	StorageType string `json:"storageType,omitempty"`
}

// Validate validates this create cos image import job
func (m *CreateCosImageImportJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBucketAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBucketName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageFilename(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegion(formats); err != nil {
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

var createCosImageImportJobTypeBucketAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","private"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createCosImageImportJobTypeBucketAccessPropEnum = append(createCosImageImportJobTypeBucketAccessPropEnum, v)
	}
}

const (

	// CreateCosImageImportJobBucketAccessPublic captures enum value "public"
	CreateCosImageImportJobBucketAccessPublic string = "public"

	// CreateCosImageImportJobBucketAccessPrivate captures enum value "private"
	CreateCosImageImportJobBucketAccessPrivate string = "private"
)

// prop value enum
func (m *CreateCosImageImportJob) validateBucketAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createCosImageImportJobTypeBucketAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateCosImageImportJob) validateBucketAccess(formats strfmt.Registry) error {
	if swag.IsZero(m.BucketAccess) { // not required
		return nil
	}

	// value enum
	if err := m.validateBucketAccessEnum("bucketAccess", "body", *m.BucketAccess); err != nil {
		return err
	}

	return nil
}

func (m *CreateCosImageImportJob) validateBucketName(formats strfmt.Registry) error {

	if err := validate.Required("bucketName", "body", m.BucketName); err != nil {
		return err
	}

	return nil
}

func (m *CreateCosImageImportJob) validateImageFilename(formats strfmt.Registry) error {

	if err := validate.Required("imageFilename", "body", m.ImageFilename); err != nil {
		return err
	}

	return nil
}

func (m *CreateCosImageImportJob) validateImageName(formats strfmt.Registry) error {

	if err := validate.Required("imageName", "body", m.ImageName); err != nil {
		return err
	}

	return nil
}

var createCosImageImportJobTypeOsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["aix","ibmi","rhel","sles"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createCosImageImportJobTypeOsTypePropEnum = append(createCosImageImportJobTypeOsTypePropEnum, v)
	}
}

const (

	// CreateCosImageImportJobOsTypeAix captures enum value "aix"
	CreateCosImageImportJobOsTypeAix string = "aix"

	// CreateCosImageImportJobOsTypeIbmi captures enum value "ibmi"
	CreateCosImageImportJobOsTypeIbmi string = "ibmi"

	// CreateCosImageImportJobOsTypeRhel captures enum value "rhel"
	CreateCosImageImportJobOsTypeRhel string = "rhel"

	// CreateCosImageImportJobOsTypeSles captures enum value "sles"
	CreateCosImageImportJobOsTypeSles string = "sles"
)

// prop value enum
func (m *CreateCosImageImportJob) validateOsTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createCosImageImportJobTypeOsTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateCosImageImportJob) validateOsType(formats strfmt.Registry) error {
	if swag.IsZero(m.OsType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOsTypeEnum("osType", "body", m.OsType); err != nil {
		return err
	}

	return nil
}

func (m *CreateCosImageImportJob) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

func (m *CreateCosImageImportJob) validateStorageAffinity(formats strfmt.Registry) error {
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

// ContextValidate validate this create cos image import job based on the context it is used
func (m *CreateCosImageImportJob) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateStorageAffinity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCosImageImportJob) contextValidateStorageAffinity(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CreateCosImageImportJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCosImageImportJob) UnmarshalBinary(b []byte) error {
	var res CreateCosImageImportJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
