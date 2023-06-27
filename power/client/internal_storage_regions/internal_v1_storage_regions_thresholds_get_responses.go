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

// InternalV1StorageRegionsThresholdsGetReader is a Reader for the InternalV1StorageRegionsThresholdsGet structure.
type InternalV1StorageRegionsThresholdsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1StorageRegionsThresholdsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInternalV1StorageRegionsThresholdsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewInternalV1StorageRegionsThresholdsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInternalV1StorageRegionsThresholdsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInternalV1StorageRegionsThresholdsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1StorageRegionsThresholdsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds] internal.v1.storage.regions.thresholds.get", response, response.Code())
	}
}

// NewInternalV1StorageRegionsThresholdsGetOK creates a InternalV1StorageRegionsThresholdsGetOK with default headers values
func NewInternalV1StorageRegionsThresholdsGetOK() *InternalV1StorageRegionsThresholdsGetOK {
	return &InternalV1StorageRegionsThresholdsGetOK{}
}

/*
InternalV1StorageRegionsThresholdsGetOK describes a response with status code 200, with default header values.

OK
*/
type InternalV1StorageRegionsThresholdsGetOK struct {
	Payload *models.Thresholds
}

// IsSuccess returns true when this internal v1 storage regions thresholds get o k response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 storage regions thresholds get o k response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds get o k response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 storage regions thresholds get o k response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions thresholds get o k response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the internal v1 storage regions thresholds get o k response
func (o *InternalV1StorageRegionsThresholdsGetOK) Code() int {
	return 200
}

func (o *InternalV1StorageRegionsThresholdsGetOK) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetOK  %+v", 200, o.Payload)
}

func (o *InternalV1StorageRegionsThresholdsGetOK) String() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetOK  %+v", 200, o.Payload)
}

func (o *InternalV1StorageRegionsThresholdsGetOK) GetPayload() *models.Thresholds {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Thresholds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsGetUnauthorized creates a InternalV1StorageRegionsThresholdsGetUnauthorized with default headers values
func NewInternalV1StorageRegionsThresholdsGetUnauthorized() *InternalV1StorageRegionsThresholdsGetUnauthorized {
	return &InternalV1StorageRegionsThresholdsGetUnauthorized{}
}

/*
InternalV1StorageRegionsThresholdsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InternalV1StorageRegionsThresholdsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions thresholds get unauthorized response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions thresholds get unauthorized response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds get unauthorized response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions thresholds get unauthorized response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions thresholds get unauthorized response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the internal v1 storage regions thresholds get unauthorized response
func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) Code() int {
	return 401
}

func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsGetForbidden creates a InternalV1StorageRegionsThresholdsGetForbidden with default headers values
func NewInternalV1StorageRegionsThresholdsGetForbidden() *InternalV1StorageRegionsThresholdsGetForbidden {
	return &InternalV1StorageRegionsThresholdsGetForbidden{}
}

/*
InternalV1StorageRegionsThresholdsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InternalV1StorageRegionsThresholdsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions thresholds get forbidden response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions thresholds get forbidden response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds get forbidden response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions thresholds get forbidden response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions thresholds get forbidden response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the internal v1 storage regions thresholds get forbidden response
func (o *InternalV1StorageRegionsThresholdsGetForbidden) Code() int {
	return 403
}

func (o *InternalV1StorageRegionsThresholdsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetForbidden  %+v", 403, o.Payload)
}

func (o *InternalV1StorageRegionsThresholdsGetForbidden) String() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetForbidden  %+v", 403, o.Payload)
}

func (o *InternalV1StorageRegionsThresholdsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsGetNotFound creates a InternalV1StorageRegionsThresholdsGetNotFound with default headers values
func NewInternalV1StorageRegionsThresholdsGetNotFound() *InternalV1StorageRegionsThresholdsGetNotFound {
	return &InternalV1StorageRegionsThresholdsGetNotFound{}
}

/*
InternalV1StorageRegionsThresholdsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InternalV1StorageRegionsThresholdsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions thresholds get not found response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions thresholds get not found response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds get not found response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 storage regions thresholds get not found response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 storage regions thresholds get not found response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the internal v1 storage regions thresholds get not found response
func (o *InternalV1StorageRegionsThresholdsGetNotFound) Code() int {
	return 404
}

func (o *InternalV1StorageRegionsThresholdsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetNotFound  %+v", 404, o.Payload)
}

func (o *InternalV1StorageRegionsThresholdsGetNotFound) String() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetNotFound  %+v", 404, o.Payload)
}

func (o *InternalV1StorageRegionsThresholdsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1StorageRegionsThresholdsGetInternalServerError creates a InternalV1StorageRegionsThresholdsGetInternalServerError with default headers values
func NewInternalV1StorageRegionsThresholdsGetInternalServerError() *InternalV1StorageRegionsThresholdsGetInternalServerError {
	return &InternalV1StorageRegionsThresholdsGetInternalServerError{}
}

/*
InternalV1StorageRegionsThresholdsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1StorageRegionsThresholdsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 storage regions thresholds get internal server error response has a 2xx status code
func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 storage regions thresholds get internal server error response has a 3xx status code
func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 storage regions thresholds get internal server error response has a 4xx status code
func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 storage regions thresholds get internal server error response has a 5xx status code
func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 storage regions thresholds get internal server error response a status code equal to that given
func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 storage regions thresholds get internal server error response
func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) Code() int {
	return 500
}

func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /internal/v1/storage/regions/{region_zone_id}/thresholds][%d] internalV1StorageRegionsThresholdsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1StorageRegionsThresholdsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
