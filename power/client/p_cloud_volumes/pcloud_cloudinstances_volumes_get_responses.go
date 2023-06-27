// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.ibm.com/power-iaas/service-broker/power/models"
)

// PcloudCloudinstancesVolumesGetReader is a Reader for the PcloudCloudinstancesVolumesGet structure.
type PcloudCloudinstancesVolumesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesVolumesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesVolumesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesVolumesGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesVolumesGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesVolumesGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesVolumesGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesVolumesGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}] pcloud.cloudinstances.volumes.get", response, response.Code())
	}
}

// NewPcloudCloudinstancesVolumesGetOK creates a PcloudCloudinstancesVolumesGetOK with default headers values
func NewPcloudCloudinstancesVolumesGetOK() *PcloudCloudinstancesVolumesGetOK {
	return &PcloudCloudinstancesVolumesGetOK{}
}

/*
PcloudCloudinstancesVolumesGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesVolumesGetOK struct {
	Payload *models.Volume
}

// IsSuccess returns true when this pcloud cloudinstances volumes get o k response has a 2xx status code
func (o *PcloudCloudinstancesVolumesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances volumes get o k response has a 3xx status code
func (o *PcloudCloudinstancesVolumesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes get o k response has a 4xx status code
func (o *PcloudCloudinstancesVolumesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances volumes get o k response has a 5xx status code
func (o *PcloudCloudinstancesVolumesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes get o k response a status code equal to that given
func (o *PcloudCloudinstancesVolumesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudinstances volumes get o k response
func (o *PcloudCloudinstancesVolumesGetOK) Code() int {
	return 200
}

func (o *PcloudCloudinstancesVolumesGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetOK) GetPayload() *models.Volume {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesGetBadRequest creates a PcloudCloudinstancesVolumesGetBadRequest with default headers values
func NewPcloudCloudinstancesVolumesGetBadRequest() *PcloudCloudinstancesVolumesGetBadRequest {
	return &PcloudCloudinstancesVolumesGetBadRequest{}
}

/*
PcloudCloudinstancesVolumesGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesVolumesGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes get bad request response has a 2xx status code
func (o *PcloudCloudinstancesVolumesGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes get bad request response has a 3xx status code
func (o *PcloudCloudinstancesVolumesGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes get bad request response has a 4xx status code
func (o *PcloudCloudinstancesVolumesGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes get bad request response has a 5xx status code
func (o *PcloudCloudinstancesVolumesGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes get bad request response a status code equal to that given
func (o *PcloudCloudinstancesVolumesGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudinstances volumes get bad request response
func (o *PcloudCloudinstancesVolumesGetBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudinstancesVolumesGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesGetUnauthorized creates a PcloudCloudinstancesVolumesGetUnauthorized with default headers values
func NewPcloudCloudinstancesVolumesGetUnauthorized() *PcloudCloudinstancesVolumesGetUnauthorized {
	return &PcloudCloudinstancesVolumesGetUnauthorized{}
}

/*
PcloudCloudinstancesVolumesGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesVolumesGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes get unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesVolumesGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes get unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesVolumesGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes get unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesVolumesGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes get unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesVolumesGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes get unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesVolumesGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudinstances volumes get unauthorized response
func (o *PcloudCloudinstancesVolumesGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudinstancesVolumesGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesGetForbidden creates a PcloudCloudinstancesVolumesGetForbidden with default headers values
func NewPcloudCloudinstancesVolumesGetForbidden() *PcloudCloudinstancesVolumesGetForbidden {
	return &PcloudCloudinstancesVolumesGetForbidden{}
}

/*
PcloudCloudinstancesVolumesGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesVolumesGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes get forbidden response has a 2xx status code
func (o *PcloudCloudinstancesVolumesGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes get forbidden response has a 3xx status code
func (o *PcloudCloudinstancesVolumesGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes get forbidden response has a 4xx status code
func (o *PcloudCloudinstancesVolumesGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes get forbidden response has a 5xx status code
func (o *PcloudCloudinstancesVolumesGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes get forbidden response a status code equal to that given
func (o *PcloudCloudinstancesVolumesGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudinstances volumes get forbidden response
func (o *PcloudCloudinstancesVolumesGetForbidden) Code() int {
	return 403
}

func (o *PcloudCloudinstancesVolumesGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesGetNotFound creates a PcloudCloudinstancesVolumesGetNotFound with default headers values
func NewPcloudCloudinstancesVolumesGetNotFound() *PcloudCloudinstancesVolumesGetNotFound {
	return &PcloudCloudinstancesVolumesGetNotFound{}
}

/*
PcloudCloudinstancesVolumesGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesVolumesGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes get not found response has a 2xx status code
func (o *PcloudCloudinstancesVolumesGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes get not found response has a 3xx status code
func (o *PcloudCloudinstancesVolumesGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes get not found response has a 4xx status code
func (o *PcloudCloudinstancesVolumesGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances volumes get not found response has a 5xx status code
func (o *PcloudCloudinstancesVolumesGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances volumes get not found response a status code equal to that given
func (o *PcloudCloudinstancesVolumesGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudinstances volumes get not found response
func (o *PcloudCloudinstancesVolumesGetNotFound) Code() int {
	return 404
}

func (o *PcloudCloudinstancesVolumesGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesVolumesGetInternalServerError creates a PcloudCloudinstancesVolumesGetInternalServerError with default headers values
func NewPcloudCloudinstancesVolumesGetInternalServerError() *PcloudCloudinstancesVolumesGetInternalServerError {
	return &PcloudCloudinstancesVolumesGetInternalServerError{}
}

/*
PcloudCloudinstancesVolumesGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesVolumesGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances volumes get internal server error response has a 2xx status code
func (o *PcloudCloudinstancesVolumesGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances volumes get internal server error response has a 3xx status code
func (o *PcloudCloudinstancesVolumesGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances volumes get internal server error response has a 4xx status code
func (o *PcloudCloudinstancesVolumesGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances volumes get internal server error response has a 5xx status code
func (o *PcloudCloudinstancesVolumesGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances volumes get internal server error response a status code equal to that given
func (o *PcloudCloudinstancesVolumesGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudinstances volumes get internal server error response
func (o *PcloudCloudinstancesVolumesGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudinstancesVolumesGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/volumes/{volume_id}][%d] pcloudCloudinstancesVolumesGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesVolumesGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesVolumesGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
