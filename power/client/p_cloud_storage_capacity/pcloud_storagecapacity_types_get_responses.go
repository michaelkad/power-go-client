// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_storage_capacity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudStoragecapacityTypesGetReader is a Reader for the PcloudStoragecapacityTypesGet structure.
type PcloudStoragecapacityTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudStoragecapacityTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudStoragecapacityTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPcloudStoragecapacityTypesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudStoragecapacityTypesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudStoragecapacityTypesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudStoragecapacityTypesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}] pcloud.storagecapacity.types.get", response, response.Code())
	}
}

// NewPcloudStoragecapacityTypesGetOK creates a PcloudStoragecapacityTypesGetOK with default headers values
func NewPcloudStoragecapacityTypesGetOK() *PcloudStoragecapacityTypesGetOK {
	return &PcloudStoragecapacityTypesGetOK{}
}

/*
PcloudStoragecapacityTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudStoragecapacityTypesGetOK struct {
	Payload *models.StorageTypeCapacity
}

// IsSuccess returns true when this pcloud storagecapacity types get o k response has a 2xx status code
func (o *PcloudStoragecapacityTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud storagecapacity types get o k response has a 3xx status code
func (o *PcloudStoragecapacityTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity types get o k response has a 4xx status code
func (o *PcloudStoragecapacityTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud storagecapacity types get o k response has a 5xx status code
func (o *PcloudStoragecapacityTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity types get o k response a status code equal to that given
func (o *PcloudStoragecapacityTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud storagecapacity types get o k response
func (o *PcloudStoragecapacityTypesGetOK) Code() int {
	return 200
}

func (o *PcloudStoragecapacityTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetOK) GetPayload() *models.StorageTypeCapacity {
	return o.Payload
}

func (o *PcloudStoragecapacityTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StorageTypeCapacity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityTypesGetUnauthorized creates a PcloudStoragecapacityTypesGetUnauthorized with default headers values
func NewPcloudStoragecapacityTypesGetUnauthorized() *PcloudStoragecapacityTypesGetUnauthorized {
	return &PcloudStoragecapacityTypesGetUnauthorized{}
}

/*
PcloudStoragecapacityTypesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudStoragecapacityTypesGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity types get unauthorized response has a 2xx status code
func (o *PcloudStoragecapacityTypesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity types get unauthorized response has a 3xx status code
func (o *PcloudStoragecapacityTypesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity types get unauthorized response has a 4xx status code
func (o *PcloudStoragecapacityTypesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud storagecapacity types get unauthorized response has a 5xx status code
func (o *PcloudStoragecapacityTypesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity types get unauthorized response a status code equal to that given
func (o *PcloudStoragecapacityTypesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud storagecapacity types get unauthorized response
func (o *PcloudStoragecapacityTypesGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudStoragecapacityTypesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityTypesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityTypesGetForbidden creates a PcloudStoragecapacityTypesGetForbidden with default headers values
func NewPcloudStoragecapacityTypesGetForbidden() *PcloudStoragecapacityTypesGetForbidden {
	return &PcloudStoragecapacityTypesGetForbidden{}
}

/*
PcloudStoragecapacityTypesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudStoragecapacityTypesGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity types get forbidden response has a 2xx status code
func (o *PcloudStoragecapacityTypesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity types get forbidden response has a 3xx status code
func (o *PcloudStoragecapacityTypesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity types get forbidden response has a 4xx status code
func (o *PcloudStoragecapacityTypesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud storagecapacity types get forbidden response has a 5xx status code
func (o *PcloudStoragecapacityTypesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity types get forbidden response a status code equal to that given
func (o *PcloudStoragecapacityTypesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud storagecapacity types get forbidden response
func (o *PcloudStoragecapacityTypesGetForbidden) Code() int {
	return 403
}

func (o *PcloudStoragecapacityTypesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityTypesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityTypesGetNotFound creates a PcloudStoragecapacityTypesGetNotFound with default headers values
func NewPcloudStoragecapacityTypesGetNotFound() *PcloudStoragecapacityTypesGetNotFound {
	return &PcloudStoragecapacityTypesGetNotFound{}
}

/*
PcloudStoragecapacityTypesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudStoragecapacityTypesGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity types get not found response has a 2xx status code
func (o *PcloudStoragecapacityTypesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity types get not found response has a 3xx status code
func (o *PcloudStoragecapacityTypesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity types get not found response has a 4xx status code
func (o *PcloudStoragecapacityTypesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud storagecapacity types get not found response has a 5xx status code
func (o *PcloudStoragecapacityTypesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud storagecapacity types get not found response a status code equal to that given
func (o *PcloudStoragecapacityTypesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud storagecapacity types get not found response
func (o *PcloudStoragecapacityTypesGetNotFound) Code() int {
	return 404
}

func (o *PcloudStoragecapacityTypesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityTypesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudStoragecapacityTypesGetInternalServerError creates a PcloudStoragecapacityTypesGetInternalServerError with default headers values
func NewPcloudStoragecapacityTypesGetInternalServerError() *PcloudStoragecapacityTypesGetInternalServerError {
	return &PcloudStoragecapacityTypesGetInternalServerError{}
}

/*
PcloudStoragecapacityTypesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudStoragecapacityTypesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud storagecapacity types get internal server error response has a 2xx status code
func (o *PcloudStoragecapacityTypesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud storagecapacity types get internal server error response has a 3xx status code
func (o *PcloudStoragecapacityTypesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud storagecapacity types get internal server error response has a 4xx status code
func (o *PcloudStoragecapacityTypesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud storagecapacity types get internal server error response has a 5xx status code
func (o *PcloudStoragecapacityTypesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud storagecapacity types get internal server error response a status code equal to that given
func (o *PcloudStoragecapacityTypesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud storagecapacity types get internal server error response
func (o *PcloudStoragecapacityTypesGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudStoragecapacityTypesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/storage-capacity/storage-types/{storage_type_name}][%d] pcloudStoragecapacityTypesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudStoragecapacityTypesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudStoragecapacityTypesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
