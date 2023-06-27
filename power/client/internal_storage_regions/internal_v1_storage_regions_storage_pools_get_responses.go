// Code generated by go-swagger; DO NOT EDIT.

package internal_storage_regions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// InternalV1StorageRegionsStoragePoolsGetReader is a Reader for the InternalV1StorageRegionsStoragePoolsGet structure.
type InternalV1StorageRegionsStoragePoolsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1StorageRegionsStoragePoolsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInternalV1StorageRegionsStoragePoolsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewInternalV1StorageRegionsStoragePoolsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInternalV1StorageRegionsStoragePoolsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInternalV1StorageRegionsStoragePoolsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1StorageRegionsStoragePoolsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}] internal.v1.storage.regions.storage-pools.get", response, response.Code())
	}
}

// NewInternalV1StorageRegionsStoragePoolsGetOK creates a InternalV1StorageRegionsStoragePoolsGetOK with default headers values
func NewInternalV1StorageRegionsStoragePoolsGetOK() *InternalV1StorageRegionsStoragePoolsGetOK {
	return &InternalV1StorageRegionsStoragePoolsGetOK{}
}

/*
InternalV1StorageRegionsStoragePoolsGetOK describes a response with status code 200, with default header values.

OK
*/
type InternalV1StorageRegionsStoragePoolsGetOK struct {
	Payload models.StoragePools
}

// IsSuccess returns true when this internal v1 storage regions storage pools get o k response has a 2xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 storage regions storage pools get o k response has a 3xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions storage pools get o k response has a 4xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 storage regions storage pools get o k response has a 5xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions storage pools get o k response a status code equal to that given
func (o *InternalV1StorageRegionsStoragePoolsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the internal v1 storage regions storage pools get o k response
func (o *InternalV1StorageRegionsStoragePoolsGetOK) Code() int {
	return 200
}

func (o *InternalV1StorageRegionsStoragePoolsGetOK) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}][%d] internalV1StorageRegionsStoragePoolsGetOK  %+v", 200, o.Payload)
}

func (o *InternalV1StorageRegionsStoragePoolsGetOK) String() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}][%d] internalV1StorageRegionsStoragePoolsGetOK  %+v", 200, o.Payload)
}

func (o *InternalV1StorageRegionsStoragePoolsGetOK) GetPayload() models.StoragePools {
	return o.Payload
}

func (o *InternalV1StorageRegionsStoragePoolsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsStoragePoolsGetUnauthorized creates a InternalV1StorageRegionsStoragePoolsGetUnauthorized with default headers values
func NewInternalV1StorageRegionsStoragePoolsGetUnauthorized() *InternalV1StorageRegionsStoragePoolsGetUnauthorized {
	return &InternalV1StorageRegionsStoragePoolsGetUnauthorized{}
}

/*
InternalV1StorageRegionsStoragePoolsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InternalV1StorageRegionsStoragePoolsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions storage pools get unauthorized response has a 2xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions storage pools get unauthorized response has a 3xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions storage pools get unauthorized response has a 4xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions storage pools get unauthorized response has a 5xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions storage pools get unauthorized response a status code equal to that given
func (o *InternalV1StorageRegionsStoragePoolsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the internal v1 storage regions storage pools get unauthorized response
func (o *InternalV1StorageRegionsStoragePoolsGetUnauthorized) Code() int {
	return 401
}

func (o *InternalV1StorageRegionsStoragePoolsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}][%d] internalV1StorageRegionsStoragePoolsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *InternalV1StorageRegionsStoragePoolsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}][%d] internalV1StorageRegionsStoragePoolsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *InternalV1StorageRegionsStoragePoolsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsStoragePoolsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsStoragePoolsGetForbidden creates a InternalV1StorageRegionsStoragePoolsGetForbidden with default headers values
func NewInternalV1StorageRegionsStoragePoolsGetForbidden() *InternalV1StorageRegionsStoragePoolsGetForbidden {
	return &InternalV1StorageRegionsStoragePoolsGetForbidden{}
}

/*
InternalV1StorageRegionsStoragePoolsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InternalV1StorageRegionsStoragePoolsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions storage pools get forbidden response has a 2xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions storage pools get forbidden response has a 3xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions storage pools get forbidden response has a 4xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions storage pools get forbidden response has a 5xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions storage pools get forbidden response a status code equal to that given
func (o *InternalV1StorageRegionsStoragePoolsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the internal v1 storage regions storage pools get forbidden response
func (o *InternalV1StorageRegionsStoragePoolsGetForbidden) Code() int {
	return 403
}

func (o *InternalV1StorageRegionsStoragePoolsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}][%d] internalV1StorageRegionsStoragePoolsGetForbidden  %+v", 403, o.Payload)
}

func (o *InternalV1StorageRegionsStoragePoolsGetForbidden) String() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}][%d] internalV1StorageRegionsStoragePoolsGetForbidden  %+v", 403, o.Payload)
}

func (o *InternalV1StorageRegionsStoragePoolsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsStoragePoolsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsStoragePoolsGetNotFound creates a InternalV1StorageRegionsStoragePoolsGetNotFound with default headers values
func NewInternalV1StorageRegionsStoragePoolsGetNotFound() *InternalV1StorageRegionsStoragePoolsGetNotFound {
	return &InternalV1StorageRegionsStoragePoolsGetNotFound{}
}

/*
InternalV1StorageRegionsStoragePoolsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InternalV1StorageRegionsStoragePoolsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions storage pools get not found response has a 2xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions storage pools get not found response has a 3xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions storage pools get not found response has a 4xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions storage pools get not found response has a 5xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions storage pools get not found response a status code equal to that given
func (o *InternalV1StorageRegionsStoragePoolsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the internal v1 storage regions storage pools get not found response
func (o *InternalV1StorageRegionsStoragePoolsGetNotFound) Code() int {
	return 404
}

func (o *InternalV1StorageRegionsStoragePoolsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}][%d] internalV1StorageRegionsStoragePoolsGetNotFound  %+v", 404, o.Payload)
}

func (o *InternalV1StorageRegionsStoragePoolsGetNotFound) String() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}][%d] internalV1StorageRegionsStoragePoolsGetNotFound  %+v", 404, o.Payload)
}

func (o *InternalV1StorageRegionsStoragePoolsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsStoragePoolsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsStoragePoolsGetInternalServerError creates a InternalV1StorageRegionsStoragePoolsGetInternalServerError with default headers values
func NewInternalV1StorageRegionsStoragePoolsGetInternalServerError() *InternalV1StorageRegionsStoragePoolsGetInternalServerError {
	return &InternalV1StorageRegionsStoragePoolsGetInternalServerError{}
}

/*
InternalV1StorageRegionsStoragePoolsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1StorageRegionsStoragePoolsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions storage pools get internal server error response has a 2xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions storage pools get internal server error response has a 3xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions storage pools get internal server error response has a 4xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 storage regions storage pools get internal server error response has a 5xx status code
func (o *InternalV1StorageRegionsStoragePoolsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 storage regions storage pools get internal server error response a status code equal to that given
func (o *InternalV1StorageRegionsStoragePoolsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 storage regions storage pools get internal server error response
func (o *InternalV1StorageRegionsStoragePoolsGetInternalServerError) Code() int {
	return 500
}

func (o *InternalV1StorageRegionsStoragePoolsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}][%d] internalV1StorageRegionsStoragePoolsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalV1StorageRegionsStoragePoolsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/storage-pools/{storage_pool_name}][%d] internalV1StorageRegionsStoragePoolsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalV1StorageRegionsStoragePoolsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsStoragePoolsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
